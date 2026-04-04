---
title: DissolveDefocus
---

## S_DissolveDefocus

在两个输入素材之间转场，同时对每个素材应用散焦。第一个素材被散焦并淡出，第二个素材恢复焦点并淡入。应通过动画 Dissolve Percent 参数来控制转场速度。

在 Sapphire Transitions 效果子菜单中。

![DissolveDefocus](../_static/DissolveDefocus.jpg)


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

- **Defocus Width** (Default: 0.8, Range: 0 or greater)
  散焦的宽度。

- **Defocus Rel From** (Default: 1, Range: 0 or greater)
  缩放应用于第一个素材的散焦量。设为 0 可在没有散焦的情况下淡出。

- **Defocus Rel To** (Default: 1, Range: 0 or greater)
  缩放应用于第二个素材的散焦量。设为 0 可在没有散焦的情况下淡入。

- **Rel Height** (Default: 1, Range: 0.01 or greater)
  光圈形状的相对高度。如果不为 1，圆形会变成椭圆形，等等。

- **Shape** (Popup menu, Default: Circle)
  决定模拟相机光圈的形状。
  - **Circle**: 圆形。
  - **3 sides**: 三角形。
  - **4 sides**: 正方形。
  - **5 sides**: 五边形。
  - **6 sides**: 六边形。
  - **7 sides**: 等等。

- **Show Shape** (Check-box, Default: off)
  显示光圈形状而非散焦后的图像。

- **Roundness** (Default: 0, Range: any)
  修改模拟相机光圈的形状。值为 1 产生圆形；0 产生由 Shape 参数决定边数的平边多边形。小于 0 时边向内挤压形成星形，大于 1 时角向内挤压形成花形。当 Shape 设为 Circle 时无效。

- **Rotate** (Default: 0, Range: any)
  旋转光圈形状。

- **Bokeh** (Default: 0, Range: any)
  柔化光圈形状的外边缘，使散焦高光看起来更柔和。负值会使光圈形状中心变暗，产生环形散焦形状。

- **Lens Noise** (Default: 0, Range: 0 or greater)
  增大可向光圈形状添加噪声，使散焦效果略带脏感。可使结果更逼真。超过 1 可获得更具风格化的效果。

- **Noise Freq** (Default: 40, Range: 0.01 or greater)
  噪声的空间频率。

- **Noise Freq Rel X** (Default: 1, Range: 0.01 or greater)
  添加的光圈噪声的相对水平频率。增大可使其垂直拉伸，减小可使其水平拉伸。

- **Noise Seed** (Default: 0.123, Range: 0 or greater)
  添加噪声的种子值。要使噪声在每帧看起来不同，请将此参数动画为每帧不同的值。实际值不重要；重要的是每帧不同。

- **Use Gamma** (Default: 1, Range: 0.1 or greater)
  大于 1 的值会使源素材中的高光在应用散焦后保持其亮度。

- **Boost Highlights** (Default: 0, Range: 0 or greater)
  增加源素材中高光的亮度值。增大此参数可在不影响暗部或中间调的情况下使高光过曝。

- **Hilight Threshold** (Default: 0.9, Range: 0 or greater)
  高光的最小亮度值。比此值更亮的像素将根据 Boost Highlights 参数被增亮。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果的亮度。

- **Edge Mode** (Popup menu, Default: Reflect)
  决定访问源图像外部区域时的行为。
  - **Transparent**: 源图像外部区域被视为透明，这可能在图像边缘产生透明效果。选择此选项可获得最快的渲染速度。
  - **Repeat**: 在图像边界外重复最后一个像素。
  - **Reflect**: 在边界外反射图像。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度 (alpha=1) 时使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时不太准确。

- **Show Defocus Width** (Check-box, Default: on)
  开启或关闭用于调整 Defocus Width 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些软件支持屏幕控件。
