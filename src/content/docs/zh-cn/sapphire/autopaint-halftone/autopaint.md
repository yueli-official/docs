---
title: AutoPaint
---

## S_AutoPaint

生成源素材的"画笔绘制"版本。使用 Frequency 和 Stroke Length 参数来调整笔触的密度和形状。如果您希望每帧重新随机化笔触图案，可以将 Jitter Frames 设置为 1。

位于 Sapphire Stylize 效果子菜单中。

![AutoPaint](../_static/AutoPaint.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。在结果和源输入之间进行插值。白色区域使用效果结果。黑色区域使用源素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此数值模糊 Mocha 遮罩。可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 遮罩的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  如果启用，在应用效果之前会反转 Mocha 遮罩的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素量扩展或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下扩展 Mocha 遮罩，用于快速调整。
  - **High**: 在 High 质量模式下扩展 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用到整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩同时提供给效果时，决定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Style** (Popup menu, Default: Van Gogh)
  选择笔触的风格。
  - **Van Gogh**: 笔触方向与图像中发现的边缘对齐。
  - **Hairy Paint**: 笔触与图像中的边缘垂直。
  - **Pointalize**: 笔触为无方向的蜂窝状尖锐形状。

- **Frequency** (Default: 50, Range: 0.1 or greater)
  画面中笔触的密度。增大可获得更小的笔触。

- **Stroke Length** (Default: 2, Range: any)
  确定笔触沿源素材边缘方向的长度。如果为负值，可以在 VanGogh 和 HairyPaint 风格之间切换。

- **Stroke Align** (Default: 0.2, Range: 0 or greater)
  增大以平滑笔触方向，使相邻笔触更加平行。

- **Smooth Colors** (Default: 0, Range: 0 or greater)
  在生成笔触之前按此数值模糊源素材。增大可使相邻笔触的颜色更加一致。

- **Seed** (Default: 0, Range: 0 or greater)
  用于初始化随机数生成器。实际种子值并不重要，但不同的种子会产生不同的结果，相同的值应该会产生可重复的结果。

- **Jitter Frames** (Integer, Default: 0, Range: 0 or greater)
  如果为 0，笔触的位置在每一帧处理时保持不变。如果为 1，笔触位置在每帧重新随机化。如果为 2，则每隔一帧重新随机化，以此类推。

- **Sharpen** (Default: 1, Range: any)
  应用的后处理锐化量。

- **Sharpen Width** (Default: 0.1, Range: 0 or greater)
  应用后处理锐化滤镜的宽度，相对于笔触大小。较高的值影响边缘周围更宽的区域，较低的值仅影响锐利边缘附近的区域。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已为预乘形式处理（颜色已按不透明度缩放）。此选项的渲染速度也略快于 Normal 模式，但结果也将为预乘形式，这有时不太准确。

- **Mix With Source** (Default: 0, Range: 0 to 1)
  在结果（0）和原始源素材（1）之间进行插值。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用 Mask 输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊遮罩输入。可提供遮罩区域和非遮罩区域之间更平滑的过渡。除非提供了遮罩输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果开启，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则无效。
