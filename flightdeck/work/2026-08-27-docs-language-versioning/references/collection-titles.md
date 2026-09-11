# 文档集多语言标题

2026-09-08 用户授权实现并询问迁移规模。本次只改 Docs，不涉及共享基础服务。

## 数据与行为

新增 0018_collection_locale_titles：collection_locales 增加 title TEXT，以所属 collections.title 回填现有语言，不修改既有迁移文件。没有正文、图片、ID 或 URL 迁移。正式当前 schema_migrations=18、dirty=false，18 条语言记录无空标题。用户已经新增/调整文档集，正式验证使用当前 adobe-ae-scripting，不能继续假定旧 aftereffects-scripting 路径或 220 篇初始规模。

每种文档语言单独保存标题。Upsert 请求显式空标题拒绝；省略标题时保存已有值，新语言初始化为当前集合标题。不是运行时猜测翻译。语言派生初始化源标题，版本派生保留其他语言标题并写入目标默认标题。

默认语言标题与 collections.title 在同一 DAO 事务同步；基础表单在语言保存后刷新标题，防止旧表单值覆盖。公开/管理树按已解析语言投影集合标题，因此页面、目录、面包屑及集合 SEO 随语言切换。后台入口：文档集 → 编辑 → 语言 → 选择语言 → 文档集标题。已有译名不会自动生成。

## 验证

- TestLocaleViewPreservesCollectionTitle 修改前红、修改后绿；controller/catalog/dao/server 测试通过。
- OpenAPI 更新，72 operation 合同检查通过；无新增错误码或 Envelope。
- Web 最终 typecheck/build 通过，git diff --check 通过。
- CLI Playwright + 真实本地 PostgreSQL：创建双语集合，API 独立标题存储、显式空值 400；后台编辑英文标题；1440/390 H1/SEO/刷新；真实下拉语言切换；默认标题同步；版本派生保留英文标题且使用新中文标题。测试集合均清理。
- 正式中英文 API locales.title 与公开集合 H1/SEO 一致，1440/390 渲染无 pageerror。当前旧标题原值回填，不代用户翻译或改名。
- 本次本地 Docs 隔离组合 5 个进程已通过 Workspace CLI 停止，其他 Provider 保持运行。

## 部署

Docs API/Web：yueli/docs-api:server-20260908-titles-1、yueli/docs-web:server-20260908-titles-1，均 healthy。
本地制品与浏览器证据：E:/tmp/yueli-media-preset-20260908，collection-titles-report.json、production-titles-report.json、screenshots/titles-*、docs-titles-source-sha256.json。
服务器同名部署目录：/projects/yuelili.com/.deploy/media-preset-20260908；deploy-titles.py、docs-titles.sha256、docs-titles-web.sha256。先构建镜像，备份 docs-before-titles.sql.gz（600），再通过现有 docs-migrate 执行 0018，最后 --no-deps --wait 更新 API/Web。未重跑 Bootstrap、认领或修改身份权限。没有 Git 提交/推送、正式 SDK 发布。
