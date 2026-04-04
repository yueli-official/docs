---
title: Etching
---

## S_Etching

Generates a version of the source clip using two sets of black and white
lines of varying thickness to give an 'etching' or 'lithograph' look. Use
the Smooth Source parameter to remove some details and make the lines more
evenly shaped. Use the Lines Frequency parameter to adjust the density of
all lines.

In the Sapphire Stylize effects submenu.

![Etching](../_static/Etching.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

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

- **Lines Frequency** (Default: 50, Range: 0 or greater)
  The frequency of the etched lines. Increase for a finer line pattern, decrease for fewer lines.

- **Lines1 Frequency** (Default: 1, Range: 0 or greater)
  Scales the frequency of the first set of etched lines. Increase for a finer line pattern, decrease for fewer lines.

- **Lines2 Frequency** (Default: 1, Range: 0 or greater)
  Scales the frequency of the second set of etched lines.

- **Lines Angle** (Default: 0, Range: any)
  Rotation of the etched lines pattern in degrees.

- **Lines1 Angle** (Default: 30, Range: any)
  The relative angle of the first set of etched lines in degrees.

- **Lines2 Angle** (Default: -20, Range: any)
  The relative angle of the second set of etched lines in degrees.

- **Lines Shift** (X & Y, Default: [0 0], Range: any)
  Shifts the pattern of lines. This location will also be the center of rotation when the line angle parameters are adjusted.

- **Lines Sharpness** (Default: 4, Range: 0 or greater)
  The sharpness of the etched lines. Decrease for softer edges.

- **Lines Add Width** (Default: 0, Range: any)
  Increase for thicker lines.

- **Smooth Source** (Default: 0, Range: 0 or greater)
  If positive, the source is blurred by this amount before the etching is applied.

- **Color1** (Default rgb: [1 1 1])
  The 'brighter' color of the lines pattern.

- **Color0** (Default rgb: [0 0 0])
  The 'darker' color of the lines pattern.

- **Wave Amp** (Default: 0.1, Range: 0 or greater)
  The amplitude of the waviness of the sets of etched lines.

- **Wave Frequency** (Default: 2, Range: 0 or greater)
  The frequency of the waviness of the etched lines. Increase for more waves.

- **Warp Amp** (Default: 0.04, Range: any)
  The amount the output is warped using the source brightness.

- **Warp Smooth** (Default: 0.044, Range: 0 or greater)
  The smoothness of the warping. This has no effect if Warp Amp is 0.

- **Edges Scale** (Default: 0.5, Range: 0 or greater)
  Adjusts the amount of source edges to be included in the result. If positive, edges in the source image are found and added to the etching pattern.

- **Edges Threshold** (Default: 0.3, Range: 0 or greater)
  Determines which edges are included in the result. Increase to remove minor edges and speckles. This has no effect unless Edges Scale is positive.

- **Edges Width** (Default: 0, Range: 0 or greater)
  The width of the edges added to the result. Increase for wider edges. This has no effect unless Edges Scale is positive.

- **Edges Sharpness** (Default: 3, Range: 0 or greater)
  Increase for sharper edges, decrease for softer edges. This has no effect unless Edges Scale is positive.

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

