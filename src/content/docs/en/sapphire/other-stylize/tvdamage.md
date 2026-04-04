---
title: TVDamage
---

## S_TVDamage

Simulates a TV with transmission and reception problems,
VCR issues, and TV hardware difficulties. Simulates static, interference,
ghosting, horizontal and vertical hold, hum bars, color stripes, visible scanlines,
VCR fast-forward, dropouts, vignetting, orthicon, fisheye, and turn-off.

In the Sapphire Stylize effects submenu.

![TVDamage](../_static/TVDamage.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Mask**: Defaults to None. Interpolate between the result and the Source input. White areas use the result of the effect. Black areas use the Source clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: TVDamage Color)
  Type of TV to simulate: color or black & white.
  - **TVDamage Color**: simulates a color TV.
  - **TVDamage Mono**: simulates a black & white TV.

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

- **Reception Master** (Default: 0.4, Range: 0 or greater)
  Master control for all reception-oriented artifacts: static, interference, ghosting, horizontal and vertical hold, hum bars, and color stripes. Turn to zero to get perfect reception, i.e. zero of each of the above artifacts.

- **Interference Amp** (Default: 0.6, Range: 0 or greater)
  Simulates interference from nearby electrical devices (electric motors, cordless phones, and so on). The look is a pattern of semi-regularly spaced random color dots. The dot size is controlled by the TV Pixels parameter. Scaled by Reception Master.

- **Ghost Amp** (Default: 0.6, Range: 0 or greater)
  Ghosts are copies of the image that result from multipath distortion between the transmitter and the TV. Turn up this parameter to get stronger ghosts. Scaled by Reception Master.

- **Horizontal Hold** (Default: 0.5, Range: 0 or greater)
  This causes the image to shift horizontally in a semi-random way, simulating a TV with a bad horizontal hold circuit, or a signal not strong enough to engage the horizontal hold. Scaled by Reception Master.

- **Vertical Hold** (Default: 0.8, Range: 0 or greater)
  This causes the image to shift vertically in a rolling motion, and is normally caused by a weak signal preventing the TV from locking on. This parameter controls the fraction of the time that the image is having hold problems. Set to zero for no vertical hold problems. Scaled by Reception Master.

- **Bars Brightness** (Default: 0.3, Range: 0 or greater)
  Power line hum and other TV problems can cause rolling light and dark bars to crawl up the screen. This can also be caused by failure to synchronize a video camera to the TV output. This parameter controls the overall strength of these bars. There are two sets of bars, one large and one small, that mutually interfere. This parameter controls the overall brightness scale of the bars. Turn to zero for no bars. Scaled by Reception Master.

- **Color Stripes Amplitude** (Default: 0.2, Range: 0 or greater)
  Another common form of interference, color stripes are caused by phase shifts in the chroma signal, among other things. This parameter controls the overall brightness of the color stripes. Scaled by Reception Master.

- **Fast Forward Amount** (Default: 0, Range: 0 or greater)
  Generates a VCR fast-forward look with torn bars across the screen.

- **Tape Dropout Brightness** (Default: 0, Range: 0 or greater)
  Generates VCR dropouts on random frames, at random times.

- **Vignette Darkness** (Default: 0, Range: 0 to 1)
  Vignetting is darkening of the image towards the corners and sides of the image. This parameter controls how much the outer corners of the screen should be darkened (vignetted). 0 gives no vignetting, 1 gives maximum darkening.

- **Static Amplitude** (Default: 0.8, Range: 0 or greater)
  Scales the brightness of the static noise. Scaled by Reception Master. The static dot size is controlled by the TV Pixels parameter.

- **Static Density** (Default: 0.7, Range: 0.01 to 1)
  Density of the static; turn up to get more static pixels; turn down to get only occasional static pixels.

- **Frequency** (Default: 1.28, Range: 0 to 500)
  Interference frequency. The look is very sensitive to this parameter. Fractional values like 0.3 or 1.23 look better than integers. Animating it very slightly, say from 1.27 to 1.3 gives a nice look.

- **Dots Speed** (X & Y, Default: [100 -10], Range: any)
  The dot pattern moves with this speed over time in X and Y.

- **Jitter Amount** (Default: 10, Range: 0 to 1000)
  Turning this up makes the dot pattern jitter randomly between frames for more realism.

