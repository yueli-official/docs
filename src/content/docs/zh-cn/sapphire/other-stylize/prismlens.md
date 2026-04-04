---
title: PrismLens
---

## S_PrismLens

模拟通过各种不同形状的棱镜镜头拍摄的素材效果。

在 Sapphire Stylize 效果子菜单中。

### Inputs:

- **Source**: 当前图层。决定棱镜位置和颜色的输入素材。

- **Reflection**: 默认为无。用于反射的素材。如果未连接输入，将使用源素材进行反射。

- **Mask**: 默认为无。如果提供，遮罩可以控制两件事之一。默认情况下，遮罩将控制受 PrismLens 影响的区域。如果"Use Mask For"更改为"Shape"，则遮罩将用于计算反射使用的形状。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: PrismLensSingle)
  决定棱镜位置和颜色的输入素材。
  - **PrismLensSingle**: 生成单个反射。
  - **PrismLensLinear**: 生成一排反射，类似于直棱镜产生的效果。
  - **PrismLensRadial**: 生成几个圆形反射，类似于圆形棱镜产生的效果。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前对 Mocha 遮罩进行此量的模糊处理。可用于柔化遮罩的边缘或量化伪像，并平滑时间位移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 遮罩的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  如果启用，在应用效果之前反转 Mocha 遮罩的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前对 Mocha 遮罩进行此像素量的膨胀或腐蚀处理。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认 Fast 模式下快速调整还是在 High 质量模式下获得更好效果。
  - **Fast**: 以 Fast 模式进行 Dilate Mocha，用于快速调整。
  - **High**: 以 High 质量模式进行 Dilate Mocha，获得更好看的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩并将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果并仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  决定当效果同时提供 Mocha 遮罩和输入遮罩时如何合并它们。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Use Mask For** (Popup menu, Default: Effect)
  指定遮罩应用于控制反射形状还是效果输出影响的区域。
  - **Shape**: 将遮罩用于反射形状。
  - **Effect**: 将遮罩用于反射输出。

- **Shape Source** (Popup menu, Default: Shape Only)
  应使用哪个输入来指定反射形状。
  - **Shape Only**: 使用形状参数并忽略遮罩输入。
  - **Shape and Mask**: 合并从参数生成的形状和遮罩输入。
  - **Mask Only**: 使用遮罩输入控制反射形状并忽略形状参数。
  - **None**: 不限制反射的边界。

- **Show** (Popup menu, Default: Result)
  选择输出类型。
  - **Result**: 显示最终结果。
  - **Source**: 显示用作源的输入图像。
  - **Reflection**: 显示用于反射的输入图像。
  - **Shape Mask**: 显示形状参数的结果。
  - **All Masks**: 显示所有合并的遮罩。
  - **Gradient**: 单独显示渐变，在应用到反射之前。
  - **Single Reflection**: 显示提取后的单个反射。
  - **All Reflections**: 显示未经后处理的反射。
  - **Final Reflections**: 显示经过所有后处理但在最终合成步骤之前的反射。
  - **Lights Only**: 显示灯光效果，不含源或反射。

- **Combine With Source** (Popup menu, Default: Composite)
  决定反射如何与源合成。
  - **Composite**: 将反射合成在源素材上方。
  - **Add**: 将反射颜色添加到源素材。如果暗角颜色为黑色，此操作无效。
  - **Screen**: 使用叠加操作将反射颜色与源素材合并。如果反射颜色为黑色，此操作无效。
  - **Reflections Only**: 显示不含源素材的反射图案。反射量最大的地方（例如源素材完全变暗的地方）输出将为白色。

- **Geometry** (Popup menu, Default: Polygon)
  指定反射使用的形状。
  - **Polygon**: 将反射设为多边形形状。
  - **Ellipse**: 将反射设为椭圆形。

- **Sides** (Integer, Default: 4, Range: 3 or greater)
  反射形状中的点数。除非 Pointiness 为零，否则形状将在边缘周围有此数量的点。

