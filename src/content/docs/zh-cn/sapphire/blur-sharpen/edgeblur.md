---
title: EdgeBlur
---

## S_EdgeBlur

查找 Matte 素材中的边缘，并在这些边缘处模糊源素材。使用 Show Edges 选项可查看在调整边缘参数时哪些区域将接受模糊。然后调整 Blur Width 以控制模糊量。

位于 Sapphire Blur+Sharpen 效果子菜单中。

![EdgeBlur](../_static/EdgeBlur.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Edge_Source**: 默认为无。用于确定源素材应被模糊的边缘位置的素材。如果未连接此输入，则使用主源素材来确定边缘。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此量模糊 Mocha 遮罩。可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 遮罩的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  如果启用，在应用效果前反转 Mocha 遮罩的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素量膨胀或腐蚀 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下膨胀 Mocha 遮罩，便于快速调整。
  - **High**: 在 High 质量模式下膨胀 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，只显示 Mocha 遮罩本身。

- **Blur Width** (Default: 0.112, Range: 0 or greater)
  模糊的宽度。通常不应大于 Edge Width。此参数可通过 Blur Width Widget 调整。

- **Edge Width** (Default: 0.112, Range: 0 or greater)
  要在其中进行模糊的边缘区域宽度。

- **Edge Strength** (Default: 0.5, Range: 0 or greater)
  边缘的强度决定了替换边缘的模糊源素材的量。

- **Edge Threshold** (Default: 0, Range: 0 or greater)
  确定哪些边缘被模糊。增大可去除次要边缘或斑点。

- **Show** (Popup menu, Default: Result)
  在输出选项之间选择。
  - **Result**: 输出带有模糊边缘的源图像。
  - **Edges**: 仅输出边缘图像。在调整边缘参数时很有用。

- **Subpixel** (Check-box, Default: on)
  启用亚像素量的模糊。使用此选项可使 Blur Width 或 Edge Width 参数的动画更平滑。

- **Matte Use** (Popup menu, Default: Luma)
  确定如何使用 Matte 输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且无透明度 (alpha=1) 时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已为预乘形式处理（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时可能不太准确。

- **Show Blur Width** (Check-box, Default: on)
  打开或关闭用于调整 Blur Width 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

