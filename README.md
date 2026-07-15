# Docs product

- Lifecycle: active reusable product type
- Authority: Catalog product type `docs`, `api/` migrations/OpenAPI, `web/` UI
- Consumers: Docs site instances such as `docs-main`
- Verify: `pnpm platformctl verify product --file catalog/overlays/local.yaml --root . docs`

Docs owns documentation collections, hierarchical pages, navigation and search
presentation. `api/` owns durable content/domain behavior; `web/` owns public
manual reading and management. Identity and Asset remain platform dependencies;
instance URLs, DB and OIDC values come from Catalog.
