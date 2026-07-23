package controller

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/google/uuid"
	"github.com/yueli-official/foundation/go/abuse"
	"github.com/yueli-official/foundation/go/authorization"

	v1 "platform/products/docs/api/api/v1"
	"platform/products/docs/api/internal/docsabuse"
	"platform/products/docs/api/internal/docsauthz"
	"platform/products/docs/api/internal/docserr"
)

type Authorization struct{}

func NewAuthorization() *Authorization { return &Authorization{} }

func (controller *Authorization) ListRequestableRoles(
	ctx context.Context,
	_ *v1.ListRequestableRolesReq,
) (*v1.ListRequestableRolesRes, error) {
	service := authorizationService(ctx)
	roles, err := service.Runtime().ListRequestableRoles(ctx, authorization.RequestableRoleQuery{
		Subject: service.Subject(ctx), ScopeID: docsauthz.RootScopeID,
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	items := make([]v1.AuthorizationRoleView, len(roles))
	for index, role := range roles {
		items[index] = authorizationRoleView(role)
	}
	return &v1.ListRequestableRolesRes{Items: items}, nil
}

func (controller *Authorization) ApplyForRole(
	ctx context.Context,
	req *v1.ApplyForRoleReq,
) (*v1.ApplyForRoleRes, error) {
	service := authorizationService(ctx)
	attemptID := strings.TrimSpace(req.AbuseAttemptID)
	if attemptID == "" {
		attemptID = strings.TrimSpace(idempotencyKeyOf(ctx))
	}
	if attemptID == "" {
		attemptID = uuid.NewString()
	}
	if action := service.AuthorApplicationAction(); action != nil {
		request := ghttp.RequestFromCtx(ctx)
		if request == nil {
			return nil, docserr.AbuseUnavailable()
		}
		network, err := docsabuse.NetworkPrefix(request.GetClientIp())
		if err != nil {
			return nil, docserr.AbuseUnavailable()
		}
		input := abuse.Input{
			ID: abuse.AttemptID(attemptID),
			Signals: abuse.Signals{
				Network: network,
				Actor:   service.Subject(ctx).ID,
			},
		}
		if proof := strings.TrimSpace(req.ChallengeProof); proof != "" {
			input.Proof = &abuse.Proof{Kind: "turnstile", Token: proof}
		}
		admission, err := action.Admit(ctx, input)
		if err != nil {
			if abuse.IsKind(err, abuse.ErrorConflict) {
				return nil, docserr.AbuseAttemptReplayed()
			}
			return nil, docserr.AbuseUnavailable()
		}
		switch admission.Disposition {
		case abuse.DispositionAllow:
			if admission.Replay {
				return nil, docserr.AbuseAttemptReplayed()
			}
		case abuse.DispositionChallenge:
			return nil, docserr.ChallengeRequired(attemptID)
		default:
			return nil, docserr.RateLimited()
		}
	}
	application, err := service.Runtime().Apply(ctx, authorization.ApplyCommand{
		Actor: service.Subject(ctx), Role: authorization.RoleKey(req.Role),
		ScopeID: docsauthz.RootScopeID, Reason: req.Reason,
		RequestGroupID: authorization.RequestGroupID(req.RequestGroupID),
		IdempotencyKey: idempotencyKeyOf(ctx),
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	return &v1.ApplyForRoleRes{Application: authorizationApplicationView(application)}, nil
}

func (controller *Authorization) ListMyApplications(
	ctx context.Context,
	req *v1.ListMyApplicationsReq,
) (*v1.ListMyApplicationsRes, error) {
	service := authorizationService(ctx)
	subject := service.Subject(ctx)
	page, err := service.Runtime().ListApplications(ctx, authorization.ApplicationListQuery{
		Actor: subject, Subject: subject, ScopeID: docsauthz.RootScopeID,
		State: authorization.ApplicationState(req.State), Offset: req.Offset, Limit: req.Limit,
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	return &v1.ListMyApplicationsRes{
		Items: authorizationApplicationViews(page.Applications), Total: page.Total,
	}, nil
}

func (controller *Authorization) WithdrawRoleApplication(
	ctx context.Context,
	req *v1.WithdrawRoleApplicationReq,
) (*v1.WithdrawRoleApplicationRes, error) {
	service := authorizationService(ctx)
	application, err := service.Runtime().WithdrawApplication(ctx, authorization.WithdrawApplicationCommand{
		Actor: service.Subject(ctx), ApplicationID: authorization.ApplicationID(req.ID),
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	return &v1.WithdrawRoleApplicationRes{Application: authorizationApplicationView(application)}, nil
}

func (controller *Authorization) ListRoleApplications(
	ctx context.Context,
	req *v1.ListRoleApplicationsReq,
) (*v1.ListRoleApplicationsRes, error) {
	service := authorizationService(ctx)
	page, err := service.Runtime().ListApplications(ctx, authorization.ApplicationListQuery{
		Actor: service.Subject(ctx), ScopeID: docsauthz.RootScopeID,
		State: authorization.ApplicationState(req.State), Offset: req.Offset, Limit: req.Limit,
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	return &v1.ListRoleApplicationsRes{
		Items: authorizationApplicationViews(page.Applications), Total: page.Total,
	}, nil
}

func (controller *Authorization) ReviewRoleApplication(
	ctx context.Context,
	req *v1.ReviewRoleApplicationReq,
) (*v1.ReviewRoleApplicationRes, error) {
	service := authorizationService(ctx)
	application, err := service.Runtime().ReviewApplication(ctx, authorization.ReviewApplicationCommand{
		Actor: service.Subject(ctx), ApplicationID: authorization.ApplicationID(req.ID),
		Decision: authorization.ReviewDecision(req.Decision), Reason: req.Reason,
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	return &v1.ReviewRoleApplicationRes{Application: authorizationApplicationView(application)}, nil
}

func (controller *Authorization) ListAuthorizationRoles(
	ctx context.Context,
	req *v1.ListAuthorizationRolesReq,
) (*v1.ListAuthorizationRolesRes, error) {
	service := authorizationService(ctx)
	page, err := service.Runtime().ListRoles(ctx, authorization.RoleListQuery{
		Actor: service.Subject(ctx), ScopeID: docsauthz.RootScopeID,
		IncludeRetired: req.IncludeRetired, Limit: 500,
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	items := make([]v1.AuthorizationRoleView, len(page.Roles))
	for index, role := range page.Roles {
		items[index] = authorizationRoleView(role)
	}
	return &v1.ListAuthorizationRolesRes{Items: items, Total: page.Total}, nil
}

func (controller *Authorization) ListAuthorizationPolicies(
	ctx context.Context,
	_ *v1.ListAuthorizationPoliciesReq,
) (*v1.ListAuthorizationPoliciesRes, error) {
	service := authorizationService(ctx)
	page, err := service.Runtime().ListPolicyRevisions(ctx, authorization.PolicyRevisionListQuery{
		Actor: service.Subject(ctx), ScopeID: docsauthz.RootScopeID, Limit: 500,
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	items := make([]v1.AuthorizationPolicyView, len(page.Revisions))
	for index, revision := range page.Revisions {
		items[index] = authorizationPolicyView(revision)
	}
	return &v1.ListAuthorizationPoliciesRes{Items: items, Total: page.Total}, nil
}

func (controller *Authorization) GetAuthorizationPolicy(
	ctx context.Context,
	req *v1.GetAuthorizationPolicyReq,
) (*v1.GetAuthorizationPolicyRes, error) {
	service := authorizationService(ctx)
	snapshot, err := service.Runtime().GetPolicySnapshot(ctx, authorization.PolicySnapshotQuery{
		Actor: service.Subject(ctx), Revision: req.Revision, ScopeID: docsauthz.RootScopeID,
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	roles := make([]v1.AuthorizationRoleView, len(snapshot.Roles))
	for index, role := range snapshot.Roles {
		roles[index] = authorizationRoleView(role)
	}
	rules := make([]v1.AuthorizationAutomaticRuleView, len(snapshot.AutomaticRules))
	for index, rule := range snapshot.AutomaticRules {
		rules[index] = v1.AuthorizationAutomaticRuleView{Key: rule.Key, Enabled: rule.Enabled}
	}
	return &v1.GetAuthorizationPolicyRes{
		Policy: authorizationPolicyView(snapshot.Revision), Roles: roles, AutomaticRules: rules,
	}, nil
}

func (controller *Authorization) CreateAuthorizationPolicyDraft(
	ctx context.Context,
	req *v1.CreateAuthorizationPolicyDraftReq,
) (*v1.CreateAuthorizationPolicyDraftRes, error) {
	service := authorizationService(ctx)
	revision, err := service.Runtime().CreatePolicyDraft(ctx, authorization.CreatePolicyDraftCommand{
		Actor: service.Subject(ctx), ScopeID: docsauthz.RootScopeID,
		ExpectedActiveRevision: req.ExpectedActiveRevision,
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	return &v1.CreateAuthorizationPolicyDraftRes{Policy: authorizationPolicyView(revision)}, nil
}

func (controller *Authorization) SetAuthorizationRoleCapabilities(
	ctx context.Context,
	req *v1.SetAuthorizationRoleCapabilitiesReq,
) (*v1.SetAuthorizationRoleCapabilitiesRes, error) {
	service := authorizationService(ctx)
	revision, err := service.Runtime().SetRoleCapabilities(ctx, authorization.SetRoleCapabilitiesCommand{
		Actor: service.Subject(ctx), Revision: req.Revision, Role: authorization.RoleKey(req.Role),
		Capabilities: capabilityKeys(req.Capabilities),
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	return &v1.SetAuthorizationRoleCapabilitiesRes{Policy: authorizationPolicyView(revision)}, nil
}

func (controller *Authorization) CreateAuthorizationRole(
	ctx context.Context,
	req *v1.CreateAuthorizationRoleReq,
) (*v1.CreateAuthorizationRoleRes, error) {
	service := authorizationService(ctx)
	role, err := service.Runtime().CreateRole(ctx, authorization.CreateRoleCommand{
		Actor: service.Subject(ctx), Revision: req.Revision,
		Key: authorization.RoleKey(req.Key), DisplayName: req.DisplayName,
		ScopeID: docsauthz.RootScopeID, Capabilities: capabilityKeys(req.Capabilities),
		Assignment: authorization.AssignmentPolicy{Sources: grantSources(req.Sources)},
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	return &v1.CreateAuthorizationRoleRes{Role: authorizationRoleView(role)}, nil
}

func (controller *Authorization) RetireAuthorizationRole(
	ctx context.Context,
	req *v1.RetireAuthorizationRoleReq,
) (*v1.RetireAuthorizationRoleRes, error) {
	service := authorizationService(ctx)
	role, err := service.Runtime().RetireRole(ctx, authorization.RetireRoleCommand{
		Actor: service.Subject(ctx), Revision: req.Revision, Role: authorization.RoleKey(req.Role),
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	return &v1.RetireAuthorizationRoleRes{Role: authorizationRoleView(role)}, nil
}

func (controller *Authorization) SetAuthorizationAutomaticRule(
	ctx context.Context,
	req *v1.SetAuthorizationAutomaticRuleReq,
) (*v1.SetAuthorizationAutomaticRuleRes, error) {
	service := authorizationService(ctx)
	revision, err := service.Runtime().SetAutomaticRuleEnabled(ctx, authorization.SetAutomaticRuleEnabledCommand{
		Actor: service.Subject(ctx), Revision: req.Revision, Rule: req.Rule, Enabled: req.Enabled,
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	return &v1.SetAuthorizationAutomaticRuleRes{Policy: authorizationPolicyView(revision)}, nil
}

func (controller *Authorization) ValidateAuthorizationPolicy(
	ctx context.Context,
	req *v1.ValidateAuthorizationPolicyReq,
) (*v1.ValidateAuthorizationPolicyRes, error) {
	service := authorizationService(ctx)
	validation, err := service.Runtime().ValidatePolicy(ctx, authorization.ValidatePolicyCommand{
		Actor: service.Subject(ctx), Revision: req.Revision,
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	return &v1.ValidateAuthorizationPolicyRes{
		Valid: validation.Valid, Violations: validation.Violations,
	}, nil
}

func (controller *Authorization) PreviewAuthorizationPolicy(
	ctx context.Context,
	req *v1.PreviewAuthorizationPolicyReq,
) (*v1.PreviewAuthorizationPolicyRes, error) {
	service := authorizationService(ctx)
	impact, err := service.Runtime().PreviewPolicy(ctx, authorization.PreviewPolicyCommand{
		Actor: service.Subject(ctx), Revision: req.Revision,
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	return &v1.PreviewAuthorizationPolicyRes{
		AddedBindings: impact.AddedBindings, RemovedBindings: impact.RemovedBindings,
	}, nil
}

func (controller *Authorization) ActivateAuthorizationPolicy(
	ctx context.Context,
	req *v1.ActivateAuthorizationPolicyReq,
) (*v1.ActivateAuthorizationPolicyRes, error) {
	service := authorizationService(ctx)
	revision, err := service.Runtime().ActivatePolicy(ctx, authorization.ActivatePolicyCommand{
		Actor: service.Subject(ctx), Revision: req.Revision,
		ExpectedActiveRevision: req.ExpectedActiveRevision,
	})
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	return &v1.ActivateAuthorizationPolicyRes{Policy: authorizationPolicyView(revision)}, nil
}

func authorizationRoleView(role authorization.Role) v1.AuthorizationRoleView {
	capabilities := make([]string, len(role.Capabilities))
	for index, capability := range role.Capabilities {
		capabilities[index] = string(capability)
	}
	sources := make([]string, len(role.Assignment.Sources))
	for index, source := range role.Assignment.Sources {
		sources[index] = string(source)
	}
	return v1.AuthorizationRoleView{
		ID: string(role.ID), Key: string(role.Key), DisplayName: role.DisplayName,
		ScopeID: string(role.ScopeID), Kind: string(role.Kind), Status: string(role.Status),
		Protected: role.Protected, Capabilities: capabilities, Sources: sources,
	}
}

func authorizationPolicyView(revision authorization.PolicyRevision) v1.AuthorizationPolicyView {
	return v1.AuthorizationPolicyView{
		Number: revision.Number, Base: revision.Base, State: string(revision.State),
		CreatedBy: revision.CreatedBy.ID, CreatedAt: revision.CreatedAt, ActivatedAt: revision.ActivatedAt,
	}
}

func capabilityKeys(values []string) []authorization.CapabilityKey {
	keys := make([]authorization.CapabilityKey, len(values))
	for index, value := range values {
		keys[index] = authorization.CapabilityKey(value)
	}
	return keys
}

func grantSources(values []string) []authorization.GrantSource {
	sources := make([]authorization.GrantSource, len(values))
	for index, value := range values {
		sources[index] = authorization.GrantSource(value)
	}
	return sources
}

func authorizationApplicationViews(applications []authorization.Application) []v1.AuthorizationApplicationView {
	items := make([]v1.AuthorizationApplicationView, len(applications))
	for index, application := range applications {
		items[index] = authorizationApplicationView(application)
	}
	return items
}

func authorizationApplicationView(application authorization.Application) v1.AuthorizationApplicationView {
	return v1.AuthorizationApplicationView{
		ID: string(application.ID), RequestGroupID: string(application.RequestGroupID),
		Subject: application.Subject.ID, Role: string(application.Role), ScopeID: string(application.ScopeID),
		Reason: application.Reason, State: string(application.State), GrantID: string(application.GrantID),
		CreatedAt: application.CreatedAt, ReviewedAt: application.ReviewedAt,
		ReviewedBy: application.ReviewedBy.ID, ReviewReason: application.ReviewReason,
	}
}
