---
title: TextureTiles
---

## S_TextureTiles

TextureTiles draws a repeating pattern of tiles.
The shapes can be hexagons, triangles, diamonds, stars, or variations on those,
depending on the Morph parameters.

In the Sapphire Render effects submenu.

![TextureTiles](../_static/TextureTiles.jpg)


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

- **Size** (Default: 0.5, Range: 0 or greater)
  The size of each tile, within its cell. Zero will give all color0, one will give all color1. This doesn't change the overall size of the pattern; use Frequency for that.

- **Frequency** (Default: 5, Range: 0.01 or greater)
  Spatial frequency of the tile pattern; increase for many smaller tiles, decrease for fewer large tiles. This parameter can be adjusted using the Frequency Widget.

- **Angle** (Default: 0, Range: any)
  Rotates the whole pattern around the center point. Use Shift to adjust the center of rotation.

- **Rel Width** (Default: 1, Range: 0.2 or greater)
  Squashes or stretches the pattern.

- **Rel Wid Pre Rot** (Default: 1, Range: 0.1 or greater)
  Squashes or stretches the pattern before rotating by Angle. Use this if you want to squash or stretch and have the whole squashed/stretched pattern rotate around the center. If Angle is zero, this has the same effect as Rel Width.

- **Shift** (X & Y, Default: [0 0], Range: any)
  Shift the whole pattern on the screen. Also sets the center point for rotation, Morph Radial, and Size Radial.

- **Morph Shapes** (Default: 0, Range: any)
  Changes the shapes of the tiles smoothly, from hexagons to triangles, diamonds, and stars.

- **Morph Speed** (Default: 0.5, Range: any)
  Automatically animates the shape morphing over time. A value of one means a complete morph cycle once per second.

- **Morph Grad Add** (Default: 0, Range: any)
  Change the shape morphing across the image, so the left side has one shape, and the right side another. See Morph Grad Angle to change the angle of this gradient.

- **Morph Grad Angle** (Default: 0, Range: any)
  Angle of the morph gradient. If Morph Grad Add is zero, this has no effect.

- **Morph Radial** (Default: 0, Range: any)
  Morph the shapes radially away from the center point; the shapes will be (for instance) hexagons in the center, smoothly becoming different toward the edges of the image. Morph Shapes and Morph Speed also interact with this parameter.

- **Size Grad Add** (Default: 0, Range: -10 to 10)
  Change the size of the shapes (like the Size parameter) differently across the image.

- **Size Grad Angle** (Default: 0, Range: any)
  Angle of the size gradient. If Size Grad Add is zero, this has no effect.

- **Size Radial** (Default: 0, Range: any)
  Change the size of the shapes (like the Size parameter) according to the distance from the center point. Increase to make the sizes smaller around the edges.

- **Edge Softness** (Default: 0.17, Range: 0 or greater)
  Softens the edges of each tile. If Softness Red/Green/Blue are not one, there will be some color fringing around the edges of the tiles when this is on.

- **Softness Red** (Default: 0, Range: 0 or greater)
  Relative softness of the red channel; see Edge Softness. To remove the color fringing around the edges of the tiles, set all the Softness Red/Green/Blue to one.

- **Softness Green** (Default: 1, Range: 0 or greater)
  Relative softness of the green channel; see Edge Softness. To remove the color fringing around the edges of the tiles, set all the Softness Red/Green/Blue to one.

- **Softness Blue** (Default: 2, Range: 0 or greater)
  Relative softness of the blue channel; see Edge Softness. To remove the color fringing around the edges of the tiles, set all the Softness Red/Green/Blue to one.

- **Invert** (Check-box, Default: off)
  Invert the whole pattern, swapping the dark and bright areas.

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

- **Show Frequency** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Frequency parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

