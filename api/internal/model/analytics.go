package model

type AnalyticsDocument struct {
	ID             string `orm:"id"`
	Title          string `orm:"title"`
	CollectionSlug string `orm:"collection_slug"`
	VersionKey     string `orm:"version_key"`
	SlugPath       string `orm:"slug_path"`
	Locale         string `orm:"locale"`
}

type AnalyticsTrafficSource struct {
	Source string `orm:"source"`
	Views  int64  `orm:"views"`
}

type AnalyticsSearchSummary struct {
	Searches    int64 `orm:"searches"`
	ZeroResults int64 `orm:"zero_results"`
}

type AnalyticsSearchQuery struct {
	Query       string `orm:"query"`
	Searches    int64  `orm:"searches"`
	ZeroResults int64  `orm:"zero_results"`
}
