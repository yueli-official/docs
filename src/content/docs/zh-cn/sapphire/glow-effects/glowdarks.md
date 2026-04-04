---
title: GlowDarks
---

## S_GlowDarks

Areas of the source clip darker than the given threshold are blurred and
combined with the input clip to give a deep smoky look. Adjust the
Darkness, Width, and Threshold parameters to give different types of looks.

In the Sapphire Lighting effects submenu.

![GlowDarks](../_static/GlowDarks.jpg)


### Inputs:

- **Source**: The current layer. The input clip that determines the glow locations and colors.

- **Background**: Defaults to None. The clip to combine the glows with. If no background is given, the Source is also used as the Background.

- **Matte**: Defaults to None. If provided, the source glow colors are scaled by this input. A monochrome matte can be used to choose a subset of Source areas that will generate glows. A color matte can be used to selectively adjust the glow colors in different regions. The matte is applied to the source before the glows are generated so it will not clip the resulting glows.


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

- **Darkness** (Default: 0.5, Range: 0 or greater)
  The magnitude of the dark glows.

- **Threshold** (Default: 0.5, Range: 0 or greater)
  Dark glows will be generated from locations in the source clip that are darker than this value. A value of 0.1 causes glows at only the darkest areas. A value of 1.0 causes glows on every non-white area.

- **Glow Saturation** (Default: 1, Range: -2 to 8)
  Scales the saturation of the dark colors. Increase for more intense colors.

- **Glow Width** (Default: 1, Range: 0 or greater)
  Scales the glow distance. This and all the width parameters can be adjusted using the Width Widget. Note that a zero glow width still affects the dark areas; set the darkness parameter to zero if you want to pass the Source through unchanged.

- **Width X** (Default: 1, Range: 0 or greater)
  Scales the horizontal glow width. Set to 0 for vertical only.

- **Width Y** (Default: 1, Range: 0 or greater)
  Scales the vertical glow width. Set to 0 for horizontal only.

- **Subpixel** (Check-box, Default: on)
  Enables glowing by subpixel widths. Use this for smoother animation of the Width parameters.

- **Glow From Alpha** (Default: 0, Range: 0 to 1)
  Set to 1 to generate glows from the alpha channel of the source input instead of the RGB channels. In this case the glows will not pick up color from the source and will typically be brighter. Values between 0 and 1 interpolate between using the RGB and the Alpha.

- **Glow Under Source** (Default: 0, Range: 0 to 1)
  Set to 1 to composite the Source input over the glows.

- **Source Opacity** (Default: 1, Range: 0 to 1)
  Scales the opacity of the Source input when combined with the glows. This does not affect the generation of the glows themselves.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the background input clip.

- **Show** (Popup menu, Default: Result)
  Selects the type of output
  - **Result**: Shows the final result of combining the glow, source, and background.
  - **Threshold**: Shows the thresholded image that is used to generate the glow.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Atmosphere Amp** (Default: 0, Range: 0 or greater)
  Atmosphere gives the effect of the glow shining through a dusty atmosphere and picking up light or getting shadowed. This parameter adjusts the amount, or amplitude, of the atmospheric effect. Zero gives a smooth glow, higher values give more dusty look.

- **Atmosphere Freq** (Default: 1, Range: 0.1 to 20)
  Controls the spatial frequency of the atmospheric noise. Turn this up higher to get finer details, turn down for broader overall variation.

- **Atmosphere Detail** (Default: 0.6, Range: 0 to 1)
  Controls the amount of fine detail in the atmosphere simulation. Decrease to get smoother atmosphere, increase for a more crunchy or grainy look.

- **Atmosphere Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator for the atmospheric noise. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Atmosphere Speed** (Default: 1, Range: any)
  The cloudy noise in the atmosphere evolves over time like real dust clouds; this parameter controls how fast the cloud pattern changes over time. Set to zero for a static pattern.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Glow Width** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Glow Width parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

