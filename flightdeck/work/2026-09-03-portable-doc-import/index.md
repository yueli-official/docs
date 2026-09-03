# Docs Import Package

## Goal

交付一个有文档、可验证、便于其他文档系统转换的 Markdown-first ZIP 导入格式；普通目录使用纯 Markdown，多语言、显式导航
和重定向使用可选 `docs.json`，正文、层级、元数据、图片和内部链接可预检与导入。

## Status

Open

## Current

Docs Import Package v1 已实现：单语言 ZIP 无需 JSON；多语言/导航包使用 `docs.json`，默认语言位于根目录，其他语言目录由
`locales` 映射声明。正式 Front Matter 为 `id/title/description/slug/order/draft`；未发布的 `manifest.json` 与
`portable-docs.json` 试验格式已删除，不保留兼容层。上传表单显式选择目标文档集、默认语言和导入模式。

Foundation/Identity/Asset/Docs 依赖合同已对齐，Docs Shared Session `20260903T133805Z-21912` ready，入口
`http://192.168.5.7:3003`。管理页双 Vue runtime hydration 已修复。Markdown 图片 destination 现正确剥离可选标题；AE 双语 ZIP
真实预检为 120 篇、4 张图片、零阻塞。新建文档集已在同一事务创建默认语言和不可变语义版本，真实 `zh-CN + 1.0.0`
创建、查询和删除通过。Importkit、Catalog 定向、Controller、Server、Web 52 tests、两条真实 Playwright 与 Impeccable detector
均通过。Web typecheck 仍被 Foundation/Overlay 的 Tiptap 3.28.0/3.30.2 既存类型漂移阻塞。

随后真实确认链路已完成：补齐 `docs-import-image` Asset Profile，服务端 PUT 改为发送原始字节与准确 Content-Length；导航虚拟组
物化为双语章节入口，文档按父层级批量插入；新增 0017 让 `updated_at` 与 Search revision 保持一致，重复 upsert 不再触发幂等冲突。
当前 Session `20260903T164405Z-26776` ready。真实 AE 包以 120 个页面 + 24 个双语章节入口、4 张图片完成导入；确认失败只显示
一个 `duration: 0` 的可手动关闭 Toast，不再重复渲染右侧 Alert。

AE 源仓库现提供一键导出脚本，将 `note/info/tip/important/warning/caution/danger` 容器转换为 GFM Alerts，复制包内图片并从
`src/nav/ae-scripting.ts` 生成顺序；未知或未闭合容器会直接中止。Docs 图片仍按 10 MiB 文件大小限制，产品原有 3000 万像素
限制已取消，`maxPixels` 与 Asset 基础设施的 8000 万安全护栏对齐。

## Next

用户已确认重新导入转换后的 AE 双语包有效；同一脚本现已批量生成其余 8 套双语文档包。后续按文档集逐一导入验收。验收继续遵守
[媒体合同](../../../../workspace/flightdeck/knowledge/asset/consumer-media-contract.md)，不得移除 Asset 的解码炸弹基础设施护栏。

## Progress

- 2026-09-03：审计现有 ZIP importer、Front Matter、图片/链接校验、批次事务与管理入口，确认问题主要在交换格式而非导入生命周期。
- 2026-09-03：完成 Docusaurus、VitePress、Nextra、GitBook、Mintlify 官方资料对照，选定“纯 Markdown 默认 + 可选包级 JSON”。
- 2026-09-03：实现 Markdown-first parser、上传目标参数、v1 兼容、格式文档、Schema、双语示例和静态门禁；真实启动被 Lock 不兼容阻塞。
- 2026-09-03：以 `yozya/docs` 的 120 篇中英文 AE Scripting 源文档和 `ae-scripting.ts` 导航为只读输入，生成
  `exports/ae-scripting-docs-v1.zip`：中英文各 60 篇、10 个顶层导航组、60 个有序页面和每语言 2 张图片；
  源内容与导航工作树保持未修改。
- 2026-09-03：修复 Markdown 图片可选标题误报，并将默认语言和语义版本纳入文档集首次创建事务与表单；真实浏览器证明
  临时 `zh-CN + 1.0.0` 文档集生命周期及 AE 包 120 文档/4 图片预检全绿。
- 2026-09-03：修复 Asset Profile、原始 PUT、导航父级、分层批量事务和 Search revision；真实 Playwright 证明 AE 双语包确认
  `completed`，模拟失败只有一个持久错误 Toast。
- 2026-09-03：长 AE API 页右侧 TOC 改为页面单一滚动轴：共享 Module 删除固定高度和内部滚动，Docs 阅读布局删除固定 240px
  列并改为 3:1 自适应比例；长标题省略并以原生 title 提供完整 hover 文本。桌面/390px Playwright 与 detector 全绿。
- 2026-09-03：按用户复核将右侧 TOC 改为紧凑的视口内纵向滚动轨，增加左侧 1px 分隔并继续禁止横向滚动；公开文档标题区
  对具备 `docs.document.update` 的用户显示直达语义后台路由的“编辑”按钮。管理员/匿名/390px Playwright、Foundation 单测、
  Docs 52 tests 与 detector 全绿。
