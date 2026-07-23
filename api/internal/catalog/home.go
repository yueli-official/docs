package catalog

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/google/uuid"
	"github.com/yueli-official/foundation/go/audit"

	"platform/products/docs/api/internal/dao"
	"platform/products/docs/api/internal/docsaudit"
	"platform/products/docs/api/internal/model"
)

func (s *Service) GetHomeConfig(ctx context.Context) (*model.HomeConfig, error) {
	return s.dao.GetHomeConfig(ctx)
}

func (s *Service) UpdateHomeConfig(ctx context.Context, cfg *model.HomeConfig) (*model.HomeConfig, error) {
	if cfg == nil {
		return nil, gerror.New("docs site configuration is required")
	}
	clean := &model.HomeConfig{
		QuickLinks:          sanitizeHomeQuickLinks(cfg.QuickLinks),
		FeaturedCollections: sanitizeFeaturedCollections(cfg.FeaturedCollections),
		HomeEyebrow:         strings.TrimSpace(cfg.HomeEyebrow),
		HomeTitle:           strings.TrimSpace(cfg.HomeTitle),
		HomeSubtitle:        strings.TrimSpace(cfg.HomeSubtitle),
		SiteTitle:           strings.TrimSpace(cfg.SiteTitle),
		SiteDescription:     strings.TrimSpace(cfg.SiteDescription),
		SupportEmail:        strings.TrimSpace(cfg.SupportEmail),
		FooterTagline:       strings.TrimSpace(cfg.FooterTagline),
		FooterCopyright:     strings.TrimSpace(cfg.FooterCopyright),
	}
	if clean.HomeEyebrow == "" || clean.HomeTitle == "" || clean.HomeSubtitle == "" || clean.SiteTitle == "" || clean.SiteDescription == "" || clean.FooterTagline == "" || clean.FooterCopyright == "" {
		return nil, gerror.New("docs homepage, site, and footer content must be configured")
	}
	var auditHook dao.TransactionHook
	if s.audit != nil {
		raw, _ := json.Marshal(clean)
		sum := sha256.Sum256(raw)
		auditHook = s.audit.Hook(
			ctx, docsaudit.ActionSiteProfilePublished, uuid.NewString(),
			audit.Target{Type: "docs.site_profile", ID: "default"},
			docsaudit.Evidence{Digest: hex.EncodeToString(sum[:])}, "",
		)
	}
	if err := s.dao.UpsertHomeConfigWithHook(ctx, clean, auditHook); err != nil {
		return nil, err
	}
	return s.dao.GetHomeConfig(ctx)
}

func sanitizeHomeQuickLinks(in []*model.HomeQuickLink) []*model.HomeQuickLink {
	out := make([]*model.HomeQuickLink, 0, len(in))
	for i, link := range in {
		if link == nil {
			continue
		}
		title := strings.TrimSpace(link.Title)
		to := strings.TrimSpace(link.To)
		collectionSlug := strings.TrimSpace(link.CollectionSlug)
		if title == "" && to == "" && collectionSlug == "" {
			continue
		}
		id := strings.TrimSpace(link.ID)
		if id == "" {
			id = uuid.NewString()
		}
		out = append(out, &model.HomeQuickLink{
			ID:             id,
			Title:          title,
			Description:    strings.TrimSpace(link.Description),
			Icon:           strings.TrimSpace(link.Icon),
			To:             to,
			CollectionSlug: collectionSlug,
			SortOrder:      i,
			Enabled:        link.Enabled,
		})
	}
	return out
}

func sanitizeFeaturedCollections(in []string) []string {
	seen := map[string]bool{}
	out := make([]string, 0, len(in))
	for _, slug := range in {
		slug = strings.TrimSpace(slug)
		if slug == "" || seen[slug] {
			continue
		}
		seen[slug] = true
		out = append(out, slug)
	}
	return out
}
