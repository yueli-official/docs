package catalog

import (
	"context"
	"regexp"

	"github.com/google/uuid"

	"github.com/yueli-official/docs/api/internal/docserr"
	"github.com/yueli-official/docs/api/internal/model"
)

var versionKeyPattern = regexp.MustCompile(`^(v[0-9][0-9a-z.-]*|next|latest|default)$`)

type CreateVersionInput struct {
	CollectionID    string
	Key             string
	Label           string
	Status          string
	SourceVersionID string
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
		ID:              uuid.NewString(),
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
	if publicOnly && v.Status != "published" {
		return nil, docserr.NotFound(key)
	}
	return v, nil
}
