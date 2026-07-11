package catalog

import (
	"archive/zip"
	"bytes"
	"context"
	"strings"
	"testing"

	"platform/products/docs/api/internal/assetclient"
	"platform/products/docs/api/internal/model"
)

func TestImportAssetUsesDedicatedImageProfile(t *testing.T) {
	in := assetclientInput(&model.ImportAsset{SourcePath: "guide/images/hero.gif", Data: []byte("gif")})
	if in.Category != "docs-import-image" || in.Mime != "image/gif" || in.Visibility != "public" {
		t.Fatalf("import asset input = %+v", in)
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
		"manifest.json":                    `{"schemaVersion":1,"collection":"quickstart","version":"default","defaultLocale":"zh-CN","locales":["zh-CN"],"mode":"upsert"}`,
		"zh-CN/default/guide/index.md":     "---\ntitle: 指南\norder: 2\n---\n![图](./images/a.png)",
		"zh-CN/default/guide/images/a.png": "png",
	})
	batch, summary, err := svc.PreflightImport(ctx, ImportUploadInput{
		Filename: "docs.zip",
		Data:     data,
		Author:   "owner-sub",
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
}

func TestPreflightImportMissingImageBlocksConfirm(t *testing.T) {
	svc := newTestService(t)
	ctx := context.Background()
	if _, err := svc.CreateCollection(ctx, "owner-sub", "Quickstart", "quickstart", "desc", "", ""); err != nil {
		t.Fatal(err)
	}
	data := importZipBytes(t, map[string]string{
		"manifest.json":                `{"schemaVersion":1,"collection":"quickstart","version":"default","defaultLocale":"zh-CN","locales":["zh-CN"],"mode":"upsert"}`,
		"zh-CN/default/guide/index.md": "---\ntitle: 指南\n---\n![图](./images/missing.png)",
	})
	batch, summary, err := svc.PreflightImport(ctx, ImportUploadInput{
		Filename: "docs.zip",
		Data:     data,
		Author:   "owner-sub",
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
		"manifest.json":                    `{"schemaVersion":1,"collection":"quickstart","version":"default","defaultLocale":"zh-CN","locales":["zh-CN"],"mode":"upsert"}`,
		"zh-CN/default/guide/index.md":     "---\ntitle: 指南\norder: 2\n---\n![图](./images/a.png)",
		"zh-CN/default/guide/images/a.png": "png",
	})
	batch, summary, err := svc.PreflightImport(ctx, ImportUploadInput{
		Filename: "docs.zip",
		Data:     data,
		Author:   "owner-sub",
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
