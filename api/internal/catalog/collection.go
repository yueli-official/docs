package catalog

import (
	"context"
	"strings"

	"github.com/yueli-official/docs/api/internal/assetclient"
	"github.com/yueli-official/docs/api/internal/dao"
	"github.com/yueli-official/docs/api/internal/docserr"
	"github.com/yueli-official/docs/api/internal/model"
	"github.com/yueli-official/foundation/go/identifier"
)

// CreateCollection generates a slug from the title, guards duplicates, and
// inserts a new collection owned by authorSub.
func (s *Service) CreateCollection(ctx context.Context, authorSub, title, slug, description, cover, icon string) (*model.Collection, error) {
	return s.CreateCollectionWithSetup(ctx, CreateCollectionInput{
		AuthorSub: authorSub, Title: title, Slug: slug, Description: description,
		Cover: cover, Icon: icon, DefaultLocale: "en",
	})
}

type CreateCollectionInput struct {
	AuthorSub, Title, Slug, Description, Cover, Icon string
	DefaultLocale, SemanticVersion                   string
}

func (s *Service) CreateCollectionWithSetup(ctx context.Context, in CreateCollectionInput) (*model.Collection, error) {
	title, slug := strings.TrimSpace(in.Title), in.Slug
	if title == "" {
		return nil, docserr.InvalidInput("title required")
	}
	slugSource := title
	if strings.TrimSpace(slug) != "" {
		slugSource = slug
	}
	slug = slugify(slugSource)
	if slug == "" {
		return nil, docserr.InvalidInput("slug produces an empty value")
	}
	if existing, err := s.dao.GetCollectionBySlug(ctx, slug); err != nil {
		return nil, err
	} else if existing != nil {
		return nil, docserr.SlugTaken(slug)
	}
	m := &model.Collection{
		ID:          identifier.MustNew().String(),
		Slug:        slug,
		Title:       title,
		Description: in.Description,
		CoverURL:    in.Cover,
		Icon:        in.Icon,
		AuthorSub:   in.AuthorSub,
	}
	locale := strings.TrimSpace(in.DefaultLocale)
	if !localePattern.MatchString(locale) {
		return nil, docserr.InvalidInput("invalid documentation locale")
	}
	semanticVersion := strings.TrimSpace(in.SemanticVersion)
	if semanticVersion != "" {
		if !semanticVersionPattern.MatchString(semanticVersion) {
			return nil, docserr.InvalidInput("semantic version must use major.minor.patch")
		}
		m.ReleaseFamilyID = identifier.MustNew().String()
		m.ReleaseFamilyName = title
		m.SemanticVersion = semanticVersion
	}
	version := &model.CollectionVersion{
		ID:           identifier.MustNew().String(),
		CollectionID: m.ID,
		Key:          "default",
		Label:        "默认版本",
		Status:       "published",
		IsDefault:    true,
	}
	if err := s.dao.InsertCollectionWithDefaultVersion(
		ctx, m, version, &model.CollectionLocale{
			CollectionID: m.ID, Locale: locale, Label: localeLabel(locale), HTMLLang: locale,
			Direction: localeDirection(locale), IsDefault: true, Enabled: true,
		}, s.urlReconcileHook(m.ID, "docs collection created"),
	); err != nil {
		return nil, err
	}
	return s.dao.GetCollectionByID(ctx, m.ID)
}

func localeLabel(locale string) string {
	switch locale {
	case "zh-CN":
		return "简体中文"
	case "zh-TW":
		return "繁體中文"
	case "en", "en-US":
		return "English"
	case "ja", "ja-JP":
		return "日本語"
	default:
		return locale
	}
}

func localeDirection(locale string) string {
	base := strings.ToLower(strings.SplitN(locale, "-", 2)[0])
	if base == "ar" || base == "fa" || base == "he" || base == "ur" {
		return "rtl"
	}
	return "ltr"
}

// ListCollections returns all collections ordered by sort_order then created_at.
func (s *Service) ListCollections(ctx context.Context) ([]*model.Collection, error) {
	return s.dao.ListCollections(ctx)
}

// GetCollectionBySlug returns the collection with the given slug, or a
// NotFound error when no such collection exists.
func (s *Service) GetCollectionBySlug(ctx context.Context, slug string) (*model.Collection, error) {
	c, err := s.dao.GetCollectionBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, docserr.NotFound(slug)
	}
	return c, nil
}

// UpdateCollection overwrites the mutable fields of the collection identified by id.
func (s *Service) UpdateCollection(ctx context.Context, id, title, slug, description, cover, icon string) (*model.Collection, error) {
	return s.UpdateCollectionWithBearer(ctx, id, "", title, slug, description, cover, icon)
}

