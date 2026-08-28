package dao

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/yueli-official/docs/api/internal/model"
)

const tCollectionLocales = "collection_locales"

func (p *PG) ListCollectionLocales(ctx context.Context, collectionID string, enabledOnly bool) ([]*model.CollectionLocale, error) {
	query := p.db.Model(tCollectionLocales+" l").Ctx(ctx).
		Fields("l.*", "(SELECT COUNT(*) FROM docs d WHERE d.collection_id = l.collection_id AND d.locale = l.locale AND d.deleted_at IS NULL) AS doc_count").
		Where("l.collection_id", collectionID)
	if enabledOnly {
		query = query.Where("l.enabled", true)
	}
	var out []*model.CollectionLocale
	err := query.OrderDesc("l.is_default").OrderAsc("l.sort_order").OrderAsc("l.locale").Scan(&out)
	if out == nil {
		out = []*model.CollectionLocale{}
	}
	return out, err
}

func (p *PG) GetCollectionLocale(ctx context.Context, collectionID, locale string) (*model.CollectionLocale, error) {
	var out *model.CollectionLocale
	err := p.db.Model(tCollectionLocales).Ctx(ctx).
		Where("collection_id", collectionID).
		Where("locale", locale).
		Limit(1).
		Scan(&out)
	return out, err
}

func (p *PG) GetDefaultCollectionLocale(ctx context.Context, collectionID string) (*model.CollectionLocale, error) {
	var out *model.CollectionLocale
	err := p.db.Model(tCollectionLocales).Ctx(ctx).
		Where("collection_id", collectionID).
		Where("is_default", true).
		Where("enabled", true).
		Limit(1).
		Scan(&out)
	return out, err
}

func (p *PG) UpsertCollectionLocale(ctx context.Context, locale *model.CollectionLocale, hook TransactionHook) error {
	return p.db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if locale.IsDefault {
			if _, err := tx.Model(tCollectionLocales).Ctx(ctx).
				Where("collection_id", locale.CollectionID).
				Data(g.Map{"is_default": false, "updated_at": gtime.Now()}).Update(); err != nil {
				return err
			}
		}
		if _, err := tx.Model(tCollectionLocales).Ctx(ctx).Data(g.Map{
			"collection_id": locale.CollectionID,
			"locale":        locale.Locale,
			"label":         locale.Label,
			"html_lang":     locale.HTMLLang,
			"direction":     locale.Direction,
			"is_default":    locale.IsDefault,
			"enabled":       locale.Enabled,
			"sort_order":    locale.SortOrder,
			"updated_at":    gtime.Now(),
		}).OnConflict("collection_id,locale").Save(); err != nil {
			return err
		}
		return runTransactionHook(ctx, tx, hook)
	})
}

func (p *PG) DeleteCollectionLocale(ctx context.Context, collectionID, locale string) error {
	_, err := p.db.Model(tCollectionLocales).Ctx(ctx).
		Where("collection_id", collectionID).
		Where("locale", locale).
		Where("is_default", false).
		Delete()
	return err
}

func (p *PG) CountCollectionLocaleDocs(ctx context.Context, collectionID, locale string) (int, error) {
	return p.db.Model("docs").Ctx(ctx).
		Where("collection_id", collectionID).
		Where("locale", locale).
		WhereNull("deleted_at").
		Count()
}

func (p *PG) CloneCollectionLocale(
	ctx context.Context,
	collectionID, versionID, sourceLocale string,
	targetLocale *model.CollectionLocale,
	authorSub string,
	hook TransactionHook,
) (int, error) {
	count := 0
	err := p.db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if existing, err := tx.Model(tCollectionLocales).Ctx(ctx).
			Where("collection_id", collectionID).
			Where("locale", targetLocale.Locale).
			Count(); err != nil {
			return err
		} else if existing == 0 {
			if _, err := tx.Model(tCollectionLocales).Ctx(ctx).Data(g.Map{
				"collection_id": collectionID,
				"locale":        targetLocale.Locale, "label": targetLocale.Label,
				"html_lang": targetLocale.HTMLLang, "direction": targetLocale.Direction,
				"is_default": false, "enabled": true, "sort_order": targetLocale.SortOrder,
			}).Insert(); err != nil {
				return err
			}
		}

		var sourceDocs []*model.Doc
		if err := tx.Model(tDocs).Ctx(ctx).
			Where("collection_id", collectionID).
			Where("version_id", versionID).
			Where("locale", sourceLocale).
			WhereNull("deleted_at").
			OrderAsc("created_at").
			Scan(&sourceDocs); err != nil {
			return err
		}
		cloned, err := cloneDocumentSet(ctx, tx, sourceDocs, documentCloneTarget{
			CollectionID: collectionID, VersionID: versionID,
			Locale: targetLocale.Locale, AuthorSub: authorSub, Status: "draft",
		})
		if err != nil {
			return err
		}
		count = len(cloned)
		return runTransactionHook(ctx, tx, hook)
	})
	return count, err
}
