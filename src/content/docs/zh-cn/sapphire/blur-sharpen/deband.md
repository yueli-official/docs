---
title: Deband
---

## S_Deband

通过在色带区域扩散像素来去除素材中的色带伪影，同时保持原始边缘完整。要使用此效果，首先选择 Show:Edges 并调整边缘阈值，直到色带边缘刚好消失，仅留下所需的真实边缘。然后选择 Show:Result 查看结果。如果仍然看到一些色带，请增大 Diffuse Threshold 和/或 Diffuse Radius。

位于 Sapphire Blur+Sharpen 效果子菜单中。

![Deband](../_static/Deband.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。在结果和源素材输入之间插值。白色区域使用效果结果。黑色区域使用源素材。


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

- **Combine Masks** (Popup menu, Default: Union)
  当同时提供 Mocha 遮罩和输入遮罩时，确定如何组合它们。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，只使用
Mocha 遮罩。

- **Edge Threshold** (Default: 2, Range: 0 to 255)
  相邻像素必须相差的量才能构成真正的所需边缘。值为 1.0 表示 8 位下可能的最小差异。此参数应设置得足够高，使色带不会显示为边缘，但又足够低，以便仍能检测到所有真实边缘。

- **Grow Edges** (Default: 0, Range: 0 or greater)
  以近似像素为单位扩展检测到的边缘。增大此参数可防止在边缘附近但不在边缘上的区域发生扩散。

- **Show** (Popup menu, Default: Result)
  选择输出类型。
  - **Result**: 显示最终结果。
  - **Edges**: 显示图像的边缘，即相邻像素差异大于 Edge Threshold 的位置。使用此模式可帮助微调边缘检测参数。

- **Diffuse Threshold** (Default: 1, Range: 0 or greater)
  扩散像素时允许的最大颜色差异。此参数由边缘阈值自动缩放。增大此值可在色带内有渐变时获得更好的结果。减小此值可减少无边缘区域中的扩散。

- **Diffuse Radius** (Default: 12, Range: 0 or greater)
  像素扩散的最大半径（以近似像素为单位）。较大的值可以更有效地去除大面积均匀颜色区域中的色带，而较小的值在具有许多小颜色区域的地方效果更好。

- **Pre Blur** (Default: 0, Range: 0 or greater)
  在扩散像素前模糊源素材。

- **Post Blur** (Default: 0.5, Range: 0 or greater)
  在扩散像素后模糊结果。使用此参数可减少结果中的噪点。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且无透明度 (alpha=1) 时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已为预乘形式处理（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时可能不太准确。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用 Mask 输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此量模糊 Matte 输入。这可以在遮罩区域和非遮罩区域之间提供更平滑的过渡。除非提供了 Matte 输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果开启，反转 Matte 输入，使效果应用于 Matte 为黑色而非白色的区域。除非提供了 Matte 输入，否则无效。

