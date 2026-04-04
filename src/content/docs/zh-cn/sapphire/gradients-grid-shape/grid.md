---
title: Grid
---

## S_Grid

Generates a grid of lines and combines it with a background clip. Adjust the
Latitude, Swing, and Roll parameters to rotate the grid on various axes,
and adjust Shift and Z Dist to translate and zoom.

In the Sapphire Render effects submenu.

![Grid](../_static/Grid.jpg)


### Inputs:

- **Background**: The current layer. The clip to draw the grid on.

- **Mask**: Defaults to None. Interpolate between the result and the Source input. White areas use the result of the effect. Black areas use the Source clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mocha Project** (Default: 0, Range: 0 or greater)
  Brings up the Mocha window for tracking footage and generating masks.

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  Blurs the Mocha Mask by this amount before using. This can be used to soften the edges or quantization artifacts of the mask, and smooth out the time displacements.

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  Controls the strength of the Mocha mask. Lower values reduce the intensity of the effect.

- **Invert Mocha** (Check-box, Default: off)
  If enabled, the black and white of the Mocha Mask are inverted before applying the effect.

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  Scales the Mocha Mask. 1.0 is the original size.

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  The relative horizontal size of the Mocha Mask.

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  The relative vertical size of the Mocha Mask.

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  Offsets the position of the Mocha Mask.

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  Dilates or erodes the Mocha Mask by this pixel amount before using.

- **Dilation Quality** (Popup menu, Default: Fast)
  Selects whether Dilate Mocha adusts quickly in default Fast mode or looks better in High quality mode.
  - **Fast**: Dilate Mocha in Fast mode for quick adjustments.
  - **High**: Dilate Mocha in High quality mode for a better looking mask shape.

- **Bypass Mocha** (Check-box, Default: off)
  Ignore the Mocha Mask and apply the effect to the entire source clip.

- **Show Mocha Only** (Check-box, Default: off)
  Bypass the effect and show the Mocha Mask itself.

- **Combine Masks** (Popup menu, Default: Union)
  Determines how to combine the Mocha Mask and Input Mask when both are supplied to the effect.
  - **Union**: Uses the area covered by both masks together.
  - **Intersect**: Uses the area that overlaps between the two masks.
  - **Mocha Only**: Ignore the Input Mask and only use the
Mocha Mask.

- **Boxes** (X & Y, Integer, Default: [24 16], Range: 1 or greater)
  The total number of grid cells in the horizontal and vertical directions.

- **Grid Size** (Default: 1, Range: 0 or greater)
  Scales the size of the grid object.

- **Grid Size X** (Default: 1, Range: 0 or greater)
  Scales the relative horizontal size of the grid.

- **Grid Size Y** (Default: 0.75, Range: 0 or greater)
  Scales the relative vertical size of the grid.

- **Shift** (X & Y, Default: [0 0], Range: any)
  Translates the grid by this amount.

- **Line Width** (Default: 1.16, Range: 0 or greater)
  Scales the thickness of all the grid lines.

- **H Line Rel Width** (Default: 1, Range: 0 or greater)
  Scales the relative thickness of the horizontal lines.

- **V Line Rel Width** (Default: 1, Range: 0 or greater)
  Scales the relative thickness of the vertical lines.

- **Major Line Spacing** (Integer, Default: 4, Range: 0 or greater)
  Thicker lines are drawn at each interval of this many lines. If zero, the major lines are disabled and all lines will be equal width.

- **Major Line Width** (Default: 2.5, Range: 1 or greater)
  The relative thickness of the major lines.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the grid color.

- **Color** (Default rgb: [1 1 1])
  The color of the grid.

- **Grid Opacity** (Default: 1, Range: 0 to 1)
  The opacity of the grid. Lower values allow more background to show through.

- **Latitude** (Default: 0, Range: -89 to 89)
  Tilts the grid up or down by this many degrees.

- **Swing** (Default: 0, Range: any)
  Rotation of the grid in degrees in its initial frame.

- **Roll** (Default: 0, Range: any)
  Tilts the grid from side to side, in degrees. If Latitude is 0, the effects of Swing and Roll are the same.

- **Z Dist** (Default: 1, Range: 0.01 or greater)
  Scales the 'distance' of the grid. Values greater than 1.0 move it farther away and make it smaller. Values less then 1.0 move it closer and enlarge it.

- **Tele Lens Width** (Default: 1, Range: 0.2 to 3)
  The amount of lens telescoping. Increase to zoom in with less perspective, decrease for a wider viewing angle with more perspective.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the background before combining with the grid. If 0, the result will contain only the grid image over black.

- **Combine** (Popup menu, Default: Over)
  Determines how the grid is combined with the Background.
  - **Over**: composites the grid over the background.
  - **Exclusion**: combines the grid and the Background with a difference operator.
  - **Grid Only**: displays the grid over black, ignoring the Background.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Mask Use** (Popup menu, Default: Luma)
  Determines how the Mask input channels are used to make a monochrome mask.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Mask** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

