---
title: FreeLens
---

## S_FreeLens

生成源素材的扭曲、散焦和漏光版本，以模拟将分离的镜头放在摄像机前并移动以创建对焦和光线效果的机内技术。

位于 Sapphire Blur+Sharpen 效果子菜单中。

![FreeLens](../_static/FreeLens.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Matte**: 默认为无。如果提供，效果仅在此输入的亮区指定的源素材区域上应用。此遮罩外的像素不受影响，也不会参与遮罩内的受影响像素。此输入可通过 Invert Matte 或 Matte Use 参数进行调整。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此量模糊 Mocha 遮罩。可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 遮罩的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  如果启用，在应用效果前反转 Mocha 遮罩的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素量膨胀或腐蚀 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下膨胀 Mocha 遮罩，便于快速调整。
  - **High**: 在 High 质量模式下膨胀 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，只显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当同时提供 Mocha 遮罩和输入遮罩时，确定如何组合它们。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，只使用
Mocha 遮罩。

- **Shift** (X & Y, Default: [0 0], Range: 2 or less)
  水平或垂直偏移镜头。注意图像的移动方向与镜头运动方向相反。此参数可通过 Tilt X Widget 调整。

- **Tilt X** (Default: 12.5, Range: -45 to 45)
  围绕垂直轴向左或向右旋转镜头。此参数可通过 Tilt X Widget 调整。

- **Tilt Y** (Default: 0, Range: -45 to 45)
  围绕水平轴向上或向下旋转镜头。您可以同时使用 Tilt X 和 Tilt Y 来围绕任意对角轴旋转。此参数可通过 Tilt X Widget 调整。

- **Distance** (Default: 0, Range: -0.3 to 0.5)
  将镜头从摄像机移开或靠近。此参数可通过 Distance Widget 调整。

- **Rotate Highlights** (Default: 0, Range: -180 to 180)
  围绕视线旋转镜头。在真实镜头中这会旋转光圈，因此旋转被感知为图像中任何散焦高光的旋转。

- **Perspective Amount** (Default: 1, Range: 0.25 to 4)
  控制应用 Tilt X 和 Tilt Y 时镜头伸缩的量。增大可获得更多 3D 透视效果。

- **Wrap** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  确定访问源图像边界外的方法。
  - **No**: 边界外为黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复镜像副本。使用此方法时边缘通常不太明显。

- **Filter** (Check-box, Default: on)
  如果启用，在重新采样时对图像进行自适应滤波。当图像的某些部分被缩小时，可获得更好的质量结果。

- **Defocus Width** (Default: 0.104, Range: 0 or greater)
  缩放散焦模糊和高光的总量。

- **Focal Point Offset** (X & Y, Default: [0 0], Range: 2 or less)
  偏移图像中最佳焦点的位置。此参数可通过 Focal Point Offset Widget 调整。

- **Focal Uses Mocha** (Check-box, Default: off)
  控制焦点是由 Focal Point Offset 参数控制还是跟随在 Mocha 内部跟踪的 Focal Point Offset。

- **Smooth Focal Track** (Integer, Default: 0, Range: 0 or greater)
  控制稳定 Mocha 点跟踪时要平均的点数。

- **Chroma Distort** (Default: 0, Range: any)
  在图像边缘周围添加一些色差；红色和蓝色光波在真实镜头中的折射方式不同，当光线以倾斜角度照射到镜头时会产生彩色边缘。

- **Rel Height** (Default: 1, Range: 0.01 or greater)
  光圈形状的相对高度。如果不为 1，圆形会变成椭圆形等。

- **Show** (Popup menu, Default: Result)
  选择输出类型。
  - **Result**: 显示最终输出。
  - **DepthMap**: 显示当前镜头位置和方向生成的深度图。
  - **Shape**: 显示光圈形状。

- **Highlight Shape** (Popup menu, Default: 6 sides)
  确定模拟摄像机光圈的形状。
  - **Circle**: 圆形。
  - **3 sides**: 三角形。
  - **4 sides**: 正方形。
  - **5 sides**: 五边形。
  - **6 sides**: 六边形。
  - **7 sides**: 七边形。
  - **8 sides**: 八边形。
  - **9 sides**: 九边形。
  - **10 sides**: 十边形。
  - **11 sides**: 十一边形。
  - **12 sides**: 十二边形。

- **Highlight Roundness** (Default: 0, Range: any)
  修改模拟摄像机光圈的形状。值为 1 产生圆形；0 产生由 Shape 参数决定边数的平边多边形。小于 0 会使边向内挤压产生星形，大于 1 会使角向内挤压产生花形。如果 Shape 设为 Circle 则无效。

- **Boost Highlights** (Default: 1, Range: 0 or greater)
  增加源素材中高光亮度的量。增大此参数可使高光过曝而不影响暗部或中间调。

- **Highlight Threshold** (Default: 0.9, Range: 0 or greater)
  高光的最低亮度值。亮度高于此值的像素将根据 Boost Highlights 参数被增亮。

- **Distortion** (Check-box, Default: on)
  启用畸变。

- **Link Distortion To Lens** (Check-box, Default: on)
  控制畸变是由镜头调整还是手动设置。

- **Distortion Amount** (Default: 0, Range: -0.15 to 0.15)
  径向扭曲图像以产生桶形或枕形畸变。

- **Scale Width** (Default: 1, Range: 0 or greater)
  在水平方向缩放畸变。

- **Scale Height** (Default: 1, Range: 0 or greater)
  在垂直方向缩放畸变。

- **Light Leak** (Check-box, Default: on)
  启用漏光。

- **Link Leak To Lens** (Check-box, Default: on)
  控制漏光是由镜头调整还是手动设置。

- **Leak Intensity** (Default: 0, Range: any)
  缩放漏光元素的强度。

- **Leak Rel Height** (Default: 2, Range: 0 or greater)
  控制漏光元素的宽高比。

- **Leak Size** (Default: 0.35, Range: 0 or greater)
  缩放漏光元素的宽度和高度。

- **Vary Size** (Default: 1, Range: 0 or greater)
  从一个漏光元素到下一个漏光元素之间大小的变化量。

- **Light Leak Hotspot** (X & Y, Default: [-0.5 0.25], Range: any)
  定位漏光元素线的起点。此参数可通过 Light Leak Hotspot Widget 调整。

- **Light Leak Pivot** (X & Y, Default: [0 0.25], Range: any)
  定位漏光元素线的终点。此参数可通过 Light Leak Pivot Widget 调整。

- **Leak Roundness** (Default: 0.7, Range: -1 to 1)
  使漏光元素的角变圆。

- **Sides** (Integer, Default: 4, Range: 3 to 16)
  控制每个漏光元素有多少条边。

- **Copies** (Integer, Default: 5, Range: 1 to 34)
  控制有多少个漏光元素。

- **Spread** (Default: 4, Range: 0 or greater)
  控制漏光元素之间的距离。

- **Outer Color** (Default rgb: [1 0.8 1])
  控制每个漏光元素外边缘的颜色。

- **Mid Color** (Default rgb: [1 0.9 1])
  控制每个漏光元素从中心到外边缘中间位置的颜色。

- **Center Color** (Default rgb: [1 1 1])
  控制每个漏光元素中心的颜色。

- **Midpoint** (Default: 0.5, Range: 0 to 1)
  移动每个漏光元素中心和外边缘之间 Mid Color 的位置。设为 0 将 Mid Color 放在中心，设为 1 将其放在边缘。

- **Softness** (Default: 4, Range: 0 or greater)
  模糊每个漏光元素的颜色渐变。增大可获得更平滑的渐变，减小可获得更锐利的色带。

- **Glow Brightness** (Default: 1, Range: 0 or greater)
  缩放将漏光与背景合并后应用于整个图像的辉光亮度。

- **Glow Width** (Default: 0.4, Range: 0 or greater)
  辉光的宽度。增大可获得更柔和的辉光，减小可获得更锐利、更明亮的辉光。

- **Glow Threshold** (Default: 0.8, Range: 0 or greater)
  亮度高于此值的图像部分会产生辉光。

- **Shake** (Check-box, Default: on)
  启用抖动。

- **Shake Mode** (Popup menu, Default: Normal)
  控制抖动类型。
  - **Normal**: 稳定的摄像机抖动。
  - **Twitchy**: 静止期被快速抖动的爆发所打断。
  - **Jumpy**: 从一个位置突然跳到另一个位置，中间有较慢的漂移。

- **Jump Drift** (Default: 0.3, Range: 0 to 1)
  在 Jumpy 模式下，控制跳跃之间的运动速度。

- **Jump Center Bias** (Default: 0, Range: 0 or greater)
  在 Jumpy 模式下，调整每次跳跃将图像重置到原始位置的可能性。如果设为零，每次跳跃都是随机的。如果设为一，每次跳跃都会回到中心。

- **Twitch Stillness** (Default: 0.7, Range: 0 to 1)
  在 Twitchy 模式下，调整图像静止的时间比例。增大可获得更频繁的抖动。

- **Twitch Frequency** (Default: 2, Range: 0 or greater)
  在 Twitchy 模式下，控制运动和静止周期的长度。增大可获得更短、更频繁的运动爆发。

- **Amplitude** (Default: 1, Range: 0 or greater)
  缩放抖动运动的幅度。

- **Frequency** (Default: 8, Range: 0 or greater)
  增大可获得更快的抖动，减小可获得更慢的抖动。（如果对频率值进行动画，请注意结果的抖动频率也会受到值变化率的影响。）

- **Phase** (Default: 0, Range: any)
  抖动运动的时间偏移。（如果对此值进行动画，其变化率也会影响表观频率。）

- **Motion Blur** (Check-box, Default: on)
  抖动运动的运动模糊选项。

- **Mo Blur Length** (Default: 1, Range: 0 or greater)
  缩放运动模糊量。在场处理时使用约 0.5，在帧处理时使用 1.0 可获得逼真的运动模糊。如果 Motion Blur 为 No，此参数无效。

- **Blur Res** (Popup menu, Default: Full)
  选择运动模糊的分辨率因子。较高的分辨率提供更好的质量，较低的分辨率提供更快的处理速度。
  - **Full**: 使用全分辨率。
  - **Half**: 运动模糊在半分辨率下执行。
  - **Quarter**: 运动模糊在四分之一分辨率下执行。

- **Seed** (Default: 0, Range: 0 or greater)
  用于初始化随机数生成器。实际种子值不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **X Rand Amp** (Default: 0, Range: 0 or greater)
  水平随机抖动的幅度。

- **X Rand Freq** (Default: 1, Range: 0 or greater)
  水平随机抖动的频率。

- **Y Rand Amp** (Default: 0, Range: 0 or greater)
  垂直随机抖动的幅度。

- **Y Rand Freq** (Default: 1, Range: 0 or greater)
  垂直随机抖动的频率。

- **Tilt X Rand Amp** (Default: 0, Range: 0 or greater)
  水平角度随机抖动的幅度。

- **Tilt X Rand Freq** (Default: 1, Range: 0 or greater)
  水平角度随机抖动的频率。

- **Tilt Y Rand Amp** (Default: 0, Range: 0 or greater)
  垂直角度随机抖动的幅度。

- **Tilt Y Rand Freq** (Default: 1, Range: 0 or greater)
  垂直角度随机抖动的频率。

- **Distance Rand Amp** (Default: 0, Range: 0 or greater)
  缩放随机抖动的幅度。

- **Distance Rand Freq** (Default: 1, Range: 0 or greater)
  缩放随机抖动的频率。

- **Leak Int Rand Amp** (Default: 0, Range: 0 or greater)
  漏光强度随机抖动的幅度。

- **Leak Int Rand Freq** (Default: 1, Range: 0 or greater)
  漏光强度随机抖动的频率。

- **Vignette** (Check-box, Default: on)
  启用暗角。

- **Link Vig To Lens** (Check-box, Default: on)
  控制暗角是由镜头调整还是手动设置。

- **Vig Intensity** (Default: 0.2, Range: 0 or greater)
  暗角的不透明度；动画到 0 可淡出暗角。

- **Vig Center** (X & Y, Default: [0 0], Range: any)
  旋转和缩放的中心，以相对于帧中心的屏幕坐标表示。Shift 值应为零时此位置才有意义。此参数可通过 Vig Center Widget 调整。

- **Vig Squareness** (Default: 0, Range: 0 to 1)
  确定暗角形状的方形程度。设为 1.0 可获得正方形或矩形。设为 0 可获得圆形或椭圆形。中间值可获得不同程度圆角的矩形。

- **Vig Radius** (Default: 1.5, Range: 0 or greater)
  从中心到应用暗角的距离。

- **Vig Rel Height** (Default: 1, Range: 0.05 or greater)
  光圈形状的相对高度。如果不为 1，圆形会变成椭圆形等。

- **Vig Rel Width** (Default: 1, Range: 0.05 or greater)
  暗角形状的相对水平大小。增大可获得更宽的形状，减小可获得更高的形状。

- **Vig Rotate** (Default: 0, Range: any)
  旋转光圈形状。

- **Vig Edge Softness** (Default: 0.46, Range: 0 or greater)
  暗角柔和边缘的宽度。较大的值使边缘更柔和、更不明显。

- **Vig Smooth Curve** (Default: 1, Range: 0 to 1)
  如果为零，在柔和边缘区域内使用线性渐变。增大此值可使用更平滑的"S"形曲线进行插值，这可以减少对渐变起始和结束位置的视觉感知。

- **Vig Color** (Default rgb: [0 0 0])
  暗角的颜色。

- **Vig Blur Amount** (Default: 0, Range: 0 or greater)
  除了使边框变暗外，还模糊图像的边框。

- **Vig Blur Inside** (Check-box, Default: off)
  如果勾选，则模糊图像的中心（未变暗的）区域而非边框。

- **Vig Source Brightness** (Default: 1, Range: 0 or greater)
  缩放源素材的亮度。设为零可仅查看暗角。

- **Vig Combine** (Popup menu, Default: Mult)
  确定暗角如何与源素材组合。
  - **Composite**: 将暗角合成到源素材上。
  - **Mult**: 暗角颜色与源素材相乘。如果颜色不是黑色，这将选择性地为暗角区域着色。
  - **Add**: 暗角颜色添加到源素材。如果暗角颜色为黑色则无效。
  - **Screen**: 暗角颜色与源素材使用滤色操作组合。如果暗角颜色为黑色则无效。
  - **Subtract Inv**: 从源素材中减去暗角颜色的反转。反转意味着白色代替黑色，黄色代替蓝色等。此模式看起来类似于 Mult，但更强烈；它压碎黑色并保留更多高光。如果暗角颜色为白色则无效。
  - **Vignette Only**: 显示不含源素材的暗角图案。暗角效果最强的地方输出为白色（即源素材将被完全变暗的地方）。
  - **Vignette Only Inv**: 显示不含源素材的反转暗角图案。没有暗角的地方输出为白色（即源素材不会被变暗的地方）。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用 Mask 输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此量模糊 Matte 输入。这可以在遮罩区域和非遮罩区域之间提供更平滑的过渡。除非提供了 Matte 输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果开启，反转 Matte 输入，使效果应用于 Matte 为黑色而非白色的区域。除非提供了 Matte 输入，否则无效。

- **Show Distance** (Check-box, Default: on)
  打开或关闭用于调整 Shift 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show Tilt X** (Check-box, Default: on)
  打开或关闭用于调整 Shift 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show Focal Point Offset** (Check-box, Default: on)
  打开或关闭用于调整 Focal Point Offset 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show Light Leak Hotspot** (Check-box, Default: off)
  打开或关闭用于调整 Light Leak Hotspot 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show Light Leak Pivot** (Check-box, Default: off)
  打开或关闭用于调整 Light Leak Pivot 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show Vig Center** (Check-box, Default: off)
  打开或关闭用于调整 Vig Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

