package catalog

import (
	"strings"
	"unicode"
)

// slugify converts a title to a URL-safe slug. Unicode letters and digits are
// kept as-is so CJK titles (e.g. "技术") produce a non-empty slug instead of
// being stripped to "" by ASCII-only rules. Copied verbatim from the blog
// catalog implementation.
func slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	var b strings.Builder
	prevHyphen := false
	for _, r := range s {
		switch {
		// keep Unicode letters/digits so a pure-CJK name (e.g. 技术) yields a
		// non-empty slug instead of being stripped to "" by ASCII-only rules.
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
			prevHyphen = false
		case b.Len() > 0 && !prevHyphen:
			b.WriteByte('-')
			prevHyphen = true
		}
	}
	return strings.Trim(b.String(), "-")
}
