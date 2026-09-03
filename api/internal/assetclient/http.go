package assetclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	foundationhttpclient "github.com/yueli-official/foundation/go/httpclient"

	"github.com/yueli-official/docs/api/internal/docserr"
)

type httpClient struct {
	base     string
	siteSlug string
	spaceKey string
}

func NewHTTP(baseURL, siteSlug, spaceKey string) Client {
	return &httpClient{base: strings.TrimRight(baseURL, "/"), siteSlug: siteSlug, spaceKey: spaceKey}
}

func (c *httpClient) post(ctx context.Context, bearer, path string, body g.Map) (*gjson.Json, error) {
	raw, _ := json.Marshal(body)
	for attempt := 0; attempt < 3; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.base+path, bytes.NewReader(raw))
		if err != nil {
			return nil, docserr.UpstreamFailed("foundation.request.invalid")
		}
		req.Header.Set("Authorization", "Bearer "+bearer)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, docserr.UpstreamFailed("asset service unreachable")
		}
		if resp.StatusCode == http.StatusTooManyRequests && attempt < 2 {
			if !waitForRateLimit(ctx, resp) {
				return nil, docserr.UpstreamFailed("common.rate_limited")
			}
			continue
		}
		defer resp.Body.Close()
		out, err := foundationhttpclient.DecodeJSON[map[string]any](resp, foundationhttpclient.Limits{})
		if err != nil {
			return nil, docserr.UpstreamFailed(remoteCode(err))
		}
		return gjson.New(out), nil
	}
	return nil, docserr.UpstreamFailed("common.rate_limited")
}

func waitForRateLimit(ctx context.Context, resp *http.Response) bool {
	delay := time.Second
	for _, name := range []string{"Retry-After", "Ratelimit-Reset"} {
		if seconds, err := strconv.Atoi(strings.TrimSpace(resp.Header.Get(name))); err == nil && seconds > 0 {
			delay = time.Duration(seconds) * time.Second
			break
		}
	}
	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-timer.C:
		return true
	}
}

func (c *httpClient) UploadInit(ctx context.Context, bearer string, in InitInput) (InitOutput, error) {
	j, err := c.post(ctx, bearer, "/api/v1/assets/upload-init", g.Map{
		"filename":   in.Filename,
		"mime":       in.Mime,
		"size":       in.Size,
		"category":   in.Category,
		"siteKey":    c.siteSlug,
		"spaceKey":   c.spaceKey,
		"profileKey": in.Category,
		"visibility": in.Visibility,
	})
	if err != nil {
		return InitOutput{}, err
	}
	return InitOutput{
		UploadURL:     j.Get("uploadUrl").String(),
		UploadToken:   j.Get("uploadToken").String(),
		UploadHeaders: stringMap(j.Get("uploadHeaders").Map()),
	}, nil
}

func stringMap(raw map[string]any) map[string]string {
	if len(raw) == 0 {
		return nil
	}
	out := make(map[string]string, len(raw))
	for k, v := range raw {
		out[k] = g.NewVar(v).String()
	}
	return out
}

func (c *httpClient) Finalize(ctx context.Context, bearer, uploadToken string) (View, error) {
	j, err := c.post(ctx, bearer, "/api/v1/assets/finalize", g.Map{"uploadToken": uploadToken})
	if err != nil {
		return View{}, err
	}
	return View{
		ID:       j.Get("asset.id").String(),
		MediaKey: j.Get("asset.mediaKey").String(),
		Size:     j.Get("asset.size").Int64(),
		Mime:     j.Get("asset.mime").String(),
		Filename: j.Get("asset.filename").String(),
	}, nil
}

func (c *httpClient) RegisterReference(ctx context.Context, bearer string, in ReferenceInput) error {
	_, err := c.post(ctx, bearer, "/api/v1/asset-references", g.Map{
		"assetId": in.AssetID, "siteKey": c.siteSlug, "refType": in.RefType, "refId": in.RefID,
		"refLabel": in.RefLabel, "refUrl": in.RefURL,
	})
	return err
}

func (c *httpClient) UnregisterReference(ctx context.Context, bearer string, in ReferenceInput) error {
	q := url.Values{}
	q.Set("assetId", in.AssetID)
	q.Set("siteKey", c.siteSlug)
	q.Set("refType", in.RefType)
	q.Set("refId", in.RefID)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.base+"/api/v1/asset-references?"+q.Encode(), nil)
	if err != nil {
		return docserr.UpstreamFailed("foundation.request.invalid")
	}
	req.Header.Set("Authorization", "Bearer "+bearer)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return docserr.UpstreamFailed("asset service unreachable")
	}
	defer resp.Body.Close()
	if _, err := foundationhttpclient.DecodeJSON[any](resp, foundationhttpclient.Limits{}); err != nil {
		return docserr.UpstreamFailed(remoteCode(err))
	}
	return nil
}

func remoteCode(err error) string {
	var remote *foundationhttpclient.RemoteError
	if errors.As(err, &remote) {
		return remote.Problem.Code
	}
	return "foundation.response.invalid"
}

func (c *httpClient) Upload(ctx context.Context, bearer string, in InitInput, data []byte) (View, error) {
	out, err := c.UploadInit(ctx, bearer, in)
	if err != nil {
		return View{}, err
	}
	for attempt := 0; attempt < 3; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodPut, out.UploadURL, bytes.NewReader(data))
		if err != nil {
			return View{}, docserr.UpstreamFailed("asset blob upload failed")
		}
		req.ContentLength = int64(len(data))
		if in.Mime != "" {
			req.Header.Set("Content-Type", in.Mime)
		}
		for k, v := range out.UploadHeaders {
			req.Header.Set(k, v)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return View{}, docserr.UpstreamFailed("asset blob upload failed")
		}
		if resp.StatusCode == http.StatusTooManyRequests && attempt < 2 {
			if !waitForRateLimit(ctx, resp) {
				return View{}, docserr.UpstreamFailed("common.rate_limited")
			}
			continue
		}
		_ = resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return View{}, docserr.UpstreamFailed("asset blob upload failed")
		}
		return c.Finalize(ctx, bearer, out.UploadToken)
	}
	return View{}, docserr.UpstreamFailed("common.rate_limited")
}
