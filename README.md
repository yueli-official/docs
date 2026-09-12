# 月离文档

Docs 是独立的文档消费者产品，拥有文档集合、层级页面、版本/语言、批量导入、搜索、公开阅读和管理后台。

批量导入采用 Markdown-first 的[Docs Import Package v1](docs/import-package.md)：普通目录只需 Markdown，多语言或显式导航包
附加 `docs.json`。
`api/` 与 `web/` 是本仓唯一实现真源；仓库不依赖 Platform 源码或工作区包。

## 边界

- Docs 自己拥有领域数据、PostgreSQL migration、授权实例、Discovery 发布和界面组件。
- Identity 只通过 OIDC issuer、Discovery 和 JWKS 证明用户身份。
- Asset 只通过公开 HTTP 合同管理文档集封面；Docs 不导入 Asset 内部代码。
- Foundation 通过正式 Go module 与 JS Release 提供跨产品协议原语。
- Foundation Traffic 提供幂等浏览、访客日与聚合原语；Docs 自己拥有 `document` 资源语义、来源归因、搜索统计、
  `/api/v1/docs/{id}/view`、dashboard overview 和控制台展示。统计真值保存在每个 Docs 实例自己的 PostgreSQL。
- 本地多仓编排属于 `workspace`；生产部署属于本仓 Compose。

不可变依赖与能力绑定记录在：

- `deploy/contracts/requirements.json`：消费者需要的能力；
- `deploy/deployment.lock.json`：能力到具体生产者版本的部署锁。

## 本地开发

推荐从相邻 `workspace` 仓启动：

```powershell
# 独立 Identity + Account + Asset + Docs
.\environments\docs-local\run.ps1 -Mode Isolated

# 通过 Workspace Binding 复用兼容的基础服务
.\environments\docs-local\run.ps1 -Mode Shared

.\environments\docs-local\run.ps1 -Action Down
```

所有端口均可通过 `LOCAL_*_PORT` 覆盖；`down docs` 只停止 Docs target 对应的 Workspace 会话，不会终止其他项目。

开发者令牌的完整本地投稿使用本仓 `./scripts/dev-publishing.ps1 -Action Up`，默认使用 Docs 独立端口并接通权限目录与媒体授权。权限列表、API 及本地脚本示例见[开发者令牌投稿](docs/developer-tokens.md)。

`docs-local` 默认启用正式 Dev Seed：2 个文档集、6 篇发布文档，以及可重复执行的 7 天浏览、来源和搜索样本。Seed 只在
`DOCS_DEV_SEED=true` 时运行，使用稳定 ID 且不覆盖已编辑记录；生产 Compose 不启用。需要空白本地库时，在启动前设置：

```powershell
$env:LOCAL_DOCS_DEV_SEED = "false"
```

## Docker Compose

本仓提供三种生产/预发布拓扑：

```powershell
# 完整独立部署：Identity + Account + Docs 专属 Asset + Docs
Copy-Item .env.example .env
docker compose -f compose.yaml up -d --wait

# 复用已有 Identity，部署 Docs 专属 Asset
Copy-Item deploy/env/hybrid.env.example .env
docker compose -f compose.hybrid.yaml up -d --wait

# 复用已有 Identity 与 Asset
Copy-Item deploy/env/attach.env.example .env
docker compose -f compose.attach.yaml up -d --wait
```

生成 `.env` 后必须填写所有空 secret、管理员 Subject 和外部服务 URL。宿主端口由
`DOCS_API_PORT`、`DOCS_WEB_PORT`、`IDENTITY_PORT`、`IDENTITY_ACCOUNT_PORT`、`ASSET_PORT` 配置，没有硬编码占用。

Docs PostgreSQL 使用锁定的 `postgres-zhparser` 镜像，因为现有全文检索 migration 明确依赖 `zhparser`；普通 PostgreSQL
不能替代该契约。

## 独立命令

```powershell
cd api
go run ./cmd/docs
go run ./cmd/errorcatalog

cd ..\web
pnpm install --frozen-lockfile --ignore-workspace
pnpm dev
```

运行配置模板位于 `api/manifest/config/config.example.yaml`。本仓不使用 `doctor.yaml`。

## 验收策略

API、Web、Compose 与浏览器合同均由本仓 CI 拥有。当前迁移批次按约定暂停逐产品测试；完成所有消费者迁移后，
再统一运行 API、前端、容器、Compose 和 Playwright 验收。
