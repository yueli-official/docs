---
title: SwishPan
---

## S_SwishPan

通过将一个素材滑出画面、另一个素材滑入，并叠加运动模糊以模拟快速摇镜，实现两个输入素材之间的过渡。该效果在过渡时长较短时效果最佳。

位于 Sapphire Transitions 效果子菜单中。

![SwishPan](../_static/SwishPan.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始过渡。

- **Background**: 默认为无。以此素材结束过渡。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Transition Dir** (Popup menu, Default: Wipe Off to Bg)
  选择过渡方向。
  - **Wipe Off to Bg**: 从当前图层过渡到 Background。
  - **Wipe On from Bg**: 从 Background 过渡到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  若启用，将在图层的首帧与末帧之间自动执行一次过渡。关闭时，需要通过动画 Swish Percent 手动执行过渡。

- **Wipe Percent** (Default: 0, Range: 0 to 1)
  仅在关闭 Auto Trans 时生效。决定 From 与 To 两个输入之间的过渡比例。通常将其从 0 动画到 1 以完成一次完整的过渡。可通过曲线精细控制擦除节奏。

- **Direction** (Popup menu, Default: Left)
  过渡期间素材的移动方向。
  - **Left**: 从右向左移动。
  - **Right**: 从左向右移动。
  - **Up**: 向上移动。
  - **Down**: 向下移动。

- **Blur Amount** (Default: 2, Range: 0 or greater)
  使用的运动模糊量。若方向为左右，则为水平模糊；若方向为上下，则为垂直模糊。

- **Overlap** (Default: 0, Range: any)
  两个素材的重叠量。重叠区域将以 Screen 方式叠加，这有助于消除不良边缘。

- **Slow In** (Default: 0.5, Range: 0 to 1)
  若为正，使过渡起始更为平缓。

- **Slow Out** (Default: 0.5, Range: 0 to 1)
  若为正，使过渡结束更为平缓。

- **Opacity** (Popup menu, Default: Normal)
  处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明（alpha=1）时渲染略快。
  - **Normal**: 正常处理透明度。
  - **As Premult**: 按已预乘形式处理（颜色已按不透明度缩放），渲染略快，但结果也将是预乘形式，某些情况下 Normal 模式更佳。如果图像颜色与遮罩边缘都很锐利，Normal 模式可能更好。
