# 项目文档发布与 GitHub 同步

Status: Open

## Goal
项目在源仓库维护 Markdown，通过开发者令牌与可复用 Skill 发布到 Docs，并支持绑定 GitHub Release 文档包、一键导入和自动更新。

## Current
2026-09-11 引用登记修复已上线：Docs API `server-20260911-references-1`；393张导入图片补齐1338条引用。当前内容每15秒自动对账，覆盖GitHub同步、导入、编辑、克隆、删除/恢复与回滚，归档保留引用。见[Asset交付与验证](../../../../asset/flightdeck/work/2026-09-11-reference-reconciliation/delivery.md)。本轮未提交/推送。

2026-09-11 已同步正式部署，见[本轮记录](deployment-20260911.md)。此前“仅本地/未部署”描述是历史阶段。

最新用户反馈已修正：ImportNavigation 只适配路由，内部直接复用站点设置同款 @yueli/ui/admin TabbedSurface，导航和正文合为共享卡片；删除页面顶部长说明。Nuxt typecheck 与真实桌面/手机切换及操作复验通过，仍仅本地。

用户反馈后的本地 UI/调度配置已完成：导入与项目同步并列 Tab，检查/暂停/记录使用 outline 图标按钮；每源 autoCheck 可配置，关闭时只允许手动检查，新建关闭不会立即抓取。新增 0020 单 boolean 列迁移已本地应用，未部署这轮改动。真实 PG 调度/持久化测试、Nuxt typecheck、80 项 HTTP 合同、桌面手机 Playwright 通过。证据 sync-ui.json / sync-desktop.png / sync-mobile.png 位于 E:/tmp/yueli-docs-publish-20260909。
2026-09-09 用户授权的生产更新已完成。Docs/Blog/WWW API 与 Web、Account Web、Asset 为 server-20260909-sync-1，Identity 为 sync-2；新增迁移、配置、备份、构建及线上页面/权限目录/令牌验证已完成，九个服务 healthy。交付与限制见[部署记录](deployment.md)。未提交 Git、未发布 GitHub 正式 Release。

本地实现和代理验收完成：Docs 六类范围能力、带图导入、独立发布 Skill、GitHub 来源页面/接口、加密委托、持久化调度、去重/恢复/失败保留。新增两张同步表的 0019 迁移已用于本地及生产。

当前 session `20260909T102415Z-32708`。无待等待命令。页面 `http://docs.dev.yuelili.test:3003/manage/import-sources`，令牌页 `http://account-docs.dev.yuelili.test:3601/developer-tokens`。Skill 安装在 `C:/Users/yl/.codex/skills/yueli-docs-publish`，源在本仓 skills 同名目录。

## Next
本轮布局先在[本地项目同步](http://docs.dev.yuelili.test:3003/manage/import-sources)验收；这轮 0020 和页面改动尚未上线。上一轮正式发布状态见 deployment.md。Keycrash 尚无 docs.zip，真实外部附件正向联调等待项目发布。日志额外发现的 discovery 排序问题与验证边界见[部署记录](deployment.md)。

## References
- [执行计划](plan.md)
- [验收记录](acceptance.md)
- [接入知识](../../knowledge/project-documentation-publishing.md)
- [上下文](context.md)
- [包格式](../../../docs/import-package.md)
