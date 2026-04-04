---
title: WarpCornerPin
---

## S_WarpCornerPin

Performs a 3D perspective warp of the source image
to align the corners with the four indicated points. This can be
useful for positioning the source over an object in another clip,
such as a billboard or computer screen.

In the Sapphire Distort effects submenu.

![WarpCornerPin](../_static/WarpCornerPin.jpg)


### Inputs:

- **Source**: The current layer. The input clip to be warped.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Corner1** (X & Y, Default: [-0.806 0.133], Range: any)
  Location of the lower-left corner of the source. This parameter can be adjusted using the Corner1 Widget.

- **Corner2** (X & Y, Default: [0.389 0.577], Range: any)
  Location of the lower-right corner of the source. This parameter can be adjusted using the Corner2 Widget.

- **Corner3** (X & Y, Default: [-0.556 -0.441], Range: any)
  Location of the upper-right corner of the source. This parameter can be adjusted using the Corner3 Widget.

- **Corner4** (X & Y, Default: [0.611 -0.485], Range: any)
  Location of the upper-left corner of the source. This parameter can be adjusted using the Corner4 Widget.

- **Filter** (Check-box, Default: on)
  If enabled, the image is adaptively filtered when it is resampled. This gives a better quality result when parts of the image are warped smaller.

- **Bulge** (X & Y, Default: [0 0], Range: -1 to 1)
  Distorts the perspective of the warped image, so that it appears to bulge in one direction. A value of 1 gives no distortion. A value of less than one causes the image to stretch toward the upper/right corner, while a value of greater than one causes it to stretch to the lower/left corner.

- **Wrap** (X & Y, Popup menu, Default: [ No No ])
  Determines the method for accessing outside the borders of the source image.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less
visible with this method.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.
If your image has sharp color changes where the matte
channel also has sharp edges, you may get better results with Normal
mode.

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  These 4 parameters, Crop Top , Crop Bottom , Crop Left, and Crop Right , allow selecting a rectangular subsection of the input image to be processed. If the Wrap parameters are set to "No" the exposed borders will be transparent. If the Wrap is "Tile" or "Reflect" the source image is wrapped on the new cropped borders to fill the frame. This can make it easier to avoid artifacts due to distorting an image with bad edges.

- **Show Corner1** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Corner1 parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Corner2** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Corner2 parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Corner3** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Corner3 parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Corner4** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Corner4 parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

