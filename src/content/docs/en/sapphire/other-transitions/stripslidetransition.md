---
title: StripSlideTransition
---

## S_StripSlideTransition

Transitions between two clips by breaking
them into strips and sliding them off the screen one at a time to
reveal the incoming clip.

In the Sapphire Transitions effects submenu.

![StripSlideTransition](../_static/StripSlideTransition.jpg)


### Inputs:

- **From**: The current layer. Starts the transition with this clip.

- **To**: Defaults to None. Ends the transition with this clip.


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
  If enabled, a transition is performed automatically between the first and last frames of the layer. If this is off, the transition is performed manually by animating the Strip Percent parameter.

- **Amount** (Default: 0, Range: 0 to 1)
  Auto Trans must be disabled for this parameter to be used. It determines the transition ratio between the From and To inputs, and would normally be animated from 0 to 100 to perform a complete transition. The curve controlling this parameter can be adjusted for more detailed control over the timing of the wipe.

- **Style** (Popup menu, Default: Slide Off)
  Controls which clip the slide is applied to.
  - **Slide Off**: The outgoing clip slides off to reveal the incoming clip.
  - **Slide On**: The incoming clip slides on over the outgoing clip.
  - **Side by Side**: The outgoing clip slides off while the incoming clip slides on
next to it.

- **Motion Blur** (Default: 0.3, Range: 0 or greater)
  Scales the amount of motion blur to use.

- **Strip Size** (Default: 0.1, Range: 0.01 or greater)
  The width of the strips. This parameter can affect the timing of the strips, so animating it is not recommended.

- **Randomize Size** (Default: 0, Range: 0 or greater)
  Makes some strips larger and some smaller, at random.

- **Strip Angle** (Default: 0, Range: any)
  Controls the angle along which the strips are divided, and also the direction in which they slide. This parameter can affect the timing of the strips, so animating it is not recommended.

- **Strip Shift** (Default: 0, Range: any)
  Adjusts the position of the strip boundaries. This parameter can affect the timing of the strips, so animating it is not recommended.

- **Speed** (Default: 10, Range: 1 or greater)
  The speed at which each strip moves. As speed increases, the delay between strips becomes larger. If the speed is low, many strips will be in motion at the same time, creating a wave or ripple effect. This parameter affects the timing of the strips, so animating it is not recommended.

- **Slow Start** (Default: 1, Range: 0 to 1)
  Controls the acceleration of each strip as it moves. If set to zero, the strip will start moving at full speed. With larger values, the strip will start moving more slowly and accelerate up to its full speed, resulting in smoother motion.

- **Order** (Popup menu, Default: Top Down)
  Controls the order in which the strips slide off screen.
  - **Top Down**: in order from top to bottom.
  - **Bottom Up**: in order from bottom to top.
  - **Random**: random order.
  - **Center Out**: outward from the center, alternating strips above and below the center.
  - **Edges In**: inward from the edges, alternating strips above and below the center.

- **Seed** (Default: 0.123, Range: 0 or greater)
  Initializes the random number generator for random strips sizes and order. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  These 4 parameters, Crop Top , Crop Bottom , Crop Left, and Crop Right , allow selecting a rectangular subsection of the input image to be processed. If the Wrap parameters are set to "No" the exposed borders will be transparent. If the Wrap is "Tile" or "Reflect" the source image is wrapped on the new cropped borders to fill the frame. This can make it easier to avoid artifacts due to distorting an image with bad edges.

