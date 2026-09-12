package controller

import (
	"context"
	"strings"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
	v1 "github.com/yueli-official/docs/api/api/v1"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/docserr"
	foundationauth "github.com/yueli-official/foundation/go/auth"
	"github.com/yueli-official/foundation/go/authorization"
)

type PersonalPermissions struct {
	site    string
	service *docsauthz.Service
}

// The product owns this directory. Identity only displays capabilities that the
// account currently has; route and resource checks still run on every request.
var personalPermissions = []foundationauth.PersonalPermission{
	{Key: string(docsauthz.CapabilityCollectionManage), Label: "管理文档集", Description: "创建、编辑和删除有权管理的文档集，上传封面、管理语言及克隆发布版本。"},
	{Key: string(docsauthz.CapabilityVersionManage), Label: "管理文档版本", Description: "创建、编辑文档集内部版本，设置默认版本及发布状态；仍需当前账号的版本管理权限。"},
	{Key: string(docsauthz.CapabilityDocumentRead), Label: "读取文档", Description: "读取有权访问的文档、草稿和版本列表。"},
	{Key: string(docsauthz.CapabilityDocumentCreate), Label: "创建文档", Description: "在有权投稿的文档集中创建目录和草稿，并上传正文图片。"},
	{Key: string(docsauthz.CapabilityDocumentUpdate), Label: "编辑文档", Description: "编辑有权操作的文档、目录顺序和正文图片；发布或下架另需相应权限。"},
	{Key: string(docsauthz.CapabilityDocumentPublish), Label: "发布文档", Description: "发布有权操作的文档，不包含编辑正文权限。"},
	{Key: string(docsauthz.CapabilityDocumentArchive), Label: "下架文档", Description: "下架有权操作的文档并保留内容。"},
	{Key: string(docsauthz.CapabilityDocumentDeletePermanently), Label: "永久删除文档", Description: "永久删除文档及其子文档，无法恢复；仍需当前账号的永久删除权限。"},
	{Key: string(docsauthz.CapabilityImportManage), Label: "导入文档", Description: "预检、导入并发布文档包及图片，完整镜像模式还会下架缺失页面；仍需当前站点导入权限。"},
}

func NewPersonalPermissions(site string, service *docsauthz.Service) *PersonalPermissions {
	return &PersonalPermissions{site, service}
}
func (c *PersonalPermissions) GetPersonalPermissions(ctx context.Context, req *v1.PersonalPermissionsReq) (*v1.PersonalPermissionsRes, error) {
	p, _ := foundationauth.FromContext(ctx)
	if p == nil || p.SubjectKind != foundationauth.SubjectClient || p.ClientID != "identity-svc" || !p.HasScope(foundationauth.PersonalPermissionsScope) || c.site == "" {
		return nil, docserr.Forbidden()
	}
	if c.service == nil || c.service.Runtime() == nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	userCtx := foundationauth.NewContext(ctx, &foundationauth.Principal{Subject: req.UserKey, SubjectKind: foundationauth.SubjectUser})
	// Subject resolution accepts the typed principal, not user-controlled role claims.
	if err := c.service.ReconcileSubject(userCtx); err != nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	access, err := c.service.Runtime().EffectiveAccess(userCtx, authorization.EffectiveAccessQuery{Subject: authorization.SubjectRef{Kind: authorization.SubjectUser, ID: req.UserKey}, ScopeID: docsauthz.RootScopeID, IncludeDescendants: true})
	if err != nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	items := []foundationauth.PersonalPermission{}
	for _, permission := range personalPermissions {
		for _, key := range access.Capabilities {
			if permission.Key == string(key) {
				items = append(items, permission)
				break
			}
		}
	}
	return &v1.PersonalPermissionsRes{Site: c.site, UserKey: req.UserKey, Items: items}, nil
}
func (c *PersonalPermissions) AuthorizePersonalMedia(ctx context.Context, _ *v1.PersonalMediaAuthorizationReq) (*v1.PersonalMediaAuthorizationRes, error) {
	p, _ := foundationauth.FromContext(ctx)
	if p == nil || !p.IsPersonalToken() || c.site == "" || p.ClientID != c.site {
		return nil, docserr.Forbidden()
	}
	if c.service == nil || c.service.Runtime() == nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	if err := c.service.ReconcileSubject(ctx); err != nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	access, err := c.service.Runtime().EffectiveAccess(ctx, authorization.EffectiveAccessQuery{
		Subject: c.service.Subject(ctx), ScopeID: docsauthz.RootScopeID, IncludeDescendants: true,
	})
	if err != nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	// Asset receives only upload profiles, never content-management capabilities.
	// The Docs upload endpoints separately authorize the concrete collection/doc.
	profiles := []struct {
		profile      string
		capabilities []authorization.CapabilityKey
	}{
		{"docs-import-image", []authorization.CapabilityKey{docsauthz.CapabilityImportManage}},
		{"docs-collection-cover", []authorization.CapabilityKey{docsauthz.CapabilityCollectionManage}},
		{"docs-content-image", []authorization.CapabilityKey{docsauthz.CapabilityDocumentCreate, docsauthz.CapabilityDocumentUpdate}},
	}
	out := &v1.PersonalMediaAuthorizationRes{UserKey: p.Subject, Scopes: []string{}}
	for _, profile := range profiles {
		allowed := false
		for _, required := range profile.capabilities {
			for _, current := range access.Capabilities {
				allowed = allowed || current == required && foundationauth.AllowsPersonalCapability(ctx, string(required))
			}
		}
		if allowed {
			scope, err := foundationauth.PersonalScope(c.site, "asset.profile."+profile.profile+".upload")
			if err != nil {
				return nil, docserr.AuthorizationUnavailable()
			}
			out.Scopes = append(out.Scopes, scope)
		}
	}
	if len(out.Scopes) == 0 {
		return nil, docserr.Forbidden()
	}
	if !p.ExpiresAt.IsZero() {
		out.ExpiresAt = p.ExpiresAt.Format(time.RFC3339)
	}
	return out, nil
}
func PersonalTokenRoutes(r *ghttp.Request) {
	p, _ := foundationauth.FromContext(r.Context())
	if p != nil && p.IsPersonalToken() {
		if !allowsPersonalRoute(r.Context(), r.Method, r.URL.Path) {
			r.SetError(docserr.Forbidden())
			return
		}
	}
	r.Middleware.Next()
}

