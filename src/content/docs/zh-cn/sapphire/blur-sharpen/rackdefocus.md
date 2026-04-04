---
title: RackDefocus
---

## S_RackDefocus

使用"弥散圆"卷积生成源素材的散焦版本。此效果通常比高斯模糊更适合模拟真实的散焦相机镜头，因为亮点可以被散焦为清晰的形状而不是被平滑掉。光圈形状可以通过 Points、Pointiness 和 Rotate 来控制，Use Gamma 参数可以调整模糊高光的相对亮度。

位于 Sapphire Blur+Sharpen 效果子菜单中。

![RackDefocus](../_static/RackDefocus.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Mask**: 默认为无。在结果和源输入之间进行插值。白色区域使用效果的结果。黑色区域使用源素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: Defocus Color)
  在全色或单色散焦之间选择。
  - **Defocus Color**: 散焦源输入的所有通道。
  - **Defocus Mono**: 先将源素材转为单色，然后应用散焦（更快）。

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
  在使用前按此像素量膨胀或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的 Fast 模式下快速调整，还是在 High 质量模式下获得更好的效果。
  - **Fast**: 在 Fast 模式下膨胀 Mocha 遮罩，便于快速调整。
  - **High**: 在 High 质量模式下膨胀 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩并将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果并显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当同时提供 Mocha 遮罩和输入遮罩时，确定如何组合它们。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，只使用 Mocha 遮罩。

- **Defocus Width** (Default: 0.088, Range: 0 or greater)
  散焦的宽度。此参数可通过 Defocus Width Widget 调整。

- **Rel Height** (Default: 1, Range: 0.01 or greater)
  光圈形状的相对高度。如果不为 1，圆形会变成椭圆形，等等。

- **Shape** (Popup menu, Default: Circle)
  确定模拟相机光圈的形状。
  - **Circle**: 圆形。
  - **3 sides**: 三角形。
  - **4 sides**: 正方形。
  - **5 sides**: 五边形。
  - **6 sides**: 六边形。
  - **7 sides**: 等等。

- **Show Shape** (Check-box, Default: off)
  显示光圈形状而不是散焦图像。

- **Roundness** (Default: 0, Range: any)
  修改模拟相机光圈的形状。值为 1 产生圆形；0 产生由 Shape 参数指定边数的平边多边形。小于 0 会使边向内挤压产生星形，大于 1 会使角向内挤压产生花朵形状。如果 Shape 设置为 Circle 则无效。

- **Rotate** (Default: 0, Range: any)
  旋转光圈形状。

- **Bokeh** (Default: 0, Range: any)
  柔化光圈形状的外边缘，使散焦高光看起来更柔和。负值会使光圈形状的中心变暗，产生环形散焦形状。

- **Lens Noise** (Default: 0, Range: 0 or greater)
  增加此值可向光圈形状添加噪点，使散焦稍微变脏。可使结果更逼真。超过 1 可获得更具风格化的结果。

- **Noise Freq** (Default: 40, Range: 0.01 or greater)
  添加噪点的频率。如果 Lens Noise 为零则忽略。

- **Noise Freq Rel X** (Default: 1, Range: 0.01 or greater)
  添加光圈噪点的相对水平频率。增加此值可垂直拉伸，减少可水平拉伸。

- **Noise Seed** (Default: 0.123, Range: 0 or greater)
  添加噪点的种子值。要使噪点在每帧上看起来不同，请将此值设为每帧不同的动画。实际值无关紧要，只要每帧不同即可。

- **Gauss Blur** (Default: 0, Range: 0 or greater)
  如果为正值，还会应用高斯模糊以平滑形状的边缘。这也可能使高光变暗，因为高斯模糊中不考虑 Gamma。

- **Use Gamma** (Default: 1, Range: 0.1 or greater)
  大于 1 的值使源素材中的高光在应用散焦后保持其亮度。

- **Boost Highlights** (Default: 0, Range: 0 or greater)
  增加源素材中高光亮度的量。增加此参数可使高光过曝而不影响暗部或中间调。

- **Hilight Threshold** (Default: 0.9, Range: 0 or greater)
  高光的最小亮度值。比此值更亮的像素将根据 Boost Highlights 参数进行增亮。

- **Chroma Distort** (Default: 0, Range: any)
  在图像边缘添加一些色差；红色和蓝色波长的光在真实镜头中折射方式不同，在光线以斜角射入镜头的地方产生彩色边缘。

- **Color Fringing** (Default: 0, Range: any)
  Color Fringing 通过改变每个颜色通道的焦距，在图像中每个物体周围产生彩色环。它与 Chroma Distort 产生不同风格的色差，因为它不仅出现在图像角落。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果的亮度。

- **Offset Darks** (Default: 0, Range: any)
  将此灰度值添加到结果的较暗区域。可以为负值以增加对比度。

- **Mix With Source** (Default: 0, Range: 0 to 1)
  在散焦结果和原始源素材之间进行插值。设为 1 可获得原始源素材。

- **Edge Mode** (Popup menu, Default: Reflect)
  确定访问源图像外部区域时的行为。
  - **Transparent**: 源图像外部区域被视为透明，这可能在图像边缘产生透明度。选择此项可获得最快的渲染速度。
  - **Repeat**: 重复图像边界外的最后一个像素。
  - **Reflect**: 在边界外反射图像。

- **Soft Borders** (Check-box, Default: off)
  如果启用，在处理前向输入图像添加透明边框。这允许结果包含超出原始图像大小的柔和边缘。关闭时，效果仅在帧内发生，结果将在边界处保留边缘。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且无透明度（alpha=1）时，使用此选项可略微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式略快，但结果也将是预乘形式，有时可能不太准确。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用 Mask 输入通道生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此量模糊遮罩输入。这可以在遮罩区域和非遮罩区域之间提供更平滑的过渡。除非提供了遮罩输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果启用，反转遮罩输入，使效果应用于遮罩为黑色而非白色的区域。除非提供了遮罩输入，否则无效。

- **Show Defocus Width** (Check-box, Default: on)
  打开或关闭用于调整 Defocus Width 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为它们支持屏幕控件。
