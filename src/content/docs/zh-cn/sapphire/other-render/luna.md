---
title: Luna
---

## S_Luna

渲染地球的月亮；您可以调整月相和颜色，并添加大气效果。

在 Sapphire Render 效果子菜单中。

![Luna](../_static/Luna.jpg)


### Inputs:

- **Background**: 当前图层。用作背景的素材。

- **Mask**: 默认为无。在结果和源输入之间插值。白色区域使用效果结果。黑色区域使用源素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: Luna)
  选择如何确定月相。您可以在 Luna 模式下直接调整，也可以选择 LunaDate 模式来选择日期和时间，效果将使用该日期的正确月相。
  - **Luna**: 选择此模式手动调整月相。
  - **LunaDate**: 选择此模式让效果根据给定的日期和时间计算月相。

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

- **Center** (X & Y, Default: [-0.44 0.17], Range: any)
  月亮的中心点。此参数可通过 Center 控件调整。

- **Center Uses Mocha** (Check-box, Default: off)
  控制月亮的中心是由 Center 参数控制，还是跟随 Mocha 内部跟踪的 Center。

- **Smooth Center Track** (Integer, Default: 0, Range: 0 or greater)
  控制在稳定 Mocha 点跟踪时平均多少个点。

- **Size** (Default: 0.3, Range: 0 or greater)
  月亮的大小。此参数可通过 Size 控件调整。

- **Lunar Phase** (Default: 65, Range: any)
  月相，单位为度数；0 为新月，90 为上弦月，180 为满月，270 为下弦月。仅在 Luna 模式下可用。

- **Year** (Integer, Default: 2.02e+03, Range: 1900 to 2295)
  计算月相时使用的年份。

- **Month** (Integer, Default: 3, Range: 1 to 12)
  计算月相时使用的月份。

- **Day** (Integer, Default: 16, Range: 1 to 31)
  计算月相时使用的日期。

- **Hour** (Integer, Default: 16, Range: 0 to 23)
  计算月相时使用的小时。

- **Minute** (Default: 0, Range: any)
  计算月相时使用的分钟。

- **GMT Offset** (Default: -5, Range: -12 to 12)
  计算月相时使用的 GMT 偏移量。-5 为美国东部标准时间，-8 为美国太平洋标准时间。

- **Rotation** (Default: 30, Range: any)
  月亮图像的旋转角度，单位为度。

- **Bumpiness** (Default: 0.3, Range: 0 to 1)
  月亮有凹坑可以捕捉和反射光线。此参数可用于调整这些凹坑的凹凸程度。0 为完全平滑，1 为非常粗糙。0.3 约为晴朗夜晚的物理真实值。

- **Contrast** (Default: 1, Range: 0 to 1)
  调整月亮的对比度。接近 0 的值会使暗部区域变亮。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果的亮度。

- **Color** (Default rgb: [1 1 1])
  缩放结果的颜色。例如，如果设为黄色 [1 1 0]，结果中的蓝色将为 0。

- **Earth Glow** (Default rgb: [0 0 0])
  添加地球光照，这在日落时分月亮为弯月时经常可以看到。太阳光从地球反射，部分反射光照亮了月亮的暗面。在月食期间效果尤其好看。

- **Gamma** (Default: 1.6, Range: 0.1 or greater)
  设置月亮图像的整体伽马值。适用于以不同于对比度参数的方式降低对比度。

- **Sky Color** (Default rgb: [0 0 0])
  如果您想制作包含月亮和彩色天空的完整天空图像，可以将 Sky Color 设为蓝色来将月亮放在蓝天中。这也会使月亮偏向天空颜色。

- **Glow Brightness** (Default: 0.5, Range: 0 or greater)
  为月亮添加一些辉光。在现实生活中，当有雾气或薄云时经常可以看到这种效果。

- **Threshold** (Default: 0.01, Range: 0 or greater)
  辉光的阈值；只有月亮中亮于此阈值的部分才会发光。

- **Glow Size** (Default: 0.75, Range: 0 or greater)
  月亮辉光的大小。越大则辉光越弥散。

- **Halo Brightness** (Default: 0, Range: 0 or greater)
  在某些高空弥散云层条件下，有时可以看到月亮周围微妙的彩虹光晕。增大此参数可以看到该光晕。

- **Halo Rel Size** (Default: 1.5, Range: 0 or greater)
  设置光晕相对于月亮的大小。1.0 与月亮大小相同，2.0 为月亮的两倍大。

- **Color Fringing** (Default: 0.05, Range: 0 to 1)
  增加或减少月亮光晕中色散的程度；色散将颜色分离成彩虹。

- **Inner Softness** (Default: 0.15, Range: 0 or greater)
  设置光晕内侧（靠近月亮的一侧）的柔和度或扩散程度。

- **Outer Softness** (Default: 0.3, Range: 0 or greater)
  设置光晕外侧（远离月亮的一侧）的柔和度或扩散程度。

- **Halo Saturation** (Default: 0.5, Range: 0 or greater)
  设置月亮光晕的整体饱和度。增大可获得更图形化的外观。

- **Halo Tint** (Default rgb: [1 1 1])
  将光晕偏向此颜色。

- **Atmosphere Amp** (Default: 0.3, Range: 0 or greater)
  大气参数为辉光和光晕添加少量噪波，以获得更逼真的外观。Atmosphere Amp 控制大气噪波的量。

- **Atmosphere Freq** (Default: 2, Range: 0.1 or greater)
  控制大气噪波的频率。

- **Atmosphere Turbulence** (Default: 0.6, Range: 0 to 1)
  控制大气噪波中的湍流（细节量）。

- **Atmosphere Seed** (Default: 0.123, Range: 0 or greater)
  设置大气噪波的种子值。

- **Atmosphere Speed** (Default: 1, Range: any)
  控制大气噪波随时间变化的速度。

- **Combine** (Popup menu, Default: Overlay)
  允许您以多种方式将月亮图像与背景合成。
  - **Moon Only**: 忽略背景；仅显示月亮（及其辉光和光晕）。
  - **Overlay**: 将月亮叠加（合成）在背景上。
  - **Add**: 将月亮叠加到背景上。
  - **Screen**: 将月亮与背景进行滤色合成。适合白天拍摄的镜头。
  - **Max**: 在月亮比背景亮的地方显示月亮。这对白天有云的镜头很有用。在云比月亮亮的地方，云会遮挡月亮。
  - **Transparent Shadow**: 仅将月亮的亮部合成到背景上，暗部保持透明。这并非物理真实，因为月亮的暗部会遮挡其后方的天空和星星，但可以用于图形效果。

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  在与月亮合成之前缩放背景的亮度。如果为 0，结果将仅包含黑色背景上的月亮图像。

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  如果此值为正，输出的 Alpha 通道将包含来自月亮光晕和辉光的一些不透明度。

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

- **Show Size** (Check-box, Default: on)
  开启或关闭用于调整 Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show Center** (Check-box, Default: on)
  开启或关闭用于调整 Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。
