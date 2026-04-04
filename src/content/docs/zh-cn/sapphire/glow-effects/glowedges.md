---
title: GlowEdges
---

## S_GlowEdges

Creates glowing light from the edges of the source clip.
This differs from the default Glow in that small or thin objects
generate as much glow around their edges as large objects. Also the
glow colors are not affected by the colors of the source clip.

In the Sapphire Lighting effects submenu.

![GlowEdges](../_static/GlowEdges.jpg)


### Inputs:

- **Source**: The current layer. Edges are extracted from this input clip to determine the glow locations.

- **Background**: Defaults to None. The clip to combine the glows with. If no background is given, the Source is also used as the Background.

- **Matte**: Defaults to None. If provided, the source glow colors are scaled by this input. A monochrome matte can be used to choose a subset of Source areas that will generate glows. A color matte can be used to selectively adjust the glow colors in different regions. The matte is applied to the source before the glows are generated so it will not clip the resulting glows.


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

- **Glow Brightness** (Default: 2, Range: 0 or greater)
  Controls the overall glow brightness. Set to zero to get no glow.

- **Color** (Default rgb: [1 1 1])
  Scales the color of the glows.

- **Glow Width** (Default: 0.224, Range: 0 or greater)
  Scales the glow distance. This and all the width parameters can be adjusted using the Width Widget. Note that a zero glow width still enhances the bright areas; set the brightness parameter to zero if you want to pass the Source through unchanged.

- **Width X** (Default: 1, Range: 0 or greater)
  Scales the horizontal glow width. Set to 0 for vertical only.

- **Width Y** (Default: 1, Range: 0 or greater)
  Scales the vertical glow width. Set to 0 for horizontal only.

- **Width Red** (Default: 1, Range: 0 or greater)
  Scales the red glow width. If the red, green, and blue widths are equal, the glow colors be uniform with distance. If they are not equal, the glows will vary in color with distance.

- **Width Green** (Default: 1.2, Range: 0 or greater)
  Scales the green glow width.

- **Width Blue** (Default: 1.4, Range: 0 or greater)
  Scales the blue glow width.

- **Subpixel** (Check-box, Default: on)
  Enables glowing by subpixel widths. Use this for smoother animation of the Width parameters.

- **Edges Smooth** (Default: 0, Range: 0 or greater)
  Determines the width of the extracted edges which generate the glows.

- **Edges Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the edges before the glows are applied.

- **Edges Threshold** (Default: 0.5, Range: 0 or greater)
  Increase to remove glows on the less sharp edges.

- **Show** (Popup menu, Default: Result)
  Selects the type of output.
  - **Result**: Normally the glows are combined with the source or
background, and output.
  - **Edges**: The edge image only is output, before any glows are
applied. This can be helpful while adjusting the various edge
parameters.

- **Combine** (Popup menu, Default: Screen)
  Determines how the glow is combined with the Source or Background. This parameter has no effect if Light BG is set to 1.
  - **Mult**: the source or background is multiplied by the glow.
  - **Add**: the glow is added to the source or background.
  - **Screen**: the glow is blended with the source or background using a screen operation.
  - **Difference**: the result is the difference between the glow and the
source or background.
  - **Overlay**: the glow is combined with the source or background using an overlay function.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the glows. The maximum of the red, green, and blue glow brightness is scaled by this value and combined with the background Alpha at each pixel.

- **Glow From Alpha** (Default: 0, Range: 0 to 1)
  Set to 1 to generate glows from the alpha channel of the source input instead of the RGB channels. In this case the glows will not pick up color from the source and will typically be brighter. Values between 0 and 1 interpolate between using the RGB and the Alpha.

- **Glow Under Source** (Default: 0, Range: 0 to 1)
  Set to 1 to composite the Source input over the glows.

- **Light Background** (Default: 0, Range: 0 to 1)
  Increase this to give a look of the glow casting light onto the background image. To see this more clearly you can also lower the Background Scale parameter or raise the Brightness parameter.

- **Source Opacity** (Default: 1, Range: 0 to 1)
  Scales the opacity of the Source input when combined with the glows. This does not affect the generation of the glows themselves.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the background. This parameter only has an effect if the background input is provided, and is visible due to a partially transparent Source image or a reduced Source Opacity parameter value.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Expand Borders** (Check-box, Default: off)
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

- **Show Glow Width** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Glow Width parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

