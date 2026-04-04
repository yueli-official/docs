---
title: DissolveLensFlare
---

## S_DissolveLensFlare

使用动态镜头光晕在两个输入素材之间进行转场。
素材相互溶解的同时，镜头光晕沿直线移动。
光晕在效果持续期间逐渐增大和缩小。
应对 Dissolve Percent 参数进行动画处理以控制转场速度。

在 Sapphire Transitions 效果子菜单中。

![DissolveLensFlare](../_static/DissolveLensFlare.jpg)


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
  From 和 To 素材之间的溶解速度。设为 1 时，溶解在效果的整个持续时间内进行。设为更高值时，溶解时间更短，但镜头光晕仍在整个持续时间内改变大小和亮度。设为 10 可使转场更快捷，更像闪帧切换。

- **Hotspot Center** (X & Y, Default: [0 0], Range: any)
  转场中心点处光晕最亮点经过的位置。

- **Hotspot Speed** (Default: 1, Range: 0 to 2)
  光晕在屏幕上扫过的速度。设为零可使镜头光晕在原地增大和缩小。

- **Hotspot Angle** (Default: -25, Range: any)
  光晕在屏幕上扫过的角度。

- **Pivot** (X & Y, Default: [0 0], Range: any)
  光晕的各元素将排列在 Hotspot 和 Pivot 位置之间的连线上。Pivot 位置使用屏幕坐标表示。

- **Flare Brightness** (Default: 8, Range: 0 or greater)
  转场中心点处光晕的最大亮度。

- **Flare Fade** (Default: 1, Range: 0 to 1)
  在转场开始和结束时亮度降低的比例。

- **Flare Width** (Default: 2.5, Range: 0 or greater)
  转场中心点处光晕的最大宽度。

- **Flare Shrink** (Default: 0.5, Range: 0 to 1)
  在转场开始和结束时光晕宽度缩小的比例。

- **Rel Heights** (Default: 1, Range: 0 or greater)
  缩放所有光晕元素的垂直尺寸，使其变为椭圆形而非圆形。也可以使用 Scale Widths 控件调整此参数。

- **Lens** (Default: 0, Range: 0 or greater)
  要应用的镜头光晕类型。也可以通过在光晕设计器中编辑光晕来创建自定义光晕类型或修改现有类型。


### Flare Details Parameters:

Rays Rotate:
*Default:
*0,
*Range:
*any.以度为单位旋转镜头光晕的射线元素（如有）。

Color:
*Default rgb:
*[1 1 1].缩放所有光晕元素的颜色。

Gamma:
*Default:
*1,
*Range:
*0 or greater.增加 gamma 可使光晕变亮，尤其能提升较暗元素的亮度。

Saturation:
*Default:
*1,
*Range:
*any.缩放光晕元素的色彩饱和度。
增加可获得更鲜艳的颜色。设为 0 可获得单色镜头光晕。

Hue Shift:
*Default:
*0,
*Range:
*-1 to 1.以从红到绿到蓝再到红的旋转圈数来移动光晕的色相。

Hotspot Color:
*Default rgb:
*[1 1 1].仅缩放热点元素的颜色。

Hotspot Brightness:
*Default:
*1,
*Range:
*0 or greater.仅缩放热点元素的亮度。

Rays Brightness:
*Default:
*1,
*Range:
*0 or greater.仅缩放射线元素的亮度。

Rays Num Scale:
*Default:
*1,
*Range:
*0 or greater.增加或减少射线的数量。

Rays Length:
*Default:
*1,
*Range:
*0 or greater.在不改变射线粗细或其他光晕元素大小的情况下，调整射线的长度。

Rays Thickness:
*Default:
*1,
*Range:
*0 or greater.调整光晕中各条射线的粗细。

Other Brightness:
*Default:
*1,
*Range:
*0 or greater.缩放所有不在热点位置的光晕元素的亮度。

Other Width:
*Default:
*1,
*Range:
*0 or greater.缩放所有不在热点位置的光晕元素的宽度。

Other Color:
*Default rgb:
*[1 1 1].缩放所有不在热点位置的光晕元素的颜色。

Blur Flare:
*Default:
*0,
*Range:
*0 or greater.
如果为正值，光晕图像在与背景合成前将按此量进行模糊处理。

### Other Parameters:

Bg Brightness:
*Default:
*1,
*Range:
*0 or greater.在与光晕合成前缩放背景的亮度。如果为 0，结果将只包含黑色背景上的光晕图像。

Combine:
*Popup menu, Default: Screen
*.决定光晕图像与背景的合成方式。
*Screen:
*执行混合函数，可帮助防止结果过亮。*Add:
*将光晕图像叠加到背景上。

Tint Bg Whites:
*Check-box, Default:
*off.如果启用，仅在结果被限制到最大亮度后才添加光晕的色度。
这样即使在明亮的白色背景上也能看到光晕图像的颜色。对于大多数背景，不会有明显差异。

Affect Alpha:
*Default:
*1,
*Range:
*0 or greater.如果此值为正，输出的 Alpha 通道将包含来自光晕的部分不透明度。红、绿、蓝光晕亮度的最大值按此值缩放，并在每个像素处与背景 Alpha 合成。

Performance:
*Popup menu, Default: full flare
*.决定是渲染所有元素还是仅渲染选定元素。在光晕设计器中，某些元素被选定为对光晕外观至关重要的元素。仅渲染优先级元素应能在预览时呈现真实镜头光晕的外观和感觉，但渲染速度比完整光晕更快。
*full flare:
*渲染所有镜头光晕元素。*priority only:
*仅渲染部分镜头光晕元素以提高性能。

Opacity:
*Popup menu, Default: Normal
*.决定处理不透明度/透明度的方法。
*All Opaque:
*当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。*Normal:
*正常处理不透明度。*As Premult:
*按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

Show Flare Width:
*Check-box, Default:
*on.开启或关闭用于调整 Hotspot Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。

Show Hotspot Center:
*Check-box, Default:
*on.开启或关闭用于调整 Hotspot Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。

Show Hotspot Angle:
*Check-box, Default:
*on.开启或关闭用于调整 Hotspot Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。

Show Rays Rotate:
*Check-box, Default:
*off.开启或关闭用于调整 Hotspot Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。

Show Pivot:
*Check-box, Default:
*on.
开启或关闭用于调整 Pivot 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。参见
[Motion Blur](/en/sapphire/#motion-blur) 的通用信息。
