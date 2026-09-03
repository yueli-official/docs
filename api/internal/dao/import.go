package dao

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/yueli-official/docs/api/internal/model"
)

type ImportDocMutation struct {
	Action                 string
	Doc                    *model.Doc
	ItemID                 string
	AfterDocJSON           string
	TransformedContentHash string
}

// ApplyImportDocs commits all document rows, import receipts, and the URL
// lifecycle delta in one product transaction. Asset uploads happen before this
// boundary and are referenced only after it commits.
func (p *PG) ApplyImportDocs(
	ctx context.Context,
	mutations []ImportDocMutation,
	hook TransactionHook,
) error {
	return p.db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		createByID := map[string]*model.Doc{}
		for _, mutation := range mutations {
			if mutation.Action == "create" && mutation.Doc != nil {
				createByID[mutation.Doc.ID] = mutation.Doc
			}
		}
		inserted := map[string]bool{}
		for len(inserted) < len(createByID) {
			batch := make([]g.Map, 0, len(createByID)-len(inserted))
			batchIDs := make([]string, 0, cap(batch))
			for id, doc := range createByID {
				if inserted[id] {
					continue
				}
				if _, parentIsNew := createByID[doc.ParentID]; parentIsNew && !inserted[doc.ParentID] {
					continue
				}
				batch = append(batch, importDocInsertData(doc))
				batchIDs = append(batchIDs, id)
			}
			if len(batch) == 0 {
				return fmt.Errorf("import document hierarchy contains a cycle")
			}
			if _, err := tx.Model(tDocs).Ctx(ctx).Data(batch).Insert(); err != nil {
				return err
			}
			for _, id := range batchIDs {
				inserted[id] = true
			}
		}
		for _, mutation := range mutations {
			if mutation.Doc == nil {
				continue
			}
			switch mutation.Action {
			case "create":
				// Inserted in one batch above; IDs and parent IDs are prepared before the transaction.
			case "update", "archive":
				if err := updateImportDoc(ctx, tx, mutation.Doc); err != nil {
					return err
				}
			case "delete":
				if _, err := tx.Ctx(ctx).Exec(`
WITH RECURSIVE subtree AS (
    SELECT id FROM docs WHERE id = ?::uuid AND deleted_at IS NULL
    UNION ALL
    SELECT child.id FROM docs child
    JOIN subtree parent ON child.parent_id = parent.id
    WHERE child.deleted_at IS NULL
)
UPDATE docs SET deleted_at = NOW(), updated_at = NOW()
WHERE id IN (SELECT id FROM subtree)`, mutation.Doc.ID); err != nil {
					return err
				}
			}
			if mutation.ItemID != "" {
				if _, err := tx.Model(tDocImportItems).Ctx(ctx).Where("id", mutation.ItemID).Data(g.Map{
					"target_doc_id":            nilIfEmpty(mutation.Doc.ID),
					"after_doc_json":           jsonOrEmptyObject(mutation.AfterDocJSON),
					"transformed_content_hash": mutation.TransformedContentHash,
				}).Update(); err != nil {
					return err
				}
			}
		}
		return runTransactionHook(ctx, tx, hook)
	})
}

func importDocInsertData(doc *model.Doc) g.Map {
	data := g.Map{
		"id": doc.ID, "collection_id": doc.CollectionID, "version_id": doc.VersionID,
		"slug": doc.Slug, "title": doc.Title, "content": doc.Content,
		"excerpt": doc.Excerpt, "seo_title": doc.SEOTitle,
		"seo_description": doc.SEODescription, "status": nz(doc.Status, "draft"),
		"locale": nz(doc.Locale, "en"), "translation_key": nz(doc.TranslationKey, doc.ID),
		"sort_order": doc.SortOrder, "author_sub": doc.AuthorSub,
	}
	if doc.ParentID != "" {
		data["parent_id"] = doc.ParentID
	}
	return data
}

func updateImportDoc(ctx context.Context, tx gdb.TX, doc *model.Doc) error {
	if _, err := tx.Model(tDocs).Ctx(ctx).Where("id", doc.ID).Data(g.Map{
		"title": doc.Title, "slug": doc.Slug, "content": doc.Content,
		"excerpt": doc.Excerpt, "seo_title": doc.SEOTitle,
		"seo_description": doc.SEODescription, "version_id": doc.VersionID,
		"status": doc.Status, "locale": doc.Locale,
		"translation_key": doc.TranslationKey, "sort_order": doc.SortOrder,
		"parent_id": nilIfEmpty(doc.ParentID), "deleted_at": nil,
		"updated_at": gtime.Now(),
	}).Update(); err != nil {
		return err
	}
	_, err := tx.Ctx(ctx).Exec(`
WITH RECURSIVE descendants AS (
    SELECT id FROM docs WHERE parent_id = ?::uuid AND deleted_at IS NULL
    UNION ALL
    SELECT child.id FROM docs child
    JOIN descendants parent ON child.parent_id = parent.id
    WHERE child.deleted_at IS NULL
)
UPDATE docs SET version_id = ?::uuid, locale = ?, updated_at = NOW()
WHERE id IN (SELECT id FROM descendants)`, doc.ID, doc.VersionID, doc.Locale)
	return err
}

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

func (p *PG) ListImportBatches(ctx context.Context, collectionID string, page, size int) ([]*model.ImportBatch, int, error) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 || size > 50 {
		size = 8
	}
	var out []*model.ImportBatch
	q := p.db.Model(tDocImportBatches).Ctx(ctx)
	if collectionID != "" {
		q = q.Where("collection_id", collectionID)
	}
	total, err := q.Clone().Count()
	if err != nil {
		return nil, 0, err
	}
	err = q.OrderDesc("created_at").Page(page, size).Scan(&out)
	if out == nil {
		out = []*model.ImportBatch{}
	}
	return out, total, err
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
