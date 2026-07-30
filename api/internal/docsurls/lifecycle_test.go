package docsurls

import (
	"testing"

	"github.com/yueli-official/foundation/go/urllifecycle"
)

func TestPlanReconciliationMovesSubtreeAtomically(t *testing.T) {
	collectionID := "collection-1"
	parentKey := urllifecycle.RouteKey{
		Resource: urllifecycle.ResourceKey{Kind: DocKind, ID: collectionID + "/parent"},
		Variant:  "version-1:en",
	}
	childKey := urllifecycle.RouteKey{
		Resource: urllifecycle.ResourceKey{Kind: DocKind, ID: collectionID + "/child"},
		Variant:  "version-1:en",
	}
	active := map[string]urllifecycle.Inspection{
		routeID(parentKey): {
			Route: &parentKey, Revision: 1,
			Active: &urllifecycle.ActiveRoute{Canonical: urllifecycle.LocalRef{Path: "/old/guide"}},
		},
		routeID(childKey): {
			Route: &childKey, Revision: 2,
			Active: &urllifecycle.ActiveRoute{Canonical: urllifecycle.LocalRef{Path: "/old/guide/install"}},
		},
	}
	change := planReconciliation([]desiredRoute{
		{Key: parentKey, Resource: parentKey.Resource.ID, Ref: urllifecycle.LocalRef{Path: "/new/manual"}, Public: true},
		{Key: childKey, Resource: childKey.Resource.ID, Ref: urllifecycle.LocalRef{Path: "/new/manual/install"}, Public: true},
	}, active, "move")
	if len(change.ResourceChanges) != 2 {
		t.Fatalf("changes = %d, want 2", len(change.ResourceChanges))
	}
	for _, item := range change.ResourceChanges {
		if item.Departures.Canonical.Kind != urllifecycle.FormerRedirectToCurrent {
			t.Fatalf("departure = %q", item.Departures.Canonical.Kind)
		}
	}
}

func TestPlanReconciliationMovesVariantToStableTarget(t *testing.T) {
	resource := "collection-1/doc-1"
	oldKey := urllifecycle.RouteKey{
		Resource: urllifecycle.ResourceKey{Kind: DocKind, ID: resource},
		Variant:  "v1:en",
	}
	newKey := urllifecycle.RouteKey{
		Resource: oldKey.Resource,
		Variant:  "v2:zh-CN",
	}
	active := map[string]urllifecycle.Inspection{
		routeID(oldKey): {
			Route: &oldKey, Revision: 3,
			Active: &urllifecycle.ActiveRoute{Canonical: urllifecycle.LocalRef{Path: "/docs/guide"}},
		},
	}
	change := planReconciliation([]desiredRoute{{
		Key: newKey, Resource: resource, Public: true,
		Ref: urllifecycle.LocalRef{
			Path: "/docs/guide",
			Query: []urllifecycle.QueryValue{
				{Key: "locale", Value: "zh-CN"},
				{Key: "version", Value: "v2"},
			},
		},
	}}, active, "variant move")
	if len(change.ResourceChanges) != 2 {
		t.Fatalf("changes = %d, want claim + redirect", len(change.ResourceChanges))
	}
	if change.ResourceChanges[1].Departures.Canonical.Target != newKey {
		t.Fatal("former variant does not target the new stable route")
	}
}
