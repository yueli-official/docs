---
name: web-doc-scraper
description: 从外部网站爬取文档并转换为 Markdown 集成到 Astro Starlight 文档站。用于抓取、爬取、采集外部文档内容。
argument-hint: "[入口URL]"
---

## 整体流程

### 第 1 步：探测目标站点

用户提供文档入口 URL 后，**不要立即大规模爬取**，先做小范围探测：

1. 用 `WebFetch` 尝试获取页面内容
2. 如果返回空壳（纯 JS 渲染站点），改用 `curl -s -L <url> -A "Mozilla/5.0"` 获取
3. 分析返回的 HTML，确定：
   - **正文容器**：正文在哪个标签/选择器内（如 `<div id=article>`、`<main>`、`<article>` 等）
   - **非正文区域**：侧栏、页脚、导航、语言选择等需要排除的部分
   - **URL 模式**：子页面 URL 的规律（路径结构、slug 规则）
   - **图片来源**：图片 URL 的模式（CDN 路径、相对路径等）
   - **页面结构**：标题层级、内容分区（描述、参数、代码块等）

4. 随机抽取 2-3 个子页面验证结构一致性

### 第 2 步：分析文档结构

从入口页或侧栏提取完整的文档目录树：

- **分类**（一级导航）
- **子页面列表**（每个分类下的具体页面）
- **特殊页面**（索引页、概览页、单页多锚点页等）

输出目录树给用户确认，包含：
- 分类划分是否合理
- 是否有需要跳过的页面（如 license、changelog 等）
- 目录命名是否符合预期

### 第 3 步：用户确认

**必须等用户确认后才开始大规模爬取**，确认内容包括：

- 哪些分类/页面需要采集，哪些跳过
- 目录结构和命名规范
- 特殊页面的处理方式（保留/拆分/跳过）
- 索引页是否需要

### 第 4 步：编写爬取脚本

根据探测结果编写 **Python 脚本**（使用 `re` + `subprocess` 调用 curl），实现：

1. 批量获取所有页面 HTML
2. 提取正文区域
3. HTML → Markdown 转换
4. 图片 URL 提取与路径替换
5. 写入 `.md` 文件

**重要**：脚本放在系统临时目录（`os.environ['TEMP']`），不污染项目目录。

### 第 5 步：下载图片并校验

1. 从页面 HTML 中提取所有图片 URL
2. 批量下载到 `_static/` 目录
3. **校验图片有效性**：检查文件头（如 JPEG 应以 `FF D8 FF` 开头），删除无效文件（CDN 可能返回 XML 错误页面）
4. 清除 `.md` 文件中对无效图片的引用

### 第 6 步：生成导航并注册

1. 根据分类数据生成 `src/nav/<topic>.ts` 导航文件
2. 在 `astro.config.mjs` 中导入并注册到 `starlightSidebarTopics`

### 第 7 步：注册到首页

在以下三个首页文件的 `<CardGrid>` 中添加新文档的卡片入口：

- `src/content/docs/index.mdx` — 默认首页（中文）
- `src/content/docs/en.mdx` — 英文首页
- `src/content/docs/zh-cn.mdx` — 中文首页

卡片格式：

```jsx
<a href="/en/<topic>/intro/overview">
  <Card title="文档名称" icon="open-book">
    文档简要描述。
  </Card>
</a>
```

中英文首页分别使用对应语言的链接前缀（`/en/` 或 `/zh-cn/`）和文案。

### 第 8 步：用户验收

启动 dev server 让用户检查，根据反馈迭代修复。

---

## 本项目结构约定

```
src/
├── content/docs/
│   ├── en/<topic>/<category>/<page>.md    # 英文文档（默认语言）
│   └── zh-cn/<topic>/<category>/<page>.md # 中文文档
├── nav/<topic>.ts                          # 侧栏导航配置
astro.config.mjs                            # 注册导航到 starlightSidebarTopics
```

- **Frontmatter**：仅需 `title` 字段
- **导航格式**：TypeScript 对象，支持 `label`、`link`、`translations` 嵌套结构
- **i18n**：en 为默认语言，zh-CN 为翻译语言

---

## HTML → Markdown 转换规则

### 标签转换

| HTML | Markdown |
| --- | --- |
| `<h1>` ~ `<h6>` | `#` ~ `######`（根据实际需要调整层级） |
| `<b>` / `<strong>` | `**粗体**` |
| `<i>` / `<em>` | `*斜体*` |
| `<a href=...>` | `[文本](链接)` |
| `<br>` | 换行 |
| `<p>` | 空行分段 |
| `<code>` / `<pre>` | `` `行内代码` `` 或 ` ```代码块``` ` |
| `<ul><li>` | `- 列表项` |
| `<ol><li>` | `1. 列表项` |
| `<table>` | Markdown 表格 或保留 HTML（复杂表格） |
| `<img>` | `![alt](src)` 或删除后单独处理 |
| `<dl><dd>` | 根据语义转换为列表或参数格式 |

