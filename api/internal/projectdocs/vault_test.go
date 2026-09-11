package projectdocs

import (
	"encoding/base64"
	"strings"
	"testing"
)

func TestVaultBindsCredentialToSource(t *testing.T) {
	v, err := NewVault(base64.StdEncoding.EncodeToString([]byte(strings.Repeat("x", 32))))
	if err != nil {
		t.Fatal(err)
	}
	a, err := v.Seal("source-a", "pat_secret")
	if err != nil {
		t.Fatal(err)
	}
	b, _ := v.Seal("source-a", "pat_secret")
	if a == b || strings.Contains(a, "pat_secret") {
		t.Fatal("credential not sealed with fresh nonce")
	}
	token, err := v.Open("source-a", a)
	if err != nil || token != "pat_secret" {
		t.Fatal(err)
	}
	if _, err = v.Open("source-b", a); err == nil {
		t.Fatal("cross-source credential accepted")
	}
	if _, err = NewVault("short-key"); err == nil {
		t.Fatal("weak key accepted")
	}
}
