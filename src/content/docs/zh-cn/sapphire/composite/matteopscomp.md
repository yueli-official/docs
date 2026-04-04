---
title: MatteOpsComp
---

## S_MatteOpsComp

扩展、收缩或向前景 Alpha 通道的边缘添加噪波，然后使用该结果将前景合成到背景之上。这对于去除色键抠像中的蓝色或绿色溢出很有用。

在 Sapphire Composite 效果子菜单中。

![MatteOpsComp](../_static/MatteOpsComp.jpg)


### Inputs:

- **Foreground**: 当前图层。用作前景的片段。

- **Background**: 默认为无。用作背景的片段。

- **Matte**: 默认为无。要处理的遮罩输入片段。如果未提供此输入，则使用前景输入代替。假定遮罩具有抗锯齿但硬的边缘，因为非常柔和的边缘可能不会以有用的方式受到影响。此输入可受 Invert Matte 或 Matte Use 参数的影响。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此量模糊 Mocha 遮罩。这可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 遮罩的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  如果启用，则在应用效果之前反转 Mocha 遮罩的黑白。

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
  选择 Dilate Mocha 是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在高质量模式下膨胀 Mocha 遮罩，以获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源片段。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  确定当两个遮罩同时提供给效果时，如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Shrink- Grow+** (Default: 0, Range: any)
  以近似像素为单位扩展遮罩边缘，如果为负值则收缩。

- **Edge Softness** (Default: 1, Range: 0.01 or greater)
  边缘的最终柔和度。

- **Post Blur** (Default: 0, Range: 0 or greater)
  如果为正值，则按此量模糊结果。这是柔化边缘的另一种方法。

- **Filter** (Popup menu, Default: Triangle)
  用于收缩或扩展过程的模糊滤镜类型。
  - **Box**: 使用矩形滤镜。
  - **Triangle**: 更平滑，使用金字塔形滤镜。
  - **Gauss**: 最平滑，使用高斯形滤镜。

- **Matte Use** (Popup menu, Default: Alpha)
  确定如何使用遮罩输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Invert Matte** (Check-box, Default: off)
  如果启用，则反转输出遮罩的黑白。

- **Soft Borders** (Check-box, Default: off)
  如果启用，在处理之前向输入图像添加透明边框。这允许结果包含超出原始图像大小的柔和边缘。关闭时，效果仅发生在帧内，结果将在边界处保留边缘。

- **Noise Amplitude** (Default: 0, Range: 0 or greater)
  添加到边缘的噪波纹理量。

- **Noise Width** (Default: 0.0224, Range: 0 or greater)
  在遮罩边缘包含噪波的区域宽度。除非 Noise Amplitude 为正值，否则此参数无效。

- **Frequency** (Default: 100, Range: 0.1 or greater)
  噪波的频率。增大以获得更细的颗粒噪波，减小以获得更粗的噪波。除非 Noise Amplitude 为正值，否则此参数无效。

- **Frequency Rel X** (Default: 1, Range: 0.01 or greater)
  噪波的相对水平频率。增大以垂直拉伸噪波，减小以水平拉伸噪波。除非 Noise Amplitude 为正值，否则此参数无效。

- **Octaves** (Integer, Default: 1, Range: 1 to 10)
  噪波的叠加层数。每个八度的频率是前一个的两倍，幅度是前一个的一半。除非 Noise Amplitude 为正值，否则此参数无效。

- **Seed** (Default: 0.23, Range: 0 or greater)
  用于初始化随机数生成器。实际的种子值并不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Noise Shift** (X & Y, Default: [0 0], Range: any)
  噪波纹理的水平和垂直位移。

- **Jitter Frames** (Integer, Default: 1, Range: 0 or greater)
  如果为 0，噪波纹理将在每一帧处理时保持不变。如果为 1，每一帧使用新的噪波纹理。如果为 2，每隔一帧使用新的噪波纹理，以此类推。
