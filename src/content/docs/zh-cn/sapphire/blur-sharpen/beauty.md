---
title: Beauty
---

## S_Beauty

对皮肤区域应用平滑、色彩校正、柔焦和辉光。皮肤区域的确定方式取决于 Enable Skin Detection 的值以及是否提供了第二个输入，共有四种情况：Enable Skin Detection
OFF
，无第二个输入：效果应用于整个图像。Enable Skin Detection
ON
，无第二个输入：使用指定的 Skin Color、
Luma 和 Chroma Range 参数生成内部遮罩。Enable Skin Detection
OFF
，提供了第二个输入：第二个输入用作外部遮罩。效果
应用于遮罩的亮区（若要应用于暗区，请参见 Invert Matte 参数）。Enable Skin Detection
ON
，提供了第二个输入：生成内部遮罩并与
外部遮罩相乘。

位于 Sapphire Blur+Sharpen 效果子菜单中。

![Beauty](../_static/Beauty.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **Matte**: 默认为无。与内部皮肤检测结合使用的垃圾遮罩。


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

- **Enable Skin Detection** (Check-box, Default: on)
  根据 Skin Color、Luma 和 Chroma Range 等参数生成皮肤区域的内部遮罩。禁用时，效果均匀地应用于遮罩输入（详见上述说明）。

- **Skin Color** (Default rgb: [0.749 0.498 0.345])
  用于检测的代表性皮肤颜色。此参数可通过 Skin Color Widget 调整。

- **Luma Range** (Default: 0.4, Range: 0 or greater)
  与 Skin Color 的亮度差异，在此范围内视为皮肤。

- **Chroma Range** (Default: 0.2, Range: 0 or greater)
  与 Skin Color 的色度差异，在此范围内视为皮肤。此参数可通过 Chroma Range Widget 调整。

- **Rel Orange** (Default: 1, Range: 0 to 1)
  Chroma Range 中橙色的相对量。橙色轴沿人类肤色线，这是所有人种共有的，在矢量示波器术语中也称为 I 线。降低此参数可帮助从皮肤遮罩中排除金色或红色头发（以及其他发色中的类似高光）、一些红/橙/黄色运动制服、暖色背景等。

- **Rel Purple** (Default: 0.3, Range: 0 to 1)
  Chroma Range 中紫色的相对量。紫色轴垂直于橙色轴，因此最不像皮肤色。在大多数情况下，可以减小此参数以从皮肤遮罩中排除嘴唇、眼睛、衣服和珠宝。增大此参数可将紫色和绿色色调添加到皮肤检测遮罩中，例如眼影、不良照明或外星人肤色（即非人类）。

- **Range Softness** (Default: 0.75, Range: 0 to 1)
  控制皮肤检测遮罩的柔和度。值为 1 表示只有与 Skin Color 完全匹配的像素才会产生为 1 的遮罩值，其他每个像素的遮罩值与其到 Skin Color 的距离成正比。值为 0 表示硬遮罩，在 Skin Color 的亮度/色度范围内的所有像素都会产生为 1 的遮罩值。

- **Clip White** (Default: 1, Range: 0 to 1)
  大于此值的皮肤检测遮罩值将被设为 1。

- **Clip Black** (Default: 0, Range: 0 to 1)
  小于此值的皮肤检测遮罩值将被设为 0。

- **Post Blur** (Default: 0, Range: 0 or greater)
  按此量模糊皮肤检测遮罩。

- **Show** (Popup menu, Default: Final)
  选择输出类型。
  - **Final**: 显示最终输出。
  - **Skin Detect Matte**: 显示由内部皮肤检测器生成的遮罩。
  - **With Garbage Matte**: 显示输入垃圾遮罩和内部皮肤检测器遮罩的组合结果。
  - **Skin**: 显示将皮肤检测遮罩应用于源素材的结果。
  - **Skin with Garbage Matte**: 显示将输入垃圾遮罩和内部皮肤检测器遮罩组合后应用于源素材的结果。

- **Show Color Helper** (Check-box, Default: off)
  显示交互式叠加层以帮助设置 Skin Color、Chroma Range、Rel Orange 和 Rel Purple。叠加层显示与 Skin Color 参数的亮度（亮度值）匹配的所有可能颜色。橙色在左上角，紫色在右上角（此方向类似于传统的广播矢量示波器）。与皮肤检测算法匹配的颜色会被高亮显示。更改 Skin Color 将移动高亮区域，调整 Chroma Range 会改变高亮区域的大小，调整 Rel Orange/Purple 会沿正方形的对角线拉伸区域。

- **Matte Use** (Popup menu, Default: Luma)
  确定如何使用 Matte 输入通道来生成单色遮罩。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Blur Matte** (Default: 0, Range: 0 or greater)
  在使用前按此量模糊 Matte 输入。这可以在遮罩区域和非遮罩区域之间提供更平滑的过渡。除非提供了 Matte 输入，否则无效。

- **Invert Matte** (Check-box, Default: off)
  如果开启，反转 Matte 输入，使效果应用于 Matte 为黑色而非白色的区域。除非提供了 Matte 输入，否则无效。

- **Suppress BG** (Check-box, Default: off)
  仅将 Beauty 应用于由 Face Center 及相关参数指定的区域。

- **Face Center** (X & Y, Default: [0 0], Range: any)
  启用 Suppress BG 时脸部区域的中心位置。此参数可通过 Face Center Widget 调整。

- **Face Softness** (Default: 0.1, Range: 0 or greater)
  启用 Suppress BG 时使脸部区域更柔和。这将在脸部区域到背景之间提供更平滑的过渡，但也可能降低脸部区域中 Beauty 的强度。

- **Face Radius** (Default: 0.4, Range: 0 or greater)
  启用 Suppress BG 时脸部区域的大小。此参数可通过 Face Radius Widget 调整。

- **Face Rel Height** (Default: 1.33, Range: 0.05 or greater)
  启用 Suppress BG 时脸部区域的相对高度。

- **Face Rotate** (Default: 0, Range: any)
  启用 Suppress BG 时脸部区域的旋转。此参数可通过 Face Rotate Widget 调整。

- **Show Face Widget** (Check-box, Default: off)
  显示交互式叠加层以帮助放置和调整脸部区域的大小。

- **Pore Size** (Default: 0.01, Range: 0 or greater)
  小于此大小的特征（毛孔等）即使在模糊时也会被保留。

- **Blur Amount** (Default: 0.056, Range: 0 or greater)
  缩放模糊的宽度。

- **Edge Threshold** (Default: 0.1, Range: 0 or greater)
  被大于此值的边缘分隔的颜色区域不会相互模糊。

- **Soften Shadows** (Default: 0.2, Range: -1 to 1)
  正值减少阴影的可见度，负值使阴影更明显。减少阴影可使对象看起来更年轻，而加深阴影会使其看起来更老。

- **Shadow Thresh** (Default: 0.6, Range: 0 or greater)
  小于此值的暗区将被 Soften Shadows 增强/减弱。

- **Reduce Shine** (Default: 0, Range: 0 to 1)
  使明亮、有光泽的区域变暗。变暗过程可能导致缺乏颜色，使用 Shine Saturation 可在受影响区域恢复自然肤色。

- **Shine Saturation** (Default: 1, Range: 0 or greater)
  缩放明亮区域的颜色饱和度。适用于在需要变暗的光亮区域添加自然肤色。

- **Shine Thresh** (Default: 0.9, Range: 0 or greater)
  亮度高于此值的区域将受到 Reduce Shine 的影响。

- **Hue Shift** (Default: 0, Range: any)
  偏移源颜色的色相，以从红到绿到蓝再到红的圈数计。

- **Saturation** (Default: 1.1, Range: -2 to 8)
  缩放结果的颜色饱和度。增大可获得更强烈的颜色。设为 0 则为单色。也可以通过设为负值来反转结果的色度。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果的亮度。

- **Tint** (Default rgb: [1 1 1])
  用此颜色缩放结果，从而为较亮区域着色。

- **Soft Focus** (Default: 0, Range: 0 or greater)
  缩放柔焦模糊的宽度。

- **Glow Brightness** (Default: 0.1, Range: 0 or greater)
  缩放皮肤辉光的亮度。

- **Glow Color** (Default rgb: [1 1 1])
  缩放皮肤辉光的颜色。

- **Glow Threshold** (Default: 0.2, Range: 0 or greater)
  辉光从皮肤区域中亮度高于此值的位置生成。值为 0.9 时仅在最亮的点产生辉光。值为 0 时为每个非黑色区域产生辉光。

- **Glow Width** (Default: 0.1, Range: 0 or greater)
  缩放皮肤辉光的距离。

- **Mix With Source** (Default: 0, Range: 0 to 1)
  在模糊结果 (0) 和原始源素材 (1) 之间插值。0.1 可以产生不错的朦胧效果，因为它只混入少量源素材。

- **Opacity** (Popup menu, Default: Normal)
  确定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且无透明度 (alpha=1) 时，使用此选项可稍微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按图像已为预乘形式处理（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式稍快，但结果也将是预乘形式，有时可能不太准确。

