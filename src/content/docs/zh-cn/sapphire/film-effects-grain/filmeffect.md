---
title: FilmEffect
---

## S_FilmEffect

提供一个物理精确的胶片曝光和冲洗模型，使您的视频素材看起来像是在特定的胶片上拍摄的。它可以去除场伪影、针对特定胶片类型进行色彩校正、添加胶片颗粒，并应用辉光或柔焦效果。色彩校正和颗粒可以通过 Scale CC 和 Grain Amp 参数选择性地禁用。

在 Sapphire Stylize 效果子菜单中。

![FilmEffect](../_static/FilmEffect.jpg)


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

- **Neg Film** (Popup menu, Default: Kodak 5245)
  选择负片胶片类型。
  - **None**: 忽略负片胶片的任何效果。除非您同时将正片胶片也选为 None 以禁用两者，否则通常不太有用。
  - **Kodak 5245**: Eastman EXR 50D，低速，日光平衡，颗粒极细。
  - **Kodak 5246**: Kodak VISION 250D，较高对比度，中速，日光平衡，细颗粒。
  - **Kodak 5248**: Eastman EXR 100T，中速，钨丝灯平衡，颗粒极细。
  - **Kodak 5274**: Kodak VISION 200T，中速，钨丝灯平衡，细颗粒。
  - **Kodak 5277**: Kodak VISION 320T，较低对比度，中速，钨丝灯平衡，中细颗粒。
  - **Kodak 5279**: Kodak VISION 500T，高速，钨丝灯平衡，颗粒较明显。
  - **Kodak 5284**: Kodak VISION Expression 500T，较低对比度，高速，钨丝灯平衡，中等颗粒。
  - **Kodak 5289**: Kodak VISION 800T，超高速，钨丝灯平衡，颗粒明显。
  - **Kodak 5293**: Eastman EXR 200T，降低对比度，钨丝灯平衡，中等颗粒。
  - **Kodak 5298**: Eastman EXR 500T，高速，钨丝灯平衡，颗粒明显。
  - **K SFX200T**: 特效胶片，中等颗粒。
  - **Kodak 5217**: Kodak Vision2 200T，钨丝灯平衡，细颗粒。
  - **Kodak 5218**: Kodak Vision2 500T，钨丝灯平衡，细颗粒。

- **Print Film** (Popup menu, Default: Kodak 2383)
  选择正片胶片类型。
  - **None**: 忽略正片胶片的任何效果。这将直接输出负片。如果负片胶片也设为 None，则色彩校正和颗粒将被禁用。
  - **Kodak 2383**: Kodak VISION 彩色正片，浓郁黑色。
  - **Kodak 2393**: Kodak VISION Premier 彩色正片，浓郁黑色，有些颗粒。
  - **Kodak 2395**: Kodak VISION 彩色电视正片，低对比度。
  - **Kodak 5386**: Eastman EXR 彩色正片（已被柯达停产，由 2383 取代）。
  - **Kodak 5285 Rev**: Ektachrome 100D 反转片，日光平衡，高对比度且颗粒明显。请注意，使用反转片时负片胶片将被忽略。
  - **Kodak 7270 Rev**: Kodachrome 40 电影胶片，钨丝灯平衡反转片，高对比度且颗粒较明显。请注意，使用反转片时负片胶片将被忽略。

- **Blur Input** (Default: 0, Range: 0 or greater)
  按此数量平滑输入。可用于在处理前去除视频噪声或压缩伪影。


### Color Correct Parameters:

Scale CC:
*Default:
*1,
*Range:
*0 to 5.缩放由胶片类型、伽马值和曝光值产生的色彩校正量。设为 0 可禁用色彩校正。如果将此值增大到 1.0 以上，则会夸大色彩校正，通常会增加对比度。

Input Gamma:
*Default:
*2.2,
*Range:
*0.1 to 10.原始片段拍摄时的伽马值。对于视频，通常为 2.2；对于合成计算机图形，可能更低。

Output Gamma:
*Default:
*2.2,
*Range:
*0.1 to 10.输出的预期观看伽马值。

Neg Exposure:
*Default:
*0,
*Range:
*any.调整负片胶片的模拟曝光，以档为单位。增大表示过度曝光且更亮。

Print Exposure:
*Default:
*0,
*Range:
*any.调整正片胶片的模拟曝光，以档为单位。增大表示过度曝光且更暗。

Print Lights Red:
*Default:
*25,
*Range:
*0 to 50.调整正片胶片的红色曝光，以打印灯光点为单位。1 个光点等于 1/12 档。增大以过度曝光红色并获得更偏青色的结果。

Print Lights Green:
*Default:
*25,
*Range:
*0 to 50.调整正片胶片的绿色曝光，以打印灯光点为单位。1 个光点等于 1/12 档。增大以过度曝光绿色并获得更偏洋红的结果。

