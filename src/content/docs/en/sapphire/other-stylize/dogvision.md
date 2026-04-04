---
title: DogVision
---

## S_DogVision

Generates a dual color-channel version of the input
image, as might be perceived by the limited color vision system of
dogs. Humans have three color receptors (for red, green, and blue)
while dogs have only two receptors (for yellow and blue).

In the Sapphire Stylize effects submenu.

![DogVision](../_static/DogVision.jpg)


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

- **Channels** (Popup menu, Default: Yellow-Blue)
  Selects which two complementary color channels to use.
  - **Yellow-Blue**: the result is made using yellow and blue.
  - **Cyan-Red**: the result is made using cyan and red.
  - **Magenta-Green**: the result is made using magenta and purple.

- **Rotate Channels** (Default: 0, Range: any)
  Allows hue shifting the two color channels selected above. Note that when this is non-zero, the channels may no longer match the name selected.

- **Blur Channel1** (Default: 0, Range: 0 or greater)
  Smooths the first color channel by this amount.

- **Blur Channel2** (Default: 0, Range: 0 or greater)
  Smooths the second color channel by this amount.

- **Mix Original** (Default: 0, Range: any)
  Interpolates between the 2-color result and the original source. Set this to 1 for the original, or use negative values to exaggerate the dog vision effect.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Offset Darks** (Default: 0, Range: -8 to 2)
  Adds this gray value to the darker regions of the result. This can be negative to increase contrast.

- **Saturation** (Default: 1, Range: -2 to 8)
  Scales the color saturation. Increase for more intense colors. Set to 0 for monochrome.

- **Weight Source R** (Default: 1, Range: any)
  Scales the red of the input clip before processing.

- **Weight Source G** (Default: 1, Range: any)
  Scales the green of the input clip before processing.

- **Weight Source B** (Default: 1, Range: any)
  Scales the blue of the input clip before processing.

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

