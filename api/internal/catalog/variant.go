package catalog

import (
	"context"
	"sort"

	"github.com/yueli-official/docs/api/internal/docserr"
	"github.com/yueli-official/docs/api/internal/model"
)

type VariantCatalog struct {
	Locales  []*model.CollectionLocale
	Versions []*model.CollectionVersion
}

type VariantResolution struct {
	Path           string
	Locale         string
	Version        string
	TranslationKey string
	Fallback       string
}

func (s *Service) PublicVariants(ctx context.Context, collectionSlug string) (*VariantCatalog, error) {
	collection, err := s.GetCollectionBySlug(ctx, collectionSlug)
	if err != nil {
		return nil, err
	}
	locales, err := s.dao.ListCollectionLocales(ctx, collection.ID, true)
	if err != nil {
		return nil, err
	}
	versions, err := s.dao.ListCollectionVersions(ctx, collection.ID)
	if err != nil {
		return nil, err
	}
	publicVersions := make([]*model.CollectionVersion, 0, len(versions))
	for _, version := range versions {
		if version.Status == "published" || version.Status == "archived" {
			publicVersions = append(publicVersions, version)
		}
	}
	sort.SliceStable(publicVersions, func(left, right int) bool {
		if publicVersions[left].IsDefault != publicVersions[right].IsDefault {
			return publicVersions[left].IsDefault
		}
		if (publicVersions[left].Status == "archived") != (publicVersions[right].Status == "archived") {
			return publicVersions[right].Status == "archived"
		}
		if publicVersions[left].SortOrder != publicVersions[right].SortOrder {
			return publicVersions[left].SortOrder < publicVersions[right].SortOrder
		}
		return publicVersions[left].Key < publicVersions[right].Key
	})
	return &VariantCatalog{Locales: locales, Versions: publicVersions}, nil
}

func (s *Service) ResolveDocumentVariant(ctx context.Context, collectionSlug, translationKey, locale, versionKey string) (*VariantResolution, error) {
	collection, err := s.GetCollectionBySlug(ctx, collectionSlug)
	if err != nil {
		return nil, err
	}
	version, err := s.ResolveVersion(ctx, collection.ID, versionKey, true)
	if err != nil {
		return nil, err
	}
	defaultLocale, err := s.dao.GetDefaultCollectionLocale(ctx, collection.ID)
	if err != nil {
		return nil, err
	}
	if defaultLocale == nil {
		return nil, docserr.InvalidInput("collection has no default documentation locale")
	}
	selectedLocale, err := s.dao.GetCollectionLocale(ctx, collection.ID, locale)
	if err != nil {
		return nil, err
	}
	localeFallback := false
	if selectedLocale == nil || !selectedLocale.Enabled {
		selectedLocale = defaultLocale
		localeFallback = true
	}
	result := &VariantResolution{
		Locale:         selectedLocale.Locale,
		Version:        version.Key,
		TranslationKey: translationKey,
	}
	if localeFallback {
		result.Fallback = "default_locale"
	}
	if translationKey == "" {
		if result.Fallback == "" {
			result.Fallback = "collection"
		}
		return result, nil
	}
	doc, err := s.dao.GetPublishedDocByLogicalKey(ctx, collection.ID, version.ID, selectedLocale.Locale, translationKey)
	if err != nil {
		return nil, err
	}
	if doc == nil && selectedLocale.Locale != defaultLocale.Locale {
		doc, err = s.dao.GetPublishedDocByLogicalKey(ctx, collection.ID, version.ID, defaultLocale.Locale, translationKey)
		if err != nil {
			return nil, err
		}
		if doc != nil {
			result.Locale = defaultLocale.Locale
			result.Fallback = "default_locale"
		}
	}
	if doc == nil {
		result.Fallback = "collection"
		return result, nil
	}
	path, err := s.dao.GetDocPath(ctx, doc.ID)
	if err != nil {
		return nil, err
	}
	result.Path = path
	return result, nil
}
