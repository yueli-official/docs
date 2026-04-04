---
title: Dissolve
---

## S_Dissolve

两个输入素材之间的基本交叉溶解。

在 Sapphire Transitions 效果子菜单中。

![Dissolve](../_static/Dissolve.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始转场。

- **Background**: 默认为无。以此素材结束转场。如果未提供此输入，将使用完全透明的背景，显示其后面的内容。请注意，除非提供了此输入，否则背景在转场过程中无法被扭曲。


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
  必须禁用 Auto Trans 才能使用此参数。它决定前景和背景输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制溶解的时间。Slow In 和 Slow Out 参数如果为正值，也会在内部调整转场比例，以获得更平滑的转场开始和/或结束。

- **Dissolve Speed** (Default: 1, Range: 1 or greater)
  从一个素材到另一个素材的溶解速度。设为 1 时，溶解在效果的整个持续时间内进行。设为更高值时，溶解更短，但辉光的渐入和渐出仍占据整个持续时间。设为 10 可使转场更快捷，更像闪帧切换。

- **Slow In** (Default: 0.2, Range: 0 to 1)
  如果为正值，使转场开始更加渐进。

- **Slow Out** (Default: 0.2, Range: 0 to 1)
  如果为正值，使转场结束更加渐进。
