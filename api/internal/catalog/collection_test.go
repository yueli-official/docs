package catalog

import (
	"context"
	"os"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/database/gdb"

	"github.com/yueli-official/docs/api/internal/dao"
)

func colTestEnvOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

// newTestService opens a live-PG connection, resets the schema, and returns a
// Service backed by it. Skips the test when DOCS_PG=1 is not set.
func newTestService(t *testing.T) *Service {
	t.Helper()
	if os.Getenv("DOCS_PG") == "" {
		t.Skip("set DOCS_PG=1 for live PG")
	}
	host := colTestEnvOr("DOCS_PG_HOST", "192.168.5.5")
	port := colTestEnvOr("DOCS_PG_PORT", "5432")
	user := colTestEnvOr("DOCS_PG_USER", "postgres")
	pass := colTestEnvOr("DOCS_PG_PASS", "postgres")
	db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	const dir = "../../manifest/sql/migrations/"
	ctx := context.Background()
	for _, name := range []string{"0004_doc_imports.down.sql", "0003_versions_i18n.down.sql", "0002_home_config.down.sql", "0001_init.down.sql"} {
		down, _ := os.ReadFile(dir + name)
		_, _ = db.Exec(ctx, string(down)) // ignore: tables may not exist yet
	}
	up, err := os.ReadFile(dir + "0001_init.up.sql")
	if err != nil {
		t.Fatalf("read up migration: %v", err)
	}
	if _, err := db.Exec(ctx, string(up)); err != nil {
		t.Fatalf("apply up migration: %v", err)
	}
	for _, name := range []string{"0003_versions_i18n.up.sql", "0004_doc_imports.up.sql"} {
		next, err := os.ReadFile(dir + name)
		if err != nil {
			t.Fatalf("read migration %s: %v", name, err)
		}
		if _, err := db.Exec(ctx, string(next)); err != nil {
			t.Fatalf("apply migration %s: %v", name, err)
		}
	}
	return New(dao.NewPG(db))
}

func TestCreateCollection_PersistsCoverAndIcon(t *testing.T) {
	svc := newTestService(t)
	ctx := context.Background()
	col, err := svc.CreateCollection(ctx, "owner-sub", "Sapphire Guide", "", "desc", "https://x/cover.jpg", "i-tabler-diamond")
	if err != nil {
		t.Fatal(err)
	}
	got, err := svc.GetCollectionBySlug(ctx, col.Slug)
	if err != nil {
		t.Fatal(err)
	}
	if got.CoverURL != "https://x/cover.jpg" || got.Icon != "i-tabler-diamond" {
		t.Fatalf("cover/icon not persisted: coverURL=%q icon=%q", got.CoverURL, got.Icon)
	}
}

func TestUpdateCollection_OverwritesCoverAndIcon(t *testing.T) {
	svc := newTestService(t)
	ctx := context.Background()
	// create with initial cover/icon
	col, err := svc.CreateCollection(ctx, "owner-sub", "Ruby Guide", "", "desc", "https://x/old.jpg", "i-old")
	if err != nil {
		t.Fatal(err)
	}
	// overwrite with new cover/icon
	updated, err := svc.UpdateCollection(ctx, col.ID, col.Title, "ruby-manual", col.Description, "https://x/new.jpg", "i-new")
	if err != nil {
		t.Fatal(err)
	}
	if updated.Slug != "ruby-manual" {
		t.Fatalf("slug not overwritten: %q", updated.Slug)
	}
	if updated.CoverURL != "https://x/new.jpg" || updated.Icon != "i-new" {
		t.Fatalf("cover/icon not overwritten: coverURL=%q icon=%q", updated.CoverURL, updated.Icon)
	}
}
