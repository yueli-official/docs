---
title: Luna
---

## S_Luna

Renders the Earth's Moon; you can adjust phase and colors, and add atmospheric effects.

In the Sapphire Render effects submenu.

![Luna](../_static/Luna.jpg)


### Inputs:

- **Background**: The current layer. The clip to use as background.

- **Mask**: Defaults to None. Interpolate between the result and the Source input. White areas use the result of the effect. Black areas use the Source clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: Luna)
  Selects how the moon's phase is chosen. You can adjust it directly in Luna mode, or select LunaDate mode to choose a date and time and the effect will use the proper phase for that date.
  - **Luna**: Select this mode to adjust the moon phase manually.
  - **LunaDate**: Select this mode to have the effect compute the phase from the given date and time.

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

- **Center** (X & Y, Default: [-0.44 0.17], Range: any)
  Center point of the moon. This parameter can be adjusted using the Center Widget.

- **Center Uses Mocha** (Check-box, Default: off)
  Controls whether the moon's center is controlled by the Center parameter or follows the Center tracked inside of Mocha.

- **Smooth Center Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Size** (Default: 0.3, Range: 0 or greater)
  Size of the moon. This parameter can be adjusted using the Size Widget.

- **Lunar Phase** (Default: 65, Range: any)
  Phase of the moon, in degrees; 0 is new, 90 is first quarter, 180 is full, and 270 is last quarter. Only available in Luna mode.

- **Year** (Integer, Default: 2.02e+03, Range: 1900 to 2295)
  Year to use when computing the phase.

- **Month** (Integer, Default: 3, Range: 1 to 12)
  Month to use when computing the phase.

- **Day** (Integer, Default: 16, Range: 1 to 31)
  Day to use when computing the phase.

- **Hour** (Integer, Default: 16, Range: 0 to 23)
  Hour to use when computing the phase.

- **Minute** (Default: 0, Range: any)
  Minute to use when computing the phase.

- **GMT Offset** (Default: -5, Range: -12 to 12)
  GMT offset to use when computing the phase. -5 is Eastern Standard Time, -8 is Pacific Standard Time.

- **Rotation** (Default: 30, Range: any)
  Rotation of the moon image, in degrees.

- **Bumpiness** (Default: 0.3, Range: 0 to 1)
  The moon has craters that catch and reflect light. This parameter can be used to adjust how bumpy those craters look. 0 is completely smooth, 1 is very rough. 0.3 is about physically realistic on a clear night.

- **Contrast** (Default: 1, Range: 0 to 1)
  Adjusts the contrast of the moon. Values toward 0 brighten the dark areas.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Color** (Default rgb: [1 1 1])
  Scales the color of the result. For example, if it is yellow [1 1 0], the blue of the result will be 0.

- **Earth Glow** (Default rgb: [0 0 0])
  Adds earth-glow, which you often see near sunset when the moon is crescent. The sun's light reflects off the earth, and some of that reflected light illuminates even the dark part of the moon. This gives an especially nice look during a lunar eclipse.

- **Gamma** (Default: 1.6, Range: 0.1 or greater)
  Sets the overall gamma of the moon image. Good for reducing contrast in a different way from the contrast parameter.

- **Sky Color** (Default rgb: [0 0 0])
  If you want to make a complete sky image with the moon and a colored sky, you can put the moon in a blue sky by setting Sky Color to blue. This will also tint the moon toward the sky color.

- **Glow Brightness** (Default: 0.5, Range: 0 or greater)
  Adds some glow to the moon. You can see this often in real life when there's some haze or light clouds.

- **Threshold** (Default: 0.01, Range: 0 or greater)
  Threshold for the glow; only parts of the moon brighter than this threshold will glow.

- **Glow Size** (Default: 0.75, Range: 0 or greater)
  Size of the moon glow. Larger creates a more diffuse glow.

- **Halo Brightness** (Default: 0, Range: 0 or greater)
  With certain kinds of high, diffuse clouds, you can sometimes see a subtle rainbow halo around the moon. Increase this parameter to see that halo.

- **Halo Rel Size** (Default: 1.5, Range: 0 or greater)
  Sets the size of the halo, relative to the moon. 1.0 would be the same size as the moon, 2.0 is twice as large.

- **Color Fringing** (Default: 0.05, Range: 0 to 1)
  Increases or decreases the amount of color fringing in the moon halo; fringing separates the colors into a rainbow.

- **Inner Softness** (Default: 0.15, Range: 0 or greater)
  Sets the softness or spread of the inside of the halo, closest to the moon.

- **Outer Softness** (Default: 0.3, Range: 0 or greater)
  Sets the softness or spread of the outside of the halo, farthest from the moon.

- **Halo Saturation** (Default: 0.5, Range: 0 or greater)
  Sets the overall saturation of the moon halo. Increase for a more graphic look.

- **Halo Tint** (Default rgb: [1 1 1])
  Tints the halo toward this color.

- **Atmosphere Amp** (Default: 0.3, Range: 0 or greater)
  The Atmosphere params add a little noise to the glow and halo, for a more realistic look. Atmosphere Amp controls the amount of atmospheric noise.

- **Atmosphere Freq** (Default: 2, Range: 0.1 or greater)
  Controls the frequency of the atmospheric noise.

- **Atmosphere Turbulence** (Default: 0.6, Range: 0 to 1)
  Controls the turbulence (amount of detail) in the atmospheric noise.

- **Atmosphere Seed** (Default: 0.123, Range: 0 or greater)
  Sets the seed of the atmospheric noise.

- **Atmosphere Speed** (Default: 1, Range: any)
  Controls how fast the atmospheric noise changes over time.

- **Combine** (Popup menu, Default: Overlay)
  Combine allows you to combine the moon image with the background in various ways.
  - **Moon Only**: Ignore the background; show the moon (and its glow and halo) only.
  - **Overlay**: Overlay (composite) the moon over the background.
  - **Add**: Add the moon to the background.
  - **Screen**: Screen the moon with the background. Nice for daytime shots.
  - **Max**: Where the moon is brighter than the background, show
it. This can be useful for daytime shots with clouds. Where the
clouds are brighter than the moon, they'll obscure it.
  - **Transparent Shadow**: Composite only the lit part of the moon
over the background, leaving the dark part transparent. This is not
physically realistic, since the dark part of the moon obscures the
sky and stars behind it, but it can be used for graphic effect.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the background before combining with the moon. If 0, the result will contain only the moon image over black.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive, the output Alpha channel will include some opacity from the moon's halo and glow.

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

- **Show Size** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Center** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

