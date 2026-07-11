package catalog

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"platform/products/docs/api/internal/model"
)

func (s *Service) GetHomeConfig(ctx context.Context) (*model.HomeConfig, error) {
	return s.dao.GetHomeConfig(ctx)
}

func (s *Service) UpdateHomeConfig(ctx context.Context, cfg *model.HomeConfig) (*model.HomeConfig, error) {
	clean := &model.HomeConfig{
		QuickLinks:          sanitizeHomeQuickLinks(cfg.QuickLinks),
		FeaturedCollections: sanitizeFeaturedCollections(cfg.FeaturedCollections),
	}
	if err := s.dao.UpsertHomeConfig(ctx, clean); err != nil {
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
