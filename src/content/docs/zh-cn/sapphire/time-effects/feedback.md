---
title: Feedback
---

## S_Feedback

输入素材的先前帧经过变换后与当前帧组合，产生多种受视频反馈启发的效果。每个处理后帧的输出会被存储，然后与下一帧组合。每当处理任何非连续帧时，反馈都会重新初始化：包括第一帧、重新处理某一帧或跳转到另一帧。您必须连续处理素材的多个帧才能观察到效果，渲染前清除图像缓存有时可能是必要的。

在 Sapphire Time effects 子菜单中。

![Feedback](../_static/Feedback.jpg)


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
  在使用前按此像素量扩展或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下扩展 Mocha 遮罩以进行快速调整。
  - **High**: 在高质量模式下扩展 Mocha 遮罩以获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩同时提供给效果时，决定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Max Steps** (Integer, Default: 15, Range: 2 or greater)
  调整渲染的反馈循环次数。

- **Prev Brightness** (Default: 0.8, Range: 0 or greater)
  对于每一帧，先前的输出在与新输入帧组合之前按此数值缩放。通常此值应小于 1.0，这会使先前的帧随时间逐渐淡出。值为 1.0 时不会淡出，大于 1.0 的值会使先前的帧随时间变得更亮。

- **Prev Color** (Default rgb: [1 1 1])
  对于每一帧，先前的输出在与新输入帧组合之前按此颜色缩放。这与 Prev Brightness 类似，但影响先前帧的颜色而不仅仅是亮度。

- **Prev Hue Shift** (Default: 0, Range: any)
  对于每个新帧，偏移先前帧颜色的色相。

- **Combine New** (Popup menu, Default: Ave)
  选择将先前帧与当前帧组合的方法。
  - **Ave**: 当前帧与先前输出进行平均，使运动对象随时间产生拖尾。输出按 Fade 缩放，输入按 1.0-Fade 缩放以进行加权平均，因此 Fade 必须小于 1.0 才能正常工作。与其他组合选项不同，Ave 不应影响素材中静止对象的亮度。
  - **Max**: 当前帧和先前帧的颜色通过最大值函数组合。这使输出帧至少与当前帧一样亮，例如在暗背景上移动的亮对象会产生更亮的"拖尾"效果。
  - **Screen**: 当前帧和先前帧的颜色通过混合函数组合。这可用于积累运动素材的颜色。但是，非黑色区域会随每一帧变得更亮。
  - **Add**: 当前帧和先前帧的颜色被相加。这也可用于积累运动素材的颜色，非黑色区域在每一帧都会变得更亮。
  - **Over**: 当前帧使用其 Alpha 通道合成在先前帧之上。这使用预乘合成，因此在 Alpha 为黑色的地方，源图像通常也应为黑色。如果输入素材不包含 Alpha 通道，则使用亮度代替。
  - **Under**: 当前帧合成在先前帧之下。
  - **Min**: 当前帧和先前帧的颜色通过最小值函数组合。这使输出帧不会比当前帧更亮，通常会快速淡化为黑帧。

- **New Color** (Default rgb: [1 1 1])
  缩放当前帧的颜色。将其设置为 Old Color 的补色可抵消过度着色的拖尾。

- **New Opacity** (Default: 1, Range: 0 to 10)
  缩放当前帧的不透明度和亮度。

- **Blur Amount** (Default: 0, Range: 0 or greater)
  对于每个新帧，先前帧按此数值进行模糊。除非此值为正数，否则不起作用。

- **Diffuse Amount** (Default: 0, Range: 0 or greater)
  对于每个新帧，先前帧通过此强度的像素扩散处理。除非此值为正数，否则不起作用。

- **Amount Rel** (X & Y, Default: [1 1], Range: 0 or greater)
  水平和垂直模糊和/或扩散的相对量。除非 Blur Amount 或 Diffuse Amount 为正数，否则不起作用。

- **Center** (X & Y, Default: [0 0], Range: any)
  旋转和缩放的中心位置，以相对于帧中心的屏幕坐标表示。此参数可通过 Center Widget 调整。

- **Z Dist** (Default: 0.95, Range: 0.001 to 10)
  对于每个新帧，先前帧的"距离"按此数值缩放。这会在反馈过程中产生缩放效果。大于 1.0 的值会缩小先前帧使其变小，小于 1.0 的值会放大先前帧使其变大。此参数可通过 Transform Widget 调整。

- **Rotate** (Default: 3, Range: any)
  对于每个新帧，应用于先前帧的旋转角度（度）。此参数可通过 Transform Widget 调整。

- **Shift** (X & Y, Default: [0 0], Range: any)
  对于每个新帧，按此数值偏移先前帧。如果此值非零，则 Center 位置的意义较小。此参数可通过 Transform Widget 调整。

- **Wrap** (X & Y, Popup menu, Default: [ No No ])
  确定访问源图像边界之外区域的方法。
  - **No**: 边界之外显示黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复镜像副本。使用此方法时边缘通常不太明显。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按照图像已为预乘形式（颜色已按不透明度缩放）进行处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将为预乘形式，有时不太准确。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用 Mask 输入通道来创建单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊遮罩输入。这可以在遮罩区域和非遮罩区域之间提供更平滑的过渡。除非提供了遮罩输入，否则不起作用。

- **Invert Mask** (Check-box, Default: off)
  如果启用，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则不起作用。

- **Show Center** (Check-box, Default: on)
  开启或关闭用于调整 Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些宿主支持屏幕小工具。

- **Show Transform** (Check-box, Default: on)
  开启或关闭用于调整 Z Dist 和 Rotate 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些宿主支持屏幕小工具。

- **Show Shift** (Check-box, Default: off)
  开启或关闭用于调整 Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些宿主支持屏幕小工具。

