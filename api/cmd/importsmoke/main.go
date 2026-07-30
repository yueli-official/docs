// Command importsmoke is a one-off real-data smoke (P1 Task 7): it imports a
// single Astro docs collection (yozya en/sapphire) into the docs DB to validate
// the collection + multi-level tree modeling against real content — tree depth,
// slug collisions, CJK/encoding. It is NOT the full P4 importer (that gets its
// own plan); delete or fold this in once P4 lands.
//
// Run (connects to the shared dev PG; host defaults to 192.168.5.5):
//
//	go run ./cmd/importsmoke
//	SMOKE_DIR=... DOCS_PG_HOST=... go run ./cmd/importsmoke
package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/yueli-official/docs/api/internal/catalog"
	"github.com/yueli-official/docs/api/internal/dao"
)

const (
	defaultDir   = `E:/projects/yozya/docs/src/content/docs/en/sapphire`
	collTitle    = "Sapphire"
	collSlug     = "sapphire"
	importerSub  = "importsmoke"
	importLocale = "en"
)

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func openDB() (gdb.DB, error) {
	return gdb.New(gdb.ConfigNode{
		Type: "pgsql",
		Host: envOr("DOCS_PG_HOST", "192.168.5.5"),
		Port: envOr("DOCS_PG_PORT", "5432"),
		User: envOr("DOCS_PG_USER", "postgres"),
		Pass: envOr("DOCS_PG_PASS", "postgres"),
		Name: envOr("DOCS_PG_NAME", "docs"),
	})
}

// stats tracks what the import produced so the final summary can flag modeling
// surprises (collisions, depth) instead of silently swallowing them.
type stats struct {
	categories int
	leaves     int
	collisions []string // "category/file: <slug>" that hit SlugTaken
	maxTitle   string   // longest title seen (CJK/encoding eyeball)
}

func main() {
	dir := envOr("SMOKE_DIR", defaultDir)
	ctx := gctx.New()

	db, err := openDB()
	if err != nil {
		fatalf("open db: %v", err)
	}
	cat := catalog.New(dao.NewPG(db))

	// Idempotent: drop a prior sapphire collection (cascade removes its docs).
	if existing, err := cat.GetCollectionBySlug(ctx, collSlug); err == nil && existing != nil {
		if err := cat.DeleteCollection(ctx, existing.ID); err != nil {
			fatalf("delete prior collection: %v", err)
		}
		fmt.Printf("dropped prior collection %q (id=%s)\n", collSlug, existing.ID)
	}

	col, err := cat.CreateCollection(ctx, importerSub, collTitle, "", "Sapphire plugin docs (en), smoke import", "", "")
	if err != nil {
		fatalf("create collection: %v", err)
	}
	fmt.Printf("collection: %s (id=%s)\n", col.Slug, col.ID)

	var st stats

	// Each immediate subdir of `dir` (except _static) becomes a category node at
	// the root; its .md files become leaves under it. A two-level tree, which is
	// exactly the parent_id nesting we want to exercise.
	entries, err := os.ReadDir(dir)
	if err != nil {
		fatalf("read dir %s: %v", dir, err)
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })

	for _, e := range entries {
		name := e.Name()
		if e.IsDir() {
			if name == "_static" {
				continue
			}
			importCategory(ctx, cat, col.ID, filepath.Join(dir, name), name, &st)
			continue
		}
		// A stray top-level .md → leaf at the root (parent="").
		if strings.HasSuffix(name, ".md") {
			importFile(ctx, cat, col.ID, "", filepath.Join(dir, name), &st)
		}
	}

	fmt.Printf("\n── summary ──\n")
	fmt.Printf("categories: %d\n", st.categories)
	fmt.Printf("leaf docs:  %d\n", st.leaves)
	fmt.Printf("total docs: %d\n", st.categories+st.leaves)
	if len(st.collisions) > 0 {
		fmt.Printf("⚠ slug collisions (%d):\n", len(st.collisions))
		for _, c := range st.collisions {
			fmt.Printf("  - %s\n", c)
		}
	} else {
		fmt.Printf("slug collisions: none\n")
	}
	fmt.Printf("longest title seen: %q\n", st.maxTitle)
}

// importCategory creates the category parent node, then imports every .md under
// that directory as a child leaf.
func importCategory(ctx context.Context, cat *catalog.Service, colID, path, name string, st *stats) {
	node, err := cat.CreateDoc(ctx, importerSub, catalog.CreateDocInput{
		CollectionID: colID,
		ParentID:     "",
		Title:        name,
		Content:      "",
		Locale:       importLocale,
	})
	if err != nil {
		st.collisions = append(st.collisions, fmt.Sprintf("[category] %s: %v", name, err))
		return
	}
	st.categories++

	files, err := os.ReadDir(path)
	if err != nil {
		fatalf("read category dir %s: %v", path, err)
	}
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })
	for _, f := range files {
		if f.IsDir() || !strings.HasSuffix(f.Name(), ".md") {
			continue
		}
		importFile(ctx, cat, colID, node.ID, filepath.Join(path, f.Name()), st)
	}
}

// importFile parses a single Astro .md (frontmatter title + body) and inserts it
// as a doc under parentID.
func importFile(ctx context.Context, cat *catalog.Service, colID, parentID, path string, st *stats) {
	raw, err := os.ReadFile(path)
	if err != nil {
		fatalf("read file %s: %v", path, err)
	}
	title, body := parseFrontmatter(string(raw))
	if title == "" {
		// Fall back to the filename (minus .md) so a missing/odd frontmatter
		// still yields a non-empty slug.
		title = strings.TrimSuffix(filepath.Base(path), ".md")
	}
	if len([]rune(title)) > len([]rune(st.maxTitle)) {
		st.maxTitle = title
	}
	_, err = cat.CreateDoc(ctx, importerSub, catalog.CreateDocInput{
		CollectionID: colID,
		ParentID:     parentID,
		Title:        title,
		Content:      body,
		Locale:       importLocale,
	})
	if err != nil {
		st.collisions = append(st.collisions, fmt.Sprintf("%s -> title %q: %v", filepath.Base(path), title, err))
		return
	}
	st.leaves++
}

// parseFrontmatter splits a leading "---\n...\n---\n" YAML block from the body and
// pulls the `title:` value. Minimal by design (Astro frontmatter here is just a
// title line); a full YAML parse is P4's job.
func parseFrontmatter(s string) (title, body string) {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	if !strings.HasPrefix(s, "---\n") {
		return "", s
	}
	rest := s[len("---\n"):]
	end := strings.Index(rest, "\n---")
	if end < 0 {
		return "", s
	}
	fm := rest[:end]
	body = strings.TrimPrefix(rest[end+len("\n---"):], "\n")
	for _, line := range strings.Split(fm, "\n") {
		if v, ok := strings.CutPrefix(strings.TrimSpace(line), "title:"); ok {
			title = strings.Trim(strings.TrimSpace(v), `"'`)
			break
		}
	}
	return title, body
}

func fatalf(format string, a ...any) {
	fmt.Fprintf(os.Stderr, "importsmoke: "+format+"\n", a...)
	os.Exit(1)
}
