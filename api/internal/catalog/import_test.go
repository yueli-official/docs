package catalog

import (
	"archive/zip"
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/yueli-official/docs/api/internal/assetclient"
	"github.com/yueli-official/docs/api/internal/importkit"
	"github.com/yueli-official/docs/api/internal/model"
	"github.com/yueli-official/foundation/go/problem"
)

func TestImportPackageErrorPreservesUnsupportedCompressionSemantics(t *testing.T) {
	err := importPackageError(&importkit.UnsupportedCompressionError{Method: 99})
	value, ok, resolveErr := problem.FromError(err, "trace-import")
	if resolveErr != nil || !ok {
		t.Fatalf("FromError() = %#v, %v, %v", value, ok, resolveErr)
	}
	if value.Code != "docs.import.compression_unsupported" || value.Params["method"] != 99 {
		t.Fatalf("Problem = %#v", value)
	}
}

func TestImportAssetUsesDedicatedImageProfile(t *testing.T) {
	in := assetclientInput(&model.ImportAsset{SourcePath: "guide/images/hero.gif", Data: []byte("gif")})
	if in.Category != "docs-import-image" || in.Mime != "image/gif" || in.Visibility != "public" {
		t.Fatalf("import asset input = %+v", in)
	}
}

func TestImportAssetsDeduplicateIdenticalContentAcrossLocalePaths(t *testing.T) {
	pkg := &importkit.Package{
		Docs: []importkit.DocFile{
			{SourcePath: "guide.md", ImageRefs: []importkit.ImageRef{{Original: "./images/a.png", ResolvedPath: "images/a.png"}}},
			{SourcePath: "locales/en-US/guide.md", ImageRefs: []importkit.ImageRef{{Original: "../../images/a.png", ResolvedPath: "locales/en-US/images/a.png"}}},
		},
		Assets: map[string]importkit.AssetFile{
			"images/a.png":               {SourcePath: "images/a.png", Bytes: []byte("same"), SHA256: "same-hash"},
			"locales/en-US/images/a.png": {SourcePath: "locales/en-US/images/a.png", Bytes: []byte("same"), SHA256: "same-hash"},
		},
	}
	assets, refs := importAssetsAndRefs("batch", pkg, map[string]string{
		"guide.md": "zh-item", "locales/en-US/guide.md": "en-item",
	})
	if len(assets) != 1 {
		t.Fatalf("assets=%d, want one content-addressed upload", len(assets))
	}
	if len(refs) != 2 || refs[0].AssetID != refs[1].AssetID {
		t.Fatalf("refs did not reuse the same asset: %+v", refs)
	}
}

