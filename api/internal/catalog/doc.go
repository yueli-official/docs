package catalog

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"platform/products/docs/api/internal/docserr"
	"platform/products/docs/api/internal/model"
)

// CreateDocInput carries the caller-supplied fields for a new doc.
type CreateDocInput struct {
	CollectionID   string
	VersionID      string
	ParentID       string
	Slug           string
	Title          string
	Content        string
	SEOTitle       string
	SEODescription string
	Locale         string
	TranslationKey string
	SortOrder      int
}

// PatchDocInput carries only the mutable fields present in a PATCH request.
type PatchDocInput struct {
	Title          *string
	Slug           *string
	Content        *string
	Excerpt        *string
	SEOTitle       *string
	SEODescription *string
	Status         *string
	Locale         *string
	VersionID      *string
	TranslationKey *string
	SortOrder      *int
	ParentID       *string
}

// CreateDoc validates the collection exists, derives a slug from the title or
// caller-supplied slug, and inserts the doc. A unique-constraint violation on
// (collection_id, locale, parent_id, slug) is mapped to docserr.SlugTaken.
func (s *Service) CreateDoc(ctx context.Context, authorSub string, in CreateDocInput) (*model.Doc, error) {
	if in.Title == "" {
		return nil, docserr.InvalidInput("title required")
	}
	col, err := s.dao.GetCollectionByID(ctx, in.CollectionID)
	if err != nil {
		return nil, err
	}
	if col == nil {
		return nil, docserr.NotFound(in.CollectionID)
	}
	locale := in.Locale
	if locale == "" {
		locale = "en"
	}
	versionID := in.VersionID
	if versionID == "" {
		v, err := s.ResolveVersion(ctx, in.CollectionID, "", false)
		if err != nil {
			return nil, err
		}
		versionID = v.ID
	}
	translationKey := in.TranslationKey
	if translationKey == "" {
		translationKey = uuid.NewString()
	}
	slugSource := in.Title
	if strings.TrimSpace(in.Slug) != "" {
		slugSource = in.Slug
	}
	slug := slugify(slugSource)
	if slug == "" {
		return nil, docserr.InvalidInput("slug produces empty value")
	}
	m := &model.Doc{
		ID:             uuid.NewString(),
		CollectionID:   in.CollectionID,
		VersionID:      versionID,
		ParentID:       in.ParentID,
		Slug:           slug,
		Title:          in.Title,
		Content:        in.Content,
		SEOTitle:       in.SEOTitle,
		SEODescription: in.SEODescription,
		Locale:         locale,
		TranslationKey: translationKey,
		SortOrder:      in.SortOrder,
		AuthorSub:      authorSub,
	}
	if err := s.validateDocParent(ctx, m, m.ParentID); err != nil {
		return nil, err
	}
	if err := s.dao.InsertDocWithHook(
		ctx, m, s.urlReconcileHook(m.CollectionID, "docs document created"),
	); err != nil {
		return nil, docserr.SlugTaken(slug)
	}
	return s.dao.GetDocByID(ctx, m.ID)
}

// GetDoc returns the doc by id, or NotFound when it does not exist.
func (s *Service) GetDoc(ctx context.Context, id string) (*model.Doc, error) {
	d, err := s.dao.GetDocByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, docserr.NotFound(id)
	}
	return d, nil
}

func (s *Service) GetPublishedDocByPath(ctx context.Context, collectionSlug, versionKey, path, locale string) (*model.Doc, error) {
	col, err := s.dao.GetCollectionBySlug(ctx, collectionSlug)
	if err != nil {
		return nil, err
	}
	if col == nil {
		return nil, docserr.NotFound(collectionSlug)
	}
	if locale == "" {
		locale = "en"
	}
	version, err := s.ResolveVersion(ctx, col.ID, versionKey, true)
	if err != nil {
		return nil, err
	}
	d, err := s.dao.GetPublishedDocByCollectionPath(ctx, col.ID, version.ID, locale, path)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, docserr.NotFound(path)
	}
	return d, nil
}

