---
title: FlickerMatch
---

## S_FlickerMatch

使用第二个匹配片段的闪烁为源片段添加闪烁。例如，可以使一个片段的亮度与另一个片段中闪烁的灯光同步变化。要使用此效果，首先将矩形的角定位到匹配片段中具有您想要复制的亮度变化的区域。中灰或浅灰区域最适合此操作。然后选择一帧您希望源亮度保持不变的帧，并点击 Set Match Level 按钮。处理其他帧时，源亮度将按矩形内匹配片段的平均亮度相对于匹配级别进行缩放。

在 Sapphire Time 效果子菜单中。

![FlickerMatch](../_static/FlickerMatch.jpg)


### Inputs:

- **Source**: 当前图层。要添加闪烁的片段。

- **Match**: 默认为无。要从中复制闪烁的片段。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Rect Corner1** (X & Y, Default: [-0.583 -0.441], Range: any)
  用于测量闪烁的矩形的左上角，以屏幕坐标表示。

- **Rect Corner2** (X & Y, Default: [0.583 0.441], Range: any)
  用于测量闪烁的矩形的右下角，以屏幕坐标表示。

- **Match Level** (Default: 0.5, Range: 0.01 or greater)
  矩形内匹配片段的平均亮度，在此亮度下源输入保持不变。

- **Set Match Level** (Push-button)
  按下此按钮会将 Match Level 参数设置为当前帧矩形内匹配片段的平均亮度。这会使输出在此帧等于源。此按钮本身不保留任何值，按下后会立即恢复关闭状态。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Show Rect** (Check-box, Default: on)
  开启或关闭用于调整矩形角参数的屏幕用户界面控件。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
