package controller

import (
	"context"

	"github.com/yueli-official/foundation/go/authorization"

	v1 "github.com/yueli-official/docs/api/api/v1"
	"github.com/yueli-official/docs/api/internal/catalog"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/docserr"
	"github.com/yueli-official/docs/api/internal/model"
	foundationauth "github.com/yueli-official/foundation/go/auth"
)

type Versions struct{ svc *catalog.Service }

func NewVersions(svc *catalog.Service) *Versions { return &Versions{svc: svc} }

func (c *Versions) ListCollectionVersions(ctx context.Context, req *v1.ListCollectionVersionsReq) (*v1.ListCollectionVersionsRes, error) {
	if err := authorizePersonalVersionRead(ctx, req.CollectionID); err != nil {
		return nil, err
	}
	items, err := c.svc.ListVersions(ctx, req.CollectionID)
	if err != nil {
		return nil, err
	}
	out := make([]*v1.CollectionVersionView, len(items))
	for i, item := range items {
		out[i] = versionView(item)
	}
	return &v1.ListCollectionVersionsRes{Items: out}, nil
}

func authorizePersonalVersionRead(ctx context.Context, collectionID string) error {
	principal, _ := foundationauth.FromContext(ctx)
	if principal == nil || !principal.IsPersonalToken() {
		return nil
	}
	if err := ensureCollectionScope(ctx, collectionID); err != nil {
		return err
	}
	service := authorizationService(ctx)
	if err := service.ReconcileSubject(ctx); err != nil {
		return docserr.AuthorizationUnavailable()
	}
	access, err := service.Runtime().EffectiveAccess(ctx, authorization.EffectiveAccessQuery{
		Subject: service.Subject(ctx), ScopeID: docsauthz.CollectionScopeID(collectionID), IncludeDescendants: true,
	})
	if err != nil {
		return docserr.AuthorizationUnavailable()
	}
	for _, capability := range access.Capabilities {
		switch capability {
		case docsauthz.CapabilityVersionManage, docsauthz.CapabilityCollectionManage, docsauthz.CapabilityDocumentRead, docsauthz.CapabilityDocumentCreate:
			if foundationauth.AllowsPersonalCapability(ctx, string(capability)) {
				return nil
			}
		}
	}
	return docserr.Forbidden()
}

func (c *Versions) CreateCollectionVersion(ctx context.Context, req *v1.CreateCollectionVersionReq) (*v1.CreateCollectionVersionRes, error) {
	if err := ensureCollectionScope(ctx, req.CollectionID); err != nil {
		return nil, err
	}
	if err := requireCapability(ctx, docsauthz.CapabilityVersionManage, docsauthz.RootScopeID, authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	v, err := c.svc.CreateVersion(ctx, catalog.CreateVersionInput{
		CollectionID:    req.CollectionID,
		Key:             req.Key,
		Label:           req.Label,
		Status:          req.Status,
		SourceVersionID: req.SourceVersionID,
	})
	if err != nil {
		return nil, err
	}
	writeCreated(ctx)
	return &v1.CreateCollectionVersionRes{Version: versionView(v)}, nil
}

func (c *Versions) UpdateCollectionVersion(ctx context.Context, req *v1.UpdateCollectionVersionReq) (*v1.UpdateCollectionVersionRes, error) {
	if err := ensureCollectionScope(ctx, req.CollectionID); err != nil {
		return nil, err
	}
	if err := requireCapability(ctx, docsauthz.CapabilityVersionManage, docsauthz.RootScopeID, authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	value, err := c.svc.UpdateVersion(ctx, catalog.UpdateVersionInput{
		CollectionID: req.CollectionID,
		VersionID:    req.VersionID,
		Label:        req.Label,
		Status:       req.Status,
		IsDefault:    req.IsDefault,
		SortOrder:    req.SortOrder,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateCollectionVersionRes{Version: versionView(value)}, nil
}

func versionView(m *model.CollectionVersion) *v1.CollectionVersionView {
	if m == nil {
		return nil
	}
	return &v1.CollectionVersionView{
		ID:              m.ID,
		CollectionID:    m.CollectionID,
		Key:             m.Key,
		Label:           m.Label,
		Status:          m.Status,
		IsDefault:       m.IsDefault,
		SortOrder:       m.SortOrder,
		SourceVersionID: m.SourceVersionID,
	}
}
