package v1

import "github.com/gogf/gf/v2/frame/g"

type CollectionVersionView struct {
	ID              string `json:"id"`
	CollectionID    string `json:"collectionId"`
	Key             string `json:"key"`
	Label           string `json:"label"`
	Status          string `json:"status"`
	IsDefault       bool   `json:"isDefault"`
	SortOrder       int    `json:"sortOrder"`
	SourceVersionID string `json:"sourceVersionId,omitempty"`
}

type ListCollectionVersionsReq struct {
	g.Meta       `path:"/api/v1/collections/{id}/versions" method:"get" tags:"docs" summary:"List collection versions"`
	CollectionID string `json:"id" in:"path" v:"required"`
}
type ListCollectionVersionsRes struct {
	Items []*CollectionVersionView `json:"items"`
}

type CreateCollectionVersionReq struct {
	g.Meta          `path:"/api/v1/collections/{id}/versions" method:"post" tags:"docs" summary:"Create collection version (admin)"`
	CollectionID    string `json:"id" in:"path" v:"required"`
	Key             string `json:"key" v:"required"`
	Label           string `json:"label"`
	Status          string `json:"status"`
	SourceVersionID string `json:"sourceVersionId"`
}
type CreateCollectionVersionRes struct {
	g.Meta  `status:"201"`
	Version *CollectionVersionView `json:"version"`
}

type UpdateCollectionVersionReq struct {
	g.Meta       `path:"/api/v1/manage/collections/{id}/versions/{versionId}" method:"patch" tags:"docs" summary:"Update collection version metadata"`
	CollectionID string `json:"id" in:"path" v:"required"`
	VersionID    string `json:"versionId" in:"path" v:"required"`
	Label        string `json:"label" v:"required"`
	Status       string `json:"status" v:"required"`
	IsDefault    bool   `json:"isDefault"`
	SortOrder    int    `json:"sortOrder"`
}
type UpdateCollectionVersionRes struct {
	Version *CollectionVersionView `json:"version"`
}
