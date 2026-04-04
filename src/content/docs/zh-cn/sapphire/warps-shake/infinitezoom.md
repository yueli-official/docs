---
title: InfiniteZoom
---

## S_InfiniteZoom

无限缩放到图像的无尽重复副本中，
让人联想到某些 M.C. Escher 的画作。最适合具有透明边缘的素材，
例如时钟或盘子；或具有透明中心的素材，例如相框。透明度可以来自
源素材的 Alpha 或遮罩。对 Zoom 参数设置动画以获得完整效果。

在 Sapphire Distort 效果子菜单中。

![InfiniteZoom](../_static/InfiniteZoom.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。定义源素材的透明区域。


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

- **Transparent Area** (Popup menu, Default: Outside)
  如果源素材中心是透明的（例如相框），请将此设置为 Inside。如果边缘是透明的（例如绿幕人物或带遮罩的花朵），或没有透明度，请将此设置为 Outside。
  - **Inside**: 用于相框或任何具有透明中心区域的素材，使副本出现在空白区域内。这使副本出现在原始图像后面。
  - **Outside**: 用于花朵、盘子或时钟等对象外部背景已设为透明的素材，使较小的副本出现在原始图像前面。如果没有任何透明度，也使用此模式。

- **Shrink Per Level** (Default: 0.5, Range: 0.001 to 0.9)
  每个较小副本相对于前一个较大副本的缩小量。0.8 表示每一级将是前一级的 0.8 倍大小，因此较小的值意味着每一级的副本明显更小。较大的值使级别之间的间距更近。

- **Zoom** (Default: 1, Range: 0.001 or greater)
  图像的整体缩放。通常需要对此参数设置动画以获得无限缩放效果。线性动画应能产生平滑的缩放效果。

- **Zoom Center** (X & Y, Default: [0 0], Range: any)
  无限缩放的中心点。在缩放的同时对此参数设置动画可以获得有趣的效果。此参数可以使用 Zoom Center 控件调整。

- **Twist** (Default: 0, Range: -5 or greater)
  级别之间的扭曲量。增大或减小以获得螺旋缩放效果。在 No Spiral 模式下，以每单位扭曲 30 度为单位。在螺旋模式下是非线性的，最好通过目视调整。

- **Spiral Strands** (Popup menu, Default: 1 Counterclockwise)
  设置为 No Spiral 时，效果在每个级别制作图像的直接副本（根据 Twist 值仍可能有扭曲）。使用其他螺旋选项时，会变形每个图像副本，使每个级别无缝连接到下一个级别，形成不断缩小的螺旋。
  - **No Spiral**: 不对图像进行变形以提供连续螺旋。适用于相框。
  - **1 Clockwise**: 变形图像以创建一个顺时针方向的连续螺旋线。
  - **1 Counterclockwise**: 变形图像以创建一个逆时针方向的连续螺旋线。
  - **2 Clockwise**: 变形图像以创建两个顺时针方向的连续螺旋线。
  - **2 Counterclockwise**: 变形图像以创建两个逆时针方向的连续螺旋线。

- **Rotate** (Default: 0, Range: any)
  结果图像的整体旋转。

- **Shift** (X & Y, Default: [0 0], Range: any)
  结果图像的整体偏移。

- **Wrap** (Popup menu, Default: No)
  设置源图像边界外像素的处理方式。Reflect 在 Transparent Area: Inside 时可用于填充小黑色区域，Tile 在 Transparent Area: Outside 时可产生有趣的效果，只要图像边缘有一些透明度。如果素材没有透明度，请保持默认值（None）。
  - **No**: 在边界外显示黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复图像的镜像副本。使用此方法时边缘通常不太明显。

- **Matte Use** (Popup menu, Default: Luma)
  确定如何使用 Matte 输入通道来创建单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Invert Matte** (Check-box, Default: off)
  如果启用，反转 Matte 输入，使效果应用于 Matte 为黑色而非白色的区域。除非提供了 Matte 输入，否则此选项无效。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可略微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项也比 Normal 模式渲染略快，但结果也将是预乘形式，有时不太准确。

- **Show Zoom Center** (Check-box, Default: on)
  打开或关闭用于调整 Zoom Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，支持屏幕控件。

