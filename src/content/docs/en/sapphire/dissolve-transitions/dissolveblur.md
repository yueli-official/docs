---
title: DissolveBlur
---

## S_DissolveBlur

Transitions between two input clips while blurring each.
The first clip is blurred and faded out while the second clip is
unblurred and faded in. The Dissolve
Percent parameter should be animated to control the transition speed.

In the Sapphire Transitions effects submenu.

![DissolveBlur](../_static/DissolveBlur.jpg)


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

- **Blur Amount** (Default: 2, Range: 0 or greater)
  Scales the width of the blur.

- **Blur Rel** (X & Y, Default: [1 0], Range: 0 or greater)
  The relative horizontal and vertical blur widths. Set Blur Rel X to 0 for a vertical-only blur, or set Blur Rel Y to 0 for a horizontal-only blur.

- **Blur Rel From** (Default: 1, Range: 0 or greater)
  Scales the amount of blur applied to the first clip. Set to 0 to fade out with no blur.

- **Blur Rel To** (Default: 1, Range: 0 or greater)
  Scales the amount of blur applied to the second clip. Set to 0 to fade in with no blur.

- **Blur Filter** (Popup menu, Default: Gauss)
  The type of convolution filter to blur with.
  - **Box**: uses a rectangular shaped filter.
  - **Triangle**: smoother, uses a pyramid shaped filter.
  - **Gauss**: smoothest, uses a gaussian shaped filter.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

