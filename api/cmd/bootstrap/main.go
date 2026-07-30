package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const initialHomeConfig = `
INSERT INTO home_config (
	key, quick_links, featured_collections,
	home_eyebrow, home_title, home_subtitle,
	site_title, site_description, support_email,
	footer_tagline, footer_copyright
)
VALUES (
	'default', '[]', '[]',
	'知识库', '把复杂知识，整理成可持续维护的文档',
	'指南、参考与教程，在同一个清晰的阅读系统中持续演进。',
	'月离文档', '指南、参考与教程', '',
	'持续整理值得长期维护的知识。', ''
)
ON CONFLICT (key) DO NOTHING`

func main() {
	databaseURL := strings.TrimSpace(os.Getenv("DOCS_DATABASE_URL"))
	if databaseURL == "" {
		fail("DOCS_DATABASE_URL is required")
	}
	database, err := sql.Open("postgres", databaseURL)
	if err != nil {
		fail("open database: %v", err)
	}
	defer database.Close()
	ctx := context.Background()
	if err := database.PingContext(ctx); err != nil {
		fail("connect database: %v", err)
	}
	if _, err := database.ExecContext(ctx, initialHomeConfig); err != nil {
		fail("install initial home configuration: %v", err)
	}
	fmt.Println("Docs initial records are ready")
}

func fail(format string, arguments ...any) {
	fmt.Fprintf(os.Stderr, "bootstrap: "+format+"\n", arguments...)
	os.Exit(1)
}
