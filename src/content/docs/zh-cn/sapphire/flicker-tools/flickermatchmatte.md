---
title: FlickerMatchMatte
---

## S_FlickerMatchMatte

使用第二个匹配片段的闪烁，在遮罩指定的区域内为源片段添加闪烁。要使用此效果，选择一帧您希望源亮度保持不变的帧，并点击 Set Match Level 按钮。处理其他帧时，源亮度将按遮罩内匹配片段的平均亮度相对于匹配级别进行缩放。

在 Sapphire Time 效果子菜单中。

![FlickerMatchMatte](../_static/FlickerMatch.jpg)


### Inputs:

- **Source**: 当前图层。要添加闪烁的片段。

- **Match**: 默认为无。要从中复制闪烁的片段。

- **Matte**: 默认为无。此片段指定从哪些源区域测量闪烁。如果未提供此输入，则使用匹配输入的 Alpha 通道作为遮罩。可以使用 Invert Matte 参数进行反转。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Match Level** (Default: 0.5, Range: 0.01 or greater)
  遮罩内匹配片段的平均亮度，在此亮度下源输入保持不变。

- **Set Match Level** (Push-button)
  按下此按钮会将 Match Level 参数设置为当前帧遮罩内匹配片段的平均亮度。这会使输出在此帧等于源。此按钮本身不保留任何值，按下后会立即恢复关闭状态。

- **Matte Use** (Popup menu, Default: Alpha)
  确定如何使用遮罩输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Invert Matte** (Check-box, Default: off)
  如果开启，将反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则此参数无效。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。
