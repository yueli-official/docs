---
title: WarpMagnify
---

## S_WarpMagnify

放大源素材的椭圆形区域，以创建玻璃镜片折射效果。

在 Sapphire Distort 效果子菜单中。

![WarpMagnify](../_static/WarpMagnify.jpg)


### Inputs:

- **Source**: 当前图层。要进行变形的输入素材。

- **Matte**: 默认为无。如果提供，变形的幅度将按此输入素材的值进行缩放。灰度值在内部缩放变形幅度，而不是简单地在效果和原始源素材之间交叉淡化，以在遮罩边缘获得更连续的结果，并对变形量进行更精细的控制。此输入可通过 Blur Matte、Invert Matte 或 Matte Use 参数进行调整。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用之前按此量模糊 Mocha 遮罩。这可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 遮罩的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  如果启用，Mocha 遮罩的黑白将在应用效果之前反转。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用之前按此像素量膨胀或侵蚀 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在 High 质量模式下膨胀 Mocha 遮罩，以获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩都提供给效果时，确定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Magnify Amount** (Default: 2, Range: any)
  放大区域内图像的缩放量。使用低于 1.0 的值可在镜片内缩小图像。

- **Magnify Rel** (X & Y, Default: [1 1], Range: 0 or greater)
  相对水平和垂直放大倍率。

- **Lens Center** (X & Y, Default: [0 0], Range: any)
  镜片形状的中心。此参数可以使用 Lens Center 控件调整。

- **Lens Radius** (Default: 0.4, Range: 0 or greater)
  镜片内部的半径。在此区域内，源素材按完整放大量缩放。此参数可以使用 Edge Width 控件调整。

- **Lens Edge Width** (Default: 0.5, Range: 0 or greater)
  镜片边缘的宽度，作为内半径的分数。在镜片的边缘区域，放大倍率从完整放大量逐渐减小到无放大。此参数可以使用 Edge Width 控件调整。

- **Lens Rel Height** (Default: 1, Range: 0.1 or greater)
  镜片的相对垂直大小。增大以获得更高的椭圆，减小以获得更宽的椭圆。

- **Lens Rel Width** (Default: 1, Range: 0.1 or greater)
  镜片的相对水平大小。增大以获得更宽的椭圆，减小以获得更高的椭圆。

- **Lens Rotate** (Default: 0, Range: any)
  镜片的旋转角度（以度为单位）。请注意，当 Rel Width 和 Rel Height 相等且形状为完美圆形时，旋转无效。此参数可以使用 Lens Rotate 控件调整。

- **Lens Edge Shape** (Default: 1, Range: -1 to 1)
  确定镜片边缘内放大量的曲线。设置为零时，放大倍率线性递减。设置为一时，放大倍率以更平滑的曲线递减，这可以减少镜片边界的视觉感知。其他值在两者之间插值。

- **Wrap** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  确定访问源图像边界外区域的方法。
  - **No**: 在边界外显示黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复图像的镜像副本。使用此方法时边缘通常不太明显。

- **Filter** (Check-box, Default: on)
  如果启用，图像在重新采样时会进行自适应滤波。当图像的某些部分被变形缩小时，这会产生更好的质量结果。

- **Blur Matte** (Default: 0, Range: 0 or greater)
  在使用之前按此量模糊 Matte 输入。这可以在遮罩区域和非遮罩区域之间提供更平滑的过渡。除非提供了 Matte 输入，否则此选项无效。

- **Invert Matte** (Check-box, Default: off)
  如果启用，反转 Matte 输入，使效果应用于 Matte 为黑色而非白色的区域。除非提供了 Matte 输入，否则此选项无效。

- **Matte Use** (Popup menu, Default: Luma)
  确定如何使用 Matte 输入通道来创建单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可略微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项也比 Normal 模式渲染略快，但结果也将是预乘形式，有时不太准确。如果图像中遮罩通道也有锐利边缘的区域存在明显的颜色变化，使用 Normal 模式可能会获得更好的结果。

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  这 4 个参数，Crop Top、Crop Bottom、Crop Left 和 Crop Right，允许选择输入图像的矩形子区域进行处理。如果 Wrap 参数设置为 "No"，则暴露的边框将是透明的。如果 Wrap 为 "Tile" 或 "Reflect"，源图像将在新的裁剪边框上包裹以填充画面。这可以更容易地避免因变形具有不良边缘的图像而产生的伪影。

- **Show Edge Width** (Check-box, Default: on)
  打开或关闭用于调整 Lens Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，支持屏幕控件。

- **Show Lens Radius** (Check-box, Default: on)
  打开或关闭用于调整 Lens Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，支持屏幕控件。

- **Show Lens Rotate** (Check-box, Default: on)
  打开或关闭用于调整 Lens Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，支持屏幕控件。

- **Show Lens Center** (Check-box, Default: on)
  打开或关闭用于调整 Lens Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，支持屏幕控件。

