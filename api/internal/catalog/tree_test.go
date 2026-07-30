package catalog

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"

	"platform/products/docs/api/internal/model"
)

func TestBuildTree(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		flat := []*model.Doc{
			{ID: "a", Title: "A"},                // root
			{ID: "b", ParentID: "a", Title: "B"}, // a > b
			{ID: "c", ParentID: "b", Title: "C"}, // a > b > c (3rd level)
			{ID: "d", Title: "D"},                // root
		}
		roots := BuildTree(flat)
		t.Assert(len(roots), 2) // a, d
		t.Assert(roots[0].ID, "a")
		t.Assert(len(roots[0].Children), 1)                // b
		t.Assert(roots[0].Children[0].Children[0].ID, "c") // depth ≥3 preserved
	})
}
