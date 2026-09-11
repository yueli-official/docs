# Docs 服务器部署

## Goal

将新版 Docs 作为全新实例部署到 https://docs.yuelili.com，接入现有 Account 和 Asset，完成构建与浏览器验收。

## Status

Finished

## Current

- 2026-09-08 已上线：API/Web 当前均为 `server-20260908-users-1` 且 healthy（含标题、300px 导航和权限页资料修复）。数据库、迁移、OIDC、Asset、Nginx 已完成。
- 用户明确旧站已有本地存档，本次空站部署，不迁移旧文档。
- 用户选择 Blog 式首位管理员认领，已实现并部署；正式 `/manage/setup` 登录后需要显式点击，初次交付未认领，代理未代认领；后续已出现 2 个文档集共 220 篇公开文档，当前管理员状态需按实际会话判断。
- 当前 Go 全量、72 operation 合同检查、Web typecheck / 54 单测 / build 通过；Memory 20 并发与实际 PostgreSQL 并发/重启/管理员丢失测试通过。
- 本地认领 Playwright 15 项及重新登录通过；正式 HTTPS 13 项通过。旧综合套件陈旧断言与生产写入尚未验收的边界详见[部署记录](references/server-deployment.md)。
- [部署知识](../../knowledge/deployment/windows-to-linux-compose.md)已总结 Blog / WWW / Docs 的本地构建到服务器方法。
- 本会话本地隔离组合已停止，其他共享 Provider 保留；没有 Git 提交或推送。

## Progress

2026-09-08 用户明确确认“站点验收算成功了”，本部署 Work 完成。随后按用户要求将文档桌面左侧导航从 256px 调到 300px，已构建部署。正式 CLI Playwright 在 1440/1024 实测宽 300px、768/390 目录抽屉可打开/关闭，各宽度无横向溢出及 pageerror。未运行数据库迁移或改动后端。此前“生产写入尚未验收”的阶段记录由本次用户验收确认取代；不将用户确认冒充代理重新执行了所有后台旅程。

制品和证据位于 E:/tmp/yueli-media-preset-20260908：docs-sidebar-web.sha256、deploy-sidebar.py、sidebar-width-report.json、screenshots/sidebar-*。服务器镜像 yueli/docs-web:server-20260908-sidebar-1 healthy；API 保持 titles-1。无 Git 提交/推送。

## Next

None

## References

- [多语言标题与当前迁移](../2026-08-27-docs-language-versioning/references/collection-titles.md)

- [媒体 preset 更新与当前部署](../../../../asset/flightdeck/work/2026-09-08-media-preset-query/index.md)（跨仓路径见 Asset 仓工作记录）

- [背景](context.md)
- [部署与验收](references/server-deployment.md)

## 2026-09-11 引用机制更新
当前API为server-20260911-lifecycle-1，Web保留server-20260911-pay-1。原正文1338条引用与封面1条保持，注册声明revision 2；数据配置保留，真实运行及桌面/手机通过。详见[记录](deployment-references-20260911.md)。
