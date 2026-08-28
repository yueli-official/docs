package dao

import (
	"context"
	"fmt"
	"strings"

	"github.com/yueli-official/docs/api/internal/model"
)

const manageDocsCTE = `WITH RECURSIVE doc_paths AS (
    SELECT id, parent_id, slug, slug::text AS slug_path
    FROM docs
    WHERE parent_id IS NULL
    UNION ALL
    SELECT child.id, child.parent_id, child.slug, (parent.slug_path || '/' || child.slug)::text
    FROM docs child
    JOIN doc_paths parent ON parent.id = child.parent_id
)
`

const manageDocsFrom = `FROM docs d
JOIN collections c ON c.id = d.collection_id
JOIN collection_versions v ON v.id = d.version_id
LEFT JOIN docs parent_doc ON parent_doc.id = d.parent_id
LEFT JOIN doc_paths paths ON paths.id = d.id
`

func manageDocsOrder(query model.ManageDocsQuery) (string, error) {
	column := map[string]string{
		"updatedAt": "d.updated_at",
		"title":     "LOWER(d.title)",
		"path":      "LOWER(COALESCE(paths.slug_path, d.slug))",
		"sortOrder": "d.sort_order",
	}[query.Sort]
	if column == "" {
		return "", fmt.Errorf("unsupported manage docs sort %q", query.Sort)
	}
	direction := strings.ToUpper(query.Direction)
	if direction != "ASC" && direction != "DESC" {
		return "", fmt.Errorf("unsupported manage docs direction %q", query.Direction)
	}
	return column + " " + direction + ", d.id ASC", nil
}

func manageDocsConditions(query model.ManageDocsQuery, includeLifecycle bool) (string, []any) {
	conditions := []string{"d.deleted_at IS NULL"}
	args := []any{}
	if query.CollectionID != "" {
		conditions = append(conditions, "d.collection_id = ?")
		args = append(args, query.CollectionID)
	}
	if query.OwnerSub != "" {
		conditions = append(conditions, "d.author_sub = ?")
		args = append(args, query.OwnerSub)
	}
	if query.Version != "" {
		conditions = append(conditions, "v.key = ?")
		args = append(args, query.Version)
	}
	if query.Locale != "" {
		conditions = append(conditions, "d.locale = ?")
		args = append(args, query.Locale)
	}
	if query.ParentID == "root" {
		conditions = append(conditions, "d.parent_id IS NULL")
	} else if query.ParentID != "" {
		conditions = append(conditions, "d.parent_id = ?")
		args = append(args, query.ParentID)
	}
	if query.Q != "" {
		pattern := likePattern(query.Q)
		conditions = append(conditions, `(d.title ILIKE ? ESCAPE '\' OR d.slug ILIKE ? ESCAPE '\' OR d.excerpt ILIKE ? ESCAPE '\' OR c.title ILIKE ? ESCAPE '\' OR COALESCE(paths.slug_path, d.slug) ILIKE ? ESCAPE '\')`)
		args = append(args, pattern, pattern, pattern, pattern, pattern)
	}
	if includeLifecycle {
		if query.Status != "all" {
			conditions = append(conditions, "d.status = ?")
			args = append(args, query.Status)
		}
		if query.Quality == "issues" {
			conditions = append(conditions, "(BTRIM(COALESCE(d.title, '')) = '' OR BTRIM(COALESCE(d.slug, '')) = '' OR BTRIM(COALESCE(d.excerpt, '')) = '')")
		}
	}
	return strings.Join(conditions, " AND "), args
}

func likePattern(q string) string {
	q = strings.ReplaceAll(q, `\`, `\\`)
	q = strings.ReplaceAll(q, `%`, `\%`)
	q = strings.ReplaceAll(q, `_`, `\_`)
	return "%" + q + "%"
}

// ManageDocs applies the allowlisted admin workbench query in PostgreSQL. It
// returns lifecycle counts under the non-lifecycle filters so tabs stay stable
// while the selected status or quality view changes.
func (p *PG) ManageDocs(ctx context.Context, query model.ManageDocsQuery) (*model.ManageDocsResult, error) {
	order, err := manageDocsOrder(query)
	if err != nil {
		return nil, err
	}
	where, args := manageDocsConditions(query, true)
	total, err := p.db.GetValue(ctx, manageDocsCTE+" SELECT COUNT(*) "+manageDocsFrom+" WHERE "+where, args...)
	if err != nil {
		return nil, err
	}

	rowsSQL := manageDocsCTE + ` SELECT
    d.id, d.collection_id, c.slug AS collection_slug, c.title AS collection_title,
    d.version_id, v.key AS version_key, v.label AS version_label,
    COALESCE(d.parent_id::text, '') AS parent_id, COALESCE(parent_doc.title, '') AS parent_title,
    d.slug, COALESCE(paths.slug_path, d.slug) AS slug_path,
    d.title, d.excerpt, d.status, d.locale, d.badge_text, d.badge_icon, d.sort_order, d.updated_at
` + manageDocsFrom + " WHERE " + where + " ORDER BY " + order + " LIMIT ? OFFSET ?"
	rowArgs := append(append([]any{}, args...), query.Size, (query.Page-1)*query.Size)
	var items []*model.ManageDoc
	if err := p.db.Ctx(ctx).Raw(rowsSQL, rowArgs...).Scan(&items); err != nil {
		return nil, err
	}
	if items == nil {
		items = []*model.ManageDoc{}
	}

	countWhere, countArgs := manageDocsConditions(query, false)
	countsSQL := manageDocsCTE + ` SELECT
    COUNT(*) AS all_count,
    COUNT(*) FILTER (WHERE d.status = 'draft') AS draft_count,
    COUNT(*) FILTER (WHERE d.status = 'published') AS published_count,
    COUNT(*) FILTER (WHERE d.status = 'archived') AS archived_count,
    COUNT(*) FILTER (WHERE BTRIM(COALESCE(d.title, '')) = '' OR BTRIM(COALESCE(d.slug, '')) = '' OR BTRIM(COALESCE(d.excerpt, '')) = '') AS issues_count
` + manageDocsFrom + " WHERE " + countWhere
	var counts model.ManageDocCounts
	if err := p.db.Ctx(ctx).Raw(countsSQL, countArgs...).Scan(&counts); err != nil {
		return nil, err
	}

	return &model.ManageDocsResult{Items: items, Total: total.Int(), Counts: counts}, nil
}
