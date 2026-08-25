package catalog

import (
	"testing"

	"github.com/yueli-official/foundation/go/identifier"

	"github.com/yueli-official/docs/api/internal/model"
)

func TestSanitizeHomeQuickLinksAssignsIdentifierModuleUUIDs(t *testing.T) {
	existing := identifier.MustNew().String()
	links := sanitizeHomeQuickLinks([]*model.HomeQuickLink{
		{ID: existing, Title: "Existing", To: "/existing", Enabled: true},
		{ID: "draft-link", Title: "Draft", To: "/draft", Enabled: true},
	})
	if len(links) != 2 {
		t.Fatalf("links = %d, want 2", len(links))
	}
	if links[0].ID != existing {
		t.Fatalf("existing ID = %q, want %q", links[0].ID, existing)
	}
	if links[1].ID == "draft-link" || links[1].ID == existing {
		t.Fatalf("draft ID was not replaced: %q", links[1].ID)
	}
	value, err := identifier.Parse(links[1].ID)
	if err != nil || value.Version() != 7 {
		t.Fatalf("assigned ID = %q, error = %v, want UUIDv7", links[1].ID, err)
	}
}

func TestNormalizeHomeConfigAllowsOptionalCopyright(t *testing.T) {
	got, err := normalizeHomeConfig(&model.HomeConfig{
		HomeTitle:       " 长期维护的文档 ",
		HomeSubtitle:    " 指南与参考 ",
		SiteTitle:       " 月离文档 ",
		SiteDescription: " 文档站 ",
		FooterTagline:   " 持续整理知识 ",
		FooterCopyright: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	if got.SiteTitle != "月离文档" || got.FooterCopyright != "" {
		t.Fatalf("normalized config = %#v", got)
	}
}

func TestNormalizeHomeConfigRejectsMissingRequiredContent(t *testing.T) {
	_, err := normalizeHomeConfig(&model.HomeConfig{
		HomeEyebrow:     "知识库",
		HomeTitle:       "长期维护的文档",
		HomeSubtitle:    "指南与参考",
		SiteTitle:       "月离文档",
		SiteDescription: "文档站",
	})
	if err == nil {
		t.Fatal("expected missing footer tagline to be rejected")
	}
}
