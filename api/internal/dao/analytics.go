package dao

import (
	"context"
	"strings"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/google/uuid"

	"github.com/yueli-official/docs/api/internal/model"
)

func (p *PG) RecordAnalyticsSearch(ctx context.Context, query, collectionSlug, locale string, resultCount int) error {
	_, err := p.db.Model("doc_search_events").Ctx(ctx).Data(map[string]any{
		"id": uuid.NewString(), "query": query, "collection_slug": collectionSlug,
		"locale": locale, "result_count": resultCount,
	}).Insert()
	return err
}

func (p *PG) SeedAnalyticsSearch(ctx context.Context, id, query, collectionSlug, locale string, resultCount int, createdAt time.Time) error {
	_, err := p.db.Exec(ctx, `
INSERT INTO doc_search_events (
    id, query, collection_slug, locale, result_count, created_at
) VALUES (?, ?, ?, ?, ?, ?)
ON CONFLICT (id) DO NOTHING`, id, query, collectionSlug, locale, resultCount, createdAt)
	return err
}

func (p *PG) PublishedAnalyticsDocumentsByIDs(ctx context.Context, ids []string) ([]*model.AnalyticsDocument, error) {
	if len(ids) == 0 {
		return []*model.AnalyticsDocument{}, nil
	}
	placeholders := strings.TrimSuffix(strings.Repeat("?,", len(ids)), ",")
	arguments := make([]any, 0, len(ids))
	for _, id := range ids {
		arguments = append(arguments, id)
	}
	query := manageDocsCTE + ` SELECT
    d.id, d.title, c.slug AS collection_slug,
    v.key AS version_key, COALESCE(paths.slug_path, d.slug) AS slug_path, d.locale
` + manageDocsFrom + ` WHERE d.id IN (` + placeholders + `)
    AND d.status = 'published' AND v.status = 'published' AND d.deleted_at IS NULL`
	var documents []*model.AnalyticsDocument
	if err := p.db.Ctx(ctx).Raw(query, arguments...).Scan(&documents); err != nil {
		return nil, err
	}
	if documents == nil {
		documents = []*model.AnalyticsDocument{}
	}
	return documents, nil
}

func (p *PG) RecordAnalyticsTrafficSource(ctx context.Context, eventID, day, source string) error {
	return p.db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		result, err := tx.Ctx(ctx).Exec(`
INSERT INTO docs_traffic_source_receipts (event_id, day, source)
VALUES (?, ?::date, ?)
ON CONFLICT (event_id) DO NOTHING`, eventID, day, source)
		if err != nil {
			return err
		}
		inserted, err := result.RowsAffected()
		if err != nil || inserted == 0 {
			return err
		}
		_, err = tx.Ctx(ctx).Exec(`
INSERT INTO docs_traffic_source_daily (day, source, views)
VALUES (?::date, ?, 1)
ON CONFLICT (day, source) DO UPDATE
SET views = docs_traffic_source_daily.views + 1`, day, source)
		return err
	})
}

func (p *PG) AnalyticsTrafficSources(ctx context.Context, from, to time.Time, limit int) ([]model.AnalyticsTrafficSource, error) {
	var sources []model.AnalyticsTrafficSource
	err := p.db.Ctx(ctx).Raw(`
SELECT source, SUM(views)::bigint AS views
FROM docs_traffic_source_daily
WHERE day >= ?::date AND day < ?::date AND source <> 'internal'
GROUP BY source
ORDER BY views DESC, source ASC
LIMIT ?`, from.Format(time.DateOnly), to.Format(time.DateOnly), limit).Scan(&sources)
	if sources == nil {
		sources = []model.AnalyticsTrafficSource{}
	}
	return sources, err
}

func (p *PG) AnalyticsSearchSummary(ctx context.Context, from, to time.Time) (model.AnalyticsSearchSummary, error) {
	var summary model.AnalyticsSearchSummary
	err := p.db.Ctx(ctx).Raw(`
SELECT COUNT(*)::bigint AS searches,
       COUNT(*) FILTER (WHERE result_count = 0)::bigint AS zero_results
FROM doc_search_events
WHERE created_at >= ? AND created_at < ?`, from, to).Scan(&summary)
	return summary, err
}

func (p *PG) AnalyticsTopSearches(ctx context.Context, from, to time.Time, limit int) ([]model.AnalyticsSearchQuery, error) {
	var searches []model.AnalyticsSearchQuery
	err := p.db.Ctx(ctx).Raw(`
SELECT LOWER(BTRIM(query)) AS query,
       COUNT(*)::bigint AS searches,
       COUNT(*) FILTER (WHERE result_count = 0)::bigint AS zero_results
FROM doc_search_events
WHERE created_at >= ? AND created_at < ? AND BTRIM(query) <> ''
GROUP BY LOWER(BTRIM(query))
ORDER BY searches DESC, query ASC
LIMIT ?`, from, to, limit).Scan(&searches)
	if searches == nil {
		searches = []model.AnalyticsSearchQuery{}
	}
	return searches, err
}

func (p *PG) PruneAnalyticsTrafficSourceReceipts(ctx context.Context, before time.Time) error {
	_, err := p.db.Exec(ctx,
		"DELETE FROM docs_traffic_source_receipts WHERE day < ?::date",
		before.Format(time.DateOnly),
	)
	return err
}
