---
title: TVDamage
---

## S_TVDamage

模拟具有传输和接收问题的电视机、录像机故障以及电视硬件问题。模拟静电、干扰、重影、水平和垂直同步、嗡嗡条纹、彩色条纹、可见扫描线、VCR 快进、信号丢失、暗角、正交像管、鱼眼以及关机效果。

在 Sapphire Stylize 效果子菜单中。

![TVDamage](../_static/TVDamage.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。在结果和源输入之间进行插值。白色区域使用效果结果，黑色区域使用源素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: TVDamage Color)
  要模拟的电视类型：彩色或黑白。
  - **TVDamage Color**: 模拟彩色电视。
  - **TVDamage Mono**: 模拟黑白电视。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前对 Mocha 遮罩进行此量的模糊处理。可用于柔化遮罩的边缘或量化伪像，并平滑时间位移。

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
  在使用前对 Mocha 遮罩进行此像素量的膨胀或腐蚀处理。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认 Fast 模式下快速调整还是在 High 质量模式下获得更好效果。
  - **Fast**: 以 Fast 模式进行 Dilate Mocha，用于快速调整。
  - **High**: 以 High 质量模式进行 Dilate Mocha，获得更好看的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩并将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果并仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  决定当效果同时提供 Mocha 遮罩和输入遮罩时如何合并它们。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Reception Master** (Default: 0.4, Range: 0 or greater)
  所有接收相关伪像的主控制：静电、干扰、重影、水平和垂直同步、嗡嗡条纹和彩色条纹。调为零可获得完美接收效果，即上述所有伪像均为零。

- **Interference Amp** (Default: 0.6, Range: 0 or greater)
  模拟附近电气设备（电动机、无绳电话等）的干扰。外观为半规律间距的随机彩色点图案。点的大小由 TV Pixels 参数控制。受 Reception Master 缩放影响。

- **Ghost Amp** (Default: 0.6, Range: 0 or greater)
  重影是由发射机和电视之间的多径失真产生的图像副本。增大此参数可获得更强的重影效果。受 Reception Master 缩放影响。

- **Horizontal Hold** (Default: 0.5, Range: 0 or greater)
  导致图像以半随机方式水平偏移，模拟水平同步电路故障的电视机，或信号不够强无法锁定水平同步的情况。受 Reception Master 缩放影响。

- **Vertical Hold** (Default: 0.8, Range: 0 or greater)
  导致图像以滚动方式垂直偏移，通常由信号过弱导致电视无法锁定引起。此参数控制图像出现同步问题的时间比例。设为零可消除垂直同步问题。受 Reception Master 缩放影响。

- **Bars Brightness** (Default: 0.3, Range: 0 or greater)
  电源嗡嗡声和其他电视问题可导致滚动的明暗条纹在屏幕上爬行。将摄像机与电视输出同步失败也会造成此现象。此参数控制这些条纹的整体强度。有两组条纹，一大一小，相互干扰。此参数控制条纹的整体亮度比例。调为零可消除条纹。受 Reception Master 缩放影响。

- **Color Stripes Amplitude** (Default: 0.2, Range: 0 or greater)
  另一种常见的干扰形式，彩色条纹由色度信号的相位偏移等原因引起。此参数控制彩色条纹的整体亮度。受 Reception Master 缩放影响。

- **Fast Forward Amount** (Default: 0, Range: 0 or greater)
  生成屏幕上带有撕裂条纹的 VCR 快进效果。

- **Tape Dropout Brightness** (Default: 0, Range: 0 or greater)
  在随机帧的随机时间生成 VCR 信号丢失效果。

- **Vignette Darkness** (Default: 0, Range: 0 to 1)
  暗角是图像朝向角落和边缘的变暗效果。此参数控制屏幕外角变暗（暗角）的程度。0 表示无暗角，1 表示最大变暗效果。

- **Static Amplitude** (Default: 0.8, Range: 0 or greater)
  缩放静态噪点的亮度。受 Reception Master 缩放影响。静态点的大小由 TV Pixels 参数控制。

- **Static Density** (Default: 0.7, Range: 0.01 to 1)
  静态噪点的密度；增大可获得更多静态像素；减小可获得只有偶尔出现的静态像素。

- **Frequency** (Default: 1.28, Range: 0 to 500)
  干扰频率。外观对此参数非常敏感。0.3 或 1.23 等分数值比整数值效果更好。轻微动画处理，例如从 1.27 到 1.3，可获得不错的效果。

- **Dots Speed** (X & Y, Default: [100 -10], Range: any)
  点图案随时间在 X 和 Y 方向上以此速度移动。

- **Jitter Amount** (Default: 10, Range: 0 to 1000)
  增大此值可使点图案在帧间随机抖动，以获得更逼真的效果。

