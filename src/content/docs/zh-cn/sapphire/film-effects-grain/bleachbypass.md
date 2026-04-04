---
title: BleachBypass
---

## S_BleachBypass

Simulates a film process in which silver is not
removed from the negative. The result has increased contrast and
reduced color saturation.

In the Sapphire Stylize effects submenu.

![BleachBypass](../_static/BleachBypass.jpg)


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

- **Amount** (Default: 1, Range: 0 or greater)
  Controls the intensity of the effect by interpolating between the original source and the result.

- **Soft Focus** (Default: 0, Range: 0 or greater)
  If positive, a soft focus effect is also applied. Increase for a broader soft focus look.

- **Sharpen** (Default: 0, Range: any)
  The amount of post-process sharpening applied.

- **Saturation** (Default: 1, Range: 0 to 10)
  Scales the color saturation. Increase for more intense colors. Set to 0 for monochrome.

- **Scale Lights** (Default: 1, Range: 0 or greater)
  Scales the result by this value. Increase for a brighter result.

- **Offset Darks** (Default: 0, Range: -8 to 2)
  Adds this gray value to the darker regions of the result. This can be negative to increase contrast.


### Grain Parameters:

Grain Amp:
*Default:
*0,
*Range:
*0 or greater.Scales the amplitude of the film grain that is added to
the result. Set this to 0 to disable all grain.

Grain Amp Red:
*Default:
*0.9,
*Range:
*0 or greater.Scales the red grain amplitude.

Grain Amp Green:
*Default:
*1,
*Range:
*0 or greater.Scales the green grain amplitude.

Grain Amp Blue:
*Default:
*1.6,
*Range:
*0 or greater.Scales the blue grain amplitude. Note that grain is
added and subtracted from the image, so for example, increasing
Grain Amp Blue will amplify both the blue and yellow speckles.

Grain Amp Darks:
*Default:
*0.2,
*Range:
*0 to 2.The relative amount of grain applied to the darkest
regions of the image, per channel. This defaults to less than 1.0
because dark areas usually have less grain than midtones.

Grain Amp Brights:
*Default:
*0,
*Range:
*0 to 2.The relative amount of grain applied to the
brightest regions of the image, per channel. This defaults to zero
because bright areas usually have less grain than midtones. Note
that highly saturated colors can be affected by both Grain Amp Darks
and Grain Amp Brights because they are dark in some color channels
and bright in others.

Grain Blur:
*Default:
*0,
*Range:
*0 or greater.The grain is smoothed by this amount. Increase for
coarser grain.

Grain Blur Red:
*Default:
*1,
*Range:
*0 or greater.The relative blur amount for the red grain.

Grain Blur Green:
*Default:
*0.9,
*Range:
*0 or greater.The relative blur amount for the green grain.

Grain Blur Blue:
*Default:
*1.2,
*Range:
*0 or greater.The relative blur amount for the blue grain.

Grain Mono:
*Check-box, Default:
*off.When enabled, the same grain pattern is used for the
red, green, and blue channels. To make truly monochrome grain you
should also set Grain Amp Red/Green/Blue equal to each other, make
sure Midtone Pos Red/Green/Blue are equal, and if GrainBlur is
positive also set Grain Blur Red/Green/Blue equal

Grain Seed:
*Default:
*0,
*Range:
*0 or greater.
Initializes the random number generator for the grain
generation. The actual seed value is not significant, but different
seeds give different grain patterns and the same value should give a
repeatable pattern.

### Other Parameters:

Scale Colors:
*Default rgb:
*[1 1 1].Scales the color of the result. For example, if it is yellow
[1 1 0], the blue of the result will be 0.

Opacity:
*Popup menu, Default: Normal
*.Determines the method used for dealing with
opacity/transparency.
*All Opaque:
*Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).*Normal:
*Process opacity normally.*As Premult:
*Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

Mask Use:
*Popup menu, Default: Luma
*.Determines how the Mask input channels are used to make a
monochrome mask.
*Luma:
*the luminance of the RGB channels is used.*Alpha:
*only the Alpha channel is used.

Blur Mask:
*Default:
*0.05,
*Range:
*0 or greater.Blurs the Matte input by this amount before using. This
can provide a smoother transition between the matted and unmatted
areas. It has no effect unless the Matte input is provided.

Invert Mask:
*Check-box, Default:
*off.
If on, inverts the Matte input so the effect is applied
to areas where the Matte is black instead of white. This has no effect
unless the Matte input is provided.
