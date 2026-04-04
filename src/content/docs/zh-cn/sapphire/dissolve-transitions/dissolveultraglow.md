---
title: DissolveUltraGlow
---

## S_DissolveUltraGlow

在两个输入素材之间进行转场，同时生成各种发光闪光效果。
素材相互溶解，同时每个素材都获得一个在效果持续期间逐渐增强和减弱的辉光。
调整 After Glow 参数可在主辉光结果上生成次级辉光。
还可以选择增强边缘或将结果与大气噪声合并。
应对 Dissolve Percent 参数进行动画处理以控制转场速度。

在 Sapphire Transitions 效果子菜单中。

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
  From 和 To 素材之间的溶解速度。设为 1 时，溶解在效果的整个持续时间内进行。设为更高值时，溶解时间更短，但辉光的渐入渐出仍占用整个持续时间。设为 10 可使转场更快捷，更像闪帧切换。

- **Brightness** (Default: 1.7, Range: 0 or greater)
  缩放所有辉光的亮度。

- **Color** (Default rgb: [1 1 1])
  缩放主辉光的颜色。

- **Threshold** (Default: 0.2, Range: 0 or greater)
  辉光从源素材中亮度超过此值的位置生成。值为 0.9 时，仅从最亮的点生成辉光。值为 0 时，从每个非黑色区域生成辉光。

- **Threshold Add Color** (Default rgb: [0 0 0])
  可用于提高特定颜色的阈值，从而减少从包含该颜色的源素材区域生成的辉光。

- **Glow Width** (Default: 0.371, Range: 0 or greater)
  发光边框的宽度。

- **Glow Falloff** (Default: 0.35, Range: -2 to 2)
  增强或削减辉光延伸的距离。

- **Glow Bias** (Default: 0, Range: -3 to 3)
  扩大阈值化结果的外围区域的量，若为负值则收缩。

- **Glow From Alpha** (Default: 0, Range: 0 to 1)
  设为 1 可从源输入的 alpha 通道而非 RGB 通道生成辉光。这种情况下辉光不会从源获取颜色，通常会更亮。0 到 1 之间的值在使用 RGB 和 Alpha 之间进行插值。

- **Width X** (Default: 1, Range: 0 or greater)
  缩放水平辉光宽度。设为 0 仅保留垂直辉光。

- **Width Y** (Default: 1, Range: 0 or greater)
  缩放垂直辉光宽度。设为 0 仅保留水平辉光。

- **Width Red** (Default: 1, Range: 0 or greater)
  缩放红色辉光宽度。如果红、绿、蓝宽度都相等，辉光将与 Glow Color 匹配。否则将有颜色变化的边缘。

- **Width Green** (Default: 1, Range: 0 or greater)
  缩放绿色辉光宽度。

- **Width Blue** (Default: 1, Range: 0 or greater)
  缩放蓝色辉光宽度。

- **Show** (Popup menu, Default: Result)
  选择输出类型。
  - **Result**: 显示合并辉光、源和背景的最终结果。
  - **Threshold**: 显示用于生成辉光的阈值化图像。

- **After Glow Width** (Default: 1.8, Range: 0 or greater)
  缩放次级辉光的辉光距离。

- **After Glow Color** (Default rgb: [1 1 1])
  缩放次级辉光的颜色。

- **After Glow Stretch X** (Default: 0.3, Range: 0 or greater)
  缩放水平次级辉光宽度。

- **After Glow Stretch Y** (Default: 0.1, Range: 0 or greater)
  缩放垂直次级辉光宽度。

- **Horizontal Streaks** (Default: 0, Range: 0 or greater)
  缩放水平方向细条纹的外观。

- **Vertical Streaks** (Default: 0, Range: 0 or greater)
  缩放垂直方向细条纹的外观。

- **Edge Detect** (Check-box, Default: off)
  启用边缘检测。

- **Edge Combine** (Popup menu, Default: Edges Only)
  决定检测到的边缘如何与源合成。
  - **Screen**: 使用叠加操作将检测到的边缘与源混合。
  - **Add**: 将检测到的边缘添加到源上。
  - **Edges Only**: 仅给出检测到的边缘，不包含源。

- **Edge Smooth** (Default: 0, Range: 0 or greater)
  增加可获得更粗更平滑的边缘。

- **Edge Mode** (Popup menu, Default: Transparent)
  决定访问源图像外部区域时的行为。
  - **Transparent**: 源图像外部的区域被视为透明，可能在图像边缘产生透明区域。选择此项可获得最快渲染速度。
  - **Reflect**: 在边界外反射图像。

- **Edge Fill** (Check-box, Default: on)
  使检测到的边缘内部的区域不透明。

- **Edge Thin** (Default: 0.1, Range: 0 or greater)
  从检测到的边缘结果中减去此值。增加可去除次要边缘中不需要的噪声。

- **Atmosphere** (Check-box, Default: off)
  大气效果使辉光呈现穿越尘雾并拾取光线或产生阴影的效果。此参数调整大气效果的量或幅度。零值给出平滑的辉光，较高值给出更多尘雾感。

- **Atmosphere Amp** (Default: 1, Range: 0 or greater)
  大气效果使辉光呈现穿越尘雾并拾取光线或产生阴影的效果。此参数调整大气效果的量或幅度。零值给出平滑的辉光，较高值给出更多尘雾感。

- **Atmosphere Freq** (Default: 11.6, Range: 0.1 to 20)
  控制大气噪声的空间频率。调高可获得更精细的细节，调低可获得更宽泛的整体变化。

- **Atmosphere Detail** (Default: 0.506, Range: 0 to 1)
  控制大气模拟中精细细节的量。减少可获得更平滑的大气，增加可获得更粗糙或颗粒状的外观。

- **Atmosphere Speed** (Default: 1, Range: any)
  大气中的云状噪声随时间演变，就像真实的尘云；此参数控制云图案随时间变化的速度。设为零可获得静态图案。

- **Atmosphere Lights** (Default: 0.5, Range: 0 or greater)
  按此值缩放大气层。增加可获得更强烈的结果。

- **Atmosphere Darks** (Default: 0, Range: 0 or greater)
  向大气层的较暗区域添加此灰度值。此值可以为负以增加对比度。

- **Atmosphere Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化大气噪声的随机数生成器。实际种子值并不重要，但不同的种子会给出不同的结果，相同的值应给出可重复的结果。

- **Apply Pre-Glow** (Check-box, Default: off)
  启用在任何辉光之前将大气与源合并。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Show Glow Width** (Check-box, Default: on)
  开启或关闭用于调整 Glow Width 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
