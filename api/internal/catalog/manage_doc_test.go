package catalog

import (
	"strings"
	"testing"
)

func TestNormalizeManageDocsQueryDefaultsAndTrims(t *testing.T) {
	query, err := normalizeManageDocsQuery(ManageDocsInput{Q: "  guide  "})
	if err != nil {
		t.Fatal(err)
	}
	if query.Q != "guide" || query.Status != "all" || query.Quality != "all" || query.Sort != "updatedAt" || query.Direction != "desc" || query.Page != 1 || query.Size != 30 {
		t.Fatalf("query = %+v", query)
	}
}

func TestNormalizeManageDocsQueryRejectsUnsafeOrAmbiguousValues(t *testing.T) {
	tests := []struct {
		name  string
		input ManageDocsInput
	}{
		{name: "long query", input: ManageDocsInput{Q: strings.Repeat("x", 201)}},
		{name: "status", input: ManageDocsInput{Status: "deleted"}},
		{name: "quality", input: ManageDocsInput{Quality: "broken"}},
		{name: "sort injection", input: ManageDocsInput{Sort: "updated_at desc; drop table docs"}},
		{name: "direction", input: ManageDocsInput{Direction: "sideways"}},
		{name: "negative page", input: ManageDocsInput{Page: -1}},
		{name: "oversized page", input: ManageDocsInput{Size: 101}},
		{name: "version without collection", input: ManageDocsInput{Version: "v2"}},
		{name: "invalid collection", input: ManageDocsInput{CollectionID: "docs"}},
		{name: "invalid exact id", input: ManageDocsInput{ID: "bad-id"}},
		{name: "invalid excluded document", input: ManageDocsInput{ExcludeID: "bad-id"}},
		{name: "path without collection", input: ManageDocsInput{Path: "guide/start"}},
		{name: "invalid parent", input: ManageDocsInput{ParentID: "top"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := normalizeManageDocsQuery(test.input); err == nil {
				t.Fatalf("input accepted: %+v", test.input)
			}
		})
	}
}

func TestNormalizeManageDocsQueryAcceptsAllowlistedFilters(t *testing.T) {
	query, err := normalizeManageDocsQuery(ManageDocsInput{
		Status: "published", Quality: "issues",
		CollectionID: "4f553f75-e2d9-4f21-8d12-5f43659504f2", Version: "v2",
		Locale: "zh-CN", ParentID: "root", Sort: "path", Direction: "asc", Page: 3, Size: 60,
	})
	if err != nil {
		t.Fatal(err)
	}
	if query.ParentID != "root" || query.Sort != "path" || query.Direction != "asc" || query.Page != 3 || query.Size != 60 {
		t.Fatalf("query = %+v", query)
	}
}
