---
title: Rays
---

## S_Rays

Generates beams of light emitting from the bright areas of
the source clip. Lower the Threshold parameter to generate rays from
more areas or raise it to generate rays from only the brightest
areas. Set the Rays Res parameter to 1/2 for faster rendering with
slightly softer rays.

In the Sapphire Lighting effects submenu.

![Rays](../_static/Rays.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Background**: Defaults to None. The clip to use as background.

- **Matte**: Defaults to None. If provided, the ray colors are scaled by this input. A monochrome matte can be used to choose a subset of areas that will generate rays. If the Matte Type is set to Color, a color matte input can be used to selectively adjust the ray colors in different regions. This input can optionally be blurred or inverted using the Blur Matte or Invert Matte parameters.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: Light Rays)
  Selects between light and dark rays.
  - **Light Rays**: Generates beams of light emitting from the bright areas of the source.
  - **Dark Rays**: Generates beams of darkness emitting from the dark areas of the source.

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
  The location from which the rays beam outwards. This parameter can be adjusted using the Center Widget.

- **Center Uses Mocha** (Check-box, Default: off)
  Controls whether the center is controlled by the Center parameter or follows the Center tracked inside of Mocha.

- **Smooth Center Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Rays Length** (Default: 0.25, Range: -5 to 1)
  The length of the rays. A length of 1.0 gives rays that continue forever, although they may still fade out as they go. To make the rays look longer you can also increase the Bias Outer Bright parameter. If Rays Length is negative the rays can beam inwards instead of outwards. Note that processing times increase for longer rays. This parameter can be adjusted using the Center Widget.

- **Length Red** (Default: 1, Range: 0 or greater)
  The relative length of the red channel of the rays. Adjust this, along with Length Green and Length Blue, to create color fringing effects.

- **Length Green** (Default: 1, Range: 0 or greater)
  The relative length of the green channel of the rays.

- **Length Blue** (Default: 1, Range: 0 or greater)
  The relative length of the blue channel of the rays.

- **Reverse Rays** (Default: 0, Range: 0 or greater)
  Extend rays inward as well as outward. The length of the reversed rays is controlled by Rays Length as well as this parameter.

- **Rays Brightness** (Default: 3, Range: 0 or greater)
  Scales the brightness of the ray beams.

- **Rays Darkness** (Default: 3, Range: 0 or greater)
  Scales the intensity of the dark ray beams.

- **Blur Rays** (Default: 0, Range: 0 or greater)
  Blur the rays only, before applied to the Source image.

- **Blur Rays Rel** (X & Y, Default: [1 1], Range: 0 or greater)
  The relative horizontal and vertical blur widths applied to the rays. Set Blur Rays Rel X to 0 for a vertical-only blur, or set Blur Rays Rel Y to 0 for a horizontal-only blur.

- **Rays Color** (Default rgb: [1 1 1])
  Scales the color of the ray beams.

- **Rays Color** (Default rgb: [0 0 0])
  Scales the color of the ray beams.

- **Bias Outer Bright** (Default: 0, Range: 0 to 1)
  Determines the variable amount of brightness along the rays. This is normally near 0 so the rays fade away at their outer ends, 0.5 causes equal brightness along the rays, and 1.0 causes maximum brightness at the ends.

- **Rays Res** (Popup menu, Default: Full)
  Selects the resolution factor for the rays. Higher resolutions give sharper rays, lower resolutions give smoother rays and faster processing. This 'Res' factor only affects the rays: the background is still combined with the rays at full resolution.
  - **Full**: Full resolution is used.
  - **Half**: The rays are calculated at half resolution.
  - **Quarter**: The rays are calculated at quarter resolution.

- **Threshold** (Default: 0.5, Range: 0 or greater)
  Rays are generated from locations in the source clip that are brighter than this value. A value of 0.9 causes rays at only the brightest spots. A value of 0 causes rays for every non-black area.

- **Threshold Add Color** (Default rgb: [0 0 0])
  This can be used to raise the threshold on a specific color and thereby reduce the rays generated on areas of the source clip containing that color.

- **Shimmer Amp** (Default: 0.5, Range: 0 or greater)
  Modulates the ray source image with this amount of noise texture to give the rays a shimmering look.

- **Shimmer Freq** (Default: 40, Range: 0.1 or greater)
  The frequency of the shimmer texture. Increase for a finer grained shimmer effect, decrease for larger, softer shimmer. This has no effect unless Shimmer Amp is positive.

- **Shimmer Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator for the shimmer texture. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Shimmer Shift** (X & Y, Default: [0 0], Range: any)
  Translation of the shimmer texture. This has no effect unless Shimmer Amp is positive.

- **Shimmer Speed** (X & Y, Default: [0 0], Range: any)
  Translation speed of the shimmer texture. If non-zero, the shimmering is automatically animated to shift at this rate.

- **Atmosphere Amp** (Default: 0, Range: 0 or greater)
  Atmosphere gives the effect of rays shining through a dusty atmosphere and picking up light or getting shadowed. This parameter adjusts the amount, or amplitude, of the atmospheric effect. Zero gives smooth rays, higher values give more dusty look.

- **Atmosphere Freq** (Default: 1, Range: 0.1 to 20)
  Controls the spatial frequency of the atmospheric noise. Turn this up higher to get finer details, turn down for broader overall variation.

- **Atmosphere Detail** (Default: 0.6, Range: 0 to 1)
  Controls the amount of fine detail in the atmosphere simulation. Decrease to get smoother atmosphere, increase for a more crunchy or grainy look.

- **Atmosphere Speed** (Default: 1, Range: any)
  The cloudy noise in the atmosphere evolves over time like real dust clouds; this parameter controls how fast the cloud pattern changes over time. Set to zero for a static pattern.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the rays. The maximum of the red, green, and blue ray brightness is scaled by this value and combined with the background Alpha at each pixel.

- **Rays From Alpha** (Default: 0, Range: 0 to 1)
  Set to 1 to generate rays from the source's alpha channel instead of its RGB channels. This will typically cause many more rays to be generated. Values between 0 and 1 interpolate between using the RGB and the Alpha.

- **Rays Under Source** (Default: 0, Range: 0 to 1)
  Set to 1 to composite the Source input over the rays.

- **Source Opacity** (Default: 1, Range: 0 to 1)
  Scales the opacity of the Source input when combined with the rays. This does not affect the generation of the rays themselves.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the background before combining with the rays. This parameter only has an effect if the background input is provided.

- **Use Source Chroma** (Default: 1, Range: 0 or greater)
  If this is 1, the chroma of the Source input affects the chroma of the resulting rays. If it is 0, only the brightness of the Source input affects the brightness of the rays, and the rendering speed should also be faster. Values between 0 and 1 interpolate between these two options.

- **Matte Type** (Popup menu, Default: Luma)
  This setting is ignored unless the Matte input is provided.
  - **Luma**: uses the luminance of the Matte input to scale the brightness of the rays.
  - **Color**: uses the RGB channels of the Matte input to scale the colors of the rays.
  - **Alpha**: uses the alpha channel of the Matte input to scale the brightness of the rays.

- **Blur Matte** (Default: 0, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Center** (Check-box, Default: on)
  Turns on or off the screen user interface widget for adjusting the Center and Rays Length parameters.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

