package dao

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"platform/products/docs/api/internal/model"
)

const tDocs = "docs"
const tDocSearchEvents = "doc_search_events"

// InsertDoc inserts a new doc. parent_id is left SQL NULL when m.ParentID is empty.
func (p *PG) InsertDoc(ctx context.Context, m *model.Doc) error {
	data := g.Map{
		"id":              m.ID,
		"collection_id":   m.CollectionID,
		"version_id":      m.VersionID,
		"slug":            m.Slug,
		"title":           m.Title,
		"content":         m.Content,
		"excerpt":         m.Excerpt,
		"seo_title":       m.SEOTitle,
		"seo_description": m.SEODescription,
		"status":          nz(m.Status, "draft"),
		"locale":          nz(m.Locale, "en"),
		"translation_key": nz(m.TranslationKey, m.ID),
		"sort_order":      m.SortOrder,
		"author_sub":      m.AuthorSub,
	}
	if m.ParentID != "" {
		data["parent_id"] = m.ParentID // else leaves SQL NULL
	}
	_, err := p.db.Model(tDocs).Ctx(ctx).Data(data).Insert()
	return err
}

// ListDocsByCollection returns non-deleted docs for the given collection + version + locale.
func (p *PG) ListDocsByCollection(ctx context.Context, collectionID, versionID, locale string) ([]*model.Doc, error) {
	var out []*model.Doc
	err := p.db.Model(tDocs).Ctx(ctx).
		Where("collection_id", collectionID).
		Where("version_id", versionID).
		Where("locale", locale).
		WhereNull("deleted_at").
		OrderAsc("sort_order").
		OrderAsc("created_at").
		Scan(&out)
	if out == nil {
		out = []*model.Doc{}
	}
	return out, err
}

// ListDocsByCollectionStatus returns non-deleted docs for one publication status.
func (p *PG) ListDocsByCollectionStatus(ctx context.Context, collectionID, versionID, locale, status string) ([]*model.Doc, error) {
	var out []*model.Doc
	err := p.db.Model(tDocs).Ctx(ctx).
		Where("collection_id", collectionID).
		Where("version_id", versionID).
		Where("locale", locale).
		Where("status", status).
		WhereNull("deleted_at").
		OrderAsc("sort_order").
		OrderAsc("created_at").
		Scan(&out)
	if out == nil {
		out = []*model.Doc{}
	}
	return out, err
}

// SearchPublishedDocs returns lightweight published doc matches for public search.
func (p *PG) SearchPublishedDocs(ctx context.Context, collectionID, versionID, locale, query string, limit int) (*model.SearchResult, error) {
	q := strings.TrimSpace(query)
	if q == "" {
		return &model.SearchResult{Items: []*model.Doc{}, CollectionFacets: []*model.SearchCollectionFacet{}}, nil
	}
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	result, err := p.searchPublishedDocsFTS(ctx, collectionID, versionID, locale, q, limit)
	if err == nil {
		return result, nil
	}
	if !isSearchConfigError(err) {
		return nil, err
	}
	return p.searchPublishedDocsLike(ctx, collectionID, versionID, locale, q, limit)
}

func (p *PG) searchPublishedDocsFTS(ctx context.Context, collectionID, versionID, locale, query string, limit int) (*model.SearchResult, error) {
	conds := []string{"v.status = ?", "d.locale = ?", "d.status = ?", "d.deleted_at IS NULL", "d.search_vector @@ tsq"}
	args := []any{query, "published", locale, "published"}
	if versionID != "" {
		conds = append(conds, "d.version_id = ?")
		args = append(args, versionID)
	}
	if collectionID != "" {
		conds = append(conds, "d.collection_id = ?")
		args = append(args, collectionID)
	}
	from := "docs d JOIN collection_versions v ON v.id = d.version_id, websearch_to_tsquery('chinese_zh', ?) tsq"
	where := strings.Join(conds, " AND ")

	total, err := p.db.GetValue(ctx, "SELECT COUNT(*) FROM "+from+" WHERE "+where, args...)
	if err != nil {
		return nil, err
	}
	var out []*model.Doc
	rowsSQL := "SELECT d.* FROM " + from + " WHERE " + where +
		" ORDER BY ts_rank(d.search_vector, tsq) DESC, d.updated_at DESC LIMIT ?"
	err = p.db.Raw(rowsSQL, append(args, limit)...).Scan(&out)
	if out == nil {
		out = []*model.Doc{}
	}
	if err != nil {
		return nil, err
	}

	var facets []*model.SearchCollectionFacet
	facetSQL := "SELECT c.id, c.slug, c.title, COUNT(*) AS count " +
		"FROM docs d JOIN collections c ON c.id = d.collection_id JOIN collection_versions v ON v.id = d.version_id, websearch_to_tsquery('chinese_zh', ?) tsq " +
		"WHERE " + where + " GROUP BY c.id, c.slug, c.title ORDER BY count DESC, c.title ASC"
	if err := p.db.Raw(facetSQL, args...).Scan(&facets); err != nil {
		return nil, err
	}
	if facets == nil {
		facets = []*model.SearchCollectionFacet{}
	}

	return &model.SearchResult{
		Items:            out,
		Total:            total.Int(),
		CollectionFacets: facets,
	}, nil
}

