package docsdiscovery

import (
	"testing"

	"github.com/yueli-official/foundation/go/discovery"

	"platform/products/docs/api/internal/model"
)

func TestProjectDocUsesSEOFieldsAndVersionLocaleIdentity(t *testing.T) {
	module := discovery.MustCompile(discovery.Definition{
		ContractVersion: discovery.ContractVersion,
		Site: discovery.SiteProfile{
			Origin: "https://docs.example", Name: "Docs", DefaultLocale: "en",
		},
		URLPolicy: discovery.URLPolicy{PreserveQuery: true},
	})
	projection, err := ProjectDoc(module, &model.Doc{
		ID: "one", Title: "Install", Excerpt: "Excerpt",
		SEOTitle: "Install Yueli", SEODescription: "SEO description",
		Status: "published", Locale: "zh-CN",
	}, &model.Collection{
		ID: "collection", Slug: "guide", Title: "Guide",
	}, &model.CollectionVersion{
		ID: "version", Key: "v2", Status: "published", IsDefault: false,
	}, "start/install", "en")
	if err != nil {
		t.Fatal(err)
	}
	if projection.Head.Title != "Install Yueli" ||
		projection.CanonicalURL != "https://docs.example/guide/start/install?locale=zh-CN&version=v2" {
		t.Fatalf("docs discovery facts drifted: %#v", projection)
	}
	if projection.Sitemap == nil || projection.Sitemap.Location != projection.CanonicalURL {
		t.Fatalf("versioned doc sitemap identity drifted: %#v", projection.Sitemap)
	}
}
