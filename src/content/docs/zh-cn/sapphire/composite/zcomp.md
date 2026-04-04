---
title: ZComp
---

## S_ZComp

Layers a source input over or under a second source input based on the
difference of two depth images. The DepthA input should be a 'z' depth
image corresponding to the objects in the first input, and DepthB should be
a 'z' depth image corresponding to the objects in the second input.

In the Sapphire Composite effects submenu.

![ZComp](../_static/ZComp.jpg)


### Inputs:

- **SourceA**: The current layer. The first input image.

- **SourceB**: Defaults to None. The second input image.

- **DepthA**: Defaults to None. The depth image corresponding to the objects in SourceA

- **DepthB**: Defaults to None. The depth image corresponding to the objects in SourceB


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Anti Alias** (Default: 0, Range: 0 or greater)
  The amount of depth difference over which to interpolate the source inputs instead of taking just the closer one. Specified as a fraction of the entire depth range: 0 does no antialiasing, 1 interpolates over the entire depth range.

- **Invert Z** (Check-box, Default: off)
  Normally larger depth values (white) are treated as farther away and smaller values (black) are treated as near. When this is enabled, these depth values are reversed.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

