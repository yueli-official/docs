// Package docserr declares the docs-site error codes (namespace docs.*) and
// their HTTP status, registered with the shared gokit/errs catalog.
package docserr

import (
	"net/http"

	"platform/gokit/errs"
)

var (
	CodeNotFound      = errs.Register("docs.not_found", http.StatusNotFound)
	CodeForbidden     = errs.Register("docs.forbidden", http.StatusForbidden)
	CodeSlugTaken     = errs.Register("docs.slug_taken", http.StatusConflict)
	CodeInvalidInput  = errs.Register("docs.invalid_input", http.StatusBadRequest)
	CodeUpstream      = errs.Register("docs.upstream_failed", http.StatusBadGateway)
	CodeImportBlocked = errs.Register("docs.import_blocked", http.StatusBadRequest)
	CodeUnavailable   = errs.Register("docs.authorization_unavailable", http.StatusServiceUnavailable)
)

// NotFound is returned when a collection or doc id/slug does not exist.
func NotFound(id string) *errs.Coded {
	return errs.New(CodeNotFound, "not found", map[string]any{"id": id})
}

// Forbidden is returned when the caller lacks the required permission.
func Forbidden() *errs.Coded { return errs.New(CodeForbidden, "forbidden", nil) }

// AuthorizationUnavailable fails closed without disguising an infrastructure
// outage as a permissions decision.
func AuthorizationUnavailable() *errs.Coded {
	return errs.New(CodeUnavailable, "authorization unavailable", nil)
}

// SlugTaken is returned when a generated slug collides with an existing one.
func SlugTaken(slug string) *errs.Coded {
	return errs.New(CodeSlugTaken, "slug already taken", map[string]any{"slug": slug})
}

// InvalidInput is returned for malformed request input not caught by binding.
func InvalidInput(detail string) *errs.Coded {
	return errs.New(CodeInvalidInput, "invalid input: "+detail, nil)
}

// UpstreamFailed is returned when a dependency, such as the asset service, fails.
func UpstreamFailed(detail string) *errs.Coded {
	return errs.New(CodeUpstream, "upstream failed: "+detail, nil)
}

// ImportBlocked is returned when a preflight report contains blocking issues.
func ImportBlocked(detail string) *errs.Coded {
	return errs.New(CodeImportBlocked, "import blocked: "+detail, nil)
}
