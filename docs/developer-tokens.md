# 开发者令牌投稿

开发者令牌通过 Docs 站点的 `/api/v1` 调用内容接口，使用 `Authorization: Bearer <PAT>`。Nuxt 同源入口会保留显式 PAT；过期或撤销后不会回退到浏览器登录会话。

账户中心按账号当前权限显示可选操作。令牌权限与账号当前的站点、文档集及文档权限取交集；勾选权限不会授予账号原本没有的角色。旧令牌不会自动取得新增权限，需要在账户中心创建包含所需操作的新令牌。

## 可选操作

| 操作 | Capability | 范围 |
| --- | --- | --- |
| 管理文档集 | `docs.collection.manage` | 创建、编辑、删除文档集，封面上传、语言注册/删除/翻译草稿克隆、发布版本初始化/克隆、管理目录树 |
| 管理文档版本 | `docs.version.manage` | 创建及编辑文档集内部版本，设置默认版本、发布/归档状态 |
| 读取文档 | `docs.document.read` | 读取有权访问的文档和草稿、查询文档列表、读取版本元数据 |
| 创建文档 | `docs.document.create` | 创建目录/草稿，在目标文档集中上传正文图片 |
| 编辑文档 | `docs.document.update` | 编辑正文、元数据、目录顺序，以及目标文档的正文图片 |
| 发布文档 | `docs.document.publish` | 发布文档 |
| 下架文档 | `docs.document.archive` | 下架并保留文档内容 |
| 永久删除文档 | `docs.document.delete_permanently` | 永久删除文档及其子文档，单独勾选 |
| 导入文档 | `docs.import.manage` | 上传/预检文档包、查询批次、确认导入及图片上传 |

创建或克隆新的文档集需要站点级文档集管理权限；只管理某个文档集的账号不能据此创建新的站点级文档集。内部版本管理和导入沿用现有站点管理员权限。普通作者继续受到文档集授权及文档所有者约束。

编辑权限不能代替发布/下架权限；通过 `PATCH` 修改状态也会检查对应权限。站点设置、权限治理、GitHub 同步源及委托凭证管理不在发布令牌范围内。

## 本地启动

从 Docs 仓库运行：

```powershell
./scripts/dev-publishing.ps1 -Action Up
```

这个入口复用 Workspace 的 Docs `Isolated` 合同，并通过产品支持的 `GF_*` 环境变量连接令牌目录和媒体授权。默认入口：

- Docs：`http://docs.dev.yuelili.test:3003`
- 账户中心令牌页：`http://account-docs.dev.yuelili.test:3601/developer-tokens`
- 回环后端：Docs `8086`、Identity `8781`、Asset `8782`

它与 Blog 的组合独立。停止只使用 `./scripts/dev-publishing.ps1 -Action Down`；端口可通过 `LOCAL_*_PORT` 显式覆盖。本地依赖采用 Workspace 声明的相邻源码覆盖，不代表正式依赖制品已发布。

## 从本地创建并发布文档

在账户中心创建至少包含“管理文档集、创建文档、发布文档”的令牌，将其通过当前终端的 `YUELI_DOCS_TOKEN` 环境变量提供给脚本。下面请求创建一个新文档集和首篇文档；已有文档集可通过公开的 `GET /api/v1/collections/{slug}` 查找 ID。

```powershell
$base = 'http://docs.dev.yuelili.test:3003'
$headers = @{ Authorization = "Bearer $env:YUELI_DOCS_TOKEN" }
$collectionBody = @{
    title = '我的项目'; slug = 'my-project'
    defaultLocale = 'zh-CN'; semanticVersion = '1.0.0'
} | ConvertTo-Json
$collection = Invoke-RestMethod "$base/api/v1/collections" -Method Post `
    -Headers $headers -ContentType 'application/json; charset=utf-8' -Body $collectionBody

