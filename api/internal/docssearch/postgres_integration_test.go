package docssearch

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func TestPostgresCollectionCascadeAndProjectionShareTransaction(t *testing.T) {
	dsn := os.Getenv("SEARCH_CONSUMER_PG_DSN")
	if dsn == "" {
		t.Skip("SEARCH_CONSUMER_PG_DSN is not set")
	}
	ctx := context.Background()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	site := "test-" + uuid.NewString()
	index, err := NewPostgres(ctx, db, site)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_, _ = db.ExecContext(ctx, `DELETE FROM search_instances WHERE instance_key=$1`, "docs.search."+site)
	})

	collectionID, versionID, docID := uuid.NewString(), uuid.NewString(), uuid.NewString()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tx.ExecContext(ctx,
		`INSERT INTO collections (id,slug,title,author_sub) VALUES ($1::uuid,$2,'Search Docs','author')`,
		collectionID, "search-"+collectionID,
	); err != nil {
		t.Fatal(err)
	}
	if _, err := tx.ExecContext(ctx, `
		INSERT INTO collection_versions (id,collection_id,key,label,status,is_default)
		VALUES ($2::uuid,$1::uuid,'default','Default','published',TRUE)
	`, collectionID, versionID); err != nil {
		t.Fatal(err)
	}
	if _, err := tx.ExecContext(ctx, `
		INSERT INTO docs (
			id,collection_id,version_id,slug,title,content,status,locale,
			translation_key,author_sub
		) VALUES (
			$3::uuid,$1::uuid,$2::uuid,'guide','中文检索能力','事务一致的搜索投影',
			'published','en',$3::text,'author'
		)
	`, collectionID, versionID, docID); err != nil {
		t.Fatal(err)
	}
	if err := index.SubtreeHook(docID)(ctx, tx); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _, _ = db.ExecContext(ctx, `DELETE FROM collections WHERE id=$1::uuid`, collectionID) })

	assertHits := func(want int) {
		t.Helper()
		page, err := index.Search(ctx, "中文检索", collectionID, versionID, "en", 20)
		if err != nil {
			t.Fatal(err)
		}
		if len(page.Hits) != want {
			t.Fatalf("search hits = %#v, want %d", page.Hits, want)
		}
	}
	assertHits(1)
	facetPage, err := index.Search(ctx, "中文检索", "", "", "en", 20)
	if err != nil {
		t.Fatal(err)
	}
	if len(facetPage.Facets) != 1 || len(facetPage.Facets[0].Buckets) != 1 ||
		facetPage.Facets[0].Buckets[0].Value != collectionID ||
		facetPage.Facets[0].Buckets[0].Count != 1 {
		t.Fatalf("collection facets = %#v", facetPage.Facets)
	}

	tx, err = db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := index.DeleteCollectionHook(collectionID)(ctx, tx); err != nil {
		t.Fatal(err)
	}
	if _, err := tx.ExecContext(ctx, `DELETE FROM collections WHERE id=$1::uuid`, collectionID); err != nil {
		t.Fatal(err)
	}
	if err := tx.Rollback(); err != nil {
		t.Fatal(err)
	}
	assertHits(1)

	tx, err = db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := index.DeleteCollectionHook(collectionID)(ctx, tx); err != nil {
		t.Fatal(err)
	}
	if _, err := tx.ExecContext(ctx, `DELETE FROM collections WHERE id=$1::uuid`, collectionID); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
	assertHits(0)
}
