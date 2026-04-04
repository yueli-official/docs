---
title: CloudsPsyko
---

## S_CloudsPsyko

Generates a procedural noise texture, and passes this through a colorizing
process. The Shift Speed parameters cause the pattern to automatically
translate over time, and Phase Speed causes the colors to rotate over
time.

In the Sapphire Render effects submenu.

![CloudsPsyko](../_static/CloudsPsyko.jpg)


### Inputs:

- **Background**: The current layer. The clip to combine the clouds image with. This may be ignored if the Combine option is set to Clouds Only.

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

- **Frequency** (Default: 2, Range: 0.01 or greater)
  The spatial frequency of the clouds. Increase to zoom out, decrease to zoom in. Very high values of Frequency are clamped internally so the grain size is no smaller than a few pixels. If you want even finer grain use S_Grain or S_Clouds:Perspective instead.

- **Frequency Rel X** (Default: 0.4, Range: 0.01 or greater)
  The relative horizontal frequency of the texture. Increase to stretch it vertically or decrease to stretch it horizontally.

- **Octaves** (Integer, Default: 8, Range: 1 to 10)
  The number of summed layers of noise. Each octave is twice the frequency and half the amplitude of the previous. A single octave gives a smooth texture. Adding octaves makes the result approach a fractal (1/f) noise texture.

- **Seed** (Default: 0.23, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Boiling Mode** (Check-box, Default: off)
  When enabled, the clouds will boil, or evolve, over time. This mode takes slightly more computation but usually looks better.

- **Boil Details** (Default: 0.55, Range: 0 to 1)
  Increases or decreases the amount of fine detail in the clouds. Decrease to get a smoother look, increase to get a more high-frequency, noisy look. Only used when Boiling Mode is enabled.

- **Boil Speed** (Default: 1, Range: any)
  Sets the speed of the cloud boiling. Zero gives no boiling at all. Only used when Boiling Mode is enabled.

- **Color Freq** (Default: 4, Range: 0.01 or greater)
  The frequency of the color pattern. Increase for a busier texture with more cycles through the spectrum.

- **Freq Red** (Default: 1, Range: 0 or greater)
  The frequency of the red color component. Increase for more cycles in the red channel.

- **Freq Green** (Default: 1.1, Range: 0 or greater)
  The frequency of the green color component. Increase for more cycles in the green channel.

- **Freq Blue** (Default: 1.2, Range: 0 or greater)
  The frequency of the blue color component. Increase for more cycles in the blue channel.

- **Shift Start** (X & Y, Default: [0 0], Range: any)
  Translation offset of the texture. Since the texture is procedurally generated it can be shifted with no repeating units or seams occurring.

- **Shift Speed** (X & Y, Default: [0.5 0], Range: any)
  Translation speed of the texture. If non-zero, the result is automatically animated to shift at this rate. The result of animated Speed values may not be intuitive, so for variable speed motion it is usually best to set this to 0 and animate the Shift Start values instead.

- **Phase Start** (Default: -0.5, Range: any)
  The phase offset of the color patterns.

- **Phase Speed** (Default: 0.3, Range: any)
  The phase speed of the color patterns. If non-zero, the phase is automatically animated to give the color pattern a boiling look.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Scale Color** (Default rgb: [1 1 1])
  Scales the color of the result. For example, if it is yellow [1 1 0], the blue of the result will be 0.

- **Saturation** (Default: 1, Range: 0 or greater)
  Scales the strength of the colors. Increase for more intense colors, or decrease for muted colors.

- **Offset** (Default: 0, Range: any)
  Adds this gray value to the result (or subtracts if negative). 0 has no effect, .5 is middle gray, and 1 is white.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  The background brightness is scaled by this value before being combined with the clouds.

- **Combine** (Popup menu, Default: Clouds Only)
  Determines how the texture is combined with the Background.
  - **Clouds Only**: gives only the clouds texture with no Background.
  - **Mult**: the texture is multiplied by the Background.
  - **Add**: the texture is added to the Background.
  - **Screen**: the texture is blended with the Background using a screen operation.
  - **Difference**: the result is the difference between the texture and Background.
  - **Overlay**: the texture is combined with the Background using an overlay function.

- **Input Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Output Opacity** (Popup menu, Default: Copy From Input)
  Determines the opacity/transparency of the result. This effect does not process the opacity (alpha channel) of its input but it can either copy the opacity from the input, or output a fully opaque result.
  - **All Opaque**: Makes the result fully opaque with no
transparency.
  - **Copy From Input**: Copies the opacity/transparency from
the current layer given to this effect.

- **Mask Use** (Popup menu, Default: Luma)
  Determines how the Mask input channels are used to make a monochrome mask.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Mask** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

