# Windows 本地构建部署到 Linux Compose

## 适用范围

Blog、WWW 和 Docs 已使用此路径向现有 1Panel / OpenResty 服务器部署：本地准备可审计源码快照与构建产物，经 SSH/SCP 传输，服务器组装并运行产品自己的 Docker Compose，再切换已有域名反向代理。它不是 git push 触发 CI，也不是把本地 node_modules、开发配置或数据库复制到生产。

这份知识由 Docs 维护，汇总已验证的方法；逐次镜像版本、目录、端口与验收记录留在所属 Work。共享 Provider 的发布与全局编排合同仍归所属仓库拥有。

## 构建输入与证据

1. 记录产品 Git SHA、未提交补丁、Go 依赖快照 SHA、前端包 SHA256、锁文件和运行时版本。Git SHA 本身不能代表有未提交改动的制品。
2. 候选源码快照放在仓库外。本轮使用 `E:/tmp/yueli-<site>-deploy-<date>`；放入产品子目录的多份 go.mod 会被 Workspace 递归发现，导致模块冲突。
3. 本地 Workspace 的 sibling overlay 只证明联合开发可用。服务器候选需要独立依赖：七个 @yueli 包在本轮锁为 `server.20260907.1` tgz，Go 指向固定 Foundation 快照；没有冒充正式 Release 或移动 Tag。
4. 内容编辑器依赖必须检查整个图。Tiptap 扩展统一到兼容版本，本轮为 3.31.3；独立版本的 y-tiptap 不能机械一起覆盖。
5. 运行产品 Go 检查、HTTP 合同生成与覆盖检查、Web typecheck / unit / production build，再执行真实 Playwright。旧综合测试含陈旧断言时记录明确失败，不将定向通过说成全套通过。

## Windows 产物到 Linux

- Go API 使用 `CGO_ENABLED=0 GOOS=linux GOARCH=amd64` 交叉编译；healthcheck、bootstrap、bindingcheck 命令随产品一起构建。运行镜像使用 distroless nonroot，COPY 可执行位为 755。
- Nuxt 交付 Nitro `.output`，使用匹配的 Node 24 Linux 运行镜像和非 root 用户；不上传开发 node_modules。
- Windows pnpm junction 不能原样归档。遍历 `.output` 时把 junction/symlink 转成产物内部相对 symlink，并校验目标仍在 `.output` 内。本轮不含 `.node` 原生模块；将来出现原生依赖时，必须在 Linux 构建相应制品，不能继续宣称 Windows 输出可直接移植。
- 将产物压缩为 tar.gz，生成 SHA256SUMS；SCP 后在服务器执行 `sha256sum -c`。成功上传不能代替校验。
- 基础镜像已缓存时可 `docker build --network=none --pull=false`，避免服务器再次安装依赖；镜像按部署版本显式标记，不用模糊的 latest。
- Windows 向 SSH bash 传脚本用 UTF-8 字节和 LF。Python 的 text=True stdin 在 Windows 可能转换为 CRLF，造成远端 shell 难以定位的语法错误。不要拼接含密钥的 shell 命令，也不要打印生产 env。

## 持久配置与服务边界

- 服务目录 `/projects/yuelili.com/<site>` 保存所属 Compose、权限 0600 的 .env、迁移和必要持久文件。临时归档在 `.deploy/<site-date>`，不能覆盖正式运行配置。
- 数据库和媒体归基础服务/数据库 Compose 持久化管理，应用替换不包含删除这些数据。Blog 与 Docs 复用 content PostgreSQL 的 zhparser 能力，但使用独立数据库、独立非超级用户；数据库不发布公网端口。
- 初装脚本与更新脚本分开。已有库、Client、Asset Registration 完成后不原样重跑初装；后续只部署变更制品并按需执行新增迁移。迁移原始字节不可修改。
- Web seal secret、数据库密码、OAuth 服务凭据稳定保留，不能每次部署重生成导致全部会话失效。
- Web 和 API 仅映射宿主机 127.0.0.1，OpenResty 接收公网 HTTPS；进程之间通过声明的 Docker 网络通信。不要把本地 LAN/localhost issuer 写入生产。

