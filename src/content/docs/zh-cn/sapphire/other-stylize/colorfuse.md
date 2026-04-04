---
title: ColorFuse
---

## S_ColorFuse

ColorFuse 允许将最多三个 LUT 组合在一起，以创建独特的风格化外观。提供了 Host-colorspace
和 lut-colorspace 参数，用于将素材从宿主色彩空间转换为 ColorFuse 中三个风格化 LUT 所使用的色彩空间。ColorFuse 最常使用 sRGB 作为内部 LUT 色彩空间。

位于 Sapphire Stylize 效果子菜单中。
### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。如果提供，效果仅应用于此输入的亮区所指定的源素材区域。此遮罩之外的像素不受影响，也不会对其内部的受影响像素产生贡献。此输入可通过 Invert Mask 或 Mask Use 参数进行调整。


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

- **Host Colorspace** (Popup menu, Default: sRGB)
  素材应从中转换的色彩空间。提供了一组常用 LUT，用于将素材从宿主应用程序中的色彩空间转换为 LUT 设计使用的色彩空间。如果宿主中使用的色彩空间在此预设中不可用，可以在 S_ColorFuse 之前应用 S_OCIOTransform 以获得更全面的色彩空间列表。
  - **linear**: 使用 linear 作为宿主色彩空间
  - **sRGB**: 使用 sRGB 作为宿主色彩空间
  - **rec709**: 使用 rec709 作为宿主色彩空间
  - **rec2020**: 使用 rec2020 作为宿主色彩空间
  - **rec1886**: 使用 rec1886 作为宿主色彩空间
  - **S-Log1**: 使用 S-Log1 作为宿主色彩空间
  - **S-Log2**: 使用 S-Log2 作为宿主色彩空间
  - **S-Log3**: 使用 S-Log3 作为宿主色彩空间

- **Lut Colorspace** (Popup menu, Default: sRGB)
  ColorFuse 应使用的色彩空间。提供了一组常用 LUT，用于定义效果 LUT 期望的色彩空间。
  - **linear**: 使用 linear 作为 LUT 色彩空间
  - **sRGB**: 使用 sRGB 作为 LUT 色彩空间
  - **rec709**: 使用 rec709 作为 LUT 色彩空间
  - **rec2020**: 使用 rec2020 作为 LUT 色彩空间
  - **rec1886**: 使用 rec1886 作为 LUT 色彩空间
  - **S-Log1**: 使用 S-Log1 作为 LUT 色彩空间
  - **S-Log2**: 使用 S-Log2 作为 LUT 色彩空间
  - **S-Log3**: 使用 S-Log3 作为 LUT 色彩空间

- **Choose Lut1** (Push-button)
  显示文件对话框以选择第一个 LUT。

- **Lut1 Strength** (Default: 0.5, Range: 0 to 1)
  第一个 LUT 应用于素材时的强度。

- **Choose Lut2** (Push-button)
  显示文件对话框以选择第二个 LUT。

- **Lut2 Strength** (Default: 1, Range: 0 to 1)
  第二个 LUT 应用于素材时的强度。

- **Choose Lut3** (Push-button)
  显示文件对话框以选择第三个 LUT。

- **Lut3 Strength** (Default: 1, Range: 0 to 1)
  第三个 LUT 应用于素材时的强度。

- **Mask Use** (Popup menu, Default: Luma)
  决定如何使用 Mask 输入通道生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊蒙版输入。可以在蒙版区域和非蒙版区域之间提供更平滑的过渡。除非提供了蒙版输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果开启，反转蒙版输入，使效果应用于蒙版为黑色而非白色的区域。除非提供了蒙版输入，否则无效。

