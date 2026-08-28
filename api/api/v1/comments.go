package v1

import "github.com/gogf/gf/v2/frame/g"

type CommentView struct {
	ID         string         `json:"id"`
	ParentID   string         `json:"parentId,omitempty"`
	AuthorName string         `json:"authorName"`
	AvatarURL  string         `json:"avatarUrl,omitempty"`
	Content    string         `json:"content"`
	CreatedAt  string         `json:"createdAt"`
	Replies    []*CommentView `json:"replies,omitempty"`
}

type CommentAdminView struct {
	ID             string `json:"id"`
	DocumentID     string `json:"documentId"`
	DocumentTitle  string `json:"documentTitle"`
	CollectionSlug string `json:"collectionSlug"`
	VersionKey     string `json:"versionKey"`
	SlugPath       string `json:"slugPath"`
	Locale         string `json:"locale"`
	ParentID       string `json:"parentId,omitempty"`
	AuthorName     string `json:"authorName"`
	AvatarURL      string `json:"avatarUrl,omitempty"`
	UserSub        string `json:"userSub"`
	Content        string `json:"content"`
	Status         string `json:"status"`
	CreatedAt      string `json:"createdAt"`
}

type ListDocumentCommentsReq struct {
	g.Meta     `path:"/api/v1/docs/{documentId}/comments" method:"get" tags:"comments" summary:"List approved document comments"`
	DocumentID string `json:"documentId" in:"path" v:"required"`
	SortOrder  string `json:"sortOrder" in:"query" d:"asc" v:"in:asc,desc"`
	Page       int    `json:"page" in:"query" d:"1"`
	Size       int    `json:"size" in:"query" d:"20"`
}

type ListDocumentCommentsRes struct {
	Items []*CommentView `json:"items"`
	Total int            `json:"total"`
	Page  int            `json:"page"`
	Size  int            `json:"size"`
}

type CreateDocumentCommentReq struct {
	g.Meta     `path:"/api/v1/docs/{documentId}/comments" method:"post" tags:"comments" summary:"Create a document comment"`
	DocumentID string `json:"documentId" in:"path" v:"required"`
	Content    string `json:"content" v:"required|length:1,2000"`
	ParentID   string `json:"parentId"`
}

type CreateDocumentCommentRes struct {
	Comment *CommentView `json:"comment"`
}

type ManageCommentsReq struct {
	g.Meta    `path:"/api/v1/manage/comments" method:"get" tags:"comments" summary:"List comments for moderation"`
	Status    string `json:"status" in:"query"`
	Q         string `json:"q" in:"query"`
	SortBy    string `json:"sortBy" in:"query" d:"createdAt" v:"in:createdAt"`
	SortOrder string `json:"sortOrder" in:"query" d:"desc" v:"in:asc,desc"`
	Page      int    `json:"page" in:"query" d:"1"`
	Size      int    `json:"size" in:"query" d:"20"`
}

type ManageCommentsRes struct {
	Items []*CommentAdminView `json:"items"`
	Total int                 `json:"total"`
	Page  int                 `json:"page"`
	Size  int                 `json:"size"`
}

type ModerateCommentReq struct {
	g.Meta `path:"/api/v1/manage/comments/{id}" method:"patch" tags:"comments" summary:"Moderate a document comment"`
	ID     string `json:"id" in:"path" v:"required"`
	Status string `json:"status" v:"required|in:approved,pending,spam,trash"`
}

type ModerateCommentRes struct {
	Comment *CommentAdminView `json:"comment"`
}

type DeleteCommentReq struct {
	g.Meta `path:"/api/v1/manage/comments/{id}" method:"delete" tags:"comments" summary:"Delete a document comment"`
	ID     string `json:"id" in:"path" v:"required"`
}

type DeleteCommentRes struct {
	Deleted bool `json:"deleted"`
}
