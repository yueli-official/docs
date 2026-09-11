package projectdocs

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

type Vault struct{ aead cipher.AEAD }

// The configured key is independent of browser cookies and database credentials.
func NewVault(key string) (*Vault, error) {
	raw, err := base64.StdEncoding.DecodeString(key)
	if err != nil || len(raw) != 32 {
		return nil, errors.New("project documentation sync requires a base64-encoded 32-byte key")
	}
	block, err := aes.NewCipher(raw)
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return &Vault{aead: aead}, nil
}
func (v *Vault) Seal(sourceID, token string) (string, error) {
	nonce := make([]byte, v.aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}
	sealed := v.aead.Seal(nonce, nonce, []byte(token), []byte("docs-project-source/v1:"+sourceID))
	return base64.RawURLEncoding.EncodeToString(sealed), nil
}
func (v *Vault) Open(sourceID, sealed string) (string, error) {
	raw, err := base64.RawURLEncoding.DecodeString(sealed)
	if err != nil || len(raw) < v.aead.NonceSize()+v.aead.Overhead() {
		return "", errors.New("invalid sync credential")
	}
	n := v.aead.NonceSize()
	plain, err := v.aead.Open(nil, raw[:n], raw[n:], []byte("docs-project-source/v1:"+sourceID))
	if err != nil {
		return "", errors.New("invalid sync credential")
	}
	return string(plain), nil
}
