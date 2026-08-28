package dao

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/yueli-official/docs/api/internal/model"
)

const tCollectionVersions = "collection_versions"

func (p *PG) InsertCollectionVersion(ctx context.Context, m *model.CollectionVersion) error {
	_, err := p.db.Model(tCollectionVersions).Ctx(ctx).Data(g.Map{
		"id":                m.ID,
		"collection_id":     m.CollectionID,
		"key":               m.Key,
		"label":             m.Label,
		"status":            nz(m.Status, "draft"),
		"is_default":        m.IsDefault,
		"sort_order":        m.SortOrder,
		"source_version_id": nilIfEmpty(m.SourceVersionID),
	}).Insert()
	return err
}

func (p *PG) UpdateCollectionVersion(ctx context.Context, version *model.CollectionVersion, hook TransactionHook) error {
	return p.db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if version.IsDefault {
			if _, err := tx.Model(tCollectionVersions).Ctx(ctx).
				Where("collection_id", version.CollectionID).
				Data(g.Map{"is_default": false, "updated_at": gtime.Now()}).Update(); err != nil {
				return err
			}
		}
		if _, err := tx.Model(tCollectionVersions).Ctx(ctx).
			Where("id", version.ID).
			Where("collection_id", version.CollectionID).
			Data(g.Map{
				"label":      version.Label,
				"status":     version.Status,
				"is_default": version.IsDefault,
				"sort_order": version.SortOrder,
				"updated_at": gtime.Now(),
			}).Update(); err != nil {
			return err
		}
		return runTransactionHook(ctx, tx, hook)
	})
}

func (p *PG) ListCollectionVersions(ctx context.Context, collectionID string) ([]*model.CollectionVersion, error) {
	var out []*model.CollectionVersion
	err := p.db.Model(tCollectionVersions).Ctx(ctx).
		Where("collection_id", collectionID).
		OrderAsc("sort_order").
		OrderAsc("created_at").
		Scan(&out)
	if out == nil {
		out = []*model.CollectionVersion{}
	}
	return out, err
}

func (p *PG) GetCollectionVersionByID(ctx context.Context, id string) (*model.CollectionVersion, error) {
	var out *model.CollectionVersion
	err := p.db.Model(tCollectionVersions).Ctx(ctx).Where("id", id).Limit(1).Scan(&out)
	return out, err
}

func (p *PG) GetCollectionVersionByKey(ctx context.Context, collectionID, key string) (*model.CollectionVersion, error) {
	var out *model.CollectionVersion
	err := p.db.Model(tCollectionVersions).Ctx(ctx).
		Where("collection_id", collectionID).
		Where("key", key).
		Limit(1).
		Scan(&out)
	return out, err
}

func (p *PG) GetDefaultCollectionVersion(ctx context.Context, collectionID string) (*model.CollectionVersion, error) {
	var out *model.CollectionVersion
	err := p.db.Model(tCollectionVersions).Ctx(ctx).
		Where("collection_id", collectionID).
		Where("is_default", true).
		Limit(1).
		Scan(&out)
	return out, err
}
