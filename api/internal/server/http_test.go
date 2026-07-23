package server_test

// Full HTTP round-trip against a loopback server + live PG (database "docs").
// Skipped unless DOCS_PG_HOST is set:
//
//	DOCS_PG_HOST=192.168.5.5 DOCS_PG_USER=postgres DOCS_PG_PASS=postgres \
//	  go test -run TestDocsRoundTrip ./products/docs/api/internal/server/... -p 1

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/test/gtest"
	_ "github.com/lib/pq"

	"platform/products/docs/api/internal/assetclient"
	"platform/products/docs/api/internal/catalog"
	"platform/products/docs/api/internal/dao"
	"platform/products/docs/api/internal/server"
)

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func bodyContains(s, needle string) bool {
	return len(needle) > 0 && strings.Contains(s, needle)
}

func docsHTTPImportZip(t *gtest.T) []byte {
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	files := map[string]string{
		"manifest.json":              `{"schemaVersion":1,"collection":"importable","version":"default","defaultLocale":"zh-CN","locales":["zh-CN"],"mode":"upsert"}`,
		"zh-CN/default/guide.md":     "---\ntitle: 导入指南\ntranslationKey: import-guide\n---\n# 导入指南\n![截图](./images/a.png)\n",
		"zh-CN/default/images/a.png": "png-bytes",
	}
	for name, body := range files {
		w, err := zw.Create(name)
		t.AssertNil(err)
		_, err = w.Write([]byte(body))
		t.AssertNil(err)
	}
	t.AssertNil(zw.Close())
	return buf.Bytes()
}

// TestHealthz boots a minimal server (nil Catalog) and asserts the enveloped
// health response: {"code":"ok","message":"ok","data":{"status":"ok"},...}
func TestHealthz(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		c := g.Client()
		c.SetPrefix(prefix(s))
		resp, err := c.Get(context.Background(), "/healthz")
		t.AssertNil(err)
		defer resp.Close()
		t.Assert(resp.StatusCode, 200)
		t.Assert(gjson.New(resp.ReadAllString()).Get("status").String(), "ok")
	})
}

// TestMe_AnonymousIsUnauthenticated calls GET /api/v1/me without a bearer token
// and expects HTTP 200 with authenticated=false. No database is
// needed — the me endpoint is catalog-independent.
func TestMe_AnonymousIsUnauthenticated(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{
			Verifier: mustVerifier(t, priv), Authorization: mustAuthorization(t, "identity-probe-admin"),
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		c := g.Client()
		c.SetPrefix(prefix(s))
		resp, err := c.Get(context.Background(), "/api/v1/me")
		t.AssertNil(err)
		defer resp.Close()

		t.Assert(resp.StatusCode, 200)
		j := gjson.New(resp.ReadAllString())
		t.Assert(j.Get("me.authenticated").Bool(), false)
		t.Assert(j.Get("me.isAdministrator").Bool(), false)
	})
}

func TestAuthorizationApplicationFlow(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		authz := mustAuthorization(t, testSub)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{
			Verifier: mustVerifier(t, priv), Authorization: authz,
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		userSub := "22222222-2222-4222-8222-222222222222"
		token := func(sub string) string {
			return signToken(t, priv, sub, time.Now().Add(time.Hour))
		}
		client := func(sub string) *gclient.Client {
			c := g.Client()
			c.SetPrefix(prefix(s))
			c.SetHeader("Authorization", "Bearer "+token(sub))
			return c
		}
		ctx := context.Background()

		requestable, err := client(userSub).Get(ctx, "/api/v1/authorization/requestable-roles")
		t.AssertNil(err)
		defer requestable.Close()
		t.Assert(requestable.StatusCode, http.StatusOK)
		t.Assert(gjson.New(requestable.ReadAllString()).Get("items.0.key").String(), "author")

		applicantClient := client(userSub)
		applicantClient.SetHeader("Idempotency-Key", "apply-author-1")
		applied, err := applicantClient.Post(ctx, "/api/v1/authorization/applications", g.Map{
			"role": "author", "reason": "I maintain this guide",
		})
		t.AssertNil(err)
		defer applied.Close()
		t.Assert(applied.StatusCode, http.StatusOK)
		applicationID := gjson.New(applied.ReadAllString()).Get("application.id").String()
		t.AssertNE(applicationID, "")
		replayed, err := applicantClient.Post(ctx, "/api/v1/authorization/applications", g.Map{
			"role": "author", "reason": "I maintain this guide",
		})
		t.AssertNil(err)
		defer replayed.Close()
		t.Assert(replayed.StatusCode, http.StatusOK)
		t.Assert(gjson.New(replayed.ReadAllString()).Get("application.id").String(), applicationID)

		reviewed, err := client(testSub).Post(
			ctx, "/api/v1/authorization/manage/applications/"+applicationID+"/review",
			g.Map{"decision": "approve", "reason": "accepted"},
		)
		t.AssertNil(err)
		defer reviewed.Close()
		t.Assert(reviewed.StatusCode, http.StatusOK)
		t.Assert(gjson.New(reviewed.ReadAllString()).Get("application.state").String(), "approved")

		me, err := client(userSub).Get(ctx, "/api/v1/me")
		t.AssertNil(err)
		defer me.Close()
		t.Assert(me.StatusCode, http.StatusOK)
		t.Assert(gjson.New(me.ReadAllString()).Get("me.roles.0").String(), "author")
	})
}

