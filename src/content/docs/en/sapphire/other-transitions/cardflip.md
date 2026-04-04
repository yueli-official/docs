---
title: CardFlip
---

## S_CardFlip

Transitions between two clips by sliding or spinning the outgoing clip
to reveal the incoming clip behind it. The Amount parameter should be animated to control
the transition speed. Adjusting Revolutions and Shift will give different kinds of transitions.

In the Sapphire Transitions effects submenu.

![CardFlip](../_static/CardFlip.jpg)


### Inputs:

- **Foreground**: The current layer. Starts the transition with this clip.

- **Background**: Defaults to None. Ends the transition with this clip. If this input is not provided, a fully transparent background is used, showing whatever is behind it. Note that the background can not be warped during the transition unless this input is provided.


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
  If enabled, a transition is performed automatically between the first and last frames of the layer. If this is off, the transition is performed manually by animating the Card Percent parameter.

- **Amount** (Default: 0, Range: 0 to 1)
  Auto Trans must be disabled for this parameter to be used. It determines the transition ratio between the From and To inputs, and would normally be animated from 0 to 100 to perform a complete transition. The curve controlling this parameter can be adjusted for more detailed control over the timing of the wipe.

- **Slow In** (Default: 0.5, Range: 0 to 1)
  If positive, causes the transition to start more gradually.

- **Slow Out** (Default: 0.5, Range: 0 to 1)
  If positive, causes the transition to end more gradually.

- **Revolutions** (Integer, Default: 1, Range: 0 or greater)
  The number of times the clip should flip over during the transition. Set this to 1 for a simple flip, 2 or more for a spinning transition, and 0 for a slide/shuffle.

- **Spin Direction** (Popup menu, Default: Left)
  The direction of spin.
  - **Left**: horizontal spin to the left.
  - **Right**: horizontal spin to the right.
  - **Up**: vertical spin upward.
  - **Down**: vertical spin downward.

- **Shift** (Default: 0, Range: 0 or greater)
  Slides the clips horizontally or vertically away from each other during the first half of the transition, then toward each other in the second half. Both clips end in the same position in which they started. Set to a value or 1 or greater to prevent the clips from overlapping at the mid-point of the transition.

- **Shift Direction** (Popup menu, Default: Left)
  The direction of shifting.
  - **Left**: The outgoing clip shifts left and the incoming clip shifts right.
  - **Right**: The outgoing clip shifts right and the incoming clip shifts left.
  - **Up**: The outgoing clip shifts up and the incoming clip shifts down.
  - **Down**: The outgoing clip shifts down and the incoming clip shifts up.

- **Perspective Amount** (Default: 1, Range: 0.25 to 4)
  Controls the amount of lens telescoping while the clips are flipping over. Increase for more 3D perspective.

- **Shadow Color** (Default rgb: [0 0 0])
  The color of the drop shadow cast from the front clip onto the back clip.

- **Shadow Opacity** (Default: 2, Range: 0 or greater)
  The opacity of the shadow, use values near 0 for subtle transparent shadows, or values near 1.0 for stronger shadows.

- **Shadow Blur** (Default: 0.088, Range: 0 or greater)
  Determines the softness of the shadow.

- **Shadow Shift** (X & Y, Default: [0 0], Range: any)
  The horizontal and vertical offset of the shadow.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

