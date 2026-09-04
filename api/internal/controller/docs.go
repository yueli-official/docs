package controller

import (
	"context"

	"github.com/yueli-official/foundation/go/authorization"

	v1 "github.com/yueli-official/docs/api/api/v1"
	"github.com/yueli-official/docs/api/internal/catalog"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/docserr"
	"github.com/yueli-official/docs/api/internal/model"
)

// Docs handles the admin (JWT) doc-management endpoints.
type Docs struct{ svc *catalog.Service }

// NewDocs returns a new Docs controller backed by the given catalog service.
func NewDocs(svc *catalog.Service) *Docs { return &Docs{svc: svc} }

func (c *Docs) ListDocs(ctx context.Context, req *v1.ListDocsReq) (*v1.ListDocsRes, error) {
	if err := ensureCollectionScope(ctx, req.CollectionID); err != nil {
		return nil, err
	}
	items, err := c.svc.ListDocs(ctx, req.CollectionID, req.Version, req.Locale)
	if err != nil {
		return nil, err
	}
	visible := make([]*model.Doc, 0, len(items))
	for _, item := range items {
		if err := ensureDocumentScope(ctx, item.ID, item.CollectionID); err != nil {
			return nil, err
		}
		if err := requireCapability(
			ctx, docsauthz.CapabilityDocumentRead, docsauthz.DocumentScopeID(item.ID),
			docsauthz.DocumentResource(item.ID, item.AuthorSub),
		); err == nil {
			visible = append(visible, item)
		}
	}
	return &v1.ListDocsRes{Items: docViews(visible)}, nil
}

func (c *Docs) ManageDocs(ctx context.Context, req *v1.ManageDocsReq) (*v1.ManageDocsRes, error) {
	scopeID := docsauthz.RootScopeID
	if req.CollectionID != "" {
		if err := ensureCollectionScope(ctx, req.CollectionID); err != nil {
			return nil, err
		}
		scopeID = docsauthz.CollectionScopeID(req.CollectionID)
	}
	constraint, err := authorizationService(ctx).Runtime().Plan(ctx, authorization.QueryRequest{
		Subject:    authorizationService(ctx).Subject(ctx),
		Capability: docsauthz.CapabilityDocumentRead,
		ScopeID:    scopeID,
	})
	if err != nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	if constraint.Kind == authorization.QueryNone {
		return nil, docserr.Forbidden()
	}
	ownerSub := ""
	if constraint.Kind == authorization.QueryRelation && constraint.Relation == docsauthz.RelationOwner {
		ownerSub = constraint.Subject.ID
	} else if constraint.Kind != authorization.QueryAll {
		return nil, docserr.AuthorizationUnavailable()
	}
	result, err := c.svc.ManageDocs(ctx, catalog.ManageDocsInput{
		Q: req.Q, Status: req.Status, Quality: req.Quality,
		CollectionID: req.CollectionID, Version: req.Version, Locale: req.Locale, ParentID: req.ParentID,
		Sort: req.Sort, Direction: req.Direction, Page: req.Page, Size: req.Size, OwnerSub: ownerSub,
	})
	if err != nil {
		return nil, err
	}
	items := make([]*v1.ManageDocView, len(result.Items))
	for index, item := range result.Items {
		items[index] = &v1.ManageDocView{
			ID: item.ID, CollectionID: item.CollectionID, CollectionSlug: item.CollectionSlug, CollectionTitle: item.CollectionTitle,
			VersionID: item.VersionID, VersionKey: item.VersionKey, VersionLabel: item.VersionLabel,
			ParentID: item.ParentID, ParentTitle: item.ParentTitle, Slug: item.Slug, SlugPath: item.SlugPath,
			Title: item.Title, Excerpt: item.Excerpt, Status: item.Status, Locale: item.Locale,
			BadgeText: item.BadgeText, BadgeIcon: item.BadgeIcon,
			SortOrder: item.SortOrder, UpdatedAt: item.UpdatedAt,
		}
	}
	return &v1.ManageDocsRes{
		Items: items, Total: result.Total, Page: max(req.Page, 1), Size: normalizedManageDocsSize(req.Size),
		Counts: v1.ManageDocCountsView{
			All: result.Counts.All, Draft: result.Counts.Draft, Published: result.Counts.Published,
			Archived: result.Counts.Archived, Issues: result.Counts.Issues,
		},
	}, nil
}

func normalizedManageDocsSize(size int) int {
	if size == 0 {
		return 30
	}
	return size
}

func (c *Docs) GetDoc(ctx context.Context, req *v1.GetDocReq) (*v1.GetDocRes, error) {
	d, err := c.svc.GetDoc(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if err := ensureDocumentHierarchy(ctx, d.ID, d.CollectionID); err != nil {
		return nil, err
	}
	if err := requireCapability(
		ctx, docsauthz.CapabilityDocumentRead, docsauthz.DocumentScopeID(d.ID),
		docsauthz.DocumentResource(d.ID, d.AuthorSub),
	); err != nil {
		return nil, err
	}
	return &v1.GetDocRes{Doc: docView(d)}, nil
}

func (c *Docs) CreateDoc(ctx context.Context, req *v1.CreateDocReq) (*v1.CreateDocRes, error) {
	if err := ensureCollectionScope(ctx, req.CollectionID); err != nil {
		return nil, err
	}
	if err := requireCapability(ctx, docsauthz.CapabilityDocumentCreate, docsauthz.CollectionScopeID(req.CollectionID), authorization.ResourceFacts{}); err != nil {
		return nil, err
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
		BadgeText:      req.BadgeText,
		BadgeIcon:      req.BadgeIcon,
		SortOrder:      req.SortOrder,
	})
	if err != nil {
		return nil, err
	}
	if err := ensureDocumentScope(ctx, d.ID, d.CollectionID); err != nil {
		return nil, err
	}
	writeCreated(ctx)
	return &v1.CreateDocRes{Doc: docView(d)}, nil
}

