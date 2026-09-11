# 背景

- 用户在 WWW 上线后明确要求继续部署 Docs。旧站不迁移，存档由用户本地保留。
- SSH `bt`，生产服务根目录 `/projects/yuelili.com`；1Panel OpenResty 管理正式域名证书与路由。
- 复用 Account、Identity、Asset 和现有 content PostgreSQL（含 zhparser）；Docs 使用独立数据库与稳定会话密钥。
- 不自动认领首位管理员，不启用生产 Dev Seed，不提交或推送 Git。
- 候选快照放在产品目录外，避免 Workspace 将其中 go.mod 误识别为重复模块。
- 相邻 Workspace 和 Provider 源码只读，本地进程由 Workspace CLI 管理。当前只选择 Docs 部署 Work，保留既有文档导入与语言版本 Work。
