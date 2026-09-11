# Docs 本地资产引用与独立库存修复

## Status
Finished

## Goal
当前本地文档使用的媒体在独立 Asset 库可交付并完整登记；保持导入、同步、编辑和删除/恢复共享事实源。

## Current
2026-09-11 当前源码组合 20260911T065120Z-38676 已恢复原独立端口。保留原 ID 复制 1424 份旧共享 Docs 库存，原有 1 份素材保留；1425 个文件逐一校验并放入独立 .data/docs-assets，避免跨数据库删除保护互相绕过，原共享记录/文件未改。
已补齐机器对账配置，313 份当前使用图片登记 1170 条正文引用。CLI Playwright 桌面/移动端引用详情、Account 缩略图实际解码（240×162）和 Docs 同源媒体 200 通过。旧导入副本无当前使用，正确保持未引用。
两处旧集合封面共用一个失效 ID 01a0380d-fa36-7bfa-8e11-0bd3692ed56c，已查本地 25 个 Asset 库均无记录；保留引用告警，未清空用户内容或伪造素材。生产此前修复未改，本轮未发布或提交。

## Next
None。现存素材引用、交付和本地页面验收完成；丢失旧封面源文件不能从现有本地库存恢复。

## References
- 本地证据：E:/tmp/yueli-consumer-references-20260911/docs-inventory-report.json、docs-existing-report.json、docs-media-report.json 及对应截图。
- [已上线的导入与同步引用修复](../2026-09-09-project-docs-publishing/index.md)

## Git 交付（2026-09-11）
用户已授权本地提交。本次纳入本 Work 的实现、相关验证和部署记录；其他工作改动保留，未推送或发布正式 SDK。此前“未提交”为对应阶段的历史状态。提交及范围汇总见 Workspace `flightdeck/work/2026-09-11-target-stop-isolation/commits.md`。
