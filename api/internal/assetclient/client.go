// Package assetclient is the docs site's gateway to the asset service for
// collection cover images. Collection covers are public assets, so docs stores
// the finalized CDN URL on the collection record.
package assetclient

import "context"

type InitInput struct {
	Filename   string
	Mime       string
	Category   string
	Visibility string
	Size       int64
}

type InitOutput struct {
	UploadURL     string
	UploadToken   string
	UploadHeaders map[string]string
}

type View struct {
	ID       string
	MediaKey string
	CdnURL   string
	Size     int64
	Mime     string
	Filename string
}

type ReferenceInput struct {
	AssetID, RefType, RefID string
	RefLabel, RefURL        string
}

type Client interface {
	UploadInit(ctx context.Context, bearer string, in InitInput) (InitOutput, error)
	Finalize(ctx context.Context, bearer, uploadToken string) (View, error)
	Upload(ctx context.Context, bearer string, in InitInput, data []byte) (View, error)
	RegisterReference(ctx context.Context, bearer string, in ReferenceInput) error
	UnregisterReference(ctx context.Context, bearer string, in ReferenceInput) error
}
