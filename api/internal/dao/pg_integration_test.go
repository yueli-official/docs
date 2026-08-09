package dao_test

// Schema smoke against a live PG (database "docs"). Skipped unless DOCS_PG=1.
//
//	DOCS_PG=1 DOCS_PG_HOST=192.168.5.5 go test ./internal/dao/ -p 1 -v

import (
	"context"
	"os"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"

	"github.com/yueli-official/docs/api/internal/dao"
	"github.com/yueli-official/docs/api/internal/model"
	"github.com/yueli-official/foundation/go/identifier"
)

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

// openTestDB connects to the "docs" PG database using DOCS_PG_* env vars.
func openTestDB(t *gtest.T) gdb.DB {
	host := envOr("DOCS_PG_HOST", "192.168.5.5")
	port := envOr("DOCS_PG_PORT", "5432")
	user := envOr("DOCS_PG_USER", "postgres")
	pass := envOr("DOCS_PG_PASS", "postgres")
	db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
	t.AssertNil(err)
	return db
}

// resetSchema tears down and re-applies core migrations so every run starts clean.
func resetSchema(ctx context.Context, t *gtest.T, db gdb.DB) {
	const dir = "../../manifest/sql/migrations/"
	down4, _ := os.ReadFile(dir + "0004_doc_imports.down.sql")
	_, _ = db.Exec(ctx, string(down4))
	down3, _ := os.ReadFile(dir + "0003_versions_i18n.down.sql")
	_, _ = db.Exec(ctx, string(down3))
	down, err := os.ReadFile(dir + "0001_init.down.sql")
	t.AssertNil(err)
	_, _ = db.Exec(ctx, string(down)) // ignore: tables may not exist yet
	up, err := os.ReadFile(dir + "0001_init.up.sql")
	t.AssertNil(err)
	_, err = db.Exec(ctx, string(up))
	t.AssertNil(err)
	up3, err := os.ReadFile(dir + "0003_versions_i18n.up.sql")
	t.AssertNil(err)
	_, err = db.Exec(ctx, string(up3))
	t.AssertNil(err)
}

func TestPGCollectionCRUD(t *testing.T) {
	if os.Getenv("DOCS_PG") == "" {
		t.Skip("set DOCS_PG=1 for live PG")
	}
	gtest.C(t, func(t *gtest.T) {
		ctx := gctx.New()
		db := openTestDB(t)
		resetSchema(ctx, t, db)
		p := dao.NewPG(db)
		c := &model.Collection{ID: identifier.MustNew().String(), Slug: "ps", Title: "PS教程", AuthorSub: "u1"}
		t.AssertNil(p.InsertCollection(ctx, c))
		got, err := p.GetCollectionBySlug(ctx, "ps")
		t.AssertNil(err)
		t.Assert(got.Title, "PS教程")
		c.Title = "PS Tutorial"
		t.AssertNil(p.UpdateCollection(ctx, c))
		list, err := p.ListCollections(ctx)
		t.AssertNil(err)
		t.Assert(len(list), 1)
		t.AssertNil(p.DeleteCollection(ctx, c.ID))
	})
}

func TestPGDocCRUD(t *testing.T) {
	if os.Getenv("DOCS_PG") == "" {
		t.Skip("set DOCS_PG=1 for live PG")
	}
	gtest.C(t, func(t *gtest.T) {
		ctx := gctx.New()
		db := openTestDB(t)
		resetSchema(ctx, t, db)
		p := dao.NewPG(db)
		col := &model.Collection{ID: identifier.MustNew().String(), Slug: "yueli", Title: "YUELI", AuthorSub: "u1"}
		t.AssertNil(p.InsertCollection(ctx, col))
		version := &model.CollectionVersion{
			ID:           identifier.MustNew().String(),
			CollectionID: col.ID,
			Key:          "default",
			Label:        "默认版本",
			Status:       "published",
			IsDefault:    true,
		}
		t.AssertNil(p.InsertCollectionVersion(ctx, version))
		d := &model.Doc{
			ID:             identifier.MustNew().String(),
			CollectionID:   col.ID,
			VersionID:      version.ID,
			Slug:           "intro",
			Title:          "Intro",
			Locale:         "en",
			TranslationKey: "intro",
			AuthorSub:      "u1",
		}
		t.AssertNil(p.InsertDoc(ctx, d))
		got, err := p.GetDocByID(ctx, d.ID)
		t.AssertNil(err)
		t.Assert(got.Title, "Intro")
		t.AssertNil(p.SoftDeleteDoc(ctx, d.ID))
		list, err := p.ListDocsByCollection(ctx, col.ID, version.ID, "en")
		t.AssertNil(err)
		t.Assert(len(list), 0) // soft-deleted excluded
	})
}