// SearchPublishedDocs returns published docs matching q, optionally scoped to a collection.
func (s *Service) SearchPublishedDocs(ctx context.Context, collectionSlug, versionKey, locale, q string) (*model.SearchResult, error) {
	if locale == "" {
		locale = "en"
	}
	collectionID := ""
	versionID := ""
	if collectionSlug != "" {
		col, err := s.dao.GetCollectionBySlug(ctx, collectionSlug)
		if err != nil {
			return nil, err
		}
		if col == nil {
			return nil, docserr.NotFound(collectionSlug)
		}
		collectionID = col.ID
		version, err := s.ResolveVersion(ctx, col.ID, versionKey, true)
		if err != nil {
			return nil, err
		}
		versionID = version.ID
	}
	result, err := s.dao.SearchPublishedDocs(ctx, collectionID, versionID, locale, q, 50)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(q) != "" {
		_ = s.dao.InsertSearchEvent(ctx, strings.TrimSpace(q), collectionSlug, locale, result.Total)
	}
	return result, nil
}

// ListDocs returns all non-deleted docs in the given collection + locale.
func (s *Service) ListDocs(ctx context.Context, collectionID, versionKey, locale string) ([]*model.Doc, error) {
	version, err := s.ResolveVersion(ctx, collectionID, versionKey, false)
	if err != nil {
		return nil, err
	}
	return s.dao.ListDocsByCollection(ctx, collectionID, version.ID, locale)
}

// UpdateDoc overwrites every mutable field on an existing doc.
func (s *Service) UpdateDoc(ctx context.Context, id, title, content, excerpt, status string, sortOrder int, parentID string) (*model.Doc, error) {
	d, err := s.dao.GetDocByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, docserr.NotFound(id)
	}
	d.Title = title
	d.Content = content
	d.Excerpt = excerpt
	d.Status = status
	d.SortOrder = sortOrder
	d.ParentID = parentID
	if err := s.validateDocParent(ctx, d, d.ParentID); err != nil {
		return nil, err
	}
	if err := s.dao.UpdateDocWithHook(
		ctx, d, s.urlReconcileHook(d.CollectionID, "docs document updated"),
	); err != nil {
		return nil, err
	}
	return s.dao.GetDocByID(ctx, id)
}

// PatchDoc updates only the fields explicitly supplied by the caller.
func (s *Service) PatchDoc(ctx context.Context, id string, in PatchDocInput) (*model.Doc, error) {
	d, err := s.dao.GetDocByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, docserr.NotFound(id)
	}
	if in.Title != nil {
		title := strings.TrimSpace(*in.Title)
		if title == "" {
			return nil, docserr.InvalidInput("title required")
		}
		d.Title = title
	}
	slugStr := ""
	if in.Slug != nil {
		slugStr = slugify(*in.Slug)
		if slugStr == "" {
			return nil, docserr.InvalidInput("slug produces empty value")
		}
		d.Slug = slugStr
	}
	if in.Content != nil {
		d.Content = *in.Content
	}
	if in.Excerpt != nil {
		d.Excerpt = *in.Excerpt
	}
	if in.SEOTitle != nil {
		d.SEOTitle = *in.SEOTitle
	}
	if in.SEODescription != nil {
		d.SEODescription = *in.SEODescription
	}
	if in.Status != nil {
		status := strings.TrimSpace(*in.Status)
		if status != "draft" && status != "published" && status != "archived" {
			return nil, docserr.InvalidInput("status must be draft, published or archived")
		}
		d.Status = status
	}
	if in.Locale != nil {
		d.Locale = *in.Locale
	}
	if in.VersionID != nil {
		d.VersionID = *in.VersionID
	}
	if in.TranslationKey != nil {
		d.TranslationKey = *in.TranslationKey
	}
	if in.SortOrder != nil {
		d.SortOrder = *in.SortOrder
	}
	if in.ParentID != nil {
		parentID := strings.TrimSpace(*in.ParentID)
		if err := s.validateDocParent(ctx, d, parentID); err != nil {
			return nil, err
		}
		d.ParentID = parentID
	}
	if (in.VersionID != nil || in.Locale != nil) && d.ParentID != "" {
		if err := s.validateDocParent(ctx, d, d.ParentID); err != nil {
			return nil, err
		}
	}
	if err := s.dao.UpdateDocWithHook(
		ctx, d, s.urlReconcileHook(d.CollectionID, "docs document patched"),
	); err != nil {
		if slugStr != "" {
			return nil, docserr.SlugTaken(slugStr)
		}
		return nil, err
	}
	return s.dao.GetDocByID(ctx, id)
}