func (s *Service) UpdateCollectionWithBearer(ctx context.Context, id, bearer, title, slug, description, cover, icon string) (*model.Collection, error) {
	c, err := s.dao.GetCollectionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, docserr.NotFound(id)
	}
	if strings.TrimSpace(slug) != "" {
		nextSlug := slugify(slug)
		if nextSlug == "" {
			return nil, docserr.InvalidInput("slug produces an empty value")
		}
		if existing, err := s.dao.GetCollectionBySlug(ctx, nextSlug); err != nil {
			return nil, err
		} else if existing != nil && existing.ID != id {
			return nil, docserr.SlugTaken(nextSlug)
		}
		c.Slug = nextSlug
	}
	oldAssetID := c.CoverAssetID
	oldCoverURL := c.CoverURL
	c.Title, c.Description, c.CoverURL, c.Icon = title, description, cover, icon
	if oldCoverURL != cover {
		c.CoverAssetID = ""
	}
	if err := s.dao.UpdateCollectionWithHook(
		ctx, c, s.urlReconcileHook(c.ID, "docs collection updated"),
	); err != nil {
		return nil, err
	}
	if oldAssetID != "" && oldCoverURL != cover && s.asset != nil && bearer != "" {
		_ = s.asset.UnregisterReference(ctx, bearer, assetclient.ReferenceInput{
			AssetID: oldAssetID, RefType: "collection-cover", RefID: id,
		})
	}
	return s.dao.GetCollectionByID(ctx, id)
}

// AddCollectionCover opens an upload for a collection cover image.
func (s *Service) AddCollectionCover(ctx context.Context, id, bearer, filename, mime string, size int64) (assetclient.InitOutput, error) {
	if s.asset == nil {
		return assetclient.InitOutput{}, docserr.UpstreamFailed("asset client not configured")
	}
	c, err := s.dao.GetCollectionByID(ctx, id)
	if err != nil {
		return assetclient.InitOutput{}, err
	}
	if c == nil {
		return assetclient.InitOutput{}, docserr.NotFound(id)
	}
	if mime == "" {
		mime = "application/octet-stream"
	}
	return s.asset.UploadInit(ctx, bearer, assetclient.InitInput{
		Filename: filename, Mime: mime, Size: size, Category: s.coverCategory, Visibility: "public",
	})
}

// FinalizeCollectionCover finalizes the uploaded image and stores its public URL
// on the collection.
func (s *Service) FinalizeCollectionCover(ctx context.Context, id, bearer, uploadToken string) (*model.Collection, error) {
	if s.asset == nil {
		return nil, docserr.UpstreamFailed("asset client not configured")
	}
	c, err := s.dao.GetCollectionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, docserr.NotFound(id)
	}
	view, err := s.asset.Finalize(ctx, bearer, uploadToken)
	if err != nil {
		return nil, err
	}
	old := c.CoverAssetID
	coverURL, err := publicImageURL(view, "cover")
	if err != nil {
		return nil, err
	}
	c.CoverURL = coverURL
	c.CoverAssetID = view.ID
	if err := s.dao.UpdateCollection(ctx, c); err != nil {
		return nil, err
	}
	if err := s.asset.RegisterReference(ctx, bearer, assetclient.ReferenceInput{
		AssetID: view.ID, RefType: "collection-cover", RefID: c.ID,
		RefLabel: c.Title, RefURL: "/" + c.Slug,
	}); err != nil {
		return nil, err
	}
	if old != "" && old != view.ID {
		_ = s.asset.UnregisterReference(ctx, bearer, assetclient.ReferenceInput{
			AssetID: old, RefType: "collection-cover", RefID: c.ID,
		})
	}
	return s.dao.GetCollectionByID(ctx, id)
}

// DeleteCollection removes the collection (and its docs via CASCADE).
func (s *Service) DeleteCollection(ctx context.Context, id string) error {
	return s.DeleteCollectionWithBearer(ctx, id, "")
}

func (s *Service) DeleteCollectionWithBearer(ctx context.Context, id, bearer string) error {
	c, err := s.dao.GetCollectionByID(ctx, id)
	if err != nil {
		return err
	}
	if c != nil && c.CoverAssetID != "" && s.asset != nil && bearer != "" {
		_ = s.asset.UnregisterReference(ctx, bearer, assetclient.ReferenceInput{
			AssetID: c.CoverAssetID, RefType: "collection-cover", RefID: id,
		})
	}
	var searchHook dao.TransactionHook
	if s.search != nil {
		searchHook = s.search.DeleteCollectionHook(id)
	}
	return s.dao.DeleteCollectionWithHooks(
		ctx, id, searchHook, s.urlReconcileHook(id, "docs collection deleted"),
	)
}
