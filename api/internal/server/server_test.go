// Package server_test contains shared JWT-signing helpers used by all server
// integration tests (TestHealthz, TestDocsRoundTrip, …).
package server_test

import (
	"crypto/rsa"
	"fmt"
	"time"

	jose "github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"

	"platform/gokit/authjwt"
)

const (
	testIssuer = "http://localhost:8081"
	testKID    = "docs-test-kid"
	testSub    = "11111111-1111-1111-1111-111111111111"
)

func prefix(s *ghttp.Server) string {
	return fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
}

func mustVerifier(t *gtest.T, priv *rsa.PrivateKey) *authjwt.Verifier {
	set := jose.JSONWebKeySet{Keys: []jose.JSONWebKey{{
		Key: priv.Public(), KeyID: testKID, Algorithm: "RS256", Use: "sig",
	}}}
	v, err := authjwt.NewVerifier(authjwt.VerifierConfig{
		Keys:   authjwt.NewStaticKeySource(set),
		Issuer: testIssuer,
	})
	t.AssertNil(err)
	return v
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
	}).CompactSerialize()
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
		CompactSerialize()
	t.AssertNil(err)
	return raw
}
