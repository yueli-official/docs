// Package identityclient resolves the Identity-owned public display fields
// needed by Docs without importing Identity service source.
package identityclient

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	foundationhttpclient "github.com/yueli-official/foundation/go/httpclient"
)

type MediaRef struct {
	MediaKey string `json:"mediaKey"`
}

type PublicUser struct {
	UserKey     string    `json:"userKey"`
	DisplayName string    `json:"displayName"`
	Avatar      *MediaRef `json:"avatar"`
}

type Client interface {
	GetMany(context.Context, []string) map[string]PublicUser
}

type httpClient struct{ base string }

func NewHTTP(baseURL string) Client {
	return &httpClient{base: strings.TrimRight(baseURL, "/")}
}

func (client *httpClient) GetMany(ctx context.Context, userKeys []string) map[string]PublicUser {
	out := make(map[string]PublicUser, len(userKeys))
	unique := make([]string, 0, len(userKeys))
	seen := map[string]bool{}
	for _, userKey := range userKeys {
		userKey = strings.TrimSpace(userKey)
		if userKey != "" && !seen[userKey] {
			seen[userKey] = true
			unique = append(unique, userKey)
		}
	}
	if len(unique) == 0 || client.base == "" {
		return out
	}
	query := url.Values{"ids": {strings.Join(unique, ",")}}
	request, err := http.NewRequestWithContext(
		ctx, http.MethodGet, client.base+"/api/v1/users?"+query.Encode(), nil,
	)
	if err != nil {
		return out
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return out
	}
	defer response.Body.Close()
	result, err := foundationhttpclient.DecodeJSON[struct {
		Users []PublicUser `json:"users"`
	}](response, foundationhttpclient.Limits{})
	if err != nil {
		return out
	}
	for _, user := range result.Users {
		if user.UserKey != "" {
			out[user.UserKey] = user
		}
	}
	return out
}
