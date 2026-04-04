---
title: TextureNoiseEmboss
---

## S_TextureNoiseEmboss

通过对程序化生成的噪波纹理应用 EmbossShiny 效果来创建抽象纹理。调整 Light Dir 可从不同角度照亮图案。

在 Sapphire Render 效果子菜单中。

![TextureNoiseEmboss](../_static/TextureNoiseEmboss.jpg)


### Inputs:

- **Background**: 当前图层。用于与纹理图像合成的素材。如果 Combine 选项设置为 Texture Only，则可能忽略此输入。

- **Mask**: 默认为无。如果提供，浮雕效果仅应用于此输入指定的区域。此输入可通过 Blur Mask、Invert Mask 或 Mask Use 参数进行调整。


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

- **Frequency** (Default: 2, Range: 0.1 or greater)
  纹理的空间频率。增大可缩小视图，减小可放大视图。

- **Frequency Rel X** (Default: 1.5, Range: 0.01 or greater)
  纹理的相对水平频率。增大可垂直拉伸，减小可水平拉伸。

- **Octaves** (Integer, Default: 5, Range: 1 to 10)
  叠加的噪波层数。每个倍频程的频率是前一个的两倍，振幅是前一个的一半。单个倍频程产生平滑的纹理。添加倍频程使结果趋近于分形（1/f）噪波纹理。

- **Seed** (Default: 0.23, Range: 0 or greater)
  用于初始化随机数生成器。实际的种子值本身不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Bumps Scale** (Default: 2.5, Range: any)
  缩放凹凸贴图的振幅。

- **Bumps Threshold** (Default: 0, Range: 0 or greater)
  在使用前从凹凸输入中减去此值。可用于创建类似"湖泊"的平坦区域。

- **Bumps Smooth** (Default: 0, Range: 0 or greater)
  在应用浮雕前平滑噪波纹理。这有助于消除噪波生成算法产生的不需要的伪影。

- **Shift** (X & Y, Default: [0 0], Range: any)
  纹理的平移偏移。由于纹理是程序化生成的，可以平移而不会出现重复单元或接缝。

- **Light Dir** (X & Y, Default: [0.389 0.33], Range: any)
  光源的方向向量。使用来自此方向的光照射生成的凹凸贴图来计算表面着色。可通过 Light Dir Widget 调整此参数。

- **Hilight Brightness** (Default: 0.5, Range: 0 to 1)
  缩放镜面高光的亮度。

- **Hilight Size** (Default: 0.5, Range: 0.1 or greater)
  调整镜面高光的大小。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果的亮度。

- **Surface Color** (Default rgb: [0.75 0.75 0.75])
  表面的颜色。最终颜色受此参数和 Light Color 的共同影响。

- **Light Color** (Default rgb: [1 1 1])
  创建浮雕结果的光源颜色。

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

- **Show Light Dir** (Check-box, Default: on)
  打开或关闭用于调整 Light Dir 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕小部件。
