---
title: OCIOTransform
---

## S_OCIOTransform

使用 OpenColorIO 将输入从一个色彩空间转换到另一个色彩空间。

使用 OCIO 环境变量查找您的 OpenColorIO 配置。如果未设置，将回退到默认配置。不建议在将此效果添加到项目后更改环境变量，因为这可能会改变渲染结果。如果发生这种情况，效果会显示警告对话框。

在 Sapphire Adjust 效果子菜单中。

![OCIOTransform](../_static/OCIOTransform.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。如果提供，效果仅应用于此输入的亮区所指定的源素材区域。此遮罩之外的像素不受影响，也不会对其内部的受影响像素产生贡献。可以使用 Invert Mask 或 Mask Use 参数来影响此输入。


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
  在使用前按此像素量膨胀或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下进行快速调整。
  - **High**: 在 High 质量模式下获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩同时提供给效果时，确定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **From** (Popup menu)
  输入素材的原始色彩空间。此菜单列出您的 OpenColorIO 配置中所有可用的色彩空间。

- **To** (Popup menu)
  所需的输出色彩空间。此菜单列出您的 OpenColorIO 配置中所有可用的色彩空间。

- **Switch To And From** (Check-box, Default: off)
  如果勾选，通过交换 To 和 From 色彩空间来反转变换方向。

- **Config** (Popup menu, Default: ConfigPath)
  此菜单显示当前 OpenColorIO 配置文件的路径。
  - **ConfigPath**: 当前 OpenColorIO 配置文件的位置和名称。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用遮罩输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊遮罩输入。这可以在遮罩区域和非遮罩区域之间提供更平滑的过渡。除非提供了遮罩输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果启用，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则无效。
