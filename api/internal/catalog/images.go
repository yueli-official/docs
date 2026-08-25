package catalog

import (
	"context"
	"net/url"

	"github.com/yueli-official/docs/api/internal/assetclient"
	"github.com/yueli-official/docs/api/internal/docserr"
)

const docsContentImageProfile = "docs-content-image"

func (s *Service) InitDocumentImage(ctx context.Context, bearer, filename, mime string, size int64) (assetclient.InitOutput, error) {
	if s.asset == nil {
		return assetclient.InitOutput{}, docserr.UpstreamFailed("asset client not configured")
	}
	if mime == "" {
		mime = "application/octet-stream"
	}
	return s.asset.UploadInit(ctx, bearer, assetclient.InitInput{
		Filename:   filename,
		Mime:       mime,
		Size:       size,
		Category:   docsContentImageProfile,
		Visibility: "public",
	})
}

func (s *Service) FinalizeDocumentImage(ctx context.Context, bearer, uploadToken string) (string, error) {
	if s.asset == nil {
		return "", docserr.UpstreamFailed("asset client not configured")
	}
	view, err := s.asset.Finalize(ctx, bearer, uploadToken)
	if err != nil {
		return "", err
	}
	if view.MediaKey == "" {
		return view.CdnURL, nil
	}
	return "/media/" + url.PathEscape(view.MediaKey) + "?format=webp&name=inline", nil
}
