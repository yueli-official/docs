package model

type ImportBatch struct {
	ID            string `json:"id" orm:"id"`
	CollectionID  string `json:"collectionId" orm:"collection_id"`
	VersionID     string `json:"versionId" orm:"version_id"`
	DefaultLocale string `json:"defaultLocale" orm:"default_locale"`
	Mode          string `json:"mode" orm:"mode"`
	Status        string `json:"status" orm:"status"`
	SummaryJSON   string `json:"summaryJson" orm:"summary_json"`
	ErrorMessage  string `json:"errorMessage" orm:"error_message"`
	CreatedBy     string `json:"createdBy" orm:"created_by"`
}

type ImportItem struct {
	ID                     string `json:"id" orm:"id"`
	BatchID                string `json:"batchId" orm:"batch_id"`
	Locale                 string `json:"locale" orm:"locale"`
	VersionKey             string `json:"versionKey" orm:"version_key"`
	Path                   string `json:"path" orm:"path"`
	SourceMarkdownPath     string `json:"sourceMarkdownPath" orm:"source_markdown_path"`
	Title                  string `json:"title" orm:"title"`
	Slug                   string `json:"slug" orm:"slug"`
	TranslationKey         string `json:"translationKey" orm:"translation_key"`
	Action                 string `json:"action" orm:"action"`
	TargetDocID            string `json:"targetDocId" orm:"target_doc_id"`
	BeforeDocJSON          string `json:"beforeDocJson" orm:"before_doc_json"`
	AfterDocJSON           string `json:"afterDocJson" orm:"after_doc_json"`
	IssuesJSON             string `json:"issuesJson" orm:"issues_json"`
	SourceContentHash      string `json:"sourceContentHash" orm:"source_content_hash"`
	TransformedContentHash string `json:"transformedContentHash" orm:"transformed_content_hash"`
}

type ImportAsset struct {
	ID           string `json:"id" orm:"id"`
	BatchID      string `json:"batchId" orm:"batch_id"`
	SourcePath   string `json:"sourcePath" orm:"source_path"`
	Data         []byte `json:"-" orm:"data"`
	AssetURL     string `json:"assetUrl" orm:"asset_url"`
	ContentHash  string `json:"contentHash" orm:"content_hash"`
	Status       string `json:"status" orm:"status"`
	ErrorMessage string `json:"errorMessage" orm:"error_message"`
}

type ImportAssetRef struct {
	ID               string `json:"id" orm:"id"`
	BatchID          string `json:"batchId" orm:"batch_id"`
	AssetID          string `json:"assetId" orm:"asset_id"`
	ItemID           string `json:"itemId" orm:"item_id"`
	MarkdownFilePath string `json:"markdownFilePath" orm:"markdown_file_path"`
	OriginalRef      string `json:"originalRef" orm:"original_ref"`
	RewrittenRef     string `json:"rewrittenRef" orm:"rewritten_ref"`
}
