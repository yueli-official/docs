# Plan

- [x] 1. 调研主流文档系统的一手内容格式，确定 Portable Docs 1.0 规范与旧 `manifest.json` v1 迁移策略。
- [x] 2. 用解析器测试固定纯 Markdown、可选 manifest、Front Matter、目录层级、图片和链接行为。
- [x] 3. 实现 Portable Docs 1.0 解析与旧包兼容，提供规范、JSON Schema 和可下载示例包。
- [x] 4. 对齐管理界面的格式说明、错误信息和 OpenAPI，不改变现有预检/确认/回滚安全边界。
- [ ] 5. 运行 Go/Node 门禁，并用 CLI Playwright 完成真实上传、预检、确认、公开阅读与回滚验收。
