# 2026-09-11 部署补充

docs 当前验收源码已同步正式服务器。API/Web 镜像、备份、验证和边界见 [Commerce 同轮部署记录](../../../../commerce/flightdeck/work/2026-09-10-digital-commerce/deployment-20260911.md)。

Docs API pay-2、Web pay-1；0020 自动检查配置迁移已生效，旧 SQL 保持不变。修复同步接口错误合同遗漏，独立 Go 测试、Web 类型与构建通过，正式公开页面和匿名权限边界通过。未用测试账户操作生产同步来源。
