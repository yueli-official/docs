---
title: LightLeak
---

## S_LightLeak

Renders abstract patterns of color that simulate light leaking through
gaps in a camera body. The light leak consists of three distinct elements which can be
adjusted individually.

In the Sapphire Lighting effects submenu.

![LightLeak](../_static/LightLeak.jpg)


### Inputs:

- **Background**: The current layer. The clip to use as background.

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

- **Scale Lights** (Default: 1, Range: 0 or greater)
  Scales the light leak by this value. Increase for a brighter result.

- **Offset Darks** (Default: 0, Range: -8 to 2)
  Adds this gray value to the darker regions of the result. This can be negative to increase contrast.

- **Color** (Default rgb: [1 1 1])
  The overall color of the light leak.

- **Hue Shift** (Default: 0, Range: any)
  Shifts the hue of the light leak, in revolutions from red to green to blue to red.

- **Saturation** (Default: 1, Range: -2 to 8)
  Scales the color saturation of the light leak. Increase for more intense colors. Set to 0 for a monochrome light leak.

- **Gamma** (Default: 1, Range: 0.1 or greater)
  Increasing gamma brightens the light leak, and especially boosts the darker areas.

- **Speed** (Default: 1, Range: 0 or greater)
  Scales the speed of all elements.

- **Shift** (X & Y, Default: [0 0], Range: any)
  Shifts the position of all elements.

- **Flicker Amp** (Default: 0.2, Range: 0 or greater)
  The amount of random flickering in the light leak brightness.

- **Flicker Freq** (Default: 4, Range: 0 or greater)
  The frequency of the random flickering. Increase for more variation between frames. Decrease for slower flickering.

- **Random Motion** (Default: 0, Range: 0 or greater)
  The amount of random motion of each element.

- **Random Frequency** (Default: 4, Range: 0 or greater)
  The frequency of random motion. Increase for faster, more frenetic motion. Decrease for slower, smoother motion.

- **Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the background before combining with the lights. If 0, the result will contain only the light image over black.

- **Combine** (Popup menu, Default: Screen)
  Determines how the light leak is combined with the background image.
  - **Screen**: the light leak is blended with the background using a function that
helps prevent overly bright results
  - **Add**: the light leak is added to the background.
  - **Leaks Only**: the light leak is shown on its own, with no background

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the light leak. The maximum of the red, green, and blue light leak brightness is scaled by this value and combined with the Background Alpha at each pixel.

- **Glow Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the glow which is applied to the entire image after combining the light leak with the background.

- **Glow Width** (Default: 0.4, Range: 0 or greater)
  The width of the glow. Increase for a softer glow and decrease for a sharper, brighter glow.

- **Glow Threshold** (Default: 0.8, Range: 0 or greater)
  Parts of the image that are brighter than this value get glowed.


### Element1 Parameters:

Element1 Enable:
*Check-box, Default:
*on.Turns this element on and off.

Size1:
*Default:
*2,
*Range:
*0 or greater.Adjusts the size of this element. This parameter can be adjusted using the Size1 Widget.

Rel Height1:
*Default:
*2,
*Range:
*0 or greater.Scales the vertical dimension of this element, making it elliptical
instead of circular.

Brightness1:
*Default:
*0.2,
*Range:
*0 or greater.Scales the brightness of this element.

Speed1:
*Default:
*1,
*Range:
*0 or greater.The speed at which this element moves across the screen. Set to zero
to keep the element stationary at its center position. A faster speed means the element will
start and end farther from its center.

Angle1:
*Default:
*175,
*Range:
*any.The angle of this element's path across the screen. The element moves along
a line at this angle, passing through the center position at the midpoint of the clip. This parameter can be adjusted using the Angle1 Widget.

Center1:
*X & Y, Default:
*[0 0],
*Range:
*any.The center point of this element's motion. The element will reach this
location half way through the clip. This parameter can be adjusted using the Center1 Widget.

Outer Color1:
*Default rgb:
*[1 1 0.15].The color at the outer edge of this element.

Mid Color1:
*Default rgb:
*[1 0.7 0.3].The color at the midpoint of this element, between the Outer and
Center Colors. The exact location depends on the Midpoint parameter.

Center Color1:
*Default rgb:
*[1 0.1 0].The color at the center of this element.

Midpoint1:
*Default:
*0.5,
*Range:
*0 to 1.Moves the location of the Mid Color between the center and outer edge of the
element. Set to 0 to place the Mid Color at the center, or 1 to place it at the edge.

Softness1:
*Default:
*0.4,
*Range:
*0 or greater.Blurs the color gradient of this element. Increase for a smoother gradient, or
decrease for sharper bands of color.

Noise Amp1:
*Default:
*2,
*Range:
*0 or greater.The amount of noise applied to this element.

Noise Freq1:
*Default:
*1.5,
*Range:
*0 or greater.The frequency of the noise applied to this element. Increase
for smaller blobs, or decrease for larger ones.

Noise Freq Rel Y1:
*Default:
*1,
*Range:
*0 or greater.The relative vertical frequency of the noise pattern. Increase
to flatten the noise, or decrease to stretch it out vertically.

Noise Detail1:
*Default:
*0,
*Range:
*0 to 1.Controls the amount of fine detail in the
noise simulation. Decrease to get smoother noise,
increase for a more crunchy or grainy look.

Noise Boil Speed1:
*Default:
*1,
*Range:
*0 or greater.
Sets the speed of noise boiling or evolving as the element moves.
Set to 0 for a static noise pattern.

