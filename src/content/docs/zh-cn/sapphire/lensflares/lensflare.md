---
title: LensFlare
---

## S_LensFlare

在背景素材上渲染镜头光晕图像，将各种光晕元素在热点和枢轴位置之间对齐。使用镜头菜单选择不同类型的镜头光晕。

在 Sapphire Lighting 效果子菜单中。

![LensFlare](../_static/LensFlare.jpg)


### Inputs:

- **Background**: 当前图层。要在其上应用镜头光晕的素材。

- **Occlusion**: 默认为无。遮挡并着色光晕。光晕的亮度会根据此素材的不透明度降低，热点位置处的颜色用于为光晕着色。这些行为可以通过 Occlusion Softness、Occlusion From、Invert Occlusion 和 Use Color 参数进行调整。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: 2D)
  在 LensFlare 效果的多个变体之间选择。
  - **2D**: 手动定位热点。光晕可以被遮罩输入遮挡。
  - **3D**: 一个或多个热点可以连接到 3D 合成中的灯光。光晕可以被合成中的 3D 图层遮挡。

- **Lens** (Default: 0, Range: 0 or greater)
  要应用的镜头光晕类型。也可以通过在光晕设计器中编辑光晕来创建自定义镜头光晕类型或修改现有类型。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此数值模糊 Mocha 遮罩。这可用于柔化遮罩的边缘或量化伪影，并平滑时间位移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 遮罩的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  如果启用，Mocha 遮罩的黑白将在应用效果前反转。

- **Resize Mocha** (Default: 1, Range: 0 or greater)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 or greater)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 or greater)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素值扩展或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下进行 Dilate Mocha，以便快速调整。
  - **High**: 在高质量模式下进行 Dilate Mocha，以获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用到整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  绕过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩都提供给效果时，决定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Scale Widths** (Default: 1.15, Range: 0 or greater)
  缩放所有光晕元素的大小。此参数可以使用 Scale Widths 控件进行调整。

- **Rel Heights** (Default: 1, Range: 0 or greater)
  缩放所有光晕元素的垂直尺寸，使它们变为椭圆形而非圆形。也可以使用 Scale Widths 控件进行调整。

- **Rays Rotate** (Default: 0, Range: any)
  旋转镜头光晕的光线元素（如果有），单位为度。

- **Hotspot** (X & Y, Default: [-0.444 0.176], Range: any)
  光晕中最亮点的位置，使用屏幕坐标。可以通过启用并移动热点控件来设置。

- **Hotspot Uses Mocha** (Check-box, Default: off)
  控制 LensFlare 热点是由 Hotspot 参数控制还是跟随 Mocha 内部跟踪的热点。

- **Smooth Hotspot Track** (Integer, Default: 0, Range: 0 or greater)
  控制在稳定 Mocha 点跟踪时要平均多少个点。

- **Hotspots** (Default: 0, Range: 0 or greater)
  要将光晕中心附加到的 AE 灯光。

- **Pivot** (X & Y, Default: [0 0], Range: any)
  光晕的元素将排列在热点和枢轴位置之间的一条线上。枢轴位置使用屏幕坐标。

- **Pivot Uses Mocha** (Check-box, Default: off)
  控制 LensFlare 枢轴是由 Pivot 参数控制还是跟随 Mocha 内部跟踪的枢轴。

- **Smooth Pivot Track** (Integer, Default: 0, Range: 0 or greater)
  控制在稳定 Mocha 点跟踪时要平均多少个点。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放所有光晕元素的亮度。

- **Color** (Default rgb: [1 1 1])
  缩放所有光晕元素的颜色。

- **Gamma** (Default: 1, Range: 0.1 or greater)
  增加 Gamma 会使光晕变亮，尤其会增强较暗的元素。

- **Saturation** (Default: 1, Range: -2 to 8)
  缩放光晕元素的色彩饱和度。增加以获得更强烈的颜色。设为 0 可获得单色镜头光晕。

- **Hue Shift** (Default: 0, Range: any)
  偏移光晕的色相，从红到绿到蓝再到红循环。

- **Hotspot Bright** (Default: 1, Range: 0 or greater)
  仅缩放热点元素的亮度。

- **Hotspot Color** (Default rgb: [1 1 1])
  仅缩放热点元素的颜色。

