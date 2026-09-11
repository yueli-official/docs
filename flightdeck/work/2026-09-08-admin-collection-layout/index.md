# docs 后台集合布局

## Goal
将本站管理集合接入用户确认的共享分页、网格、评论和标题工具规则，保留权限与查询行为。

## Status
Finished

## Current
已接入共享分页与标题工具区。CLI Playwright 已验证 /manage/collections, /manage/comments, /manage/docs, /manage/import，覆盖 390/1440 宽度，无 pageerror。admin.4 独立候选构建和类型检查通过。 Web 已部署 server-20260908-admin-4，容器健康；正式公开页面 390/1440 检查通过。生产后台管理员交互由本地权限夹具覆盖，未使用真实管理员会话代替验收。

## Next
None

## References
- [共享规则](../../../../foundation/flightdeck/knowledge/frontend/compact-admin-collections.md)
- 本次浏览器证据：`E:/tmp/yueli-media-preset-20260908/docs-all-admin-browser.json` 与 screenshots。
- 本地 URL Catalog 从旧 IP Origin 精确迁移到声明的开发域名：先比对旧 version/digest，事务更新 catalog digest/revision，再由新 Catalog 构造器确认；历史 URL 保留，未修改 SQL 迁移文件。备份在本次外部验收目录。

## Final verification
- 最终生产构建通过：`docs-admin-4-build.log`；本轮仅推广后台布局，未重新验收支付等无关业务。
- 类型检查、CLI Playwright 页面与相关交互、改动文件空白检查通过。额外验收组合已通过 Workspace CLI 停止，原有共享服务与 Gallery / Blog 保留。
- [页面验收结果](references/browser-layout.json)；截图及交互日志保存在本轮外部制品目录。
