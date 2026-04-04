---
title: VintageColor2Strip
---

## S_VintageColor2Strip

模拟 1920 年代的老式双色胶片工艺。场景通过红色和绿色滤镜两次曝光到单色胶片条的交替帧上。然后红色印刷品用红色染料染色，绿色印刷品用青色染料染色。这两条胶片背靠背粘合在一起形成最终的印刷品。结果主要包含红色和绿色，以及一些由染料的蓝色成分合成的蓝色。此效果模拟两种滤镜颜色和两种染料颜色，还允许添加颗粒和色彩校正。

在 Sapphire Stylize 效果子菜单中。

![VintageColor2Strip](../_static/VintageColor2Strip.jpg)


### Inputs:

- **Source**: 当前图层。要处理的片段。

- **Mask**: 默认为无。在结果和 Source 输入之间进行插值。白色区域使用效果的结果。黑色区域使用 Source 片段。


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

- **Amount** (Default: 1, Range: 0 or greater)
  使用的效果量。设为零可获得原始源。增大超过 1 可过度饱和。

- **Red Filter** (Default rgb: [1 0 0])
  红色滤镜的颜色。

- **Bluegreen Filter** (Default rgb: [0 1 0.5])
  绿色滤镜的颜色。

- **Red Dye** (Default rgb: [1 0 0])
  红色条的染料颜色。

- **Cyan Dye** (Default rgb: [0.02 1 0.91])
  青色条的染料颜色。略微调绿可获得更暖的外观。

- **Grain Amp** (Default: 0, Range: 0 or greater)
  缩放添加到结果中的胶片颗粒的振幅。设为 0 可禁用所有颗粒。

- **Grain Blur** (Default: 0, Range: 0 or greater)
  按此数量平滑颗粒。增大以获得更粗的颗粒。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果的亮度。

- **Saturation** (Default: 1, Range: -2 to 10)
  缩放色彩饱和度。增大以获得更鲜艳的颜色。设为 0 可获得单色效果。

- **Offset Darks** (Default: 0, Range: -8 to 2)
  将此灰度值添加到结果的较暗区域。可以为负值以增加对比度。

- **Show** (Popup menu, Default: Result)
  显示最终结果或过程中各种中间部分。
  - **Result**: 显示最终结果。
  - **Red Strip**: 显示红色滤镜后的源为单色，如同在实际胶片上的样子。
  - **BlueGreen Strip**: 显示蓝绿色滤镜后的源为单色，如同在实际胶片上的样子。
  - **Red Dye**: 显示红色染料染色的红色条。
  - **Cyan Dye**: 显示青色染料染色的绿色条。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用 Mask 输入通道来创建单色蒙版。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数量模糊 Matte 输入。可以在蒙版区域和非蒙版区域之间提供更平滑的过渡。除非提供了 Matte 输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果开启，反转 Matte 输入，使效果应用于 Matte 为黑色而非白色的区域。除非提供了 Matte 输入，否则无效。

