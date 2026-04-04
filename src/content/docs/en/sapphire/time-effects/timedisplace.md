---
title: TimeDisplace
---

## S_TimeDisplace

Displaces the Source clip by variable amounts in
time depending on the brightness values of a Matte input.

In the Sapphire Time effects submenu.

![TimeDisplace](../_static/TimeDisplace.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Matte**: Defaults to None. Determines the amount of time displacement. Where the Matte is white the Source is time-shifted by a number of frames given by White Time Shift, and where it is black the Source is shifted by Black Time Shift. Gray areas are time-shifted by the appropriately interpolated amount. This input can be optionally blurred using the Blur Matte parameter. If this input is not provided, the Source input is used for the displacement matte instead.


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

- **Black Time Shift** (Default: -10, Range: any)
  Time shift by this many frames where the Matte is black.

- **White Time Shift** (Default: 10, Range: any)
  Time shift by this many frames where the Matte is white.

- **Shift Relative To** (Popup menu, Default: Current Frame)
  Selects relative or absolute time-shifting.
  - **Frame 0**: Time shift to an absolute frame number, relative to thefirst frame.
  - **Current Frame**: Time shift relative to the current frame.

- **Interp Frames** (Check-box, Default: off)
  Selects the method to use for non-integer frame number references. If disabled, the nearest integer frame number is used with no interpolation, which usually gives visible edges between the time slices. If enabled, a weighted interpolation is performed between the two nearest integer frame numbers, which smooths the results between the time slices.

- **Matte Use** (Popup menu, Default: Luma)
  Determines how the Matte input channels are used to make a monochrome matte.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Matte** (Default: 0.224, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can be used to soften the edges or quantization artifacts of the matte, and smooth out the time displacements.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