// TestDocsRoundTrip exercises the full HTTP stack: create collection, create a
// root doc, create a child doc, then verify the tree endpoint returns the
// nested structure.
func TestDocsRoundTrip(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		// Apply migration 0001 via database/sql so the schema is clean for this run.
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down))
		_, err = sdb.Exec(string(up))
		t.AssertNil(err)
		sdb.Close()

		// Boot the GoFrame docs server in-process on a random port.
		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db))

		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		exp := time.Now().UTC().Add(10 * time.Minute)
		adminJWT := signToken(t, priv, testSub, exp)

		// admin client factory — JWT-authenticated, JSON content-type.
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}
		// anon client factory — no auth header.
		anon := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			return c
		}

		// 1. Create collection "Sapphire"; slug auto-derived to "sapphire".
		rc, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Sapphire"})
		t.AssertNil(err)
		jc := gjson.New(rc.ReadAllString())
		rc.Close()
		t.Assert(rc.StatusCode, 200)
		colID := jc.Get("collection.id").String()
		t.AssertNE(colID, "")
		t.Assert(jc.Get("collection.slug").String(), "sapphire")

		// 2. Create root doc "Adjust" inside the collection.
		rd, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": colID,
			"title":        "Adjust",
			"locale":       "en",
		})
		t.AssertNil(err)
		jd := gjson.New(rd.ReadAllString())
		rd.Close()
		t.Assert(rd.StatusCode, 200)
		rootID := jd.Get("doc.id").String()
		t.AssertNE(rootID, "")
		rp, err := admin().Post(ctx, "/api/v1/docs/"+rootID+"/publish", nil)
		t.AssertNil(err)
		rp.Close()
		t.Assert(rp.StatusCode, 200)

		// 3. Create child doc "Gamma" nested under the root doc.
		rc2, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": colID,
			"parentId":     rootID,
			"title":        "Gamma",
			"locale":       "en",
		})
		t.AssertNil(err)
		jc2 := gjson.New(rc2.ReadAllString())
		rc2.Close()
		t.Assert(rc2.StatusCode, 200)
		childID := jc2.Get("doc.id").String()
		t.AssertNE(childID, "")
		cp, err := admin().Post(ctx, "/api/v1/docs/"+childID+"/publish", nil)
		t.AssertNil(err)
		cp.Close()
		t.Assert(cp.StatusCode, 200)

		// 4. GET tree (public, no auth) and assert nested structure depth ≥ 2.
		rt, err := anon().Get(ctx, "/api/v1/collections/sapphire/tree?locale=en")
		t.AssertNil(err)
		jt := gjson.New(rt.ReadAllString())
		rt.Close()
		t.Assert(rt.StatusCode, 200)
		t.Assert(jt.Get("tree.0.title").String(), "Adjust")
		t.Assert(jt.Get("tree.0.children.0.title").String(), "Gamma")
	})
}

func TestDocsImportHTTPRoundTrip(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP import integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down4, _ := os.ReadFile("../../manifest/sql/migrations/0004_doc_imports.down.sql")
		down3, _ := os.ReadFile("../../manifest/sql/migrations/0003_versions_i18n.down.sql")
		down1, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up1, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		up3, err := os.ReadFile("../../manifest/sql/migrations/0003_versions_i18n.up.sql")
		t.AssertNil(err)
		up4, err := os.ReadFile("../../manifest/sql/migrations/0004_doc_imports.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down4))
		_, _ = sdb.Exec(string(down3))
		_, _ = sdb.Exec(string(down1))
		_, err = sdb.Exec(string(up1))
		t.AssertNil(err)
		_, err = sdb.Exec(string(up3))
		t.AssertNil(err)
		_, err = sdb.Exec(string(up4))
		t.AssertNil(err)
		sdb.Close()

		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db)).WithAssets(&assetclient.Fake{}, "")

		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		adminJWT := signToken(t, priv, testSub, time.Now().UTC().Add(10*time.Minute))
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}
		anon := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			return c
		}

		rc, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Importable", "slug": "importable"})
		t.AssertNil(err)
		rc.Close()
		t.Assert(rc.StatusCode, 200)

		var body bytes.Buffer
		mw := multipart.NewWriter(&body)
		part, err := mw.CreateFormFile("file", "docs.zip")
		t.AssertNil(err)
		_, err = part.Write(docsHTTPImportZip(t))
		t.AssertNil(err)
		t.AssertNil(mw.Close())

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, base+"/api/v1/imports/docs", &body)
		t.AssertNil(err)
		req.Header.Set("Authorization", "Bearer "+adminJWT)
		req.Header.Set("Content-Type", mw.FormDataContentType())
		uploadResp, err := http.DefaultClient.Do(req)
		t.AssertNil(err)
		uploadBody, err := io.ReadAll(uploadResp.Body)
		t.AssertNil(err)
		uploadResp.Body.Close()
		t.Assert(uploadResp.StatusCode, 200)
		uploadJSON := gjson.New(uploadBody)
		batchID := uploadJSON.Get("batch.id").String()
		t.AssertNE(batchID, "")
		t.Assert(uploadJSON.Get("batch.status").String(), "checked")
		t.Assert(uploadJSON.Get("summary.creates").Int(), 1)
		t.Assert(uploadJSON.Get("summary.images").Int(), 1)

		confirmResp, err := admin().Post(ctx, "/api/v1/imports/docs/"+batchID+"/confirm", nil)
		t.AssertNil(err)
		confirmJSON := gjson.New(confirmResp.ReadAllString())
		confirmResp.Close()
		t.Assert(confirmResp.StatusCode, 200)
		t.Assert(confirmJSON.Get("batch.status").String(), "completed")

		treeResp, err := anon().Get(ctx, "/api/v1/collections/importable/tree?locale=zh-CN")
		t.AssertNil(err)
		treeBody := treeResp.ReadAllString()
		treeResp.Close()
		t.Assert(treeResp.StatusCode, 200)
		t.Assert(bodyContains(treeBody, "导入指南"), true)

		docResp, err := anon().Get(ctx, "/api/v1/docs/by-path?collection=importable&path=guide&locale=zh-CN")
		t.AssertNil(err)
		docBody := docResp.ReadAllString()
		docResp.Close()
		t.Assert(docResp.StatusCode, 200)
		t.Assert(bodyContains(docBody, "https://asset.test/docs/a.png"), true)
	})
}

