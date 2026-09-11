package projectdocs

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	_ "github.com/lib/pq"
	"github.com/yueli-official/docs/api/internal/catalog"
	"github.com/yueli-official/docs/api/internal/dao"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	auth "github.com/yueli-official/foundation/go/auth"
	"github.com/yueli-official/foundation/go/authorization"
	"github.com/yueli-official/foundation/go/identifier"
)

type releaseFixture struct {
	release       Release
	asset         ReleaseAsset
	data          []byte
	failure       error
	downloads     int
	afterDownload func()
}

func (f *releaseFixture) Latest(context.Context, string, string) (Release, ReleaseAsset, error) {
	return f.release, f.asset, f.failure
}
func (f *releaseFixture) Download(context.Context, ReleaseAsset) ([]byte, string, error) {
	f.downloads++
	if f.afterDownload != nil {
		f.afterDownload()
	}
	sum := sha256.Sum256(f.data)
	return f.data, "sha256:" + hex.EncodeToString(sum[:]), f.failure
}
func (f *releaseFixture) publish(t *testing.T, n int64, title string) {
	t.Helper()
	var b bytes.Buffer
	w := zip.NewWriter(&b)
	for _, name := range []string{"index.md", "guide.md"} {
		p, e := w.Create(name)
		if e != nil {
			t.Fatal(e)
		}
		_, _ = p.Write([]byte("---\ntitle: " + title + "\nslug: " + strings.TrimSuffix(name, ".md") + "-custom\n---\n# " + title))
	}
	if e := w.Close(); e != nil {
		t.Fatal(e)
	}
	f.data = b.Bytes()
	f.release = Release{ID: n, Tag: "v" + string(rune('0'+n))}
	f.asset = ReleaseAsset{ID: n, Name: "docs.zip", Size: int64(len(f.data))}
}

