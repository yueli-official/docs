# Docs 语言与版本设计调研

> 调研日期：2026-08-27
> 范围：文档集语言预设、前台语言切换、版本管理、前台版本切换，以及语言 × 版本组合。
> 本文只给出产品与数据设计建议，不包含实现修改。

> **2026-08-28 决策更新：** 下文关于“一个文档集内维护多个 `CollectionVersion`”的阶段性建议已被产品决策取代。
> 正式模型改为“一个版本一个独立文档集，多个文档集组成版本族”；文档集使用不可变 `major.minor.patch`，新版本通过
> 克隆现有文档集建立。Python 官方文档的 [3.11](https://docs.python.org/3.11/) 与
> [3.9](https://docs.python.org/3.9/) 入口也是彼此独立的版本根路径。旧 `collection_versions/default` 仅作为迁移期的
> 内部存储分区，不再是管理端或读者可见的版本产品模型。领域真值以 `CONTEXT.md` 和当前 Flightdeck 为准。

## 结论

现有 Docs 的数据底座没有走错方向：`CollectionVersion` 已经区分稳定标识、展示标签、状态、默认项和排序；`Doc` 已经通过 `versionId + locale + translationKey` 表达版本、语言与逻辑文档之间的关系。真正缺失的是：

1. 文档集级语言注册表与默认语言，而不是编辑器里写死的两个语言选项；
2. 语言和版本的管理入口；
3. 前台语言、版本选择器，以及“尽量停留在同一篇文档”的解析规则；
4. 缺失翻译、缺失版本对应页时的明确回退；
5. 版本元数据更新接口；当前只有创建和列表；
6. `sourceVersionId` 目前只是记录来源，并没有真正复制文档，不能把它包装成“版本快照”。

推荐先补齐上述管理和切换闭环，不引入 Git 分支、翻译平台、自动翻译或复杂发布流水线。

## 一手资料对比

| 项目 | 语言模型 | URL 与切换 | 缺失翻译 | 版本模型 |
| --- | --- | --- | --- | --- |
| Docusaurus | 注册 `defaultLocale`、`locales`，每个 locale 可配置 `label`、`htmlLang`、方向等 | 默认语言不带前缀，其他语言通常进入语言前缀；提供语言下拉 | 翻译文件按 locale 与版本组织；官方强调译文需要维护同步 | 版本名作为内部标识；`label`、`path`、`banner`、`badge` 独立配置；当前开发版与已发布快照分离 |
| Starlight | `defaultLocale` + locale 表；支持 `root` 默认语言 | `root` 语言无前缀，其他语言使用目录前缀 | 自动显示默认语言同页，并明确提示该页尚未翻译 | 没有内建完整版本管理，重点参考其语言回退 |
| VitePress | `locales` 中配置 `root`、`label`、`lang`、`dir` | 默认把 `/foo` 切为 `/zh/foo`；可用 `i18nRouting` 自定义并保持当前相对路径 | UI Markdown 文案可以回退到 root 配置；内容组织由目录负责 | 没有内建版本生命周期，通常由部署或站点结构补充 |
| Material for MkDocs + mike | 每个构建只有一个 canonical language；多语言通过独立项目/目录互链 | 语言和版本切换都优先保持当前相对页面 | 语言项目自行维护 | 每次部署保存一个不可变的已生成版本；版本标识与 title/alias 分离，支持 retitle、alias、default |
| GitHub Docs | URL 同时携带语言与产品版本；文章通过 frontmatter 声明适用版本 | 内部链接会按当前语言和版本重写；支持链接到“当前文章”的另一版本 | 翻译独立维护 | 不是快照复制，而是同一内容按产品/版本条件渲染；适合复杂产品矩阵，不适合当前项目直接照搬 |

参考：

- [Docusaurus i18n 配置](https://www.docusaurus.io/docs/api/docusaurus-config#i18n)：`defaultLocale` 必须属于 `locales`，locale 可配置显示标签、BCP 47 `htmlLang` 与方向；默认语言不进入 URL 前缀。
- [Docusaurus i18n 工作流](https://docusaurus.io/docs/i18n/introduction)：Markdown/MDX 按整篇翻译，以保留上下文；翻译文件按 locale 和插件组织。
- [Docusaurus 版本管理](https://docusaurus.io/docs/versioning)：区分 current、latest 和历史快照，并把 `label`、`path`、`banner`、`badge` 作为版本标识之外的元数据。
- [Docusaurus Docs 插件配置](https://docusaurus.io/docs/api/plugins/%40docusaurus/plugin-content-docs)：`lastVersion`、`onlyIncludeVersions` 与 `versions` 用来控制默认版本、发布集合和版本展示。
- [Starlight i18n](https://starlight.astro.build/guides/i18n/)：相同文件名关联不同语言页面；缺失翻译时显示默认语言内容并给出未翻译提示。
- [VitePress i18n](https://vitepress.dev/guide/i18n)：locale 配置包含 `label` 和 `lang`，内容按语言目录组织。
- [VitePress `i18nRouting`](https://vitepress.dev/reference/default-theme-config.html#i18nrouting)：切换语言时默认保留当前相对路径，也支持自定义目标地址。
- [Material for MkDocs 语言设置](https://squidfunk.github.io/mkdocs-material/setup/changing-the-language/)：语言选择器使用非空名称、绝对目标和语言代码；同路径存在时切换后停留在当前页。
- [Material for MkDocs 版本设置](https://squidfunk.github.io/mkdocs-material/setup/setting-up-versioning/)：版本作为独立部署结果，旧版本不随新版本改变；版本切换默认尝试停留在同页。
- [mike 官方仓库](https://github.com/jimporter/mike)：版本标识、显示 title、alias 和 default 分开管理；提供 `retitle` 而不是把版本标识当普通标题编辑。
- [GitHub Docs 内容模型](https://github.com/github/docs/blob/main/content/README.md)：URL 会带语言与版本上下文，内部链接按当前上下文重写；文章通过 `versions` 声明适用范围。
- [GitHub Docs 版本说明](https://docs.github.com/en/get-started/using-github-docs/about-versions-of-github-docs)：读者通过页头选择器切换适用的产品版本。

## 本地现状

### 已有能力

- `CollectionVersion` 已有：
  - `key`：版本稳定标识；
  - `label`：前台展示名；
  - `status`：`draft | published | archived`；
  - `isDefault`：默认版本；
  - `sortOrder`：展示顺序；
  - `sourceVersionId`：来源版本记录。
- 数据库保证一个文档集只有一个默认版本，并保证 `(collection_id, key)` 唯一。
- `Doc` 已有 `versionId`、`locale`、`translationKey`。
- 文档路径在 `(collection, version, locale, parent, slug)` 内唯一。
- 公共查询和搜索已经接受 `locale`、`version`；默认项在 URL query 中省略。
- 当前文档、目录、搜索、评论和分析已经携带版本与语言上下文。

### 当前断点

- `locale` 仍是自由字符串，没有文档集支持语言、默认语言、显示名称和排序。
- 编辑器仅写死 `English / en` 与 `简体中文 / zh-CN`。
- 版本 API 只有 `GET list` 和 `POST create`，没有更新 label、status、default、order 的能力。
- 公共页面读取 query，但没有语言和版本选择器。
- 切换只靠 path/query 无法覆盖译文 slug 不同、版本结构变化的情况。
- `sourceVersionId` 没有触发文档复制；现在创建的版本只是空壳元数据。
- `translationKey` 没有清晰写成跨语言、跨版本的稳定页面身份，也缺少对应约束。

## 语言设计建议

### 1. 语言属于文档集，不属于全站固定枚举

不同文档集可能只需要中文，也可能需要中英双语。建议增加文档集语言注册：

```text
collection_locales
├── collection_id
├── locale          # 稳定标识，如 zh-CN、en、ja
├── label           # 本地语言名称，如 简体中文、English、日本語
├── html_lang       # BCP 47，用于 <html lang> / hreflang
├── direction       # ltr / rtl
├── is_default
├── enabled
└── sort_order
```

语言选择器提供常用预设，例如 `zh-CN`、`zh-TW`、`en`、`ja`、`ko`，预设填充本地名称、`htmlLang` 和方向；高级入口允许合法的自定义 BCP 47 标签。不要把全部语言一次性铺在表单里，也不要继续允许作者随手输入任意字符串。

`locale` 一旦已有文档，就视为稳定标识；日常只允许修改 `label`、启用状态和排序。需要改代码时走专门迁移操作，而不是普通编辑。

### 2. 默认语言

- 每个文档集恰好一个默认语言。
- 默认语言必须启用。
- 新文档默认使用文档集默认语言，不再全局硬编码 `en`。
- 公开 URL 继续省略默认语言；非默认语言携带语言上下文。
- 修改默认语言会改变 canonical URL 与回退目标，应是显式管理操作。

### 3. 前台语言切换

前台应该支持语言切换，但只在当前文档集启用了两个及以上语言时显示。切换目标不能只替换 query 或 slug，应该按稳定页面身份查找：

```text
(collection, targetVersion, targetLocale, translationKey)
```

建议解析顺序：

1. 目标语言、当前版本、相同 `translationKey`；
2. 当前版本、默认语言、相同 `translationKey`，并显示“此页暂无该语言版本，正在显示默认语言”；
3. 目标语言的当前版本文档集首页；
4. 默认语言的当前版本文档集首页。

不要静默把默认语言内容伪装成用户选择的语言。Starlight 的“回退内容 + 明确提示”比直接 404 或无提示混用更适合渐进翻译。

### 4. `translationKey` 的语义

将它明确为“逻辑页面身份”，而不是只服务当前版本的一次性翻译键：

- 同一篇内容的不同语言、不同版本保留相同 `translationKey`；
- 新建译文时从源文档继承；
- 从旧版本复制出新版本时继承；
- slug、标题和父级可以随语言、版本变化；
- 至少增加 `(collection_id, version_id, locale, translation_key)` 唯一约束。

这样语言和版本切换都可以保持当前文档上下文，而不是假定路径永远一致。

## 版本设计建议

### 1. 版本 key 不作为普通可编辑字段

`key` 已经参与 API 查询、URL、唯一约束、搜索与分析，是稳定标识。主流系统也普遍把稳定版本标识与展示名分开：Docusaurus 允许自定义 label/path，mike 提供 `retitle` 修改 title。

因此版本编辑应支持：

- `label`：可改，例如 `2.1` 改为 `2.1 LTS`；
- `status`：`draft / published / archived`；
- `isDefault`：在已发布版本中选择唯一默认项；
- `sortOrder`：控制选择器顺序。

`key` 创建后只读。未来确实需要改 key 时，再增加低频的“更改版本标识”迁移操作，并同时生成旧 URL 重定向；不应为当前阶段先做。

### 2. 状态语义

| 状态 | 管理端 | 前台 |
| --- | --- | --- |
| `draft` | 可编辑 | 不出现在公共选择器，不允许公共读取 |
| `published` | 可编辑 | 可公开读取，可成为默认版本 |
| `archived` | 低频维护 | 仍可读取并出现在选择器尾部，页面显示“旧版本”提示 |

默认版本必须是 `published`。归档默认版本前必须先选择另一个默认版本。当前 `ResolveVersion(..., publicOnly=true)` 只允许 `published`，若采用上述语义，需要显式放行 `archived`，而不是误以为现状已经支持。

### 3. 默认版本与展示顺序

- 一个文档集只有一个默认版本。
- 默认版本的 URL 不携带 `version`，其他版本携带稳定 key。
- 前台选择器按 `sortOrder` 展示，默认版本优先，归档版本置后。
- 版本显示 label，不显示内部 key；key 可以作为次要信息出现在管理端。
- 只有一个公开版本时不显示版本选择器。

### 4. 版本切换

版本切换也使用 `translationKey` 保持当前页面：

```text
(collection, targetVersion, currentLocale, translationKey)
```

建议回退顺序：

1. 目标版本、当前语言、同一 `translationKey`；
2. 目标版本、默认语言、同一 `translationKey`，显示语言回退提示；
3. 目标版本、当前语言的文档集首页；
4. 目标版本、默认语言的文档集首页。

Material for MkDocs 的版本切换和 VitePress 的语言切换都把“切换后尽量停留在当前页”作为默认体验；路径不存在时回到目标版本首页比留在旧版本或直接 404 更可预测。

### 5. 版本快照与 `sourceVersionId`

当前创建版本只插入 `collection_versions`，不会复制文档。`sourceVersionId` 只是来源记录，不能表示一个真实快照。

当前阶段建议：

- 新建版本默认是空版本；
- 暂不在 UI 中宣称“从版本复制”或“创建快照”；
- 等真实需求出现后，再增加“从现有版本创建”：事务内复制树、正文与元数据，生成新文档 ID，但保留 `translationKey`；
- 已发布历史版本仍允许管理员修正内容，不必照搬 mike 的严格不可变构建产物；修改要保留审计记录即可。

不要把 Git branch 变成产品模型。Docusaurus/mike 的分支或构建快照适合文件型静态站，当前 Docs 是数据库 CMS，版本应该是数据库中的内容集合。

## 语言 × 版本组合

不要提前创建完整笛卡尔积。一个文档存在于：

```text
(collectionId, versionId, locale, translationKey)
```

文档集注册“允许的语言”，版本注册生命周期；某个版本是否真的有某语言内容，以该组合下是否存在已发布文档为准。

管理端可以展示简洁的覆盖情况，例如：

| 版本 | 简体中文 | English |
| --- | --- | --- |
| 2.0（默认） | 12 篇 | 8 / 12 篇 |
| 1.0（旧版） | 10 篇 | 未启用内容 |

但不要让管理员逐格创建“版本语言实例”。创建文档或复制版本时自然产生内容即可。

公共页面不直接展示版本 × 语言组合矩阵，而是分别提供两个选择器：

- 语言菜单显示文档集已启用语言；没有当前页译文时仍允许选择，但使用明确回退；
- 版本菜单显示所有公开/归档版本；没有当前页对应版本时回到目标版本首页；
- 搜索、目录、前后篇、评论和分析始终继承当前解析后的版本与语言。

## URL 设计

主流文档站更常用路径前缀，例如：

```text
/{locale}/{version}/{collection}/{path}
```

并省略默认语言、默认版本。GitHub Docs 采用语言 + 产品版本路径，Docusaurus、Starlight、VitePress 也通常使用语言目录。

但是本地 Docs 已经有 URL 生命周期与 query 形式：

```text
/{collection}/{path}?locale=zh-CN&version=v2
```

现阶段不建议为了“更像主流”立即改路由。先用现有 query 完成管理、切换、回退和 canonical 规则，避免同时引入路径迁移。等正式发布前再单独决定是否切为路径前缀；如果切换，必须通过现有 URL lifecycle 为旧 query URL 建立 canonical/redirect，而不是两套 URL 长期并存。

## 分阶段实施建议

### 阶段 1：补齐语言闭环

1. 增加文档集语言注册与默认语言。
2. 管理端提供语言预设、启用、默认、排序；locale code 创建后只读。
3. 编辑器语言下拉读取当前文档集配置。
4. 前台增加语言选择器，使用 `translationKey` 保持当前页。
5. 实现默认语言回退与明确提示。
6. 搜索、目录、评论、分析继承解析后的语言。

### 阶段 2：补齐版本闭环

1. 增加版本更新接口：label、status、default、sortOrder。
2. 管理端增加版本管理，不允许普通修改 key。
3. 前台增加版本选择器与归档版本提示。
4. 使用 `translationKey` 保持当前页；缺页回到目标版本首页。
5. 约束默认版本必须已发布。

### 阶段 3：确有需要再做

- 从现有版本复制完整文档树；
- 翻译完成度与源文档变更后“译文可能过期”的标记；
- 版本 key 迁移与 URL 重定向；
- 翻译平台或自动翻译接入；
- 路径式 locale/version URL。

## 不建议当前实现

- 不为每个语言创建独立文档集；这会破坏 `translationKey` 与统一导航。
- 不允许任意 locale 字符串继续进入数据库。
- 不根据浏览器语言自动强制跳转；优先使用明确的用户选择，必要时记住选择。
- 不把缺失译文静默当成目标语言内容。
- 不允许普通编辑版本 key。
- 不把 Git branch、静态构建目录或第三方翻译平台引入核心模型。
- 不在第一阶段创建版本 × 语言的完整实例表。

## 推荐的产品决策

1. **语言功能保留并补齐。** 文档站应支持语言预设和前台切换；现有模型已具备大半底座。
2. **版本功能也保留，但先做元数据管理与切换。** 不需要删除已有版本模型，也不需要立刻做复杂快照系统。
3. **版本 key 稳定，label 可修改。** 这是避免链接、搜索、评论和分析记录失联的关键。
4. **默认语言、默认版本省略 URL 参数。** 当前 query 路由先保持，路径化另立任务。
5. **`translationKey` 作为跨语言、跨版本逻辑页面 ID。** 语言和版本切换都围绕它解析。
6. **缺失内容必须可解释。** 语言缺失显示默认语言同页并提示；版本缺页进入目标版本首页。
