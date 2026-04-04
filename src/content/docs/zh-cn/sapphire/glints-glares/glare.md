---
title: Glare
---

## S_Glare

在源剪辑亮度超过阈值的位置合成彩虹光晕和/或类似闪光的射线。
降低阈值参数可在更多区域产生眩光。
使用 Style 菜单选择不同的眩光类型。
将 Glare Res 参数设置为 1/2 可以获得更快的渲染速度，眩光效果会略微柔和。
使用 Convolve 选项可获得更平滑的效果。
眩光效果在具有少量亮点的暗图像上观察效果最佳。

在 Sapphire Lighting 效果子菜单中。

![Glare](../_static/Glare.jpg)


### Inputs:

- **Source**: 当前图层。用于确定眩光位置和颜色的输入剪辑。

- **Background**: 默认为无。与眩光合成的剪辑。如果未提供背景，则源也用作背景。

- **Matte**: 默认为无。如果提供，源眩光颜色将按此输入进行缩放。单色遮罩可用于选择源中将生成眩光的区域子集。彩色遮罩可用于选择性地调整不同区域中的眩光颜色。遮罩在生成眩光之前应用于源，因此不会裁剪生成的眩光。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此量模糊 Mocha 遮罩。可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

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
  在使用前按此像素量膨胀或腐蚀 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在高质量模式下膨胀 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源剪辑。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩都提供给效果时，确定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Style** (Default: 0, Range: 0 or greater)
  要应用的眩光样式。也可以通过编辑 "s_glares.text" 文件来制作自定义眩光类型或修改现有类型。

- **Convolve** (Check-box, Default: off)
  确定将眩光应用于背景的方法。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放所有眩光的亮度。

- **Scale Colors** (Default rgb: [1 1 1])
  缩放眩光的颜色。眩光的颜色和亮度也受源和遮罩输入的影响。

- **Saturation** (Default: 1, Range: -2 to 8)
  缩放眩光元素的颜色饱和度。增大可获得更强烈的颜色。设置为 0 可获得单色眩光。

- **Hue Shift** (Default: 0, Range: -1 to 1)
  移动眩光的色相，以从红到绿到蓝再到红的圈数表示。

- **Threshold** (Default: 0.8, Range: 0 or greater)
  眩光从源剪辑中亮度超过此值的位置生成。值为 0.9 时仅在最亮的点产生眩光。值为 0 时每个非黑色区域都会产生眩光。

- **Threshold Add Color** (Default rgb: [0 0 0])
  可用于提高特定颜色的阈值，从而减少在源剪辑中包含该颜色区域生成的眩光。

- **Threshold Blur** (Default: 0.0896, Range: 0 or greater)
  增大此值可平滑产生眩光的区域。可用于消除由小斑点生成的眩光或简单地柔化眩光。增大此值可能会使更多高光低于阈值并使生成的眩光变暗，但可以降低阈值参数进行补偿。

- **Size** (Default: 0.8, Range: 0 or greater)
  缩放眩光的大小。此参数可使用 Size Widget 进行调整。

- **Rel Height** (Default: 1, Range: 0 or greater)
  缩放眩光的垂直维度，使其变为椭圆形而非圆形。

- **Rotate** (Default: 0, Range: any)
  旋转眩光的射线元素（如果有），以度为单位。

- **Rays Num Scale** (Default: 1, Range: 0 or greater)
  增加或减少射线的数量。

- **Rays Length** (Default: 1, Range: 0 or greater)
  调整射线的长度而不改变其粗细。

- **Rays Thickness** (Default: 1, Range: 0 or greater)
  调整各条射线的粗细。

- **Blur Glare** (Default: 0, Range: 0 or greater)
  眩光在与背景合成之前按此量进行模糊。

- **Glare Res** (Popup menu, Default: Full)
  选择眩光的分辨率因子。较高的分辨率产生更锐利的眩光，较低的分辨率产生更平滑的眩光和更快的处理速度。此"分辨率"因子仅影响眩光：背景仍以全分辨率与眩光合成。
  - **Full**: 使用全分辨率。
  - **Half**: 眩光以半分辨率计算。
  - **Quarter**: 眩光以四分之一分辨率计算。

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  如果此值为正，输出 Alpha 通道将包含来自眩光的一些不透明度。红、绿、蓝眩光亮度的最大值按此值缩放，并在每个像素处与背景 Alpha 合成。

- **Glare From Alpha** (Default: 0, Range: 0 to 1)
  设置为 1 可从源输入的 Alpha 通道而非 RGB 通道生成眩光。在这种情况下，眩光不会从源获取颜色，通常会更亮。0 到 1 之间的值在使用 RGB 和 Alpha 之间进行插值。

- **Glare Under Source** (Default: 0, Range: 0 to 1)
  设置为 1 可将源输入合成在眩光之上。

- **Source Opacity** (Default: 1, Range: 0 to 1)
  缩放与眩光合成时源输入的不透明度。这不影响眩光本身的生成。

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  缩放背景的亮度。此参数仅在提供了背景输入且由于部分透明的源图像或降低的 Source Opacity 参数值而可见时才有效果。

- **Invert Matte** (Check-box, Default: off)
  如果开启，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则此选项无效。

- **Matte Type** (Popup menu, Default: Luma)
  除非提供了遮罩输入，否则此设置将被忽略。
  - **Luma**: 使用遮罩输入的亮度来缩放眩光的亮度。
  - **Color**: 使用遮罩输入的 RGB 通道来缩放眩光的颜色。
  - **Alpha**: 使用遮罩输入的 Alpha 通道来缩放眩光的亮度。

- **Expand Borders** (Check-box, Default: off)
  如果启用，在处理前向输入图像添加透明边框。这允许结果包含超出原始图像大小的柔和边缘。关闭时，效果仅在帧内发生，结果将在边界处保留边缘。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按照图像已经是预乘形式（颜色已按不透明度缩放）进行处理。此选项的渲染速度也比 Normal 模式略快，但结果也将是预乘形式，有时不太准确。

- **Show Size** (Check-box, Default: on)
  打开或关闭用于调整 Size 和 Rel Height 参数的屏幕用户界面控件。此参数仅在支持屏幕控件的 AE 和 Premiere 中出现。
