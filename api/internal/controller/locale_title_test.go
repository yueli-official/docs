package controller

import (
	"encoding/json"
	"github.com/yueli-official/docs/api/internal/model"
	"testing"
)

func TestLocaleViewPreservesCollectionTitle(t *testing.T) {
	var locale model.CollectionLocale
	if err := json.Unmarshal([]byte(`{"locale":"en-US","title":"After Effects Scripting","label":"English"}`), &locale); err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(localeView(&locale))
	if err != nil {
		t.Fatal(err)
	}
	var view map[string]any
	if err := json.Unmarshal(data, &view); err != nil {
		t.Fatal(err)
	}
	if view["title"] != "After Effects Scripting" || view["label"] != "English" {
		t.Fatalf("locale view lost localized title: %s", data)
	}
}