func TestCollectionDocCount(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down))
		_, err = sdb.Exec(string(up))
		t.AssertNil(err)
		sdb.Close()

		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db))

		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		adminJWT := signToken(t, priv, testSub, time.Now().UTC().Add(10*time.Minute))
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}
		anon := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			return c
		}

		rc, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Counted"})
		t.AssertNil(err)
		jc := gjson.New(rc.ReadAllString())
		rc.Close()
		colID := jc.Get("collection.id").String()

		for _, title := range []string{"Draft Doc", "Published Doc", "Archived Doc"} {
			rd, err := admin().Post(ctx, "/api/v1/docs", g.Map{
				"collectionId": colID,
				"title":        title,
				"locale":       "en",
			})
			t.AssertNil(err)
			jd := gjson.New(rd.ReadAllString())
			rd.Close()
			docID := jd.Get("doc.id").String()
			if title == "Published Doc" {
				rp, err := admin().Post(ctx, "/api/v1/docs/"+docID+"/publish", nil)
				t.AssertNil(err)
				rp.Close()
			}
			if title == "Archived Doc" {
				ra, err := admin().Post(ctx, "/api/v1/docs/"+docID+"/archive", nil)
				t.AssertNil(err)
				ra.Close()
			}
		}

		listResp, err := anon().Get(ctx, "/api/v1/collections")
		t.AssertNil(err)
		listJSON := gjson.New(listResp.ReadAllString())
		listResp.Close()
		t.Assert(listResp.StatusCode, 200)
		t.Assert(listJSON.Get("items.0.docCount").Int(), 3)

		getResp, err := anon().Get(ctx, "/api/v1/collections/counted")
		t.AssertNil(err)
		getJSON := gjson.New(getResp.ReadAllString())
		getResp.Close()
		t.Assert(getResp.StatusCode, 200)
		t.Assert(getJSON.Get("collection.docCount").Int(), 3)
	})
}

func TestUpdateHomeConfig(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down2, _ := os.ReadFile("../../manifest/sql/migrations/0002_home_config.down.sql")
		down6, _ := os.ReadFile("../../manifest/sql/migrations/0006_site_settings.down.sql")
		down1, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up1, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		up2, err := os.ReadFile("../../manifest/sql/migrations/0002_home_config.up.sql")
		t.AssertNil(err)
		up6, err := os.ReadFile("../../manifest/sql/migrations/0006_site_settings.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down6))
		_, _ = sdb.Exec(string(down2))
		_, _ = sdb.Exec(string(down1))
		_, err = sdb.Exec(string(up1))
		t.AssertNil(err)
		_, err = sdb.Exec(string(up2))
		t.AssertNil(err)
		_, err = sdb.Exec(string(up6))
		t.AssertNil(err)
		sdb.Close()

		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db))

		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		adminJWT := signToken(t, priv, testSub, time.Now().UTC().Add(10*time.Minute))
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}
		anon := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			return c
		}

		save, err := admin().Patch(ctx, "/api/v1/home", g.Map{
			"quickLinks": []g.Map{{
				"id":             "quickstart",
				"title":          "Quickstart",
				"description":    "Start here",
				"icon":           "i-tabler-rocket",
				"to":             "",
				"collectionSlug": "quickstart",
				"sortOrder":      0,
				"enabled":        true,
			}},
			"featuredCollections": []string{"quickstart", "reference"},
			"homeEyebrow":         "Manual",
			"homeTitle":           "开发文档",
			"homeSubtitle":        "从任务开始阅读。",
			"siteTitle":           "Yueli Docs",
			"siteDescription":     "产品与开发文档",
			"supportEmail":        "docs@example.com",
			"footerTagline":       "可靠的产品文档",
			"footerCopyright":     "© 2026 Yueli",
		})
		t.AssertNil(err)
		saveBody := save.ReadAllString()
		saveJSON := gjson.New(saveBody)
		save.Close()
		t.Assert(save.StatusCode, 200)
		t.Assert(saveJSON.Get("config.quickLinks.0.title").String(), "Quickstart")
		t.Assert(saveJSON.Get("config.featuredCollections.1").String(), "reference")
		t.Assert(saveJSON.Get("config.siteTitle").String(), "Yueli Docs")
		t.Assert(saveJSON.Get("config.footerTagline").String(), "可靠的产品文档")

		read, err := anon().Get(ctx, "/api/v1/home")
		t.AssertNil(err)
		readJSON := gjson.New(read.ReadAllString())
		read.Close()
		t.Assert(read.StatusCode, 200)
		t.Assert(readJSON.Get("config.quickLinks.0.icon").String(), "i-tabler-rocket")
		t.Assert(readJSON.Get("config.featuredCollections.0").String(), "quickstart")
		t.Assert(readJSON.Get("config.homeTitle").String(), "开发文档")
		t.Assert(readJSON.Get("config.supportEmail").String(), "docs@example.com")
	})
}

