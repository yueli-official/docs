---
title: DissolveLuma
---

## S_DissolveLuma

使用从两个素材的亮度值衍生的图案在它们之间进行转场。
一个素材通常看起来像是从另一个素材中浮现出来。
应对 Dissolve Percent 参数进行动画处理以控制转场速度。

在 Sapphire Transitions 效果子菜单中。

![DissolveLuma](../_static/DissolveLuma.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始转场。

- **Background**: 默认为无。以此素材结束转场。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Transition Dir** (Popup menu, Default: Dissolve Off to Bg)
  选择转场的方向。
  - **Dissolve Off to Bg**: 从当前图层转场到背景。
  - **Dissolve On from Bg**: 从背景转场到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  如果启用，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过动画 Dissolve Percent 参数手动执行转场。

- **Dissolve Percent** (Default: 0, Range: 0 to 1)
  必须禁用 Auto Trans 才能使用此参数。它决定前景和背景输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制溶解的时间。

- **Softness** (Default: 0.1, Range: 0 to 1)
  增加可获得更柔和、更缓慢的转场。

- **Use Luma Of** (Popup menu, Default: Difference)
  决定如何从素材的亮度值生成转场图案。
  - **Difference**: 相似区域优先转场，差异区域最后转场。
  - **Subtract**: 第一个素材较亮的区域优先转场，第二个素材较亮的区域最后转场。
  - **Mult**: 两个图像都较亮的区域优先转场，任一图像较暗的区域最后转场。
  - **Screen**: 任一图像较亮的区域优先转场，两个图像都较暗的区域最后转场。
  - **Foreground**: 第一个素材的暗区域最先消失，亮区域最后消失。
  - **Background**: 第二个素材的亮区域最先出现，暗区域最后出现。

- **Invert Pattern** (Check-box, Default: off)
  如果启用，转场图案在时间上反转。

- **Smooth Pattern** (Default: 0, Range: 0 or greater)
  如果为正值，则对转场图案应用模糊。这可以减少噪点并使转场线条具有更清晰的边缘。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。
