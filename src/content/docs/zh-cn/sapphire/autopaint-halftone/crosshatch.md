---
title: Crosshatch
---

## S_Crosshatch

使用重叠笔触模拟钢笔素描交叉影线效果。源图像根据亮度分为四个色带；从暗到亮的每个色带获得不同的笔触图案。

位于 Sapphire Stylize 效果子菜单中。

![Crosshatch](../_static/Crosshatch.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。在结果和源输入之间进行插值。白色区域使用效果结果。黑色区域使用源素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: CrosshatchPencil)
  选择铅笔或粉笔模式。
  - **CrosshatchPencil**: 模拟白纸上的深色铅笔或钢笔笔触。
  - **CrosshatchChalk**: 模拟深色纸张上的白色粉笔笔触。

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

- **Stroke Frequency** (Default: 100, Range: 1 to 500)
  增大可获得更小、更细的笔触；减小可获得更宽的笔触。

- **Stroke Length** (Default: 10, Range: 0.1 or greater)
  笔触的平均长度，与宽度相比。

- **Stroke Strength** (Default: 0.55, Range: 0 to 1)
  整体大小和强度；减小可获得更少、更小的笔触。为零时，笔触将消失。增大可获得更粗、更重叠的笔触。为一时，到处都有笔触，因此看不到笔触图案。

- **Stroke Softness** (Default: 0.1, Range: 0.001 to 1)
  每个笔触边缘的柔和度。减小可获得硬边钢笔笔触；增大可获得更柔和的粉笔效果。

- **Stroke Angle** (Default: 45, Range: any)
  笔触的角度，以度为单位；零使笔触水平和垂直。

- **Stroke Shift** (X & Y, Default: [0 0], Range: any)
  移动整体笔触图案；这可以帮助将笔触图案与素材中的整体摄像机运动匹配。

- **Animate Speed** (Default: 1, Range: 0 to 5)
  笔触通常会随时间微妙变化；此参数控制该动画的速度。设为零可获得不移动的静态笔触。

- **Threshold Darks** (Default: 0.15, Range: 0 to 1)
  最暗区域获得双重重叠笔触（或粉笔模式下的纯黑）；亮度低于此阈值的源区域被视为最暗色带，获得双重笔触。增大此值（或任何阈值）将使整体结果变暗，因为更多图像将落入最暗色带。

- **Threshold Mids** (Default: 0.35, Range: 0 to 1)
  中间调分为较暗中间调和较亮中间调；此阈值设置分隔这两个色带的亮度值。较暗中间调获得更暗和更密集的笔触。

- **Threshold Brights** (Default: 0.6, Range: 0 to 1)
  最亮区域获得最淡的笔触（通常只是白色，除非在粉笔模式下）；亮度高于此阈值的区域被视为亮色。

- **Thresholds Add** (Default: 0, Range: any)
  对所有阈值进行加减；增大可使整体结果变暗（因为提高了阈值），减小可使整体结果变亮（因为降低了阈值）。

- **Mix Threshold** (Default: 0.005, Range: 0 to 0.1)
  柔化暗/中/亮亮度色带之间的边界。

- **Strokes Use Source** (Default: 0, Range: 0 to 1)
  增大可使用更多源颜色为笔触着色。零表示使用笔触颜色；一表示使用底层源素材的颜色。在笔触之间，背景颜色会透出；如果您将 Back Style 设置为 Source，当此值设为一时笔触将消失。

- **Stroke Color** (Default rgb: [0 0 0])
  用于笔触的颜色。在铅笔模式下默认为黑色；在粉笔模式下默认为白色。

- **Posterize Amount** (Default: 0, Range: 0 to 1)
  对源进行色调分离，产生更具卡通感的纯色区域效果。仅在使用源着色笔触或使用源作为背景时有效。

- **Posterize Smooth** (Default: 0, Range: 0 to 1)
  色调分离时，平滑纯色区域的边缘。这可以避免锯齿，通常效果更好。

- **Posterize Phase** (Default: 0, Range: any)
  调整色调分离的相位。使用此参数定位平面颜色区域并避免在您希望保持平整的区域中出现边缘。

- **Edge Strength** (Default: 0, Range: 0 or greater)
  为效果添加卡通风格的边缘。

- **Edge Width** (Default: 0.002, Range: 0 or greater)
  调整边缘笔触的宽度；增大此值也会柔化边缘。

- **Edge Threshold** (Default: 0.5, Range: 0 or greater)
  增大此值可去除次要的、不重要的边缘笔触，产生更大胆的外观。

- **Edge Color** (Default rgb: [0 0 0])
  设置边缘笔触的颜色。

- **Suppress Small Edges** (Default: 0.5, Range: 0 or greater)
  增大可抑制小的、次要的边缘。

- **Edge Sharpen** (Default: 0, Range: 0 or greater)
  锐化边缘笔触。

- **Back Style** (Popup menu, Default: Solid Color)
  在笔触下方用作背景的内容。
  - **Source**: 使用源作为背景。这会产生更丰富多彩的外观，就像笔触画在原始素材上一样。使用此选项时您可能需要调整 Stroke Color。
  - **Solid Color**: 使用指定的纯色背景。

- **Solid Color** (Default rgb: [1 1 1])
  在 Solid Color 模式下用于背景的颜色。

- **Pre Blur Bg** (Default: 0, Range: 0 or greater)
  在用作背景或为笔触着色之前模糊源。这可以帮助减少因噪声或颗粒源导致的闪烁。

- **Use Source Alpha** (Default: 1, Range: 0 to 1)
  使用源的 Alpha 裁剪笔触。为一时，在源 Alpha 为零的地方抑制笔触；即通过 Alpha 裁剪它们。为零时，笔触在任何地方都会绘制，即使源 Alpha 为零。如果您希望帧中所有地方都有笔触纹理，请设为零。当源完全不透明时，此参数无效。

- **Saturation** (Default: 1, Range: 0 or greater)
  增大或减小输出的整体饱和度。

- **Scale Lights** (Default: 1, Range: 0 or greater)
  按此数值缩放结果的亮度。

- **Offset Darks** (Default: 0, Range: any)
  将此灰度值添加到源的较暗区域。可以为负值以增加对比度。

- **Tint Lights** (Default rgb: [1 1 1])
  按此颜色缩放结果，从而为较亮区域着色。

- **Tint Darks** (Default rgb: [0 0 0])
  将此颜色添加到结果的较暗区域。

- **Mix With Source** (Default: 0, Range: 0 to 1)
  在结果（设为 0 时）和原始源（设为 1 时）之间进行插值。0.7 可以通过将部分源混入笔触产生不错的效果。

- **Seed** (Default: 0.123, Range: 0 or greater)
  初始化笔触的随机数生成器。不同的值会产生不同的随机笔触图案。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已为预乘形式处理（颜色已按不透明度缩放）。此选项的渲染速度也略快于 Normal 模式，但结果也将为预乘形式，这有时不太准确。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用 Mask 输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊遮罩输入。可提供遮罩区域和非遮罩区域之间更平滑的过渡。除非提供了遮罩输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果开启，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则无效。
