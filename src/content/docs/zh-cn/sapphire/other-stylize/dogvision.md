---
title: DogVision
---

## S_DogVision

生成输入图像的双色通道版本，模拟狗的有限色觉系统所感知到的图像。人类有三种颜色感受器（红色、绿色和蓝色），而狗只有两种感受器（黄色和蓝色）。

位于 Sapphire Stylize 效果子菜单中。

![DogVision](../_static/DogVision.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。在结果和源素材输入之间进行插值。白色区域使用效果结果。黑色区域使用源素材。


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

- **Channels** (Popup menu, Default: Yellow-Blue)
  选择使用哪两个互补色通道。
  - **Yellow-Blue**: 使用黄色和蓝色生成结果。
  - **Cyan-Red**: 使用青色和红色生成结果。
  - **Magenta-Green**: 使用品红色和紫色生成结果。

- **Rotate Channels** (Default: 0, Range: any)
  允许对上面选择的两个色通道进行色相偏移。注意，当此值非零时，通道可能不再与所选名称匹配。

- **Blur Channel1** (Default: 0, Range: 0 or greater)
  按此数值平滑第一个颜色通道。

- **Blur Channel2** (Default: 0, Range: 0 or greater)
  按此数值平滑第二个颜色通道。

- **Mix Original** (Default: 0, Range: any)
  在双色结果和原始源素材之间进行插值。设为 1 显示原始图像，或使用负值来增强狗视觉效果。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果的亮度。

- **Offset Darks** (Default: 0, Range: -8 to 2)
  将此灰度值添加到结果的较暗区域。可以为负值以增加对比度。

- **Saturation** (Default: 1, Range: -2 to 8)
  缩放颜色饱和度。增大以获得更浓烈的颜色。设为 0 则为单色。

- **Weight Source R** (Default: 1, Range: any)
  在处理前缩放输入素材的红色分量。

- **Weight Source G** (Default: 1, Range: any)
  在处理前缩放输入素材的绿色分量。

- **Weight Source B** (Default: 1, Range: any)
  在处理前缩放输入素材的蓝色分量。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可略微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 将图像视为已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式略快，但结果也将是预乘形式，有时不太正确。

- **Mask Use** (Popup menu, Default: Luma)
  决定如何使用 Mask 输入通道生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊蒙版输入。可以在蒙版区域和非蒙版区域之间提供更平滑的过渡。除非提供了蒙版输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果开启，反转蒙版输入，使效果应用于蒙版为黑色而非白色的区域。除非提供了蒙版输入，否则无效。

