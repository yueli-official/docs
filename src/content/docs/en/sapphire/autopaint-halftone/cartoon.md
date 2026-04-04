---
title: Cartoon
---

## S_Cartoon

Generates a version of the source clip with a
cartoon look. Finds the edges in the image and draws new outlines
for those edges. Smooths the colors of the areas between the edges,
and optionally posterizes the colors into fewer color values.

In the Sapphire Stylize effects submenu.

![Cartoon](../_static/Cartoon.jpg)


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

- **Edge Width** (Default: 0.02, Range: 0 or greater)
  The width of the outlined edges. Increase for thicker outlines.

- **Edge Strength** (Default: 2, Range: 0 or greater)
  Scales the strength of the outlined edges by this amount. Increase for heavier edges.

- **Edge Threshold** (Default: 0.1, Range: 0 or greater)
  Subtracts this value from outline image. Increase to remove unwanted noise and minor edges.

- **Edge Color** (Default rgb: [0 0 0])
  Outline the edges of the clip in this color.

- **Suppress Small Edges** (Default: 0.5, Range: 0 or greater)
  Increase this value to remove smaller edges while keeping the larger edges.

- **Edge Sharpen** (Default: 0, Range: 0 or greater)
  Amount to sharpen the outlines. Increase this value for sharper sides to the edges.

- **Smooth** (Default: 0.1, Range: 0 or greater)
  The amount to blur the colors in the non-edge regions.


### Posterize Parameters:

Posterize Amount:
*Default:
*0,
*Range:
*0 to 1.If positive, generates a posterized look by
limiting the number of colors in the result. Increase this for
fewer and larger regions of solid colors. Decrease for more colors
and more steps between colors.

Posterize Smooth:
*Default:
*0.1,
*Range:
*0 to 1.Amount to smooth the edges between color regions
when posterizing. Increase this value to reduce aliasing between
the colored areas. If set to 1, the areas will be completely
smoothed together and no posterize effect will occur.

Posterize Phase:
*Default:
*0,
*Range:
*any.
Amount to shift color boundaries when posterizing. Adjust
this to fine-tune the location of the edges between the color
regions. A phase of 1 is equivalent to 0.

### Color Correct Parameters:

Saturation:
*Default:
*1,
*Range:
*any.Scales the color saturation. Increase for more intense
colors. Set to 0 for monochrome.

Scale Lights:
*Default:
*1,
*Range:
*0 or greater.Scales the result by this value. Increase for a
brighter result.

Tint Lights:
*Default rgb:
*[1 1 1].Scales the result by this color, thus tinting the
lighter regions.

Tint Darks:
*Default rgb:
*[0 0 0].Adds this color to the darker regions of the source.

Offset Darks:
*Default:
*0,
*Range:
*any.Adds this gray value to the darker regions of the source.
This can be negative to increase contrast.

Mix With Source:
*Default:
*0,
*Range:
*0 to 1.Interpolates between the result (0) and the
original source (1).

Opacity:
*Popup menu, Default: Normal
*.Determines the method used for dealing with
opacity/transparency.
*All Opaque:
*Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).*Normal:
*Process opacity normally.*As Premult:
*Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

Mask Use:
*Popup menu, Default: Luma
*.Determines how the Mask input channels are used to make a
monochrome mask.
*Luma:
*the luminance of the RGB channels is used.*Alpha:
*only the Alpha channel is used.

Blur Mask:
*Default:
*0.05,
*Range:
*0 or greater.Blurs the Matte input by this amount before using. This
can provide a smoother transition between the matted and unmatted
areas. It has no effect unless the Matte input is provided.

Invert Mask:
*Check-box, Default:
*off.
If on, inverts the Matte input so the effect is applied
to areas where the Matte is black instead of white. This has no effect
unless the Matte input is provided.
