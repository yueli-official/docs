---
title: WarpPerspective
---

## S_WarpPerspective

Transforms the source clip onto a 3D plane with perspective. Adjust the
Latitude, Swing, and Roll parameters to rotate the image on various axes,
and adjust Shift and Z Dist to translate and zoom. Turn off the Wrap
options to give a single non-repeated copy of the image.

In the Sapphire Distort effects submenu.

![WarpPerspective](../_static/WarpPerspective.jpg)


### Inputs:

- **Source**: The current layer. The input clip to be warped.

- **Matte**: Defaults to None. If provided, the amplitude of warping is scaled by the values of this input clip. Gray values internally scale the warping amplitude rather than simply cross-fading between the effect and the original source to allow more continuous results at the matte edges and more detailed control over the warping amounts. This input can be affected using the Blur Matte, Invert Matte, or Matte Use parameters.


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

- **Latitude** (Default: 35, Range: -89 to 89)
  Tilts the image up or down in 3D. Positive latitude tilts the image down and negative tilts it up.

- **Swing** (Default: 0, Range: any)
  Rotation of the image in degrees in its initial frame.

- **Roll** (Default: 0, Range: any)
  Tilts the result from side to side, in degrees.

- **Z Dist** (Default: 3, Range: 0.001 or greater)
  Scales the 'distance' of the image. Values greater than 1.0 move it farther away and make it smaller. Values less then 1.0 move the image closer and enlarge it.

- **Tele Lens Width** (Default: 1, Range: 0.2 to 3)
  The amount of lens telescoping. Increase to zoom in with less perspective, decrease for a wider viewing angle with more perspective.

- **Shift Orig** (X & Y, Default: [0 0], Range: any)
  Translates the image in its initial frame.

- **Wrap** (X & Y, Popup menu, Default: [ Tile Tile ])
  Determines the method for accessing outside the borders of the source image.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less
visible with this method.

- **Wrap Above Horizon** (Check-box, Default: off)
  When the image is sufficiently tilted using the Latitude parameter, a horizon can be seen. When Wrap Above Horizon is check, the image is repeated on both sides of the horizon.

- **Filter** (Check-box, Default: on)
  If enabled, the image is adaptively filtered when it is resampled. This gives a better quality result when parts of the image are warped smaller.

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
If your image has sharp color changes where the matte
channel also has sharp edges, you may get better results with Normal
mode.

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  These 4 parameters, Crop Top , Crop Bottom , Crop Left, and Crop Right , allow selecting a rectangular subsection of the input image to be processed. If the Wrap parameters are set to "No" the exposed borders will be transparent. If the Wrap is "Tile" or "Reflect" the source image is wrapped on the new cropped borders to fill the frame. This can make it easier to avoid artifacts due to distorting an image with bad edges.

