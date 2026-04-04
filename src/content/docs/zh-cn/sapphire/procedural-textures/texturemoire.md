---
title: TextureMoire
---

## S_TextureMoire

通过将两组同心圆环图案叠加在一起来创建抽象的摩尔纹纹理。Phase Speed 和 Moire Speed 参数使圆环随时间自动动画。

在 Sapphire Render 效果子菜单中。

![TextureMoire](../_static/TextureMoire.jpg)


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

- **A Center** (X & Y, Default: [-0.0833 -0.0926], Range: any)
  A 圆环图案的中心位置。

- **B Center** (X & Y, Default: [0.0833 0.0926], Range: any)
  B 圆环图案的中心位置。

- **Frequency** (Default: 20, Range: 0.5 or greater)
  圆环的频率。增大可获得更多更小的圆环，减小可获得更少更大的圆环。

- **Rel Freq Red** (Default: 1, Range: 0.1 or greater)
  仅缩放红色通道的圆环频率。

- **Rel Freq Green** (Default: 1, Range: 0.1 or greater)
  仅缩放绿色通道的圆环频率。

- **Rel Freq Blue** (Default: 1, Range: 0.1 or greater)
  仅缩放蓝色通道的圆环频率。

- **Double Space Rings** (Check-box, Default: off)
  如果勾选，每隔一个圆环为负值，产生双倍间距的外观。如果未勾选，则使用波形的绝对值，产生两倍数量的可见圆环。

- **Phase Start** (Default: 0, Range: any)
  圆环图案的相位。增大可从中心向外移动，减小可向中心内移动。相位参数相对于圆环的周期（1/frequency），因此将任何值精确更改 1 应产生相同的结果。

- **Phase Speed** (Default: 1, Range: any)
  每秒的相位自动变化。

- **Phase Red** (Default: 0.2, Range: any)
  仅移动红色通道的圆环相位。

- **Phase Green** (Default: 0.1, Range: any)
  仅移动绿色通道的圆环相位。

- **Phase Blue** (Default: 0, Range: any)
  仅移动蓝色通道的圆环相位。

- **Moire Phase** (Default: 0, Range: any)
  两组圆环图案的相对起始相位。将 A 圆环图案向外移动，同时将 B 圆环图案向内移动相同的量，从而改变摩尔纹图案本身。

- **Moire Speed** (Default: 1, Range: any)
  两组圆环图案相对相位的每秒自动变化。

- **A Brightness** (Default: 1, Range: 0 or greater)
  缩放 A 圆环图案的亮度。设为零可禁用并仅查看 B 圆环。

- **A Color** (Default rgb: [0.5 0.5 0.5])
  缩放 A 圆环图案的颜色。

- **A Rel Freq** (Default: 1, Range: 0.1 or greater)
  缩放 A 圆环图案的圆环频率。

- **A Rel Width** (Default: 1, Range: 0.2 or greater)
  A 圆环图案的相对水平大小。增大可获得更宽的圆环形状，减小可获得更高的形状。

- **A Rotate** (Default: 0, Range: any)
  A 圆环图案的旋转角度（度）。请注意，当 A Rel Width 为 1 时，此参数无效。

- **B Brightness** (Default: 1, Range: 0 or greater)
  缩放 B 圆环图案的亮度。设为零可禁用并仅查看 A 圆环。

- **B Color** (Default rgb: [0.5 0.5 0.5])
  缩放 B 圆环图案的颜色。

- **B Rel Freq** (Default: 1, Range: 0.1 or greater)
  缩放 B 圆环图案的圆环频率。

- **B Rel Width** (Default: 1, Range: 0.2 or greater)
  B 圆环图案的相对水平大小。增大可获得更宽的圆环形状，减小可获得更高的形状。

- **B Rotate** (Default: 0, Range: any)
  B 圆环图案的旋转角度（度）。请注意，当 A Rel Width 为 1 时，此参数无效。

- **Brightness1** (Default: 1, Range: 0 or greater)
  缩放 Color1 的亮度。增大以获得更多对比度。

- **Color1** (Default rgb: [1 1 1])
  纹理"较亮"部分的颜色。结果的颜色由 Color0 和 Color1 之间的插值决定。

- **Color0** (Default rgb: [0 0 0])
  纹理"较暗"部分的颜色。

- **Offset0** (Default: 0, Range: any)
  将此值添加到 color0。减小为负值以获得更多对比度。

- **Saturation** (Default: 1, Range: 0 to 10)
  缩放颜色饱和度。增大可获得更鲜艳的颜色。设为 0 可获得单色效果。

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
