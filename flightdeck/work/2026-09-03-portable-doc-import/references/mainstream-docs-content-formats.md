# 主流文档系统内容格式与可移植导入建议

> 调研时间：2026-09-03。资料只采用各产品官方文档或官方仓库；这里讨论的是内容交换格式，不是本站内部存储模型。

## 结论

本站宜采用 **Markdown-first 的目录包**，而不是“每页 Markdown 再配一份 JSON”，也不宜把完整文档树塞进一个 JSON：

- 页面正文使用 UTF-8 `.md`，元数据使用文件顶部的 YAML front matter。
- 图片、附件与页面一起放在目录包中，正文优先使用相对路径引用。
- 单页导入只需要一个 `.md` 文件，完全不要求 JSON。
- 多页导入默认可由目录结构推导路由和导航，因此仍可做到纯 Markdown。
- 只有需要**无损表达显式导航顺序、虚拟分组、隐藏页、多语言/版本入口、重定向或站点级元数据**时，才附带一个可选的 `portable-docs.json` 清单。

也就是说，对用户可宣传为“Markdown 或 Markdown 文档包”；JSON 是多页包的可选增强层，不是内容本体，也不是每页伴生文件。

这一选择与主流实现的交集最大：Docusaurus 和 VitePress以 Markdown 为核心；Nextra、Mintlify 以 MDX 为主但仍建立在 Markdown 上；GitBook 明确称自己以 Markdown 为基础，并建议迁移时优先使用 Markdown。各系统差异主要集中在导航和站点配置，而不是正文交换。因此应把正文与站点结构分层，并将 MDX/JSX 等执行性扩展视为源系统专有能力，导入时降级或报告，而不能纳入通用格式核心。

## 横向比较

| 系统 | 内容源与路由 | 导航/目录 | Front matter | 多语言 | 资产 | 对通用格式的启示 |
| --- | --- | --- | --- | --- | --- | --- |
| Docusaurus | 主要创作格式是 Markdown；`.md`/`.mdx` 均经 MDX 编译器处理，也可选择 CommonMark 解析。文档默认来自 `docs` 目录，路径参与文档 ID/URL 推导。 | Sidebar 可显式配置，也可由文件系统自动生成；自动生成时目录成为分类、文件成为页面。 | 文件顶部 YAML；文档字段包含 `id`、`title`、`description`、`slug`、`sidebar_label`、`sidebar_position` 等，且 front matter 可选，必要元数据能被推导。 | 官方 i18n 流程把主题/UI 翻译与内容翻译分开；内容通常按 locale 目录维护。 | 支持正文相对图片，也支持 `static` 下根路径资产。 | 核心应兼容普通 Markdown + YAML；排序和 slug 可保留，但不应把 Docusaurus 的 MDX/React 组件当作可移植语法。 |
| VitePress | 文件系统路由：源目录下 Markdown 路径映射为页面 URL；官方建议页间链接省略扩展名。 | 默认主题的 sidebar 是站点/locale 级配置；也可由目录扫描生成配置，但不是正文的一部分。 | 所有 Markdown 原生支持 YAML；也接受 JSON 形式的 front matter，并允许自定义字段。 | locale 可在站点配置中分别定义语言、主题文案和导航。 | 官方优先建议相对 URL；被引用资源由 Vite 处理。需保留原名或未被引用的文件放 `public`，PDF 等普通链接文件也需显式放入 `public`。 | 相对资产路径是最稳妥的交换约定；JSON front matter 虽可用，但 YAML 的生态交集更大。 |
| Nextra | 从 `content` 目录收集 `.md`/`.mdx`，或使用 Next.js App Router 的 `page.md(x)`；文件树生成 route/page map。 | 同目录 `_meta.js` 控制顺序、标题、显示与目录；未列出的路由会追加并按字母排序。 | 页面信息进入生成的 `pageMap.frontMatter`；例如 `title`、`sidebarTitle`、`asIndexPage`。 | 国际化由 Next.js/Nextra 配置和 locale 布局协作，并非一套独立于框架的正文格式。 | 使用 Next.js `public` 目录及常规 Markdown/MDX 图片引用。 | `_meta.js` 是可执行配置、不可作为交换格式；导入器应读取其中可静态提取的信息并写入中立清单。 |
| GitBook | Git Sync 将 GitHub/GitLab 仓库中的 Markdown 转为文档；内建导入支持 Markdown、HTML、DOCX，以及若干第三方来源。官方明确称 Markdown 导入效果最好。 | 可选 `SUMMARY.md` 镜像目录树；缺少时从文件夹和 Markdown 推导。`.gitbook.yaml` 指定内容根、README、SUMMARY 和 redirects。 | GitBook 页面元数据会映射到 Markdown/其扩展语法，但其核心迁移入口不是一套强制通用 front matter schema。 | 可用独立 translation space/variant 表达语言版本，并可持续同步翻译；语言维度高于单页。 | ZIP 可携带 Markdown 和图片等文件；某些 GitBook 自定义 block 无 Markdown 表示，导出时可能成为 HTML。 | `SUMMARY.md` 证明纯 Markdown 可表达基本导航；专有 block 必须有“保留 HTML、降级或报错”策略。 |
| Mintlify | 每页是带组件能力的 `.mdx`。 | `docs.json` 的 `navigation` 显式描述 pages、groups、tabs、versions、languages 等完整信息架构。 | 每页以 YAML front matter 开始；`title` 必需，另有 `description`、`sidebarTitle`、`icon`、`tag`、布局模式等。 | `docs.json.navigation.languages` 在导航树层面分区，每种语言可拥有自己的结构与站点配置。 | 编辑器将图片/视频作为仓库文件管理，页面引用仓库中的媒体。 | JSON 很适合表达复杂站点树，但它应是可选清单；Mintlify MDX 组件仍需转换为标准块或明确标为不支持。 |

