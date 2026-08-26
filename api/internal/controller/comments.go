package controller

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/yueli-official/foundation/go/authorization"

	v1 "github.com/yueli-official/docs/api/api/v1"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/docscomments"
	"github.com/yueli-official/docs/api/internal/docserr"
	"github.com/yueli-official/docs/api/internal/identityclient"
	foundationauth "github.com/yueli-official/foundation/go/auth"
)

type PublicComments struct {
	comments *docscomments.Module
	profiles identityclient.Client
}

func NewPublicComments(comments *docscomments.Module, profiles identityclient.Client) *PublicComments {
	return &PublicComments{comments: comments, profiles: profiles}
}

func (controller *PublicComments) ListDocumentComments(
	ctx context.Context,
	req *v1.ListDocumentCommentsReq,
) (*v1.ListDocumentCommentsRes, error) {
	threads, total, page, size, err := controller.comments.List(
		ctx, req.DocumentID, req.Page, req.Size,
	)
	if err != nil {
		return nil, mapCommentError(err, req.DocumentID)
	}
	return &v1.ListDocumentCommentsRes{
		Items: commentThreadViews(threads, resolveCommentProfiles(ctx, controller.profiles, commentThreadUserKeys(threads))),
		Total: total, Page: page, Size: size,
	}, nil
}

type Comments struct {
	comments *docscomments.Module
	profiles identityclient.Client
}

func NewComments(comments *docscomments.Module, profiles identityclient.Client) *Comments {
	return &Comments{comments: comments, profiles: profiles}
}

func (controller *Comments) CreateDocumentComment(
	ctx context.Context,
	req *v1.CreateDocumentCommentReq,
) (*v1.CreateDocumentCommentRes, error) {
	userSub, err := commentUserKey(ctx)
	if err != nil {
		return nil, err
	}
	comment, err := controller.comments.Create(
		ctx, req.DocumentID, userSub, commentAuthorName(ctx, userSub),
		req.Content, req.ParentID,
	)
	if err != nil {
		return nil, mapCommentError(err, req.DocumentID)
	}
	profiles := resolveCommentProfiles(ctx, controller.profiles, []string{comment.UserSub})
	return &v1.CreateDocumentCommentRes{Comment: commentView(comment, profiles)}, nil
}

func (controller *Comments) ManageComments(
	ctx context.Context,
	req *v1.ManageCommentsReq,
) (*v1.ManageCommentsRes, error) {
	if err := requireCommentManagement(ctx); err != nil {
		return nil, err
	}
	result, err := controller.comments.Manage(ctx, docscomments.AdminQuery{
		Status: docscomments.Status(req.Status), Q: req.Q,
		SortBy: req.SortBy, SortOrder: req.SortOrder,
		Page: req.Page, Size: req.Size,
	})
	if err != nil {
		return nil, mapCommentError(err, "comments")
	}
	return &v1.ManageCommentsRes{
		Items: adminCommentViews(
			result.Items,
			resolveCommentProfiles(ctx, controller.profiles, adminCommentUserKeys(result.Items)),
		), Total: result.Total,
		Page: result.Page, Size: result.Size,
	}, nil
}

func (controller *Comments) ModerateComment(
	ctx context.Context,
	req *v1.ModerateCommentReq,
) (*v1.ModerateCommentRes, error) {
	if err := requireCommentManagement(ctx); err != nil {
		return nil, err
	}
	comment, err := controller.comments.Moderate(ctx, req.ID, docscomments.Status(req.Status))
	if err != nil {
		return nil, mapCommentError(err, req.ID)
	}
	return &v1.ModerateCommentRes{
		Comment: adminCommentView(
			&docscomments.AdminComment{Comment: comment},
			resolveCommentProfiles(ctx, controller.profiles, []string{comment.UserSub}),
		),
	}, nil
}

func (controller *Comments) DeleteComment(
	ctx context.Context,
	req *v1.DeleteCommentReq,
) (*v1.DeleteCommentRes, error) {
	if err := requireCommentManagement(ctx); err != nil {
		return nil, err
	}
	if err := controller.comments.Delete(ctx, req.ID); err != nil {
		return nil, mapCommentError(err, req.ID)
	}
	return &v1.DeleteCommentRes{Deleted: true}, nil
}

func requireCommentManagement(ctx context.Context) error {
	return requireCapability(
		ctx, authorization.CapabilityManage, docsauthz.RootScopeID,
		authorization.ResourceFacts{},
	)
}

func commentAuthorName(ctx context.Context, subject string) string {
	principal, _ := foundationauth.FromContext(ctx)
	for _, claim := range []string{"name", "preferred_username", "nickname"} {
		value, ok := principal.Claim(claim)
		name, valid := value.(string)
		if ok && valid && strings.TrimSpace(name) != "" {
			return truncateRunes(strings.TrimSpace(name), 80)
		}
	}
	if len(subject) > 8 {
		return subject[:8]
	}
	return subject
}

