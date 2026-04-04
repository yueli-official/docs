---
title: VintageColor2Strip
---

## S_VintageColor2Strip

Simulates the old color 2-strip film
process from the 1920s. The scene is exposed twice, through red and
green filters, onto alternating frames of a monochrome film strip.
Then the red print is dyed with a red dye, and the green print is
dyed cyan. Those two strips are cemented together back-to-back to
form the final print. The result contains mostly red and green
colors, with some synthetic blue from the blue components of the
dyes.
This effect simulates the two filter colors and the two dye colors,
and also allows adding grain and color correction.

In the Sapphire Stylize effects submenu.

![VintageColor2Strip](../_static/VintageColor2Strip.jpg)


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

- **Red Filter** (Default rgb: [1 0 0])
  The color of the red filter.

- **Bluegreen Filter** (Default rgb: [0 1 0.5])
  The color of the green filter.

- **Red Dye** (Default rgb: [1 0 0])
  The dye color for the red strip.

- **Cyan Dye** (Default rgb: [0.02 1 0.91])
  The dye color for the cyan strip. Adjust slightly greener for a warmer look.

- **Grain Amp** (Default: 0, Range: 0 or greater)
  Scales the amplitude of the film grain that is added to the result. Set this to 0 to disable all grain.

- **Grain Blur** (Default: 0, Range: 0 or greater)
  The grain is smoothed by this amount. Increase for coarser grain.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Saturation** (Default: 1, Range: -2 to 10)
  Scales the color saturation. Increase for more intense colors. Set to 0 for monochrome.

- **Offset Darks** (Default: 0, Range: -8 to 2)
  Adds this gray value to the darker regions of the result. This can be negative to increase contrast.

- **Show** (Popup menu, Default: Result)
  Shows either the final result, or any of various intermediate parts of the process.
  - **Result**: Shows the final result.
  - **Red Strip**: Shows the red-filtered source as monochrome, as it would be on the real film.
  - **BlueGreen Strip**: Shows the blue-green-filtered source as monochrome, as it would be on the real film.
  - **Red Dye**: Shows the red-dyed red strip.
  - **Cyan Dye**: Shows the cyan-dyed green strip.

- **Mask Use** (Popup menu, Default: Luma)
  Determines how the Mask input channels are used to make a monochrome mask.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Mask** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

