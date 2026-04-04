---
title: TextureLoops
---

## S_TextureLoops

创建重叠环形形状的抽象纹理。三组形状可以分别调整、着色，然后组合在一起。Phase Speed 参数使图案随时间自动变化。

在 Sapphire Render 效果子菜单中。

![TextureLoops](../_static/TextureLoops.jpg)


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

- **Frequency** (Default: 3, Range: 0.01 or greater)
  纹理的空间频率。增大可缩小视图，减小可放大视图。

- **Frequency Rel X** (Default: 1, Range: 0.01 or greater)
  纹理的相对水平频率。增大可垂直拉伸，减小可水平拉伸。

- **Loop Freq** (Default: 4, Range: 1 or greater)
  噪波图案内环形的频率。增大可获得更多同心环，减小可获得更少。

- **Phase Start** (Default: 0, Range: any)
  环形的相位。向内或向外移动。

- **Phase Speed** (Default: 0.1, Range: any)
  相位随时间的自动变化。

- **Seed** (Default: 1, Range: 0 or greater)
  用于初始化随机数生成器。实际的种子值本身不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Thickness** (Default: 0.1, Range: -1 to 2)
  控制环形的厚度。

- **Softness** (Default: 0.2, Range: 0.01 or greater)
  环形边缘的柔和度。增大可获得更平滑的边缘或减少锯齿。

- **Smooth** (Default: 0, Range: 0 or greater)
  在组合前模糊环形形状的程度。增大可获得散焦外观，或帮助消除锯齿伪影。

- **Shift** (X & Y, Default: [0 0], Range: any)
  纹理的平移偏移。由于纹理是程序化生成的，可以平移而不会出现重复单元或接缝。

- **Brightness1** (Default: 1, Range: 0 or greater)
  缩放 Color1 的亮度。增大以获得更多对比度。

- **Color1** (Default rgb: [1 1 1])
  纹理"较亮"部分的颜色。结果的颜色由 Color0 和 Color1 之间的插值决定。

- **Color0** (Default rgb: [0 0 0])
  纹理"较暗"部分的颜色。

- **Offset0** (Default: 0, Range: any)
  将此值添加到 color0。减小为负值以获得更多对比度。

- **Saturation** (Default: 1, Range: any)
  缩放颜色饱和度。增大可获得更鲜艳的颜色。设为 0 可获得单色效果。

- **Loops1 Freq** (Default: 1, Range: 0.01 or greater)
  第一组环形的相对频率。

- **Loops2 Freq** (Default: 1, Range: 0.01 or greater)
  第二组环形的相对频率。

- **Loops3 Freq** (Default: 1, Range: 0.01 or greater)
  第三组环形的相对频率。

- **Loops1 Thick** (Default: 0, Range: -1 to 1)
  向第一组环形的厚度添加此数值。

- **Loops2 Thick** (Default: 0, Range: -1 to 1)
  向第二组环形的厚度添加此数值。

- **Loops3 Thick** (Default: 0, Range: -1 to 1)
  向第三组环形的厚度添加此数值。

- **Loops1 Bright** (Default: 1, Range: 0 or greater)
  缩放第一组环形的亮度。设为零可移除它们。

- **Loops2 Bright** (Default: 1, Range: 0 or greater)
  缩放第二组环形的亮度。设为零可移除它们。

- **Loops3 Bright** (Default: 1, Range: 0 or greater)
  缩放第三组环形的亮度。设为零可移除它们。

- **Loops1 Color** (Default rgb: [1 1 1])
  第一组环形的颜色。

- **Loops2 Color** (Default rgb: [1 1 1])
  第二组环形的颜色。

- **Loops3 Color** (Default rgb: [1 1 1])
  第三组环形的颜色。

- **Invert** (Check-box, Default: off)
  如果启用，生成的纹理颜色将被反转。这类似于交换 Color0 和 Color1。

- **Combine Loops** (Popup menu, Default: Diff)
  用于组合三组环形颜色的操作。
  - **Add**: 将它们相加。
  - **Screen**: 使用滤色混合模式组合它们。
  - **Diff**: 使用差值运算符组合它们。
  - **Comp**: 将第二组合成到第三组之上，再将第一组合成到其之上。

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