func TestWorkerPostgresLifecycle(t *testing.T) {
	dsn := os.Getenv("DOCS_SYNC_TEST_DSN")
	if dsn == "" {
		t.Skip("set DOCS_SYNC_TEST_DSN; creates an isolated temporary schema")
	}
	ctx := context.Background()
	u, e := url.Parse(dsn)
	if e != nil {
		t.Fatal(e)
	}
	admin, e := sql.Open("postgres", dsn)
	if e != nil {
		t.Fatal(e)
	}
	defer admin.Close()
	schema := "docs_sync_test_" + strings.ReplaceAll(identifier.MustNew().String(), "-", "")
	if _, e = admin.ExecContext(ctx, `CREATE SCHEMA `+schema); e != nil {
		t.Fatal(e)
	}
	defer func() {
		if _, e := admin.ExecContext(ctx, `DROP SCHEMA `+schema+` CASCADE`); e != nil {
			t.Error(e)
		}
	}()
	q := u.Query()
	q.Set("search_path", schema+",public")
	u.RawQuery = q.Encode()
	db, e := sql.Open("postgres", u.String())
	if e != nil {
		t.Fatal(e)
	}
	defer db.Close()
	migrations, e := filepath.Glob("../../manifest/sql/migrations/*.up.sql")
	if e != nil {
		t.Fatal(e)
	}
	for _, file := range migrations {
		data, e := os.ReadFile(file)
		if e != nil {
			t.Fatal(e)
		}
		if _, e = db.ExecContext(ctx, string(data)); e != nil {
			t.Fatalf("migration %s: %v", filepath.Base(file), e)
		}
	}
	password, _ := u.User.Password()
	gf, e := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: u.Hostname(), Port: u.Port(), User: u.User.Username(), Pass: password, Name: strings.TrimPrefix(u.Path, "/"), Extra: "search_path=" + schema + ",public sslmode=disable"})
	if e != nil {
		t.Fatal(e)
	}
	defer gf.Close(ctx)
	cat := catalog.New(dao.NewPG(gf))
	col, e := cat.CreateCollectionWithSetup(ctx, catalog.CreateCollectionInput{AuthorSub: "owner", Title: "Source Test", Slug: "source-test", DefaultLocale: "zh-CN", SemanticVersion: "1.0.0"})
	if e != nil {
		t.Fatal(e)
	}
	module, e := authorization.NewMemory(authorization.MustCompile(docsauthz.Definition()), authorization.MemoryOptions{RootScopeID: docsauthz.RootScopeID, ProtectedSubjects: []authorization.SubjectRef{{Kind: authorization.SubjectUser, ID: "owner"}}, Constraints: docsauthz.ConstraintEvaluators(), Predicates: docsauthz.PredicateEvaluators()})
	if e != nil {
		t.Fatal(e)
	}
	scope, _ := auth.PersonalScope("docs-main-web", string(docsauthz.CapabilityImportManage))
	revoked := false
	endpoint := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if revoked {
			w.WriteHeader(401)
			return
		}
		json.NewEncoder(w).Encode(map[string]any{"userKey": "owner", "scopes": []string{scope}})
	}))
	defer endpoint.Close()
	verifier, e := auth.NewPersonalTokenVerifier(endpoint.URL, "docs-main-web", nil)
	if e != nil {
		t.Fatal(e)
	}
	vault, e := NewVault(base64.StdEncoding.EncodeToString([]byte(strings.Repeat("x", 32))))
	if e != nil {
		t.Fatal(e)
	}
	service := New(db, cat, docsauthz.New(module), verifier, vault)
	fixture := &releaseFixture{}
	fixture.publish(t, 1, "first")
	service.github = fixture
	cat.WithImportGuard(func(ctx context.Context, token string) error {
		_, e := service.credential(ctx, token, "owner")
		return e
	})
	source, e := service.Create(ctx, "owner", CreateInput{Repository: "example/project", Collection: col.Slug, DefaultLocale: "zh-CN", Token: "pat_test"})
	if e != nil {
		t.Fatal(e)
	}
	if source.AutoCheck || source.Status != "idle" {
		t.Fatalf("manual source should start idle: %#v", source)
	}
	service.runDue(ctx)
	service.check(ctx, source.ID)
	initialRuns, e := service.Store.Runs(ctx, source.ID)
	if e != nil || len(initialRuns) != 0 {
		t.Fatalf("manual source was automatically checked: %d runs, %v", len(initialRuns), e)
	}
	if e = service.Queue(ctx, source.ID); e != nil {
		t.Fatal(e)
	}
	service.check(ctx, source.ID)
	assertState := func(status string) {
		t.Helper()
		s, e := service.Store.Get(ctx, source.ID)
		if e != nil || s.Status != status {
			t.Fatalf("source status=%s error=%s err=%v", s.Status, s.LastError, e)
		}
	}
	assertCount := func(want int) {
		t.Helper()
		var count int
		if e := db.QueryRowContext(ctx, `SELECT count(*) FROM docs WHERE collection_id=$1 AND deleted_at IS NULL`, col.ID).Scan(&count); e != nil || count != want {
			t.Fatalf("count=%d want=%d err=%v", count, want, e)
		}
	}
	assertState("completed")
	assertCount(2)
	_, e = db.ExecContext(ctx, `UPDATE project_doc_sources SET next_check_at=NOW() WHERE id=$1`, source.ID)
	if e != nil {
		t.Fatal(e)
	}
	service.runDue(ctx)
	manualRuns, _ := service.Store.Runs(ctx, source.ID)
	if len(manualRuns) != 1 {
		t.Fatal("manual check scheduled another run")
	}
	automatic := true
	if e = service.Update(ctx, source.ID, "owner", "", nil, &automatic); e != nil {
		t.Fatal(e)
	}
	service.runDue(ctx)
	assertState("skipped")
	automatic = false
	if e = service.Update(ctx, source.ID, "owner", "", nil, &automatic); e != nil {
		t.Fatal(e)
	}
	saved, e := service.Store.Get(ctx, source.ID)
	if e != nil || saved.AutoCheck || !saved.Enabled {
		t.Fatal("automatic check preference did not persist independently of pause")
	}
	if e = service.Queue(ctx, source.ID); e != nil {
		t.Fatal(e)
	}
	service.check(ctx, source.ID)
	assertState("skipped")
	if fixture.downloads != 1 {
		t.Fatal("unchanged asset downloaded again")
	}
	fixture.publish(t, 2, "second")
	_ = service.Queue(ctx, source.ID)
	service.check(ctx, source.ID)
	assertState("completed")
	assertCount(2)
	fixture.failure = errors.New("fixture download failed")
	_ = service.Queue(ctx, source.ID)
	service.check(ctx, source.ID)
	assertState("failed")
	assertCount(2)
	fixture.failure = nil
	fixture.publish(t, 3, "third")
	fixture.afterDownload = func() { revoked = true }
	_ = service.Queue(ctx, source.ID)
	service.check(ctx, source.ID)
	assertState("failed")
	assertCount(2)
	var title string
	if e = db.QueryRowContext(ctx, `SELECT title FROM docs WHERE collection_id=$1 LIMIT 1`, col.ID).Scan(&title); e != nil || title != "second" {
		t.Fatalf("late revocation committed: %s %v", title, e)
	}
	revoked = false
	fixture.afterDownload = nil
	_ = service.Queue(ctx, source.ID)
	service.check(ctx, source.ID)
	assertState("completed")
	assertCount(2)
	// Simulate a crash after the catalog transaction but before source completion.
	_, e = db.ExecContext(ctx, `UPDATE project_doc_source_runs SET status='running' WHERE id=(SELECT id FROM project_doc_source_runs WHERE source_id=$1 ORDER BY started_at DESC LIMIT 1)`, source.ID)
	if e != nil {
		t.Fatal(e)
	}
	_, _ = db.ExecContext(ctx, `UPDATE project_doc_sources SET last_digest='',last_asset_id=0 WHERE id=$1`, source.ID)
	_ = service.Queue(ctx, source.ID)
	restarted := New(db, cat, docsauthz.New(module), verifier, vault)
	restarted.github = fixture
	restarted.check(ctx, source.ID)
	assertState("skipped")
	assertCount(2)
	fixture.publish(t, 4, "bad package")
	fixture.data = []byte("not a ZIP archive")
	_ = service.Queue(ctx, source.ID)
	service.check(ctx, source.ID)
	assertState("failed")
	assertCount(2)
	// Concurrent workers are fenced by the database, not a process-local mutex.
	conn, e := db.Conn(ctx)
	if e != nil {
		t.Fatal(e)
	}
	_, e = conn.ExecContext(ctx, `SELECT pg_advisory_lock(hashtextextended('docs-project-source:' || $1,0))`, source.ID)
	if e != nil {
		t.Fatal(e)
	}
	_ = service.Queue(ctx, source.ID)
	before := fixture.downloads
	service.check(ctx, source.ID)
	if fixture.downloads != before {
		t.Fatal("lock ignored")
	}
	_, _ = conn.ExecContext(ctx, `SELECT pg_advisory_unlock(hashtextextended('docs-project-source:' || $1,0))`, source.ID)
	conn.Close()
	revoked = true
	service.check(ctx, source.ID)
	assertState("paused")
	assertCount(2)
	// New process/service instance observes persisted pause and cannot import.
	recovered := New(db, cat, docsauthz.New(module), verifier, vault)
	recovered.github = fixture
	recovered.check(ctx, source.ID)
	assertState("paused")
}
