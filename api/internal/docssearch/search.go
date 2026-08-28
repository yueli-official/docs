package docssearch

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/yueli-official/foundation/go/search"
)

const analyzer search.AnalyzerKey = "content-v1"

type Index struct {
	module   search.Module
	postgres *search.Postgres
}

func Definition() search.Definition {
	return search.Definition{
		Consumer: "docs.public", Version: 1,
		Analyzers: []search.AnalyzerDefinition{{Key: analyzer, QueryMode: search.QueryWeb, Required: []search.Capability{search.CapabilityFullText}}},
		Filters: []search.FilterDefinition{
			{Name: "collection", Facetable: true}, {Name: "version"}, {Name: "locale"},
		},
		Limits: search.Limits{MaxPageSize: 100, MaxFacetBuckets: 100},
	}
}

func NewPostgres(ctx context.Context, db *sql.DB, site string) (*Index, error) {
	catalog, err := search.Compile(Definition())
	if err != nil {
		return nil, err
	}
	module, err := search.NewPostgres(ctx, catalog, search.PostgresOptions{
		DB: db, InstanceKey: "docs.search." + site,
		AnalyzerBindings: map[search.AnalyzerKey]string{analyzer: "chinese_zh"},
	})
	if err != nil {
		return nil, err
	}
	return &Index{module: module, postgres: module}, nil
}

func NewMemory() *Index { return &Index{module: search.NewMemory(search.MustCompile(Definition()))} }

type docRow struct {
	ID, CollectionID, VersionID, Locale, Title, Excerpt, Content string
	Status, VersionStatus                                        string
	Revision                                                     uint64
	UpdatedAt                                                    time.Time
	DeletedAt                                                    sql.NullTime
}

const docSelect = `SELECT d.id,d.collection_id,d.version_id,d.locale,d.title,d.excerpt,d.content,
	d.status,v.status,d.search_revision,d.updated_at,d.deleted_at
	FROM docs d JOIN collection_versions v ON v.id=d.version_id`

func scanDoc(scanner interface{ Scan(...any) error }) (docRow, error) {
	var row docRow
	err := scanner.Scan(&row.ID, &row.CollectionID, &row.VersionID, &row.Locale, &row.Title,
		&row.Excerpt, &row.Content, &row.Status, &row.VersionStatus, &row.Revision, &row.UpdatedAt, &row.DeletedAt)
	return row, err
}

func (index *Index) SubtreeHook(rootID string) func(context.Context, *sql.Tx) error {
	return func(ctx context.Context, tx *sql.Tx) error {
		rows, err := tx.QueryContext(ctx, `
			WITH RECURSIVE subtree AS (
				SELECT id FROM docs WHERE id=$1
				UNION ALL SELECT child.id FROM docs child JOIN subtree parent ON child.parent_id=parent.id
			)
			`+docSelect+` WHERE d.id IN (SELECT id FROM subtree)
		`, rootID)
		if err != nil {
			return err
		}
		var values []docRow
		for rows.Next() {
			row, scanErr := scanDoc(rows)
			if scanErr != nil {
				_ = rows.Close()
				return scanErr
			}
			values = append(values, row)
		}
		if err := rows.Close(); err != nil {
			return err
		}
		return index.applyRowsTx(ctx, tx, values)
	}
}

func (index *Index) CollectionHook(collectionID string) func(context.Context, *sql.Tx) error {
	return func(ctx context.Context, tx *sql.Tx) error {
		rows, err := tx.QueryContext(ctx, docSelect+" WHERE d.collection_id=$1", collectionID)
		if err != nil {
			return err
		}
		var values []docRow
		for rows.Next() {
			row, scanErr := scanDoc(rows)
			if scanErr != nil {
				_ = rows.Close()
				return scanErr
			}
			values = append(values, row)
		}
		if err := rows.Close(); err != nil {
			return err
		}
		return index.applyRowsTx(ctx, tx, values)
	}
}

