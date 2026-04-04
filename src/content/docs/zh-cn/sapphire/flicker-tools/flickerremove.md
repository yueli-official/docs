---
title: FlickerRemove
---

## S_FlickerRemove

Removes temporal flickering from the Source clip.
For example, old footage with uneven exposure times can be smoothed out with
this effect. To use this effect, first position the corners of the rectangle
over an area where the average brightness should remain constant. A middle
or light gray area is best for this. Then select a Source frame that has the
desired brightness within the rectangle, and hit the Set Hold Level button.
When other frames are processed, their brightness will be scaled so the
average brightness within the rectangle is equal to the Hold Level. You can
keyframe different Hold Level values over time to account for desirable
brightness changes.

In the Sapphire Time effects submenu.

![FlickerRemove](../_static/FlickerRemove.jpg)


### Inputs:

- **Source**: The current layer. The clip to remove flicker from.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Rect Corner1** (X & Y, Default: [-0.583 -0.441], Range: any)
  The upper left corner of the rectangle which is used to measure the flicker, in screen coordinates.

- **Rect Corner2** (X & Y, Default: [0.583 0.441], Range: any)
  The lower right corner of the rectangle which is used to measure the flicker, in screen coordinates.

- **Hold Level** (Default: 0.5, Range: 0.01 or greater)
  The requested average output brightness for the area within the rectangle.

- **Set Hold Level** (Push-button)
  Pressing this button has a side effect of setting the Hold Level parameter to the average Source brightness in the rectangle at the current frame. It causes the output to equal the Source at this frame. This button retains no value itself, and is turned back off immediately after being pushed.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Rect** (Check-box, Default: on)
  Turns on or off the screen user interface widget for adjusting the Rect Corner corner parameters.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