- **Num Ghosts** (Integer, Default: 5, Range: 0 to 30)
  重影图像的数量。一些可能在源图像的前方（左侧），大多数在右侧。一些为正值，一些为负值（反转）。请参阅下方的 Shift 和 Negative Ghosts 参数。

- **Negative Ghosts** (Default: 0.5, Range: 0 to 1)
  平均而言，负值（反转）重影占总重影数的比例。

- **Spacing** (Default: 0.2, Range: 0 or greater)
  重影图像分布的图像宽度比例。

- **Vary Position** (Default: 0.3, Range: 0 to 1)
  控制重影图像间距的规律性。设为零可获得规律间距的重影；设为一可获得随机定位。

- **Shift** (Default: 0.5, Range: -1 to 1)
  将重影图像向左或向右偏移，而不移动主图像。

- **Blur** (Default: 0, Range: 0 or greater)
  模糊重影图像，但不模糊主图像或其他伪像。

- **H Frequency** (Default: 1.25, Range: 0 or greater)
  水平同步波的垂直频率。

- **H Time Vary** (Default: 0.5, Range: 0 or greater)
  随时间按此量调制水平同步波。增大时，某些帧的水平偏移更多，而其他帧更少。

- **H Octaves** (Integer, Default: 3, Range: 1 to 10)
  水平同步波的倍频程数。增加可获得更尖锐的外观，减少可获得更平滑的波形。

- **Border Width** (Default: 0.05, Range: 0 or greater)
  电视信号在显示区域外有一个黑色边框；当水平同步不工作时此边框变得可见。此参数控制该黑色边框的宽度。在边框的另一侧，可以看到图像的另一个副本。

- **V Frequency** (Default: 2, Range: 0 or greater)
  垂直同步跳动的频率。减小可获得更连贯的滚动运动，增大可获得更跳动的外观。

- **V Speed** (Default: 2, Range: 0 or greater)
  垂直同步滚动运动随时间的平均速度。

- **V Random** (Default: 0.1, Range: 0 or greater)
  控制垂直同步滚动运动中的随机性大小。设为零可获得平滑滚动，设为 1 或更大可获得抖动行为。

- **Border Height** (Default: 0.1, Range: 0 or greater)
  与 Border Width 类似，此参数控制垂直同步未锁定时变得可见的帧间垂直边框高度。此边框中通常会显示一些静态、隐藏字幕和时间码信息。

- **Border Data** (Default: 1, Range: 0 to 10)
  在 Border Width 指定的垂直消隐区间内出现的点和线的亮度。

- **Bar Roll Speed** (Default: 0.5, Range: -10 to 10)
  条纹向上滚动的速度。设为负值可获得向下滚动效果。

- **Bar Sharpness** (Default: 0.5, Range: 0.1 to 10)
  锐化或平滑主条纹的上下边缘。设为零表示无主条纹；只能看到较小的条纹。

- **Bar Frequency** (Default: 1, Range: 0.1 or greater)
  条纹的频率；增大可获得更多更细的条纹，减小可获得更少更粗的条纹。

- **Bar1 Width** (Default: 0.35, Range: 0 to 1)
  主条纹中亮色部分的比例；其余为暗色。

- **Bar2 Rel Frequency** (Default: 6, Range: 1 or greater)
  控制较小条纹的频率。

- **Bar2 Sharpness** (Default: 0.5, Range: 0.01 to 10)
  锐化或平滑较小条纹的上下边缘。设为零表示无小条纹；只能看到主条纹。

- **Color Frequency** (Default: 10, Range: 1 to 40)
  彩色条纹的空间频率。

- **Color Angle** (Default: 160, Range: any)
  条纹的角度。

- **Roll Speed** (Default: 3, Range: 0 or greater)
  控制条纹随时间滚动的速度。

- **Band Frequency** (Default: 4, Range: 0 or greater)
  要创建的快进条带数量。

- **Band Shift** (Default: 0.1, Range: 0 or greater)
  将快进条带向上或向下偏移。

- **Band Height** (Default: 0.16, Range: 0 to 1)
  每个快进条带的高度。

- **Dropout Length** (Default: 0.25, Range: 0 to 1)
  每条信号丢失扫描线的平均长度。

- **Dropout Gap Length** (Default: 0.2, Range: 0 to 2)
  信号丢失之间间隙的平均长度。

- **Dropout Y Freq** (Default: 5, Range: 0 to 50)
  信号丢失按此频率的噪声函数出现在随机扫描线上。减小可获得少量大范围信号丢失带；增大可获得大量小范围信号丢失带。

