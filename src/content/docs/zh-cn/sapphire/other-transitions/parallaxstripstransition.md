---
title: ParallaxStripsTransition
---

## S_ParallaxStripsTransition

应用一组 3D 折射玻璃条来分解图像。图像在每个条带内偏移，条带随时间移动。条带逐渐淡入或淡出，因此与源图像的转场是无缝的。
注意：由于您可以控制条带的大小，因此有可能设置的转场无法完全覆盖次要（传入或传出）片段，从而在开始或结束时产生"跳跃"。
要确保平滑的转场，请转到条带最大的转场末尾，选择 Show: Strips 模式，并调整条带大小以确保它们完全覆盖图像。这样后方片段就不会在该帧上透漏出来。或者，点击"Ensure Full Coverage"按钮，条带数量将自动调整为完全覆盖图像所需的最小值。

在 Sapphire Transitions 效果子菜单中。

![ParallaxStripsTransition](../_static/ParallaxStripsTransition.jpg)


### Inputs:

- **Foreground**: 当前图层。以此片段开始转场。

- **Background**: 默认为无。以此片段结束转场。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: Rectangular Strips)
  在 ParallaxStripsTransition 的多个变体之间选择。
  - **Rectangular Strips**: 条带在转场中滑动并溶解为次要片段。
  - **Linear Strips**: 全高条带逐渐展开以显示次要片段。

- **Dissolve Speed** (Default: 3, Range: 1 or greater)
  Panning Strips 中 From 和 To 片段之间溶解的速度。

- **Transition Dir** (Popup menu, Default: Transition Off to Bg)
  选择转场方向。
  - **Transition Off to Bg**: 从当前图层转场到背景。
  - **Transition On from Bg**: 从背景转场到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  启用后，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过对 Parallax Percent 参数进行动画设置来手动执行转场。

- **Trans Amount** (Default: 0.5, Range: 0 to 1)
  From 和 To 输入之间的转场比例。值为 0 时仅显示 From 输入，值为 1 时仅显示 To 输入。默认情况下，此参数将自动从 0 动画到 1 以执行完整的转场。

- **Ensure Full Coverage** (Push-button)
  按下此按钮可将条带数量调整为完全覆盖最后一帧所需的最小值。

- **N Strips** (Integer, Default: 50, Range: 1 to 1000)
  要应用的折射条带数量。条带在整个图像上随机定位。

- **Size** (Default: 0.35, Range: 0 or greater)
  条带的大小，以图像宽度为单位。

- **Rel Height** (Default: 0.3, Range: 0.001 or greater)
  条带的高度，相对于其宽度。增大以使条带更高。

- **Size Vary** (X & Y, Default: [0.1 0.1], Range: 0 to 1)
  增大以使每个条带随机变大或变小。

- **Angle** (Default: 0, Range: any)
  条带的角度；0 为水平。条带沿其角度移动，并沿相同角度偏移图像。

- **Depth** (Default: 2, Range: 0 or greater)
  使最前面的条带更大并移动得更快，使其看起来在前方，产生 3D 效果。

- **Strip Speed** (Default: 0.4, Range: any)
  设置条带沿其主轴移动的速度。请注意，这不会影响图像在条带内的折射或偏移方式，只影响条带本身的移动速度。

- **Strip Speed Vary** (Default: 0, Range: 0 or greater)
  增大以使每个条带的速度有一些随机性。

- **Shift Amount** (Default: 0.6, Range: any)
  设置图像在每个条带内偏移或折射的程度。偏移始终沿条带的主轴方向。随着效果的进行，偏移量逐渐降至零，无缝过渡到原始片段。有关详细信息，请参见 Fade 和 Slow Fade 参数。

- **Shift Vary** (Default: 0, Range: 0 or greater)
  增大以使每个条带中的偏移量更加随机。

- **All Strips Shift** (X & Y, Default: [0 0], Range: any)
  在屏幕上移动所有条带。

- **Z Dist** (Default: 1, Range: 0.001 or greater)
  在应用视差条带之前，对源图像进行缩放。

- **Show** (Popup menu, Default: Result)
  显示效果结果或条带本身，这在效果设置期间很有用。
  - **Result**: 显示效果的结果。
  - **Strips Over Source**: 将每个条带显示为灰色矩形，亮度由深度设置。未覆盖的区域显示源图像。
  - **Strips Over Black**: 将每个条带显示为灰色矩形，亮度由深度设置。未覆盖的区域显示为黑色。

- **Slow Fade** (Default: 0.9, Range: 0 to 2)
  增大以使淡入或淡出更慢。设置为 0 表示线性淡化。

- **Slow Grow** (Default: 0.9, Range: 0 to 2)
  增大以使条带开始增长更慢，以获得更好的效果。设置为 0 表示在效果过程中线性增长。

- **Wrap** (Popup menu, Default: Reflect)
  确定访问源图像边界之外区域的方法。
  - **No**: 边界之外显示黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复镜像副本。使用此方法时边缘通常不太明显。

- **Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化随机数生成器。实际种子值本身不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Flip Tiles** (Check-box, Default: off)
  如有需要，垂直翻转图块以获得一致的外观。
