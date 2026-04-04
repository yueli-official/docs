---
title: Etching
---

## S_Etching

使用两组粗细不同的黑白线条生成源素材的版本，呈现"蚀刻"或"石版画"效果。使用 Smooth Source 参数可去除一些细节并使线条形状更均匀。使用 Lines Frequency 参数可调整所有线条的密度。

位于 Sapphire Stylize 效果子菜单中。

![Etching](../_static/Etching.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。在结果和源输入之间进行插值。白色区域使用效果结果。黑色区域使用源素材。


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
  如果启用，在应用效果之前会反转 Mocha 遮罩的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素量扩展或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下扩展 Mocha 遮罩，用于快速调整。
  - **High**: 在 High 质量模式下扩展 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用到整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩同时提供给效果时，决定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Lines Frequency** (Default: 50, Range: 0 or greater)
  蚀刻线条的频率。增大可获得更细的线条图案，减小可获得更少的线条。

- **Lines1 Frequency** (Default: 1, Range: 0 or greater)
  缩放第一组蚀刻线条的频率。增大可获得更细的线条图案，减小可获得更少的线条。

- **Lines2 Frequency** (Default: 1, Range: 0 or greater)
  缩放第二组蚀刻线条的频率。

- **Lines Angle** (Default: 0, Range: any)
  蚀刻线条图案的旋转角度（度）。

- **Lines1 Angle** (Default: 30, Range: any)
  第一组蚀刻线条的相对角度（度）。

- **Lines2 Angle** (Default: -20, Range: any)
  第二组蚀刻线条的相对角度（度）。

- **Lines Shift** (X & Y, Default: [0 0], Range: any)
  移动线条图案。当调整线条角度参数时，此位置也将作为旋转中心。

- **Lines Sharpness** (Default: 4, Range: 0 or greater)
  蚀刻线条的锐利度。减小可获得更柔和的边缘。

- **Lines Add Width** (Default: 0, Range: any)
  增大可获得更粗的线条。

- **Smooth Source** (Default: 0, Range: 0 or greater)
  如果为正值，在应用蚀刻之前按此数值模糊源。

- **Color1** (Default rgb: [1 1 1])
  线条图案的"较亮"颜色。

- **Color0** (Default rgb: [0 0 0])
  线条图案的"较暗"颜色。

- **Wave Amp** (Default: 0.1, Range: 0 or greater)
  蚀刻线条组波浪的振幅。

- **Wave Frequency** (Default: 2, Range: 0 or greater)
  蚀刻线条波浪的频率。增大可获得更多波浪。

- **Warp Amp** (Default: 0.04, Range: any)
  使用源亮度扭曲输出的量。

- **Warp Smooth** (Default: 0.044, Range: 0 or greater)
  扭曲的平滑度。如果 Warp Amp 为 0，此参数无效。

- **Edges Scale** (Default: 0.5, Range: 0 or greater)
  调整结果中包含的源边缘量。如果为正值，将查找源图像中的边缘并添加到蚀刻图案中。

- **Edges Threshold** (Default: 0.3, Range: 0 or greater)
  确定结果中包含哪些边缘。增大可去除次要边缘和斑点。除非 Edges Scale 为正值，否则此参数无效。

- **Edges Width** (Default: 0, Range: 0 or greater)
  添加到结果中的边缘宽度。增大可获得更宽的边缘。除非 Edges Scale 为正值，否则此参数无效。

- **Edges Sharpness** (Default: 3, Range: 0 or greater)
  增大可获得更锐利的边缘，减小可获得更柔和的边缘。除非 Edges Scale 为正值，否则此参数无效。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已为预乘形式处理（颜色已按不透明度缩放）。此选项的渲染速度也略快于 Normal 模式，但结果也将为预乘形式，这有时不太准确。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用 Mask 输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊遮罩输入。可提供遮罩区域和非遮罩区域之间更平滑的过渡。除非提供了遮罩输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果开启，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则无效。
