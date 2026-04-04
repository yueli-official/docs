---
title: MatteOps
---

## S_MatteOps

Grows, shrinks, or adds noise to the alpha channel of the
Source input. This can be useful for removing blue or green spill
from a chroma key.

In the Sapphire Composite effects submenu.

![MatteOps](../_static/MatteOps.jpg)


### Inputs:

- **Source**: The current layer. The input clip containing the matte to process. The matte is assumed to have anti-aliased but hard edges, because very soft edges might not be affected in a useful way.


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

- **Shrink- Grow+** (Default: 0, Range: any)
  Amount to grow the matte edges in approximate pixels, or shrink if negative.

- **Edge Softness** (Default: 1, Range: 0.01 or greater)
  The resulting softness of the edges.

- **Post Blur** (Default: 0, Range: 0 or greater)
  If positive, the result is blurred by this amount. This is an alternative method for softening the edges.

- **Filter** (Popup menu, Default: Triangle)
  The type of blur filter to use for the shrink or grow process.
  - **Box**: uses a rectangular shaped filter.
  - **Triangle**: smoother, uses a pyramid shaped filter.
  - **Gauss**: smoothest, uses a gaussian shaped filter.

- **Matte Use** (Popup menu, Default: Alpha)
  Determines how the Matte input channels are used to make a monochrome matte.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Invert Matte** (Check-box, Default: off)
  If enabled, the black and white of the output matte are inverted.

- **Soft Borders** (Check-box, Default: off)
  If enabled, transparent borders are added to the input image before processing. This allows the result to include soft edges beyond the original image size. When off, the effect only occurs within the frame and the result will retain an edge at the borders.

- **Output** (Popup menu, Default: RGBA)
  Selects the format of the output.
  - **Matte**: the processed Matte is output as white.
  - **RGBA**: the Alpha output channel receives the processed
Matte, and the RGB channels are passed through from the input
unchanged.
  - **RGBA Premult**: the Alpha output channel receives the
processed Matte, and the RGB channels receive the input multiplied
by the new Matte. This option can be appropriate if you are
shrinking a matte and need an RGBA result for pre-multiplied
compositing.
  - **Matte Premult**: the processed Matte is output on all channels.

- **Noise Amplitude** (Default: 0, Range: 0 or greater)
  The amount of noise texture to add to the edges.

- **Noise Width** (Default: 0.0224, Range: 0 or greater)
  The width of the area at the matte edges where the noise is included. This has no effect unless Noise Amplitude is positive

- **Frequency** (Default: 100, Range: 0.1 or greater)
  The frequency of the noise. Increase for finer grain noise, decrease for coarser noise. This has no effect unless Noise Amplitude is positive.

- **Frequency Rel X** (Default: 1, Range: 0.01 or greater)
  The relative horizontal frequency of the noise. Increase to stretch the noise vertically, decrease to stretch it horizontally. This has no effect unless Noise Amplitude is positive.

- **Octaves** (Integer, Default: 1, Range: 1 to 10)
  The number of summed layers of noise. Each octave is twice the frequency and half the magnitude of the previous. This has no effect unless Noise Amplitude is positive.

- **Seed** (Default: 0.23, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Noise Shift** (X & Y, Default: [0 0], Range: any)
  The horizontal and vertical translation of the noise texture.

- **Jitter Frames** (Integer, Default: 1, Range: 0 or greater)
  If this is 0, the noise texture will remain the same for every frame processed. If it is 1, a new noise texture is used for each frame. If it is 2, a new noise texture is used for every other frame, and so on.

