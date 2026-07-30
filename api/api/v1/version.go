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
	Version *CollectionVersionView `json:"version"`
}
