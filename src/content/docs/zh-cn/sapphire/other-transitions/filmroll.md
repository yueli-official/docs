---
title: FilmRoll
---

## S_FilmRoll

通过将一个片段垂直滚出屏幕同时将另一个片段滚入来实现两个片段之间的转场，同时应用各种胶片损伤效果，如抖动、污渍、划痕和闪烁。

在 Sapphire Transitions 效果子菜单中。

![FilmRoll](../_static/FilmRoll.jpg)


### Inputs:

- **Foreground**: 当前图层。以此片段开始转场。

- **Background**: 默认为无。以此片段结束转场。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Transition Dir** (Popup menu, Default: Wipe Off to Bg)
  选择转场方向。
  - **Wipe Off to Bg**: 从当前图层转场到背景。
  - **Wipe On from Bg**: 从背景转场到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  启用后，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过对 Film Percent 参数进行动画设置来手动执行转场。

- **Amount** (Default: 0, Range: 0 to 1)
  必须禁用 Auto Trans 才能使用此参数。它确定 From 和 To 输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制擦除的时序。

- **Slow In** (Default: 0.5, Range: 0 or greater)
  如果为正值，使转场开始更加平缓。

- **Slow Out** (Default: 0.5, Range: 0 or greater)
  如果为正值，使转场结束更加平缓。

- **Roll Speed** (Integer, Default: 1, Range: any)
  垂直滚动的量，以屏幕高度为单位。片段将在整个转场过程中移动此距离，最终以 To 片段的正常位置结束。

- **Motion Blur** (Default: 0.5, Range: 0 or greater)
  按抖动量的比例模糊结果。

- **Border Height** (Default: 0.1, Range: 0 or greater)
  滚动时 From 和 To 片段之间出现的边框高度。

- **Glow Brightness** (Default: 0.5, Range: 0 or greater)
  调整辉光的峰值亮度。辉光会在效果的开始和结束时自动衰减到零，以提供平滑的转场。

- **Glow Width** (Default: 0.224, Range: 0 or greater)
  发光边框的宽度。

- **Damage Amount** (Default: 2, Range: 0 or greater)
  调整所有损伤效果的峰值量。增大以获得更多损伤效果，减小以获得更干净的效果。损伤会在转场的开始和结束时自动衰减。各个损伤类型也可以通过其各自的参数进行调整，如 Stain Density、Hairs、Scratches 等。


### Stains Parameters:

Stain Density:
*Default:
*2,
*Range:
*0 to 500.每帧上的污渍数量。小数值被视为单个污渍出现在任意给定帧上的概率。

Vary Stain Density:
*Default:
*0.2,
*Range:
*0 or greater.逐帧变化污渍密度的量。

Stain Print:
*Default:
*1,
*Range:
*0 to 1.正片上污渍的相对密度。

Stain Negative:
*Default:
*0,
*Range:
*0 to 1.底片上污渍的相对密度。

Stain Size:
*Default:
*1,
*Range:
*0 or greater.缩放污渍的宽度和高度。

Vary Stain Size:
*Default:
*0.5,
*Range:
*0 or greater.每个污渍之间大小的变化量。

Stain Opacity:
*Default:
*0.5,
*Range:
*0 to 1.缩放污渍的不透明度。

Vary Stain Opacity:
*Default:
*0.5,
*Range:
*0 or greater.每个污渍之间不透明度的变化量。

Vary Stain Brightness:
*Default:
*0,
*Range:
*0 or greater.每个污渍之间亮度的变化量。

Vary Stain Color:
*Default:
*0,
*Range:
*0 or greater.每个污渍的额外随机颜色变化量。如果此参数大于零，污渍颜色可以超出 color1 和 color2 定义的范围。

Stain Color1:
*Default rgb:
*[0 0 0].污渍颜色范围的起始值。

Stain Color2:
*Default rgb:
*[0.25 0.125 0].
污渍颜色范围的终止值。每个污渍将在 color1 和 color2 之间随机取色。

### Dust Parameters:

Dust Density:
*Default:
*60,
*Range:
*0 or greater.每帧上灰尘颗粒的平均数量。小数值被视为单个灰尘斑点出现在任意给定帧上的概率。

Vary Dust Density:
*Default:
*0.2,
*Range:
*0 or greater.逐帧变化灰尘密度的量。

Dust On Print:
*Default:
*1,
*Range:
*0 to 1.正片上灰尘的相对密度。

Dust On Negative:
*Default:
*0,
*Range:
*0 to 1.底片上灰尘的相对密度。

Dust Size:
*Default:
*1,
*Range:
*0 or greater.缩放灰尘的宽度和高度。

Vary Dust Size:
*Default:
*0.5,
*Range:
*0 or greater.每个灰尘颗粒之间大小的变化量。

Dust Opacity:
*Default:
*0.8,
*Range:
*0 to 1.缩放灰尘的不透明度。

Vary Dust Opacity:
*Default:
*0.5,
*Range:
*0 or greater.每个灰尘颗粒之间不透明度的变化量。

Vary Dust Brightness:
*Default:
*0,
*Range:
*0 or greater.每个灰尘颗粒之间亮度的变化量。

Vary Dust Color:
*Default:
*0,
*Range:
*0 or greater.每个灰尘颗粒的额外随机颜色变化量。如果此参数大于零，灰尘颜色可以超出 color1 和 color2 定义的范围。

Dust Color1:
*Default rgb:
*[0 0 0].灰尘颜色范围的起始值。