- **Size** (Default: 0.724, Range: 0 or greater)
  单个反射的整体大小。

- **Rel Width** (Default: 0.42, Range: 0 or greater)
  增加可使反射更宽。

- **Position** (X & Y, Default: [0.227 -0.024], Range: any)
  反射或反射图案中心的位置。

- **Position Uses Mocha** (Check-box, Default: off)
  使用 mocha 控制反射的位置。

- **Position Mocha Shift** (X & Y, Default: [0 0], Range: any)
  偏移来自 mocha 的位置数据。

- **Rotate** (Default: -33, Range: any)
  将整个反射围绕其中心旋转。

- **Softness** (Default: 0.2, Range: 0 or greater)
  对反射进行模糊处理的量。

- **Copies** (Integer, Default: 5, Range: 1 to 10)
  每行的反射数量。

- **Angle** (Default: 0, Range: any)
  反射行的角度。

- **Distance** (Default: 0.2, Range: 0.01 to 1)
  反射之间的距离。

- **Start Offset** (Default: 0.2, Range: 0 to 1)
  中心到第一个反射的距离。

- **Bias** (Default: 0, Range: -1 to 1)
  左右反射行之间间距的大小。0 表示间距相等，负值会将间距偏向中心一侧，正值会偏向另一方向。

- **Opacity Fade** (Default: 0, Range: 0 to 1)
  随着反射离中心越远，缩放其透明度。值越大，反射越透明。

- **Size Fade** (Default: 0, Range: -3 to 3)
  随着反射离中心越远，缩放其大小。正值导致大小减小，负值导致大小增加。

- **Reflect Both Directions** (Check-box, Default: on)
  是否在中心两侧都绘制一排反射。

- **Radial Copies** (Integer, Default: 6, Range: 0 to 10)
  单个环中应生成多少个反射。

- **Start Angle** (Default: 15, Range: any)
  环中第一个反射的角度。

- **Total Angle** (Default: 360, Range: 0 to 360)
  反射环的总角度覆盖范围。

- **Rows** (Integer, Default: 1, Range: 1 to 3)
  在反射中心周围生成多少个环。

- **Row Distance** (Default: 0.2, Range: 0 to 1)
  反射行之间的距离。

- **Row Start Offset** (Default: 0.2, Range: 0 to 1)
  中心到第一行反射的距离。

- **Radial Opacity Fade** (Default: 0, Range: 0 to 1)
  改变圆圈周围反射的透明度。正值使大小缩小，负值使大小增加。

- **Radial Size Fade** (Default: 0, Range: -3 to 3)
  改变圆圈周围反射的大小。正值使大小缩小，负值使大小增加。

- **Row Opacity Fade** (Default: 0, Range: 0 to 1)
  改变从中心向外辐射的反射的透明度。正值使大小缩小，负值使大小增加。

- **Row Size Fade** (Default: 0, Range: -3 to 3)
  改变从中心向外辐射的反射的大小。正值使大小缩小，负值使大小增加。

- **Reflection Opacity** (Default: 0.9, Range: 0 to 1)
  控制每个单独反射在与其他反射合并之前的透明度。

- **Source Position** (X & Y, Default: [-0.128 0.04], Range: any)
  用于反射图像的源位置。注意，这应用于指定反射源的中心，无论使用何种方法来指定反射形状。

- **Source Pos Uses Mocha** (Check-box, Default: off)
  使用 mocha 控制从反射输入中提取反射的位置。

- **Source Pos Mocha Shift** (X & Y, Default: [0 0], Range: any)
  偏移来自 mocha 的源位置数据。

- **Source Z Dist** (Default: 1.05, Range: 0 or greater)
  缩放反射的"距离"。大于 1.0 的值使其更远并更小。小于 1.0 的值使图像更近并放大。请注意，Scale X 和 Y 也会缩放图像大小，但方式相反且分别作用于每个轴。

