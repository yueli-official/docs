---
title: Convolve
---

## S_Convolve

使用卷积核对源图像进行卷积。卷积是一种数学运算符，使用一个图像（卷积核）作为另一个图像（源图像）的滤波器形状。卷积实际上是在源图像的每个点上印刻一份卷积核的副本，使用该点的源图像亮度。效果是卷积核的副本将出现在源图像的所有亮点上。形状像圆形或多边形的卷积核图像将产生类似于 RackDefocus 的效果；形状像星爆的卷积核图像可以产生类似 Glare 的效果。

位于 Sapphire Blur+Sharpen 效果子菜单中。

![Convolve](../_static/Convolve.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Kernel**: 默认为无。用于卷积的滤波核或形状。通常应在边缘（指定的 Kernel Crop 区域之外）全部为黑色，中心部分为非黑色。较大的形状通常产生更模糊的结果。只有两个 Kernel Crop 参数范围内的卷积核部分会被考虑；该边界之外的部分被忽略。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Kernel Size** (Default: 1, Range: 0 or greater)
  Kernel Size 将卷积核放大或缩小。1.0 为原始大小。此参数可通过 Kernel Size Widget 调整。

- **Kernel Rel Width** (Default: 1, Range: 0 or greater)
  增大可使卷积核更宽而不改变高度。减小可水平缩小，使其更窄。

- **Kernel Rel Height** (Default: 1, Range: 0 or greater)
  增大可使卷积核更高而不改变宽度。减小可垂直缩小，使其更扁。

- **Kernel Center** (X & Y, Default: [0 0], Range: any)
  卷积核的中心点；如果将卷积想象为在源图像的每个点上重复印刻卷积核，中心就是印章与其覆盖的源像素对齐的位置。如果将中心向右移动，整个结果图像将向左移动，上下方向同理。如果 AutoCenter 开启，此参数将被忽略。在调整此参数时打开 Show Kernel 可能会有帮助。注意，如果 Autocenter 关闭，无论此参数设为何值，中心点始终包含在卷积核中。

- **Autocenter** (Check-box, Default: on)
  自动查找卷积核图像的中心。开启此选项会使效果忽略 Kernel Center 参数。

- **Use Color Kernel** (Check-box, Default: off)
  独立使用卷积核的每个颜色通道。如果您的卷积核不仅仅是黑白的，并且希望在卷积中使用卷积核的颜色，请开启此选项。关闭可获得最快的渲染速度。

- **Show Kernel** (Check-box, Default: off)
  在结果上显示卷积核，以便更轻松地调整卷积核参数。最终渲染时请关闭此选项。

- **Use Gamma** (Default: 1, Range: 0.1 or greater)
  大于 1 的值使源素材中的高光在应用卷积滤波器后保持其亮度。

- **Boost Highlights** (Default: 0, Range: 0 or greater)
  增加源素材中高光亮度的量。增大此参数可使高光过曝而不影响暗部或中间调。

- **Hilight Threshold** (Default: 0.9, Range: 0 or greater)
  高光的最低亮度值。亮度高于此值的像素将根据 Boost Highlights 参数被增亮。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果的亮度。

- **Chroma Distort** (Default: 0, Range: any)
  在图像边缘周围添加一些色差；红色和蓝色光波在真实镜头中的折射方式不同，当光线以倾斜角度照射到镜头时会产生彩色边缘。

- **Threshold** (Default: 0, Range: 0 or greater)
  低于此值的任何源值都将被视为黑色。当将卷积结果与原始图像组合时，可以增大此值以仅卷积源图像的亮区。通常使用此参数时，还需将 Combine 设为 Screen 或 Add 以获得类似眩光的效果。

- **Threshold Add Color** (Default rgb: [0 0 0])
  可用于提高特定颜色的阈值，从而减少在源素材中包含该颜色的区域上生成的卷积结果。

- **Combine** (Popup menu, Default: Convolve Only)
  确定卷积图像如何与原始源图像组合。
  - **Convolve Only**: 仅显示卷积图像。用于模糊或散焦类效果。
  - **Screen**: 将卷积图像与原始源素材进行滤色叠加。用于辉光或眩光类效果。
  - **Add**: 将卷积图像添加到原始源素材。
  - **Difference**: 显示卷积图像与源素材之间的差异。

- **Mix With Source** (Default: 0, Range: 0 to 1)
  在卷积结果 (0) 和原始源素材 (1) 之间插值。0.1 可以产生不错的朦胧效果，因为它只混入少量源素材。

- **Edge Mode** (X & Y, Popup menu, Default: [ Transparent Transparent ])
  确定访问源图像边界外区域时的行为。
  - **Transparent**: 源图像外的区域被视为透明，这可能在图像边缘产生透明度。选择此项可获得最快的渲染速度。
  - **Repeat**: 在图像边界外重复最后一个像素。
  - **Reflect**: 在边界外反射图像。

- **Kernel Threshold** (Default: 0.001, Range: 0 or greater)
  低于此值的任何卷积核值都将被视为黑色。卷积核图像的边缘必须完全为黑色，否则结果会有灰色调。如果您的卷积核图像在黑色区域可能有少量噪点，请稍微增大阈值以去除该背景噪点。

- **Clamp Below Thresh** (Check-box, Default: on)
  开启时，低于阈值的值被钳制为零。这通常给出最佳结果。对于某些具有部分负值卷积核的特殊情况，关闭此选项可为您设计卷积核提供额外的灵活性。

- **Kernel Crop1** (X & Y, Default: [-0.997 -0.747], Range: any)
  卷积核区域的左上角。卷积核图像中 Kernel Crop1 和 Kernel Crop2 定义的矩形之外的部分被视为黑色。缩小此区域以避免处理卷积核的黑色边缘可以在一定程度上加速卷积。在调整此参数时打开 Show Kernel 可能会有帮助。注意，如果 Autocenter 关闭，无论此参数设为何值，中心点始终包含在卷积核中。

- **Kernel Crop2** (X & Y, Default: [0.997 0.747], Range: any)
  卷积核区域的右下角。

- **Autoscale Mode** (Popup menu, Default: Max Channel)
  在卷积中，较大或较亮的卷积核都会使结果图像更亮。必须对卷积核进行自动缩放或归一化，使结果平均与输入一样亮。自动缩放可以通过几种方式完成，每种方式在某些情况下最佳。对于单色卷积核或关闭 Color Kernel 时，Max Channel、Luma 和 Indep Channels 给出相同的结果。
  - **Max Channel**: 通过对每个通道的元素求和来自动缩放卷积核，并使用最亮的通道作为整体卷积核缩放因子。这将暗淡的卷积核归一化到全亮度，通常保留卷积核的颜色，但允许较暗通道中的亮度变化显示在结果中。
  - **Luma**: 通过对每个卷积核像素的亮度求和来自动缩放卷积核。此方法保留卷积核色相的变化，但归一化亮度，因此较亮或较暗的卷积核没有效果。使用 Scale 参数调整结果亮度。
  - **Indep Channels**: 独立归一化卷积核的每个颜色通道。有色卷积核使用此方法会产生白色/灰色结果。如果您的卷积核通道彼此独立（即 R、G 和 B 中有不同的内容），但您希望每个通道的结果都被归一化，请使用此方法。
  - **Count Nonzero**: 计算有多少卷积核像素非零（亮于黑色），但忽略它们的亮度。如果您希望卷积核色相和亮度的变化显示在结果中，此方法最佳。但模糊卷积核会产生较暗的结果，因为将有更多非零像素。
  - **Kernel Size**: 完全忽略像素；仅使用卷积核矩形的大小进行自动缩放。如果您希望所有卷积核变化都显示在结果中，请使用此选项，但如果您打算对 Kernel Crop1 和 Crop2 进行动画，请不要使用它，因为那会影响结果的亮度。

- **Soft Borders** (Check-box, Default: off)
  如果启用，在处理前向输入图像添加透明边框。这允许结果包含超出原始图像大小的柔和边缘。关闭时，效果仅在帧内发生，结果将在边框处保留边缘。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且无透明度 (alpha=1) 时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已为预乘形式处理（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时可能不太准确。

- **Show Kernel Size** (Check-box, Default: on)
  打开或关闭用于调整 Kernel Size 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show Kernel Crop** (Check-box, Default: off)
  打开或关闭用于调整 Kernel Crop1 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

