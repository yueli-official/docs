package model

// CollectionLocale is one language explicitly enabled by a document collection.
type CollectionLocale struct {
	CollectionID string `json:"collectionId" orm:"collection_id"`
	Locale       string `json:"locale" orm:"locale"`
	Label        string `json:"label" orm:"label"`
	HTMLLang     string `json:"htmlLang" orm:"html_lang"`
	Direction    string `json:"direction" orm:"direction"`
	IsDefault    bool   `json:"isDefault" orm:"is_default"`
	Enabled      bool   `json:"enabled" orm:"enabled"`
	SortOrder    int    `json:"sortOrder" orm:"sort_order"`
	DocCount     int    `json:"docCount" orm:"doc_count"`
}
