---
title: WhipLash
---

## S_WhipLash

Swish3D 的纯 2D 版本，带有可选的 Whip Out 运动和 RGB 分离效果。在对两个输入片段执行 2D 运动的同时进行切换转场。在转场过程中，片段由 Rotate、Shift 和 Scale 参数进行变换。

在 Sapphire Transitions 效果子菜单中。

![WhipLash](../_static/WhipLash.jpg)


### Inputs:

- **Foreground**: 当前图层。以此片段开始转场。

- **Background**: 默认为无。以此片段结束转场。


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
  启用后，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过对 Whip Percent 参数进行动画设置来手动执行转场。

- **Wipe Percent** (Default: 0, Range: 0 to 1)
  必须禁用 Auto Trans 才能使用此参数。它确定 From 和 To 输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制擦除的时序。

- **Center** (X & Y, Default: [0 0], Range: any)
  缩放或旋转的中心位置。

- **Motion Blur** (Default: 1, Range: 0 or greater)
  缩放要使用的运动模糊量。

- **Rotate** (Default: 0, Range: any)
  按指定角度旋转，以度为单位。

- **Shift** (X & Y, Default: [-4 0], Range: any)
  水平或垂直平移。

- **Scale** (Default: 1, Range: 0 to 2)
  缩放片段的大小。

- **Whip Out** (Popup menu, Default: Smooth)
  转场结束时的运动方式。
  - **Smooth**: 平滑减速。
  - **Bounce**: 超调后弹跳至停止。
  - **Snap**: 超调后迅速卡到停止位置。

- **Mix RGB** (Default: 0, Range: 0 to 1)
  混入可选的 RGB 分离/模糊效果。

- **Blur Amount** (Default: 1.25, Range: 0 or greater)
  缩放模糊的宽度。

- **Angle** (Default: 0, Range: any)
  用于擦除的整体鞭动模式的旋转角度，以度为单位。

- **Shift RGB** (Default: 2, Range: any)
  沿模糊方向偏移图像。负偏移量将图像沿相反方向偏移。

- **Bias** (Default: 0.5, Range: 0 to 1)
  改变沿模糊路径上像素的权重，使其呈现出单方向的拖尾或条纹效果。值为 0.5 时所有像素权重均等。值为 1 时权重向模糊方向增加，值为 0 时效果相反。

- **Blur Red** (Default: 1, Range: 0 or greater)
  红色通道的模糊宽度，相对于 Blur Amount。

- **Blur Green** (Default: 0.5, Range: 0 or greater)
  绿色通道的模糊宽度，相对于 Blur Amount。

- **Blur Blue** (Default: 0, Range: 0 or greater)
  蓝色通道的模糊宽度，相对于 Blur Amount。

- **Shift Red** (Default: 0.5, Range: any)
  红色通道的额外偏移量。

- **Shift Green** (Default: 0.25, Range: any)
  绿色通道的额外偏移量。

- **Shift Blue** (Default: 0, Range: any)
  蓝色通道的额外偏移量。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果的亮度。

- **Offset Darks** (Default: 0, Range: -8 to 2)
  向结果的较暗区域添加此灰度值。可以为负值以增加对比度。

- **Edge Mode** (Popup menu, Default: Reflect)
  确定访问源图像之外区域时的行为。
  - **Transparent**: 源图像之外的区域被视为透明，这可能在图像边缘周围产生透明区域。选择此选项可获得最快的渲染速度。
  - **Repeat**: 重复图像边界之外的最后一个像素。
  - **Reflect**: 在边界之外反射图像。

- **Soft Borders** (Check-box, Default: off)
  启用后，在处理之前向输入图像添加透明边框。这允许结果包含超出原始图像大小的柔和边缘。关闭时，效果仅在画面内发生，结果将在边界处保留硬边。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时，使用此选项可以稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将以预乘形式呈现，有时不太准确。
