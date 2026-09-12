package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func personalContext(t *testing.T, user string, capabilities ...authorization.CapabilityKey) context.Context {
	t.Helper()
	scopes := make([]string, 0, len(capabilities))
	for _, capability := range capabilities {
		scope, err := auth.PersonalScope("docs-main-web", string(capability))
		if err != nil {
			t.Fatal(err)
		}
		scopes = append(scopes, scope)
	}
	endpoint := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"userKey": user, "scopes": scopes})
	}))
	defer endpoint.Close()
	verifier, err := auth.NewPersonalTokenVerifier(endpoint.URL, "docs-main-web", nil)
	if err != nil {
		t.Fatal(err)
	}
	p, err := verifier.Verify(context.Background(), "pat_test")
	if err != nil {
		t.Fatal(err)
	}
	return auth.NewContext(context.Background(), p)
}

func TestPersonalPublishingRoutesRequireSelectedCapability(t *testing.T) {
	cases := []struct {
		method, path string
		capability   authorization.CapabilityKey
	}{
		{"POST", "/api/v1/collections", docsauthz.CapabilityCollectionManage},
		{"PATCH", "/api/v1/collections/c", docsauthz.CapabilityCollectionManage},
		{"DELETE", "/api/v1/collections/c", docsauthz.CapabilityCollectionManage},
		{"POST", "/api/v1/collections/c/cover", docsauthz.CapabilityCollectionManage},
		{"POST", "/api/v1/collections/c/cover/finalize", docsauthz.CapabilityCollectionManage},
		{"GET", "/api/v1/manage/collections/c/tree", docsauthz.CapabilityCollectionManage},
		{"GET", "/api/v1/manage/collections/c/locales", docsauthz.CapabilityCollectionManage},
		{"POST", "/api/v1/manage/collections/c/locales", docsauthz.CapabilityCollectionManage},
		{"POST", "/api/v1/manage/collections/c/locales/clone", docsauthz.CapabilityCollectionManage},
		{"DELETE", "/api/v1/manage/collections/c/locales/en", docsauthz.CapabilityCollectionManage},
		{"POST", "/api/v1/manage/collections/c/clone-release", docsauthz.CapabilityCollectionManage},
		{"POST", "/api/v1/manage/collections/c/release", docsauthz.CapabilityCollectionManage},
		{"GET", "/api/v1/collections/c/versions", docsauthz.CapabilityDocumentRead},
		{"POST", "/api/v1/collections/c/versions", docsauthz.CapabilityVersionManage},
		{"PATCH", "/api/v1/manage/collections/c/versions/v", docsauthz.CapabilityVersionManage},
		{"GET", "/api/v1/docs", docsauthz.CapabilityDocumentRead},
		{"GET", "/api/v1/manage/docs", docsauthz.CapabilityDocumentRead},
		{"GET", "/api/v1/docs/d", docsauthz.CapabilityDocumentRead},
		{"POST", "/api/v1/docs", docsauthz.CapabilityDocumentCreate},
		{"PATCH", "/api/v1/docs/d", docsauthz.CapabilityDocumentUpdate},
		{"POST", "/api/v1/docs/d/publish", docsauthz.CapabilityDocumentPublish},
		{"POST", "/api/v1/docs/d/archive", docsauthz.CapabilityDocumentArchive},
		{"DELETE", "/api/v1/docs/d", docsauthz.CapabilityDocumentDeletePermanently},
		{"POST", "/api/v1/images", docsauthz.CapabilityDocumentCreate},
		{"POST", "/api/v1/images/finalize", docsauthz.CapabilityDocumentUpdate},
		{"GET", "/api/v1/imports/docs", docsauthz.CapabilityImportManage},
		{"POST", "/api/v1/imports/docs", docsauthz.CapabilityImportManage},
		{"GET", "/api/v1/imports/docs/b", docsauthz.CapabilityImportManage},
		{"POST", "/api/v1/imports/docs/b/confirm", docsauthz.CapabilityImportManage},
	}
	for _, tc := range cases {
		t.Run(tc.method+" "+tc.path, func(t *testing.T) {
			if !allowsPersonalRoute(personalContext(t, "owner", tc.capability), tc.method, tc.path) {
				t.Fatal("selected capability did not open route")
			}
			if allowsPersonalRoute(personalContext(t, "owner", docsauthz.CapabilityPublicRead), tc.method, tc.path) {
				t.Fatal("route allowed without selected capability")
			}
		})
	}
	var keys []authorization.CapabilityKey
	for _, permission := range personalPermissions {
		keys = append(keys, authorization.CapabilityKey(permission.Key))
	}
	ctx := personalContext(t, "owner", keys...)
	for _, path := range []string{"/api/v1/home", "/api/v1/import-sources", "/api/v1/authorization/roles", "/api/v1/manage/collections/c/extra/tree", "/api/v1/manage/collections//tree", "/api/v1/docs/d/publish/extra", "/api/v1/imports/docs/b/confirm/extra"} {
		for _, method := range []string{"GET", "POST", "PATCH", "DELETE"} {
			if allowsPersonalRoute(ctx, method, path) {
				t.Fatalf("unexpected route allowed: %s %s", method, path)
			}
		}
	}
}

