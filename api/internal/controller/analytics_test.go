package controller

import (
	"testing"

	"github.com/yueli-official/foundation/go/traffic"
)

func TestAnalyticsTransportNormalizesSourcesAndVisitClasses(t *testing.T) {
	for input, expected := range map[string]string{
		"": "direct", " WWW.Google.COM. ": "google.com",
		"https://evil.test/path": "direct", "docs.example.test": "docs.example.test",
	} {
		if actual := normalizeAnalyticsSource(input); actual != expected {
			t.Fatalf("normalizeAnalyticsSource(%q) = %q, want %q", input, actual, expected)
		}
	}
	if actual := classifyAnalyticsVisit("Mozilla/5.0"); actual != traffic.VisitHuman {
		t.Fatalf("human class = %q", actual)
	}
	if actual := classifyAnalyticsVisit("Googlebot"); actual != traffic.VisitBot {
		t.Fatalf("bot class = %q", actual)
	}
}
