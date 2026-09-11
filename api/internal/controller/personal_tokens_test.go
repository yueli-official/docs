package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/yueli-official/docs/api/api/v1"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/testidentity"
	auth "github.com/yueli-official/foundation/go/auth"
	"github.com/yueli-official/foundation/go/authorization"
)

func TestPersonalImportRequiresScopeAndCurrentAdministrator(t *testing.T) {
	module, err := authorization.NewMemory(authorization.MustCompile(docsauthz.Definition()), authorization.MemoryOptions{
		RootScopeID: docsauthz.RootScopeID, ProtectedSubjects: []authorization.SubjectRef{{Kind: authorization.SubjectUser, ID: "owner"}},
		Constraints: docsauthz.ConstraintEvaluators(), Predicates: docsauthz.PredicateEvaluators(),
	})
	if err != nil {
		t.Fatal(err)
	}
	service := docsauthz.New(module)
	scope, _ := auth.PersonalScope("docs-main-web", string(docsauthz.CapabilityImportManage))
	user, scopes := "owner", []string{scope}
	endpoint := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"userKey": user, "scopes": scopes})
	}))
	defer endpoint.Close()
	verifier, err := auth.NewPersonalTokenVerifier(endpoint.URL, "docs-main-web", nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name, user string
		scopes     []string
		allowed    bool
	}{
		{"current administrator", "owner", []string{scope}, true},
		{"missing selected scope", "owner", []string{"site:ZG9jcy1tYWluLXdlYg:docs.doc.read"}, false},
		{"scope without administrator", "other", []string{scope}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			user, scopes = tc.user, tc.scopes
			p, e := verifier.Verify(context.Background(), "pat_test")
			if e != nil {
				t.Fatal(e)
			}
			ctx := auth.NewContext(context.Background(), p)
			d, e := service.Decide(ctx, docsauthz.CapabilityImportManage, docsauthz.RootScopeID, authorization.ResourceFacts{})
			if e != nil || d.Allowed != tc.allowed {
				t.Fatalf("decision=%+v error=%v", d, e)
			}
			_, e = NewPersonalPermissions("docs-main-web", service).AuthorizePersonalMedia(ctx, &v1.PersonalMediaAuthorizationReq{})
			if (e == nil) != tc.allowed {
				t.Fatalf("media authorization=%v", e)
			}
		})
	}
}

func TestPersonalPermissionDirectoryIsServiceOnly(t *testing.T) {
	c := NewPersonalPermissions("docs-main-web", nil)
	for _, ctx := range []context.Context{context.Background(), auth.NewContext(context.Background(), testidentity.User(t, "owner", nil, nil)), auth.NewContext(context.Background(), testidentity.Client(t, "other", nil))} {
		if _, err := c.GetPersonalPermissions(ctx, &v1.PersonalPermissionsReq{UserKey: "owner"}); err == nil {
			t.Fatal("directory exposed")
		}
	}
}

func TestPersonalDocumentScopeDoesNotGrantSiblingCollections(t *testing.T) {
	ctx := context.Background()
	admin := authorization.SubjectRef{Kind: authorization.SubjectUser, ID: "admin"}
	writer := authorization.SubjectRef{Kind: authorization.SubjectUser, ID: "writer"}
	module, err := authorization.NewMemory(authorization.MustCompile(docsauthz.Definition()), authorization.MemoryOptions{RootScopeID: docsauthz.RootScopeID, ProtectedSubjects: []authorization.SubjectRef{admin}, Constraints: docsauthz.ConstraintEvaluators(), Predicates: docsauthz.PredicateEvaluators()})
	if err != nil {
		t.Fatal(err)
	}
	service := docsauthz.New(module)
	for _, id := range []string{"a", "b"} {
		if err := service.EnsureCollectionScope(ctx, id); err != nil {
			t.Fatal(err)
		}
	}
	grant, err := module.Grant(ctx, authorization.GrantCommand{Actor: admin, Target: writer, Role: docsauthz.RoleAuthor, ScopeID: docsauthz.CollectionScopeID("a"), Source: authorization.GrantSourceDirect})
	if err != nil {
		t.Fatal(err)
	}
	catalogCtx := auth.NewContext(ctx, testidentity.Client(t, "identity-svc", []string{auth.PersonalPermissionsScope}))
	directory, err := NewPersonalPermissions("docs-main-web", service).GetPersonalPermissions(catalogCtx, &v1.PersonalPermissionsReq{UserKey: "writer"})
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, item := range directory.Items {
		found = found || item.Key == string(docsauthz.CapabilityDocumentCreate)
	}
	if !found {
		t.Fatal("collection capability absent from directory")
	}
	scope, _ := auth.PersonalScope("docs-main-web", string(docsauthz.CapabilityDocumentCreate))
	endpoint := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"userKey": "writer", "scopes": []string{scope}})
	}))
	defer endpoint.Close()
	verifier, err := auth.NewPersonalTokenVerifier(endpoint.URL, "docs-main-web", nil)
	if err != nil {
		t.Fatal(err)
	}
	p, err := verifier.Verify(ctx, "pat_test")
	if err != nil {
		t.Fatal(err)
	}
	userCtx := auth.NewContext(ctx, p)
	for _, id := range []string{"a", "b"} {
		d, err := service.Decide(userCtx, docsauthz.CapabilityDocumentCreate, docsauthz.CollectionScopeID(id), authorization.ResourceFacts{})
		if err != nil || d.Allowed != (id == "a") {
			t.Fatalf("collection %s: %+v %v", id, d, err)
		}
	}
	if _, err = module.Revoke(ctx, authorization.RevokeCommand{Actor: admin, GrantID: grant.ID}); err != nil {
		t.Fatal(err)
	}
	d, err := service.Decide(userCtx, docsauthz.CapabilityDocumentCreate, docsauthz.CollectionScopeID("a"), authorization.ResourceFacts{})
	if err != nil || d.Allowed {
		t.Fatalf("revoked role still granted: %+v %v", d, err)
	}
}
