---
title: Diffuse
---

## S_Diffuse

在由 Diffuse Amount 决定的区域内打乱源输入的像素。使用 Blur Rel X 和 Y 参数可获得更水平或垂直的扩散方向。此效果的像素化外观取决于图像分辨率，因此建议在处理前测试最终分辨率。

在 Sapphire Stylize 效果子菜单中。

![Diffuse](../_static/Diffuse.jpg)


### Inputs:

- **Source**: 当前图层。要处理的片段。

- **Matte**: 默认为无。如果提供，则确定图像的哪些区域接收扩散像素。灰度值在内部缩放 Diffuse Amount 参数，而不是简单地在效果和原始源之间交叉淡化。这可以在蒙版边缘产生更连续的结果，并对扩散量进行更精细的控制。此输入可通过 Blur Matte、Invert Matte 或 Matte Use 参数来调整。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成蒙版。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此数量模糊 Mocha 蒙版。可用于柔化蒙版的边缘或量化伪影，并平滑时间位移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 蒙版的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  如果启用，在应用效果之前反转 Mocha 蒙版的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 蒙版。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 蒙版的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 蒙版的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 蒙版的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素数量膨胀或腐蚀 Mocha 蒙版。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下膨胀 Mocha 蒙版，以便快速调整。
  - **High**: 在 High 质量模式下膨胀 Mocha 蒙版，以获得更好的蒙版形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 蒙版，将效果应用于整个源片段。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 蒙版本身。

- **Combine Masks** (Popup menu, Default: Union)
  当同时提供 Mocha 蒙版和输入蒙版时，确定如何组合它们。
  - **Union**: 使用两个蒙版共同覆盖的区域。
  - **Intersect**: 使用两个蒙版之间重叠的区域。
  - **Mocha Only**: 忽略输入蒙版，仅使用 Mocha 蒙版。

- **Diffuse Amount** (Default: 0.2, Range: 0 or greater)
  像素扩散处理的振幅。此参数可通过 Diffuse Amount 小部件进行调整。

- **Rel Amount** (X & Y, Default: [1 1], Range: 0 or greater)
  缩放水平和垂直扩散的相对量。此参数可通过 Diffuse Amount 小部件进行调整。

- **Wrap** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  确定访问源图像边界外部的方法。
  - **No**: 边界外部显示为黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复镜像的副本。使用此方法时边缘通常不太明显。

- **Blur Matte** (Default: 0, Range: 0 or greater)
  在使用前按此数量模糊 Matte 输入。可以在蒙版区域和非蒙版区域之间提供更平滑的过渡。除非提供了 Matte 输入，否则无效。

- **Invert Matte** (Check-box, Default: off)
  如果开启，反转 Matte 输入，使效果应用于 Matte 为黑色而非白色的区域。除非提供了 Matte 输入，否则无效。

- **Matte Use** (Popup menu, Default: Luma)
  确定如何使用 Matte 输入通道来创建单色蒙版。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  这 4 个参数——Crop Top、Crop Bottom、Crop Left 和 Crop Right——允许选择输入图像的矩形子区域进行处理。如果 Wrap 参数设为 "No"，则暴露的边框将是透明的。如果 Wrap 为 "Tile" 或 "Reflect"，则源图像在新的裁剪边框上进行环绕以填充画面。这可以更容易地避免因扭曲边缘不佳的图像而产生的伪影。

- **Show Diffuse Amount** (Check-box, Default: on)
  打开或关闭用于调整 Diffuse Amount 参数的屏幕用户界面。此参数仅出现在支持屏幕小部件的 AE 和 Premiere 中。

