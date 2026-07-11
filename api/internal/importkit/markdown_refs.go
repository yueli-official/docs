package importkit

import (
	"net/url"
	"path"
	"regexp"
	"strings"
)

var (
	mdImageRe = regexp.MustCompile(`!\[[^\]]*\]\(([^)]+)\)`)
	imgSrcRe  = regexp.MustCompile(`<img[^>]+src=["']([^"']+)["']`)
	mdLinkRe  = regexp.MustCompile(`(?m)(^|[^!])\[[^\]]+\]\(([^)]+)\)`)
)

func imageRefs(markdownPath, body string) []ImageRef {
	var refs []ImageRef
	for _, m := range mdImageRe.FindAllStringSubmatch(body, -1) {
		refs = appendImageRef(refs, markdownPath, strings.TrimSpace(m[1]))
	}
	for _, m := range imgSrcRe.FindAllStringSubmatch(body, -1) {
		refs = appendImageRef(refs, markdownPath, strings.TrimSpace(m[1]))
	}
	return refs
}

func appendImageRef(refs []ImageRef, markdownPath, raw string) []ImageRef {
	if raw == "" || strings.HasPrefix(raw, "http://") || strings.HasPrefix(raw, "https://") {
		return refs
	}
	resolved, ok := safeJoin(markdownPath, raw)
	if !ok {
		return append(refs, ImageRef{Original: raw})
	}
	return append(refs, ImageRef{Original: raw, ResolvedPath: resolved})
}

func docLinks(markdownPath, locale, version, body string) []DocLink {
	var out []DocLink
	for _, m := range mdLinkRe.FindAllStringSubmatch(body, -1) {
		raw := strings.TrimSpace(m[2])
		if raw == "" || strings.HasPrefix(raw, "http://") || strings.HasPrefix(raw, "https://") || strings.HasPrefix(raw, "#") {
			continue
		}
		withoutAnchor, anchor, _ := strings.Cut(raw, "#")
		if !strings.HasSuffix(withoutAnchor, ".md") && !strings.HasSuffix(withoutAnchor, ".mdx") {
			continue
		}
		resolved, ok := safeJoin(markdownPath, withoutAnchor)
		if !ok {
			out = append(out, DocLink{Original: raw, Anchor: anchor})
			continue
		}
		target := docPathFromMarkdown(locale, version, resolved)
		if unescaped, err := url.PathUnescape(target); err == nil {
			target = unescaped
		}
		out = append(out, DocLink{Original: raw, TargetPath: path.Clean(target), Anchor: anchor})
	}
	return out
}
