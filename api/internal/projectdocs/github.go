// Package projectdocs imports released documentation without executing source repositories.
package projectdocs

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const MaxPackageBytes = 100 << 20

type Release struct {
	ID         int64          `json:"id"`
	Tag        string         `json:"tag_name"`
	Draft      bool           `json:"draft"`
	Prerelease bool           `json:"prerelease"`
	Assets     []ReleaseAsset `json:"assets"`
}
type ReleaseAsset struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	URL    string `json:"browser_download_url"`
	Size   int64  `json:"size"`
	Digest string `json:"digest"`
}
type GitHub struct {
	client *http.Client
	api    string
	token  string
}

type RateLimitError struct{ RetryAt time.Time }

func (e *RateLimitError) Error() string {
	return "GitHub 请求额度已用完，将在额度恢复后继续检查。"
}

var repositoryPattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9_.-]*/[A-Za-z0-9][A-Za-z0-9_.-]*$`)

func Repository(input string) (string, error) {
	input = strings.TrimSpace(input)
	if strings.HasPrefix(input, "https://") {
		u, err := url.Parse(input)
		if err != nil || u.Host != "github.com" || u.User != nil || u.RawQuery != "" || u.Fragment != "" {
			return "", errors.New("expected a public GitHub repository URL")
		}
		input = strings.Trim(u.Path, "/")
	}
	input = strings.TrimSuffix(input, ".git")
	if !repositoryPattern.MatchString(input) || strings.Contains(input, "..") {
		return "", errors.New("expected owner/repository")
	}
	return input, nil
}

func allowedHost(host string) bool {
	return host == "api.github.com" || host == "github.com" || host == "release-assets.githubusercontent.com" || host == "objects.githubusercontent.com"
}
func allowedURL(u *url.URL) bool {
	return u.Scheme == "https" && u.User == nil && u.Port() == "" && allowedHost(u.Hostname())
}
func publicIP(ip net.IP) bool {
	return ip.IsGlobalUnicast() && !ip.IsPrivate() && !ip.IsLoopback() && !ip.IsLinkLocalUnicast()
}

func NewGitHub(token string) *GitHub {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	// Dial only the addresses we checked; do not resolve again after the check.
	transport.Proxy = nil
	transport.DialContext = func(ctx context.Context, network, address string) (net.Conn, error) {
		host, port, err := net.SplitHostPort(address)
		if err != nil || !allowedHost(host) || port != "443" {
			return nil, errors.New("GitHub network target refused")
		}
		ips, err := net.DefaultResolver.LookupIPAddr(ctx, host)
		if err != nil {
			return nil, err
		}
		for _, candidate := range ips {
			if !publicIP(candidate.IP) {
				return nil, errors.New("GitHub resolved to a non-public address")
			}
		}
		dialer := net.Dialer{Timeout: 15 * time.Second}
		for _, candidate := range ips {
			conn, e := dialer.DialContext(ctx, network, net.JoinHostPort(candidate.IP.String(), port))
			if e == nil {
				return conn, nil
			}
			err = e
		}
		if err == nil {
			err = errors.New("GitHub has no address")
		}
		return nil, err
	}
	client := &http.Client{Transport: transport, Timeout: 2 * time.Minute, CheckRedirect: func(req *http.Request, via []*http.Request) error {
		if len(via) > 3 || !allowedURL(req.URL) {
			return errors.New("GitHub redirect refused")
		}
		return nil
	}}
	return &GitHub{client: client, api: "https://api.github.com", token: token}
}

func (g *GitHub) read(ctx context.Context, target string, limit int64) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Yueli-Docs-Importer")
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	if req.URL.Host == "api.github.com" && g.token != "" {
		req.Header.Set("Authorization", "Bearer "+g.token)
	}
	response, err := g.client.Do(req)
	if err != nil {
		return nil, errors.New("GitHub request failed")
	}
	defer response.Body.Close()
	if response.StatusCode == 429 || (response.StatusCode == 403 && response.Header.Get("X-RateLimit-Remaining") == "0") {
		retry := time.Now().Add(time.Hour)
		if seconds, e := strconv.ParseInt(response.Header.Get("X-RateLimit-Reset"), 10, 64); e == nil && seconds > time.Now().Unix() {
			retry = time.Unix(seconds, 0)
		}
		if seconds, e := strconv.Atoi(response.Header.Get("Retry-After")); e == nil && seconds > 0 {
			retry = time.Now().Add(time.Duration(seconds) * time.Second)
		}
		return nil, &RateLimitError{RetryAt: retry.Add(time.Minute)}
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub returned HTTP %d", response.StatusCode)
	}
	if response.ContentLength > limit {
		return nil, errors.New("GitHub content exceeds size limit")
	}
	data, err := io.ReadAll(io.LimitReader(response.Body, limit+1))
	if err != nil {
		return nil, errors.New("GitHub download interrupted")
	}
	if int64(len(data)) > limit {
		return nil, errors.New("GitHub content exceeds size limit")
	}
	return data, nil
}

func (g *GitHub) Latest(ctx context.Context, repository, assetName string) (Release, ReleaseAsset, error) {
	repo, err := Repository(repository)
	if err != nil {
		return Release{}, ReleaseAsset{}, err
	}
	if assetName == "" || strings.ContainsAny(assetName, "/\\") || !strings.HasSuffix(strings.ToLower(assetName), ".zip") {
		return Release{}, ReleaseAsset{}, errors.New("expected a ZIP asset filename")
	}
	raw, err := g.read(ctx, g.api+"/repos/"+repo+"/releases/latest", 2<<20)
	if err != nil {
		return Release{}, ReleaseAsset{}, err
	}
	var release Release
	if err = json.Unmarshal(raw, &release); err != nil || release.ID <= 0 || release.Draft || release.Prerelease {
		return Release{}, ReleaseAsset{}, errors.New("no eligible published release")
	}
	for _, asset := range release.Assets {
		if asset.Name != assetName {
			continue
		}
		u, e := url.Parse(asset.URL)
		validPath := false
		if e == nil {
			parts := strings.Split(strings.TrimPrefix(u.Path, "/"), "/")
			validPath = len(parts) >= 6 && strings.EqualFold(parts[0]+"/"+parts[1], repo) && parts[2] == "releases" && parts[3] == "download"
		}
		if e != nil || !allowedURL(u) || u.Host != "github.com" || !validPath || asset.ID <= 0 || asset.Size <= 0 || asset.Size > MaxPackageBytes {
			return Release{}, ReleaseAsset{}, errors.New("invalid documentation release asset")
		}
		return release, asset, nil
	}
	return Release{}, ReleaseAsset{}, errors.New("documentation ZIP is missing from the latest release")
}

func (g *GitHub) Download(ctx context.Context, asset ReleaseAsset) ([]byte, string, error) {
	u, err := url.Parse(asset.URL)
	if err != nil || !allowedURL(u) {
		return nil, "", errors.New("invalid GitHub asset URL")
	}
	data, err := g.read(ctx, asset.URL, MaxPackageBytes)
	if err != nil {
		return nil, "", err
	}
	if int64(len(data)) != asset.Size {
		return nil, "", errors.New("release asset size changed")
	}
	sum := sha256.Sum256(data)
	digest := "sha256:" + hex.EncodeToString(sum[:])
	if asset.Digest != "" && asset.Digest != digest {
		return nil, "", errors.New("release asset digest mismatch")
	}
	return data, digest, nil
}