- **Rays Brightness** (Default: 1, Range: 0 or greater)
  仅缩放光线元素的亮度。

- **Rays Num Scale** (Default: 1, Range: 0 or greater)
  增加或减少光线的数量。

- **Rays Length** (Default: 1, Range: 0 or greater)
  调整光线的长度，而不改变其粗细或其他光晕元素的大小。

- **Rays Thickness** (Default: 1, Range: 0 or greater)
  调整光晕中单根光线的粗细。

- **Other Brightness** (Default: 1, Range: 0 or greater)
  缩放所有不在热点位置的光晕元素的亮度。

- **Other Color** (Default rgb: [1 1 1])
  缩放所有不在热点位置的光晕元素的颜色。

- **Other Width** (Default: 1, Range: 0 or greater)
  缩放所有不在热点位置的光晕元素的宽度。

- **Atmosphere Amp** (Default: 0, Range: 0 or greater)
  大气效果模拟光晕穿过尘埃大气时拾取光线或被遮蔽的效果。此参数调整大气效果的数量或振幅。零值产生更平滑的光晕，较高值产生更脏的外观。大气效果影响光晕的某些部分但不影响其他部分。可以在光晕设计器中通过切换每个元素的忽略大气设置来调整此行为。通常，源自相机内部的元素（如二次反射）应使用忽略大气，而源自相机外部的元素（如辉光）则不应使用。

- **Atmosphere Freq** (Default: 1, Range: 0.1 to 20)
  控制大气噪声的空间频率。调高以获得更精细的细节，调低以获得更宽泛的整体变化。

- **Atmosphere Detail** (Default: 0.6, Range: 0 to 1)
  控制大气模拟中的精细细节量。减小以获得更平滑的大气效果，增大以获得更粗糙或颗粒感的外观。

- **Atmosphere Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化大气噪声的随机数生成器。实际种子值并不重要，但不同的种子会产生不同的结果，相同的值应该产生可重复的结果。

- **Atmosphere Speed** (Default: 1, Range: any)
  大气中的云状噪声会随时间演变，如同真实的尘埃云；此参数控制云图案随时间变化的速度。设为零可获得静态图案。

- **Flicker Amp** (Default: 0, Range: 0 or greater)
  光晕亮度随机闪烁的量。

- **Flicker Speed** (Default: 1, Range: 0 or greater)
  随机闪烁的速度。

- **Flicker Randomness** (Default: 0.6, Range: 0 to 1)
  控制闪烁的变化程度。设为零时，光晕将持续闪烁，带有少量随机变化。在较高值时，闪烁将有较长的稳定期，偶尔出现大幅波动。

- **Atmosphere Amp** (Default: 0, Range: 0 or greater)
  大气效果模拟光晕穿过尘埃大气时拾取光线或被遮蔽的效果。此参数调整大气效果的数量或振幅。零值产生更平滑的光晕，较高值产生更脏的外观。大气效果影响光晕的某些部分但不影响其他部分。可以在光晕设计器中通过切换每个元素的忽略大气设置来调整此行为。通常，源自相机内部的元素（如二次反射）应使用忽略大气，而源自相机外部的元素（如辉光）则不应使用。

- **Atmosphere Freq** (Default: 1, Range: 0.1 to 20)
  控制大气噪声的空间频率。调高以获得更精细的细节，调低以获得更宽泛的整体变化。

- **Atmosphere Detail** (Default: 0.6, Range: 0 to 1)
  控制大气模拟中的精细细节量。减小以获得更平滑的大气效果，增大以获得更粗糙或颗粒感的外观。

- **Atmosphere Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化大气噪声的随机数生成器。实际种子值并不重要，但不同的种子会产生不同的结果，相同的值应该产生可重复的结果。

- **Atmosphere Speed** (Default: 1, Range: any)
  大气中的云状噪声会随时间演变，如同真实的尘埃云；此参数控制云图案随时间变化的速度。设为零可获得静态图案。

- **Flicker Amp** (Default: 0, Range: 0 or greater)
  光晕亮度随机闪烁的量。

- **Flicker Speed** (Default: 1, Range: 0 or greater)
  随机闪烁的速度。