type personalRoute struct {
	method, path string
	capabilities []authorization.CapabilityKey // Any selected capability; controllers enforce resource rights.
}

var personalRoutes = []personalRoute{
	{"POST", "/api/v1/collections", []authorization.CapabilityKey{docsauthz.CapabilityCollectionManage}},
	{"PATCH", "/api/v1/collections/{id}", []authorization.CapabilityKey{docsauthz.CapabilityCollectionManage}},
	{"DELETE", "/api/v1/collections/{id}", []authorization.CapabilityKey{docsauthz.CapabilityCollectionManage}},
	{"POST", "/api/v1/collections/{id}/cover", []authorization.CapabilityKey{docsauthz.CapabilityCollectionManage}},
	{"POST", "/api/v1/collections/{id}/cover/finalize", []authorization.CapabilityKey{docsauthz.CapabilityCollectionManage}},
	{"GET", "/api/v1/collections/{id}/versions", []authorization.CapabilityKey{docsauthz.CapabilityVersionManage, docsauthz.CapabilityCollectionManage, docsauthz.CapabilityDocumentRead, docsauthz.CapabilityDocumentCreate}},
	{"POST", "/api/v1/collections/{id}/versions", []authorization.CapabilityKey{docsauthz.CapabilityVersionManage}},
	{"PATCH", "/api/v1/manage/collections/{id}/versions/{versionId}", []authorization.CapabilityKey{docsauthz.CapabilityVersionManage}},
	{"POST", "/api/v1/manage/collections/{id}/clone-release", []authorization.CapabilityKey{docsauthz.CapabilityCollectionManage}},
	{"POST", "/api/v1/manage/collections/{id}/release", []authorization.CapabilityKey{docsauthz.CapabilityCollectionManage}},
	{"GET", "/api/v1/manage/collections/{slug}/tree", []authorization.CapabilityKey{docsauthz.CapabilityCollectionManage}},
	{"GET", "/api/v1/manage/collections/{id}/locales", []authorization.CapabilityKey{docsauthz.CapabilityCollectionManage}},
	{"POST", "/api/v1/manage/collections/{id}/locales", []authorization.CapabilityKey{docsauthz.CapabilityCollectionManage}},
	{"DELETE", "/api/v1/manage/collections/{id}/locales/{locale}", []authorization.CapabilityKey{docsauthz.CapabilityCollectionManage}},
	{"POST", "/api/v1/manage/collections/{id}/locales/clone", []authorization.CapabilityKey{docsauthz.CapabilityCollectionManage}},
	{"GET", "/api/v1/docs", []authorization.CapabilityKey{docsauthz.CapabilityDocumentRead}},
	{"GET", "/api/v1/manage/docs", []authorization.CapabilityKey{docsauthz.CapabilityDocumentRead}},
	{"GET", "/api/v1/docs/{id}", []authorization.CapabilityKey{docsauthz.CapabilityDocumentRead}},
	{"POST", "/api/v1/docs", []authorization.CapabilityKey{docsauthz.CapabilityDocumentCreate}},
	{"PATCH", "/api/v1/docs/{id}", []authorization.CapabilityKey{docsauthz.CapabilityDocumentUpdate}},
	{"POST", "/api/v1/docs/{id}/publish", []authorization.CapabilityKey{docsauthz.CapabilityDocumentPublish}},
	{"POST", "/api/v1/docs/{id}/archive", []authorization.CapabilityKey{docsauthz.CapabilityDocumentArchive}},
	{"DELETE", "/api/v1/docs/{id}", []authorization.CapabilityKey{docsauthz.CapabilityDocumentDeletePermanently}},
	{"POST", "/api/v1/images", []authorization.CapabilityKey{docsauthz.CapabilityDocumentCreate, docsauthz.CapabilityDocumentUpdate}},
	{"POST", "/api/v1/images/finalize", []authorization.CapabilityKey{docsauthz.CapabilityDocumentCreate, docsauthz.CapabilityDocumentUpdate}},
	{"GET", "/api/v1/imports/docs", []authorization.CapabilityKey{docsauthz.CapabilityImportManage}},
	{"POST", "/api/v1/imports/docs", []authorization.CapabilityKey{docsauthz.CapabilityImportManage}},
	{"GET", "/api/v1/imports/docs/{id}", []authorization.CapabilityKey{docsauthz.CapabilityImportManage}},
	{"POST", "/api/v1/imports/docs/{id}/confirm", []authorization.CapabilityKey{docsauthz.CapabilityImportManage}},
	{"POST", "/api/v1/personal-token/media-authorization", []authorization.CapabilityKey{docsauthz.CapabilityImportManage, docsauthz.CapabilityCollectionManage, docsauthz.CapabilityDocumentCreate, docsauthz.CapabilityDocumentUpdate}},
}

