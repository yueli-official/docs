---
title: ZapTo
---

## S_ZapTo

从指定点向 ToObject 输入片段中物体的边缘生成分叉闪电，并将其渲染在背景输入上。使用 Show:Edges 选项可以在调整 Threshold 和 Blur To Obj 参数时查看目标边缘。

在 Sapphire Render 效果子菜单中。

![ZapTo](../_static/ZapTo.jpg)


### Inputs:

- **Background**: 当前图层。用作背景的片段。如果未选择，主输入（当前图层）也用作背景。

- **ToObject**: 默认为无。提取此片段中物体的边缘，闪电连接到这些边缘上面向起始点的点。

- **Mask**: 默认为无。如果提供，此输入用于控制哪些边缘用于吸引闪电。黑色区域中找到的边缘将不被使用。白色区域中的边缘正常使用。灰色区域中的边缘将被缩减。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于追踪素材和生成蒙版。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此量模糊 Mocha 蒙版。可用于柔化蒙版的边缘或量化伪影，并平滑时间偏移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 蒙版的强度。较低的值会减弱效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  如果启用，在应用效果之前反转 Mocha 蒙版的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 蒙版。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 蒙版的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 蒙版的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 蒙版的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素量膨胀或腐蚀 Mocha 蒙版。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下膨胀 Mocha 蒙版，便于快速调整。
  - **High**: 在高质量模式下膨胀 Mocha 蒙版，获得更好的蒙版形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 蒙版，将效果应用于整个源片段。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 蒙版本身。

- **Combine Masks** (Popup menu, Default: Union)
  当同时提供 Mocha 蒙版和输入蒙版时，确定如何组合它们。
  - **Union**: 使用两个蒙版共同覆盖的区域。
  - **Intersect**: 使用两个蒙版重叠的区域。
  - **Mocha Only**: 忽略输入蒙版，仅使用 Mocha 蒙版。

- **Surface Points** (Integer, Default: 10, Range: 1 to 500)
  沿边缘连接闪电的点数。这些表面点按 ToObject 输入中各形状的比例分配。如果请求的表面点数等于 ToObject 输入中独立形状的数量，则每个形状连接一条闪电分叉。

- **Bolts** (Integer, Default: 1, Range: 1 to 200)
  要绘制的独立分叉闪电数量，每条连接起始位置和边缘点。

- **Start** (X & Y, Default: [0 0], Range: any)
  闪电的起始位置。

- **Max Dist** (Default: 2, Range: 0 or greater)
  表面点距起始位置的最大距离。超过此距离的边缘将被忽略。

- **Threshold** (Default: 0.5, Range: 0 or greater)
  用于确定边缘位置的值。比此值暗的物体将被忽略。在平滑物体上，较大的阈值会将边缘向内移动使形状变小，较小的值会将边缘向外移动。您可以在调整此参数时使用 Show Edges 选项直接查看边缘图像。

- **Blur ToObj** (Default: 0.0224, Range: 0 or greater)
  在查找边缘之前模糊 ToObject 输入片段。这有助于去除噪点并减少独立形状的数量。您可以在调整此参数时使用 Show Edges 选项直接查看边缘图像。

- **ToObj Use** (Popup menu, Default: Luma)
  确定使用 ToObject 输入的哪个通道。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Show** (Popup menu, Default: Result)
  选择效果的输出内容。
  - **Result**: 显示背景上的正常闪电结果。
  - **Edges**: 显示目标边缘图像。在调整 Threshold 和 Blur To Obj 参数时查看此图像很有用。

- **Bolt Width** (Default: 0.07, Range: 0 or greater)
  闪电的宽度。

- **Vary Width** (Default: 0, Range: 0 to 1)
  闪电沿其长度方向宽度的随机变化量。

- **End Pointiness** (Default: 0.25, Range: 0 to 1)
  确定闪电末端的尖锐程度。如果为 0，整条闪电宽度相等。如果为 1，闪电将沿整个长度逐渐变细形成尖端。如果为 0.5，闪电将从起点和终点的中间位置开始变细。

- **Wiggle Start** (Default: 0, Range: 0 or greater)
  默认情况下闪电会随时间自动摆动。此参数为这些闪电扰动提供起始偏移量。

- **Wiggle Speed** (Default: 1, Range: 0 or greater)
  闪电随时间自动扰动的速度。要制作速度变化的动画，请将此值设为零并改为对 Wiggle Start 参数进行动画。

- **Jitter Frames** (Integer, Default: 0, Range: 0 or greater)
  如果为 0，每个处理帧使用相同的随机闪电。如果为 1，每帧使用不同的随机闪电。如果为 2，每隔一帧使用新的随机闪电，依此类推。

- **Rand Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化随机数生成器。实际的种子值并不重要，但不同的种子会产生不同的随机闪电，相同的值应产生可重复的结果。

- **Wrinkle Amp** (Default: 1, Range: 0 or greater)
  缩放闪电中的褶皱量。减小可获得更直更平滑的闪电，增大可获得更弯曲的闪电。

