// Package appconfig builds runtime objects from the GoFrame config
// (manifest/config/config.yaml + GF_* env overrides).
package appconfig

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/url"

	"github.com/gogf/gf/v2/frame/g"
	_ "github.com/lib/pq"

	"platform/products/docs/api/internal/assetclient"
)

// JWKS is the IdP key/issuer config for the Foundation auth verifier.
type JWKS struct {
	URL               string
	Issuer            string
	Audience          string
	AllowLoopbackHTTP bool
}

// OpenAuthorizationDB opens the standard-library PostgreSQL handle required by
// the Foundation Authorization Adapter. It points at the same consumer-owned
// database as GoFrame; authorization does not use a central service or store.
func OpenAuthorizationDB(ctx context.Context) (*sql.DB, error) {
	host := g.Cfg().MustGet(ctx, "database.default.host").String()
	port := g.Cfg().MustGet(ctx, "database.default.port", "5432").String()
	name := g.Cfg().MustGet(ctx, "database.default.name").String()
	user := g.Cfg().MustGet(ctx, "database.default.user").String()
	password := g.Cfg().MustGet(ctx, "database.default.pass").String()
	if host == "" || name == "" || user == "" {
		return nil, fmt.Errorf("database.default host, name, and user are required")
	}
	dsn := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, password),
		Host:   net.JoinHostPort(host, port),
		Path:   name,
	}
	query := dsn.Query()
	query.Set("sslmode", g.Cfg().MustGet(ctx, "database.default.sslmode", "disable").String())
	dsn.RawQuery = query.Encode()
	db, err := sql.Open("postgres", dsn.String())
	if err != nil {
		return nil, fmt.Errorf("open authorization database: %w", err)
	}
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping authorization database: %w", err)
	}
	return db, nil
}

// BootstrapAdministratorSubs are identity subjects used only when the local
// authorization instance is first created. Later starts load grants from the
// consumer database and do not reinterpret this configuration as authority.
func BootstrapAdministratorSubs(ctx context.Context) []string {
	return g.Cfg().MustGet(ctx, "docs.authorization.bootstrapAdministratorSubs").Strings()
}

// LoadJWKS reads docs.jwks.{url,issuer,audience} from the GoFrame config.
func LoadJWKS(ctx context.Context) JWKS {
	return JWKS{
		URL:               g.Cfg().MustGet(ctx, "docs.jwks.url").String(),
		Issuer:            g.Cfg().MustGet(ctx, "docs.jwks.issuer").String(),
		Audience:          g.Cfg().MustGet(ctx, "docs.jwks.audience").String(),
		AllowLoopbackHTTP: g.Cfg().MustGet(ctx, "docs.jwks.allowLoopbackHttp", false).Bool(),
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
