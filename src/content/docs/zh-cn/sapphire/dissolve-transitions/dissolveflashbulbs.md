---
title: DissolveFlashbulbs
---

## S_DissolveFlashbulbs

在两个素材之间溶解的同时模拟大量闪光灯闪烁。使用许多小闪光时，看起来像体育场场景。使用少量大闪光时，适合名人红毯素材。

在 Sapphire Transitions 效果子菜单中。

![DissolveFlashbulbs](../_static/DissolveFlashbulbs.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始转场。

- **Background**: 默认为无。以此素材结束转场。


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

- **Dissolve Speed** (Default: 3, Range: 1 or greater)
  从一个素材到另一个素材的溶解速度。设为 1 时，溶解在效果的整个持续时间内进行。设为更高值时，溶解更短，但闪光灯的渐入和渐出仍占据整个持续时间。设为 10 可使转场更快捷，更像快速切换。

- **Flash Style** (Default: 0, Range: 0 or greater)
  要使用的闪光灯风格。有多种风格可供选择，也可以尝试一些眩光以获得不同的外观。

- **Max Flashes** (Default: 20, Range: 0 or greater)
  溶解中间每帧的最大闪光数量。

- **Flash Randomness** (Default: 0.2, Range: 0 to 1)
  增大可使某些帧产生更多闪光（最多达到 Flashes 的值），而另一些帧产生更少闪光。

- **Flash Size** (Default: 0.4, Range: 0 or greater)
  闪光的平均大小。

- **Flash Rel Height** (Default: 1, Range: 0 or greater)
  用于压缩或拉伸闪光。

- **Brightness** (Default: 3, Range: 0 or greater)
  闪光的整体亮度。

- **Vary Brightness** (Default: 0.2, Range: 0 to 1)
  增大可使每帧中每个闪光灯的亮度产生变化。

- **Flash Gamma** (Default: 1, Range: 0.1 or greater)
  增亮或变暗闪光的中间调。可以产生圆润、硬边的外观，或使闪光更加柔和和微妙。

- **Hold Frames** (Integer, Default: 1, Range: 0 or greater)
  每次闪光会随时间略微衰减，以模拟视觉暂留以及老式闪光灯灯丝冷却的效果。Hold Frames 控制该拖尾持续的时间。

- **Flash Decay Rate** (Default: 0.1, Range: 0 to 1)
  闪光在 Hold Frames 时间内衰减的速度。增大可使其在屏幕上保持更亮、更长时间；减小可使其快速消失。注意可能需要增加 Hold Frames 才能看到长时间的闪光拖尾。

- **Combine** (Popup menu, Default: Add)
  决定闪光图像如何与背景合成。
  - **Screen**: 将闪光与背景混合，有助于防止结果过亮。
  - **Add**: 将闪光图像添加到背景上。

- **Seed** (Default: 0.1, Range: 0 or greater)
  用于初始化随机数生成器。实际种子值本身不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  如果此值为正，输出的 Alpha 通道将包含来自闪光的一些不透明度。红、绿、蓝闪光亮度的最大值按此值缩放，并在每个像素处与背景 Alpha 合成。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Flip Vertically** (Check-box, Default: off)
  如需保持一致的外观，可垂直翻转闪光。

- **Show Flash Size** (Check-box, Default: on)
  开启或关闭用于调整 Flash Size 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
