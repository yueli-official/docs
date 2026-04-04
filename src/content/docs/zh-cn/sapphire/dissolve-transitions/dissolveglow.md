---
title: DissolveGlow
---

## S_DissolveGlow

使用明亮的辉光闪烁在两个输入素材之间转场。素材相互溶解，同时每个素材都会获得一个在效果持续时间内渐入和渐出的辉光。应通过动画 Dissolve Percent 参数来控制转场速度。

在 Sapphire Transitions 效果子菜单中。

![DissolveGlow](../_static/DissolveGlow.jpg)


### Inputs:

- **Foreground**: 当前图层。以此素材开始转场。

- **Background**: 默认为无。以此素材结束转场。


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
  必须禁用 Auto Trans 才能使用此参数。它决定前景和背景输入之间的转场比例，通常从 0 动画到 100 以执行完整的转场。可以调整控制此参数的曲线，以更精细地控制溶解的时间。

- **Dissolve Speed** (Default: 3, Range: 1 or greater)
  从一个素材到另一个素材的溶解速度。设为 1 时，溶解在效果的整个持续时间内进行。设为更高值时，溶解更短，但辉光的渐入和渐出仍占据整个持续时间。设为 10 可使转场更快捷，更像闪帧切换。

- **Glow Brightness** (Default: 6, Range: 0 or greater)
  辉光的整体最大亮度。

- **Glow Threshold** (Default: 0.2, Range: 0 or greater)
  源素材中亮度超过此值的部分将产生辉光。值为 0.9 时仅最亮的位置产生辉光。值为 0 时每个非黑色区域都产生辉光。

- **Glow Color** (Default rgb: [1 1 1])
  辉光的整体颜色。

- **Glow Width** (Default: 0.4, Range: 0 or greater)
  辉光的宽度。此参数及所有宽度参数均可通过宽度控件进行调整。请注意，辉光宽度为零时仍会增强明亮区域；如果要原样传递源素材，请将辉光亮度参数设为零。

- **Width X** (Default: 1, Range: 0 or greater)
  缩放水平辉光宽度。设为 0 则仅显示垂直方向。

- **Width Y** (Default: 1, Range: 0 or greater)
  缩放垂直辉光宽度。设为 0 则仅显示水平方向。

- **Width Red** (Default: 1, Range: 0 or greater)
  缩放红色辉光宽度。如果红、绿、蓝宽度相等，辉光将与源素材的颜色一致。如果不相等，辉光颜色将随距离变化。

- **Width Green** (Default: 1.2, Range: 0 or greater)
  缩放绿色辉光宽度。

- **Width Blue** (Default: 1.4, Range: 0 or greater)
  缩放蓝色辉光宽度。

- **Rel From Brightness** (Default: 1, Range: 0 or greater)
  出场 (From) 素材上辉光的相对亮度。

- **Rel From Width** (Default: 1, Range: 0 or greater)
  出场 (From) 素材上辉光的相对宽度。

- **From Offset Threshold** (Default: 0, Range: any)
  应用于出场 (From) 素材辉光的额外阈值。

- **Rel From Color** (Default rgb: [1 1 1])
  出场 (From) 素材上辉光的相对颜色。

- **Rel To Brightness** (Default: 1, Range: 0 or greater)
  入场 (To) 素材上辉光的相对亮度。

- **Rel To Width** (Default: 1, Range: 0 or greater)
  入场 (To) 素材上辉光的相对亮度。

- **To Offset Threshold** (Default: 0, Range: any)
  应用于入场 (To) 素材辉光的额外阈值。

- **Rel To Color** (Default rgb: [1 1 1])
  入场 (To) 素材上辉光的相对颜色。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Show Glow Width** (Check-box, Default: on)
  开启或关闭用于调整 Glow Width 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
