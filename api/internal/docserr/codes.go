// Package docserr declares Docs' immutable public Problem contract.
package docserr

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/yueli-official/foundation/go/problem"
)

const (
	CodeNotFound                 = "docs.not_found"
	CodeForbidden                = "docs.forbidden"
	CodeSlugTaken                = "docs.slug_taken"
	CodeInvalidInput             = "docs.invalid_input"
	CodeUpstream                 = "docs.upstream_failed"
	CodeImportBlocked            = "docs.import_blocked"
	CodeAuthorizationUnavailable = "docs.authorization_unavailable"
	CodeRateLimited              = "docs.rate_limited"
	CodeChallengeRequired        = "docs.challenge_required"
	CodeAbuseUnavailable         = "docs.abuse_unavailable"
	CodeAbuseReplay              = "docs.abuse_attempt_replayed"
)

var (
	DescriptorRateLimited = descriptor("common.rate_limited", http.StatusTooManyRequests)
	DescriptorValidation  = descriptor("common.validation_failed", http.StatusBadRequest)
	DescriptorInternal    = descriptor("common.internal", http.StatusInternalServerError)

	descriptors = map[string]problem.Descriptor{
		CodeNotFound:                 descriptor(CodeNotFound, http.StatusNotFound),
		CodeForbidden:                descriptor(CodeForbidden, http.StatusForbidden),
		CodeSlugTaken:                descriptor(CodeSlugTaken, http.StatusConflict),
		CodeInvalidInput:             descriptor(CodeInvalidInput, http.StatusBadRequest),
		CodeUpstream:                 descriptor(CodeUpstream, http.StatusBadGateway),
		CodeImportBlocked:            descriptor(CodeImportBlocked, http.StatusBadRequest),
		CodeAuthorizationUnavailable: descriptor(CodeAuthorizationUnavailable, http.StatusServiceUnavailable),
		CodeRateLimited:              descriptor(CodeRateLimited, http.StatusTooManyRequests),
		CodeChallengeRequired:        descriptor(CodeChallengeRequired, http.StatusForbidden),
		CodeAbuseUnavailable:         descriptor(CodeAbuseUnavailable, http.StatusServiceUnavailable),
		CodeAbuseReplay:              descriptor(CodeAbuseReplay, http.StatusConflict),
	}
)

func descriptor(code string, status int) problem.Descriptor {
	return problem.MustDescriptor(
		problem.MustKind(code, status),
		"https://errors.yueli.dev/problems/"+code,
	)
}

func DescriptorForCode(code string) (problem.Descriptor, bool) {
	value, ok := descriptors[code]
	return value, ok
}

type CatalogEntry struct {
	Code   string `json:"code"`
	Status int    `json:"status"`
}

func Catalog() []CatalogEntry {
	result := make([]CatalogEntry, 0, len(descriptors))
	for code, value := range descriptors {
		result = append(result, CatalogEntry{Code: code, Status: value.Kind().Status()})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Code < result[j].Code })
	return result
}

func mapped(code string, params problem.Parameters) error {
	value, ok := DescriptorForCode(code)
	if !ok {
		return fmt.Errorf("docs public error code is not declared: %s", code)
	}
	result, err := problem.NewError(value, params)
	if err != nil {
		return fmt.Errorf("docs public error %s: %w", code, err)
	}
	return result
}

func NotFound(id string) error {
	return mapped(CodeNotFound, map[string]any{"id": id})
}

func Forbidden() error {
	return mapped(CodeForbidden, nil)
}

func AuthorizationUnavailable() error {
	return mapped(CodeAuthorizationUnavailable, nil)
}

func RateLimited() error {
	return mapped(CodeRateLimited, nil)
}

func ChallengeRequired(attemptID string) error {
	return mapped(CodeChallengeRequired, map[string]any{
		"attemptId": attemptID,
		"challenge": "turnstile",
	})
}

func AbuseUnavailable() error {
	return mapped(CodeAbuseUnavailable, nil)
}

func AbuseAttemptReplayed() error {
	return mapped(CodeAbuseReplay, nil)
}

func SlugTaken(slug string) error {
	return mapped(CodeSlugTaken, map[string]any{"slug": slug})
}

func InvalidInput(detail string) error {
	return mapped(CodeInvalidInput, map[string]any{"detail": detail})
}

func UpstreamFailed(detail string) error {
	return mapped(CodeUpstream, map[string]any{"detail": detail})
}

func ImportBlocked(detail string) error {
	return mapped(CodeImportBlocked, map[string]any{"detail": detail})
}
