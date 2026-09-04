package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	steps := []struct {
		name string
		env  []string
		args []string
	}{
		{name: "OpenAPI", env: []string{"DOCS_OPENAPI_OUTPUT=../contracts/openapi/docs.json"}, args: []string{"run", "./cmd/docs"}},
		{name: "operation manifest", args: []string{"run", "./cmd/httpoperations"}},
		{name: "generated contracts", args: []string{
			"run", "github.com/yueli-official/foundation/go/httpcontract/cmd/httpcontract@v0.4.1",
			"-errors", "contracts/http-result/error-catalog.json",
			"-operations", "contracts/http-result/operations.json",
			"-generate-go", "internal/docserr/catalog_gen.go", "-package", "docserr",
			"-generate-ts", "../web/app/generated/docsFailure.ts", "-ts-type", "DocsFailure",
			"-generate-i18n", "../web/app/generated/error-i18n-inventory.json",
		}},
		{name: "legacy error catalog", args: []string{"run", "./cmd/errorcatalog"}},
	}
	for _, step := range steps {
		command := exec.Command("go", step.args...)
		command.Env = append(os.Environ(), step.env...)
		command.Stdout, command.Stderr = os.Stdout, os.Stderr
		if err := command.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "httpcontracts: %s: %v\n", step.name, err)
			os.Exit(1)
		}
	}
	fmt.Println("Docs HTTP contracts generated")
}
