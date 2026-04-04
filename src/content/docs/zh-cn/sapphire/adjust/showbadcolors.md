---
title: ShowBadColors
---

## S_ShowBadColors

识别所有超出给定颜色范围的像素，并用相同的颜色标记它们以便于查看。

在 Sapphire Adjust 效果子菜单中。

![ShowBadColors](../_static/ShowBadColors.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Min** (Default: 0, Range: 0 to 1)
  最小颜色值。任何颜色通道小于此值的像素将被标记为 Low Color。

- **Max** (Default: 1, Range: 0 to 1)
  最大颜色值。任何颜色通道大于此值的像素将被标记为 High Color。

- **Min Luma** (Default: 0, Range: 0 to 1)
  最小亮度值。亮度小于此值的像素将被标记为 Low Color。

- **Max Luma** (Default: 1, Range: 0 to 1)
  最大亮度值。亮度大于此值的像素将被标记为 High Color。

- **Min Chroma** (Default: 0, Range: 0 to 1)
  最小色度值。色度小于此值的像素将被标记为 Low Color。

- **Max Chroma** (Default: 1, Range: 0 to 1)
  最大色度值。色度大于此值的像素将被标记为 High Color。

- **Min Rgb** (Default rgb: [0 0 0])
  每个颜色通道的最小值。任何颜色通道低于此参数对应通道的像素将被标记为 Low Color。

- **Max Rgb** (Default rgb: [1 1 1])
  每个颜色通道的最大值。任何颜色通道高于此参数对应通道的像素将被标记为 High Color。

- **High Color** (Default rgb: [1 0 0])
  用于标记高值像素的颜色。任何高于 Max 参数之一的像素将被设置为此颜色。

- **Low Color** (Default rgb: [0 0 1])
  用于标记低值像素的颜色。任何低于 Min 参数之一的像素将被设置为此颜色。

- **Output Matte** (Check-box, Default: off)
  如果启用，输出一个遮罩，其中不良像素为白色，其他像素为黑色。

- **Invert Matte** (Check-box, Default: off)
  如果启用，遮罩被反转，不良像素显示为黑色，其他像素显示为白色。除非同时启用了 Output Matte，否则无效。