func allowsPersonalRoute(ctx context.Context, method, path string) bool {
	parts := strings.Split(path, "/")
	for _, route := range personalRoutes {
		if route.method != method {
			continue
		}
		pattern := strings.Split(route.path, "/")
		if len(parts) != len(pattern) {
			continue
		}
		matches := true
		for i, part := range pattern {
			if strings.HasPrefix(part, "{") {
				matches = matches && parts[i] != "" && parts[i] != "." && parts[i] != ".."
			} else {
				matches = matches && parts[i] == part
			}
		}
		if matches {
			for _, capability := range route.capabilities {
				if foundationauth.AllowsPersonalCapability(ctx, string(capability)) {
					return true
				}
			}
		}
	}
	return false
}
func ImportAuthorizationGuard(verifier *foundationauth.PersonalTokenVerifier, service *docsauthz.Service) func(context.Context, string) error {
	return func(ctx context.Context, bearer string) error {
		p, _ := foundationauth.FromContext(ctx)
		if p != nil && p.IsPersonalToken() {
			if verifier == nil {
				return docserr.AuthorizationUnavailable()
			}
			fresh, err := verifier.Verify(ctx, bearer)
			if err != nil {
				return docserr.Forbidden()
			}
			ctx = foundationauth.NewContext(ctx, fresh)
		}
		if service == nil {
			return docserr.AuthorizationUnavailable()
		}
		d, err := service.Decide(ctx, docsauthz.CapabilityImportManage, docsauthz.RootScopeID, authorization.ResourceFacts{})
		if err != nil {
			return docserr.AuthorizationUnavailable()
		}
		if !d.Allowed {
			return docserr.Forbidden()
		}
		return nil
	}
}
