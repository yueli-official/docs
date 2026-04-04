---
title: CartoonPaint
---

## S_CartoonPaint

自动生成具有卡通画笔外观的源素材版本。查找图像中的边缘并为这些边缘绘制新的轮廓线。用画笔笔触形状替换边缘之间区域的颜色。

位于 Sapphire Stylize 效果子菜单中。

![CartoonPaint](../_static/CartoonPaint.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。在结果和源输入之间进行插值。白色区域使用效果结果。黑色区域使用源素材。


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
  如果启用，在应用效果之前会反转 Mocha 遮罩的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素量扩展或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下扩展 Mocha 遮罩，用于快速调整。
  - **High**: 在 High 质量模式下扩展 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用到整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩同时提供给效果时，决定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Edge Width** (Default: 0.02, Range: 0 or greater)
  轮廓线边缘的宽度。增大可获得更粗的轮廓线。

- **Edge Strength** (Default: 2, Range: 0 or greater)
  按此数值缩放轮廓线边缘的强度。增大可获得更重的边缘。

- **Edge Threshold** (Default: 0.1, Range: 0 or greater)
  从轮廓线图像中减去此值。增大可去除不需要的噪声和次要边缘。

- **Edge Color** (Default rgb: [0 0 0])
  用此颜色勾画素材的边缘轮廓。

- **Suppress Small Edges** (Default: 0.5, Range: 0 or greater)
  增大此值可去除较小的边缘，同时保留较大的边缘。

- **Edge Sharpen** (Default: 0, Range: 0 or greater)
  锐化轮廓线的程度。增大此值可使边缘两侧更锐利。


### Paint Parameters:

Frequency:
*Default:
*50,
*Range:
*0.3 or greater.画面中笔触的密度。增大可获得更小的笔触。

Stroke Length:
*Default:
*2,
*Range:
*any.确定笔触沿源素材边缘方向的长度。如果为负值，可以在 VanGogh 和 HairyPaint 风格之间切换。

Stroke Align:
*Default:
*0.2,
*Range:
*0 or greater.增大以平滑笔触方向，使相邻笔触更加平行。

Smooth Colors:
*Default:
*0,
*Range:
*0 or greater.在生成笔触之前按此数值模糊源素材。增大可使相邻笔触的颜色更加一致。

Seed:
*Default:
*0,
*Range:
*0 or greater.用于初始化随机数生成器。实际种子值并不重要，但不同的种子会产生不同的结果，相同的值应该会产生可重复的结果。

Jitter Frames:
*Integer, Default:
*0,
*Range:
*0 or greater.
如果为 0，笔触的位置在每一帧处理时保持不变。如果为 1，笔触位置在每帧重新随机化。如果为 2，则每隔一帧重新随机化，以此类推。

### Posterize Parameters:

Posterize Amount:
*Default:
*0,
*Range:
*0 to 1.如果为正值，通过限制结果中的颜色数量生成色调分离效果。增大此值可获得更少和更大的纯色区域。减小可获得更多颜色和颜色之间更多的过渡步骤。

Posterize Smooth:
*Default:
*0.1,
*Range:
*0 to 1.色调分离时平滑颜色区域之间边缘的量。增大此值可减少彩色区域之间的锯齿。如果设为 1，区域将完全平滑在一起，不会产生色调分离效果。

Posterize Phase:
*Default:
*0,
*Range:
*any.
色调分离时移动颜色边界的量。调整此值可微调颜色区域之间边缘的位置。相位为 1 等同于 0。

### Color Correct Parameters:

Saturation:
*Default:
*1,
*Range:
*any.缩放颜色饱和度。增大可获得更强烈的颜色。设为 0 可获得单色效果。

Scale Lights:
*Default:
*1,
*Range:
*0 or greater.按此值缩放结果。增大可获得更亮的结果。

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
*[0 0 0].将此颜色添加到源的较暗区域。

Opacity:
*Popup menu, Default: Normal
*.确定处理不透明度/透明度的方法。
*All Opaque:
*当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。*Normal:
*正常处理不透明度。*As Premult:
*按图像已为预乘形式处理（颜色已按不透明度缩放）。此选项的渲染速度也略快于 Normal 模式，但结果也将为预乘形式，这有时不太准确。

Mask Use:
*Popup menu, Default: Luma
*.确定如何使用 Mask 输入通道来生成单色遮罩。
*Luma:
*使用 RGB 通道的亮度。*Alpha:
*仅使用 Alpha 通道。

Blur Mask:
*Default:
*0.05,
*Range:
*0 or greater.在使用前按此数值模糊遮罩输入。可提供遮罩区域和非遮罩区域之间更平滑的过渡。除非提供了遮罩输入，否则无效。

Invert Mask:
*Check-box, Default:
*off.
如果开启，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则无效。
