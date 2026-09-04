# 统一错误与 HTTP Result 合同

## Goal

让 Docs API 与 Web 完整采用 Foundation 声明式错误目录、operation manifest、直接成功 DTO 和统一失败反馈，并通过本地 Identity/Asset 组合完成真实浏览器验收。

## Status

Open

## Current

Docs 的 13 个业务错误已由 Foundation v1 catalog 单向生成 Go/TypeScript/i18n，旧目录仅作兼容投影；70/70 operation success 已覆盖，其中 66 个声明业务 errors，4 个纯公开读取端点仅使用 Foundation 通用失败。前端 raw message fallback 已清零，导入批次与 Front Matter issue 不再保存或展示内部错误文本。

## Next

依据 [执行计划](plan.md) 修正 canonical OpenAPI 中仍错误保留的创建/删除 200：为真实创建写 201，为无正文操作写 204，并补运行时状态测试；不发布版本。

## Progress

- v1 catalog 已成为唯一事实源，旧目录和生成 Go/TypeScript/i18n freshness 通过。
- 公开动态错误参数改为命名类型；已知 `err.Error()` 与上游响应文本不再进入 Problem。
- operation manifest 已自动生成 70/70，Foundation 校验、覆盖门禁、Go 全量测试/vet 与 Web 53 tests/typecheck 通过。
- 授权域 operation errors 已声明；门禁会拒绝 catalog error 无消费者或错误表残留失效路由。
- 66/70 operation 已声明业务 errors；Web 统一 failure resolver、54 项测试、typecheck/build 与 raw message 扫描通过。

## References

- [稳定上下文](context.md)
- [Foundation HTTP Result Contract](../../../../foundation/flightdeck/knowledge/errors/http-result-contract.md)
- [Foundation Error Catalog](../../../../foundation/flightdeck/knowledge/errors/error-catalog.md)
