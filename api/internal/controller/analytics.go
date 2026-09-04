package controller

import (
	"context"
	"strings"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
	foundationauth "github.com/yueli-official/foundation/go/auth"
	"github.com/yueli-official/foundation/go/authorization"
	"github.com/yueli-official/foundation/go/traffic"

	v1 "github.com/yueli-official/docs/api/api/v1"
	"github.com/yueli-official/docs/api/internal/docsanalytics"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/docserr"
)

type PublicAnalytics struct {
	module   *docsanalytics.Module
	verifier *foundationauth.Verifier
}

func NewPublicAnalytics(module *docsanalytics.Module, verifier *foundationauth.Verifier) *PublicAnalytics {
	return &PublicAnalytics{module: module, verifier: verifier}
}

func (controller *PublicAnalytics) RecordDocumentView(ctx context.Context, req *v1.RecordDocumentViewReq) (*v1.RecordDocumentViewRes, error) {
	occurredAt, err := time.Parse(time.RFC3339Nano, req.OccurredAt)
	if err != nil {
		return nil, docserr.InvalidInput("occurredAt must be an RFC3339 timestamp")
	}
	ip, userAgent := analyticsClientMeta(ctx)
	visitorSeed := anonymousAnalyticsVisitorSeed(ip, userAgent)
	if subject := optionalAnalyticsSubject(ctx, controller.verifier); subject != "" {
		visitorSeed = []byte("subject\x00" + subject)
	}
	result, err := controller.module.RecordView(ctx, req.ID, docsanalytics.ViewInput{
		EventID: req.EventID, OccurredAt: occurredAt,
		Class: classifyAnalyticsVisit(userAgent), VisitorSeed: visitorSeed,
		Source: normalizeAnalyticsSource(req.Source),
	})
	if err != nil {
		switch {
		case err == docsanalytics.ErrDocumentNotFound:
			return nil, docserr.NotFound(req.ID)
		case traffic.IsKind(err, traffic.ErrorInvalidInput), traffic.IsKind(err, traffic.ErrorConflict):
			return nil, docserr.InvalidInput("traffic_event_invalid")
		default:
			return nil, err
		}
	}
	return &v1.RecordDocumentViewRes{
		Ok: true, Counted: result.Counted, Replay: result.Replay,
		ViewCount: result.ResourceTotals.Views,
	}, nil
}

type Dashboard struct{ module *docsanalytics.Module }

func NewDashboard(module *docsanalytics.Module) *Dashboard { return &Dashboard{module: module} }

