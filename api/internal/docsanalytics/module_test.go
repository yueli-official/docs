package docsanalytics

import (
	"context"
	"testing"
	"time"

	"github.com/yueli-official/docs/api/internal/docstraffic"
	"github.com/yueli-official/docs/api/internal/model"
	"github.com/yueli-official/foundation/go/traffic"
)

type fakeStore struct {
	documents map[string]*model.AnalyticsDocument
	sources   []model.AnalyticsTrafficSource
	search    model.AnalyticsSearchSummary
	topSearch []model.AnalyticsSearchQuery
}

func (store *fakeStore) PublishedAnalyticsDocumentsByIDs(_ context.Context, ids []string) ([]*model.AnalyticsDocument, error) {
	result := make([]*model.AnalyticsDocument, 0, len(ids))
	for _, id := range ids {
		if document := store.documents[id]; document != nil {
			result = append(result, document)
		}
	}
	return result, nil
}

func (store *fakeStore) RecordAnalyticsTrafficSource(_ context.Context, _, _, source string) error {
	store.sources = append(store.sources, model.AnalyticsTrafficSource{Source: source, Views: 1})
	return nil
}

func (store *fakeStore) AnalyticsTrafficSources(context.Context, time.Time, time.Time, int) ([]model.AnalyticsTrafficSource, error) {
	return store.sources, nil
}

func (store *fakeStore) AnalyticsSearchSummary(context.Context, time.Time, time.Time) (model.AnalyticsSearchSummary, error) {
	return store.search, nil
}

func (store *fakeStore) AnalyticsTopSearches(context.Context, time.Time, time.Time, int) ([]model.AnalyticsSearchQuery, error) {
	return store.topSearch, nil
}

func TestModuleRecordsViewsAndBuildsOneDashboardOverview(t *testing.T) {
	clock := time.Date(2026, 8, 24, 12, 0, 0, 0, time.UTC)
	catalog := traffic.MustCompile(docstraffic.Definition("Asia/Shanghai"))
	trafficModule, err := traffic.NewMemory(catalog, traffic.MemoryOptions{
		Clock:  func() time.Time { return clock },
		Secret: []byte("docs-analytics-test-secret-32-bytes"),
	})
	if err != nil {
		t.Fatal(err)
	}
	store := &fakeStore{
		documents: map[string]*model.AnalyticsDocument{
			"doc-a": {ID: "doc-a", Title: "安装", CollectionSlug: "guide", VersionKey: "v2", SlugPath: "start/install", Locale: "zh-CN"},
			"doc-b": {ID: "doc-b", Title: "部署", CollectionSlug: "guide", VersionKey: "default", SlugPath: "deploy", Locale: "en"},
		},
		search:    model.AnalyticsSearchSummary{Searches: 8, ZeroResults: 2},
		topSearch: []model.AnalyticsSearchQuery{{Query: "部署", Searches: 3, ZeroResults: 0}},
	}
	module, err := New(store, trafficModule, "Asia/Shanghai")
	if err != nil {
		t.Fatal(err)
	}

	for _, input := range []struct{ id, event, source string }{
		{"doc-a", "019c0000-0000-7000-8000-000000000001", "google.com"},
		{"doc-b", "019c0000-0000-7000-8000-000000000002", "direct"},
	} {
		result, recordErr := module.RecordView(context.Background(), input.id, ViewInput{
			EventID: input.event, OccurredAt: clock.Add(-time.Hour),
			Class: traffic.VisitHuman, VisitorSeed: []byte("visitor-1"), Source: input.source,
		})
		if recordErr != nil {
			t.Fatal(recordErr)
		}
		if !result.Counted {
			t.Fatalf("view %s was not counted", input.id)
		}
	}

	overview, err := module.Overview(context.Background(), 14, clock)
	if err != nil {
		t.Fatal(err)
	}
	if overview.AllTime.Views != 2 || overview.Current.Views != 2 {
		t.Fatalf("overview totals = %#v", overview)
	}
	if len(overview.Series) != 14 || len(overview.TopDocuments) != 2 {
		t.Fatalf("overview series/top = %d/%d", len(overview.Series), len(overview.TopDocuments))
	}
	if overview.TopDocuments[0].Document.Locale == "" || overview.TopDocuments[0].Document.VersionKey == "" {
		t.Fatalf("top document routing metadata = %#v", overview.TopDocuments[0].Document)
	}
	if overview.Searches.Searches != 8 || overview.Searches.ZeroResults != 2 || len(overview.TopSearches) != 1 {
		t.Fatalf("overview search = %#v", overview)
	}
}

func TestModuleRejectsUnknownDocumentsAndUnsupportedWindows(t *testing.T) {
	clock := time.Date(2026, 8, 24, 12, 0, 0, 0, time.UTC)
	catalog := traffic.MustCompile(docstraffic.Definition("UTC"))
	trafficModule, err := traffic.NewMemory(catalog, traffic.MemoryOptions{
		Clock: func() time.Time { return clock }, Secret: []byte("docs-analytics-test-secret-32-bytes"),
	})
	if err != nil {
		t.Fatal(err)
	}
	module, err := New(&fakeStore{documents: map[string]*model.AnalyticsDocument{}}, trafficModule, "UTC")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := module.RecordView(context.Background(), "missing", ViewInput{EventID: "019c0000-0000-7000-8000-000000000003", OccurredAt: clock}); err != ErrDocumentNotFound {
		t.Fatalf("RecordView error = %v", err)
	}
	if _, err := module.Overview(context.Background(), 9, clock); err != ErrUnsupportedWindow {
		t.Fatalf("Overview error = %v", err)
	}
}
