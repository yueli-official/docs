---
title: TextureFlux
---

## S_TextureFlux

Creates abstract textures of fluctuating liquid or cellular patterns.
The Morph Speed parameter causes the pattern to automatically undulate over time.

In the Sapphire Render effects submenu.

![TextureFlux](../_static/TextureFlux.jpg)


### Inputs:

- **Background**: The current layer. The clip to combine the texture image with. This may be ignored if the Combine option is set to Texture Only.

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

- **Frequency** (Default: 4, Range: 0.01 to 100)
  The spatial frequency of the texture. Increase to zoom out, decrease to zoom in.

- **Frequency Rel X** (Default: 1, Range: 0.01 or greater)
  The relative horizontal frequency of the texture. Increase to stretch it vertically or decrease to stretch it horizontally.

- **Octaves** (Integer, Default: 2, Range: 1 to 10)
  The number of summed layers of noise. Each octave is twice the frequency and half the amplitude of the previous. A single octave gives a smooth texture. Adding octaves makes the result approach a fractal (1/f) noise texture.

- **Seed** (Default: 0.234, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Shift** (X & Y, Default: [0 0], Range: any)
  Translation offset of the texture. Since the texture is procedurally generated it can be shifted with no repeating units or seams occurring.

- **Bubble Amount** (Default: 1, Range: any)
  Amplitude of warping applied to generate the bubble shapes. This can be negative to generate brighter bubble shapes with dark edges.

- **Bubble Smooth** (Default: 0.25, Range: 0.008 or greater)
  Smooths the warping pattern by this amount.

- **Rotate Warp Dir** (Default: 0, Range: any)
  Rotates the direction of the warping. This can cause a twisting effect or an inverted effect when set to 180.

- **Morph Speed** (Default: 0.3, Range: any)
  Speed to automatically undulate the underlying noise pattern over time.

- **Morph** (X & Y, Default: [0 1], Range: any)
  The horizontal and vertical directions to undulate the underlying noise pattern, when using Morph Speed.

- **Brightness1** (Default: 1, Range: 0 or greater)
  Scales the brightness of Color1. Increase for more contrast.

- **Color1** (Default rgb: [1 1 1])
  The color of the 'brighter' parts of the texture. The colors of the result are determined by an interpolation between Color0 and Color1.

- **Color0** (Default rgb: [0 0.12 0.15])
  The color of the 'darker' parts of the texture.

- **Offset0** (Default: 0, Range: any)
  Adds this value to color0. Decrease to a negative value for more contrast.

- **Filter** (Check-box, Default: off)
  The type of convolution filter to blur with.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  The background brightness is scaled by this value before being combined with the texture.

- **Combine** (Popup menu, Default: Texture Only)
  Determines how the texture is combined with the Background.
  - **Texture Only**: gives only the texture image with no Background.
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

