---
title: Mosaic
---

## S_Mosaic

生成源素材的像素化版本。使用 Pixel Frequency 和 Pixel Rel Height 参数调整色块的大小和形状。增大 Smooth Colors 参数可使相邻像素色块的颜色更一致，并减少随时间的闪烁。

位于 Sapphire Stylize 效果子菜单中。

![Mosaic](../_static/Mosaic.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Matte**: 默认为无。如果提供，效果仅应用于此输入指定的源素材区域。对于蒙版中的灰度值，像素色块与原始源素材混合，使色块淡出但保持完整。此输入可通过 Blur Matte、Invert Matte 或 Matte Use 参数进行调整。


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
  在使用前按此像素数扩展或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下扩展 Mocha 遮罩，用于快速调整。
  - **High**: 在高质量模式下扩展 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩同时提供给效果时，决定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Pixel Frequency** (Default: 40, Range: 1 or greater)
  像素色块的频率。增大以获得更多、更小的像素。

- **Pixel Rel Height** (Default: 1, Range: 0.01 or greater)
  像素色块的相对高度。增大以获得更高的色块，减小以获得更宽的色块。

- **Pixel Shift** (X & Y, Default: [0 0], Range: any)
  像素图案的平移。

- **Smooth Colors** (Default: 0, Range: 0 or greater)
  在像素化之前模糊源素材。增大可使相邻像素色块的颜色更一致，并减少随时间的闪烁。

- **Mix With Source** (Default: 0, Range: 0 to 1)
  在结果（0）和原始源素材（1）之间进行插值。

- **Blur Matte** (Default: 0, Range: 0 or greater)
  在使用前按此数值模糊蒙版输入。可以在蒙版区域和非蒙版区域之间提供更平滑的过渡。除非提供了蒙版输入，否则无效。

- **Invert Matte** (Check-box, Default: off)
  如果开启，反转蒙版输入，使效果应用于蒙版为黑色而非白色的区域。除非提供了蒙版输入，否则无效。

- **Matte Use** (Popup menu, Default: Luma)
  决定如何使用 Matte 输入通道生成单色蒙版。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可略微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 将图像视为已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式略快，但结果也将是预乘形式，有时不太正确。

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  这 4 个参数 Crop Top、Crop Bottom、Crop Left 和 Crop Right 允许选择输入图像的矩形子区域进行处理。如果 Wrap 参数设为 "No"，露出的边框将是透明的。如果 Wrap 为 "Tile" 或 "Reflect"，源图像将在新的裁剪边框上环绕以填充画面。这可以更容易地避免因扭曲边缘不良的图像而产生的伪影。

