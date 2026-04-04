---
title: Shape
---

## S_Shape

在图像中绘制形状。可以生成多种形状，从多边形和圆形到星形、花形和螺旋海星形。
主要参数包括 Points、Pointiness、Roundness 和 Swirl。

在 Sapphire Render 效果子菜单中。

![Shape](../_static/Shape.jpg)


### Inputs:

- **Background**: 当前图层。用作背景的素材。

- **Mask**: 默认为无。在结果和源输入之间进行插值。白色区域使用效果结果，黑色区域使用源素材。


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
  如果启用，在应用效果前反转 Mocha 遮罩的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素量膨胀或侵蚀 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在高质量模式下膨胀 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩同时提供给效果时，确定如何合并 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Center** (X & Y, Default: [0 0], Range: any)
  形状的中心点。可使用 Center Widget 调整此参数。

- **Size** (Default: 0.5, Range: 0 or greater)
  形状的整体大小。可使用 Size Widget 调整此参数。

- **Rel Width** (Default: 1.4, Range: 0 or greater)
  增加以使形状更宽。

- **Rel Height** (Default: 1, Range: 0 or greater)
  增加以使形状更高。

- **Points** (Integer, Default: 5, Range: 3 to 500)
  形状中的点数。除非 Pointiness 为零，否则形状边缘将有此数量的尖角。

- **Pointiness** (Default: 0, Range: any)
  形状的尖锐程度。0 为圆形（只要 Roundness 为 1）；1 为正多边形。大于 1 为星形，小于零为具有向外凸出瓣的花形。

- **Roundness** (Default: 1, Range: 0 to 1)
  尖角之间形状边缘的圆滑程度。0 表示直线，1 表示平滑曲线。当 Pointiness 为 1 时，此参数无效。

- **Swirl** (Default: 0, Range: -5 to 5)
  设为非零值会使整个形状产生旋转效果；外边缘的旋转程度大于中心，呈现漩涡般的外观。可尝试配合较大的 Pointiness 值使用。

- **Rotate** (Default: 0, Range: any)
  围绕中心旋转整个形状。可使用 Rotate Widget 调整此参数。

- **Rotate Pre Scale** (Default: 0, Range: any)
  在应用 Rel Width 和 Rel Height 之前围绕中心旋转图形。可同时使用两种旋转以获得有趣的效果。

- **Blur** (Default: 0, Range: 0 or greater)
  模糊整个形状。

- **Brightness1** (Default: 1, Range: 0 or greater)
  缩放形状的亮度。

- **Color1** (Default rgb: [1 1 1])
  形状的颜色。

- **Color0** (Default rgb: [0 0 0])
  形状图像背景的颜色。

- **Offset0** (Default: 0, Range: any)
  将此值添加到 color0。

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  在与形状合成之前缩放背景的亮度。如果为 0，结果将仅包含黑色背景上的形状图像。

- **Combine** (Popup menu, Default: Over)
  确定形状图像与背景的合成方式。
  - **Shape Only**: 仅显示形状图像，不包含背景。
  - **Mult**: 形状图像与背景相乘。
  - **Add**: 形状图像与背景相加。
  - **Screen**: 形状图像与背景使用滤色操作混合。
  - **Difference**: 结果为形状图像与背景的差值。
  - **Overlay**: 形状图像与背景使用叠加功能合成。
  - **Over**: 将形状图像合成在背景之上。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且无透明度（alpha=1）时，使用此选项可略微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也略快于正常模式，但结果也将为预乘形式，有时可能不太准确。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用遮罩输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊蒙版输入。这可以提供遮罩区域和非遮罩区域之间更平滑的过渡。除非提供了蒙版输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果开启，反转蒙版输入，使效果应用于蒙版为黑色而非白色的区域。除非提供了蒙版输入，否则无效。

- **Show Size** (Check-box, Default: on)
  开启或关闭用于调整 Center 参数的屏幕用户界面。此参数仅在支持屏幕控件的 AE 和 Premiere 中显示。

- **Show Rotate** (Check-box, Default: on)
  开启或关闭用于调整 Center 参数的屏幕用户界面。此参数仅在支持屏幕控件的 AE 和 Premiere 中显示。

- **Show Center** (Check-box, Default: on)
  开启或关闭用于调整 Center 参数的屏幕用户界面。此参数仅在支持屏幕控件的 AE 和 Premiere 中显示。