### HTML 实体解码

`&amp;` → `&`、`&lt;` → `<`、`&gt;` → `>`、`&quot;` → `"`、`&#39;` → `'`、`&nbsp;` → 空格

### 清理规则

- 删除空的格式标记（如 `**` 内无内容、`*` 内无内容）
- 连续 3 行以上空行压缩为 2 行
- 行内连续空格压缩为单个
- 去除行首行尾多余空白

---

## 图片处理规则

### 存放位置

所有图片统一放在主题根目录的 `_static/` 下，不按分类拆分：

```
src/content/docs/en/<topic>/
├── _static/          # 所有图片
│   ├── image1.jpg
│   └── image2.png
├── <category>/
│   └── page.md       # 引用图片：../_static/image1.jpg
```

### 相对路径计算

页面在分类子目录中时，引用图片使用 **一层上级** 路径：`../_static/filename.jpg`

页面在主题根目录中时，引用图片使用 **同级** 路径：`./_static/filename.jpg`

### 图片校验

下载后必须验证文件有效性（检查文件头魔数），CDN/服务器可能返回错误页面（HTML/XML）伪装为图片文件。无效文件删除并清除 `.md` 中的引用。

---

## 链接处理规则

### 内部链接转换

将源站链接转换为本站路径格式：

```
源站：/documentation/product/section/page/
本站：/en/<topic>/<category>/<slug>
```

### 分类路径映射

如果本站按分类组织目录但源站 URL 没有分类层级，需要构建 `slug → category` 映射表，确保所有内部链接都包含正确的分类路径。

### 外部链接

保持原样，不做转换。

---

## 导航文件生成

导航文件 `src/nav/<topic>.ts` 格式：

```typescript
export const topicName = {
  label: "Topic Name",
  link: "<topic>/<first-page>",
  icon: "open-book",
  items: [
    {
      label: "Category Name",
      translations: { "zh-CN": "分类中文名" },
      items: [
        {
          label: "Page Name",
          link: "/<topic>/category/page",
        },
      ],
    },
  ],
};
```

---

## 索引页处理

如果源站有索引/目录页，根据页面类型选择合适的展示格式：

| 索引类型 | 推荐格式 |
| --- | --- |
| 纯名称列表 | `- [Name](link)` Markdown 列表 |
| 带图片的网格 | 3 列 Markdown 表格，每个单元格 `![img](src)<br>[Name](link)` |
| 带摘要的列表 | `- [Name](link) — 摘要描述` |

---

## 命名规范

- **目录名**：小写，连字符分隔（kebab-case），如 `glow-effects`
- **文件名**：小写，连字符分隔或源站 slug 小写，如 `blur.md`
- **图片文件名**：保持原始文件名（保留大小写），如 `Glow.jpg`

---

## 工具选择优先级

| 步骤 | 首选工具 | 备选工具 | 说明 |
| --- | --- | --- | --- |
| 页面获取 | WebFetch | Bash (curl) | JS 渲染站点 WebFetch 无法获取时用 curl |
| HTML 解析 | Python (re) | — | 正则提取，不依赖外部库 |
| 图片下载 | Bash (curl) | Python | curl 更方便处理 HTTP 头 |
| 文件写入 | Write | Python | 少量文件用 Write，批量用 Python 脚本 |
| 导航生成 | Python → Write | — | Python 生成 TS 内容，Write 写入文件 |

---

## 注意事项

1. **先探测后爬取**：永远先小范围探测，确认结构后再批量处理
2. **用户确认**：分类结构、跳过内容、命名规范必须用户确认后才执行
3. **迭代修复**：第一版产出后让用户检查，根据反馈修复格式问题
4. **JS 渲染站点**：WebFetch 返回空内容时，用 `curl` 作为后备方案
5. **编码**：全程 UTF-8
6. **速率控制**：避免短时间内大量请求同一站点
7. **脚本存放**：爬取脚本放在系统临时目录，分类数据（JSON）也放临时目录
8. **Windows 兼容**：使用 `python` 而非 `python3`，路径用正斜杠，临时目录用 `os.environ['TEMP']`
9. **不保留的内容**：导航/侧栏、页脚、"See Also"等关联链接区域、版权声明等装饰性内容按需跳过（与用户确认）
