package importkit

import (
	"archive/zip"
	"bytes"
	"testing"
)

func zipBytes(t *testing.T, files map[string]string) []byte {
	t.Helper()
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	for name, body := range files {
		w, err := zw.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := w.Write([]byte(body)); err != nil {
			t.Fatal(err)
		}
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func TestParseMarkdownDirectoryWithoutManifest(t *testing.T) {
	data := zipBytes(t, map[string]string{
		"index.md":         "---\nid: home\ntitle: 首页\ndescription: 入口\norder: 1\n---\n[安装](./guide/install.md)",
		"guide/install.md": "---\nid: install\ntitle: 安装\ndraft: true\n---\n正文",
	})
	pkg, err := ParseZip(data, Options{Collection: "quickstart", DefaultLocale: "zh-CN", Mode: "upsert"})
	if err != nil {
		t.Fatalf("ParseZip: %v", err)
	}
	if len(pkg.Issues) != 0 || len(pkg.Docs) != 2 {
		t.Fatalf("package: docs=%d issues=%+v", len(pkg.Docs), pkg.Issues)
	}
	var home, install DocFile
	for _, doc := range pkg.Docs {
		if doc.Path == "" {
			home = doc
		}
		if doc.Path == "guide/install" {
			install = doc
		}
	}
	if home.Excerpt != "入口" || home.TranslationKey != "home" || home.Order != 1 {
		t.Fatalf("front matter not mapped: %+v", home)
	}
	if !install.Draft {
		t.Fatalf("draft not mapped: %+v", install)
	}
}

func TestParseDocsJSONLocalesAndNavigation(t *testing.T) {
	data := zipBytes(t, map[string]string{
		"docs.json":                    `{"schemaVersion":1,"defaultLocale":"zh-CN","locales":{"zh-CN":".","en-US":"locales/en-US"},"navigation":[{"group":"指南","translations":{"en-US":"Guide"},"children":[{"page":"guide/start.md"}]}]}`,
		"guide/start.md":               "---\nid: guide-start\ntitle: 指南\n---\n正文",
		"locales/en-US/guide/start.md": "---\nid: guide-start\ntitle: Guide\n---\nBody",
	})
	pkg, err := ParseZip(data, Options{Collection: "quickstart"})
	if err != nil {
		t.Fatalf("ParseZip: %v", err)
	}
	if len(pkg.Issues) != 0 || len(pkg.Docs) != 4 {
		t.Fatalf("package: docs=%d issues=%+v", len(pkg.Docs), pkg.Issues)
	}
	groups := 0
	for _, doc := range pkg.Docs {
		if doc.Path == "guide" && doc.TranslationKey == "section:guide" {
			groups++
		}
	}
	if groups != 2 {
		t.Fatalf("localized navigation groups = %d, docs=%+v", groups, pkg.Docs)
	}
}

func TestParseDocsJSONRejectsMissingNavigationPage(t *testing.T) {
	data := zipBytes(t, map[string]string{
		"docs.json": `{"schemaVersion":1,"defaultLocale":"zh-CN","locales":{"zh-CN":"."},"navigation":[{"page":"missing.md"}]}`,
		"guide.md":  "# 指南",
	})
	pkg, err := ParseZip(data, Options{Collection: "quickstart"})
	if err != nil {
		t.Fatalf("ParseZip: %v", err)
	}
	if len(pkg.Issues) != 1 || pkg.Issues[0].Code != "navigation_page_missing" {
		t.Fatalf("issues: %+v", pkg.Issues)
	}
}

func TestParsePackageRejectsUnsafeImagePath(t *testing.T) {
	data := zipBytes(t, map[string]string{"install.md": "![bad](../../outside.png)"})
	pkg, err := ParseZip(data, Options{Collection: "quickstart", DefaultLocale: "zh-CN"})
	if err != nil {
		t.Fatalf("ParseZip: %v", err)
	}
	if len(pkg.Issues) == 0 || pkg.Issues[0].Code != "image_path_unsafe" {
		t.Fatalf("issues: %+v", pkg.Issues)
	}
}

func TestParsePackageRequiresUploadTarget(t *testing.T) {
	data := zipBytes(t, map[string]string{"index.md": "# 首页"})
	if _, err := ParseZip(data, Options{DefaultLocale: "zh-CN"}); err == nil {
		t.Fatal("expected target collection error")
	}
}

func TestParseImageDestinationIgnoresOptionalMarkdownTitle(t *testing.T) {
	data := zipBytes(t, map[string]string{
		"introduction/objectmodel.md": `![对象模型](../_static/objectmodel.png "After Effects 对象模型")`,
		"_static/objectmodel.png":     "png",
	})
	pkg, err := ParseZip(data, Options{Collection: "ae", DefaultLocale: "zh-CN"})
	if err != nil {
		t.Fatalf("ParseZip: %v", err)
	}
	if len(pkg.Issues) != 0 {
		t.Fatalf("issues: %+v", pkg.Issues)
	}
	if got := pkg.Docs[0].ImageRefs[0].ResolvedPath; got != "_static/objectmodel.png" {
		t.Fatalf("resolved image path = %q", got)
	}
}
