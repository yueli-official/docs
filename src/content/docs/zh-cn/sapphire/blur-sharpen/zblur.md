---
title: ZBlur
---

## S_ZBlur

Blurs areas of the source clip by different amounts
using depth values from a ZBuffer input. Separates the input into a
number of layers in depth and blurs them by different amounts
depending on each layer's depth. Linear fog can also be mixed into
the result. To use this effect, first set ZBuffer:Black Is Near or White Is
Near according to your Z buffer, then adjust the focus depth and
depth of field parameters to get the look you want. To help set the
focus depth, you can use Show: In Focus Zone.

In the Sapphire Blur+Sharpen effects submenu.

![ZBlur](../_static/ZBlur.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **ZBuffer**: Defaults to None. The input clip containing depth values for each Source pixel. These values should be in the range of black to white, and it is best if not anti-aliased. Normally black corresponds to the farthest objects and white to the nearest, though this can be adjusted using Z Buffer parameter.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Focal Depth** (Default: 0, Range: any)
  The depth of the focus plane; 0 is near and 1 is far. Areas with this Z value will be in focus. Objects near this depth may be in focus depending on the Depth of Field parameter. You can use Show: In Focus Zone to show the Focal Depth when adjusting. If the effect of this parameter seems backwards, you can invert the depth values using the Z Buffer parameter.

- **Depth Of Field** (Default: 0.1, Range: 0 to 1)
  Specifies how wide a range of depths near the Focal Depth will be in focus. If the Focal Depth is 0.5 and Depth of Field is 0.2, all objects with Z values from 0.4 to 0.6 will be in focus. Set to zero to have only objects exactly at the Focal Depth in focus. You can use Show: In Focus Zone to show this when adjusting.

- **Blur Width** (Default: 0.168, Range: 0 or greater)
  Scales the overall amount of blur. This parameter can be adjusted using the Blur Width Widget.

- **Blur Rel** (X & Y, Default: [1 1], Range: 0 or greater)
  The relative horizontal and vertical blur widths. Set Blur Rel X to 0 for a vertical-only blur, or set Blur Rel Y to 0 for a horizontal-only blur.

- **Z Buffer Type** (Popup menu, Default: White is Near)
  How to interpret the values in the Z buffer.
  - **Black is Near**: Black pixels in the Z buffer indicate that
the object at that point is near (close to you), and white means far
away.
  - **White is Near**: White pixels in the Z buffer indicate that
the object at that point is near (close to you), and black means far
away.

- **Show** (Popup menu, Default: Result)
  Selects the type of output
  - **Result**: Shows the normal result of the effect.
  - **In Focus Zone**: Highlights the in-focus areas of the clip to make it easier to select the focal point and depth.

- **Layers** (Integer, Default: 5, Range: 2 to 50)
  The number of depth layers to separate the source into. More layers require more processing but give smoother results in Z. More layers are sometimes needed to avoid visible seams between the layers.

- **Layer Mode** (Popup menu, Default: Interp)
  Determines how the differently blurred layers are combined.
  - **Comp**: the closer layers are composited over the farther
layers. This method often gives better results if you have objects
at different depths overlapping each other with discontinuous values
in your depth image. However, this option can be slower, and sometimes
artifacts between layers are visible.
  - **Interp**: the layers are interpolated using depth
image values. This method gives smoother transitions between
layers, and is usually better if there are no sharp changes in your
depth image.

- **Width Rel Near** (Default: 1, Range: 0 or greater)
  Scales the blur width for parts of the image that are nearer than the focal plane.

- **Width Rel Far** (Default: 1, Range: 0 or greater)
  Scales the blur width for parts of the image that are farther away than the focal plane.

- **Fog Near** (Default: 0, Range: 0 to 1)
  The amount of fog to add to nearby (close) objects.

- **Fog Far** (Default: 0, Range: 0 to 1)
  The amount of fog to add to far away objects.

- **Fog Color** (Default rgb: [0.5 0.5 0.5])
  The fog color should normally match the sky or background color of the source clip. Use gray for mist, brown for smog, blue for underwater, etc.

- **Zbuffer Use** (Popup menu, Default: Luma)
  Determines how the ZBuffer input channels make a monochrome z image.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Soft Borders** (Check-box, Default: off)
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

- **Show Blur Width** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Blur Width parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

