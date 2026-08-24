// Package docsanalytics owns Docs reading and search analytics over Foundation Traffic.
package docsanalytics

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/yueli-official/docs/api/internal/docstraffic"
	"github.com/yueli-official/docs/api/internal/model"
	"github.com/yueli-official/foundation/go/traffic"
)

var (
	ErrDocumentNotFound  = errors.New("analytics document not found")
	ErrUnsupportedWindow = errors.New("analytics window must be 7, 14, or 30 days")
)

type Store interface {
	PublishedAnalyticsDocumentsByIDs(context.Context, []string) ([]*model.AnalyticsDocument, error)
	RecordAnalyticsTrafficSource(context.Context, string, string, string) error
	AnalyticsTrafficSources(context.Context, time.Time, time.Time, int) ([]model.AnalyticsTrafficSource, error)
	AnalyticsSearchSummary(context.Context, time.Time, time.Time) (model.AnalyticsSearchSummary, error)
	AnalyticsTopSearches(context.Context, time.Time, time.Time, int) ([]model.AnalyticsSearchQuery, error)
}

type ViewInput struct {
	EventID     string
	OccurredAt  time.Time
	Class       traffic.VisitClass
	VisitorSeed []byte
	Source      string
}

type TopDocument struct {
	Document *model.AnalyticsDocument
	Totals   traffic.Totals
}

type Overview struct {
	Days             int
	AllTime          traffic.Totals
	Current          traffic.Totals
	Previous         traffic.Totals
	Series           []traffic.SeriesPoint
	TopDocuments     []TopDocument
	TopSources       []model.AnalyticsTrafficSource
	Searches         model.AnalyticsSearchSummary
	PreviousSearches model.AnalyticsSearchSummary
	TopSearches      []model.AnalyticsSearchQuery
}

type Module struct {
	store    Store
	traffic  traffic.Module
	location *time.Location
}

func New(store Store, trafficModule traffic.Module, timeZone string) (*Module, error) {
	if store == nil {
		return nil, errors.New("analytics store is required")
	}
	if trafficModule == nil {
		return nil, errors.New("traffic module is required")
	}
	location, err := time.LoadLocation(timeZone)
	if err != nil {
		return nil, fmt.Errorf("load analytics time zone: %w", err)
	}
	return &Module{store: store, traffic: trafficModule, location: location}, nil
}

func (module *Module) RecordView(ctx context.Context, documentID string, input ViewInput) (traffic.RecordResult, error) {
	documents, err := module.store.PublishedAnalyticsDocumentsByIDs(ctx, []string{documentID})
	if err != nil {
		return traffic.RecordResult{}, err
	}
	if len(documents) != 1 || documents[0].ID != documentID {
		return traffic.RecordResult{}, ErrDocumentNotFound
	}
	observation := traffic.Observation{
		EventID:    traffic.EventID(input.EventID),
		Resource:   traffic.Resource{Kind: docstraffic.ResourceDocument, ID: documentID},
		OccurredAt: input.OccurredAt,
		Class:      input.Class,
	}
	if len(input.VisitorSeed) > 0 {
		token, tokenErr := module.traffic.TokenizeVisitor(ctx, input.OccurredAt, input.VisitorSeed)
		if tokenErr != nil {
			return traffic.RecordResult{}, tokenErr
		}
		observation.HasVisitor = true
		observation.VisitorToken = token
	}
	result, err := module.traffic.Record(ctx, observation)
	if err != nil {
		return traffic.RecordResult{}, err
	}
	if result.Counted {
		day := input.OccurredAt.In(module.location).Format(time.DateOnly)
		if err := module.store.RecordAnalyticsTrafficSource(ctx, input.EventID, day, input.Source); err != nil {
			return traffic.RecordResult{}, err
		}
	}
	return result, nil
}

func (module *Module) Overview(ctx context.Context, days int, now time.Time) (Overview, error) {
	if days == 0 {
		days = 14
	}
	if days != 7 && days != 14 && days != 30 {
		return Overview{}, ErrUnsupportedWindow
	}
	localNow := now.In(module.location)
	to := time.Date(localNow.Year(), localNow.Month(), localNow.Day(), 0, 0, 0, 0, module.location).AddDate(0, 0, 1)
	from := to.AddDate(0, 0, -days)
	previousFrom := from.AddDate(0, 0, -days)
	currentRange := traffic.DateRange{
		From: traffic.MustParseDay(from.Format(time.DateOnly)),
		To:   traffic.MustParseDay(to.Format(time.DateOnly)),
	}
	previousRange := traffic.DateRange{
		From: traffic.MustParseDay(previousFrom.Format(time.DateOnly)),
		To:   currentRange.From,
	}
	allTime, err := module.traffic.Summary(ctx, traffic.SummaryQuery{Scope: traffic.InstanceScope()})
	if err != nil {
		return Overview{}, err
	}
	current, err := module.traffic.Summary(ctx, traffic.SummaryQuery{Scope: traffic.InstanceScope(), Range: &currentRange})
	if err != nil {
		return Overview{}, err
	}
	previous, err := module.traffic.Summary(ctx, traffic.SummaryQuery{Scope: traffic.InstanceScope(), Range: &previousRange})
	if err != nil {
		return Overview{}, err
	}
	series, err := module.traffic.Series(ctx, traffic.SeriesQuery{Scope: traffic.InstanceScope(), Range: currentRange})
	if err != nil {
		return Overview{}, err
	}
	topEntries, err := module.traffic.Top(ctx, traffic.TopQuery{
		ResourceKind: docstraffic.ResourceDocument, Range: &currentRange,
		Metric: traffic.RankViews, Limit: 5,
	})
	if err != nil {
		return Overview{}, err
	}
	ids := make([]string, 0, len(topEntries))
	for _, entry := range topEntries {
		ids = append(ids, entry.Resource.ID)
	}
	documents, err := module.store.PublishedAnalyticsDocumentsByIDs(ctx, ids)
	if err != nil {
		return Overview{}, err
	}
	byID := make(map[string]*model.AnalyticsDocument, len(documents))
	for _, document := range documents {
		byID[document.ID] = document
	}
	topDocuments := make([]TopDocument, 0, len(topEntries))
	for _, entry := range topEntries {
		if document := byID[entry.Resource.ID]; document != nil {
			topDocuments = append(topDocuments, TopDocument{Document: document, Totals: entry.Totals})
		}
	}
	topSources, err := module.store.AnalyticsTrafficSources(ctx, from, to, 5)
	if err != nil {
		return Overview{}, err
	}
	searches, err := module.store.AnalyticsSearchSummary(ctx, from, to)
	if err != nil {
		return Overview{}, err
	}
	previousSearches, err := module.store.AnalyticsSearchSummary(ctx, previousFrom, from)
	if err != nil {
		return Overview{}, err
	}
	topSearches, err := module.store.AnalyticsTopSearches(ctx, from, to, 5)
	if err != nil {
		return Overview{}, err
	}
	return Overview{
		Days: days, AllTime: allTime.Totals, Current: current.Totals, Previous: previous.Totals,
		Series: series, TopDocuments: topDocuments, TopSources: topSources,
		Searches: searches, PreviousSearches: previousSearches, TopSearches: topSearches,
	}, nil
}
