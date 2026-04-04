---
title: Light3D
---

## S_Light3D

使用最多 4 个可单独控制的光源执行 3D 重新打光。Source 输入通常是来自 3D 渲染器的环境光或漫射通道，显示表面颜色。法线向量输入确定每个像素处的表面方向。源和法线应由 3D 程序一起生成，以确保匹配。

在 Sapphire Lighting 效果子菜单中。

![Light3D](../_static/Light3D.jpg)


### Inputs:

- **Source**: 当前图层。3D 表面颜色。

- **Normals**: 默认为无。包含与源素材片段匹配的法线向量。通常红色通道包含法线的 X 分量，绿色通道包含 Y 分量，蓝色通道包含 Z 分量，但您可以使用第二页的 Normal Offset 和 Invert 参数调整此映射。

- **Matte**: 默认为无。用于在原始图像和结果之间进行插值。在遮罩为黑色的区域，不应用光照，可以看到原始的源图像。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览该效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存该效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此数值模糊 Mocha 遮罩。可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 遮罩的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  启用后，在应用效果之前反转 Mocha 遮罩的黑白区域。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素值膨胀或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是以默认的快速模式进行快速调整，还是以高质量模式获得更好的效果。
  - **Fast**: 以快速模式膨胀 Mocha 遮罩，便于快速调整。
  - **High**: 以高质量模式膨胀 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材片段。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当同时提供 Mocha 遮罩和输入遮罩时，确定如何组合它们。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Brightness** (Default: 1, Range: 0 or greater)
  统一缩放所有灯光的亮度。

- **Ambient Bright** (Default: 0.2, Range: any)
  整个画面中包含的环境光量。这使得没有灯光照射的源区域也能可见。

- **Diffuse Bright** (Default: 0.5, Range: 0.1 or greater)
  缩放所有光源的漫射光。

- **Hilight Bright** (Default: 0.8, Range: 0 or greater)
  缩放所有镜面高光的亮度。

- **Hilight Size** (Default: 0.5, Range: 0.1 or greater)
  调整所有镜面高光的大小。


### Light1 Parameters:

Light1 Enable:
*Check-box, Default:
*on.启用第一个光源。

Light1 Dir:
*X & Y, Default:
*[-0.806 0.608],
*Range:
*any.第一个光源的 x 和 y 位置。此参数可通过 Light1 Dir 控件进行调整。

Light1 Z:
*Default:
*0.5,
*Range:
*any.第一个光源的 z 位置。

Diffuse Bright 1:
*Default:
*0.5,
*Range:
*0.1 or greater.仅缩放 Light 1 的漫射亮度。

Hilight Bright 1:
*Default:
*1,
*Range:
*0 or greater.仅缩放 Light 1 的镜面高光亮度。

Hilight Size 1:
*Default:
*1,
*Range:
*0 or greater.仅调整 Light 1 的镜面高光大小。

Light1 Color:
*Default rgb:
*[1 1 1].
第一个光源的颜色。

### Light2 Parameters:

Light2 Enable:
*Check-box, Default:
*off.启用第二个光源。

Light2 Dir:
*X & Y, Default:
*[0.806 0.608],
*Range:
*any.第二个光源的 x 和 y 位置。此参数可通过 Light2 Dir 控件进行调整。

Light2 Z:
*Default:
*0.5,
*Range:
*any.第二个光源的 z 位置。

Diffuse Bright 2:
*Default:
*0.5,
*Range:
*0.1 or greater.仅缩放 Light 2 的漫射亮度。

Hilight Bright 2:
*Default:
*1,
*Range:
*0 or greater.仅缩放 Light 2 的镜面高光亮度。

Hilight Size 2:
*Default:
*1,
*Range:
*0 or greater.仅调整 Light 2 的镜面高光大小。

Light2 Color:
*Default rgb:
*[1 1 1].
第二个光源的颜色。

### Light3 Parameters:

Light3 Enable:
*Check-box, Default:
*off.启用第三个光源。

Light3 Dir:
*X & Y, Default:
*[-0.806 -0.627],
*Range:
*any.第三个光源的 x 和 y 位置。此参数可通过 Light2 Dir 控件进行调整。

Light3 Z:
*Default:
*0.5,
*Range:
*any.第三个光源的 z 位置。

Diffuse Bright 3:
*Default:
*0.5,
*Range:
*0.1 or greater.仅缩放 Light 3 的漫射亮度。

Hilight Bright 3:
*Default:
*1,
*Range:
*0 or greater.仅缩放 Light 3 的镜面高光亮度。

Hilight Size 3:
*Default:
*1,
*Range:
*0 or greater.仅调整 Light 3 的镜面高光大小。

Light3 Color:
*Default rgb:
*[1 1 1].
第三个光源的颜色。

### Light4 Parameters:

Light4 Enable:
*Check-box, Default:
*off.启用第四个光源。

Light4 Dir:
*X & Y, Default:
*[0.806 -0.627],
*Range:
*any.第四个光源的 x 和 y 位置。此参数可通过 Light2 Dir 控件进行调整。

Light4 Z:
*Default:
*0.5,
*Range:
*any.第四个光源的 z 位置。

Diffuse Bright 4:
*Default:
*0.5,
*Range:
*0.1 or greater.仅缩放 Light 4 的漫射亮度。

Hilight Bright 4:
*Default:
*1,
*Range:
*0 or greater.仅缩放 Light 4 的镜面高光亮度。

Hilight Size 4:
*Default:
*1,
*Range:
*0 or greater.仅调整 Light 4 的镜面高光大小。

Light4 Color:
*Default rgb:
*[1 1 1].第四个光源的颜色。

Normal Offset:
*Default:
*-0.5,
*Range:
*any.添加到法线输入中的值。

Normal X <-:
*Popup menu, Default: Red
*.确定哪个颜色通道用于法线向量的水平分量。
*Red:
*使用红色通道。*Green:
*使用绿色通道。*Blue:
*使用蓝色通道。

Normal Y <-:
*Popup menu, Default: Green
*.确定哪个颜色通道用于法线向量的垂直分量。
*Red:
*使用红色通道。*Green:
*使用绿色通道。*Blue:
*使用蓝色通道。

Normal Z <-:
*Popup menu, Default: Blue
*.确定哪个颜色通道用于法线向量的深度分量。
*Red:
*使用红色通道。*Green:
*使用绿色通道。*Blue:
*使用蓝色通道。

Invert X:
*Check-box, Default:
*off.勾选后，反转法线向量的水平分量。

Invert Y:
*Check-box, Default:
*off.勾选后，反转法线向量的垂直分量。

Invert Z:
*Check-box, Default:
*off.勾选后，反转法线向量的深度分量。

Blur Matte:
*Default:
*0,
*Range:
*0 or greater.在使用前按此数值模糊遮罩输入。这可以提供遮罩区域和非遮罩区域之间更平滑的过渡。除非提供了遮罩输入，否则无效。

Invert Matte:
*Check-box, Default:
*off.启用后，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则无效。

Matte Use:
*Popup menu, Default: Luma
*.确定如何使用遮罩输入通道来生成单色遮罩。
*Luma:
*使用 RGB 通道的亮度。*Alpha:
*仅使用 Alpha 通道。

Show Light1 Dir:
*Check-box, Default:
*on.开启或关闭用于调整 Light1 Dir 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为那里支持屏幕控件。

Show Light2 Dir:
*Check-box, Default:
*on.
开启或关闭用于调整 Light2 Dir 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为那里支持屏幕控件。
