---
title: FilmEffect
---

## S_FilmEffect

Provides a physically accurate model of film exposure
and processing to make your video footage look like it was shot on
particular film stocks. It can remove field artifacts, perform
color correction for specific film types, add film grain, and apply
glow or soft focus effects. The color correction and grain can be
selectively disabled using the Scale CC and Grain Amp parameters.

In the Sapphire Stylize effects submenu.

![FilmEffect](../_static/FilmEffect.jpg)


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

- **Neg Film** (Popup menu, Default: Kodak 5245)
  Selects the negative film stock.
  - **None**: Ignore any effect of negative film. This is not
normally useful unless you also select None for the print film to
disable both.
  - **Kodak 5245**: Eastman EXR 50D, low speed, daylight
balanced, very fine grain.
  - **Kodak 5246**: Kodak VISION 250D, higher contrast, medium
speed, daylight balanced, fine grain.
  - **Kodak 5248**: Eastman EXR 100T, medium speed, tungsten
light balanced, very fine grain.
  - **Kodak 5274**: Kodak VISION 200T, medium speed, tungsten
light balanced, fine grain.
  - **Kodak 5277**: Kodak VISION 320T, lower contrast, medium
speed, tungsten light balanced, medium-fine grain.
  - **Kodak 5279**: Kodak VISION 500T, high speed, tungsten light
balanced, somewhat grainy.
  - **Kodak 5284**: Kodak VISION Expression 500T, lower contrast,
high speed, tungsten light balanced, medium grain.
  - **Kodak 5289**: Kodak VISION 800T, very fast, tungsten light
balanced, grainy.
  - **Kodak 5293**: Eastman EXR 200T, reduced
contrast, tungsten light balanced, medium grain.
  - **Kodak 5298**: Eastman EXR 500T, high speed, tungsten light
balanced, grainy.
  - **K SFX200T**: Special effects film, medium grain.
  - **Kodak 5217**: Kodak Vision2 200T, tungsten light
balanced, fine grain.
  - **Kodak 5218**: Kodak Vision2 500T, tungsten light
balanced, fine grain.

- **Print Film** (Popup menu, Default: Kodak 2383)
  Selects the print film stock.
  - **None**: Ignore any effect of the print film. This causes
the negative to be output directly. If the negative film is also
set to None, the color correction and grain are disabled.
  - **Kodak 2383**: Kodak VISION Color Print Film, rich blacks.
  - **Kodak 2393**: Kodak VISION Premier Color Print Film, rich
blacks, some grain.
  - **Kodak 2395**: Kodak VISION Color Teleprint Film, low
contrast.
  - **Kodak 5386**: Eastman EXR Color Print Film
(discontinued by Kodak, replaced by 2383).
  - **Kodak 5285 Rev**: Ektachrome 100D Reversal film,
daylight balanced, high contrast and grainy.
Note that the negative film is ignored when using reversal film.
  - **Kodak 7270 Rev**: Kodachrome 40 Movie Film,
tungsten balanced reversal film, high contrast and somewhat grainy.
Note that the negative film is ignored when using reversal film.

- **Blur Input** (Default: 0, Range: 0 or greater)
  The input is smoothed by this amount. This can be used to remove video noise or compression artifacts before processing.


### Color Correct Parameters:

Scale CC:
*Default:
*1,
*Range:
*0 to 5.Scales the amount of color correction performed due to
the film types, gamma values, and exposure values. Set to 0 to
disable color correction. If you increase this above 1.0 it
exaggerates the color correction, which normally increases the
contrast.

Input Gamma:
*Default:
*2.2,
*Range:
*0.1 to 10.The gamma that your original clip was shot for. For
video this is normally 2.2; for synthetic computer graphics it may
be less.

Output Gamma:
*Default:
*2.2,
*Range:
*0.1 to 10.The intended viewing gamma of the output.

Neg Exposure:
*Default:
*0,
*Range:
*any.Adjusts the simulated exposure of the negative film,
in stops. Increase for over-exposed and brighter.

Print Exposure:
*Default:
*0,
*Range:
*any.Adjusts the simulated exposure of the print film,
in stops. Increase for over-exposed and darker.

Print Lights Red:
*Default:
*25,
*Range:
*0 to 50.Adjusts the red exposure of the print film, in
printer light points. 1 light point is 1/12 stop. Increase to
over-expose red and give a more cyan result.

Print Lights Green:
*Default:
*25,
*Range:
*0 to 50.Adjusts the green exposure of the print film, in
printer light points. 1 light point is 1/12 stop. Increase to
over-expose green and give a more magenta result.

Print Lights Blue:
*Default:
*25,
*Range:
*0 to 50.Adjusts the blue exposure of the print film, in
printer light points. 1 light point is 1/12 stop. Increase to
over-expose blue and give a more yellow result.

Scale Brights:
*Default:
*1,
*Range:
*0 or greater.Scales the bright areas of the final result after
the other color correction, glow, and grain are applied. (This
parameter is not affected by Scale CC.)

Offset Darks:
*Default:
*0,
*Range:
*-8 to 2.
Adds this gray value to the darker regions of the
final result after the other color correction, glow, and grain are
applied. This can be negative to increase contrast. (This
parameter is not affected by Scale CC.)

