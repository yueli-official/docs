package importkit

type Options struct {
	MaxImageBytes int64
}

type Manifest struct {
	SchemaVersion int      `json:"schemaVersion"`
	Collection    string   `json:"collection"`
	Version       string   `json:"version"`
	DefaultLocale string   `json:"defaultLocale"`
	Locales       []string `json:"locales"`
	Mode          string   `json:"mode"`
}

type Package struct {
	Manifest Manifest
	Docs     []DocFile
	Assets   map[string]AssetFile
	Issues   []Issue
}

type DocFile struct {
	Locale         string
	VersionKey     string
	SourcePath     string
	Path           string
	Slug           string
	Title          string
	Content        string
	RawContent     string
	Excerpt        string
	TranslationKey string
	Order          int
	ImageRefs      []ImageRef
	Links          []DocLink
}

type AssetFile struct {
	SourcePath string
	Bytes      []byte
	Mime       string
	Size       int64
	SHA256     string
}

type ImageRef struct {
	Original     string
	ResolvedPath string
}

type DocLink struct {
	Original   string
	TargetPath string
	Anchor     string
}

type Issue struct {
	Severity string `json:"severity"`
	Code     string `json:"code"`
	Message  string `json:"message"`
	Path     string `json:"path,omitempty"`
}
