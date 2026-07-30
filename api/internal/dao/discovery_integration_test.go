package dao

import (
	"os"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gctx"
)

func TestPGDiscoveryQuery(t *testing.T) {
	if os.Getenv("DOCS_PG") == "" {
		t.Skip("set DOCS_PG=1 to run the docs discovery integration test")
	}
	host := docsDiscoveryEnvOr("DOCS_PG_HOST", "192.168.5.5")
	db, err := gdb.New(gdb.ConfigNode{
		Type: "pgsql", Host: host, Port: docsDiscoveryEnvOr("DOCS_PG_PORT", "5432"),
		User: docsDiscoveryEnvOr("DOCS_PG_USER", "postgres"),
		Pass: docsDiscoveryEnvOr("DOCS_PG_PASS", "postgres"), Name: "docs",
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx := gctx.New()
	tx, err := db.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = tx.Rollback() }()

	relationExists, err := tx.GetValue(`SELECT to_regclass('collection_versions') IS NOT NULL`)
	if err != nil {
		t.Fatal(err)
	}
	if !relationExists.Bool() {
		migration, readErr := os.ReadFile("../../manifest/sql/migrations/0003_versions_i18n.up.sql")
		if readErr != nil {
			t.Fatal(readErr)
		}
		if _, err := tx.Exec(string(migration)); err != nil {
			t.Fatal(err)
		}
	}
	var rows []DiscoveryRow
	if err := tx.Ctx(ctx).Raw(
		discoveryPagesSQL,
		"en", "en", "en", "https://docs.example.com", "", "https://docs.example.com", 1,
	).Scan(&rows); err != nil {
		t.Fatal(err)
	}
}

func docsDiscoveryEnvOr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