func TestPersonalDirectoryAndMediaUseCurrentCapabilities(t *testing.T) {
	admin := authorization.SubjectRef{Kind: authorization.SubjectUser, ID: "admin"}
	module, err := authorization.NewMemory(authorization.MustCompile(docsauthz.Definition()), authorization.MemoryOptions{
		RootScopeID: docsauthz.RootScopeID, ProtectedSubjects: []authorization.SubjectRef{admin},
		Constraints: docsauthz.ConstraintEvaluators(), Predicates: docsauthz.PredicateEvaluators(),
	})
	if err != nil {
		t.Fatal(err)
	}
	service := docsauthz.New(module)
	c := NewPersonalPermissions("docs-main-web", service)
	catalogCtx := auth.NewContext(context.Background(), testidentity.Client(t, "identity-svc", []string{auth.PersonalPermissionsScope}))
	directory, err := c.GetPersonalPermissions(catalogCtx, &v1.PersonalPermissionsReq{UserKey: "admin"})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(directory.Items, personalPermissions) || len(directory.Items) != 9 {
		t.Fatalf("directory: %+v", directory)
	}
	for _, tc := range []struct {
		capability authorization.CapabilityKey
		profiles   []string
	}{
		{docsauthz.CapabilityCollectionManage, []string{"docs-collection-cover"}},
		{docsauthz.CapabilityDocumentCreate, []string{"docs-content-image"}},
		{docsauthz.CapabilityDocumentUpdate, []string{"docs-content-image"}},
		{docsauthz.CapabilityImportManage, []string{"docs-import-image"}},
		{docsauthz.CapabilityDocumentRead, nil},
		{docsauthz.CapabilityVersionManage, nil},
	} {
		out, err := c.AuthorizePersonalMedia(personalContext(t, "admin", tc.capability), &v1.PersonalMediaAuthorizationReq{})
		if len(tc.profiles) == 0 {
			if err == nil {
				t.Fatalf("%s allowed uploads", tc.capability)
			}
			continue
		}
		if err != nil {
			t.Fatal(err)
		}
		var want []string
		for _, profile := range tc.profiles {
			scope, _ := auth.PersonalScope("docs-main-web", "asset.profile."+profile+".upload")
			want = append(want, scope)
		}
		if !reflect.DeepEqual(out.Scopes, want) {
			t.Fatalf("%s: got %v want %v", tc.capability, out.Scopes, want)
		}
	}
	if err := service.EnsureCollectionScope(context.Background(), "a"); err != nil {
		t.Fatal(err)
	}
	grant, err := module.Grant(context.Background(), authorization.GrantCommand{Actor: admin, Target: authorization.SubjectRef{Kind: authorization.SubjectUser, ID: "writer"}, Role: docsauthz.RoleAuthor, ScopeID: docsauthz.CollectionScopeID("a"), Source: authorization.GrantSourceDirect})
	if err != nil {
		t.Fatal(err)
	}
	writer := personalContext(t, "writer", docsauthz.CapabilityDocumentCreate, docsauthz.CapabilityImportManage, docsauthz.CapabilityCollectionManage)
	media, err := c.AuthorizePersonalMedia(writer, &v1.PersonalMediaAuthorizationReq{})
	if err != nil || len(media.Scopes) != 1 {
		t.Fatalf("scoped author upload: %+v %v", media, err)
	}
	want, _ := auth.PersonalScope("docs-main-web", "asset.profile.docs-content-image.upload")
	if media.Scopes[0] != want {
		t.Fatalf("scoped author got unrelated profiles: %+v", media)
	}
	if _, err := module.Revoke(context.Background(), authorization.RevokeCommand{Actor: admin, GrantID: grant.ID}); err != nil {
		t.Fatal(err)
	}
	if _, err := c.AuthorizePersonalMedia(writer, &v1.PersonalMediaAuthorizationReq{}); err == nil {
		t.Fatal("revoked author can still upload")
	}
}

