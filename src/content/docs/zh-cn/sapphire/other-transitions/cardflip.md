---
title: CardFlip
---

## S_CardFlip

通过滑动或旋转传出片段来显示其后面的传入片段，实现两个片段之间的转场。Amount 参数应进行动画设置以控制转场速度。调整 Revolutions 和 Shift 可以产生不同类型的转场效果。

在 Sapphire Transitions 效果子菜单中。

![CardFlip](../_static/CardFlip.jpg)


### Inputs:

- **Foreground**: 当前图层。以此片段开始转场。

- **Background**: 默认为无。以此片段结束转场。如果未提供此输入，则使用完全透明的背景，显示其后面的内容。请注意，除非提供此输入，否则背景在转场过程中无法进行变形。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Transition Dir** (Popup menu, Default: Wipe Off to Bg)
  选择转场方向。
  - **Wipe Off to Bg**: 从当前图层转场到背景。
  - **Wipe On from Bg**: 从背景转场到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  启用后，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过对 Card Percent 参数进行动画设置来手动执行转场。

- **Amount** (Default: 0, Range: 0 to 1)
  必须禁用 Auto Trans 才能使用此参数。它确定 From 和 To 输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制擦除的时序。

- **Slow In** (Default: 0.5, Range: 0 to 1)
  如果为正值，使转场开始更加平缓。

- **Slow Out** (Default: 0.5, Range: 0 to 1)
  如果为正值，使转场结束更加平缓。

- **Revolutions** (Integer, Default: 1, Range: 0 or greater)
  转场期间片段翻转的次数。设置为 1 表示简单翻转，设置为 2 或更多表示旋转转场，设置为 0 表示滑动/洗牌。

- **Spin Direction** (Popup menu, Default: Left)
  旋转方向。
  - **Left**: 向左水平旋转。
  - **Right**: 向右水平旋转。
  - **Up**: 向上垂直旋转。
  - **Down**: 向下垂直旋转。

- **Shift** (Default: 0, Range: 0 or greater)
  在转场的前半段，使片段水平或垂直方向相互远离滑动，然后在后半段相互靠近。两个片段最终停在其起始位置。设置为 1 或更大的值可以防止片段在转场中点重叠。

- **Shift Direction** (Popup menu, Default: Left)
  偏移方向。
  - **Left**: 传出片段向左偏移，传入片段向右偏移。
  - **Right**: 传出片段向右偏移，传入片段向左偏移。
  - **Up**: 传出片段向上偏移，传入片段向下偏移。
  - **Down**: 传出片段向下偏移，传入片段向上偏移。

- **Perspective Amount** (Default: 1, Range: 0.25 to 4)
  控制片段翻转时的镜头伸缩程度。增大以获得更强的 3D 透视效果。

- **Shadow Color** (Default rgb: [0 0 0])
  前方片段投射到后方片段上的阴影颜色。

- **Shadow Opacity** (Default: 2, Range: 0 or greater)
  阴影的不透明度，使用接近 0 的值获得微妙的透明阴影，或使用接近 1.0 的值获得更强烈的阴影。

- **Shadow Blur** (Default: 0.088, Range: 0 or greater)
  确定阴影的柔和度。

- **Shadow Shift** (X & Y, Default: [0 0], Range: any)
  阴影的水平和垂直偏移量。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时，使用此选项可以稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将以预乘形式呈现，有时不太准确。
