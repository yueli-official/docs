// Package docsurls maps Docs collections and locale/version document
// representations onto the Foundation URL lifecycle module.
package docsurls

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/yueli-official/foundation/go/identifier"
	"github.com/yueli-official/foundation/go/urllifecycle"
)

const (
	CollectionKind urllifecycle.ResourceKind = "docs.collection"
	DocKind        urllifecycle.ResourceKind = "docs.doc"
)

type Lifecycle struct {
	module        urllifecycle.Module
	postgres      *urllifecycle.PostgresAdapter
	defaultLocale string
}

type desiredRoute struct {
	Key      urllifecycle.RouteKey
	Ref      urllifecycle.LocalRef
	Public   bool
	Resource string
}

func Definition(origin, defaultLocale string) urllifecycle.Definition {
	return urllifecycle.Definition{
		Version:       urllifecycle.DefinitionVersion,
		TrustedOrigin: strings.TrimRight(origin, "/"),
		ResourceKinds: []urllifecycle.ResourceKindDefinition{
			{Key: CollectionKind},
			{Key: DocKind},
		},
		Namespaces: []urllifecycle.NamespaceDefinition{{
			Key: "docs.public", PathPrefix: "/",
			IdentityQuery: urllifecycle.QueryIdentityDefinition{
				Keys: []urllifecycle.QueryKeyDefinition{
					{Key: "locale", Default: defaultLocale, OmitDefault: true},
					{Key: "version", Default: "default", OmitDefault: true},
				},
			},
		}},
		Limits: urllifecycle.Limits{MaxPageSize: 1000},
	}
}

func NewPostgres(ctx context.Context, db *sql.DB, instanceKey, origin, defaultLocale string) (*Lifecycle, error) {
	catalog, err := urllifecycle.Compile(Definition(origin, defaultLocale))
	if err != nil {
		return nil, err
	}
	adapter, err := urllifecycle.NewPostgres(ctx, catalog, urllifecycle.PostgresOptions{
		DB: db, InstanceKey: instanceKey,
	})
	if err != nil {
		return nil, err
	}
	return &Lifecycle{module: adapter, postgres: adapter, defaultLocale: defaultLocale}, nil
}

func NewMemory(origin, defaultLocale string) (*Lifecycle, error) {
	catalog, err := urllifecycle.Compile(Definition(origin, defaultLocale))
	if err != nil {
		return nil, err
	}
	module, err := urllifecycle.NewMemory(catalog, urllifecycle.MemoryOptions{})
	if err != nil {
		return nil, err
	}
	return &Lifecycle{module: module, defaultLocale: defaultLocale}, nil
}

func (l *Lifecycle) Resolver() urllifecycle.Resolver {
	if l == nil {
		return nil
	}
	return l.module
}

// ReconcileCollection reads product state through tx and applies one final-state
// transition for the collection route and every document representation in it.
func (l *Lifecycle) ReconcileCollection(ctx context.Context, tx *sql.Tx, collectionID, reason string) error {
	module, err := l.bound(tx)
	if err != nil {
		return err
	}
	desired, err := l.loadDesired(ctx, tx, collectionID)
	if err != nil {
		return err
	}
	active, err := listActive(ctx, module, collectionID)
	if err != nil {
		return err
	}
	change := planReconciliation(desired, active, reason)
	if len(change.ResourceChanges) == 0 {
		return nil
	}
	_, err = module.Apply(ctx, change, urllifecycle.ApplyOptions{})
	return err
}

// ReconcileAll is the startup backfill. It uses one transaction per collection
// so a large site is bounded by collection size rather than global page count.
func (l *Lifecycle) ReconcileAll(ctx context.Context, db *sql.DB) error {
	rows, err := db.QueryContext(ctx, `SELECT id::text FROM collections ORDER BY id`)
	if err != nil {
		return err
	}
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			_ = rows.Close()
			return err
		}
		ids = append(ids, id)
	}
	if err := rows.Close(); err != nil {
		return err
	}
	for _, id := range ids {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		if err := l.ReconcileCollection(ctx, tx, id, "docs startup reconciliation"); err != nil {
			_ = tx.Rollback()
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
	}
	return nil
}

