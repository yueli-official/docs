package v1

import "github.com/gogf/gf/v2/frame/g"

// CollectionView is the public JSON projection of a collection.
type CollectionView struct {
	ID                      string `json:"id"`
	Slug                    string `json:"slug"`
	Title                   string `json:"title"`
	Description             string `json:"description"`
	CoverAssetID            string `json:"coverAssetId"`
	CoverURL                string `json:"coverUrl"`
	Icon                    string `json:"icon"`
	SortOrder               int    `json:"sortOrder"`
	ReleaseFamilyID         string `json:"releaseFamilyId,omitempty"`
	ReleaseFamilyName       string `json:"releaseFamilyName,omitempty"`
	SemanticVersion         string `json:"semanticVersion,omitempty"`
	DerivedFromCollectionID string `json:"derivedFromCollectionId,omitempty"`
	DocCount                int    `json:"docCount"`
}

type ListCollectionsReq struct {
	g.Meta `path:"/api/v1/collections" method:"get" tags:"docs" summary:"List collections"`
}
type ListCollectionsRes struct {
	Items []*CollectionView `json:"items"`
}

type GetCollectionReq struct {
	g.Meta `path:"/api/v1/collections/{slug}" method:"get" tags:"docs" summary:"Get collection by slug"`
	Slug   string `json:"slug" in:"path" v:"required"`
}
type GetCollectionRes struct {
	*CollectionView
}

type GetCollectionReleasesReq struct {
	g.Meta `path:"/api/v1/collections/{slug}/releases" method:"get" tags:"docs" summary:"List related collection releases"`
	Slug   string `json:"slug" in:"path" v:"required"`
}
type GetCollectionReleasesRes struct {
	Items []*CollectionView `json:"items"`
}

type CreateCollectionReq struct {
	g.Meta          `path:"/api/v1/collections" method:"post" tags:"docs" summary:"Create collection (admin)"`
	Title           string `json:"title" v:"required"`
	Slug            string `json:"slug"`
	Description     string `json:"description"`
	Cover           string `json:"cover"`
	Icon            string `json:"icon"`
	DefaultLocale   string `json:"defaultLocale" v:"required"`
	SemanticVersion string `json:"semanticVersion" v:"required"`
}
type CreateCollectionRes struct {
	g.Meta `status:"201"`
	*CollectionView
}

type UpdateCollectionReq struct {
	g.Meta      `path:"/api/v1/collections/{id}" method:"patch" tags:"docs" summary:"Update collection (admin)"`
	ID          string `json:"id" in:"path" v:"required"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	Icon        string `json:"icon"`
}
type UpdateCollectionRes struct {
	Collection *CollectionView `json:"collection"`
}

type CollectionCoverInitReq struct {
	g.Meta   `path:"/api/v1/collections/{id}/cover" method:"post" tags:"docs" summary:"Begin a collection cover image upload"`
	ID       string `json:"id" in:"path" v:"required"`
	Filename string `json:"filename" v:"required"`
	Mime     string `json:"mime"`
	Size     int64  `json:"size"`
}
type CollectionCoverInitRes struct {
	g.Meta        `status:"201"`
	UploadURL     string            `json:"uploadUrl"`
	UploadToken   string            `json:"uploadToken"`
	UploadHeaders map[string]string `json:"uploadHeaders,omitempty"`
}

type CollectionCoverFinalizeReq struct {
	g.Meta      `path:"/api/v1/collections/{id}/cover/finalize" method:"post" tags:"docs" summary:"Finalize a collection cover image upload"`
	ID          string `json:"id" in:"path" v:"required"`
	UploadToken string `json:"uploadToken" v:"required"`
}
type CollectionCoverFinalizeRes struct {
	Collection *CollectionView `json:"collection"`
	CoverURL   string          `json:"coverUrl"`
}

type DeleteCollectionReq struct {
	g.Meta `path:"/api/v1/collections/{id}" method:"delete" tags:"docs" summary:"Delete collection (admin)"`
	ID     string `json:"id" in:"path" v:"required"`
}
type DeleteCollectionRes struct {
	g.Meta `status:"204"`
}

type CloneCollectionReleaseReq struct {
	g.Meta                `path:"/api/v1/manage/collections/{id}/clone-release" method:"post" tags:"docs" summary:"Clone a collection as a related semantic release"`
	ID                    string `json:"id" in:"path" v:"required"`
	SourceSemanticVersion string `json:"sourceSemanticVersion"`
	TargetSemanticVersion string `json:"targetSemanticVersion" v:"required"`
	Title                 string `json:"title" v:"required"`
	Slug                  string `json:"slug"`
}
type CloneCollectionReleaseRes struct {
	g.Meta     `status:"201"`
	Collection *CollectionView `json:"collection"`
}

type InitializeCollectionReleaseReq struct {
	g.Meta          `path:"/api/v1/manage/collections/{id}/release" method:"post" tags:"docs" summary:"Initialize a collection as a semantic release"`
	ID              string `json:"id" in:"path" v:"required"`
	SemanticVersion string `json:"semanticVersion" v:"required"`
}
type InitializeCollectionReleaseRes struct {
	Collection *CollectionView `json:"collection"`
}
