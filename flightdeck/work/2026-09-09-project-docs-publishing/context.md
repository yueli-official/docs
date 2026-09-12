# 上下文

用户已验收 Identity/Blog 的个人令牌能力，包括分类权限卡片、防抖搜索、分组摘要、hover card 和列表复制；不再提供查看按钮或创建后的原文弹窗。下一消费者为 Docs。

用户希望文档唯一维护位置是项目仓库。认可 Markdown → Release docs.zip → Docs 自动同步；上传 Skill 给其他项目复用，不依赖本机源码 checkout。Skill 辅助整理/校验/发布，持续更新交给确定性自动化。

现有 Docs Import Package v1 支持 Markdown、相对图片、目录、多语言和可选 docs.json，使用稳定 source_id 关联重复导入。现有导入模式包含 upsert 与 replace-version，后者归档未包含页面。代码、正式格式文档与旧 Work 对 Front Matter 的描述存在时间差，实现前以当前 parser/测试为证据核对，不照抄旧摘要。

GitHub 集成第一版限定公开仓库的正式 Release 附件，默认 docs.zip；预发布和私有仓库不自动纳入。默认 upsert；完整镜像下架必须在绑定中明确选择，且目标版本归该来源管理。允许配置目标版本，不把软件 tag 自动等同于 Docs 版本。未经要求不提交、发布或改动其他项目的 GitHub 配置。

Workspace 合同和顶层 Flightdeck 保持只读；本 Work 属于 Docs 单站产品能力。已有 Foundation PAT 共享认证与 Asset 受限上传协议可复用。所有进程通过 Workspace CLI，本地验收必须使用 CLI Playwright。

2026-09-12 扩展范围：用户希望自定义投稿可以从本地脚本完成文档集创建和整个内容发布流程；隔壁会话负责 Blog 权限，禁止修改它的 checkout 或运行组合。新增令牌入口复用 Docs 现有账号/资源能力，每次操作取令牌勾选权限与账号当前权限的交集。本轮本地交付，不部署。
