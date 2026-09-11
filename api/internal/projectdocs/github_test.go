package projectdocs

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func TestRepositoryBoundary(t *testing.T) {
	for _, input := range []string{"https://github.com/Yuelioi/keycrash", "Yuelioi/keycrash.git"} {
		repo, e := Repository(input)
		if e != nil || repo != "Yuelioi/keycrash" {
			t.Fatalf("%s: %s %v", input, repo, e)
		}
	}
	for _, input := range []string{"https://github.com.evil.test/a/b", "http://127.0.0.1/a/b", "https://github.com/a/b/issues", "../repo", "https://user@github.com/a/b", "https://github.com/a/b?x=1"} {
		if _, e := Repository(input); e == nil {
			t.Fatal(input)
		}
	}
	for _, ip := range []string{"127.0.0.1", "192.168.5.5", "169.254.169.254", "::1", "fc00::1"} {
		if publicIP(net.ParseIP(ip)) {
			t.Fatal(ip)
		}
	}
}

func TestLatestAcceptsCanonicalGitHubCaseAndBacksOff(t *testing.T) {
	limited := false
	g := &GitHub{api: "https://api.github.com", client: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if limited {
			return &http.Response{StatusCode: 429, Body: io.NopCloser(strings.NewReader("")), Header: http.Header{"Retry-After": []string{"3600"}}}, nil
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"id":1,"tag_name":"v1","assets":[{"id":1,"name":"docs.zip","size":3,"browser_download_url":"https://github.com/Owner/Project/releases/download/v1/docs.zip"}]}`)), Header: make(http.Header)}, nil
	})}}
	if _, _, err := g.Latest(context.Background(), "owner/project", "docs.zip"); err != nil {
		t.Fatal(err)
	}
	limited = true
	_, _, err := g.Latest(context.Background(), "owner/project", "docs.zip")
	var limit *RateLimitError
	if !errors.As(err, &limit) || limit.RetryAt.Before(time.Now().Add(59*time.Minute)) {
		t.Fatalf("rate limit=%v", err)
	}
}
func TestDownloadSizeDigestAndOrigin(t *testing.T) {
	client := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader("zip")), Header: make(http.Header)}, nil
	})}
	g := &GitHub{client: client}
	asset := ReleaseAsset{URL: "https://github.com/a/b/releases/download/v1/docs.zip", Size: 3}
	data, digest, e := g.Download(context.Background(), asset)
	if e != nil || string(data) != "zip" || !strings.HasPrefix(digest, "sha256:") {
		t.Fatal(e)
	}
	asset.Digest = "sha256:invalid"
	if _, _, e = g.Download(context.Background(), asset); e == nil {
		t.Fatal("digest accepted")
	}
	asset.Digest = ""
	asset.Size = 4
	if _, _, e = g.Download(context.Background(), asset); e == nil {
		t.Fatal("size accepted")
	}
	asset.URL = "https://127.0.0.1/private"
	if _, _, e = g.Download(context.Background(), asset); e == nil {
		t.Fatal("origin accepted")
	}
}