Print Lights Blue:
*Default:
*25,
*Range:
*0 to 50.调整正片胶片的蓝色曝光，以打印灯光点为单位。1 个光点等于 1/12 档。增大以过度曝光蓝色并获得更偏黄色的结果。

Scale Brights:
*Default:
*1,
*Range:
*0 or greater.在其他色彩校正、辉光和颗粒应用之后，缩放最终结果的亮部区域。（此参数不受 Scale CC 影响。）

Offset Darks:
*Default:
*0,
*Range:
*-8 to 2.
在其他色彩校正、辉光和颗粒应用之后，将此灰度值添加到最终结果的较暗区域。可以为负值以增加对比度。（此参数不受 Scale CC 影响。）

### Glow Parameters:

Glow Brightness:
*Default:
*0,
*Range:
*0 or greater.如果为正值，图像将与其模糊版本合并以产生辉光效果。增大以获得更亮的辉光。

Glow Soft Focus:
*Default:
*0,
*Range:
*0 to 1.如果为正值，图像将与其模糊版本混合以产生柔焦效果。此参数的效果类似于 Glow Brightness，但不会使整体结果变亮。增大此值可混入更多模糊版本和更少原始图像。如果此值为 1 且 Glow Brightness 为 0，您将只获得模糊版本。

Glow Width:
*Default:
*0.224,
*Range:
*0 or greater.辉光和/或柔焦使用的模糊宽度。

Glow Width Red:
*Default:
*1,
*Range:
*0 or greater.红色通道的相对辉光宽度。

Glow Width Green:
*Default:
*1,
*Range:
*0 or greater.绿色通道的相对辉光宽度。

Glow Width Blue:
*Default:
*1,
*Range:
*0 or greater.
蓝色通道的相对辉光宽度。

### Grain Parameters:

Grain Amp:
*Default:
*1,
*Range:
*0 or greater.缩放添加到结果中的胶片颗粒的振幅。设为 0 可禁用所有颗粒。

Grain Amp Red:
*Default:
*0.9,
*Range:
*0 or greater.缩放红色颗粒振幅。

Grain Amp Green:
*Default:
*1,
*Range:
*0 or greater.缩放绿色颗粒振幅。

Grain Amp Blue:
*Default:
*1.6,
*Range:
*0 or greater.缩放蓝色颗粒振幅。请注意，颗粒会在图像上加减，因此例如增大 Grain Amp Blue 会同时放大蓝色和黄色斑点。

Grain Amp Darks:
*Default:
*0.2,
*Range:
*0 to 2.每个通道中应用于图像最暗区域的颗粒相对量。此值默认小于 1.0，因为暗区通常比中间调有更少的颗粒。

Grain Amp Brights:
*Default:
*0,
*Range:
*0 to 2.每个通道中应用于图像最亮区域的颗粒相对量。此值默认为零，因为亮区通常比中间调有更少的颗粒。请注意，高饱和度颜色可能同时受到 Grain Amp Darks 和 Grain Amp Brights 的影响，因为它们在某些颜色通道中较暗，在其他通道中较亮。

Midtone Pos Red:
*Default:
*0.5,
*Range:
*0 to 1.红色通道中通常接收最大颗粒量的中间调位置。红色颗粒振幅从黑色处的 Grain Amp Darks 插值到此中间调位置处的 1.0，然后下降到白色处的 Grain Amp Brights。整条曲线再由 Grain Amp Red 参数缩放。

Midtone Pos Green:
*Default:
*0.5,
*Range:
*0 to 1.绿色通道中通常接收最大颗粒量的中间调位置。绿色颗粒振幅从黑色处的 Grain Amp Darks 插值到此中间调位置处的 1.0，然后下降到白色处的 Grain Amp Brights。整条曲线再由 Grain Amp Green 参数缩放。

Midtone Pos Blue:
*Default:
*0.5,
*Range:
*0 to 1.蓝色通道中通常接收最大颗粒量的中间调位置。蓝色颗粒振幅从黑色处的 Grain Amp Darks 插值到此中间调位置处的 1.0，然后下降到白色处的 Grain Amp Brights。整条曲线再由 Grain Amp Blue 参数缩放。

Grain Blur:
*Default:
*0,
*Range:
*0 or greater.按此数量平滑颗粒。增大以获得更粗的颗粒。

Grain Blur Red:
*Default:
*1,
*Range:
*0 or greater.红色颗粒的相对模糊量。

Grain Blur Green:
*Default:
*0.9,
*Range:
*0 or greater.绿色颗粒的相对模糊量。

Grain Blur Blue:
*Default:
*1.2,
*Range:
*0 or greater.蓝色颗粒的相对模糊量。