func importZipBytes(t *testing.T, files map[string]string) []byte {
	t.Helper()
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	for name, body := range files {
		w, err := zw.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := w.Write([]byte(body)); err != nil {
			t.Fatal(err)
		}
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func TestPreflightImportCreatesCheckedBatch(t *testing.T) {
	svc := newTestService(t)
	ctx := context.Background()
	col, err := svc.CreateCollection(ctx, "owner-sub", "Quickstart", "quickstart", "desc", "", "")
	if err != nil {
		t.Fatal(err)
	}
	data := importZipBytes(t, map[string]string{
		"docs.json":          `{"schemaVersion":1,"defaultLocale":"zh-CN","locales":{"zh-CN":"."}}`,
		"guide/index.md":     "---\nid: guide\ntitle: 指南\norder: 2\n---\n![图](./images/a.png)",
		"guide/images/a.png": "png",
	})
	batch, summary, err := svc.PreflightImport(ctx, ImportUploadInput{
		Filename:   "docs.zip",
		Data:       data,
		Author:     "owner-sub",
		Collection: "quickstart", DefaultLocale: "zh-CN", Mode: "upsert",
	})
	if err != nil {
		t.Fatal(err)
	}
	if batch.Status != "checked" || batch.CollectionID != col.ID {
		t.Fatalf("batch not checked: %+v", batch)
	}
	if summary.Blocking || summary.Creates != 1 || summary.Images != 1 {
		t.Fatalf("summary mismatch: %+v", summary)
	}
	items, err := svc.dao.ListImportItems(ctx, batch.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 1 || items[0].Action != "create" || items[0].Path != "guide" {
		t.Fatalf("items mismatch: %+v", items)
	}
	refs, err := svc.dao.ListImportAssetRefs(ctx, batch.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(refs) != 1 || refs[0].OriginalRef != "./images/a.png" {
		t.Fatalf("refs mismatch: %+v", refs)
	}
	history, total, err := svc.ListImports(ctx, "", 1, 20)
	if err != nil {
		t.Fatal(err)
	}
	if total != 1 || len(history) != 1 || history[0].ID != batch.ID || history[0].CreatedAt.IsZero() {
		t.Fatalf("import history mismatch: %+v", history)
	}
}

func TestPreflightImportMissingImageBlocksConfirm(t *testing.T) {
	svc := newTestService(t)
	ctx := context.Background()
	if _, err := svc.CreateCollection(ctx, "owner-sub", "Quickstart", "quickstart", "desc", "", ""); err != nil {
		t.Fatal(err)
	}
	data := importZipBytes(t, map[string]string{
		"docs.json":      `{"schemaVersion":1,"defaultLocale":"zh-CN","locales":{"zh-CN":"."}}`,
		"guide/index.md": "---\nid: guide\ntitle: 指南\n---\n![图](./images/missing.png)",
	})
	batch, summary, err := svc.PreflightImport(ctx, ImportUploadInput{
		Filename:   "docs.zip",
		Data:       data,
		Author:     "owner-sub",
		Collection: "quickstart", DefaultLocale: "zh-CN", Mode: "upsert",
	})
	if err != nil {
		t.Fatal(err)
	}
	if batch.Status != "checked" {
		t.Fatalf("batch status=%q", batch.Status)
	}
	if !summary.Blocking || summary.Errors == 0 {
		t.Fatalf("expected blocking summary: %+v", summary)
	}
}

func TestConfirmImportAndRollback(t *testing.T) {
	fakeAssets := &assetclient.Fake{}
	svc := newTestService(t).WithAssets(fakeAssets, "")
	ctx := context.Background()
	if _, err := svc.CreateCollection(ctx, "owner-sub", "Quickstart", "quickstart", "desc", "", ""); err != nil {
		t.Fatal(err)
	}
	data := importZipBytes(t, map[string]string{
		"docs.json":          `{"schemaVersion":1,"defaultLocale":"zh-CN","locales":{"zh-CN":"."}}`,
		"guide/index.md":     "---\nid: guide\ntitle: 指南\norder: 2\n---\n![图](./images/a.png)",
		"guide/images/a.png": "png",
	})
	batch, summary, err := svc.PreflightImport(ctx, ImportUploadInput{
		Filename:   "docs.zip",
		Data:       data,
		Author:     "owner-sub",
		Collection: "quickstart", DefaultLocale: "zh-CN", Mode: "upsert",
	})
	if err != nil {
		t.Fatal(err)
	}
	if summary.Blocking {
		t.Fatalf("unexpected blocking summary: %+v", summary)
	}
	done, _, err := svc.ConfirmImport(ctx, batch.ID, "bearer", "owner-sub")
	if err != nil {
		t.Fatal(err)
	}
	if done.Status != "completed" {
		t.Fatalf("status=%q", done.Status)
	}
	doc, err := svc.GetPublishedDocByPath(ctx, "quickstart", "default", "guide", "zh-CN")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(doc.Content, "https://asset.test/docs/a.png") {
		t.Fatalf("image was not rewritten: %q", doc.Content)
	}
	rolled, err := svc.RollbackImport(ctx, batch.ID, "owner-sub")
	if err != nil {
		t.Fatal(err)
	}
	if rolled.Status != "rolled_back" {
		t.Fatalf("rollback status=%q", rolled.Status)
	}
	if _, err := svc.GetPublishedDocByPath(ctx, "quickstart", "default", "guide", "zh-CN"); err == nil {
		t.Fatal("expected imported doc to disappear after rollback")
	}
}
