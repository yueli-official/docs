package docscomments

import (
	"context"
	"sort"
	"strings"
	"sync"
	"time"
)

type Memory struct {
	mu        sync.RWMutex
	documents map[string]Document
	comments  map[string]*Comment
}

func NewMemory() *Memory {
	return &Memory{
		documents: map[string]Document{},
		comments:  map[string]*Comment{},
	}
}

func (store *Memory) SeedDocument(document Document) {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.documents[document.ID] = document
}

func (store *Memory) PublishedDocument(
	_ context.Context,
	id string,
) (Document, bool, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	document, exists := store.documents[id]
	return document, exists, nil
}

func (store *Memory) Comment(_ context.Context, id string) (*Comment, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return cloneComment(store.comments[id]), nil
}

func (store *Memory) Insert(_ context.Context, comment *Comment) error {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.comments[comment.ID] = cloneComment(comment)
	return nil
}

func (store *Memory) ApprovedTop(
	_ context.Context,
	documentID string,
	limit int,
	offset int,
) ([]*Comment, int, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	comments := []*Comment{}
	for _, comment := range store.comments {
		if comment.DocumentID == documentID && comment.ParentID == "" && comment.Status == StatusApproved {
			comments = append(comments, cloneComment(comment))
		}
	}
	sort.Slice(comments, func(left, right int) bool {
		return comments[left].CreatedAt.After(comments[right].CreatedAt)
	})
	return pageComments(comments, limit, offset), len(comments), nil
}

func (store *Memory) ApprovedReplies(
	_ context.Context,
	parentIDs []string,
) (map[string][]*Comment, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	parents := map[string]bool{}
	for _, id := range parentIDs {
		parents[id] = true
	}
	result := map[string][]*Comment{}
	for _, comment := range store.comments {
		if parents[comment.ParentID] && comment.Status == StatusApproved {
			result[comment.ParentID] = append(result[comment.ParentID], cloneComment(comment))
		}
	}
	for parent := range result {
		sort.Slice(result[parent], func(left, right int) bool {
			return result[parent][left].CreatedAt.Before(result[parent][right].CreatedAt)
		})
	}
	return result, nil
}

func (store *Memory) AdminComments(
	_ context.Context,
	query AdminQuery,
) ([]*AdminComment, int, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	keyword := strings.ToLower(query.Q)
	items := []*AdminComment{}
	for _, comment := range store.comments {
		if query.Status != "" && comment.Status != query.Status {
			continue
		}
		document := store.documents[comment.DocumentID]
		if keyword != "" && !strings.Contains(strings.ToLower(strings.Join([]string{
			comment.Content, comment.AuthorName, document.Title,
		}, " ")), keyword) {
			continue
		}
		items = append(items, &AdminComment{
			Comment: cloneComment(comment), Document: document,
		})
	}
	sort.Slice(items, func(left, right int) bool {
		if query.SortOrder == "asc" {
			return items[left].Comment.CreatedAt.Before(items[right].Comment.CreatedAt)
		}
		return items[left].Comment.CreatedAt.After(items[right].Comment.CreatedAt)
	})
	total := len(items)
	start := (query.Page - 1) * query.Size
	if start >= total {
		return []*AdminComment{}, total, nil
	}
	end := min(start+query.Size, total)
	return items[start:end], total, nil
}

func (store *Memory) SetStatus(
	_ context.Context,
	id string,
	status Status,
	updatedAt time.Time,
) (*Comment, error) {
	store.mu.Lock()
	defer store.mu.Unlock()
	comment := store.comments[id]
	if comment == nil {
		return nil, nil
	}
	comment.Status = status
	comment.UpdatedAt = updatedAt
	return cloneComment(comment), nil
}

func (store *Memory) SoftDelete(
	_ context.Context,
	id string,
	_ time.Time,
) (bool, error) {
	store.mu.Lock()
	defer store.mu.Unlock()
	if store.comments[id] == nil {
		return false, nil
	}
	delete(store.comments, id)
	for commentID, comment := range store.comments {
		if comment.ParentID == id {
			delete(store.comments, commentID)
		}
	}
	return true, nil
}

func cloneComment(comment *Comment) *Comment {
	if comment == nil {
		return nil
	}
	copy := *comment
	return &copy
}

func pageComments(comments []*Comment, limit, offset int) []*Comment {
	if offset >= len(comments) {
		return []*Comment{}
	}
	return comments[offset:min(offset+limit, len(comments))]
}
