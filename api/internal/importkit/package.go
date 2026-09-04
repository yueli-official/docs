package importkit

import (
	"fmt"
	"path"
	"sort"
	"strings"
)

func ParseZip(data []byte, opts Options) (*Package, error) {
	opts = withDefaultLimits(opts)
	if int64(len(data)) > opts.MaxArchiveBytes {
		return nil, fmt.Errorf("archive exceeds %d bytes", opts.MaxArchiveBytes)
	}
	files, err := readZipFiles(data, opts)
	if err != nil {
		return nil, err
	}
	var manifest Manifest
	if manifestRaw, ok := files["docs.json"]; ok {
		manifest, err = parseDocsManifest(manifestRaw, opts)
	} else {
		manifest, err = manifestFromOptions(opts, "", nil)
	}
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
		locale, version, root, ok := docLocationForPath(manifest, key)
		if !ok {
			continue
		}
		doc, issues, err := parseDocFile(manifest, locale, version, root, key, string(files[key]))
		if err != nil {
			pkg.Issues = append(pkg.Issues, Issue{Severity: "error", Code: "frontmatter_invalid", Message: "front matter is invalid", Path: key})
			continue
		}
		pkg.Docs = append(pkg.Docs, doc)
		docPathSet[doc.Locale+"/"+doc.VersionKey+"/"+doc.Path] = true
		pkg.Issues = append(pkg.Issues, issues...)
	}
	addNavigationGroups(pkg)
	docPathSet = map[string]bool{}
	for _, doc := range pkg.Docs {
		docPathSet[doc.Locale+"/"+doc.VersionKey+"/"+doc.Path] = true
	}
	validateDocCollisions(pkg)
	validateImageRefs(pkg, opts)
	validateDocLinks(pkg, docPathSet)
	validateNavigation(pkg, docPathSet)
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

func withDefaultLimits(opts Options) Options {
	if opts.MaxArchiveBytes <= 0 {
		opts.MaxArchiveBytes = 100 * 1024 * 1024
	}
	if opts.MaxExtractedBytes <= 0 {
		opts.MaxExtractedBytes = 1024 * 1024 * 1024
	}
	if opts.MaxEntryBytes <= 0 {
		opts.MaxEntryBytes = 25 * 1024 * 1024
	}
	if opts.MaxEntries <= 0 {
		opts.MaxEntries = 20_000
	}
	return opts
}

func addNavigationGroups(pkg *Package) {
	existing := map[string]bool{}
	for _, doc := range pkg.Docs {
		existing[doc.Locale+"/"+doc.Path] = true
	}
	var walk func([]NavigationNode)
	walk = func(nodes []NavigationNode) {
		for index, node := range nodes {
			if node.Group == "" {
				continue
			}
			dir := navigationCommonDir(node.Children)
			if dir == "" {
				walk(node.Children)
				continue
			}
			for _, locale := range pkg.Manifest.Locales {
				key := locale + "/" + dir
				if existing[key] {
					continue
				}
				title := node.Group
				if translated := node.Translations[locale]; translated != "" {
					title = translated
				}
				pkg.Docs = append(pkg.Docs, DocFile{
					Locale: locale, SourcePath: "docs.json#navigation/" + locale + "/" + dir,
					Path: dir, Slug: path.Base(dir), Title: title,
					TranslationKey: "section:" + dir, Order: index + 1,
				})
				existing[key] = true
			}
			walk(node.Children)
		}
	}
	walk(pkg.Manifest.Navigation)
}

func navigationCommonDir(nodes []NavigationNode) string {
	var pages []string
	var collect func([]NavigationNode)
	collect = func(items []NavigationNode) {
		for _, item := range items {
			if item.Page != "" {
				pages = append(pages, cleanZipPath(item.Page))
			}
			collect(item.Children)
		}
	}
	collect(nodes)
	if len(pages) == 0 {
		return ""
	}
	common := path.Dir(pages[0])
	for _, page := range pages[1:] {
		dir := path.Dir(page)
		for common != "." && common != "" && dir != common && !strings.HasPrefix(dir, common+"/") {
			common = path.Dir(common)
		}
	}
	if common == "." {
		return ""
	}
	return common
}

func docLocationForPath(manifest Manifest, source string) (string, string, string, bool) {
	clean := cleanZipPath(source)
	for locale, root := range manifest.LocaleRoots {
		if root != "" && strings.HasPrefix(clean, root+"/") {
			return locale, "", root, true
		}
	}
	defaultRoot := manifest.LocaleRoots[manifest.DefaultLocale]
	if defaultRoot == "" {
		return manifest.DefaultLocale, "", "", true
	}
	if strings.HasPrefix(clean, defaultRoot+"/") {
		return manifest.DefaultLocale, "", defaultRoot, true
	}
	return "", "", "", false
}

func parseDocFile(manifest Manifest, locale, version, root, source, raw string) (DocFile, []Issue, error) {
	fm, body, err := splitFrontmatter(raw)
	if err != nil {
		return DocFile{}, nil, err
	}
	docPath := docPathFromRoot(root, source)
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
		ContentRoot:    root,
		Path:           docPath,
		Slug:           slug,
		Title:          title,
		Content:        body,
		RawContent:     raw,
		Excerpt:        fm.Description,
		TranslationKey: fm.ID,
		Order:          fm.Order,
		Draft:          fm.Draft,
		ImageRefs:      imageRefs(source, body),
		Links:          docLinks(source, root, body),
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

func validateNavigation(pkg *Package, docPathSet map[string]bool) {
	seen := map[string]bool{}
	var walk func([]NavigationNode)
	walk = func(nodes []NavigationNode) {
		for _, node := range nodes {
			if node.Page != "" {
				page := strings.TrimSuffix(strings.TrimSuffix(cleanZipPath(node.Page), ".md"), ".mdx")
				if path.Base(page) == "index" {
					page = path.Dir(page)
					if page == "." {
						page = ""
					}
				}
				key := pkg.Manifest.DefaultLocale + "//" + page
				if seen[page] {
					pkg.Issues = append(pkg.Issues, Issue{Severity: "error", Code: "navigation_page_duplicate", Message: "navigation page is duplicated", Path: node.Page})
				} else if !docPathSet[key] {
					pkg.Issues = append(pkg.Issues, Issue{Severity: "error", Code: "navigation_page_missing", Message: "navigation page does not exist", Path: node.Page})
				}
				seen[page] = true
			}
			walk(node.Children)
		}
	}
	walk(pkg.Manifest.Navigation)
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
