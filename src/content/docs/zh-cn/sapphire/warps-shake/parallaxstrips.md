---
title: ParallaxStrips
---

## S_ParallaxStrips

应用一组 3D 折射玻璃条纹来分解图像。图像在每个条纹内偏移，条纹随时间移动。条纹逐渐淡入或淡出，因此与源素材的过渡是无缝的。

在 Sapphire Distort 效果子菜单中。

![ParallaxStrips](../_static/ParallaxStrips.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。在结果和源输入之间进行插值。白色区域使用效果的结果。黑色区域使用源素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

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

- **Mode** (Popup menu, Default: Automatic)
  设置效果是随时间自动演变，还是可以手动控制。
  - **Automatic**: 根据素材中的时间自动移动条纹。在此模式下忽略 Manual Amount。
  - **Manual**: 使用 Manual Amount 指定控制效果的演变。0 为开始，1 为结束。

- **Manual Amount** (Default: 1, Range: 0 or greater)
  在 Manual 模式下，控制效果的强度和演变。选择 Fade:End 时，1 将条纹放在起始位置，偏移量最大（即效果最强）。0 给出原始源素材：将条纹放在结束位置，它们变得不可见，因为偏移底层图像的量变为零。选择 Fade:Start 时，值相反，0 最强，1 淡出。

- **Ensure Full Coverage** (Push-button)
  按下此按钮将条纹数量调整为完全覆盖最后一帧所需的最小数量。

- **N Strips** (Integer, Default: 50, Range: 1 to 1000)
  要应用的折射条纹数量。条纹在图像上随机定位。

- **Size** (Default: 0.35, Range: 0 or greater)
  条纹的大小，以图像宽度为单位。

- **Rel Height** (Default: 0.3, Range: 0.001 or greater)
  条纹的高度，相对于其宽度。增大以使条纹更高。

- **Size Vary** (X & Y, Default: [0.1 0.1], Range: 0 to 1)
  增大以使每个条纹随机变大或变小。

- **Angle** (Default: 0, Range: any)
  条纹的角度；0 为水平。条纹沿其角度移动，也沿同一角度偏移图像。

- **Depth** (Default: 2, Range: 0 or greater)
  使最前面的条纹更大并移动更快，看起来像在前面，产生 3D 效果。

- **Strip Speed** (Default: 0.4, Range: any)
  设置条纹沿其主轴移动的速度。请注意，这不影响条纹内图像的折射或偏移方式，只影响条纹本身移动的速度。

- **Strip Speed Vary** (Default: 0, Range: 0 or greater)
  增大以使每个条纹的速度有一些随机性。

- **Shift Amount** (Default: 0.6, Range: any)
  设置每个条纹内图像偏移或折射的程度。偏移始终沿条纹的主轴方向。随着效果进展，偏移量逐渐变为零，无缝过渡到原始素材。详情请参阅 Fade 和 Slow Fade 参数。

- **Shift Vary** (Default: 0, Range: 0 or greater)
  增大以使每个条纹中的偏移量更加随机。

- **All Strips Shift** (X & Y, Default: [0 0], Range: any)
  在屏幕上移动所有条纹。

- **Z Dist** (Default: 1, Range: 0.001 or greater)
  在应用视差条纹之前放大或缩小源图像。

- **Show** (Popup menu, Default: Result)
  显示效果结果或条纹本身（在效果设置期间很有用）。
  - **Result**: 显示效果的结果。
  - **Strips Over Source**: 将每个条纹显示为灰色矩形，亮度由深度设置。未覆盖的区域显示源图像。
  - **Strips Over Black**: 将每个条纹显示为灰色矩形，亮度由深度设置。未覆盖的区域显示为黑色。

- **Slow Fade** (Default: 0.9, Range: 0 to 2)
  增大以使淡入或淡出更慢。设置为 0 以获得线性淡化。

- **Full Height** (Check-box, Default: off)
  打开以使条纹始终为全高，忽略 Rel Height。这可以产生不错的滑动条纹效果。

- **Fade** (Popup menu, Default: End)
  效果在一端淡出，以实现无缝的开始或结束。
  - **Start**: 使效果在开始时缓慢淡入。
  - **End**: 使效果在结束时缓慢淡出。

- **Wrap** (Popup menu, Default: Reflect)
  确定访问源图像边界外区域的方法。
  - **No**: 在边界外显示黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复图像的镜像副本。使用此方法时边缘通常不太明显。

- **Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化随机数生成器。实际种子值并不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Flip Tiles** (Check-box, Default: off)
  如有需要垂直翻转图块以获得一致的外观。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用 Mask 输入通道来创建单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用之前按此量模糊 Matte 输入。这可以在遮罩区域和非遮罩区域之间提供更平滑的过渡。除非提供了 Matte 输入，否则此选项无效。

- **Invert Mask** (Check-box, Default: off)
  如果启用，反转 Matte 输入，使效果应用于 Matte 为黑色而非白色的区域。除非提供了 Matte 输入，否则此选项无效。

