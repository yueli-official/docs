---
title: Flicker
---

## S_Flicker

按不同的量随时间缩放源片段的颜色，以产生闪烁效果。闪烁的模式可以是随机的、周期性波形的，或两者的组合。

在 Sapphire Time 效果子菜单中。

![Flicker](../_static/Flicker.jpg)


### Inputs:

- **Source**: 当前图层。要处理的片段。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Amplitude** (Default: 0.2, Range: 0 or greater)
  缩放所有闪烁的振幅。

- **Rand Luma Amp** (Default: 1, Range: 0 or greater)
  影响亮度的平滑随机闪烁的振幅。

- **Rand Color Amp** (Default: 0, Range: 0 or greater)
  独立影响各颜色通道的随机闪烁的振幅。

- **Rand Freq** (Default: 30, Range: 0 or greater)
  随机闪烁的频率。增大可使帧间变化更多，减小可使闪烁更慢。

- **Wave Amp** (Default: 0, Range: 0 or greater)
  周期性波形闪烁的振幅。

- **Wave Freq** (Default: 5, Range: 0 or greater)
  波形闪烁的频率。增大可加快闪烁，减小可减慢。如果 Wave Amp 为 0，则此参数无效。

- **Wave R Phase** (Default: 0, Range: any)
  在时间上偏移红色通道的波形模式。

- **Wave G Phase** (Default: 0, Range: any)
  在时间上偏移绿色通道的波形模式。

- **Wave B Phase** (Default: 0, Range: any)
  在时间上偏移蓝色通道的波形模式。

- **Red Amp** (Default: 1, Range: 0 or greater)
  缩放应用于红色通道的闪烁量。

- **Green Amp** (Default: 1, Range: 0 or greater)
  缩放应用于绿色通道的闪烁量。

- **Blue Amp** (Default: 1, Range: 0 or greater)
  缩放应用于蓝色通道的闪烁量。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果的亮度。

- **Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化随机数生成器。实际的种子值并不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。
