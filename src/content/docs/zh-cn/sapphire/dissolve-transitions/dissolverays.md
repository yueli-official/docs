---
title: DissolveRays
---

## S_DissolveRays

在两个输入素材之间进行转场，同时添加动态光线效果。
素材相互溶解，光线叠加到结果上。
光线在效果持续期间逐渐增强和减弱。
光线通过沿直线移动其起点来实现动画效果。
应对 Dissolve Percent 参数进行动画处理以控制转场速度。

在 Sapphire Transitions 效果子菜单中。

![DissolveRays](../_static/DissolveRays.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始转场。

- **Background**: 默认为无。以此素材结束转场。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: Light Rays)
  在亮光线和暗光线之间进行选择。
  - **Light Rays**: 从源素材的明亮区域生成光束。
  - **Dark Rays**: 从源素材的黑暗区域生成暗光束。

- **Transition Dir** (Popup menu, Default: Dissolve Off to Bg)
  选择转场的方向。
  - **Dissolve Off to Bg**: 从当前图层转场到背景。
  - **Dissolve On from Bg**: 从背景转场到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  如果启用，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过动画 Dissolve Percent 参数手动执行转场。

- **Dissolve Percent** (Default: 0, Range: 0 to 1)
  必须禁用 Auto Trans 才能使用此参数。它决定前景和背景输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制溶解的时间。

- **Dissolve Speed** (Default: 3, Range: 1 or greater)
  From 和 To 素材之间的溶解速度。设为 1 时，溶解在效果的整个持续时间内进行。设为更高值时，溶解时间更短，但光线的渐入渐出仍占用整个持续时间。设为 10 可使转场更快捷，更像闪帧切换。

- **Rays Center** (X & Y, Default: [0 0], Range: any)
  转场中点处光线向外辐射的位置。

- **Rays Center Speed** (Default: 0.2, Range: 0 to 2)
  光线中心在屏幕上移动的速度。

- **Rays Center Angle** (Default: 0, Range: any)
  光线中心在屏幕上移动的角度。

- **Rays Length** (Default: 0.75, Range: 2 or less)
  转场中点处光线的最大长度。

- **Length Red** (Default: 1, Range: 0 or greater)
  光线红色通道的相对长度。调整此参数以及 Length Green 和 Length Blue，可创建色彩边缘效果。

- **Length Green** (Default: 1, Range: 0 or greater)
  光线绿色通道的相对长度。

- **Length Blue** (Default: 1, Range: 0 or greater)
  光线蓝色通道的相对长度。

- **Reverse Rays** (Default: 0, Range: 0 or greater)
  同时向内和向外延伸光线。反向光线的长度由 Rays Length 和此参数共同控制。

- **Rays Shrink** (Default: 0, Range: 0 to 1)
  在转场开始和结束时光线长度缩短的比例。

- **Rays Brightness** (Default: 8, Range: 0 or greater)
  转场中点处光线的最大亮度。

- **Rays Darkness** (Default: 8, Range: 0 or greater)
  缩放暗光束的强度。

- **Rays Fade** (Default: 1, Range: 0 to 1)
  在转场开始和结束时光线亮度降低的比例。

- **Blur Rays** (Default: 0, Range: 0 or greater)
  仅对光线进行模糊处理，在应用到源图像之前。

- **Blur Rays Rel** (X & Y, Default: [1 1], Range: 0 or greater)
  应用于光线的相对水平和垂直模糊宽度。将 Blur Rays Rel X 设为 0 可实现仅垂直模糊，将 Blur Rays Rel Y 设为 0 可实现仅水平模糊。

- **Rays Color** (Default rgb: [1 1 1])
  缩放光束的颜色。

- **Rays Color** (Default rgb: [0 0 0])
  缩放光束的颜色。

- **Bias Outer Bright** (Default: 0, Range: 0 to 1)
  决定沿光线的可变亮度量。通常接近 0，使光线在末端逐渐消失；0.5 使光线沿途亮度相等；1.0 使末端亮度最大。

- **Rays Res** (Popup menu, Default: Full)
  选择光线的分辨率系数。更高分辨率产生更清晰的光线，更低分辨率产生更平滑的光线且处理更快。此"Res"系数仅影响光线，背景仍以全分辨率与光线合成。
  - **Full**: 使用全分辨率。
  - **Half**: 光线以半分辨率计算。
  - **Quarter**: 光线以四分之一分辨率计算。

- **Threshold** (Default: 0.5, Range: 0 or greater)
  溶解从源素材中亮度超过此值的位置生成。值为 0.9 时，仅从最亮的点生成溶解。值为 0 时，从每个非黑色区域生成溶解。

- **Threshold Add Color** (Default rgb: [0 0 0])
  可用于提高特定颜色的阈值，从而减少从包含该颜色的源素材区域生成的溶解。

- **Shimmer Amp** (Default: 0.5, Range: 0 or greater)
  用此量的噪声纹理调制光线源图像，使光线呈现闪光效果。

- **Shimmer Freq** (Default: 40, Range: 0.1 or greater)
  闪光纹理的频率。增加可获得更细腻的闪光效果，减少可获得更大、更柔和的闪光。除非 Shimmer Amp 为正值，否则此参数无效。

- **Shimmer Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化闪光纹理的随机数生成器。实际种子值并不重要，但不同的种子会给出不同的结果，相同的值应给出可重复的结果。

- **Shimmer Shift** (X & Y, Default: [0 0], Range: any)
  闪光纹理的平移偏移。除非 Shimmer Amp 为正值，否则此参数无效。

- **Shimmer Speed** (X & Y, Default: [0 0], Range: any)
  闪光纹理的平移速度。如果非零，闪光将自动以此速度动画偏移。

- **Atmosphere Amp** (Default: 0, Range: 0 or greater)
  大气效果使光线呈现穿越尘雾并拾取光线或产生阴影的效果。此参数调整大气效果的量或幅度。零值给出平滑的光线，较高值给出更多尘雾感。

- **Atmosphere Freq** (Default: 1, Range: 0.1 to 20)
  控制大气噪声的空间频率。调高可获得更精细的细节，调低可获得更宽泛的整体变化。

- **Atmosphere Detail** (Default: 0.6, Range: 0 to 1)
  控制大气模拟中精细细节的量。减少可获得更平滑的大气，增加可获得更粗糙或颗粒状的外观。

- **Atmosphere Speed** (Default: 1, Range: any)
  大气中的云状噪声随时间演变，就像真实的尘云；此参数控制云图案随时间变化的速度。设为零可获得静态图案。

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  如果此值为正，输出的 Alpha 通道将包含来自光线的部分不透明度。红、绿、蓝光线亮度的最大值按此值缩放，并在每个像素处与背景 Alpha 合成。

- **Rays From Alpha** (Default: 0, Range: 0 to 1)
  设为 1 可从源的 alpha 通道边缘而非 RGB 通道生成光线。这通常会减少从内部边缘生成的光线。0 到 1 之间的值在使用 RGB 和 Alpha 之间进行插值。

- **Source Opacity** (Default: 1, Range: 0 to 1)
  在与光线合成时，缩放源输入的不透明度。这不影响光线本身的生成。

- **Use Source Chroma** (Default: 1, Range: 0 or greater)
  如果为 1，源输入的色度影响生成光线的色度。如果为 0，仅源输入的亮度影响光线的亮度，渲染速度也应更快。0 到 1 之间的值在这两个选项之间进行插值。

- **Show Rays Length** (Check-box, Default: on)
  开启或关闭用于调整 Rays Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。

- **Show Rays Center Angle** (Check-box, Default: on)
  开启或关闭用于调整 Rays Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
