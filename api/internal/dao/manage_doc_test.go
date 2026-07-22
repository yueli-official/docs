package dao

import (
	"strings"
	"testing"

	"platform/products/docs/api/internal/model"
)

func TestManageDocsOrderUsesAllowlistedColumns(t *testing.T) {
	tests := []struct {
		sort, direction, want string
	}{
		{"updatedAt", "desc", "d.updated_at DESC, d.id ASC"},
		{"title", "asc", "LOWER(d.title) ASC, d.id ASC"},
		{"path", "asc", "LOWER(COALESCE(paths.slug_path, d.slug)) ASC, d.id ASC"},
		{"sortOrder", "desc", "d.sort_order DESC, d.id ASC"},
	}
	for _, test := range tests {
		got, err := manageDocsOrder(model.ManageDocsQuery{Sort: test.sort, Direction: test.direction})
		if err != nil || got != test.want {
			t.Fatalf("%s/%s => %q, %v", test.sort, test.direction, got, err)
		}
	}
	for _, query := range []model.ManageDocsQuery{
		{Sort: "updated_at; DROP TABLE docs", Direction: "asc"},
		{Sort: "updatedAt", Direction: "desc nulls last"},
	} {
		if _, err := manageDocsOrder(query); err == nil {
			t.Fatalf("unsafe order accepted: %+v", query)
		}
	}
}

func TestManageDocsConditionsSeparatesLifecycleFromStableCounts(t *testing.T) {
	query := model.ManageDocsQuery{
		Q: "guide", Status: "published", Quality: "issues",
		CollectionID: "collection-id", Version: "v2", Locale: "zh-CN", ParentID: "root",
	}
	filtered, args := manageDocsConditions(query, true)
	for _, fragment := range []string{"d.collection_id = ?", "v.key = ?", "d.locale = ?", "d.parent_id IS NULL", "d.status = ?", "COALESCE(d.excerpt", "paths.slug_path"} {
		if !strings.Contains(filtered, fragment) {
			t.Fatalf("filtered conditions missing %q: %s", fragment, filtered)
		}
	}
	if len(args) != 9 {
		t.Fatalf("filtered args = %#v", args)
	}
	stable, stableArgs := manageDocsConditions(query, false)
	if strings.Contains(stable, "d.status = ?") || strings.Contains(stable, "COALESCE(d.excerpt") {
		t.Fatalf("stable count conditions include lifecycle filters: %s", stable)
	}
	if len(stableArgs) != 8 {
		t.Fatalf("stable args = %#v", stableArgs)
	}
}
