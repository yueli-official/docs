---
title: Grunge
---

## S_Grunge

模拟多种不同类型的做旧效果，包括污垢、污渍、斑点、污迹、划痕和油漆。最多可以组合三种不同类型的做旧效果。提供了用于整体调整所有做旧效果的主控制，以及用于调整每个做旧效果集合外观的详细控制。

位于 Sapphire Render 效果子菜单中。

![Grunge](../_static/Grunge.jpg)


### Inputs:

- **Background**: 当前图层。用作背景的素材。

- **Matte**: 默认为无。如果提供，模糊仅在此输入的亮区所指定的源素材区域上执行。此蒙版之外的像素不会被模糊，也不会对其内部的模糊像素产生贡献。此输入可通过 Invert Matte 或 Matte Use 参数进行调整。


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
  如果启用，在应用效果之前反转 Mocha 遮罩的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素数扩展或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Dilate Mocha 是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下扩展 Mocha 遮罩，用于快速调整。
  - **High**: 在高质量模式下扩展 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当两个遮罩同时提供给效果时，决定如何组合 Mocha 遮罩和输入遮罩。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Stamp Density** (Default: 500, Range: 0 or greater)
  画面上印章的整体数量。增大以获得更多印章，减小以获得更少。

- **Stamp Size** (Default: 1, Range: 0 or greater)
  缩放做旧印章的整体大小。

- **Stamp Opacity** (Default: 0.8, Range: 0 to 1)
  做旧印章的整体不透明度。

- **Stamp Brightness** (Default: 1, Range: 0 or greater)
  控制做旧印章的亮度。

- **Stamp Center** (X & Y, Default: [0 0], Range: any)
  做旧印章中心的整体位置。

- **Stamp Uses Mocha** (Check-box, Default: off)
  控制印章中心是由 Stamp Center 参数控制，还是跟随在 Mocha 内部跟踪的 Stamp Center 点。

- **Smooth Stamp Track** (Integer, Default: 0, Range: 0 or greater)
  控制稳定 Mocha 点跟踪时平均多少个点。

- **Fade Softness** (Default: 0.3, Range: 0 to 1)
  控制单个做旧印章淡入淡出的速度。设为零时，印章会突然出现和消失。设为一时，它们会淡入淡出。印章在闪烁时、在围绕背景的做旧边框边缘以及蒙版的灰色区域会淡出。

- **Blink Amount** (Default: 0, Range: 0 to 1)
  控制单帧上有多少印章在闪烁。闪烁量为 0 时所有印章都可见。闪烁量为 1 时大约一半的印章在给定帧上会闪烁。

- **Blink Coherence** (Default: 0.5, Range: 0 to 1)
  改变闪烁印章的图案。

- **Blink Speed** (Default: 6, Range: 0 or greater)
  控制印章闪烁的速度。

- **Frame Amount** (Default: 0.35, Range: 0 to 1)
  控制边框或边界内做旧效果的亮度。如果边框量设为非零值，Grunge 将在边框中心周围创建一个做旧边框。边框量为 1 时，边框内的印章完全不可见；边框量为 0 时，没有边框。

- **Frame Softness** (Default: 0.5, Range: 0 or greater)
  边框柔边的宽度。较大的值使做旧边框在边缘处的亮度渐变更柔和。

- **Frame Center** (X & Y, Default: [0 0], Range: any)
  做旧边框中心的位置。此参数可以使用 Frame Center 控件调整。

- **Frame Uses Mocha** (Check-box, Default: off)
  控制边框中心是由 Frame Center 参数控制，还是跟随在 Mocha 内部跟踪的 Frame Center 点。

- **Smooth Frame Track** (Integer, Default: 0, Range: 0 or greater)
  控制稳定 Mocha 点跟踪时平均多少个点。

- **Frame Radius** (Default: 1, Range: 0 or greater)
  从中心到应用做旧边框的距离。此参数可以使用 Frame Radius 控件调整。

- **Frame Rel Height** (Default: 0.75, Range: 0.1 or greater)
  边框形状的相对垂直大小。增大以获得更高的形状，减小以获得更宽的形状。

- **Invert Frame** (Check-box, Default: off)
  如果启用，在边框中心显示做旧效果，而不是在边框边缘。

