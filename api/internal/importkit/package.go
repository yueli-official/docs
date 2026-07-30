package importkit

import (
	"fmt"
	"path"
	"sort"
	"strings"
)

func ParseZip(data []byte, opts Options) (*Package, error) {
	files, err := readZipFiles(data)
	if err != nil {
		return nil, err
	}
	manifestRaw, ok := files["manifest.json"]
	if !ok {
		return nil, fmt.Errorf("manifest.json is required")
	}
	manifest, err := parseManifest(manifestRaw)
	if err != nil {
		return nil, err
	}
	pkg := &Package{
		Manifest: manifest,
		Assets:   map[string]AssetFile{},
	}
	for _, key := range sortedKeys(files) {
		if isImagePath(key) {
			pkg.Assets[key] = assetFromFile(key, files[key])
		}
	}
	docPathSet := map[string]bool{}
	for _, key := range sortedKeys(files) {
		if !isMarkdownPath(key) {
			continue
		}
		locale, version, ok := localeVersionForPath(manifest, key)
		if !ok {
			continue
		}
		doc, issues, err := parseDocFile(manifest, locale, version, key, string(files[key]))
		if err != nil {
			pkg.Issues = append(pkg.Issues, Issue{Severity: "error", Code: "frontmatter_invalid", Message: err.Error(), Path: key})
			continue
		}
		pkg.Docs = append(pkg.Docs, doc)
		docPathSet[doc.Locale+"/"+doc.VersionKey+"/"+doc.Path] = true
		pkg.Issues = append(pkg.Issues, issues...)
	}
	validateDocCollisions(pkg)
	validateImageRefs(pkg, opts)
	validateDocLinks(pkg, docPathSet)
	sort.SliceStable(pkg.Docs, func(i, j int) bool {
		if pkg.Docs[i].Locale != pkg.Docs[j].Locale {
			return pkg.Docs[i].Locale < pkg.Docs[j].Locale
		}
		if pkg.Docs[i].Order != pkg.Docs[j].Order {
			return pkg.Docs[i].Order < pkg.Docs[j].Order
		}
		return pkg.Docs[i].Path < pkg.Docs[j].Path
	})
	return pkg, nil
}

func localeVersionForPath(manifest Manifest, source string) (string, string, bool) {
	for _, locale := range manifest.Locales {
		prefix := cleanZipPath(locale + "/" + manifest.Version + "/")
		if strings.HasPrefix(cleanZipPath(source), prefix) {
			return locale, manifest.Version, true
		}
	}
	return "", "", false
}

func parseDocFile(manifest Manifest, locale, version, source, raw string) (DocFile, []Issue, error) {
	fm, body, err := splitFrontmatter(raw)
	if err != nil {
		return DocFile{}, nil, err
	}
	docPath := docPathFromMarkdown(locale, version, source)
	slug := fm.Slug
	if slug == "" {
		slug = slugFromPath(docPath)
	}
	title := fm.Title
	if title == "" {
		title = slug
	}
	doc := DocFile{
		Locale:         locale,
		VersionKey:     version,
		SourcePath:     source,
		Path:           docPath,
		Slug:           slug,
		Title:          title,
		Content:        body,
		RawContent:     raw,
		Excerpt:        fm.Excerpt,
		TranslationKey: fm.TranslationKey,
		Order:          fm.Order,
		ImageRefs:      imageRefs(source, body),
		Links:          docLinks(source, locale, version, body),
	}
	var issues []Issue
	if doc.Path == "" && path.Base(source) != "index.md" && path.Base(source) != "index.mdx" {
		issues = append(issues, Issue{Severity: "error", Code: "doc_path_empty", Message: "document path is empty", Path: source})
	}
	return doc, issues, nil
}

func validateDocCollisions(pkg *Package) {
	seen := map[string]string{}
	for _, doc := range pkg.Docs {
		key := doc.Locale + "/" + doc.VersionKey + "/" + doc.Path
		if other := seen[key]; other != "" {
			pkg.Issues = append(pkg.Issues, Issue{
				Severity: "error",
				Code:     "duplicate_doc_path",
				Message:  "duplicate document path also used by " + other,
				Path:     doc.SourcePath,
			})
			continue
		}
		seen[key] = doc.SourcePath
	}
}

func validateImageRefs(pkg *Package, opts Options) {
	for _, doc := range pkg.Docs {
		for _, ref := range doc.ImageRefs {
			if ref.ResolvedPath == "" {
				pkg.Issues = append(pkg.Issues, Issue{Severity: "error", Code: "image_path_unsafe", Message: "image path is unsafe or unsupported", Path: doc.SourcePath})
				continue
			}
			asset, ok := pkg.Assets[ref.ResolvedPath]
			if !ok {
				pkg.Issues = append(pkg.Issues, Issue{Severity: "error", Code: "image_missing", Message: "image file not found: " + ref.Original, Path: doc.SourcePath})
				continue
			}
			if !isImagePath(asset.SourcePath) {
				pkg.Issues = append(pkg.Issues, Issue{Severity: "error", Code: "image_type_invalid", Message: "image type is not allowed", Path: ref.ResolvedPath})
			}
			if opts.MaxImageBytes > 0 && asset.Size > opts.MaxImageBytes {
				pkg.Issues = append(pkg.Issues, Issue{Severity: "error", Code: "image_too_large", Message: "image exceeds size limit", Path: ref.ResolvedPath})
			}
		}
	}
}

func validateDocLinks(pkg *Package, docPathSet map[string]bool) {
	for _, doc := range pkg.Docs {
		for _, link := range doc.Links {
			if link.TargetPath == "" {
				pkg.Issues = append(pkg.Issues, Issue{Severity: "error", Code: "link_path_unsafe", Message: "internal link path is unsafe", Path: doc.SourcePath})
				continue
			}
			key := doc.Locale + "/" + doc.VersionKey + "/" + path.Clean(link.TargetPath)
			if !docPathSet[key] {
				pkg.Issues = append(pkg.Issues, Issue{Severity: "error", Code: "link_target_missing", Message: "internal link target not found: " + link.Original, Path: doc.SourcePath})
			}
		}
	}
}
