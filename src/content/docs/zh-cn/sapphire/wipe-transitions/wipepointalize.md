---
title: WipePointalize
---

## S_WipePointalize

通过以半随机顺序将类似笔刷的多边形图形从一个素材“点彩”到另一个素材上，实现两个输入素材之间的过渡。应对 Wipe Percent 进行动画以控制过渡速度。通过调整 Frequency 来改变形状大小，并调整 Edge Width 与 Chunky 以获得不同的图案风格。

位于 Sapphire Transitions 效果子菜单中。

![WipePointalize](../_static/WipePointalize.jpg)


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
  若启用，将在图层的首帧与末帧之间自动执行一次过渡。关闭时，需要通过动画 Wipe Percent 手动控制过渡。

- **Wipe Percent** (Default: 0, Range: 0 to 1)
  仅在关闭 Auto Trans 时生效。决定 From 与 To 两个输入之间的过渡比例。通常将其从 0 动画到 1 以完成一次完整的过渡。可通过曲线精细控制擦除节奏。

- **Edge Width** (Default: 2, Range: 0.05 or greater)
  过渡区域的宽度。

- **Angle** (Default: 0, Range: any)
  擦除方向角度（度）。0 为自左向右；90 或 -90 为竖直擦除；180 为自右向左。

- **Frequency** (Default: 20, Range: 5 or greater)
  增大得到更多且更细小的多边形；减小得到更少且更大的多边形。

- **Chunky** (Default: 0, Range: 0 or greater)
  增大使形状以更成团的顺序加入。

- **Stroke Length** (Default: 0, Range: any)
  控制笔刷笔触形状的长度。0 为规则多边形；增大得到更长且更随机的笔触；负值会反向定向。注意当该值非 0 时，笔触会如同被重新“绘制”而随时间变化。

- **Stroke Align** (Default: 0.5, Range: 0 or greater)
  增大以平滑笔触方向，使相邻笔触更加平行。

- **Seed** (Default: 0.23, Range: 0 or greater)
  随机种子。不同种子得到不同结果，相同种子应得到可复现的结果。

- **Opacity** (Popup menu, Default: Normal)
  处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明（alpha=1）时渲染略快。
  - **Normal**: 正常处理透明度。
  - **As Premult**: 按已预乘形式处理（颜色已按不透明度缩放），渲染略快，但结果也将是预乘形式，精确性可能略差。

- **Show Wipe** (Check-box, Default: on)
  打开或关闭用于调整 Grad Add、Grad Angle 与 Wipe Percent 的屏幕控件。须先将 Grad Add 设为正值以显示该控件。此参数仅在支持屏幕控件的 AE 与 Premiere 中出现。
