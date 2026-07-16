# 文档产品

- 生命周期：活跃的可复用产品类型
- 权威来源：Catalog 产品类型 `docs`、`api/` 迁移/OpenAPI、`web/` 界面
- 消费者：`docs-main` 等文档站点实例
- 验证：`pnpm platformctl verify product --file catalog/overlays/local.yaml --root . docs`

Docs 负责文档集合、层级页面、导航和搜索呈现。`api/` 负责持久内容与领域行为，`web/` 负责公开阅读和管理。Identity 与 Asset 仍是平台依赖；实例 URL、数据库与 OIDC 值来自 Catalog。
