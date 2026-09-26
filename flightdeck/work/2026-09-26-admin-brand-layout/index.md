# Docs 后台品牌布局

Status: Finished（已验收并部署 `server-20260926-admin-brand-2`）

## 交付

2026-09-26 用户确认 Commerce 方案后，授权推广本产品后台。顶部标题卡、真实统计、集合卡与紧凑工具栏沿用本站靛紫色。覆盖总览及文档、文档集、评论、导入、项目同步、首页设置、资源策略、权限与申请。公开页面和沉浸式编辑器不启用后台品牌主题。

- Foundation `AdminOverview` 共享排版与窄屏让位逻辑，`AdminMetricCard` 共享统计卡；产品通过 `artwork` 插槽提供独立的展开书页与索引书册 SVG，图案不放进公共库。
- 搜索/筛选移到集合内部，表头保持轻背景，保留原有权限、筛选、排序、分页和编辑流程。
- 窄屏按内容摘要组织，列表操作使用共享 `AdminRowActions` 的更多菜单。

## 验收证据

- 当前代码 `pnpm typecheck` 退出码 0。
- 共享 UI typecheck、141/141 单测、tarball 独立消费者安装/构建通过，包含两个新组件及主题 CSS。
- 三站接入共享组件后的生产构建均通过。随后本产品窄屏样式、菜单及文案微调通过最终类型检查、Nuxt 实页编译和 Playwright；未对最终微调重新生成生产制品。
- CLI Playwright 全后台检查 34 项通过；最终受影响列表复查 12 项通过。覆盖 320/390/768/1024/1440，深浅色偏好、主容器与表格溢出、Tabs、真实统计与装饰插槽、搜索/清空、筛选、更多菜单快速编辑/取消、公开页主题隔离。代表页 Axe WCAG AA 零违规；最终运行无捕获到的页面错误或水合告警。WWW 按主站设计始终保持黑银色。
- 证据：[完整报告](../../../web/test-results/admin-brand-20260926/report.json)、[最终列表复查](../../../web/test-results/admin-brand-20260926/report-list.json)、[桌面总览](../../../web/test-results/admin-brand-20260926/light-dashboard-1440.png)。生成物在 Git 忽略目录。

## 恢复与边界

## 追加复查

- 账户入口移到共享壳顶栏右侧；评论搜索与 Blog 一致地从集合左侧开始排列。
- Playwright 实测账户下拉可打开；390px 下账户按钮可见且页面无横向溢出，评论工具栏从集合左侧起排。Foundation 与三站类型检查通过。

- 预览：http://docs.dev.yuelili.test:3003/manage
- Workspace Session：`20260926T051755Z-38308`；最后 `dev status --check` 为 ready。保留供用户预览。
- 使用 Workspace 显式 `--local foundation` 的依赖覆盖验证新共享接口；生产站点使用固定预览 tarball 随 Web 镜像交付，未发布 npm 包。
- 不修改 Workspace 合同或顶层 Flightdeck，不停其他产品进程；Commerce 原预览保留。

## 2026-09-26 正式部署

Docs Web 已部署 `yueli/docs-web:server-20260926-admin-brand-2`，容器 healthy、RestartCount=0；API 未变。修正版将 Tabler 图标集合纳入 Nitro server bundle，正式 `/api/_nuxt_icon/tabler.json?icons=cube` 返回 200。正式首页及 `/manage` 登录跳转通过 Playwright，桌面/390px 无横向溢出或 pageerror。`/robots.txt` 仍由未修改的 Docs API 发现服务返回 500；没有使用生产管理员会话验证写操作。

文档列表按容器宽度收起次要列，保留标题、状态、文档集、质量提示和更多操作，修复固定操作列在手机上挤出可视区；缺少摘要时不再重复显示“暂无摘要”。初次 dev 预编译期间遇到网络中断/水合告警，稳定启动后的完整与最终复查均通过。

## 2026-09-26 文档集分页收口

分页移动至文档集同一主卡片内，仅保留顶部分隔线；修复开发预览裁剪依赖预编译，生产类型检查与构建通过。真实本地 Playwright 验证 1440/768/390 无横向溢出，分页父级属于 data-admin-collection；每页 24 切为 12 后显示两页。

已部署 yueli/docs-web:server-20260926-admin-brand-3，healthy，restart=0。发布包 SHA256 fc9fc3aa57846429efbc5c29930c0588c1eee45882f1d2d036d2147ca8e1e5f9。旁路候选健康与 Tabler 图标通过；备份在 /projects/yuelili.com/backups/docs-admin-brand-3-20260926，正式域名最终浏览器复查进行中。

正式域名 Playwright 最终复查通过：桌面 1440 / 手机 390 无页面横向溢出、无 pageerror，healthz=200，后台入口进入统一账户登录。生产未使用管理员会话执行写操作。
