---
title: ZGlow
---

## S_ZGlow

Glows areas of the source clip with varying widths
depending on the depth values from a ZBuffer input. Separates the
input into a number of layers and applies different amounts of glow
depending on Width Near, Width Far, Brightness Near, and Brightness
Far parameters.

In the Sapphire Lighting effects submenu.

![ZGlow](../_static/ZGlow.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **ZBuffer**: Defaults to None. The input clip containing depth values for each Source pixel. These values should be in the range of black to white, and it is best if not anti-aliased. Normally black corresponds to the farthest objects and white to the nearest, though this can be adjusted using Z Buffer parameter.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Brightness** (Default: 2, Range: 0 or greater)
  Scales the brightness of all the glows.

- **Color** (Default rgb: [1 1 1])
  Scales the color of the glow. The colors and brightnesses of the glow are also affected by the Source input.

- **Width Near** (Default: 0.0336, Range: 0 or greater)
  The glow width of near (close) objects.

- **Width Far** (Default: 0.4, Range: 0 or greater)
  The glow width of far away objects.

- **Threshold** (Default: 0.5, Range: 0 or greater)
  Glows are generated from locations in the source clip that are brighter than this value. A value of 0.9 causes glows at only the brightest spots. A value of 0 causes glows for every non-black area.

- **Threshold Add Color** (Default rgb: [0 0 0])
  This can be used to raise the threshold on a specific color and thereby reduce the glow generated on areas of the source clip containing that color.

- **Z Buffer Type** (Popup menu, Default: White is Near)
  How to interpret the values in the Z buffer.
  - **Black is Near**: Black pixels in the Z buffer indicate that
the object at that point is near (close to you), and white means far
away.
  - **White is Near**: White pixels in the Z buffer indicate that
the object at that point is near (close to you), and black means far
away.

- **Z Min** (Default: 0, Range: 0 to 1)
  Clamps all Z values to this minimum bound. Use this parameter to create a constant glow on all parts of the image nearer than Z Min.

- **Z Max** (Default: 1, Range: 0 to 1)
  Clamps all Z values to this maximum bound. Use this parameter to create a constant glow on all parts of the image farther than Z Max.

- **Layers** (Integer, Default: 5, Range: 2 to 50)
  The number of depth layers to separate the source into. More layers require more processing but give smoother results in Z. More layers are sometimes needed to avoid visible seams between the layers.

- **Brightness Near** (Default: 1, Range: 0 or greater)
  Scales the glow brightness for near objects.

- **Color Near** (Default rgb: [1 1 1])
  Scales the glow color for near objects.

- **Width Red Near** (Default: 1, Range: 0 or greater)
  Scales the red glow width for near objects.

- **Width Green Near** (Default: 1, Range: 0 or greater)
  Scales the green glow width for near objects.

- **Width Blue Near** (Default: 1, Range: 0 or greater)
  Scales the blue glow width for near objects.

- **Brightness Far** (Default: 1, Range: 0 or greater)
  Scales the glow brightness for far objects.

- **Color Far** (Default rgb: [1 1 1])
  Scales the glow color for far objects.

- **Width Red Far** (Default: 1, Range: 0 or greater)
  Scales the red glow width for far objects.

- **Width Green Far** (Default: 1, Range: 0 or greater)
  Scales the green glow width for far objects.

- **Width Blue Far** (Default: 1, Range: 0 or greater)
  Scales the blue glow width for far objects.

- **Width X** (Default: 1, Range: 0 or greater)
  Scales the horizontal glow width. Set to 0 for vertical only.

- **Width Y** (Default: 1, Range: 0 or greater)
  Scales the vertical glow width. Set to 0 for horizontal only.

- **Width Red** (Default: 1, Range: 0 or greater)
  Scales the red glow width. If the red, green, and blue widths are equal, the glows will match the color of the source clip. If they are not equal, the glows will vary in color with distance.

- **Width Green** (Default: 1.2, Range: 0 or greater)
  Scales the green glow width.

- **Width Blue** (Default: 1.4, Range: 0 or greater)
  Scales the blue glow width.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the zs. The maximum of the red, green, and blue z brightness is scaled by this value and combined with the background Alpha at each pixel.

- **Source Opacity** (Default: 1, Range: 0 to 1)
  Scales the opacity of the Source input when combined with the zs. This does not affect the generation of the zs themselves.

- **Zbuffer Use** (Popup menu, Default: Luma)
  Determines how the ZBuffer input channels make a monochrome z image.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

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

- **Show Width Near** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Width Near parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Width Far** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Width Far parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

