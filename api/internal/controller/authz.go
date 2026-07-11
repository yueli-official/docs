package controller

import (
	"context"
	"slices"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"platform/gokit/authjwt"
	"platform/products/docs/api/internal/docserr"
)

// subject extracts the authenticated subject (JWT group), or a forbidden error.
func subject(ctx context.Context) (string, error) {
	p, ok := authjwt.From(ctx)
	if !ok {
		return "", docserr.Forbidden()
	}
	return p.Subject, nil
}

func bearerOf(ctx context.Context) string {
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		return ""
	}
	return stripBearer(r.Request.Header.Get("Authorization"))
}

func stripBearer(h string) string {
	const p = "bearer "
	if len(h) < len(p) || !strings.EqualFold(h[:len(p)], p) {
		return ""
	}
	return strings.TrimSpace(h[len(p):])
}

// isAdmin reports whether the caller is a docs site owner — the site's top
// role, held by whoever the catalog lists in docs.operatorSubs (their identity
// sub). Per-site authz: deliberately does NOT read any IdP global role.
func isAdmin(ctx context.Context) bool {
	p, ok := authjwt.From(ctx)
	if !ok {
		return false
	}
	return slices.Contains(g.Cfg().MustGet(ctx, "docs.operatorSubs").Strings(), p.Subject)
}