- **Num Ghosts** (Integer, Default: 5, Range: 0 to 30)
  The number of ghost images. Some may be ahead (to the left of) the source image, most will be to the right. Some will be positive and some negative (inverted). See Shift and Negative Ghosts below.

- **Negative Ghosts** (Default: 0.5, Range: 0 to 1)
  The fraction of the ghosts that are negative (inverted), on average.

- **Spacing** (Default: 0.2, Range: 0 or greater)
  The fraction of the image width over which ghost images are spread out.

- **Vary Position** (Default: 0.3, Range: 0 to 1)
  Controls the regularity of the ghost image spacing. Set to zero for regularly spaced ghosts; set to one for random positioning.

- **Shift** (Default: 0.5, Range: -1 to 1)
  Shifts the ghost images to the left or right, without shifting the main image.

- **Blur** (Default: 0, Range: 0 or greater)
  Blurs the ghost images without blurring the main image or any other artifacts.

- **H Frequency** (Default: 1.25, Range: 0 or greater)
  Vertical frequency of the horizontal-hold waves.

- **H Time Vary** (Default: 0.5, Range: 0 or greater)
  Modulates the horizontal-hold waves over time by this amount. When increased, some frames will have more horizontal shifting while other frames will have less.

- **H Octaves** (Integer, Default: 3, Range: 1 to 10)
  Octaves for the horizontal hold waves. Increase for spikier look, decrease for smoother waves.

- **Border Width** (Default: 0.05, Range: 0 or greater)
  A TV signal has a black border outside the displayed area; this becomes visible when the horizontal hold isn't working. This parameter controls the width of that black border. On the other side of the border, you see another copy of the image.

- **V Frequency** (Default: 2, Range: 0 or greater)
  The frequency of vertical hold jumps. Decrease to get a more consistent rolling motion, or increase to get a jumpier look.

- **V Speed** (Default: 2, Range: 0 or greater)
  The average speed of the vertical hold rolling motion over time.

- **V Random** (Default: 0.1, Range: 0 or greater)
  Controls how much randomness there is in the vertical hold rolling motion. Set to zero for smooth rolling, 1 or more for jittery behavior.

- **Border Height** (Default: 0.1, Range: 0 or greater)
  Like Border Width, this controls the vertical border between frames that becomes visible when vertical hold is not locked. Some static and closed-captioning and timecode information will typically be visible in this border.

- **Border Data** (Default: 1, Range: 0 to 10)
  Brightness of the dots and lines that appear in the vertical blanking interval specified by Border Width.

- **Bar Roll Speed** (Default: 0.5, Range: -10 to 10)
  The speed of the bars rolling up the screen. Turn negative for downward rolling.

- **Bar Sharpness** (Default: 0.5, Range: 0.1 to 10)
  Sharpens or smooths the top and bottom edges of the main bars. Set to zero for no main bars; you will only see the smaller bars.

- **Bar Frequency** (Default: 1, Range: 0.1 or greater)
  The frequency of the bars; turn up for more thinner bars, turn down for fewer fat bars.

- **Bar1 Width** (Default: 0.35, Range: 0 to 1)
  Fraction of the main bar that is light; the rest is dark.

- **Bar2 Rel Frequency** (Default: 6, Range: 1 or greater)
  Controls the frequency of the smaller bars.

- **Bar2 Sharpness** (Default: 0.5, Range: 0.01 to 10)
  Sharpens or smooths the top and bottom edges of the smaller bars. Set to zero for no small bars; you will only see the main bars.

- **Color Frequency** (Default: 10, Range: 1 to 40)
  Spatial frequency of the color stripes.

- **Color Angle** (Default: 160, Range: any)
  Angle of the stripes.

- **Roll Speed** (Default: 3, Range: 0 or greater)
  Controls how fast the stripes roll over time.

- **Band Frequency** (Default: 4, Range: 0 or greater)
  How many fast-forward bands to create.

- **Band Shift** (Default: 0.1, Range: 0 or greater)
  Shifts the fast-forward bands up or down.

- **Band Height** (Default: 0.16, Range: 0 to 1)
  The height of each fast-forward band.

- **Dropout Length** (Default: 0.25, Range: 0 to 1)
  The average length of each dropout scanline.

- **Dropout Gap Length** (Default: 0.2, Range: 0 to 2)
  The average length of the gaps between the dropouts.

