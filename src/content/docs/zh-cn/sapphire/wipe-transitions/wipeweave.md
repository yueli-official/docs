---
title: WipeWeave
---

## S_WipeWeave

Performs a wipe transition between two input clips
using a texture resembling perpendicular woven strands. The Wipe
Percent parameter should be animated to control the transition
speed.
Increase the Grad Add parameter to make the timing of the weave
pattern move across the screen during the wipe.
Increase the Border Width parameter to draw a border at the
wipe transition edges.

In the Sapphire Transitions effects submenu.

![WipeWeave](../_static/WipeWeave.jpg)


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
  If enabled, a transition is performed automatically between the first and last frames of the layer. If this is off, the transition is performed manually by animating the Wipe Percent parameter.

- **Wipe Percent** (Default: 0, Range: 0 to 1)
  Auto Trans must be disabled for this parameter to be used. It determines the transition ratio between the From and To inputs, and would normally be animated from 0 to 100 to perform a complete transition. The curve controlling this parameter can be adjusted for more detailed control over the timing of the wipe.

- **Strands** (Popup menu, Default: Grow)
  The direction of the weave pattern transition.
  - **Shrink**: The strands start large and shrink inwards.
  - **Grow**: The strands start small and grow outwards.

- **Edge Softness** (Default: 0, Range: 0 or greater)
  The width of the transition edges. Larger values will cause softer, less visible edges in the wipe pattern.

- **Frequency** (Default: 20, Range: 0.01 or greater)
  The frequency of the weave pattern. Increase for more and smaller elements, or decrease for fewer and larger.

- **Rel Length** (Default: 10, Range: 0.1 or greater)
  The relative length of the strands. Increase for longer thinner strands. Decrease for shorter thicker strands.

- **Octaves** (Integer, Default: 2, Range: 1 to 10)
  The number of summed layers of noise. Each octave is twice the frequency and half the amplitude of the previous. A single octave gives a smooth texture. Adding octaves makes the result approach a fractal (1/f) noise texture.

- **Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Shift** (X & Y, Default: [0 0], Range: any)
  Translation of the weave pattern.

- **H Speed X** (Default: 0, Range: any)
  The horizontal speed of the horizontal strands. If non-zero, the horizontal strands will automatically crawl along their lengths at this rate.

- **V Speed Y** (Default: 0, Range: any)
  The vertical speed of the vertical strands. If non-zero, the vertical strands will automatically crawl along their lengths at this rate.

- **Grad Add** (Default: 0, Range: -10 to 10)
  If positive, a gradient will be added to the timing of the transition pattern so it moves across the screen during the wipe. This parameter can be adjusted using the Wipe Widget if enabled, but the value must be positive to make this widget visible.

- **Grad Angle** (Default: 0, Range: any)
  The direction of the wipe gradient in degrees. This will have no effect unless Grad Add is positive. The Wipe Widget also allows adjusting this parameter.

- **Border Width** (Default: 0, Range: 0 or greater)
  If positive, a colored border is drawn at the wipe transition edges, using the border color, opacity, softness, and shift parameters below.

- **Border Color** (Default rgb: [0.75 0 0])
  The color of the border. This has no effect unless Border Width is positive.

- **Border Opacity** (Default: 1, Range: 0 to 1)
  The opacity of the border. Decrease to make the border transparent and allow the image under it to show through. This has no effect unless Border Width is positive.

- **Border Softness** (Default: 0, Range: 0 or greater)
  The softness of the border edges. This has no effect unless Border Width is positive.

- **Border Shift** (Default: 0, Range: any)
  Shifts the border ahead of or behind the transition edge. This has no effect unless Border Width is positive.

- **Border Glow** (Default: 0, Range: 0 or greater)
  Adds a glow along the border of the wipe. The value determines the brightness of the glow.

- **Glow Width** (Default: 0.1, Range: 0 or greater)
  The width of the glowing border.

- **Width Red** (Default: 1, Range: 0 or greater)
  Scales the red glow width. If the red, green, and blue widths are all equal, the glow will match Glow Color. Otherwise it will have a fringe of varying color.

- **Width Green** (Default: 1.2, Range: 0 or greater)
  Scales the green glow width.

- **Width Blue** (Default: 1.4, Range: 0 or greater)
  Scales the blue glow width.

- **Glow Color** (Default rgb: [1 1 1])
  The color of the glowing border.

- **Noise Amp** (Default: 1, Range: 0 or greater)
  The amount of noise to add to the glowing border.

- **Noise Freq** (Default: 16, Range: 0.1 to 20)
  The spatial frequency of the noise.

- **Noise Speed** (Default: 2, Range: any)
  The speed with which the noise changes or boils over time.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Wipe** (Check-box, Default: on)
  Turns on or off the screen user interface widget for adjusting the Grad Add, Grad Angle, and Wipe Percent parameters. The value of the Grad Add parameter must first be positive for this widget to be visible.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Glow Width** (Check-box, Default: off)
  Turns on or off the screen user interface for adjusting the Glow Width parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

