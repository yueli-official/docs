# 权限页用户资料

## Goal

Docs 权限申请与用户管理显示昵称、头像、用户主页和真实时间，修复 Identity 批量资料响应读取错误。

## Status

Finished

## Current
2026-09-09 授权角色标签统一：使用共享 AuthorizationGrantBadge，标签只显示角色、中文来源支持 hover/focus（含 initial_claim），删除产品内 grantSourceLabel。真实本地桌面/手机验收、独立 typecheck/build 通过。线上 Web `server-20260909-grants-1` healthy，正式组件 SHA256 验证通过。无授权/API/数据库修改。


用户指出 Docs/Blog 都只显示 UserKey。正式 Identity 返回 items（含 displayName=yueli 和 avatar），页面读取旧 users 导致空映射；Docs 申请行还直接显示 subject。已修正 items，使用 Nuxt UI UUser、Asset SDK 头像 URL、Account /u/{userKey} 主页。申请时间 createdAt；授权时间使用既有 ValidFrom 并标为“授权生效”，不伪称注册/加入时间。无数据库迁移。

已部署 API/Web `server-20260908-users-1`，容器健康；真实本地浏览器、生产只读依赖检查与构建/合同检查均通过。完整证据和夹具边界见验收记录。

## Next

None

## References

- [背景](context.md)

- [修复与验收](acceptance.md)
