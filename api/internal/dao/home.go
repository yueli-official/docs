package dao

import (
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/yueli-official/docs/api/internal/model"
)

const tHomeConfig = "home_config"

type homeConfigRow struct {
	QuickLinks          string `orm:"quick_links"`
	FeaturedCollections string `orm:"featured_collections"`
	HomeEyebrow         string `orm:"home_eyebrow"`
	HomeTitle           string `orm:"home_title"`
	HomeSubtitle        string `orm:"home_subtitle"`
	SiteTitle           string `orm:"site_title"`
	SiteDescription     string `orm:"site_description"`
	SupportEmail        string `orm:"support_email"`
	FooterTagline       string `orm:"footer_tagline"`
	FooterCopyright     string `orm:"footer_copyright"`
}

func (p *PG) GetHomeConfig(ctx context.Context) (*model.HomeConfig, error) {
	var row *homeConfigRow
	err := p.db.Model(tHomeConfig).Ctx(ctx).
		Fields("quick_links::text AS quick_links", "featured_collections::text AS featured_collections", "home_eyebrow", "home_title", "home_subtitle", "site_title", "site_description", "support_email", "footer_tagline", "footer_copyright").
		Where("key", "default").
		Limit(1).
		Scan(&row)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, gerror.New("docs site configuration is not seeded")
	}
	out := &model.HomeConfig{QuickLinks: []*model.HomeQuickLink{}, FeaturedCollections: []string{}, HomeEyebrow: row.HomeEyebrow, HomeTitle: row.HomeTitle, HomeSubtitle: row.HomeSubtitle, SiteTitle: row.SiteTitle, SiteDescription: row.SiteDescription, SupportEmail: row.SupportEmail, FooterTagline: row.FooterTagline, FooterCopyright: row.FooterCopyright}
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
	return p.UpsertHomeConfigWithHook(ctx, cfg, nil)
}

func (p *PG) UpsertHomeConfigWithHook(
	ctx context.Context,
	cfg *model.HomeConfig,
	hook TransactionHook,
) error {
	quickLinks, err := json.Marshal(cfg.QuickLinks)
	if err != nil {
		return err
	}
	featured, err := json.Marshal(cfg.FeaturedCollections)
	if err != nil {
		return err
	}
	return p.db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = tx.Ctx(ctx).Exec(`INSERT INTO home_config (key, quick_links, featured_collections, home_eyebrow, home_title, home_subtitle, site_title, site_description, support_email, footer_tagline, footer_copyright, updated_at)
		VALUES ('default', ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, now())
		ON CONFLICT (key) DO UPDATE SET
			quick_links = EXCLUDED.quick_links,
			featured_collections = EXCLUDED.featured_collections,
			home_eyebrow = EXCLUDED.home_eyebrow,
			home_title = EXCLUDED.home_title,
			home_subtitle = EXCLUDED.home_subtitle,
			site_title = EXCLUDED.site_title,
			site_description = EXCLUDED.site_description,
			support_email = EXCLUDED.support_email,
			footer_tagline = EXCLUDED.footer_tagline,
			footer_copyright = EXCLUDED.footer_copyright,
			updated_at = now()`,
			string(quickLinks), string(featured), cfg.HomeEyebrow, cfg.HomeTitle, cfg.HomeSubtitle, cfg.SiteTitle, cfg.SiteDescription, cfg.SupportEmail, cfg.FooterTagline, cfg.FooterCopyright)
		if err != nil {
			return err
		}
		return runTransactionHook(ctx, tx, hook)
	})
}
