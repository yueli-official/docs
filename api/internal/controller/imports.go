package controller

import (
	"context"
	"encoding/json"
	"io"

	v1 "platform/products/docs/api/api/v1"
	"platform/products/docs/api/internal/catalog"
	"platform/products/docs/api/internal/docserr"
	"platform/products/docs/api/internal/importkit"
	"platform/products/docs/api/internal/model"
)

type Imports struct{ svc *catalog.Service }

func NewImports(svc *catalog.Service) *Imports { return &Imports{svc: svc} }

func (c *Imports) UploadDocsImport(ctx context.Context, req *v1.UploadDocsImportReq) (*v1.UploadDocsImportRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
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
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
	}
	batch, items, err := c.svc.GetImport(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &v1.GetDocsImportRes{Batch: importBatchView(batch), Items: importItemViews(items)}, nil
}

func (c *Imports) ConfirmDocsImport(ctx context.Context, req *v1.ConfirmDocsImportReq) (*v1.ConfirmDocsImportRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
	}
	author, err := subject(ctx)
	if err != nil {
		return nil, err
	}
	batch, summary, err := c.svc.ConfirmImport(ctx, req.ID, bearerOf(ctx), author)
	if err != nil {
		return nil, err
	}
	return &v1.ConfirmDocsImportRes{Batch: importBatchView(batch), Summary: importSummaryView(summary)}, nil
}

func (c *Imports) RollbackDocsImport(ctx context.Context, req *v1.RollbackDocsImportReq) (*v1.RollbackDocsImportRes, error) {
	if !isAdmin(ctx) {
		return nil, docserr.Forbidden()
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
	}
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
