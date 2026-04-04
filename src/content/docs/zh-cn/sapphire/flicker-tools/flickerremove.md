---
title: FlickerRemove
---

## S_FlickerRemove

从源片段中移除时间上的闪烁。例如，曝光时间不均匀的旧素材可以通过此效果进行平滑处理。要使用此效果，首先将矩形的角定位到平均亮度应保持恒定的区域。中灰或浅灰区域最适合此操作。然后选择矩形内具有所需亮度的源帧，并点击 Set Hold Level 按钮。处理其他帧时，其亮度将被缩放，使矩形内的平均亮度等于保持级别。您可以随时间对不同的 Hold Level 值设置关键帧，以适应所需的亮度变化。

在 Sapphire Time 效果子菜单中。

![FlickerRemove](../_static/FlickerRemove.jpg)


### Inputs:

- **Source**: 当前图层。要移除闪烁的片段。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Rect Corner1** (X & Y, Default: [-0.583 -0.441], Range: any)
  用于测量闪烁的矩形的左上角，以屏幕坐标表示。

- **Rect Corner2** (X & Y, Default: [0.583 0.441], Range: any)
  用于测量闪烁的矩形的右下角，以屏幕坐标表示。

- **Hold Level** (Default: 0.5, Range: 0.01 or greater)
  矩形内区域所请求的平均输出亮度。

- **Set Hold Level** (Push-button)
  按下此按钮会将 Hold Level 参数设置为当前帧矩形内源片段的平均亮度。这会使输出在此帧等于源。此按钮本身不保留任何值，按下后会立即恢复关闭状态。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Show Rect** (Check-box, Default: on)
  开启或关闭用于调整矩形角参数的屏幕用户界面控件。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
