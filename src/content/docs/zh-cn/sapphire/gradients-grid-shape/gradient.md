---
title: Gradient
---

## S_Gradient

使用给定的起始和结束位置及颜色，在屏幕上生成平滑的颜色渐变，然后可选择将渐变与背景素材合成。增加 Add Noise 可减少因颜色量化导致的渐变色带伪影。

在 Sapphire Render 效果子菜单中。

![Gradient](../_static/Gradient.jpg)


### Inputs:

- **Background**: 当前图层。用于与渐变合成的素材。

- **Mask**: 默认为无。在结果和源输入之间进行插值。白色区域使用效果结果，黑色区域使用源素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此数值模糊 Mocha 遮罩。可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

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
  在使用前按此像素量膨胀或侵蚀 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在高质量模式下膨胀 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩同时提供给效果时，确定如何合并 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Start** (X & Y, Default: [0 0.596], Range: any)
  渐变的起始位置。可使用 Start Widget 调整此参数。

- **End** (X & Y, Default: [0 -0.596], Range: any)
  渐变的结束位置。可使用 End Widget 调整此参数。

- **Start Color** (Default rgb: [1 1 1])
  渐变在起始位置的颜色。

- **End Color** (Default rgb: [0 0 0])
  渐变在结束位置的颜色。

- **Swap Colors** (Check-box, Default: off)
  如果勾选，通过交换起始和结束颜色来反转渐变方向。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放渐变图像的亮度（包括起始颜色和结束颜色）。

- **Add Noise** (Default: 0, Range: 0 or greater)
  如果为正值，会向渐变添加此数量的噪点。这可以产生颗粒效果并消除因量化导致的渐变色带。将此值设为 1.0 可为 8 位结果启用有效的去色带处理。

- **Smooth Curve** (Default: 0, Range: 0 to 1)
  如果为零，在起始颜色和结束颜色之间使用线性插值。增加此值以使用更平滑的"S"形曲线进行插值，可以减少对渐变起始和结束位置的视觉感知。

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  在与渐变合成之前缩放背景的亮度。

- **Combine** (Popup menu, Default: Grad Only)
  确定渐变与背景的合成方式。
  - **Grad Only**: 仅显示渐变图像，不包含背景。
  - **Mult**: 背景与渐变相乘。
  - **Add**: 背景与渐变相加。
  - **Screen**: 背景与渐变使用滤色操作混合。
  - **Difference**: 结果为背景与渐变的差值。
  - **Overlay**: 使用叠加功能合并渐变和背景。

- **Input Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且无透明度（alpha=1）时，使用此选项可略微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也略快于正常模式，但结果也将为预乘形式，有时可能不太准确。

- **Output Opacity** (Popup menu, Default: Copy From Input)
  确定结果的不透明度/透明度。此效果不处理输入的不透明度（Alpha 通道），但可以从输入复制不透明度，或输出完全不透明的结果。
  - **All Opaque**: 使结果完全不透明，无透明度。
  - **Copy From Input**: 从提供给此效果的当前图层复制不透明度/透明度。

- **Mask Use** (Popup menu, Default: Luma)
  确定如何使用遮罩输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  在使用前按此数值模糊蒙版输入。这可以提供遮罩区域和非遮罩区域之间更平滑的过渡。除非提供了蒙版输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  如果开启，反转蒙版输入，使效果应用于蒙版为黑色而非白色的区域。除非提供了蒙版输入，否则无效。

- **Show Start** (Check-box, Default: on)
  开启或关闭用于调整 Start 参数的屏幕用户界面。此参数仅在支持屏幕控件的 AE 和 Premiere 中显示。

- **Show End** (Check-box, Default: on)
  开启或关闭用于调整 End 参数的屏幕用户界面。此参数仅在支持屏幕控件的 AE 和 Premiere 中显示。
