---
title: Deband
---

## S_Deband

Removes banding artifacts from a clip by diffusing pixels
across the banded areas, while keeping the original edges intact.
To use this effect, first select Show:Edges and adjust the edge
threshold until the banding edges just disappear, leaving only the
desired real edges. Then select Show:Result to see the result. If
you still see some banding, increase Diffuse Threshold and/or
Diffuse Radius.

In the Sapphire Blur+Sharpen effects submenu.

![Deband](../_static/Deband.jpg)


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

- **Edge Threshold** (Default: 2, Range: 0 to 255)
  The amount by which adjacent pixels must differ to constitute a real desired edge. A value of 1.0 represents the smallest possible difference at 8 bits. This parameter should be set high enough that none of the bands appear as edges, but low enough that all the real edges are still detected.

- **Grow Edges** (Default: 0, Range: 0 or greater)
  Amount to grow the detected edges in approximate pixels. Increasing this parameter can prevent diffusion in areas that are near edges, but not on an edge.

- **Show** (Popup menu, Default: Result)
  Selects the type of output.
  - **Result**: Shows the final result.
  - **Edges**: Shows the edges of the image, where adjacent pixels differ by more than Edge
Threshold. Use this mode to help fine-tune the edge detection parameters.

- **Diffuse Threshold** (Default: 1, Range: 0 or greater)
  The maximum color difference allowed when diffusing pixels. This parameter is automatically scaled by the edge threshold. Increasing it can give better results when there is a gradient within the bands. Decreasing it will reduce diffusion in areas where there are no edges.

- **Diffuse Radius** (Default: 12, Range: 0 or greater)
  The maximum radius of pixel diffusion, in approximate pixels. A larger value will remove banding more effectively in large areas with uniform colors, while a smaller value will give a better result in areas with many small color regions.

- **Pre Blur** (Default: 0, Range: 0 or greater)
  Blurs the source before diffusing pixels.

- **Post Blur** (Default: 0.5, Range: 0 or greater)
  Blurs the result after diffusing pixels. Use this parameter to reduce noisiness in the result.

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

