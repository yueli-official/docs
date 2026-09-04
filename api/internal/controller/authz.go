package controller

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/docserr"
	foundationauth "github.com/yueli-official/foundation/go/auth"
	"github.com/yueli-official/foundation/go/authorization"
)

type authorizationContextKey struct{}

func AuthorizationMiddleware(service *docsauthz.Service) ghttp.HandlerFunc {
	return func(request *ghttp.Request) {
		ctx := context.WithValue(request.Context(), authorizationContextKey{}, service)
		correlationID := strings.TrimSpace(request.Header.Get("X-Trace-Id"))
		if correlationID == "" {
			correlationID = strings.TrimSpace(request.Header.Get("X-Request-Id"))
		}
		ctx = authorization.WithRequestMetadata(ctx, authorization.RequestMetadata{
			CorrelationID: correlationID,
		})
		request.SetCtx(ctx)
		request.Middleware.Next()
	}
}

func authorizationService(ctx context.Context) *docsauthz.Service {
	service, _ := ctx.Value(authorizationContextKey{}).(*docsauthz.Service)
	return service
}

// subject extracts the authenticated subject (JWT group), or a forbidden error.
func subject(ctx context.Context) (string, error) {
	p, ok := foundationauth.FromContext(ctx)
	if !ok || p == nil || strings.TrimSpace(p.Subject) == "" {
		return "", docserr.Forbidden()
	}
	kind, _ := p.Claim("subject_kind")
	if kind != "user" {
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

func idempotencyKeyOf(ctx context.Context) string {
	request := ghttp.RequestFromCtx(ctx)
	if request == nil {
		return ""
	}
	return strings.TrimSpace(request.Request.Header.Get("Idempotency-Key"))
}

func stripBearer(h string) string {
	const p = "bearer "
	if len(h) < len(p) || !strings.EqualFold(h[:len(p)], p) {
		return ""
	}
	return strings.TrimSpace(h[len(p):])
}

func isAdministrator(ctx context.Context) bool {
	return authorizationService(ctx).IsAdministrator(ctx)
}

func requireCapability(
	ctx context.Context,
	capability authorization.CapabilityKey,
	scopeID authorization.ScopeID,
	resource authorization.ResourceFacts,
) error {
	service := authorizationService(ctx)
	if service == nil {
		return docserr.AuthorizationUnavailable()
	}
	decision, err := service.Decide(ctx, capability, scopeID, resource)
	if err != nil {
		if authorization.Is(err, authorization.ErrorUnavailable) {
			return docserr.AuthorizationUnavailable()
		}
		return docserr.Forbidden()
	}
	if !decision.Allowed {
		return docserr.Forbidden()
	}
	return nil
}

func ensureCollectionScope(ctx context.Context, collectionID string) error {
	service := authorizationService(ctx)
	if service == nil {
		return docserr.AuthorizationUnavailable()
	}
	if err := service.EnsureCollectionScope(ctx, collectionID); err != nil {
		return docserr.AuthorizationUnavailable()
	}
	return nil
}

func ensureDocumentScope(ctx context.Context, documentID, collectionID string) error {
	service := authorizationService(ctx)
	if service == nil {
		return docserr.AuthorizationUnavailable()
	}
	if err := service.EnsureDocumentScope(ctx, documentID, collectionID); err != nil {
		return docserr.AuthorizationUnavailable()
	}
	return nil
}

func ensureDocumentHierarchy(ctx context.Context, documentID, collectionID string) error {
	if err := ensureCollectionScope(ctx, collectionID); err != nil {
		return err
	}
	return ensureDocumentScope(ctx, documentID, collectionID)
}

func mapAuthorizationError(err error) error {
	switch {
	case authorization.Is(err, authorization.ErrorDenied):
		return docserr.Forbidden()
	case authorization.Is(err, authorization.ErrorUnavailable):
		return docserr.AuthorizationUnavailable()
	case authorization.Is(err, authorization.ErrorNotFound):
		return docserr.NotFound("authorization")
	case authorization.Is(err, authorization.ErrorInvalidInput),
		authorization.Is(err, authorization.ErrorConflict),
		authorization.Is(err, authorization.ErrorExpired):
		return docserr.InvalidInput("authorization_request_invalid")
	default:
		return docserr.AuthorizationUnavailable()
	}
}
