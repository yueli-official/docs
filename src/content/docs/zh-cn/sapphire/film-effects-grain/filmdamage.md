---
title: FilmDamage
---

## S_FilmDamage

模拟受损胶片，提供多种选项，包括灰尘、毛发、污渍、划痕、散焦、闪烁和抖动。每个选项都有一个主控制和一组详细控制，用于调整该类型损伤的外观。

在 Sapphire Stylize 效果子菜单中。

![FilmDamage](../_static/FilmDamage.jpg)


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


### Grain Parameters:

Grain Amp:
*Default:
*0.1,
*Range:
*0 to 2.缩放添加到结果中的胶片颗粒的振幅。设为 0 可禁用所有颗粒。

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
*保持颗粒图案一帧（2 个场）。*3:2 Pulldown at 0:
*以 3:2 下拉模式保持颗粒，第一个下拉帧在 0。如果您的片段是以 24 fps 创建但现在处于 30 fps 下拉形式，则这些选项是合适的。如果您的片段是 24P，则不适用。3:2 下拉模式每 5 帧重复一次，因此如果帧 1:00:23 是三个正常帧之后第一个带有场伪影的帧，则应指定 3 作为第一个下拉帧。*3:2 Pulldown at 1:
*以 3:2 下拉模式保持颗粒，第一个下拉帧在 1。*3:2 Pulldown at 2:
*以 3:2 下拉模式保持颗粒，第一个下拉帧在 2。*3:2 Pulldown at 3:
*以 3:2 下拉模式保持颗粒，第一个下拉帧在 3。*3:2 Pulldown at 4:
*以 3:2 下拉模式保持颗粒，第一个下拉帧在 4。

### Color Correct Parameters:

Saturation:
*Default:
*1,
*Range:
*any.缩放色彩饱和度。增大以获得更鲜艳的颜色。设为 0 可获得单色效果。

Scale Lights:
*Default:
*1,
*Range:
*0 or greater.按此灰度值缩放结果。增大以获得更亮的结果。

Offset Darks:
*Default:
*0,
*Range:
*any.将此灰度值添加到源的较暗区域。可以为负值以增加对比度。

Tint Lights:
*Default rgb:
*[1 1 1].按此颜色缩放结果，从而为较亮区域着色。

Tint Darks:
*Default rgb:
*[0 0 0].
将此颜色添加到源的较暗区域。

### Stains Parameters:

Stain Density:
*Default:
*0.2,
*Range:
*0 or greater.每帧上的污渍数量。小数值被视为单个污渍出现在任何给定帧上的概率。

Vary Stain Density:
*Default:
*0.2,
*Range:
*0 or greater.帧与帧之间污渍密度的变化量。

Stain Print:
*Default:
*1,
*Range:
*0 to 1.正片上的污渍相对密度。

Stain Negative:
*Default:
*0,
*Range:
*0 to 1.底片上的污渍相对密度。

Stain Size:
*Default:
*1,
*Range:
*0 or greater.缩放污渍的宽度和高度。

Vary Stain Size:
*Default:
*0.5,
*Range:
*0 or greater.各污渍之间大小的变化量。

Stain Opacity:
*Default:
*0.5,
*Range:
*0 to 1.缩放污渍的不透明度。

Vary Stain Opacity:
*Default:
*0.5,
*Range:
*0 or greater.各污渍之间不透明度的变化量。

Vary Stain Brightness:
*Default:
*0,
*Range:
*0 or greater.各污渍之间亮度的变化量。

Vary Stain Color:
*Default:
*0,
*Range:
*0 or greater.每个污渍的额外随机颜色变化量。如果此参数大于零，污渍颜色可能会超出 color1 和 color2 定义的范围。

Stain Color1:
*Default rgb:
*[0 0 0].污渍颜色范围的起始值。

Stain Color2:
*Default rgb:
*[0.25 0.125 0].
污渍颜色范围的结束值。每个污渍将在 color1 和 color2 之间获得一个随机颜色。

### Dust Parameters:

Dust Density:
*Default:
*30,
*Range:
*0 or greater.每帧上的平均灰尘颗粒数。小数值被视为单个灰尘颗粒出现在任何给定帧上的概率。

Vary Dust Density:
*Default:
*0.2,
*Range:
*0 or greater.帧与帧之间灰尘密度的变化量。

Dust On Print:
*Default:
*1,
*Range:
*0 to 1.正片上的灰尘相对密度。

