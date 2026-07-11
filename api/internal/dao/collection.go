package dao

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"platform/products/docs/api/internal/model"
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
	_, err := p.db.Model(tCollections).Ctx(ctx).Where("id", m.ID).Data(g.Map{
		"title": m.Title, "slug": m.Slug, "description": m.Description,
		"cover_asset_id": m.CoverAssetID, "cover_url": m.CoverURL,
		"icon": m.Icon, "sort_order": m.SortOrder, "updated_at": gtime.Now(),
	}).Update()
	return err
}

func (p *PG) DeleteCollection(ctx context.Context, id string) error {
	_, err := p.db.Model(tCollections).Ctx(ctx).Where("id", id).Delete()
	return err
}
