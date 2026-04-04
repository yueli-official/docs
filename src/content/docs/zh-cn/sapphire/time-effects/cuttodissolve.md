---
title: CutToDissolve
---

## S_CutToDissolve

将单个素材中的硬切转换为叠化过渡。
不需要额外的头尾素材；只需设置切点（帧），
CutToDissolve 就会在该切点周围创建叠化效果。
请注意，此效果不需要两个素材；只需一个已包含硬切的素材。Cut Point 参数是使其正常工作的关键；切点两侧的帧将被视为硬切，叠化效果将围绕它们创建。

在 Sapphire Time effects 子菜单中。

![CutToDissolve](../_static/CutToDissolve.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Cut Point** (Integer, Default: 5, Range: 2 or greater)
  硬切发生的帧。按 Find Cut 按钮可自动查找切点帧。

- **Find Cut** (Push-button)
  按此按钮可自动搜索素材并尝试查找硬切，从当前 Cut Frame 开始搜索。帧号将存储在上方的 Cut Point 中。如果几秒内未找到切点，搜索将停止。再次点击 Find Cut 可继续搜索。

- **Dissolve Length** (Integer, Default: 6, Range: 2 or greater)
  叠化的总长度。一半在切点左侧，一半在右侧。

- **Slow In Out** (Default: 2, Range: 0.1 to 10)
  设为 0 表示线性叠化，增大到 2 则产生更微妙的慢入慢出过渡效果。

- **Gamma** (Default: 1, Range: 0.1 to 10)
  设为 1 表示视频叠化，稍微增大可获得更具胶片感的效果。

- **Show** (Popup menu, Default: Result)
  此选项可帮助您找到切点帧；设为 Cut Frames 可查看基于 Cut Point 的最后一帧出画和第一帧入画的分屏视图。您也可以用它仅显示一侧或另一侧，以及插值的叠化帧。
  - **Result**: 显示结果素材，包含叠化效果。
  - **Cut Frames**: 显示两个切点帧的分屏视图，无论播放头当前在哪个位置。
  - **A**: 显示切点的 A 侧（出画）：就好像 B 侧是黑色的。
  - **B**: 显示切点的 B 侧（入画）：就好像 A 侧是黑色的。

- **Split For Cut** (Default: 0, Range: -1 to 1)
  在 Show:Cut Frames 模式下，分屏视图中切点的分割位置。通常此参数不起作用。

