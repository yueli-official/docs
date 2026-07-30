package controller

import (
	"context"

	"github.com/yueli-official/foundation/go/authorization"

	v1 "github.com/yueli-official/docs/api/api/v1"
	"github.com/yueli-official/docs/api/internal/catalog"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/model"
)

type Versions struct{ svc *catalog.Service }

func NewVersions(svc *catalog.Service) *Versions { return &Versions{svc: svc} }

func (c *Versions) ListCollectionVersions(ctx context.Context, req *v1.ListCollectionVersionsReq) (*v1.ListCollectionVersionsRes, error) {
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

func (c *Versions) CreateCollectionVersion(ctx context.Context, req *v1.CreateCollectionVersionReq) (*v1.CreateCollectionVersionRes, error) {
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
	return &v1.CreateCollectionVersionRes{Version: versionView(v)}, nil
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
