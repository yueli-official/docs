package catalog

import (
	"context"
	"regexp"
	"strings"

	"github.com/yueli-official/docs/api/internal/assetclient"
	"github.com/yueli-official/docs/api/internal/dao"
	"github.com/yueli-official/docs/api/internal/docserr"
	"github.com/yueli-official/docs/api/internal/model"
	"github.com/yueli-official/foundation/go/identifier"
)

var semanticVersionPattern = regexp.MustCompile(`^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)$`)

type CloneCollectionReleaseInput struct {
	SourceCollectionID    string
	SourceSemanticVersion string
	TargetSemanticVersion string
	Title                 string
	Slug                  string
}

func (s *Service) InitializeCollectionRelease(ctx context.Context, collectionID, semanticVersion string) (*model.Collection, error) {
	collection, err := s.dao.GetCollectionByID(ctx, collectionID)
	if err != nil {
		return nil, err
	}
	if collection == nil {
		return nil, docserr.NotFound(collectionID)
	}
	semanticVersion = strings.TrimSpace(semanticVersion)
	if !semanticVersionPattern.MatchString(semanticVersion) {
		return nil, docserr.InvalidInput("semantic version must use major.minor.patch")
	}
	if collection.ReleaseFamilyID != "" {
		if collection.SemanticVersion == semanticVersion {
			return collection, nil
		}
		return nil, docserr.InvalidInput("semantic version is immutable after initialization")
	}
	familyID := identifier.MustNew().String()
	if err := s.dao.InitializeCollectionRelease(
		ctx, collection.ID, familyID, collection.Title, semanticVersion,
		s.urlReconcileHook(collection.ID, "docs collection release initialized"),
	); err != nil {
		return nil, err
	}
	return s.dao.GetCollectionByID(ctx, collection.ID)
}

func (s *Service) CollectionReleases(ctx context.Context, slug string) ([]*model.Collection, error) {
	collection, err := s.GetCollectionBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	if collection.ReleaseFamilyID == "" {
		return []*model.Collection{collection}, nil
	}
	return s.dao.ListCollectionReleases(ctx, collection.ReleaseFamilyID)
}

func (s *Service) CloneCollectionAsRelease(
	ctx context.Context,
	authorSub, bearer string,
	in CloneCollectionReleaseInput,
) (*model.Collection, error) {
	source, err := s.dao.GetCollectionByID(ctx, in.SourceCollectionID)
	if err != nil {
		return nil, err
	}
	if source == nil {
		return nil, docserr.NotFound(in.SourceCollectionID)
	}
	targetVersion := strings.TrimSpace(in.TargetSemanticVersion)
	if !semanticVersionPattern.MatchString(targetVersion) {
		return nil, docserr.InvalidInput("target semantic version must use major.minor.patch")
	}
	sourceVersion := source.SemanticVersion
	if source.ReleaseFamilyID == "" {
		sourceVersion = strings.TrimSpace(in.SourceSemanticVersion)
		if !semanticVersionPattern.MatchString(sourceVersion) {
			return nil, docserr.InvalidInput("source semantic version must use major.minor.patch")
		}
	}
	if sourceVersion == targetVersion {
		return nil, docserr.InvalidInput("target semantic version must differ from source")
	}
	if source.ReleaseFamilyID != "" {
		releases, err := s.dao.ListCollectionReleases(ctx, source.ReleaseFamilyID)
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			if release.SemanticVersion == targetVersion {
				return nil, docserr.InvalidInput("target semantic version already exists in this release family")
			}
		}
	}

	slugSource := in.Slug
	if strings.TrimSpace(slugSource) == "" {
		slugSource = in.Title
	}
	targetSlug := slugify(slugSource)
	if targetSlug == "" {
		return nil, docserr.InvalidInput("slug produces an empty value")
	}
	if existing, err := s.dao.GetCollectionBySlug(ctx, targetSlug); err != nil {
		return nil, err
	} else if existing != nil {
		return nil, docserr.SlugTaken(targetSlug)
	}
	title := strings.TrimSpace(in.Title)
	if title == "" {
		return nil, docserr.InvalidInput("title required")
	}
	internalVersion, err := s.dao.GetDefaultCollectionVersion(ctx, source.ID)
	if err != nil {
		return nil, err
	}
	if internalVersion == nil {
		return nil, docserr.InvalidInput("source collection has no storage partition")
	}

	familyID := source.ReleaseFamilyID
	if familyID == "" {
		familyID = identifier.MustNew().String()
	}
	target := &model.Collection{
		ID: identifier.MustNew().String(), Slug: targetSlug, Title: title,
		Description: source.Description, CoverAssetID: source.CoverAssetID,
		CoverURL: source.CoverURL, Icon: source.Icon, SortOrder: source.SortOrder,
		AuthorSub: authorSub, ReleaseFamilyID: familyID,
		ReleaseFamilyName: source.ReleaseFamilyName, SemanticVersion: targetVersion,
		DerivedFromCollectionID: source.ID,
	}
	if target.ReleaseFamilyName == "" {
		target.ReleaseFamilyName = source.Title
	}
	targetInternalVersion := &model.CollectionVersion{
		ID: identifier.MustNew().String(), CollectionID: target.ID,
		Key: "default", Label: "默认版本", Status: "published", IsDefault: true,
	}
	if _, err := s.dao.CloneCollectionAsRelease(ctx, dao.CloneCollectionReleaseInput{
		Source: source, SourceInternalVersion: internalVersion.ID,
		SourceSemanticVersion: sourceVersion, Target: target,
		TargetInternalVersion: targetInternalVersion,
		FamilyID:              familyID, FamilyName: target.ReleaseFamilyName,
	}, s.urlReconcileHook(target.ID, "docs collection release cloned")); err != nil {
		return nil, err
	}
	if target.CoverAssetID != "" && s.asset != nil && bearer != "" {
		_ = s.asset.RegisterReference(ctx, bearer, assetclient.ReferenceInput{
			AssetID: target.CoverAssetID, RefType: "collection-cover", RefID: target.ID,
			RefLabel: target.Title, RefURL: "/" + target.Slug,
		})
	}
	return s.dao.GetCollectionByID(ctx, target.ID)
}
