package controller

import (
	"context"
	"fmt"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/model"
	"github.com/yueli-official/docs/api/internal/testidentity"
	foundationauth "github.com/yueli-official/foundation/go/auth"
	"github.com/yueli-official/foundation/go/authorization"
	"testing"
)

type countedReads struct {
	docsauthz.Runtime
	decisions int
}

func (r *countedReads) Decide(ctx context.Context, req authorization.DecisionRequest) (authorization.Decision, error) {
	r.decisions++
	return r.Runtime.Decide(ctx, req)
}

func TestLargeDocumentCollectionUsesProvenAccessWithoutPerDocumentDecisions(t *testing.T) {
	admin := authorization.SubjectRef{Kind: authorization.SubjectUser, ID: "admin"}
	m, err := authorization.NewMemory(authorization.MustCompile(docsauthz.Definition()), authorization.MemoryOptions{RootScopeID: docsauthz.RootScopeID, ProtectedSubjects: []authorization.SubjectRef{admin}, Constraints: docsauthz.ConstraintEvaluators(), Predicates: docsauthz.PredicateEvaluators()})
	if err != nil {
		t.Fatal(err)
	}
	runtime := &countedReads{Runtime: m}
	service := docsauthz.New(runtime)
	if err := service.EnsureCollectionScope(context.Background(), "large"); err != nil {
		t.Fatal(err)
	}
	items := make([]*model.Doc, 1119)
	for i := range items {
		items[i] = &model.Doc{ID: fmt.Sprint(i), CollectionID: "large", AuthorSub: "owner"}
	}
	for _, subject := range []string{"admin", "ungranted"} {
		ctx := foundationauth.NewContext(context.Background(), testidentity.User(t, subject, nil, nil))
		ctx = context.WithValue(ctx, authorizationContextKey{}, service)
		runtime.decisions = 0
		visible, err := readableDocuments(ctx, "large", items)
		if err != nil {
			t.Fatal(err)
		}
		if subject == "admin" {
			if len(visible) != 1119 || runtime.decisions != 0 {
				t.Fatalf("collection access: visible=%d decisions=%d", len(visible), runtime.decisions)
			}
		} else if len(visible) != 0 || runtime.decisions != 1119 {
			t.Fatalf("restricted access bypassed: visible=%d decisions=%d", len(visible), runtime.decisions)
		}
	}
}