Dust On Negative:
*Default:
*0,
*Range:
*0 to 1.底片上的灰尘相对密度。

Dust Size:
*Default:
*1,
*Range:
*0 or greater.缩放灰尘的宽度和高度。

Vary Dust Size:
*Default:
*0.5,
*Range:
*0 or greater.各灰尘颗粒之间大小的变化量。

Dust Opacity:
*Default:
*0.8,
*Range:
*0 to 1.缩放灰尘的不透明度。

Vary Dust Opacity:
*Default:
*0.5,
*Range:
*0 or greater.各灰尘颗粒之间不透明度的变化量。

Vary Dust Brightness:
*Default:
*0,
*Range:
*0 or greater.各灰尘颗粒之间亮度的变化量。

Vary Dust Color:
*Default:
*0,
*Range:
*0 or greater.每个灰尘颗粒的额外随机颜色变化量。如果此参数大于零，灰尘颜色可能会超出 color1 和 color2 定义的范围。

Dust Color1:
*Default rgb:
*[0 0 0].灰尘颜色范围的起始值。

Dust Color2:
*Default rgb:
*[0 0 0].
灰尘颜色范围的结束值。每个灰尘颗粒将在 color1 和 color2 之间获得一个随机颜色。

### Hairs Parameters:

Hairs:
*Default:
*2,
*Range:
*0 or greater.卡在放映机片门中的毛发数量。

Hair Persistence:
*Default:
*3,
*Range:
*0.1 or greater.控制毛发持续存在的时间长度以及新毛发出现的频率。增大此值可获得长时间存在的毛发，减小此值可更频繁地出现新毛发。

Hair Wiggle Amp:
*Default:
*0.1,
*Range:
*0 or greater.控制每根毛发表现出的随机运动和拉伸量。

Hair Wiggle Freq:
*Default:
*1,
*Range:
*0 or greater.控制毛发摆动的频率。

Hair Opacity:
*Default:
*1,
*Range:
*0 to 1.缩放毛发的不透明度。

Hair Size:
*Default:
*1,
*Range:
*0 or greater.缩放毛发的宽度和高度。

Vary Hair Size:
*Default:
*1,
*Range:
*0 or greater.各毛发之间大小的变化量。

Hair Color:
*Default rgb:
*[0 0 0].
毛发的颜色。

### Scratches Parameters:

Scratches:
*Integer, Default:
*5,
*Range:
*0 or greater.控制每帧上的平均划痕数量。

Black Scratches:
*Default:
*1,
*Range:
*0 to 1.相对于 Scratches 参数值的黑色划痕数量。

White Scratches:
*Default:
*0.1,
*Range:
*0 to 1.相对于 Scratches 参数值的白色划痕数量。

Black Scratch Length:
*Default:
*10,
*Range:
*0 or greater.黑色划痕的平均长度（以帧为单位）。

White Scratch Length:
*Default:
*2,
*Range:
*0 or greater.白色划痕的平均长度（以帧为单位）。

Scratch Width:
*Default:
*0.15,
*Range:
*0 or greater.平均划痕的宽度，以近似 NTSC 大小的像素为单位。

Vary Scratches Width:
*Default:
*1,
*Range:
*0 to 1.如果为 0，所有划痕将具有相同的宽度。增大以让每条划痕有自己的宽度。

Scratches Taper:
*Default:
*0.1,
*Range:
*0 to 1.控制每条划痕末端的尖锐程度。较大的值在每个末端产生更长的锥形。

Scratch Opacity:
*Default:
*1,
*Range:
*0 to 1.划痕的最大不透明度。设为 0 将使划痕淡出。

Scratch Roughness:
*Default:
*1,
*Range:
*0 or greater.粗糙化每条划痕边缘的量，以模拟真实划痕的随机特征。

Scratch Rough Freq:
*Default:
*150,
*Range:
*0.01 or greater.设置划痕边缘粗糙度的频率。

Gaps:
*Default:
*0.28,
*Range:
*0 to 1.与真实的模拟划痕一样，产生划痕的灰尘颗粒有时会滚动，划痕会"跳过"。此参数控制这种情况发生的程度。

Gaps Freq:
*Default:
*120,
*Range:
*0 or greater.划痕间隙出现的频率。

Scratch Area Center:
*Default:
*0,
*Range:
*any.划痕覆盖的屏幕区域的中心坐标。0 在屏幕中间，-1 在左边缘，1 在右边缘。

