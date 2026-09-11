# 文档包

ZIP 根目录直接放 Markdown，图片使用相对路径。支持 PNG、JPEG、WebP、GIF；例如 `![界面](assets/screenshot.png)`。Markdown 使用 UTF-8。文件目录形成层级，`index.md` 是目录入口。保持文件路径稳定，以便重复导入更新同一页。

可选 Front Matter：

```markdown
---
id: getting-started
title: 快速开始
description: 安装与运行
slug: getting-started
order: 10
draft: false
---
# 快速开始
![界面](assets/screenshot.png)
```

`id` 用于跨语言关联；相同逻辑页不同语言使用同一个 id。它不是重命名文件的迁移键：重复导入按语言和路径匹配，移动文件要单独处理旧页。`draft: true` 导入为草稿；导入已下架页面的发布行为由服务端规则决定。

多语言在根目录放 `docs.json`：

```json
{
  "schemaVersion": 1,
  "defaultLocale": "zh-CN",
  "locales": {"zh-CN": ".", "en-US": "locales/en-US"}
}
```

根目录 `getting-started.md` 对应 `locales/en-US/getting-started.md`。清单只描述文档结构，目标文档集、语言和导入策略由上传参数指定。更复杂的 navigation / redirects 以 Docs 仓库的 `docs/import-package.md` 为准。

包最大 100 MiB，解压后 1 GiB，单文件 25 MiB，最多 20,000 个文件，单图片 10 MiB。MDX/JSX 与可执行配置不执行，先转换为 Markdown。外部 URL 图片不会由这个打包脚本下载；需要托管图片时先将图片下载到所选目录并改为相对引用。
