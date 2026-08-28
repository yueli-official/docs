package catalog

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"path"
	"sort"
	"strings"

	"github.com/yueli-official/docs/api/internal/assetclient"
	"github.com/yueli-official/docs/api/internal/dao"
	"github.com/yueli-official/docs/api/internal/docsaudit"
	"github.com/yueli-official/docs/api/internal/docserr"
	"github.com/yueli-official/docs/api/internal/importkit"
	"github.com/yueli-official/docs/api/internal/model"
	"github.com/yueli-official/foundation/go/identifier"
)

type ImportUploadInput struct {
	Filename string
	Data     []byte
	Bearer   string
	Author   string
}

type ImportSummary struct {
	Creates   int               `json:"creates"`
	Updates   int               `json:"updates"`
	Archives  int               `json:"archives"`
	Skips     int               `json:"skips"`
	Conflicts int               `json:"conflicts"`
	Errors    int               `json:"errors"`
	Images    int               `json:"images"`
	Warnings  int               `json:"warnings"`
	Blocking  bool              `json:"blocking"`
	Issues    []importkit.Issue `json:"issues"`
}

type plannedImportDoc struct {
	CollectionID   string `json:"collectionId"`
	VersionID      string `json:"versionId"`
	ParentPath     string `json:"parentPath"`
	ParentID       string `json:"parentId"`
	Slug           string `json:"slug"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	Excerpt        string `json:"excerpt"`
	Status         string `json:"status"`
	Locale         string `json:"locale"`
	TranslationKey string `json:"translationKey"`
	SortOrder      int    `json:"sortOrder"`
	SourcePath     string `json:"sourcePath"`
}

func (s *Service) PreflightImport(ctx context.Context, in ImportUploadInput) (*model.ImportBatch, ImportSummary, error) {
	pkg, err := importkit.ParseZip(in.Data, importkit.Options{MaxImageBytes: 10 << 20})
	if err != nil {
		return nil, ImportSummary{}, docserr.InvalidInput(err.Error())
	}
	col, err := s.dao.GetCollectionBySlug(ctx, pkg.Manifest.Collection)
	if err != nil {
		return nil, ImportSummary{}, err
	}
	if col == nil {
		return nil, ImportSummary{}, docserr.NotFound(pkg.Manifest.Collection)
	}
	version, err := s.ResolveVersion(ctx, col.ID, pkg.Manifest.Version, false)
	if err != nil {
		return nil, ImportSummary{}, err
	}

	existing, err := s.existingImportDocs(ctx, col.ID, version.ID, pkg.Manifest.Locales)
	if err != nil {
		return nil, ImportSummary{}, err
	}
	batchID := identifier.MustNew().String()
	summary := summarizePackage(pkg)
	items := make([]*model.ImportItem, 0, len(pkg.Docs))
	itemIDBySource := map[string]string{}
	incoming := map[string]bool{}

	for _, doc := range pkg.Docs {
		key := importDocKey(doc.Locale, doc.Path)
		incoming[key] = true
		current := existing[key]
		action := "create"
		if current != nil {
			if pkg.Manifest.Mode == "create-only" {
				action = "conflict"
				summary.Conflicts++
				summary.Blocking = true
			} else {
				action = "update"
				summary.Updates++
			}
		} else {
			summary.Creates++
		}
		itemID := identifier.MustNew().String()
		itemIDBySource[doc.SourcePath] = itemID
		planned := plannedImportDoc{
			CollectionID:   col.ID,
			VersionID:      version.ID,
			ParentPath:     parentPath(doc.Path),
			Slug:           doc.Slug,
			Title:          doc.Title,
			Content:        doc.Content,
			Excerpt:        doc.Excerpt,
			Status:         plannedStatus(current),
			Locale:         doc.Locale,
			TranslationKey: doc.TranslationKey,
			SortOrder:      doc.Order,
			SourcePath:     doc.SourcePath,
		}
		items = append(items, &model.ImportItem{
			ID:                 itemID,
			BatchID:            batchID,
			Locale:             doc.Locale,
			VersionKey:         doc.VersionKey,
			Path:               doc.Path,
			SourceMarkdownPath: doc.SourcePath,
			Title:              doc.Title,
			Slug:               doc.Slug,
			TranslationKey:     doc.TranslationKey,
			Action:             action,
			TargetDocID:        docID(current),
			BeforeDocJSON:      mustJSON(current),
			AfterDocJSON:       mustJSON(planned),
			IssuesJSON:         mustJSON(issuesForPath(pkg.Issues, doc.SourcePath)),
			SourceContentHash:  sha256Hex([]byte(doc.RawContent)),
		})
	}

	if pkg.Manifest.Mode == "replace-version" {
		for key, current := range existing {
			if incoming[key] || current == nil {
				continue
			}
			summary.Archives++
			items = append(items, &model.ImportItem{
				ID:             identifier.MustNew().String(),
				BatchID:        batchID,
				Locale:         current.Locale,
				VersionKey:     pkg.Manifest.Version,
				Path:           keyPath(key),
				Title:          current.Title,
				Slug:           current.Slug,
				TranslationKey: current.TranslationKey,
				Action:         "archive",
				TargetDocID:    current.ID,
				BeforeDocJSON:  mustJSON(current),
				AfterDocJSON:   mustJSON(plannedImportDoc{Status: "archived"}),
				IssuesJSON:     "[]",
			})
		}
	}

	assets, refs := importAssetsAndRefs(batchID, pkg, itemIDBySource)
	summary.Images = len(assets)
	summaryJSON := mustJSON(summary)
	batch := &model.ImportBatch{
		ID:            batchID,
		CollectionID:  col.ID,
		VersionID:     version.ID,
		DefaultLocale: pkg.Manifest.DefaultLocale,
		Mode:          pkg.Manifest.Mode,
		Status:        "checked",
		SummaryJSON:   summaryJSON,
		CreatedBy:     in.Author,
	}
	if err := s.dao.InsertImportBatch(ctx, batch); err != nil {
		return nil, ImportSummary{}, err
	}
	if err := s.dao.InsertImportItems(ctx, items); err != nil {
		return nil, ImportSummary{}, err
	}
	if err := s.dao.InsertImportAssets(ctx, assets); err != nil {
		return nil, ImportSummary{}, err
	}
	if err := s.dao.InsertImportAssetRefs(ctx, refs); err != nil {
		return nil, ImportSummary{}, err
	}
	created, err := s.dao.GetImportBatch(ctx, batchID)
	return created, summary, err
}

func (s *Service) ConfirmImport(ctx context.Context, batchID, bearer, author string) (*model.ImportBatch, ImportSummary, error) {
	batch, err := s.dao.GetImportBatch(ctx, batchID)
	if err != nil {
		return nil, ImportSummary{}, err
	}
	if batch == nil {
		return nil, ImportSummary{}, docserr.NotFound(batchID)
	}
	summary := parseImportSummary(batch.SummaryJSON)
	if batch.Status != "checked" {
		return nil, summary, docserr.InvalidInput("import batch is not ready")
	}
	if summary.Blocking {
		return nil, summary, docserr.ImportBlocked("preflight has blocking issues")
	}
	if err := s.dao.UpdateImportBatchStatus(ctx, batchID, "running", batch.SummaryJSON, ""); err != nil {
		return nil, summary, err
	}
	if err := s.executeImport(ctx, batch, bearer, author); err != nil {
		_ = s.dao.UpdateImportBatchStatus(ctx, batchID, "failed", batch.SummaryJSON, err.Error())
		return nil, summary, err
	}
	if err := s.dao.UpdateImportBatchStatus(ctx, batchID, "completed", batch.SummaryJSON, ""); err != nil {
		return nil, summary, err
	}
	done, err := s.dao.GetImportBatch(ctx, batchID)
	return done, summary, err
}

func (s *Service) RollbackImport(ctx context.Context, batchID, author string) (*model.ImportBatch, error) {
	batch, err := s.dao.GetImportBatch(ctx, batchID)
	if err != nil {
		return nil, err
	}
	if batch == nil {
		return nil, docserr.NotFound(batchID)
	}
	if batch.Status != "completed" {
		return nil, docserr.InvalidInput("import batch is not completed")
	}
	items, err := s.dao.ListImportItems(ctx, batchID)
	if err != nil {
		return nil, err
	}
	sort.SliceStable(items, func(i, j int) bool {
		return pathDepth(items[i].Path) > pathDepth(items[j].Path)
	})
	mutations := make([]dao.ImportDocMutation, 0, len(items))
	for _, item := range items {
		switch item.Action {
		case "create":
			if item.TargetDocID != "" {
				mutations = append(mutations, dao.ImportDocMutation{
					Action: "delete", Doc: &model.Doc{ID: item.TargetDocID},
					ItemID: item.ID, AfterDocJSON: "{}",
				})
			}
		case "update", "archive":
			var before model.Doc
			if err := json.Unmarshal([]byte(item.BeforeDocJSON), &before); err != nil || before.ID == "" {
				continue
			}
			mutations = append(mutations, dao.ImportDocMutation{
				Action: "update", Doc: &before, ItemID: item.ID,
				AfterDocJSON:           item.BeforeDocJSON,
				TransformedContentHash: sha256Hex([]byte(before.Content)),
			})
		}
	}
	if err := s.dao.ApplyImportDocs(
		ctx, mutations, s.importMutationHook(
			ctx, docsaudit.ActionImportRolledBack, batch, len(mutations), author,
			"docs import rolled back",
		),
	); err != nil {
		return nil, err
	}
	if err := s.dao.UpdateImportBatchStatus(ctx, batchID, "rolled_back", batch.SummaryJSON, ""); err != nil {
		return nil, err
	}
	return s.dao.GetImportBatch(ctx, batchID)
}

func (s *Service) GetImport(ctx context.Context, batchID string) (*model.ImportBatch, []*model.ImportItem, error) {
	batch, err := s.dao.GetImportBatch(ctx, batchID)
	if err != nil {
		return nil, nil, err
	}
	if batch == nil {
		return nil, nil, docserr.NotFound(batchID)
	}
	items, err := s.dao.ListImportItems(ctx, batchID)
	if err != nil {
		return nil, nil, err
	}
	return batch, items, nil
}

func (s *Service) ListImports(ctx context.Context, collectionID string, limit int) ([]*model.ImportBatch, error) {
	return s.dao.ListImportBatches(ctx, strings.TrimSpace(collectionID), limit)
}

func (s *Service) executeImport(ctx context.Context, batch *model.ImportBatch, bearer, author string) error {
	items, err := s.dao.ListImportItems(ctx, batch.ID)
	if err != nil {
		return err
	}
	localeSet := map[string]struct{}{}
	for _, item := range items {
		localeSet[item.Locale] = struct{}{}
	}
	for locale := range localeSet {
		existing, err := s.dao.GetCollectionLocale(ctx, batch.CollectionID, locale)
		if err != nil {
			return err
		}
		if existing != nil {
			continue
		}
		if _, err := s.UpsertLocale(ctx, UpsertLocaleInput{
			CollectionID: batch.CollectionID,
			Locale:       locale,
			Label:        locale,
			HTMLLang:     locale,
			Direction:    "ltr",
			Enabled:      true,
			SortOrder:    100,
		}); err != nil {
			return err
		}
	}
	assets, err := s.dao.ListImportAssets(ctx, batch.ID)
	if err != nil {
		return err
	}
	refs, err := s.dao.ListImportAssetRefs(ctx, batch.ID)
	if err != nil {
		return err
	}
	assetURLs, err := s.uploadImportAssets(ctx, bearer, assets)
	if err != nil {
		return err
	}
	refsByItem := map[string][]*model.ImportAssetRef{}
	for _, ref := range refs {
		refsByItem[ref.ItemID] = append(refsByItem[ref.ItemID], ref)
	}
	parentIDs, err := s.existingParentIDs(ctx, batch.CollectionID, batch.VersionID, items)
	if err != nil {
		return err
	}
	sort.SliceStable(items, func(i, j int) bool {
		return pathDepth(items[i].Path) < pathDepth(items[j].Path)
	})
	mutations := make([]dao.ImportDocMutation, 0, len(items))
	for _, item := range items {
		switch item.Action {
		case "create", "update":
			mutation, err := s.prepareImportMutation(item, refsByItem[item.ID], assetURLs, parentIDs)
			if err != nil {
				return err
			}
			mutations = append(mutations, mutation)
		case "archive":
			doc, err := s.dao.GetDocByID(ctx, item.TargetDocID)
			if err != nil {
				return err
			}
			if doc == nil {
				return docserr.NotFound(item.TargetDocID)
			}
			doc.Status = "archived"
			mutations = append(mutations, dao.ImportDocMutation{
				Action: "archive", Doc: doc, ItemID: item.ID,
				AfterDocJSON:           mustJSON(doc),
				TransformedContentHash: sha256Hex([]byte(doc.Content)),
			})
		}
	}
	return s.dao.ApplyImportDocs(
		ctx, mutations, s.importMutationHook(
			ctx, docsaudit.ActionImportConfirmed, batch, len(mutations), author,
			"docs import applied",
		),
	)
}

func summarizePackage(pkg *importkit.Package) ImportSummary {
	summary := ImportSummary{Issues: pkg.Issues}
	for _, issue := range pkg.Issues {
		if issue.Severity == "warning" {
			summary.Warnings++
			continue
		}
		summary.Errors++
		summary.Blocking = true
	}
	return summary
}

func (s *Service) existingImportDocs(ctx context.Context, collectionID, versionID string, locales []string) (map[string]*model.Doc, error) {
	out := map[string]*model.Doc{}
	for _, locale := range locales {
		docs, err := s.dao.ListDocsByCollection(ctx, collectionID, versionID, locale)
		if err != nil {
			return nil, err
		}
		for _, item := range flattenDocPaths(BuildTree(docs), "") {
			out[importDocKey(item.doc.Locale, item.path)] = item.doc
		}
	}
	return out, nil
}

type importDocPath struct {
	path string
	doc  *model.Doc
}

func flattenDocPaths(nodes []*model.Doc, base string) []importDocPath {
	var out []importDocPath
	for _, node := range nodes {
		p := node.Slug
		if base != "" {
			p = base + "/" + node.Slug
		}
		out = append(out, importDocPath{path: p, doc: node})
		out = append(out, flattenDocPaths(node.Children, p)...)
	}
	return out
}

func importDocKey(locale, p string) string {
	return locale + "\x00" + path.Clean(strings.Trim(p, "/"))
}

func keyPath(key string) string {
	_, p, _ := strings.Cut(key, "\x00")
	return p
}

func docID(doc *model.Doc) string {
	if doc == nil {
		return ""
	}
	return doc.ID
}

func issuesForPath(issues []importkit.Issue, p string) []importkit.Issue {
	var out []importkit.Issue
	for _, issue := range issues {
		if issue.Path == p {
			out = append(out, issue)
		}
	}
	return out
}

func importAssetsAndRefs(batchID string, pkg *importkit.Package, itemIDBySource map[string]string) ([]*model.ImportAsset, []*model.ImportAssetRef) {
	assetIDBySource := map[string]string{}
	var assets []*model.ImportAsset
	var refs []*model.ImportAssetRef
	for _, doc := range pkg.Docs {
		itemID := itemIDBySource[doc.SourcePath]
		for _, ref := range doc.ImageRefs {
			asset := pkg.Assets[ref.ResolvedPath]
			assetID := assetIDBySource[ref.ResolvedPath]
			if assetID == "" && asset.SourcePath != "" {
				assetID = identifier.MustNew().String()
				assetIDBySource[ref.ResolvedPath] = assetID
				assets = append(assets, &model.ImportAsset{
					ID:          assetID,
					BatchID:     batchID,
					SourcePath:  asset.SourcePath,
					Data:        asset.Bytes,
					ContentHash: asset.SHA256,
					Status:      "pending",
				})
			}
			refs = append(refs, &model.ImportAssetRef{
				ID:               identifier.MustNew().String(),
				BatchID:          batchID,
				AssetID:          assetID,
				ItemID:           itemID,
				MarkdownFilePath: doc.SourcePath,
				OriginalRef:      ref.Original,
			})
		}
	}
	return assets, refs
}

func (s *Service) uploadImportAssets(ctx context.Context, bearer string, assets []*model.ImportAsset) (map[string]string, error) {
	out := map[string]string{}
	for _, asset := range assets {
		if asset.AssetURL != "" {
			out[asset.ID] = asset.AssetURL
			continue
		}
		if s.asset == nil {
			return nil, docserr.UpstreamFailed("asset service not configured")
		}
		view, err := s.asset.Upload(ctx, bearer, assetclientInput(asset), asset.Data)
		if err != nil {
			return nil, err
		}
		out[asset.ID] = view.CdnURL
		if err := s.dao.UpdateImportAssetUploaded(ctx, asset.ID, view.CdnURL); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func assetclientInput(asset *model.ImportAsset) assetclient.InitInput {
	filename := path.Base(asset.SourcePath)
	mimeType := mimeFromPath(asset.SourcePath)
	return assetclient.InitInput{
		Filename:   filename,
		Mime:       mimeType,
		Category:   "docs-import-image",
		Visibility: "public",
		Size:       int64(len(asset.Data)),
	}
}

func mimeFromPath(p string) string {
	switch strings.ToLower(path.Ext(p)) {
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".webp":
		return "image/webp"
	case ".gif":
		return "image/gif"
	case ".svg":
		return "image/svg+xml"
	default:
		return "application/octet-stream"
	}
}

func (s *Service) existingParentIDs(ctx context.Context, collectionID, versionID string, items []*model.ImportItem) (map[string]string, error) {
	locales := map[string]bool{}
	for _, item := range items {
		locales[item.Locale] = true
	}
	out := map[string]string{}
	for locale := range locales {
		docs, err := s.dao.ListDocsByCollection(ctx, collectionID, versionID, locale)
		if err != nil {
			return nil, err
		}
		for _, item := range flattenDocPaths(BuildTree(docs), "") {
			out[importDocKey(item.doc.Locale, item.path)] = item.doc.ID
		}
	}
	return out, nil
}

func (s *Service) prepareImportMutation(
	item *model.ImportItem,
	refs []*model.ImportAssetRef,
	assetURLs map[string]string,
	parentIDs map[string]string,
) (dao.ImportDocMutation, error) {
	var planned plannedImportDoc
	if err := json.Unmarshal([]byte(item.AfterDocJSON), &planned); err != nil {
		return dao.ImportDocMutation{}, err
	}
	content := rewriteImageRefs(planned.Content, refs, assetURLs)
	planned.Content = content
	planned.ParentID = parentIDs[importDocKey(planned.Locale, planned.ParentPath)]
	doc := &model.Doc{}
	if item.Action == "create" {
		doc.ID = identifier.MustNew().String()
		doc.AuthorSub = "import"
	} else {
		if err := json.Unmarshal([]byte(item.BeforeDocJSON), doc); err != nil {
			return dao.ImportDocMutation{}, err
		}
		doc.ID = item.TargetDocID
	}
	doc.CollectionID = planned.CollectionID
	doc.VersionID = planned.VersionID
	doc.ParentID = planned.ParentID
	doc.Slug = planned.Slug
	doc.Title = planned.Title
	doc.Content = planned.Content
	doc.Excerpt = planned.Excerpt
	doc.Status = planned.Status
	doc.Locale = planned.Locale
	doc.TranslationKey = planned.TranslationKey
	doc.SortOrder = planned.SortOrder
	parentIDs[importDocKey(doc.Locale, item.Path)] = doc.ID
	after := mustJSON(doc)
	return dao.ImportDocMutation{
		Action: item.Action, Doc: doc, ItemID: item.ID,
		AfterDocJSON:           after,
		TransformedContentHash: sha256Hex([]byte(planned.Content)),
	}, nil
}

func rewriteImageRefs(content string, refs []*model.ImportAssetRef, assetURLs map[string]string) string {
	out := content
	for _, ref := range refs {
		url := assetURLs[ref.AssetID]
		if url == "" {
			continue
		}
		out = strings.ReplaceAll(out, "("+ref.OriginalRef+")", "("+url+")")
		out = strings.ReplaceAll(out, `src="`+ref.OriginalRef+`"`, `src="`+url+`"`)
		out = strings.ReplaceAll(out, `src='`+ref.OriginalRef+`'`, `src='`+url+`'`)
	}
	return out
}

func parseImportSummary(raw string) ImportSummary {
	var summary ImportSummary
	_ = json.Unmarshal([]byte(raw), &summary)
	return summary
}

func plannedStatus(doc *model.Doc) string {
	if doc == nil || doc.Status == "" {
		return "published"
	}
	return doc.Status
}

func parentPath(p string) string {
	dir := path.Dir(strings.Trim(p, "/"))
	if dir == "." {
		return ""
	}
	return dir
}

func pathDepth(p string) int {
	p = strings.Trim(p, "/")
	if p == "" {
		return 0
	}
	return strings.Count(p, "/") + 1
}

func strPtr(s string) *string {
	return &s
}

func intPtr(n int) *int {
	return &n
}

func sha256Hex(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func mustJSON(v any) string {
	if v == nil {
		return "{}"
	}
	body, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(body)
}
