# Docs 独立消费者迁移

## Goal

将 Platform 中的 Docs 产品历史无损迁入独立仓库，并让 Docs 只依赖正式发布的 Foundation、
Identity 与 Asset 合同；本地开发由 Workspace 统一编排，生产部署由本仓 Compose 独立负责。

## Status

Finished

## Current

独立仓库已合并 Platform 产品历史并保留旧公开仓库恢复标签。API 与 Web 已移除 Platform
源码和工作区包依赖，补齐产品自有 runtime、管理组件、部署绑定检查、错误目录、OpenAPI、
Dockerfile、三种 Compose 拓扑及本地 E2E 合同。OpenAPI 导出不再依赖数据库或部署配置。
CI 在迁移期只允许手动触发，避免提前执行逐产品测试。

## Next

把本次发布的精确 revision 锁入 Workspace，并从 Platform 退役旧副本。所有消费者迁移完成后
统一执行 API、Web、容器、Compose 与 Playwright 验收，再开启 CI 的 push/PR 自动触发。

## Invariants

- Docs 不导入 Platform 源码，不使用 `workspace:*`。
- Identity 是唯一身份生产者；Docs 仅消费 OIDC/JWKS 合同。
- Asset 可由 Docs 独立部署，也可复用现有实例；Docs 仅提交自身 profile 与资源语义。
- 本仓不使用 `doctor.yaml`，部署入口是标准 Docker Compose。
- 迁移批次不逐产品启动服务或运行测试，统一验收另行执行。
