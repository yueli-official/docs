package v1

import "github.com/gogf/gf/v2/frame/g"

type RecordDocumentViewReq struct {
	g.Meta     `path:"/api/v1/docs/{id}/view" method:"post" tags:"docs" summary:"Record an idempotent public document view"`
	ID         string `json:"id" in:"path" v:"required"`
	EventID    string `json:"eventId" v:"required|length:16,200"`
	OccurredAt string `json:"occurredAt" v:"required"`
	Source     string `json:"source" v:"length:0,200"`
}

type RecordDocumentViewRes struct {
	Ok        bool  `json:"ok"`
	Counted   bool  `json:"counted"`
	Replay    bool  `json:"replay"`
	ViewCount int64 `json:"viewCount"`
}

type DashboardOverviewReq struct {
	g.Meta `path:"/api/v1/dashboard/overview" method:"get" tags:"docs" summary:"Get documentation usage analytics"`
	Days   int `json:"days"`
}

type DashboardTrafficPointView struct {
	Day               string `json:"day"`
	Views             int64  `json:"views"`
	UniqueVisitorDays int64  `json:"uniqueVisitorDays"`
}

type DashboardTopDocumentView struct {
	ID                string `json:"id"`
	Title             string `json:"title"`
	CollectionSlug    string `json:"collectionSlug"`
	VersionKey        string `json:"versionKey"`
	SlugPath          string `json:"slugPath"`
	Locale            string `json:"locale"`
	Views             int64  `json:"views"`
	UniqueVisitorDays int64  `json:"uniqueVisitorDays"`
}

type DashboardTrafficSourceView struct {
	Source string `json:"source"`
	Views  int64  `json:"views"`
}

type DashboardSearchQueryView struct {
	Query       string `json:"query"`
	Searches    int64  `json:"searches"`
	ZeroResults int64  `json:"zeroResults"`
}

type DashboardOverviewRes struct {
	Days                      int                           `json:"days"`
	AllTimeViews              int64                         `json:"allTimeViews"`
	AllTimeUniqueVisitorDays  int64                         `json:"allTimeUniqueVisitorDays"`
	PeriodViews               int64                         `json:"periodViews"`
	PeriodUniqueVisitorDays   int64                         `json:"periodUniqueVisitorDays"`
	PreviousPeriodViews       int64                         `json:"previousPeriodViews"`
	PreviousUniqueVisitorDays int64                         `json:"previousUniqueVisitorDays"`
	PeriodSearches            int64                         `json:"periodSearches"`
	PreviousPeriodSearches    int64                         `json:"previousPeriodSearches"`
	ZeroResultSearches        int64                         `json:"zeroResultSearches"`
	Series                    []*DashboardTrafficPointView  `json:"series"`
	TopDocuments              []*DashboardTopDocumentView   `json:"topDocuments"`
	TopSources                []*DashboardTrafficSourceView `json:"topSources"`
	TopSearches               []*DashboardSearchQueryView   `json:"topSearches"`
}
