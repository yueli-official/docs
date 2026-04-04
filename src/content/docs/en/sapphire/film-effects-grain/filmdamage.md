---
title: FilmDamage
---

## S_FilmDamage

Simulates damaged film with many options,
including dust, hairs, stains, scratches, defocusing, flicker,
and shake. Each option has a master control and a set of detailed
controls for adjusting the look of that type of damage.

In the Sapphire Stylize effects submenu.

![FilmDamage](../_static/FilmDamage.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

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


### Grain Parameters:

Grain Amp:
*Default:
*0.1,
*Range:
*0 to 2.Scales the amplitude of the film grain that is added to
the result. Set this to 0 to disable all grain.

Grain Amp Red:
*Default:
*0.9,
*Range:
*0 or greater.Scales the red grain amplitude.

Grain Amp Green:
*Default:
*1,
*Range:
*0 or greater.Scales the green grain amplitude.

Grain Amp Blue:
*Default:
*1.6,
*Range:
*0 or greater.Scales the blue grain amplitude. Note that grain is
added and subtracted from the image, so for example, increasing
Grain Amp Blue will amplify both the blue and yellow speckles.

Grain Amp Darks:
*Default:
*0.2,
*Range:
*0 to 2.The relative amount of grain applied to the darkest
regions of the image, per channel. This defaults to less than 1.0
because dark areas usually have less grain than midtones.

Grain Amp Brights:
*Default:
*0,
*Range:
*0 to 2.The relative amount of grain applied to the
brightest regions of the image, per channel. This defaults to zero
because bright areas usually have less grain than midtones. Note
that highly saturated colors can be affected by both Grain Amp Darks
and Grain Amp Brights because they are dark in some color channels
and bright in others.

Grain Blur:
*Default:
*0,
*Range:
*0 or greater.The grain is smoothed by this amount. Increase for
coarser grain.

Grain Blur Red:
*Default:
*1,
*Range:
*0 or greater.The relative blur amount for the red grain.

Grain Blur Green:
*Default:
*0.9,
*Range:
*0 or greater.The relative blur amount for the green grain.

Grain Blur Blue:
*Default:
*1.2,
*Range:
*0 or greater.The relative blur amount for the blue grain.

Grain Mono:
*Check-box, Default:
*off.When enabled, the same grain pattern is used for the
red, green, and blue channels. To make truly monochrome grain you
should also set Grain Amp Red/Green/Blue equal to each other, make
sure Midtone Pos Red/Green/Blue are equal, and if GrainBlur is
positive also set Grain Blur Red/Green/Blue equal

Grain Hold:
*Popup menu, Default: Frame
*.Indicates how often a new grain pattern should be
generated. You will probably only notice a difference between these
options if Grain Blur is positive to make the grain size larger than
one pixel.
*Field:
*holds the grain pattern for one field.*Frame:
*holds the grain pattern for one frame (2
fields).*3:2 Pulldown at 0:
*holds the grain in a 3:2 pulldown
pattern with the first pulldown frame at 0. These options are
appropriate if your clip was created at 24 fps but is now in 30 fps
pulldown form. They will not make sense if your clip is 24P. A 3:2
pulldown pattern repeats every 5 frames, so if frame 1:00:23 is the
first frame with field artifacts after three normal frames, then you
should specify 3 as the first pulldown frame.*3:2 Pulldown at 1:
*holds the grain in a 3:2 pulldown
pattern with the first pulldown frame at 1.*3:2 Pulldown at 2:
*holds the grain in a 3:2 pulldown
pattern with the first pulldown frame at 2.*3:2 Pulldown at 3:
*holds the grain in a 3:2 pulldown
pattern with the first pulldown frame at 3.*3:2 Pulldown at 4:
*holds the grain in a 3:2 pulldown
pattern with the first pulldown frame at 4.

### Color Correct Parameters:

Saturation:
*Default:
*1,
*Range:
*any.Scales the color saturation. Increase for more intense
colors. Set to 0 for monochrome.

Scale Lights:
*Default:
*1,
*Range:
*0 or greater.Scales the result by this gray value. Increase for
a brighter result.

Offset Darks:
*Default:
*0,
*Range:
*any.Adds this gray value to the darker regions of the source.
This can be negative to increase contrast.

Tint Lights:
*Default rgb:
*[1 1 1].Scales the result by this color, thus tinting the
lighter regions.

Tint Darks:
*Default rgb:
*[0 0 0].
Adds this color to the darker regions of the source.

### Stains Parameters:

Stain Density:
*Default:
*0.2,
*Range:
*0 or greater.The number of stains on each frame. A fractional value is treated as
the probability of a single stain appearing on any given frame.

Vary Stain Density:
*Default:
*0.2,
*Range:
*0 or greater.Amount to vary the stain density from frame to frame.

Stain Print:
*Default:
*1,
*Range:
*0 to 1.Relative density of stains on the print.

Stain Negative:
*Default:
*0,
*Range:
*0 to 1.Relative density of stains on the negative.

Stain Size:
*Default:
*1,
*Range:
*0 or greater.Scales the width and height of stains.

Vary Stain Size:
*Default:
*0.5,
*Range:
*0 or greater.Amount to vary the size from one stain to the next.

Stain Opacity:
*Default:
*0.5,
*Range:
*0 to 1.Scales the opacity of the stains.

Vary Stain Opacity:
*Default:
*0.5,
*Range:
*0 or greater.Amount to vary opacity from one stain to the next.

Vary Stain Brightness:
*Default:
*0,
*Range:
*0 or greater.Amount to vary brightness from one stain to the next.

Vary Stain Color:
*Default:
*0,
*Range:
*0 or greater.Amount of additional, random color variation for each stain. If
this parameter is greater than zero, stain colors can vary outside the
range defined by color1 and color2.

Stain Color1:
*Default rgb:
*[0 0 0].Beginning of the range of colors for stains.

Stain Color2:
*Default rgb:
*[0.25 0.125 0].
End of the range of colors for stains. Each stain will have a random
color between color1 and color2.

### Dust Parameters:

Dust Density:
*Default:
*30,
*Range:
*0 or greater.The average number of dust pieces on each frame. A fractional value is
treated as the probability of a single dust speck appearing on any
given frame.

Vary Dust Density:
*Default:
*0.2,
*Range:
*0 or greater.Amount to vary the dust density from frame to frame.

Dust On Print:
*Default:
*1,
*Range:
*0 to 1.Relative density of dust on the print.

Dust On Negative:
*Default:
*0,
*Range:
*0 to 1.Relative density of dust on the negative.

Dust Size:
*Default:
*1,
*Range:
*0 or greater.Scales the width and height of dust.

Vary Dust Size:
*Default:
*0.5,
*Range:
*0 or greater.Amount to vary the size from one piece of dust to the next.

Dust Opacity:
*Default:
*0.8,
*Range:
*0 to 1.Scales the opacity of the dust.

Vary Dust Opacity:
*Default:
*0.5,
*Range:
*0 or greater.Amount to vary opacity from one piece of dust to the next.

Vary Dust Brightness:
*Default:
*0,
*Range:
*0 or greater.Amount to vary brightness from one piece of dust to the next.

Vary Dust Color:
*Default:
*0,
*Range:
*0 or greater.Amount of additional, random color variation for each piece of dust. If
this parameter is greater than zero, dust colors can vary outside the
range defined by color1 and color2.

Dust Color1:
*Default rgb:
*[0 0 0].Beginning of the range of colors for dust.

Dust Color2:
*Default rgb:
*[0 0 0].
End of the range of colors for dust. Each piece of dust will have a random
color between color1 and color2.

### Hairs Parameters:

Hairs:
*Default:
*2,
*Range:
*0 or greater.Number of hairs stuck in the projector gate.

Hair Persistence:
*Default:
*3,
*Range:
*0.1 or greater.Controls the length of time that hairs persist,
and the frequency with which new hairs appear. Increase this value
for long-lived hairs, and decrease it to get new hairs more often.

Hair Wiggle Amp:
*Default:
*0.1,
*Range:
*0 or greater.Controls the amount of random movement and
stretching that each hair exhibits.

Hair Wiggle Freq:
*Default:
*1,
*Range:
*0 or greater.Controls the frequency of the hair wiggle.

Hair Opacity:
*Default:
*1,
*Range:
*0 to 1.Scales the opacity of the hairs.

Hair Size:
*Default:
*1,
*Range:
*0 or greater.Scales the width and height of the hairs.

Vary Hair Size:
*Default:
*1,
*Range:
*0 or greater.Amount to vary the size from one hair to the next.

Hair Color:
*Default rgb:
*[0 0 0].
The color of the hairs.

### Scratches Parameters:

Scratches:
*Integer, Default:
*5,
*Range:
*0 or greater.Controls the number of scratches on each frame, on
average.

Black Scratches:
*Default:
*1,
*Range:
*0 to 1.Number of black scratches, relative to the Scratches parameter value.

White Scratches:
*Default:
*0.1,
*Range:
*0 to 1.Number of white scratches, relative to the Scratches parameter value.

Black Scratch Length:
*Default:
*10,
*Range:
*0 or greater.The length of the black scratches in frames,
on average.

White Scratch Length:
*Default:
*2,
*Range:
*0 or greater.The length of the white scratches in frames,
on average.

Scratch Width:
*Default:
*0.15,
*Range:
*0 or greater.Width of the average scratch, in approximate
NTSC-sized pixels.

Vary Scratches Width:
*Default:
*1,
*Range:
*0 to 1.If this is 0 all the scratches will be the same
width. Increase to let each scratch have its own width.

Scratches Taper:
*Default:
*0.1,
*Range:
*0 to 1.Controls the pointiness of the ends of each
scratch. Larger value makes a longer taper on each end.

Scratch Opacity:
*Default:
*1,
*Range:
*0 to 1.Maximum opacity of the scratches. Setting
this to 0 will fade the scratches out.

Scratch Roughness:
*Default:
*1,
*Range:
*0 or greater.Amount to roughen the edges of each scratch to
simulate the random character of a real scratch.

Scratch Rough Freq:
*Default:
*150,
*Range:
*0.01 or greater.Sets the frequency of the roughness on the
scratch edges.

Gaps:
*Default:
*0.28,
*Range:
*0 to 1.Like real analog scratches, the dust particle
creating the scratch sometimes rolls around and the scratch 'skips'.
This controls how much that happens.

Gaps Freq:
*Default:
*120,
*Range:
*0 or greater.How often do the scratch gaps occur.

Scratch Area Center:
*Default:
*0,
*Range:
*any.The center coordinate of the area of the screen
covered by the scratches. 0 is in the middle of the screen, -1 is
the left edge, and 1 is the right edge.

Scratch Area Width:
*Default:
*1,
*Range:
*0 or greater.The width of the area of the screen covered by
scratches. 1 means the scratches cover the full screen area. To
get scratches only in one strip, adjust scratch area width smaller.

Weave Amount:
*Default:
*1,
*Range:
*0 or greater.How much does each scratch weave around on the screen,
on average. This is in frame-widths, so 1.0 will let a scratch
wander all over the screen. If set to zero, the scratches will all
be straight vertical.

Weave Frequency:
*Default:
*0.1,
*Range:
*0.01 or greater.
How fast do the scratches weave around on the
screen, in cycles per frame. Normally less than one.

### Shake Parameters:

Shake Amplitude:
*Default:
*0,
*Range:
*0 or greater.Amount of vertical shaking to add.

Shake Frequency:
*Default:
*1,
*Range:
*0 or greater.Scales the frequency of the shaking. Increase for faster shaking
with more frequent hops and changes in direction.

Shake Jumpiness:
*Default:
*1,
*Range:
*0 or greater.Amount of large-scale, jumpy shaking.

Shake Random:
*Default:
*0.1,
*Range:
*0 or greater.Amount of small-scale, random shaking.

Shake Always:
*Default:
*0.5,
*Range:
*0 to 1.Controls how often shaking occurs. If set to 1, the clip shakes constantly.
If set to 0, the clip never shakes. Values in between cause the clip to
shake some of the time, and to stay still at other times.

Interframe Border Height:
*Default:
*0.1,
*Range:
*0 or greater.Size of the black bar in between frames (the
unexposed part of the film).

Shake Time Offset:
*Default:
*0,
*Range:
*any.Offsets the shake pattern in time. Adjust this value to control the
exact time when shaking occurs.

Shake Motion Blur:
*Default:
*0.1,
*Range:
*0 or greater.
Blurs the result proportionally to the amount of shaking.

### Vignette Parameters:

Vignette Darkness:
*Default:
*0.1,
*Range:
*0 to 1.Vignetting is darkening of the image towards the
corners and sides of the image. This parameter controls how much
the outer corners of the screen should be darkened (vignetted). 0
gives no vignetting, 1 gives maximum darkening.

Vignette Radius:
*Default:
*1,
*Range:
*0 or greater.Distance from the center to apply the vignette.

Vignette Edge Softness:
*Default:
*0.5,
*Range:
*0 or greater.The width of the vignette's soft edge. Larger values
give softer, less visible edges.

Vignette Rel Height:
*Default:
*0.75,
*Range:
*0.1 or greater.
Controls the aspect ratio of the vignette
ellipse. This should normally be set to the aspect ratio of the
image, e.g. .75 for NTSC.

### Flicker Parameters:

Flicker:
*Default:
*0.2,
*Range:
*0 or greater.Scales the colors of the source clip by different
amounts over time for a flickering effect. The pattern of flickering
can be random, a periodic wave, or a combination of the two.

Flicker Rand Amp:
*Default:
*1,
*Range:
*0 or greater.The amplitude of random brightness flickering.

Flicker Rand Freq:
*Default:
*10,
*Range:
*0 or greater.The frequency of the random flickering. Increase for
more variation between frames. Decrease for slower flickering.

Flicker Wave Amp:
*Default:
*0,
*Range:
*0 or greater.The amplitude of periodic wave flickering.

Flicker Wave Freq:
*Default:
*5,
*Range:
*0 or greater.
The frequency of the wave flickering. Increase for
faster flickering, decrease for slower. This has no effect if Wave
Amp is 0.

### Defocus Parameters:

Defocus:
*Default:
*0,
*Range:
*0 or greater.Blurs the source clip by different amounts over time to
simulate focus problems in the projector. The pattern of defocus
can be random, a periodic wave, or a combination of the two.

Defocus Rand Amp:
*Default:
*1,
*Range:
*0 or greater.The amplitude of defocusing that changes randomly over time.

Defocus Rand Freq:
*Default:
*10,
*Range:
*0 or greater.Scales the frequency of the random
defocus. Increase for more variation between frames. Decrease for slower defocus changes over time.

Defocus Wave Amp:
*Default:
*0,
*Range:
*0 or greater.The amplitude of periodic wave defocus.

Defocus Wave Freq:
*Default:
*5,
*Range:
*0 or greater.
The frequency of the wave defocus. Increase for
more variation between frames.

### Other Parameters:

Seed:
*Default:
*0.123,
*Range:
*0 or greater.Used to initialize the random number generator. The actual
seed value is not significant, but different seeds give different
results and the same value should give a repeatable result.

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

Flip Stamps Vertically:
*Check-box, Default:
*off.Flip all stamps (hairs, scratches, dust, etc) vertically.

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
