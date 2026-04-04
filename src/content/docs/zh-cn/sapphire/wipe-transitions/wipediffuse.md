---
title: WipeDiffuse
---

## S_WipeDiffuse

在转场区域内执行像素扩散处理，实现两个输入素材之间的擦除转场。应对 Wipe Percent 参数设置动画以控制转场速度。此效果的像素化外观取决于图像分辨率，因此建议在处理前测试最终分辨率。

在 Sapphire Transitions 效果子菜单中。

![WipeDiffuse](../_static/WipeDiffuse.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始转场。

- **Background**: 默认为无。以此素材结束转场。如果未提供此输入，将使用完全透明的背景，显示其后面的内容。请注意，除非提供此输入，否则在转场过程中无法对背景进行扩散处理。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Transition Dir** (Popup menu, Default: Wipe Off to Bg)
  选择转场的方向。
  - **Wipe Off to Bg**: 从当前图层转场到背景。
  - **Wipe On from Bg**: 从背景转场到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  如果启用，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过对 Wipe Percent 参数设置动画来手动执行转场。

- **Wipe Percent** (Default: 0, Range: 0 to 1)
  必须禁用 Auto Trans 才能使用此参数。它决定了 From 和 To 输入之间的转场比例，通常应从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线以更精细地控制擦除时序。

- **Edge Width** (Default: 1.4, Range: 0.0138 or greater)
  转场区域的宽度。可以使用 Wipe Widget 调整此参数。

- **Angle** (Default: 0, Range: any)
  擦除方向的角度（从右侧起以度为单位）。可以使用 Wipe Widget 调整此参数。

- **Diffuse Amount** (Default: 0.4, Range: 0 or greater)
  像素扩散的幅度。

- **Wrap** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  确定访问源图像边界之外区域的方法。
  - **No**: 边界之外显示为黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复图像的镜像副本。使用此方法边缘通常不太明显。

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  这4个参数（Crop Top、Crop Bottom、Crop Left 和 Crop Right）允许选择要处理的输入图像的矩形子区域。如果 Wrap 参数设置为"No"，则暴露的边界将是透明的。如果 Wrap 设置为"Tile"或"Reflect"，则源图像将在新裁剪的边界上环绕以填充画面。这可以更容易地避免因扭曲具有不良边缘的图像而产生的伪影。

- **Show Wipe** (Check-box, Default: on)
  打开或关闭用于调整 Wipe Amt、Angle 和 Edge Width 参数的屏幕用户界面控件。此参数仅在 AE 和 Premiere 中显示，因为这些软件支持屏幕控件。
