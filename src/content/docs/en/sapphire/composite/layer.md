---
title: Layer
---

## S_Layer

Layers the Foreground image over the Background using one of a variety of
blending operations. The colors of each input can also be adjusted using
the lights, darks, and saturation parameters.

In the Sapphire Composite effects submenu.

![Layer](../_static/Layer.jpg)


### Inputs:

- **Foreground**: The current layer. The clip to use as foreground.

- **Background**: Defaults to None. The clip to use as background.

- **Matte**: Defaults to None. Specifies the opacities of the Foreground clip. If this input is not provided the Foreground alpha is used instead. These values are scaled by the Fg Opacity parameter before being used.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: Normal)
  Determines which blending method is used to combine the foreground and background pixel colors.
  - **Normal**: a normal composite. This will just give
the foreground as the result unless the Opacity is below 1.0
or an alpha channel is given.
  - **Dissolve**: randomly replaces background pixels with
foreground. The opacity determines the probability, so the
foreground is more likely to replace the background for higher
values of Opacity.
  - **Multiply**: this can be used as an 'intersection' operation
on matte images. White is the identity for Multiply, where one
image contains white the other is not affected, so the result only
contains white where both inputs are white.
  - **Screen**: this can be useful for combining the bright areas
of two clips. It can also be used as a 'union' operation on matte
images. Black is the identity for Screen, where one image contains
black the other is not affected, so the result is white where either
of the input images is white.
  - **Overlay**: combines foreground and background using an
overlay function.
  - **Soft Light**: darkens or lightens the background depending on
the foreground.
  - **Hard Light**: similar to overlay but with foreground and
background swapped.
  - **Color Dodge**: brightens the background depending on the foreground.
  - **Color Burn**: darkens the background depending on the foreground.
  - **Darken**: the minimum of foreground and background. This can
also be used as an 'intersection' operation with slightly different
results than Multiply.
  - **Lighten**: the maximum of foreground and background. This
can also be used as a 'union' operation with slightly different
results than Screen.
  - **Add**: adds the foreground to the background.
  - **Subtract**: subtracts the foreground from the background.
  - **Difference**: similar to Subtract but the absolute value of
the result is used, which tends to give more resulting colors in
bounds. This can be used to select the regions of two matte
images where one or the other is white, but not both.
  - **Exclusion**: similar to Difference but with smoother results.
  - **Hue**: combines the hue of the foreground with the saturation
and luminance of the background.
  - **Saturation**: combines the saturation of the foreground with
the hue and luminance of the background.
  - **Chroma**: combines the hue and saturation of the foreground
with the luminance of the background.
  - **Luminance**: combines the luminance of the foreground with
the hue and saturation of the background.
  - **Linear Dodge**: adds foreground and background and clamps the
result at white.
  - **Linear Burn**: adds foreground and background but offsets to
make the result darker. Similar to multiply in that combining with
white gives no change and combining with black gives black.
  - **Linear Light**: performs a linear burn or linear dodge
depending on if the foreground is more or less than 50 percent gray.
  - **Vivid Light**: performs a color burn or color dodge
depending on if the foreground is more or less than 50 percent gray.
  - **Pin Light**: performs a lighten or darken depending on if the
foreground is more or less than 50 percent gray.

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

- **Swap Inputs** (Check-box, Default: off)
  If enabled, effectively swaps the Background and Foreground inputs, and can be helpful for non-commutative operations like subtract. Note that this also causes parameters labeled 'Front' to affect the 'Back' input instead, and vice versa.

- **Fg Opacity** (Default: 1, Range: 0 to 1)
  Scales the opacity of the effect. When this is decreased the result approaches the background. At zero, the result will equal the background.

- **Fg Lights** (Default: 1, Range: any)
  Scales the Foreground before performing the effect.

- **Fg Darks** (Default: 0, Range: any)
  Offsets the darker regions of the Foreground before performing the effect. This can be negative to increase contrast.

- **Fg Saturation** (Default: 1, Range: any)
  Scales the color saturation of the Foreground before performing the effect. 0.0 makes it monochromatic, 1.0 has no effect.

- **Fg Hue Shift** (Default: 0, Range: any)
  Shifts the hue of the colors in the Front clip, in revolutions from red to green to blue to red.

- **Fg Blur** (Default: 0, Range: 0 or greater)
  Amount to blur the foreground.

- **Bg Lights** (Default: 1, Range: any)
  Scales the Background before performing the effect.

- **Bg Darks** (Default: 0, Range: any)
  Offsets the darker regions of the Background before performing the effect. This can be negative to increase contrast.

- **Bg Saturation** (Default: 1, Range: any)
  Scales the color saturation of the Background before performing the effect. 0.0 makes it monochromatic, 1.0 has no effect.

- **Bg Hue Shift** (Default: 0, Range: any)
  Shifts the hue of the colors in the Back clip, in revolutions from red to green to blue to red.

- **Bg Blur** (Default: 0, Range: 0 or greater)
  Amount to blur the background.

- **Result Lights** (Default: 1, Range: any)
  Scales the result after performing the effect.

- **Result Darks** (Default: 0, Range: any)
  Offsets the darker regions of the result after performing the effect. This can be negative to increase contrast.

- **Result Saturation** (Default: 1, Range: any)
  Scales the color saturation of the result after performing the effect. 0.0 makes it monochromatic, 1.0 has no effect.

- **Result Hue Shift** (Default: 0, Range: any)
  Shifts the hue of the colors in the result, in revolutions from red to green to blue to red.

- **Matte Use** (Popup menu, Default: Alpha)
  Determines how the Matte input channels are used to make a monochrome matte.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.
  - **All Opaque**: an alpha of 1.0 is used as if the Matte were fully opaque.

- **Blur Subpixel** (Check-box, Default: on)
  Enables blurring by subpixel amounts. Use this for smoother animation of the Blur Front and Blur Back parameters.

- **Soft Borders** (Check-box, Default: off)
  If enabled, transparent borders are added to the input image before processing. This allows the result to include soft edges beyond the original image size. When off, the effect only occurs within the frame and the result will retain an edge at the borders.

- **Comp Premult** (Check-box, Default: on)
  Disable this if you have provided a separate Matte input and the Foreground pixel values have not been pre-multiplied by this Matte.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

