---
title: RackDfComp
---

## S_RackDfComp

Composites the Foreground over the Background
while defocusing both layers by different amounts. The Foreground
alpha channel is used as the matte. If the Middle input is provided, it
is composited between the Foreground and Background.

In the Sapphire Blur+Sharpen effects submenu.

![RackDfComp](../_static/RackDfComp.jpg)


### Inputs:

- **Foreground**: The current layer. The clip to use as foreground, and the alpha channel of this clip is used as the matte.

- **Background**: Defaults to None. The clip to use as background.

- **Matte**: Defaults to None. The alpha channel of this input specifies the opacities of the Foreground input. If this input is not provided, the alpha channel of the Foreground input is used instead. This input can be affected by the Invert Matte or Matte Use parameters.

- **Middle**: Defaults to None. The clip to composite between the Foreground and Background.

- **Mid_Matte**: Defaults to None. The alpha channel of this input specifies the opacities of the Middle input. If this input is not provided, the alpha channel of the Middle input is used instead. This input can be affected by the Invert Matte or Matte Use parameters.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Defocus Foreground** (Default: 0.088, Range: 0 or greater)
  The amount to defocus the Foreground and its Matte. This parameter can be adjusted using the Fg Defocus Widget.

- **Defocus Background** (Default: 0, Range: 0 or greater)
  The amount to defocus the Background. This parameter can be adjusted using the Bg Defocus Widget.

- **Defocus Middle** (Default: 0, Range: 0 or greater)
  The amount to defocus the Middle and its Matte.

- **Rel Height** (Default: 1, Range: 0.01 or greater)
  The relative height of the iris shape. If it is not 1, circles become ellipses, etc.

- **Shape** (Popup menu, Default: Circle)
  Determines the shape of the simulated camera iris.
  - **Circle**: round.
  - **3 sides**: triangle.
  - **4 sides**: square.
  - **5 sides**: pentagon.
  - **6 sides**: hexagon.
  - **7 sides**: etc.

- **Show Shape** (Check-box, Default: off)
  Show the iris shape instead of the defocused image.

- **Roundness** (Default: 0, Range: any)
  Modifies the shape of the simulated camera iris. A value of 1 produces a circle; 0 gives a flat-sided polygon with a number of sides given by the Shape parameter. Less than 0 causes the sides to squeeze inward giving a star shape, while a value greater than 1 causes the corners to squeeze inward, giving a flowery shape. Has no effect if the Shape is set to Circle.

- **Rotate** (Default: 0, Range: any)
  Rotates the iris shape.

- **Bokeh** (Default: 0, Range: any)
  Softens the outer edge of the iris shape, which gives a softer look to the defocused highlights. A negative value darkens the center of the iris shape, producing a ring-like defocus shape.

- **Lens Noise** (Default: 0, Range: 0 or greater)
  Increase to add noise to the iris shape, dirtying up the defocus a little. Can make the result more realistic. Turn up past 1 for a more stylistic result.

- **Noise Freq** (Default: 40, Range: 0.01 or greater)
  The frequency of the added noise. Ignored if Lens Noise is zero.

- **Noise Freq Rel X** (Default: 1, Range: 0.01 or greater)
  The relative horizontal frequency of the added iris noise. Increase to stretch it vertically or decrease to stretch it horizontally.

- **Noise Seed** (Default: 0.123, Range: 0 or greater)
  The seed value for the added noise. To make the noise appear different on each frame, animate this to be different on each frame. The actual value doesn't matter; only that it's different.

- **Use Gamma** (Default: 1, Range: 0.1 or greater)
  Values above 1 cause highlights in the source clip to keep their brightness after the defocus is applied.

- **Matte Gamma** (Default: 1, Range: 0.1 or greater)
  The gamma value to use for the defocus of the Matte.

- **Boost Highlights** (Default: 0, Range: 0 or greater)
  The amount to increase the luma of the highlights in the source clip. Increase this parameter to blow out the highlights without affecting the darks or mid-tones.

- **Hilight Threshold** (Default: 0.9, Range: 0 or greater)
  The minimum luma value for highlights. Pixels brighter than this will be brightened according to the Boost Highlights parameter.

- **Comp Premult** (Check-box, Default: on)
  Disable this if you have provided a separate Matte input and the Foreground pixel values have not been pre-multiplied by this Matte.

- **Matte Use** (Popup menu, Default: Alpha)
  Determines how the Foreground or Matte input channels are used to make a monochrome matte.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Edge Mode** (Popup menu, Default: Reflect)
  Determines the behavior when accessing areas outside the source image.
  - **Transparent**: Areas outside the source image are treated as transparent, which can produce
transparency around the edges of the image.
Select this for fastest rendering.
  - **Repeat**: Repeats the last pixel outside the border of the image.
  - **Reflect**: Reflects the image outside the border.

- **Soft Borders** (Check-box, Default: off)
  If enabled, transparent borders are added to the input image before processing. This allows the result to include soft edges beyond the original image size. When off, the effect only occurs within the frame and the result will retain an edge at the borders.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Fg Defocus** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Defocus Foreground parameter. Its value should first be made positive to adjust this more easily.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Bg Defocus** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Defocus Background parameter. Its value should first be made positive to adjust this more easily.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

