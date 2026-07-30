package catalog

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/yueli-official/docs/api/internal/docserr"
	"github.com/yueli-official/docs/api/internal/model"
)

const (
	defaultManageDocsPage = 1
	defaultManageDocsSize = 30
	maxManageDocsSize     = 100
)

var manageDocsStatuses = map[string]bool{"all": true, "draft": true, "published": true, "archived": true}
var manageDocsQualities = map[string]bool{"all": true, "issues": true}
var manageDocsSorts = map[string]bool{"updatedAt": true, "title": true, "path": true, "sortOrder": true}
var manageDocsDirections = map[string]bool{"asc": true, "desc": true}

type ManageDocsInput struct {
	Q            string
	Status       string
	Quality      string
	CollectionID string
	Version      string
	Locale       string
	ParentID     string
	Sort         string
	Direction    string
	Page         int
	Size         int
	OwnerSub     string
}

func normalizeManageDocsQuery(input ManageDocsInput) (model.ManageDocsQuery, error) {
	query := model.ManageDocsQuery{
		Q:            strings.TrimSpace(input.Q),
		Status:       input.Status,
		Quality:      input.Quality,
		CollectionID: strings.TrimSpace(input.CollectionID),
		Version:      strings.TrimSpace(input.Version),
		Locale:       strings.TrimSpace(input.Locale),
		ParentID:     strings.TrimSpace(input.ParentID),
		Sort:         input.Sort,
		Direction:    input.Direction,
		Page:         input.Page,
		Size:         input.Size,
		OwnerSub:     strings.TrimSpace(input.OwnerSub),
	}
	if query.Status == "" {
		query.Status = "all"
	}
	if query.Quality == "" {
		query.Quality = "all"
	}
	if query.Sort == "" {
		query.Sort = "updatedAt"
	}
	if query.Direction == "" {
		query.Direction = "desc"
	}
	if query.Page == 0 {
		query.Page = defaultManageDocsPage
	}
	if query.Size == 0 {
		query.Size = defaultManageDocsSize
	}

	switch {
	case len(query.Q) > 200:
		return model.ManageDocsQuery{}, fmt.Errorf("q must not exceed 200 characters")
	case !manageDocsStatuses[query.Status]:
		return model.ManageDocsQuery{}, fmt.Errorf("unsupported status %q", query.Status)
	case !manageDocsQualities[query.Quality]:
		return model.ManageDocsQuery{}, fmt.Errorf("unsupported quality %q", query.Quality)
	case !manageDocsSorts[query.Sort]:
		return model.ManageDocsQuery{}, fmt.Errorf("unsupported sort %q", query.Sort)
	case !manageDocsDirections[query.Direction]:
		return model.ManageDocsQuery{}, fmt.Errorf("unsupported direction %q", query.Direction)
	case query.Page < 1:
		return model.ManageDocsQuery{}, fmt.Errorf("page must be positive")
	case query.Size < 1 || query.Size > maxManageDocsSize:
		return model.ManageDocsQuery{}, fmt.Errorf("size must be between 1 and %d", maxManageDocsSize)
	case query.Version != "" && query.CollectionID == "":
		return model.ManageDocsQuery{}, fmt.Errorf("version requires collectionId")
	case query.CollectionID != "" && uuid.Validate(query.CollectionID) != nil:
		return model.ManageDocsQuery{}, fmt.Errorf("collectionId must be a UUID")
	case query.ParentID != "" && query.ParentID != "root" && uuid.Validate(query.ParentID) != nil:
		return model.ManageDocsQuery{}, fmt.Errorf("parentId must be root or a UUID")
	}
	return query, nil
}

func (s *Service) ManageDocs(ctx context.Context, input ManageDocsInput) (*model.ManageDocsResult, error) {
	query, err := normalizeManageDocsQuery(input)
	if err != nil {
		return nil, docserr.InvalidInput(err.Error())
	}
	return s.dao.ManageDocs(ctx, query)
}