Dust Color2:
*Default rgb:
*[0 0 0].
灰尘颜色范围的终止值。每个灰尘颗粒将在 color1 和 color2 之间随机取色。

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
*0.1 or greater.控制毛发持续存在的时间长度以及新毛发出现的频率。增大此值可使毛发存在更长时间，减小此值可更频繁地出现新毛发。

Hair Wiggle Amp:
*Default:
*0.1,
*Range:
*0 or greater.控制每根毛发的随机运动和拉伸量。

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
*0 or greater.每根毛发之间大小的变化量。

Hair Color:
*Default rgb:
*[0 0 0].
毛发的颜色。

### Scratches Parameters:

Scratches:
*Integer, Default:
*5,
*Range:
*0 or greater.控制每帧上划痕的平均数量。

Black Scratches:
*Default:
*1,
*Range:
*0 to 1.黑色划痕的数量，相对于 Scratches 参数值。

White Scratches:
*Default:
*0.1,
*Range:
*0 to 1.白色划痕的数量，相对于 Scratches 参数值。

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
*0 to 1.如果为 0，所有划痕将具有相同的宽度。增大以使每条划痕拥有各自的宽度。

Scratches Taper:
*Default:
*0.1,
*Range:
*0 to 1.控制每条划痕末端的尖锐程度。较大的值使每端的锥度更长。

Scratch Opacity:
*Default:
*1,
*Range:
*0 to 1.划痕的最大不透明度。将其设置为 0 将使划痕淡出。

Scratch Roughness:
*Default:
*1,
*Range:
*0 or greater.使每条划痕边缘粗糙化的量，以模拟真实划痕的随机特性。

Scratch Roughness Freq:
*Default:
*150,
*Range:
*0.01 or greater.设置划痕边缘粗糙度的频率。

Gaps:
*Default:
*0.28,
*Range:
*0 to 1.与真实模拟划痕类似，产生划痕的灰尘颗粒有时会滚动，划痕会"跳过"。此参数控制发生这种情况的程度。

Gaps Freq:
*Default:
*120,
*Range:
*0 or greater.划痕间隙出现的频率。

Scratch Area Center:
*Default:
*0,
*Range:
*-2 or greater.划痕覆盖的屏幕区域的中心坐标。0 在屏幕中央，-1 在左边缘，1 在右边缘。

Scratch Area Width:
*Default:
*1,
*Range:
*0 or greater.划痕覆盖的屏幕区域宽度。1 表示划痕覆盖整个屏幕区域。要使划痕仅出现在一个条带中，请将划痕区域宽度调小。

Weave Amount:
*Default:
*1,
*Range:
*0 or greater.每条划痕在屏幕上平均摆动的幅度。以画面宽度为单位，因此 1.0 将允许划痕在整个屏幕上游走。如果设置为零，所有划痕将是直线垂直的。

Weave Frequency:
*Default:
*0.1,
*Range:
*0.01 or greater.
划痕在屏幕上摆动的速度，以每帧的周期数为单位。通常小于 1。

### Shake Parameters:

Shake Amplitude:
*Default:
*0.2,
*Range:
*0 or greater.要添加的垂直抖动量。

Shake Frequency:
*Default:
*1,
*Range:
*0 or greater.缩放抖动的频率。增大以获得更快的抖动和更频繁的跳跃及方向变化。

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

Shake Time Offset:
*Default:
*0,
*Range:
*any.
在时间上偏移抖动模式。调整此值以控制抖动发生的确切时间。

### Vignette Parameters:

Vignette Darkness:
*Default:
*0.5,
*Range:
*0 to 1.暗角效果是图像向角落和边缘方向变暗的效果。此参数控制屏幕外围角落应该变暗（暗角化）的程度。0 表示无暗角效果，1 表示最大暗角效果。

Vignette Radius:
*Default:
*1,
*Range:
*0 or greater.从中心到应用暗角效果的距离。

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
控制暗角椭圆的宽高比。通常应设置为图像的宽高比，例如 NTSC 为 0.75。

### Flicker Parameters:

Flicker:
*Default:
*1,
*Range:
*0 or greater.随时间以不同程度缩放源片段的颜色，产生闪烁效果。闪烁模式可以是随机的、周期性波形的，或两者的组合。

Flicker Rand Amp:
*Default:
*1,
*Range:
*0 or greater.随机亮度闪烁的振幅。

Flicker Rand Freq:
*Default:
*10,
*Range:
*0 or greater.随机闪烁的频率。增大以获得帧与帧之间更多的变化。减小以获得更慢的闪烁。

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
*0.5,
*Range:
*0 or greater.随时间以不同程度模糊源片段，以模拟放映机的对焦问题。散焦模式可以是随机的、周期性波形的，或两者的组合。

Defocus Rand Amp:
*Default:
*1,
*Range:
*0 or greater.随时间随机变化的散焦振幅。

Defocus Rand Freq:
*Default:
*10,
*Range:
*0 or greater.缩放随机散焦的频率。增大以获得帧与帧之间更多的变化。减小以获得更慢的散焦变化。

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
波形散焦的频率。增大以获得帧与帧之间更多的变化。

### Other Parameters:

Seed:
*Default:
*0.123,
*Range:
*0 or greater.用于初始化随机数生成器。实际种子值本身不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

Flip Stamps Vertically:
*Check-box, Default:
*off.
垂直翻转所有印记（毛发、划痕、灰尘等）。