func (l *Lifecycle) loadDesired(ctx context.Context, tx *sql.Tx, collectionID string) ([]desiredRoute, error) {
	var collectionSlug, collectionDefaultLocale string
	err := tx.QueryRowContext(ctx, `
SELECT collection.slug,
       COALESCE(locale.locale, $2)
FROM collections collection
LEFT JOIN collection_locales locale
  ON locale.collection_id = collection.id
 AND locale.is_default = true
 AND locale.enabled = true
WHERE collection.id = $1::uuid`, collectionID, l.defaultLocale).Scan(&collectionSlug, &collectionDefaultLocale)
	if errors.Is(err, sql.ErrNoRows) {
		return []desiredRoute{}, nil
	}
	if err != nil {
		return nil, err
	}
	desired := []desiredRoute{{
		Key: urllifecycle.RouteKey{
			Resource: urllifecycle.ResourceKey{Kind: CollectionKind, ID: collectionID},
		},
		Ref:      urllifecycle.LocalRef{Path: "/" + collectionSlug},
		Public:   true,
		Resource: collectionID,
	}}
	rows, err := tx.QueryContext(ctx, `
WITH RECURSIVE paths AS (
    SELECT doc.id, doc.parent_id, doc.version_id, doc.locale, doc.slug,
           doc.slug::text AS slug_path,
           (doc.status = 'published') AS public_chain
    FROM docs doc
    WHERE doc.collection_id = $1::uuid
      AND doc.parent_id IS NULL
      AND doc.deleted_at IS NULL
    UNION ALL
    SELECT child.id, child.parent_id, child.version_id, child.locale, child.slug,
           (parent.slug_path || '/' || child.slug)::text,
           (parent.public_chain AND child.status = 'published')
    FROM docs child
    JOIN paths parent ON parent.id = child.parent_id
    WHERE child.collection_id = $1::uuid
      AND child.deleted_at IS NULL
)
SELECT path.id::text, path.version_id::text, path.locale, path.slug_path,
       version.key, version.is_default,
       (path.public_chain AND version.status IN ('published', 'archived')) AS public
FROM paths path
JOIN collection_versions version ON version.id = path.version_id
ORDER BY path.id`, collectionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, versionID, locale, path, versionKey string
		var isDefault, public bool
		if err := rows.Scan(&id, &versionID, &locale, &path, &versionKey, &isDefault, &public); err != nil {
			return nil, err
		}
		resourceID := collectionID + "/" + id
		query := make([]urllifecycle.QueryValue, 0, 2)
		if locale != collectionDefaultLocale {
			query = append(query, urllifecycle.QueryValue{Key: "locale", Value: locale})
		}
		if !isDefault {
			query = append(query, urllifecycle.QueryValue{Key: "version", Value: versionKey})
		}
		desired = append(desired, desiredRoute{
			Key: urllifecycle.RouteKey{
				Resource: urllifecycle.ResourceKey{Kind: DocKind, ID: resourceID},
				Variant:  versionID + ":" + locale,
			},
			Ref: urllifecycle.LocalRef{
				Path: "/" + collectionSlug + "/" + path, Query: query,
			},
			Public:   public,
			Resource: resourceID,
		})
	}
	return desired, rows.Err()
}

func (l *Lifecycle) bound(tx *sql.Tx) (urllifecycle.Module, error) {
	if l == nil || l.module == nil {
		return nil, fmt.Errorf("docs URL lifecycle is not configured")
	}
	if l.postgres == nil {
		return l.module, nil
	}
	if tx == nil {
		return nil, fmt.Errorf("docs URL lifecycle requires the product transaction")
	}
	return l.postgres.Bind(tx)
}

func listActive(ctx context.Context, module urllifecycle.Module, collectionID string) (map[string]urllifecycle.Inspection, error) {
	result := map[string]urllifecycle.Inspection{}
	after := ""
	for {
		page, err := module.List(ctx, urllifecycle.ListQuery{After: after, Limit: 1000})
		if err != nil {
			return nil, err
		}
		for _, item := range page.Items {
			if item.Route == nil || item.Active == nil || !inCollection(*item.Route, collectionID) {
				continue
			}
			result[routeID(*item.Route)] = item
		}
		if page.Next == "" {
			break
		}
		after = page.Next
	}
	return result, nil
}

