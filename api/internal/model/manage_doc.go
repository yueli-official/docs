package model

import "time"

// ManageDocsQuery is the normalized, allowlisted query accepted by the admin
// document workbench. Page is one-based.
type ManageDocsQuery struct {
	Q            string
	Status       string
	Quality      string
	CollectionID string
	Version      string
	Locale       string
	ParentID     string
	Sort         string
	Direction    string
	Page         int
	Size         int
	OwnerSub     string
}

// ManageDoc is the denormalized row required by the admin list. Tree editing
// continues to use the dedicated collection tree endpoint.
type ManageDoc struct {
	ID              string    `json:"id" orm:"id"`
	CollectionID    string    `json:"collectionId" orm:"collection_id"`
	CollectionSlug  string    `json:"collectionSlug" orm:"collection_slug"`
	CollectionTitle string    `json:"collectionTitle" orm:"collection_title"`
	VersionID       string    `json:"versionId" orm:"version_id"`
	VersionKey      string    `json:"versionKey" orm:"version_key"`
	VersionLabel    string    `json:"versionLabel" orm:"version_label"`
	ParentID        string    `json:"parentId" orm:"parent_id"`
	ParentTitle     string    `json:"parentTitle" orm:"parent_title"`
	Slug            string    `json:"slug" orm:"slug"`
	SlugPath        string    `json:"slugPath" orm:"slug_path"`
	Title           string    `json:"title" orm:"title"`
	Excerpt         string    `json:"excerpt" orm:"excerpt"`
	Status          string    `json:"status" orm:"status"`
	Locale          string    `json:"locale" orm:"locale"`
	SortOrder       int       `json:"sortOrder" orm:"sort_order"`
	UpdatedAt       time.Time `json:"updatedAt" orm:"updated_at"`
}

type ManageDocCounts struct {
	All       int `json:"all" orm:"all_count"`
	Draft     int `json:"draft" orm:"draft_count"`
	Published int `json:"published" orm:"published_count"`
	Archived  int `json:"archived" orm:"archived_count"`
	Issues    int `json:"issues" orm:"issues_count"`
}

type ManageDocsResult struct {
	Items  []*ManageDoc
	Total  int
	Counts ManageDocCounts
}