func TestPersonalCollectionManagerCannotCreateAtRootOrManageSibling(t *testing.T) {
	ctx := context.Background()
	admin := authorization.SubjectRef{Kind: authorization.SubjectUser, ID: "admin"}
	definition := docsauthz.Definition()
	definition.Roles = append(definition.Roles, authorization.RoleDefinition{
		Key: "collection-editor", DisplayName: "文档集编辑", Capabilities: []authorization.CapabilityKey{docsauthz.CapabilityCollectionManage},
		Assignment: authorization.AssignmentPolicy{Sources: []authorization.GrantSource{authorization.GrantSourceDirect}},
	})
	module, err := authorization.NewMemory(authorization.MustCompile(definition), authorization.MemoryOptions{RootScopeID: docsauthz.RootScopeID, ProtectedSubjects: []authorization.SubjectRef{admin}, Constraints: docsauthz.ConstraintEvaluators(), Predicates: docsauthz.PredicateEvaluators()})
	if err != nil {
		t.Fatal(err)
	}
	service := docsauthz.New(module)
	for _, id := range []string{"a", "b"} {
		if err := service.EnsureCollectionScope(ctx, id); err != nil {
			t.Fatal(err)
		}
	}
	grant, err := module.Grant(ctx, authorization.GrantCommand{Actor: admin, Target: authorization.SubjectRef{Kind: authorization.SubjectUser, ID: "manager"}, Role: "collection-editor", ScopeID: docsauthz.CollectionScopeID("a"), Source: authorization.GrantSourceDirect})
	if err != nil {
		t.Fatal(err)
	}
	manager := personalContext(t, "manager", docsauthz.CapabilityCollectionManage)
	manager = context.WithValue(manager, authorizationContextKey{}, service)
	if _, err := NewCollections(nil).CloneCollectionRelease(manager, &v1.CloneCollectionReleaseReq{ID: "a", TargetSemanticVersion: "2.0.0", Title: "clone"}); err == nil {
		t.Fatal("collection-scoped manager created a root release clone")
	}
	for _, scope := range []authorization.ScopeID{docsauthz.RootScopeID, docsauthz.CollectionScopeID("a"), docsauthz.CollectionScopeID("b")} {
		d, err := service.Decide(manager, docsauthz.CapabilityCollectionManage, scope, authorization.ResourceFacts{})
		if err != nil || d.Allowed != (scope == docsauthz.CollectionScopeID("a")) {
			t.Fatalf("scope %s: %+v %v", scope, d, err)
		}
	}
	if _, err := module.Revoke(ctx, authorization.RevokeCommand{Actor: admin, GrantID: grant.ID}); err != nil {
		t.Fatal(err)
	}
	d, err := service.Decide(manager, docsauthz.CapabilityCollectionManage, docsauthz.CollectionScopeID("a"), authorization.ResourceFacts{})
	if err != nil || d.Allowed {
		t.Fatalf("revoked manager allowed: %+v %v", d, err)
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
	userCtx = context.WithValue(userCtx, authorizationContextKey{}, service)
	for _, id := range []string{"a", "b"} {
		d, err := service.Decide(userCtx, docsauthz.CapabilityDocumentCreate, docsauthz.CollectionScopeID(id), authorization.ResourceFacts{})
		if err != nil || d.Allowed != (id == "a") {
			t.Fatalf("collection %s: %+v %v", id, d, err)
		}
		if err := authorizePersonalVersionRead(userCtx, id); (err == nil) != (id == "a") {
			t.Fatalf("version metadata collection %s: %v", id, err)
		}
	}
	if _, err = module.Revoke(ctx, authorization.RevokeCommand{Actor: admin, GrantID: grant.ID}); err != nil {
		t.Fatal(err)
	}
	d, err := service.Decide(userCtx, docsauthz.CapabilityDocumentCreate, docsauthz.CollectionScopeID("a"), authorization.ResourceFacts{})
	if err != nil || d.Allowed {
		t.Fatalf("revoked role still granted: %+v %v", d, err)
	}
	if err := authorizePersonalVersionRead(userCtx, "a"); err == nil {
		t.Fatal("revoked role still reads version metadata")
	}
}
