---
title: GlowOrthicon
---

## S_GlowOrthicon

The source clip is darkened at areas around parts of the source clip that are
brighter than the given threshold, to give an 'orthicon' or 'dark glow'
look. Lower the Threshold parameter to produce the orthicon effect in more
areas. Adjust the Darkness and Width parameters to give different types of
looks.

In the Sapphire Lighting effects submenu.

![GlowOrthicon](../_static/GlowOrthicon.jpg)


### Inputs:

- **Source**: The current layer. The input clip that determines the locations to be darkened.

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

- **Darkness** (Default: 1, Range: 0 or greater)
  Scales the amount of darkening.

- **Color** (Default rgb: [0 0 0])
  Scales the color of the glows. The colors and brightnesses of the glows are also affected by the Source and Matte inputs.

- **Threshold** (Default: 0.7, Range: 0 or greater)
  Darkening will occur around locations in the source clip that are brighter than this value. A value of 0.9 causes dark glows from only the brightest spots. A value of 0 causes glows for every non-black area.

- **Threshold Add Color** (Default rgb: [0 0 0])
  This can be used to raise the threshold on a specific color and thereby reduce the glows generated on areas of the source clip containing that color.

- **Darks Width** (Default: 0.224, Range: 0 or greater)
  Scales the dark glow distance. This and all the width parameters can be adjusted using the Width Widget.

- **Protect Width** (Default: 0.1, Range: 0 or greater)
  The distance around the bright areas that is protected from darkening. This should normally be less than the value of Darks Width.

- **Protect Amount** (Default: 1, Range: 0 or greater)
  The amount that the bright areas are protected from darkening.

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
  Scales the brightness of the background. This parameter only has an effect if the background input is provided, and is visible due to a partially transparent Source image or a reduced Source Opacity parameter value.

- **Show** (Popup menu, Default: Result)
  Selects the type of output
  - **Result**: Shows the final result of combining the glow, source, and background.
  - **Threshold**: Shows the thresholded image that is used to generate the glow.

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

- **Show Darks Width** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Darks Width parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

