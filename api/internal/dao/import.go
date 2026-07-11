package dao

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"platform/products/docs/api/internal/model"
)

const (
	tDocImportBatches   = "doc_import_batches"
	tDocImportItems     = "doc_import_items"
	tDocImportAssets    = "doc_import_assets"
	tDocImportAssetRefs = "doc_import_asset_refs"
)

func (p *PG) InsertImportBatch(ctx context.Context, m *model.ImportBatch) error {
	_, err := p.db.Model(tDocImportBatches).Ctx(ctx).Data(g.Map{
		"id":             m.ID,
		"collection_id":  m.CollectionID,
		"version_id":     m.VersionID,
		"default_locale": m.DefaultLocale,
		"mode":           m.Mode,
		"status":         m.Status,
		"summary_json":   jsonOrEmptyObject(m.SummaryJSON),
		"error_message":  m.ErrorMessage,
		"created_by":     m.CreatedBy,
	}).Insert()
	return err
}

func (p *PG) UpdateImportBatchStatus(ctx context.Context, id, status, summaryJSON, errorMessage string) error {
	data := g.Map{
		"status":        status,
		"summary_json":  jsonOrEmptyObject(summaryJSON),
		"error_message": errorMessage,
		"updated_at":    gtime.Now(),
	}
	if status == "completed" || status == "failed" || status == "cancelled" || status == "rolled_back" {
		data["completed_at"] = gtime.Now()
	}
	_, err := p.db.Model(tDocImportBatches).Ctx(ctx).Where("id", id).Data(data).Update()
	return err
}

func (p *PG) GetImportBatch(ctx context.Context, id string) (*model.ImportBatch, error) {
	var out *model.ImportBatch
	err := p.db.Model(tDocImportBatches).Ctx(ctx).Where("id", id).Limit(1).Scan(&out)
	return out, err
}

func (p *PG) ListImportBatches(ctx context.Context, collectionID string, limit int) ([]*model.ImportBatch, error) {
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	var out []*model.ImportBatch
	q := p.db.Model(tDocImportBatches).Ctx(ctx)
	if collectionID != "" {
		q = q.Where("collection_id", collectionID)
	}
	err := q.OrderDesc("created_at").Limit(limit).Scan(&out)
	if out == nil {
		out = []*model.ImportBatch{}
	}
	return out, err
}

func (p *PG) InsertImportItems(ctx context.Context, items []*model.ImportItem) error {
	if len(items) == 0 {
		return nil
	}
	rows := make([]g.Map, 0, len(items))
	for _, m := range items {
		rows = append(rows, g.Map{
			"id":                       m.ID,
			"batch_id":                 m.BatchID,
			"locale":                   m.Locale,
			"version_key":              m.VersionKey,
			"path":                     m.Path,
			"source_markdown_path":     m.SourceMarkdownPath,
			"title":                    m.Title,
			"slug":                     m.Slug,
			"translation_key":          m.TranslationKey,
			"action":                   m.Action,
			"target_doc_id":            nilIfEmpty(m.TargetDocID),
			"before_doc_json":          jsonOrEmptyObject(m.BeforeDocJSON),
			"after_doc_json":           jsonOrEmptyObject(m.AfterDocJSON),
			"issues_json":              jsonOrEmptyArray(m.IssuesJSON),
			"source_content_hash":      m.SourceContentHash,
			"transformed_content_hash": m.TransformedContentHash,
		})
	}
	_, err := p.db.Model(tDocImportItems).Ctx(ctx).Data(rows).Insert()
	return err
}

func (p *PG) ListImportItems(ctx context.Context, batchID string) ([]*model.ImportItem, error) {
	var out []*model.ImportItem
	err := p.db.Model(tDocImportItems).Ctx(ctx).
		Where("batch_id", batchID).
		OrderAsc("locale").
		OrderAsc("path").
		Scan(&out)
	if out == nil {
		out = []*model.ImportItem{}
	}
	return out, err
}

func (p *PG) InsertImportAssets(ctx context.Context, assets []*model.ImportAsset) error {
	if len(assets) == 0 {
		return nil
	}
	rows := make([]g.Map, 0, len(assets))
	for _, m := range assets {
		rows = append(rows, g.Map{
			"id":            m.ID,
			"batch_id":      m.BatchID,
			"source_path":   m.SourcePath,
			"data":          m.Data,
			"asset_url":     m.AssetURL,
			"content_hash":  m.ContentHash,
			"status":        nz(m.Status, "pending"),
			"error_message": m.ErrorMessage,
		})
	}
	_, err := p.db.Model(tDocImportAssets).Ctx(ctx).Data(rows).Insert()
	return err
}

func (p *PG) ListImportAssets(ctx context.Context, batchID string) ([]*model.ImportAsset, error) {
	var out []*model.ImportAsset
	err := p.db.Model(tDocImportAssets).Ctx(ctx).
		Where("batch_id", batchID).
		OrderAsc("source_path").
		Scan(&out)
	if out == nil {
		out = []*model.ImportAsset{}
	}
	return out, err
}

func (p *PG) UpdateImportAssetUploaded(ctx context.Context, id, assetURL string) error {
	_, err := p.db.Model(tDocImportAssets).Ctx(ctx).Where("id", id).Data(g.Map{
		"asset_url": assetURL,
		"status":    "uploaded",
	}).Update()
	return err
}

func (p *PG) UpdateImportItemResult(ctx context.Context, id, targetDocID, afterDocJSON, transformedContentHash string) error {
	_, err := p.db.Model(tDocImportItems).Ctx(ctx).Where("id", id).Data(g.Map{
		"target_doc_id":            nilIfEmpty(targetDocID),
		"after_doc_json":           jsonOrEmptyObject(afterDocJSON),
		"transformed_content_hash": transformedContentHash,
	}).Update()
	return err
}

func (p *PG) InsertImportAssetRefs(ctx context.Context, refs []*model.ImportAssetRef) error {
	if len(refs) == 0 {
		return nil
	}
	rows := make([]g.Map, 0, len(refs))
	for _, m := range refs {
		rows = append(rows, g.Map{
			"id":                 m.ID,
			"batch_id":           m.BatchID,
			"asset_id":           nilIfEmpty(m.AssetID),
			"item_id":            m.ItemID,
			"markdown_file_path": m.MarkdownFilePath,
			"original_ref":       m.OriginalRef,
			"rewritten_ref":      m.RewrittenRef,
		})
	}
	_, err := p.db.Model(tDocImportAssetRefs).Ctx(ctx).Data(rows).Insert()
	return err
}

func (p *PG) ListImportAssetRefs(ctx context.Context, batchID string) ([]*model.ImportAssetRef, error) {
	var out []*model.ImportAssetRef
	err := p.db.Model(tDocImportAssetRefs).Ctx(ctx).
		Where("batch_id", batchID).
		OrderAsc("markdown_file_path").
		OrderAsc("original_ref").
		Scan(&out)
	if out == nil {
		out = []*model.ImportAssetRef{}
	}
	return out, err
}

func jsonOrEmptyObject(s string) string {
	if s == "" {
		return "{}"
	}
	return s
}

func jsonOrEmptyArray(s string) string {
	if s == "" {
		return "[]"
	}
	return s
}
