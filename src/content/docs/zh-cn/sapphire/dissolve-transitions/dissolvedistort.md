---
title: DissolveDistort
---

## S_DissolveDistort

Transitions between two input clips while
distorting each using the gradient of the other. The first clip is
warped away and faded out while the second clip is unwarped into
place and faded in. The Dissolve Percent parameter should be
animated to control the transition speed. Note that the Background
input must be provided or this effect will just perform a simple
dissolve without any distortion.

In the Sapphire Transitions effects submenu.

![DissolveDistort](../_static/DissolveDistort.jpg)


### Inputs:

- **Foreground**: The current layer. Starts the transition with this clip.

- **Background**: Defaults to None. Ends the transition with this clip. If this input is not provided, a fully transparent background is used, showing whatever is behind it. Note that the background can not be warped during the transition unless this input is provided.


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
  Auto Trans must be disabled for this parameter to be used. It determines the transition ratio between the Foreground and Background inputs, and would normally be animated from 0 to 100 to perform a complete transition. The curve controlling this parameter can be adjusted for more detailed control over the timing of the dissolve. The Slow In and Slow Out parameters, if positive, also adjust the transition ratio internally for a smoother start and/or end to the transition.

- **Amplitude** (Default: 1, Range: any)
  Scales the amount of distortion applied to both input clips. This can also be negative to turn expansions into contractions and vice versa.

- **Rel Amp From** (Default: 1, Range: any)
  Scales the relative distortion amplitude of the From clip.

- **Rel Amp To** (Default: -1, Range: any)
  Scales the relative distortion amplitude of the To clip.

- **Smoothness** (Default: 0.25, Range: 0 or greater)
  Smooths the distortions by this amount. Increase for large scale distortion, decrease for finer detailed distortion.

- **Rotate Warp Dir** (Default: 0, Range: any)
  Rotates the direction of the distortion. This can cause areas of similar brightness to be twisted instead of just expanded or shrunk.

- **Slow In** (Default: 0.2, Range: 0 to 1)
  If positive, causes the transition to start more gradually.

- **Slow Out** (Default: 0.2, Range: 0 to 1)
  If positive, causes the transition to end more gradually.

- **Wrap** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  Determines the method for accessing outside the borders of the source images.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less
visible with this method.

- **Filter** (Check-box, Default: on)
  If enabled, the image is adaptively filtered when it is resampled. This gives a better quality result when parts of the image are warped smaller.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

