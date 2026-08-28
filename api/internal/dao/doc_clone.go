package dao

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yueli-official/docs/api/internal/model"
	"github.com/yueli-official/foundation/go/identifier"
)

type documentCloneTarget struct {
	CollectionID string
	VersionID    string
	Locale       string
	AuthorSub    string
	Status       string
}

func cloneDocumentSet(
	ctx context.Context,
	tx gdb.TX,
	sourceDocs []*model.Doc,
	target documentCloneTarget,
) ([]*model.Doc, error) {
	cloned := make([]*model.Doc, 0, len(sourceDocs))
	idMap := make(map[string]string, len(sourceDocs))
	remaining := append([]*model.Doc(nil), sourceDocs...)
	for len(remaining) > 0 {
		progress := false
		next := make([]*model.Doc, 0, len(remaining))
		for _, source := range remaining {
			parentID := ""
			if source.ParentID != "" {
				mapped, ok := idMap[source.ParentID]
				if !ok {
					next = append(next, source)
					continue
				}
				parentID = mapped
			}
			locale := target.Locale
			if locale == "" {
				locale = source.Locale
			}
			copy := &model.Doc{
				ID: identifier.MustNew().String(), CollectionID: target.CollectionID,
				VersionID: target.VersionID, ParentID: parentID,
				Slug: source.Slug, Title: source.Title, Content: source.Content,
				Excerpt: source.Excerpt, SEOTitle: source.SEOTitle,
				SEODescription: source.SEODescription, Status: nz(target.Status, "draft"),
				Locale: locale, TranslationKey: source.TranslationKey,
				BadgeText: source.BadgeText, BadgeIcon: source.BadgeIcon,
				SortOrder: source.SortOrder, AuthorSub: target.AuthorSub,
			}
			if _, err := tx.Model(tDocs).Ctx(ctx).Data(g.Map{
				"id": copy.ID, "collection_id": copy.CollectionID,
				"version_id": copy.VersionID, "parent_id": nilIfEmpty(copy.ParentID),
				"slug": copy.Slug, "title": copy.Title, "content": copy.Content,
				"excerpt": copy.Excerpt, "seo_title": copy.SEOTitle,
				"seo_description": copy.SEODescription, "status": copy.Status,
				"locale": copy.Locale, "translation_key": copy.TranslationKey,
				"badge_text": copy.BadgeText, "badge_icon": copy.BadgeIcon,
				"sort_order": copy.SortOrder, "author_sub": copy.AuthorSub,
			}).Insert(); err != nil {
				return nil, err
			}
			idMap[source.ID] = copy.ID
			cloned = append(cloned, copy)
			progress = true
		}
		if !progress {
			return nil, fmt.Errorf("cannot clone document set: parent graph is invalid")
		}
		remaining = next
	}
	return cloned, nil
}
