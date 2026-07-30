package importkit

import (
	"strings"

	"gopkg.in/yaml.v3"
)

type frontmatter struct {
	Title          string `yaml:"title"`
	Slug           string `yaml:"slug"`
	Excerpt        string `yaml:"excerpt"`
	TranslationKey string `yaml:"translationKey"`
	Order          int    `yaml:"order"`
}

func splitFrontmatter(s string) (frontmatter, string, error) {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	if !strings.HasPrefix(s, "---\n") {
		return frontmatter{}, s, nil
	}
	rest := s[len("---\n"):]
	end := strings.Index(rest, "\n---")
	if end < 0 {
		return frontmatter{}, s, nil
	}
	var fm frontmatter
	if err := yaml.Unmarshal([]byte(rest[:end]), &fm); err != nil {
		return frontmatter{}, "", err
	}
	body := strings.TrimPrefix(rest[end+len("\n---"):], "\n")
	return fm, body, nil
}
