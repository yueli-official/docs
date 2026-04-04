---
title: ZFogExponential
---

## S_ZFogExponential

使用 ZBuffer 输入的深度值将雾色混合到源素材中。雾从 Z Near 开始，根据 Fog Density 以指数方式增加到 Z Far。如果未提供 ZBuffer 输入，它将为纯黑色，因此应为此效果指定该输入以使其产生实际效果。

在 Sapphire Stylize 效果子菜单中。

![ZFogExponential](../_static/ZFogExponential.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **ZBuffer**: 默认为无。包含每个源像素深度值的输入素材。这些值应在黑色到白色的范围内，最好不要进行抗锯齿处理。通常黑色对应最远的物体，白色对应最近的物体，但这可以通过 Z Buffer 参数进行调整。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Fog Density** (Default: 0.7, Range: 0 to 1)
  雾的密度。

- **Fog Color** (Default rgb: [0.5 0.5 0.5])
  雾的颜色通常应与源素材的天空或背景颜色匹配。雾气用灰色，烟雾用棕色，水下效果用蓝色等。

- **Z Buffer Type** (Popup menu, Default: White is Near)
  如何解释 Z 缓冲区中的值。
  - **Black is Near**: Z 缓冲区中的黑色像素表示该点的物体较近（靠近您），白色表示较远。
  - **White is Near**: Z 缓冲区中的白色像素表示该点的物体较近（靠近您），黑色表示较远。

- **Z Buffer Use** (Popup menu, Default: Luma)
  决定如何使用 ZBuffer 输入通道生成单色 Z 图像。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。