func TestPublicTreeOnlyPublishedAndLightweight(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down))
		_, err = sdb.Exec(string(up))
		t.AssertNil(err)
		sdb.Close()

		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db))

		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		exp := time.Now().UTC().Add(10 * time.Minute)
		adminJWT := signToken(t, priv, testSub, exp)
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}
		anon := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			return c
		}

		rc, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Public Safety"})
		t.AssertNil(err)
		jc := gjson.New(rc.ReadAllString())
		rc.Close()
		colID := jc.Get("collection.id").String()

		pubResp, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": colID,
			"title":        "Visible",
			"content":      "published body must not be in tree",
			"locale":       "en",
		})
		t.AssertNil(err)
		pubJSON := gjson.New(pubResp.ReadAllString())
		pubResp.Close()
		pubID := pubJSON.Get("doc.id").String()
		patchResp, err := admin().Patch(ctx, "/api/v1/docs/"+pubID, g.Map{
			"title":   "Visible",
			"content": "published body must not be in tree",
			"status":  "published",
			"locale":  "en",
		})
		t.AssertNil(err)
		patchResp.Close()

		draftResp, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": colID,
			"title":        "Hidden Draft",
			"content":      "draft body",
			"locale":       "en",
		})
		t.AssertNil(err)
		draftResp.Close()

		treeResp, err := anon().Get(ctx, "/api/v1/collections/public-safety/tree?locale=en")
		t.AssertNil(err)
		body := treeResp.ReadAllString()
		treeResp.Close()
		j := gjson.New(body)
		t.Assert(treeResp.StatusCode, 200)
		t.Assert(j.Get("tree.0.title").String(), "Visible")
		t.Assert(j.Get("tree.0.content").String(), "")
		t.Assert(bodyContains(body, "Hidden Draft"), false)
		t.Assert(bodyContains(body, "published body must not be in tree"), false)
	})
}

func TestManageTreeIncludesDraftDocs(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down))
		_, err = sdb.Exec(string(up))
		t.AssertNil(err)
		sdb.Close()

		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db))
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		adminJWT := signToken(t, priv, testSub, time.Now().UTC().Add(10*time.Minute))
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}
		anon := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			return c
		}

		rc, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Manage Safety"})
		t.AssertNil(err)
		jc := gjson.New(rc.ReadAllString())
		rc.Close()
		colID := jc.Get("collection.id").String()
		draftResp, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": colID,
			"title":        "Visible To Managers",
			"content":      "draft manager body",
			"locale":       "en",
		})
		t.AssertNil(err)
		draftResp.Close()

		publicTree, err := anon().Get(ctx, "/api/v1/collections/manage-safety/tree?locale=en")
		t.AssertNil(err)
		publicBody := publicTree.ReadAllString()
		publicTree.Close()
		t.Assert(publicTree.StatusCode, 200)
		t.Assert(bodyContains(publicBody, "Visible To Managers"), false)

		manageTree, err := admin().Get(ctx, "/api/v1/manage/collections/manage-safety/tree?locale=en")
		t.AssertNil(err)
		manageBody := manageTree.ReadAllString()
		manageTree.Close()
		t.Assert(manageTree.StatusCode, 200)
		t.Assert(bodyContains(manageBody, "Visible To Managers"), true)
		t.Assert(bodyContains(manageBody, "draft manager body"), true)
	})
}

func TestPublicDocByPathReturnsOnlyPublishedBody(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down))
		_, err = sdb.Exec(string(up))
		t.AssertNil(err)
		sdb.Close()

		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db))
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		adminJWT := signToken(t, priv, testSub, time.Now().UTC().Add(10*time.Minute))
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}
		anon := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			return c
		}

		rc, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Body Lookup"})
		t.AssertNil(err)
		jc := gjson.New(rc.ReadAllString())
		rc.Close()
		colID := jc.Get("collection.id").String()

		rootResp, err := admin().Post(ctx, "/api/v1/docs", g.Map{"collectionId": colID, "title": "Guide", "locale": "en"})
		t.AssertNil(err)
		rootJSON := gjson.New(rootResp.ReadAllString())
		rootResp.Close()
		rootID := rootJSON.Get("doc.id").String()
		rootPatch, err := admin().Patch(ctx, "/api/v1/docs/"+rootID, g.Map{"title": "Guide", "status": "published"})
		t.AssertNil(err)
		rootPatch.Close()

		childResp, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": colID,
			"parentId":     rootID,
			"title":        "Install",
			"content":      "install body",
			"locale":       "en",
		})
		t.AssertNil(err)
		childJSON := gjson.New(childResp.ReadAllString())
		childResp.Close()
		childID := childJSON.Get("doc.id").String()
		childPatch, err := admin().Patch(ctx, "/api/v1/docs/"+childID, g.Map{
			"title":    "Install",
			"content":  "install body",
			"parentId": rootID,
			"status":   "published",
		})
		t.AssertNil(err)
		childPatch.Close()

		found, err := anon().Get(ctx, "/api/v1/docs/by-path?collection=body-lookup&path=guide/install&locale=en")
		t.AssertNil(err)
		foundJSON := gjson.New(found.ReadAllString())
		found.Close()
		t.Assert(found.StatusCode, 200)
		t.Assert(foundJSON.Get("doc.title").String(), "Install")
		t.Assert(foundJSON.Get("doc.content").String(), "install body")

		draftResp, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": colID,
			"title":        "Draft Only",
			"content":      "draft body",
			"locale":       "en",
		})
		t.AssertNil(err)
		draftResp.Close()
		missing, err := anon().Get(ctx, "/api/v1/docs/by-path?collection=body-lookup&path=draft-only&locale=en")
		t.AssertNil(err)
		missing.Close()
		t.Assert(missing.StatusCode, 404)
	})
}