func (controller *Dashboard) Overview(ctx context.Context, req *v1.DashboardOverviewReq) (*v1.DashboardOverviewRes, error) {
	service := authorizationService(ctx)
	if service == nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	constraint, err := service.Runtime().Plan(ctx, authorization.QueryRequest{
		Subject: service.Subject(ctx), Capability: docsauthz.CapabilityDocumentRead,
		ScopeID: docsauthz.RootScopeID,
	})
	if err != nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	if constraint.Kind != authorization.QueryAll {
		return nil, docserr.Forbidden()
	}
	overview, err := controller.module.Overview(ctx, req.Days, time.Now())
	if err != nil {
		if err == docsanalytics.ErrUnsupportedWindow {
			return nil, docserr.InvalidInput("traffic_query_invalid")
		}
		return nil, err
	}
	series := make([]*v1.DashboardTrafficPointView, 0, len(overview.Series))
	for _, point := range overview.Series {
		series = append(series, &v1.DashboardTrafficPointView{
			Day: point.Day.String(), Views: point.Totals.Views,
			UniqueVisitorDays: point.Totals.UniqueVisitorDays,
		})
	}
	topDocuments := make([]*v1.DashboardTopDocumentView, 0, len(overview.TopDocuments))
	for _, item := range overview.TopDocuments {
		topDocuments = append(topDocuments, &v1.DashboardTopDocumentView{
			ID: item.Document.ID, Title: item.Document.Title,
			CollectionSlug: item.Document.CollectionSlug, VersionKey: item.Document.VersionKey,
			SlugPath: item.Document.SlugPath, Locale: item.Document.Locale,
			Views: item.Totals.Views, UniqueVisitorDays: item.Totals.UniqueVisitorDays,
		})
	}
	topSources := make([]*v1.DashboardTrafficSourceView, 0, len(overview.TopSources))
	for _, item := range overview.TopSources {
		topSources = append(topSources, &v1.DashboardTrafficSourceView{Source: item.Source, Views: item.Views})
	}
	topSearches := make([]*v1.DashboardSearchQueryView, 0, len(overview.TopSearches))
	for _, item := range overview.TopSearches {
		topSearches = append(topSearches, &v1.DashboardSearchQueryView{
			Query: item.Query, Searches: item.Searches, ZeroResults: item.ZeroResults,
		})
	}
	return &v1.DashboardOverviewRes{
		Days:         overview.Days,
		AllTimeViews: overview.AllTime.Views, AllTimeUniqueVisitorDays: overview.AllTime.UniqueVisitorDays,
		PeriodViews: overview.Current.Views, PeriodUniqueVisitorDays: overview.Current.UniqueVisitorDays,
		PreviousPeriodViews: overview.Previous.Views, PreviousUniqueVisitorDays: overview.Previous.UniqueVisitorDays,
		PeriodSearches: overview.Searches.Searches, PreviousPeriodSearches: overview.PreviousSearches.Searches,
		ZeroResultSearches: overview.Searches.ZeroResults,
		Series:             series, TopDocuments: topDocuments, TopSources: topSources, TopSearches: topSearches,
	}, nil
}

func normalizeAnalyticsSource(value string) string {
	value = strings.TrimSuffix(strings.TrimPrefix(strings.ToLower(strings.TrimSpace(value)), "www."), ".")
	if value == "" || value == "direct" {
		return "direct"
	}
	if value == "internal" {
		return value
	}
	if len(value) > 200 || !strings.Contains(value, ".") || strings.Contains(value, "..") {
		return "direct"
	}
	for _, character := range value {
		if (character < 'a' || character > 'z') && (character < '0' || character > '9') && character != '.' && character != '-' {
			return "direct"
		}
	}
	return value
}

func classifyAnalyticsVisit(userAgent string) traffic.VisitClass {
	value := strings.ToLower(userAgent)
	for _, marker := range []string{
		"bot", "crawler", "spider", "slurp", "headless", "monitoring",
		"facebookexternalhit", "twitterbot", "preview",
	} {
		if strings.Contains(value, marker) {
			return traffic.VisitBot
		}
	}
	if strings.TrimSpace(value) == "" {
		return traffic.VisitUnknown
	}
	return traffic.VisitHuman
}

func analyticsClientMeta(ctx context.Context) (string, string) {
	request := ghttp.RequestFromCtx(ctx)
	if request == nil {
		return "", ""
	}
	return request.GetClientIp(), request.Request.UserAgent()
}

func anonymousAnalyticsVisitorSeed(ip, userAgent string) []byte {
	ip = strings.TrimSpace(ip)
	userAgent = strings.TrimSpace(userAgent)
	if ip == "" && userAgent == "" {
		return nil
	}
	return []byte("network\x00" + ip + "\x00" + userAgent)
}

func optionalAnalyticsSubject(ctx context.Context, verifier *foundationauth.Verifier) string {
	raw := bearerOf(ctx)
	if raw == "" || verifier == nil {
		return ""
	}
	principal, err := verifier.Verify(ctx, raw)
	if err != nil || principal == nil || strings.TrimSpace(principal.Subject) == "" {
		return ""
	}
	kind, _ := principal.Claim("subject_kind")
	if kind != "user" {
		return ""
	}
	return principal.Subject
}