// DeleteCollectionHook emits revisioned tombstones before the product
// transaction cascades the collection's source documents.
func (index *Index) DeleteCollectionHook(collectionID string) func(context.Context, *sql.Tx) error {
	return func(ctx context.Context, tx *sql.Tx) error {
		rows, err := tx.QueryContext(ctx, `
			SELECT id,search_revision FROM docs WHERE collection_id=$1 ORDER BY id
		`, collectionID)
		if err != nil {
			return err
		}
		type ref struct {
			id       string
			revision uint64
		}
		var refs []ref
		for rows.Next() {
			var value ref
			if err := rows.Scan(&value.id, &value.revision); err != nil {
				_ = rows.Close()
				return err
			}
			refs = append(refs, value)
		}
		if err := rows.Close(); err != nil {
			return err
		}
		if len(refs) == 0 {
			return nil
		}
		projector, err := index.postgres.Bind(tx)
		if err != nil {
			return err
		}
		const batchSize = 100
		for start := 0; start < len(refs); start += batchSize {
			end := min(start+batchSize, len(refs))
			changes := make([]search.Change, 0, end-start)
			for _, value := range refs[start:end] {
				changes = append(changes, search.Remove(
					search.DocumentKey{Kind: "document", ID: search.DocumentID(value.id)},
					search.ProjectionRevision(value.revision+1),
				))
			}
			_, err = projector.Apply(ctx, search.Batch{
				ID: search.BatchID(fmt.Sprintf(
					"docs.collection.%s.delete.%d.%d", collectionID, start, end,
				)),
				Changes: changes,
			})
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (index *Index) applyRowsTx(ctx context.Context, tx *sql.Tx, values []docRow) error {
	projector, err := index.postgres.Bind(tx)
	if err != nil {
		return err
	}
	for _, row := range values {
		if err := index.applyRow(ctx, projector, row); err != nil {
			return err
		}
	}
	return nil
}

func (index *Index) applyRow(ctx context.Context, projector search.Projector, row docRow) error {
	key := search.DocumentKey{Kind: "document", ID: search.DocumentID(row.ID)}
	var change search.Change
	if row.Status == "published" && (row.VersionStatus == "published" || row.VersionStatus == "archived") && !row.DeletedAt.Valid {
		change = search.Upsert(search.SourceDocument{
			Key: key, Revision: search.ProjectionRevision(row.Revision), Analyzer: analyzer,
			Title: row.Title, Summary: row.Excerpt, Body: row.Content, SortAt: row.UpdatedAt.UTC(),
			Filters: search.FieldValues{
				"collection": search.Keyword(row.CollectionID), "version": search.Keyword(row.VersionID),
				"locale": search.Keyword(row.Locale),
			},
			Visibility: search.VisibilityReference{ResourceType: "docs.document", ResourceID: row.ID},
		})
	} else {
		change = search.Remove(key, search.ProjectionRevision(row.Revision))
	}
	_, err := projector.Apply(ctx, search.Batch{
		ID:      search.BatchID(fmt.Sprintf("docs.document.%s.%d", row.ID, row.Revision)),
		Changes: []search.Change{change},
	})
	return err
}

func (index *Index) Reconcile(ctx context.Context, db *sql.DB) error {
	rows, err := db.QueryContext(ctx, docSelect)
	if err != nil {
		return err
	}
	var values []docRow
	for rows.Next() {
		row, scanErr := scanDoc(rows)
		if scanErr != nil {
			_ = rows.Close()
			return scanErr
		}
		values = append(values, row)
	}
	if err := rows.Close(); err != nil {
		return err
	}
	for _, row := range values {
		if err := index.applyRow(ctx, index.module, row); err != nil {
			return err
		}
	}
	return nil
}

func (index *Index) Search(ctx context.Context, text, collectionID, versionID, locale string, limit int) (search.Page, error) {
	filters := []search.Filter{search.Equal("locale", locale)}
	if collectionID != "" {
		filters = append(filters, search.Equal("collection", collectionID))
	}
	if versionID != "" {
		filters = append(filters, search.Equal("version", versionID))
	}
	facets := []search.FacetRequest{}
	if collectionID == "" {
		facets = append(facets, search.FacetRequest{Name: "collection", Limit: 100})
	}
	return index.module.Search(ctx, search.Query{
		Text: text, Analyzer: analyzer, Filters: filters, Facets: facets,
		Page:      search.PageRequest{Size: limit},
		Highlight: search.HighlightRequest{Fields: []search.TextField{search.TextTitle, search.TextSummary}},
	})
}
