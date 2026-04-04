---
title: DissolveZap
---

## S_DissolveZap

使用动态闪电在两个素材之间进行转场。素材相互溶解，同时闪电逐渐增强。
应对 Dissolve Percent 参数进行动画处理以控制转场速度。

在 Sapphire Transitions 效果子菜单中。

![DissolveZap](../_static/DissolveZap.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始转场。

- **Background**: 默认为无。用于与溶解合成的素材。如果未提供背景，则源也用作背景。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Transition Dir** (Popup menu, Default: Dissolve Off to Bg)
  选择转场的方向。
  - **Dissolve Off to Bg**: 从当前图层转场到背景。
  - **Dissolve On from Bg**: 从背景转场到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  如果启用，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过动画 Dissolve Percent 参数手动执行转场。

- **Dissolve Percent** (Default: 0, Range: 0 to 1)
  必须禁用 Auto Trans 才能使用此参数。它决定前景和背景输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制溶解的时间。

- **Dissolve Speed** (Default: 5, Range: 1 or greater)
  From 和 To 素材之间的溶解速度。设为 1 时，溶解在效果的整个持续时间内进行。设为更高值时，溶解时间更短，但闪电仍在整个持续时间内改变大小和亮度。设为 10 可使转场更快捷，更像闪帧切换。

- **Max Bolts** (Integer, Default: 35, Range: 1 to 500)
  转场中点处闪电的最大数量。

- **Start** (X & Y, Default: [0 0], Range: any)
  闪电的起始点。

- **End** (X & Y, Default: [0 0], Range: any)
  闪电的终点。可以使用 End 控件调整此参数。

- **Vary Endpoint** (Default: 1.4, Range: 0 or greater)
  在此半径的圆内随机偏移终点位置。如果 Max Bolts 大于 1，这可以用来分散不同的终点。

- **Bolt Width** (Default: 0.112, Range: 0 or greater)
  闪电的宽度。

- **Branchiness** (Default: 5, Range: 0 to 20)
  缩放从主闪电分支出的附加闪电数量。设为 0 可获得没有额外分支的基本闪电。

- **Zap Bright** (Default: 1, Range: 0 or greater)
  缩放闪电的亮度。

- **Zap Color** (Default rgb: [1 1 1])
  闪电的颜色。如果想保持闪电本身的亮白色，可以通过调整 Glow Color 来影响感知颜色。

- **Zap Glow Bright** (Default: 2, Range: 0 or greater)
  缩放应用于闪电的辉光亮度。

- **Zap Glow Color** (Default rgb: [0.5 0.5 1])
  应用于闪电的辉光颜色。

- **Zap Glow Width** (Default: 0.224, Range: 0 or greater)
  应用于闪电的辉光宽度。

- **Bg Glow Bright** (Default: 8, Range: 0 or greater)
  缩放转场中点处背景辉光的亮度。

- **Bg Glow Color** (Default rgb: [1 1 1])
  缩放转场中点处背景辉光的颜色。辉光的颜色和亮度也受输入影响。

- **Bg Glow Width** (Default: 0.4, Range: 0 or greater)
  缩放转场中点处背景辉光的距离。请注意，零辉光宽度仍会增强明亮区域；如果不需要背景辉光，请将亮度参数设为零。

- **Start Offset** (Default: 0, Range: 0 or greater)
  从起始点开始绘制闪电的偏移量。这对于动画化闪电打击效果很有用。

- **Length** (Default: 1, Range: 0 or greater)
  从 Start Offset 开始的闪电长度。如果小于 1，闪电将不会从起点完整绘制到终点。这对于动画化闪电打击效果很有用。

- **Rand Seed** (Default: 0, Range: 0 or greater)
  用于初始化随机数生成器。实际种子值并不重要，但不同的种子会给出不同的随机闪电，相同的值应给出可重复的结果。

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  如果此值为正，输出的 Alpha 通道将包含来自闪电及其辉光的部分不透明度。红、绿、蓝亮度的最大值按此值缩放，并在每个像素处与背景 Alpha 合成。

- **Show Start** (Check-box, Default: on)
  开启或关闭用于调整 Start 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。

- **Show Vary Endpoint** (Check-box, Default: on)
  开启或关闭用于调整 End 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
