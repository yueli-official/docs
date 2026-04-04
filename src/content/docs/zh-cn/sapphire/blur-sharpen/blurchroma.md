---
title: BlurChroma
---

## S_BlurChroma

将源素材分离为亮度和色度分量，分别独立模糊色度和/或亮度，然后重新组合。您还可以独立缩放亮度和色度以增强或去除其中一个。

位于 Sapphire Blur+Sharpen 效果子菜单中。

![BlurChroma](../_static/BlurChroma.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Matte**: 默认为无。如果提供，模糊仅在此输入的亮区指定的源素材区域上执行。此遮罩外的像素不会被模糊，也不会参与遮罩内的模糊结果像素。此输入可通过 Invert Matte 或 Matte Use 参数进行调整。


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

- **Blur Chroma** (Default: 0.4, Range: 0 or greater)
  色度的模糊量。此参数可通过 Blur Chroma Widget 调整。

- **Blur Luminance** (Default: 0, Range: 0 or greater)
  亮度的模糊量。此参数可通过 Blur Luminance Widget 调整。

- **Blur Rel** (X & Y, Default: [1 1], Range: 0 or greater)
  相对水平和垂直模糊宽度。将 Blur Rel X 设为 0 可仅进行垂直模糊，将 Blur Rel Y 设为 0 可仅进行水平模糊。此参数可通过 Blur Chroma Widget 调整。

- **Scale Chroma** (Default: 1, Range: 0 or greater)
  按此量缩放色度。增大可获得更强烈的颜色，减小可获得更柔和的颜色。

- **Scale Luminance** (Default: 1, Range: 0 or greater)
  缩放结果的亮度。

- **Offset Result** (Default: 0, Range: any)
  向结果添加此灰度值（或减去，如果为负值）。0 无效果，0.5 为中灰，1 为白色。

- **Mix With Source** (Default: 0, Range: 0 to 1)
  在模糊结果 (0) 和原始源素材 (1) 之间插值。0.1 可以产生不错的朦胧效果，因为它只混入少量源素材。

- **Filter** (Popup menu, Default: Gauss)
  用于模糊的卷积滤波器类型。
  - **Box**: 使用矩形滤波器。
  - **Triangle**: 更平滑，使用金字塔形滤波器。
  - **Gauss**: 最平滑，使用高斯形滤波器。

- **Subpixel** (Check-box, Default: on)
  启用亚像素量的模糊。使用此选项可使 Blur Chroma 或 Blur Luminance 参数的动画更平滑。

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

- **Show Blur Chroma** (Check-box, Default: on)
  打开或关闭用于调整 Blur Chroma 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show Blur Luminance** (Check-box, Default: on)
  打开或关闭用于调整 Blur Luminance 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

