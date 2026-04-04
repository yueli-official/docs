---
title: FlickerMchMatteColor
---

## S_FlickerMchMatteColor

Adds color changes to the Source clip using
the color changes from a second Match clip, in the areas specified by
a Matte. To use this effect, select a frame where you want the Source
color unchanged, and hit the Set Match Color button. When other
frames are processed, the Source color will be scaled by the
average Match color within the Matte, relative to the Match
Color.

In the Sapphire Time effects submenu.

![FlickerMchMatteColor](../_static/FlickerMatchColor.jpg)


### Inputs:

- **Source**: The current layer. The clip to add flicker to.

- **Match**: Defaults to None. The clip to copy flicker from.

- **Matte**: Defaults to None. This clip specifies which Source areas to measure the flicker from. If this input is not provided, the Alpha of the Match input is used as the Matte instead. It can be inverted with the Invert Matte parameter.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Match Color** (Default rgb: [0.5 0.5 0.5])
  The average Match color in the Matte for which the Source input is unchanged.

- **Set Match Color** (Push-button)
  Pressing this button has a side effect of setting the Match Color parameter to the average Match clip color within the Matte at the current frame. It causes the output to equal the Source at this frame. This button retains no value itself, and is turned back off immediately after being pushed.

- **Matte Use** (Popup menu, Default: Alpha)
  Determines how the Matte input channels are used to make a monochrome matte.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

