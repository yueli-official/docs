---
title: Sharpen
---

## S_Sharpen

Amplifies the high frequencies in the source clip
such as edges and details. Increase the Sharpen Width parameter
to sharpen more of the mid range frequencies, and adjust Sharpen Amp
to control the amount of sharpening applied.

In the Sapphire Blur+Sharpen effects submenu.

![Sharpen](../_static/Sharpen.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Matte**: Defaults to None. Interpolate between the result and the Source input. White areas use the result of the effect. Black areas use the Source clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Quality** (Popup menu, Default: Best)
  Sharpen filter to apply.
  - **Best**: Advanced sharpen filter that has significantly fewer artifacts
  - **Fast**: Classic sharpen filter

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

- **Sharpen Amp** (Default: 1, Range: any)
  The amount of sharpening to apply.

- **Edge Threshold** (Default: 0.15, Range: 0 or greater)
  Edges stronger than this will not be sharpened. For standard dynamic range footage (not HDR) Edge Thresholds between 0.05 and 0.3 usually produce good results. Increasing the threshold will lead to a stronger effect (more edges will be sharpened), but can introduce dark bands around objects.

- **Small Detail Size** (Default: 0.01, Range: 0 or greater)
  The size in pixels of small details. This parameter can be adjusted using the Small Detail Size Widget.

- **Scale Tiny Details** (Default: 3, Range: 0 or greater)
  Values less than one will make tiny details less visible, values greater than one will make them more visible. Tiny details are about half the size of small details.

- **Scale Small Details** (Default: 1.5, Range: 0 or greater)
  Values less than one will make small details less visible, values greater than one will make them more visible.

- **Scale Medium Details** (Default: 1, Range: 0 or greater)
  Values less than one will make medium details less visible, values greater than one will make them more visible. Medium details are about twice the size of small details.

- **Scale Large Details** (Default: 2, Range: 0 or greater)
  Values less than one will make medium details less visible, values greater than one will make them more visible. Large details are about four times the size of small details.

- **Sharpen Width** (Default: 0.112, Range: 0 or greater)
  The width in pixels to perform the sharpen. Increase to sharpen softer edges, decrease to sharpen only the sharper edges.

- **Sharpen Luma** (Default: 1, Range: 0 or greater)
  The relative amount of sharpening to apply to the luminance of the source.

- **Sharpen Chroma** (Default: 1, Range: 0 or greater)
  The relative amount of sharpening to apply to the chroma of the source.

- **Sharpen Red** (Default: 1, Range: any)
  The relative amount of sharpening to apply to the red color channel.

- **Sharpen Green** (Default: 1, Range: any)
  The relative amount of sharpening to apply to the red color channel.

- **Sharpen Blue** (Default: 1, Range: any)
  The relative amount of sharpening to apply to the red color channel.

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

- **Show Small Detail Size** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Small Detail Size parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