Scratch Area Width:
*Default:
*1,
*Range:
*0 or greater.划痕覆盖的屏幕区域宽度。1 表示划痕覆盖整个屏幕区域。要仅在一个条带中获得划痕，请将划痕区域宽度调小。

Weave Amount:
*Default:
*1,
*Range:
*0 or greater.每条划痕在屏幕上平均蜿蜒的程度。这以画面宽度为单位，因此 1.0 将让划痕在整个屏幕上蜿蜒。如果设为零，所有划痕将是垂直直线。

Weave Frequency:
*Default:
*0.1,
*Range:
*0.01 or greater.
划痕在屏幕上蜿蜒的速度，以每帧周期数为单位。通常小于 1。

### Shake Parameters:

Shake Amplitude:
*Default:
*0,
*Range:
*0 or greater.要添加的垂直抖动量。

Shake Frequency:
*Default:
*1,
*Range:
*0 or greater.缩放抖动的频率。增大以获得更快的抖动，具有更频繁的跳跃和方向变化。

Shake Jumpiness:
*Default:
*1,
*Range:
*0 or greater.大尺度跳跃式抖动的量。

Shake Random:
*Default:
*0.1,
*Range:
*0 or greater.小尺度随机抖动的量。

Shake Always:
*Default:
*0.5,
*Range:
*0 to 1.控制抖动发生的频率。如果设为 1，片段持续抖动。如果设为 0，片段从不抖动。介于两者之间的值会使片段在部分时间抖动，在其他时间保持静止。

Interframe Border Height:
*Default:
*0.1,
*Range:
*0 or greater.帧之间黑色条的大小（胶片的未曝光部分）。

Shake Time Offset:
*Default:
*0,
*Range:
*any.在时间上偏移抖动图案。调整此值以控制抖动发生的确切时间。

Shake Motion Blur:
*Default:
*0.1,
*Range:
*0 or greater.
按与抖动量成比例的程度模糊结果。

### Vignette Parameters:

Vignette Darkness:
*Default:
*0.1,
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

### Flicker Parameters:

Flicker:
*Default:
*0.2,
*Range:
*0 or greater.随时间以不同的量缩放源片段的颜色，产生闪烁效果。闪烁模式可以是随机的、周期性波形的，或两者的组合。

Flicker Rand Amp:
*Default:
*1,
*Range:
*0 or greater.随机亮度闪烁的振幅。

Flicker Rand Freq:
*Default:
*10,
*Range:
*0 or greater.随机闪烁的频率。增大以获得帧间更多变化。减小以获得更慢的闪烁。

Flicker Wave Amp:
*Default:
*0,
*Range:
*0 or greater.周期性波形闪烁的振幅。

Flicker Wave Freq:
*Default:
*5,
*Range:
*0 or greater.
波形闪烁的频率。增大以获得更快的闪烁，减小以获得更慢的闪烁。如果 Wave Amp 为 0，则此参数无效。

### Defocus Parameters:

Defocus:
*Default:
*0,
*Range:
*0 or greater.随时间以不同的量模糊源片段，以模拟放映机中的对焦问题。散焦模式可以是随机的、周期性波形的，或两者的组合。

Defocus Rand Amp:
*Default:
*1,
*Range:
*0 or greater.随时间随机变化的散焦振幅。

Defocus Rand Freq:
*Default:
*10,
*Range:
*0 or greater.缩放随机散焦的频率。增大以获得帧间更多变化。减小以获得更慢的散焦变化。

Defocus Wave Amp:
*Default:
*0,
*Range:
*0 or greater.周期性波形散焦的振幅。

Defocus Wave Freq:
*Default:
*5,
*Range:
*0 or greater.
波形散焦的频率。增大以获得帧间更多变化。

### Other Parameters:

Seed:
*Default:
*0.123,
*Range:
*0 or greater.用于初始化随机数生成器。实际种子值并不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

Opacity:
*Popup menu, Default: Normal
*.确定处理不透明度/透明度的方法。
*All Opaque:
*当输入图像完全不透明且无透明度 (alpha=1) 时，使用此选项可稍微加快渲染速度。*Normal:
*正常处理不透明度。*As Premult:
*按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将为预乘形式，有时可能不太准确。

Flip Stamps Vertically:
*Check-box, Default:
*off.垂直翻转所有图章（毛发、划痕、灰尘等）。

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
