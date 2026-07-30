// Package server_test contains shared JWT-signing helpers used by all server
// integration tests (TestHealthz, TestDocsRoundTrip, …).
package server_test

import (
	"crypto/rsa"
	"fmt"
	"time"

	jose "github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"

	foundationauth "github.com/yueli-official/foundation/go/auth"
	"github.com/yueli-official/foundation/go/authorization"
	"platform/gokit/authsetup"
	"platform/products/docs/api/internal/docsauthz"
)

const (
	testIssuer = "http://localhost:8081"
	testKID    = "docs-test-kid"
	testSub    = "11111111-1111-1111-1111-111111111111"
)

func prefix(s *ghttp.Server) string {
	return fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
}

func mustVerifier(t *gtest.T, priv *rsa.PrivateKey) *foundationauth.Verifier {
	set := jose.JSONWebKeySet{Keys: []jose.JSONWebKey{{
		Key: priv.Public(), KeyID: testKID, Algorithm: "RS256", Use: "sig",
	}}}
	v, err := authsetup.NewStaticVerifier(authsetup.StaticVerifierConfig{
		Keys:   set,
		Issuer: testIssuer,
	})
	t.AssertNil(err)
	return v
}

func mustAuthorization(t *gtest.T, administrators ...string) *docsauthz.Service {
	t.Helper()
	subjects := make([]authorization.SubjectRef, 0, len(administrators))
	for _, subject := range administrators {
		subjects = append(subjects, authorization.SubjectRef{Kind: authorization.SubjectUser, ID: subject})
	}
	module, err := authorization.NewMemory(authorization.MustCompile(docsauthz.Definition()), authorization.MemoryOptions{
		RootScopeID: docsauthz.RootScopeID, ProtectedSubjects: subjects,
		Constraints: docsauthz.ConstraintEvaluators(),
		Predicates:  docsauthz.PredicateEvaluators(),
	})
	t.AssertNil(err)
	return docsauthz.New(module)
}

func signToken(t *gtest.T, priv *rsa.PrivateKey, sub string, exp time.Time) string {
	signer, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.RS256, Key: priv},
		(&jose.SignerOptions{}).WithType("JWT").WithHeader("kid", testKID),
	)
	t.AssertNil(err)
	now := time.Now().UTC()
	raw, err := jwt.Signed(signer).Claims(jwt.Claims{
		Issuer:   testIssuer,
		Subject:  sub,
		IssuedAt: jwt.NewNumericDate(now.Add(-time.Minute)),
		Expiry:   jwt.NewNumericDate(exp),
	}).Serialize()
	t.AssertNil(err)
	return raw
}

func signTokenRoles(t *gtest.T, priv *rsa.PrivateKey, sub string, roles []string, exp time.Time) string {
	t.Helper()
	signer, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.RS256, Key: priv},
		(&jose.SignerOptions{}).WithType("JWT").WithHeader("kid", testKID),
	)
	t.AssertNil(err)
	now := time.Now().UTC()
	raw, err := jwt.Signed(signer).
		Claims(jwt.Claims{Issuer: testIssuer, Subject: sub, IssuedAt: jwt.NewNumericDate(now.Add(-time.Minute)), Expiry: jwt.NewNumericDate(exp)}).
		Claims(map[string]any{"scope": "openid profile roles", "roles": roles}).
		Serialize()
	t.AssertNil(err)
	return raw
}
