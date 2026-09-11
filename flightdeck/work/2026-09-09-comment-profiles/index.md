# 评论用户资料修复

## Goal

Docs 公开评论和后台审核显示 Identity 昵称及头像，继续使用 Foundation 共享评论组件。

## Status

Finished

## Current

Identity `/api/v1/users?ids=` 的正式合同为 `items`，Docs Go 客户端错误读取 `users`，导致两个评论入口均使用保存的 UserKey 回退。已修复批量解码；不改评论数据、不新增 DTO、不复制 UI。回归测试使用真实服务响应形状，修改前失败、修改后通过。Identity 单用户接口仍为 `user`，不可一并机械替换。

## Next

无。本次修复已部署；先前导入/同步 UI 和 0020 迁移仍保持本地状态。

## Evidence

- 用户确认 Gallery 后，Docs 后台切到共享 columns 分栏（作者正文、来源、状态、菜单）；状态包含已通过。真实本地后台验证菜单、状态、排序、390/1440 无溢出、标题工具贴右。
- Web 已更新 `server-20260909-columns-1`，API 保持 comments-1；独立 UI 固定候选、frozen 安装、typecheck/build 通过，线上评论 chunk SHA256 与构建一致、公开页双宽度 Playwright 通过。产物与记录 `E:/tmp/yueli-sites-sync-20260909/columns`。未部署尚在本地的导入/同步 UI 和 0020 迁移。

- 本地脚本与截图：`E:/tmp/yueli-docs-publish-20260909/comments-ui.mjs`。
- 服务端候选：`E:/tmp/yueli-sites-sync-20260909/comments-fix`，基于已上线 sync-1 快照仅修复 Identity 客户端，不携带尚在本地验收的导入/同步 UI 和 0020 迁移。
- 当前 Docs 本地 Session：`20260909T105743Z-52232`。
- 定向 Go identityclient/controller 测试通过。本地真实 Playwright 覆盖创建评论、公开/管理列表昵称、390/1440 布局；本地测试账号没有头像，未把它记作图片验收。上传本地头像的尝试返回 500，未扩大成账户媒体功能修复。
- 正式 `docs-api` 镜像 `server-20260909-comments-1`，健康。基于 sync-1 API 镜像仅替换应用二进制，无迁移、配置秘密或数据变更。Compose 备份 `/projects/yuelili.com/backups/comments-20260909`。
- 正式 Playwright 打开 `https://docs.yuelili.com/test/ces`，现有评论显示“月离”，头像 img.complete 且 naturalWidth > 0，390/1440 无横向溢出，无 pageerror。证据 `comments-fix/live.json`、`docs-production-*.png`。生产管理员会话不可用，后台通过本地真实登录/API 验收，未冒称线上管理员验收。
