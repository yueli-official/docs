// Command docs is the docs-site backend: documentation hub.
package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"

	"platform/gokit/authjwt"
	"platform/gokit/openapiexport"
	"platform/products/docs/api/internal/appconfig"
	"platform/products/docs/api/internal/catalog"
	"platform/products/docs/api/internal/dao"
	"platform/products/docs/api/internal/server"
)

func main() {
	ctx := gctx.New()

	// ── Catalog logic (DB access) ─────────────────────────────────────────────
	cat := catalog.New(dao.NewPG(g.DB())).
		WithAssets(appconfig.BuildAssetClient(ctx), appconfig.CoverCategory(ctx))

	// ── JWT verifier (IdP JWKS, lazy) ────────────────────────────────────────
	jw := appconfig.LoadJWKS(ctx)
	verifier, err := authjwt.NewVerifier(authjwt.VerifierConfig{
		Keys: authjwt.NewRemoteKeySource(jw.URL), Issuer: jw.Issuer, Audience: jw.Audience,
	})
	if err != nil {
		panic(err)
	}

	s := g.Server()
	server.Configure(s, server.Deps{Verifier: verifier, Catalog: cat})
	if handled, err := openapiexport.ExportIfRequested(s); handled {
		if err != nil {
			panic(err)
		}
		return
	}
	g.Log().Info(ctx, "docs-service starting")
	s.Run()
}
