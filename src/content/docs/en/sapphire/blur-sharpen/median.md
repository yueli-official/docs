---
title: Median
---

## S_Median

Applies a median filter to the source image. Median filters are useful for cleaning up isolated spots and noise.

In the Sapphire Blur+Sharpen effects submenu.

![Median](../_static/Median.jpg)


### Inputs:

- **Source**: The current layer. The clip to be filtered.

- **Mask**: Defaults to None. Interpolate between the result and the Source input. White areas use the result of the effect. Black areas use the Source clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: Median)
  Selects whether to apply the same median filter to all channels, or a separate median filter per channel.
  - **Median**: Apply the same median filter to each channel.
  - **MedianChannels**: Apply different median channels to the r, g, and b channels.

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

- **Size** (Default: 1, Range: 0.1 to 40)
  Size of the median filter.

- **Subpixel** (Check-box, Default: off)
  Enables subpixel-width filtering. Use this for smoother animation of the Size parameter.

- **Size Rel X** (Default: 1, Range: 0.1 to 5)
  The relative horizontal size of the filter.

- **Size Rel Y** (Default: 1, Range: 0.1 to 5)
  The relative vertical size of the filter.

- **Red Rel Size X** (Default: 1, Range: 0 to 2)
  The relative horizontal size of the filter in the red channel.

- **Red Rel Size Y** (Default: 1, Range: 0 to 2)
  The relative vertical size of the filter in the red channel.

- **Green Rel Size X** (Default: 1, Range: 0 to 2)
  The relative horizontal size of the filter in the green channel.

- **Green Rel Size Y** (Default: 1, Range: 0 to 2)
  The relative vertical size of the filter in the green channel.

- **Blue Rel Size X** (Default: 1, Range: 0 to 2)
  The relative horizontal size of the filter in the blue channel.

- **Blue Rel Size Y** (Default: 1, Range: 0 to 2)
  The relative vertical size of the filter in the blue channel.

- **Alpha Rel Size X** (Default: 1, Range: 0 to 2)
  The relative horizontal size of the filter in the alpha channel.

- **Alpha Rel Size Y** (Default: 1, Range: 0 to 2)
  The relative vertical size of the filter in the alpha channel.

- **Mask Use** (Popup menu, Default: Luma)
  Determines how the Mask input channels are used to make a monochrome mask.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Mask** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

