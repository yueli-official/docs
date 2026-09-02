package catalog

import (
	"context"
	"testing"

	"github.com/yueli-official/docs/api/internal/assetclient"
)

type imageAssetFake struct {
	init assetclient.InitInput
	view assetclient.View
}

func (f *imageAssetFake) UploadInit(_ context.Context, _ string, in assetclient.InitInput) (assetclient.InitOutput, error) {
	f.init = in
	return assetclient.InitOutput{UploadURL: "https://asset.test/upload", UploadToken: "upload-token"}, nil
}

func (f *imageAssetFake) Finalize(context.Context, string, string) (assetclient.View, error) {
	return f.view, nil
}

func (*imageAssetFake) Upload(context.Context, string, assetclient.InitInput, []byte) (assetclient.View, error) {
	return assetclient.View{}, nil
}

func (*imageAssetFake) RegisterReference(context.Context, string, assetclient.ReferenceInput) error {
	return nil
}

func (*imageAssetFake) UnregisterReference(context.Context, string, assetclient.ReferenceInput) error {
	return nil
}

func TestDocumentImageUsesConsumerProfileAndStableDeliveryURL(t *testing.T) {
	fake := &imageAssetFake{view: assetclient.View{
		ID: "asset-1", MediaKey: "docs_content_image_opaque",
	}}
	svc := New(nil).WithAssets(fake, "")

	init, err := svc.InitDocumentImage(context.Background(), "bearer", "diagram.webp", "image/webp", 1234)
	if err != nil {
		t.Fatalf("InitDocumentImage() error = %v", err)
	}
	if init.UploadToken != "upload-token" || fake.init.Category != "docs-content-image" || fake.init.Visibility != "public" {
		t.Fatalf("unexpected image init: output=%+v input=%+v", init, fake.init)
	}

	got, err := svc.FinalizeDocumentImage(context.Background(), "bearer", "upload-token")
	if err != nil {
		t.Fatalf("FinalizeDocumentImage() error = %v", err)
	}
	if want := "/media/docs%2Fdocs-content-image%2Fopaque?format=webp&name=inline&v=1"; got != want {
		t.Fatalf("FinalizeDocumentImage() = %q, want %q", got, want)
	}
}