- **Flicker Randomness** (Default: 0.6, Range: 0 to 1)
  控制闪烁的变化程度。设为零时，光晕将持续闪烁，带有少量随机变化。在较高值时，闪烁将有较长的稳定期，偶尔出现大幅波动。

- **Blur Flare** (Default: 0, Range: 0 or greater)
  如果为正值，光晕图像在与背景合成之前会按此数值进行模糊处理。

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  在与光晕合成之前缩放背景的亮度。如果为 0，结果将仅包含黑色背景上的光晕图像。

- **Combine** (Popup menu, Default: Screen)
  决定光晕图像如何与背景合成。
  - **Screen**: 执行混合功能，有助于防止过亮的结果。
  - **Add**: 将光晕图像添加到背景上。
  - **Flare Only**: 仅显示光晕图像，没有背景。

- **Tint Bg Whites** (Check-box, Default: off)
  如果启用，光晕的色度仅在结果被钳制到最大亮度之后才添加。这允许光晕图像的颜色即使在明亮的白色背景上也能可见。对于大多数背景，不会有明显差异。

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  如果此值为正，输出 Alpha 通道将包含来自光晕的一些不透明度。光晕红、绿、蓝亮度的最大值按此值缩放，并在每个像素处与背景 Alpha 合成。


### Edge Triggers Parameters:

Edge Width:
*Default:
*0.3,
*Range:
*0 or greater.在屏幕边缘创建一个触发区域，影响光晕的亮度和大小。此参数控制区域的宽度。设为零将禁用边缘触发器。

Edge Falloff:
*Default:
*2,
*Range:
*0.01 or greater.控制触发器强度在远离边缘时下降的速度。值为 1 产生线性渐变。大于 1 的值产生较陡的初始下降，然后逐渐趋于平缓。小于 1 的值产生起初平缓、末端变陡的斜率。

Shift Out:
*Default:
*0,
*Range:
*any.将触发区域从屏幕边缘向外移动，使峰值位于屏幕外。负值将触发区域向图像中心内移。

One Way:
*Check-box, Default:
*off.如果选中此框，当热点移出屏幕时触发器将保持最大强度，而不是随着远离边缘而逐渐降低。

Edge Scale Brightness:
*Default:
*1.5,
*Range:
*0 or greater.当触发器处于峰值强度时，按此数值缩放亮度。例如，值为 3 将使光晕在通过触发区域时亮度提升到正常亮度的 3 倍。如果小于 1，触发器将使光晕变暗而不是变亮。

Edge Scale Widths:
*Default:
*1.2,
*Range:
*0 or greater.当触发器处于峰值强度时，按此数值缩放光晕的宽度。例如，值为 3 将使光晕在通过触发区域时扩展到正常大小的 3 倍。如果小于 1，触发器将缩小光晕而不是放大。

Show Edge Zones:
*Check-box, Default:
*off.在输出上叠加灰度图像，显示边缘触发器的位置和强度。

Center Radius:
*Default:
*0.3,
*Range:
*0 or greater.创建一个以枢轴为中心的触发区域，影响光晕的亮度和大小。此参数控制区域的半径。设为零将禁用中心触发器。

Center Falloff:
*Default:
*1,
*Range:
*0.01 or greater.控制触发器强度在远离边缘时下降的速度。值为 1 产生线性渐变。大于 1 的值产生较陡的初始下降，然后逐渐趋于平缓。小于 1 的值产生起初平缓、末端变陡的斜率。

Center Scale Brightness:
*Default:
*1.5,
*Range:
*0 or greater.当触发器处于峰值强度时，按此数值缩放亮度。例如，值为 3 将使光晕在通过触发区域时亮度提升到正常亮度的 3 倍。如果小于 1，触发器将使光晕变暗而不是变亮。

Center Scale Widths:
*Default:
*1.2,
*Range:
*0 or greater.当触发器处于峰值强度时，按此数值缩放光晕的宽度。例如，值为 3 将使光晕在通过触发区域时扩展到正常大小的 3 倍。如果小于 1，触发器将缩小光晕而不是放大。

Show Center Zone:
*Check-box, Default:
*off.
在输出上叠加灰度图像，显示中心触发器的位置和强度。

### Occlusion Parameters:

