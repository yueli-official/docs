---
title: WipeClouds
---

## S_WipeClouds

使用移动的云纹理从第一个素材转场到第二个素材。应对 Wipe Percent 参数设置动画以控制转场速度。

在 Sapphire Transitions 效果子菜单中。

![WipeClouds](../_static/WipeClouds.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始转场。

- **Background**: 默认为无。以此素材结束转场。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Transition Dir** (Popup menu, Default: Wipe Off to Bg)
  选择转场的方向。
  - **Wipe Off to Bg**: 从当前图层转场到背景。
  - **Wipe On from Bg**: 从背景转场到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  如果启用，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过对 Wipe Percent 参数设置动画来手动执行转场。

- **Wipe Percent** (Default: 0, Range: 0 to 1)
  必须禁用 Auto Trans 才能使用此参数。它决定了 From 和 To 输入之间的转场比例，通常应从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线以更精细地控制擦除时序。

- **Frequency** (Default: 2, Range: 0.1 or greater)
  云图案的频率。增大可获得更多更小的元素，减小可获得更少更大的元素。

- **Frequency Rel X** (Default: 0.4, Range: 0.01 or greater)
  纹理的相对水平频率。增大可将其垂直拉伸，减小可将其水平拉伸。

- **Octaves** (Integer, Default: 8, Range: 1 to 10)
  噪声叠加层的数量。每个八度是前一个的两倍频率和一半振幅。单个八度产生平滑的纹理。添加八度使结果趋近于分形（1/f）噪声纹理。

- **Seed** (Default: 0.23, Range: 0 or greater)
  用于初始化随机数生成器。实际种子值并不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Shift Start** (X & Y, Default: [0 0], Range: any)
  纹理的平移偏移。由于纹理是程序化生成的，因此可以在不出现重复单元或接缝的情况下进行移动。

- **Shift Speed** (X & Y, Default: [2 0], Range: any)
  纹理的平移速度。如果非零，结果将自动以此速率动画移动。动画化 Speed 值的结果可能不直观，因此对于可变速度运动，通常最好将此设置为 0 并改为对 Shift Start 值设置动画。

- **Grad Add** (Default: 0, Range: -10 to 10)
  如果为正值，将在转场图案的时序中添加渐变，使其在擦除过程中移动到屏幕上。如果启用了 Wipe Widget，可以使用它调整此参数，但值必须为正才能使此控件可见。

- **Grad Angle** (Default: 0, Range: any)
  擦除渐变的方向（以度为单位）。除非 Grad Add 为正值，否则此参数无效。Wipe Widget 也允许调整此参数。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（Alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按照图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项也比 Normal 模式渲染速度稍快，但结果也将是预乘形式，这有时不太正确。

- **Show Wipe** (Check-box, Default: on)
  打开或关闭用于调整 Grad Add、Grad Angle 和 Wipe Percent 参数的屏幕用户界面控件。Grad Add 参数的值必须首先为正值才能使此控件可见。此参数仅在 AE 和 Premiere 中显示，因为这些软件支持屏幕控件。
