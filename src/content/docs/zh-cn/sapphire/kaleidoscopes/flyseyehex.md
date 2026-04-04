---
title: FlysEyeHex
---

## S_FlysEyeHex

将图像分割成六边形瓦片，并对每个形状内的图像进行变换，以创建蝇眼视觉效果。增大 Edge Softness 可使瓦片之间的重叠更加平滑。"Inside"参数在源图像被平铺到图案之前对其进行变换，"Tile"参数则变换整个蝇眼图案。

在 Sapphire Stylize 效果子菜单中。

![FlysEyeHex](../_static/FlysEyeHex.jpg)


### Inputs:

- **Source**: 当前图层。要处理的片段。

- **Mask**: 默认为无。在结果和源输入之间进行插值。白色区域使用效果结果，黑色区域使用源片段。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此数量模糊 Mocha 遮罩。可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

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
  选择 Dilate Mocha 是在默认的 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在 High 质量模式下膨胀 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源片段。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩都提供给效果时，确定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Tile Frequency** (Default: 12, Range: 0.1 or greater)
  瓦片图案的频率，增大可获得更多更小的瓦片。此参数可通过 Tile Freq 控件调整。

- **Tile Rel Height** (Default: 1, Range: 0.01 or greater)
  瓦片形状的相对高度，增大可获得更高的瓦片。

- **Tile Shift** (X & Y, Default: [0 0], Range: any)
  平移瓦片图案。此参数可通过 Tile Shift 控件调整。

- **Tile Rotate** (Default: 0, Range: any)
  瓦片图案的旋转角度，以度为单位。

- **Edge Softness** (Default: 0, Range: 0 to 1)
  瓦片形状之间边缘的柔和度。增大可使形状之间的混合更加平滑。

- **Inside Zdist** (Default: 2, Range: 0 or greater)
  确定每个瓦片内图像的缩放因子。大于 1 的值缩小，小于 1 的值放大。如果此值为 1、Inside Rotate 为 0 且 Overall Zdist 为 1，则结果应与输入图像相同。

- **Inside Rotate** (Default: 0, Range: any)
  每个瓦片内图像的旋转角度，以度为单位。

- **Overall Zdist** (Default: 1, Range: any)
  通过使每个瓦片朝向或远离图像中心来创建整体缩放效果。减小以放大，增大以缩小。当为 0 时，所有瓦片应包含相同的图像。

- **Wrap** (Popup menu, Default: Reflect)
  确定访问源图像边界之外区域的方法。
  - **No**: 边界之外显示黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复图像的镜像副本。使用此方法时边缘通常不太明显。

- **Filter** (Check-box, Default: on)
  如果启用，源图像将使用像素平均进行重新采样。这可以消除锯齿并获得更高质量的结果，尤其是当 Inside Zdist 较大时。如果输入图像平滑或 Inside Zdist 较小，则可能不需要。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式处理（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式略快，但结果也将是预乘形式，有时可能不太准确。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用遮罩输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数量模糊蒙版输入。这可以在蒙版和非蒙版区域之间提供更平滑的过渡。除非提供了蒙版输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果启用，反转蒙版输入，使效果应用于蒙版为黑色而非白色的区域。除非提供了蒙版输入，否则无效。

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  这 4 个参数（Crop Top、Crop Bottom、Crop Left 和 Crop Right）允许选择输入图像的矩形子区域进行处理。如果 Wrap 参数设置为"No"，暴露的边界将是透明的。如果 Wrap 为"Tile"或"Reflect"，源图像将在新的裁剪边界上进行包裹以填充画面。这可以更容易地避免因扭曲具有不良边缘的图像而产生的伪影。

- **Show Tile Freq** (Check-box, Default: off)
  打开或关闭用于调整 Tile Frequency 参数的屏幕用户界面。此参数仅在支持屏幕控件的 AE 和 Premiere 中显示。

- **Show Tile Shift** (Check-box, Default: off)
  打开或关闭用于调整 Tile Shift 参数的屏幕用户界面。此参数仅在支持屏幕控件的 AE 和 Premiere 中显示。
