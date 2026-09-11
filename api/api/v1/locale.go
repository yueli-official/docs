package v1

import "github.com/gogf/gf/v2/frame/g"

type CollectionLocaleView struct {
	Title        string `json:"title"`
	CollectionID string `json:"collectionId"`
	Locale       string `json:"locale"`
	Label        string `json:"label"`
	HTMLLang     string `json:"htmlLang"`
	Direction    string `json:"direction"`
	IsDefault    bool   `json:"isDefault"`
	Enabled      bool   `json:"enabled"`
	SortOrder    int    `json:"sortOrder"`
	DocCount     int    `json:"docCount"`
}

type ListManageCollectionLocalesReq struct {
	g.Meta       `path:"/api/v1/manage/collections/{id}/locales" method:"get" tags:"docs" summary:"List collection locales (admin)"`
	CollectionID string `json:"id" in:"path" v:"required"`
}
type ListManageCollectionLocalesRes struct {
	Items []*CollectionLocaleView `json:"items"`
}

type UpsertCollectionLocaleReq struct {
	Title        *string `json:"title"`
	g.Meta       `path:"/api/v1/manage/collections/{id}/locales" method:"post" tags:"docs" summary:"Create or update a collection locale"`
	CollectionID string `json:"id" in:"path" v:"required"`
	Locale       string `json:"locale" v:"required"`
	Label        string `json:"label" v:"required"`
	HTMLLang     string `json:"htmlLang"`
	Direction    string `json:"direction"`
	IsDefault    bool   `json:"isDefault"`
	Enabled      bool   `json:"enabled"`
	SortOrder    int    `json:"sortOrder"`
}
type UpsertCollectionLocaleRes struct {
	Locale *CollectionLocaleView `json:"locale"`
}

type DeleteCollectionLocaleReq struct {
	g.Meta       `path:"/api/v1/manage/collections/{id}/locales/{locale}" method:"delete" tags:"docs" summary:"Delete an unused non-default collection locale"`
	CollectionID string `json:"id" in:"path" v:"required"`
	Locale       string `json:"locale" in:"path" v:"required"`
}
type DeleteCollectionLocaleRes struct {
	g.Meta `status:"204"`
}

type CloneCollectionLocaleReq struct {
	g.Meta          `path:"/api/v1/manage/collections/{id}/locales/clone" method:"post" tags:"docs" summary:"Clone one collection locale into translation drafts"`
	CollectionID    string `json:"id" in:"path" v:"required"`
	SourceLocale    string `json:"sourceLocale" v:"required"`
	TargetLocale    string `json:"targetLocale" v:"required"`
	TargetLabel     string `json:"targetLabel" v:"required"`
	TargetHTMLLang  string `json:"targetHtmlLang"`
	TargetDirection string `json:"targetDirection"`
	TargetSortOrder int    `json:"targetSortOrder"`
}
type CloneCollectionLocaleRes struct {
	Created int `json:"created"`
}
