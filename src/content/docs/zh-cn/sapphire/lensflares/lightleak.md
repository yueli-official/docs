---
title: LightLeak
---

## S_LightLeak

渲染抽象的色彩图案，模拟光线通过相机机身缝隙泄漏的效果。漏光由三个独立的元素组成，可以分别调整。

在 Sapphire Lighting 效果子菜单中。

![LightLeak](../_static/LightLeak.jpg)


### Inputs:

- **Background**: 当前图层。用作背景的素材。

- **Mask**: 默认为无。在结果和源输入之间插值。白色区域使用效果结果。黑色区域使用源素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此数值模糊 Mocha 遮罩。这可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 遮罩的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  如果启用，Mocha 遮罩的黑白将在应用效果前反转。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素值扩展或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下进行 Dilate Mocha，以便快速调整。
  - **High**: 在高质量模式下进行 Dilate Mocha，以获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用到整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩都提供给效果时，决定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Scale Lights** (Default: 1, Range: 0 or greater)
  按此值缩放漏光。增大可获得更亮的结果。

- **Offset Darks** (Default: 0, Range: -8 to 2)
  将此灰度值添加到结果的较暗区域。可以为负值以增加对比度。

- **Color** (Default rgb: [1 1 1])
  漏光的整体颜色。

- **Hue Shift** (Default: 0, Range: any)
  偏移漏光的色相，从红到绿到蓝再到红循环。

- **Saturation** (Default: 1, Range: -2 to 8)
  缩放漏光的色彩饱和度。增加以获得更强烈的颜色。设为 0 可获得单色漏光。

- **Gamma** (Default: 1, Range: 0.1 or greater)
  增加 Gamma 会使漏光变亮，尤其会增强较暗的区域。

- **Speed** (Default: 1, Range: 0 or greater)
  缩放所有元素的速度。

- **Shift** (X & Y, Default: [0 0], Range: any)
  偏移所有元素的位置。

- **Flicker Amp** (Default: 0.2, Range: 0 or greater)
  漏光亮度中随机闪烁的量。

- **Flicker Freq** (Default: 4, Range: 0 or greater)
  随机闪烁的频率。增加以获得帧间更多的变化。减少以获得更慢的闪烁。

- **Random Motion** (Default: 0, Range: 0 or greater)
  每个元素的随机运动量。

- **Random Frequency** (Default: 4, Range: 0 or greater)
  随机运动的频率。增加以获得更快、更狂乱的运动。减少以获得更慢、更平滑的运动。

- **Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化随机数生成器。实际种子值并不重要，但不同的种子会产生不同的结果，相同的值应该产生可重复的结果。

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  在与漏光合成之前缩放背景的亮度。如果为 0，结果将仅包含黑色背景上的漏光图像。

- **Combine** (Popup menu, Default: Screen)
  决定漏光如何与背景图像合成。
  - **Screen**: 漏光使用有助于防止过亮结果的函数与背景混合。
  - **Add**: 漏光被添加到背景上。
  - **Leaks Only**: 仅显示漏光，没有背景。

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  如果此值为正，输出 Alpha 通道将包含来自漏光的一些不透明度。漏光红、绿、蓝亮度的最大值按此值缩放，并在每个像素处与背景 Alpha 合成。

- **Glow Brightness** (Default: 1, Range: 0 or greater)
  缩放在漏光与背景合成后应用于整个图像的辉光亮度。

- **Glow Width** (Default: 0.4, Range: 0 or greater)
  辉光的宽度。增大可获得更柔和的辉光，减小可获得更锐利、更亮的辉光。

- **Glow Threshold** (Default: 0.8, Range: 0 or greater)
  图像中亮度高于此值的部分会被添加辉光。


### Element1 Parameters:

Element1 Enable:
*Check-box, Default:
*on.打开或关闭此元素。

Size1:
*Default:
*2,
*Range:
*0 or greater.调整此元素的大小。此参数可以使用 Size1 控件进行调整。

Rel Height1:
*Default:
*2,
*Range:
*0 or greater.缩放此元素的垂直尺寸，使其变为椭圆形而非圆形。

Brightness1:
*Default:
*0.2,
*Range:
*0 or greater.缩放此元素的亮度。

Speed1:
*Default:
*1,
*Range:
*0 or greater.此元素在屏幕上移动的速度。设为零可使元素固定在其中心位置。速度越快，元素的起始和结束位置距其中心越远。

