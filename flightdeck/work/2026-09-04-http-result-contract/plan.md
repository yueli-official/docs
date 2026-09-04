# Plan

## P0 — 事实源与覆盖

- [x] 以 v1 catalog 为唯一事实源，旧目录改为单向兼容投影。
- [ ] 清点 canonical OpenAPI 全部 operation 的 success/status/errors 并建立覆盖门禁。

## P1 — 实现与前端

- [ ] 审计 typed cause、稳定公开参数和 raw message fallback。
- [ ] 接入生成物 freshness、统一 failure feedback 和 compatibility diff。

## P2 — 验证

- [ ] Go test/race/vet/govuln、Web tests/typecheck/build 与合同门禁全绿。
- [ ] Workspace 本地组合完成公开阅读、管理、导入错误和 201/204 Playwright。
