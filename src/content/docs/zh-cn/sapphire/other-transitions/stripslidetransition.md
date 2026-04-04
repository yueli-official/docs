---
title: StripSlideTransition
---

## S_StripSlideTransition

通过将片段分割成条带并逐一滑出屏幕以显示传入片段来实现两个片段之间的转场。

在 Sapphire Transitions 效果子菜单中。

![StripSlideTransition](../_static/StripSlideTransition.jpg)


### Inputs:

- **From**: 当前图层。以此片段开始转场。

- **To**: 默认为无。以此片段结束转场。


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
  启用后，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过对 Strip Percent 参数进行动画设置来手动执行转场。

- **Amount** (Default: 0, Range: 0 to 1)
  必须禁用 Auto Trans 才能使用此参数。它确定 From 和 To 输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制擦除的时序。

- **Style** (Popup menu, Default: Slide Off)
  控制滑动应用于哪个片段。
  - **Slide Off**: 传出片段滑出以显示传入片段。
  - **Slide On**: 传入片段滑入覆盖传出片段。
  - **Side by Side**: 传出片段滑出的同时传入片段在其旁边滑入。

- **Motion Blur** (Default: 0.3, Range: 0 or greater)
  缩放要使用的运动模糊量。

- **Strip Size** (Default: 0.1, Range: 0.01 or greater)
  条带的宽度。此参数可能影响条带的时序，因此不建议对其进行动画设置。

- **Randomize Size** (Default: 0, Range: 0 or greater)
  随机使某些条带变大，某些条带变小。

- **Strip Angle** (Default: 0, Range: any)
  控制条带分割的角度以及滑动的方向。此参数可能影响条带的时序，因此不建议对其进行动画设置。

- **Strip Shift** (Default: 0, Range: any)
  调整条带边界的位置。此参数可能影响条带的时序，因此不建议对其进行动画设置。

- **Speed** (Default: 10, Range: 1 or greater)
  每个条带的移动速度。随着速度增大，条带之间的延迟变大。如果速度较低，许多条带将同时运动，产生波浪或涟漪效果。此参数影响条带的时序，因此不建议对其进行动画设置。

- **Slow Start** (Default: 1, Range: 0 to 1)
  控制每个条带移动时的加速度。如果设置为零，条带将以全速开始移动。值越大，条带开始移动越慢并逐渐加速到全速，从而产生更平滑的运动。

- **Order** (Popup menu, Default: Top Down)
  控制条带滑出屏幕的顺序。
  - **Top Down**: 从上到下依次进行。
  - **Bottom Up**: 从下到上依次进行。
  - **Random**: 随机顺序。
  - **Center Out**: 从中心向外，交替处理中心上方和下方的条带。
  - **Edges In**: 从边缘向内，交替处理中心上方和下方的条带。

- **Seed** (Default: 0.123, Range: 0 or greater)
  初始化随机条带大小和顺序的随机数生成器。实际种子值本身不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时，使用此选项可以稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将以预乘形式呈现，有时不太准确。

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  这 4 个参数 Crop Top、Crop Bottom、Crop Left 和 Crop Right 允许选择输入图像的矩形子区域进行处理。如果 Wrap 参数设置为"No"，暴露的边框将是透明的。如果 Wrap 设置为"Tile"或"Reflect"，源图像将在新裁剪的边框上进行包裹以填充画面。这可以更容易地避免由于扭曲具有不良边缘的图像而产生的瑕疵。
