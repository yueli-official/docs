package dao

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

const discoveryPagesSQL = `
WITH RECURSIVE doc_paths AS (
    SELECT doc.id, doc.parent_id, doc.slug, doc.slug::text AS slug_path
    FROM docs doc
    WHERE doc.parent_id IS NULL
    UNION ALL
    SELECT child.id, child.parent_id, child.slug,
           (parent.slug_path || '/' || child.slug)::text
    FROM docs child
    JOIN doc_paths parent ON parent.id = child.parent_id
),
pages AS (
    SELECT 'collection:' || collection.id::text AS key,
           '/' || collection.slug AS path,
           'collection'::text AS kind,
           collection.title,
           collection.description,
           collection.cover_url AS image_url,
           COALESCE((
               SELECT locale.locale FROM collection_locales locale
               WHERE locale.collection_id = collection.id
                 AND locale.is_default = true
                 AND locale.enabled = true
               LIMIT 1
           ), ?::text) AS locale,
           collection.updated_at
    FROM collections collection
    WHERE EXISTS (
        SELECT 1 FROM docs doc
        JOIN collection_versions version ON version.id = doc.version_id
        WHERE doc.collection_id = collection.id
          AND doc.status = 'published'
          AND doc.deleted_at IS NULL
          AND version.status IN ('published', 'archived')
    )
    UNION ALL
    SELECT 'doc:' || doc.id::text,
           '/' || collection.slug || '/' || path.slug_path ||
           CASE
               WHEN doc.locale <> COALESCE(default_locale.locale, ?) AND NOT version.is_default
                   THEN CHR(63) || 'locale=' || doc.locale || CHR(38) || 'version=' || version.key
               WHEN doc.locale <> COALESCE(default_locale.locale, ?)
                   THEN CHR(63) || 'locale=' || doc.locale
               WHEN NOT version.is_default
                   THEN CHR(63) || 'version=' || version.key
               ELSE ''
           END,
           'doc',
           COALESCE(NULLIF(doc.seo_title, ''), doc.title),
           COALESCE(NULLIF(doc.seo_description, ''), doc.excerpt),
           '',
           doc.locale,
           doc.updated_at
    FROM docs doc
    JOIN doc_paths path ON path.id = doc.id
    JOIN collections collection ON collection.id = doc.collection_id
    JOIN collection_versions version ON version.id = doc.version_id
    LEFT JOIN collection_locales default_locale
      ON default_locale.collection_id = collection.id
     AND default_locale.is_default = true
     AND default_locale.enabled = true
    WHERE doc.status = 'published'
      AND doc.deleted_at IS NULL
      AND version.status IN ('published', 'archived')
)
SELECT * FROM pages
WHERE ? || path > ?
ORDER BY ? || path ASC
LIMIT ?`

type DiscoveryRow struct {
	Key         string      `orm:"key"`
	Path        string      `orm:"path"`
	Kind        string      `orm:"kind"`
	Title       string      `orm:"title"`
	Description string      `orm:"description"`
	ImageURL    string      `orm:"image_url"`
	Locale      string      `orm:"locale"`
	UpdatedAt   *gtime.Time `orm:"updated_at"`
}

func (p *PG) ListDiscoveryPages(
	ctx context.Context,
	origin string,
	defaultLocale string,
	afterURL string,
	limit int,
) ([]DiscoveryRow, error) {
	var rows []DiscoveryRow
	err := p.db.Ctx(ctx).Raw(discoveryPagesSQL,
		defaultLocale, defaultLocale, defaultLocale,
		origin, afterURL, origin, limit,
	).Scan(&rows)
	return rows, err
}
