---
title: FlysEyeCircles
---

## S_FlysEyeCircles

Breaks the image into circle shaped tiles and transforms the image within
each shape, to create a fly's eye view effect. The Overlap options allow
the circles to be combined in different ways where they overlap. The
'Inside' parameters transform the Source image before it is tiled into the
pattern, and the 'Tile' parameters transform the entire fly's eye pattern.

In the Sapphire Stylize effects submenu.

![FlysEyeCircles](../_static/FlysEyeCircles.jpg)


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

- **Tile Frequency** (Default: 12, Range: 0.1 or greater)
  The frequency of the tile pattern, increase for more smaller tiles. This parameter can be adjusted using the Tile Freq Widget.

- **Tile Rel Height** (Default: 1, Range: 0.01 or greater)
  The relative height of the tile shapes, increase for taller tiles.

- **Tile Shift** (X & Y, Default: [0 0], Range: any)
  Translates the tile pattern. This parameter can be adjusted using the Tile Shift Widget.

- **Circle Overlap** (Popup menu, Default: Ave)
  Determines the method used to combine the overlapping regions of the circles.
  - **Ave**: uses a weighted average across the overlapping region
for a smooth transition.
  - **Screen**: uses a screen operation.
  - **Max**: uses the lighter.
  - **Min**: uses the darker.
  - **Mult**: uses a multiply operation.

- **Circle Radius** (Default: 1, Range: 0 to 1)
  The radius of the circles relative to each other. If this is less than 1.0 you will get empty spaces between the circles. The color of these empty spaces will be either transparent, black, or white depending on the combine mode.

- **Edge Softness** (Default: 0, Range: 0 to 1)
  The softness of the edges of the circles. If this is increased, it may also be necessary to lower the Circle Radius to avoid rectangular artifacts where the soft edges overlap.

- **Inside Zdist** (Default: 2, Range: 0 or greater)
  Determines the zoom factor of the image inside each tile. Values greater than 1 zoom out, values less than 1 zoom in. If this is 1, Inside Rotate is 0, and Overall Zdist is 1, the result should be the same as the input image.

- **Inside Rotate** (Default: 0, Range: any)
  The rotation angle of the image inside each tile, in degrees.

- **Overall Zdist** (Default: 1, Range: any)
  Creates an overall zooming effect by making each tile look toward or away from the image center. Decrease to zoom in, increase to zoom out. When 0 all tiles should contain identical images.

- **Wrap** (Popup menu, Default: Reflect)
  Determines the method for accessing outside the borders of the source image.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less
visible with this method.

- **Filter** (Check-box, Default: on)
  If enabled, the Source image is resampled using pixel averaging. This removes aliasing and gives a higher quality result especially when Inside Zdist is large. It may not be necessary if your input image is smooth or Inside Zdist is small.

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

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  These 4 parameters, Crop Top , Crop Bottom , Crop Left, and Crop Right , allow selecting a rectangular subsection of the input image to be processed. If the Wrap parameters are set to "No" the exposed borders will be transparent. If the Wrap is "Tile" or "Reflect" the source image is wrapped on the new cropped borders to fill the frame. This can make it easier to avoid artifacts due to distorting an image with bad edges.

- **Show Tile Freq** (Check-box, Default: off)
  Turns on or off the screen user interface for adjusting the Tile Frequency parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Tile Shift** (Check-box, Default: off)
  Turns on or off the screen user interface for adjusting the Tile Shift parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

