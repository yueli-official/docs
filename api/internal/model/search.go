package model

type SearchCollectionFacet struct {
	ID    string `json:"id" orm:"id"`
	Slug  string `json:"slug" orm:"slug"`
	Title string `json:"title" orm:"title"`
	Count int    `json:"count" orm:"count"`
}

type SearchResult struct {
	Items            []*Doc
	Total            int
	CollectionFacets []*SearchCollectionFacet
}
