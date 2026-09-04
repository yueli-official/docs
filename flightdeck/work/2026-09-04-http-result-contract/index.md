# 统一错误与 HTTP Result 合同

## Goal

让 Docs API 与 Web 完整采用 Foundation 声明式错误目录、operation manifest、直接成功 DTO 和统一失败反馈，并通过本地 Identity/Asset 组合完成真实浏览器验收。

## Status

Open

## Current

Docs 的 13 个业务错误、70/70 operation 和 66 个业务 errors 已迁入 Foundation Project v1：`project.json` 声明 producer/输出，`operation-errors.json` 拥有路由错误事实。本地 Foundation CLI generate 与 `-check` 均通过，不需要发布。成功结果为 200×54、201×11、202×1、204×4；前端和导入 DTO 不泄漏内部错误文本。

## Next

通过 Workspace local checkout 重建 Docs 组合，完成公开阅读、管理、导入错误及 201/202/204 Playwright；发布 Foundation 前保留旧生成命令供远端 CI 使用，不发布版本。

## Progress

- v1 catalog 已成为唯一事实源，旧目录和生成 Go/TypeScript/i18n freshness 通过。
- 公开动态错误参数改为命名类型；已知 `err.Error()` 与上游响应文本不再进入 Problem。
- operation manifest 已自动生成 70/70，Foundation 校验、覆盖门禁、Go 全量测试/vet 与 Web 53 tests/typecheck 通过。
- 授权域 operation errors 已声明；门禁会拒绝 catalog error 无消费者或错误表残留失效路由。
- 66/70 operation 已声明业务 errors；Web 统一 failure resolver、54 项测试、typecheck/build 与 raw message 扫描通过。
- 明确创建与无正文删除已在 DTO 和 controller 同步状态；Idempotency-Key 角色申请保留 200。Go test/race/vet/govulncheck 全绿，gRPC 已升级到修复版本。
- Foundation 本地固定 CLI 已从旧 manifest 提取 operation-errors，并用 Project v1 对全部六类产物完成 generate/check；旧命令仅作为 v0.4.1 远端 CI 兼容入口暂留。

## References

- [稳定上下文](context.md)
- [Foundation HTTP Result Contract](../../../../foundation/flightdeck/knowledge/errors/http-result-contract.md)
- [Foundation Error Catalog](../../../../foundation/flightdeck/knowledge/errors/error-catalog.md)
