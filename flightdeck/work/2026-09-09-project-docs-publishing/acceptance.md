# 本地验收记录

2026-09-09。实现已完成并通过代理本地验收，等待用户体验及真实项目 docs.zip 联调。没有提交、发布依赖或部署生产。

## 实际通过

- Docs / Foundation 相关 Go 测试：controller、docsauthz、catalog、dao、assetclient、server、importkit、projectdocs 和共享 auth。
- `projectdocs.TestWorkerPostgresLifecycle` 用真实 PostgreSQL、独立临时 schema 和受控 Release 数据验证首次导入、同附件免下载、第二版更新不重复、自定义 slug / 根 index、损坏 ZIP 和下载失败保留、最终提交前撤销回滚、并发 advisory lock、重启恢复已提交批次、撤销暂停。测试 schema 按本次生成名称清理，没有重置现有产品库。
- 旧 catalog 的 live-PG 测试默认跳过，未启用会重置固定 docs 数据库的旧测试入口。不能把普通 Go 测试结果当作旧 PG 测试已执行。
- 真实 Identity / Asset / Docs 与 CLI Playwright：带图 PAT 上传/确认、重复 updates=1 creates=0、BFF 透传、PAT 撤销、未开放路由拒绝。
- 真实文档 CRUD：读取/创建/编辑/发布/下架可按 scope 分配；编辑不能附带发布，发布不能编辑正文，永久删除不对 PAT 开放。内存授权模块另测集合 A 的作者不能写集合 B、撤销角色后立即拒绝；目录能发现集合级能力。
- 来源 UI：绑定、暂停、恢复、运行记录、解除绑定、凭证不出现在返回 DTO、PAT 不能管理来源；1440px / 390px 渲染和交互通过，无页面异常与横向溢出。移动端主按钮与标题对齐，目标选项显示文档集版本。
- 独立 Python 标准库发布脚本成功发布带图文档包；3 个脚本测试通过，Skill 校验通过，已安装。
- Web typecheck 通过。HTTP 声明/错误/DTO 生成与检查通过，共 80 operations。

## GitHub 真实性边界

PowerShell 代理出口曾返回匿名额度不足；正式 Go 客户端实际读取到 Keycrash 最新正式 Release，但没有 docs.zip 附件。缺附件路径已在真实本地后台观察到；成功附件下载→发布、二次版本、并发与恢复使用受控 Release + 真实数据库验证。未伪造真实 GitHub 正向附件验收，也未修改任何远程 Release。

## 证据

本机目录 `E:/tmp/yueli-docs-publish-20260909`：

- `live.mjs` / `live.json`：带图 PAT 闭环。
- `crud.mjs` / `crud.json`：文档 scope API。
- `source-ui.mjs` / `source-ui.json`、`source-form-desktop.png`、`source-desktop.png`、`source-mobile.png`、`source-history-mobile.png`：来源 UI。
- `skill-live.jsonl`：独立发布脚本输出。

测试 PAT 均撤销，来源 UI 测试绑定已解除。保留测试文档集和已导入样本供查看。部分早期失败脚本输出也在临时目录，最终通过结果以以上 JSON 为准。

## 当前运行

Workspace CLI 的 Docs 独立 session `20260909T084527Z-26152`，显式 local docs / foundation / identity / asset。

- Docs：`http://docs.dev.yuelili.test:3003/manage/import-sources`
- 手动导入：`http://docs.dev.yuelili.test:3003/manage/import`
- 令牌：`http://account-docs.dev.yuelili.test:3601/developer-tokens`
- 回环 API：Docs 8086，Identity 8683，Asset 8684。

本地同步密钥保存在会话临时目录 `sync-key.local.txt`，通过 `GF_DOCS_PROJECTDOCS_ENCRYPTIONKEY` 注入；不能提交该文件或把它当作正式部署密钥。重启仍须经 Workspace CLI，并复用同一密钥才能读取本地来源凭证。所有其他本地站点和共享 Provider 未停止。

## 导入与同步导航、自动检查配置

2026-09-09 本地完成：共享 ImportNavigation 展示导入/项目同步两个路由 Tab；移除页头手动导入链接，导航归为导入与同步；outline 图标操作、自动检查开关与创建表单设置。0020 为来源表追加 auto_check BOOLEAN，旧来源保留自动检查。新建关闭时 idle，不运行首次抓取；PATCH enabled/autoCheck 均可独立省略，自动检查不等于暂停。

真实 PostgreSQL TestWorkerPostgresLifecycle 验证手动源不被轮询、显式排队执行、执行后不重复轮询、开启自动立即到期、偏好持久化且不误暂停；原同步/回滚/恢复测试仍通过。Controller 测试、80 操作合同生成、Nuxt typecheck 通过。CLI Playwright 验证 Tab URL、旧链接消失、新建手动无运行记录、刷新保留、开关/暂停/恢复/检查/记录/解除绑定、手机390无溢出和无pageerror。视觉已查看1440/390截图；机械扫描无问题。临时来源和PAT已清理，测试文档集保留。session 20260909T102415Z-32708；无运行中的构建或验收命令。尚未部署本轮更改。