func TestPatchDocIsPartialSafe(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down))
		_, err = sdb.Exec(string(up))
		t.AssertNil(err)
		sdb.Close()

		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db))
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		adminJWT := signToken(t, priv, testSub, time.Now().UTC().Add(10*time.Minute))
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}
		anon := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			return c
		}

		rc, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Patch Safety"})
		t.AssertNil(err)
		jc := gjson.New(rc.ReadAllString())
		rc.Close()
		colID := jc.Get("collection.id").String()

		rd, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": colID,
			"title":        "Keep Me",
			"content":      "body stays",
			"locale":       "en",
		})
		t.AssertNil(err)
		jd := gjson.New(rd.ReadAllString())
		rd.Close()
		id := jd.Get("doc.id").String()

		rp, err := admin().Patch(ctx, "/api/v1/docs/"+id, g.Map{"status": "published"})
		t.AssertNil(err)
		jp := gjson.New(rp.ReadAllString())
		rp.Close()
		t.Assert(rp.StatusCode, 200)
		t.Assert(jp.Get("doc.title").String(), "Keep Me")
		t.Assert(jp.Get("doc.content").String(), "body stays")
		t.Assert(jp.Get("doc.status").String(), "published")

		rs, err := admin().Patch(ctx, "/api/v1/docs/"+id, g.Map{"slug": "Custom Guide"})
		t.AssertNil(err)
		js := gjson.New(rs.ReadAllString())
		rs.Close()
		t.Assert(rs.StatusCode, 200)
		t.Assert(js.Get("doc.slug").String(), "custom-guide")
		t.Assert(js.Get("doc.title").String(), "Keep Me")
		t.Assert(js.Get("doc.content").String(), "body stays")

		oldPath, err := anon().Get(ctx, "/api/v1/docs/by-path?collection=patch-safety&path=keep-me&locale=en")
		t.AssertNil(err)
		oldPath.Close()
		t.Assert(oldPath.StatusCode, 404)
		newPath, err := anon().Get(ctx, "/api/v1/docs/by-path?collection=patch-safety&path=custom-guide&locale=en")
		t.AssertNil(err)
		newJSON := gjson.New(newPath.ReadAllString())
		newPath.Close()
		t.Assert(newPath.StatusCode, 200)
		t.Assert(newJSON.Get("doc.id").String(), id)
	})
}

func TestDocSEOFieldsArePartialSafe(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down))
		_, err = sdb.Exec(string(up))
		t.AssertNil(err)
		sdb.Close()

		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db))
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		adminJWT := signToken(t, priv, testSub, time.Now().UTC().Add(10*time.Minute))
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}

		rc, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "SEO"})
		t.AssertNil(err)
		jc := gjson.New(rc.ReadAllString())
		rc.Close()
		colID := jc.Get("collection.id").String()

		rd, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId":   colID,
			"title":          "Canonical Title",
			"content":        "body stays",
			"locale":         "en",
			"seoTitle":       "Initial SEO Title",
			"seoDescription": "Initial SEO description",
		})
		t.AssertNil(err)
		jd := gjson.New(rd.ReadAllString())
		rd.Close()
		id := jd.Get("doc.id").String()
		t.Assert(jd.Get("doc.seoTitle").String(), "Initial SEO Title")
		t.Assert(jd.Get("doc.seoDescription").String(), "Initial SEO description")

		rp, err := admin().Patch(ctx, "/api/v1/docs/"+id, g.Map{"seoTitle": "Updated SEO Title"})
		t.AssertNil(err)
		jp := gjson.New(rp.ReadAllString())
		rp.Close()
		t.Assert(rp.StatusCode, 200)
		t.Assert(jp.Get("doc.title").String(), "Canonical Title")
		t.Assert(jp.Get("doc.content").String(), "body stays")
		t.Assert(jp.Get("doc.seoTitle").String(), "Updated SEO Title")
		t.Assert(jp.Get("doc.seoDescription").String(), "Initial SEO description")
	})
}

