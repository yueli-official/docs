---
title: Shake
---

## S_Shake

随时间对源素材应用平移、缩放和/或旋转的抖动运动。抖动是随机但可重复的，因此使用相同的参数每次都会生成相同的抖动运动。打开 Motion Blur 并调整 Mo Blur Length 以获得不同程度的运动模糊。调整 Amplitude 和 Frequency 以获得不同的抖动速度和程度。Rand 参数提供对随机非周期抖动的详细控制，Wave 参数调整规则周期抖动。X、Y、Z 和 Tilt 参数分别控制水平、垂直、缩放和旋转的抖动量。

在 Sapphire Distort 效果子菜单中。

![Shake](../_static/Shake.jpg)


### Inputs:

- **Source**: 当前图层。要进行抖动的素材。

- **Mask**: 默认为无。在结果和源输入之间进行插值。白色区域使用效果的结果。黑色区域使用源素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Style** (Popup menu, Default: Normal)
  控制抖动类型。
  - **Normal**: 稳定的摄像机抖动。
  - **Twitchy**: 静止期间被突然的快速抖动打断。
  - **Jumpy**: 从一个位置突然跳到另一个位置，中间有较慢的漂移。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用之前按此量模糊 Mocha 遮罩。这可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

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
  在使用之前按此像素量膨胀或侵蚀 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在 High 质量模式下膨胀 Mocha 遮罩，以获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩都提供给效果时，确定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Apply Mask** (Popup menu, Default: Post-effect)
  控制遮罩在效果中的应用位置 - 这会影响输入遮罩和 Mocha 遮罩。
  - **Post-effect**: 在所有效果渲染完成后应用遮罩。
  - **Pre-effect**: 在处理效果之前将遮罩应用于源素材。

- **Amplitude** (Default: 1, Range: 0 or greater)
  缩放抖动运动的幅度。

- **Frequency** (Default: 8, Range: 0 or greater)
  增大以加快抖动，减小以减慢抖动。（如果对频率值设置动画，请注意产生的抖动频率也受值变化率的影响。）

- **Phase** (Default: 0, Range: any)
  抖动运动的时间偏移。（如果对此值设置动画，其变化率也会影响表观频率。）

- **Stillness** (Default: 0.7, Range: 0 to 1)
  在 Twitchy 模式下，调整图像静止的时间比例。增大以获得更频繁的抖动。

- **Twitch Frequency** (Default: 2, Range: 0 or greater)
  在 Twitchy 模式下，控制运动和静止期间的长度。增大以获得更短、更频繁的运动突发。

- **Drift** (Default: 0.3, Range: 0 to 1)
  在 Jumpy 模式下，控制跳跃之间的运动速度。

- **Center Bias** (Default: 0, Range: 0 or greater)
  在 Jumpy 模式下，调整每次跳跃将图像重置到原始位置的可能性。如果设置为零，每次跳跃都是随机的。如果设置为一，每次跳跃都会回到中心。

- **Z Dist** (Default: 1, Range: 0.001 or greater)
  缩放图像的"距离"。大于 1.0 的值使其更远更小。小于 1.0 的值使图像更近更大。稍微放大有时可用于隐藏边缘伪影。

- **Motion Blur** (Check-box, Default: off)
  抖动运动的运动模糊选项。

- **Mo Blur Length** (Default: 1, Range: 0 or greater)
  缩放运动模糊的量。在处理场时使用约 0.5，处理帧时使用 1.0 以获得逼真的运动模糊。如果 Motion Blur 为 No，此参数无效。

- **Seed** (Default: 0, Range: 0 or greater)
  用于初始化随机数生成器。实际种子值并不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Wrap** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  确定访问源图像边界外区域的方法。
  - **No**: 在边界外显示黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复图像的镜像副本。使用此方法时边缘通常不太明显。


### X Shake Parameters:

X Rand Amp:
*Default:
*0.2,
*Range:
*0 or greater.水平随机抖动的幅度。

X Rand Freq:
*Default:
*1,
*Range:
*0 or greater.水平随机抖动的频率。

X Wave Amp:
*Default:
*0,
*Range:
*0 or greater.水平规则波动抖动的幅度。

X Wave Freq:
*Default:
*0.5,
*Range:
*0 or greater.水平规则波动抖动的频率，以每秒周期为单位。

X Phase:
*Default:
*0,
*Range:
*any.
水平抖动的时间偏移。

### Y Shake Parameters:

Y Rand Amp:
*Default:
*0.1,
*Range:
*0 or greater.垂直随机抖动的幅度。

Y Rand Freq:
*Default:
*1,
*Range:
*0 or greater.垂直随机抖动的频率。

