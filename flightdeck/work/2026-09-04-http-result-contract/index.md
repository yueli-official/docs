# 统一错误与 HTTP Result 合同

## Goal

让 Docs API 与 Web 完整采用 Foundation 声明式错误目录、operation manifest、直接成功 DTO 和统一失败反馈，并通过本地 Identity/Asset 组合完成真实浏览器验收。

## Status

Open

## Current

Docs 的 13 个业务错误已由 Foundation v1 catalog 单向生成 Go/TypeScript/i18n，旧目录仅作兼容投影；`invalid_input`、`import_blocked` 与 `upstream_failed` 已改用稳定 reason/dependency，已移除导入、文档树、授权、流量和 Asset seam 的原始错误文本。canonical OpenAPI 有 70 个 operation，当前 manifest 仍仅覆盖 7 个。

## Next

依据 [执行计划](plan.md) 将 operation manifest 扩展到 canonical OpenAPI 的 70/70，并补齐 coverage/freshness/compatibility CI；不发布版本。

## Progress

- v1 catalog 已成为唯一事实源，旧目录和生成 Go/TypeScript/i18n freshness 通过。
- 公开动态错误参数改为命名类型；已知 `err.Error()` 与上游响应文本不再进入 Problem。

## References

- [稳定上下文](context.md)
- [Foundation HTTP Result Contract](../../../../foundation/flightdeck/knowledge/errors/http-result-contract.md)
- [Foundation Error Catalog](../../../../foundation/flightdeck/knowledge/errors/error-catalog.md)
