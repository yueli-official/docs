---
title: NearestColor
---

## S_NearestColor

Collects pixel colors from the input clip's frames that
are closest to the given Match Color. This can create, for example, a
background-only image from a clip with objects moving over a blue or
green-screen background. It can also be used to accumulate the color of a
moving object over a non-colored background. The collected colors are
reinitialized whenever any non-consecutive frame is processed, either the
first frame, reprocessing a given frame, or jumping to another frame. You
must process multiple frames of a clip in a row to observe the effect, and
clearing your image cache before rendering may sometimes be necessary.

In the Sapphire Time effects submenu.

![NearestColor](../_static/NearestColor.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Steps** (Integer, Default: 15, Range: 2 or greater)
  Adjusts the number of input frames from which color is collected.

- **Match Color** (Default rgb: [0 0 1])
  Pixel colors are kept that are 'nearest' to this color.

- **Chroma Weight** (Default: 1, Range: 0 or greater)
  The amount of influence hue has on the color matching. If this is 0, the pixels with the closest brightness to Match Color will be kept; if it is 2, the hue will have more influence and the brightness will have less.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

