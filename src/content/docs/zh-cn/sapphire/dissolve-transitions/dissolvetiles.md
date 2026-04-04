---
title: DissolveTiles
---

## S_DissolveTiles

Transitions between two input clips while
breaking each up into tiles and scrambling them. The first clip
breaks apart and spreads out while the second clip coalesces
behind the first. The Dissolve Percent parameter should be
animated to control the transition speed.

In the Sapphire Transitions effects submenu.

![DissolveTiles](../_static/DissolveTiles.jpg)


### Inputs:

- **Foreground**: The current layer. Starts the transition with this clip.

- **Background**: Defaults to None. Ends the transition with this clip. If this input is not provided, a fully transparent background is used, showing whatever is behind it. Note that the background can not be warped during the transition unless this input is provided.


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
  Auto Trans must be disabled for this parameter to be used. It determines the transition ratio between the Foreground and Background inputs, and would normally be animated from 0 to 100 to perform a complete transition. The curve controlling this parameter can be adjusted for more detailed control over the timing of the dissolve. The Slow In and Slow Out parameters, if positive, also adjust the transition ratio internally for a smoother start and/or end to the transition.

- **Scramble Speed** (Default: 2, Range: any)
  The amount each input should be scrambled at the edges of the transiton. The incoming clip is scrambled by this amount at the beginning of the transition, and the outgoing clip is scramble by this amount at the end. Setting this to zero will result in no tiling on either clip.

- **Scramble Rel** (X & Y, Default: [1 1], Range: 0 or greater)
  The relative amounts of horiztonal and vertical scrambling.

- **Scramble Rel From** (Default: 1, Range: any)
  The relative amount of scrambling in the outgoing clip. Set this to zero if the outgoing clip shouldn't be scrambled at all.

- **Scramble Rel To** (Default: -1, Range: any)
  The relative amount of scrambling in the incoming clip. Set this to zero if the incoming clip shouldn't be scrambled at all.

- **Slow In** (Default: 0.2, Range: 0 to 1)
  If positive, causes the transition to start more gradually.

- **Slow Out** (Default: 0.2, Range: 0 to 1)
  If positive, causes the transition to end more gradually.

- **Tiles** (Default: 10, Range: 1 or greater)
  How many tiles across the image. Increase for many tiny tiles; decrease for a few large ones.

- **Tile Rel Width** (Default: 1, Range: 0.01 or greater)
  Scales the height of each tile.

- **Tile Rel Height** (Default: 1, Range: 0.01 or greater)
  Scales the width of each tile.

- **Dissolve Delay** (Default: 0.6, Range: 0 to 1)
  The delay before cross-dissolving between the From and To clips. If this is set to 1, the outgoing clip does not fade at all. If set to 0, the outgoing and incoming clips will dissolve smoothly throughout the transition.

- **Combine** (Popup menu, Default: From Over To)
  By default the outgoing From clip scrambles away, revealing the To clip scrambling in underneath it. Set this to To Over From to have the To clip scramble in on top of the From clip. Adjusting Scramble Rel From and Scramble Rel To along with this can give nice results.
  - **From Over To**: Composites the From (outgoing) clip over the To (incoming) clip,
which reveals the To clip as the From clip scrambles away. Works well with default settings
or with Scramble Rel To set to zero.
  - **To Over From**: Composites the To (incoming) clip over the From (outgoing) clip,
which scrambles the To clip in over the From clip. Works well with default settings or
with Scramble Rel From set to zero.

- **Rotate Warp Dir** (Default: 0, Range: any)
  Rotates the warping direction by this many degrees. Animate to rotate the tiles around for an interesting effect.

- **Seed** (Default: 0.5, Range: 0 or greater)
  Used to initialize the random number generator for tiling the clips. The actual seed value is not significant, but different values will give different results.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.
If your image has sharp color changes where the matte
channel also has sharp edges, you may get better results with Normal
mode.

