---
title: EmbossShiny
---

## S_EmbossShiny

Embosses the Source clip using the Bumps input as a relief map. A lighting
model is used which includes highlights from specular reflections.
Increase the Bumps Smooth parameter for bolder bumps, and adjust the Light
Dir to illuminate the bumps from different angles.

In the Sapphire Stylize effects submenu.

![EmbossShiny](../_static/EmbossShiny.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Bumps**: Defaults to None. The bump map for the emboss. Only the luminance of this input is used.

- **Matte**: Defaults to None. If provided, the emboss is applied only at the areas specified by this input. This input can be affected using the Blur Matte, Invert Matte, or Matte Use parameters.


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

- **Light Dir** (X & Y, Default: [-0.5 0.361], Range: any)
  The direction vector for the light source. Surface shading is calculated using light from this direction shining onto the Bumps input. This parameter can be adjusted using the Light Dir Widget.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Light Color** (Default rgb: [1 1 1])
  The color of the light source that creates the embossed result.

- **Bumps Scale** (Default: 1, Range: any)
  Scales the amplitude of the bump map.

- **Bumps Threshold** (Default: 0, Range: 0 or greater)
  This value is subtracted from the Bumps input before it is used.

- **Bumps Smooth** (Default: 0.01, Range: 0 or greater)
  If positive, the Bumps input is blurred by this amount before being used. Increase for a softer emboss effect.

- **Subpixel Smooth** (Check-box, Default: on)
  If enabled, the amount of pre-smoothing of the Bumps input is performed at subpixel accuracy. It can be helpful if Bumps Smooth is very small or is being animated. This parameter has no effect unless Bumps Smooth is positive.

- **Hilight Brightness** (Default: 0.8, Range: 0 to 1)
  Scales the brightness of the specular highlights.

- **Hilight Size** (Default: 0.5, Range: 0.1 or greater)
  Adjusts the size of the specular highlights.

- **Blur Matte** (Default: 0, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

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

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  These 4 parameters, Crop Top , Crop Bottom , Crop Left, and Crop Right , allow selecting a rectangular subsection of the input image to be processed. If the Wrap parameters are set to "No" the exposed borders will be transparent. If the Wrap is "Tile" or "Reflect" the source image is wrapped on the new cropped borders to fill the frame. This can make it easier to avoid artifacts due to distorting an image with bad edges.

- **Show Light Dir** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Light Dir parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

