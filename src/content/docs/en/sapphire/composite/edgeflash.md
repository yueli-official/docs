---
title: EdgeFlash
---

## S_EdgeFlash

Adds a glow from the Front clip onto the Back clip, and
vice versa, then composites the Front over the Back. This can be
used to make a composite look more natural with light flashing
between the layers as if exposed on film together.

In the Sapphire Composite effects submenu.

![EdgeFlash](../_static/EdgeFlash.jpg)


### Inputs:

- **Foreground**: The current layer. The clip to use as foreground.

- **Background**: Defaults to None. The clip to use as background.

- **Matte**: Defaults to None. The alpha channel of this input specifies the opacities of the Foreground input. If this input is not provided, the alpha channel of the Foreground input is used instead. This input can be affected by the Invert Matte or Matte Use parameters.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: Original)
  Select which processing method to use
  - **Original**: The original method
  - **LightWrap**: An improved method that handles some edge conditions better

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

- **Fg Flash Amp** (Default: 0.8, Range: 0 or greater)
  The amount of flashing from the Front onto the Back.

- **Bg Flash Amp** (Default: 0.8, Range: 0 or greater)
  The amount of flashing from the Front onto the Back.

- **Flash Width** (Default: 0.088, Range: 0 or greater)
  The width of the flashing. This parameter can be adjusted using the Flash Width Widget.

- **Fg Lights** (Default: 1, Range: any)
  Scales the Front input by this value. Increase for a brighter result

- **Fg Darks** (Default: 0, Range: any)
  Adds this gray value to the darker regions of the Front input. This can be negative to increase contrast.

- **Fg Saturation** (Default: 1, Range: 0 or greater)
  Scales the color saturation of the Front input. Increase for more intense colors. Set to 0 for monochrome.

- **Bg Lights** (Default: 1, Range: any)
  Scales the Back input by this value. Increase for a brighter result

- **Bg Darks** (Default: 0, Range: any)
  Adds this gray value to the darker regions of the Back input. This can be negative to increase contrast.

- **Bg Saturation** (Default: 1, Range: 0 or greater)
  Scales the color saturation of the Back input. Increase for more intense colors. Set to 0 for monochrome.

- **Output** (Popup menu, Default: Comp)
  Selects between different output options.
  - **Foreground**: outputs only the Front clip with flashing from the Back.
  - **Background**: outputs only the Back clip with flashing from the Front.
  - **Comp**: flashes both, composites the Front over the Back, and
outputs the result.

- **Subpixel Widths** (Check-box, Default: off)
  Enables flashing by subpixel amounts. Use this for smoother animation of the flash width.

- **Comp Premult** (Check-box, Default: on)
  Disable this if you have provided a separate Matte input and the Foreground pixel values have not been pre-multiplied by this Matte.

- **Matte Use** (Popup menu, Default: Alpha)
  Determines how the Matte input channels are used to make a monochrome matte.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Invert Matte** (Check-box, Default: off)
  If enabled, the black and white of the output matte are inverted.

- **Fg Flash Amp** (Default: 0.8, Range: 0 or greater)
  The amount of flashing from the Front onto the Back.

- **Bg Flash Amp** (Default: 0.8, Range: 0 or greater)
  The amount of flashing from the Front onto the Back.

- **Flash Width** (Default: 0.088, Range: 0 or greater)
  The width of the flashing. This parameter can be adjusted using the Flash Width Widget.

- **Fg Lights** (Default: 1, Range: any)
  Scales the Front input by this value. Increase for a brighter result

- **Fg Darks** (Default: 0, Range: any)
  Adds this gray value to the darker regions of the Front input. This can be negative to increase contrast.

- **Fg Saturation** (Default: 1, Range: 0 or greater)
  Scales the color saturation of the Front input. Increase for more intense colors. Set to 0 for monochrome.

- **Bg Lights** (Default: 1, Range: any)
  Scales the Back input by this value. Increase for a brighter result

- **Bg Darks** (Default: 0, Range: any)
  Adds this gray value to the darker regions of the Back input. This can be negative to increase contrast.

- **Bg Saturation** (Default: 1, Range: 0 or greater)
  Scales the color saturation of the Back input. Increase for more intense colors. Set to 0 for monochrome.

- **Output** (Popup menu, Default: Comp)
  Selects between different output options.
  - **Foreground**: outputs only the Front clip with flashing from the Back.
  - **Background**: outputs only the Back clip with flashing from the Front.
  - **Comp**: flashes both, composites the Front over the Back, and
outputs the result.

- **Subpixel Widths** (Check-box, Default: off)
  Enables flashing by subpixel amounts. Use this for smoother animation of the flash width.

- **Comp Premult** (Check-box, Default: on)
  Disable this if you have provided a separate Matte input and the Foreground pixel values have not been pre-multiplied by this Matte.

- **Matte Use** (Popup menu, Default: Alpha)
  Determines how the Matte input channels are used to make a monochrome matte.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Invert Matte** (Check-box, Default: off)
  If enabled, the black and white of the output matte are inverted.

- **Shrink- Grow+** (Default: 0, Range: any)
  Amount to grow the matte edges in approximate pixels, or shrink if negative.

- **Edge Softness** (Default: 1, Range: 0.01 or greater)
  The resulting softness of the edges.

- **Post Blur** (Default: 0, Range: 0 or greater)
  If positive, the result is blurred by this amount. This is an alternative method for softening the edges.

- **Noise Amplitude** (Default: 0, Range: 0 or greater)
  The amount of noise texture to add to the edges.

- **Noise Width** (Default: 0.0224, Range: 0 or greater)
  The width of the area at the matte edges where the noise is included. This has no effect unless Noise Amplitude is positive

- **Frequency** (Default: 100, Range: 0.1 or greater)
  The frequency of the noise. Increase for finer grain noise, decrease for coarser noise. This has no effect unless Noise Amplitude is positive.

- **Octaves** (Integer, Default: 1, Range: 1 to 10)
  The number of summed layers of noise. Each octave is twice the frequency and half the magnitude of the previous. This has no effect unless Noise Amplitude is positive.

- **Seed** (Default: 0.23, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Noise Shift** (X & Y, Default: [0 0], Range: any)
  The horizontal and vertical translation of the noise texture.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Flash Width** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Flash Width parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

