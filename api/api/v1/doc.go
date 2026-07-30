package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/yueli-official/foundation/go/discovery"
)

// DocView is the public JSON projection of a doc.
type DocView struct {
	ID             string     `json:"id"`
	CollectionID   string     `json:"collectionId"`
	VersionID      string     `json:"versionId"`
	ParentID       string     `json:"parentId,omitempty"`
	Slug           string     `json:"slug"`
	Title          string     `json:"title"`
	Content        string     `json:"content"`
	Excerpt        string     `json:"excerpt"`
	SEOTitle       string     `json:"seoTitle"`
	SEODescription string     `json:"seoDescription"`
	Status         string     `json:"status"`
	Locale         string     `json:"locale"`
	TranslationKey string     `json:"translationKey"`
	SortOrder      int        `json:"sortOrder"`
	Children       []*DocView `json:"children,omitempty"`
}

// DocTreeNodeView is the lightweight JSON projection used for collection navigation.
type DocTreeNodeView struct {
	ID             string             `json:"id"`
	CollectionID   string             `json:"collectionId"`
	VersionID      string             `json:"versionId"`
	ParentID       string             `json:"parentId,omitempty"`
	Slug           string             `json:"slug"`
	Title          string             `json:"title"`
	Excerpt        string             `json:"excerpt"`
	Status         string             `json:"status"`
	Locale         string             `json:"locale"`
	TranslationKey string             `json:"translationKey"`
	SortOrder      int                `json:"sortOrder"`
	Children       []*DocTreeNodeView `json:"children,omitempty"`
}

type ListDocsReq struct {
	g.Meta       `path:"/api/v1/docs" method:"get" tags:"docs" summary:"List docs in a collection"`
	CollectionID string `json:"collectionId" in:"query" v:"required"`
	Locale       string `json:"locale" in:"query"`
	Version      string `json:"version" in:"query"`
}
type ListDocsRes struct {
	Items []*DocView `json:"items"`
}

type ManageDocView struct {
	ID              string    `json:"id"`
	CollectionID    string    `json:"collectionId"`
	CollectionSlug  string    `json:"collectionSlug"`
	CollectionTitle string    `json:"collectionTitle"`
	VersionID       string    `json:"versionId"`
	VersionKey      string    `json:"versionKey"`
	VersionLabel    string    `json:"versionLabel"`
	ParentID        string    `json:"parentId,omitempty"`
	ParentTitle     string    `json:"parentTitle"`
	Slug            string    `json:"slug"`
	SlugPath        string    `json:"slugPath"`
	Title           string    `json:"title"`
	Excerpt         string    `json:"excerpt"`
	Status          string    `json:"status"`
	Locale          string    `json:"locale"`
	SortOrder       int       `json:"sortOrder"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type ManageDocCountsView struct {
	All       int `json:"all"`
	Draft     int `json:"draft"`
	Published int `json:"published"`
	Archived  int `json:"archived"`
	Issues    int `json:"issues"`
}

type ManageDocsReq struct {
	g.Meta       `path:"/api/v1/manage/docs" method:"get" tags:"docs" summary:"Query docs for the admin workbench"`
	Q            string `json:"q" in:"query"`
	Status       string `json:"status" in:"query"`
	Quality      string `json:"quality" in:"query"`
	CollectionID string `json:"collectionId" in:"query"`
	Version      string `json:"version" in:"query"`
	Locale       string `json:"locale" in:"query"`
	ParentID     string `json:"parentId" in:"query"`
	Sort         string `json:"sort" in:"query"`
	Direction    string `json:"direction" in:"query"`
	Page         int    `json:"page" in:"query"`
	Size         int    `json:"size" in:"query"`
}

type ManageDocsRes struct {
	Items  []*ManageDocView    `json:"items"`
	Total  int                 `json:"total"`
	Page   int                 `json:"page"`
	Size   int                 `json:"size"`
	Counts ManageDocCountsView `json:"counts"`
}

type GetDocReq struct {
	g.Meta `path:"/api/v1/docs/{id}" method:"get" tags:"docs" summary:"Get doc by ID"`
	ID     string `json:"id" in:"path" v:"required"`
}
type GetDocRes struct {
	Doc *DocView `json:"doc"`
}

type GetPublicDocByPathReq struct {
	g.Meta     `path:"/api/v1/docs/by-path" method:"get" tags:"docs" summary:"Get published doc by collection path"`
	Collection string `json:"collection" in:"query" v:"required"`
	Path       string `json:"path" in:"query" v:"required"`
	Locale     string `json:"locale" in:"query"`
	Version    string `json:"version" in:"query"`
}
type GetPublicDocByPathRes struct {
	Doc       *DocView                  `json:"doc"`
	Discovery *discovery.PageProjection `json:"discovery,omitempty"`
}

type SearchDocsReq struct {
	g.Meta     `path:"/api/v1/docs/search" method:"get" tags:"docs" summary:"Search published docs"`
	Q          string `json:"q" in:"query"`
	Collection string `json:"collection" in:"query"`
	Locale     string `json:"locale" in:"query"`
	Version    string `json:"version" in:"query"`
}
type SearchDocsRes struct {
	Items  []*DocTreeNodeView `json:"items"`
	Total  int                `json:"total"`
	Facets SearchFacetsView   `json:"facets"`
}

type SearchFacetsView struct {
	Collections []*SearchCollectionFacetView `json:"collections"`
}

type SearchCollectionFacetView struct {
	ID    string `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Count int    `json:"count"`
}