func commentUserKey(ctx context.Context) (string, error) {
	principal, _ := foundationauth.FromContext(ctx)
	if principal != nil {
		if value, ok := principal.Claim("user_key"); ok {
			if userKey, valid := value.(string); valid && strings.TrimSpace(userKey) != "" {
				return strings.TrimSpace(userKey), nil
			}
		}
	}
	return subject(ctx)
}

func truncateRunes(value string, maximum int) string {
	runes := []rune(value)
	if len(runes) <= maximum {
		return value
	}
	return string(runes[:maximum])
}

func commentView(comment *docscomments.Comment, profiles map[string]identityclient.PublicUser) *v1.CommentView {
	if comment == nil {
		return nil
	}
	name, avatarURL := commentPresentation(comment, profiles)
	return &v1.CommentView{
		ID: comment.ID, ParentID: comment.ParentID,
		AuthorName: name, AvatarURL: avatarURL,
		Content: comment.Content, CreatedAt: comment.CreatedAt.UTC().Format(time.RFC3339),
	}
}

func commentThreadViews(threads []*docscomments.Thread, profiles map[string]identityclient.PublicUser) []*v1.CommentView {
	items := make([]*v1.CommentView, 0, len(threads))
	for _, thread := range threads {
		item := commentView(thread.Comment, profiles)
		for _, reply := range thread.Replies {
			item.Replies = append(item.Replies, commentView(reply, profiles))
		}
		items = append(items, item)
	}
	return items
}

func adminCommentView(item *docscomments.AdminComment, profiles map[string]identityclient.PublicUser) *v1.CommentAdminView {
	if item == nil || item.Comment == nil {
		return nil
	}
	comment := item.Comment
	document := item.Document
	name, avatarURL := commentPresentation(comment, profiles)
	return &v1.CommentAdminView{
		ID: comment.ID, DocumentID: comment.DocumentID,
		DocumentTitle: document.Title, CollectionSlug: document.CollectionSlug,
		VersionKey: document.VersionKey, SlugPath: document.SlugPath, Locale: document.Locale,
		ParentID: comment.ParentID, AuthorName: name, AvatarURL: avatarURL, UserSub: comment.UserSub,
		Content: comment.Content, Status: string(comment.Status),
		CreatedAt: comment.CreatedAt.UTC().Format(time.RFC3339),
	}
}

func adminCommentViews(items []*docscomments.AdminComment, profiles map[string]identityclient.PublicUser) []*v1.CommentAdminView {
	views := make([]*v1.CommentAdminView, 0, len(items))
	for _, item := range items {
		views = append(views, adminCommentView(item, profiles))
	}
	return views
}

func commentPresentation(comment *docscomments.Comment, profiles map[string]identityclient.PublicUser) (string, string) {
	name := comment.AuthorName
	profile, ok := profiles[comment.UserSub]
	if !ok {
		return name, ""
	}
	if strings.TrimSpace(profile.DisplayName) != "" {
		name = strings.TrimSpace(profile.DisplayName)
	}
	return name, publicCommentMediaURL(profile.Avatar)
}

func publicCommentMediaURL(reference *identityclient.MediaRef) string {
	if reference == nil || reference.MediaKey == "" {
		return ""
	}
	return "/media/" + reference.MediaKey + "?format=webp&name=thumbnail"
}

func resolveCommentProfiles(ctx context.Context, client identityclient.Client, userKeys []string) map[string]identityclient.PublicUser {
	if client == nil {
		return map[string]identityclient.PublicUser{}
	}
	return client.GetMany(ctx, userKeys)
}

func commentThreadUserKeys(threads []*docscomments.Thread) []string {
	keys := make([]string, 0, len(threads))
	for _, thread := range threads {
		keys = append(keys, thread.Comment.UserSub)
		for _, reply := range thread.Replies {
			keys = append(keys, reply.UserSub)
		}
	}
	return keys
}

func adminCommentUserKeys(items []*docscomments.AdminComment) []string {
	keys := make([]string, 0, len(items))
	for _, item := range items {
		if item != nil && item.Comment != nil {
			keys = append(keys, item.Comment.UserSub)
		}
	}
	return keys
}

func mapCommentError(err error, id string) error {
	switch {
	case errors.Is(err, docscomments.ErrDocumentNotFound),
		errors.Is(err, docscomments.ErrCommentNotFound):
		return docserr.NotFound(id)
	case errors.Is(err, docscomments.ErrInvalidInput):
		return docserr.InvalidInput("comment input is invalid")
	default:
		return docserr.UpstreamFailed("comment storage unavailable")
	}
}