func TestPublicSearchPublishedDocs(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down))
		_, err = sdb.Exec(string(up))
		t.AssertNil(err)
		sdb.Close()

		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db))
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		adminJWT := signToken(t, priv, testSub, time.Now().UTC().Add(10*time.Minute))
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}
		anon := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			return c
		}

		rc, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Searchable"})
		t.AssertNil(err)
		jc := gjson.New(rc.ReadAllString())
		rc.Close()
		searchableID := jc.Get("collection.id").String()
		rcOther, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Other"})
		t.AssertNil(err)
		jo := gjson.New(rcOther.ReadAllString())
		rcOther.Close()
		otherID := jo.Get("collection.id").String()

		publishedResp, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": searchableID,
			"title":        "Install Search Needle",
			"content":      "needle body",
			"locale":       "en",
		})
		t.AssertNil(err)
		publishedJSON := gjson.New(publishedResp.ReadAllString())
		publishedResp.Close()
		publishedID := publishedJSON.Get("doc.id").String()
		pub, err := admin().Post(ctx, "/api/v1/docs/"+publishedID+"/publish", nil)
		t.AssertNil(err)
		pub.Close()

		draftResp, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": searchableID,
			"title":        "Draft Search Needle",
			"content":      "needle draft body",
			"locale":       "en",
		})
		t.AssertNil(err)
		draftResp.Close()

		otherResp, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": otherID,
			"title":        "Other Search Needle",
			"content":      "needle other body",
			"locale":       "en",
		})
		t.AssertNil(err)
		otherJSON := gjson.New(otherResp.ReadAllString())
		otherResp.Close()
		otherDocID := otherJSON.Get("doc.id").String()
		otherPub, err := admin().Post(ctx, "/api/v1/docs/"+otherDocID+"/publish", nil)
		t.AssertNil(err)
		otherPub.Close()

		found, err := anon().Get(ctx, "/api/v1/docs/search?collection=searchable&locale=en&q=needle")
		t.AssertNil(err)
		body := found.ReadAllString()
		j := gjson.New(body)
		found.Close()
		t.Assert(found.StatusCode, 200)
		t.Assert(j.Get("items.0.title").String(), "Install Search Needle")
		t.Assert(j.Get("items.0.content").String(), "")
		t.Assert(bodyContains(body, "Draft Search Needle"), false)
		t.Assert(bodyContains(body, "Other Search Needle"), false)
	})
}

func TestPublicSearchRanksTitleMatchesFirst(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down))
		_, err = sdb.Exec(string(up))
		t.AssertNil(err)
		sdb.Close()

		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db))
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		adminJWT := signToken(t, priv, testSub, time.Now().UTC().Add(10*time.Minute))
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}
		anon := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			return c
		}

		rc, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Ranking"})
		t.AssertNil(err)
		jc := gjson.New(rc.ReadAllString())
		rc.Close()
		colID := jc.Get("collection.id").String()

		titleResp, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": colID,
			"title":        "Needle Quickstart",
			"content":      "plain body",
			"locale":       "en",
		})
		t.AssertNil(err)
		titleJSON := gjson.New(titleResp.ReadAllString())
		titleResp.Close()
		titleID := titleJSON.Get("doc.id").String()
		titlePub, err := admin().Post(ctx, "/api/v1/docs/"+titleID+"/publish", nil)
		t.AssertNil(err)
		titlePub.Close()

		bodyResp, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": colID,
			"title":        "Background",
			"content":      "needle appears only in the body",
			"locale":       "en",
		})
		t.AssertNil(err)
		bodyJSON := gjson.New(bodyResp.ReadAllString())
		bodyResp.Close()
		bodyID := bodyJSON.Get("doc.id").String()
		bodyPub, err := admin().Post(ctx, "/api/v1/docs/"+bodyID+"/publish", nil)
		t.AssertNil(err)
		bodyPub.Close()

		found, err := anon().Get(ctx, "/api/v1/docs/search?collection=ranking&locale=en&q=needle")
		t.AssertNil(err)
		j := gjson.New(found.ReadAllString())
		found.Close()
		t.Assert(found.StatusCode, 200)
		t.Assert(j.Get("items.0.title").String(), "Needle Quickstart")
		t.Assert(j.Get("items.1.title").String(), "Background")
	})
}

func TestPublicSearchReturnsTotalsAndCollectionFacets(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down))
		_, err = sdb.Exec(string(up))
		t.AssertNil(err)
		sdb.Close()

		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db))
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		adminJWT := signToken(t, priv, testSub, time.Now().UTC().Add(10*time.Minute))
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}
		anon := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			return c
		}

		alphaResp, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Alpha"})
		t.AssertNil(err)
		alphaJSON := gjson.New(alphaResp.ReadAllString())
		alphaResp.Close()
		alphaID := alphaJSON.Get("collection.id").String()
		betaResp, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Beta"})
		t.AssertNil(err)
		betaJSON := gjson.New(betaResp.ReadAllString())
		betaResp.Close()
		betaID := betaJSON.Get("collection.id").String()

		for _, in := range []struct {
			collectionID string
			title        string
			publish      bool
		}{
			{alphaID, "Needle Alpha", true},
			{betaID, "Needle Beta", true},
			{betaID, "Needle Draft", false},
		} {
			rd, err := admin().Post(ctx, "/api/v1/docs", g.Map{
				"collectionId": in.collectionID,
				"title":        in.title,
				"locale":       "en",
			})
			t.AssertNil(err)
			jd := gjson.New(rd.ReadAllString())
			rd.Close()
			if in.publish {
				rp, err := admin().Post(ctx, "/api/v1/docs/"+jd.Get("doc.id").String()+"/publish", nil)
				t.AssertNil(err)
				rp.Close()
			}
		}

		found, err := anon().Get(ctx, "/api/v1/docs/search?locale=en&q=needle")
		t.AssertNil(err)
		j := gjson.New(found.ReadAllString())
		found.Close()
		t.Assert(found.StatusCode, 200)
		t.Assert(j.Get("total").Int(), 2)
		t.Assert(j.Get("facets.collections.0.slug").String(), "alpha")
		t.Assert(j.Get("facets.collections.0.count").Int(), 1)
		t.Assert(j.Get("facets.collections.1.slug").String(), "beta")
		t.Assert(j.Get("facets.collections.1.count").Int(), 1)
	})
}

