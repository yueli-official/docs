package docsauthz_test

import (
	"context"
	"testing"

	"github.com/yueli-official/foundation/go/authorization"
	"platform/products/docs/api/internal/docsauthz"
)

func TestDefinitionEnforcesVisitorAuthorOwnerAndAdministratorContract(t *testing.T) {
	catalog := authorization.MustCompile(docsauthz.Definition())
	admin := authorization.SubjectRef{Kind: authorization.SubjectUser, ID: "admin"}
	author := authorization.SubjectRef{Kind: authorization.SubjectUser, ID: "author"}
	module, err := authorization.NewMemory(catalog, authorization.MemoryOptions{
		RootScopeID:       docsauthz.RootScopeID,
		ProtectedSubjects: []authorization.SubjectRef{admin},
		Constraints:       docsauthz.ConstraintEvaluators(),
		Predicates:        docsauthz.PredicateEvaluators(),
	})
	if err != nil {
		t.Fatalf("NewMemory() error = %v", err)
	}
	ctx := context.Background()
	if _, err := module.RegisterScope(ctx, authorization.RegisterScopeCommand{
		ID: docsauthz.CollectionScopeID("collection-1"),
		Type: docsauthz.ScopeCollection, ParentID: docsauthz.RootScopeID,
	}); err != nil {
		t.Fatalf("CreateScope() collection error = %v", err)
	}
	if _, err := module.RegisterScope(ctx, authorization.RegisterScopeCommand{
		ID: docsauthz.DocumentScopeID("doc-1"),
		Type: docsauthz.ScopeDocument, ParentID: docsauthz.CollectionScopeID("collection-1"),
	}); err != nil {
		t.Fatalf("CreateScope() document error = %v", err)
	}
	if _, err := module.Grant(ctx, authorization.GrantCommand{
		Actor: admin, Target: author, Role: docsauthz.RoleAuthor,
		ScopeID: docsauthz.RootScopeID, Source: authorization.GrantSourceDirect,
	}); err != nil {
		t.Fatalf("Grant() author error = %v", err)
	}
	visitor, err := module.Decide(ctx, authorization.DecisionRequest{
		Subject:    authorization.SubjectRef{Kind: authorization.SubjectAnonymous},
		Capability: docsauthz.CapabilityPublicRead,
		ScopeID:    docsauthz.DocumentScopeID("doc-1"),
	})
	if err != nil || !visitor.Allowed {
		t.Fatalf("Decide() visitor = %#v, %v, want allow", visitor, err)
	}
	otherOwner := authorization.ResourceFacts{
		Type: "document", ID: "doc-1", ScopeID: docsauthz.DocumentScopeID("doc-1"),
		Relations: map[authorization.RelationKind][]authorization.SubjectRef{
			docsauthz.RelationOwner: {{Kind: authorization.SubjectUser, ID: "someone-else"}},
		},
	}
	denied, err := module.Decide(ctx, authorization.DecisionRequest{
		Subject: author, Capability: docsauthz.CapabilityDocumentPublish,
		ScopeID: docsauthz.DocumentScopeID("doc-1"), Resource: otherOwner,
	})
	if err != nil {
		t.Fatalf("Decide() author other error = %v", err)
	}
	if denied.Allowed || denied.Reason != authorization.ReasonConstraint {
		t.Fatalf("Decide() author other = %#v, want owner constraint deny", denied)
	}
	otherOwner.Relations[docsauthz.RelationOwner] = []authorization.SubjectRef{author}
	own, err := module.Decide(ctx, authorization.DecisionRequest{
		Subject: author, Capability: docsauthz.CapabilityDocumentPublish,
		ScopeID: docsauthz.DocumentScopeID("doc-1"), Resource: otherOwner,
	})
	if err != nil || !own.Allowed {
		t.Fatalf("Decide() author own = %#v, %v, want allow", own, err)
	}
	plan, err := module.Plan(ctx, authorization.QueryRequest{
		Subject: author, Capability: docsauthz.CapabilityDocumentRead,
		ScopeID: docsauthz.CollectionScopeID("collection-1"),
	})
	if err != nil || plan.Kind != authorization.QueryRelation ||
		plan.Relation != docsauthz.RelationOwner || plan.Subject != author {
		t.Fatalf("Plan() = %#v, %v; want owner relation", plan, err)
	}
	otherOwner.Relations[docsauthz.RelationOwner] = []authorization.SubjectRef{
		{Kind: authorization.SubjectUser, ID: "someone-else"},
	}
	adminDecision, err := module.Decide(ctx, authorization.DecisionRequest{
		Subject: admin, Capability: docsauthz.CapabilityDocumentPublish,
		ScopeID: docsauthz.DocumentScopeID("doc-1"), Resource: otherOwner,
	})
	if err != nil || !adminDecision.Allowed {
		t.Fatalf("Decide() administrator = %#v, %v, want independent allow", adminDecision, err)
	}
}
