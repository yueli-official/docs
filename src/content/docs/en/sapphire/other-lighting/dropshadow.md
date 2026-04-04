---
title: DropShadow
---

## S_DropShadow

Generates a shadow on the Background clip using the
alpha channel of the Foreground or an optional Matte, then composites the Foreground over
the Background to give the final result.

In the Sapphire Lighting effects submenu.

![DropShadow](../_static/DropShadow.jpg)


### Inputs:

- **Foreground**: The current layer. The clip to use as foreground, and the alpha channel of this clip is used as the matte to generate the shadow.

- **Background**: Defaults to None. The shadow is drawn onto this Background clip.

- **Matte**: Defaults to None. If this is provided, its alpha channel is used instead of the Foreground to generate the shadow. This input can be affected by the Invert Matte or Matte Use parameters.


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

- **Shadow Color** (Default rgb: [0 0 0])
  The color of the shadow.

- **Shadow Opacity** (Default: 1, Range: 0 or greater)
  The opacity of the shadow, use values near 0 for subtle transparent shadows, or values near 1.0 for stronger shadows.

- **Shadow Blur** (Default: 0.1, Range: 0 or greater)
  Determines the softness of the shadow. This parameter can be adjusted using the Shift Widget.

- **Shift** (X & Y, Default: [0.042 -0.042], Range: any)
  The horizontal and vertical offset of the shadow. This parameter can be adjusted using the Shift Widget.

- **Fg Opacity** (Default: 1, Range: 0 to 1)
  Scales the opacity of the Foreground without affecting the shadow. Lowering this can be used to fade out the Foreground, or setting it to zero prevents the Foreground from being composited over the result at all.

- **Comp Premult** (Check-box, Default: on)
  Disable this if you have provided a separate Matte input and the Foreground pixel values have not been pre-multiplied by this Matte.

- **Matte Use** (Popup menu, Default: Alpha)
  Determines which Foreground or Matte input channels are used to make the shadow.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Invert Matte** (Check-box, Default: off)
  If enabled, the black and white of the Foreground alpha channel are inverted before use.

- **Expand Borders** (Check-box, Default: on)
  If enabled, transparent borders are added to the input image before processing. This allows the result to include soft edges beyond the original image size. When off, the effect only occurs within the frame and the result will retain an edge at the borders.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Shift** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Shadow Blur parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

