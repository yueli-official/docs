package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type AuthorizationRoleView struct {
	ID           string   `json:"id"`
	Key          string   `json:"key"`
	DisplayName  string   `json:"displayName"`
	ScopeID      string   `json:"scopeId"`
	Kind         string   `json:"kind"`
	Status       string   `json:"status"`
	Protected    bool     `json:"protected"`
	Capabilities []string `json:"capabilities"`
	Sources      []string `json:"assignmentSources"`
}

type AuthorizationApplicationView struct {
	ID             string    `json:"id"`
	RequestGroupID string    `json:"requestGroupId"`
	Subject        string    `json:"subject"`
	Role           string    `json:"role"`
	ScopeID        string    `json:"scopeId"`
	Reason         string    `json:"reason"`
	State          string    `json:"state"`
	GrantID        string    `json:"grantId"`
	CreatedAt      time.Time `json:"createdAt"`
	ReviewedAt     time.Time `json:"reviewedAt,omitempty"`
	ReviewedBy     string    `json:"reviewedBy,omitempty"`
	ReviewReason   string    `json:"reviewReason,omitempty"`
}

type ListRequestableRolesReq struct {
	g.Meta `path:"/api/v1/authorization/requestable-roles" method:"get" tags:"authorization" summary:"Requestable roles"`
}
type ListRequestableRolesRes struct {
	Items []AuthorizationRoleView `json:"items"`
}

type ApplyForRoleReq struct {
	g.Meta         `path:"/api/v1/authorization/applications" method:"post" tags:"authorization" summary:"Apply for a role"`
	Role           string `json:"role" v:"required"`
	Reason         string `json:"reason" v:"length:0,2000"`
	RequestGroupID string `json:"requestGroupId" v:"length:0,128"`
	AbuseAttemptID string `json:"abuseAttemptId" v:"length:0,128"`
	ChallengeProof string `json:"challengeProof" v:"length:0,4096"`
}
type ApplyForRoleRes struct {
	Application AuthorizationApplicationView `json:"application"`
}

type ListMyApplicationsReq struct {
	g.Meta `path:"/api/v1/authorization/applications/mine" method:"get" tags:"authorization" summary:"My role applications"`
	State  string `json:"state" in:"query"`
	Offset int    `json:"offset" in:"query" d:"0"`
	Limit  int    `json:"limit" in:"query" d:"100"`
}
type ListMyApplicationsRes struct {
	Items []AuthorizationApplicationView `json:"items"`
	Total int                            `json:"total"`
}

type WithdrawRoleApplicationReq struct {
	g.Meta `path:"/api/v1/authorization/applications/{id}/withdraw" method:"post" tags:"authorization" summary:"Withdraw my role application"`
	ID     string `json:"id" in:"path" v:"required"`
}
type WithdrawRoleApplicationRes struct {
	Application AuthorizationApplicationView `json:"application"`
}

type ListRoleApplicationsReq struct {
	g.Meta `path:"/api/v1/authorization/manage/applications" method:"get" tags:"authorization" summary:"Review queue"`
	State  string `json:"state" in:"query"`
	Offset int    `json:"offset" in:"query" d:"0"`
	Limit  int    `json:"limit" in:"query" d:"100"`
}
type ListRoleApplicationsRes struct {
	Items []AuthorizationApplicationView `json:"items"`
	Total int                            `json:"total"`
}

type ReviewRoleApplicationReq struct {
	g.Meta   `path:"/api/v1/authorization/manage/applications/{id}/review" method:"post" tags:"authorization" summary:"Review a role application"`
	ID       string `json:"id" in:"path" v:"required"`
	Decision string `json:"decision" v:"required|in:approve,reject"`
	Reason   string `json:"reason" v:"length:0,2000"`
}
type ReviewRoleApplicationRes struct {
	Application AuthorizationApplicationView `json:"application"`
}

