package model

// HomeQuickLink is one configurable shortcut on the public docs homepage.
type HomeQuickLink struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Icon           string `json:"icon"`
	To             string `json:"to"`
	CollectionSlug string `json:"collectionSlug"`
	SortOrder      int    `json:"sortOrder"`
	Enabled        bool   `json:"enabled"`
}

// HomeConfig stores the docs homepage curation controlled from manage.
type HomeConfig struct {
	QuickLinks          []*HomeQuickLink `json:"quickLinks"`
	FeaturedCollections []string         `json:"featuredCollections"`
}
