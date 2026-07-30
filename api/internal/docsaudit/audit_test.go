package docsaudit_test

import (
	"testing"

	"github.com/yueli-official/foundation/go/audit"

	"github.com/yueli-official/docs/api/internal/docsaudit"
)

func TestDefinitionCompilesStableConsumerActions(t *testing.T) {
	catalog, err := audit.Compile(docsaudit.Definition())
	if err != nil {
		t.Fatal(err)
	}
	if catalog.Digest() == "" {
		t.Fatal("compiled definition has no digest")
	}
}
