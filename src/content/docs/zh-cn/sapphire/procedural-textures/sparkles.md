---
title: Sparkles
---

## S_Sparkles

Generates a field of sparkling glint effects. Adjust the Frequency, Density,
and Size parameters for different types of sparkling patterns. Use the
Matte input to only generate sparkles in specified areas.

In the Sapphire Render effects submenu.

![Sparkles](../_static/Sparkles.jpg)


### Inputs:

- **Background**: The current layer. The clip to combine the sparkles with.

- **Matte**: Defaults to None. If provided, the sparkle colors are scaled by this input. A monochrome matte can be used to choose the areas that will generate sparkles. A color matte can be used to selectively adjust the sparkle colors in different regions. The matte is applied before the sparkles are generated so it will not clip the resulting glint rays.


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

- **Frequency** (Default: 25, Range: 0.01 or greater)
  The frequency of the sparkles. Increase to zoom out, decrease to zoom in.

- **Density** (Default: 0.65, Range: 0 to 1)
  Increase to add more sparkles.

- **Seed** (Default: 0.23, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of all the sparkles.

- **Color** (Default rgb: [1 1 1])
  Scales the color of all the sparkles.

- **Brightness X** (Default: 1, Range: 0 or greater)
  Scales the brightness of the horizontal glint rays.

- **Brightness Y** (Default: 1, Range: 0 or greater)
  Scales the brightness of the vertical glint rays.

- **Brightness Diag1** (Default: 1, Range: 0 or greater)
  Scales the brightness of the diagonal rays from top right to bottom left.

- **Brightness Diag2** (Default: 1, Range: 0 or greater)
  Scales the brightness of the diagonal rays from top left to bottom right.

- **Size** (Default: 1, Range: 0 or greater)
  Scales the length of all the glint rays. This and all the size parameters can be adjusted using the Size Widget.

- **Size X** (Default: 1, Range: 0 or greater)
  Scales the length of the horizontal glint rays.

- **Size Y** (Default: 1, Range: 0 or greater)
  Scales the length of the vertical glint rays.

- **Size Diag1** (Default: 0.5, Range: 0 or greater)
  Scales the length of the diagonal rays from top left to bottom right.

- **Size Diag2** (Default: 0.5, Range: 0 or greater)
  Scales the length of the diagonal rays from top right to bottom left.

- **Size Red** (Default: 0.6, Range: 0 or greater)
  Scales the length of the red component of the rays. If the red, green, and blue sizes are equal the sparkles will be monochrome.

- **Size Green** (Default: 0.8, Range: 0 or greater)
  Scales the length of the green component of the rays.

- **Size Blue** (Default: 1, Range: 0 or greater)
  Scales the length of the blue component of the rays.

- **Shift Start** (X & Y, Default: [0 0], Range: any)
  Translation offset of the result.

- **Shift Speed** (X & Y, Default: [0 0], Range: any)
  Translation speed of the result. If non-zero, the result is automatically animated to shift at this rate. The result of animated Speed values may not be intuitive, so for variable speed motion it is usually best to set this to 0 and animate the Shift Start values instead.

- **Sparkle Speed** (X & Y, Default: [0.1 0], Range: any)
  If non-zero, the sparkles automatically twinkle on and off at this rate.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the sparkles. The maximum of the red, green, and blue sparkle brightness is scaled by this value and combined with the background Alpha at each pixel.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the background before combining with the Sparkles. If 0, the result will contain only the sparkles image over black.

- **Smooth Anim** (Check-box, Default: off)
  Enable for more steady animation, especially at high values of Frequency.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Matte Use** (Popup menu, Default: RGB)
  Determines how the Matte input channels are used to make a monochrome matte.
  - **RGB**: The red, green, and blue channels are used.
  - **Alpha**: only the Alpha channel is used.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Swap Diagonals** (Check-box, Default: off)
  Flips sparkles vertically if needed to achieve a consistent look.

- **Show Size** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the size parameters.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

