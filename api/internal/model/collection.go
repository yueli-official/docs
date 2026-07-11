package model

// Collection is a tutorial set that groups related docs.
type Collection struct {
	ID           string `json:"id" orm:"id"`
	Slug         string `json:"slug" orm:"slug"`
	Title        string `json:"title" orm:"title"`
	Description  string `json:"description" orm:"description"`
	CoverAssetID string `json:"coverAssetId" orm:"cover_asset_id"`
	CoverURL     string `json:"coverUrl" orm:"cover_url"`
	Icon         string `json:"icon" orm:"icon"`
	SortOrder    int    `json:"sortOrder" orm:"sort_order"`
	AuthorSub    string `json:"authorSub" orm:"author_sub"`
	DocCount     int    `json:"docCount" orm:"-"`
}
