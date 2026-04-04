---
title: ZGlow
---

## S_ZGlow

根据 ZBuffer 输入的深度值，以不同的宽度对源素材的区域进行辉光处理。将输入分成若干层，并根据 Width Near、Width Far、Brightness Near 和 Brightness Far 参数应用不同程度的辉光。

在 Sapphire Lighting 效果子菜单中。

![ZGlow](../_static/ZGlow.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **ZBuffer**: 默认为无。包含每个源像素深度值的输入素材。这些值应在黑色到白色的范围内，最好没有抗锯齿。通常黑色对应最远的物体，白色对应最近的物体，但可以通过 Z Buffer 参数进行调整。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Brightness** (Default: 2, Range: 0 or greater)
  缩放所有辉光的亮度。

- **Color** (Default rgb: [1 1 1])
  缩放辉光的颜色。辉光的颜色和亮度也受源素材的影响。

- **Width Near** (Default: 0.0336, Range: 0 or greater)
  近处（靠近）物体的辉光宽度。

- **Width Far** (Default: 0.4, Range: 0 or greater)
  远处物体的辉光宽度。

- **Threshold** (Default: 0.5, Range: 0 or greater)
  从源素材中亮度超过此值的位置生成辉光。值为 0.9 时仅在最亮的位置产生辉光。值为 0 时在每个非黑色区域都产生辉光。

- **Threshold Add Color** (Default rgb: [0 0 0])
  可用于提高特定颜色的阈值，从而减少源素材中包含该颜色的区域所产生的辉光。

- **Z Buffer Type** (Popup menu, Default: White is Near)
  如何解读 Z 缓冲区中的值。
  - **Black is Near**: Z 缓冲区中的黑色像素表示该点的物体距离近（靠近你），白色表示远。
  - **White is Near**: Z 缓冲区中的白色像素表示该点的物体距离近（靠近你），黑色表示远。

- **Z Min** (Default: 0, Range: 0 to 1)
  将所有 Z 值钳制到此最小边界。使用此参数可在图像中比 Z Min 更近的所有部分创建恒定辉光。

- **Z Max** (Default: 1, Range: 0 to 1)
  将所有 Z 值钳制到此最大边界。使用此参数可在图像中比 Z Max 更远的所有部分创建恒定辉光。

- **Layers** (Integer, Default: 5, Range: 2 to 50)
  将源素材分成的深度层数。更多层需要更多处理，但在 Z 方向上产生更平滑的结果。有时需要更多层以避免层之间出现可见的接缝。

- **Brightness Near** (Default: 1, Range: 0 or greater)
  缩放近处物体的辉光亮度。

- **Color Near** (Default rgb: [1 1 1])
  缩放近处物体的辉光颜色。

- **Width Red Near** (Default: 1, Range: 0 or greater)
  缩放近处物体的红色辉光宽度。

- **Width Green Near** (Default: 1, Range: 0 or greater)
  缩放近处物体的绿色辉光宽度。

- **Width Blue Near** (Default: 1, Range: 0 or greater)
  缩放近处物体的蓝色辉光宽度。

- **Brightness Far** (Default: 1, Range: 0 or greater)
  缩放远处物体的辉光亮度。

- **Color Far** (Default rgb: [1 1 1])
  缩放远处物体的辉光颜色。

- **Width Red Far** (Default: 1, Range: 0 or greater)
  缩放远处物体的红色辉光宽度。

- **Width Green Far** (Default: 1, Range: 0 or greater)
  缩放远处物体的绿色辉光宽度。

- **Width Blue Far** (Default: 1, Range: 0 or greater)
  缩放远处物体的蓝色辉光宽度。

- **Width X** (Default: 1, Range: 0 or greater)
  缩放水平辉光宽度。设为 0 则仅显示垂直方向。

- **Width Y** (Default: 1, Range: 0 or greater)
  缩放垂直辉光宽度。设为 0 则仅显示水平方向。

- **Width Red** (Default: 1, Range: 0 or greater)
  缩放红色辉光宽度。如果红、绿、蓝宽度相等，辉光将与源素材的颜色一致。如果不相等，辉光颜色将随距离变化。

- **Width Green** (Default: 1.2, Range: 0 or greater)
  缩放绿色辉光宽度。

- **Width Blue** (Default: 1.4, Range: 0 or greater)
  缩放蓝色辉光宽度。

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  如果此值为正，输出的 Alpha 通道将包含来自辉光的一些不透明度。红、绿、蓝辉光亮度的最大值按此值缩放，并在每个像素处与背景 Alpha 合成。

- **Source Opacity** (Default: 1, Range: 0 to 1)
  缩放源输入与辉光合成时的不透明度。这不影响辉光本身的生成。

- **Zbuffer Use** (Popup menu, Default: Luma)
  决定 ZBuffer 输入通道如何生成单色深度图像。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Expand Borders** (Check-box, Default: off)
  如果启用，在处理前向输入图像添加透明边框。这允许结果包含超出原始图像大小的柔和边缘。关闭时，效果仅在画面内发生，结果将在边界处保留硬边。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Show Width Near** (Check-box, Default: on)
  开启或关闭用于调整 Width Near 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。

- **Show Width Far** (Check-box, Default: on)
  开启或关闭用于调整 Width Far 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