Y Wave Amp:
*Default:
*0,
*Range:
*0 or greater.垂直规则波动抖动的幅度。

Y Wave Freq:
*Default:
*0.5,
*Range:
*0 or greater.垂直规则波动抖动的频率，以每秒周期为单位。

Y Phase:
*Default:
*0,
*Range:
*any.
垂直抖动的时间偏移。

### Z Shake Parameters:

Z Rand Amp:
*Default:
*0,
*Range:
*0 or greater.缩放随机抖动的幅度。

Z Rand Freq:
*Default:
*1,
*Range:
*0 or greater.缩放随机抖动的频率。

Z Wave Amp:
*Default:
*0,
*Range:
*0 or greater.缩放规则波动抖动的幅度。

Z Wave Freq:
*Default:
*0.5,
*Range:
*0 or greater.缩放规则波动抖动的频率，以每秒周期为单位。

Z Phase:
*Default:
*0,
*Range:
*any.
缩放抖动的时间偏移。

### Tilt Shake Parameters:

Tilt Rand Amp:
*Default:
*0,
*Range:
*0 or greater.旋转随机抖动的幅度（以度为单位）。

Tilt Rand Freq:
*Default:
*1,
*Range:
*0 or greater.旋转随机抖动的频率。

Tilt Wave Amp:
*Default:
*0,
*Range:
*0 or greater.旋转规则波动抖动的幅度（以度为单位）。

Tilt Wave Freq:
*Default:
*0.5,
*Range:
*0 or greater.旋转规则波动抖动的频率，以每秒周期为单位。

Tilt Phase:
*Default:
*0,
*Range:
*any.
旋转抖动的时间偏移。

### Channels Parameters:

Red Amplitude:
*Default:
*1,
*Range:
*0 or greater.红色通道中抖动的相对量。将此值从默认值更改将导致红色通道比其他颜色通道移动更多或更少，产生色彩边纹或通道分离效果。

Green Amplitude:
*Default:
*1,
*Range:
*0 or greater.绿色通道中抖动的相对量。

Blue Amplitude:
*Default:
*1,
*Range:
*0 or greater.蓝色通道中抖动的相对量。

Red Phase:
*Default:
*0,
*Range:
*any.红色通道的相对相位。正值将使红色通道在时间上领先于其他通道，使其首先移动，其他通道跟随。负值效果相反，使红色通道落后于其他通道。较小的值通常产生最佳效果。

Green Phase:
*Default:
*0,
*Range:
*any.绿色通道的相对相位。

Blue Phase:
*Default:
*0,
*Range:
*any.蓝色通道的相对相位。

RGB Randomness:
*Default:
*0,
*Range:
*0 or greater.每个颜色通道中随机运动的量。增大此参数以使所有三个颜色通道在不同路径上独立于整体抖动进行随机移动。此运动由 X Rand Amp、Y Rand Amp、Z Rand Amp 和 Tilt Rand Amp 缩放。

RGB Frequency:
*Default:
*2,
*Range:
*0 or greater.
随机颜色通道抖动的频率。

### Other Parameters:

Opacity:
*Popup menu, Default: Normal
*.确定处理不透明度/透明度的方法。
*All Opaque:
*当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可略微加快渲染速度。*Normal:
*正常处理不透明度。*As Premult:
*按预乘形式处理图像（颜色已按不透明度缩放）。此选项也比 Normal 模式渲染略快，但结果也将是预乘形式，有时不太准确。

Mask Use:
*Popup menu, Default: Luma
*.确定如何使用 Mask 输入通道来创建单色遮罩。
*Luma:
*使用 RGB 通道的亮度。*Alpha:
*仅使用 Alpha 通道。

Blur Mask:
*Default:
*0.05,
*Range:
*0 or greater.在使用之前按此量模糊 Matte 输入。这可以在遮罩区域和非遮罩区域之间提供更平滑的过渡。除非提供了 Matte 输入，否则此选项无效。

Invert Mask:
*Check-box, Default:
*off.如果启用，反转 Matte 输入，使效果应用于 Matte 为黑色而非白色的区域。除非提供了 Matte 输入，否则此选项无效。

Crop Input Parameters:
*Default:
*0,
*Range:
*0 or greater.这 4 个参数，
Crop Top
,
Crop Bottom
,
Crop Left,
和
Crop Right
,
允许选择输入图像的矩形子区域进行处理。
如果 Wrap 参数设置为 "No"，则暴露的边框将是透明的。
如果 Wrap 为 "Tile" 或 "Reflect"，源图像将在新的裁剪边框上包裹以填充画面。
这可以更容易地避免因变形具有不良边缘的图像而产生的伪影。
