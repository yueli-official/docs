package assetreferences

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func TestSourceTracksImportSyncDeleteAndRestore(t *testing.T) {
	dsn := os.Getenv("REFERENCE_TEST_DATABASE_URL")
	if dsn == "" {
		t.Skip("set REFERENCE_TEST_DATABASE_URL")
	}
	db, e := sql.Open("postgres", dsn)
	if e != nil {
		t.Fatal(e)
	}
	defer db.Close()
	ctx := context.Background()
	tx, e := db.BeginTx(ctx, nil)
	if e != nil {
		t.Fatal(e)
	}
	defer tx.Rollback()
	_, e = tx.Exec(`CREATE TEMP TABLE collections(id text,title text,slug text,cover_asset_id text) ON COMMIT DROP; CREATE TEMP TABLE docs(id text,title text,content text,collection_id text,deleted_at timestamptz,status text) ON COMMIT DROP; INSERT INTO collections VALUES('c','Collection','collection',NULL); INSERT INTO docs VALUES('a','Imported','![image](/media/34kRPmE5SHCmWUG0tPXBi?v=1)','c',NULL,'published'),('b','Synced copy','![image](/media/34kRPmE5SHCmWUG0tPXBi?v=1)','c',NULL,'archived')`)
	if e != nil {
		t.Fatal(e)
	}
	assert := func(want int) {
		t.Helper()
		snapshots, err := Source("https://docs.test", "")(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if len(snapshots[0].References) != want {
			t.Fatalf("references=%d want=%d", len(snapshots[0].References), want)
		}
	}
	assert(2) // Repeated GitHub imports/cloned versions remain separate business objects.
	if _, e = tx.Exec(`UPDATE docs SET deleted_at=NOW() WHERE id='a'`); e != nil {
		t.Fatal(e)
	}
	assert(1)
	if _, e = tx.Exec(`UPDATE docs SET deleted_at=NULL WHERE id='a'`); e != nil {
		t.Fatal(e)
	}
	assert(2)
	if _, e = tx.Exec(`UPDATE docs SET content='' WHERE id='a'`); e != nil {
		t.Fatal(e)
	}
	assert(1)
	if _, e = tx.Exec(`DELETE FROM collections`); e != nil {
		t.Fatal(e)
	}
	assert(0)
}
