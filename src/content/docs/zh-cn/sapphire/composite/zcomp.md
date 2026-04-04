---
title: ZComp
---

## S_ZComp

根据两个深度图像的差异，将一个源输入叠加到第二个源输入的上方或下方。DepthA 输入应为与第一个输入中的对象对应的"z"深度图像，DepthB 应为与第二个输入中的对象对应的"z"深度图像。

在 Sapphire Composite 效果子菜单中。

![ZComp](../_static/ZComp.jpg)


### Inputs:

- **SourceA**: 当前图层。第一个输入图像。

- **SourceB**: 默认为无。第二个输入图像。

- **DepthA**: 默认为无。与 SourceA 中对象对应的深度图像。

- **DepthB**: 默认为无。与 SourceB 中对象对应的深度图像。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Anti Alias** (Default: 0, Range: 0 or greater)
  在源输入之间进行插值而非仅取较近者的深度差异量。以整个深度范围的分数指定：0 表示不进行抗锯齿，1 表示在整个深度范围内进行插值。

- **Invert Z** (Check-box, Default: off)
  通常较大的深度值（白色）被视为较远，较小的值（黑色）被视为较近。启用此选项时，这些深度值将反转。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍快地渲染。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按照图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比普通模式稍快，但结果也将是预乘形式，有时不太准确。