Grain Mono:
*Check-box, Default:
*off.启用后，红色、绿色和蓝色通道使用相同的颗粒图案。要制作真正的单色颗粒，还应将 Grain Amp Red/Green/Blue 设为相等，确保 Midtone Pos Red/Green/Blue 相等，如果 GrainBlur 为正值，还应将 Grain Blur Red/Green/Blue 设为相等。

Grain Hold:
*Popup menu, Default: Frame
*.指示应多久生成一次新的颗粒图案。只有当 Grain Blur 为正值使颗粒大小大于一个像素时，您才可能注意到这些选项之间的差异。
*Field:
*保持颗粒图案一个场。*Frame:
*保持颗粒图案一帧（2 个场）。*3:2 Stutter at 0:
*以 3:2 下拉模式保持颗粒，第一个下拉帧在 0。如果您的片段是以 24 fps 创建但现在处于 30 fps 下拉形式，则这些选项是合适的。如果您的片段是 24P，则不适用。3:2 下拉模式每 5 帧重复一次，因此如果帧 1:00:23 是三个正常帧之后第一个带有场伪影的帧，则应指定 3 作为第一个下拉帧。*3:2 Stutter at 1:
*以 3:2 下拉模式保持颗粒，第一个下拉帧在 1。*3:2 Stutter at 2:
*以 3:2 下拉模式保持颗粒，第一个下拉帧在 2。*3:2 Stutter at 3:
*以 3:2 下拉模式保持颗粒，第一个下拉帧在 3。*3:2 Stutter at 4:
*以 3:2 下拉模式保持颗粒，第一个下拉帧在 4。

Grain Seed:
*Default:
*0.123,
*Range:
*0 or greater.
初始化颗粒生成的随机数生成器。实际种子值并不重要，但不同的种子会产生不同的颗粒图案，相同的值应产生可重复的图案。

### Vignette Parameters:

Vignette Darkness:
*Default:
*0,
*Range:
*0 to 1.暗角是图像朝向角落和边缘变暗的效果。此参数控制屏幕外部角落应变暗（暗角化）多少。0 表示无暗角，1 表示最大变暗。

Vignette Radius:
*Default:
*1,
*Range:
*0 or greater.从中心到应用暗角的距离。

Vignette Edge Softness:
*Default:
*0.5,
*Range:
*0 or greater.暗角柔和边缘的宽度。较大的值产生更柔和、更不明显的边缘。

Vignette Rel Height:
*Default:
*0.75,
*Range:
*0.1 or greater.
控制暗角椭圆的纵横比。通常应设置为图像的纵横比，例如 NTSC 为 .75。

### Field Parameters:

Fields:
*Popup menu, Default: As Is
*.允许从输入片段中去除场伪影。如果您希望片段看起来像是以帧而非场拍摄的，这很有用。您可以显示单个场、合并两个场，或模拟 3:2 下拉抖动模式。
*As Is:
*保持场不变。*Keep Lower Only:
*仅显示下场，去除上场。*Keep Upper Only:
*仅显示上场，去除下场。*Merge Fields:
*混合两个场以去除隔行扫描伪影。*3:2 Stutter at 0:
*模拟时间抖动效果，就好像片段已从 24P 通过 3:2 下拉转换为 NTSC 视频，第一个下拉帧在 0。如果您将此选项与非零 Grain Blur 一起使用，可能还需要将 Grain Hold 设置为相应的值。*3:2 Stutter at 1:
*模拟 3:2 下拉效果，第一个下拉帧在 1。*3:2 Stutter at 2:
*模拟 3:2 下拉效果，第一个下拉帧在 2。*3:2 Stutter at 3:
*模拟 3:2 下拉效果，第一个下拉帧在 3。*3:2 Stutter at 4:
*模拟 3:2 下拉效果，第一个下拉帧在 4。

Field Dominance:
*Popup menu, Default: Lower First
*.指定在模拟 3:2 下拉模式时哪个场应在时间上排在前面。仅当在 Fields 和/或 Grain Hold 选项中选择了 3:2 抖动选项时才使用。
*Lower First:
*下场在时间上排在前面。*Upper First:
*上场在时间上排在前面。

Mask Use:
*Popup menu, Default: Luma
*.确定如何使用 Mask 输入通道来创建单色蒙版。
*Luma:
*使用 RGB 通道的亮度。*Alpha:
*仅使用 Alpha 通道。

Blur Mask:
*Default:
*0.05,
*Range:
*0 or greater.在使用前按此数量模糊 Matte 输入。可以在蒙版区域和非蒙版区域之间提供更平滑的过渡。除非提供了 Matte 输入，否则无效。

Invert Mask:
*Check-box, Default:
*off.
如果开启，反转 Matte 输入，使效果应用于 Matte 为黑色而非白色的区域。除非提供了 Matte 输入，否则无效。
