---
title: Median
---

## S_Median

对源图像应用中值滤波器。中值滤波器适用于清除孤立斑点和噪点。

位于 Sapphire Blur+Sharpen 效果子菜单中。

![Median](../_static/Median.jpg)


### Inputs:

- **Source**: 当前图层。要过滤的素材。

- **Mask**: 默认为无。在结果和 Source 输入之间进行插值。白色区域使用效果的结果。黑色区域使用 Source 素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: Median)
  选择是对所有通道应用相同的中值滤波器，还是对每个通道分别应用中值滤波器。
  - **Median**: 对每个通道应用相同的中值滤波器。
  - **MedianChannels**: 对红、绿、蓝通道分别应用不同的中值滤波器。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此量模糊 Mocha 遮罩。可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

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
  在使用前按此像素量膨胀或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在 High 质量模式下膨胀 Mocha 遮罩，以获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩并将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果并显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当同时为效果提供 Mocha 遮罩和输入遮罩时，确定如何组合它们。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Size** (Default: 1, Range: 0.1 to 40)
  中值滤波器的大小。

- **Subpixel** (Check-box, Default: off)
  启用亚像素宽度滤波。使用此选项可使 Size 参数的动画更加平滑。

- **Size Rel X** (Default: 1, Range: 0.1 to 5)
  滤波器的相对水平大小。

- **Size Rel Y** (Default: 1, Range: 0.1 to 5)
  滤波器的相对垂直大小。

- **Red Rel Size X** (Default: 1, Range: 0 to 2)
  红色通道中滤波器的相对水平大小。

- **Red Rel Size Y** (Default: 1, Range: 0 to 2)
  红色通道中滤波器的相对垂直大小。

- **Green Rel Size X** (Default: 1, Range: 0 to 2)
  绿色通道中滤波器的相对水平大小。

- **Green Rel Size Y** (Default: 1, Range: 0 to 2)
  绿色通道中滤波器的相对垂直大小。

- **Blue Rel Size X** (Default: 1, Range: 0 to 2)
  蓝色通道中滤波器的相对水平大小。

- **Blue Rel Size Y** (Default: 1, Range: 0 to 2)
  蓝色通道中滤波器的相对垂直大小。

- **Alpha Rel Size X** (Default: 1, Range: 0 to 2)
  Alpha 通道中滤波器的相对水平大小。

- **Alpha Rel Size Y** (Default: 1, Range: 0 to 2)
  Alpha 通道中滤波器的相对垂直大小。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用 Mask 输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此量模糊 Matte 输入。这可以在遮罩区域和非遮罩区域之间提供更平滑的过渡。除非提供了 Matte 输入，否则此参数无效。

- **Invert Mask** (Check-box, Default: off)
  如果开启，反转 Matte 输入，使效果应用于 Matte 为黑色而非白色的区域。除非提供了 Matte 输入，否则此参数无效。