- **Source Rotate** (Default: -2, Range: any)
  以指定的角度（度）旋转反射。

- **Source Tilt** (Default: -10, Range: any)
  在 3D 中沿水平轴向上或向下旋转反射。可以将 Swivel 和 Tilt 结合使用以围绕任意对角轴旋转。

- **Source Swivel** (Default: -12, Range: any)
  在 3D 中沿垂直轴向左或向右旋转。

- **Flip Horizontal** (Check-box, Default: off)
  水平翻转反射内的源。

- **Flip Vertical** (Check-box, Default: off)
  垂直翻转反射内的源。

- **Blur Source** (Default: 0.0084, Range: 0 or greater)
  仅对反射内的源进行模糊处理。

- **Source Threshold** (Default: 0, Range: 0 or greater)
  反射从反射输入中大于源阈值的位置生成。

- **Threshold Softness** (Default: 0.05, Range: 0 or greater)
  在阈值化后对反射进行模糊处理以获得更柔和的边缘。

- **Source Wrap** (Popup menu, Default: Reflect)
  决定访问源图像边界之外区域的方法。
  - **No**: 边界外呈现黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复镜像副本。此方法通常边缘不太明显。

- **CC Mix With Source** (Default: 0, Range: 0 to 1)
  在颜色校正后的反射和原始反射之间进行插值。

- **CC Brightness** (Default: 1, Range: 0 or greater)
  缩放反射的颜色饱和度。增加可获得更鲜艳的颜色。设为 0 可获得单色效果。也可以通过将此值设为负值来反转结果的色度。

- **Saturation** (Default: 1, Range: -2 to 8)
  缩放反射的颜色饱和度。增加可获得更鲜艳的颜色。设为 0 可获得单色效果。也可以通过将此值设为负值来反转结果的色度。

- **Hue Shift** (Default: 0, Range: any)
  以从红到绿到蓝再到红的旋转圈数来移动反射的色相。

- **Offset Darks** (Default: 0, Range: -8 to 2)
  向结果的较暗区域添加此灰度值。此值可以为负以增加对比度。

- **Amount** (Default: 0.05, Range: 0 or greater)
  通过缩放变换来调整整体扭曲量。将此值设为零可禁用两个变换并使图像保持不变。

- **Chromatic Brightness** (Default: 1, Range: 0 or greater)
  缩放渐变图像（包括 Start Color 和 End Color）的亮度。

- **Chromatic Steps** (Integer, Default: 12, Range: 3 to 100)
  反射上色差的光谱采样数量。更多步骤给出更平滑的结果，但需要更多处理时间。

- **Chromatic Shift** (X & Y, Default: [0 0], Range: any)
  控制色差的方向。

- **Chromatic Rotate** (Default: 0, Range: any)
  色差的旋转角度。

- **Chromatic Z Dist** (Default: 1, Range: 0.001 or greater)
  色差的距离或比例。

- **Chromatic Color1** (Default rgb: [1 0 0])
  色差中起始点的颜色。

- **Chromatic Color2** (Default rgb: [0 1 0])
  色差中起点和终点之间中间点的颜色。

- **Chromatic Color3** (Default rgb: [0 0 1])
  色差中终点的颜色。

- **Glow Brightness** (Default: 1.56, Range: 0 or greater)
  缩放渐变图像（包括 Start Color 和 End Color）的亮度。

- **Glow Threshold** (Default: 0.6, Range: 0 or greater)
  棱镜从源素材中亮度超过此值的位置生成。值为 0.9 时，仅从最亮的点生成棱镜。值为 0 时，从每个非黑色区域生成棱镜。

- **Glow Width** (Default: 0.2, Range: 0 or greater)
  缩放辉光距离。此参数及所有宽度参数都可以使用 Width 控件调整。请注意，零辉光宽度仍会增强明亮区域；如果想让源素材直接通过而不做更改，请将亮度参数设为零。