- **Branchiness** (Default: 1, Range: 0 to 20)
  缩放从主闪电分支出的额外闪电数量。设为 0 可获得没有额外分支的基本闪电。

- **Branch Angle** (Default: 65, Range: 0 to 180)
  随机分支相对于其所分支的闪电的最大角度。如果为 0，分支将更贴近主闪电方向。值越大，分支越垂直于主闪电。

- **Branch Length** (Default: 0.5, Range: 0 to 3)
  分支相对于起点和终点之间距离的缩放长度。


### Glow Parameters:

Glow Bright:
*Default:
*2,
*Range:
*0 or greater.缩放应用于闪电的辉光亮度。

Glow Color:
*Default rgb:
*[0.5 0.5 1].应用于闪电的辉光颜色。

Glow Width:
*Default:
*0.224,
*Range:
*0 or greater.应用于闪电的辉光宽度。

Glow Width Red:
*Default:
*0.5,
*Range:
*0 or greater.辉光的相对红色宽度。

Glow Width Grn:
*Default:
*1,
*Range:
*0 or greater.辉光的相对绿色宽度。

Glow Width Blue:
*Default:
*1.5,
*Range:
*0 or greater.辉光的相对蓝色宽度。

Affect Alpha:
*Default:
*1,
*Range:
*0 or greater.
如果此值为正，输出 Alpha 通道将包含来自闪电及其辉光的一些不透明度。红、绿、蓝亮度的最大值按此值缩放，并与每个像素处的背景 Alpha 合成。

### Other Parameters:

Zap Bright:
*Default:
*1,
*Range:
*0 or greater.缩放闪电的亮度。

Zap Color:
*Default rgb:
*[1 1 1].闪电的颜色。如果您想保持闪电本身为亮白色，可以通过调整 Glow Color 来影响感知颜色。

Start Offset:
*Default:
*0,
*Range:
*0 to 1.从起点开始绘制闪电的偏移量。这对于制作闪电打击的动画很有用。

Length:
*Default:
*1,
*Range:
*0 to 1.闪电的长度，从 Start Offset 开始。如果小于 1，闪电将不会从起点到终点完全绘制。这对于制作闪电打击的动画很有用。

Bg Brightness:
*Default:
*1,
*Range:
*0 or greater.在与闪电合成之前缩放背景的亮度。如果为 0，结果将只包含黑色背景上的闪电图像。

Combine:
*Popup menu, Default: Screen
*.确定闪电和辉光与背景的合成方式。
*Screen:
*执行混合功能，有助于防止过亮的结果。*Add:
*将闪电添加到背景上。在明亮背景上会产生更亮的辉光。*Zap Only:
*在黑色背景上显示闪电，忽略源图像。

Blur Matte:
*Default:
*0.05,
*Range:
*0 or greater.在使用前按此量模糊蒙版输入。这可以提供蒙版区域和非蒙版区域之间更平滑的过渡。除非提供了蒙版输入，否则无效。

Invert Matte:
*Check-box, Default:
*off.如果开启，反转蒙版输入，使效果应用于蒙版为黑色的区域而非白色区域。除非提供了蒙版输入，否则无效。

Matte Use:
*Popup menu, Default: Luma
*.确定如何使用蒙版输入通道来生成单色蒙版。
*Luma:
*使用 RGB 通道的亮度。*Alpha:
*仅使用 Alpha 通道。

Atmosphere Amp:
*Default:
*0,
*Range:
*0 or greater.大气效果模拟电击效果穿过尘埃大气时被照亮或遮蔽的效果。此参数调整大气效果的量或幅度。零值产生平滑的电击效果，较高值产生更多尘埃感。

Atmosphere Freq:
*Default:
*2,
*Range:
*0.1 to 20.控制大气噪点的空间频率。调高可获得更精细的细节，调低可获得更宽泛的整体变化。

Atmosphere Detail:
*Default:
*0.7,
*Range:
*0 to 1.控制大气模拟中精细细节的量。减小可获得更平滑的大气效果，增大可获得更粗糙或颗粒感的外观。

Atmosphere Seed:
*Default:
*0.123,
*Range:
*0 or greater.用于初始化大气噪点的随机数生成器。实际的种子值并不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

Atmosphere Speed:
*Default:
*1,
*Range:
*any.大气中的云状噪点会像真实的尘埃云一样随时间变化；此参数控制云图案随时间变化的速度。设为零可获得静态图案。

Opacity:
*Popup menu, Default: Normal
*.确定处理不透明度/透明度的方法。
*All Opaque:
*当输入图像完全不透明（无透明度，alpha=1）时使用此选项可略微加快渲染速度。*Normal:
*正常处理不透明度。*As Premult:
*按照图像已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比正常模式略快，但结果也将是预乘形式，有时不太准确。

Show Max Dist:
*Check-box, Default:
*on.
打开或关闭用于调整 Start 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，因为这些平台支持屏幕控件。参见
[Motion Blur](/en/sapphire/#motion-blur) 的一般信息
