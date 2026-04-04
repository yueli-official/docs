---
title: BlurMoCurves
---

## S_BlurMoCurves

Performs a motion blur and optionally transforms
the source clip using the animated curves of the Z Dist, Rotate and
Shift parameters. If these parameters are constant, no motion blur
will occur.

In the Sapphire Blur+Sharpen effects submenu.

![BlurMoCurves](../_static/BlurMoCurves.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Mask**: Defaults to None. If provided, the amount of motion blur is scaled by this input for each destination pixel. This input can be affected using the Blur Mask, Invert Mask, or Mask Use parameters.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: Transform and Blur)
  Allows disabling of the transformation.
  - **Transform and Blur**: transforms the Source as well as blurring.
  - **Blur Only**: this can be useful if the motions have already
occurred. The curves are used only to apply the corresponding motion
blur in place, and no transformation is performed.

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

- **Center** (X & Y, Default: [0 0], Range: any)
  The center of rotation and zooming, in screen coordinates relative to the center of the frame. The shift values should be zero for this location to make sense. This parameter can be adjusted using the Center Widget.

- **Z Dist** (Default: 1, Range: 0.001 or greater)
  The 'distance' of the image from the camera, about the Center position. The rate of change of this parameter is also used for the motion blur. Values greater than 1.0 move it farther away and make it smaller. Values less than 1.0 move the image closer and enlarge it. This parameter can be adjusted using the Transform Widget.

- **Rotate** (Default: 0, Range: any)
  Rotates the image by this amount in degrees, about the Center. The rate of change of this parameter is also used for the motion blur. Note that for high rotation speeds, the motion blur will become less accurate. This parameter can be adjusted using the Transform Widget.

- **Shift** (X & Y, Default: [0 0], Range: any)
  Translates the source image by this amount. The rate of change of this parameter is also used for the motion blur. It is in screen coordinates for easy use with tracker data. This parameter can be adjusted using the Transform Widget.

- **Shutter Duration** (Default: 1, Range: 0 or greater)
  The amount of time, in frames, to apply the motion blur over. Larger values cause more blurring, smaller values cause less. The curves are sampled at plus and minus half of this value.

- **Shutter Shift** (Default: 0, Range: any)
  The time-shift in frames of the motion blur. If the Shutter Speed is 1.0 and Shutter Shift is 0, the blur is calculated between the current frame -.5 and +.5. If the Shutter Shift is instead .5 then the motion blur would be calculated between the current frame and the next frame.

- **Exposure Bias** (Default: 0.5, Range: 0 to 1)
  Determines the variable amount of exposure along the path between the From and To transformations. A value of 0 causes more exposure at the From end, 0.5 causes equal exposure along the path, and 1.0 causes more exposure at the To end. If you have bright spots on a dark background, a 0 value would cause the processed spots to be brighter at the From end and dark at the To end, and a 1.0 value would cause the opposite.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Wrap** (X & Y, Popup menu, Default: [ No No ])
  Determines the method for accessing outside the borders of the source image.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less
visible with this method.

- **Blur Res** (Popup menu, Default: Full)
  Selects the resolution factor for the motion blur. This is similar to the general 'Res' factor parameter, but does a better job of averaging down to lower resolution and interpolating back up to the result. Higher resolutions give better quality, lower resolutions give faster processing.
  - **Full**: Full resolution is used.
  - **Half**: The motion blurring is performed at half resolution.
  - **Quarter**: The motion blurring is performed at quarter resolution.

- **Subpixel** (Check-box, Default: on)
  If enabled, uses a better quality but slightly slower method for performing the blur.

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

- **Show Center** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Transform** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Z Dist and Rotate parameters.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Shift** (Check-box, Default: off)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

