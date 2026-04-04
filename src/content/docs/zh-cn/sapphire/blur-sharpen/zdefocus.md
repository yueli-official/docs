---
title: ZDefocus
---

## S_ZDefocus

根据 ZBuffer 输入的深度值，对源素材的不同区域施加不同强度的失焦。将输入分成若干深度层，并根据各层深度应用不同的失焦量。使用此效果时，首先根据你的 Z 缓冲设置 ZBuffer: Black Is Near 或 White Is Near，然后调整 Focus Depth 与 Depth Of Field 以获得所需外观。为便于设置 Focus Depth，可使用 Show: In Focus Zone。

位于 Sapphire Blur+Sharpen 效果子菜单中。

![ZDefocus](../_static/ZDefocus.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。

- **ZBuffer**: 默认为无。包含每个 Source 像素深度值的输入素材。取值应在黑到白之间，且最好不要抗锯齿。通常黑色表示最远处，白色表示最近处，可通过 Z Buffer 参数进行调整。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Focal Depth** (Default: 0, Range: any)
  焦平面的深度；0 为近，1 为远。具有该 Z 值的区域将处于对焦状态。接近该深度的物体是否清晰取决于 Depth Of Field。可用 Show: In Focus Zone 显示以辅助调节。若该参数效果与预期相反，可用 Z Buffer 参数反转深度值。

- **Depth Of Field** (Default: 0.1, Range: 0 to 1)
  指定 Focal Depth 附近被视为清晰的深度范围宽度。例如 Focal Depth=0.5 且 Depth Of Field=0.2 时，Z 值在 0.4–0.6 的物体都将清晰。设为 0 则只有恰好在 Focal Depth 处的物体清晰。可用 Show: In Focus Zone 显示以辅助调节。

- **Defocus Width** (Default: 0.35, Range: 0 or greater)
  缩放总体失焦宽度。此参数可通过 Defocus Width Widget 调整。

- **Rel Height** (Default: 1, Range: 0.01 or greater)
  光圈形状的相对高度。若不为 1，则圆形会变为椭圆等。

- **Z Buffer Type** (Popup menu, Default: White is Near)
  解释 Z 缓冲中的取值方式。
  - **Black is Near**: Z 缓冲中黑色表示近处，白色表示远处。
  - **White is Near**: Z 缓冲中白色表示近处，黑色表示远处。

- **Show** (Popup menu, Default: Result)
  选择输出类型。
  - **Result**: 显示效果最终结果。
  - **In Focus Zone**: 高亮显示图像中的对焦区域，以便选择焦点与深度。
  - **Shape**: 显示光圈形状而非失焦图像。

- **Layers** (Integer, Default: 5, Range: 2 to 50)
  将源素材按深度分层的层数。层数越多处理越慢，但 Z 方向过渡更平滑。有时需要更多层以避免层间接缝可见。

- **Layer Mode** (Popup menu, Default: Interp)
  确定不同失焦层的合成方式。
  - **Comp**: 近处层合成在远处层之上。若不同深度的物体相互遮挡且深度图存在不连续，此方式通常更好，但可能更慢，且偶尔会看到层间伪影。
  - **Interp**: 依据深度图对各层进行插值。层间过渡更平滑，通常在深度图无剧烈变化时更佳。

- **Shape** (Popup menu, Default: Circle)
  确定模拟相机光圈的形状。
  - **Circle**: 圆形。
  - **3 sides**: 三角形。
  - **4 sides**: 正方形。
  - **5 sides**: 五边形。
  - **6 sides**: 六边形。
  - **7 sides**: 等等。

- **Roundness** (Default: 0, Range: any)
  修改模拟相机光圈的形状。值为 1 产生圆形；0 产生由 Shape 指定边数的平边多边形；小于 0 使边向内收缩产生星形；大于 1 使角向内收缩产生花状。若 Shape 为 Circle 则无效。

- **Rotate** (Default: 0, Range: any)
  旋转光圈形状。

- **Bokeh** (Default: 0, Range: any)
  软化光圈外缘，使失焦高光更柔。负值会使形状中心变暗，产生环状失焦。

- **Lens Noise** (Default: 0, Range: 0 or greater)
  增加光圈形状噪声，使失焦略显“脏”。可提升真实感。大于 1 可获得更风格化的效果。

- **Noise Freq** (Default: 40, Range: 0.01 or greater)
  添加噪声的频率。若 Lens Noise 为 0 则忽略。

- **Noise Freq Rel X** (Default: 1, Range: 0.01 or greater)
  添加的光圈噪声的相对水平频率。增大以纵向拉伸；减小以横向拉伸。

- **Noise Seed** (Default: 0.123, Range: 0 or greater)
  噪声种子。要让每帧噪声不同，可逐帧动画此值。

- **Use Gamma** (Default: 1, Range: 0.1 or greater)
  大于 1 时，源素材中的高光在失焦后保持其亮度。

- **Boost Highlights** (Default: 0, Range: 0 or greater)
  提升源素材高光的亮度。增大此值可在不影响暗部与中间调的情况下“拉亮”高光。

- **Highlight Threshold** (Default: 0.9, Range: 0 or greater)
  高光的最小亮度阈值。高于该阈值的像素将按 Boost Highlights 提升亮度。

- **Brightness** (Default: 1, Range: 0 or greater)
  缩放结果亮度。

- **Offset Darks** (Default: 0, Range: any)
  向较暗区域添加该灰度值。可为负以增加对比度。

- **Mix With Source** (Default: 0, Range: 0 to 1)
  在失焦结果与原始源之间插值。设为 1 可得到原图。

- **Width Rel Near** (Default: 1, Range: 0 or greater)
  缩放焦平面近侧区域的失焦宽度。

- **Width Rel Far** (Default: 1, Range: 0 or greater)
  缩放焦平面远侧区域的失焦宽度。

- **Fog Near** (Default: 0, Range: 0 to 1)
  近处雾效强度。

- **Fog Far** (Default: 0, Range: 0 to 1)
  远处雾效强度。

- **Fog Color** (Default rgb: [0.5 0.5 0.5])
  雾的颜色通常应与源素材的天空或背景相匹配。灰色用于薄雾，棕色用于烟尘，蓝色用于水下等。

- **Edge Mode** (Popup menu, Default: Reflect)
  确定访问源图像边界之外的行为。
  - **Transparent**: 边界之外视为透明，可能在图像边缘产生透明度；渲染最快。
  - **Repeat**: 重复图像边界外的最后一个像素。
  - **Reflect**: 在边界外镜像翻转图像。

- **Zbuffer Use** (Popup menu, Default: Luma)
  决定如何由 ZBuffer 输入通道生成单通道深度图。
  - **Luma**: 使用 RGB 通道的亮度。
  - **Alpha**: 仅使用 Alpha 通道。

- **Soft Borders** (Check-box, Default: off)
  若启用，在处理前为输入图像添加透明边框，从而允许结果包含超出原始图像尺寸的柔和边缘。关闭时效果仅在画面内发生，结果边缘将保留边界。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入完全不透明（alpha=1）时可稍微加快渲染。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按已预乘形式处理（颜色已按不透明度缩放）。渲染略快于 Normal，但结果也将是预乘形式，某些情况下精确性较差。

- **Show Defocus Width** (Check-box, Default: on)
  打开或关闭用于调整 Defocus Width 的屏幕控件。此参数仅在支持屏幕控件的 AE 与 Premiere 中出现。
