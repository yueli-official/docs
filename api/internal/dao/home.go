package dao

import (
	"context"
	"encoding/json"

	"platform/products/docs/api/internal/model"
)

const tHomeConfig = "home_config"

type homeConfigRow struct {
	QuickLinks          string `orm:"quick_links"`
	FeaturedCollections string `orm:"featured_collections"`
}

func defaultHomeConfig() *model.HomeConfig {
	return &model.HomeConfig{
		QuickLinks:          []*model.HomeQuickLink{},
		FeaturedCollections: []string{},
	}
}

func (p *PG) GetHomeConfig(ctx context.Context) (*model.HomeConfig, error) {
	var row *homeConfigRow
	err := p.db.Model(tHomeConfig).Ctx(ctx).
		Fields("quick_links::text AS quick_links", "featured_collections::text AS featured_collections").
		Where("key", "default").
		Limit(1).
		Scan(&row)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return defaultHomeConfig(), nil
	}
	out := defaultHomeConfig()
	if row.QuickLinks != "" {
		if err := json.Unmarshal([]byte(row.QuickLinks), &out.QuickLinks); err != nil {
			return nil, err
		}
	}
	if row.FeaturedCollections != "" {
		if err := json.Unmarshal([]byte(row.FeaturedCollections), &out.FeaturedCollections); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (p *PG) UpsertHomeConfig(ctx context.Context, cfg *model.HomeConfig) error {
	quickLinks, err := json.Marshal(cfg.QuickLinks)
	if err != nil {
		return err
	}
	featured, err := json.Marshal(cfg.FeaturedCollections)
	if err != nil {
		return err
	}
	_, err = p.db.Exec(ctx, `INSERT INTO home_config (key, quick_links, featured_collections, updated_at)
		VALUES ('default', ?, ?, now())
		ON CONFLICT (key) DO UPDATE SET
			quick_links = EXCLUDED.quick_links,
			featured_collections = EXCLUDED.featured_collections,
			updated_at = now()`,
		string(quickLinks), string(featured))
	return err
}
