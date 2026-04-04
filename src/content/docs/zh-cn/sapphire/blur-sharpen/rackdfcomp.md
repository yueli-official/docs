---
title: RackDfComp
---

## S_RackDfComp

将前景合成到背景上，同时对两个图层进行不同程度的散焦。前景的 Alpha 通道用作遮罩。如果提供了 Middle 输入，它将合成在前景和背景之间。

位于 Sapphire Blur+Sharpen 效果子菜单中。

![RackDfComp](../_static/RackDfComp.jpg)


### Inputs:

- **Foreground**: 当前图层。用作前景的素材，此素材的 Alpha 通道用作遮罩。

- **Background**: 默认为无。用作背景的素材。

- **Matte**: 默认为无。此输入的 Alpha 通道指定前景输入的不透明度。如果未提供此输入，则使用前景输入的 Alpha 通道。此输入可通过 Invert Matte 或 Matte Use 参数进行调整。

- **Middle**: 默认为无。合成在前景和背景之间的素材。

- **Mid_Matte**: 默认为无。此输入的 Alpha 通道指定 Middle 输入的不透明度。如果未提供此输入，则使用 Middle 输入的 Alpha 通道。此输入可通过 Invert Matte 或 Matte Use 参数进行调整。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Defocus Foreground** (Default: 0.088, Range: 0 or greater)
  前景及其遮罩的散焦量。此参数可通过 Fg Defocus Widget 调整。

- **Defocus Background** (Default: 0, Range: 0 or greater)
  背景的散焦量。此参数可通过 Bg Defocus Widget 调整。

- **Defocus Middle** (Default: 0, Range: 0 or greater)
  Middle 及其遮罩的散焦量。

- **Rel Height** (Default: 1, Range: 0.01 or greater)
  光圈形状的相对高度。如果不为 1，圆形会变成椭圆形，等等。

- **Shape** (Popup menu, Default: Circle)
  确定模拟相机光圈的形状。
  - **Circle**: 圆形。
  - **3 sides**: 三角形。
  - **4 sides**: 正方形。
  - **5 sides**: 五边形。
  - **6 sides**: 六边形。
  - **7 sides**: 等等。

- **Show Shape** (Check-box, Default: off)
  显示光圈形状而不是散焦图像。

- **Roundness** (Default: 0, Range: any)
  修改模拟相机光圈的形状。值为 1 产生圆形；0 产生由 Shape 参数指定边数的平边多边形。小于 0 会使边向内挤压产生星形，大于 1 会使角向内挤压产生花朵形状。如果 Shape 设置为 Circle 则无效。

- **Rotate** (Default: 0, Range: any)
  旋转光圈形状。

- **Bokeh** (Default: 0, Range: any)
  柔化光圈形状的外边缘，使散焦高光看起来更柔和。负值会使光圈形状的中心变暗，产生环形散焦形状。

- **Lens Noise** (Default: 0, Range: 0 or greater)
  增加此值可向光圈形状添加噪点，使散焦稍微变脏。可使结果更逼真。超过 1 可获得更具风格化的结果。

- **Noise Freq** (Default: 40, Range: 0.01 or greater)
  添加噪点的频率。如果 Lens Noise 为零则忽略。

- **Noise Freq Rel X** (Default: 1, Range: 0.01 or greater)
  添加光圈噪点的相对水平频率。增加此值可垂直拉伸，减少可水平拉伸。

- **Noise Seed** (Default: 0.123, Range: 0 or greater)
  添加噪点的种子值。要使噪点在每帧上看起来不同，请将此值设为每帧不同的动画。实际值无关紧要，只要每帧不同即可。

- **Use Gamma** (Default: 1, Range: 0.1 or greater)
  大于 1 的值使源素材中的高光在应用散焦后保持其亮度。

- **Matte Gamma** (Default: 1, Range: 0.1 or greater)
  用于遮罩散焦的 Gamma 值。

- **Boost Highlights** (Default: 0, Range: 0 or greater)
  增加源素材中高光亮度的量。增加此参数可使高光过曝而不影响暗部或中间调。

- **Hilight Threshold** (Default: 0.9, Range: 0 or greater)
  高光的最小亮度值。比此值更亮的像素将根据 Boost Highlights 参数进行增亮。

- **Comp Premult** (Check-box, Default: on)
  如果您提供了单独的 Matte 输入且前景像素值未按此遮罩进行预乘，请禁用此选项。

- **Matte Use** (Popup menu, Default: Alpha)
  确定如何使用前景或 Matte 输入通道生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Invert Matte** (Check-box, Default: off)
  如果启用，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则无效。

- **Edge Mode** (Popup menu, Default: Reflect)
  确定访问源图像外部区域时的行为。
  - **Transparent**: 源图像外部区域被视为透明，这可能在图像边缘产生透明度。选择此项可获得最快的渲染速度。
  - **Repeat**: 重复图像边界外的最后一个像素。
  - **Reflect**: 在边界外反射图像。

- **Soft Borders** (Check-box, Default: off)
  如果启用，在处理前向输入图像添加透明边框。这允许结果包含超出原始图像大小的柔和边缘。关闭时，效果仅在帧内发生，结果将在边界处保留边缘。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且无透明度（alpha=1）时，使用此选项可略微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式略快，但结果也将是预乘形式，有时可能不太准确。

- **Show Fg Defocus** (Check-box, Default: on)
  打开或关闭用于调整 Defocus Foreground 参数的屏幕用户界面。其值应先设为正值以便更轻松地调整。此参数仅在 AE 和 Premiere 中出现，因为它们支持屏幕控件。

- **Show Bg Defocus** (Check-box, Default: on)
  打开或关闭用于调整 Defocus Background 参数的屏幕用户界面。其值应先设为正值以便更轻松地调整。此参数仅在 AE 和 Premiere 中出现，因为它们支持屏幕控件。
