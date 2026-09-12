# Docs PAT 与项目文档包接入

## 配置

Docs `docs.personalTokens.siteId` 必须等于自己的 OIDC audience，例如 `docs-main-web`。Identity 的可信 `pat.applications` 添加：

```json
{"id":"docs-main-web","name":"月离文档","permissionsUrl":"https://docs.yuelili.com/api/v1/internal/personal-token/permissions","audience":"docs-main-web"}
```

服务间使用可达的内部地址，HTTPS 为默认要求；本机 loopback HTTP 允许。不要把 PAT 校验地址指向用户提供的 URL。

Asset 的 `asset.personalTokens.authorities` 将 `docs-main-web` 映射到 Docs `/api/v1/personal-token/media-authorization`。按令牌勾选及账号当前权限派生 Profile：导入对应 `docs-import-image`，文档集管理对应 `docs-collection-cover`，文档创建/编辑对应 `docs-content-image`。具体上传端点还必须检查目标文档集/文档；媒体 Profile 不代表内容管理授权。

本地通过 Workspace CLI 的进程环境设置 `GF_PAT_APPLICATIONS`、`GF_DOCS_PERSONALTOKENS_SITEID`、`GF_ASSET_PERSONALTOKENS_AUTHORITIES`，不手改 Workspace 生成配置。

Docs 的 `scripts/dev-publishing.ps1` 统一设置上述环境并调用 Workspace Isolated 入口。9 项令牌能力、完整投稿接口与脚本示例见[开发者令牌投稿](../../docs/developer-tokens.md)。创建/克隆新文档集检查站点根范围；内部版本管理按站点级能力检查。读取版本元数据也必须取目标范围的当前账号权限与 PAT 权限交集。

## 验证范围

当前开放文档读取、创建、编辑、发布、下架，以及导入的上传、批次读取、确认和专用媒体授权。治理、永久删除、来源凭证管理等路由对 PAT 拒绝。能力目录包含用户在子文档集上的能力，但每次请求仍按实际资源检查，不能扩大到相邻集合。`docs.import.manage` 当前为站点管理员能力，不代表集合作者自动具有批量导入能力。JWT/Cookie 原流程保留。

确认仍为同步长请求：完成后返回 202 / `batch.status=completed`。脚本中断后按 batch ID 查询，不重新上传。已完成批次确认可重复返回结果；并发确认通过数据库 checked→running 条件更新只允许一次执行。

在文档写入事务提交前再检查令牌在线状态与当前 Docs 权限，失权时连同本次文档变更一起回滚；批次 completed 与文档写入同事务完成。导入完成只登记本批次目标的授权 scope，不能再次调用全站 bootstrap 同步。

页面稳定匹配使用最近一次成功批次记录的语言/源路径，不使用公开 slug 代替源路径。根 index 不得把自身设置为父文档；包根级页面保持根级。

## 发布 Skill

分发源：`skills/yueli-docs-publish`。运行脚本只依赖 Python 标准库；完整目录可复制到其他机器的 skills 目录。

`id` 是翻译关联键，重复导入按语言和源路径。删除源文件默认保留站内页面，`replace-version` 才会归档缺失页面。脚本只打包明确文档目录内的 Markdown、docs.json 和支持图片，避免包含仓库秘密。

## GitHub 参考

下载模块依据 [GitHub Releases API](https://docs.github.com/en/rest/releases/releases) 与 [Release Assets API](https://docs.github.com/en/rest/releases/assets)；正式 latest Release、附件 ID/大小/digest 为检查依据。来源绑定、定时调度和运行记录已接入本地。真实 GitHub Release 元数据已读取，目标 Keycrash Release 暂无 docs.zip；成功导入、二次更新与恢复用受控 Release 数据和真实隔离 PostgreSQL schema 验证。2026-09-09 已随 Docs server-20260909-sync-1 生产部署；真实外部 docs.zip 正向联调仍等待项目发布附件。


## GitHub 来源运行配置

新增迁移 `0019_project_doc_sources`，包含来源绑定和运行记录两张表，不重写已有文档表。每个目标文档集绑定一个来源；目标版本由所选文档集的发布版本确定。

- `docs.projectDocs.encryptionKey`：独立随机 32 字节密钥的标准 Base64；持久保存在部署秘密存储。空值禁用来源同步，不影响手动 ZIP 导入。
- `docs.projectDocs.githubToken`：可选，仅发送给 api.github.com，用于提高公开 API 额度，不发送到附件地址。
- 定时检查 15 分钟；收到 GitHub 限流后按恢复时间延后，手动检查与新建会唤醒调度。
- 来源令牌必须属于配置者，且勾选 Docs 导入能力；密封存储，列表与详情不返回凭证。更换令牌可转交给当前操作者，提交前检查来源仍启用且凭证未改变。
- 当前源和相同附件重复检查不会重复下载；新附件内容摘要相同则跳过导入。失败保留上次成功内容。
- PostgreSQL advisory lock 串行执行同一来源，进程退出自动释放。发现已提交但来源状态未完成的旧任务时，依据原批次完成恢复，不重复创建文档。

API：`/api/v1/import-sources` 列表/创建；`/{id}` 详情/暂停恢复/解除绑定；`/{id}/check` 立即检查。来源管理使用正常站点会话，不开放给发布 PAT。页面为 `/manage/import-sources`，入口在批量导入页右上角。

## 自动检查与手动检查

来源配置 autoCheck 独立于 enabled（暂停/恢复）。新增 0020_project_doc_auto_check：已有来源默认 true；新建请求明确提交 autoCheck。关闭后不参与周期轮询，仍可显式 POST /{id}/check；新建时关闭则不立即导入。运行中任务按既有提交检查完成，故障恢复仍能处理已领取任务。PATCH 可只提交 autoCheck，不应默认为 enabled=false。导入与同步页面通过 ImportNavigation 路由适配并列呈现，容器必须直接复用站点设置使用的 Foundation TabbedSurface；不在产品层重新写 UTabs 或复制容器样式。页面不重复展示自动检查说明，选项提示留在创建表单。此节功能本地已验收，待后续部署。