Occlusion Softness:
*Default:
*0.0224,
*Range:
*0 or greater.增加此值可使光晕在热点移到物体后方时逐渐淡出。设为零可使光晕突然消失。在较大的柔和度值下，热点位置会自动调整到未被遮挡的区域。这准确模拟了部分被遮挡的大型光源，并防止热点出现在遮挡物体前方。

Occlusion From:
*Popup menu, Default: Alpha
*.选择 Occlusion 图像中控制光晕遮挡的通道。
*None:
*光晕完全不被遮挡。*Luma:
*光晕被 Occlusion 素材的亮度遮挡。黑色区域显示光晕，白色区域遮挡光晕。*Alpha:
*光晕被 Occlusion 素材的 Alpha 通道遮挡。如果你有一个应该出现在光晕前方的 RGBA 物体素材，请使用此选项。

Invert Occlusion:
*Check-box, Default:
*off.如果启用，反转 Occlusion 素材，使光晕被黑色区域而非白色区域遮挡。

Use Color:
*Check-box, Default:
*on.如果启用，根据 Occlusion 素材的颜色为光晕着色。当热点经过透明物体后方时，可用于创建彩色玻璃效果。

Diffraction Glow:
*Default:
*0,
*Range:
*0 or greater.在热点附近的遮挡物体边缘创建辉光，模拟光线绕过物体边缘溢出的效果。此参数控制辉光的亮度。

Glow Width:
*Default:
*0.4,
*Range:
*0 or greater.衍射辉光的宽度。

Glow Color:
*Default rgb:
*[1 1 1].衍射辉光的颜色。

Glow Radius:
*Default:
*0.2,
*Range:
*0 or greater.
衍射辉光可见的距热点距离。辉光会随着距热点距离增加而柔和地衰减。此半径之外的物体不会产生任何辉光。

### Occlusion Triggers Parameters:

Occlusion Scale Brightness:
*Default:
*1,
*Range:
*0 or greater.启用一个触发器，在光晕被遮挡时调整其亮度。当热点被遮挡 50% 时，亮度会上升到此值，然后回落到正常亮度。同时，光晕会因遮挡而淡出，因此其亮度会在最后快速下降。与 Occlusion Softness 调高时配合使用效果最佳。

Occlusion Scale Widths:
*Default:
*1,
*Range:
*0 or greater.启用一个触发器，在光晕被遮挡时调整其大小。当热点被遮挡 50% 时，光晕会扩展到此数值，然后缩回正常大小。同时，光晕会因遮挡而淡出，因此最后看起来会快速缩小。与 Occlusion Softness 调高时配合使用效果最佳。

Occlusion Layers:
*Default:
*0,
*Range:
*0 or greater.用于遮挡光晕热点的 AE 图层。

Performance:
*Popup menu, Default: full flare
*.决定是渲染所有元素还是仅渲染选定的元素。某些元素在光晕设计器中被选为对光晕外观很重要的元素。仅使用优先级渲染应该能呈现真实 LensFlare 的外观和感觉，用于预览目的，但渲染速度比完整光晕更快。
*full flare:
*渲染所有 LensFlare 元素。*priority only:
*仅渲染 LensFlare 元素的子集以提高性能。

Opacity:
*Popup menu, Default: Normal
*.决定处理不透明度/透明度的方法。
*All Opaque:
*当输入图像完全不透明且没有透明度 (alpha=1) 时，使用此选项可稍微加快渲染速度。*Normal:
*正常处理不透明度。*As Premult:
*按图像已为预乘形式处理（颜色已按不透明度缩放）。此选项的渲染速度也略快于 Normal 模式，但结果也将为预乘形式，有时不太正确。

Show Scale Widths:
*Check-box, Default:
*on.打开或关闭用于调整 Scale Widths 和 Rel Heights 参数的屏幕界面控件。此参数仅出现在 AE 和 Premiere 上，这些平台支持屏幕控件。

Show Hotspot:
*Check-box, Default:
*on.打开或关闭用于调整 Hotspot 参数的屏幕用户界面。此参数仅出现在 AE 和 Premiere 上，这些平台支持屏幕控件。

Show Pivot:
*Check-box, Default:
*on.
打开或关闭用于调整 Pivot 参数的屏幕用户界面。此参数仅出现在 AE 和 Premiere 上，这些平台支持屏幕控件。参见
[Motion Blur](/en/sapphire/#motion-blur) 的一般信息