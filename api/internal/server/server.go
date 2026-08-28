// Package server wires the docs-site HTTP routes onto a GoFrame server.
// Shared by cmd/docs and integration tests so they exercise the same wiring.
package server

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/yueli-official/docs/api/internal/catalog"
	"github.com/yueli-official/docs/api/internal/controller"
	"github.com/yueli-official/docs/api/internal/docsanalytics"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/docscomments"
	"github.com/yueli-official/docs/api/internal/identityclient"
	"github.com/yueli-official/docs/api/internal/runtime"
	foundationauth "github.com/yueli-official/foundation/go/auth"
	"github.com/yueli-official/foundation/go/discovery"
	"github.com/yueli-official/foundation/go/urllifecycle"
)

// Deps are the wiring dependencies. Catalog may be nil for a minimal
// health-only server.
type Deps struct {
	Verifier       *foundationauth.Verifier
	Catalog        *catalog.Service
	Authorization  *docsauthz.Service
	Discovery      *discovery.Module
	DiscoveryCache *discovery.Cache
	URLResolver    urllifecycle.Resolver
	Analytics      *docsanalytics.Module
	Comments       *docscomments.Module
	Identity       identityclient.Client
}

// Configure mounts: public health, identity probe, and the catalog API (if Catalog is set).
func Configure(s *ghttp.Server, d Deps) {
	apiMiddleware := runtime.MustAPIMiddleware(runtime.MustRateLimiterFromEnvironment())
	s.Use(runtime.TraceRouteMiddleware)
	s.Group("/", func(grp *ghttp.RouterGroup) {
		grp.Middleware(apiMiddleware.Handle)
		grp.GET("/healthz", controller.Healthz)
		grp.GET("/readyz", runtime.ReadinessHandler(map[string]runtime.ReadinessCheck{"database": runtime.DatabaseReadiness}))
	})

	// Identity probe: JWT parsed-if-present, never 401. Available regardless of
	// whether a Catalog is configured so health + auth probes work standalone.
	s.Group("/", func(grp *ghttp.RouterGroup) {
		if d.Verifier != nil {
			grp.Middleware(apiMiddleware.Handle, runtime.OptionalAuth(d.Verifier), controller.AuthorizationMiddleware(d.Authorization))
		} else {
			grp.Middleware(apiMiddleware.Handle, controller.AuthorizationMiddleware(d.Authorization))
		}
		grp.Bind(controller.NewMe())
	})

	if d.Authorization != nil {
		s.Group("/", func(grp *ghttp.RouterGroup) {
			if d.Verifier != nil {
				grp.Middleware(apiMiddleware.Handle, runtime.RequiredAuth(d.Verifier), controller.AuthorizationMiddleware(d.Authorization))
			} else {
				// OpenAPI export has no runtime verifier, but protected route shapes
				// still belong in the generated contract.
				grp.Middleware(apiMiddleware.Handle, controller.AuthorizationMiddleware(d.Authorization))
			}
			grp.Bind(controller.NewAuthorization())
		})
	}

	if d.Comments != nil {
		s.Group("/", func(grp *ghttp.RouterGroup) {
			grp.Middleware(apiMiddleware.Handle)
			grp.Bind(controller.NewPublicComments(d.Comments, d.Identity))
		})
		s.Group("/", func(grp *ghttp.RouterGroup) {
			if d.Verifier != nil {
				grp.Middleware(apiMiddleware.Handle, runtime.RequiredAuth(d.Verifier), controller.AuthorizationMiddleware(d.Authorization))
			} else {
				grp.Middleware(apiMiddleware.Handle, controller.AuthorizationMiddleware(d.Authorization))
			}
			grp.Bind(controller.NewComments(d.Comments, d.Identity))
		})
	}

	if d.Catalog == nil {
		return
	}

	// Public browse: enveloped, no mandatory auth.
	s.Group("/", func(grp *ghttp.RouterGroup) {
		grp.Middleware(apiMiddleware.Handle)
		grp.Bind(controller.NewPublicCollections(d.Catalog, d.Discovery))
		if d.Analytics != nil {
			grp.Bind(controller.NewPublicAnalytics(d.Analytics, d.Verifier))
		}
		if d.DiscoveryCache != nil {
			grp.Bind(controller.NewPublicDiscovery(d.DiscoveryCache))
		}
		if d.URLResolver != nil {
			grp.Bind(controller.NewPublicURLLifecycle(d.URLResolver))
		}
	})

	// Admin API: envelope first, then mandatory JWT.
	s.Group("/", func(grp *ghttp.RouterGroup) {
		if d.Verifier != nil {
			grp.Middleware(apiMiddleware.Handle, runtime.RequiredAuth(d.Verifier), controller.AuthorizationMiddleware(d.Authorization))
		} else {
			// OpenAPI export has no runtime verifier, but protected route shapes
			// still belong in the generated contract.
			grp.Middleware(apiMiddleware.Handle, controller.AuthorizationMiddleware(d.Authorization))
		}
		grp.Bind(controller.NewCollections(d.Catalog))
		grp.Bind(controller.NewVersions(d.Catalog))
		grp.Bind(controller.NewLocales(d.Catalog))
		grp.Bind(controller.NewDocs(d.Catalog))
		grp.Bind(controller.NewImports(d.Catalog))
		if d.Analytics != nil {
			grp.Bind(controller.NewDashboard(d.Analytics))
		}
	})
}