- 2026-09-03：对照 Docusaurus、VitePress、GitBook、Mintlify、GitHub Docs 与 MDN 官方实现后，将公开 TOC 限定为 H2/H3，
  排除 API 页重复的“描述/参数/返回/示例/类型”H4；视觉改为浅左轨、主色文字和短 active marker，取消整行紫色高亮，并使用
  细低对比滚动条。真实 1440px/390px、管理员/匿名 Playwright 与 detector 全绿，调研见 `references/long-toc-patterns.md`。
- 2026-09-03：修复 Foundation `ContentProse` 在 LAN HTTP 下 Clipboard API 不可用时静默失败；增加复制 fallback、原按钮成功/失败
  状态和清理逻辑。真实 AE 代码块 Playwright 已证明点击后出现“已复制”，通用 Content 9 tests 全绿。
- 2026-09-03：编辑页公开链接复制改用共享 LAN-safe clipboard；Slug 与摘要迁入“文档设置 → 内容”，标题输入并入始终可见的
  顶部命令栏。新增“沉浸式协作”：全屏固定工作区只保留顶部命令与下方独立滚动正文；普通模式的共享编辑工具栏固定在命令栏
  下方。1440px/390px 长文滚动、沉浸切换和复制 Playwright 2/2、Docs 52 tests、Content 9 tests 与 detector 全绿。
- 2026-09-03：按用户复核将沉浸式协作进一步收敛为纯编辑器画布：Docs 命令栏和标题输入完全不渲染，退出动作进入共享编辑器
  toolbar slot，正文占满视口。共享 AdminIconPicker 从固定 5 列改为 auto-fill 容器网格；Inspector 实测超过 5 列。桌面/390px
  Playwright、UI/Icon Picker 3 tests、Content 9 tests、Docs 52 tests 与 detector 全绿。
- 2026-09-03：通用编辑工作区规则已推广到 Blog：保留 Blog 封面/分类/标签/系列/发布/SEO Adapter，标题进入 sticky 命令栏，
  Slug/摘要进入 Inspector，普通工具栏 sticky，沉浸模式只保留编辑器。Blog API 全量 Go、编辑器源码合同、1440px/390px
  Playwright 与 detector 全绿；全量 Web 仍有既存 Asset 页面旧断言和 Tiptap 双版本类型漂移。
- 2026-09-03：新增可复用 `export-docs-package.mjs` 与测试，将 AE 双语源的 616 个 Docusaurus 容器转换为 GFM Alerts，保留
  120 篇 Markdown、双语导航顺序和 4 个图片条目；源目录与导航文件不变。真实 Docs Playwright 预检得到 144 个节点、4 张图片、
  零问题。Docs 专属 3000 万像素限制改为 Asset 通用 8000 万安全护栏，10 MiB 文件大小限制不变。
- 2026-09-03：用户确认 AE 转换包实际导入有效。导出器新增 `--all` 自动发现及批量报告，兼容中文 `注意`、`bug`、`quote`、
  旧式反向容器、行尾关闭标记，以及导航/文件名的大小写与括号差异；9 套双语包共 3818 篇 Markdown、804 个图片条目全部生成，
  ZIP 内 `:::` 指令残留为零，源文档和导航保持未修改。
- 2026-09-04：批量导入页支持具备 `docs.collection.manage` 权限的用户原地新建文档集，同时填写标题、路径、默认语言和首个语义版本；
  创建成功后自动选为当前导入目标并同步默认语言，无需离开或刷新页面。桌面真实创建/选中/删除和 390px 响应式 Playwright、
  Web 52 tests 与 Impeccable detector 全绿；typecheck 仅余 Foundation/Overlay Tiptap 3.28.0/3.30.2 既存类型漂移。
- 2026-09-04：修复 5.58 MiB `sapphire-docs-v1.zip` 被通用 BFF 的 1 MiB 缓冲正文上限拒绝并显示
  `foundation.request.network`：只为 `POST /api/v1/imports/docs` 增加复用 Identity 会话 Adapter 的流式代理，保留 32 MiB
  Content-Length 门禁和 5 分钟截止时间，其他 API 安全预算不变。原始 Sapphire 浏览器复现由连接重置转为 582 篇、546 张图片、
  零问题的成功预检，回归 Playwright 9.0 秒通过；Web 52 tests 全绿，typecheck 仅余既存 Tiptap 双版本漂移。
- 2026-09-04：修复 Sapphire 确认阶段的 `docs.upstream_failed(common.rate_limited)`：图片按 SHA-256 跨语言去重，546 个路径归并为
  270 份物理上传；Asset HTTP 客户端遵守 429 的 `Retry-After`/`Ratelimit-Reset` 后自动续传，确认端点使用独立 5 分钟代理预算。
  临时文档集真实 Playwright 完成上传、零问题预检、270 张图片上传、582 篇文档确认与清理，耗时 2.3 分钟；Asset retry、内容去重、
  导入确认定向 Go tests 全绿。

## References

- [格式调研](references/mainstream-docs-content-formats.md)
- [长 TOC 模式调研](references/long-toc-patterns.md)
- [通用编辑工作区知识](../../../../workspace/flightdeck/knowledge/frontend/editor-workspace.md)