$documentBody = @{
    collectionId = $collection.id; title = '快速开始'; slug = 'getting-started'
    locale = 'zh-CN'; content = '<h2>安装</h2><p>这里是正文。</p>'
} | ConvertTo-Json
$created = Invoke-RestMethod "$base/api/v1/docs" -Method Post `
    -Headers $headers -ContentType 'application/json; charset=utf-8' -Body $documentBody
Invoke-RestMethod "$base/api/v1/docs/$($created.doc.id)/publish" -Method Post -Headers $headers
```

公开页面为 `/<collection-slug>/<document-path>`。创建子文档时传 `parentId`；未传 `versionId` 时使用默认内部版本。Markdown 项目及相对图片使用 [Docs Import Package](import-package.md) 打包导入；`POST /api/v1/imports/docs/{id}/confirm` 返回 `202` 后，通过批次详情查询 `completed` 或 `failed`，不能只按受理状态判断发布成功。

## 自动化接口

| 环节 | 接口 | 令牌操作 |
| --- | --- | --- |
| 文档集发现 | `GET /api/v1/collections`、`GET /api/v1/collections/{slug}` | 公开读取 |
| 创建 / 编辑 / 删除文档集 | `POST /api/v1/collections`、`PATCH / DELETE /api/v1/collections/{id}` | 管理文档集 |
| 封面 | `POST /api/v1/collections/{id}/cover`、`POST .../cover/finalize` | 管理文档集 |
| 管理目录树 | `GET /api/v1/manage/collections/{slug}/tree` | 管理文档集 |
| 语言列表 / 注册与更新 | `GET / POST /api/v1/manage/collections/{id}/locales` | 管理文档集 |
| 删除未使用的非默认语言 | `DELETE /api/v1/manage/collections/{id}/locales/{locale}` | 管理文档集 |
| 克隆翻译草稿 | `POST /api/v1/manage/collections/{id}/locales/clone` | 管理文档集 |
| 克隆发布版本 / 初始化版本号 | `POST /api/v1/manage/collections/{id}/clone-release`、`POST .../release` | 管理文档集 |
| 读取内部版本 | `GET /api/v1/collections/{id}/versions` | 读取文档、创建文档、管理文档集或管理文档版本之一，且账号当前有目标范围权限 |
| 创建 / 更新内部版本 | `POST /api/v1/collections/{id}/versions`、`PATCH /api/v1/manage/collections/{id}/versions/{versionId}` | 管理文档版本 |
| 文档列表 / 详情 | `GET /api/v1/docs?collectionId=...`、`GET /api/v1/manage/docs`、`GET /api/v1/docs/{id}` | 读取文档 |
| 创建 / 编辑 | `POST /api/v1/docs`、`PATCH /api/v1/docs/{id}` | 创建文档 / 编辑文档 |
| 发布 / 下架 | `POST /api/v1/docs/{id}/publish`、`POST .../archive` | 发布文档 / 下架文档 |
| 永久删除 | `DELETE /api/v1/docs/{id}` | 永久删除文档 |
| 正文图片 | `POST /api/v1/images`、`POST /api/v1/images/finalize` | 使用 `collectionId` 检查创建权限；使用 `documentId` 检查编辑权限 |
| 文档包 | `POST / GET /api/v1/imports/docs`、`GET .../{id}`、`POST .../{id}/confirm` | 导入文档 |

图片初始化请求包含 `filename`、`mime`、`size` 和相应的目标 ID；按返回的 `uploadUrl` 与 `uploadHeaders` 执行真实 `PUT`，再带 `uploadToken` 和相同目标调用 finalize。正文保存 finalize 返回的 `url`；封面 finalize 自动保存。媒体 Profile 由 Docs 的当前权限推导，脚本无需自行申请 Asset 管理权限或构造对象 Key。业务引用由现有引用对账服务登记。

当前正文图片直传 Profile 接受 WebP / GIF；封面接受 JPEG / PNG / WebP；文档包导入图片接受 JPEG / PNG / WebP / GIF。它们继续遵守站点已注册的格式和大小限制。

成功响应直接返回端点 DTO，失败使用 HTTP Problem 的 `status/code/traceId`。新客户端以 HTTP 状态判断结果，不解析错误描述文本。完整结构见仓库 [OpenAPI](../contracts/openapi/docs.json)。
