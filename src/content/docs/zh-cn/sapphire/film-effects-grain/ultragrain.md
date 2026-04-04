---
title: UltraGrain
---

## S_UltraGrain

向源片段添加模拟的数字摄像机颗粒。

在 Sapphire Stylize 效果子菜单中。
### Inputs:

- **Source**: 当前图层。要处理的片段。

- **Matte**: 默认为无。在结果和 Source 输入之间进行插值。白色区域使用效果的结果。黑色区域使用 Source 片段。


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


### Grain Parameters:

Grain Amp:
*Default:
*1,
*Range:
*0 to 2.缩放添加到结果中的颗粒的振幅。设为 0 可禁用所有颗粒。

Grain Size:
*Default:
*1,
*Range:
*0.001 or greater.缩放添加到结果中的颗粒大小。在更改分辨率时很有用，将此缩放因子设置为与分辨率变化匹配。

Grain Blur:
*Default:
*0.00567,
*Range:
*0 or greater.按此数量平滑颗粒。增大以获得更粗的颗粒。

Grain Amp Red:
*Default:
*3.29,
*Range:
*0 or greater.缩放红色颗粒振幅。

Grain Amp Green:
*Default:
*3,
*Range:
*0 or greater.缩放绿色颗粒振幅。

Grain Amp Blue:
*Default:
*3.52,
*Range:
*0 or greater.缩放蓝色颗粒振幅。请注意，颗粒会在图像上加减，因此例如增大 Grain Amp Blue 会同时放大蓝色和黄色斑点。

Grain Blur Red:
*Default:
*0.572,
*Range:
*0 or greater.红色颗粒的相对模糊量。

Grain Blur Green:
*Default:
*0.561,
*Range:
*0 or greater.绿色颗粒的相对模糊量。

Grain Blur Blue:
*Default:
*0.528,
*Range:
*0 or greater.蓝色颗粒的相对模糊量。

Grain Amp Darks:
*Default:
*0.37,
*Range:
*0 to 2.每个通道中应用于图像最暗区域的颗粒相对量。图像中的暗源强度定义为黑色 (0 0 0)。

Grain Amp Mids:
*Default:
*1,
*Range:
*0 to 2.每个通道中应用于图像中间调区域的颗粒相对量。图像中的中间调源强度由 Midtone Pos 参数定义。

Grain Amp Brights:
*Default:
*0,
*Range:
*0 to 2.每个通道中应用于图像最亮区域的颗粒相对量。图像中的亮源强度定义为白色 (1 1 1)。

Midtone Pos Red:
*Default:
*0.27,
*Range:
*0 to 1.红色通道中的中间调位置。红色颗粒振幅从黑色处的 Grain Amp Darks 插值到此中间调位置处的 Grain Amp Mids，然后到白色处的 Grain Amp Brights。整条曲线再由 Grain Amp Red 参数缩放。

Midtone Pos Green:
*Default:
*0.3,
*Range:
*0 to 1.绿色通道中的中间调位置。绿色颗粒振幅从黑色处的 Grain Amp Darks 插值到此中间调位置处的 Grain Amp Mids，然后到白色处的 Grain Amp Brights。整条曲线再由 Grain Amp Green 参数缩放。

Midtone Pos Blue:
*Default:
*0.31,
*Range:
*0 to 1.蓝色通道中的中间调位置。蓝色颗粒振幅从黑色处的 Grain Amp Darks 插值到此中间调位置处的 Grain Amp Mids，然后到白色处的 Grain Amp Brights。整条曲线再由 Grain Amp Blue 参数缩放。

Grain Mono:
*Check-box, Default:
*off.启用后，红色、绿色和蓝色通道使用相同的颗粒图案。要制作真正的单色颗粒，还应将 Grain Amp Red/Green/Blue 设为相等，确保 Midtone Pos Red/Green/Blue 相等，如果 GrainBlur 为正值，还应将 Grain Blur Red/Green/Blue 设为相等。

Grain Seed:
*Default:
*0.123,
*Range:
*0 or greater.
初始化颗粒生成的随机数生成器。实际种子值并不重要，但不同的种子会产生不同的颗粒图案，相同的值应产生可重复的图案。
