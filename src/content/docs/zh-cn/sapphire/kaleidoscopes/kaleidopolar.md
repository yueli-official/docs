---
title: KaleidoPolar
---

## S_KaleidoPolar

Warps the source clip around in a disk shape and
reflects radially as if viewed through a reflecting cylinder.

In the Sapphire Stylize effects submenu.

![KaleidoPolar](../_static/KaleidoPolar.jpg)


### Inputs:

- **Source**: The current layer. The input clip to be warped.

- **Mask**: Defaults to None. If provided, the amplitude of warping is scaled by the values of this input clip. Gray values internally scale the warping amplitude rather than simply cross-fading between the effect and the original source to allow more continuous results at the mask edges and more detailed control over the warping amounts. This input can be affected using the Blur Mask, Invert Mask, or Mask Use parameters.


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

- **Apply Mask** (Popup menu, Default: Post-effect)
  Control where in the effect the mask is applied - this affects both the input mask and the mocha mask.
  - **Post-effect**: Applies the masks after all the effect has been rendered.
  - **Pre-effect**: Applies the mask to the source before processing the effect.

- **Center** (X & Y, Default: [0 0], Range: any)
  Center location of the kaleidoscoped image in screen coordinates relative to the center of the frame. The entire result will be shifted by this amount.

- **Z Dist** (Default: 1, Range: 0.001 or greater)
  Scales the 'distance' of the entire result in or out from the Center. Increase to zoom out, decrease to zoom in.

- **Rotate** (Default: 0, Range: any)
  Rotates the kaleidoscope's reflection pattern about the Center by this many degrees.

- **Stretch** (X & Y, Default: [1 1], Range: 0.1 or greater)
  Scales the horizontal or vertical size of the result.

- **Inside Shift Y** (Default: 0, Range: any)
  Shifts the source image up by this amount before it is reflected. This causes the resulting pattern of images to radiate outward from the center.

- **Angle Repeats** (Default: 6, Range: 0.01 or greater)
  The number of copies of the source image to wrap around. This should be an even integer to avoid a seam where the first copy connects to the last.

- **Kaleido Amount** (Default: 1, Range: 0 or greater)
  Adjusts the overall amount of distortion applied to the Source clip. Set this to zero to leave the source unchanged or to one for a normal kaleidoscope pattern.

- **Filter** (Check-box, Default: on)
  If enabled, the image is adaptively filtered when it is resampled. This gives a better quality result when parts of the image are warped smaller.

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

