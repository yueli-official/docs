---
title: StretchFrameEdges
---

## S_StretchFrameEdges

Stretch the edges of a 4x3 image while preserving the center, to hide the black pillars in a 16x9 comp. This effect takes the middle part of the Source clip and squeezes it, since viewing a 4x3 image in a 16x9 comp normally stretches it out to fit. The edges are not squeezed, so the image goes all the way out to the edges. The left and right edge portions of the image will appear stretched horizontally. Although intended for 4x3 conversion, it can work with any aspect ratios.

In the Sapphire Distort effects submenu.

![StretchFrameEdges](../_static/StretchFrameEdges.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Center Squeeze** (Default: 1.33, Range: 0 to 2)
  Amount to squeeze the center portion of the image. To fit a 4x3 image into a 16x9, normally squeeze by 4/3 or 1.333.

- **Center Width** (Default: 0.4, Range: 0 to 0.99)
  The center of the image (the middle half, if this parameter is 0.5) is squeezed uniformly, with no distortion. This parameter defines the un-distorted part of the image.

- **Border Width** (Default: 0, Range: 0 to 1)
  To reduce distortion at the edges, you can turn up Border Width to allow some black borders. The edges of the clip won't be distorted as severely because they don't have to stretch as far.

- **Shift X** (Default: 0, Range: -1 to 1)
  Shift the entire image left or right, to keep the interesting part of the image in the non-distorted part of the frame. Setting this to non-zero will reveal the edge of the clip, unless Wrap is set to Tile or Reflect.

- **Smooth** (Default: 0.5, Range: 0 to 1)
  When set to zero, the edges of the image are linearly stretched. This produces the least distortion at the very edge, but can give a visible seam where the center meets the edge zone. When set to one, the seam is fully hidden, but the very edges of the image will be fairly seriously distorted. Compromise values are in between zero and one.

- **Wrap** (Popup menu, Default: No)
  Determines the method for accessing outside the borders of the source image.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less
visible with this method.

