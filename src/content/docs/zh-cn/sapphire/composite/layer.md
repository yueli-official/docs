---
title: Layer
---

## S_Layer

使用多种混合操作之一将前景图像叠加到背景之上。每个输入的颜色也可以通过亮部、暗部和饱和度参数进行调整。

在 Sapphire Composite 效果子菜单中。

![Layer](../_static/Layer.jpg)


### Inputs:

- **Foreground**: 当前图层。用作前景的片段。

- **Background**: 默认为无。用作背景的片段。

- **Matte**: 默认为无。指定前景片段的不透明度。如果未提供此输入，则使用前景的 Alpha 通道代替。这些值在使用前会按 Fg Opacity 参数进行缩放。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: Normal)
  确定使用哪种混合方法来组合前景和背景像素颜色。
  - **Normal**: 正常合成。除非不透明度低于 1.0 或提供了 Alpha 通道，否则结果就是前景。
  - **Dissolve**: 随机用前景像素替换背景像素。不透明度决定概率，因此不透明度值越高，前景替换背景的可能性越大。
  - **Multiply**: 可用作遮罩图像的"交集"操作。白色是 Multiply 的恒等元素，其中一个图像包含白色时，另一个不受影响，因此结果仅在两个输入都为白色的地方包含白色。
  - **Screen**: 可用于组合两个片段的亮部区域。也可用作遮罩图像的"并集"操作。黑色是 Screen 的恒等元素，其中一个图像包含黑色时，另一个不受影响，因此任一输入图像为白色的地方结果都为白色。
  - **Overlay**: 使用叠加函数组合前景和背景。
  - **Soft Light**: 根据前景使背景变亮或变暗。
  - **Hard Light**: 类似于叠加，但前景和背景互换。
  - **Color Dodge**: 根据前景使背景变亮。
  - **Color Burn**: 根据前景使背景变暗。
  - **Darken**: 前景和背景的最小值。也可用作"交集"操作，结果与 Multiply 略有不同。
  - **Lighten**: 前景和背景的最大值。也可用作"并集"操作，结果与 Screen 略有不同。
  - **Add**: 将前景添加到背景。
  - **Subtract**: 从背景中减去前景。
  - **Difference**: 类似于 Subtract，但使用结果的绝对值，这往往会使更多结果颜色保持在范围内。可用于选择两个遮罩图像中一个为白色而另一个不为白色的区域。
  - **Exclusion**: 类似于 Difference，但结果更平滑。
  - **Hue**: 将前景的色相与背景的饱和度和亮度组合。
  - **Saturation**: 将前景的饱和度与背景的色相和亮度组合。
  - **Chroma**: 将前景的色相和饱和度与背景的亮度组合。
  - **Luminance**: 将前景的亮度与背景的色相和饱和度组合。
  - **Linear Dodge**: 将前景和背景相加，并将结果限制在白色。
  - **Linear Burn**: 将前景和背景相加，但进行偏移使结果更暗。类似于 Multiply，与白色组合不产生变化，与黑色组合得到黑色。
  - **Linear Light**: 根据前景是否大于或小于 50% 灰度，执行线性加深或线性减淡。
  - **Vivid Light**: 根据前景是否大于或小于 50% 灰度，执行颜色加深或颜色减淡。
  - **Pin Light**: 根据前景是否大于或小于 50% 灰度，执行变亮或变暗。

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

- **Swap Inputs** (Check-box, Default: off)
  如果启用，则有效地交换背景和前景输入，对于减法等非交换操作很有帮助。请注意，这也会导致标记为"Front"的参数影响"Back"输入，反之亦然。

- **Fg Opacity** (Default: 1, Range: 0 to 1)
  缩放效果的不透明度。降低此值时，结果会接近背景。为零时，结果将等于背景。

- **Fg Lights** (Default: 1, Range: any)
  在执行效果之前缩放前景。

- **Fg Darks** (Default: 0, Range: any)
  在执行效果之前偏移前景的较暗区域。可以为负值以增加对比度。

- **Fg Saturation** (Default: 1, Range: any)
  在执行效果之前缩放前景的色彩饱和度。0.0 使其变为单色，1.0 无效果。

- **Fg Hue Shift** (Default: 0, Range: any)
  偏移前景片段颜色的色相，以从红到绿到蓝再到红的圈数为单位。

- **Fg Blur** (Default: 0, Range: 0 or greater)
  前景的模糊量。

- **Bg Lights** (Default: 1, Range: any)
  在执行效果之前缩放背景。

- **Bg Darks** (Default: 0, Range: any)
  在执行效果之前偏移背景的较暗区域。可以为负值以增加对比度。

- **Bg Saturation** (Default: 1, Range: any)
  在执行效果之前缩放背景的色彩饱和度。0.0 使其变为单色，1.0 无效果。

- **Bg Hue Shift** (Default: 0, Range: any)
  偏移背景片段颜色的色相，以从红到绿到蓝再到红的圈数为单位。

- **Bg Blur** (Default: 0, Range: 0 or greater)
  背景的模糊量。

- **Result Lights** (Default: 1, Range: any)
  在执行效果之后缩放结果。

- **Result Darks** (Default: 0, Range: any)
  在执行效果之后偏移结果的较暗区域。可以为负值以增加对比度。

- **Result Saturation** (Default: 1, Range: any)
  在执行效果之后缩放结果的色彩饱和度。0.0 使其变为单色，1.0 无效果。

- **Result Hue Shift** (Default: 0, Range: any)
  偏移结果颜色的色相，以从红到绿到蓝再到红的圈数为单位。

- **Matte Use** (Popup menu, Default: Alpha)
  确定如何使用遮罩输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。
  - **All Opaque**: 使用 1.0 的 Alpha 值，如同遮罩完全不透明。

- **Blur Subpixel** (Check-box, Default: on)
  启用亚像素级别的模糊。使用此选项可获得更平滑的 Blur Front 和 Blur Back 参数动画。

- **Soft Borders** (Check-box, Default: off)
  如果启用，在处理之前向输入图像添加透明边框。这允许结果包含超出原始图像大小的柔和边缘。关闭时，效果仅发生在帧内，结果将在边界处保留边缘。

- **Comp Premult** (Check-box, Default: on)
  如果您提供了单独的遮罩输入，且前景像素值尚未被此遮罩预乘，请禁用此选项。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍快地渲染。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按照图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比普通模式稍快，但结果也将是预乘形式，有时不太准确。
