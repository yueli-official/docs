# Docs 服务器部署记录

## 当前结果

2026-09-08 03:07（Asia/Shanghai）新版已切换至 https://docs.yuelili.com。用户明确旧版已删除且本地留档，本次是空站部署，没有迁移旧文档或本地数据库。API / Web 均 healthy。

用户已选择与 Blog 相同的首位管理员认领：访问 https://docs.yuelili.com/manage/setup，真实登录后显式点击“初始化站点并成为管理员”。生产仍为 `claimed=false, canClaim=true`，没有预设 Bootstrap subject，代理没有代认领或授予全局权限。

## 源码与制品

- 产品基础 Git SHA：`10d5a2ff2c0401b18132c556c16ebeb5614b4ee4`，当前本地修改未提交、未推送。
- 本轮修改：四个页面补显式 import；首页失效封面扫描移到 mounted 后；Docs 接入 Foundation 一次性认领、状态/命令 API、SSR 管理入口 gate、独立响应式 setup 页、声明式 409 与生成合同；部署默认 Bootstrap 为空。
- Foundation Go 归档 SHA：`28a4f092d2c0dbeb12e65a908b9593a7aca30dc8`，位于 `E:/tmp/yueli-www-deploy-20260908/foundation/go`。
- 前端复用 Blog/WWW 七个 `server.20260907.1` 候选包；独立 lock 统一 Tiptap 3.31.3，保留独立版本 y-tiptap。候选不是正式 Release。
- 本地快照与证据：`E:/tmp/yueli-docs-deploy-20260908`，`source-manifest.json`、`source-fixes.patch`、`source/`、vendor tgz、日志和截图。候选源码放在仓库外，避免 Workspace 递归识别重复 Go modules。
- 最终 API：`yueli/docs-api:server-20260908-2`，镜像 SHA `9c55200306a2de279abd730e4b26af8f6695df2564b4e9dad6b7fb2a70c81c77`。
- 最终 Web：`yueli/docs-web:server-20260908-4`，镜像 SHA `c67b7350a5df1577ef82afbb476bb09db22b33af779af4f1567412027871c4a3`。此前 Web 1/2/3 都不是最终交付版本。
- 远端归档：`/projects/yuelili.com/.deploy/docs-20260908`；最终 `api.tar.gz` SHA256 `4cb760d7788cc11b2d98438f2fa0e0972c339070b2d568bb8fba3a5afea6b8da`，`web-final.tar.gz` SHA256 `dec3e909e5986441970901645e9c25da683be24aa736d101323efde28a28128d`，上传后校验通过。
- Nitro 归档将 Windows junction 转成 `.output` 内部相对 symlink，拒绝 Windows 原生 `.node` 模块。Go 为 Linux amd64 / CGO=0；镜像非 root，服务器使用已有基础镜像离线组装。

## 服务器配置

- SSH 别名 `bt`。产品目录 `/projects/yuelili.com/docs`，独立 Compose，复用 `yueli-services`、`yueli-content-db` 网络。
- content PostgreSQL 中独立 `docs` 数据库及 `docs_app` 非超级用户，pgcrypto/zhparser 已安装。迁移保持原字节，生产未开启 Dev Seed，bootstrap 仅建立默认首页配置。
- `.env` 权限 0600，数据库密码、Web seal secret、服务 Client secret 稳定保留；`DOCS_BOOTSTRAP_ADMIN_SUBS=[]`。
- API 宿主机 `127.0.0.1:18086 → 8086`；Web `127.0.0.1:13003 → 3003`，无直接公网端口。
- OIDC `docs-main-web`，callback `https://docs.yuelili.com/auth/callback`，audiences docs-main-web / identity-api / asset-api。服务 Client `docs-asset-svc` 的 secret 留在服务器。
- Asset `consumer=docs namespace=docs revision=1` 注册完成，实际声明与候选一致，bindingcheck 通过。
- 现有 Asset 默认公开图片护栏 3000 万像素拒绝 Docs 导入 8000 万合同。已把生产 public-image.maxPixels 对齐 8000 万，保留其他默认 MIME、大小、存储类别和后端映射；备份为远端 `asset-compose-before-guardrail.yaml`（0600）。只重建 Asset 后 healthy，未修改 Provider 仓库源码。
- OpenResty 原证书与域名规则保留，根路径代理至 13003，100 MiB 上传限制及 1900 秒超时。旧静态 `.txt` location 已移除，`nginx -t` 成功后 reload。
- Nginx 切换前备份在 `before-activation/docs.yuelili.com.conf`，Compose 备份为 `compose-before-claim.yaml`。
- `activate.py` 已执行，保留空管理员列表。禁止继续使用旧 `--admin-sub` 指令或原样重跑初装脚本。日常更新使用当前产品 Compose 和明确的新镜像版本。