- **Dropout Y Freq** (Default: 5, Range: 0 to 50)
  The dropouts appear on random scanlines according to a noise function with this frequency. Decrease to get a few large bands of dropouts; increase to get lots of small bands of dropouts.

- **Dropout Y Threshold** (Default: 0.75, Range: 0 to 1)
  Increase to cover more of the screen with dropouts (on average); decrease to cover less of it. If you don't see any dropouts at all on some frames, increase this parameter.

- **Dropouts Always** (Default: 1, Range: 0 to 1)
  Dropouts only appear on some frames; increase this parameter to see dropouts on more frames, so they occur more frequently in time. If you don't see dropouts on any frames, increase this parameter.

- **Vignette Radius** (Default: 1, Range: 0 or greater)
  Distance from the center where the vignetting starts.

- **Vignette Edge Softness** (Default: 0.5, Range: 0 or greater)
  The width of the vignette's soft edge. Larger values give softer, less visible edges.

- **Vignette Rel Height** (Default: 0.75, Range: 0.1 or greater)
  Controls the aspect ratio of the vignette ellipse. This should normally be set to the aspect ratio of the image, e.g. .75 for NTSC.

- **Scanlines** (Default: 0.1, Range: 0 or greater)
  Creates visible scanlines in the image. Increase to get more intense scanlines, or set to zero for no scanlines. The width of the scanlines is controlled by the TV Pixels parameter and Scanlines Rel Freq.

- **Scanlines Rel Freq** (Default: 1, Range: 0 or greater)
  Relative frequency of the TV scanlines. Increase to get more scanlines, decrease to get fewer large scanlines. Note that the number of scanlines is also controlled by the TV Pixels parameter.

- **Orthicon** (Default: 0, Range: 0 or greater)
  Darkens the clip at areas around parts of the source clip that are brighter than the given threshold, to simulate a 1950s 'orthicon' TV camera look. Most useful in black & white mode.

- **Threshold** (Default: 0.7, Range: 0 or greater)
  Darkening will occur around locations in the source clip that are brighter than this value. A value of 0.9 causes dark glows from only the brightest spots. A value of 0 causes glows for every non-black area.

- **Darks Width** (Default: 0.2, Range: 0 or greater)
  Scales the dark glow distance.


### Color Correct Parameters:

Hue Shift:
*Default:
*0,
*Range:
*any.Shift the color hues by this amount.

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
*any.Adds this gray value to the darker regions of
the result. This can be negative to increase contrast.

Tint Lights:
*Default rgb:
*[1 1 1].Scales the result by this color, thus tinting the
lighter regions.

Tint Darks:
*Default rgb:
*[0 0 0].Adds this color to the darker regions of the result.
Set this to a dark red-orange color for a negative-film effect look.

Turn Off:
*Default:
*0,
*Range:
*0 to 1.Animate this parameter from 0 to 1 to simulate the TV turning off.
The image will turn white and shrink to a dot in the center, with a flash near the end.

Flare Width:
*Default:
*1,
*Range:
*0 or greater.Width of the flare or flash near the end of the turn-off sequence.
Set to zero to omit this flash.

Flare Brightness:
*Default:
*2,
*Range:
*0 or greater.Brightness of the flare or flash near the end of the turn-off sequence.

Fade Out Time:
*Default:
*0.7,
*Range:
*0 to 1.The length of the fade out during the Turn Off
sequence. A value of one produces a smooth, gradual fade as Turn Off increases. As the value decreases,
the fade will start later (at higher values of Turn Off), and progress more quickly. When Fade Out Time is
set to zero, the fade will cut to black instantly when Turn Off reaches one.

Fish Eye:
*Default:
*0,
*Range:
*any.Expands the center of the source clip as if viewed
through a fish-eye lens. This gives an old-time slightly rounded TV
look.

Tv Pixels:
*Default:
*720,
*Range:
*1 or greater.The number of 'TV pixels' across the screen. Controls the size of the static,
interference, scanlines, and dropouts. Lower this to simulate a lower resolution TV.

Downsample:
*Check-box, Default:
*on.If checked, reduces output resolution according to TV Pixels.

Seed:
*Default:
*0.123,
*Range:
*0 or greater.Used to initialize the random number generator. The actual
seed value is not significant, but different seeds give different
results and the same value should give a repeatable result.

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
