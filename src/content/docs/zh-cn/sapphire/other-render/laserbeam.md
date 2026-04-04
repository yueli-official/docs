---
title: LaserBeam
---

## S_LaserBeam

模拟科幻风格激光枪的光束。光束在若干帧内从源点移动到目标点。还可以添加透视效果。

在 Sapphire Render 效果子菜单中。

![LaserBeam](../_static/LaserBeam.jpg)


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
  在使用前按此数值模糊 Mocha 遮罩。可用于柔化遮罩的边缘或量化伪影，并平滑时间偏移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 遮罩的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  启用后，在应用效果之前反转 Mocha 遮罩的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素值膨胀或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Mocha 膨胀是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在高质量模式下膨胀 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当同时提供 Mocha 遮罩和输入遮罩时，决定如何组合它们。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Start** (X & Y, Default: [-0.444 -0.287], Range: any)
  光束的起始点。此参数可通过 Start 控件调整。

- **Stop** (X & Y, Default: [0.111 0.33], Range: any)
  光束的目标点。此参数可通过 Stop 控件调整。

- **Shift** (X & Y, Default: [0 0], Range: any)
  移动 Start 和 Stop 点，将整个光束移到不同位置。此参数可通过 Shift 控件调整。

- **Position** (Default: 0.5, Range: 0 or greater)
  绘制的光束段应出现在光束轨迹上的位置。值为 0 时，绘制的光束段尚未从源点射出（注意：此时不会绘制任何内容）。值为 1 时，绘制的光束段已进入目标点（注意：同样不会绘制任何内容）。在中间值时，光束段将出现在光束轨迹上。光束位置可通过关键帧手动控制此参数，也可通过时间控件使用帧编号来控制。

- **Length** (Default: 0.5, Range: 0.001 or greater)
  绘制的光束段的长度。

- **Width** (Default: 0.01, Range: 0.001 or greater)
  绘制的光束段的宽度。

- **Core Color** (Default rgb: [1 1 0])
  光束中心的颜色。

- **Edge Color** (Default rgb: [1 0 0])
  光束边缘的颜色。

- **Color Balance** (Default: 0.5, Range: 0 to 1)
  调整核心颜色和边缘颜色之间的平衡。

- **Brightness** (Default: 1, Range: 0 or greater)
  光束的亮度。

- **Softness** (Default: 1, Range: 0 or greater)
  光束中纹理的柔和度。

- **Fade Back** (Default: 0, Range: 0 or greater)
  淡出光束后半部分的亮度。设为 1 将在光束最后端淡至黑色。较高的值会更快地淡出。

- **Fade Front** (Default: 0, Range: 0 or greater)
  淡出光束前半部分的亮度。设为 1 将在光束最前端淡至黑色。较高的值会更快地淡出。

- **Start Time** (Default: 0, Range: 0 or greater)
  如果使用时间模式操作，这是光束从源点出发的时间。

- **Duration** (Default: 10, Range: 0 to 500)
  如果使用时间模式操作，这是激光束的飞行持续时间。在起始时间 + 持续时间后，光束将消失在目标处。

- **Laser Shape** (Popup menu, Default: Forward)
  绘制的光束段的形状。
  - **Smooth**: 光束段以圆形轮廓绘制。
  - **Spear**: 光束段绘制为两端尖锐的对称形状。
  - **Forward**: 光束段绘制为前端较宽、尾端较窄的不对称形状。
  - **Backward**: 光束段绘制为尾端较宽、前端较窄的不对称形状。

- **Perspective** (Default: 0, Range: -1 to 1)
  透视效果的强度。为零时无效果。正值在给定光束位置（或时间）下将光束移近目标，负值将其移近源点。

- **Use Time** (Check-box, Default: off)
  使用帧编号自动控制光束段的运动，而不是通过位置控件手动控制。

- **Breakup** (Default: 0.05, Range: 0 to 1)
  随着此值增大，绘制的光束段变得越来越参差不齐。

- **Smooth** (Default: 0, Range: 0 or greater)
  应用于光束段的整体平滑度。

- **Atmosphere Amp** (Default: 0, Range: 0 or greater)
  大气效果模拟激光穿过尘埃大气并拾取光线或被遮蔽的效果。此参数调整大气效果的量或振幅。零值产生平滑光束，较高值产生更多尘埃外观。

- **Atmosphere Freq** (Default: 4, Range: 0.1 to 20)
  控制大气噪波的空间频率。调高可获得更精细的细节，调低可获得更宽泛的整体变化。

- **Atmosphere Detail** (Default: 0.8, Range: 0 to 1)
  控制大气模拟中精细细节的数量。减小可获得更平滑的大气，增大可获得更粗糙或颗粒感的外观。

- **Atmosphere Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化大气噪波的随机数生成器。实际种子值本身并不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Atmosphere Speed** (Default: 1, Range: any)
  大气中的云状噪波会像真实尘埃云一样随时间变化；此参数控制云图案随时间变化的速度。设为零可获得静态图案。

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  如果此值为正，输出的 Alpha 通道将包含来自激光束的一些不透明度。红、绿、蓝激光束亮度的最大值按此值缩放，并在每个像素处与背景 Alpha 合成。

- **Combine** (Popup menu, Default: Screen)
  决定光束图像如何与背景合成。
  - **Screen**: 将光束与背景混合，有助于防止过亮的结果。
  - **Add**: 将光束图像叠加到背景上。
  - **Beam Only**: 在透明黑色背景上仅显示光束。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且无透明度 (alpha=1) 时，使用此选项可略微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式略快，但结果也将是预乘形式，有时不够精确。

- **Mask Use** (Popup menu, Default: Luma)
  决定如何使用遮罩输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊蒙版输入。可在蒙版区域和非蒙版区域之间提供更平滑的过渡。除非提供了蒙版输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  开启后，反转蒙版输入，使效果应用于蒙版为黑色而非白色的区域。除非提供了蒙版输入，否则无效。

- **Show Start** (Check-box, Default: on)
  开启或关闭用于调整 Start 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show Stop** (Check-box, Default: on)
  开启或关闭用于调整 Stop 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show Shift** (Check-box, Default: off)
  开启或关闭用于调整 Shift 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。
