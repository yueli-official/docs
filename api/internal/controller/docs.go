package controller

import (
	"context"

	v1 "platform/products/docs/api/api/v1"
	"platform/products/docs/api/internal/catalog"
	"platform/products/docs/api/internal/docserr"
	"platform/products/docs/api/internal/model"
)

// Docs handles the admin (JWT) doc-management endpoints.
type Docs struct{ svc *catalog.Service }

// NewDocs returns a new Docs controller backed by the given catalog service.
func NewDocs(svc *catalog.Service) *Docs { return &Docs{svc: svc} }

func (c *Docs) ListDocs(ctx context.Context, req *v1.ListDocsReq) (*v1.ListDocsRes, error) {
	locale := req.Locale
	if locale == "" {
		locale = "en"
	}
	items, err := c.svc.ListDocs(ctx, req.CollectionID, req.Version, locale)
	if err != nil {
		return nil, err
	}
	return &v1.ListDocsRes{Items: docViews(items)}, nil
}

func (c *Docs) GetDoc(ctx context.Context, req *v1.GetDocReq) (*v1.GetDocRes, error) {
	d, err := c.svc.GetDoc(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &v1.GetDocRes{Doc: docView(d)}, nil
}

func (c *Docs) CreateDoc(ctx context.Context, req *v1.CreateDocReq) (*v1.CreateDocRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
	}
	author, err := subject(ctx)
	if err != nil {
		return nil, err
	}
	d, err := c.svc.CreateDoc(ctx, author, catalog.CreateDocInput{
		CollectionID:   req.CollectionID,
		VersionID:      req.VersionID,
		ParentID:       req.ParentID,
		Slug:           req.Slug,
		Title:          req.Title,
		Content:        req.Content,
		SEOTitle:       req.SEOTitle,
		SEODescription: req.SEODescription,
		Locale:         req.Locale,
		TranslationKey: req.TranslationKey,
		SortOrder:      req.SortOrder,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateDocRes{Doc: docView(d)}, nil
}

func (c *Docs) UpdateDoc(ctx context.Context, req *v1.UpdateDocReq) (*v1.UpdateDocRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
	}
	d, err := c.svc.PatchDoc(ctx, req.ID, catalog.PatchDocInput{
		Title:          req.Title,
		Slug:           req.Slug,
		Content:        req.Content,
		Excerpt:        req.Excerpt,
		SEOTitle:       req.SEOTitle,
		SEODescription: req.SEODescription,
		Status:         req.Status,
		Locale:         req.Locale,
		VersionID:      req.VersionID,
		TranslationKey: req.TranslationKey,
		SortOrder:      req.SortOrder,
		ParentID:       req.ParentID,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateDocRes{Doc: docView(d)}, nil
}

func (c *Docs) PublishDoc(ctx context.Context, req *v1.PublishDocReq) (*v1.PublishDocRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
	}
	d, err := c.svc.PublishDoc(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &v1.PublishDocRes{Doc: docView(d)}, nil
}

func (c *Docs) ArchiveDoc(ctx context.Context, req *v1.ArchiveDocReq) (*v1.ArchiveDocRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
	}
	d, err := c.svc.ArchiveDoc(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &v1.ArchiveDocRes{Doc: docView(d)}, nil
}

func (c *Docs) DeleteDoc(ctx context.Context, req *v1.DeleteDocReq) (*v1.DeleteDocRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
	}
	if err := c.svc.DeleteDoc(ctx, req.ID); err != nil {
		return nil, err
	}
	return &v1.DeleteDocRes{Deleted: true}, nil
}

// ── view helpers ──────────────────────────────────────────────────────────────

func docView(m *model.Doc) *v1.DocView {
	if m == nil {
		return nil
	}
	return &v1.DocView{
		ID:             m.ID,
		CollectionID:   m.CollectionID,
		VersionID:      m.VersionID,
		ParentID:       m.ParentID,
		Slug:           m.Slug,
		Title:          m.Title,
		Content:        m.Content,
		Excerpt:        m.Excerpt,
		SEOTitle:       m.SEOTitle,
		SEODescription: m.SEODescription,
		Status:         m.Status,
		Locale:         m.Locale,
		TranslationKey: m.TranslationKey,
		SortOrder:      m.SortOrder,
	}
}

func docViews(ms []*model.Doc) []*v1.DocView {
	out := make([]*v1.DocView, len(ms))
	for i, m := range ms {
		out[i] = docView(m)
	}
	return out
}
