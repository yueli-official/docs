---
title: Retime
---

## S_Retime

使用基于光流的运动估计和帧插值对素材进行变速处理。

在 Sapphire Time effects 子菜单中。

![Retime](../_static/Retime.jpg)


### Inputs:

- **Source**: 当前图层。要进行变速处理的素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: Retime Speed)
  素材变速的方式。
  - **Retime Speed**: 输出素材以输入素材速度的指定比例运行。
  - **Retime Longer**: 输出素材比输入素材长指定的百分比。
  - **Retime Scale**: 输出素材为输入素材长度的指定百分比倍。
  - **Retime FPS**: 输出素材速度适合在指定的帧率下播放，给定输入素材以指定帧率的速度。
  - **Retime Length**: 使用光流运动估计对素材进行变速，使输出素材为指定长度（以帧为单位）。
  - **Retime Variable Speed**: 使用光流运动估计对素材进行变速，使每个输出帧的输入帧可通过关键帧参数控制。这允许变速效果、倒放运动等。

- **Speed** (Default: 0.5, Range: 0.01 to 10)
  在 Speed 模式下使用。输出素材以此值乘以输入素材速度的速率运行。例如，0.2 将使输出比输入慢 5 倍。

- **Percent Longer** (Default: 100, Range: 0 to 1e+04)
  在 Longer 模式下使用。输出相对于输入要长多少百分比。0% 使输出与输入长度相同。100% 使输出为输入的两倍长。

- **Percent** (Default: 200, Range: 1 to 1e+04)
  在 Scale 模式下使用。输出素材长度占输入素材长度的百分比。100% 使输出与输入长度相同。200% 使输出为输入的两倍长。50% 使输出为输入的一半长。

- **Input Fps** (Default: 30, Range: 1 to 1000)
  在 FPS 模式下使用。输入素材的帧率。

- **Output Fps** (Default: 60, Range: 1 to 1000)
  在 FPS 模式下使用。输出素材的帧率。如果此值大于 Input Fps 参数，输出素材将比输入素材更长。如果此值小于 Input Fps 参数，输出素材将比输入素材更短。如果输出素材以指定的输出帧率播放，其持续时间（秒）将与输入素材相同。

- **Output Length** (Integer, Default: 30, Range: 2 to 1e+04)
  在 Length 模式下使用。所需的输出素材帧长度。

- **Input Frame** (Default: 1, Range: 0 to 1e+04)
  在 Curve 模式下使用。要输出的输入帧。此参数通常应设置关键帧，设置关键帧的帧指定输出帧，该关键帧处设置的值表示在该输出帧处所需的输入帧。

- **Result** (Popup menu, Default: Result)
  输出变速素材或运动向量的表示。
  - **Result**: 输出变速后的素材。
  - **Flow Vectors**: 输出运动向量的表示。像素处向量的方向显示为颜色，向量的长度显示为该颜色的饱和度。

- **Start Frame** (Integer, Default: 0, Range: 0 or greater)
  素材中第一帧的帧号。

- **Motion Blur** (Default: 0, Range: 0 to 2)
  如果大于 0.0，沿运动向量长度的此比例进行模糊。

- **Field Dominance** (Popup menu, Default: Normal)
  对于隔行输入素材，指定场顺序。
  - **Normal**: 使用正常场顺序。
  - **Reverse**: 反转场顺序。

- **Flow Field Smoothness** (Default: 0.8, Range: 0 or greater)
  光流参数：遵循图像数据的相对重要性与保持平滑流场的相对重要性。此值越高，分配给流场平滑度的相对重要性越大。

- **Emphasize Edges** (Default: 0, Range: 0 to 1)
  光流参数：指定在求解运动向量场之前应用于输入素材的结构/纹理分解程度。这有助于减少光照变化引起的问题，并可能提供更好的局部流和其他改进。当输入素材中有快速运动时不建议使用。但在某些情况下，它对获得可接受的结果非常有帮助。建议用法是首先尝试 0.0（关闭结构/纹理分解）。如果结果不满意，尝试 0.4。最后尝试 1.0。

- **Flow View Scale** (Default: 0.9, Range: 0 or greater)
  输出运动向量表示时，这会增加与单位长度运动对应的颜色饱和度。这使表示更灵敏，更清晰地显示较短的向量。

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  这 4 个参数（Crop Top、Crop Bottom、Crop Left 和 Crop Right）允许选择要处理的输入图像的矩形子区域。这可以避免边缘处拉入黑色的伪影。当在更高分辨率的项目中对较低分辨率的元素进行变速时，可能会发生这种情况。

