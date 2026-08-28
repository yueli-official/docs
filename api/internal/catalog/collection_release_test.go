package catalog

import "testing"

func TestSemanticVersionPattern(t *testing.T) {
	valid := []string{"0.0.0", "1.0.0", "3.11.9", "2026.8.28"}
	invalid := []string{"v1.0.0", "1.0", "1", "01.0.0", "1.0.0-beta", "latest"}
	for _, value := range valid {
		if !semanticVersionPattern.MatchString(value) {
			t.Fatalf("expected %q to be a semantic documentation version", value)
		}
	}
	for _, value := range invalid {
		if semanticVersionPattern.MatchString(value) {
			t.Fatalf("expected %q to be rejected", value)
		}
	}
}

func TestValidateDocumentBadge(t *testing.T) {
	for _, test := range []struct {
		name      string
		text      string
		icon      string
		wantError bool
	}{
		{name: "none"},
		{name: "text", text: "已废弃"},
		{name: "icon", icon: "i-tabler-alert-triangle"},
		{name: "mutually exclusive", text: "v2", icon: "i-tabler-sparkles", wantError: true},
		{name: "wrong icon set", icon: "i-lucide-fish", wantError: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			err := validateDocumentBadge(test.text, test.icon)
			if (err != nil) != test.wantError {
				t.Fatalf("validateDocumentBadge() error = %v, wantError = %v", err, test.wantError)
			}
		})
	}
}
