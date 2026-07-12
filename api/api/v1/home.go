package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"platform/products/docs/api/internal/model"
)

type HomeQuickLinkView struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Icon           string `json:"icon"`
	To             string `json:"to"`
	CollectionSlug string `json:"collectionSlug"`
	SortOrder      int    `json:"sortOrder"`
	Enabled        bool   `json:"enabled"`
}

type HomeConfigView struct {
	QuickLinks          []*HomeQuickLinkView `json:"quickLinks"`
	FeaturedCollections []string             `json:"featuredCollections"`
	HomeEyebrow         string               `json:"homeEyebrow"`
	HomeTitle           string               `json:"homeTitle"`
	HomeSubtitle        string               `json:"homeSubtitle"`
	SiteTitle           string               `json:"siteTitle"`
	SiteDescription     string               `json:"siteDescription"`
	SupportEmail        string               `json:"supportEmail"`
	FooterTagline       string               `json:"footerTagline"`
	FooterCopyright     string               `json:"footerCopyright"`
}

type GetHomeConfigReq struct {
	g.Meta `path:"/api/v1/home" method:"get" tags:"docs" summary:"Get homepage configuration"`
}
type GetHomeConfigRes struct {
	Config *HomeConfigView `json:"config"`
}

type UpdateHomeConfigReq struct {
	g.Meta              `path:"/api/v1/home" method:"patch" tags:"docs" summary:"Update homepage configuration"`
	QuickLinks          []*model.HomeQuickLink `json:"quickLinks"`
	FeaturedCollections []string               `json:"featuredCollections"`
	HomeEyebrow         string                 `json:"homeEyebrow"`
	HomeTitle           string                 `json:"homeTitle"`
	HomeSubtitle        string                 `json:"homeSubtitle"`
	SiteTitle           string                 `json:"siteTitle"`
	SiteDescription     string                 `json:"siteDescription"`
	SupportEmail        string                 `json:"supportEmail"`
	FooterTagline       string                 `json:"footerTagline"`
	FooterCopyright     string                 `json:"footerCopyright"`
}
type UpdateHomeConfigRes struct {
	Config *HomeConfigView `json:"config"`
}
