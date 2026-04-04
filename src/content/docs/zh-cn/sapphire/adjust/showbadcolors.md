---
title: ShowBadColors
---

## S_ShowBadColors

Identifies all pixels that fall outside a given
color range, and flags them with the same color so they can be seen
easily.

In the Sapphire Adjust effects submenu.

![ShowBadColors](../_static/ShowBadColors.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Min** (Default: 0, Range: 0 to 1)
  Minimum color value. Pixels where any color channel is less than this value will be marked with Low Color.

- **Max** (Default: 1, Range: 0 to 1)
  Maximum color value. Pixels where any color channel is greater than this value will be marked with High Color.

- **Min Luma** (Default: 0, Range: 0 to 1)
  Minimum luminance value. Pixels where the luminance is less than this value will be marked with Low Color.

- **Max Luma** (Default: 1, Range: 0 to 1)
  Maximum luminance value. Pixels where the luminance is greater than this value will be marked with High Color.

- **Min Chroma** (Default: 0, Range: 0 to 1)
  Minimum chrominance value. Pixels where the chroma is less than this value will be marked with Low Color.

- **Max Chroma** (Default: 1, Range: 0 to 1)
  Maximum chrominance value. Pixels where the chroma is greater than this value will be marked with High Color.

- **Min Rgb** (Default rgb: [0 0 0])
  Minimum values per color channel. Pixels where any color channel is below the corresponding channel of this parameter will be marked with Low Color.

- **Max Rgb** (Default rgb: [1 1 1])
  Maximum values per color channel. Pixels where any color channel is above the corresponding channel of this parameter will be marked with High Color.

- **High Color** (Default rgb: [1 0 0])
  Color to mark high pixels with. Any pixel that is above one of the Max parameters will be set to this color.

- **Low Color** (Default rgb: [0 0 1])
  Color to mark low pixels with. Any pixel that is below one of the Min parameters will be set to this color.

- **Output Matte** (Check-box, Default: off)
  If enabled, output a matte which is set to white for bad pixels and black otherwise.

- **Invert Matte** (Check-box, Default: off)
  If enabled, the matte is inverted to show black for bad pixels and white otherwise. Has no effect unless Output Matte is also enabled.