- **Stamp1** (Popup menu, Default: Garage Floor)
  要应用的做旧风格。最多可选择三种风格。
  - **None**: 无做旧效果。
  - **Plaster**: 模拟石膏的大面积低细节做旧。
  - **Garage Floor**: 类似路面的大面积斑点做旧。
  - **Speckles**: 由小型、大小相近的做旧圆点组成的群组，类似喷漆。
  - **Paint Spray**: 由中小型油漆飞溅组成的群组，类似斑点。
  - **Paint Splatters**: 长条状的油漆飞溅。
  - **Hairline Cracks**: 细长、弯曲、无分支的裂纹。
  - **Tile Cracks**: 长直锯齿状裂纹。
  - **Pavement Cracks**: 分支众多、宽度变化很大的裂纹。
  - **Hairs**: 各种大小的卷发和直发。
  - **Scratches**: 偏向对角线方向的直线划痕。
  - **Frost**: 霜冻斑块。
  - **Glass Cracks**: 蛛网状玻璃裂纹斑块。
  - **Clouds**: 云朵状，类似烟雾和水彩滴落。
  - **Smoke**: 缕缕烟雾，类似云朵和水彩滴落。
  - **Splotches**: 带有大量向外扩散条纹的油漆斑点。
  - **Corrosion**: 类似锈蚀损坏的斑块。
  - **Watercolor Drops**: 低细节的做旧飞溅，类似云朵和烟雾。
  - **Dust**: 形状各异的小颗粒做旧，类似碎片。
  - **Stains**: 咖啡渍和水渍。
  - **Flecks**: 形状和大小各异的中小型做旧颗粒，类似灰尘。
  - **Circles**: 大型不完整圆环。
  - **Geometric Lines**: 大型锐利线条和棱角做旧。
  - **Liquid Lino**: 湿润外观的斑块。
  - **Geometric Star Map**: 星座外观的做旧。
  - **Moon Dust**: 小点条纹。
  - **Starfield Spray**: 小点群组。
  - **Spray Paint Stars**: 圆点和尖星群组。
  - **Spray Paint**: 各种小点群组。
  - **Splatter And Spray**: 油漆涂抹和飞溅。
  - **Hand Painted**: 各种油漆斑块。
  - **Ink Drops**: 各种墨滴图案。
  - **Large Sponges**: 海绵印迹图案。
  - **Paper Sponges**: 褶皱纸张。
  - **Decaying Damask**: 各种褪色锦缎图案。
  - **Stylized Blooms**: 风格化盛开的花朵。
  - **Oil Paint Floral**: 花朵群组。

- **Stamp1 Rel Density** (Default: 1, Range: 0 or greater)
  缩放特定印章集合的密度。

- **Stamp1 Color1** (Default rgb: [0.01 0.01 0.01])
  印章集合颜色范围的起始色。

- **Stamp1 Color2** (Default rgb: [0.3 0.3 0.3])
  印章集合颜色范围的结束色。每个做旧元素将在 color1 和 color2 之间随机取色。

- **Stamp1 Rel Brightness** (Default: 1, Range: 0 or greater)
  缩放特定印章集合的亮度。

- **Stamp1 Rel Opacity** (Default: 1, Range: 0 or greater)
  缩放特定印章集合的不透明度。

- **Stamp1 Rel Size** (Default: 1, Range: 0 or greater)
  缩放特定印章集合的大小。

- **Vary Stamp1 Brightness** (Default: 0, Range: 0 or greater)
  各做旧元素之间亮度的变化量。

- **Vary Stamp1 Opacity** (Default: 0.75, Range: 0 or greater)
  各做旧元素之间不透明度的变化量。

- **Vary Stamp1 Size** (Default: 0.5, Range: 0 or greater)
  各做旧元素之间大小的变化量。