### Element2 Parameters:

Element2 Enable:
*Check-box, Default:
*off.Turns this element on and off.

Size2:
*Default:
*2,
*Range:
*0 or greater.Adjusts the size of this element. This parameter can be adjusted using the Size2 Widget.

Rel Height2:
*Default:
*2,
*Range:
*0 or greater.Scales the vertical dimension of this element, making it elliptical
instead of circular.

Brightness2:
*Default:
*1,
*Range:
*0 or greater.Scales the brightness of this element.

Speed2:
*Default:
*1,
*Range:
*0 or greater.The speed at which this element moves across the screen. Set to zero
to keep the element stationary at its center position. A faster speed means the element will
start and end farther from its center.

Angle2:
*Default:
*175,
*Range:
*any.The angle of this element's path across the screen. The element moves along
a line at this angle, passing through the center position at the midpoint of the clip. This parameter can be adjusted using the Angle2 Widget.

Center2:
*X & Y, Default:
*[-0.5 -0.5],
*Range:
*any.The center point of this element's motion. The element will reach this
location half way through the clip. This parameter can be adjusted using the Center2 Widget.

Outer Color2:
*Default rgb:
*[0.25 0.15 0.1].The color at the outer edge of this element.

Mid Color2:
*Default rgb:
*[0.7 0.12 0.22].The color at the midpoint of this element, between the Outer and
Center Colors. The exact location depends on the Midpoint parameter.

Center Color2:
*Default rgb:
*[1 1 1].The color at the center of this element.

Midpoint2:
*Default:
*0.5,
*Range:
*0 to 1.Moves the location of the Mid Color between the center and outer edge of the
element. Set to 0 to place the Mid Color at the center, or 1 to place it at the edge.

Softness2:
*Default:
*0.4,
*Range:
*0 or greater.Blurs the color gradient of this element. Increase for a smoother gradient, or
decrease for sharper bands of color.

Noise Amp2:
*Default:
*1,
*Range:
*0 or greater.The amount of noise applied to this element.

Noise Freq2:
*Default:
*3,
*Range:
*0 or greater.The frequency of the noise applied to this element. Increase
for smaller blobs, or decrease for larger ones.

Noise Freq Rel Y2:
*Default:
*1,
*Range:
*0 or greater.The relative vertical frequency of the noise pattern. Increase
to flatten the noise, or decrease to stretch it out vertically.

Noise Detail2:
*Default:
*0,
*Range:
*0 to 1.Controls the amount of fine detail in the
noise simulation. Decrease to get smoother noise,
increase for a more crunchy or grainy look.

Noise Boil Speed2:
*Default:
*0,
*Range:
*0 or greater.
Sets the speed of noise boiling or evolving as the element moves.
Set to 0 for a static noise pattern.

### Element3 Parameters:

Element3 Enable:
*Check-box, Default:
*off.Turns this element on and off.

Size3:
*Default:
*0.25,
*Range:
*0 or greater.Adjusts the size of this element. This parameter can be adjusted using the Size3 Widget.

Rel Height3:
*Default:
*1,
*Range:
*0 or greater.Scales the vertical dimension of this element, making it elliptical
instead of circular.

Brightness3:
*Default:
*1,
*Range:
*0 or greater.Scales the brightness of this element.

Speed3:
*Default:
*1,
*Range:
*0 or greater.The speed at which this element moves across the screen. Set to zero
to keep the element stationary at its center position. A faster speed means the element will
start and end farther from its center.

Angle3:
*Default:
*175,
*Range:
*any.The angle of this element's path across the screen. The element moves along
a line at this angle, passing through the center position at the midpoint of the clip. This parameter can be adjusted using the Angle3 Widget.

Center3:
*X & Y, Default:
*[0.5 0.5],
*Range:
*any.The center point of this element's motion. The element will reach this
location half way through the clip. This parameter can be adjusted using the Center3 Widget.

Outer Color3:
*Default rgb:
*[0.2 0.2 0].The color at the outer edge of this element.

Mid Color3:
*Default rgb:
*[0.55 0.4 0].The color at the midpoint of this element, between the Outer and
Center Colors. The exact location depends on the Midpoint parameter.

Center Color3:
*Default rgb:
*[1 0.7 0].The color at the center of this element.

Midpoint3:
*Default:
*0.5,
*Range:
*0 to 1.Moves the location of the Mid Color between the center and outer edge of the
element. Set to 0 to place the Mid Color at the center, or 1 to place it at the edge.

Softness3:
*Default:
*0.4,
*Range:
*0 or greater.Blurs the color gradient of this element. Increase for a smoother gradient, or
decrease for sharper bands of color.

Noise Amp3:
*Default:
*0.5,
*Range:
*0 or greater.The amount of noise applied to this element.

Noise Freq3:
*Default:
*1,
*Range:
*0 or greater.The frequency of the noise applied to this element. Increase
for smaller blobs, or decrease for larger ones.

Noise Freq Rel Y3:
*Default:
*1,
*Range:
*0 or greater.The relative vertical frequency of the noise pattern. Increase
to flatten the noise, or decrease to stretch it out vertically.

Noise Detail3:
*Default:
*0,
*Range:
*0 to 1.Controls the amount of fine detail in the
noise simulation. Decrease to get smoother noise,
increase for a more crunchy or grainy look.

Noise Boil Speed3:
*Default:
*0,
*Range:
*0 or greater.Sets the speed of noise boiling or evolving as the element moves.
Set to 0 for a static noise pattern.

Opacity:
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
