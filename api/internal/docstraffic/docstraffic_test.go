package docstraffic

import (
	"testing"

	"github.com/yueli-official/foundation/go/traffic"
)

func TestDefinitionOwnsDocumentTrafficVocabulary(t *testing.T) {
	catalog, err := traffic.Compile(Definition("Asia/Shanghai"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := catalog.NormalizeResource(traffic.Resource{Kind: ResourceDocument, ID: "doc-1"}); err != nil {
		t.Fatalf("document resource rejected: %v", err)
	}
	if _, err := catalog.NormalizeResource(traffic.Resource{Kind: "post", ID: "doc-1"}); err == nil {
		t.Fatal("Blog post resource leaked into Docs traffic vocabulary")
	}
}
