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
	original := strings.TrimSpace(raw)
	destination := markdownDestination(original)
	if destination == "" || strings.HasPrefix(destination, "http://") || strings.HasPrefix(destination, "https://") {
		return refs
	}
	resolved, ok := safeJoin(markdownPath, destination)
	if !ok {
		return append(refs, ImageRef{Original: original})
	}
	return append(refs, ImageRef{Original: original, ResolvedPath: resolved})
}

func markdownDestination(raw string) string {
	raw = strings.TrimSpace(raw)
	if strings.HasPrefix(raw, "<") {
		if end := strings.Index(raw, ">"); end > 0 {
			return raw[1:end]
		}
	}
	for i, r := range raw {
		if (r == ' ' || r == '\t' || r == '\n') && (i == 0 || raw[i-1] != '\\') {
			return raw[:i]
		}
	}
	return raw
}

func docLinks(markdownPath, root, body string) []DocLink {
	var out []DocLink
	for _, m := range mdLinkRe.FindAllStringSubmatch(body, -1) {
		raw := markdownDestination(m[2])
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
		target := docPathFromRoot(root, resolved)
		if unescaped, err := url.PathUnescape(target); err == nil {
			target = unescaped
		}
		out = append(out, DocLink{Original: raw, TargetPath: path.Clean(target), Anchor: anchor})
	}
	return out
}
