# 开发者令牌完整投稿验收

2026-09-12，本轮仅本地交付。Docs 工作树包含实现、回归测试、启动入口和使用文档；用户已授权本地提交；未推送或部署生产。Blog、Identity、Asset、Foundation 及 Workspace 的版本化实现/合同未由本轮修改。

## 交付

- Docs 权限目录从 6 项扩至 9 项，新增文档集管理、内部版本管理、独立永久删除文档。文档集管理覆盖创建/更新/删除、封面、语言管理、翻译草稿和语义发布版本克隆；正文图片沿用文档创建或编辑权限。
- PAT 使用精确 method/path 白名单，账号/资源授权继续由原控制器执行。版本列表同时检查当前账号在目标范围的权限；撤销集合授权后不能继续读取。
- 修复内部版本接口把站点级 `docs.version.manage` 用于文档集 scope、导致管理员被拒绝的问题。克隆新的文档集发布版本额外检查站点级创建能力，避免集合级管理员通过克隆创建根文档集。
- 媒体授权只派生当前账号及令牌允许的 `docs-import-image`、`docs-collection-cover`、`docs-content-image` Profile。PAT 封面持久化后由既有服务身份引用对账处理，不再调用只接受普通凭证的旧引用接口而返回失败。
- 新增 [本地启动入口](../../../scripts/dev-publishing.ps1) 和 [权限/API 使用说明](../../../docs/developer-tokens.md)。沿用 Workspace 单一 Isolated 合同和产品 GF 环境配置，不手写 `.doctor` 状态或另一份进程合同。

## 实际验证

1. Go controller、docsauthz、catalog、server 测试通过。新回归覆盖 29 个操作的勾选范围、拒绝未知/多余路径、目录仅供 Identity service 读取、媒体 Profile 隔离、集合级管理员不能创建根文档集/越权访问兄弟集合、作者及管理权限撤销。
2. HTTP operation 声明检查与 80 个 OpenAPI operation 覆盖通过；没有新增 DTO/错误目录或迁移。
3. 真实 PostgreSQL `TestSourceTracksImportSyncDeleteAndRestore` 通过，使用事务内临时表并回滚，验证既有媒体引用快照生命周期。
4. CLI Playwright Chrome 完整流程通过（1 项，约 1.4 分钟）：账户中心实际勾选 9 项并创建 PAT → 创建/更新文档集 → 注册/删除语言 → 创建/发布内部版本 → PNG 封面 PUT/finalize/交付 → 创建父子文档 → GIF 正文图片 PUT/finalize/渲染 → 编辑/发布 → 翻译克隆/文档集发布版本克隆 → 两次带 PNG 的 ZIP 导入确认与完成 → 下架/删除 → 撤销令牌后拒绝请求。带浏览器会话的已撤销 PAT 同样不能回退登录凭证。
5. 真实账户中心与阅读页桌面、390px 手机截图已查看；无水平溢出、页面异常或失败资源。测试生成的文档集与令牌由 finally 清理，上传对象按现有 Asset 引用/回收生命周期管理。
6. 新增 Playwright 用例及引用的 runtime 工具通过独立严格 TypeScript 检查：`tsc --ignoreConfig --noEmit --strict --module esnext --moduleResolution bundler --target ES2022 --esModuleInterop --skipLibCheck --types node test/e2e/personal-token-publishing.spec.ts`。

截图和结果位于忽略的 `web/test-results/e2e/pat-publishing-20260912-verified/`：

- [JUnit](../../../web/test-results/e2e/pat-publishing-20260912-verified/junit.xml)
- [令牌桌面](../../../web/test-results/e2e/pat-publishing-20260912-verified/artifacts/personal-token-publishing--39de9-ument-and-import-publishing/developer-token-desktop.png)
- [令牌手机](../../../web/test-results/e2e/pat-publishing-20260912-verified/artifacts/personal-token-publishing--39de9-ument-and-import-publishing/developer-token-mobile.png)
- [文档桌面](../../../web/test-results/e2e/pat-publishing-20260912-verified/artifacts/personal-token-publishing--39de9-ument-and-import-publishing/published-document.png)
- [文档手机](../../../web/test-results/e2e/pat-publishing-20260912-verified/artifacts/personal-token-publishing--39de9-ument-and-import-publishing/published-document-mobile.png)

## 验证限制

- 运行及 Go 测试使用 Workspace 生成的本地源码依赖 overlay。仓库声明的 Foundation Go v0.4.1 尚缺本仓已有的 `IsPersonalToken` / `AllowsPersonalCapability` API，直接按正式依赖运行测试不能编译；本轮未发布/替换共享依赖。
- 全站 Nuxt typecheck 未通过：本地 Foundation 源码包缺少 `vue`、`@nuxt/ui/composables`、`@nuxt/kit` 等类型解析，并连带产生未修改的 Docs 管理页及共享组件隐式 any 错误。没有为此修改共享包或无关页面。独立 E2E 类型检查及真实浏览器运行已通过，不能将其表述为全站 typecheck 通过。
- 正文直传沿用现有 WebP/GIF Profile；PNG/JPEG 可用于封面或文档包导入，格式没有被令牌权限放宽。
- 历史 Work 中的真实外部 GitHub Release `docs.zip` 正向联调仍保留，未计入本轮本地任务。

## 当前运行

Docs Session：`20260912T073711Z-2200`，Docs/Account/Identity/Asset 均 ready。Docs `3003/8086`、Account `3601`、Identity `8781`、Asset `8782`。本地站点 `http://docs.dev.yuelili.test:3003`，令牌页 `http://account-docs.dev.yuelili.test:3601/developer-tokens`。

Blog 的 `20260912T071938Z-54828` 保持原 PID、端口和 ready 状态；只通过 Workspace 精确停止/重启了 Docs。