## Account、Asset 与管理员

- Identity issuer 使用正式 Account HTTPS 域名；产品 callback、logout origin 必须与正式域名完全匹配。产品 OAuth Client audience 同时覆盖产品及实际调用的 Identity/Asset，服务 Client 单独注册并留在服务器。
- Asset 消费者通过产品声明注册，再运行 bindingcheck 验证实际能力。不能手工推导对象 Key 或直接配置存储 Backend。
- Docs 导入声明 8000 万像素，现有 Asset 默认公开图片护栏为 3000 万会拒绝注册；应检查声明与生产 Provider guardrails 的差异，在已授权范围内明确对齐并保留其他限制，不绕过注册校验。本次修改记录在 Docs 部署 Work。
- 默认 `DOCS_BOOTSTRAP_ADMIN_SUBS=[]`。已登录用户访问 `/manage/setup`，点击“初始化站点并成为管理员”才能认领；第一个普通登录或 Identity 全局管理员都不会自动提权。显式 Bootstrap JSON 只用于无人值守或受控恢复。
- 代理不能为了验收代替用户认领生产站。一次性认领的成功、冲突与后台交互在隔离本地实例验证；生产未认领时只验证公开状态、匿名拒绝、登录跳转和页面。

## 代理切换与验收

启用了出站代理的服务新增私网调用时，同步维护 NO_PROXY 的精确 Docker 服务主机。Identity 的 PAT 权限目录使用 blog-api、docs-api，需将这两个主机加入已有列表；只启动健康并不能证明权限目录可达。实际用普通生产验收账号检查 unavailableSites、令牌创建/读取/撤销，避免漏掉只在真实 PostgreSQL 缺失记录时出现的错误映射。新来源密钥仅生成一次，保存在产品持久秘密配置，并补入部署后配置备份。

1. 验证迁移、默认空站设置、OIDC 与 Asset 绑定。生产不开 Dev Seed，也不把本地文档当成部署数据。
2. 启动所属产品 API/Web，等待健康并请求回环入口；诊断区分 API、BFF、上游 Provider、反向代理和浏览器。
3. 备份当前站点 OpenResty conf，保留证书和现有域名规则；更新根 location 指向 Web，先 `nginx -t` 后 reload。
4. 清理旧静态站特定匹配，例如 `location ~* \.txt$` 会截走 llms.txt；不要全局改其他域名。Docs 需要 100 MiB 上传入口及与 30 分钟导入任务匹配的超时。
5. 正式 HTTPS 使用 CLI Playwright 检查页面渲染、关键交互、390/1440px 和按需 light/dark、控制台错误、溢出、静态资源、OIDC origin/callback。已初始化后台还需三次刷新验证 SSR shell 和持久会话。
6. 图片上传必须走 init → PUT → finalize → 媒体交付才算通过，空站首页健康不代表图片链路通过。验收创建的临时内容只清理自己创建的对象。
7. 保存源码补丁、依赖清单、哈希、命令结果和截图；停止本会话通过 Workspace CLI 拥有的本地组合，保留其他会话/共享 Provider。

## 事实来源

- [Blog 服务器部署](../../../../blog/flightdeck/work/2026-09-04-http-result-contract/references/server-deployment.md)：候选包、Linux 制品、持久目录、正式登录/上传/发布证据。
- [WWW 服务器部署](../../../../www/flightdeck/work/2026-09-08-server-deployment/references/server-deployment.md)：相同部署路径与用户自行认领记录。
- [Docs 本次部署](../../work/2026-09-08-server-deployment/references/server-deployment.md)：具体服务器状态、制品、认领接入和验收限制。
- [Workspace 首位管理员合同](../../../../workspace/flightdeck/knowledge/authorization/initial-administrator-claim.md)。

本轮候选包未正式发布；正式 Release 安装、整站灾难恢复和持续 CI/CD 不属于本轮已完成结论。