## 验证

### 构建与合同

- 当前候选 `go test ./...` 全部通过；Linux 四个命令构建通过。
- 公开 HTTP 合同共 72 operations，覆盖与生成漂移检查通过。
- Web 最终 typecheck、54/54 单测、production build 通过。
- Go 单测验证：普通登录/Identity 管理员角色不触发认领；匿名和服务主体不能认领；20 个不同用户并发只有一位赢家；同赢家重试幂等；后续用户冲突；Bootstrap 关闭入口。
- Foundation 实际 PostgreSQL 测试使用临时 schema，验证两个 Adapter 并发、重启保持状态、失去管理员后不重新开放。不是只跑默认 skip 的数据库测试。

### 本地浏览器

- `claim-local-report.json`：15 项通过。真实 Account 登录保持未认领；SSR 直接进入 setup，不先输出后台壳；390/768/1024/1440 × light/dark，无溢出、pageerror、hydration error，Axe serious/critical 为 0。
- 显式点击成功后进入后台，三次完整刷新 HTML 都含 `data-admin-shell`；同赢家重试 `created=false`。
- 冲突恢复 UI 使用受控 409 响应验证；真实并发授权结果由上述 Go / PostgreSQL 测试验证，不混称两者。
- `claim-relogin.json`：新浏览器会话重新登录，仍有管理员权限，setup 自动进入后台。
- 本地验收通过环境变量使用独立授权实例 `docs:docs-claim-acceptance-20260908`，没有清空本地原有管理员或文档；生产实例是 `docs:docs-main`。
- 原部署 Smoke 18 项已通过，覆盖公开页/后台、默认 zh-CN + 1.0.0 文档集创建 201 和删除 204；原公开页/管理员定向回归 2/2 通过。
- 旧综合测试有陈旧 dashboard 文案、编辑器几何断言、缺少现行 defaultLocale/semanticVersion 创建参数；不能声明旧综合套件全通过。

### 正式 HTTPS

- `production-report.json`：13 项通过，包括未认领状态、匿名 POST 401、桌面/手机首页和文档集、真实导航、healthz / robots.txt / sitemap.xml、正确 OIDC issuer / client / callback，以及管理入口实际进入 Account 登录页。
- 页面无 pageerror / 水合错误 / 横向溢出，Axe serious/critical 为 0；截图位于 `screenshots/production`。
- Docs 当前未配置 llms Publication，`/llms.txt` 返回产品 404；不是静态 Nginx location 截走，也未把未实现端点算作成功发布。
- `server-final-status.log`：两容器 healthy，migration no change，默认记录就绪，bindings compatible，认领状态仍 false。

## 交付边界与继续入口

- 本会话的本地隔离组合已由 Workspace CLI `dev down docs` 停止，共 5 个进程；原共享 Identity/Asset/Blog 未停止。前台 supervisor 被正常停止时外层 PowerShell 返回 1，不是产品验收失败。
- 用户亲自认领生产后，再验证正式后台会话、封面上传与文档创建/发布；当前本地通过不能代替尚未执行的生产写入链路。无需再问管理员方式，不要代理代认领。
- [可复用部署知识](../../../knowledge/deployment/windows-to-linux-compose.md)汇总 Blog、WWW 与 Docs 的方法及限制。


## 2026-09-08 权限页资料更新

当前本站 API/Web 均为 `server-20260908-users-1`，容器 healthy；包含此前媒体 preset、Docs 多语言标题及侧栏更新。权限页面补齐用户资料与时间，无数据库迁移。制品与部署校验保存在 `E:/tmp/yueli-media-preset-20260908`。