// TestSchemaApplies is a raw-gdb schema smoke: insert via db.Model(...).Data(g.Map{...}).
// It does NOT reference model.Collection / model.Doc or CRUD dao methods (Task 3/4).
func TestSchemaApplies(t *testing.T) {
	if os.Getenv("DOCS_PG") == "" {
		t.Skip("set DOCS_PG=1 for live PG")
	}
	gtest.C(t, func(t *gtest.T) {
		ctx := gctx.New()
		db := openTestDB(t)
		resetSchema(ctx, t, db)

		colID := identifier.MustNew().String()
		_, err := db.Model("collections").Ctx(ctx).Data(g.Map{
			"id": colID, "slug": "sapphire", "title": "Sapphire", "author_sub": "u1",
		}).Insert()
		t.AssertNil(err)
		versionID := identifier.MustNew().String()
		_, err = db.Model("collection_versions").Ctx(ctx).Data(g.Map{
			"id": versionID, "collection_id": colID, "key": "default", "label": "默认版本",
			"status": "published", "is_default": true,
		}).Insert()
		t.AssertNil(err)

		rootID := identifier.MustNew().String()
		_, err = db.Model("docs").Ctx(ctx).Data(g.Map{
			"id": rootID, "collection_id": colID, "version_id": versionID,
			"slug": "adjust", "title": "Adjust", "locale": "en",
			"translation_key": "adjust", "author_sub": "u1",
		}).Insert()
		t.AssertNil(err)

		// child references parent → exercises parent_id self-FK
		_, err = db.Model("docs").Ctx(ctx).Data(g.Map{
			"id": identifier.MustNew().String(), "collection_id": colID, "version_id": versionID,
			"parent_id": rootID, "slug": "gamma", "title": "Gamma", "locale": "en",
			"translation_key": "gamma", "author_sub": "u1",
		}).Insert()
		t.AssertNil(err)

		n, err := db.Model("docs").Ctx(ctx).Where("collection_id", colID).Count()
		t.AssertNil(err)
		t.Assert(n, 2)

		// cascade: deleting the collection removes its docs
		_, err = db.Model("collections").Ctx(ctx).Where("id", colID).Delete()
		t.AssertNil(err)
		n2, _ := db.Model("docs").Ctx(ctx).Where("collection_id", colID).Count()
		t.Assert(n2, 0)
	})
}

func TestPGImportSchemaApplies(t *testing.T) {
	if os.Getenv("DOCS_PG") == "" {
		t.Skip("set DOCS_PG=1 for live PG")
	}
	gtest.C(t, func(t *gtest.T) {
		ctx := gctx.New()
		db := openTestDB(t)
		resetSchema(ctx, t, db)

		up4, err := os.ReadFile("../../manifest/sql/migrations/0004_doc_imports.up.sql")
		t.AssertNil(err)
		_, err = db.Exec(ctx, string(up4))
		t.AssertNil(err)

		colID := identifier.MustNew().String()
		versionID := identifier.MustNew().String()
		batchID := identifier.MustNew().String()
		itemID := identifier.MustNew().String()
		assetID := identifier.MustNew().String()

		_, err = db.Model("collections").Ctx(ctx).Data(g.Map{
			"id": colID, "slug": "imported", "title": "Imported", "author_sub": "u1",
		}).Insert()
		t.AssertNil(err)
		_, err = db.Model("collection_versions").Ctx(ctx).Data(g.Map{
			"id": versionID, "collection_id": colID, "key": "default", "label": "默认版本",
			"status": "published", "is_default": true,
		}).Insert()
		t.AssertNil(err)
		_, err = db.Model("doc_import_batches").Ctx(ctx).Data(g.Map{
			"id": batchID, "collection_id": colID, "version_id": versionID,
			"default_locale": "zh-CN", "mode": "upsert", "status": "checked",
			"created_by": "u1", "summary_json": g.Map{"creates": 1},
		}).Insert()
		t.AssertNil(err)
		_, err = db.Model("doc_import_items").Ctx(ctx).Data(g.Map{
			"id": itemID, "batch_id": batchID, "locale": "zh-CN", "version_key": "default",
			"path": "guide/install", "source_markdown_path": "zh-CN/default/guide/install.md",
			"title": "安装", "slug": "install", "translation_key": "install-guide",
			"action": "create", "issues_json": []string{},
		}).Insert()
		t.AssertNil(err)
		_, err = db.Model("doc_import_assets").Ctx(ctx).Data(g.Map{
			"id": assetID, "batch_id": batchID, "source_path": "zh-CN/default/guide/images/a.png",
			"content_hash": "abc", "status": "pending",
		}).Insert()
		t.AssertNil(err)
		_, err = db.Model("doc_import_asset_refs").Ctx(ctx).Data(g.Map{
			"batch_id": batchID, "asset_id": assetID, "item_id": itemID,
			"markdown_file_path": "zh-CN/default/guide/install.md",
			"original_ref":       "./images/a.png", "rewritten_ref": "",
		}).Insert()
		t.AssertNil(err)

		n, err := db.Model("doc_import_asset_refs").Ctx(ctx).Where("batch_id", batchID).Count()
		t.AssertNil(err)
		t.Assert(n, 1)

		down4, err := os.ReadFile("../../manifest/sql/migrations/0004_doc_imports.down.sql")
		t.AssertNil(err)
		_, err = db.Exec(ctx, string(down4))
		t.AssertNil(err)
	})
}
