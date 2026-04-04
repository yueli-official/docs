---
title: DissolveUltraGlow
---

## S_DissolveUltraGlow

Transitions between two input clips while generating varieties of
glowing flashes. The clips dissolve into each other, while each one gets a glow which
ramps up and down over the duration of the effect.
Adjust After Glow parameters to generate a secondary glow on the result of the
primary glow. Optionally enhance the edges or combine the result with atmospheric noise.
The Dissolve Percent parameter should be animated
to control the transition speed.

In the Sapphire Transitions effects submenu.
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
  The speed of the dissolve between the From and To clips. When set to 1, the dissolve takes place over the entire duration of the effect. When set higher, the dissolve is shorter, although the glow ramp-up and ramp-down still takes the entire duration. Setting this to 10 can make the transition snappier and more like a flash-frame cut.

- **Brightness** (Default: 1.7, Range: 0 or greater)
  Scales the brightness of all the glows.

- **Color** (Default rgb: [1 1 1])
  Scales the color of the primary glow.

- **Threshold** (Default: 0.2, Range: 0 or greater)
  Glows are generated from locations in the source clip that are brighter than this value. A value of 0.9 causes glows at only the brightest spots. A value of 0 causes glows for every non-black area.

- **Threshold Add Color** (Default rgb: [0 0 0])
  This can be used to raise the threshold on a specific color and thereby reduce the glows generated on areas of the source clip containing that color.

- **Glow Width** (Default: 0.371, Range: 0 or greater)
  The width of the glowing border.

- **Glow Falloff** (Default: 0.35, Range: -2 to 2)
  Boost or cut the distance that the glow extends.

- **Glow Bias** (Default: 0, Range: -3 to 3)
  Amount to grow the outskirts of the thresholded result, or shrink if negative.

- **Glow From Alpha** (Default: 0, Range: 0 to 1)
  Set to 1 to generate glows from the alpha channel of the source input instead of the RGB channels. In this case the glows will not pick up color from the source and will typically be brighter. Values between 0 and 1 interpolate between using the RGB and the Alpha.

- **Width X** (Default: 1, Range: 0 or greater)
  Scales the horizontal glow width. Set to 0 for vertical only.

- **Width Y** (Default: 1, Range: 0 or greater)
  Scales the vertical glow width. Set to 0 for horizontal only.

- **Width Red** (Default: 1, Range: 0 or greater)
  Scales the red glow width. If the red, green, and blue widths are all equal, the glow will match Glow Color. Otherwise it will have a fringe of varying color.

- **Width Green** (Default: 1, Range: 0 or greater)
  Scales the green glow width.

- **Width Blue** (Default: 1, Range: 0 or greater)
  Scales the blue glow width.

- **Show** (Popup menu, Default: Result)
  Selects the type of output
  - **Result**: Shows the final result of combining the glow, source, and background.
  - **Threshold**: Shows the thresholded image that is used to generate the glow.

- **After Glow Width** (Default: 1.8, Range: 0 or greater)
  Scales the glow distance for the secondary glow.

- **After Glow Color** (Default rgb: [1 1 1])
  Scales the color of the secondary glow.

- **After Glow Stretch X** (Default: 0.3, Range: 0 or greater)
  Scales the horizontal secondary glow width.

- **After Glow Stretch Y** (Default: 0.1, Range: 0 or greater)
  Scales the vertical secondary glow width.

- **Horizontal Streaks** (Default: 0, Range: 0 or greater)
  Scales the appearance of narrow trails in the horizontal direction.

- **Vertical Streaks** (Default: 0, Range: 0 or greater)
  Scales the appearance of narrow trails in the vertical direction.

- **Edge Detect** (Check-box, Default: off)
  Enables edge detection.

- **Edge Combine** (Popup menu, Default: Edges Only)
  Determines how the detected edges are combined with the Source.
  - **Screen**: detected edges are blended with the Source using a screen operation.
  - **Add**: detected edges are added to the Source.
  - **Edges Only**: gives only the detected edges with no Source.

- **Edge Smooth** (Default: 0, Range: 0 or greater)
  Increase for thicker and smoother edges.

- **Edge Mode** (Popup menu, Default: Transparent)
  Determines the behavior when accessing areas outside the source image.
  - **Transparent**: Areas outside the source image are treated as transparent, which can produce
transparency around the edges of the image.
Select this for fastest rendering.
  - **Reflect**: Reflects the image outside the border.

- **Edge Fill** (Check-box, Default: on)
  Make areas within detected edges opaque

- **Edge Thin** (Default: 0.1, Range: 0 or greater)
  Subtracts this value from the detected edge result. Increase to remove unwanted noise from minor edges.

- **Atmosphere** (Check-box, Default: off)
  Atmosphere gives the effect of the glow shining through a dusty atmosphere and picking up light or getting shadowed. This parameter adjusts the amount, or amplitude, of the atmospheric effect. Zero gives a smooth glow, higher values give more dusty look.

- **Atmosphere Amp** (Default: 1, Range: 0 or greater)
  Atmosphere gives the effect of the glow shining through a dusty atmosphere and picking up light or getting shadowed. This parameter adjusts the amount, or amplitude, of the atmospheric effect. Zero gives a smooth glow, higher values give more dusty look.

- **Atmosphere Freq** (Default: 11.6, Range: 0.1 to 20)
  Controls the spatial frequency of the atmospheric noise. Turn this up higher to get finer details, turn down for broader overall variation.

- **Atmosphere Detail** (Default: 0.506, Range: 0 to 1)
  Controls the amount of fine detail in the atmosphere simulation. Decrease to get smoother atmosphere, increase for a more crunchy or grainy look.

- **Atmosphere Speed** (Default: 1, Range: any)
  The cloudy noise in the atmosphere evolves over time like real dust clouds; this parameter controls how fast the cloud pattern changes over time. Set to zero for a static pattern.

- **Atmosphere Lights** (Default: 0.5, Range: 0 or greater)
  Scales the atmosphere layer by this value. Increase for a more intense result.

- **Atmosphere Darks** (Default: 0, Range: 0 or greater)
  Adds this gray value to the darker regions of the atmosphere layer. This can be negative to increase contrast.

- **Atmosphere Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator for the atmospheric noise. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Apply Pre-Glow** (Check-box, Default: off)
  Enables combining atmosphere with the Source prior to any glows.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Glow Width** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Glow Width parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

