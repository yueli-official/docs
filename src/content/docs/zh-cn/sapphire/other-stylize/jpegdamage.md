---
title: JpegDamage
---

## S_JpegDamage

创建源素材输入经过 Jpeg 压缩伪影和错误处理后的版本。可用于呈现各种低质量数字传输的外观。提供了三种图像处理方法：可以调整 Jpeg 质量、缩放各种内部频率，以及引入随机解压缩错误。在所有情况下，降低分辨率因子也可以创建更大、更明显的 Jpeg 色块。

位于 Sapphire Stylize 效果子菜单中。

![JpegDamage](../_static/JpegDamage.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。在结果和源素材输入之间进行插值。白色区域使用效果结果。黑色区域使用源素材。


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
  在使用前按此像素数扩展或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下扩展 Mocha 遮罩，用于快速调整。
  - **High**: 在高质量模式下扩展 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩同时提供给效果时，决定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Quality** (Default: 0.1, Range: 0.01 to 1)
  决定正常 Jpeg 伪影的程度。使用较低的值获得更多压缩。

- **Res Factor** (Integer, Default: 1, Range: 1 or greater)
  按此数值的倒数降低结果的分辨率，因此 1 为全分辨率，2 为 1/2，3 为 1/3，依此类推。增大时像素形状会变大。除非该值超过当前查看的降分辨率因子，否则不会注意到此参数的效果。

- **Res Rel X** (Default: 1, Range: 0.01 or greater)
  在水平方向按此数值的倒数降低结果的分辨率。如果不为 1，jpeg 色块形状将变为矩形。

- **All Freq Scale** (Default: 1, Range: 0 or greater)
  缩放所有 Jpeg 系数的频率。非 1 的值会导致异常结果，创建不寻常的块状版本的输入。

- **X Freq Scale** (Default: 1, Range: 0 or greater)
  缩放水平 Jpeg 频率。非 1 的值会导致异常结果。

- **Y Freq Scale** (Default: 1, Range: 0 or greater)
  缩放垂直 Jpeg 频率。

- **Low Freq Scale** (Default: 1, Range: 0 or greater)
  缩放较柔和的低频率。

- **Mid Freq Scale** (Default: 1, Range: 0 or greater)
  缩放中间范围的频率。

- **High Freq Scale** (Default: 1, Range: 0 or greater)
  缩放较锐利的高频率。您可能需要高 Quality 设置才能看到高频率。

- **Affect Luma** (Default: 1, Range: 0 or greater)
  决定上面的 Freq Scale 参数对亮度通道的影响程度。零值不会导致亮度变化。大于 1.0 的值会增强变化。

- **Affect Chroma** (Default: 0.5, Range: 0 or greater)
  决定上面的 Freq Scale 参数对色度通道的影响程度。零值不会导致色度变化。大于 1.0 的值会增强变化。

- **Error Rate** (Default: 0, Range: 0 or greater)
  如果为正值，将引入随机解压缩错误。该值决定接收错误的色块中的平均错误数。较大的值产生更均匀的颗粒外观。

- **Err Block Density** (Default: 0.75, Range: 0 to 1)
  决定有错误的 Jpeg 色块的百分比。值为 .5 时一半的色块有错误，1.0 时所有色块都有错误。

- **Error Amp** (Default: 1, Range: 0 or greater)
  解压缩错误的振幅。较大的值产生更明显的视觉错误。除非 Error Rate 也为正值，否则无效。

- **Error Coherence** (Default: 1, Range: 0 or greater)
  决定有错误的色块聚集在一起的程度。为零时，错误均匀分布在整个帧中。增大时，错误聚集成更大的群组。除非 Error Rate 为正值且 Err Block Density 小于 1，否则无效。

- **Jitter Frames** (Integer, Default: 1, Range: 0 or greater)
  如果为 0，随机错误将在每帧处理中保持不变。如果为 1，每帧使用不同的错误。如果为 2，每隔一帧使用新的错误，依此类推。除非 Error Rate 也为正值，否则无效。

- **Rand Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化随机数生成器。实际的种子值并不重要，但不同的种子会产生不同的随机错误图案，相同的值应产生可重复的结果。除非 Error Rate 也为正值，否则无效。

- **Scale Lights** (Default: 1, Range: 0 or greater)
  按此数值缩放结果的亮度。

- **Offset Darks** (Default: 0, Range: any)
  将此灰度值添加到源素材的较暗区域。可以为负值以增加对比度。

- **Saturation** (Default: 1, Range: any)
  缩放颜色饱和度。增大以获得更浓烈的颜色。设为 0 则为单色。

- **Flip Noise Vertically** (Check-box, Default: off)
  如果需要，垂直翻转噪声以获得一致的外观。

- **Mask Use** (Popup menu, Default: Luma)
  决定如何使用 Mask 输入通道生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊蒙版输入。可以在蒙版区域和非蒙版区域之间提供更平滑的过渡。除非提供了蒙版输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果开启，反转蒙版输入，使效果应用于蒙版为黑色而非白色的区域。除非提供了蒙版输入，否则无效。

