package model

// CollectionVersion is a version line inside one docs collection.
type CollectionVersion struct {
	ID              string `json:"id" orm:"id"`
	CollectionID    string `json:"collectionId" orm:"collection_id"`
	Key             string `json:"key" orm:"key"`
	Label           string `json:"label" orm:"label"`
	Status          string `json:"status" orm:"status"`
	IsDefault       bool   `json:"isDefault" orm:"is_default"`
	SortOrder       int    `json:"sortOrder" orm:"sort_order"`
	SourceVersionID string `json:"sourceVersionId" orm:"source_version_id"`
}
