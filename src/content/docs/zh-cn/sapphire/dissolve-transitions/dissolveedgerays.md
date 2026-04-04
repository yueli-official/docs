---
title: DissolveEdgeRays
---

## S_DissolveEdgeRays

使用动画边缘光线在两个输入素材之间转场。素材相互溶解，并在结果中添加边缘光线。边缘光线在效果持续时间内渐入和渐出。边缘光线通过沿一条线移动其原点来实现动画效果。应通过动画 Dissolve Percent 参数来控制转场速度。

在 Sapphire Transitions 效果子菜单中。

![DissolveEdgeRays](../_static/DissolveEdgeRays.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始转场。

- **Background**: 默认为无。以此素材结束转场。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Transition Dir** (Popup menu, Default: Dissolve Off to Bg)
  选择转场的方向。
  - **Dissolve Off to Bg**: 从当前图层转场到背景。
  - **Dissolve On from Bg**: 从背景转场到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  如果启用，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过动画 Dissolve Percent 参数手动执行转场。

- **Dissolve Percent** (Default: 0, Range: 0 to 1)
  必须禁用 Auto Trans 才能使用此参数。它决定前景和背景输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制溶解的时间。

- **Dissolve Speed** (Default: 3, Range: 1 or greater)
  从一个素材到另一个素材的溶解速度。设为 1 时，溶解在效果的整个持续时间内进行。设为更高值时，溶解更短，但边缘光线的渐入和渐出仍占据整个持续时间。设为 10 可使转场更快捷，更像闪帧切换。

- **Rays Center** (X & Y, Default: [0 0], Range: any)
  转场中点处光线向外发射的位置。

- **Rays Center Speed** (Default: 0.2, Range: 0 to 2)
  光线中心在屏幕上移动的速度。

- **Rays Center Angle** (Default: 0, Range: any)
  光线中心在屏幕上移动的角度。

- **Rays Length** (Default: 0.75, Range: 2 or less)
  转场中点处光线的最大长度。

- **Length Red** (Default: 1, Range: 0 or greater)
  光线红色通道的相对长度。调整此参数以及 Length Green 和 Length Blue 可创建色边效果。

- **Length Green** (Default: 1, Range: 0 or greater)
  光线绿色通道的相对长度。

- **Length Blue** (Default: 1, Range: 0 or greater)
  光线蓝色通道的相对长度。

- **Reverse Rays** (Default: 0, Range: 0 or greater)
  使光线向内和向外同时延伸。反向光线的长度由 Rays Length 和此参数共同控制。

- **Rays Shrink** (Default: 0, Range: 0 to 1)
  在转场开始和结束时光线长度缩减的比例。

- **Rays Brightness** (Default: 8, Range: 0 or greater)
  转场中点处光线的最大亮度。

- **Rays Fade** (Default: 1, Range: 0 to 1)
  在转场开始和结束时光线亮度降低的比例。

- **Blur Rays** (Default: 0, Range: 0 or greater)
  仅模糊光线，在应用于源图像之前。

- **Blur Rays Rel** (X & Y, Default: [1 1], Range: 0 or greater)
  应用于光线的相对水平和垂直模糊宽度。将 Blur Rays Rel X 设为 0 可获得仅垂直方向的模糊，将 Blur Rays Rel Y 设为 0 可获得仅水平方向的模糊。

- **Rays Color** (Default rgb: [1 1 1])
  缩放光线束的颜色。

- **Enable Dark Rays** (Check-box, Default: off)
  允许光线使源素材变暗和变亮。如果启用，暗色的 Rays Color 将使源素材变暗。亮色的 Rays Color 将照常使源素材变亮。

- **Bias Outer Bright** (Default: 0, Range: 0 to 1)
  决定沿光线的亮度变化量。通常接近 0，光线在外端逐渐消失；0.5 使光线沿途亮度均匀；1.0 使末端亮度最大。

- **Rays Res** (Popup menu, Default: Full)
  选择光线的分辨率系数。较高的分辨率产生更锐利的光线，较低的分辨率产生更平滑的光线和更快的处理速度。此"Res"系数仅影响光线：背景仍以全分辨率与光线合成。
  - **Full**: 使用全分辨率。
  - **Half**: 光线以半分辨率计算。
  - **Quarter**: 光线以四分之一分辨率计算。

- **Show** (Popup menu, Default: Result)
  在输出选项之间选择。
  - **Result**: 在背景上输出光线。
  - **Edges**: 仅输出边缘图像。这在调整边缘或闪烁参数时很有用。

- **Edge Thickness** (Default: 0.022, Range: 0 or greater)
  生成光线的边缘的厚度。

- **Edge Brightness** (Default: 1, Range: 0 or greater)
  缩放生成光线的边缘的亮度。

- **Edge Subpixel** (Check-box, Default: on)
  启用亚像素 Edge Thickness 值。如果你正在对 Edge Thickness 进行动画或需要更精细地控制小数值，请开启此选项。

- **Shimmer Amp** (Default: 0.5, Range: 0 or greater)
  用此数量的噪声纹理调制光线源图像，使光线具有闪烁外观。

- **Shimmer Freq** (Default: 40, Range: 0.01 or greater)
  闪烁纹理的频率。增大可获得更精细的闪烁效果，减小可获得更大、更柔和的闪烁。除非 Shimmer Amp 为正值，否则此参数无效。

- **Shimmer Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化闪烁纹理的随机数生成器。实际种子值本身不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Shimmer Shift** (X & Y, Default: [0 0], Range: any)
  闪烁纹理的平移。除非 Shimmer Amp 为正值，否则此参数无效。

- **Shimmer Speed** (X & Y, Default: [0 0], Range: any)
  闪烁纹理的平移速度。如果非零，闪烁将自动以此速率进行动画移动。

- **Atmosphere Amp** (Default: 0, Range: 0 or greater)
  大气效果模拟光线穿过尘土飞扬的大气层并拾取光线或被遮蔽的效果。此参数调整大气效果的数量或振幅。零值产生平滑光线，较高值产生更具尘土感的外观。

- **Atmosphere Freq** (Default: 1, Range: 0.1 to 20)
  控制大气噪声的空间频率。调高可获得更精细的细节，调低可获得更宽泛的整体变化。

- **Atmosphere Detail** (Default: 0.6, Range: 0 to 1)
  控制大气模拟中精细细节的数量。降低可获得更平滑的大气效果，增加可获得更粗糙或颗粒感的外观。

- **Atmosphere Speed** (Default: 1, Range: any)
  大气中的云状噪声会像真实的尘云一样随时间演变；此参数控制云图案随时间变化的速度。设为零可获得静态图案。

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  如果此值为正，输出的 Alpha 通道将包含来自光线的一些不透明度。红、绿、蓝光线亮度的最大值按此值缩放，并在每个像素处与背景 Alpha 合成。

- **Rays From Alpha** (Default: 0, Range: 0 to 1)
  设为 1 可从源素材的 Alpha 通道边缘而非 RGB 通道生成光线。这通常会减少从内部边缘生成的光线。0 到 1 之间的值在使用 RGB 和 Alpha 之间插值。

- **Source Opacity** (Default: 1, Range: 0 to 1)
  缩放源输入与光线合成时的不透明度。这不影响光线本身的生成。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Show Rays Center** (Check-box, Default: on)
  开启或关闭用于调整 Rays Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。

- **Show Rays Center Angle** (Check-box, Default: on)
  开启或关闭用于调整 Rays Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
