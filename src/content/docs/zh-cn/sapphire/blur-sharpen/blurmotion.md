---
title: BlurMotion
---

## S_BlurMotion

在指定的 From 和 To 变换之间对源素材执行运动模糊。这可用于执行径向缩放模糊、旋转模糊、方向模糊或这些的任意组合。From 和 To 参数不指时间。它们描述空间中的两个变换，决定应用于每帧的模糊样式。

位于 Sapphire Blur+Sharpen 效果子菜单中。

![BlurMotion](../_static/BlurMotion.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Matte**: 默认为无。如果提供，每个目标像素的运动模糊量将按此输入进行缩放。此输入可通过 Blur Matte、Invert Matte 或 Matte Use 参数进行调整。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: Blur Color)
  在全色或单色结果之间选择。
  - **Blur Color**: 模糊源输入的所有通道。
  - **Blur Mono**: 先将源素材转为单色，然后模糊生成的单通道（更快）。

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

- **Center** (X & Y, Default: [0 0], Range: any)
  旋转和缩放的中心，以相对于帧中心的屏幕坐标表示。Shift 值应为零时此位置才有意义。此参数可通过 Center Widget 调整。

- **From Z Dist** (Default: 1, Range: 0.001 or greater)
  From 变换的"距离"。当 Shift 为 0 时，这围绕 Center 位置进行缩放。增大可缩小，减小可放大。此参数可通过 From Transfm Widget 调整。

- **From Rotate** (Default: 0, Range: any)
  From 变换的旋转角度（以度为单位），围绕中心。此参数可通过 From Transfm Widget 调整。

- **From Shift** (X & Y, Default: [0 0], Range: any)
  From 变换的水平和垂直平移。这可用于方向运动。如果不为零，中心位置的意义将减弱。此参数可通过 From Transfm Widget 调整。

- **To Z Dist** (Default: 0.8, Range: 0.001 or greater)
  To 变换的"距离"。增大可缩小，减小可放大。此参数可通过 To Transform Widget 调整。

- **To Rotate** (Default: 0, Range: any)
  To 变换的旋转角度（以度为单位），围绕中心。注意，如果 From 和 To 的旋转角度相差很大，它们之间的插值将变得不太准确。此参数可通过 To Transform Widget 调整。

- **To Shift** (X & Y, Default: [0 0], Range: any)
  To 变换的水平和垂直平移。这可用于方向运动。如果不为零，中心位置的意义将减弱。此参数可通过 To Transform Widget 调整。

- **Exposure Bias** (Default: 0.5, Range: 0 to 1)
  确定 From 和 To 变换之间路径上的可变曝光量。值为 0 在 From 端产生更多曝光，0.5 沿路径产生均匀曝光，1.0 在 To 端产生更多曝光。如果在暗背景上有亮点，值为 0 会使处理后的亮点在 From 端更亮、To 端更暗，值为 1.0 则效果相反。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果的亮度。

- **Wrap** (X & Y, Popup menu, Default: [ No No ])
  确定访问源图像边界外的方法。
  - **No**: 边界外为黑色。
  - **Tile**: 重复图像的副本。
  - **Reflect**: 重复镜像副本。使用此方法时边缘通常不太明显。

- **Blur Res** (Popup menu, Default: Full)
  选择运动模糊的分辨率因子。这类似于通用的 'Res' 因子参数，但在向下平均到较低分辨率并插值回结果方面做得更好。较高的分辨率提供更好的质量，较低的分辨率提供更快的处理速度。
  - **Full**: 使用全分辨率。
  - **Half**: 运动模糊在半分辨率下执行。
  - **Quarter**: 运动模糊在四分之一分辨率下执行。

- **Subpixel** (Check-box, Default: on)
  如果启用，使用质量更好但稍慢的方法执行模糊。

- **Blur Matte** (Default: 0, Range: 0 or greater)
  在使用前按此量模糊 Matte 输入。这可以在遮罩区域和非遮罩区域之间提供更平滑的过渡。除非提供了 Matte 输入，否则无效。

- **Invert Matte** (Check-box, Default: off)
  如果开启，反转 Matte 输入，使效果应用于 Matte 为黑色而非白色的区域。除非提供了 Matte 输入，否则无效。

- **Matte Use** (Popup menu, Default: Luma)
  确定如何使用 Matte 输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且无透明度 (alpha=1) 时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已为预乘形式处理（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时可能不太准确。

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  这 4 个参数 Crop Top、Crop Bottom、Crop Left 和 Crop Right 允许选择要处理的输入图像的矩形子区域。如果 Wrap 参数设为 "No"，暴露的边框将是透明的。如果 Wrap 为 "Tile" 或 "Reflect"，源图像将在新裁剪边框上包裹以填充帧。这可以更容易地避免因扭曲具有不良边缘的图像而产生的伪影。

- **Show Center** (Check-box, Default: on)
  打开或关闭用于调整 Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show From Transfm** (Check-box, Default: on)
  打开或关闭用于调整 From Z Dist 和 From Rotate 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show To Transform** (Check-box, Default: on)
  打开或关闭用于调整 To Z Dist 和 To Rotate 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show From Shift** (Check-box, Default: off)
  打开或关闭用于调整 Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show To Shift** (Check-box, Default: off)
  打开或关闭用于调整 Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