- **Width Red** (Default: 1, Range: 0 or greater)
  缩放红色辉光宽度。如果红、绿、蓝宽度相等，辉光将与源素材的颜色匹配。如果不相等，辉光颜色将随距离变化。

- **Width Green** (Default: 1, Range: 0 or greater)
  缩放绿色辉光宽度。

- **Width Blue** (Default: 1, Range: 0 or greater)
  缩放蓝色辉光宽度。

- **Gradient Opacity** (Default: 0.4, Range: 0 to 1)
  控制渐变与反射合并之前的透明度。

- **Color 1** (Default rgb: [1 0 0])
  渐变第一个控制点的颜色。

- **Point 1** (X & Y, Default: [0.599 -0.526], Range: any)
  渐变第一个控制点的位置。

- **Color 2** (Default rgb: [0 1 0])
  渐变第二个控制点的颜色。

- **Point 2** (X & Y, Default: [-0.98 0.529], Range: any)
  渐变第二个控制点的位置。

- **Color 3** (Default rgb: [0 0 1])
  渐变第三个控制点的颜色。

- **Point 3** (X & Y, Default: [0.958 0.529], Range: any)
  渐变第三个控制点的位置。

- **Combine** (Popup menu, Default: Screen)
  决定渐变如何与背景合并。
  - **Grad Only**: 仅给出渐变图像，不含背景。
  - **Mult**: 背景乘以渐变。
  - **Add**: 背景加到渐变上。
  - **Screen**: 使用叠加操作将背景与渐变混合。
  - **Difference**: 结果是背景和渐变之间的差值。
  - **Overlay**: 使用叠加函数合并渐变和背景。

- **Lights** (Default: 0, Range: 0 or greater)
  应用于生成灯光图案的镜头光晕类型。也可以通过用光晕设计器编辑光晕来创建自定义镜头光晕类型或修改现有类型。

- **Lights Brightness** (Default: 1, Range: 0 or greater)
  缩放灯光的亮度。

- **Lights Size** (Default: 2.25, Range: 0 or greater)
  缩放灯光的大小。

- **Lights Start Position** (X & Y, Default: [-1 0], Range: any)
  灯光图案热点的起始位置。

- **Lights Pos Uses Mocha** (Check-box, Default: off)
  使用 mocha 轨道控制灯光的位置。

- **Shift Speed** (X & Y, Default: [0.167 -0.178], Range: any)
  热点在素材上移动的速度。

- **Lights Pivot** (X & Y, Default: [0.565 0.328], Range: any)
  灯光图案中心的位置。

- **Pivot Uses Mocha** (Check-box, Default: off)
  使用 mocha 轨道控制灯光的轴心点。

- **Mask Use** (Popup menu, Default: Luma)
  决定如何使用遮罩输入通道来制作单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前对遮罩输入进行此量的模糊处理。这可以在遮罩和非遮罩区域之间提供更平滑的过渡。除非提供了遮罩输入，否则此参数无效。

- **Invert Mask** (Check-box, Default: off)
  如果启用，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则此参数无效。

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  如果此值为正，输出的 Alpha 通道将包含来自棱镜的部分不透明度。红、绿、蓝棱镜亮度的最大值按此值缩放，并在每个像素处与背景 Alpha 合成。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Show Position** (Check-box, Default: on)
  开启或关闭用于调整 Position 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。

- **Show Source Position** (Check-box, Default: on)
  开启或关闭用于调整 Source Position 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。

- **Show Point 1** (Check-box, Default: on)
  开启或关闭用于调整 Point 1 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。

- **Show Point 2** (Check-box, Default: on)
  开启或关闭用于调整 Point 2 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。

- **Show Point 3** (Check-box, Default: on)
  开启或关闭用于调整 Point 3 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。

- **Show Lights Start Position** (Check-box, Default: on)
  开启或关闭用于调整 Lights Start Position 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。

- **Show Lights Pivot** (Check-box, Default: on)
  开启或关闭用于调整 Lights Pivot 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
