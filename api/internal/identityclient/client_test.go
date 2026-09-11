package identityclient

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetManyReadsIdentityBatchItems(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/api/v1/users" || r.URL.Query().Get("ids") != "UserA123,UserB234" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.RequestURI())
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"items":[{"userKey":"UserA123","displayName":"月离","avatar":{"mediaKey":"34kNV1Rw14KiopnMv5xtu"}},{"userKey":"UserB234","displayName":"读者"}]}`)
	}))
	defer server.Close()

	users := NewHTTP(server.URL).GetMany(context.Background(), []string{"UserA123", "UserA123", "", "UserB234"})
	if len(users) != 2 || users["UserA123"].DisplayName != "月离" || users["UserB234"].DisplayName != "读者" {
		t.Fatalf("public profiles lost: %#v", users)
	}
	if avatar := users["UserA123"].Avatar; avatar == nil || avatar.MediaKey != "34kNV1Rw14KiopnMv5xtu" {
		t.Fatalf("avatar lost: %#v", avatar)
	}
}