- **Stamp2** (Popup menu, Default: Paint Spray)
  要应用的做旧风格。最多可选择三种风格。
  - **None**: 无做旧效果。
  - **Plaster**: 模拟石膏的大面积低细节做旧。
  - **Garage Floor**: 类似路面的大面积斑点做旧。
  - **Speckles**: 由小型、大小相近的做旧圆点组成的群组，类似喷漆。
  - **Paint Spray**: 由中小型油漆飞溅组成的群组，类似斑点。
  - **Paint Splatters**: 长条状的油漆飞溅。
  - **Hairline Cracks**: 细长、弯曲、无分支的裂纹。
  - **Tile Cracks**: 长直锯齿状裂纹。
  - **Pavement Cracks**: 分支众多、宽度变化很大的裂纹。
  - **Hairs**: 各种大小的卷发和直发。
  - **Scratches**: 偏向对角线方向的直线划痕。
  - **Frost**: 霜冻斑块。
  - **Glass Cracks**: 蛛网状玻璃裂纹斑块。
  - **Clouds**: 云朵状，类似烟雾和水彩滴落。
  - **Smoke**: 缕缕烟雾，类似云朵和水彩滴落。
  - **Splotches**: 带有大量向外扩散条纹的油漆斑点。
  - **Corrosion**: 类似锈蚀损坏的斑块。
  - **Watercolor Drops**: 低细节的做旧飞溅，类似云朵和烟雾。
  - **Dust**: 形状各异的小颗粒做旧，类似碎片。
  - **Stains**: 咖啡渍和水渍。
  - **Flecks**: 形状和大小各异的中小型做旧颗粒，类似灰尘。
  - **Circles**: 大型不完整圆环。
  - **Geometric Lines**: 大型锐利线条和棱角做旧。
  - **Liquid Lino**: 湿润外观的斑块。
  - **Geometric Star Map**: 星座外观的做旧。
  - **Moon Dust**: 小点条纹。
  - **Starfield Spray**: 小点群组。
  - **Spray Paint Stars**: 圆点和尖星群组。
  - **Spray Paint**: 各种小点群组。
  - **Splatter And Spray**: 油漆涂抹和飞溅。
  - **Hand Painted**: 各种油漆斑块。
  - **Ink Drops**: 各种墨滴图案。
  - **Large Sponges**: 海绵印迹图案。
  - **Paper Sponges**: 褶皱纸张。
  - **Decaying Damask**: 各种褪色锦缎图案。
  - **Stylized Blooms**: 风格化盛开的花朵。
  - **Oil Paint Floral**: 花朵群组。

- **Stamp2 Rel Density** (Default: 0.5, Range: 0 or greater)
  缩放特定印章集合的密度。

- **Stamp2 Color1** (Default rgb: [0.1 0 0])
  印章集合颜色范围的起始色。

- **Stamp2 Color2** (Default rgb: [0.2 0.1 0.05])
  印章集合颜色范围的结束色。每个做旧元素将在 color1 和 color2 之间随机取色。

- **Stamp2 Rel Brightness** (Default: 1, Range: 0 or greater)
  缩放特定印章集合的亮度。

- **Stamp2 Rel Opacity** (Default: 1, Range: 0 or greater)
  缩放特定印章集合的不透明度。

- **Stamp2 Rel Size** (Default: 1, Range: 0 or greater)
  缩放特定印章集合的大小。

- **Vary Stamp2 Brightness** (Default: 0, Range: 0 or greater)
  各做旧元素之间亮度的变化量。

- **Vary Stamp2 Opacity** (Default: 0.75, Range: 0 or greater)
  各做旧元素之间不透明度的变化量。

- **Vary Stamp2 Size** (Default: 0.5, Range: 0 or greater)
  各做旧元素之间大小的变化量。