### Glow Parameters:

Glow Brightness:
*Default:
*0,
*Range:
*0 or greater.If positive, the image is combined with a blurred
version of itself to give a glowing look. Increase for a brighter
glow.

Glow Soft Focus:
*Default:
*0,
*Range:
*0 to 1.If positive, the image is mixed with a blurred
version of itself to give a soft focus look. The effect of this
parameter is similar to Glow Brightness, but this does not brighten
the overall result. Increase this to mix in more of the blurred
version and less of the original. If this is 1 and Glow Brightness
is 0 you will get only the blurred version.

Glow Width:
*Default:
*0.224,
*Range:
*0 or greater.The width of the blur used by the glow and/or soft focus.

Glow Width Red:
*Default:
*1,
*Range:
*0 or greater.The relative glow width for the red channel.

Glow Width Green:
*Default:
*1,
*Range:
*0 or greater.The relative glow width for the green channel.

Glow Width Blue:
*Default:
*1,
*Range:
*0 or greater.
The relative glow width for the blue channel.

### Grain Parameters:

Grain Amp:
*Default:
*1,
*Range:
*0 or greater.Scales the amplitude of the film grain that is added to
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

Midtone Pos Red:
*Default:
*0.5,
*Range:
*0 to 1.The position of the midtones in the red channel
that will normally receive the maximum amount of grain. The red
grain amplitude is interpolated from Grain Amp Darks at black, up to
1.0 at this midtone position, then down to Grain Amp Brights at
white. This whole curve is then scaled by the Grain Amp Red
parameter.

Midtone Pos Green:
*Default:
*0.5,
*Range:
*0 to 1.The position of the midtones in the green channel
that will normally receive the maximum amount of grain. The green
grain amplitude is interpolated from Grain Amp Darks at black, up to
1.0 at this midtone position, then down to Grain Amp Brights at
white. This whole curve is then scaled by the Grain Amp Green
parameter.

Midtone Pos Blue:
*Default:
*0.5,
*Range:
*0 to 1.The position of the midtones in the blue channel
that will normally receive the maximum amount of grain. The blue
grain amplitude is interpolated from Grain Amp Darks at black, up to
1.0 at this midtone position, then down to Grain Amp Brights at
white. This whole curve is then scaled by the Grain Amp Blue
parameter.

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
fields).*3:2 Stutter at 0:
*holds the grain in a 3:2 pulldown
pattern with the first pulldown frame at 0. These options are
appropriate if your clip was created at 24 fps but is now in 30 fps
pulldown form. They will not make sense if your clip is 24P. A 3:2
pulldown pattern repeats every 5 frames, so if frame 1:00:23 is the
first frame with field artifacts after three normal frames, then you
should specify 3 as the first pulldown frame.*3:2 Stutter at 1:
*holds the grain in a 3:2 pulldown
pattern with the first pulldown frame at 1.*3:2 Stutter at 2:
*holds the grain in a 3:2 pulldown
pattern with the first pulldown frame at 2.*3:2 Stutter at 3:
*holds the grain in a 3:2 pulldown
pattern with the first pulldown frame at 3.*3:2 Stutter at 4:
*holds the grain in a 3:2 pulldown
pattern with the first pulldown frame at 4.

Grain Seed:
*Default:
*0.123,
*Range:
*0 or greater.
Initializes the random number generator for the grain
generation. The actual seed value is not significant, but different
seeds give different grain patterns and the same value should give a
repeatable pattern.

### Vignette Parameters:

Vignette Darkness:
*Default:
*0,
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

### Field Parameters:

Fields:
*Popup menu, Default: As Is
*.Allows removing field artifacts from the input clip.
This is useful if you want the clip to look like it was shot on
frames instead of fields. You can show a single field, merge the
two together, or simulate a 3:2 pulldown stutter pattern.
*As Is:
*leaves the fields unchanged.*Keep Lower Only:
*shows the lower field only, removes the upper field.*Keep Upper Only:
*shows the upper field only, removes the lower field.*Merge Fields:
*blends both fields together to remove interlacing artifacts.*3:2 Stutter at 0:
*Simulates a temporal stutter effect as if
the clip had been transferred from 24P to NSTC video using 3:2
pulldown, with the first pulldown frame at 0. If you are using this
option with non-zero Grain Blur, you may want to also set Grain Hold to
the corresponding value.*3:2 Stutter at 1:
*Simulates a 3:2 pulldown effect
with the first pulldown frame at 1.*3:2 Stutter at 2:
*Simulates a 3:2 pulldown effect
with the first pulldown frame at 2.*3:2 Stutter at 3:
*Simulates a 3:2 pulldown effect
with the first pulldown frame at 3.*3:2 Stutter at 4:
*Simulates a 3:2 pulldown effect
with the first pulldown frame at 4.

Field Dominance:
*Popup menu, Default: Lower First
*.Specifies which field should come first in time
when simulating 3:2 pulldown patterns. This is only used if a 3:2 stutter
option is selected in the Fields and/or the Grain Hold options.
*Lower First:
*the lower field is first in time.*Upper First:
*the upper field is first in time.

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
