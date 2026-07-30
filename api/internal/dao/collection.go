package dao

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/yueli-official/docs/api/internal/model"
)

const tCollections = "collections"

func collectionSelectFields(alias string) []string {
	return []string{
		alias + ".id", alias + ".slug", alias + ".title", alias + ".description",
		alias + ".cover_asset_id", alias + ".cover_url", alias + ".icon", alias + ".sort_order", alias + ".author_sub",
		"(SELECT COUNT(*) FROM docs d WHERE d.collection_id = " + alias + ".id AND d.deleted_at IS NULL) AS doc_count",
	}
}

func (p *PG) InsertCollection(ctx context.Context, m *model.Collection) error {
	_, err := p.db.Model(tCollections).Ctx(ctx).Data(g.Map{
		"id": m.ID, "slug": m.Slug, "title": m.Title, "description": m.Description,
		"cover_asset_id": m.CoverAssetID, "cover_url": m.CoverURL, "icon": m.Icon, "sort_order": m.SortOrder, "author_sub": m.AuthorSub,
	}).Insert()
	return err
}

func (p *PG) InsertCollectionWithDefaultVersion(
	ctx context.Context,
	collection *model.Collection,
	version *model.CollectionVersion,
	hook TransactionHook,
) error {
	return p.db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if _, err := tx.Model(tCollections).Ctx(ctx).Data(g.Map{
			"id": collection.ID, "slug": collection.Slug, "title": collection.Title,
			"description": collection.Description, "cover_asset_id": collection.CoverAssetID,
			"cover_url": collection.CoverURL, "icon": collection.Icon,
			"sort_order": collection.SortOrder, "author_sub": collection.AuthorSub,
		}).Insert(); err != nil {
			return err
		}
		if _, err := tx.Model(tCollectionVersions).Ctx(ctx).Data(g.Map{
			"id": version.ID, "collection_id": version.CollectionID, "key": version.Key,
			"label": version.Label, "status": nz(version.Status, "draft"),
			"is_default": version.IsDefault, "sort_order": version.SortOrder,
			"source_version_id": nilIfEmpty(version.SourceVersionID),
		}).Insert(); err != nil {
			return err
		}
		return runTransactionHook(ctx, tx, hook)
	})
}

func (p *PG) GetCollectionByID(ctx context.Context, id string) (*model.Collection, error) {
	var c *model.Collection
	err := p.db.Model(tCollections+" c").Ctx(ctx).
		Fields(collectionSelectFields("c")).
		Where("c.id", id).
		Limit(1).
		Scan(&c)
	return c, err
}

func (p *PG) GetCollectionBySlug(ctx context.Context, slug string) (*model.Collection, error) {
	var c *model.Collection
	err := p.db.Model(tCollections+" c").Ctx(ctx).
		Fields(collectionSelectFields("c")).
		Where("c.slug", slug).
		Limit(1).
		Scan(&c)
	return c, err
}

func (p *PG) ListCollections(ctx context.Context) ([]*model.Collection, error) {
	var out []*model.Collection
	err := p.db.Model(tCollections + " c").Ctx(ctx).
		Fields(collectionSelectFields("c")).
		OrderAsc("c.sort_order").
		OrderDesc("c.created_at").
		Scan(&out)
	if out == nil {
		out = []*model.Collection{}
	}
	return out, err
}

func (p *PG) UpdateCollection(ctx context.Context, m *model.Collection) error {
	return p.UpdateCollectionWithHook(ctx, m, nil)
}

func (p *PG) UpdateCollectionWithHook(ctx context.Context, m *model.Collection, hook TransactionHook) error {
	return p.db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if _, err := tx.Model(tCollections).Ctx(ctx).Where("id", m.ID).Data(g.Map{
			"title": m.Title, "slug": m.Slug, "description": m.Description,
			"cover_asset_id": m.CoverAssetID, "cover_url": m.CoverURL,
			"icon": m.Icon, "sort_order": m.SortOrder, "updated_at": gtime.Now(),
		}).Update(); err != nil {
			return err
		}
		return runTransactionHook(ctx, tx, hook)
	})
}

func (p *PG) DeleteCollection(ctx context.Context, id string) error {
	return p.DeleteCollectionWithHook(ctx, id, nil)
}

func (p *PG) DeleteCollectionWithHook(ctx context.Context, id string, hook TransactionHook) error {
	return p.DeleteCollectionWithHooks(ctx, id, nil, hook)
}

// DeleteCollectionWithHooks runs beforeDelete while the collection's
// cascade-owned rows still exist, then runs afterDelete after the collection
// has been removed. Both hooks share the deletion transaction.
func (p *PG) DeleteCollectionWithHooks(
	ctx context.Context,
	id string,
	beforeDelete TransactionHook,
	afterDelete TransactionHook,
) error {
	return p.db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if err := runTransactionHook(ctx, tx, beforeDelete); err != nil {
			return err
		}
		if _, err := tx.Model(tCollections).Ctx(ctx).Where("id", id).Delete(); err != nil {
			return err
		}
		return runTransactionHook(ctx, tx, afterDelete)
	})
}
