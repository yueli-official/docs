---
title: BlurChroma
---

## S_BlurChroma

Separates the source into luminance and chrominance
components, blurs the chrominance and/or the luminance independently, and
recombines them. You can also scale the luma and chroma independently
to enhance or remove either.

In the Sapphire Blur+Sharpen effects submenu.

![BlurChroma](../_static/BlurChroma.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Matte**: Defaults to None. If provided, the blur is only performed on regions of the source clip specified by the bright areas of this input. Pixels outside this matte are not blurred, and do not contribute to the resulting blurred pixels within it. This input can be affected using the Invert Matte, or Matte Use parameters.


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

- **Blur Chroma** (Default: 0.4, Range: 0 or greater)
  The amount to blur the chrominance. This parameter can be adjusted using the Blur Chroma Widget.

- **Blur Luminance** (Default: 0, Range: 0 or greater)
  The amount to blur the luminance. This parameter can be adjusted using the Blur Luminance Widget.

- **Blur Rel** (X & Y, Default: [1 1], Range: 0 or greater)
  The relative horizontal and vertical blur widths. Set Blur Rel X to 0 for a vertical-only blur, or set Blur Rel Y to 0 for a horizontal-only blur. This parameter can be adjusted using the Blur Chroma Widget.

- **Scale Chroma** (Default: 1, Range: 0 or greater)
  Scales the chrominance by this amount. Increase for more intense colors, decrease for muted colors.

- **Scale Luminance** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Offset Result** (Default: 0, Range: any)
  Adds this gray value to the result (or subtracts if negative). 0 has no effect, .5 is middle gray, and 1 is white.

- **Mix With Source** (Default: 0, Range: 0 to 1)
  Interpolates between the blurred result (0) and the original source (1). 0.1 can give a nice misty effect since it mixes only a little of the source in.

- **Filter** (Popup menu, Default: Gauss)
  The type of convolution filter to blur with.
  - **Box**: uses a rectangular shaped filter.
  - **Triangle**: smoother, uses a pyramid shaped filter.
  - **Gauss**: smoothest, uses a gaussian shaped filter.

- **Subpixel** (Check-box, Default: on)
  Enables blurring by subpixel amounts. Use this for smoother animation of the Blur Chroma or Blur Luminance parameters.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Matte Use** (Popup menu, Default: Luma)
  Determines how the Matte input channels are used to make a monochrome matte.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Blur Chroma** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Blur Chroma parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Blur Luminance** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Blur Luminance parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