func (p *PG) searchPublishedDocsLike(ctx context.Context, collectionID, versionID, locale, query string, limit int) (*model.SearchResult, error) {
	pattern := likePattern(query)
	conds := []string{
		"v.status = ?",
		"d.locale = ?",
		"d.status = ?",
		"d.deleted_at IS NULL",
		"(d.title ILIKE ? ESCAPE '\\' OR d.excerpt ILIKE ? ESCAPE '\\' OR d.content ILIKE ? ESCAPE '\\')",
	}
	args := []any{"published", locale, "published", pattern, pattern, pattern}
	if versionID != "" {
		conds = append(conds, "d.version_id = ?")
		args = append(args, versionID)
	}
	if collectionID != "" {
		conds = append(conds, "d.collection_id = ?")
		args = append(args, collectionID)
	}
	where := strings.Join(conds, " AND ")

	total, err := p.db.GetValue(ctx, "SELECT COUNT(*) FROM docs d JOIN collection_versions v ON v.id = d.version_id WHERE "+where, args...)
	if err != nil {
		return nil, err
	}

	var out []*model.Doc
	rowsSQL := "SELECT d.* FROM docs d JOIN collection_versions v ON v.id = d.version_id WHERE " + where +
		" ORDER BY CASE WHEN d.title ILIKE ? ESCAPE '\\' THEN 0 WHEN d.excerpt ILIKE ? ESCAPE '\\' THEN 1 ELSE 2 END, d.updated_at DESC LIMIT ?"
	rowArgs := append(append([]any{}, args...), pattern, pattern, limit)
	err = p.db.Raw(rowsSQL, rowArgs...).Scan(&out)
	if out == nil {
		out = []*model.Doc{}
	}
	if err != nil {
		return nil, err
	}

	var facets []*model.SearchCollectionFacet
	facetSQL := "SELECT c.id, c.slug, c.title, COUNT(*) AS count " +
		"FROM docs d JOIN collections c ON c.id = d.collection_id JOIN collection_versions v ON v.id = d.version_id " +
		"WHERE " + where + " GROUP BY c.id, c.slug, c.title ORDER BY count DESC, c.title ASC"
	if err := p.db.Raw(facetSQL, args...).Scan(&facets); err != nil {
		return nil, err
	}
	if facets == nil {
		facets = []*model.SearchCollectionFacet{}
	}

	return &model.SearchResult{
		Items:            out,
		Total:            total.Int(),
		CollectionFacets: facets,
	}, nil
}

func isSearchConfigError(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "text search configuration") ||
		strings.Contains(msg, "search_vector")
}

func likePattern(q string) string {
	q = strings.ReplaceAll(q, `\`, `\\`)
	q = strings.ReplaceAll(q, `%`, `\%`)
	q = strings.ReplaceAll(q, `_`, `\_`)
	return "%" + q + "%"
}

func (p *PG) InsertSearchEvent(ctx context.Context, query, collectionSlug, locale string, resultCount int) error {
	_, err := p.db.Model(tDocSearchEvents).Ctx(ctx).Data(g.Map{
		"query":           query,
		"collection_slug": collectionSlug,
		"locale":          locale,
		"result_count":    resultCount,
	}).Insert()
	return err
}

// GetDocByID returns the doc with the given id (excluding soft-deleted rows).
func (p *PG) GetDocByID(ctx context.Context, id string) (*model.Doc, error) {
	var d *model.Doc
	err := p.db.Model(tDocs).Ctx(ctx).
		Where("id", id).
		WhereNull("deleted_at").
		Limit(1).
		Scan(&d)
	return d, err
}

func (p *PG) GetPublishedDocByCollectionPath(ctx context.Context, collectionID, versionID, locale, path string) (*model.Doc, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		return nil, nil
	}
	parentID := ""
	var current *model.Doc
	for _, slug := range parts {
		var d *model.Doc
		q := p.db.Model(tDocs).Ctx(ctx).
			Where("collection_id", collectionID).
			Where("version_id", versionID).
			Where("locale", locale).
			Where("slug", slug).
			Where("status", "published").
			WhereNull("deleted_at")
		if parentID == "" {
			q = q.WhereNull("parent_id")
		} else {
			q = q.Where("parent_id", parentID)
		}
		if err := q.Limit(1).Scan(&d); err != nil {
			return nil, err
		}
		if d == nil {
			return nil, nil
		}
		current = d
		parentID = d.ID
	}
	return current, nil
}

// UpdateDoc overwrites the mutable fields of an existing doc.
func (p *PG) UpdateDoc(ctx context.Context, m *model.Doc) error {
	_, err := p.db.Model(tDocs).Ctx(ctx).Where("id", m.ID).Data(g.Map{
		"title":           m.Title,
		"slug":            m.Slug,
		"content":         m.Content,
		"excerpt":         m.Excerpt,
		"seo_title":       m.SEOTitle,
		"seo_description": m.SEODescription,
		"version_id":      m.VersionID,
		"status":          m.Status,
		"locale":          m.Locale,
		"translation_key": m.TranslationKey,
		"sort_order":      m.SortOrder,
		"parent_id":       nilIfEmpty(m.ParentID),
		"updated_at":      gtime.Now(),
	}).Update()
	return err
}

// SoftDeleteDoc marks the doc deleted without removing the row.
func (p *PG) SoftDeleteDoc(ctx context.Context, id string) error {
	_, err := p.db.Model(tDocs).Ctx(ctx).
		Where("id", id).
		Data(g.Map{"deleted_at": gtime.Now()}).
		Update()
	return err
}

// nz returns def when s is empty; used for defaulting optional string fields.
func nz(s, def string) string {
	if s == "" {
		return def
	}
	return s
}

// nilIfEmpty maps "" → nil so nullable FK columns receive SQL NULL.
func nilIfEmpty(s string) any {
	if s == "" {
		return nil
	}
	return s
}
