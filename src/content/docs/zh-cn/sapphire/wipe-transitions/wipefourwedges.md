---
title: WipeFourWedges
---

## S_WipeFourWedges

使用四个楔形合并成"X"形的图案在两个输入素材之间执行擦除转场。应对 Wipe Percent 参数设置动画以控制转场速度。增大 Border Width 参数可在擦除转场边缘绘制边框。

在 Sapphire Transitions 效果子菜单中。

![WipeFourWedges](../_static/WipeFourWedges.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始转场。

- **Background**: 默认为无。以此素材结束转场。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Transition Dir** (Popup menu, Default: Wipe Off to Bg)
  选择转场的方向。
  - **Wipe Off to Bg**: 从当前图层转场到背景。
  - **Wipe On from Bg**: 从背景转场到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  如果启用，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过对 Wipe Percent 参数设置动画来手动执行转场。

- **Wipe Percent** (Default: 0, Range: 0 to 1)
  必须禁用 Auto Trans 才能使用此参数。它决定了 From 和 To 输入之间的转场比例，通常应从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线以更精细地控制擦除时序。

- **Wipe Direction** (Popup menu, Default: Wedge In)
  选择楔形运动的方向。
  - **Wedge In**: 楔形包含第一个图像并向内缩小。
  - **Wedge Out**: 楔形包含第二个图像并向外扩大。

- **Edge Softness** (Default: 0, Range: 0 or greater)
  转场边缘的宽度。较大的值将导致擦除图案中更柔和、更不明显的边缘。

- **Angle** (Default: 0, Range: any)
  楔形图案的旋转角度（以度为单位）。

- **Aspect Scale** (Default: 1, Range: 0.01 or greater)
  缩放楔形图案的宽高比。增大可在水平方向拉伸形状。

- **Border Width** (Default: 0, Range: 0 or greater)
  如果为正值，将在擦除转场边缘使用以下边框颜色、不透明度、柔和度和偏移参数绘制彩色边框。

- **Border Color** (Default rgb: [0.75 0 0])
  边框的颜色。除非 Border Width 为正值，否则此参数无效。

- **Border Opacity** (Default: 1, Range: 0 to 1)
  边框的不透明度。减小可使边框变为透明，允许其下方的图像透过显示。除非 Border Width 为正值，否则此参数无效。

- **Border Softness** (Default: 0, Range: 0 or greater)
  边框边缘的柔和度。除非 Border Width 为正值，否则此参数无效。

- **Border Shift** (Default: 0, Range: any)
  将边框向转场边缘的前方或后方偏移。除非 Border Width 为正值，否则此参数无效。

- **Border Glow** (Default: 0, Range: 0 or greater)
  沿擦除边框添加辉光。该值决定辉光的亮度。

- **Glow Width** (Default: 0.1, Range: 0 or greater)
  辉光边框的宽度。

- **Width Red** (Default: 1, Range: 0 or greater)
  缩放红色辉光宽度。如果红色、绿色和蓝色宽度都相等，辉光将匹配 Glow Color。否则将产生不同颜色的边缘。

- **Width Green** (Default: 1.2, Range: 0 or greater)
  缩放绿色辉光宽度。

- **Width Blue** (Default: 1.4, Range: 0 or greater)
  缩放蓝色辉光宽度。

- **Glow Color** (Default rgb: [1 1 1])
  辉光边框的颜色。

- **Noise Amp** (Default: 1, Range: 0 or greater)
  添加到辉光边框的噪声量。

- **Noise Freq** (Default: 16, Range: 0.1 to 20)
  噪声的空间频率。

- **Noise Speed** (Default: 2, Range: any)
  噪声随时间变化或沸腾的速度。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（Alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按照图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项也比 Normal 模式渲染速度稍快，但结果也将是预乘形式，这有时不太正确。

- **Show Angle** (Check-box, Default: off)
  打开或关闭用于调整 Angle 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中显示，因为这些软件支持屏幕控件。

- **Show Glow Width** (Check-box, Default: off)
  打开或关闭用于调整 Glow Width 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中显示，因为这些软件支持屏幕控件。
