package controller

import (
	"context"
	"encoding/json"
	"io"

	"github.com/yueli-official/foundation/go/authorization"

	v1 "github.com/yueli-official/docs/api/api/v1"
	"github.com/yueli-official/docs/api/internal/catalog"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/docserr"
	"github.com/yueli-official/docs/api/internal/importkit"
	"github.com/yueli-official/docs/api/internal/model"
)

type Imports struct{ svc *catalog.Service }

func NewImports(svc *catalog.Service) *Imports { return &Imports{svc: svc} }

func (c *Imports) ListDocsImports(ctx context.Context, req *v1.ListDocsImportsReq) (*v1.ListDocsImportsRes, error) {
	if err := requireCapability(ctx, docsauthz.CapabilityImportManage, docsauthz.RootScopeID, authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	items, err := c.svc.ListImports(ctx, req.CollectionID, req.Limit)
	if err != nil {
		return nil, err
	}
	return &v1.ListDocsImportsRes{Items: importBatchViews(items)}, nil
}

func (c *Imports) UploadDocsImport(ctx context.Context, req *v1.UploadDocsImportReq) (*v1.UploadDocsImportRes, error) {
	if err := requireCapability(ctx, docsauthz.CapabilityImportManage, docsauthz.RootScopeID, authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	author, err := subject(ctx)
	if err != nil {
		return nil, err
	}
	if req.File == nil {
		return nil, docserr.InvalidInput("file required")
	}
	file, err := req.File.Open()
	if err != nil {
		return nil, docserr.InvalidInput("file open failed")
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, docserr.InvalidInput("file read failed")
	}
	batch, summary, err := c.svc.PreflightImport(ctx, catalog.ImportUploadInput{
		Filename: req.File.Filename,
		Data:     data,
		Bearer:   bearerOf(ctx),
		Author:   author,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UploadDocsImportRes{
		Batch:   importBatchView(batch),
		Summary: importSummaryView(summary),
	}, nil
}

func (c *Imports) GetDocsImport(ctx context.Context, req *v1.GetDocsImportReq) (*v1.GetDocsImportRes, error) {
	if err := requireCapability(ctx, docsauthz.CapabilityImportManage, docsauthz.RootScopeID, authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	batch, items, err := c.svc.GetImport(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &v1.GetDocsImportRes{Batch: importBatchView(batch), Items: importItemViews(items)}, nil
}

func (c *Imports) ConfirmDocsImport(ctx context.Context, req *v1.ConfirmDocsImportReq) (*v1.ConfirmDocsImportRes, error) {
	if err := requireCapability(ctx, docsauthz.CapabilityImportManage, docsauthz.RootScopeID, authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	author, err := subject(ctx)
	if err != nil {
		return nil, err
	}
	batch, summary, err := c.svc.ConfirmImport(ctx, req.ID, bearerOf(ctx), author)
	if err != nil {
		return nil, err
	}
	if err := authorizationService(ctx).SyncCatalogScopes(ctx); err != nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	return &v1.ConfirmDocsImportRes{Batch: importBatchView(batch), Summary: importSummaryView(summary)}, nil
}

func (c *Imports) RollbackDocsImport(ctx context.Context, req *v1.RollbackDocsImportReq) (*v1.RollbackDocsImportRes, error) {
	if err := requireCapability(ctx, docsauthz.CapabilityImportManage, docsauthz.RootScopeID, authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	author, err := subject(ctx)
	if err != nil {
		return nil, err
	}
	batch, err := c.svc.RollbackImport(ctx, req.ID, author)
	if err != nil {
		return nil, err
	}
	return &v1.RollbackDocsImportRes{Batch: importBatchView(batch)}, nil
}

func importBatchView(m *model.ImportBatch) *v1.ImportBatchView {
	if m == nil {
		return nil
	}
	return &v1.ImportBatchView{
		ID:            m.ID,
		CollectionID:  m.CollectionID,
		VersionID:     m.VersionID,
		DefaultLocale: m.DefaultLocale,
		Mode:          m.Mode,
		Status:        m.Status,
		ErrorMessage:  m.ErrorMessage,
		Summary:       importSummaryFromJSON(m.SummaryJSON),
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
		CompletedAt:   m.CompletedAt,
	}
}

func importBatchViews(items []*model.ImportBatch) []*v1.ImportBatchView {
	out := make([]*v1.ImportBatchView, len(items))
	for i, item := range items {
		out[i] = importBatchView(item)
	}
	return out
}

func importItemView(m *model.ImportItem) *v1.ImportItemView {
	if m == nil {
		return nil
	}
	return &v1.ImportItemView{
		ID:                 m.ID,
		Locale:             m.Locale,
		VersionKey:         m.VersionKey,
		Path:               m.Path,
		SourceMarkdownPath: m.SourceMarkdownPath,
		Title:              m.Title,
		Slug:               m.Slug,
		TranslationKey:     m.TranslationKey,
		Action:             m.Action,
		TargetDocID:        m.TargetDocID,
		Issues:             importIssuesFromJSON(m.IssuesJSON),
	}
}

func importItemViews(items []*model.ImportItem) []*v1.ImportItemView {
	out := make([]*v1.ImportItemView, len(items))
	for i, item := range items {
		out[i] = importItemView(item)
	}
	return out
}

func importSummaryView(s catalog.ImportSummary) v1.ImportSummaryView {
	return v1.ImportSummaryView{
		Creates:   s.Creates,
		Updates:   s.Updates,
		Archives:  s.Archives,
		Skips:     s.Skips,
		Conflicts: s.Conflicts,
		Errors:    s.Errors,
		Images:    s.Images,
		Warnings:  s.Warnings,
		Blocking:  s.Blocking,
		Issues:    importIssueViews(s.Issues),
	}
}

func importSummaryFromJSON(raw string) v1.ImportSummaryView {
	var summary catalog.ImportSummary
	_ = json.Unmarshal([]byte(raw), &summary)
	return importSummaryView(summary)
}

func importIssuesFromJSON(raw string) []v1.ImportIssueView {
	var issues []importkit.Issue
	_ = json.Unmarshal([]byte(raw), &issues)
	return importIssueViews(issues)
}

func importIssueViews(issues []importkit.Issue) []v1.ImportIssueView {
	out := make([]v1.ImportIssueView, len(issues))
	for i, issue := range issues {
		out[i] = v1.ImportIssueView{
			Severity: issue.Severity,
			Code:     issue.Code,
			Message:  issue.Message,
			Path:     issue.Path,
		}
	}
	return out
}
