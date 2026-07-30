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

func TestParsePackageHandlesIndexOrderImagesAndLinks(t *testing.T) {
	data := zipBytes(t, map[string]string{
		"manifest.json":                  `{"schemaVersion":1,"collection":"quickstart","version":"v1","defaultLocale":"zh-CN","locales":["zh-CN"],"mode":"upsert"}`,
		"zh-CN/v1/guide/index.md":        "---\ntitle: 指南\norder: 2\ntranslationKey: guide\n---\n# 指南\n![图](./images/a.png)\n[安装](./install.md#step)\n",
		"zh-CN/v1/guide/install.md":      "---\ntitle: 安装\nslug: install\norder: 3\n---\n正文",
		"zh-CN/v1/guide/images/a.png":    "png-bytes",
		"zh-CN/v1/guide/images/extra.md": "not a doc because it is under images",
	})
	pkg, err := ParseZip(data, Options{MaxImageBytes: 10 << 20})
	if err != nil {
		t.Fatalf("ParseZip: %v", err)
	}
	if len(pkg.Issues) != 0 {
		t.Fatalf("issues: %+v", pkg.Issues)
	}
	if pkg.Manifest.Collection != "quickstart" || pkg.Manifest.Version != "v1" {
		t.Fatalf("manifest not parsed: %+v", pkg.Manifest)
	}
	if len(pkg.Docs) != 2 {
		t.Fatalf("docs=%d", len(pkg.Docs))
	}
	if pkg.Docs[0].Path != "guide" || pkg.Docs[0].Order != 2 {
		t.Fatalf("index mapping failed: %+v", pkg.Docs[0])
	}
	if len(pkg.Docs[0].ImageRefs) != 1 || pkg.Docs[0].ImageRefs[0].ResolvedPath != "zh-CN/v1/guide/images/a.png" {
		t.Fatalf("image refs: %+v", pkg.Docs[0].ImageRefs)
	}
	if len(pkg.Docs[0].Links) != 1 || pkg.Docs[0].Links[0].Anchor != "step" || pkg.Docs[0].Links[0].TargetPath != "guide/install" {
		t.Fatalf("links: %+v", pkg.Docs[0].Links)
	}
}

func TestParsePackageRejectsUnsafeImagePath(t *testing.T) {
	data := zipBytes(t, map[string]string{
		"manifest.json":       `{"schemaVersion":1,"collection":"quickstart","version":"v1","defaultLocale":"zh-CN","locales":["zh-CN"],"mode":"upsert"}`,
		"zh-CN/v1/install.md": "---\ntitle: 安装\n---\n![bad](../../../../outside.png)",
	})
	pkg, err := ParseZip(data, Options{MaxImageBytes: 10 << 20})
	if err != nil {
		t.Fatalf("ParseZip returned fatal error, wanted package issue: %v", err)
	}
	if len(pkg.Issues) == 0 {
		t.Fatal("expected unsafe image issue")
	}
}

func TestParsePackageRejectsMissingManifestVersion(t *testing.T) {
	data := zipBytes(t, map[string]string{
		"manifest.json":       `{"collection":"quickstart","version":"v1","defaultLocale":"zh-CN","locales":["zh-CN"],"mode":"upsert"}`,
		"zh-CN/v1/install.md": "---\ntitle: 安装\n---\n正文",
	})
	_, err := ParseZip(data, Options{MaxImageBytes: 10 << 20})
	if err == nil {
		t.Fatal("expected schemaVersion error")
	}
}

func TestParsePackageReportsMissingInternalLinkTarget(t *testing.T) {
	data := zipBytes(t, map[string]string{
		"manifest.json":       `{"schemaVersion":1,"collection":"quickstart","version":"v1","defaultLocale":"zh-CN","locales":["zh-CN"],"mode":"upsert"}`,
		"zh-CN/v1/install.md": "---\ntitle: 安装\n---\n[缺失](./missing.md#part)",
	})
	pkg, err := ParseZip(data, Options{MaxImageBytes: 10 << 20})
	if err != nil {
		t.Fatalf("ParseZip: %v", err)
	}
	if len(pkg.Issues) == 0 {
		t.Fatal("expected missing link issue")
	}
	if pkg.Docs[0].Links[0].Anchor != "part" {
		t.Fatalf("anchor not preserved: %+v", pkg.Docs[0].Links[0])
	}
}