- **Stamp3** (Popup menu, Default: None)
  要应用的做旧风格。最多可选择三种风格。
  - **None**: 无做旧效果。
  - **Plaster**: 模拟石膏的大面积低细节做旧。
  - **Garage Floor**: 类似路面的大面积斑点做旧。
  - **Speckles**: 由小型、大小相近的做旧圆点组成的群组，类似喷漆。
  - **Paint Spray**: 由中小型油漆飞溅组成的群组，类似斑点。
  - **Paint Splatters**: 长条状的油漆飞溅。
  - **Hairline Cracks**: 细长、弯曲、无分支的裂纹。
  - **Tile Cracks**: 长直锯齿状裂纹。
  - **Pavement Cracks**: 分支众多、宽度变化很大的裂纹。
  - **Hairs**: 各种大小的卷发和直发。
  - **Scratches**: 偏向对角线方向的直线划痕。
  - **Frost**: 霜冻斑块。
  - **Glass Cracks**: 蛛网状玻璃裂纹斑块。
  - **Clouds**: 云朵状，类似烟雾和水彩滴落。
  - **Smoke**: 缕缕烟雾，类似云朵和水彩滴落。
  - **Splotches**: 带有大量向外扩散条纹的油漆斑点。
  - **Corrosion**: 类似锈蚀损坏的斑块。
  - **Watercolor Drops**: 低细节的做旧飞溅，类似云朵和烟雾。
  - **Dust**: 形状各异的小颗粒做旧，类似碎片。
  - **Stains**: 咖啡渍和水渍。
  - **Flecks**: 形状和大小各异的中小型做旧颗粒，类似灰尘。
  - **Circles**: 大型不完整圆环。
  - **Geometric Lines**: 大型锐利线条和棱角做旧。
  - **Liquid Lino**: 湿润外观的斑块。
  - **Geometric Star Map**: 星座外观的做旧。
  - **Moon Dust**: 小点条纹。
  - **Starfield Spray**: 小点群组。
  - **Spray Paint Stars**: 圆点和尖星群组。
  - **Spray Paint**: 各种小点群组。
  - **Splatter And Spray**: 油漆涂抹和飞溅。
  - **Hand Painted**: 各种油漆斑块。
  - **Ink Drops**: 各种墨滴图案。
  - **Large Sponges**: 海绵印迹图案。
  - **Paper Sponges**: 褶皱纸张。
  - **Decaying Damask**: 各种褪色锦缎图案。
  - **Stylized Blooms**: 风格化盛开的花朵。
  - **Oil Paint Floral**: 花朵群组。

- **Stamp3 Rel Density** (Default: 0.5, Range: 0 or greater)
  缩放特定印章集合的密度。

- **Stamp3 Color1** (Default rgb: [0 0 0])
  印章集合颜色范围的起始色。

- **Stamp3 Color2** (Default rgb: [0.25 0.25 0.25])
  印章集合颜色范围的结束色。每个做旧元素将在 color1 和 color2 之间随机取色。

- **Stamp3 Rel Brightness** (Default: 1, Range: 0 or greater)
  缩放特定印章集合的亮度。

- **Stamp3 Rel Opacity** (Default: 1, Range: 0 or greater)
  缩放特定印章集合的不透明度。

- **Stamp3 Rel Size** (Default: 1, Range: 0 or greater)
  缩放特定印章集合的大小。

- **Vary Stamp3 Brightness** (Default: 0, Range: 0 or greater)
  各做旧元素之间亮度的变化量。

- **Vary Stamp3 Opacity** (Default: 0.75, Range: 0 or greater)
  各做旧元素之间不透明度的变化量。

- **Vary Stamp3 Size** (Default: 0.5, Range: 0 or greater)
  各做旧元素之间大小的变化量。

- **Emboss Bumps Scale** (Default: 0.15, Range: 0 or greater)
  缩放凸起的振幅。

- **Emboss Light Angle** (Default: 135, Range: any)
  调整浮雕的光照角度。此参数可以使用 Emboss Light Angle 控件调整。

- **Emboss Smooth** (Default: 0.0001, Range: 0 or greater)
  平滑图像的细小细节，使它们不会像大特征那样被强烈浮雕。设为 0 则浮雕所有细节。增大以平滑更多细节。

- **Emboss Threshold** (Default: 0.5, Range: -0.5 to 1.5)
  亮度高于阈值的做旧将形成凸起，暗于阈值的做旧将形成凹陷。

- **Blur Grunge** (Default: 0, Range: 0 or greater)
  模糊做旧效果。增大以获得更多模糊。不影响背景。

- **Blur Grunge Rel** (X & Y, Default: [1 1], Range: 0 or greater)
  缩放模糊的宽度。

- **Shake Amplitude** (Default: 0, Range: 0 or greater)
  应用于印章的随机抖动运动的量。

- **Shake Amplitude Rel X** (Default: 1, Range: 0 or greater)
  相对水平抖动量。

- **Shake Amplitude Rel Y** (Default: 1, Range: 0 or greater)
  相对垂直抖动量。

- **Shake Frequency** (Default: 60, Range: 0 or greater)
  增大以获得更快的抖动，减小以获得更慢的抖动。（如果对频率值进行动画，请注意产生的抖动频率也受到值变化速率的影响。）

- **Per Stamp Amplitude** (Default: 0, Range: 0 or greater)
  独立应用于每个印章的随机抖动量。

