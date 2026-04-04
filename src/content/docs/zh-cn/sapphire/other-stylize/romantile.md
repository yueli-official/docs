---
title: RomanTile
---

## S_RomanTile

Generates a mosaic pattern based on the Source clip. Adjust the
Edge Attract parameter to get the tile corners to bias towards the edges in
the source. Increase Vary Shape to get a less regular tile pattern.

In the Sapphire Stylize effects submenu.

![RomanTile](../_static/RomanTile.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Matte**: Defaults to None. Defines the area that will be tiled.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mocha Project** (Default: 0, Range: 0 or greater)
  Brings up the Mocha window for tracking footage and generating masks.

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  Blurs the Mocha Mask by this amount before using. This can be used to soften the edges or quantization artifacts of the mask, and smooth out the time displacements.

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  Controls the strength of the Mocha mask. Lower values reduce the intensity of the effect.

- **Invert Mocha** (Check-box, Default: off)
  If enabled, the black and white of the Mocha Mask are inverted before applying the effect.

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  Scales the Mocha Mask. 1.0 is the original size.

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  The relative horizontal size of the Mocha Mask.

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  The relative vertical size of the Mocha Mask.

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  Offsets the position of the Mocha Mask.

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  Dilates or erodes the Mocha Mask by this pixel amount before using.

- **Dilation Quality** (Popup menu, Default: Fast)
  Selects whether Dilate Mocha adusts quickly in default Fast mode or looks better in High quality mode.
  - **Fast**: Dilate Mocha in Fast mode for quick adjustments.
  - **High**: Dilate Mocha in High quality mode for a better looking mask shape.

- **Bypass Mocha** (Check-box, Default: off)
  Ignore the Mocha Mask and apply the effect to the entire source clip.

- **Show Mocha Only** (Check-box, Default: off)
  Bypass the effect and show the Mocha Mask itself.

- **Combine Masks** (Popup menu, Default: Union)
  Determines how to combine the Mocha Mask and Input Mask when both are supplied to the effect.
  - **Union**: Uses the area covered by both masks together.
  - **Intersect**: Uses the area that overlaps between the two masks.
  - **Mocha Only**: Ignore the Input Mask and only use the
Mocha Mask.

- **Tile Size** (Default: 0.5, Range: 0 or greater)
  The width of an individual tile.

- **Tile Shape** (Popup menu, Default: Square)
  Determines the shape of the tiles.
  - **Square**: four sided tiles.
  - **Hexagon**: six sided tiles.

- **Vary Shape** (Default: 0.2, Range: 0 to 1)
  Controls the variation of the tile shape. Set to 0 for regularly shaped tiles. Set to 1 for randomly shaped tiles.

- **Tile Shift** (X & Y, Default: [0 0], Range: any)
  Translation offset of the result.

- **Tile Edge Sharpness** (Default: 0.9, Range: 0 to 1)
  How sharp to make the 3d lighting roll off on the edge of the tile. Set to 1 for a very sharp tile edge. Set to a lower number for a softer, more curved tile.

- **Tile Texture Freq** (Default: 50, Range: 0 or greater)
  The frequency controls how coarse or fine the bumpy texture on the tiles is.

- **Tile Roughness** (Default: 0.75, Range: 0 to 1)
  The height of the bumpy texture on the tiles.

- **Tile Height** (Default: 0.5, Range: 0 or greater)
  The strength of the lighting on the edge of the tiles.

- **Tile Opacity** (Default: 0.9, Range: 0 to 1)
  The opacity of the tiles. Set to 0 to show the source. Set to 1 to show only the tile.

- **Cracked Tiles** (Default: 0, Range: 0 to 1)
  How likely a tile is to crack along edges in the source. Set to 0 to get no cracked tiles. Set to 1 to see tiles with detectable edges crack. At .5 only tiles with strong edges will crack. Tiles with a very slow gradient will never crack.

- **Smooth Colors** (Default: 0.2, Range: 0 or greater)
  Control the variation in the color palette. Increase to make only very sharp image edges change tile colors.

- **Edge Attract** (Default: 0.2, Range: 0 to 1)
  How strongly the corners of the tiles should attract to the edges in the image.

- **Grout Color** (Default rgb: [0.4 0.4 0.4])
  The color of the grout between the tiles.

- **Grout Width** (Default: 0.1, Range: 0 to 1)
  The width of the grout between the tiles as a percentage of the tile size.

- **Grout Texture Freq** (Default: 150, Range: 0 or greater)
  The frequency controls how coarse or fine the bumpy texture in the grout is.

- **Grout Roughness** (Default: 0.5, Range: 0 to 1)
  The height of the bumpy texture in the grout.

- **Grout Opacity** (Default: 1, Range: 0 to 1)
  The opacity of the grout between the tiles. Set to 0 to show the source. Set to 1 to show only the grout.

- **Light Position** (X & Y, Default: [0.9 0.5], Range: any)
  The XY position of the light. This parameter can be adjusted using the Light Position Widget.

- **Light Brightness** (Default: 1, Range: 0 or greater)
  The tiles are lit with a 3d point light source. This param sets the brightness of that light. Set to 0 to disable the light. Increase value to increase the intensity of the light.

- **Light Color** (Default rgb: [0.5 0.5 0.5])
  The color of the light.

- **Light Z** (Default: 5, Range: 1 or greater)
  The height of the light source.

- **Crop To Alpha** (Check-box, Default: off)
  Crop the tiles to the source alpha. If a mask input is provided, the mosaic will be cropped to the mask as well. When turned off, tiles generated inside the opaque region of the image might stick out into the transparent regions. When turned on, the tiles themselves will be cropped at the edge of the opaque region.

- **Bg Brightness** (Default: 1, Range: 0 to 1)
  Scales the brightness of the background before combining with the romans. If 0, the result will contain only the roman image over black.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Matte Use** (Popup menu, Default: Alpha)
  Determines how the Matte input channels are used to make a monochrome matte.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Light Position** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Light Position parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

