package assetclient

import (
	"context"
	"strings"

	"github.com/yueli-official/foundation/go/identifier"
)

type Fake struct {
	Uploaded map[string][]byte
	Refs     []ReferenceInput
	Unrefs   []ReferenceInput
}

func (f *Fake) UploadInit(context.Context, string, InitInput) (InitOutput, error) {
	return InitOutput{UploadURL: "https://asset.test/blob/fake", UploadToken: "fake"}, nil
}

func (f *Fake) Finalize(context.Context, string, string) (View, error) {
	return View{ID: identifier.MustNew().String(), MediaKey: "docs_fake", Mime: "image/png", Filename: "fake.png"}, nil
}

func (f *Fake) Upload(_ context.Context, _ string, in InitInput, data []byte) (View, error) {
	if f.Uploaded == nil {
		f.Uploaded = map[string][]byte{}
	}
	key := "docs_" + strings.NewReplacer("/", "_", "\\", "_").Replace(in.Filename)
	f.Uploaded[key] = append([]byte(nil), data...)
	return View{ID: identifier.MustNew().String(), MediaKey: key, Size: in.Size, Mime: in.Mime, Filename: in.Filename}, nil
}

func (f *Fake) RegisterReference(_ context.Context, _ string, in ReferenceInput) error {
	f.Refs = append(f.Refs, in)
	return nil
}

func (f *Fake) UnregisterReference(_ context.Context, _ string, in ReferenceInput) error {
	f.Unrefs = append(f.Unrefs, in)
	return nil
}
