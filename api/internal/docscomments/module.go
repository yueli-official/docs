// Package docscomments owns document comment threads and moderation.
package docscomments

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/yueli-official/foundation/go/identifier"
)

type Status string

const (
	StatusApproved Status = "approved"
	StatusPending  Status = "pending"
	StatusSpam     Status = "spam"
	StatusTrash    Status = "trash"
)

var (
	ErrDocumentNotFound = errors.New("document not found")
	ErrCommentNotFound  = errors.New("comment not found")
	ErrInvalidInput     = errors.New("invalid comment input")
)

type Document struct {
	ID             string
	Title          string
	CollectionSlug string
	VersionKey     string
	SlugPath       string
	Locale         string
}

type Comment struct {
	ID         string
	DocumentID string
	ParentID   string
	UserSub    string
	AuthorName string
	Content    string
	Status     Status
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Thread struct {
	Comment *Comment
	Replies []*Comment
}

type AdminComment struct {
	Comment  *Comment
	Document Document
}

type AdminQuery struct {
	Status    Status
	Q         string
	SortBy    string
	SortOrder string
	Page      int
	Size      int
}

type AdminResult struct {
	Items []*AdminComment
	Total int
	Page  int
	Size  int
}

type Store interface {
	PublishedDocument(context.Context, string) (Document, bool, error)
	Comment(context.Context, string) (*Comment, error)
	Insert(context.Context, *Comment) error
	ApprovedTop(context.Context, string, int, int) ([]*Comment, int, error)
	ApprovedReplies(context.Context, []string) (map[string][]*Comment, error)
	AdminComments(context.Context, AdminQuery) ([]*AdminComment, int, error)
	SetStatus(context.Context, string, Status, time.Time) (*Comment, error)
	SoftDelete(context.Context, string, time.Time) (bool, error)
}

type Module struct {
	store Store
	clock func() time.Time
}

func New(store Store) *Module {
	return &Module{store: store, clock: time.Now}
}

func (module *Module) List(
	ctx context.Context,
	documentID string,
	page int,
	size int,
) ([]*Thread, int, int, int, error) {
	page, size = normalizePage(page, size)
	document, exists, err := module.store.PublishedDocument(ctx, documentID)
	if err != nil {
		return nil, 0, page, size, err
	}
	if !exists || document.ID == "" {
		return nil, 0, page, size, ErrDocumentNotFound
	}
	top, total, err := module.store.ApprovedTop(ctx, documentID, size, (page-1)*size)
	if err != nil {
		return nil, 0, page, size, err
	}
	ids := make([]string, 0, len(top))
	for _, comment := range top {
		ids = append(ids, comment.ID)
	}
	replies, err := module.store.ApprovedReplies(ctx, ids)
	if err != nil {
		return nil, 0, page, size, err
	}
	threads := make([]*Thread, 0, len(top))
	for _, comment := range top {
		threads = append(threads, &Thread{Comment: comment, Replies: replies[comment.ID]})
	}
	return threads, total, page, size, nil
}

func (module *Module) Create(
	ctx context.Context,
	documentID string,
	userSub string,
	authorName string,
	content string,
	parentID string,
) (*Comment, error) {
	content = strings.TrimSpace(content)
	if count := len([]rune(content)); count < 1 || count > 2000 {
		return nil, ErrInvalidInput
	}
	document, exists, err := module.store.PublishedDocument(ctx, documentID)
	if err != nil {
		return nil, err
	}
	if !exists || document.ID == "" {
		return nil, ErrDocumentNotFound
	}
	comment := &Comment{
		ID: identifier.MustNew().String(), DocumentID: documentID,
		UserSub: strings.TrimSpace(userSub), AuthorName: strings.TrimSpace(authorName),
		Content: content, Status: StatusApproved,
	}
	if comment.UserSub == "" {
		return nil, ErrInvalidInput
	}
	if comment.AuthorName == "" {
		comment.AuthorName = memberLabel(comment.UserSub)
	}
	if parentID = strings.TrimSpace(parentID); parentID != "" {
		parent, err := module.store.Comment(ctx, parentID)
		if err != nil {
			return nil, err
		}
		if parent == nil || parent.DocumentID != documentID || parent.Status != StatusApproved {
			return nil, ErrInvalidInput
		}
		if parent.ParentID != "" {
			comment.ParentID = parent.ParentID
		} else {
			comment.ParentID = parent.ID
		}
	}
	now := module.clock().UTC()
	comment.CreatedAt = now
	comment.UpdatedAt = now
	if err := module.store.Insert(ctx, comment); err != nil {
		return nil, err
	}
	return comment, nil
}

func (module *Module) Manage(
	ctx context.Context,
	query AdminQuery,
) (*AdminResult, error) {
	query.Page, query.Size = normalizePage(query.Page, query.Size)
	query.Q = strings.TrimSpace(query.Q)
	query.SortBy = strings.TrimSpace(query.SortBy)
	if query.SortBy == "" {
		query.SortBy = "createdAt"
	}
	query.SortOrder = strings.ToLower(strings.TrimSpace(query.SortOrder))
	if query.SortOrder == "" {
		query.SortOrder = "desc"
	}
	if query.Status != "" && !validStatus(query.Status) {
		return nil, ErrInvalidInput
	}
	if query.SortBy != "createdAt" || (query.SortOrder != "asc" && query.SortOrder != "desc") {
		return nil, ErrInvalidInput
	}
	items, total, err := module.store.AdminComments(ctx, query)
	if err != nil {
		return nil, err
	}
	return &AdminResult{
		Items: items, Total: total, Page: query.Page, Size: query.Size,
	}, nil
}

func (module *Module) Moderate(
	ctx context.Context,
	id string,
	status Status,
) (*Comment, error) {
	if !validStatus(status) {
		return nil, ErrInvalidInput
	}
	comment, err := module.store.SetStatus(ctx, id, status, module.clock().UTC())
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, ErrCommentNotFound
	}
	return comment, nil
}

func (module *Module) Delete(ctx context.Context, id string) error {
	deleted, err := module.store.SoftDelete(ctx, id, module.clock().UTC())
	if err != nil {
		return err
	}
	if !deleted {
		return ErrCommentNotFound
	}
	return nil
}

func validStatus(status Status) bool {
	switch status {
	case StatusApproved, StatusPending, StatusSpam, StatusTrash:
		return true
	default:
		return false
	}
}

func normalizePage(page, size int) (int, int) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	if size > 100 {
		size = 100
	}
	return page, size
}

func memberLabel(subject string) string {
	if len(subject) > 8 {
		return subject[:8]
	}
	return subject
}
