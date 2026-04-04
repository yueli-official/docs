---
title: WarpDrops
---

## S_WarpDrops

通过从多个中心位置发出的同心波图案对源素材进行变形。Centers 输入素材中每个亮度高于 Threshold Cntrs 值的区域会生成独立的同心波图案，每个区域的总亮度会缩放这些波的变形幅度。如果 Centers 图像较复杂，生成的中心数量和位置可能对阈值相当敏感。尝试使用纯黑色加几个白点作为 Centers 输入。如果只需要单组波纹，可以改用 WarpPuddle 效果。

在 Sapphire Distort 效果子菜单中。

![WarpDrops](../_static/WarpDrops.jpg)


### Inputs:

- **Source**: 当前图层。要进行变形的输入素材。

- **Centers**: 默认为无。确定波纹图案的中心。此素材中每个亮度高于 Threshold Cntrs 值的区域会生成独立的同心波图案。区域的总亮度（亮度 x 面积）会缩放这些波的变形幅度。此素材通常是绘制的不同大小和亮度的点的图像。如果绘制的中心随时间移动，效果中心将随之移动。

- **Matte**: 默认为无。如果提供，变形的幅度将按此输入素材的值进行缩放。灰度值在内部缩放变形幅度，而不是简单地在效果和原始源素材之间交叉淡化，以在遮罩边缘获得更连续的结果，并对变形量进行更精细的控制。此输入可通过 Blur Matte、Invert Matte 或 Matte Use 参数进行调整。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用之前按此量模糊 Mocha 遮罩。这可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

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
  在使用之前按此像素量膨胀或侵蚀 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在 High 质量模式下膨胀 Mocha 遮罩，以获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩都提供给效果时，确定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Amplitude** (Default: 1, Range: any)
  缩放变形扭曲的量。增大以获得更强烈的变形。

- **Frequency** (Default: 8, Range: 0.01 or greater)
  波纹的频率。增大以获得更多波纹，减小以获得更少波纹。此参数可以使用 Frequency 控件调整。

- **Rel Height** (Default: 1, Range: 0.01 or greater)
  同心波图案的相对高度。

- **Rotate Rel H** (Default: 0, Range: any)
  围绕每个中心旋转波纹图案的角度（以度为单位）。如果 Rel Height 参数为 1.0，此参数无效。此参数可以使用 Rotate Rel H 控件调整。

- **Threshold Cntrs** (Default: 0.6, Range: 0 or greater)
  亮度高于此值的区域用作波纹中心。中心由每组高于此值的连接像素的质心生成。

- **Max Centers** (Integer, Default: 20, Range: 1 or greater)
  使用的最大中心总数。可用于测试或避免过多的中心数量。

- **Phase Start** (Default: 0, Range: any)
  波纹的相位偏移。

- **Phase Speed** (Default: 0, Range: any)
  波纹的速度。如果为正值，波纹将以此速率自动从中心向外移动。

- **Inner Radius** (Default: 0, Range: any)
  从波纹中心开始引入波纹变形的距离。在此半径内不会生成波纹。此参数可以使用 Inner Radius 控件调整。

- **Inner Softness** (Default: 0.1, Range: 0.0028 or greater)
  在 Inner Radius 处引入波纹变形的区域宽度。

- **Outer Radius** (Default: 0.5, Range: 0 or greater)
  从波纹中心开始淡出波纹变形的距离。在此半径外不会生成波纹。此参数可以使用 Outer Radius 控件调整。

- **Outer Softness** (Default: 0.5, Range: 0.0056 or greater)
  在 Outer Radius 处淡出波纹变形的区域宽度。

- **Z Dist** (Default: 1, Range: 0.001 or greater)
  缩放图像的"距离"。大于 1.0 的值使其更远更小。小于 1.0 的值使图像更近更大。稍微放大有时可用于隐藏边缘伪影。

- **Wrap** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  确定访问源图像边界外区域的方法。
  - **No**: 在边界外显示黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复图像的镜像副本。使用此方法时边缘通常不太明显。

- **Filter** (Check-box, Default: on)
  如果启用，图像在重新采样时会进行自适应滤波。当图像的某些部分被变形缩小时，这会产生更好的质量结果。

- **Blur Matte** (Default: 0, Range: 0 or greater)
  在使用之前按此量模糊 Matte 输入。这可以在遮罩区域和非遮罩区域之间提供更平滑的过渡。除非提供了 Matte 输入，否则此选项无效。

- **Invert Matte** (Check-box, Default: off)
  如果启用，反转 Matte 输入，使效果应用于 Matte 为黑色而非白色的区域。除非提供了 Matte 输入，否则此选项无效。

- **Matte Use** (Popup menu, Default: Luma)
  确定如何使用 Matte 输入通道来创建单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可略微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项也比 Normal 模式渲染略快，但结果也将是预乘形式，有时不太准确。如果图像中遮罩通道也有锐利边缘的区域存在明显的颜色变化，使用 Normal 模式可能会获得更好的结果。

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  这 4 个参数，Crop Top、Crop Bottom、Crop Left 和 Crop Right，允许选择输入图像的矩形子区域进行处理。如果 Wrap 参数设置为 "No"，则暴露的边框将是透明的。如果 Wrap 为 "Tile" 或 "Reflect"，源图像将在新的裁剪边框上包裹以填充画面。这可以更容易地避免因变形具有不良边缘的图像而产生的伪影。

- **Show Frequency** (Check-box, Default: on)
  打开或关闭用于调整 Frequency 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，支持屏幕控件。

- **Show Outer Radius** (Check-box, Default: on)
  打开或关闭用于调整 Rel Height 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，支持屏幕控件。

- **Show Inner Radius** (Check-box, Default: on)
  打开或关闭用于调整 Inner Radius 的屏幕界面参数。Inner Radius 参数的值必须首先为正值才能使此控件可见。此参数仅在 AE 和 Premiere 中出现，支持屏幕控件。

- **Show Rotate Rel H** (Check-box, Default: on)
  打开或关闭用于调整 Rotate Rel H 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，支持屏幕控件。