- **Dropout Y Threshold** (Default: 0.75, Range: 0 to 1)
  增大可使屏幕上（平均而言）有更多信号丢失覆盖区域；减小可覆盖更少区域。如果在某些帧上看不到任何信号丢失，请增大此参数。

- **Dropouts Always** (Default: 1, Range: 0 to 1)
  信号丢失只出现在某些帧上；增大此参数可在更多帧上看到信号丢失，使其在时间上更频繁出现。如果在任何帧上都看不到信号丢失，请增大此参数。

- **Vignette Radius** (Default: 1, Range: 0 or greater)
  从中心开始出现暗角的距离。

- **Vignette Edge Softness** (Default: 0.5, Range: 0 or greater)
  暗角软边的宽度。值越大，边缘越柔和、越不明显。

- **Vignette Rel Height** (Default: 0.75, Range: 0.1 or greater)
  控制暗角椭圆的长宽比。通常应设置为图像的长宽比，例如 NTSC 的 0.75。

- **Scanlines** (Default: 0.1, Range: 0 or greater)
  在图像中创建可见的扫描线。增大可获得更强烈的扫描线效果，设为零可消除扫描线。扫描线的宽度由 TV Pixels 参数和 Scanlines Rel Freq 控制。

- **Scanlines Rel Freq** (Default: 1, Range: 0 or greater)
  电视扫描线的相对频率。增大可获得更多扫描线，减小可获得更少更大的扫描线。注意扫描线数量也受 TV Pixels 参数控制。

- **Orthicon** (Default: 0, Range: 0 or greater)
  在源素材中亮于给定阈值的区域周围使画面变暗，以模拟 1950 年代"正交像管"电视摄像机的外观。在黑白模式下最为有用。

- **Threshold** (Default: 0.7, Range: 0 or greater)
  变暗效果将出现在源素材中亮于此值的位置周围。值为 0.9 时仅在最亮的点周围产生暗辉光；值为 0 时在每个非黑色区域产生辉光。

- **Darks Width** (Default: 0.2, Range: 0 or greater)
  缩放暗辉光的距离。


### Color Correct Parameters:

- **Hue Shift** (Default: 0, Range: any)
  按此量偏移颜色色相。

- **Saturation** (Default: 1, Range: any)
  缩放颜色饱和度。增加可获得更鲜艳的颜色。设为 0 可获得单色效果。

- **Scale Lights** (Default: 1, Range: 0 or greater)
  按此灰度值缩放结果。增加可获得更亮的结果。

- **Offset Darks** (Default: 0, Range: any)
  向结果的较暗区域添加此灰度值。可以为负值以增加对比度。

- **Tint Lights** (Default rgb: [1 1 1])
  按此颜色缩放结果，从而给较亮区域着色。

- **Tint Darks** (Default rgb: [0 0 0])
  向结果的较暗区域添加此颜色。将此设为深红橙色可获得负片胶片效果。

- **Turn Off** (Default: 0, Range: 0 to 1)
  将此参数从 0 动画到 1 以模拟电视关机效果。图像将变白并收缩到中心的一个点，在最后附近有一道闪光。

- **Flare Width** (Default: 1, Range: 0 or greater)
  关机序列结束时耀斑或闪光的宽度。设为零可省略此闪光。

- **Flare Brightness** (Default: 2, Range: 0 or greater)
  关机序列结束时耀斑或闪光的亮度。

- **Fade Out Time** (Default: 0.7, Range: 0 to 1)
  Turn Off 序列中淡出的持续时长。值为 1 时，随着 Turn Off 增大产生平滑渐进的淡出效果。值减小时，淡出将在更晚时刻（Turn Off 值更高时）开始，并更快速地完成。当 Fade Out Time 设为零时，当 Turn Off 达到 1 时立即切换为黑色。

- **Fish Eye** (Default: 0, Range: any)
  将源素材的中心扩展，如同通过鱼眼镜头观看，呈现老式略微圆形的电视外观。

- **Tv Pixels** (Default: 720, Range: 1 or greater)
  屏幕上"电视像素"的数量。控制静电、干扰、扫描线和信号丢失的大小。降低此值可模拟分辨率较低的电视机。

- **Downsample** (Check-box, Default: on)
  如果勾选，根据 TV Pixels 降低输出分辨率。

- **Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化随机数生成器。实际种子值并不重要，但不同的种子会给出不同的结果，相同的值应给出可重复的结果。

- **Mask Use** (Popup menu, Default: Luma)
  决定如何使用遮罩输入通道来制作单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前对遮罩输入进行此量的模糊处理。这可以在遮罩和非遮罩区域之间提供更平滑的过渡。除非提供了遮罩输入，否则此参数无效。

- **Invert Mask** (Check-box, Default: off)
  如果启用，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则此参数无效。
