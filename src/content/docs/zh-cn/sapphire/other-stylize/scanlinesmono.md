---
title: ScanLinesMono
---

## S_ScanLinesMono

ScanLines 的单色版本。创建源素材类似黑白电视监视器扫描线图案的版本。增加 Add Noise 参数还可为结果添加颗粒效果。

在 Sapphire Stylize 效果子菜单中。

![ScanLinesMono](../_static/ScanLinesMono.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。在结果和源输入之间进行插值。白色区域使用效果结果，黑色区域使用源素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前对 Mocha 遮罩进行此量的模糊处理。可用于柔化遮罩的边缘或量化伪像，并平滑时间位移。

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
  在使用前对 Mocha 遮罩进行此像素量的膨胀或腐蚀处理。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认 Fast 模式下快速调整还是在 High 质量模式下获得更好效果。
  - **Fast**: 以 Fast 模式进行 Dilate Mocha，用于快速调整。
  - **High**: 以 High 质量模式进行 Dilate Mocha，获得更好看的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩并将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果并仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  决定当效果同时提供 Mocha 遮罩和输入遮罩时如何合并它们。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Lines Frequency** (Default: 50, Range: 1 or greater)
  屏幕上扫描线的频率。增加可获得更多扫描线，减少可获得更少。

- **Lines Sharpness** (Default: 1, Range: 0 or greater)
  缩放线条的强度。增加可获得更锐利的边缘，减少可获得更细微的效果。锐利度为零时扫描线效果消失。

- **Lines Angle** (Default: 0, Range: any)
  扫描线的角度（单位为度）。设为 90 可获得垂直线而非水平线。可使用 Lines Angle 控件调整此参数。

- **Lines Shift** (Default: 0, Range: any)
  偏移线条图案的位置。值为 1.0 时将整条扫描线向前偏移一格，效果与 0 相同。

- **Add Noise** (Default: 0, Range: 0 or greater)
  如果为正值，将向图像添加此量的黑白噪点。

- **Noise Freq Rel** (Default: 1, Range: 0.01 or greater)
  噪点频率，相对于线条频率。除非上方的 Add Noise 参数为正值，否则此参数无效。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果的亮度。

- **Color1** (Default rgb: [1 1 1])
  扫描线图案中较亮的颜色。

- **Color0** (Default rgb: [0 0 0])
  扫描线图案中较暗的颜色。

- **Offset** (Default: 0, Range: -8 to 2)
  向结果添加此灰度值（如为负则减去）。0 无效果，0.5 为中灰，1 为白色。

- **Gamma** (Default: 1.5, Range: 0.1 to 10)
  使用此 Gamma 值通过曲线缩放图像亮度，允许调整扫描线中的中间灰色值。这有助于使输出的平均亮度与输入匹配。

- **Smooth Source** (Default: 0, Range: 0 or greater)
  如果为正值，在处理之前将源素材模糊此量。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Mask Use** (Popup menu, Default: Luma)
  决定如何使用遮罩输入通道来制作单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前对遮罩输入进行此量的模糊处理。这可以在遮罩和非遮罩区域之间提供更平滑的过渡。除非提供了遮罩输入，否则此参数无效。

- **Invert Mask** (Check-box, Default: off)
  如果启用，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则此参数无效。

- **Show Lines Angle** (Check-box, Default: off)
  开启或关闭用于调整 Lines Angle 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
