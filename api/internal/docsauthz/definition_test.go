package docsauthz_test

import (
	"context"
	"testing"

	"github.com/yueli-official/docs/api/internal/docsauthz"
	foundationauth "github.com/yueli-official/foundation/go/auth"
	"github.com/yueli-official/foundation/go/authorization"
)

func TestServiceReconcilesEnabledAutomaticAuthorOnFirstAuthenticatedAccess(t *testing.T) {
	admin := authorization.SubjectRef{Kind: authorization.SubjectUser, ID: "admin"}
	module, err := authorization.NewMemory(
		authorization.MustCompile(docsauthz.Definition()),
		authorization.MemoryOptions{
			RootScopeID: docsauthz.RootScopeID, ProtectedSubjects: []authorization.SubjectRef{admin},
			Constraints: docsauthz.ConstraintEvaluators(), Predicates: docsauthz.PredicateEvaluators(),
		},
	)
	if err != nil {
		t.Fatalf("NewMemory() error = %v", err)
	}
	ctx := context.Background()
	draft, err := module.CreatePolicyDraft(ctx, authorization.CreatePolicyDraftCommand{
		Actor: admin, ScopeID: docsauthz.RootScopeID, ExpectedActiveRevision: 1,
	})
	if err != nil {
		t.Fatalf("CreatePolicyDraft() error = %v", err)
	}
	if _, err := module.SetAutomaticRuleEnabled(ctx, authorization.SetAutomaticRuleEnabledCommand{
		Actor: admin, Revision: draft.Number,
		Rule: docsauthz.AutomaticRegistrationAuthorKey, Enabled: true,
	}); err != nil {
		t.Fatalf("SetAutomaticRuleEnabled() error = %v", err)
	}
	if _, err := module.ActivatePolicy(ctx, authorization.ActivatePolicyCommand{
		Actor: admin, Revision: draft.Number, ExpectedActiveRevision: 1,
	}); err != nil {
		t.Fatalf("ActivatePolicy() error = %v", err)
	}
	service := docsauthz.New(module)
	userContext := foundationauth.NewContext(ctx, &foundationauth.Principal{Subject: "registered-user"})
	access, err := service.EffectiveAccess(userContext)
	if err != nil {
		t.Fatalf("EffectiveAccess() error = %v", err)
	}
	found := false
	for _, grant := range access.Grants {
		found = found || grant.Role == docsauthz.RoleAuthor && grant.Source == authorization.GrantSourceAutomatic
	}
	if !found {
		t.Fatalf("EffectiveAccess() grants = %#v, want automatic author", access.Grants)
	}
}

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
		ID:   docsauthz.CollectionScopeID("collection-1"),
		Type: docsauthz.ScopeCollection, ParentID: docsauthz.RootScopeID,
	}); err != nil {
		t.Fatalf("CreateScope() collection error = %v", err)
	}
	if _, err := module.RegisterScope(ctx, authorization.RegisterScopeCommand{
		ID:   docsauthz.DocumentScopeID("doc-1"),
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
