package model

// Doc is a single page within a collection, optionally nested under a parent.
type Doc struct {
	ID             string `json:"id" orm:"id"`
	CollectionID   string `json:"collectionId" orm:"collection_id"`
	VersionID      string `json:"versionId" orm:"version_id"`
	ParentID       string `json:"parentId" orm:"parent_id"`
	Slug           string `json:"slug" orm:"slug"`
	Title          string `json:"title" orm:"title"`
	Content        string `json:"content" orm:"content"`
	Excerpt        string `json:"excerpt" orm:"excerpt"`
	SEOTitle       string `json:"seoTitle" orm:"seo_title"`
	SEODescription string `json:"seoDescription" orm:"seo_description"`
	Status         string `json:"status" orm:"status"`
	Locale         string `json:"locale" orm:"locale"`
	TranslationKey string `json:"translationKey" orm:"translation_key"`
	BadgeText      string `json:"badgeText" orm:"badge_text"`
	BadgeIcon      string `json:"badgeIcon" orm:"badge_icon"`
	SortOrder      int    `json:"sortOrder" orm:"sort_order"`
	AuthorSub      string `json:"authorSub" orm:"author_sub"`
	SearchRevision uint64 `json:"-" orm:"search_revision"`
	Children       []*Doc `json:"children,omitempty" orm:"-"`
}
