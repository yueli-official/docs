package importkit

import (
	"path"
	"strings"
)

func cleanZipPath(p string) string {
	return strings.Trim(path.Clean(strings.ReplaceAll(p, "\\", "/")), "/")
}

func safeJoin(baseFile, ref string) (string, bool) {
	if ref == "" || strings.Contains(ref, "://") || strings.HasPrefix(ref, "file:") || strings.HasPrefix(ref, "/") {
		return "", false
	}
	p := cleanZipPath(path.Join(path.Dir(baseFile), ref))
	if p == "." || strings.HasPrefix(p, "../") || strings.Contains(p, "/../") {
		return "", false
	}
	return p, true
}

func docPathFromMarkdown(locale, version, source string) string {
	p := cleanZipPath(source)
	p = strings.TrimPrefix(p, cleanZipPath(locale+"/"+version)+"/")
	p = strings.TrimSuffix(p, ".md")
	p = strings.TrimSuffix(p, ".mdx")
	if path.Base(p) == "index" {
		p = path.Dir(p)
		if p == "." {
			return ""
		}
		return p
	}
	return p
}

func slugFromPath(docPath string) string {
	base := path.Base(cleanZipPath(docPath))
	if base == "." || base == "/" {
		return "index"
	}
	return base
}
