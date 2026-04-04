---
title: VintageColor3Strip
---

## S_VintageColor3Strip

Simulates the color 3-strip film process
from 1935 through 1955. Three-strip color was a subtractive
process which exposed three separate film strips through color
filters, then applied complementary color dyes to the print
according to the density of the original records. This process
was used for many films such as The Wizard Of Oz, Fantasia, and Gone With
The Wind.
Modern color film has much broader color filtering in the emulsion
layers, so this effect simulates the narrower filters and
sharper colored dyes of the era which gave it its characteristic
vibrancy. This effect also allows adding grain and color
correction.

In the Sapphire Stylize effects submenu.

![VintageColor3Strip](../_static/VintageColor3Strip.jpg)


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

- **Amount** (Default: 1, Range: 0 or greater)
  Amount of the effect to use. Set to zero to get the original source. Increase beyond to to oversaturate.

- **Key Layer Density** (Default: 0.1, Range: 0 or greater)
  From 1932 up to about 1945, the blank print started with a 50 percent black and white duplicate of the green original record. This increased apparent sharpness and improved contrast. Set this to 0.5 for a historically accurate key layer, but it will decrease the overall brightness. After 1945 the key layer was no longer needed due to improvements in the process.

- **Grain Amp** (Default: 0, Range: 0 or greater)
  Scales the amplitude of the film grain that is added to the result. Set this to 0 to disable all grain.

- **Grain Blur** (Default: 0, Range: 0 or greater)
  The grain is smoothed by this amount. Increase for coarser grain.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Tint** (Default rgb: [1 1 1])
  Tints the image towards the given color.

- **Saturation** (Default: 1, Range: -2 to 10)
  Scales the color saturation. Increase for more intense colors. Set to 0 for monochrome.

- **Offset Darks** (Default: 0, Range: -8 to 2)
  Adds this gray value to the darker regions of the result. This can be negative to increase contrast.

- **Show** (Popup menu, Default: Result)
  Shows either the final result, or any of various intermediate parts of the process.
  - **Result**: Shows the final result.
  - **Pure Colors**: Shows an RGB matte containing only the pure colors in the source.
  - **Complementary Masks**: Shows a matte of the complementary colors used to apply the dyes to the final print.

- **Mask Use** (Popup menu, Default: Luma)
  Determines how the Mask input channels are used to make a monochrome mask.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Mask** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

