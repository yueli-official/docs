---
title: Swish3D
---

## S_Swish3D

在两个输入素材之间进行溶解过渡的同时，对每个素材施加 3D 运动。在过渡期间，From 素材将根据 Zdist、Rotate、Swivel、Tilt、Shift、Scale 和 Shear 参数进行变换，To 素材则按这些参数的相反值进行变换。每个图像整体运动的强度可通过 Rel Amp From 与 Rel Amp To 参数进行缩放。

位于 Sapphire Transitions 效果子菜单中。

![Swish3D](../_static/Swish3D.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始过渡。

- **Background**: 默认为无。以此素材结束过渡。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: Blur Warp)
  选择在移动 From 与 To 素材时应用的运动模糊类型。
  - **Blur Warp**: 普通的运动模糊，类似 BlurMotion 效果。
  - **Chroma Warp**: 让各颜色通道以不同幅度移动，产生类似 WarpChroma 的彩色边缘效果。

- **Transition Dir** (Popup menu, Default: Wipe Off to Bg)
  选择过渡方向。
  - **Wipe Off to Bg**: 从当前图层过渡到 Background。
  - **Wipe On from Bg**: 从 Background 过渡到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  若启用，将在图层的首帧与末帧之间自动执行一次过渡。关闭时，需要通过动画 Swish3 Percent 参数手动执行过渡。

- **Wipe Percent** (Default: 0, Range: 0 to 1)
  仅在关闭 Auto Trans 时生效。决定 From 与 To 两个输入之间的过渡比例。通常将其从 0 动画到 1 以完成一次完整的过渡。可通过曲线精细控制擦除节奏。

- **Center** (X & Y, Default: [0 0], Range: any)
  d 中心在屏幕坐标中的位置（相对于画面中心）。可通过开启并拖动 Center Widget 设置。注意：移动 d 中心也可能改变 d 的大小，以保证当前 Wipe Amt 的值仍然正确。

- **Motion Blur** (Default: 1, Range: 0 or greater)
  缩放所用的运动模糊量。

- **Z Dist** (Default: 0.5, Range: 0.001 or greater)
  对 From 素材应用的“距离”变换。大于 1.0 将其推远并变小；小于 1.0 将其拉近并放大。默认情况下，To 素材会按该值的相反方向进行变换。

- **Rotate** (Default: 0, Range: any)
  以度数指定旋转角。

- **Swivel** (Default: 0, Range: any)
  围绕竖直轴在 3D 中向左/向右旋转。

- **Tilt** (Default: 0, Range: any)
  围绕水平轴在 3D 中向上/向下旋转。可与 Swivel 组合以围绕任意对角轴旋转。

- **Perspective Amount** (Default: 1, Range: 0.25 to 4)
  控制应用 Swivel 与 Tilt 时的透视拉伸强度。增大以获得更强的 3D 透视感。

- **Shift** (X & Y, Default: [0 0], Range: any)
  d 模式的平移。

- **Scale** (Default: 1, Range: 0 to 2)
  缩放素材大小。

- **Scale Rel** (X & Y, Default: [1 1], Range: 0 to 2)
  缩放素材的相对水平或垂直尺寸。

- **Shear** (X & Y, Default: [0 0], Range: any)
  水平或垂直方向剪切。

- **Rel Amp From** (Default: 1, Range: any)
  缩放对 From 素材应用的变换量。设为 0 以禁用移动 From 素材；设为负值可反向运动。

- **Rel Amp To** (Default: -1, Range: any)
  缩放对 To 素材应用的变换量。默认 To 素材的变换方向与 From 相反。设为 0 以禁用移动 To 素材；设为正值可使 To 与 From 同向移动。

- **Fade** (Popup menu, Default: From and To)
  决定在过渡期间哪些素材参与淡入/淡出。
  - **From and To**: 交叉淡入淡出两个素材。
  - **Only From**: 仅淡出 From，并将其与 To 合成。这会使 To 在 From 未覆盖的区域保持完全不透明。
  - **Only To**: 仅淡入 To，并将其与 From 合成。这会使 From 在 To 未覆盖的区域保持完全不透明。

- **Fade Mid Time** (Default: 0.5, Range: 0 to 1)
  溶解的时间中点。减小为更早的溶解；增大为更晚的溶解。若该值为 1.0，From 将在整个过渡期间保持完全不透明。可与 Combine 参数配合制作各种显隐而无需淡变，例如将 Dissolve Mid Time 设为 1.0、Combine 设为 Fade From，并通过 Shift/Rotate 使 From 移出画面。

- **Slow In** (Default: 0.5, Range: 0 to 1)
  若为正，使过渡起始更为平缓。

- **Slow Out** (Default: 0.5, Range: 0 to 1)
  若为正，使过渡结束更为平缓。

- **Wrap From** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  确定访问 From 图像边界外的方式。
  - **No**: 边界外为黑色。
  - **Tile**: 重复平铺图像。
  - **Reflect**: 镜像重复图像。通常该方式边缘更不明显。

- **Wrap To** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  确定访问 To 图像边界外的方式。
  - **No**: 边界外为黑色。
  - **Tile**: 重复平铺图像。
  - **Reflect**: 镜像重复图像。通常该方式边缘更不明显。

- **Filter** (Check-box, Default: on)
  若启用，图像在重采样时会自适应滤波。当图像被缩小或形变变小，能获得更好的质量。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果亮度。可在过渡期间做动画以提亮，但通常起止应为 1.0 以避免过渡开始或结束时的跳变。

- **Mid Brightness** (Default: 1, Range: 0 or greater)
  在过渡中点按该量缩放结果亮度。在过渡过程中会自动渐变到此亮度并返回。

- **Steps** (Integer, Default: 8, Range: 3 to 100)
  在 From（红）与 To（蓝）变换路径之间采样的光谱步数。步数越多越平滑，但处理时间更长。

- **Color1** (Default rgb: [1 0 0])
  From 变换处的颜色。

- **Color2** (Default rgb: [0 1 0])
  From 与 To 变换之间中点处的颜色。

- **Color3** (Default rgb: [0 0 1])
  To 变换处的颜色。

- **White Balance** (Check-box, Default: off)
  若启用，内部会调整三色使其相加为白色。此时未形变区域的颜色不受影响，结果的平均颜色保持不变。

- **Opacity** (Popup menu, Default: Normal)
  处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明（alpha=1）时渲染略快。
  - **Normal**: 正常处理透明度。
  - **As Premult**: 按已预乘形式处理（颜色已按不透明度缩放），渲染略快，但结果也将是预乘形式，某些情况下 Normal 模式更佳。如果图像颜色与遮罩边缘都很锐利，Normal 模式可能更好。

- **Show To Shift** (Check-box, Default: off)
  打开或关闭用于调整 Center 的屏幕控件。此参数仅在 AE 与 Premiere 中出现（支持屏幕控件）。

- **Show To Transform** (Check-box, Default: on)
  打开或关闭用于调整 To Z Dist 与 To Rotate 的屏幕控件。仅在 AE 与 Premiere 中出现。

- **Show From Shift** (Check-box, Default: off)
  打开或关闭用于调整 Center 的屏幕控件。仅在 AE 与 Premiere 中出现。

- **Show From Transform** (Check-box, Default: on)
  打开或关闭用于调整 Center 的屏幕控件。仅在 AE 与 Premiere 中出现。

- **Show Center** (Check-box, Default: on)
  打开或关闭用于调整 Center 的屏幕控件。仅在 AE 与 Premiere 中出现。
