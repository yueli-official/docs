# 2026-09-09 线上部署

用户明确授权更新 Blog、Docs、WWW。本记录属于 Docs 项目文档发布 Work；其他产品使用各自原有部署 Work 留交付指针，没有创建全站 Work。

## 结果

Docs API/Web、Blog API/Web、WWW API/Web、Account Web、Asset 为 `server-20260909-sync-1`；Identity API 为 `server-20260909-sync-2`。九个服务均 healthy。Compose 位于服务器 `/projects/yuelili.com/{docs,blog,www,identity,asset}`，SSH 别名 bt。

本机证据 `E:/tmp/yueli-sites-sync-20260909`，服务器 `/projects/yuelili.com/.deploy/sites-sync-20260909`。`source-manifest.json`、Git 补丁、候选 package/lock、SHA256、构建/迁移/运行结果均保留。七个 JS 包是固定 `server.20260909.sync.1` 私有候选归档，Go 使用独立 Foundation 源码快照与 GOWORK=off；不是 GitHub 正式 Release，也不是直接部署本地运行 overlay。未提交或推送 Git。

五个相关数据库已 pg_dump 并验证 pg_restore --list；配置和迁移备份在 `/projects/yuelili.com/backups/sites-sync-20260909`，目录 0700。config-after.tar.gz 包含最终配置及新 Docs 同步密钥。保持原生产账号、内容、媒体、OIDC/seal/PAT 主密钥和其他服务配置。

仅追加 Identity 0035_pat_recovery、Docs 0019_project_doc_sources，迁移成功；服务器原有迁移逐文件 SHA256 保持不变。本地若干旧 SQL 的差异仅是行尾，未覆盖服务器执行过的字节。

## 运行接入

Identity `pat.applications` 注册 blog-main-web / docs-main-web，权限目录使用 Docker 私网 blog-api:8085 / docs-api:8086；对应 PAT allowHTTP 仅用于这组受信私网地址。Identity NO_PROXY 增加两个内部主机，防止凭证查询误经现有出站代理。Asset authorities 同样指向两个消费者的媒体授权接口。

Docs personalTokens.siteId=docs-main-web；独立随机 32 字节来源加密密钥保存在 Docs `.env` 的 DOCS_PROJECTDOCS_ENCRYPTION_KEY，Compose 注入 GF_DOCS_PROJECTDOCS_ENCRYPTIONKEY。没有把本地测试密钥复制到生产。项目同步入口：https://docs.yuelili.com/manage/import-sources 。

## 验证

- 五个后端相关包测试和 Linux amd64 构建通过；四个 Web 独立安装、冻结锁安装、Nuxt 类型检查与生产构建通过。
- 服务器所有制品 SHA256 校验、九个服务 healthy、旧迁移字节复核通过。
- CLI Playwright 正式 Blog/Docs/WWW 390/1440 首页均 200，无横向溢出或 pageerror；三站 OIDC 正式 Account 跳转及匿名受保护接口拒绝通过；WWW Keycrash release API 200。
- 使用此前生产验收账号，两个权限目录 unavailableSites=[]；临时 PAT 创建、加密恢复/no-store、用户信息读取、未选 Blog/Docs 权限拒绝、撤销后拒绝通过。测试令牌均已撤销，没有新建生产账号或修改产品角色。
- 首轮发现 GetPATByHash 在 PostgreSQL 未找到记录时返回 sql.ErrNoRows 导致 500。现映射为 found=false；已有真实 PG TestPGPATRoundTrip 先失败后通过，TestPGPAT 全部通过，线上撤销复验返回 401。修复源码在 Identity internal/dao/pg_pat.go，单独 Identity sync-2 制品已上线。
- 证据：acceptance.json（首页通过但首轮目录失败）、acceptance-auth.json（修复后全部授权检查通过）、remote-after.json、pat-pg-red.log、pat-pg-green.log 与截图。不要把首轮失败记录当最终结果。

## 验证边界与后续

正式验收账号不具备 Docs/Blog 站点管理权，未代用户修改生产文档或绑定 GitHub 来源。真实 GitHub docs.zip 正向联调仍待项目发布附件；已有受控 Release + PostgreSQL 成功测试不冒充外部生产同步。

检查日志时另见 Docs discovery 发布器 `source_order_violation`（记录排序不严格递增）。实际请求 https://docs.yuelili.com/sitemap.xml 返回 500。未改 discovery 代码，本次首页与所列接口验证不覆盖该发布器；后续需单独复现检查 PostgreSQL 排序与 canonical URL 游标的一致性。
