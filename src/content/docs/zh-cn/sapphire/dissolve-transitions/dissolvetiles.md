---
title: DissolveTiles
---

## S_DissolveTiles

在两个输入素材之间进行转场，同时将每个素材分解为瓦片并打乱排列。
第一个素材分崩离析并向外扩散，而第二个素材在第一个的背后逐渐凝聚。
应对 Dissolve Percent 参数进行动画处理以控制转场速度。

在 Sapphire Transitions 效果子菜单中。

![DissolveTiles](../_static/DissolveTiles.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始转场。

- **Background**: 默认为无。以此素材结束转场。如果未提供此输入，将使用完全透明的背景，显示其后面的内容。请注意，除非提供此输入，否则背景在转场期间无法被扭曲。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Transition Dir** (Popup menu, Default: Dissolve Off to Bg)
  选择转场的方向。
  - **Dissolve Off to Bg**: 从当前图层转场到背景。
  - **Dissolve On from Bg**: 从背景转场到当前图层。

- **Auto Trans** (Popup YES-NO, Default: No)
  如果启用，将在图层的第一帧和最后一帧之间自动执行转场。如果关闭，则通过动画 Dissolve Percent 参数手动执行转场。

- **Dissolve Percent** (Default: 0, Range: 0 to 1)
  必须禁用 Auto Trans 才能使用此参数。它决定前景和背景输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制溶解的时间。如果 Slow In 和 Slow Out 参数为正值，它们也会在内部调整转场比例，使转场开始和/或结束更平滑。

- **Scramble Speed** (Default: 2, Range: any)
  在转场边缘处每个输入应被打乱的量。传入素材在转场开始时按此量打乱，传出素材在转场结束时按此量打乱。将此值设为零将使两个素材都不产生瓦片效果。

- **Scramble Rel** (X & Y, Default: [1 1], Range: 0 or greater)
  水平和垂直打乱的相对量。

- **Scramble Rel From** (Default: 1, Range: any)
  传出素材中打乱的相对量。如果传出素材不应被打乱，请将此值设为零。

- **Scramble Rel To** (Default: -1, Range: any)
  传入素材中打乱的相对量。如果传入素材不应被打乱，请将此值设为零。

- **Slow In** (Default: 0.2, Range: 0 to 1)
  如果为正值，使转场开始更加缓和。

- **Slow Out** (Default: 0.2, Range: 0 to 1)
  如果为正值，使转场结束更加缓和。

- **Tiles** (Default: 10, Range: 1 or greater)
  图像横向的瓦片数量。增加可获得更多更小的瓦片，减少可获得更少更大的瓦片。

- **Tile Rel Width** (Default: 1, Range: 0.01 or greater)
  缩放每个瓦片的高度。

- **Tile Rel Height** (Default: 1, Range: 0.01 or greater)
  缩放每个瓦片的宽度。

- **Dissolve Delay** (Default: 0.6, Range: 0 to 1)
  在 From 和 To 素材之间交叉溶解之前的延迟。如果设为 1，传出素材完全不淡出。如果设为 0，传出和传入素材在整个转场过程中平滑溶解。

- **Combine** (Popup menu, Default: From Over To)
  默认情况下传出的 From 素材打乱消失，露出下方正在打乱进入的 To 素材。将此设为 To Over From 可使 To 素材打乱进入并覆盖在 From 素材上方。同时调整 Scramble Rel From 和 Scramble Rel To 可获得不错的效果。
  - **From Over To**: 将 From（传出）素材合成在 To（传入）素材上方，随着 From 素材打乱消失而露出 To 素材。在默认设置或将 Scramble Rel To 设为零时效果良好。
  - **To Over From**: 将 To（传入）素材合成在 From（传出）素材上方，将 To 素材打乱叠入 From 素材。在默认设置或将 Scramble Rel From 设为零时效果良好。

- **Rotate Warp Dir** (Default: 0, Range: any)
  将扭曲方向旋转此度数。可以通过动画旋转瓦片来获得有趣的效果。

- **Seed** (Default: 0.5, Range: 0 or greater)
  用于初始化素材瓦片化的随机数生成器。实际种子值并不重要，但不同的值会给出不同的结果。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。如果您的图像在遮罩通道也有锐利边缘的地方有锐利的颜色变化，Normal 模式可能会给出更好的结果。
