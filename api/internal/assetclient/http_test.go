package assetclient

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHTTPClientUsesConfiguredSiteContextForAssetWrites(t *testing.T) {
	var bodies []map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			bodies = append(bodies, map[string]any{"siteKey": r.URL.Query().Get("siteKey")})
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{}`))
			return
		}
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("read request: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var body map[string]any
		if err := json.Unmarshal(raw, &body); err != nil {
			values, parseErr := url.ParseQuery(string(raw))
			if parseErr != nil {
				t.Errorf("decode request: json=%v form=%v", err, parseErr)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			body = make(map[string]any, len(values))
			for key := range values {
				body[key] = values.Get(key)
			}
		}
		bodies = append(bodies, body)
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/v1/assets/upload-init" {
			_, _ = w.Write([]byte(`{"uploadUrl":"http://upload.test","uploadToken":"token"}`))
			return
		}
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHTTP(server.URL, "docs-main", "yueli")
	if _, err := client.UploadInit(context.Background(), "token", InitInput{
		Filename: "cover.png", Mime: "image/png", Category: "docs-collection-cover", Visibility: "public", Size: 12,
	}); err != nil {
		t.Fatalf("UploadInit() error = %v", err)
	}
	if err := client.RegisterReference(context.Background(), "token", ReferenceInput{
		AssetID: "asset-1", RefType: "collection-cover", RefID: "collection-1",
	}); err != nil {
		t.Fatalf("RegisterReference() error = %v", err)
	}
	if err := client.UnregisterReference(context.Background(), "token", ReferenceInput{
		AssetID: "asset-1", RefType: "collection-cover", RefID: "collection-1",
	}); err != nil {
		t.Fatalf("UnregisterReference() error = %v", err)
	}

	if len(bodies) != 3 {
		t.Fatalf("request count = %d, want 3", len(bodies))
	}
	for i, body := range bodies {
		if body["siteKey"] != "docs-main" {
			t.Fatalf("request %d siteKey = %#v, want docs-main", i, body["siteKey"])
		}
	}
	if bodies[0]["spaceKey"] != "yueli" {
		t.Fatalf("upload spaceKey = %#v, want yueli", bodies[0]["spaceKey"])
	}
}

func TestHTTPClientUploadPutsExactBytesBeforeFinalize(t *testing.T) {
	payload := []byte("exact-image-bytes")
	var uploaded []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/v1/assets/upload-init":
			_, _ = w.Write([]byte(`{"uploadUrl":"` + "http://" + r.Host + `/blob","uploadToken":"token"}`))
		case "/blob":
			if r.ContentLength != int64(len(payload)) {
				t.Errorf("Content-Length = %d, want %d", r.ContentLength, len(payload))
			}
			uploaded, _ = io.ReadAll(r.Body)
		case "/api/v1/assets/finalize":
			_, _ = w.Write([]byte(`{"asset":{"id":"a1","mediaKey":"docs/test","size":17,"mime":"image/png","filename":"a.png"}}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()
	client := NewHTTP(server.URL, "docs-main", "yueli")
	if _, err := client.Upload(context.Background(), "token", InitInput{Filename: "a.png", Mime: "image/png", Category: "docs-import-image", Size: int64(len(payload))}, payload); err != nil {
		t.Fatalf("Upload() error = %v", err)
	}
	if string(uploaded) != string(payload) {
		t.Fatalf("uploaded = %q, want %q", uploaded, payload)
	}
}

func TestHTTPClientRetriesAssetRateLimit(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		w.Header().Set("Content-Type", "application/json")
		if requests == 1 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"type":"about:blank","status":429,"code":"common.rate_limited","traceId":"rate-test"}`))
			return
		}
		_, _ = w.Write([]byte(`{"uploadUrl":"http://upload.test","uploadToken":"token"}`))
	}))
	defer server.Close()

	client := NewHTTP(server.URL, "docs-main", "yueli")
	if _, err := client.UploadInit(context.Background(), "token", InitInput{Filename: "a.png", Mime: "image/png", Category: "docs-import-image", Size: 1}); err != nil {
		t.Fatalf("UploadInit() error = %v", err)
	}
	if requests != 2 {
		t.Fatalf("requests=%d, want one retry", requests)
	}
}
