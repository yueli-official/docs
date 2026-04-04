---
title: Grain
---

## S_Grain

向源片段添加彩色和/或单色颗粒。振幅和频率参数允许独立调整所有颜色的颗粒纹理、每个颜色通道的颗粒纹理，或黑白颗粒。

在 Sapphire Stylize 效果子菜单中。

![Grain](../_static/Grain.jpg)


### Inputs:

- **Source**: 当前图层。要处理的片段。

- **Mask**: 默认为无。在结果和 Source 输入之间进行插值。白色区域使用效果的结果。黑色区域使用 Source 片段。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成蒙版。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此数量模糊 Mocha 蒙版。可用于柔化蒙版的边缘或量化伪影，并平滑时间位移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 蒙版的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  如果启用，在应用效果之前反转 Mocha 蒙版的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 蒙版。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 蒙版的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 蒙版的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 蒙版的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素数量膨胀或腐蚀 Mocha 蒙版。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下膨胀 Mocha 蒙版，以便快速调整。
  - **High**: 在 High 质量模式下膨胀 Mocha 蒙版，以获得更好的蒙版形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 蒙版，将效果应用于整个源片段。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 蒙版本身。

- **Combine Masks** (Popup menu, Default: Union)
  当同时提供 Mocha 蒙版和输入蒙版时，确定如何组合它们。
  - **Union**: 使用两个蒙版共同覆盖的区域。
  - **Intersect**: 使用两个蒙版之间重叠的区域。
  - **Mocha Only**: 忽略输入蒙版，仅使用 Mocha 蒙版。

- **Color Scale** (Default rgb: [1 1 1])
  按此值缩放颗粒的颜色。颗粒将包含此颜色的正值和负值。

- **Color Amplitude** (Default: 0.1, Range: 0 or greater)
  要包含的彩色颗粒的振幅。

- **Color Frequency** (Default: 100, Range: 0.1 or greater)
  彩色颗粒的频率。增大以获得更细的彩色颗粒，减小以获得更粗的彩色颗粒。

- **Red Freq** (Default: 1, Range: 0.01 or greater)
  红色通道颗粒的相对频率。

- **Green Freq** (Default: 1, Range: 0.01 or greater)
  绿色通道颗粒的相对频率。

- **Blue Freq** (Default: 1, Range: 0.01 or greater)
  蓝色通道颗粒的相对频率。

- **Color Octaves** (Integer, Default: 1, Range: 1 to 10)
  要包含的彩色颗粒八度数。每个八度的频率是前一个的两倍，振幅是前一个的一半。

- **Bw Amplitude** (Default: 0, Range: 0 or greater)
  要包含的黑白颗粒的振幅。

- **Bw Frequency** (Default: 100, Range: 0.1 or greater)
  黑白颗粒的频率。增大以获得更细的颗粒，减小以获得更粗的颗粒。

- **Bw Octaves** (Integer, Default: 1, Range: 1 to 10)
  要包含的黑白颗粒八度数。每个八度的频率是前一个的两倍，振幅是前一个的一半。

- **Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化随机数生成器。实际种子值并不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Jitter Frames** (Integer, Default: 1, Range: 0 or greater)
  如果为 0，则每帧处理时噪声纹理保持不变。如果为 1，则每帧使用新的噪声纹理。如果为 2，则每隔一帧使用新的噪声纹理，依此类推。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用 Mask 输入通道来创建单色蒙版。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数量模糊 Matte 输入。可以在蒙版区域和非蒙版区域之间提供更平滑的过渡。除非提供了 Matte 输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果开启，反转 Matte 输入，使效果应用于 Matte 为黑色而非白色的区域。除非提供了 Matte 输入，否则无效。

