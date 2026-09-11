---
name: yueli-docs-publish
description: 将项目 Markdown 和图片打包并发布到月离文档，支持预检、重复更新和 GitHub Release 文档包。用于其他项目维护并上传文档，不用于通用文章编辑。
---

# 发布项目文档

从当前项目选择文档目录和目标文档集 slug。使用账户中心创建、勾选“月离文档 → 导入文档”的开发者令牌，通过 `YUELI_DOCS_TOKEN` 环境变量读取。不要将令牌写入项目、命令行参数或输出。

脚本路径相对于本 Skill 目录；与项目 checkout 无关，只需 Python 3.10+ 标准库。

```sh
python scripts/publish.py pack ./docs --output ./docs.zip
python scripts/publish.py preflight ./docs.zip --base https://docs.yuelili.com --collection my-project --locale zh-CN
python scripts/publish.py confirm BATCH_ID --base https://docs.yuelili.com
```

用户已要求发布时，可执行 `publish`，它会上传、检查阻塞问题后确认。仅要求检查时使用 `preflight`。确认请求超时后使用 `status BATCH_ID` 查询，不重新上传或盲目重试确认。

```sh
python scripts/publish.py publish ./docs.zip --base https://docs.yuelili.com --collection my-project --locale zh-CN
python scripts/publish.py status BATCH_ID --base https://docs.yuelili.com
```

默认 `upsert` 按语言和页面路径更新，未出现的页面保留。不要自行使用 `--mode replace-version`：它会下架目标中缺失的页面，只有用户要求完整镜像且目标专属本项目时才使用。令牌持有者仍须拥有站点当次导入权限。

页面、图片和多语言格式见 [references/package.md](references/package.md)。缺图、断链等以服务端预检报告为准；不要为了通过预检悄悄删除正文。打包脚本仅收录 Markdown、docs.json 和支持的图片，不打包密钥或整个仓库。

需要在项目发版时提供 docs.zip 时，参考 [assets/docs-release.yml](assets/docs-release.yml)。将脚本复制到项目的 `scripts/yueli_docs_publish.py`，按实际目录调整工作流。创建模板不等于授权推送或发布 Release。


GitHub 自动更新：项目 Release 提供 docs.zip 后，在 Docs「批量导入 → 项目同步」绑定公开仓库、附件和目标文档集。默认每 15 分钟检查，可立即检查、暂停、查看报告或更换令牌；用户不必保持页面打开。绑定所选文档集即决定发布版本。包有阻塞问题时保留上次成功内容。
