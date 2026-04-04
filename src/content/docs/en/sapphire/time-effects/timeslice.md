---
title: TimeSlice
---

## S_TimeSlice

Divides the output frame into slices, where each slice receives
a different frame from the source clip. An example use of this effect
might be to make a turning object twist into a helix shape instead of
rigidly rotating. The slices are oriented depending on Slice Direction,
and receive relative frame numbers between plus and minus half of Slice
Number. For example if the current frame number is 30, Slice Direction is
-90 degrees, Slice Number is 12, and Frame Offset is 0, the result will
consist of horizontal slices containing approximately frames 30-6 to 30+6
from bottom to top.

In the Sapphire Time effects submenu.

![TimeSlice](../_static/TimeSlice.jpg)


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

- **Slice Direction** (Default: 90, Range: any)
  The orientation of the slices, in degrees. If this is 0 the slices will go from left to right. If it is 90 they will go from top to bottom. This parameter can be adjusted using the Slice Widget.

- **Slice Number** (Default: 12, Range: 1 or greater)
  The number of time slices to slice the frame into. This parameter can be adjusted using the Slice Widget.

- **Frame Offset** (Default: 0, Range: any)
  Shifts all frame numbers in time that the slices receive. This parameter can be adjusted using the Slice Widget.

- **Interp Frames** (Check-box, Default: off)
  Selects the method to use for non-integer frame number references. If disabled, the nearest integer frame number is used with no interpolation, which usually gives visible edges between the time slices. If enabled, a weighted interpolation is performed between the two nearest integer frame numbers, which smooths the results between the time slices.

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

- **Show Slice** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Slice Direction, Slice Number, and Frame Offset parameters. This widget visually shows the single slice where the result equals the current frame of the source.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

