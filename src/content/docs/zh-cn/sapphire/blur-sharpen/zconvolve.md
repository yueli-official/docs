---
title: ZConvolve
---

## S_ZConvolve

使用来自 ZBuffer 输入的深度值对卷积核进行放大或缩小，从而对源素材的局部进行卷积模糊。将输入分成若干层，并依据与焦点深度的距离及景深，对不同深度层应用不同尺寸的卷积模糊。此效果与 ZDefocus 类似，但光圈形状（或“Kernel”）来自一个素材片段。

位于 Sapphire Blur+Sharpen 效果子菜单中。

![ZConvolve](../_static/ZConvolve.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Kernel**: 默认为无。用于卷积的滤波核或形状。正常情况下，边缘（超出 Kernel Crop 指定区域）应为全黑，中心为非黑色区域。形状越大，通常模糊越强。仅考虑位于 Kernel Crop1 与 Kernel Crop2 两个参数围成区域内的核，边界之外会被忽略。

- **ZBuffer**: 默认为无。包含每个 Source 像素深度值的输入素材。取值应在黑到白之间，且最好不要抗锯齿。通常黑色表示最远处，白色表示最近处，可通过 Z Buffer 参数进行调整。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Focal Depth** (Default: 0, Range: any)
  焦平面的深度；0 为近，1 为远。具有该 Z 值的区域将处于对焦状态。接近该深度的物体是否清晰取决于 Depth Of Field。可用 Show: In Focus Zone 显示以辅助调节。若该参数效果与预期相反，可用 Z Buffer 参数反转深度值。

- **Depth Of Field** (Default: 0.1, Range: 0 to 1)
  指定 Focal Depth 附近被视为清晰的深度范围宽度。例如 Focal Depth=0.5 且 Depth Of Field=0.2 时，Z 值在 0.4–0.6 的物体都将清晰。设为 0 则只有恰好在 Focal Depth 处的物体清晰。可用 Show: In Focus Zone 显示以辅助调节。

- **Size** (Default: 1, Range: 0 or greater)
  内核缩放的最大量，既可放大也可缩小。1.0 为原始尺寸。此参数可通过 Size Widget 调整。

- **Size Rel X** (Default: 1, Range: 0 or greater)
  增大可在不改变高度的情况下让内核更“胖/宽”，减小则水平方向收缩使其更“瘦窄”。

- **Size Rel Y** (Default: 1, Range: 0 or greater)
  增大可在不改变宽度的情况下让内核更“高”，减小则垂直方向收缩使其更“扁平”。

- **Z Buffer Type** (Popup menu, Default: White is Near)
  解释 Z 缓冲中的取值方式。
  - **Black is Near**: Z 缓冲中黑色表示近处，白色表示远处。
  - **White is Near**: Z 缓冲中白色表示近处，黑色表示远处。

- **Show** (Popup menu, Default: Result)
  选择输出类型。
  - **Result**: 显示最终输出。
  - **Kernel**: 在最终输出上显示卷积核，以便调整裁剪与阈值参数。
  - **In Focus Zone**: 在原始图像上高亮显示对焦区域，以便调整焦点深度与景深。

- **Layers** (Integer, Default: 5, Range: 2 to 50)
  将源素材按深度分层的层数。层数越多处理越慢，但 Z 方向过渡更平滑。有时需要更多层以避免层间接缝可见。

- **Layer Mode** (Popup menu, Default: Interp)
  确定不同模糊层的合成方式。
  - **Comp**: 近处层合成在远处层之上。若不同深度的物体相互遮挡且深度图存在不连续，此方式通常更好，但可能更慢，且偶尔会看到层间伪影。
  - **Interp**: 依据深度图对各层进行插值。层间过渡更平滑，通常在深度图无剧烈变化时更佳。

- **Kernel Center** (X 6 Y, Default: [0 0], Range: any)
  内核的中心点；将卷积理解为在源图每一点上反复“盖章”内核，中心即内核与被卷积源像素对齐的位置。若将中心在内核中向右移动，整体结果图像会向左移动；上下同理。若 AutoCenter 为开，此参数被忽略。调整时可开启 Show Kernel 辅助。注意：若关闭 Autocenter，无论如何设置，该中心点总在内核中被包含。

- **Autocenter** (Check-box, Default: on)
  自动寻找内核图像的中心。开启后将忽略 Kernel Center 参数。

- **Use Gamma** (Default: 1, Range: 0.1 or greater)
  大于 1 时，卷积后源素材中的高光可保持其亮度。

- **Boost Highlights** (Default: 0, Range: 0 or greater)
  提升源素材高光的亮度。增大此值可在不影响暗部与中间调的情况下“拉亮”高光。

- **Highlight Threshold** (Default: 0.9, Range: 0 or greater)
  高光的最小亮度阈值。高于该阈值的像素将按 Boost Highlights 提升亮度。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果亮度。

- **Threshold** (Default: 0, Range: 0 or greater)
  低于此值的源像素将视为黑色。将卷积结果与原图合成时，可增大该值以仅卷积源中较亮区域。通常使用此参数时，也会将 Combine 设为 Screen 或 Add 以获得眩光/辉光风格。

- **Threshold Add Color** (Default rgb: [0 0 0])
  可对特定颜色提高阈值，从而减少该颜色区域生成的卷积结果。

- **Combine** (Popup menu, Default: Convolve Only)
  确定如何将卷积图与原图合成。
  - **Convolve Only**: 仅显示卷积图。用于纯模糊/失焦风格。
  - **Screen**: 将卷积图与原图相加并压暗（Screen）。用于辉光/眩光风格。
  - **Add**: 将卷积图直接与原图相加。
  - **Difference**: 显示卷积图与原图的差值。

- **Mix With Source** (Default: 0, Range: 0 to 1)
  在卷积结果（0）与原图（1）之间插值。设为约 0.1 可得到轻雾感，仅混入少量原图。

- **Edge Mode** (X 6 Y, Popup menu, Default: [ Transparent Transparent ])
  确定访问源图像边界之外的行为。
  - **Transparent**: 边界之外视为透明，可能在图像边缘产生透明度；渲染最快。
  - **Repeat**: 重复图像边界外的最后一个像素。
  - **Reflect**: 在边界外镜像翻转图像。

- **Size Rel Near** (Default: 1, Range: 0 or greater)
  缩放焦平面近侧区域的内核尺寸。

- **Size Rel Far** (Default: 1, Range: 0 or greater)
  缩放焦平面远侧区域的内核尺寸。

- **Kernel Threshold** (Default: 0.001, Range: 0 or greater)
  内核中低于此值的像素视为黑色。核图像边缘必须完全为黑，否则结果会有灰雾。若核图像的黑色区域有微弱噪声，可略微增大该阈值以去除背景噪声。

- **Clamp Below Threshold** (Check-box, Default: on)
  开启后，低于阈值的数值将被夹紧为 0，通常能得到更佳效果。对于部分为负值的特殊核图，关闭该选项可获得更多设计自由度。

- **Kernel Crop1** (X 6 Y, Default: [-0.997 -0.747], Range: any)
  内核区域的左上角。Kernel Crop1 与 Kernel Crop2 围成的矩形之外被视为黑色。适当缩小该区域以避开核图的全黑边缘，可略微提升速度。调整时可开启 Show Kernel 辅助。注意：若关闭 Autocenter，无论如何设置，中心点总在内核中被包含。

- **Kernel Crop2** (X 6 Y, Default: [0.997 0.747], Range: any)
  内核区域的右下角。

- **Autoscale Mode** (Popup menu, Default: Max Channel)
  在卷积中，核越大或越亮，结果图就越亮。必须对核进行自动缩放/归一化，使结果平均亮度与输入相当。不同方式适用于不同场景。对单色核或关闭 Color Kernel 时，Max Channel、Luma、Indep Channels 的结果相同。
  - **Max Channel**: 对各通道求和并选取最亮通道作为整体缩放因子。可将较暗的核归一化到满亮度，一般能保留核的颜色，但允许较暗通道的亮度变化体现在结果中。
  - **Luma**: 对各像素的亮度求和进行归一化。保持色相变化，但归一化亮度，使更亮/更暗的核不影响整体亮度。用 Scale 参数调节结果亮度。
  - **Indep Channels**: 分别对核的每个颜色通道做独立归一化。彩色核在此模式下会得到灰白结果。适用于 R/G/B 通道相互独立但希望各通道结果都被归一化的场景。
  - **Count Nonzero**: 仅统计核中非黑像素的数量（忽略其亮度）。适合让核的色相与亮度变化都体现在结果中。但当你对核进行模糊时，非黑像素增多，会导致结果变暗。
  - **Kernel Size**: 完全忽略像素值，只使用核矩形的面积进行归一化。若希望所有核形态变化都反映到结果中可用，但若动画 Kernel Crop1/2，将影响结果亮度，不建议配合动画使用。

- **Zbuffer Use** (Popup menu, Default: Luma)
  决定如何由 ZBuffer 输入通道生成单通道深度图。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入完全不透明（alpha=1）时可稍微加快渲染。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按已预乘形式处理（颜色已按不透明度缩放）。渲染略快于 Normal，但结果也将是预乘形式，某些情况下精确性较差。

- **Show Size** (Check-box, Default: on)
  打开或关闭用于调整 Size 的屏幕控件。此参数仅在支持屏幕控件的 AE 与 Premiere 中出现。

- **Show Kernel Crop** (Check-box, Default: off)
  打开或关闭用于调整 Kernel Crop1 的屏幕控件。此参数仅在支持屏幕控件的 AE 与 Premiere 中出现。
