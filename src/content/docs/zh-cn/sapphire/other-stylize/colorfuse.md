---
title: ColorFuse
---

## S_ColorFuse

ColorFuse allows up to three LUTs to be combined to create unique and stylized looks. Host-colorspace
and lut-colorspace parameters are provided to convert footage from the host colorspace into the colorspace used in
the three stylized LUTs in ColorFuse. ColorFuse most commonly uses sRGB for the internal LUT colorspace.

In the Sapphire Stylize effects submenu.
### Inputs:

- **Source**: The current layer. The clip to process.

- **Mask**: Defaults to None. If provided, the effect is only applied on regions of the source clip specified by the bright areas of this input. Pixels outside this mask are not affected, and do not contribute to the resulting affected pixels within it. This input can be affected using the Invert Mask, or Mask Use parameters.


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

- **Host Colorspace** (Popup menu, Default: sRGB)
  Colorspace that the footage should be converted from. A list of common LUTs for converting footage from the colorspace in the host application to to the colorspace that the LUT was designed to operate in. If the colorspace used in the host isn't available in this preset, S_OCIOTransform may be applied before S_ColorFuse to get a more comprehensive colorspace list.
  - **linear**: Use linear for the host colorspace
  - **sRGB**: Use sRGB for the host colorspace
  - **rec709**: Use rec709 for the host colorspace
  - **rec2020**: Use rec2020 for the host colorspace
  - **rec1886**: Use rec1886 for the host colorspace
  - **S-Log1**: Use S-Log1 for the host colorspace
  - **S-Log2**: Use S-Log2 for the host colorspace
  - **S-Log3**: Use S-Log3 for the host colorspace

- **Lut Colorspace** (Popup menu, Default: sRGB)
  Colorspace ColorFuse should operate in. A list of common LUTs for defining the colorspace the effect LUTs expect.
  - **linear**: Use linear for the LUT colorspace
  - **sRGB**: Use sRGB for the LUT colorspace
  - **rec709**: Use rec709 for the LUT colorspace
  - **rec2020**: Use rec2020 for the LUT colorspace
  - **rec1886**: Use rec1886 for the LUT colorspace
  - **S-Log1**: Use S-Log1 for the LUT colorspace
  - **S-Log2**: Use S-Log2 for the LUT colorspace
  - **S-Log3**: Use S-Log3 for the LUT colorspace

- **Choose Lut1** (Push-button)
  Displays a file dialog to select the first LUT.

- **Lut1 Strength** (Default: 0.5, Range: 0 to 1)
  Intensity of the first LUT when applied to the footage.

- **Choose Lut2** (Push-button)
  Displays a file dialog to select the second LUT.

- **Lut2 Strength** (Default: 1, Range: 0 to 1)
  Intensity of second LUT when applied to the footage.

- **Choose Lut3** (Push-button)
  Displays a file dialog to select the third LUT.

- **Lut3 Strength** (Default: 1, Range: 0 to 1)
  Intensity of third LUT when applied to the footage.

- **Mask Use** (Popup menu, Default: Luma)
  Determines how the Mask input channels are used to make a monochrome mask.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Mask** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

