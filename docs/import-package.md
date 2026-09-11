# Docs Import Package v1

Docs 接受一个 ZIP 文档包，支持条目使用 Store、Deflate 或 XZ（ZIP method 95）压缩。正文以 UTF-8 Markdown 为唯一内容真值；普通单语言目录不需要 JSON。多语言包添加
`docs.json`；显式导航和重定向也由这份包级清单表达。

## 最小文档包

```text
my-docs.zip
├─ index.md
├─ getting-started.md
├─ guides/
│  └─ install.md
└─ assets/
   └─ overview.png
```

上传时在管理后台选择目标文档集、默认语言和导入模式。目录形成文档层级，`index.md` 表示所在目录的入口，相对 Markdown 链接
和图片引用会在预检时校验。

## 页面格式

页面使用 CommonMark/GFM 和可选 YAML Front Matter：

```markdown
---
title: 快速开始
description: 安装并运行示例项目
slug: getting-started
order: 10
id: getting-started
---

# 快速开始

![界面概览](./assets/overview.png)
```

支持字段为 `title`、`description`、`slug`、`order`、`draft` 和 `id`。缺失标题和 slug 从文件路径推导。

`id` 用于关联不同语言的同一逻辑页面。重复导入按语言和页面路径匹配；`id` 不是重命名文件的迁移键。
相对引用不能越出 ZIP 根目录；外部资源使用完整
`https` URL。MDX/JSX、Vue 组件和来源系统的可执行配置不属于通用格式，转换器必须先转为 Markdown/受限 HTML，或在预检报告
中明确标记为不支持。

## 多语言包

默认语言仍放在根目录，其他语言放在 `locales/<BCP-47>/` 的镜像路径中，并添加清单：

```json
{
  "$schema": "./docs-import.schema.json",
  "schemaVersion": 1,
  "defaultLocale": "zh-CN",
  "locales": {
    "zh-CN": ".",
    "en-US": "locales/en-US"
  }
}
```

```text
docs.json
getting-started.md
locales/en-US/getting-started.md
```

清单不保存正文，也不指定目标文档集或覆盖策略；这些属于本次上传操作。未发布试验格式不提供兼容层。

## 导入安全边界

上传只创建预检批次；存在路径冲突、断链、缺图、越界引用或超限图片时不能确认。用户确认后才上传图片并事务写入文档；完成
批次可以回滚。`replace-version` 会归档目标中未出现在包里的文档，应只用于完整镜像同步。

管理员上传包采用面向迁移任务的资源预算，而不是限制正常文档数量：ZIP 最大 100 MiB，解压后总大小最大 1 GiB，单个文件最大
25 MiB，文件条目最多 20,000 个。图片数量没有单独上限；相同内容按 SHA-256 复用一次上传，Asset 限流时导入任务会按服务端
`Retry-After` 自动等待并继续。
