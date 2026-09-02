package catalog

import (
	"context"
	"errors"
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
	return publicImageURL(view, "inline")
}

func publicImageURL(view assetclient.View, rendition string) (string, error) {
	if view.MediaKey == "" {
		return "", errors.New("asset finalize did not return mediaKey")
	}
	return "/media/" + url.PathEscape(view.MediaKey) + "?format=webp&name=" + url.QueryEscape(rendition) + "&v=1", nil
}
