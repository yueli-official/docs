---
title: WarpCornerPin
---

## S_WarpCornerPin

对源图像执行 3D 透视变形，将角点与四个指定点对齐。这对于将源素材定位在另一个素材中的对象上非常有用，例如广告牌或电脑屏幕。

在 Sapphire Distort 效果子菜单中。

![WarpCornerPin](../_static/WarpCornerPin.jpg)


### Inputs:

- **Source**: 当前图层。要进行变形的输入素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Corner1** (X & Y, Default: [-0.806 0.133], Range: any)
  源素材左下角的位置。此参数可以使用 Corner1 控件调整。

- **Corner2** (X & Y, Default: [0.389 0.577], Range: any)
  源素材右下角的位置。此参数可以使用 Corner2 控件调整。

- **Corner3** (X & Y, Default: [-0.556 -0.441], Range: any)
  源素材右上角的位置。此参数可以使用 Corner3 控件调整。

- **Corner4** (X & Y, Default: [0.611 -0.485], Range: any)
  源素材左上角的位置。此参数可以使用 Corner4 控件调整。

- **Filter** (Check-box, Default: on)
  如果启用，图像在重新采样时会进行自适应滤波。当图像的某些部分被变形缩小时，这会产生更好的质量结果。

- **Bulge** (X & Y, Default: [0 0], Range: -1 to 1)
  扭曲变形图像的透视效果，使其看起来向一个方向膨胀。值为 1 不产生变形。小于 1 的值使图像向右上角拉伸，大于 1 的值使其向左下角拉伸。

- **Wrap** (X & Y, Popup menu, Default: [ No No ])
  确定访问源图像边界外区域的方法。
  - **No**: 在边界外显示黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复图像的镜像副本。使用此方法时边缘通常不太明显。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可略微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项也比 Normal 模式渲染略快，但结果也将是预乘形式，有时不太准确。如果图像中遮罩通道也有锐利边缘的区域存在明显的颜色变化，使用 Normal 模式可能会获得更好的结果。

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  这 4 个参数，Crop Top、Crop Bottom、Crop Left 和 Crop Right，允许选择输入图像的矩形子区域进行处理。如果 Wrap 参数设置为 "No"，则暴露的边框将是透明的。如果 Wrap 为 "Tile" 或 "Reflect"，源图像将在新的裁剪边框上包裹以填充画面。这可以更容易地避免因变形具有不良边缘的图像而产生的伪影。

- **Show Corner1** (Check-box, Default: on)
  打开或关闭用于调整 Corner1 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，支持屏幕控件。

- **Show Corner2** (Check-box, Default: on)
  打开或关闭用于调整 Corner2 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，支持屏幕控件。

- **Show Corner3** (Check-box, Default: on)
  打开或关闭用于调整 Corner3 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，支持屏幕控件。

- **Show Corner4** (Check-box, Default: on)
  打开或关闭用于调整 Corner4 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，支持屏幕控件。