func TestPublicSearchRecordsQueryEvent(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down))
		_, err = sdb.Exec(string(up))
		t.AssertNil(err)
		sdb.Close()

		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db))
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		adminJWT := signToken(t, priv, testSub, time.Now().UTC().Add(10*time.Minute))
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}
		anon := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			return c
		}

		rc, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Events"})
		t.AssertNil(err)
		jc := gjson.New(rc.ReadAllString())
		rc.Close()
		colID := jc.Get("collection.id").String()
		rd, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": colID,
			"title":        "Needle Event",
			"locale":       "en",
		})
		t.AssertNil(err)
		jd := gjson.New(rd.ReadAllString())
		rd.Close()
		rp, err := admin().Post(ctx, "/api/v1/docs/"+jd.Get("doc.id").String()+"/publish", nil)
		t.AssertNil(err)
		rp.Close()

		found, err := anon().Get(ctx, "/api/v1/docs/search?collection=events&locale=en&q=needle")
		t.AssertNil(err)
		found.Close()
		t.Assert(found.StatusCode, 200)

		var event struct {
			Query          string `orm:"query"`
			CollectionSlug string `orm:"collection_slug"`
			Locale         string `orm:"locale"`
			ResultCount    int    `orm:"result_count"`
		}
		err = db.Model("doc_search_events").Ctx(ctx).
			Fields("query", "collection_slug", "locale", "result_count").
			OrderDesc("created_at").
			Limit(1).
			Scan(&event)
		t.AssertNil(err)
		t.Assert(event.Query, "needle")
		t.Assert(event.CollectionSlug, "events")
		t.Assert(event.Locale, "en")
		t.Assert(event.ResultCount, 1)

		empty, err := anon().Get(ctx, "/api/v1/docs/search?collection=events&locale=en&q=missing")
		t.AssertNil(err)
		empty.Close()
		t.Assert(empty.StatusCode, 200)
		err = db.Model("doc_search_events").Ctx(ctx).
			Fields("query", "collection_slug", "locale", "result_count").
			OrderDesc("created_at").
			Limit(1).
			Scan(&event)
		t.AssertNil(err)
		t.Assert(event.Query, "missing")
		t.Assert(event.CollectionSlug, "events")
		t.Assert(event.Locale, "en")
		t.Assert(event.ResultCount, 0)
	})
}

func TestPublishArchiveEndpointsPreserveContent(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down))
		_, err = sdb.Exec(string(up))
		t.AssertNil(err)
		sdb.Close()

		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db))
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		adminJWT := signToken(t, priv, testSub, time.Now().UTC().Add(10*time.Minute))
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}

		rc, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Lifecycle"})
		t.AssertNil(err)
		jc := gjson.New(rc.ReadAllString())
		rc.Close()
		colID := jc.Get("collection.id").String()
		rd, err := admin().Post(ctx, "/api/v1/docs", g.Map{"collectionId": colID, "title": "Lifecycle Doc", "content": "keep", "locale": "en"})
		t.AssertNil(err)
		jd := gjson.New(rd.ReadAllString())
		rd.Close()
		id := jd.Get("doc.id").String()

		pub, err := admin().Post(ctx, "/api/v1/docs/"+id+"/publish", nil)
		t.AssertNil(err)
		jp := gjson.New(pub.ReadAllString())
		pub.Close()
		t.Assert(pub.StatusCode, 200)
		t.Assert(jp.Get("doc.status").String(), "published")
		t.Assert(jp.Get("doc.content").String(), "keep")

		arc, err := admin().Post(ctx, "/api/v1/docs/"+id+"/archive", nil)
		t.AssertNil(err)
		ja := gjson.New(arc.ReadAllString())
		arc.Close()
		t.Assert(arc.StatusCode, 200)
		t.Assert(ja.Get("doc.status").String(), "archived")
		t.Assert(ja.Get("doc.content").String(), "keep")
	})
}

func TestVersionsMigrationBackfillsDefaultVersion(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down3, _ := os.ReadFile("../../manifest/sql/migrations/0003_versions_i18n.down.sql")
		down2, _ := os.ReadFile("../../manifest/sql/migrations/0002_home_config.down.sql")
		down1, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up1, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		up3, err := os.ReadFile("../../manifest/sql/migrations/0003_versions_i18n.up.sql")
		t.AssertNil(err)

		_, _ = sdb.Exec(string(down3))
		_, _ = sdb.Exec(string(down2))
		_, _ = sdb.Exec(string(down1))
		_, err = sdb.Exec(string(up1))
		t.AssertNil(err)

		_, err = sdb.Exec(`INSERT INTO collections (id, slug, title, author_sub) VALUES
			('00000000-0000-0000-0000-000000000101', 'alpha', 'Alpha', 'u1'),
			('00000000-0000-0000-0000-000000000102', 'beta', 'Beta', 'u1')`)
		t.AssertNil(err)
		_, err = sdb.Exec(`INSERT INTO docs (id, collection_id, slug, title, locale, author_sub, status) VALUES
			('00000000-0000-0000-0000-000000000201', '00000000-0000-0000-0000-000000000101', 'intro', 'Intro', 'en', 'u1', 'published')`)
		t.AssertNil(err)

		_, err = sdb.Exec(string(up3))
		t.AssertNil(err)
		defer sdb.Close()

		var versionCount int
		err = sdb.QueryRow(`SELECT COUNT(*) FROM collection_versions`).Scan(&versionCount)
		t.AssertNil(err)
		t.Assert(versionCount, 2)

		var missingVersion int
		err = sdb.QueryRow(`SELECT COUNT(*) FROM docs WHERE version_id IS NULL`).Scan(&missingVersion)
		t.AssertNil(err)
		t.Assert(missingVersion, 0)

		var key string
		err = sdb.QueryRow(`SELECT translation_key FROM docs WHERE slug='intro'`).Scan(&key)
		t.AssertNil(err)
		t.AssertNE(key, "")

		_, err = sdb.Exec(string(down3))
		t.AssertNil(err)
	})
}