func (c *Docs) authorizeImageUpload(ctx context.Context, collectionID, documentID string) error {
	if documentID != "" {
		current, err := c.svc.GetDoc(ctx, documentID)
		if err != nil {
			return err
		}
		if err := ensureDocumentHierarchy(ctx, current.ID, current.CollectionID); err != nil {
			return err
		}
		return requireCapability(
			ctx,
			docsauthz.CapabilityDocumentUpdate,
			docsauthz.DocumentScopeID(current.ID),
			docsauthz.DocumentResource(current.ID, current.AuthorSub),
		)
	}
	if collectionID == "" {
		return docserr.InvalidInput("collectionId or documentId is required")
	}
	if err := ensureCollectionScope(ctx, collectionID); err != nil {
		return err
	}
	return requireCapability(
		ctx,
		docsauthz.CapabilityDocumentCreate,
		docsauthz.CollectionScopeID(collectionID),
		authorization.ResourceFacts{},
	)
}

func (c *Docs) ImageInit(ctx context.Context, req *v1.ImageInitReq) (*v1.ImageInitRes, error) {
	if err := c.authorizeImageUpload(ctx, req.CollectionID, req.DocumentID); err != nil {
		return nil, err
	}
	out, err := c.svc.InitDocumentImage(ctx, bearerOf(ctx), req.Filename, req.Mime, req.Size)
	if err != nil {
		return nil, err
	}
	writeCreated(ctx)
	return &v1.ImageInitRes{
		UploadURL:     out.UploadURL,
		UploadToken:   out.UploadToken,
		UploadHeaders: out.UploadHeaders,
	}, nil
}

func (c *Docs) ImageFinalize(ctx context.Context, req *v1.ImageFinalizeReq) (*v1.ImageFinalizeRes, error) {
	if err := c.authorizeImageUpload(ctx, req.CollectionID, req.DocumentID); err != nil {
		return nil, err
	}
	imageURL, err := c.svc.FinalizeDocumentImage(ctx, bearerOf(ctx), req.UploadToken)
	if err != nil {
		return nil, err
	}
	return &v1.ImageFinalizeRes{URL: imageURL}, nil
}

func (c *Docs) UpdateDoc(ctx context.Context, req *v1.UpdateDocReq) (*v1.UpdateDocRes, error) {
	current, err := c.svc.GetDoc(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if err := ensureDocumentHierarchy(ctx, current.ID, current.CollectionID); err != nil {
		return nil, err
	}
	if err := requireCapability(
		ctx, docsauthz.CapabilityDocumentUpdate, docsauthz.DocumentScopeID(current.ID),
		docsauthz.DocumentResource(current.ID, current.AuthorSub),
	); err != nil {
		return nil, err
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
		BadgeText:      req.BadgeText,
		BadgeIcon:      req.BadgeIcon,
		SortOrder:      req.SortOrder,
		ParentID:       req.ParentID,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateDocRes{Doc: docView(d)}, nil
}

func (c *Docs) PublishDoc(ctx context.Context, req *v1.PublishDocReq) (*v1.PublishDocRes, error) {
	current, err := c.svc.GetDoc(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if err := ensureDocumentHierarchy(ctx, current.ID, current.CollectionID); err != nil {
		return nil, err
	}
	if err := requireCapability(
		ctx, docsauthz.CapabilityDocumentPublish, docsauthz.DocumentScopeID(current.ID),
		docsauthz.DocumentResource(current.ID, current.AuthorSub),
	); err != nil {
		return nil, err
	}
	d, err := c.svc.PublishDoc(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &v1.PublishDocRes{Doc: docView(d)}, nil
}

func (c *Docs) ArchiveDoc(ctx context.Context, req *v1.ArchiveDocReq) (*v1.ArchiveDocRes, error) {
	current, err := c.svc.GetDoc(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if err := ensureDocumentHierarchy(ctx, current.ID, current.CollectionID); err != nil {
		return nil, err
	}
	if err := requireCapability(
		ctx, docsauthz.CapabilityDocumentArchive, docsauthz.DocumentScopeID(current.ID),
		docsauthz.DocumentResource(current.ID, current.AuthorSub),
	); err != nil {
		return nil, err
	}
	d, err := c.svc.ArchiveDoc(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &v1.ArchiveDocRes{Doc: docView(d)}, nil
}

func (c *Docs) DeleteDoc(ctx context.Context, req *v1.DeleteDocReq) (*v1.DeleteDocRes, error) {
	current, err := c.svc.GetDoc(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if err := ensureDocumentHierarchy(ctx, current.ID, current.CollectionID); err != nil {
		return nil, err
	}
	if err := requireCapability(
		ctx, docsauthz.CapabilityDocumentDeletePermanently, docsauthz.DocumentScopeID(current.ID),
		docsauthz.DocumentResource(current.ID, current.AuthorSub),
	); err != nil {
		return nil, err
	}
	if err := c.svc.DeleteDoc(ctx, req.ID); err != nil {
		return nil, err
	}
	writeNoContent(ctx)
	return &v1.DeleteDocRes{}, nil
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
		BadgeText:      m.BadgeText,
		BadgeIcon:      m.BadgeIcon,
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
