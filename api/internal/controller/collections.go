package controller

import (
	"context"

	"github.com/yueli-official/foundation/go/authorization"

	v1 "github.com/yueli-official/docs/api/api/v1"
	"github.com/yueli-official/docs/api/internal/catalog"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/model"
)

// Collections handles the admin (JWT) collection-management endpoints.
type Collections struct{ svc *catalog.Service }

func NewCollections(svc *catalog.Service) *Collections { return &Collections{svc: svc} }

func (c *Collections) CreateCollection(ctx context.Context, req *v1.CreateCollectionReq) (*v1.CreateCollectionRes, error) {
	if err := requireCapability(ctx, docsauthz.CapabilityCollectionManage, docsauthz.RootScopeID, authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	author, err := subject(ctx)
	if err != nil {
		return nil, err
	}
	col, err := c.svc.CreateCollectionWithSetup(ctx, catalog.CreateCollectionInput{
		AuthorSub: author, Title: req.Title, Slug: req.Slug, Description: req.Description,
		Cover: req.Cover, Icon: req.Icon, DefaultLocale: req.DefaultLocale, SemanticVersion: req.SemanticVersion,
	})
	if err != nil {
		return nil, err
	}
	if err := ensureCollectionScope(ctx, col.ID); err != nil {
		return nil, err
	}
	return &v1.CreateCollectionRes{Collection: collectionView(col)}, nil
}

func (c *Collections) UpdateHomeConfig(ctx context.Context, req *v1.UpdateHomeConfigReq) (*v1.UpdateHomeConfigRes, error) {
	if err := requireCapability(ctx, docsauthz.CapabilitySiteSettingsManage, docsauthz.RootScopeID, authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	cfg, err := c.svc.UpdateHomeConfig(ctx, &model.HomeConfig{
		QuickLinks:          req.QuickLinks,
		FeaturedCollections: req.FeaturedCollections,
		HomeEyebrow:         req.HomeEyebrow,
		HomeTitle:           req.HomeTitle,
		HomeSubtitle:        req.HomeSubtitle,
		SiteTitle:           req.SiteTitle,
		SiteDescription:     req.SiteDescription,
		SupportEmail:        req.SupportEmail,
		FooterTagline:       req.FooterTagline,
		FooterCopyright:     req.FooterCopyright,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateHomeConfigRes{Config: homeConfigView(cfg)}, nil
}

func (c *Collections) UpdateCollection(ctx context.Context, req *v1.UpdateCollectionReq) (*v1.UpdateCollectionRes, error) {
	if err := ensureCollectionScope(ctx, req.ID); err != nil {
		return nil, err
	}
	if err := requireCapability(ctx, docsauthz.CapabilityCollectionManage, docsauthz.CollectionScopeID(req.ID), authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	col, err := c.svc.UpdateCollectionWithBearer(ctx, req.ID, bearerOf(ctx), req.Title, req.Slug, req.Description, req.Cover, req.Icon)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateCollectionRes{Collection: collectionView(col)}, nil
}

func (c *Collections) CollectionCoverInit(ctx context.Context, req *v1.CollectionCoverInitReq) (*v1.CollectionCoverInitRes, error) {
	if err := ensureCollectionScope(ctx, req.ID); err != nil {
		return nil, err
	}
	if err := requireCapability(ctx, docsauthz.CapabilityCollectionManage, docsauthz.CollectionScopeID(req.ID), authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	out, err := c.svc.AddCollectionCover(ctx, req.ID, bearerOf(ctx), req.Filename, req.Mime, req.Size)
	if err != nil {
		return nil, err
	}
	return &v1.CollectionCoverInitRes{UploadURL: out.UploadURL, UploadToken: out.UploadToken, UploadHeaders: out.UploadHeaders}, nil
}

func (c *Collections) CollectionCoverFinalize(ctx context.Context, req *v1.CollectionCoverFinalizeReq) (*v1.CollectionCoverFinalizeRes, error) {
	if err := ensureCollectionScope(ctx, req.ID); err != nil {
		return nil, err
	}
	if err := requireCapability(ctx, docsauthz.CapabilityCollectionManage, docsauthz.CollectionScopeID(req.ID), authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	col, err := c.svc.FinalizeCollectionCover(ctx, req.ID, bearerOf(ctx), req.UploadToken)
	if err != nil {
		return nil, err
	}
	return &v1.CollectionCoverFinalizeRes{Collection: collectionView(col), CoverURL: col.CoverURL}, nil
}

func (c *Collections) DeleteCollection(ctx context.Context, req *v1.DeleteCollectionReq) (*v1.DeleteCollectionRes, error) {
	if err := ensureCollectionScope(ctx, req.ID); err != nil {
		return nil, err
	}
	if err := requireCapability(ctx, docsauthz.CapabilityCollectionManage, docsauthz.CollectionScopeID(req.ID), authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	if err := c.svc.DeleteCollectionWithBearer(ctx, req.ID, bearerOf(ctx)); err != nil {
		return nil, err
	}
	return &v1.DeleteCollectionRes{Deleted: true}, nil
}

func (c *Collections) CloneCollectionRelease(ctx context.Context, req *v1.CloneCollectionReleaseReq) (*v1.CloneCollectionReleaseRes, error) {
	if err := ensureCollectionScope(ctx, req.ID); err != nil {
		return nil, err
	}
	if err := requireCapability(ctx, docsauthz.CapabilityCollectionManage, docsauthz.CollectionScopeID(req.ID), authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	author, err := subject(ctx)
	if err != nil {
		return nil, err
	}
	collection, err := c.svc.CloneCollectionAsRelease(ctx, author, bearerOf(ctx), catalog.CloneCollectionReleaseInput{
		SourceCollectionID: req.ID, SourceSemanticVersion: req.SourceSemanticVersion,
		TargetSemanticVersion: req.TargetSemanticVersion, Title: req.Title, Slug: req.Slug,
	})
	if err != nil {
		return nil, err
	}
	if err := ensureCollectionScope(ctx, collection.ID); err != nil {
		return nil, err
	}
	return &v1.CloneCollectionReleaseRes{Collection: collectionView(collection)}, nil
}

func (c *Collections) InitializeCollectionRelease(ctx context.Context, req *v1.InitializeCollectionReleaseReq) (*v1.InitializeCollectionReleaseRes, error) {
	if err := ensureCollectionScope(ctx, req.ID); err != nil {
		return nil, err
	}
	if err := requireCapability(ctx, docsauthz.CapabilityCollectionManage, docsauthz.CollectionScopeID(req.ID), authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	collection, err := c.svc.InitializeCollectionRelease(ctx, req.ID, req.SemanticVersion)
	if err != nil {
		return nil, err
	}
	return &v1.InitializeCollectionReleaseRes{Collection: collectionView(collection)}, nil
}

func (c *Collections) GetManageCollectionTree(ctx context.Context, req *v1.GetManageCollectionTreeReq) (*v1.GetManageCollectionTreeRes, error) {
	col, tree, err := c.svc.ManageDocTree(ctx, req.Slug, req.Version, req.Locale)
	if err != nil {
		return nil, err
	}
	if err := ensureCollectionScope(ctx, col.ID); err != nil {
		return nil, err
	}
	if err := requireCapability(ctx, docsauthz.CapabilityCollectionManage, docsauthz.CollectionScopeID(col.ID), authorization.ResourceFacts{}); err != nil {
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
