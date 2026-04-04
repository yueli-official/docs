---
title: DissolveDefocus
---

## S_DissolveDefocus

Transitions between two input clips while defocusing each.
The first clip is defocused and faded out while the second clip is
brought into focus and faded in. The Dissolve
Percent parameter should be animated to control the transition speed.

In the Sapphire Transitions effects submenu.

![DissolveDefocus](../_static/DissolveDefocus.jpg)


### Inputs:

- **Foreground**: The current layer. Starts the transition with this clip.

- **Background**: Defaults to None. Ends the transition with this clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Transition Dir** (Popup menu, Default: Dissolve Off to Bg)
  Selects the direction of the transition.
  - **Dissolve Off to Bg**: transitions from the current layer to the Background.
  - **Dissolve On from Bg**: transitions from the Background to the current layer.

- **Auto Trans** (Popup YES-NO, Default: No)
  If enabled, a transition is performed automatically between the first and last frames of the layer. If this is off, the transition is performed manually by animating the Dissolve Percent parameter.

- **Dissolve Percent** (Default: 0, Range: 0 to 1)
  Auto Trans must be disabled for this parameter to be used. It determines the transition ratio between the Foreground and Background inputs, and would normally be animated from 0 to 100 to perform a complete transition. The curve controlling this parameter can be adjusted for more detailed control over the timing of the dissolve.

- **Defocus Width** (Default: 0.8, Range: 0 or greater)
  The width of the defocus.

- **Defocus Rel From** (Default: 1, Range: 0 or greater)
  Scales the amount of defocus applied to the first clip. Set to 0 to fade out with no defocus.

- **Defocus Rel To** (Default: 1, Range: 0 or greater)
  Scales the amount of defocus applied to the second clip. Set to 0 to fade in with no defocus.

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
  The spatial frequency of the noise.

- **Noise Freq Rel X** (Default: 1, Range: 0.01 or greater)
  The relative horizontal frequency of the added iris noise. Increase to stretch it vertically or decrease to stretch it horizontally.

- **Noise Seed** (Default: 0.123, Range: 0 or greater)
  The seed value for the added noise. To make the noise appear different on each frame, animate this to be different on each frame. The actual value doesn't matter; only that it's different.

- **Use Gamma** (Default: 1, Range: 0.1 or greater)
  Values above 1 cause highlights in the source clip to keep their brightness after the defocus is applied.

- **Boost Highlights** (Default: 0, Range: 0 or greater)
  The amount to increase the luma of the highlights in the source clip. Increase this parameter to blow out the highlights without affecting the darks or mid-tones.

- **Hilight Threshold** (Default: 0.9, Range: 0 or greater)
  The minimum luma value for highlights. Pixels brighter than this will be brightened according to the Boost Highlights parameter.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Edge Mode** (Popup menu, Default: Reflect)
  Determines the behavior when accessing areas outside the source image.
  - **Transparent**: Areas outside the source image are treated as transparent, which can produce
transparency around the edges of the image.
Select this for fastest rendering.
  - **Repeat**: Repeats the last pixel outside the border of the image.
  - **Reflect**: Reflects the image outside the border.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Defocus Width** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Defocus Width parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

