---
title: DissolveLuma
---

## S_DissolveLuma

Transitions between two input clips using a pattern
derived from their luminances. One clip often appears to emerge
through the other. The Dissolve
Percent parameter should be animated to control the transition speed.

In the Sapphire Transitions effects submenu.

![DissolveLuma](../_static/DissolveLuma.jpg)


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

- **Softness** (Default: 0.1, Range: 0 to 1)
  Increase for softer and slower transitions.

- **Use Luma Of** (Popup menu, Default: Difference)
  Determines how the transition pattern is generated from the clips' luminance values.
  - **Difference**: similar areas transition first, different areas last.
  - **Subtract**: areas where the first clip is brighter
transition first, and areas where the second clip is brighter
transition last.
  - **Mult**: areas where both images are bright transition first, and
areas where either is dark are last.
  - **Screen**: areas where either image is bright transition first,
and areas where both are dark transition last.
  - **Foreground**: dark areas of the first clip disappear first,
bright areas last.
  - **Background**: bright areas of the second clip appear first, dark
areas last.

- **Invert Pattern** (Check-box, Default: off)
  If enabled, the transition pattern is reversed in time.

- **Smooth Pattern** (Default: 0, Range: 0 or greater)
  If positive, a blur is applied to the transition pattern. This can reduce noise and give clearer edges to transition lines.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

