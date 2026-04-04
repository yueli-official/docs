---
title: GradientMulti
---

## S_GradientMulti

Generates a smooth multi-color gradient across the screen using multiple control
points, and optionally combines the gradient with a background clip.

In the Sapphire Render effects submenu.

![GradientMulti](../_static/GradientMulti.jpg)


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

- **Softness** (Default: 1, Range: 0.01 or greater)
  The softness of the edges between color regions. Increasing this parameter will create a smoother gradient, while decreasing it will create sharper edges and more well-defined colors.

- **Softness Falloff** (Default: 0, Range: 0 or greater)
  Reduces the softness as the distance from the control points increases. Higher values will create more well-defined color regions near the edges of the image, while lower values will cause the colors to blend together more.


### Point 1 Parameters:

Point 1 Enable:
*Check-box, Default:
*on.Turns on or off the first control point.

Color 1:
*Default rgb:
*[1 0 0].The color at Point 1.

Point 1:
*X & Y, Default:
*[-0.972 -0.719],
*Range:
*any.First control point. This parameter can be adjusted using the Point 1 Widget.

Softness 1:
*Default:
*1,
*Range:
*0.1 or greater.The relative softness of color 1.

Size 1:
*Default:
*1,
*Range:
*0.1 or greater.
Scales the size of the color centered at Point 1.

### Point 2 Parameters:

Point 2 Enable:
*Check-box, Default:
*on.Turns on or off the second control point.

Color 2:
*Default rgb:
*[0 1 0].The color at Point 2.

Point 2:
*X & Y, Default:
*[-0.972 0.701],
*Range:
*any.Second control point. This parameter can be adjusted using the Point 2 Widget.

Softness 2:
*Default:
*1,
*Range:
*0.1 or greater.The relative softness of color 2.

Size 2:
*Default:
*1,
*Range:
*0.1 or greater.
Scales the size of the color centered at Point 2.

### Point 3 Parameters:

Point 3 Enable:
*Check-box, Default:
*on.Turns on or off the third control point.

Color 3:
*Default rgb:
*[0 0 1].The color at Point 3.

Point 3:
*X & Y, Default:
*[0.972 0.701],
*Range:
*any.Third control point. This parameter can be adjusted using the Point 3 Widget.

Softness 3:
*Default:
*1,
*Range:
*0.1 or greater.The relative softness of color 3.

Size 3:
*Default:
*1,
*Range:
*0.1 or greater.
Scales the size of the color centered at Point 3.

### Point 4 Parameters:

Point 4 Enable:
*Check-box, Default:
*off.Turns on or off the fourth control point.

Color 4:
*Default rgb:
*[1 1 1].The color at Point 4.

Point 4:
*X & Y, Default:
*[0.972 -0.719],
*Range:
*any.Fourth control point. This parameter can be adjusted using the Point 4 Widget.

Softness 4:
*Default:
*1,
*Range:
*0.1 or greater.The relative softness of color 4.

Size 4:
*Default:
*1,
*Range:
*0.1 or greater.
Scales the size of the color centered at Point 4.

### Point 5 Parameters:

Point 5 Enable:
*Check-box, Default:
*off.Turns on or off the fifth control point.

Color 5:
*Default rgb:
*[1 1 0].The color at Point 5.

Point 5:
*X & Y, Default:
*[-0.167 0],
*Range:
*any.Fifth control point. This parameter can be adjusted using the Point 5 Widget.

Softness 5:
*Default:
*1,
*Range:
*0.1 or greater.The relative softness of color 5.

Size 5:
*Default:
*1,
*Range:
*0.1 or greater.
Scales the size of the color centered at Point 5.

### Point 6 Parameters:

Point 6 Enable:
*Check-box, Default:
*off.Turns on or off the sixth control point.

Color 6:
*Default rgb:
*[0 1 1].The color at Point 6.

Point 6:
*X & Y, Default:
*[0.167 0],
*Range:
*any.Sixth control point. This parameter can be adjusted using the Point 6 Widget.

Softness 6:
*Default:
*1,
*Range:
*0.1 or greater.The relative softness of color 6.

Size 6:
*Default:
*1,
*Range:
*0.1 or greater.Scales the size of the color centered at Point 6.

Combine:
*Popup menu, Default: Grad Only
*.Determines how the gradient is combined with the background.
*Grad Only:
*gives the gradient image alone with no background.*Mult:
*the background is multiplied by the gradient.*Add:
*the background is added to the gradient.*Screen:
*the background is blended with the gradient using a
screen operation.*Difference:
*the result is the difference between the background
and gradient.*Overlay:
*combines gradient and background using an overlay function.

Bg Brightness:
*Default:
*1,
*Range:
*0 or greater.Scales the brightness of the background before
combining it with the gradient.

Input Opacity:
*Popup menu, Default: Normal
*.Determines the method used for dealing with
opacity/transparency.
*All Opaque:
*Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).*Normal:
*Process opacity normally.*As Premult:
*Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

Output Opacity:
*Popup menu, Default: Copy From Input
*.Determines the opacity/transparency of the result.
This effect does not process the opacity (alpha channel) of its input
but it can either copy the opacity from the input, or output a fully
opaque result.
*All Opaque:
*Makes the result fully opaque with no
transparency.*Copy From Input:
*Copies the opacity/transparency from
the current layer given to this effect.

Mask Use:
*Popup menu, Default: Luma
*.Determines how the Mask input channels are used to make a
monochrome mask.
*Luma:
*the luminance of the RGB channels is used.*Alpha:
*only the Alpha channel is used.

Blur Mask:
*Default:
*0.05,
*Range:
*0 or greater.Blurs the Matte input by this amount before using. This
can provide a smoother transition between the matted and unmatted
areas. It has no effect unless the Matte input is provided.

Invert Mask:
*Check-box, Default:
*off.
If on, inverts the Matte input so the effect is applied
to areas where the Matte is black instead of white. This has no effect
unless the Matte input is provided.
