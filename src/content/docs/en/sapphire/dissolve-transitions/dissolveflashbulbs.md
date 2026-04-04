---
title: DissolveFlashbulbs
---

## S_DissolveFlashbulbs

Simulates lots of flashbulbs going off while
dissolving between two clips. With many small flashes, can look
like a stadium scene. With a few large flashes, works well on a
celebrity red carpet clip.

In the Sapphire Transitions effects submenu.

![DissolveFlashbulbs](../_static/DissolveFlashbulbs.jpg)


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
  The speed of the dissolve between the From and To clips. When set to 1, the dissolve takes place over the entire duration of the effect. When set higher, the dissolve is shorter, although the flashbulb ramp-up and ramp-down still takes the entire duration. Setting this to 10 can make the transition snappier and more like a quick cut.

- **Flash Style** (Default: 0, Range: 0 or greater)
  Style of flashbulb to use. Several styles are available, or you can try some of the glares for a different look.

- **Max Flashes** (Default: 20, Range: 0 or greater)
  Maximum number of flashes per frame, in the middle of the dissolve.

- **Flash Randomness** (Default: 0.2, Range: 0 to 1)
  Increase to get more flashes on some frames (up to the values of Flashes) and fewer on others.

- **Flash Size** (Default: 0.4, Range: 0 or greater)
  Average size of flashes.

- **Flash Rel Height** (Default: 1, Range: 0 or greater)
  Use to squash or stretch flashes.

- **Brightness** (Default: 3, Range: 0 or greater)
  Overall brightness of the flashes.

- **Vary Brightness** (Default: 0.2, Range: 0 to 1)
  Increase to vary the brightness of each flashbulb in each frame.

- **Flash Gamma** (Default: 1, Range: 0.1 or greater)
  Brightens or darkens the midtones of the flashes. Can give a round, hard-edged look, or make the flashes more soft and subtle.

- **Hold Frames** (Integer, Default: 1, Range: 0 or greater)
  Each flash trails off slightly in time, to simulate persistence of vision as well as the effect of the filament cooling off in old-time flashbulbs. Hold Frames controls how long that trail lasts.

- **Flash Decay Rate** (Default: 0.1, Range: 0 to 1)
  How quickly the flashes decay over the Hold Frames time. Increase to make them stay on screen brighter, for longer; decrease to make them disappear quickly. Note that you may have to increase Hold Frames to see long-lived flash trails.

- **Combine** (Popup menu, Default: Add)
  Determines how the flash image is combined with the background.
  - **Screen**: blends the flashes with the background, which can help prevent
overly bright results.
  - **Add**: causes the flash image to be added to the background.

- **Seed** (Default: 0.1, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the flashes. The maximum of the red, green, and blue flash brightness is scaled by this value and combined with the background Alpha at each pixel.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Flip Vertically** (Check-box, Default: off)
  Flips flashes vertically if needed to achieve a consistent look.

- **Show Flash Size** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Flash Size parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

