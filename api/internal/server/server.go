// Package server wires the docs-site HTTP routes onto a GoFrame server.
// Shared by cmd/docs and integration tests so they exercise the same wiring.
package server

import (
	"github.com/gogf/gf/v2/net/ghttp"

	foundationauth "github.com/yueli-official/foundation/go/auth"
	"platform/gokit/authhttp"
	"platform/gokit/ghttpx"
	"platform/gokit/healthcheck"
	"platform/products/docs/api/internal/catalog"
	"platform/products/docs/api/internal/controller"
	"platform/products/docs/api/internal/docsauthz"
)

// Deps are the wiring dependencies. Catalog may be nil for a minimal
// health-only server.
type Deps struct {
	Verifier      *foundationauth.Verifier
	Catalog       *catalog.Service
	Authorization *docsauthz.Service
}

// Configure mounts: public health, identity probe, and the catalog API (if Catalog is set).
func Configure(s *ghttp.Server, d Deps) {
	apiMiddleware := ghttpx.NewMiddleware(ghttpx.MustRateLimiterFromEnvironment(), ghttpx.ForwardedClientIPKey)
	s.Use(ghttpx.TraceRouteMiddleware)
	s.Group("/", func(grp *ghttp.RouterGroup) {
		grp.Middleware(apiMiddleware)
		grp.GET("/healthz", controller.Healthz)
		grp.GET("/readyz", healthcheck.Handler(map[string]healthcheck.Check{"database": healthcheck.Database}))
	})

	// Identity probe: JWT parsed-if-present, never 401. Available regardless of
	// whether a Catalog is configured so health + auth probes work standalone.
	s.Group("/", func(grp *ghttp.RouterGroup) {
		grp.Middleware(apiMiddleware, authhttp.Optional(d.Verifier), controller.AuthorizationMiddleware(d.Authorization))
		grp.Bind(controller.NewMe())
	})

	if d.Authorization != nil {
		s.Group("/", func(grp *ghttp.RouterGroup) {
			grp.Middleware(apiMiddleware, authhttp.Required(d.Verifier), controller.AuthorizationMiddleware(d.Authorization))
			grp.Bind(controller.NewAuthorization())
		})
	}

	if d.Catalog == nil {
		return
	}

	// Public browse: enveloped, no mandatory auth.
	s.Group("/", func(grp *ghttp.RouterGroup) {
		grp.Middleware(apiMiddleware)
		grp.Bind(controller.NewPublicCollections(d.Catalog))
	})

	// Admin API: envelope first, then mandatory JWT.
	s.Group("/", func(grp *ghttp.RouterGroup) {
		grp.Middleware(apiMiddleware, authhttp.Required(d.Verifier), controller.AuthorizationMiddleware(d.Authorization))
		grp.Bind(controller.NewCollections(d.Catalog))
		grp.Bind(controller.NewVersions(d.Catalog))
		grp.Bind(controller.NewDocs(d.Catalog))
		grp.Bind(controller.NewImports(d.Catalog))
	})
}
