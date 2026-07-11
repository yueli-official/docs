package assetclient

import (
	"bytes"
	"context"
	"net/url"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"

	"platform/products/docs/api/internal/docserr"
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
	cli := g.Client()
	cli.SetHeader("Authorization", "Bearer "+bearer)
	cli.ContentJson()
	resp, err := cli.Post(ctx, c.base+path, body)
	if err != nil {
		return nil, docserr.UpstreamFailed("asset service unreachable")
	}
	defer resp.Close()
	j := gjson.New(resp.ReadAllString())
	if code := j.Get("code").String(); code != "ok" {
		return nil, docserr.UpstreamFailed(code)
	}
	return j, nil
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
		UploadURL:     j.Get("data.uploadUrl").String(),
		UploadToken:   j.Get("data.uploadToken").String(),
		UploadHeaders: stringMap(j.Get("data.uploadHeaders").Map()),
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
		ID:       j.Get("data.asset.id").String(),
		CdnURL:   j.Get("data.asset.cdnUrl").String(),
		Size:     j.Get("data.asset.size").Int64(),
		Mime:     j.Get("data.asset.mime").String(),
		Filename: j.Get("data.asset.filename").String(),
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
	cli := g.Client()
	cli.SetHeader("Authorization", "Bearer "+bearer)
	q := url.Values{}
	q.Set("assetId", in.AssetID)
	q.Set("siteKey", c.siteSlug)
	q.Set("refType", in.RefType)
	q.Set("refId", in.RefID)
	resp, err := cli.Delete(ctx, c.base+"/api/v1/asset-references?"+q.Encode())
	if err != nil {
		return docserr.UpstreamFailed("asset service unreachable")
	}
	defer resp.Close()
	if code := gjson.New(resp.ReadAllString()).Get("code").String(); code != "ok" {
		return docserr.UpstreamFailed(code)
	}
	return nil
}

func (c *httpClient) Upload(ctx context.Context, bearer string, in InitInput, data []byte) (View, error) {
	out, err := c.UploadInit(ctx, bearer, in)
	if err != nil {
		return View{}, err
	}
	cli := g.Client()
	if in.Mime != "" {
		cli.SetHeader("Content-Type", in.Mime)
	}
	for k, v := range out.UploadHeaders {
		cli.SetHeader(k, v)
	}
	resp, err := cli.Put(ctx, out.UploadURL, bytes.NewReader(data))
	if err != nil {
		return View{}, docserr.UpstreamFailed("asset blob upload failed")
	}
	defer resp.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return View{}, docserr.UpstreamFailed("asset blob upload failed")
	}
	return c.Finalize(ctx, bearer, out.UploadToken)
}
