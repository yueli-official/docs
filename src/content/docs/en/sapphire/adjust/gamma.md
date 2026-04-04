---
title: Gamma
---

## S_Gamma

Applies a gamma correction to the input clip. The red, green, and blue
channels can be adjusted independently. From Gamma just causes the inverse effect
of adjusting Gamma.

In the Sapphire Adjust effects submenu.

![Gamma](../_static/Gamma.jpg)


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

- **Gamma** (Default: 1, Range: 0.1 to 10)
  Values greater than 1.0 make the mid-tones brighter, values less than 1.0 make them darker, 1.0 leaves the input unchanged.

- **Gamma Red** (Default: 1, Range: 0.1 to 10)
  Brightens or darkens the red mid-tones.

- **Gamma Green** (Default: 1, Range: 0.1 to 10)
  Brightens or darkens the green mid-tones.

- **Gamma Blue** (Default: 1, Range: 0.1 to 10)
  Brightens or darkens the blue mid-tones.

- **From Gamma** (Default: 1, Range: 0.1 to 10)
  Divides the Gamma by this value before processing. This can be useful if your image was correct at this gamma, but needs to be adjusted from this to a new gamma.

- **From Gamma Red** (Default: 1, Range: 0.1 to 10)
  Darkens or brightens the red mid-tones.

- **From Gamma Green** (Default: 1, Range: 0.1 to 10)
  Darkens or brightens the green mid-tones.

- **From Gamma Blue** (Default: 1, Range: 0.1 to 10)
  Darkens or brightens the blue mid-tones.

- **Scale Lights** (Default: 1, Range: 0 or greater)
  Scales the brightness by this amount after the gamma correction. Increase for a brighter result.

- **Scale Lights Red** (Default: 1, Range: 0 or greater)
  Scales the red by this amount after the gamma correction.

- **Scale Lights Green** (Default: 1, Range: 0 or greater)
  Scales the red by this amount after the gamma correction.

- **Scale Lights Blue** (Default: 1, Range: 0 or greater)
  Scales the red by this amount after the gamma correction.

- **Offset Darks** (Default: 0, Range: -8 to 2)
  Adds this gray value to the darker regions after the gamma correction. This can be negative to increase contrast.

- **Offset Darks Red** (Default: 0, Range: -8 to 2)
  Adds this red value to the darker red regions after the gamma correction. This can be negative to increase contrast.

- **Offset Darks Green** (Default: 0, Range: -8 to 2)
  Adds this green value to the darker green regions after the gamma correction. This can be negative to increase contrast.

- **Offset Darks Blue** (Default: 0, Range: -8 to 2)
  Adds this blue value to the darker blue regions after the gamma correction. This can be negative to increase contrast.

- **Mask Use** (Popup menu, Default: Luma)
  Determines how the Mask input channels are used to make a monochrome mask.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Mask** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

