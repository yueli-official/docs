package controller

import (
	"context"

	"github.com/yueli-official/foundation/go/discovery"

	v1 "github.com/yueli-official/docs/api/api/v1"
	"github.com/yueli-official/docs/api/internal/catalog"
	"github.com/yueli-official/docs/api/internal/docsdiscovery"
	"github.com/yueli-official/docs/api/internal/model"
)

// PublicCollections handles the public collection browse endpoints (no mandatory auth).
type PublicCollections struct {
	svc       *catalog.Service
	discovery *discovery.Module
}

func NewPublicCollections(svc *catalog.Service, modules ...*discovery.Module) *PublicCollections {
	controller := &PublicCollections{svc: svc}
	if len(modules) > 0 {
		controller.discovery = modules[0]
	}
	return controller
}

func (c *PublicCollections) ListCollections(ctx context.Context, _ *v1.ListCollectionsReq) (*v1.ListCollectionsRes, error) {
	items, err := c.svc.ListCollections(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ListCollectionsRes{Items: collectionViews(items)}, nil
}

func (c *PublicCollections) GetHomeConfig(ctx context.Context, _ *v1.GetHomeConfigReq) (*v1.GetHomeConfigRes, error) {
	cfg, err := c.svc.GetHomeConfig(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetHomeConfigRes{Config: homeConfigView(cfg)}, nil
}

func (c *PublicCollections) GetCollection(ctx context.Context, req *v1.GetCollectionReq) (*v1.GetCollectionRes, error) {
	col, err := c.svc.GetCollectionBySlug(ctx, req.Slug)
	if err != nil {
		return nil, err
	}
	return &v1.GetCollectionRes{Collection: collectionView(col)}, nil
}

func (c *PublicCollections) GetCollectionReleases(ctx context.Context, req *v1.GetCollectionReleasesReq) (*v1.GetCollectionReleasesRes, error) {
	items, err := c.svc.CollectionReleases(ctx, req.Slug)
	if err != nil {
		return nil, err
	}
	return &v1.GetCollectionReleasesRes{Items: collectionViews(items)}, nil
}

func (c *PublicCollections) GetCollectionVariants(ctx context.Context, req *v1.GetCollectionVariantsReq) (*v1.GetCollectionVariantsRes, error) {
	variants, err := c.svc.PublicVariants(ctx, req.Slug)
	if err != nil {
		return nil, err
	}
	locales := make([]*v1.CollectionLocaleView, len(variants.Locales))
	for index, item := range variants.Locales {
		locales[index] = localeView(item)
	}
	versions := make([]*v1.CollectionVersionView, len(variants.Versions))
	for index, item := range variants.Versions {
		versions[index] = versionView(item)
	}
	return &v1.GetCollectionVariantsRes{Locales: locales, Versions: versions}, nil
}

func (c *PublicCollections) ResolveDocumentVariant(ctx context.Context, req *v1.ResolveDocumentVariantReq) (*v1.ResolveDocumentVariantRes, error) {
	resolution, err := c.svc.ResolveDocumentVariant(ctx, req.Slug, req.TranslationKey, req.Locale, req.Version)
	if err != nil {
		return nil, err
	}
	return &v1.ResolveDocumentVariantRes{
		Path:           resolution.Path,
		Locale:         resolution.Locale,
		Version:        resolution.Version,
		TranslationKey: resolution.TranslationKey,
		Fallback:       resolution.Fallback,
	}, nil
}

func (c *PublicCollections) GetCollectionTree(ctx context.Context, req *v1.GetCollectionTreeReq) (*v1.GetCollectionTreeRes, error) {
	col, tree, err := c.svc.PublicDocTree(ctx, req.Slug, req.Version, req.Locale)
	if err != nil {
		return nil, err
	}
	return &v1.GetCollectionTreeRes{
		Collection: collectionView(col),
		Tree:       docTreeNodeViews(tree),
	}, nil
}

func (c *PublicCollections) GetPublicDocByPath(ctx context.Context, req *v1.GetPublicDocByPathReq) (*v1.GetPublicDocByPathRes, error) {
	d, err := c.svc.GetPublishedDocByPath(ctx, req.Collection, req.Version, req.Path, req.Locale)
	if err != nil {
		return nil, err
	}
	response := &v1.GetPublicDocByPathRes{Doc: docView(d)}
	if c.discovery != nil {
		collection, err := c.svc.GetCollectionBySlug(ctx, req.Collection)
		if err != nil {
			return nil, err
		}
		version, err := c.svc.ResolveVersion(ctx, collection.ID, req.Version, true)
		if err != nil {
			return nil, err
		}
		defaultLocale, err := c.svc.DefaultLocale(ctx, collection.ID)
		if err != nil {
			return nil, err
		}
		defaultLocaleKey := "en"
		if defaultLocale != nil {
			defaultLocaleKey = defaultLocale.Locale
		}
		projection, err := docsdiscovery.ProjectDoc(c.discovery, d, collection, version, req.Path, defaultLocaleKey)
		if err != nil {
			return nil, err
		}
		response.Discovery = &projection
	}
	return response, nil
}

func (c *PublicCollections) SearchDocs(ctx context.Context, req *v1.SearchDocsReq) (*v1.SearchDocsRes, error) {
	result, err := c.svc.SearchPublishedDocs(ctx, req.Collection, req.Version, req.Locale, req.Q)
	if err != nil {
		return nil, err
	}
	return &v1.SearchDocsRes{
		Items: docTreeNodeViews(result.Items),
		Total: result.Total,
		Facets: v1.SearchFacetsView{
			Collections: searchCollectionFacetViews(result.CollectionFacets),
		},
	}, nil
}

// ── view helpers (public) ─────────────────────────────────────────────────────

func docTreeNodeView(m *model.Doc) *v1.DocTreeNodeView {
	if m == nil {
		return nil
	}
	dv := &v1.DocTreeNodeView{
		ID:             m.ID,
		CollectionID:   m.CollectionID,
		VersionID:      m.VersionID,
		ParentID:       m.ParentID,
		Slug:           m.Slug,
		Title:          m.Title,
		Excerpt:        m.Excerpt,
		Status:         m.Status,
		Locale:         m.Locale,
		TranslationKey: m.TranslationKey,
		BadgeText:      m.BadgeText,
		BadgeIcon:      m.BadgeIcon,
		SortOrder:      m.SortOrder,
	}
	if len(m.Children) > 0 {
		dv.Children = docTreeNodeViews(m.Children)
	}
	return dv
}

func docTreeNodeViews(ms []*model.Doc) []*v1.DocTreeNodeView {
	out := make([]*v1.DocTreeNodeView, len(ms))
	for i, m := range ms {
		out[i] = docTreeNodeView(m)
	}
	return out
}

func searchCollectionFacetViews(ms []*model.SearchCollectionFacet) []*v1.SearchCollectionFacetView {
	out := make([]*v1.SearchCollectionFacetView, len(ms))
	for i, m := range ms {
		out[i] = &v1.SearchCollectionFacetView{
			ID:    m.ID,
			Slug:  m.Slug,
			Title: m.Title,
			Count: m.Count,
		}
	}
	return out
}

// ── collection view helpers ───────────────────────────────────────────────────

func collectionView(m *model.Collection) *v1.CollectionView {
	if m == nil {
		return nil
	}
	return &v1.CollectionView{
		ID: m.ID, Slug: m.Slug, Title: m.Title, Description: m.Description,
		CoverAssetID: m.CoverAssetID, CoverURL: m.CoverURL, Icon: m.Icon,
		SortOrder: m.SortOrder, ReleaseFamilyID: m.ReleaseFamilyID,
		ReleaseFamilyName: m.ReleaseFamilyName, SemanticVersion: m.SemanticVersion,
		DerivedFromCollectionID: m.DerivedFromCollectionID, DocCount: m.DocCount,
	}
}

func collectionViews(ms []*model.Collection) []*v1.CollectionView {
	out := make([]*v1.CollectionView, len(ms))
	for i, m := range ms {
		out[i] = collectionView(m)
	}
	return out
}
