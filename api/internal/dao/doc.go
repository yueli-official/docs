package dao

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/yueli-official/docs/api/internal/model"
)

const tDocs = "docs"

// InsertDoc inserts a new doc. parent_id is left SQL NULL when m.ParentID is empty.
func (p *PG) InsertDoc(ctx context.Context, m *model.Doc) error {
	return p.InsertDocWithHook(ctx, m, nil)
}

func (p *PG) InsertDocWithHook(ctx context.Context, m *model.Doc, hook TransactionHook) error {
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
		"badge_text":      m.BadgeText,
		"badge_icon":      m.BadgeIcon,
		"sort_order":      m.SortOrder,
		"author_sub":      m.AuthorSub,
	}
	if m.ParentID != "" {
		data["parent_id"] = m.ParentID // else leaves SQL NULL
	}
	return p.db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if _, err := tx.Model(tDocs).Ctx(ctx).Data(data).Insert(); err != nil {
			return err
		}
		return runTransactionHook(ctx, tx, hook)
	})
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

func (p *PG) PublishedDocsByIDs(ctx context.Context, ids []string) ([]*model.Doc, error) {
	if len(ids) == 0 {
		return []*model.Doc{}, nil
	}
	var out []*model.Doc
	err := p.db.Model(tDocs+" d").Ctx(ctx).
		InnerJoin("collection_versions v", "v.id=d.version_id").
		WhereIn("d.id", ids).Where("d.status", "published").WhereIn("v.status", []string{"published", "archived"}).
		Where("d.deleted_at IS NULL").Fields("d.*").Scan(&out)
	return out, err
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

func (p *PG) GetPublishedDocByLogicalKey(ctx context.Context, collectionID, versionID, locale, translationKey string) (*model.Doc, error) {
	var out *model.Doc
	err := p.db.Model(tDocs).Ctx(ctx).
		Where("collection_id", collectionID).
		Where("version_id", versionID).
		Where("locale", locale).
		Where("translation_key", translationKey).
		Where("status", "published").
		WhereNull("deleted_at").
		Limit(1).
		Scan(&out)
	return out, err
}

func (p *PG) GetDocPath(ctx context.Context, id string) (string, error) {
	value, err := p.db.GetValue(ctx, `
WITH RECURSIVE ancestors AS (
    SELECT id, parent_id, slug, slug::text AS slug_path
    FROM docs
    WHERE id = ?::uuid AND deleted_at IS NULL
    UNION ALL
    SELECT parent.id, parent.parent_id, parent.slug,
           (parent.slug || '/' || child.slug_path)::text
    FROM docs parent
    JOIN ancestors child ON child.parent_id = parent.id
    WHERE parent.deleted_at IS NULL
)
SELECT slug_path FROM ancestors WHERE parent_id IS NULL LIMIT 1`, id)
	if err != nil {
		return "", err
	}
	return value.String(), nil
}

// UpdateDoc overwrites the mutable fields of an existing doc.
func (p *PG) UpdateDoc(ctx context.Context, m *model.Doc) error {
	return p.UpdateDocWithHook(ctx, m, nil)
}

func (p *PG) UpdateDocWithHook(ctx context.Context, m *model.Doc, hook TransactionHook) error {
	return p.db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if _, err := tx.Model(tDocs).Ctx(ctx).Where("id", m.ID).Data(g.Map{
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
			"badge_text":      m.BadgeText,
			"badge_icon":      m.BadgeIcon,
			"sort_order":      m.SortOrder,
			"parent_id":       nilIfEmpty(m.ParentID),
			"updated_at":      gtime.Now(),
		}).Update(); err != nil {
			return err
		}
		if _, err := tx.Ctx(ctx).Exec(`
WITH RECURSIVE descendants AS (
    SELECT id FROM docs WHERE parent_id = ?::uuid AND deleted_at IS NULL
    UNION ALL
    SELECT child.id FROM docs child
    JOIN descendants parent ON child.parent_id = parent.id
    WHERE child.deleted_at IS NULL
)
UPDATE docs
SET version_id = ?::uuid, locale = ?, updated_at = NOW()
WHERE id IN (SELECT id FROM descendants)`, m.ID, m.VersionID, m.Locale); err != nil {
			return err
		}
		return runTransactionHook(ctx, tx, hook)
	})
}

// SoftDeleteDoc marks the doc deleted without removing the row.
func (p *PG) SoftDeleteDoc(ctx context.Context, id string) error {
	return p.SoftDeleteDocWithHook(ctx, id, nil)
}

func (p *PG) SoftDeleteDocWithHook(ctx context.Context, id string, hook TransactionHook) error {
	return p.db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if _, err := tx.Ctx(ctx).Exec(`
WITH RECURSIVE subtree AS (
    SELECT id FROM docs WHERE id = ?::uuid AND deleted_at IS NULL
    UNION ALL
    SELECT child.id FROM docs child
    JOIN subtree parent ON child.parent_id = parent.id
    WHERE child.deleted_at IS NULL
)
UPDATE docs SET deleted_at = NOW(), updated_at = NOW()
WHERE id IN (SELECT id FROM subtree)`, id); err != nil {
			return err
		}
		return runTransactionHook(ctx, tx, hook)
	})
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
