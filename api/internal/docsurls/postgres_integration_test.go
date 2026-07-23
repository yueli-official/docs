package docsurls

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/yueli-official/foundation/go/urllifecycle"
)

func TestPostgresCollectionSubtreeRollbackTogether(t *testing.T) {
	dsn := os.Getenv("URL_LIFECYCLE_CONSUMER_PG_DSN")
	if dsn == "" {
		t.Skip("URL_LIFECYCLE_CONSUMER_PG_DSN is not set")
	}
	ctx := context.Background()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	lifecycle, err := NewPostgres(ctx, db, "docs-test:"+uuid.NewString(), "https://docs.test", "en")
	if err != nil {
		t.Fatal(err)
	}
	collectionID, versionID, docID := uuid.NewString(), uuid.NewString(), uuid.NewString()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tx.ExecContext(ctx,
		`INSERT INTO collections (id, slug, title, author_sub) VALUES ($1::uuid, 'atomic', 'Atomic', 'author')`,
		collectionID,
	); err != nil {
		t.Fatal(err)
	}
	if _, err := tx.ExecContext(ctx, `
INSERT INTO collection_versions (id, collection_id, key, label, status, is_default)
VALUES ($2::uuid, $1::uuid, 'default', 'Default', 'published', TRUE)`,
		collectionID, versionID,
	); err != nil {
		t.Fatal(err)
	}
	if _, err := tx.ExecContext(ctx, `
INSERT INTO docs (
    id, collection_id, version_id, slug, title, content, status, locale,
    translation_key, author_sub
) VALUES (
    $3::uuid, $1::uuid, $2::uuid, 'guide', 'Guide', 'body', 'published', 'en',
    $3::text, 'author'
)`, collectionID, versionID, docID); err != nil {
		t.Fatal(err)
	}
	if err := lifecycle.ReconcileCollection(ctx, tx, collectionID, "atomic rollback test"); err != nil {
		t.Fatal(err)
	}
	if err := tx.Rollback(); err != nil {
		t.Fatal(err)
	}
	resolution, err := lifecycle.Resolver().Resolve(ctx, urllifecycle.Lookup{EscapedPath: "/atomic/guide"})
	if err != nil {
		t.Fatal(err)
	}
	if resolution.Kind != urllifecycle.ResolutionUnknown {
		t.Fatalf("URL subtree survived rollback: %#v", resolution)
	}
}
