package docsanalytics

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yueli-official/foundation/go/traffic"
)

var localWeightedDocumentIDs = []string{
	"019c52f0-3000-7000-8000-000000000301",
	"019c52f0-3000-7000-8000-000000000301",
	"019c52f0-3000-7000-8000-000000000302",
	"019c52f0-3000-7000-8000-000000000302",
	"019c52f0-3000-7000-8000-000000000302",
	"019c52f0-3000-7000-8000-000000000303",
	"019c52f0-3000-7000-8000-000000000304",
	"019c52f0-3000-7000-8000-000000000304",
	"019c52f0-3000-7000-8000-000000000305",
	"019c52f0-3000-7000-8000-000000000306",
}

var localDailyViews = []int{8, 11, 7, 13, 9, 12, 6}
var localSources = []string{"direct", "google.com", "internal", "zhihu.com", "bing.com"}
var localSearches = []struct {
	query   string
	results int
}{
	{query: "安装", results: 1},
	{query: "部署", results: 1},
	{query: "站点配置", results: 1},
	{query: "排查", results: 1},
	{query: "不存在的命令", results: 0},
}

var localSeedEpoch = time.Date(2026, time.August, 14, 0, 0, 0, 0, time.UTC)

type localSearchSeeder interface {
	SeedAnalyticsSearch(context.Context, string, string, string, string, int, time.Time) error
}

func localSeedDayOffset(day time.Time) int {
	civilDay := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.UTC)
	return int(civilDay.Sub(localSeedEpoch) / (24 * time.Hour))
}

func localSeedIndex(offset, length int) int {
	index := offset % length
	if index < 0 {
		index += length
	}
	return index
}

// SeedLocal installs deterministic local-only reading, source and search events.
func SeedLocal(ctx context.Context, module *Module, now time.Time) error {
	if module == nil {
		return nil
	}
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	for windowOffset := range localDailyViews {
		dayStart := today.AddDate(0, 0, windowOffset-(len(localDailyViews)-1))
		stableOffset := localSeedDayOffset(dayStart)
		count := localDailyViews[localSeedIndex(stableOffset, len(localDailyViews))]
		day := dayStart.Format(time.DateOnly)
		for index := 0; index < count; index++ {
			occurredAt := dayStart.Add(time.Duration(index) * time.Minute)
			_, err := module.RecordView(ctx,
				localWeightedDocumentIDs[localSeedIndex(stableOffset+index, len(localWeightedDocumentIDs))],
				ViewInput{
					EventID:    fmt.Sprintf("docs-local-traffic:%s:%02d", day, index),
					OccurredAt: occurredAt, Class: traffic.VisitHuman,
					VisitorSeed: []byte(fmt.Sprintf("docs-local-visitor-%d", index%4)),
					Source:      localSources[localSeedIndex(stableOffset+index, len(localSources))],
				},
			)
			if err != nil {
				return err
			}
		}
		if seeder, ok := module.store.(localSearchSeeder); ok {
			for index, search := range localSearches {
				id := uuid.NewSHA1(uuid.NameSpaceOID, []byte(fmt.Sprintf("docs-local-search:%s:%02d", day, index))).String()
				if err := seeder.SeedAnalyticsSearch(
					ctx, id, search.query, "", "en", search.results,
					dayStart.Add(time.Duration(index)*time.Minute),
				); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