func TestPublicTreeUsesRequestedVersion(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		host := os.Getenv("DOCS_PG_HOST")
		if host == "" {
			t.Skip("set DOCS_PG_HOST to run the docs HTTP integration test")
		}
		ctx := context.Background()
		port := envOr("DOCS_PG_PORT", "5432")
		user := envOr("DOCS_PG_USER", "postgres")
		pass := os.Getenv("DOCS_PG_PASS")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=docs sslmode=disable", host, port, user, pass)
		sdb, err := sql.Open("postgres", dsn)
		t.AssertNil(err)
		down3, _ := os.ReadFile("../../manifest/sql/migrations/0003_versions_i18n.down.sql")
		down1, _ := os.ReadFile("../../manifest/sql/migrations/0001_init.down.sql")
		up1, err := os.ReadFile("../../manifest/sql/migrations/0001_init.up.sql")
		t.AssertNil(err)
		up3, err := os.ReadFile("../../manifest/sql/migrations/0003_versions_i18n.up.sql")
		t.AssertNil(err)
		_, _ = sdb.Exec(string(down3))
		_, _ = sdb.Exec(string(down1))
		_, err = sdb.Exec(string(up1))
		t.AssertNil(err)
		_, err = sdb.Exec(string(up3))
		t.AssertNil(err)
		sdb.Close()

		db, err := gdb.New(gdb.ConfigNode{Type: "pgsql", Host: host, Port: port, User: user, Pass: pass, Name: "docs"})
		t.AssertNil(err)
		cat := catalog.New(dao.NewPG(db))
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		t.AssertNil(err)
		s := g.Server(t.Name())
		s.SetAddr("127.0.0.1:0")
		server.Configure(s, server.Deps{Verifier: mustVerifier(t, priv), Catalog: cat, Authorization: mustAuthorization(t, testSub)})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		base := prefix(s)
		adminJWT := signToken(t, priv, testSub, time.Now().UTC().Add(10*time.Minute))
		admin := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			c.ContentJson()
			c.SetHeader("Authorization", "Bearer "+adminJWT)
			return c
		}
		anon := func() *gclient.Client {
			c := g.Client()
			c.SetPrefix(base)
			return c
		}

		rc, err := admin().Post(ctx, "/api/v1/collections", g.Map{"title": "Versioned"})
		t.AssertNil(err)
		colID := gjson.New(rc.ReadAllString()).Get("collection.id").String()
		rc.Close()

		v2, err := admin().Post(ctx, "/api/v1/collections/"+colID+"/versions", g.Map{
			"key": "v2", "label": "Version 2", "status": "published", "sourceVersionId": "",
		})
		t.AssertNil(err)
		v2ID := gjson.New(v2.ReadAllString()).Get("version.id").String()
		v2.Close()
		t.AssertNE(v2ID, "")

		defaultDoc, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": colID, "title": "Default Intro", "locale": "en",
		})
		t.AssertNil(err)
		defaultDocID := gjson.New(defaultDoc.ReadAllString()).Get("doc.id").String()
		defaultDoc.Close()
		pubDefault, err := admin().Post(ctx, "/api/v1/docs/"+defaultDocID+"/publish", nil)
		t.AssertNil(err)
		pubDefault.Close()

		v2Doc, err := admin().Post(ctx, "/api/v1/docs", g.Map{
			"collectionId": colID, "versionId": v2ID, "title": "V2 Intro", "locale": "en",
		})
		t.AssertNil(err)
		v2DocID := gjson.New(v2Doc.ReadAllString()).Get("doc.id").String()
		v2Doc.Close()
		pubV2, err := admin().Post(ctx, "/api/v1/docs/"+v2DocID+"/publish", nil)
		t.AssertNil(err)
		pubV2.Close()

		defaultTree, err := anon().Get(ctx, "/api/v1/collections/versioned/tree?locale=en")
		t.AssertNil(err)
		defaultBody := defaultTree.ReadAllString()
		defaultTree.Close()
		t.Assert(bodyContains(defaultBody, "Default Intro"), true)
		t.Assert(bodyContains(defaultBody, "V2 Intro"), false)

		v2Tree, err := anon().Get(ctx, "/api/v1/collections/versioned/tree?locale=en&version=v2")
		t.AssertNil(err)
		v2Body := v2Tree.ReadAllString()
		v2Tree.Close()
		t.Assert(bodyContains(v2Body, "V2 Intro"), true)
		t.Assert(bodyContains(v2Body, "Default Intro"), false)
	})
}
