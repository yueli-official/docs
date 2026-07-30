package docsauthz

import (
	"context"
	"database/sql"

	"github.com/yueli-official/docs/api/internal/docsabuse"
	"github.com/yueli-official/foundation/go/abuse"
	foundationauth "github.com/yueli-official/foundation/go/auth"
	"github.com/yueli-official/foundation/go/authorization"
)

type Runtime interface {
	authorization.Authorizer
	authorization.QueryPlanner
	authorization.AccessReader
	authorization.ScopeManager
	authorization.ResourceScopeRegistry
	authorization.RoleManager
	authorization.RoleReader
	authorization.GrantManager
	authorization.GrantReader
	authorization.WorkflowManager
	authorization.WorkflowReader
	authorization.PolicyManager
	authorization.PolicyReader
	authorization.Reconciler
}

type Service struct {
	runtime Runtime
	db      *sql.DB
	abuse   docsabuse.Actions
}

func New(runtime Runtime) *Service {
	return &Service{runtime: runtime}
}

func NewWithDB(runtime Runtime, db *sql.DB) *Service {
	return &Service{runtime: runtime, db: db}
}

func (service *Service) SetAbuse(module abuse.Module) error {
	actions, err := docsabuse.Bind(module)
	if err != nil {
		return err
	}
	service.abuse = actions
	return nil
}

func (service *Service) AuthorApplicationAction() abuse.Action {
	if service == nil {
		return nil
	}
	return service.abuse.AuthorApplication
}

func (service *Service) Runtime() Runtime {
	if service == nil {
		return nil
	}
	return service.runtime
}

func (service *Service) Subject(ctx context.Context) authorization.SubjectRef {
	principal, ok := foundationauth.FromContext(ctx)
	if !ok {
		return authorization.SubjectRef{Kind: authorization.SubjectAnonymous}
	}
	if principal.Subject != "" {
		return authorization.SubjectRef{Kind: authorization.SubjectUser, ID: principal.Subject}
	}
	if principal.ClientID != "" {
		return authorization.SubjectRef{Kind: authorization.SubjectService, ID: principal.ClientID}
	}
	return authorization.SubjectRef{Kind: authorization.SubjectAnonymous}
}

func (service *Service) Decide(
	ctx context.Context,
	capability authorization.CapabilityKey,
	scopeID authorization.ScopeID,
	resource authorization.ResourceFacts,
) (authorization.Decision, error) {
	if service == nil || service.runtime == nil {
		return authorization.Decision{}, &authorization.Error{
			Kind: authorization.ErrorUnavailable, Field: "runtime", Message: "is not configured",
		}
	}
	if err := service.ReconcileSubject(ctx); err != nil {
		return authorization.Decision{}, err
	}
	return service.runtime.Decide(ctx, authorization.DecisionRequest{
		Subject: service.Subject(ctx), Capability: capability, ScopeID: scopeID, Resource: resource,
	})
}

func (service *Service) EnsureCollectionScope(
	ctx context.Context,
	collectionID string,
) error {
	return ensureScope(ctx, service.runtime, authorization.RegisterScopeCommand{
		ID: CollectionScopeID(collectionID), Type: ScopeCollection, ParentID: RootScopeID,
	})
}

func (service *Service) EnsureDocumentScope(
	ctx context.Context,
	documentID string,
	collectionID string,
) error {
	return ensureScope(ctx, service.runtime, authorization.RegisterScopeCommand{
		ID: DocumentScopeID(documentID), Type: ScopeDocument, ParentID: CollectionScopeID(collectionID),
	})
}

func DocumentResource(id, ownerSub string) authorization.ResourceFacts {
	resource := authorization.ResourceFacts{
		Type: authorization.ResourceType("document"),
		ID:   authorization.ResourceID(id), ScopeID: DocumentScopeID(id),
	}
	if ownerSub != "" {
		resource.Relations = map[authorization.RelationKind][]authorization.SubjectRef{
			RelationOwner: {{Kind: authorization.SubjectUser, ID: ownerSub}},
		}
	}
	return resource
}

func (service *Service) SyncCatalogScopes(ctx context.Context) error {
	if service == nil || service.runtime == nil {
		return &authorization.Error{
			Kind: authorization.ErrorUnavailable, Field: "scope_sync", Message: "is not configured",
		}
	}
	if service.db == nil {
		return nil
	}
	return SyncResourceScopes(ctx, service.db, service.runtime)
}

func (service *Service) IsAdministrator(ctx context.Context) bool {
	if service == nil || service.runtime == nil {
		return false
	}
	decision, err := service.Decide(ctx, authorization.CapabilityManage, RootScopeID, authorization.ResourceFacts{})
	return err == nil && decision.Allowed
}

func (service *Service) EffectiveAccess(ctx context.Context) (authorization.EffectiveAccess, error) {
	if service == nil || service.runtime == nil {
		return authorization.EffectiveAccess{}, &authorization.Error{
			Kind: authorization.ErrorUnavailable, Field: "runtime", Message: "is not configured",
		}
	}
	if err := service.ReconcileSubject(ctx); err != nil {
		return authorization.EffectiveAccess{}, err
	}
	return service.runtime.EffectiveAccess(ctx, authorization.EffectiveAccessQuery{
		Subject: service.Subject(ctx), ScopeID: RootScopeID,
	})
}

func (service *Service) ReconcileSubject(ctx context.Context) error {
	if service == nil || service.runtime == nil {
		return &authorization.Error{
			Kind: authorization.ErrorUnavailable, Field: "runtime", Message: "is not configured",
		}
	}
	subject := service.Subject(ctx)
	if subject.Kind != authorization.SubjectUser || subject.ID == "" {
		return nil
	}
	preview, err := service.runtime.PreviewReconcileSubject(ctx, authorization.ReconcileSubjectCommand{
		Subject: subject,
	})
	if err != nil || preview.Created == 0 {
		return err
	}
	_, err = service.runtime.ReconcileSubject(ctx, authorization.ReconcileSubjectCommand{
		Subject: subject,
	})
	return err
}
