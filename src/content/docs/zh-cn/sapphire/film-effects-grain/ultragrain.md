---
title: UltraGrain
---

## S_UltraGrain

Adds simulated digital camera grain to the source clip.

In the Sapphire Stylize effects submenu.
### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Matte**: Defaults to None. Interpolate between the result and the Source input. White areas use the result of the effect. Black areas use the Source clip.


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


### Grain Parameters:

Grain Amp:
*Default:
*1,
*Range:
*0 to 2.Scales the amplitude of the grain that is added to
the result. Set this to 0 to disable all grain.

Grain Size:
*Default:
*1,
*Range:
*0.001 or greater.Scales the size of the grain that is added to the result.
Useful when changing resolution, set this scale factor to match the change in resolution.

Grain Blur:
*Default:
*0.00567,
*Range:
*0 or greater.The grain is smoothed by this amount. Increase for
coarser grain.

Grain Amp Red:
*Default:
*3.29,
*Range:
*0 or greater.Scales the red grain amplitude.

Grain Amp Green:
*Default:
*3,
*Range:
*0 or greater.Scales the green grain amplitude.

Grain Amp Blue:
*Default:
*3.52,
*Range:
*0 or greater.Scales the blue grain amplitude. Note that grain is
added and subtracted from the image, so for example, increasing
Grain Amp Blue will amplify both the blue and yellow speckles.

Grain Blur Red:
*Default:
*0.572,
*Range:
*0 or greater.The relative blur amount for the red grain.

Grain Blur Green:
*Default:
*0.561,
*Range:
*0 or greater.The relative blur amount for the green grain.

Grain Blur Blue:
*Default:
*0.528,
*Range:
*0 or greater.The relative blur amount for the blue grain.

Grain Amp Darks:
*Default:
*0.37,
*Range:
*0 to 2.The relative amount of grain applied to the darkest
regions of the image, per channel. Dark source intensity in the image
is defined as black (0 0 0).

Grain Amp Mids:
*Default:
*1,
*Range:
*0 to 2.The relative amount of grain applied to the midtone
regions of the image, per channel. Midtone source intensity in the image
is defined by the Midtone Pos parameters.

Grain Amp Brights:
*Default:
*0,
*Range:
*0 to 2.The relative amount of grain applied to the
brightest regions of the image, per channel. Bright source intensity in
the image is defined as white (1 1 1).

Midtone Pos Red:
*Default:
*0.27,
*Range:
*0 to 1.The position of the midtones in the red channel. The red
grain amplitude is interpolated from Grain Amp Darks at black, to
Grain Amp Mids at this midtone position, then to Grain Amp Brights at
white. This whole curve is then scaled by the Grain Amp Red
parameter.

Midtone Pos Green:
*Default:
*0.3,
*Range:
*0 to 1.The position of the midtones in the green channel. The green
grain amplitude is interpolated from Grain Amp Darks at black, to
Grain Amp Mids at this midtone position, then to Grain Amp Brights at
white. This whole curve is then scaled by the Grain Amp Green
parameter.

Midtone Pos Blue:
*Default:
*0.31,
*Range:
*0 to 1.The position of the midtones in the blue channel. The blue
grain amplitude is interpolated from Grain Amp Darks at black, to
Grain Amp Mids at this midtone position, then to Grain Amp Brights at
white. This whole curve is then scaled by the Grain Amp Blue
parameter.

Grain Mono:
*Check-box, Default:
*off.When enabled, the same grain pattern is used for the
red, green, and blue channels. To make truly monochrome grain you
should also set Grain Amp Red/Green/Blue equal to each other, make
sure Midtone Pos Red/Green/Blue are equal, and if GrainBlur is
positive also set Grain Blur Red/Green/Blue equal

Grain Seed:
*Default:
*0.123,
*Range:
*0 or greater.
Initializes the random number generator for the grain
generation. The actual seed value is not significant, but different
seeds give different grain patterns and the same value should give a
repeatable pattern.
