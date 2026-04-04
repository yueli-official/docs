---
title: FilmRoll
---

## S_FilmRoll

Transitions between two clips by rolling one off screen vertically
while rolling the other on, while applying various film damage effects such as shaking, stains,
scratches, and flicker.

In the Sapphire Transitions effects submenu.

![FilmRoll](../_static/FilmRoll.jpg)


### Inputs:

- **Foreground**: The current layer. Starts the transition with this clip.

- **Background**: Defaults to None. Ends the transition with this clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Transition Dir** (Popup menu, Default: Wipe Off to Bg)
  Selects the direction of the transition.
  - **Wipe Off to Bg**: transitions from the current layer to the Background.
  - **Wipe On from Bg**: transitions from the Background to the current layer.

- **Auto Trans** (Popup YES-NO, Default: No)
  If enabled, a transition is performed automatically between the first and last frames of the layer. If this is off, the transition is performed manually by animating the Film Percent parameter.

- **Amount** (Default: 0, Range: 0 to 1)
  Auto Trans must be disabled for this parameter to be used. It determines the transition ratio between the From and To inputs, and would normally be animated from 0 to 100 to perform a complete transition. The curve controlling this parameter can be adjusted for more detailed control over the timing of the wipe.

- **Slow In** (Default: 0.5, Range: 0 or greater)
  If positive, causes the transition to start more gradually.

- **Slow Out** (Default: 0.5, Range: 0 or greater)
  If positive, causes the transition to end more gradually.

- **Roll Speed** (Integer, Default: 1, Range: any)
  The amount of vertical rolling, in screen heights. The clips will move this distance over the full course of the transition, ending with the To clip in its normal position.

- **Motion Blur** (Default: 0.5, Range: 0 or greater)
  Blurs the result proportionally to the amount of shaking.

- **Border Height** (Default: 0.1, Range: 0 or greater)
  The height of the border that appears between the From and To clips as they are rolling.

- **Glow Brightness** (Default: 0.5, Range: 0 or greater)
  Adjusts the peak amount of glow. Glow will automatically fall off to zero at the beginning and end of the effect, to provide a smooth transition.

- **Glow Width** (Default: 0.224, Range: 0 or greater)
  The width of the glowing border.

- **Damage Amount** (Default: 2, Range: 0 or greater)
  Adjusts the peak amount of all damage effects. Increase for a more damaged look, or decrease for a cleaner look. The damage will automatically fall off at the beginning and end of the transition. Individual damage types can also be adjusted with their respective parameters, such as Stain Density, Hairs, Scratches, etc.


### Stains Parameters:

Stain Density:
*Default:
*2,
*Range:
*0 to 500.The number of stains on each frame. A fractional value is treated as
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
*60,
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

Scratch Roughness Freq:
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
*-2 or greater.The center coordinate of the area of the screen
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
*0.2,
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

Shake Time Offset:
*Default:
*0,
*Range:
*any.
Offsets the shake pattern in time. Adjust this value to control the
exact time when shaking occurs.

### Vignette Parameters:

Vignette Darkness:
*Default:
*0.5,
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
*1,
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
*0.5,
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

Flip Stamps Vertically:
*Check-box, Default:
*off.
Flip all stamps (hairs, scratches, dust, etc) vertically.
