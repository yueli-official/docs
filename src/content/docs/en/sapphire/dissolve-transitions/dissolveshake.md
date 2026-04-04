---
title: DissolveShake
---

## S_DissolveShake

Transitons between two clips by applying a
shaking motion to them, along with a quick dissolve. The shaking
uses translation, zooming, and/or rotation. It is random but
repeatable, so with the same parameters the same shaking motion is
generated each time. Turn on Motion Blur and adjust the Mo Blur
Length for different amounts of blur. Adjust the Amplitude and
Frequency for different shaking speeds and amounts. The Rand
parameters give detailed control of the random non-periodic shaking,
and the Wave parameters adjust the regular periodic shaking. The X,
Y, Z, and Tilt parameters control the horizontal, vertical, zoom,
and rotation amounts of shaking respectively.

In the Sapphire Transitions effects submenu.

![DissolveShake](../_static/DissolveShake.jpg)


### Inputs:

- **Foreground**: The current layer. Starts the transition with this clip.

- **Background**: Defaults to None. Ends the transition with this clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Transition Dir** (Popup menu, Default: Dissolve Off to Bg)
  Selects the direction of the transition.
  - **Dissolve Off to Bg**: transitions from the current layer to the Background.
  - **Dissolve On from Bg**: transitions from the Background to the current layer.

- **Auto Trans** (Popup YES-NO, Default: No)
  If enabled, a transition is performed automatically between the first and last frames of the layer. If this is off, the transition is performed manually by animating the Dissolve Percent parameter.

- **Dissolve Percent** (Default: 0, Range: 0 to 1)
  Auto Trans must be disabled for this parameter to be used. It determines the transition ratio between the Foreground and Background inputs, and would normally be animated from 0 to 100 to perform a complete transition. The curve controlling this parameter can be adjusted for more detailed control over the timing of the dissolve.

- **Dissolve Speed** (Default: 3, Range: 1 or greater)
  The speed of the dissolve between the From and To clips. When set to 1, the dissolve takes place over the entire duration of the effect. When set higher, the dissolve is shorter, although the shaking still takes place over the entire duration.

- **Amplitude** (Default: 3, Range: 0 or greater)
  Scales the amplitude of the shaking motion.

- **Frequency** (Default: 10, Range: 0 or greater)
  Increase for faster shaking, decrease for slower shaking. (Be careful if you animate frequency values because the resulting shake frequency is also affected by the rate of change of the value.)

- **Motion Blur** (Check-box, Default: on)
  Options for motion blur of the shaking motion.

- **Mo Blur Length** (Default: 0.5, Range: 0 or greater)
  Scales the amount of motion blur. Use around .5 when processing on fields or 1.0 for frames to give realistic motion blur. This parameter has no effect if Motion Blur is No .

- **Seed** (Default: 0, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Wrap** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  Determines the method for accessing outside the borders of the source images.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less
visible with this method.


### X Shake Parameters:

X Rand Amp:
*Default:
*0.2,
*Range:
*0 or greater.Amplitude of horizontal random shaking.

X Rand Freq:
*Default:
*1,
*Range:
*0 or greater.Frequency of horizontal random shaking.

X Wave Amp:
*Default:
*0,
*Range:
*0 or greater.Amplitude of horizontal regular wave shaking.

X Wave Freq:
*Default:
*0.5,
*Range:
*0 or greater.Frequency of horizontal regular wave shaking, in
cycles per second.

X Phase:
*Default:
*0,
*Range:
*any.
Time shift of the horizontal shaking.

### Y Shake Parameters:

Y Rand Amp:
*Default:
*0.1,
*Range:
*0 or greater.Amplitude of the vertical random shaking.

Y Rand Freq:
*Default:
*1,
*Range:
*0 or greater.Frequency of the vertical random shaking.

Y Wave Amp:
*Default:
*0,
*Range:
*0 or greater.Amplitude of the vertical regular wave shaking.

Y Wave Freq:
*Default:
*0.5,
*Range:
*0 or greater.Frequency of the vertical regular wave shaking, in
cycles per second.

Y Phase:
*Default:
*0,
*Range:
*any.
Time shift of the vertical shaking.

### Z Shake Parameters:

Z Rand Amp:
*Default:
*0,
*Range:
*0 or greater.Amplitude of the zoom random shaking.

Z Rand Freq:
*Default:
*1,
*Range:
*0 or greater.Frequency of the zoom random shaking.

Z Wave Amp:
*Default:
*0,
*Range:
*0 or greater.Amplitude of the zoom regular wave shaking.

Z Wave Freq:
*Default:
*0.5,
*Range:
*0 or greater.Frequency of the zoom regular wave shaking, in
cycles per second.

Z Phase:
*Default:
*0,
*Range:
*any.
Time shift of the zoom shaking.

### Tilt Shake Parameters:

Tilt Rand Amp:
*Default:
*0,
*Range:
*0 or greater.Amplitude of the rotational random shaking, in degrees.

Tilt Rand Freq:
*Default:
*1,
*Range:
*0 or greater.Frequency of the rotational random shaking.

Tilt Wave Amp:
*Default:
*0,
*Range:
*0 or greater.Amplitude of the rotational regular wave shaking, in degrees.

Tilt Wave Freq:
*Default:
*0.5,
*Range:
*0 or greater.Frequency of the rotational regular wave shaking, in
cycles per second.

Tilt Phase:
*Default:
*0,
*Range:
*any.
Time shift of the rotational shaking.

### Channels Parameters:

Red Amplitude:
*Default:
*1,
*Range:
*0 or greater.The relative amount of shaking in the red channel. Changing this value from
the default will cause the red channel to move more or less than the other color channels,
resulting in a color fringing or channel separation look.

Green Amplitude:
*Default:
*1,
*Range:
*0 or greater.The relative amount of shaking in the green channel.

Blue Amplitude:
*Default:
*1,
*Range:
*0 or greater.The relative amount of shaking in the blue channel.

Red Phase:
*Default:
*0,
*Range:
*any.The relative phase of the red channel. Positive values will move the red
channel ahead of the others in time, causing it to move first and the other channels to
follow. Negative values have the opposite effect, causing the red channel to lag behind
the others. Small values usually produce the best looks.

Green Phase:
*Default:
*0,
*Range:
*any.The relative phase of the green channel.

Blue Phase:
*Default:
*0,
*Range:
*any.The relative phase of the blue channel.

RGB Randomness:
*Default:
*0,
*Range:
*0 or greater.The amount of random motion in each color channel. Turn up this
parameter to cause all three color channels to move randomly on different paths, independent
of the overall shaking. This motion is scaled by X Rand Amp, Y Rand Amp, Z Rand Amp, and
Tilt Rand Amp.

RGB Frequency:
*Default:
*2,
*Range:
*0 or greater.
The frequency of the random color channel shaking.

### Other Parameters:

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