func (s *Service) validateDocParent(ctx context.Context, doc *model.Doc, parentID string) error {
	if parentID == "" {
		return nil
	}
	if parentID == doc.ID {
		return docserr.InvalidInput("document cannot be its own parent")
	}
	seen := map[string]struct{}{doc.ID: {}}
	currentID := parentID
	for currentID != "" {
		if _, cycle := seen[currentID]; cycle {
			return docserr.InvalidInput("document parent would create a cycle")
		}
		seen[currentID] = struct{}{}
		parent, err := s.dao.GetDocByID(ctx, currentID)
		if err != nil {
			return err
		}
		if parent == nil {
			return docserr.InvalidInput("document parent not found")
		}
		if parent.CollectionID != doc.CollectionID || parent.VersionID != doc.VersionID || parent.Locale != doc.Locale {
			return docserr.InvalidInput("document parent must use the same collection, version and locale")
		}
		currentID = parent.ParentID
	}
	return nil
}

// PublishDoc marks a doc as visible to public readers.
func (s *Service) PublishDoc(ctx context.Context, id string) (*model.Doc, error) {
	status := "published"
	return s.PatchDoc(ctx, id, PatchDocInput{Status: &status})
}

// ArchiveDoc removes a doc from public reader surfaces without deleting it.
func (s *Service) ArchiveDoc(ctx context.Context, id string) (*model.Doc, error) {
	status := "archived"
	return s.PatchDoc(ctx, id, PatchDocInput{Status: &status})
}

// DeleteDoc soft-deletes the doc with the given id.
func (s *Service) DeleteDoc(ctx context.Context, id string) error {
	doc, err := s.dao.GetDocByID(ctx, id)
	if err != nil {
		return err
	}
	if doc == nil {
		return docserr.NotFound(id)
	}
	return s.dao.SoftDeleteDocWithHook(
		ctx, id, s.urlReconcileHook(doc.CollectionID, "docs document subtree deleted"),
	)
}

// BuildTree assembles a flat doc list (already ordered) into a parent_id
// forest, preserving input order among siblings. Arbitrary depth. Orphans
// (whose parent is absent from the list) surface at the root level.
func BuildTree(flat []*model.Doc) []*model.Doc {
	byID := make(map[string]*model.Doc, len(flat))
	for _, d := range flat {
		d.Children = nil
		byID[d.ID] = d
	}
	var roots []*model.Doc
	for _, d := range flat {
		if d.ParentID == "" {
			roots = append(roots, d)
			continue
		}
		if parent, ok := byID[d.ParentID]; ok {
			parent.Children = append(parent.Children, d)
		} else {
			roots = append(roots, d) // orphan → surface at root
		}
	}
	return roots
}

// PublicDocTree returns the published-only collection tree for anonymous readers.
func (s *Service) PublicDocTree(ctx context.Context, collectionSlug, versionKey, locale string) (*model.Collection, []*model.Doc, error) {
	col, err := s.dao.GetCollectionBySlug(ctx, collectionSlug)
	if err != nil {
		return nil, nil, err
	}
	if col == nil {
		return nil, nil, docserr.NotFound(collectionSlug)
	}
	if locale == "" {
		locale = "en"
	}
	version, err := s.ResolveVersion(ctx, col.ID, versionKey, true)
	if err != nil {
		return nil, nil, err
	}
	flat, err := s.dao.ListDocsByCollectionStatus(ctx, col.ID, version.ID, locale, "published")
	if err != nil {
		return nil, nil, err
	}
	return col, BuildTree(flat), nil
}

// ManageDocTree returns the full collection tree visible to authenticated managers.
func (s *Service) ManageDocTree(ctx context.Context, collectionSlug, versionKey, locale string) (*model.Collection, []*model.Doc, error) {
	col, err := s.dao.GetCollectionBySlug(ctx, collectionSlug)
	if err != nil {
		return nil, nil, err
	}
	if col == nil {
		return nil, nil, docserr.NotFound(collectionSlug)
	}
	if locale == "" {
		locale = "en"
	}
	version, err := s.ResolveVersion(ctx, col.ID, versionKey, false)
	if err != nil {
		return nil, nil, err
	}
	flat, err := s.dao.ListDocsByCollection(ctx, col.ID, version.ID, locale)
	if err != nil {
		return nil, nil, err
	}
	return col, BuildTree(flat), nil
}

// DocTree preserves the existing full-tree behavior for current manage callers.
func (s *Service) DocTree(ctx context.Context, collectionSlug, locale string) ([]*model.Doc, error) {
	_, tree, err := s.ManageDocTree(ctx, collectionSlug, "", locale)
	return tree, err
}
