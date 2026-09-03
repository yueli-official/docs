# Context

## Product boundary

- 本 Work 只属于 Docs 单站产品，不修改 Workspace 编排合同，也不把导入实现放进 Workspace 仓库。
- 导入格式面向转换器和人工整理后的目录，不要求来源系统采用本产品内部数据库模型。
- 现有上传 → 预检 → 用户确认 → 事务应用 → 可回滚生命周期继续保留。

## Current implementation

- 当前包必须包含 `manifest.json` v1，并以 `<locale>/<version>/` 作为 Markdown 根目录。
- JSON 目前指定 collection、version、defaultLocale、locales 和 mode；文档 Front Matter 支持 title、slug、excerpt、translationKey、order。
- 相对图片随包上传到 Asset；相对 Markdown 链接在预检中验证；外部 HTTP(S) 资源保持引用。
- 语言属于文档集，版本目标正在迁移到“一个独立文档集一个版本”的产品模型，因此交换格式不能固化旧内部 Version 行模型。

## Design direction

- Markdown 是内容真值，YAML Front Matter 是单文档可选元数据；目录和文件名提供默认层级、slug 与顺序。
- JSON 只在需要覆盖包级默认值或表达无法从 Markdown 推导的导入策略时出现，不为每个页面建立 JSON 镜像。
- 单语言、单目标版本的普通 Markdown 目录应无需 manifest；复杂多语言包允许最小 manifest 显式声明 locale 根目录和稳定逻辑键。
- 未识别 Front Matter 字段默认忽略并给出非阻塞提示，来源系统私有 MDX/组件语法必须明确报告，不能静默损坏。
