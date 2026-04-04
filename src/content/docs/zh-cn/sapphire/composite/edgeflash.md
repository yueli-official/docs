---
title: EdgeFlash
---

## S_EdgeFlash

将前景片段的辉光添加到背景片段上，反之亦然，然后将前景合成到背景之上。这可以用于使合成看起来更自然，在图层之间产生光线闪烁效果，如同在胶片上一起曝光一样。

在 Sapphire Composite 效果子菜单中。

![EdgeFlash](../_static/EdgeFlash.jpg)


### Inputs:

- **Foreground**: 当前图层。用作前景的片段。

- **Background**: 默认为无。用作背景的片段。

- **Matte**: 默认为无。此输入的 Alpha 通道指定前景输入的不透明度。如果未提供此输入，则使用前景输入的 Alpha 通道代替。此输入可受 Invert Matte 或 Matte Use 参数的影响。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: Original)
  选择使用哪种处理方法
  - **Original**: 原始方法
  - **LightWrap**: 一种改进的方法，能更好地处理某些边缘条件

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此量模糊 Mocha 遮罩。这可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 遮罩的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  如果启用，则在应用效果之前反转 Mocha 遮罩的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素量膨胀或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在高质量模式下膨胀 Mocha 遮罩，以获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源片段。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  确定当两个遮罩同时提供给效果时，如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Fg Flash Amp** (Default: 0.8, Range: 0 or greater)
  从前景到背景的闪光量。

- **Bg Flash Amp** (Default: 0.8, Range: 0 or greater)
  从前景到背景的闪光量。

- **Flash Width** (Default: 0.088, Range: 0 or greater)
  闪光的宽度。此参数可通过 Flash Width 控件进行调整。

- **Fg Lights** (Default: 1, Range: any)
  按此值缩放前景输入。增大以获得更亮的结果。

- **Fg Darks** (Default: 0, Range: any)
  向前景输入的较暗区域添加此灰度值。可以为负值以增加对比度。

- **Fg Saturation** (Default: 1, Range: 0 or greater)
  缩放前景输入的色彩饱和度。增大以获得更浓烈的颜色。设为 0 以获得单色效果。

- **Bg Lights** (Default: 1, Range: any)
  按此值缩放背景输入。增大以获得更亮的结果。

- **Bg Darks** (Default: 0, Range: any)
  向背景输入的较暗区域添加此灰度值。可以为负值以增加对比度。

- **Bg Saturation** (Default: 1, Range: 0 or greater)
  缩放背景输入的色彩饱和度。增大以获得更浓烈的颜色。设为 0 以获得单色效果。

- **Output** (Popup menu, Default: Comp)
  在不同的输出选项之间进行选择。
  - **Foreground**: 仅输出带有来自背景闪光的前景片段。
  - **Background**: 仅输出带有来自前景闪光的背景片段。
  - **Comp**: 对两者都进行闪光处理，将前景合成到背景之上，并输出结果。

- **Subpixel Widths** (Check-box, Default: off)
  启用亚像素级别的闪光。使用此选项可获得更平滑的闪光宽度动画。

- **Comp Premult** (Check-box, Default: on)
  如果您提供了单独的遮罩输入，且前景像素值尚未被此遮罩预乘，请禁用此选项。

- **Matte Use** (Popup menu, Default: Alpha)
  确定如何使用遮罩输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Invert Matte** (Check-box, Default: off)
  如果启用，则反转输出遮罩的黑白。

- **Fg Flash Amp** (Default: 0.8, Range: 0 or greater)
  从前景到背景的闪光量。

- **Bg Flash Amp** (Default: 0.8, Range: 0 or greater)
  从前景到背景的闪光量。

- **Flash Width** (Default: 0.088, Range: 0 or greater)
  闪光的宽度。此参数可通过 Flash Width 控件进行调整。

- **Fg Lights** (Default: 1, Range: any)
  按此值缩放前景输入。增大以获得更亮的结果。

- **Fg Darks** (Default: 0, Range: any)
  向前景输入的较暗区域添加此灰度值。可以为负值以增加对比度。

- **Fg Saturation** (Default: 1, Range: 0 or greater)
  缩放前景输入的色彩饱和度。增大以获得更浓烈的颜色。设为 0 以获得单色效果。

- **Bg Lights** (Default: 1, Range: any)
  按此值缩放背景输入。增大以获得更亮的结果。

- **Bg Darks** (Default: 0, Range: any)
  向背景输入的较暗区域添加此灰度值。可以为负值以增加对比度。

- **Bg Saturation** (Default: 1, Range: 0 or greater)
  缩放背景输入的色彩饱和度。增大以获得更浓烈的颜色。设为 0 以获得单色效果。

- **Output** (Popup menu, Default: Comp)
  在不同的输出选项之间进行选择。
  - **Foreground**: 仅输出带有来自背景闪光的前景片段。
  - **Background**: 仅输出带有来自前景闪光的背景片段。
  - **Comp**: 对两者都进行闪光处理，将前景合成到背景之上，并输出结果。

- **Subpixel Widths** (Check-box, Default: off)
  启用亚像素级别的闪光。使用此选项可获得更平滑的闪光宽度动画。

- **Comp Premult** (Check-box, Default: on)
  如果您提供了单独的遮罩输入，且前景像素值尚未被此遮罩预乘，请禁用此选项。

- **Matte Use** (Popup menu, Default: Alpha)
  确定如何使用遮罩输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Invert Matte** (Check-box, Default: off)
  如果启用，则反转输出遮罩的黑白。

- **Shrink- Grow+** (Default: 0, Range: any)
  以近似像素为单位扩展遮罩边缘，如果为负值则收缩。

- **Edge Softness** (Default: 1, Range: 0.01 or greater)
  边缘的最终柔和度。

- **Post Blur** (Default: 0, Range: 0 or greater)
  如果为正值，则按此量模糊结果。这是柔化边缘的另一种方法。

- **Noise Amplitude** (Default: 0, Range: 0 or greater)
  添加到边缘的噪波纹理量。

- **Noise Width** (Default: 0.0224, Range: 0 or greater)
  在遮罩边缘包含噪波的区域宽度。除非 Noise Amplitude 为正值，否则此参数无效。

- **Frequency** (Default: 100, Range: 0.1 or greater)
  噪波的频率。增大以获得更细的颗粒噪波，减小以获得更粗的噪波。除非 Noise Amplitude 为正值，否则此参数无效。

- **Octaves** (Integer, Default: 1, Range: 1 to 10)
  噪波的叠加层数。每个八度的频率是前一个的两倍，幅度是前一个的一半。除非 Noise Amplitude 为正值，否则此参数无效。

- **Seed** (Default: 0.23, Range: 0 or greater)
  用于初始化随机数生成器。实际的种子值并不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Noise Shift** (X & Y, Default: [0 0], Range: any)
  噪波纹理的水平和垂直位移。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍快地渲染。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按照图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比普通模式稍快，但结果也将是预乘形式，有时不太准确。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍快地渲染。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按照图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比普通模式稍快，但结果也将是预乘形式，有时不太准确。

- **Show Flash Width** (Check-box, Default: on)
  打开或关闭用于调整 Flash Width 参数的屏幕用户界面。此参数仅出现在 AE 和 Premiere 中，因为这些平台支持屏幕控件。
