# 文档产品

- 生命周期：活跃的可复用产品类型
- 权威来源：Catalog 产品类型 `docs`、`api/` 迁移/OpenAPI、`web/` 界面
- 消费者：`docs-main` 等文档站点实例
- 验证：`pnpm platformctl verify product --file catalog/overlays/local.yaml --root . docs`

Docs 负责文档集合、层级页面、导航和搜索呈现。`api/` 负责持久内容与领域行为，`web/` 负责公开阅读和管理。Identity 与 Asset 仍是平台依赖；实例 URL、数据库与 OIDC 值来自 Catalog。

授权由实例内嵌的 Foundation Authorization Module 执行，状态保存在该 Docs 实例数据库。Identity 只证明 Subject；管理员、作者、
申请、策略和自动授权都属于 Docs。配置中的 `bootstrapAdministratorSubs` 仅在实例不存在时使用，之后修改配置不会改变授权。
作者可以创建文档并管理自己持有的文档；永久删除、重新分配、站点设置、导入和授权管理只允许管理员。Web 通过
`/api/v1/me` 的 Effective Access 显示入口，但 API 的逐能力判定始终是权威。