func planReconciliation(
	desired []desiredRoute,
	active map[string]urllifecycle.Inspection,
	reason string,
) urllifecycle.ChangeSet {
	desiredByKey := make(map[string]desiredRoute, len(desired))
	desiredByResource := make(map[string]desiredRoute, len(desired))
	for _, item := range desired {
		desiredByKey[routeID(item.Key)] = item
		desiredByResource[item.Resource] = item
	}
	consumed := map[string]bool{}
	changes := make([]urllifecycle.ResourceChange, 0)
	for _, item := range desired {
		exact, found := active[routeID(item.Key)]
		if found {
			consumed[routeID(item.Key)] = true
			if exact.Active.Canonical.Path != item.Ref.Path || !sameQuery(exact.Active.Canonical.Query, item.Ref.Query) {
				next := *exact.Active
				next.Canonical = item.Ref
				changes = append(changes, urllifecycle.ResourceChange{
					Route: item.Key, ExpectedRevision: exact.Revision, Desired: &next,
					Departures: redirectDeparture(),
				})
			}
			continue
		}
		var formerKey string
		var former urllifecycle.Inspection
		for key, candidate := range active {
			if candidate.Route != nil && candidate.Route.Resource.ID == item.Resource && !consumed[key] {
				formerKey, former = key, candidate
				break
			}
		}
		if former.Route != nil {
			consumed[formerKey] = true
			targetActive := urllifecycle.ActiveRoute{Canonical: item.Ref}
			changes = append(changes,
				urllifecycle.ResourceChange{
					Route: item.Key, Desired: &targetActive,
					Departures: releaseDeparture(),
				},
				urllifecycle.ResourceChange{
					Route: *former.Route, ExpectedRevision: former.Revision,
					Departures: urllifecycle.DeparturePolicy{
						Canonical: urllifecycle.FormerOutcome{
							Kind: urllifecycle.FormerRedirectToRoute, Target: item.Key,
							Redirect: urllifecycle.DefaultPermanentRedirect(),
						},
						Aliases: urllifecycle.FormerOutcome{
							Kind: urllifecycle.FormerRedirectToRoute, Target: item.Key,
							Redirect: urllifecycle.DefaultPermanentRedirect(),
						},
					},
				},
			)
			continue
		}
		if item.Public {
			next := urllifecycle.ActiveRoute{Canonical: item.Ref}
			changes = append(changes, urllifecycle.ResourceChange{
				Route: item.Key, Desired: &next, Departures: releaseDeparture(),
			})
		}
	}
	for key, item := range active {
		if consumed[key] || item.Route == nil {
			continue
		}
		if _, stillExists := desiredByKey[key]; stillExists {
			continue
		}
		if replacement, moved := desiredByResource[item.Route.Resource.ID]; moved {
			_ = replacement
			continue
		}
		changes = append(changes, urllifecycle.ResourceChange{
			Route: *item.Route, ExpectedRevision: item.Revision,
			Departures: urllifecycle.DeparturePolicy{
				Canonical: urllifecycle.FormerOutcome{Kind: urllifecycle.FormerGone},
				Aliases:   urllifecycle.FormerOutcome{Kind: urllifecycle.FormerGone},
			},
		})
	}
	return urllifecycle.ChangeSet{
		CommandID:       urllifecycle.CommandID(identifier.MustNew().String()),
		Actor:           urllifecycle.ActorRef{Kind: "system", ID: "docs"},
		Reason:          reason,
		ResourceChanges: changes,
	}
}

func inCollection(route urllifecycle.RouteKey, collectionID string) bool {
	if route.Resource.Kind == CollectionKind {
		return route.Resource.ID == collectionID
	}
	return route.Resource.Kind == DocKind && strings.HasPrefix(route.Resource.ID, collectionID+"/")
}

func routeID(route urllifecycle.RouteKey) string {
	return string(route.Resource.Kind) + "\x00" + route.Resource.ID + "\x00" + route.Variant
}

func sameQuery(left, right []urllifecycle.QueryValue) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}

func redirectDeparture() urllifecycle.DeparturePolicy {
	return urllifecycle.DeparturePolicy{
		Canonical: urllifecycle.FormerOutcome{
			Kind: urllifecycle.FormerRedirectToCurrent, Redirect: urllifecycle.DefaultPermanentRedirect(),
		},
		Aliases: urllifecycle.FormerOutcome{Kind: urllifecycle.FormerRelease},
	}
}

func releaseDeparture() urllifecycle.DeparturePolicy {
	return urllifecycle.DeparturePolicy{
		Canonical: urllifecycle.FormerOutcome{Kind: urllifecycle.FormerRelease},
		Aliases:   urllifecycle.FormerOutcome{Kind: urllifecycle.FormerRelease},
	}
}
