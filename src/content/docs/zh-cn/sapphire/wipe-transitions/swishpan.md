---
title: SwishPan
---

## S_SwishPan

Transitions between two input clips by sliding one clip
off the frame and the other clip on, and adding motion blur to
give the appearance of a quick pan. This works best when the duration of
the transition is short.

In the Sapphire Transitions effects submenu.

![SwishPan](../_static/SwishPan.jpg)


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
  If enabled, a transition is performed automatically between the first and last frames of the layer. If this is off, the transition is performed manually by animating the Swish Percent parameter.

- **Wipe Percent** (Default: 0, Range: 0 to 1)
  Auto Trans must be disabled for this parameter to be used. It determines the transition ratio between the From and To inputs, and would normally be animated from 0 to 100 to perform a complete transition. The curve controlling this parameter can be adjusted for more detailed control over the timing of the wipe.

- **Direction** (Popup menu, Default: Left)
  Direction that the clips move during the transition.
  - **Left**: Moves right-to-left
  - **Right**: Moves left-to-right
  - **Up**: Moves upward.
  - **Down**: Moves downward.

- **Blur Amount** (Default: 2, Range: 0 or greater)
  Amount of motion blur to use. If the direction is left or right, the blur is horizontal. If the direction is up or down, the blur is vertical.

- **Overlap** (Default: 0, Range: any)
  Amount to overlap the two clips. Where the clips overlap, they will be screened together. This is useful for eliminating bad edges.

- **Slow In** (Default: 0.5, Range: 0 to 1)
  If positive, causes the transition to start more gradually.

- **Slow Out** (Default: 0.5, Range: 0 to 1)
  If positive, causes the transition to end more gradually.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.
If your image has sharp color changes where the matte
channel also has sharp edges, you may get better results with Normal
mode.

