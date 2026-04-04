---
title: JpegDamage
---

## S_JpegDamage

Creates a version of the Source input that is
subjected to Jpeg compression artifacts and errors. This can be
used to give various looks of low quality digital transmissions.
Three methods for manipulating your image are provided: the Jpeg
quality can be adjusted, various internal frequencies can be scaled,
and random decompression errors can be introduced. In all cases it
can also be useful to lower the resolution factor to create larger, more
obvious Jpeg blocks.

In the Sapphire Stylize effects submenu.

![JpegDamage](../_static/JpegDamage.jpg)


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

- **Quality** (Default: 0.1, Range: 0.01 to 1)
  Determines the amount of normal Jpeg artifacts. Use lower values for more compression.

- **Res Factor** (Integer, Default: 1, Range: 1 or greater)
  Downres the result by the inverse of this amount, so 1 is full resolution, 2 is 1/2, 3 is 1/3, etc. The pixel shapes will be larger when this is increased. You won't notice the result of this parameter unless its value is beyond then the current viewing downres factor.

- **Res Rel X** (Default: 1, Range: 0.01 or greater)
  Downres the result by the inverse of this amount in the horizontal direction. The jpeg block shapes will become rectangular if this is not 1.

- **All Freq Scale** (Default: 1, Range: 0 or greater)
  Scales the frequencies for all Jpeg coefficients. Values other than 1 cause abnormal results, and create unusual looking blocky versions of your input.

- **X Freq Scale** (Default: 1, Range: 0 or greater)
  Scales the horizontal Jpeg frequencies. Values other than 1 cause abnormal results.

- **Y Freq Scale** (Default: 1, Range: 0 or greater)
  Scales the vertical Jpeg frequencies.

- **Low Freq Scale** (Default: 1, Range: 0 or greater)
  Scales the softer low frequencies.

- **Mid Freq Scale** (Default: 1, Range: 0 or greater)
  Scales the middle range frequencies.

- **High Freq Scale** (Default: 1, Range: 0 or greater)
  Scales the sharper high frequencies. You may need a high Quality setting to see the high frequencies at all.

- **Affect Luma** (Default: 1, Range: 0 or greater)
  Determines how much the Freq Scale parameters above affect the luminance channel. A zero value causes no luminance change. Values greater than 1.0 exaggerate the change.

- **Affect Chroma** (Default: 0.5, Range: 0 or greater)
  Determines how much the Freq Scale parameters above affect the chroma channels. A zero value causes no chroma change. Values greater than 1.0 exaggerate the change.

- **Error Rate** (Default: 0, Range: 0 or greater)
  If positive, random decompression errors are introduced. The value determines the average number of errors in those blocks that receive errors. Larger values give a more even grainy look.

- **Err Block Density** (Default: 0.75, Range: 0 to 1)
  Determines the percentage of Jpeg blocks with errors. A value of .5 will give errors in half of the blocks and 1.0 will give errors in all blocks.

- **Error Amp** (Default: 1, Range: 0 or greater)
  The amplitude of the decompression errors. Larger values give more visually obvious errors. This has no effect unless the Error Rate is also positive.

- **Error Coherence** (Default: 1, Range: 0 or greater)
  Determines how much the blocks with errors are grouped together. When zero, the errors are evenly distributed throughout the frame. When increased, the errors are clustered into larger groups. This has no effect unless the Error Rate is positive and the Err Block Density is less than 1.

- **Jitter Frames** (Integer, Default: 1, Range: 0 or greater)
  If this is 0, the random errors will remain the same for every frame processed. If it is 1, different errors are used for each frame. If it is 2, new errors are used for every other frame, and so on. This has no effect unless the Error Rate is also positive.

- **Rand Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different random error patterns, and the same value should give a repeatable result. This has no effect unless the Error Rate is also positive.

- **Scale Lights** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result by this amount.

- **Offset Darks** (Default: 0, Range: any)
  Adds this gray value to the darker regions of the source. This can be negative to increase contrast.

- **Saturation** (Default: 1, Range: any)
  Scales the color saturation. Increase for more intense colors. Set to 0 for monochrome.

- **Flip Noise Vertically** (Check-box, Default: off)
  Flips noise vertically if needed to achieve a consistent look.

- **Mask Use** (Popup menu, Default: Luma)
  Determines how the Mask input channels are used to make a monochrome mask.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Mask** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

