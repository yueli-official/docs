// Package appconfig builds runtime objects from the GoFrame config
// (manifest/config/config.yaml + GF_* env overrides).
package appconfig

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"platform/products/docs/api/internal/assetclient"
)

// JWKS is the IdP key/issuer config for the Foundation auth verifier.
type JWKS struct {
	URL      string
	Issuer   string
	Audience string
}

// LoadJWKS reads docs.jwks.{url,issuer,audience} from the GoFrame config.
func LoadJWKS(ctx context.Context) JWKS {
	return JWKS{
		URL:      g.Cfg().MustGet(ctx, "docs.jwks.url").String(),
		Issuer:   g.Cfg().MustGet(ctx, "docs.jwks.issuer").String(),
		Audience: g.Cfg().MustGet(ctx, "docs.jwks.audience").String(),
	}
}

func BuildAssetClient(ctx context.Context) assetclient.Client {
	base := g.Cfg().MustGet(ctx, "docs.assetService.baseUrl", "http://localhost:8082").String()
	return assetclient.NewHTTP(base, SiteSlug(ctx), AssetSpace(ctx))
}

func SiteSlug(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "docs.siteSlug", "docs").String()
}

func AssetSpace(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "docs.assetSpace", "default").String()
}

func CoverCategory(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "docs.coverCategory", "docs-collection-cover").String()
}