Angle1:
*Default:
*175,
*Range:
*any.此元素在屏幕上的运动路径角度。元素沿此角度的直线移动，在素材中点时经过中心位置。此参数可以使用 Angle1 控件进行调整。

Center1:
*X & Y, Default:
*[0 0],
*Range:
*any.此元素运动的中心点。元素将在素材播放到一半时到达此位置。此参数可以使用 Center1 控件进行调整。

Outer Color1:
*Default rgb:
*[1 1 0.15].此元素外边缘的颜色。

Mid Color1:
*Default rgb:
*[1 0.7 0.3].此元素中间点的颜色，位于外部颜色和中心颜色之间。确切位置取决于 Midpoint 参数。

Center Color1:
*Default rgb:
*[1 0.1 0].此元素中心的颜色。

Midpoint1:
*Default:
*0.5,
*Range:
*0 to 1.在元素的中心和外边缘之间移动中间颜色的位置。设为 0 将中间颜色放在中心，设为 1 将其放在边缘。

Softness1:
*Default:
*0.4,
*Range:
*0 or greater.模糊此元素的颜色渐变。增大可获得更平滑的渐变，减小可获得更锐利的色带。

Noise Amp1:
*Default:
*2,
*Range:
*0 or greater.应用于此元素的噪声量。

Noise Freq1:
*Default:
*1.5,
*Range:
*0 or greater.应用于此元素的噪声频率。增大可获得更小的斑点，减小可获得更大的斑点。

Noise Freq Rel Y1:
*Default:
*1,
*Range:
*0 or greater.噪声图案的相对垂直频率。增大可使噪声扁平化，减小可使其垂直拉伸。

Noise Detail1:
*Default:
*0,
*Range:
*0 to 1.控制噪声模拟中的精细细节量。减小可获得更平滑的噪声，增大可获得更粗糙或颗粒感的外观。

Noise Boil Speed1:
*Default:
*1,
*Range:
*0 or greater.
设置元素移动时噪声沸腾或演变的速度。设为 0 可获得静态噪声图案。

### Element2 Parameters:

Element2 Enable:
*Check-box, Default:
*off.打开或关闭此元素。

Size2:
*Default:
*2,
*Range:
*0 or greater.调整此元素的大小。此参数可以使用 Size2 控件进行调整。

Rel Height2:
*Default:
*2,
*Range:
*0 or greater.缩放此元素的垂直尺寸，使其变为椭圆形而非圆形。

Brightness2:
*Default:
*1,
*Range:
*0 or greater.缩放此元素的亮度。

Speed2:
*Default:
*1,
*Range:
*0 or greater.此元素在屏幕上移动的速度。设为零可使元素固定在其中心位置。速度越快，元素的起始和结束位置距其中心越远。

Angle2:
*Default:
*175,
*Range:
*any.此元素在屏幕上的运动路径角度。元素沿此角度的直线移动，在素材中点时经过中心位置。此参数可以使用 Angle2 控件进行调整。

Center2:
*X & Y, Default:
*[-0.5 -0.5],
*Range:
*any.此元素运动的中心点。元素将在素材播放到一半时到达此位置。此参数可以使用 Center2 控件进行调整。

Outer Color2:
*Default rgb:
*[0.25 0.15 0.1].此元素外边缘的颜色。

Mid Color2:
*Default rgb:
*[0.7 0.12 0.22].此元素中间点的颜色，位于外部颜色和中心颜色之间。确切位置取决于 Midpoint 参数。

Center Color2:
*Default rgb:
*[1 1 1].此元素中心的颜色。

Midpoint2:
*Default:
*0.5,
*Range:
*0 to 1.在元素的中心和外边缘之间移动中间颜色的位置。设为 0 将中间颜色放在中心，设为 1 将其放在边缘。

Softness2:
*Default:
*0.4,
*Range:
*0 or greater.模糊此元素的颜色渐变。增大可获得更平滑的渐变，减小可获得更锐利的色带。

Noise Amp2:
*Default:
*1,
*Range:
*0 or greater.应用于此元素的噪声量。

Noise Freq2:
*Default:
*3,
*Range:
*0 or greater.应用于此元素的噪声频率。增大可获得更小的斑点，减小可获得更大的斑点。

Noise Freq Rel Y2:
*Default:
*1,
*Range:
*0 or greater.噪声图案的相对垂直频率。增大可使噪声扁平化，减小可使其垂直拉伸。