type AuthorizationPolicyView struct {
	Number      uint64    `json:"number"`
	Base        uint64    `json:"base"`
	State       string    `json:"state"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	ActivatedAt time.Time `json:"activatedAt,omitempty"`
}

type ListAuthorizationRolesReq struct {
	g.Meta         `path:"/api/v1/authorization/manage/roles" method:"get" tags:"authorization" summary:"List roles"`
	IncludeRetired bool `json:"includeRetired" in:"query"`
}
type ListAuthorizationRolesRes struct {
	Items []AuthorizationRoleView `json:"items"`
	Total int                     `json:"total"`
}

type ListAuthorizationPoliciesReq struct {
	g.Meta `path:"/api/v1/authorization/manage/policies" method:"get" tags:"authorization" summary:"List policy revisions"`
}
type ListAuthorizationPoliciesRes struct {
	Items []AuthorizationPolicyView `json:"items"`
	Total int                       `json:"total"`
}

type AuthorizationAutomaticRuleView struct {
	Key     string `json:"key"`
	Enabled bool   `json:"enabled"`
}

type AuthorizationGrantView struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	Role    string `json:"role"`
	Source  string `json:"source"`
}

type AuthorizationCapabilityView struct {
	Key         string `json:"key"`
	DisplayName string `json:"displayName"`
}

type GetAuthorizationConsoleReq struct {
	g.Meta `path:"/api/v1/authorization/manage/console" method:"get" tags:"authorization" summary:"Authorization management console"`
}
type GetAuthorizationConsoleRes struct {
	ActiveRevision uint64                           `json:"activeRevision"`
	Policy         AuthorizationPolicyView          `json:"policy"`
	Roles          []AuthorizationRoleView          `json:"roles"`
	AutomaticRules []AuthorizationAutomaticRuleView `json:"automaticRules"`
	Applications   []AuthorizationApplicationView   `json:"applications"`
	Grants         []AuthorizationGrantView         `json:"grants"`
	Capabilities   []AuthorizationCapabilityView    `json:"capabilities"`
}

type GrantAuthorizationRoleReq struct {
	g.Meta  `path:"/api/v1/authorization/manage/grants" method:"post" tags:"authorization" summary:"Grant a Docs role directly"`
	Subject string `json:"subject" v:"required|length:1,255"`
	Role    string `json:"role" v:"required"`
}
type GrantAuthorizationRoleRes struct {
	Grant AuthorizationGrantView `json:"grant"`
}

type RevokeAuthorizationGrantReq struct {
	g.Meta `path:"/api/v1/authorization/manage/grants/{id}" method:"delete" tags:"authorization" summary:"Revoke a Docs role grant"`
	ID     string `json:"id" in:"path" v:"required"`
}
type RevokeAuthorizationGrantRes struct {
	Grant AuthorizationGrantView `json:"grant"`
}

type GetAuthorizationPolicyReq struct {
	g.Meta   `path:"/api/v1/authorization/manage/policies/{revision}" method:"get" tags:"authorization" summary:"Get policy revision"`
	Revision uint64 `json:"revision" in:"path" v:"required|min:1"`
}
type GetAuthorizationPolicyRes struct {
	Policy         AuthorizationPolicyView          `json:"policy"`
	Roles          []AuthorizationRoleView          `json:"roles"`
	AutomaticRules []AuthorizationAutomaticRuleView `json:"automaticRules"`
}

type CreateAuthorizationPolicyDraftReq struct {
	g.Meta                 `path:"/api/v1/authorization/manage/policies/drafts" method:"post" tags:"authorization" summary:"Create policy draft"`
	ExpectedActiveRevision uint64 `json:"expectedActiveRevision" v:"required|min:1"`
}
type CreateAuthorizationPolicyDraftRes struct {
	Policy AuthorizationPolicyView `json:"policy"`
}

type SetAuthorizationRoleCapabilitiesReq struct {
	g.Meta       `path:"/api/v1/authorization/manage/policies/{revision}/roles/{role}/capabilities" method:"put" tags:"authorization" summary:"Set role capabilities"`
	Revision     uint64   `json:"revision" in:"path" v:"required|min:1"`
	Role         string   `json:"role" in:"path" v:"required"`
	Capabilities []string `json:"capabilities"`
}
type SetAuthorizationRoleCapabilitiesRes struct {
	Policy AuthorizationPolicyView `json:"policy"`
}

type CreateAuthorizationRoleReq struct {
	g.Meta       `path:"/api/v1/authorization/manage/policies/{revision}/roles" method:"post" tags:"authorization" summary:"Create custom role"`
	Revision     uint64   `json:"revision" in:"path" v:"required|min:1"`
	Key          string   `json:"key" v:"required"`
	DisplayName  string   `json:"displayName" v:"required"`
	Capabilities []string `json:"capabilities"`
	Sources      []string `json:"assignmentSources"`
}
type CreateAuthorizationRoleRes struct {
	Role AuthorizationRoleView `json:"role"`
}

type RetireAuthorizationRoleReq struct {
	g.Meta   `path:"/api/v1/authorization/manage/policies/{revision}/roles/{role}/retire" method:"post" tags:"authorization" summary:"Retire custom role"`
	Revision uint64 `json:"revision" in:"path" v:"required|min:1"`
	Role     string `json:"role" in:"path" v:"required"`
}
type RetireAuthorizationRoleRes struct {
	Role AuthorizationRoleView `json:"role"`
}

type SetAuthorizationAutomaticRuleReq struct {
	g.Meta   `path:"/api/v1/authorization/manage/policies/{revision}/automatic/{rule}" method:"put" tags:"authorization" summary:"Configure automatic role rule"`
	Revision uint64 `json:"revision" in:"path" v:"required|min:1"`
	Rule     string `json:"rule" in:"path" v:"required"`
	Enabled  bool   `json:"enabled"`
}
type SetAuthorizationAutomaticRuleRes struct {
	Policy AuthorizationPolicyView `json:"policy"`
}

type ValidateAuthorizationPolicyReq struct {
	g.Meta   `path:"/api/v1/authorization/manage/policies/{revision}/validate" method:"post" tags:"authorization" summary:"Validate policy draft"`
	Revision uint64 `json:"revision" in:"path" v:"required|min:1"`
}
type ValidateAuthorizationPolicyRes struct {
	Valid      bool     `json:"valid"`
	Violations []string `json:"violations"`
}

type PreviewAuthorizationPolicyReq struct {
	g.Meta   `path:"/api/v1/authorization/manage/policies/{revision}/preview" method:"post" tags:"authorization" summary:"Preview policy impact"`
	Revision uint64 `json:"revision" in:"path" v:"required|min:1"`
}
type PreviewAuthorizationPolicyRes struct {
	AddedBindings   int `json:"addedBindings"`
	RemovedBindings int `json:"removedBindings"`
}

type ActivateAuthorizationPolicyReq struct {
	g.Meta                 `path:"/api/v1/authorization/manage/policies/{revision}/activate" method:"post" tags:"authorization" summary:"Activate policy draft"`
	Revision               uint64 `json:"revision" in:"path" v:"required|min:1"`
	ExpectedActiveRevision uint64 `json:"expectedActiveRevision" v:"required|min:1"`
}
type ActivateAuthorizationPolicyRes struct {
	Policy AuthorizationPolicyView `json:"policy"`
}
