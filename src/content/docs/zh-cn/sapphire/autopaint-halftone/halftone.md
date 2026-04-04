---
title: HalfTone
---

## S_HalfTone

Generates a halftone version of the source clip using a black and white
pattern of dots. Use the Smooth Source parameter to remove some details and make
the dots more consistently round.

In the Sapphire Stylize effects submenu.

![HalfTone](../_static/HalfTone.jpg)


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

- **Dots** (Popup menu, Default: Black)
  Selects the dots' color model.
  - **Black**: dark dots are used on a bright background.
  - **White**: bright dots are used on a dark background.

- **Dots Frequency** (Default: 50, Range: 0 or greater)
  The frequency of the dots pattern. Increase for finer dots, decrease for larger dots.

- **Dots Angle** (Default: 30, Range: any)
  The angle of the overall dots pattern, in degrees.

- **Dots Rel Width** (Default: 1, Range: 0.01 or greater)
  The relative width of the dots. Increase for wider dots, decrease for taller ones.

- **Dots Sharpness** (Default: 4, Range: 0 or greater)
  Scales the sharpness of the edges of the dots.

- **Dots Lighten** (Default: 0, Range: any)
  Increase to lighten the resulting dot pattern.

- **Smooth Source** (Default: 0, Range: 0 or greater)
  If positive, the source is blurred by this amount before the halftone is applied. This can be used to remove some detail in the dots and make them more consistently round.

- **Color1** (Default rgb: [1 1 1])
  The 'bright' color to use for the dots pattern.

- **Color0** (Default rgb: [0 0 0])
  The 'dark' color to use for the dots pattern.

- **Dots Shift** (X & Y, Default: [0 0], Range: any)
  The horizontal and vertical translation of the dots pattern

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

