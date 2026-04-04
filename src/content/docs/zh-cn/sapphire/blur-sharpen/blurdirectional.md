---
title: BlurDirectional
---

## S_BlurDirectional

Blurs the source clip in a given direction using a gaussian,
triangle, or box filter. It can also blur each channel by different amounts.

In the Sapphire Blur+Sharpen effects submenu.

![BlurDirectional](../_static/BlurDirectional.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Matte**: Defaults to None. If provided, the blur is only performed on regions of the source clip specified by the bright areas of this input. Pixels outside this matte are not blurred, and do not contribute to the resulting blurred pixels within it. This input can be affected using the Invert Matte, or Matte Use parameters.


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

- **Blur Amount** (Default: 0.4, Range: 0 or greater)
  Scales the width of the blur. This parameter can be adjusted using the Blur Amount Widget.

- **Angle** (Default: 45, Range: any)
  The direction of the blur. An angle of 0 produces a horizontal blur, and an angle of 90 produces a vertical blur. This parameter can be adjusted using the Angle Widget.

- **Shift** (Default: 0, Range: any)
  Shifts the image in the direction of the blur. A negative shift amount shifts the image in the opposite direction.

- **Bias** (Default: 0.5, Range: 0 to 1)
  Varies the weight of the pixels along the path of the blur, which gives the appearance of trails or streaks in a single direction. A value of 0.5 weights all pixels evenly. A value of 1 causes the weight to increase toward the direction of the blur, while a value of 0 has the opposite effect.

- **Blur Red** (Default: 1, Range: 0 or greater)
  The blur width of the red channel, relative to Blur Amount.

- **Blur Green** (Default: 1, Range: 0 or greater)
  The blur width of the green channel, relative to Blur Amount.

- **Blur Blue** (Default: 1, Range: 0 or greater)
  The blur width of the blue channel, relative to Blur Amount.

- **Shift Red** (Default: 0, Range: any)
  Additional amount to shift the red color channel.

- **Shift Green** (Default: 0, Range: any)
  Additional amount to shift the green color channel.

- **Shift Blue** (Default: 0, Range: any)
  Additional amount to shift the blue color channel.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Offset Darks** (Default: 0, Range: -8 to 2)
  Adds this gray value to the darker regions of the result. This can be negative to increase contrast.

- **Mix With Source** (Default: 0, Range: 0 to 1)
  Interpolates between the blurred result (0) and the original source (1). 0.1 can give a nice misty effect since it mixes only a little of the source in.

- **Edge Mode** (Popup menu, Default: Reflect)
  Determines the behavior when accessing areas outside the source image.
  - **Transparent**: Areas outside the source image are treated as transparent, which can produce
transparency around the edges of the image.
Select this for fastest rendering.
  - **Repeat**: Repeats the last pixel outside the border of the image.
  - **Reflect**: Reflects the image outside the border.

- **Filter** (Popup menu, Default: Box)
  The type of convolution filter to blur with.
  - **Box**: uses a rectangular shaped filter.
  - **Triangle**: smoother, uses a pyramid shaped filter.
  - **Gauss**: smoothest, uses a gaussian shaped filter.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Matte Use** (Popup menu, Default: Luma)
  Determines how the Matte input channels are used to make a monochrome matte.
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

- **Show Blur Amount** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the blur amount parameters.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Angle** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Angle parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

