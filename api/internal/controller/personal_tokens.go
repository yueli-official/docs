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

func NewPersonalPermissions(site string, service *docsauthz.Service) *PersonalPermissions {
	return &PersonalPermissions{site, service}
}
func (c *PersonalPermissions) GetPersonalPermissions(ctx context.Context, req *v1.PersonalPermissionsReq) (*v1.PersonalPermissionsRes, error) {
	p, _ := foundationauth.FromContext(ctx)
	if p == nil || p.SubjectKind != foundationauth.SubjectClient || p.ClientID != "identity-svc" || !p.HasScope(foundationauth.PersonalPermissionsScope) || c.site == "" {
		return nil, docserr.Forbidden()
	}
	if c.service == nil {
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
	for _, key := range access.Capabilities {
		if key == docsauthz.CapabilityImportManage {
			items = append(items, foundationauth.PersonalPermission{Key: string(key), Label: "导入文档", Description: "预检、导入并发布文档包及图片，完整镜像模式还会下架缺失页面；仍需当前站点导入权限。"})
		}
		labels := map[authorization.CapabilityKey]string{docsauthz.CapabilityDocumentRead: "读取文档", docsauthz.CapabilityDocumentCreate: "创建文档", docsauthz.CapabilityDocumentUpdate: "编辑文档", docsauthz.CapabilityDocumentPublish: "发布文档", docsauthz.CapabilityDocumentArchive: "下架文档"}
		if label, ok := labels[key]; ok {
			items = append(items, foundationauth.PersonalPermission{Key: string(key), Label: label, Description: "仅限当前账号有权操作的文档集和文档。"})
		}
	}
	return &v1.PersonalPermissionsRes{Site: c.site, UserKey: req.UserKey, Items: items}, nil
}
func (c *PersonalPermissions) AuthorizePersonalMedia(ctx context.Context, _ *v1.PersonalMediaAuthorizationReq) (*v1.PersonalMediaAuthorizationRes, error) {
	p, _ := foundationauth.FromContext(ctx)
	if p == nil || !p.IsPersonalToken() || c.site == "" {
		return nil, docserr.Forbidden()
	}
	if c.service == nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	decision, err := c.service.Decide(ctx, docsauthz.CapabilityImportManage, docsauthz.RootScopeID, authorization.ResourceFacts{})
	if err != nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	if !decision.Allowed {
		return nil, docserr.Forbidden()
	}
	scope, _ := foundationauth.PersonalScope(c.site, "asset.profile.docs-import-image.upload")
	out := &v1.PersonalMediaAuthorizationRes{UserKey: p.Subject, Scopes: []string{scope}}
	if !p.ExpiresAt.IsZero() {
		out.ExpiresAt = p.ExpiresAt.Format(time.RFC3339)
	}
	return out, nil
}
func PersonalTokenRoutes(r *ghttp.Request) {
	p, _ := foundationauth.FromContext(r.Context())
	if p != nil && p.IsPersonalToken() {
		path := r.URL.Path
		allowed := path == "/api/v1/imports/docs" && (r.Method == "GET" || r.Method == "POST")
		capability := ""
		if path == "/api/v1/docs" {
			if r.Method == "GET" {
				allowed = true
				capability = string(docsauthz.CapabilityDocumentRead)
			}
			if r.Method == "POST" {
				allowed = true
				capability = string(docsauthz.CapabilityDocumentCreate)
			}
		}
		if path == "/api/v1/manage/docs" && r.Method == "GET" {
			allowed = true
			capability = string(docsauthz.CapabilityDocumentRead)
		}
		if rest, ok := strings.CutPrefix(path, "/api/v1/docs/"); ok {
			parts := strings.Split(rest, "/")
			if len(parts) == 1 && r.Method == "GET" {
				allowed = true
				capability = string(docsauthz.CapabilityDocumentRead)
			}
			if len(parts) == 1 && r.Method == "PATCH" {
				allowed = true
				capability = string(docsauthz.CapabilityDocumentUpdate)
			}
			if len(parts) == 2 && r.Method == "POST" && parts[1] == "publish" {
				allowed = true
				capability = string(docsauthz.CapabilityDocumentPublish)
			}
			if len(parts) == 2 && r.Method == "POST" && parts[1] == "archive" {
				allowed = true
				capability = string(docsauthz.CapabilityDocumentArchive)
			}
		}
		if strings.HasPrefix(path, "/api/v1/manage/collections/") && strings.HasSuffix(path, "/tree") && r.Method == "GET" {
			allowed = true
			capability = string(docsauthz.CapabilityDocumentRead)
		}
		tail := strings.TrimPrefix(path, "/api/v1/imports/docs/")
		if tail != path && tail != "" {
			parts := strings.Split(tail, "/")
			allowed = r.Method == "GET" && len(parts) == 1 || r.Method == "POST" && len(parts) == 2 && parts[1] == "confirm"
		}
		if path == "/api/v1/personal-token/media-authorization" && r.Method == "POST" {
			allowed = true
		}
		if !allowed || (capability != "" && !foundationauth.AllowsPersonalCapability(r.Context(), capability)) {
			r.SetError(docserr.Forbidden())
			return
		}
	}
	r.Middleware.Next()
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
