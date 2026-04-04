---
title: TextureWeave
---

## S_TextureWeave

创建类似垂直交织编织纹路的抽象纹理。水平和垂直两组纹路可以使用频率、倍频程和速度参数分别独立调整。

在 Sapphire Render 效果子菜单中。

![TextureWeave](../_static/TextureWeave.jpg)


### Inputs:

- **Background**: 当前图层。用于与纹理图像合成的素材。如果 Combine 选项设置为 Texture Only，则可能忽略此输入。

- **Mask**: 默认为无。在结果与源输入之间进行插值。白色区域使用效果结果，黑色区域使用源素材。


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
  如果启用，则在应用效果前反转 Mocha 遮罩的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素值扩展或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下扩展 Mocha 遮罩，以便快速调整。
  - **High**: 在高质量模式下扩展 Mocha 遮罩，以获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当同时提供 Mocha 遮罩和输入遮罩时，确定如何组合它们。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Frequency** (Default: 20, Range: 0.01 or greater)
  纹理的空间频率。增大可缩小视图，减小可放大视图。

- **H Freq Rel X** (Default: 0.2, Range: 0.01 or greater)
  水平纹路的相对水平频率。增大可使纹路更短，减小可使纹路更长。

- **H Octaves** (Integer, Default: 2, Range: 1 to 10)
  用于水平纹路的噪波倍频程数。

- **H Speed X** (Default: 0, Range: any)
  水平纹路的水平速度。如果非零，水平纹路将自动沿其长度以此速率爬行。

- **V Frequency** (Default: 1, Range: 0.01 or greater)
  垂直纹路的相对频率。增大可使纹路更小，减小可使纹路更大。

- **V Freq Rel Y** (Default: 0.2, Range: 0.01 or greater)
  垂直纹路的相对垂直频率。增大可使纹路更短，减小可使纹路更长。

- **V Octaves** (Integer, Default: 2, Range: 1 to 10)
  用于垂直纹路的噪波倍频程数。

- **V Speed Y** (Default: 0, Range: any)
  垂直纹路的垂直速度。如果非零，垂直纹路将自动沿其长度以此速率爬行。

- **Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化随机数生成器。实际的种子值本身不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Shift** (X & Y, Default: [0 0], Range: any)
  纹理的平移偏移。由于纹理是程序化生成的，可以平移而不会出现重复单元或接缝。

- **Shift Speed** (X & Y, Default: [0 0], Range: any)
  纹理的平移速度。如果非零，结果将自动以此速率进行动画平移。动画速度值的结果可能不太直观，因此对于可变速度运动，通常最好将此值设为 0 并改为对 Shift Start 值进行动画。

- **Sharpen** (Default: 1, Range: any)
  应用的后处理锐化量。

- **Sharpen Width** (Default: 0.1, Range: 0 or greater)
  应用后处理锐化滤波器的宽度，相对于纹理大小。较高的值影响边缘周围更宽的区域，较低的值仅影响锐利边缘附近的区域。

- **Brightness1** (Default: 1, Range: 0 or greater)
  缩放 Color1 的亮度。增大以获得更多对比度。

- **Color1** (Default rgb: [1 1 1])
  纹理"较亮"部分的颜色。结果的颜色由 Color0 和 Color1 之间的插值决定。

- **Color0** (Default rgb: [0 0 0])
  纹理"较暗"部分的颜色。

- **Offset0** (Default: 0, Range: any)
  将此值添加到 color0。减小为负值以获得更多对比度。

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  背景亮度在与纹理合成前按此值缩放。

- **Combine** (Popup menu, Default: Texture Only)
  确定纹理如何与背景组合。
  - **Texture Only**: 仅输出纹理图像，不包含背景。
  - **Mult**: 纹理与背景相乘。
  - **Add**: 纹理与背景相加。
  - **Screen**: 纹理使用滤色操作与背景混合。
  - **Difference**: 结果为纹理与背景的差值。
  - **Overlay**: 纹理使用叠加功能与背景组合。

- **Input Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太正确。

- **Output Opacity** (Popup menu, Default: Copy From Input)
  确定结果的不透明度/透明度。此效果不处理输入的不透明度（Alpha 通道），但可以从输入复制不透明度，或输出完全不透明的结果。
  - **All Opaque**: 使结果完全不透明，没有透明度。
  - **Copy From Input**: 从给定此效果的当前图层复制不透明度/透明度。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用 Mask 输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊遮罩输入。这可以提供遮罩区域和非遮罩区域之间更平滑的过渡。除非提供了遮罩输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果启用，则反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则无效。
