---
title: FieldRemove
---

## S_FieldRemove

自适应地从运动区域移除视频场隔行扫描伪影，而不模糊图像的静止部分。内部会生成一个"运动遮罩"，运动区域以通常的垂直分辨率损失进行去隔行处理，但静止区域不会被去隔行处理，应保持清晰。

在 Sapphire Time effects 子菜单中。

![FieldRemove](../_static/FieldRemove.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。在结果和源素材输入之间进行插值。白色区域使用效果结果。黑色区域使用源素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: Same Speed)
  选择速度变换选项。
  - **Same Speed**: 不改变速度。
  - **NTSC to Film**: 将 60 场/秒的输入转换为 24 帧/秒的输出。每 5 帧输入转换为 4 帧输出，因此在此模式下只有 4/5 的输出素材是有用的。
  - **Half Speed**: 每个输入场转换为一帧输出。在此模式下，您通常应先填充输入素材使其长度加倍，以便生成正确数量的输出帧。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此数值模糊 Mocha 遮罩。可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 遮罩的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  如果启用，Mocha 遮罩的黑白将在应用效果之前反转。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素量扩展或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下扩展 Mocha 遮罩以进行快速调整。
  - **High**: 在高质量模式下扩展 Mocha 遮罩以获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩同时提供给效果时，决定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Scale Mo Matte** (Default: 4, Range: 0 or greater)
  增大以移除更多场伪影，或减小以移除较少并保持图像更清晰。

- **Threshold Matte** (Default: 0.05, Range: 0 or greater)
  从运动遮罩中减去此值，可增大以减少仅由噪声引起的不必要去隔行处理。

- **Blur Mo Matte** (Default: 0.112, Range: 0 or greater)
  确定运动遮罩平滑的程度，以避免隔行和去隔行区域之间的尖锐过渡。

- **Show** (Popup menu, Default: Result)
  选择输出选项。
  - **Result**: 正常输出去隔行结果。
  - **MotionMatte**: 允许查看运动遮罩本身，在调整上述其他参数时会有帮助。

- **Use Field** (Popup menu, Default: Lower)
  选择在有场伪影的区域中保留哪个场。此参数仅在使用 Same Speed 模式时有效。
  - **Lower**: 保留下场。
  - **Upper**: 保留上场。
  - **Merge**: 使用两个场的平均值。

- **Field Dominance** (Popup menu, Default: Lower First)
  选择输出场的顺序。此参数仅在未使用 Same Speed 模式时有效。
  - **Lower First**: 下场在时间上优先。
  - **Upper First**: 上场在时间上优先。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按照图像已为预乘形式（颜色已按不透明度缩放）进行处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将为预乘形式，有时不太准确。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用 Mask 输入通道来创建单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊遮罩输入。这可以在遮罩区域和非遮罩区域之间提供更平滑的过渡。除非提供了遮罩输入，否则不起作用。

- **Invert Mask** (Check-box, Default: off)
  如果启用，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则不起作用。

