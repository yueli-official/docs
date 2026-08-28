package v1

import "github.com/gogf/gf/v2/frame/g"

type GetCollectionVariantsReq struct {
	g.Meta `path:"/api/v1/collections/{slug}/variants" method:"get" tags:"docs" summary:"List public languages and versions for a collection"`
	Slug   string `json:"slug" in:"path" v:"required"`
}
type GetCollectionVariantsRes struct {
	Locales  []*CollectionLocaleView  `json:"locales"`
	Versions []*CollectionVersionView `json:"versions"`
}

type ResolveDocumentVariantReq struct {
	g.Meta         `path:"/api/v1/collections/{slug}/variant" method:"get" tags:"docs" summary:"Resolve the matching document in another language or version"`
	Slug           string `json:"slug" in:"path" v:"required"`
	TranslationKey string `json:"translationKey" in:"query"`
	Locale         string `json:"locale" in:"query"`
	Version        string `json:"version" in:"query"`
}
type ResolveDocumentVariantRes struct {
	Path           string `json:"path"`
	Locale         string `json:"locale"`
	Version        string `json:"version"`
	TranslationKey string `json:"translationKey,omitempty"`
	Fallback       string `json:"fallback,omitempty"`
}
