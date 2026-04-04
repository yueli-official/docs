---
title: StripSlide
---

## S_StripSlide

将素材分割成条带，逐条滑出屏幕以显示背景。

在 Sapphire Stylize 效果子菜单中。

![StripSlide](../_static/StripSlide.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Background**: 默认为无。当源素材滑走时显示此素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前对 Mocha 遮罩进行此量的模糊处理。可用于柔化遮罩的边缘或量化伪像，并平滑时间位移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 遮罩的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  如果启用，在应用效果之前反转 Mocha 遮罩的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前对 Mocha 遮罩进行此像素量的膨胀或腐蚀处理。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认 Fast 模式下快速调整还是在 High 质量模式下获得更好效果。
  - **Fast**: 以 Fast 模式进行 Dilate Mocha，用于快速调整。
  - **High**: 以 High 质量模式进行 Dilate Mocha，获得更好看的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩并将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果并仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Intersect)
  决定当效果同时提供 Mocha 遮罩和输入遮罩时如何合并它们。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Amount** (Default: 0.5, Range: 0 to 1)
  控制滑动效果的进度。为零时源素材完全可见，为一时背景完全可见。

- **Motion Blur** (Default: 0.3, Range: 0 or greater)
  缩放运动模糊的量。

- **Strip Size** (Default: 0.1, Range: 0.01 or greater)
  条带的宽度。此参数会影响条带的时序，因此不建议对其进行动画处理。

- **Randomize Size** (Default: 0, Range: 0 or greater)
  随机使某些条带变大或变小。

- **Strip Angle** (Default: 0, Range: any)
  控制条带分割的角度以及条带滑动的方向。此参数会影响条带的时序，因此不建议对其进行动画处理。

- **Strip Shift** (Default: 0, Range: any)
  调整条带边界的位置。此参数会影响条带的时序，因此不建议对其进行动画处理。

- **Speed** (Default: 10, Range: 1 or greater)
  每个条带移动的速度。速度越快，条带之间的延迟越大。如果速度较低，许多条带将同时运动，产生波浪或涟漪效果。此参数会影响条带的时序，因此不建议对其进行动画处理。

- **Slow Start** (Default: 1, Range: 0 to 1)
  控制每个条带移动时的加速度。如果设为零，条带将以全速开始移动。值越大，条带开始移动越慢并逐渐加速到全速，产生更平滑的运动效果。

- **Order** (Popup menu, Default: Top Down)
  控制条带滑出屏幕的顺序。
  - **Top Down**: 从上到下依次。
  - **Bottom Up**: 从下到上依次。
  - **Random**: 随机顺序。
  - **Center Out**: 从中心向外，交替处理中心上方和下方的条带。
  - **Edges In**: 从边缘向内，交替处理中心上方和下方的条带。

- **Seed** (Default: 0.123, Range: 0 or greater)
  初始化随机数生成器以随机确定条带大小和顺序。实际种子值并不重要，但不同的种子会给出不同的结果，相同的值应给出可重复的结果。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  这 4 个参数（Crop Top、Crop Bottom、Crop Left 和 Crop Right）允许选择要处理的输入图像的矩形子区域。如果 Wrap 参数设为"No"，暴露的边框将是透明的。如果 Wrap 为"Tile"或"Reflect"，源图像将在新裁剪的边框上进行环绕以填充帧。这可以更容易地避免因扭曲边缘不良的图像而产生的伪像。