- **Per Stamp Frequency** (Default: 5, Range: 0 or greater)
  每个印章抖动的频率。增大以获得更快的抖动，减小以获得更慢的抖动。

- **Combine** (Popup menu, Default: Comp)
  决定做旧图像如何与背景组合。
  - **Grunge Only**: 仅显示做旧图像，无背景。
  - **Comp**: 将做旧图像合成在背景之上。
  - **Mult**: 可用作蒙版图像的"交集"操作。白色是正片叠底的恒等值，当一个图像包含白色时另一个不受影响，因此结果仅在两个输入都为白色的地方包含白色。
  - **Add**: 将做旧图像添加到背景。
  - **Screen**: 执行混合功能，有助于防止过亮的结果。
  - **Difference**: 类似于减去，但使用结果的绝对值，倾向于在范围内产生更多结果颜色。可用于选择两个蒙版图像中一个或另一个为白色但不同时为白色的区域。
  - **Subtract**: 从背景中减去做旧图像。
  - **Overlay**: 使用叠加功能组合前景和背景。
  - **Hard Light**: 类似于叠加，但前景和背景互换。
  - **Soft Light**: 根据前景使背景变暗或变亮。
  - **Color Dodge**: 根据做旧图像使背景变亮。
  - **Color Burn**: 根据做旧图像使背景变暗。
  - **Darken**: 做旧图像和背景的最小值。也可用作"交集"操作，结果与正片叠底略有不同。
  - **Lighten**: 做旧图像和背景的最大值。也可用作"并集"操作，结果与滤色略有不同。
  - **Exclusion**: 类似于差值，但结果更平滑。
  - **Linear Dodge**: 将做旧图像添加到背景并将结果限制在白色。
  - **Linear Burn**: 将做旧图像添加到背景但偏移使结果更暗。类似于正片叠底，与白色组合不变，与黑色组合产生黑色。
  - **Linear Light**: 根据前景是否超过 50% 灰色执行线性加深或线性减淡。

- **Use Bg Alpha** (Check-box, Default: off)
  启用时，忽略做旧效果生成的 Alpha，使用背景的 Alpha。

- **Scale Background** (Default: 1, Range: 0 or greater)
  在与做旧效果组合之前缩放背景的亮度。如果为 0，结果将仅包含黑色背景上的做旧图像。

- **Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化印章位置、大小和变化的随机数生成器。实际的种子值并不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Blur Matte** (Default: 0, Range: 0 or greater)
  在使用前按此数值模糊蒙版输入。可以在蒙版区域和非蒙版区域之间提供更平滑的过渡。除非提供了蒙版输入，否则无效。

- **Invert Matte** (Check-box, Default: off)
  如果开启，反转蒙版输入，使效果应用于蒙版为黑色而非白色的区域。除非提供了蒙版输入，否则无效。

- **Matte Use** (Popup menu, Default: Luma)
  决定如何使用 Matte 输入通道生成单色蒙版。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Crop Grunge To Matte** (Check-box, Default: off)
  控制做旧效果是否沿遮罩边缘有清晰的边界。如果启用，做旧效果将裁剪到遮罩。这允许将做旧效果应用于输入素材的特定部分。如果禁用，每个做旧印章将根据遮罩进行检查。如果大部分做旧印章落在遮罩内，整个做旧印章将被添加。如果大部分做旧印章落在遮罩外，整个做旧印章将被隐藏。这允许类似暗角的做旧效果应用方式。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且没有透明度（alpha=1）时，使用此选项可略微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 将图像视为已经是预乘形式（颜色已按不透明度缩放）来处理。此选项的渲染速度也比 Normal 模式略快，但结果也将是预乘形式，有时不太正确。

- **Flip Stamps Vertically** (Check-box, Default: off)
  如果需要，垂直翻转印章以获得一致的外观。

- **Show Emboss Light Angle** (Check-box, Default: on)
  开启或关闭用于调整 Frame Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show Frame Radius** (Check-box, Default: on)
  开启或关闭用于调整 Frame Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

- **Show Frame Center** (Check-box, Default: on)
  开启或关闭用于调整 Frame Center 参数的屏幕用户界面。此参数仅在 AE 和 Premiere 中出现，这些软件支持屏幕控件。

