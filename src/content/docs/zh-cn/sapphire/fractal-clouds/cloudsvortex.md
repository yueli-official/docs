---
title: CloudsVortex
---

## S_CloudsVortex

生成扭曲成漩涡的程序化噪波纹理。Vortex Speed 参数可使漩涡旋转量随时间自动进行动画。

在 Sapphire Render 效果子菜单中。

![CloudsVortex](../_static/CloudsVortex.jpg)


### Inputs:

- **Background**: 当前图层。用于与纹理图像合成的素材片段。如果 Combine 选项设置为 Texture Only，则可能忽略此输入。

- **Mask**: 默认为无。如果提供，扭曲的振幅将按此输入素材的值进行缩放。灰度值在内部缩放扭曲振幅，而不是简单地在效果和原始源之间交叉淡化，从而在遮罩边缘获得更连续的结果，并对扭曲量进行更精细的控制。此输入可通过 Blur Mask、Invert Mask 或 Mask Use 参数进行调整。


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
  在使用前按此像素数量膨胀或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在 High 质量模式下膨胀 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩并将效果应用于整个源素材片段。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果并显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩都提供给效果时，确定如何合并 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用
Mocha 遮罩。

- **Frequency** (Default: 2, Range: 0.01 or greater)
  纹理的空间频率。增大可缩小视图，减小可放大视图。

- **Frequency Rel X** (Default: 1, Range: 0.01 or greater)
  纹理的相对水平频率。增大可垂直拉伸，减小可水平拉伸。

- **Octaves** (Integer, Default: 6, Range: 1 to 10)
  叠加的噪波层数。每个八度是前一个八度频率的两倍、振幅的一半。单个八度会产生平滑纹理。添加八度会使结果趋近于分形 (1/f) 噪波纹理。

- **Seed** (Default: 0.234, Range: 0 or greater)
  用于初始化随机数生成器。实际种子值本身并不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Center** (X & Y, Default: [0 0], Range: any)
  漩涡的中心，以相对于画面中心的屏幕坐标表示。此参数可通过 Center Widget 进行调整。

- **Z Dist** (Default: 1, Range: 0.001 or greater)
  缩放图像的"距离"。大于 1.0 的值会使其更远更小。小于 1.0 的值会使图像更近更大。

- **Latitude** (Default: 30, Range: -80 to 80)
  正纬度值使图像向下倾斜，负值使其向上倾斜。将纬度保持在约 -35 到 35 度范围内，以避免地平线附近的锯齿。

- **Boiling Mode** (Check-box, Default: off)
  启用后，云朵会随时间沸腾或演变。此模式需要稍多的计算量，但通常效果更好。

- **Boil Details** (Default: 0.55, Range: 0 to 1)
  增加或减少云朵中的精细细节量。减小可获得更平滑的外观，增大可获得更高频、更多噪点的外观。仅在启用 Boiling Mode 时使用。

- **Boil Speed** (Default: 1, Range: any)
  设置云朵沸腾的速度。零表示完全不沸腾。仅在启用 Boiling Mode 时使用。

- **Vortex Start** (Default: 72, Range: any)
  漩涡旋转量，以画面边缘处的近似角度（度）表示。

- **Vortex Speed** (Default: 30, Range: any)
  漩涡旋转速度，以画面边缘处每秒的近似角度（度）表示。如果非零，漩涡会自动以此速率进行动画。

- **Angle Offset** (Default: 0, Range: any)
  如果非零，则将旋转与漩涡结合。设为负值可使内部和外部区域向相反方向旋转。

- **Inner Radius** (Default: 0.04, Range: 0 or greater)
  从中心开始引入漩涡效果的半径。可用于减少漩涡最中心处的过度失真和锯齿。

- **Brightness1** (Default: 1, Range: 0 or greater)
  缩放 Color1 的亮度。增大可提高对比度。

- **Color1** (Default rgb: [1 1 1])
  纹理"较亮"部分的颜色。结果的颜色由 Color0 和 Color1 之间的插值决定。

- **Color0** (Default rgb: [0 0 0])
  纹理"较暗"部分的颜色。

- **Offset0** (Default: 0, Range: any)
  将此值添加到 color0。减小为负值可增加对比度。

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  背景亮度在与纹理合成前按此值缩放。

- **Combine** (Popup menu, Default: Clouds Only)
  确定纹理如何与背景合成。
  - **Clouds Only**: 仅输出云朵纹理，不包含背景。
  - **Mult**: 纹理与背景相乘。
  - **Add**: 纹理叠加到背景上。
  - **Screen**: 纹理使用滤色操作与背景混合。
  - **Difference**: 结果为纹理与背景的差值。
  - **Overlay**: 纹理使用叠加功能与背景合成。

- **Input Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时，使用此选项可稍快地渲染。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时可能不太准确。

- **Output Opacity** (Popup menu, Default: Copy From Input)
  确定结果的不透明度/透明度。此效果不处理输入的不透明度（Alpha 通道），但可以从输入复制不透明度，或输出完全不透明的结果。
  - **All Opaque**: 使结果完全不透明，没有透明度。
  - **Copy From Input**: 从提供给此效果的当前图层复制不透明度/透明度。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用 Mask 输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊蒙版输入。这可以在蒙版区域和非蒙版区域之间提供更平滑的过渡。除非提供了蒙版输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果开启，则反转蒙版输入，使效果应用于蒙版为黑色而非白色的区域。除非提供了蒙版输入，否则无效。

- **Show Center** (Check-box, Default: on)
  开启或关闭用于调整 Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中显示，这些平台支持屏幕控件。
