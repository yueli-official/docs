package controller

import (
	"context"

	v1 "platform/products/docs/api/api/v1"
	"platform/products/docs/api/internal/catalog"
	"platform/products/docs/api/internal/docserr"
	"platform/products/docs/api/internal/model"
)

// Collections handles the admin (JWT) collection-management endpoints.
type Collections struct{ svc *catalog.Service }

func NewCollections(svc *catalog.Service) *Collections { return &Collections{svc: svc} }

func (c *Collections) CreateCollection(ctx context.Context, req *v1.CreateCollectionReq) (*v1.CreateCollectionRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
	}
	author, err := subject(ctx)
	if err != nil {
		return nil, err
	}
	col, err := c.svc.CreateCollection(ctx, author, req.Title, req.Slug, req.Description, req.Cover, req.Icon)
	if err != nil {
		return nil, err
	}
	return &v1.CreateCollectionRes{Collection: collectionView(col)}, nil
}

func (c *Collections) UpdateHomeConfig(ctx context.Context, req *v1.UpdateHomeConfigReq) (*v1.UpdateHomeConfigRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
	}
	cfg, err := c.svc.UpdateHomeConfig(ctx, &model.HomeConfig{
		QuickLinks:          req.QuickLinks,
		FeaturedCollections: req.FeaturedCollections,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateHomeConfigRes{Config: homeConfigView(cfg)}, nil
}

func (c *Collections) UpdateCollection(ctx context.Context, req *v1.UpdateCollectionReq) (*v1.UpdateCollectionRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
	}
	col, err := c.svc.UpdateCollectionWithBearer(ctx, req.ID, bearerOf(ctx), req.Title, req.Slug, req.Description, req.Cover, req.Icon)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateCollectionRes{Collection: collectionView(col)}, nil
}

func (c *Collections) CollectionCoverInit(ctx context.Context, req *v1.CollectionCoverInitReq) (*v1.CollectionCoverInitRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
	}
	out, err := c.svc.AddCollectionCover(ctx, req.ID, bearerOf(ctx), req.Filename, req.Mime, req.Size)
	if err != nil {
		return nil, err
	}
	return &v1.CollectionCoverInitRes{UploadURL: out.UploadURL, UploadToken: out.UploadToken, UploadHeaders: out.UploadHeaders}, nil
}

func (c *Collections) CollectionCoverFinalize(ctx context.Context, req *v1.CollectionCoverFinalizeReq) (*v1.CollectionCoverFinalizeRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
	}
	col, err := c.svc.FinalizeCollectionCover(ctx, req.ID, bearerOf(ctx), req.UploadToken)
	if err != nil {
		return nil, err
	}
	return &v1.CollectionCoverFinalizeRes{Collection: collectionView(col), CoverURL: col.CoverURL}, nil
}

func (c *Collections) DeleteCollection(ctx context.Context, req *v1.DeleteCollectionReq) (*v1.DeleteCollectionRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
	}
	if err := c.svc.DeleteCollectionWithBearer(ctx, req.ID, bearerOf(ctx)); err != nil {
		return nil, err
	}
	return &v1.DeleteCollectionRes{Deleted: true}, nil
}

func (c *Collections) GetManageCollectionTree(ctx context.Context, req *v1.GetManageCollectionTreeReq) (*v1.GetManageCollectionTreeRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
	}
	col, tree, err := c.svc.ManageDocTree(ctx, req.Slug, req.Version, req.Locale)
	if err != nil {
		return nil, err
	}
	return &v1.GetManageCollectionTreeRes{
		Collection: collectionView(col),
		Tree:       docTreeViews(tree),
	}, nil
}

func docTreeView(m *model.Doc) *v1.DocView {
	if m == nil {
		return nil
	}
	dv := docView(m)
	if len(m.Children) > 0 {
		dv.Children = docTreeViews(m.Children)
	}
	return dv
}

func docTreeViews(ms []*model.Doc) []*v1.DocView {
	out := make([]*v1.DocView, len(ms))
	for i, m := range ms {
		out[i] = docTreeView(m)
	}
	return out
}
