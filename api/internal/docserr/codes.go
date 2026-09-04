// Package docserr declares Docs' immutable public Problem contract.
package docserr

import (
	"fmt"
	"net/http"

	"github.com/yueli-official/foundation/go/problem"
)

type InvalidInputReason string
type ImportBlockedReason string
type Dependency string

var (
	DescriptorRateLimited = descriptor("common.rate_limited", http.StatusTooManyRequests)
	DescriptorValidation  = descriptor("common.validation_failed", http.StatusBadRequest)
	DescriptorInternal    = descriptor("common.internal", http.StatusInternalServerError)
)

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

func AdministratorGrantProtected() error {
	return mapped(CodeAdministratorGrantProtected, nil)
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
	return mapped(CodeAbuseAttemptReplayed, nil)
}

func SlugTaken(slug string) error {
	return mapped(CodeSlugTaken, map[string]any{"slug": slug})
}

func InvalidInput(reason InvalidInputReason) error {
	return mapped(CodeInvalidInput, map[string]any{"reason": string(reason)})
}

func UpstreamFailed(dependency Dependency) error {
	return mapped(CodeUpstreamFailed, map[string]any{"dependency": string(dependency)})
}

func ImportBlocked(reason ImportBlockedReason) error {
	return mapped(CodeImportBlocked, map[string]any{"reason": string(reason)})
}

func ImportCompressionUnsupported(method int) error {
	return mapped(CodeImportCompressionUnsupported, map[string]any{"method": method})
}
