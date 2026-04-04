---
title: PixelSort
---

## S_PixelSort

沿各种图案排列的线条对超过阈值的像素进行排序。图案包括平行线、从中心点辐射的线条和圆形线条。

位于 Sapphire Stylize 效果子菜单中。

![PixelSort](../_static/PixelSort.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Matte**: 默认为无。定义将被排序的区域。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: Linear)
  在几种排序图案之间选择。
  - **Linear**: 沿平行线排序像素。
  - **Radial**: 沿从一个点辐射出的线条排序像素。
  - **Circular**: 沿同心圆排序像素。

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

- **Apply Mask** (Popup menu, Default: Post-threshold)
  控制在效果中的哪个阶段应用遮罩——同时影响输入遮罩和 Mocha 遮罩。
  - **Post-threshold**: 在排序任何像素之前将遮罩应用于阈值图。
  - **Pre-effect**: 在运行像素排序之前将遮罩应用于源素材。

- **Sort Angle** (Default: 0, Range: any)
  Linear 模式下平行线排序的角度。

- **Center** (X & Y, Default: [0 0], Range: any)
  Radial 模式下的中心点。

- **Start Angle** (Default: 0, Range: any)
  Radial 模式下开始排序的角度。

- **Degrees Sorted** (Default: 360, Range: 0 to 360)
  Radial 模式下围绕中心应排序多少度。

- **Inner Radius** (Default: 0.1, Range: 0 or greater)
  Radial 模式下中心周围未排序像素的半径。

- **Ray Length** (Default: 0.8, Range: 0 or greater)
  Radial 模式下排序像素块的长度。

- **Vary Radius** (Default: 0.1, Range: 0 to 1)
  Radial 模式下辐射线起始像素的变化量。

- **Start Angle** (Default: 0, Range: any)
  Circular 模式下开始排序的角度。

- **Degrees Sorted** (Default: 270, Range: 0 to 360)
  Circular 模式下围绕中心应排序多少度。

- **Circle Center** (X & Y, Default: [0 0], Range: any)
  Circular 模式下同心圆的中心。

- **Vary Start** (Default: 0.15, Range: 0 to 1)
  Circular 模式下起始位置的变化量。

- **Inner Radius** (Default: 0, Range: 0 or greater)
  Circular 模式下要排序的最小圆。

- **Thickness** (Default: 1.1, Range: 0 or greater)
  Circular 模式下要排序的同心圆数量。

- **Threshold** (Default: 0.3, Range: any)
  排序像素的阈值。只有阈值一侧的像素会被排序。

- **Sort Direction** (Popup menu, Default: sort above threshold)
  控制是排序阈值以上还是以下的像素。
  - **sort below threshold**: 排序值低于阈值的像素。
  - **sort above threshold**: 排序值高于阈值的像素。

- **Reverse Sort Direction** (Check-box, Default: off)
  决定是按升序还是降序排列像素。

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

- **Randomly Restart Sort** (Default: 100, Range: 0 to 1000)
  中断排序像素块的频率。

- **Soften Threshold Mask** (Default: 0.1, Range: 0 or greater)
  模糊通过阈值处理图像生成的遮罩。

- **Downsample** (Check-box, Default: off)
  如果勾选，根据 Sort Resolution 降低输出分辨率。

- **Sort Resolution** (Integer, Default: 720, Range: 1 or greater)
  勾选 Downsample 时以像素为单位的目标输出分辨率。

- **Seed** (Default: 0.273, Range: 0 or greater)
  初始化排序重启的随机数生成器。

- **Mix With Source** (Default: 0, Range: 0 to 1)
  在排序后的图像和原始源素材之间进行插值。

- **Show** (Popup menu, Default: Result)
  选择输出选项。
  - **Result**: 显示像素排序的结果。
  - **Raw Sort Values**: 显示像素排序将用于排序的值。
  - **Threshold Mask**: 显示基于阈值参数要排序的像素遮罩。
  - **Random Restart Noise**: 显示将导致排序线重新开始的像素。
  - **Combined Threshold Mask**: 如果 apply_mask 设为 Mask Threshold，则显示阈值遮罩与输入遮罩的组合。

- **Matte Use** (Popup menu, Default: Luma)
  决定如何使用 Matte 输入通道生成单色蒙版。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Matte** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊蒙版输入。可以在蒙版区域和非蒙版区域之间提供更平滑的过渡。除非提供了蒙版输入，否则无效。

- **Invert Matte** (Check-box, Default: off)
  如果开启，反转蒙版输入，使效果应用于蒙版为黑色而非白色的区域。除非提供了蒙版输入，否则无效。

- **Soft Borders** (Check-box, Default: off)
  如果启用，在处理前为输入图像添加透明边框。这允许结果包含超出原始图像大小的柔和边缘。关闭时，效果仅在画面内发生，结果将在边框处保留边缘。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可略微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 将图像视为已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式略快，但结果也将是预乘形式，有时不太正确。

