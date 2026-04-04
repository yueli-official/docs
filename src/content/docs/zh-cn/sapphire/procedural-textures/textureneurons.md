---
title: TextureNeurons
---

## S_TextureNeurons

Creates an abstract texture resembling moving nerve cell tendrils.
The Phase Speed and Morph Speed parameters cause the pattern to
automatically change over time.

In the Sapphire Render effects submenu.

![TextureNeurons](../_static/TextureNeurons.jpg)


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

- **Frequency** (Default: 3, Range: 0.01 or greater)
  The spatial frequency of the texture. Increase to zoom out, decrease to zoom in.

- **Frequency Rel X** (Default: 1, Range: 0.01 or greater)
  The relative horizontal frequency of the texture. Increase to stretch it vertically or decrease to stretch it horizontally.

- **Arms** (Integer, Default: 9, Range: 0 to 50)
  The number of tendrils emanating from each center point in the texture.

- **Softness** (Default: 0.5, Range: 0.01 or greater)
  Decrease for sharper line edges. Increase for smoother line edges or to reduce aliasing.

- **Thickness** (Default: 1.1, Range: 0 or greater)
  Decrease for thinner lines. Increase for stronger brighter lines.

- **Outer Bright** (Default: 0.4, Range: 0.01 to 1)
  Scales the brightness of the regions away from the neuron centers. Decrease to remove the connecting lines and leave only the star shapes at the centers.

- **Seed** (Default: 1, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Shift** (X & Y, Default: [0 0], Range: any)
  Translation offset of the texture. Since the texture is procedurally generated it can be shifted with no repeating units or seams occurring.

- **Phase Start** (Default: 0, Range: any)
  Amount to rotate the arms about their centers.

- **Phase Speed** (Default: 0.05, Range: any)
  Speed to automatically rotate the arms and move the lines over time.

- **Morph Speed** (Default: 0.05, Range: any)
  Speed to automatically undulate the underlying noise pattern over time.

- **Morph** (X & Y, Default: [1 0], Range: any)
  The horizontal and vertical directions to undulate the underlying noise pattern, when using Morph Speed.

- **Twist** (Default: 0, Range: any)
  Amount to rotate the centers to cause a twisting effect.

- **Wiggle Amp** (Default: 0.1, Range: 0 or greater)
  Amount of additional noise too apply along the pattern of lines. Turn down to get smoother lines.

- **Wiggle Freq Rel** (Default: 2, Range: 0 or greater)
  Frequency of the additional noise.

- **Wiggle Octaves** (Integer, Default: 4, Range: 1 to 10)
  The number of octaves to use for the additional noise.

- **Smooth** (Default: 0, Range: 0 or greater)
  Amount to blur the line pattern. Increase for a defocus look, or to help remove aliasing artifacts.

- **Brightness** (Default: 1, Range: 0 or greater)
  Brightness of the result.

- **Color** (Default rgb: [1 1 1])
  Scales the color of the result.

- **Glow Brightness** (Default: 2, Range: 0 or greater)
  Brightness of the glow applied to the texture.

- **Glow Color** (Default rgb: [1 0.8 0.8])
  Color of the glow applied to the texture.

- **Glow Width** (Default: 1, Range: 0 or greater)
  The width of the glow applied to the texture.

- **Glow Width Red** (Default: 0.4, Range: 0 or greater)
  The relative red width of the glow.

- **Glow Width Grn** (Default: 0.6, Range: 0 or greater)
  The relative green width of the glow.

- **Glow Width Blue** (Default: 0.8, Range: 0 or greater)
  The relative blue width of the glow.

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

