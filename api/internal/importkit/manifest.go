package importkit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path"
	"regexp"
	"strings"
)

type docsManifest struct {
	Schema        string            `json:"$schema"`
	SchemaVersion int               `json:"schemaVersion"`
	DefaultLocale string            `json:"defaultLocale"`
	Locales       map[string]string `json:"locales"`
	Navigation    []NavigationNode  `json:"navigation"`
	Redirects     []Redirect        `json:"redirects"`
}

type NavigationNode struct {
	Page         string            `json:"page,omitempty"`
	Group        string            `json:"group,omitempty"`
	Translations map[string]string `json:"translations,omitempty"`
	Children     []NavigationNode  `json:"children,omitempty"`
}

type Redirect struct {
	From string `json:"from"`
	To   string `json:"to"`
}

var localePattern = regexp.MustCompile(`^[A-Za-z]{2,3}(?:-[A-Za-z0-9]{2,8})*$`)

func parseDocsManifest(data []byte, opts Options) (Manifest, error) {
	var declared docsManifest
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&declared); err != nil {
		return Manifest{}, fmt.Errorf("docs.json invalid: %w", err)
	}
	if declared.SchemaVersion != 1 {
		return Manifest{}, fmt.Errorf("docs.json schemaVersion must be 1")
	}
	manifest, err := manifestFromOptions(opts, declared.DefaultLocale, declared.Locales)
	if err != nil {
		return Manifest{}, err
	}
	manifest.Navigation = declared.Navigation
	manifest.Redirects = declared.Redirects
	return manifest, nil
}

func manifestFromOptions(opts Options, declaredDefault string, declaredLocales map[string]string) (Manifest, error) {
	collection := strings.TrimSpace(opts.Collection)
	if collection == "" {
		return Manifest{}, fmt.Errorf("target collection is required")
	}
	defaultLocale := strings.TrimSpace(opts.DefaultLocale)
	if defaultLocale == "" {
		defaultLocale = strings.TrimSpace(declaredDefault)
	}
	if defaultLocale == "" || !localePattern.MatchString(defaultLocale) {
		return Manifest{}, fmt.Errorf("default locale must be a BCP 47 language tag")
	}
	localeRoots := declaredLocales
	if len(localeRoots) == 0 {
		localeRoots = map[string]string{defaultLocale: "."}
	}
	locales := make([]string, 0, len(localeRoots))
	seenRoots := map[string]string{}
	for locale, root := range localeRoots {
		if !localePattern.MatchString(locale) {
			return Manifest{}, fmt.Errorf("locale %q must be a BCP 47 language tag", locale)
		}
		cleanRoot := cleanZipPath(root)
		if root == "." {
			cleanRoot = ""
		}
		if cleanRoot != "" && (strings.HasPrefix(cleanRoot, "../") || path.IsAbs(root)) {
			return Manifest{}, fmt.Errorf("locale %q has an unsafe content root", locale)
		}
		if other, exists := seenRoots[cleanRoot]; exists {
			return Manifest{}, fmt.Errorf("locales %q and %q use the same content root", other, locale)
		}
		seenRoots[cleanRoot] = locale
		localeRoots[locale] = cleanRoot
		locales = append(locales, locale)
	}
	if _, ok := localeRoots[defaultLocale]; !ok {
		return Manifest{}, fmt.Errorf("default locale must be declared in locales")
	}
	mode := strings.TrimSpace(opts.Mode)
	if mode == "" {
		mode = "upsert"
	}
	if mode != "create-only" && mode != "upsert" && mode != "replace-version" {
		return Manifest{}, fmt.Errorf("import mode is invalid")
	}
	return Manifest{SchemaVersion: 1, Collection: collection, DefaultLocale: defaultLocale, Locales: locales, LocaleRoots: localeRoots, Mode: mode}, nil
}