### 一手资料

- Docusaurus：[Markdown Features](https://docusaurus.io/docs/markdown-features)、[Create a Doc](https://docusaurus.io/docs/create-doc)、[Docs plugin front matter](https://docusaurus.io/docs/api/plugins/@docusaurus/plugin-content-docs)、[Sidebar](https://docusaurus.io/docs/sidebar)、[i18n introduction](https://docusaurus.io/docs/i18n/introduction)、[Static assets](https://docusaurus.io/docs/static-assets)。
- VitePress：[Routing](https://vitepress.dev/guide/routing)、[Frontmatter](https://vitepress.dev/guide/frontmatter)、[Markdown Extensions](https://vitepress.dev/guide/markdown)、[Asset Handling](https://vitepress.dev/guide/asset-handling)、[Internationalization](https://vitepress.dev/guide/i18n)、[Default Theme Config: Sidebar](https://vitepress.dev/reference/default-theme-sidebar)。
- Nextra：[File Conventions](https://nextra.site/docs/file-conventions)、[`_meta.js` convention](https://nextra.site/docs/file-conventions/meta-file)、[`page.mdx` convention](https://nextra.site/docs/file-conventions/page-file)、[官方仓库](https://github.com/shuding/nextra)。
- GitBook：[Importing content](https://gitbook.com/docs/content-editor/import)、[Git Sync content configuration](https://gitbook.com/docs/getting-started/git-sync/content-configuration)、[Translations](https://gitbook.com/docs/creating-content/translations)、[Moving/exporting content](https://gitbook.com/docs/help-center/editing-content/managing-your-content)。
- Mintlify：[Pages and page metadata](https://www.mintlify.com/docs/pages)、[Navigation](https://www.mintlify.com/docs/organize/navigation)、[Site structure / `docs.json`](https://www.mintlify.com/docs/organize/settings-structure)、[Create pages and media](https://www.mintlify.com/docs/editor/pages)。

## 建议的交换包

```text
portable-docs/
├─ portable-docs.json        # 可选；复杂多页包才需要
├─ index.md
├─ getting-started.md
├─ guides/
│  ├─ index.md
│  └─ install.md
├─ locales/                  # 可选；存在翻译时使用
│  └─ en-US/
│     ├─ index.md
│     └─ getting-started.md
└─ assets/
   ├─ overview.png
   └─ reference.pdf
```

### 页面格式（必选）

正文限定为 CommonMark/GFM 的可移植子集，文件编码 UTF-8；front matter 为 YAML。建议核心字段：

```markdown
---
title: 快速开始
description: 安装并运行示例项目
slug: /getting-started
sidebar_title: 开始使用
order: 10
draft: false
tags:
  - 入门
source_id: legacy-doc-123
---

# 快速开始

![界面概览](./assets/overview.png)
```

字段规则：

- `title`：建议必填；缺失时可从首个 H1 或文件名推导，并产生 warning。
- `description`、`slug`、`sidebar_title`、`order`、`draft`、`tags`：可选、跨系统容易映射。
- `source_id`：可选但建议导入工具生成，用于幂等重导入和诊断；它不应成为公开 URL。
- 未识别字段保存在扩展命名空间（例如 `extensions.docusaurus`），不得静默覆盖核心字段。
- 页面内链接和资产优先为相对路径；外链必须是绝对 `https` URL。禁止 `..` 越出包根目录。
- HTML 可作为受限兼容输入；脚本、事件属性、iframe 等须清洗。MDX/JSX/Vue 组件不属于核心格式。

### 可选清单

`portable-docs.json` 只表达页面之外的包级关系，例如：

```json
{
  "schemaVersion": "1.0",
  "defaultLocale": "zh-CN",
  "locales": ["zh-CN", "en-US"],
  "navigation": [
    { "page": "index.md" },
    {
      "group": "指南",
      "children": [
        { "page": "getting-started.md" },
        { "page": "guides/install.md" }
      ]
    }
  ],
  "redirects": [
    { "from": "/install", "to": "/guides/install" }
  ]
}
```

建议约束：

- `schemaVersion` 必填；其余字段按需出现。
- `navigation` 中每个页面路径最多出现一次，所有路径相对包根且必须存在。
- 没有清单时按“目录优先、`index.md` 优先、其余按 `order` 后文件名”确定性推导。
- 清单存在时，它覆盖推导顺序，但未列入导航的页面仍可导入为 hidden/unlisted，并报告结果。
- locale 采用 BCP 47 标签；默认语言在根目录，翻译按 `locales/<locale>/` 镜像相同相对路径。翻译页通过“locale + 相对路径”对应，不依赖标题。
- 清单只用数据，不允许 JavaScript、模板表达式或执行钩子。

## 为什么不是纯 JSON，也不是强制 Markdown + JSON

纯 JSON 会把长文本变成转义字符串，破坏可读性、Git diff、编辑体验和现有 Markdown 工具链，也不能自然承载就近资产。主流系统没有把普通页面正文以 JSON 作为首选交换方式。

反过来，强制每次导入都提交 JSON 会给单页和简单目录制造重复真值：标题、slug、排序可能同时出现在文件名、front matter 和 JSON 中。GitBook 与 Docusaurus 已证明文件树加 Markdown 足以覆盖基本场景。

但完全排除清单也会丢失 Mintlify `docs.json`、GitBook `SUMMARY.md`、Nextra `_meta.js` 所表达的虚拟分组、精确顺序、隐藏页面、跨目录导航和语言/版本树。因此最小合理合同是：**Markdown 页面是唯一内容真值；可选 JSON 是包级结构真值；目录结构是清单缺失时的确定性回退。**

## 导入映射与降级策略

1. 先识别原生 portable 包；否则按来源适配器读取 Docusaurus/VitePress/Nextra/GitBook/Mintlify 结构。
2. 将标准 Markdown、GFM 表格、任务列表、代码块、普通 HTML 转为本站安全内容模型。
3. 将 YAML front matter 映射到核心字段，保留来源专有字段以便诊断或未来重导出。
4. 解析源系统导航配置：`sidebars.*`、VitePress sidebar、`_meta.js`、`SUMMARY.md`、`docs.json`，归一为清单导航树。可执行配置只能做安全静态提取，不能运行仓库代码。
5. 拷贝被引用资产、重写为包内相对路径，并校验文件存在、MIME、大小、哈希和路径穿越。
6. 对 MDX/JSX/Vue/专有 block 输出逐项报告：`converted`、`preserved-as-html`、`placeholder` 或 `unsupported`；不得静默删除。
7. 导入前给出 dry-run：页面数、资产数、冲突 slug、断链、未知组件、覆盖策略；确认后才落库。

## 建议的第一版范围

- 接受单个 `.md`，以及包含 `.md`、资产和可选 `portable-docs.json` 的 ZIP。
- 支持 CommonMark + GFM、YAML front matter、相对链接/资产、目录推导导航。
- 第一版清单只实现 `schemaVersion`、`defaultLocale`、`locales`、`navigation`、`redirects`。
- 明确不承诺执行 MDX/JSX/Vue、自定义 React/Vue 组件或源站构建配置；这些进入导入报告。
- 提供 JSON Schema 和示例包，但不要让 JSON Schema 演化成站点数据库导出格式。

这能先覆盖绝大多数“从其他文档站转换后上传”的需求，同时保留未来为特定来源增加适配器的空间。
