package controller

import (
	"context"

	"github.com/yueli-official/foundation/go/authorization"

	v1 "github.com/yueli-official/docs/api/api/v1"
	"github.com/yueli-official/docs/api/internal/catalog"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/model"
)

type Locales struct{ svc *catalog.Service }

func NewLocales(svc *catalog.Service) *Locales { return &Locales{svc: svc} }

func (c *Locales) ListManageCollectionLocales(ctx context.Context, req *v1.ListManageCollectionLocalesReq) (*v1.ListManageCollectionLocalesRes, error) {
	if err := ensureCollectionScope(ctx, req.CollectionID); err != nil {
		return nil, err
	}
	if err := requireCapability(ctx, docsauthz.CapabilityCollectionManage, docsauthz.CollectionScopeID(req.CollectionID), authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	items, err := c.svc.ListLocales(ctx, req.CollectionID, false)
	if err != nil {
		return nil, err
	}
	views := make([]*v1.CollectionLocaleView, len(items))
	for index, item := range items {
		views[index] = localeView(item)
	}
	return &v1.ListManageCollectionLocalesRes{Items: views}, nil
}

func (c *Locales) UpsertCollectionLocale(ctx context.Context, req *v1.UpsertCollectionLocaleReq) (*v1.UpsertCollectionLocaleRes, error) {
	if err := ensureCollectionScope(ctx, req.CollectionID); err != nil {
		return nil, err
	}
	if err := requireCapability(ctx, docsauthz.CapabilityCollectionManage, docsauthz.CollectionScopeID(req.CollectionID), authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	value, err := c.svc.UpsertLocale(ctx, catalog.UpsertLocaleInput{
		CollectionID: req.CollectionID,
		Locale:       req.Locale,
		Label:        req.Label,
		HTMLLang:     req.HTMLLang,
		Direction:    req.Direction,
		IsDefault:    req.IsDefault,
		Enabled:      req.Enabled,
		SortOrder:    req.SortOrder,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpsertCollectionLocaleRes{Locale: localeView(value)}, nil
}

func (c *Locales) DeleteCollectionLocale(ctx context.Context, req *v1.DeleteCollectionLocaleReq) (*v1.DeleteCollectionLocaleRes, error) {
	if err := ensureCollectionScope(ctx, req.CollectionID); err != nil {
		return nil, err
	}
	if err := requireCapability(ctx, docsauthz.CapabilityCollectionManage, docsauthz.CollectionScopeID(req.CollectionID), authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	if err := c.svc.DeleteLocale(ctx, req.CollectionID, req.Locale); err != nil {
		return nil, err
	}
	return &v1.DeleteCollectionLocaleRes{Deleted: true}, nil
}

func (c *Locales) CloneCollectionLocale(ctx context.Context, req *v1.CloneCollectionLocaleReq) (*v1.CloneCollectionLocaleRes, error) {
	if err := ensureCollectionScope(ctx, req.CollectionID); err != nil {
		return nil, err
	}
	if err := requireCapability(ctx, docsauthz.CapabilityCollectionManage, docsauthz.CollectionScopeID(req.CollectionID), authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	author, err := subject(ctx)
	if err != nil {
		return nil, err
	}
	created, err := c.svc.CloneLocale(ctx, author, catalog.CloneLocaleInput{
		CollectionID: req.CollectionID, SourceLocale: req.SourceLocale,
		TargetLocale: req.TargetLocale, TargetLabel: req.TargetLabel,
		TargetHTMLLang: req.TargetHTMLLang, TargetDirection: req.TargetDirection,
		TargetSortOrder: req.TargetSortOrder,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CloneCollectionLocaleRes{Created: created}, nil
}

func localeView(value *model.CollectionLocale) *v1.CollectionLocaleView {
	if value == nil {
		return nil
	}
	return &v1.CollectionLocaleView{
		CollectionID: value.CollectionID,
		Locale:       value.Locale,
		Label:        value.Label,
		HTMLLang:     value.HTMLLang,
		Direction:    value.Direction,
		IsDefault:    value.IsDefault,
		Enabled:      value.Enabled,
		SortOrder:    value.SortOrder,
		DocCount:     value.DocCount,
	}
}
