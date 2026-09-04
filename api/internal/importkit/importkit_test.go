package importkit

import (
	"archive/zip"
	"bytes"
	"errors"
	"hash/crc32"
	"strings"
	"testing"

	"github.com/ulikunitz/xz"
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

func xzZipBytes(t *testing.T, name, body string) []byte {
	t.Helper()
	var compressed bytes.Buffer
	xzWriter, err := xz.NewWriter(&compressed)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := xzWriter.Write([]byte(body)); err != nil {
		t.Fatal(err)
	}
	if err := xzWriter.Close(); err != nil {
		t.Fatal(err)
	}

	var archive bytes.Buffer
	zipWriter := zip.NewWriter(&archive)
	header := &zip.FileHeader{Name: name, Method: zipMethodXZ}
	header.CRC32 = crc32.ChecksumIEEE([]byte(body))
	header.CompressedSize64 = uint64(compressed.Len())
	header.UncompressedSize64 = uint64(len(body))
	entry, err := zipWriter.CreateRaw(header)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := entry.Write(compressed.Bytes()); err != nil {
		t.Fatal(err)
	}
	if err := zipWriter.Close(); err != nil {
		t.Fatal(err)
	}
	return archive.Bytes()
}

func TestParseZipRejectsUnknownCompressionWithTypedError(t *testing.T) {
	data := xzZipBytes(t, "index.md", "# Guide")
	// The method appears in both the local and central file headers.
	for index := 0; index+12 < len(data); index++ {
		if bytes.Equal(data[index:index+4], []byte{'P', 'K', 3, 4}) {
			data[index+8] = 99
			data[index+9] = 0
		}
		if bytes.Equal(data[index:index+4], []byte{'P', 'K', 1, 2}) {
			data[index+10] = 99
			data[index+11] = 0
		}
	}
	_, err := ParseZip(data, Options{Collection: "guide", DefaultLocale: "en", Mode: "upsert"})
	var unsupported *UnsupportedCompressionError
	if !errors.As(err, &unsupported) || unsupported.Method != 99 {
		t.Fatalf("error = %#v", err)
	}
}

func TestParseZipSupportsXZCompressedEntries(t *testing.T) {
	data := xzZipBytes(t, "index.md", "---\nid: home\ntitle: 首页\n---\n正文")
	pkg, err := ParseZip(data, Options{Collection: "guide", DefaultLocale: "zh-CN", Mode: "upsert"})
	if err != nil {
		t.Fatalf("ParseZip: %v", err)
	}
	if len(pkg.Docs) != 1 || pkg.Docs[0].Title != "首页" || pkg.Docs[0].Content != "正文" {
		t.Fatalf("package: %+v", pkg)
	}
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

func TestParseZipEnforcesArchiveResourceBudgets(t *testing.T) {
	data := zipBytes(t, map[string]string{"a.md": "12345", "b.md": "67890"})
	base := Options{Collection: "docs", DefaultLocale: "zh-CN"}

	archive := base
	archive.MaxArchiveBytes = int64(len(data) - 1)
	if _, err := ParseZip(data, archive); err == nil || !strings.Contains(err.Error(), "archive exceeds") {
		t.Fatalf("archive limit error = %v", err)
	}

	entry := base
	entry.MaxEntryBytes = 4
	if _, err := ParseZip(data, entry); err == nil || !strings.Contains(err.Error(), "entry exceeds") {
		t.Fatalf("entry limit error = %v", err)
	}

	extracted := base
	extracted.MaxExtractedBytes = 9
	if _, err := ParseZip(data, extracted); err == nil || !strings.Contains(err.Error(), "expands beyond") {
		t.Fatalf("extracted limit error = %v", err)
	}
	extracted.MaxExtractedBytes = 4
	if _, err := ParseZip(data, extracted); err == nil || !strings.Contains(err.Error(), "expands beyond") {
		t.Fatalf("single entry extracted limit error = %v", err)
	}

	entries := base
	entries.MaxEntries = 1
	if _, err := ParseZip(data, entries); err == nil || !strings.Contains(err.Error(), "file entries") {
		t.Fatalf("entry count error = %v", err)
	}
}
