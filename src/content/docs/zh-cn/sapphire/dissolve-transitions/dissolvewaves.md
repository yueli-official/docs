---
title: DissolveWaves
---

## S_DissolveWaves

使用波浪扭曲函数在两个输入素材之间进行转场。
第一个素材被扭曲消失并淡出，而第二个素材从扭曲中还原并淡入。
应对 Dissolve Percent 参数进行动画处理以控制转场速度。

在 Sapphire Transitions 效果子菜单中。

![DissolveWaves](../_static/DissolveWaves.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始转场。

- **Background**: 默认为无。以此素材结束转场。如果未提供此输入，将使用完全透明的背景，显示其后面的内容。请注意，除非提供此输入，否则背景在转场期间无法被扭曲。


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
  必须禁用 Auto Trans 才能使用此参数。它决定前景和背景输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制溶解的时间。如果 Slow In 和 Slow Out 参数为正值，它们也会在内部调整转场比例，使转场开始和/或结束更平滑。

- **Frequency** (Default: 3, Range: 0.01 or greater)
  波浪图案的频率。增加可获得更多更小的元素，减少可获得更少更大的元素。

- **Amplitude** (Default: 0.3, Range: any)
  缩放扭曲变形的程度。

- **Rel Amp2** (Default: -1, Range: any)
  第二个输入素材扭曲变形的相对幅度。如果此值为正而非负，素材将从相反方向解除扭曲。

- **Angle** (Default: 45, Range: any)
  用于擦除的整体波浪图案的旋转角度，单位为度。

- **Displace Angle** (Default: 90, Range: any)
  相对于波浪角度的扭曲方向，单位为度。0 给出压缩-扩展波浪，90 给出左右波浪。

- **Phase Start** (Default: 0, Range: any)
  波浪的相位偏移。波浪图案沿 Angle 方向平移此量。

- **Phase Speed** (Default: 0, Range: any)
  波浪的相位速度。如果非零，波浪图案将自动以此速度传播。

- **Slow In** (Default: 0.5, Range: 0 to 1)
  如果为正值，使转场开始更加缓和。

- **Slow Out** (Default: 0.5, Range: 0 to 1)
  如果为正值，使转场结束更加缓和。

- **Wrap** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  决定访问源图像边界之外区域的方法。
  - **No**: 边界外呈现黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复镜像副本。此方法通常边缘不太明显。

- **Filter** (Check-box, Default: on)
  如果启用，在重新采样时对图像进行自适应滤波。当图像的某些部分被扭曲得更小时，这可以提供更好的质量结果。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。如果您的图像在遮罩通道也有锐利边缘的地方有锐利的颜色变化，Normal 模式可能会给出更好的结果。

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  这 4 个参数（Crop Top、Crop Bottom、Crop Left 和 Crop Right）允许选择要处理的输入图像的矩形子区域。如果 Wrap 参数设为"No"，暴露的边框将是透明的。如果 Wrap 为"Tile"或"Reflect"，源图像将在新裁剪的边框上进行环绕以填充帧。这可以更容易地避免因扭曲边缘不良的图像而产生的伪像。
