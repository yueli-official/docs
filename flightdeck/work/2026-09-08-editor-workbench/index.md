# docs 内容编辑工作台

Status: Finished

## 已完成

- 接入共享 EditorCommandBar：沉浸、公开页、设置、发布/下架、保存；窄屏标题与操作分行。
- 列表提供公开页、快速编辑、完整编辑与双击编辑；选择框与按钮不会误触发导航。
- 发布前保存当前内容；下架意味着匿名原链接不可读，不提供仅链接可见选项。
- 本地 Playwright 覆盖 390/768/1024/1440px、浅深色、沉浸/设置、快速编辑、双击、上下架实际 200→404→200。
- 类型检查和构建通过；[浏览器证据](references/browser-local.json)。

## 交付边界

独立候选构建完成，线上 server-20260908-editor-3 健康；公开页面 390/1440px 复查通过。生产管理员编辑交互未冒充验证，本地已完整验证。

新建文档支持直接发布；创建后刷新文档树再跳转语义编辑路径，避免空白编辑页。

[线上公开页面复查](references/browser-production.json) · [制品与源文件校验](references/candidate.json)。

## 完整适配补齐

Docs 发布入口固定显示，加载期间禁用写入，详情加载失败明确提示。新建直接发布、已有文档下架 404 与再发布 200、390/1440px 命令栏均通过；模拟详情加载失败验证发布仍可见且保存/发布禁用。

类型检查、独立候选构建通过，修正版 editor-2 已部署。线上公开页面浏览器复查通过；线上管理员操作未代替用户登录验证。

[本地补齐验收](references/browser-local-fix.json) · [线上复查](references/browser-production-fix.json) · [修正版制品](references/candidate-fix.json)。

共享构建规则已补入 [编辑工作台知识](../../../../foundation/flightdeck/knowledge/frontend/editor-workbench.md)：正文、元信息设置与命令栏必须一起适配。

## 大集合与完全沉浸修复

列表读取先使用集合级 QueryAll 权限计划；权限不能覆盖整集时继续逐篇判断。真实本地 1119 篇目录从 10.09 秒后 504 恢复为 175ms 的 200。实际语义编辑页加载通过，只请求一次正确 zh-CN 目录；语言和版本配置就绪后才加载树，树请求失败显示错误。新增授权回归测试覆盖整集授权与未授权访问。

修正版 editor-3 已上线且 healthy，正式类型检查、构建及线上公开页浏览器复查通过，无数据库迁移。线上管理员完整交互未冒充验证，本地大集合/沉浸交互已验证。

[回归证据](references/browser-timeout-immersion.json) · [制品](references/candidate-editor-3.json) · [线上公开页](references/browser-production-editor-3.json)。

## 按需父文档搜索（当前交付）

编辑器不再调用全量文档列表、构建整集树或依赖树定位文档。管理查询新增可选 ID/精确路径/排除子树条件；精确定位使用目标文档权限检查，分页候选延续现有权限计划。打开编辑页仅定位当前文档并读取详情。

父文档选择器打开时请求 20 条轻量候选（另有固定无父文档选项），关键词防抖后服务端搜索；请求序号阻止旧响应覆盖新结果。选中项独立保存标题和路径，服务端排除当前文档及子树。保存后按目标文档 ID 查回新路径，不刷新整集树。

本地真实 1119 篇文档集验收通过：不请求全量目录、首批 20 条、搜索首批之外的文档、排除子树、手机/桌面、修改父文档后保存与刷新，以及新建直接发布。Go 全量测试、类型检查、构建与 HTTP 合同检查通过。Docs API/Web `server-20260909-parent-1` 已上线并 healthy，线上公开页检查通过；线上管理员交互没有冒充验证。无数据库迁移。

[搜索验收](references/docs-parent-search-browser.json) · [保存验收](references/docs-parent-save-browser.json) · [制品](references/candidate-parent-search.json)。规则已写入 Foundation [编辑工作台知识](../../../../foundation/flightdeck/knowledge/frontend/editor-workbench.md)。

Next: None

## 图片输入与文档导入（2026-09-09）

共享 ContentEditor 已支持图片上传框、多图拖拽/粘贴，以及 Markdown、HTML、DOCX、ZIP 和富文本剪贴板导入。图片走原有 Asset Adapter，完成后一次插入，缺图失败不覆盖原文。本站前端已更新 server-20260909-import-1；类型检查、生产构建、线上公开页与编辑器静态资源检查通过，完整真实上传/导入交互在 WWW 本地完成。没有冒充本站线上管理员写操作验收。

[共享实现与验收记录](../../../../foundation/flightdeck/work/2026-09-09-editor-import/index.md)。

## 共享编辑器依赖固定（2026-09-09 完成）
package.json 与 pnpm 锁固定 `content-nuxt 0.2.3-server.20260909.placeholder.1`，将已验证包放入 web/vendor；五个站点的编辑器包 SHA-256 一致。配套 UI、Asset、HTTP 与会话依赖固定为已经部署验收的 WWW 组合，Tiptap 全组 3.31.3，避免旧正式包缺少当前消费接口。此举不代表发布 npm/GitHub Release。
该编辑器已在本轮之前部署到线上，Web 为 server-20260909-placeholder-1；此次把仓库依赖同步到同一编辑器制品，没有重复部署。其正式构建与线上检查证据归 Foundation 编辑器导入 Work。
依赖审计：E:/tmp/yueli-editor-consumers-20260909/dependencies.json。保留其他未提交改动；本轮未提交或推送 Git。
