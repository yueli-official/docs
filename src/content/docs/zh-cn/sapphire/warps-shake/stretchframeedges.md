---
title: StretchFrameEdges
---

## S_StretchFrameEdges

拉伸 4x3 图像的边缘同时保留中心，以隐藏 16x9 合成中的黑色柱状区域。此效果获取源素材的中间部分并将其压缩，因为在 16x9 合成中查看 4x3 图像通常会将其拉伸以适应画面。边缘不会被压缩，因此图像一直延伸到边缘。图像的左右边缘部分会出现水平拉伸。虽然此效果是为 4x3 转换设计的，但它可以适用于任何宽高比。

在 Sapphire Distort 效果子菜单中。

![StretchFrameEdges](../_static/StretchFrameEdges.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Center Squeeze** (Default: 1.33, Range: 0 to 2)
  中心部分图像的压缩量。要将 4x3 图像适配到 16x9 中，通常按 4/3 或 1.333 压缩。

- **Center Width** (Default: 0.4, Range: 0 to 0.99)
  图像的中心部分（如果此参数为 0.5，则为中间一半）被均匀压缩，没有变形。此参数定义图像中不变形的部分。

- **Border Width** (Default: 0, Range: 0 to 1)
  为减少边缘变形，可以增大 Border Width 以允许一些黑色边框。素材边缘的变形不会那么严重，因为它们不需要拉伸那么远。

- **Shift X** (Default: 0, Range: -1 to 1)
  向左或向右移动整个图像，以将图像的有趣部分保持在画面的非变形区域。将此设置为非零值将显示素材的边缘，除非 Wrap 设置为 Tile 或 Reflect。

- **Smooth** (Default: 0.5, Range: 0 to 1)
  设置为零时，图像边缘被线性拉伸。这在最边缘处产生最少的变形，但可能在中心与边缘区域相交处产生可见的接缝。设置为一时，接缝完全隐藏，但图像最边缘处会有相当严重的变形。折中值介于零和一之间。

- **Wrap** (Popup menu, Default: No)
  确定访问源图像边界外区域的方法。
  - **No**: 在边界外显示黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复图像的镜像副本。使用此方法时边缘通常不太明显。

