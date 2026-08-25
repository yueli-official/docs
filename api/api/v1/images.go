package v1

import "github.com/gogf/gf/v2/frame/g"

type ImageInitReq struct {
	g.Meta       `path:"/api/v1/images" method:"post" tags:"docs" summary:"Begin a document content image upload"`
	CollectionID string `json:"collectionId"`
	DocumentID   string `json:"documentId"`
	Filename     string `json:"filename" v:"required"`
	Mime         string `json:"mime"`
	Size         int64  `json:"size"`
}

type ImageInitRes struct {
	UploadURL     string            `json:"uploadUrl"`
	UploadToken   string            `json:"uploadToken"`
	UploadHeaders map[string]string `json:"uploadHeaders,omitempty"`
}

type ImageFinalizeReq struct {
	g.Meta       `path:"/api/v1/images/finalize" method:"post" tags:"docs" summary:"Finalize a document content image"`
	CollectionID string `json:"collectionId"`
	DocumentID   string `json:"documentId"`
	UploadToken  string `json:"uploadToken" v:"required"`
}

type ImageFinalizeRes struct {
	URL string `json:"url"`
}
