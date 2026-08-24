package docsanalytics

import (
	"context"
	"testing"
	"time"

	"github.com/yueli-official/docs/api/internal/docstraffic"
	"github.com/yueli-official/docs/api/internal/model"
	"github.com/yueli-official/foundation/go/traffic"
)

type seedStore struct {
	fakeStore
	searchIDs map[string]bool
}

func (store *seedStore) SeedAnalyticsSearch(_ context.Context, id, _, _, _ string, _ int, _ time.Time) error {
	store.searchIDs[id] = true
	return nil
}

func TestSeedLocalIsIdempotentAndNeverWritesFutureEvents(t *testing.T) {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Fatal(err)
	}
	now := time.Date(2026, 8, 24, 5, 30, 0, 0, location)
	catalog := traffic.MustCompile(docstraffic.Definition("Asia/Shanghai"))
	trafficModule, err := traffic.NewMemory(catalog, traffic.MemoryOptions{
		Clock: func() time.Time { return now }, Secret: []byte("docs-local-seed-test-secret-32-bytes"),
	})
	if err != nil {
		t.Fatal(err)
	}
	documents := map[string]*model.AnalyticsDocument{}
	for _, id := range localWeightedDocumentIDs {
		documents[id] = &model.AnalyticsDocument{ID: id, Title: id, CollectionSlug: "seed", VersionKey: "default", SlugPath: id, Locale: "en"}
	}
	store := &seedStore{fakeStore: fakeStore{documents: documents}, searchIDs: map[string]bool{}}
	module, err := New(store, trafficModule, "Asia/Shanghai")
	if err != nil {
		t.Fatal(err)
	}
	if err := SeedLocal(context.Background(), module, now); err != nil {
		t.Fatal(err)
	}
	first, err := module.Overview(context.Background(), 7, now)
	if err != nil {
		t.Fatal(err)
	}
	if first.Current.Views == 0 || len(store.searchIDs) != 7*len(localSearches) {
		t.Fatalf("seed result = views %d, searches %d", first.Current.Views, len(store.searchIDs))
	}
	if err := SeedLocal(context.Background(), module, now); err != nil {
		t.Fatal(err)
	}
	second, err := module.Overview(context.Background(), 7, now)
	if err != nil {
		t.Fatal(err)
	}
	if second.Current.Views != first.Current.Views || len(store.searchIDs) != 7*len(localSearches) {
		t.Fatalf("seed replay changed totals: %d -> %d", first.Current.Views, second.Current.Views)
	}
}
