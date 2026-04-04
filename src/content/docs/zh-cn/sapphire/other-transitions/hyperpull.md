---
title: HyperPull
---

## S_HyperPull

在溶解到背景之前，将前景在 Z 轴空间中拉远，并附带一些额外的颜色和抖动效果以增添风格。

在 Sapphire Transitions 效果子菜单中。

![HyperPull](../_static/HyperPull.jpg)


### Inputs:

- **Foreground**: 当前图层。以此片段开始转场。

- **Background**: 默认为无。以此片段结束转场。如果未提供此输入，则使用完全透明的背景，显示其后面的内容。请注意，除非提供此输入，否则背景在转场过程中无法进行变形。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Transition Dir** (Popup menu, Default: Dissolve Off to Bg)
  选择转场方向。
  - **Dissolve Off to Bg**: 从当前图层转场到背景。
  - **Dissolve On from Bg**: 从背景转场到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  启用后，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过对 Dissolve Percent 参数进行动画设置来手动执行转场。

- **Dissolve Percent** (Default: 0, Range: 0 to 1)
  必须禁用 Auto Trans 才能使用此参数。它确定前景和背景输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制溶解的时序。Slow In 和 Slow Out 参数（如果为正值）也会在内部调整转场比例，以实现更平滑的转场开始和/或结束。

- **Slow In** (Default: 1, Range: 0 to 1)
  如果为正值，使转场开始更加平缓。

- **Slow Out** (Default: 1, Range: 0 to 1)
  如果为正值，使转场结束更加平缓。

- **Dissolve Speed** (Default: 5, Range: 1 or greater)
  前景和背景之间溶解的速度。

- **Z Dist From** (Default: 5, Range: 0.001 or greater)
  缩放前景的"距离"。大于 1.0 的值将其移得更远并使其变小。小于 1.0 的值将图像移得更近并放大。轻微放大有时可以用来隐藏边缘瑕疵。

- **Center XY From** (X & Y, Default: [0 0], Range: any)
  鱼眼变形函数的中心，以屏幕坐标表示，相对于前景的中心。

- **Rotate From** (Default: 0, Range: any)
  围绕中心位置旋转前景指定的度数。随着转场进行，角度逐渐增加到此值。

- **Wrap From** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  确定访问前景图像边界之外区域的方法。
  - **No**: 边界之外显示黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复镜像副本。使用此方法时边缘通常不太明显。

- **Z Dist To** (Default: 0.001, Range: 0.001 or greater)
  缩放背景的"距离"。大于 1.0 的值将其移得更远并使其变小。小于 1.0 的值将图像移得更近并放大。轻微放大有时可以用来隐藏边缘瑕疵。

- **Center XY To** (X & Y, Default: [0 0], Range: any)
  鱼眼变形函数的中心，以屏幕坐标表示，相对于背景的中心。

- **Rotate To** (Default: 0, Range: any)
  围绕中心位置旋转背景指定的度数。随着转场进行，角度逐渐增加到此值。

- **Wrap To** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  确定访问背景图像边界之外区域的方法。
  - **No**: 边界之外显示黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复镜像副本。使用此方法时边缘通常不太明显。

- **Motion Blur** (Check-box, Default: on)
  启用运动模糊。

- **Blur From Z Dist** (Default: 0.7, Range: 0.001 or greater)
  From 变换的"距离"。增大以缩小，减小以放大。

- **Blur From Rotate** (Default: 0, Range: any)
  From 变换围绕中心的旋转角度，以度为单位。

- **Blur To Z Dist** (Default: 0.9, Range: 0.001 or greater)
  To 变换的"距离"。增大以缩小，减小以放大。

- **Blur To Rotate** (Default: 0, Range: any)
  To 变换围绕中心的旋转角度，以度为单位。请注意，如果 From 和 To 的旋转角度差异很大，它们之间的插值将变得不太准确。

- **Camera Shake** (Check-box, Default: off)
  启用摄像机抖动。

- **Amplitude** (Default: 2, Range: 0 or greater)
  缩放抖动运动的振幅。

- **Frequency** (Default: 2, Range: 0 or greater)
  增大以获得更快的抖动，减小以获得更慢的抖动。

- **Glow Brights** (Check-box, Default: on)
  启用亮部辉光。

- **Glow Brightness** (Default: 3, Range: 0 or greater)
  辉光的整体最大亮度。

- **Glow Threshold** (Default: 0.2, Range: 0 or greater)
  源片段中亮于此值的部分会产生辉光。值为 0.9 时只有最亮的点发光。值为 0 时每个非黑色区域都发光。

- **Glow Width** (Default: 0, Range: 0 or greater)
  辉光的宽度。

- **Width X** (Default: 1, Range: 0 or greater)
  缩放水平辉光宽度。设置为 0 仅保留垂直方向。

- **Width Y** (Default: 1, Range: 0 or greater)
  缩放垂直辉光宽度。设置为 0 仅保留水平方向。

- **Glow Darks** (Check-box, Default: off)
  启用暗部辉光。

- **Darkness** (Default: 0.5, Range: 0 or greater)
  暗部辉光的强度。

- **Dark Threshold** (Default: 0.5, Range: 0 or greater)
  源片段中亮于此值的部分会产生辉光。值为 0.9 时只有最亮的点发光。值为 0 时每个非黑色区域都发光。

- **Dark Width** (Default: 0, Range: 0 or greater)
  缩放暗部辉光的距离。请注意，即使辉光宽度为零仍会影响暗区；如果您想原样传递源图像，请将 darkness 参数设置为零。

- **Dark Width X** (Default: 1, Range: 0 or greater)
  缩放水平暗部宽度。设置为 0 仅保留垂直方向。

- **Dark Width Y** (Default: 1, Range: 0 or greater)
  缩放垂直暗部宽度。设置为 0 仅保留水平方向。

- **Warp Chroma** (Check-box, Default: on)
  启用色度变形。

- **Warp Amount** (Default: 0.6, Range: 0 or greater)
  调整结果的整体色度变形量。随着转场进行，变形量逐渐增加到此值。将其设置为零可禁用变形并保持图像色度不变。

- **Steps** (Integer, Default: 10, Range: 3 to 100)
  沿色度变形光谱包含的颜色采样数量。更多步数产生更平滑的结果，但需要更多处理时间。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果的亮度。

- **Distortion Amount** (Default: -1, Range: any)
  鱼眼变形的振幅。

- **Distort RGB Amount** (Default: 0.2, Range: any)
  缩放所有通道的镜头畸变程度。设为负值可反转畸变方向。
