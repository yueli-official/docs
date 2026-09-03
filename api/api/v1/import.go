package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type ImportIssueView struct {
	Severity string `json:"severity"`
	Code     string `json:"code"`
	Message  string `json:"message"`
	Path     string `json:"path,omitempty"`
}

type ImportSummaryView struct {
	Creates   int               `json:"creates"`
	Updates   int               `json:"updates"`
	Archives  int               `json:"archives"`
	Skips     int               `json:"skips"`
	Conflicts int               `json:"conflicts"`
	Errors    int               `json:"errors"`
	Images    int               `json:"images"`
	Warnings  int               `json:"warnings"`
	Blocking  bool              `json:"blocking"`
	Issues    []ImportIssueView `json:"issues"`
}

type ImportBatchView struct {
	ID            string            `json:"id"`
	CollectionID  string            `json:"collectionId"`
	VersionID     string            `json:"versionId"`
	DefaultLocale string            `json:"defaultLocale"`
	Mode          string            `json:"mode"`
	Status        string            `json:"status"`
	ErrorMessage  string            `json:"errorMessage"`
	Summary       ImportSummaryView `json:"summary"`
	CreatedAt     time.Time         `json:"createdAt"`
	UpdatedAt     time.Time         `json:"updatedAt"`
	CompletedAt   *time.Time        `json:"completedAt,omitempty"`
}

type ImportItemView struct {
	ID                 string            `json:"id"`
	Locale             string            `json:"locale"`
	VersionKey         string            `json:"versionKey"`
	Path               string            `json:"path"`
	SourceMarkdownPath string            `json:"sourceMarkdownPath"`
	Title              string            `json:"title"`
	Slug               string            `json:"slug"`
	TranslationKey     string            `json:"translationKey"`
	Action             string            `json:"action"`
	TargetDocID        string            `json:"targetDocId"`
	Issues             []ImportIssueView `json:"issues"`
}

type UploadDocsImportReq struct {
	g.Meta        `path:"/api/v1/imports/docs" method:"post" mime:"multipart/form-data" tags:"docs" summary:"Upload and preflight a docs import ZIP"`
	File          *ghttp.UploadFile `json:"file" type:"file" v:"required"`
	Collection    string            `json:"collection"`
	DefaultLocale string            `json:"defaultLocale"`
	Mode          string            `json:"mode"`
}
type UploadDocsImportRes struct {
	Batch   *ImportBatchView  `json:"batch"`
	Summary ImportSummaryView `json:"summary"`
}

type ListDocsImportsReq struct {
	g.Meta       `path:"/api/v1/imports/docs" method:"get" tags:"docs" summary:"List recent docs import batches"`
	CollectionID string `json:"collectionId" in:"query"`
	Limit        int    `json:"limit" in:"query" d:"20" v:"min:1|max:100"`
}
type ListDocsImportsRes struct {
	Items []*ImportBatchView `json:"items"`
}

type GetDocsImportReq struct {
	g.Meta `path:"/api/v1/imports/docs/{id}" method:"get" tags:"docs" summary:"Get docs import batch"`
	ID     string `json:"id" in:"path" v:"required"`
}
type GetDocsImportRes struct {
	Batch *ImportBatchView  `json:"batch"`
	Items []*ImportItemView `json:"items"`
}

type ConfirmDocsImportReq struct {
	g.Meta `path:"/api/v1/imports/docs/{id}/confirm" method:"post" tags:"docs" summary:"Confirm docs import"`
	ID     string `json:"id" in:"path" v:"required"`
}
type ConfirmDocsImportRes struct {
	Batch   *ImportBatchView  `json:"batch"`
	Summary ImportSummaryView `json:"summary"`
}

type RollbackDocsImportReq struct {
	g.Meta `path:"/api/v1/imports/docs/{id}/rollback" method:"post" tags:"docs" summary:"Rollback docs import"`
	ID     string `json:"id" in:"path" v:"required"`
}
type RollbackDocsImportRes struct {
	Batch *ImportBatchView `json:"batch"`
}
