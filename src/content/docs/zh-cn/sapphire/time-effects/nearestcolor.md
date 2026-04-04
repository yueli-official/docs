---
title: NearestColor
---

## S_NearestColor

从输入素材的帧中收集最接近给定 Match Color 的像素颜色。例如，这可以从在蓝屏或绿屏背景上有对象移动的素材中创建仅背景的图像。它也可以用于在非着色背景上积累运动对象的颜色。每当处理任何非连续帧时，收集的颜色会重新初始化：包括第一帧、重新处理某一帧或跳转到另一帧。您必须连续处理素材的多个帧才能观察到效果，渲染前清除图像缓存有时可能是必要的。

在 Sapphire Time effects 子菜单中。

![NearestColor](../_static/NearestColor.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Steps** (Integer, Default: 15, Range: 2 or greater)
  调整从中收集颜色的输入帧数。

- **Match Color** (Default rgb: [0 0 1])
  保留"最接近"此颜色的像素颜色。

- **Chroma Weight** (Default: 1, Range: 0 or greater)
  色相对颜色匹配的影响程度。如果为 0，将保留与 Match Color 亮度最接近的像素；如果为 2，色相将有更大的影响而亮度的影响减小。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按照图像已为预乘形式（颜色已按不透明度缩放）进行处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将为预乘形式，有时不太准确。

