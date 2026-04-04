---
title: DissolvePixelSort
---

## S_DissolvePixelSort

在两个输入素材之间进行转场，同时对溶解结果进行像素排序。

在 Sapphire Transitions 效果子菜单中。

![DissolvePixelSort](../_static/DissolvePixelSort.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始转场。

- **Background**: 默认为无。以此素材结束转场。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: DissolvePixelSortLinear)
  在多种排序图案之间进行选择。
  - **DissolvePixelSortLinear**: 沿平行线对像素进行排序。
  - **DissolvePixelSortRadial**: 沿从某点向外辐射的线条对像素进行排序。
  - **DissolvePixelSortCircular**: 沿同心圆对像素进行排序。

- **Transition Dir** (Popup menu, Default: Dissolve Off to Bg)
  选择转场的方向。
  - **Dissolve Off to Bg**: 从当前图层转场到背景。
  - **Dissolve On from Bg**: 从背景转场到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  如果启用，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过动画 Dissolve Percent 参数手动执行转场。

- **Dissolve Percent** (Default: 0, Range: 0 to 1)
  必须禁用 Auto Trans 才能使用此参数。它决定前景和背景输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制溶解的时间。

- **Dissolve Speed** (Default: 5, Range: 1 or greater)
  From 和 To 素材之间的溶解速度。设为 1 时，溶解在效果的整个持续时间内进行。设为更高值时，溶解时间更短，但像素排序的渐入和渐出仍占用整个持续时间。设为 10 可使转场更快捷，更像闪帧切换。

- **Max Percentage** (Default: 1, Range: 0 or greater)
  转场中最大或完全排序时刻的阈值。设为 1 可在转场高峰时实现最大排序。

- **Sort Angle** (Default: 0, Range: any)
  在 Linear 模式下对平行线进行排序的角度。

- **Center** (X & Y, Default: [0 0], Range: any)
  Radial 模式下的中心点。

- **Start Angle** (Default: 0, Range: any)
  在 Radial 模式下开始排序的角度。

- **Degrees Sorted** (Default: 360, Range: 0 to 360)
  在 Radial 模式下围绕中心应排序的度数。

- **Inner Radius** (Default: 0, Range: 0 or greater)
  在 Radial 模式下中心周围未排序像素的半径。

- **Ray Length** (Default: 1, Range: 0 or greater)
  在 Radial 模式下已排序像素块的长度。

- **Vary Radius** (Default: 0, Range: 0 to 1)
  在 Radial 模式下辐射线起始像素的变化程度。

- **Start Angle** (Default: 0, Range: any)
  在 Circular 模式下开始排序的角度。

- **Degrees Sorted** (Default: 360, Range: 0 to 360)
  在 Circular 模式下围绕中心应排序的度数。

- **Circle Center** (X & Y, Default: [0 0], Range: any)
  在 Circular 模式下同心圆的中心。

- **Vary Start** (Default: 0, Range: 0 to 1)
  在 Circular 模式下起始点的变化程度。

- **Inner Radius** (Default: 0, Range: 0 or greater)
  在 Circular 模式下要排序的最小圆。

- **Thickness** (Default: 1, Range: 0 or greater)
  在 Circular 模式下要排序的同心圆数量。

- **Sort Direction** (Popup menu, Default: sort above threshold)
  控制是对阈值以上还是以下的像素进行排序。
  - **sort below threshold**: 对值低于阈值的像素进行排序。
  - **sort above threshold**: 对值高于阈值的像素进行排序。

- **Reverse Sort Direction** (Check-box, Default: off)
  决定是按升序还是降序对像素进行排序。

- **Randomly Restart Sort** (Default: 100, Range: 0 to 1000)
  中断已排序像素块的频率。

- **Sort Type** (Popup menu, Default: monochrome)
  决定如何分析像素的颜色以进行排序。
  - **monochrome**: 按像素的单色值排序。
  - **average**: 按像素颜色通道的平均值排序。
  - **minimum**: 按像素中最小的通道排序。
  - **maximum**: 按像素中最大的通道排序。
  - **red**: 按像素的红色通道排序。
  - **green**: 按像素的绿色通道排序。
  - **blue**: 按像素的蓝色通道排序。
  - **hue**: 按像素的色相排序。
  - **saturation**: 按像素的饱和度排序。
  - **brightness**: 按像素的亮度排序。

- **Seed** (Default: 0.273, Range: 0 or greater)
  初始化排序重启的随机数生成器。

- **Soften Threshold Mask** (Default: 0.1, Range: 0 or greater)
  对图像阈值化生成的遮罩进行模糊处理。

- **Downsample** (Check-box, Default: off)
  如果勾选，则根据 Sort Resolution 降低输出分辨率。

- **Sort Resolution** (Integer, Default: 720, Range: 1 or greater)
  勾选 Downsample 时的目标输出分辨率（像素）。

- **Mix With Dissolve** (Default: 0, Range: 0 to 1)
  通过在原始溶解和像素排序结果之间插值来柔化像素排序的外观。

- **Show** (Popup menu, Default: Result)
  选择输出选项。
  - **Result**: 显示像素排序的结果。
  - **Raw Sort Values**: 显示像素排序将用于排序的值。
  - **Threshold Mask**: 显示基于阈值参数要排序的像素遮罩。
  - **Random Restart Noise**: 显示将导致排序线重启的像素。
  - **Combined Mask**
