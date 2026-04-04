---
title: GlowNoise
---

## S_GlowNoise

从源素材中亮度超过给定阈值的区域生成发光效果。辉光还会被实体噪声纹理衰减，使其具有噪声或颗粒效果。如果 Jitter Frames 参数为正，每帧都会重新生成噪声，产生嘶嘶作响的外观。如果 Jitter Frames 为零，两个噪声纹理会被组合并以取决于 Spread Speed 的速率相互滑动。

在 Sapphire Lighting 效果子菜单中。

![GlowNoise](../_static/GlowNoise.jpg)


### Inputs:

- **Source**: 当前图层。用于确定辉光位置和颜色的输入素材。

- **Background**: 默认为无。用于与辉光合成的素材。如果未提供背景，则源素材也将用作背景。

- **Matte**: 默认为无。如果提供，源辉光颜色将按此输入进行缩放。单色遮罩可用于选择源素材中生成辉光的子区域。彩色遮罩可用于在不同区域选择性地调整辉光颜色。遮罩在生成辉光之前应用于源素材，因此不会裁剪生成的辉光。


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
  如果启用，Mocha 遮罩的黑白将在应用效果之前反转。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素数值膨胀或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在 High 质量模式下膨胀 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩同时提供给效果时，决定如何合并 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Brightness** (Default: 2, Range: 0 or greater)
  缩放所有辉光的亮度。

- **Color** (Default rgb: [1 1 1])
  缩放辉光的颜色。辉光的颜色和亮度也受源素材和遮罩输入的影响。

- **Threshold** (Default: 0.5, Range: 0 or greater)
  从源素材中亮度超过此值的位置生成辉光。值为 0.9 时仅在最亮的位置产生辉光。值为 0 时在每个非黑色区域都产生辉光。

- **Threshold Add Color** (Default rgb: [0 0 0])
  可用于提高特定颜色的阈值，从而减少源素材中包含该颜色的区域所产生的辉光。

- **Glow Width** (Default: 0.4, Range: 0 or greater)
  缩放辉光距离。此参数及所有宽度参数均可通过宽度控件进行调整。请注意，辉光宽度为零时仍会增强明亮区域；如果要原样传递源素材，请将亮度参数设为零。

- **Width X** (Default: 1, Range: 0 or greater)
  缩放水平辉光宽度。设为 0 则仅显示垂直方向。

- **Width Y** (Default: 1, Range: 0 or greater)
  缩放垂直辉光宽度。设为 0 则仅显示水平方向。

- **Width Red** (Default: 1, Range: 0 or greater)
  缩放红色辉光宽度。如果红、绿、蓝宽度相等，辉光将与源素材的颜色一致。如果不相等，辉光颜色将随距离变化。

- **Width Green** (Default: 1.2, Range: 0 or greater)
  缩放绿色辉光宽度。

- **Width Blue** (Default: 1.4, Range: 0 or greater)
  缩放蓝色辉光宽度。

- **Subpixel** (Check-box, Default: on)
  启用亚像素宽度辉光。用于宽度参数的更平滑动画。

- **Noise Amplitude** (Default: 1, Range: 0 or greater)
  辉光中包含的噪声振幅。

- **Noise Frequency** (Default: 40, Range: 0.1 or greater)
  噪声纹理的空间频率。增大可获得更细的颗粒，减小可获得更粗的颗粒。

- **Noise Freq Rel Y** (Default: 1, Range: 0.02 or greater)
  噪声纹理的相对垂直频率。增大可水平拉伸，减小可垂直拉伸。

- **Noise Octaves** (Integer, Default: 1, Range: 1 to 10)
  包含的噪声八度数。每个八度的频率是前一个的两倍，振幅是前一个的一半。

- **Noise Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化随机数生成器。实际种子值本身不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Noise Shift** (X & Y, Default: [0 0], Range: any)
  噪声纹理的水平和垂直平移。仅在 Jitter Frames 为零时可观察到。

- **Spread Speed** (X & Y, Default: [0.1 0], Range: any)
  两个噪声纹理相互滑动的速率和方向。除非 Jitter Frames 为零，否则此参数无效。

- **Jitter Frames** (Integer, Default: 1, Range: 0 or greater)
  如果为 0，每帧处理时噪声纹理将保持不变。如果为 1，每帧使用新的噪声纹理。如果为 2，每隔一帧使用新的噪声纹理，依此类推。

- **Show** (Popup menu, Default: Result)
  选择输出类型。
  - **Result**: 显示辉光、源素材和背景合成后的最终结果。
  - **Threshold**: 显示用于生成辉光的阈值化图像。

- **Combine** (Popup menu, Default: Screen)
  决定辉光如何与源素材或背景合成。如果 Light BG 设为 1，此参数无效。
  - **Mult**: 源素材或背景与辉光相乘。
  - **Add**: 辉光添加到源素材或背景上。
  - **Screen**: 辉光与源素材或背景使用滤色操作混合。
  - **Difference**: 结果为辉光与源素材或背景的差值。
  - **Overlay**: 辉光与源素材或背景使用叠加函数合成。

- **Edge Mode** (Popup menu, Default: Reflect)
  决定访问源图像外部区域时的行为。
  - **Transparent**: 源图像外部区域被视为透明，这可能在图像边缘产生透明效果。选择此选项可获得最快的渲染速度。
  - **Reflect**: 在边界外反射图像。

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  如果此值为正，输出的 Alpha 通道将包含来自辉光的一些不透明度。红、绿、蓝辉光亮度的最大值按此值缩放，并在每个像素处与背景 Alpha 合成。

- **Glow From Alpha** (Default: 0, Range: 0 to 1)
  设为 1 可从源输入的 Alpha 通道而非 RGB 通道生成辉光。在这种情况下，辉光不会从源素材获取颜色，通常会更亮。0 到 1 之间的值在使用 RGB 和 Alpha 之间插值。

- **Glow Under Source** (Default: 0, Range: 0 to 1)
  设为 1 可将源输入合成在辉光之上。

- **Light Background** (Default: 0, Range: 0 to 1)
  增大此值可产生辉光照亮背景图像的效果。要更清楚地看到此效果，还可以降低背景缩放参数或提高亮度参数。

- **Source Opacity** (Default: 1, Range: 0 to 1)
  缩放源输入与辉光合成时的不透明度。这不影响辉光本身的生成。

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  缩放背景的亮度。此参数仅在提供了背景输入，并且由于部分透明的源图像或降低的源不透明度参数值而可见时才有效。

- **Invert Matte** (Check-box, Default: off)
  如果开启，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则此选项无效。

- **Expand Borders** (Check-box, Default: off)
  如果启用，在处理前向输入图像添加透明边框。这允许结果包含超出原始图像大小的柔和边缘。关闭时，效果仅在画面内发生，结果将在边界处保留硬边。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Show Glow Width** (Check-box, Default: on)
  开启或关闭用于调整 Glow Width 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
