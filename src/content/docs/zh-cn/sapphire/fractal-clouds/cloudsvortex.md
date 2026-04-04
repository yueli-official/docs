---
title: CloudsVortex
---

## S_CloudsVortex

Generates a procedural noise texture twisting into a vortex. The Vortex
Speed parameter causes the amount of vortex rotation to automatically
animate over time.

In the Sapphire Render effects submenu.

![CloudsVortex](../_static/CloudsVortex.jpg)


### Inputs:

- **Background**: The current layer. The clip to combine the texture image with. This may be ignored if the Combine option is set to Texture Only.

- **Mask**: Defaults to None. If provided, the amplitude of warping is scaled by the values of this input clip. Gray values internally scale the warping amplitude rather than simply cross-fading between the effect and the original source to allow more continuous results at the mask edges and more detailed control over the warping amounts. This input can be affected using the Blur Mask, Invert Mask, or Mask Use parameters.


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
  The spatial frequency of the texture. Increase to zoom out, decrease to zoom in.

- **Frequency Rel X** (Default: 1, Range: 0.01 or greater)
  The relative horizontal frequency of the texture. Increase to stretch it vertically or decrease to stretch it horizontally.

- **Octaves** (Integer, Default: 6, Range: 1 to 10)
  The number of summed layers of noise. Each octave is twice the frequency and half the amplitude of the previous. A single octave gives a smooth texture. Adding octaves makes the result approach a fractal (1/f) noise texture.

- **Seed** (Default: 0.234, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Center** (X & Y, Default: [0 0], Range: any)
  The center of the vortex, in screen coordinates relative to the center of the frame. This parameter can be adjusted using the Center Widget.

- **Z Dist** (Default: 1, Range: 0.001 or greater)
  Scales the 'distance' of the image. Values greater than 1.0 move it farther away and make it smaller. Values less then 1.0 move the image closer and enlarge it.

- **Latitude** (Default: 30, Range: -80 to 80)
  Positive latitude tilts the image down and negative tilts it up. Keep latitude in the range of around -35 to 35 degrees to avoid aliasing towards the horizon.

- **Boiling Mode** (Check-box, Default: off)
  When enabled, the clouds will boil, or evolve, over time. This mode takes slightly more computation but usually looks better.

- **Boil Details** (Default: 0.55, Range: 0 to 1)
  Increases or decreases the amount of fine detail in the clouds. Decrease to get a smoother look, increase to get a more high-frequency, noisy look. Only used when Boiling Mode is enabled.

- **Boil Speed** (Default: 1, Range: any)
  Sets the speed of the cloud boiling. Zero gives no boiling at all. Only used when Boiling Mode is enabled.

- **Vortex Start** (Default: 72, Range: any)
  The amount of vortex rotation, in approximate degrees at the edge of the frame.

- **Vortex Speed** (Default: 30, Range: any)
  The speed of the vortex rotation, in approximate degrees per second at the edge of the frame. If non-zero, the vortexing is automatically animated at this rate.

- **Angle Offset** (Default: 0, Range: any)
  If non-zero, a rotation is combined with the vortex. Make negative to rotate the inner and outer regions in opposite directions.

- **Inner Radius** (Default: 0.04, Range: 0 or greater)
  The radius from the center at which the vortexing is phased in. This can be used to reduce excessive distortion and aliasing at the very center of the vortex.

- **Brightness1** (Default: 1, Range: 0 or greater)
  Scales the brightness of Color1. Increase for more contrast.

- **Color1** (Default rgb: [1 1 1])
  The color of the 'brighter' parts of the texture. The colors of the result are determined by an interpolation between Color0 and Color1.

- **Color0** (Default rgb: [0 0 0])
  The color of the 'darker' parts of the texture.

- **Offset0** (Default: 0, Range: any)
  Adds this value to color0. Decrease to a negative value for more contrast.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  The background brightness is scaled by this value before being combined with the texture.

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

- **Show Center** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

