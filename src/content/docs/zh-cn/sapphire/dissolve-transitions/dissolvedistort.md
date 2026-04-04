---
title: DissolveDistort
---

## S_DissolveDistort

在两个输入素材之间转场，同时使用另一个素材的梯度对每个素材进行扭曲。第一个素材被扭曲并淡出，第二个素材被反扭曲到位并淡入。应通过动画 Dissolve Percent 参数来控制转场速度。请注意，必须提供背景输入，否则此效果只会执行简单的溶解而没有任何扭曲。

在 Sapphire Transitions 效果子菜单中。

![DissolveDistort](../_static/DissolveDistort.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始转场。

- **Background**: 默认为无。以此素材结束转场。如果未提供此输入，将使用完全透明的背景，显示其后面的内容。请注意，除非提供了此输入，否则背景在转场过程中无法被扭曲。


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
  必须禁用 Auto Trans 才能使用此参数。它决定前景和背景输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制溶解的时间。Slow In 和 Slow Out 参数如果为正值，也会在内部调整转场比例，以获得更平滑的转场开始和/或结束。

- **Amplitude** (Default: 1, Range: any)
  缩放应用于两个输入素材的扭曲量。也可以为负值，将膨胀变为收缩，反之亦然。

- **Rel Amp From** (Default: 1, Range: any)
  缩放 From 素材的相对扭曲振幅。

- **Rel Amp To** (Default: -1, Range: any)
  缩放 To 素材的相对扭曲振幅。

- **Smoothness** (Default: 0.25, Range: 0 or greater)
  按此数值平滑扭曲。增大可获得大尺度扭曲，减小可获得更精细的扭曲细节。

- **Rotate Warp Dir** (Default: 0, Range: any)
  旋转扭曲的方向。这可以使亮度相似的区域发生扭转，而不仅仅是膨胀或收缩。

- **Slow In** (Default: 0.2, Range: 0 to 1)
  如果为正值，使转场开始更加渐进。

- **Slow Out** (Default: 0.2, Range: 0 to 1)
  如果为正值，使转场结束更加渐进。

- **Wrap** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  决定访问源图像边界外部区域时的方法。
  - **No**: 边界外显示黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复图像的镜像副本。使用此方法时边缘通常不太明显。

- **Filter** (Check-box, Default: on)
  如果启用，在重新采样时对图像进行自适应滤波。这在图像被扭曲缩小时能提供更好的质量效果。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。
