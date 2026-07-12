package controller

import (
	v1 "platform/products/docs/api/api/v1"
	"platform/products/docs/api/internal/model"
)

func homeConfigView(m *model.HomeConfig) *v1.HomeConfigView {
	if m == nil {
		return &v1.HomeConfigView{
			QuickLinks:          []*v1.HomeQuickLinkView{},
			FeaturedCollections: []string{},
			HomeEyebrow:         "Product manual",
			HomeSubtitle:        "搜索产品手册、集成说明和操作指南。先找到任务，再进入对应文档集继续阅读。",
		}
	}
	links := make([]*v1.HomeQuickLinkView, len(m.QuickLinks))
	for i, link := range m.QuickLinks {
		links[i] = &v1.HomeQuickLinkView{
			ID:             link.ID,
			Title:          link.Title,
			Description:    link.Description,
			Icon:           link.Icon,
			To:             link.To,
			CollectionSlug: link.CollectionSlug,
			SortOrder:      link.SortOrder,
			Enabled:        link.Enabled,
		}
	}
	return &v1.HomeConfigView{
		QuickLinks:          links,
		FeaturedCollections: append([]string(nil), m.FeaturedCollections...),
		HomeEyebrow:         m.HomeEyebrow,
		HomeTitle:           m.HomeTitle,
		HomeSubtitle:        m.HomeSubtitle,
		SiteTitle:           m.SiteTitle,
		SiteDescription:     m.SiteDescription,
		SupportEmail:        m.SupportEmail,
		FooterTagline:       m.FooterTagline,
		FooterCopyright:     m.FooterCopyright,
	}
}
