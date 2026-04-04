---
title: GlowOrthicon
---

## S_GlowOrthicon

在源素材中亮度超过给定阈值的区域周围使源素材变暗，以产生"正析像管"或"暗辉光"的外观。降低阈值参数可以在更多区域产生正析像管效果。调整 Darkness 和 Width 参数可以获得不同类型的外观。

在 Sapphire Lighting 效果子菜单中。

![GlowOrthicon](../_static/GlowOrthicon.jpg)


### Inputs:

- **Source**: 当前图层。用于确定需要变暗位置的输入素材。

- **Background**: 默认为无。用于与辉光合成的素材。如果未提供背景，则源素材也将用作背景。

- **Matte**: 默认为无。如果提供，源辉光颜色将按此输入进行缩放。单色遮罩可用于选择源素材中生成辉光的子区域。彩色遮罩可用于在不同区域选择性地调整辉光颜色。遮罩在生成辉光之前应用于源素材，因此不会裁剪生成的辉光。


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
  在使用前按此像素数值膨胀或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在 High 质量模式下膨胀 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩同时提供给效果时，决定如何合并 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Darkness** (Default: 1, Range: 0 or greater)
  缩放变暗的程度。

- **Color** (Default rgb: [0 0 0])
  缩放辉光的颜色。辉光的颜色和亮度也受源素材和遮罩输入的影响。

- **Threshold** (Default: 0.7, Range: 0 or greater)
  在源素材中亮度超过此值的位置周围会产生变暗效果。值为 0.9 时仅从最亮的位置产生暗辉光。值为 0 时在每个非黑色区域都产生辉光。

- **Threshold Add Color** (Default rgb: [0 0 0])
  可用于提高特定颜色的阈值，从而减少源素材中包含该颜色的区域所产生的辉光。

- **Darks Width** (Default: 0.224, Range: 0 or greater)
  缩放暗辉光的距离。此参数及所有宽度参数均可通过宽度控件进行调整。

- **Protect Width** (Default: 0.1, Range: 0 or greater)
  明亮区域周围受保护免于变暗的距离。此值通常应小于 Darks Width 的值。

- **Protect Amount** (Default: 1, Range: 0 or greater)
  明亮区域受保护免于变暗的程度。

- **Width X** (Default: 1, Range: 0 or greater)
  缩放水平辉光宽度。设为 0 则仅显示垂直方向。

- **Width Y** (Default: 1, Range: 0 or greater)
  缩放垂直辉光宽度。设为 0 则仅显示水平方向。

- **Subpixel** (Check-box, Default: on)
  启用亚像素宽度辉光。用于宽度参数的更平滑动画。

- **Glow From Alpha** (Default: 0, Range: 0 to 1)
  设为 1 可从源输入的 Alpha 通道而非 RGB 通道生成辉光。在这种情况下，辉光不会从源素材获取颜色，通常会更亮。0 到 1 之间的值在使用 RGB 和 Alpha 之间插值。

- **Glow Under Source** (Default: 0, Range: 0 to 1)
  设为 1 可将源输入合成在辉光之上。

- **Source Opacity** (Default: 1, Range: 0 to 1)
  缩放源输入与辉光合成时的不透明度。这不影响辉光本身的生成。

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  缩放背景的亮度。此参数仅在提供了背景输入，并且由于部分透明的源图像或降低的源不透明度参数值而可见时才有效。

- **Show** (Popup menu, Default: Result)
  选择输出类型。
  - **Result**: 显示辉光、源素材和背景合成后的最终结果。
  - **Threshold**: 显示用于生成辉光的阈值化图像。

- **Invert Matte** (Check-box, Default: off)
  如果开启，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则此选项无效。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Show Darks Width** (Check-box, Default: on)
  开启或关闭用于调整 Darks Width 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
