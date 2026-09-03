package catalog

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	_ "github.com/lib/pq"
	"github.com/yueli-official/docs/api/internal/dao"
	"github.com/yueli-official/docs/api/internal/docsurls"
	"github.com/yueli-official/foundation/go/identifier"
)

func TestCreateCollectionWithPostgresURLLifecycleDoesNotDeadlock(t *testing.T) {
	dsn := os.Getenv("DOCS_URL_INTEGRATION_PG_DSN")
	if dsn == "" {
		t.Skip("DOCS_URL_INTEGRATION_PG_DSN is not set")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	store, err := gdb.New(gdb.ConfigNode{
		Type: "pgsql", Host: colTestEnvOr("DOCS_PG_HOST", "192.168.5.5"),
		Port: colTestEnvOr("DOCS_PG_PORT", "5432"), User: colTestEnvOr("DOCS_PG_USER", "postgres"),
		Pass: colTestEnvOr("DOCS_PG_PASS", "postgres"), Name: colTestEnvOr("DOCS_PG_NAME", "doctor_docs_main"),
	})
	if err != nil {
		t.Fatal(err)
	}
	moduleDB, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer moduleDB.Close()
	lifecycle, err := docsurls.NewPostgres(ctx, moduleDB, "docs:docs-main", "http://192.168.5.7:3003", "en")
	if err != nil {
		t.Fatal(err)
	}
	svc := New(dao.NewPG(store)).WithURLLifecycle(lifecycle)
	title := "Deadlock regression " + identifier.MustNew().String()
	collection, err := svc.CreateCollection(ctx, "TestA123", title, "", "", "", "")
	if err != nil {
		t.Fatalf("create collection: %v", err)
	}
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cleanupCancel()
		if err := svc.DeleteCollection(cleanupCtx, collection.ID); err != nil {
			t.Errorf("cleanup collection: %v", err)
		}
	})
}
