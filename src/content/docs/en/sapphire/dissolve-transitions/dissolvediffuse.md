---
title: DissolveDiffuse
---

## S_DissolveDiffuse

Transitions between two input clips by scrambling the pixels of the inputs
within an area determined by Max Amount. The first clip is diffused away
while the second clip is diffused into place. The Dissolve Percent
parameter should be animated to control the transition speed. The
pixelated look of this effect depends on the image resolution, so it is
recommended to test your final resolution before processing.

In the Sapphire Transitions effects submenu.

![DissolveDiffuse](../_static/DissolveDiffuse.jpg)


### Inputs:

- **Foreground**: The current layer. Starts the transition with this clip.

- **Background**: Defaults to None. Ends the transition with this clip. If this input is not provided, a fully transparent background is used, showing whatever is behind it. Note that the background can not be diffused during the transition unless this input is provided.


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

- **Max Amount** (Default: 0.2, Range: 0 or greater)
  Scales the magnitudes of the diffusion distances.

- **Rel Amount** (X & Y, Default: [1 1], Range: 0 or greater)
  Scales the relative horizontal and vertical amounts of diffusion.

- **Wrap** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  Determines the method for accessing outside the borders of the source images.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less
visible with this method.

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  These 4 parameters, Crop Top , Crop Bottom , Crop Left, and Crop Right , allow selecting a rectangular subsection of the input image to be processed. If the Wrap parameters are set to "No" the exposed borders will be transparent. If the Wrap is "Tile" or "Reflect" the source image is wrapped on the new cropped borders to fill the frame. This can make it easier to avoid artifacts due to distorting an image with bad edges.

- **Show Max Amount** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Max Amount parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