type CreateDocReq struct {
	g.Meta         `path:"/api/v1/docs" method:"post" tags:"docs" summary:"Create doc (admin)"`
	CollectionID   string `json:"collectionId" v:"required"`
	VersionID      string `json:"versionId"`
	ParentID       string `json:"parentId"`
	Slug           string `json:"slug"`
	Title          string `json:"title" v:"required"`
	Content        string `json:"content"`
	SEOTitle       string `json:"seoTitle"`
	SEODescription string `json:"seoDescription"`
	Locale         string `json:"locale"`
	TranslationKey string `json:"translationKey"`
	SortOrder      int    `json:"sortOrder"`
}
type CreateDocRes struct {
	Doc *DocView `json:"doc"`
}

type UpdateDocReq struct {
	g.Meta         `path:"/api/v1/docs/{id}" method:"patch" tags:"docs" summary:"Update doc (admin)"`
	ID             string  `json:"id" in:"path" v:"required"`
	Title          *string `json:"title"`
	Slug           *string `json:"slug"`
	Content        *string `json:"content"`
	Excerpt        *string `json:"excerpt"`
	SEOTitle       *string `json:"seoTitle"`
	SEODescription *string `json:"seoDescription"`
	Status         *string `json:"status"`
	Locale         *string `json:"locale"`
	VersionID      *string `json:"versionId"`
	TranslationKey *string `json:"translationKey"`
	SortOrder      *int    `json:"sortOrder"`
	ParentID       *string `json:"parentId"`
}
type UpdateDocRes struct {
	Doc *DocView `json:"doc"`
}

type PublishDocReq struct {
	g.Meta `path:"/api/v1/docs/{id}/publish" method:"post" tags:"docs" summary:"Publish doc (admin)"`
	ID     string `json:"id" in:"path" v:"required"`
}
type PublishDocRes struct {
	Doc *DocView `json:"doc"`
}

type ArchiveDocReq struct {
	g.Meta `path:"/api/v1/docs/{id}/archive" method:"post" tags:"docs" summary:"Archive doc (admin)"`
	ID     string `json:"id" in:"path" v:"required"`
}
type ArchiveDocRes struct {
	Doc *DocView `json:"doc"`
}

type DeleteDocReq struct {
	g.Meta `path:"/api/v1/docs/{id}" method:"delete" tags:"docs" summary:"Delete doc (admin)"`
	ID     string `json:"id" in:"path" v:"required"`
}
type DeleteDocRes struct {
	Deleted bool `json:"deleted"`
}

type GetCollectionTreeReq struct {
	g.Meta  `path:"/api/v1/collections/{slug}/tree" method:"get" tags:"docs" summary:"Get collection doc tree (public)"`
	Slug    string `json:"slug" in:"path" v:"required"`
	Locale  string `json:"locale" in:"query"`
	Version string `json:"version" in:"query"`
}
type GetCollectionTreeRes struct {
	Collection *CollectionView    `json:"collection"`
	Tree       []*DocTreeNodeView `json:"tree"`
}

type GetManageCollectionTreeReq struct {
	g.Meta  `path:"/api/v1/manage/collections/{slug}/tree" method:"get" tags:"docs" summary:"Get collection doc tree (admin)"`
	Slug    string `json:"slug" in:"path" v:"required"`
	Locale  string `json:"locale" in:"query"`
	Version string `json:"version" in:"query"`
}
type GetManageCollectionTreeRes struct {
	Collection *CollectionView `json:"collection"`
	Tree       []*DocView      `json:"tree"`
}
