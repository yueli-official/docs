---
title: FlickerMatchColor
---

## S_FlickerMatchColor

使用第二个匹配片段的颜色变化为源片段添加颜色变化。与 FlickerMatch 类似，但该过程应用于每个颜色通道。要使用此效果，首先将矩形的角定位到匹配片段中具有您想要复制的颜色变化的区域。中灰或浅灰区域最适合此操作。然后选择一帧您希望源颜色保持不变的帧，并点击 Set Match Level 按钮。处理其他帧时，源颜色将按矩形内匹配片段的平均颜色相对于匹配颜色进行缩放。

在 Sapphire Time 效果子菜单中。

![FlickerMatchColor](../_static/FlickerMatchColor.jpg)


### Inputs:

- **Source**: 当前图层。要添加颜色变化的片段。

- **Match**: 默认为无。要从中复制颜色变化的片段。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Rect Corner1** (X & Y, Default: [-0.583 -0.441], Range: any)
  用于测量闪烁的矩形的左上角，以屏幕坐标表示。

- **Rect Corner2** (X & Y, Default: [0.583 0.441], Range: any)
  用于测量闪烁的矩形的右下角，以屏幕坐标表示。

- **Match Color** (Default rgb: [0.5 0.5 0.5])
  矩形内匹配片段的平均颜色，在此颜色下源输入保持不变。

- **Set Match Color** (Push-button)
  按下此按钮会将 Match Color 参数设置为当前帧矩形内匹配片段的平均颜色。这会使输出在此帧等于源。此按钮本身不保留任何值，按下后会立即恢复关闭状态。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Show Rect** (Check-box, Default: on)
  开启或关闭用于调整矩形角参数的屏幕用户界面控件。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
