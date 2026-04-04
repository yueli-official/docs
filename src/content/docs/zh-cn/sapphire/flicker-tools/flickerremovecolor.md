---
title: FlickerRemoveColor
---

## S_FlickerRemoveColor

从源片段中移除时间上的颜色变化。与 FlickerRemove 类似，但该过程应用于每个颜色通道。要使用此效果，首先将矩形的角定位到平均颜色应保持恒定的区域。中灰或浅灰区域最适合此操作。然后选择矩形内具有所需颜色的源帧，并点击 Set Hold Color 按钮。处理其他帧时，其颜色将被缩放，使矩形内的平均颜色等于保持颜色。

在 Sapphire Time 效果子菜单中。

![FlickerRemoveColor](../_static/FlickerRemoveColor.jpg)


### Inputs:

- **Source**: 当前图层。要移除颜色变化的片段。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Rect Corner1** (X & Y, Default: [-0.583 -0.441], Range: any)
  用于测量闪烁的矩形的左上角，以屏幕坐标表示。

- **Rect Corner2** (X & Y, Default: [0.583 0.441], Range: any)
  用于测量闪烁的矩形的右下角，以屏幕坐标表示。

- **Hold Color** (Default rgb: [0.5 0.5 0.5])
  矩形内区域所请求的平均输出颜色。

- **Set Hold Color** (Push-button)
  按下此按钮会将 Hold Color 参数设置为当前帧矩形内源片段的平均颜色。这会使输出在此帧等于源。此按钮本身不保留任何值，按下后会立即恢复关闭状态。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Show Rect** (Check-box, Default: on)
  开启或关闭用于调整矩形角参数的屏幕用户界面控件。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
