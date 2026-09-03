# 超长页内目录（TOC）模式调研

## Research Read

- **问题**：当前 AE Scripting 页面把大量重复的“描述 / 参数 / 返回 / 类型”等细粒度标题全部放进右侧 TOC，形成几十甚至上百项；独立滚动条、整行选中底色和较深缩进进一步放大了拥挤感。
- **用户任务**：读者需要快速判断“我在这个 API 页的哪个对象/方法”，并跳到相邻的主要对象或方法；不是浏览正文的每个固定模板字段。
- **约束**：桌面右栏空间有限；长页面必须保持当前项可见；标题可能很长；键盘与缩放后仍需可达；移动端不应常驻右栏。
- **研究范围**：只使用 Docusaurus、VitePress、GitBook、Mintlify、Stripe、GitHub Docs、MDN 的官方页面、官方文档或官方开源源码。

## Source Matrix

| 产品 | 一手证据 | 收录层级/降噪 | 超长与滚动 | 当前项与密度 |
| --- | --- | --- | --- | --- |
| Docusaurus | [官方 TOC 文档](https://docusaurus.io/docs/markdown-features/toc)、[官方主题源码](https://raw.githubusercontent.com/facebook/docusaurus/main/packages/docusaurus-theme-classic/src/theme/TOC/styles.module.css) | 默认只收录 H2、H3；可全局或逐页设置最小/最大层级，也允许在自定义 TOC 前过滤节点。官方明确称 H2/H3 通常足以概览页面结构。 | 右栏使用 sticky，并以视口减去导航高度计算 `max-height`，仅纵向 `overflow-y: auto`；窄于 996px 隐藏。 | 将目录视为“概览”，而不是标题全集；没有固定最大项数或自动折叠合同。
| VitePress | [官方配置类型](https://github.com/vuejs/vitepress/blob/main/types/default-theme.d.ts)、[右栏容器源码](https://github.com/vuejs/vitepress/blob/main/src/client/theme-default/components/VPDoc.vue)、[Outline 源码](https://raw.githubusercontent.com/vuejs/vitepress/main/src/client/theme-default/components/VPDocAsideOutline.vue)、[目录项源码](https://raw.githubusercontent.com/vuejs/vitepress/main/src/client/theme-default/components/VPDocOutlineItem.vue)、[激活逻辑源码](https://raw.githubusercontent.com/vuejs/vitepress/main/src/client/theme-default/composables/outline.ts) | 默认只收录 H2；可配置具体层级、范围或 `deep`（H2–H6），并支持给标题加 `ignore-header` 显式排除。 | 固定于视口、纵向 `auto`，但把 scrollbar 隐藏；激活项会 `scrollIntoView({block:'nearest'})`，因此超长目录仍可访问且不会大幅跳动。1280px 以下不显示右栏。 | 13–14px 字号、32px 行高、细左边线；当前项用 2px 品牌色滑动 marker + 文字色，不使用整行填充。长标题单行省略，并保留原生 `title`。
| GitBook | [官方界面说明](https://gitbook.com/docs/resources)、[官方标题说明](https://gitbook.com/docs/creating-content/blocks/heading) | 页面 Outline 只显示其编辑器中的 Heading 1、Heading 2；在导出的 Markdown 中，它们对应页面标题之后的 H2、H3。页面可单独关闭 Outline。 | 官方说明浏览器宽度小于 1430px 时不显示右栏；未公开最大项数/折叠算法。 | 层级限制是主要降噪手段，不试图把 Heading 3（Markdown H4）也塞入右栏。
| Mintlify | [官方自定义脚本文档](https://www.mintlify.com/docs/customize/custom-scripts)、[官方页面布局文档](https://mintlify.com/docs/pages) | 暴露独立的 TOC 容器和 item 选择器；页面可用 `center` 等布局整体关闭 TOC。 | 官方 DOM 合同明确 `#table-of-contents-content` 是可滚动区域，说明超长目录采用独立纵向滚动，而不是无限撑高页面。 | 官方还区分当前可视标题与其父标题的 active 状态，适合在层级树里同时表达“当前叶子”和“所属章节”，无需给所有项强背景。
| Stripe | [官方 API 页面](https://docs.stripe.com/api/payment_intents) | Payment Intents 聚合页把“Endpoints”“Events”作为主组，把每个操作作为下一层；操作详情拆到独立页面。也就是说，它从信息架构上避免在单页继续展开每个操作里的固定字段标题。 | 官方页面没有公开可复用的最大项数或滚动算法，不能据此臆造数值规则。 | 可借鉴的是**拆页/语义分组**，不是照抄其视觉皮肤。
| GitHub Docs | [官方 frontmatter 文档](https://docs.github.com/en/contributing/writing-for-github-docs/using-yaml-frontmatter)、[官方文章内容模型](https://docs.github.com/en/contributing/style-guide-and-content-model/contents-of-a-github-docs-article) | 每篇文章自动生成 “In this article”，只收录全部 H2。这是最激进、也最稳定的层级降噪策略。 | 目录作为文章概要；未公开右栏内部滚动或折叠规则。 | H2-only 避免把固定模板小节变成重复噪声。
| MDN | [真实官方长页面](https://developer.mozilla.org/en-US/docs/Web/CSS/Reference/Properties/overflow) | “In this article” 在长参考页仅列出 Try it、Syntax、Description、Formal definition、Examples 等主要 H2，没有把 Values、各 value 说明等 H3/H4 全部列入概要。 | 页面正文仍完整保留更深层标题；导航层级与内容层级明确分离。 | 证明“正文可以很深，右侧概要无需同样深”。

## Patterns

### 1. 先减少节点，再美化滚动条

跨产品最稳定的做法是限制 TOC 层级：VitePress 与 GitHub Docs 默认 H2-only；Docusaurus 默认 H2+H3；GitBook 也只使用两级页面标题。**没有一款成熟产品默认把 H2–H6 全部展开。**

适用于本项目：默认只收录 H2；普通叙述页允许 H2+H3。API/对象参考页的“描述 / 参数 / 返回 / 类型 / 示例”等重复模板标题不进入右栏，只保留对象、属性、方法、事件等实体标题。

避免：用更小字号、更窄行高硬塞全部标题。节点量没有下降时，视觉压缩只会变成难点按的文本墙。

### 2. TOC 是页面地图，不是标题索引

MDN、GitHub Docs 和 Stripe 都把右侧或页首目录控制在“主要章节/实体”粒度。正文保留完整语义层级，但 TOC 不必一一映射。

建议提供内容级排除合同（例如 Markdown 标题属性或 frontmatter 中的规则），并允许导入转换器对已知模板标题标记为 `toc: false`。不要按显示文本在运行时全局硬编码中文/英文词表；优先依据导入源的结构语义标记。

### 3. 超长时使用视口内纵向滚动，但弱化滚动容器感

Docusaurus、VitePress、Mintlify 都存在视口内可滚动 TOC。合理实现是：sticky/fixed 右栏、`max-block-size: calc(100dvh - header - 上下留白)`、`overflow-y: auto`、`overflow-x: clip`。使用 `auto`，不要使用会在内容不足时也常显轨道的 `scroll`。

滚动条应采用浏览器原生或窄型低对比度样式，不要做成截图中粗黑色的第二主轴。不能用 `overflow: clip` 隐藏交互项：MDN 对 overflow 的官方说明指出，被 clip 的可聚焦内容仍可能获得焦点但无法滚入视图，会造成键盘不可达的体验。

### 4. 当前项用“细 marker + 文字强调”，不用整行药丸底色

VitePress 的成熟模式是整个树左侧一条浅分割线，当前项仅有 2px 品牌色 marker，文字颜色增强；marker 随当前项移动。它比整行淡紫底更轻，也不会让长目录出现一列连续按钮的感觉。

父章节可保持次级 active 文字色，叶子项才显示 marker。Hover 只改变文字色/浅底，不与 active 竞争。键盘 focus 仍需独立、清晰的 focus ring。

### 5. 密度应紧凑但维持可操作性

VitePress 官方实现使用 14px 字号、32px 行高、一级 16px 左内边距；这是合适基线。建议本项目使用：正文项 `font-size: 13px–14px`、`line-height/block-size: 28px–32px`、二级缩进 `12px–16px`。标题和首项间距 6–8px，组间可用 4px 空隙，不要给每个项目边框或大圆角底。

### 6. 长标题省略，但完整名称必须可发现

VitePress 的官方目录项使用 `white-space: nowrap; overflow: hidden; text-overflow: ellipsis`，同时把完整标题放进 `<a title>`。本项目可保持同样模式；若以后引入自定义 tooltip，需同时支持 hover 和 keyboard focus，并以原生 `title` 作为无 JS 保底。

### 7. 当前项必须自动滚入最近可见位置

VitePress 在 active 改变时调用 `scrollIntoView({ block: 'nearest', behavior: 'smooth' })`。本项目也应只在 active 项离开 TOC 可视区时最小幅度滚动，而不是每次滚动正文都把 active 项强行居中。应尊重 `prefers-reduced-motion`，此时关闭 smooth。

### 8. 响应式直接切换形态

成熟实现不会在狭窄桌面/平板继续挤压正文：Docusaurus 在 996px 以下隐藏，VitePress 在 1280px 以下隐藏，GitBook 的编辑器右栏要求约 1430px。项目应按实际三栏最小宽度决定断点；隐藏后可在正文标题下提供可展开的 “本页目录”，不要保留窄到频繁截断的常驻右栏。

## Local Application

针对当前 AE Scripting 截图，建议按以下顺序落地：

1. **结构降噪（决定性修复）**：默认 TOC 只收录 H2；对 API 模板生成内容，仅收录类/属性/方法/事件实体标题，排除重复的“描述、参数、返回、类型”。如果实体目前被错误建成 H3 而“描述”也是 H3，应先修正文档标题层级或导入语义，不能仅靠 CSS。
2. **视觉改为 VitePress 型**：保留一条浅色左边线；删除当前项整行紫色圆角背景；用 2px 紫色 marker + 较深/品牌色文字表达当前项。
3. **滚动容器弱化**：右栏按视口剩余高度滚动，横向禁止；滚动条设为 `thin`、轨道透明、thumb 使用中性低对比色，hover 时才加深。
4. **密度基线**：14px/30px；一级不额外缩进，二级 14px；标题下 6px；不为每一项增加上下 margin。
5. **行为**：active 项只在不可见时 `nearest` 滚入；长标题 ellipsis + `title`；focus-visible 清晰；reduced motion 禁用平滑滚动。
6. **极端页兜底**：若降噪后仍超过约 40 个实体项，优先按对象/模块拆页，或在 TOC 中只显示当前实体所在组及相邻顶层实体。这个 `40` 是建议的产品实验阈值，不是一手来源声明，应通过 AE 页面实测调整。

不建议：

- 不建议继续展示截图中的所有重复模板标题；这是信息架构问题，不是间距问题。
- 不建议用永久可见的粗黑滚动条；它会与左侧 active marker 形成两个竞争竖轴。
- 不建议简单裁掉超出高度的链接；键盘焦点和深链接仍必须可达。
- 不建议默认 H2–H6 或无限嵌套；成熟文档站普遍用一至两级目录换取扫描性。
- 不建议对全部项目使用整行选中底色；长列表会看起来像表单菜单，而不是阅读辅助导航。

## Next Step

做一个最小 A/B：只在 AE Scripting 页面启用“API 精简 TOC”。A 为当前实现，B 为 H2/实体-only + VitePress 型 marker + 14/30 密度 + 细滚动条。用桌面 1440×900 验证：首屏项目数、定位一个已知方法所需时间、滚动正文时 active 可见率、键盘 Tab/Enter 可达性，并在 200% 浏览器缩放下复验。若 B 仍超过约 40 项，再进入按对象拆页或“当前组窗口化”的第二阶段。
