package assetclient

import (
	"context"
	"path"

	"github.com/google/uuid"
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
	return View{ID: uuid.NewString(), CdnURL: "https://asset.test/docs/fake.png", Mime: "image/png", Filename: "fake.png"}, nil
}

func (f *Fake) Upload(_ context.Context, _ string, in InitInput, data []byte) (View, error) {
	if f.Uploaded == nil {
		f.Uploaded = map[string][]byte{}
	}
	url := "https://asset.test/docs/" + path.Base(in.Filename)
	f.Uploaded[url] = append([]byte(nil), data...)
	return View{ID: uuid.NewString(), CdnURL: url, Size: in.Size, Mime: in.Mime, Filename: in.Filename}, nil
}

func (f *Fake) RegisterReference(_ context.Context, _ string, in ReferenceInput) error {
	f.Refs = append(f.Refs, in)
	return nil
}

func (f *Fake) UnregisterReference(_ context.Context, _ string, in ReferenceInput) error {
	f.Unrefs = append(f.Unrefs, in)
	return nil
}
