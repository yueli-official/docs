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
	'', '月离文档',
	'指南、参考与教程',
	'月离文档', '指南、参考与教程', '',
	'持续整理值得长期维护的知识。', ''
)
ON CONFLICT (key) DO NOTHING`

const localCollections = `
INSERT INTO collections (
    id, slug, title, description, icon, sort_order, author_sub,
    created_at, updated_at
)
VALUES
    ('019c52f0-3000-7000-8000-000000000101', 'getting-started',
     '快速开始', '从安装到第一次发布，快速了解文档站。',
     'i-tabler-rocket', 10, $1, TIMESTAMPTZ '2026-08-10 09:00:00+08', TIMESTAMPTZ '2026-08-10 09:00:00+08'),
    ('019c52f0-3000-7000-8000-000000000102', 'operations',
     '运维指南', '部署、观测和常见问题排查。',
     'i-tabler-tool', 20, $1, TIMESTAMPTZ '2026-08-11 09:00:00+08', TIMESTAMPTZ '2026-08-11 09:00:00+08')
ON CONFLICT DO NOTHING`

const localVersions = `
INSERT INTO collection_versions (
    id, collection_id, key, label, status, is_default, sort_order,
    created_at, updated_at
)
VALUES
    ('019c52f0-3000-7000-8000-000000000201', '019c52f0-3000-7000-8000-000000000101',
     'default', '默认版本', 'published', true, 0,
     TIMESTAMPTZ '2026-08-10 09:00:00+08', TIMESTAMPTZ '2026-08-10 09:00:00+08'),
    ('019c52f0-3000-7000-8000-000000000202', '019c52f0-3000-7000-8000-000000000102',
     'default', '默认版本', 'published', true, 0,
     TIMESTAMPTZ '2026-08-11 09:00:00+08', TIMESTAMPTZ '2026-08-11 09:00:00+08')
ON CONFLICT DO NOTHING`

const localDocuments = `
INSERT INTO docs (
    id, collection_id, version_id, slug, title, content, excerpt,
    status, locale, translation_key, sort_order, author_sub,
    created_at, updated_at
)
VALUES
    ('019c52f0-3000-7000-8000-000000000301', '019c52f0-3000-7000-8000-000000000101',
     '019c52f0-3000-7000-8000-000000000201', 'overview', '认识月离文档',
     E'# 认识月离文档\n\n月离文档把指南、参考和教程整理进统一的层级阅读系统。\n\n## 从这里开始\n\n先完成安装，再创建自己的第一个文档集。',
     '了解文档站的内容结构、阅读方式与管理入口。',
     'published', 'en', 'docs-seed-overview', 10, $1,
     TIMESTAMPTZ '2026-08-12 09:30:00+08', TIMESTAMPTZ '2026-08-12 09:30:00+08'),
    ('019c52f0-3000-7000-8000-000000000302', '019c52f0-3000-7000-8000-000000000101',
     '019c52f0-3000-7000-8000-000000000201', 'installation', '安装与首次运行',
     E'# 安装与首次运行\n\n准备 PostgreSQL、Identity 与 Asset，然后通过 Workspace 启动 Docs。\n\n## 验证\n\n打开首页和控制台，确认健康检查与登录流程正常。',
     '准备依赖并完成第一次本地启动。',
     'published', 'en', 'docs-seed-installation', 20, $1,
     TIMESTAMPTZ '2026-08-13 10:00:00+08', TIMESTAMPTZ '2026-08-13 10:00:00+08'),
    ('019c52f0-3000-7000-8000-000000000303', '019c52f0-3000-7000-8000-000000000101',
     '019c52f0-3000-7000-8000-000000000201', 'configuration', '站点配置',
     E'# 站点配置\n\n在站点设置中维护首页、页脚和基础信息。\n\n## 原则\n\n运行地址与密钥来自部署环境，后台只管理产品内容。',
     '配置首页、页脚与站点基础信息。',
     'published', 'en', 'docs-seed-configuration', 30, $1,
     TIMESTAMPTZ '2026-08-14 11:00:00+08', TIMESTAMPTZ '2026-08-14 11:00:00+08'),
    ('019c52f0-3000-7000-8000-000000000304', '019c52f0-3000-7000-8000-000000000102',
     '019c52f0-3000-7000-8000-000000000202', 'deployment', '生产环境部署',
     E'# 生产环境部署\n\n使用锁定依赖完成构建、迁移和健康检查。\n\n## 发布前\n\n验证身份回调、资源地址与公开阅读链路。',
     '完成生产构建、迁移和发布检查。',
     'published', 'en', 'docs-seed-deployment', 10, $1,
     TIMESTAMPTZ '2026-08-15 09:00:00+08', TIMESTAMPTZ '2026-08-15 09:00:00+08'),
    ('019c52f0-3000-7000-8000-000000000305', '019c52f0-3000-7000-8000-000000000102',
     '019c52f0-3000-7000-8000-000000000202', 'observability', '运行状态与统计',
     E'# 运行状态与统计\n\n控制台展示真实浏览、访客、来源和搜索统计。\n\n## 数据边界\n\n统计保存在当前 Docs 实例自己的数据库中。',
     '理解健康检查与控制台统计口径。',
     'published', 'en', 'docs-seed-observability', 20, $1,
     TIMESTAMPTZ '2026-08-16 09:00:00+08', TIMESTAMPTZ '2026-08-16 09:00:00+08'),
    ('019c52f0-3000-7000-8000-000000000306', '019c52f0-3000-7000-8000-000000000102',
     '019c52f0-3000-7000-8000-000000000202', 'troubleshooting', '常见问题排查',
     E'# 常见问题排查\n\n先检查 Workspace Session、服务健康与浏览器网络请求。\n\n## 缩小问题\n\n分别验证公开 API、同源 BFF 与页面渲染。',
     '定位启动、登录、网络和配置问题。',
     'published', 'en', 'docs-seed-troubleshooting', 30, $1,
     TIMESTAMPTZ '2026-08-17 09:00:00+08', TIMESTAMPTZ '2026-08-17 09:00:00+08')
ON CONFLICT DO NOTHING`

func installLocalSeed(ctx context.Context, tx *sql.Tx, authorSub string) error {
	for _, step := range []struct {
		name  string
		query string
		args  []any
	}{
		{name: "collections", query: localCollections, args: []any{authorSub}},
		{name: "versions", query: localVersions},
		{name: "documents", query: localDocuments, args: []any{authorSub}},
	} {
		if _, err := tx.ExecContext(ctx, step.query, step.args...); err != nil {
			return fmt.Errorf("install local Docs seed %s: %w", step.name, err)
		}
	}
	return nil
}

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
	if strings.EqualFold(strings.TrimSpace(os.Getenv("DOCS_DEV_SEED")), "true") {
		authorSub := strings.TrimSpace(os.Getenv("DOCS_DEV_SEED_AUTHOR_SUB"))
		if authorSub == "" {
			fail("DOCS_DEV_SEED_AUTHOR_SUB is required when DOCS_DEV_SEED=true")
		}
		transaction, err := database.BeginTx(ctx, nil)
		if err != nil {
			fail("begin local Docs seed: %v", err)
		}
		if err := installLocalSeed(ctx, transaction, authorSub); err != nil {
			_ = transaction.Rollback()
			fail("%v", err)
		}
		if err := transaction.Commit(); err != nil {
			fail("commit local Docs seed: %v", err)
		}
	}
	fmt.Println("Docs initial records are ready")
}

func fail(format string, arguments ...any) {
	fmt.Fprintf(os.Stderr, "bootstrap: "+format+"\n", arguments...)
	os.Exit(1)
}