Noise Detail2:
*Default:
*0,
*Range:
*0 to 1.控制噪声模拟中的精细细节量。减小可获得更平滑的噪声，增大可获得更粗糙或颗粒感的外观。

Noise Boil Speed2:
*Default:
*0,
*Range:
*0 or greater.
设置元素移动时噪声沸腾或演变的速度。设为 0 可获得静态噪声图案。

### Element3 Parameters:

Element3 Enable:
*Check-box, Default:
*off.打开或关闭此元素。

Size3:
*Default:
*0.25,
*Range:
*0 or greater.调整此元素的大小。此参数可以使用 Size3 控件进行调整。

Rel Height3:
*Default:
*1,
*Range:
*0 or greater.缩放此元素的垂直尺寸，使其变为椭圆形而非圆形。

Brightness3:
*Default:
*1,
*Range:
*0 or greater.缩放此元素的亮度。

Speed3:
*Default:
*1,
*Range:
*0 or greater.此元素在屏幕上移动的速度。设为零可使元素固定在其中心位置。速度越快，元素的起始和结束位置距其中心越远。

Angle3:
*Default:
*175,
*Range:
*any.此元素在屏幕上的运动路径角度。元素沿此角度的直线移动，在素材中点时经过中心位置。此参数可以使用 Angle3 控件进行调整。

Center3:
*X & Y, Default:
*[0.5 0.5],
*Range:
*any.此元素运动的中心点。元素将在素材播放到一半时到达此位置。此参数可以使用 Center3 控件进行调整。

Outer Color3:
*Default rgb:
*[0.2 0.2 0].此元素外边缘的颜色。

Mid Color3:
*Default rgb:
*[0.55 0.4 0].此元素中间点的颜色，位于外部颜色和中心颜色之间。确切位置取决于 Midpoint 参数。

Center Color3:
*Default rgb:
*[1 0.7 0].此元素中心的颜色。

Midpoint3:
*Default:
*0.5,
*Range:
*0 to 1.在元素的中心和外边缘之间移动中间颜色的位置。设为 0 将中间颜色放在中心，设为 1 将其放在边缘。

Softness3:
*Default:
*0.4,
*Range:
*0 or greater.模糊此元素的颜色渐变。增大可获得更平滑的渐变，减小可获得更锐利的色带。

Noise Amp3:
*Default:
*0.5,
*Range:
*0 or greater.应用于此元素的噪声量。

Noise Freq3:
*Default:
*1,
*Range:
*0 or greater.应用于此元素的噪声频率。增大可获得更小的斑点，减小可获得更大的斑点。

Noise Freq Rel Y3:
*Default:
*1,
*Range:
*0 or greater.噪声图案的相对垂直频率。增大可使噪声扁平化，减小可使其垂直拉伸。

Noise Detail3:
*Default:
*0,
*Range:
*0 to 1.控制噪声模拟中的精细细节量。减小可获得更平滑的噪声，增大可获得更粗糙或颗粒感的外观。

Noise Boil Speed3:
*Default:
*0,
*Range:
*0 or greater.设置元素移动时噪声沸腾或演变的速度。设为 0 可获得静态噪声图案。

Opacity:
*Popup menu, Default: Normal
*.决定处理不透明度/透明度的方法。
*All Opaque:
*当输入图像完全不透明且没有透明度 (alpha=1) 时，使用此选项可稍微加快渲染速度。*Normal:
*正常处理不透明度。*As Premult:
*按图像已为预乘形式处理（颜色已按不透明度缩放）。此选项的渲染速度也略快于 Normal 模式，但结果也将为预乘形式，有时不太正确。

Mask Use:
*Popup menu, Default: Luma
*.决定如何使用遮罩输入通道来创建单色遮罩。
*Luma:
*使用 RGB 通道的亮度。*Alpha:
*仅使用 Alpha 通道。

Blur Mask:
*Default:
*0.05,
*Range:
*0 or greater.在使用前按此数值模糊遮罩输入。这可以提供遮罩区域和非遮罩区域之间更平滑的过渡。除非提供了遮罩输入，否则无效。

Invert Mask:
*Check-box, Default:
*off.
如果开启，反转遮罩输入，使效果应用于遮罩为黑色的区域而非白色区域。除非提供了遮罩输入，否则无效。