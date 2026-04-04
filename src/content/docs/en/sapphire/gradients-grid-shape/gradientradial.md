---
title: GradientRadial
---

## S_GradientRadial

Makes a smooth radial color gradient in an ellipse shape, given Center, Inner
Radius, and Outer Radius parameters, and optionally combines the gradient
with a background clip. Increase Add Noise to reduce banding artifacts in
the gradient due to color quantization.

In the Sapphire Render effects submenu.

![GradientRadial](../_static/GradientRadial.jpg)


### Inputs:

- **Background**: The current layer. The clip to combine the gradient with.

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

- **Center** (X & Y, Default: [0 0], Range: any)
  The center location of the ellipse shape.

- **Inner Radius** (Default: 0.14, Range: 0 or greater)
  Distance from the center that the gradient starts. This parameter can be adjusted using the Inner Radius Widget.

- **Outer Radius** (Default: 1, Range: 0 or greater)
  Distance from the center that the gradient ends. This parameter can be adjusted using the Outer Radius Widget.

- **Rel Height** (Default: 0.75, Range: 0.1 or greater)
  The relative vertical size of the ellipse shape. Increase for a taller ellipse, decrease for a wider one.

- **Rel Width** (Default: 1, Range: 0.1 or greater)
  The relative horizontal size of the ellipse shape. Increase for a wider ellipse, decrease for a taller one.

- **Rotate** (Default: 0, Range: any)
  Rotation in degrees of the ellipse. Note that rotation will have no effect when Rel Width and Rel Height are equal and the shape is a perfect circle. This parameter can be adjusted using the Rotate Widget.

- **Inner Color** (Default rgb: [1 1 1])
  The gradient color at the Inner Radius.

- **Outer Color** (Default rgb: [0 0 0])
  The gradient color at the Outer Radius.

- **Swap Colors** (Check-box, Default: off)
  If checked, reverses the direction of the gradient by swapping the Inner and Outer colors.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the gradient image (both the Inner Color and Outer Color).

- **Add Noise** (Default: 0, Range: 0 or greater)
  If positive, this amount of noise is added to the gradient. This can create a grainy effect and eliminate banding in the gradient due to quantization. Set this to 1.0 to enable effective debanding for 8 bit results.

- **Smooth Curve** (Default: 0, Range: 0 to 1)
  If zero, a linear interpolation is used across the screen between the Start and End Color. Increase this value to use a smoother 'S' shaped curve for interpolation which can reduce the visual perception of the gradient's Start and End locations.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the background before combining it with the gradient.

- **Combine** (Popup menu, Default: Grad Only)
  Determines how the gradient is combined with the background.
  - **Grad Only**: gives the gradient image alone with no background.
  - **Mult**: the background is multiplied by the gradient.
  - **Add**: the background is added to the gradient.
  - **Screen**: the background is blended with the gradient using a
screen operation.
  - **Difference**: the result is the difference between the background
and gradient.
  - **Overlay**: combines gradient and background using an overlay function.

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

- **Show Outer Radius** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Inner Radius** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Rotate** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

