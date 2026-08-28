package catalog

import (
	"context"
	"regexp"
	"strings"

	"github.com/yueli-official/docs/api/internal/dao"
	"github.com/yueli-official/docs/api/internal/docserr"
	"github.com/yueli-official/docs/api/internal/model"
	"github.com/yueli-official/foundation/go/identifier"
)

var versionKeyPattern = regexp.MustCompile(`^(v[0-9][0-9a-z.-]*|next|latest|default)$`)

type CreateVersionInput struct {
	CollectionID    string
	Key             string
	Label           string
	Status          string
	SourceVersionID string
}

type UpdateVersionInput struct {
	CollectionID string
	VersionID    string
	Label        string
	Status       string
	IsDefault    bool
	SortOrder    int
}

func (s *Service) ListVersions(ctx context.Context, collectionID string) ([]*model.CollectionVersion, error) {
	return s.dao.ListCollectionVersions(ctx, collectionID)
}

func (s *Service) CreateVersion(ctx context.Context, in CreateVersionInput) (*model.CollectionVersion, error) {
	if !versionKeyPattern.MatchString(in.Key) {
		return nil, docserr.InvalidInput("invalid version key")
	}
	if in.Label == "" {
		in.Label = in.Key
	}
	if in.Status == "" {
		in.Status = "draft"
	}
	col, err := s.dao.GetCollectionByID(ctx, in.CollectionID)
	if err != nil {
		return nil, err
	}
	if col == nil {
		return nil, docserr.NotFound(in.CollectionID)
	}
	m := &model.CollectionVersion{
		ID:              identifier.MustNew().String(),
		CollectionID:    in.CollectionID,
		Key:             in.Key,
		Label:           in.Label,
		Status:          in.Status,
		IsDefault:       false,
		SourceVersionID: in.SourceVersionID,
	}
	if err := s.dao.InsertCollectionVersion(ctx, m); err != nil {
		return nil, docserr.SlugTaken(in.Key)
	}
	return s.dao.GetCollectionVersionByID(ctx, m.ID)
}

func (s *Service) UpdateVersion(ctx context.Context, in UpdateVersionInput) (*model.CollectionVersion, error) {
	in.Label = strings.TrimSpace(in.Label)
	if in.Label == "" {
		return nil, docserr.InvalidInput("version label is required")
	}
	if in.Status != "draft" && in.Status != "published" && in.Status != "archived" {
		return nil, docserr.InvalidInput("invalid version status")
	}
	if in.IsDefault && in.Status != "published" {
		return nil, docserr.InvalidInput("default version must be published")
	}
	value, err := s.dao.GetCollectionVersionByID(ctx, in.VersionID)
	if err != nil {
		return nil, err
	}
	if value == nil || value.CollectionID != in.CollectionID {
		return nil, docserr.NotFound(in.VersionID)
	}
	if value.IsDefault && !in.IsDefault {
		return nil, docserr.InvalidInput("choose another default version before demoting this version")
	}
	value.Label = in.Label
	value.Status = in.Status
	value.IsDefault = in.IsDefault
	value.SortOrder = in.SortOrder
	var searchHook dao.TransactionHook
	if s.search != nil {
		searchHook = s.search.CollectionHook(in.CollectionID)
	}
	if err := s.dao.UpdateCollectionVersion(ctx, value, dao.ComposeTransactionHooks(
		s.urlReconcileHook(in.CollectionID, "docs version updated"),
		searchHook,
	)); err != nil {
		return nil, err
	}
	return s.dao.GetCollectionVersionByID(ctx, value.ID)
}

func (s *Service) ResolveVersion(ctx context.Context, collectionID, key string, publicOnly bool) (*model.CollectionVersion, error) {
	var (
		v   *model.CollectionVersion
		err error
	)
	if key == "" {
		v, err = s.dao.GetDefaultCollectionVersion(ctx, collectionID)
	} else {
		v, err = s.dao.GetCollectionVersionByKey(ctx, collectionID, key)
	}
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, docserr.NotFound(key)
	}
	if publicOnly && v.Status != "published" && v.Status != "archived" {
		return nil, docserr.NotFound(key)
	}
	return v, nil
}
