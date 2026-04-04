---
title: HyperPull
---

## S_HyperPull

Pulls the foreground away in z space before dissolving to the background,
and featuring some bonus color and shake effects for added style.

In the Sapphire Transitions effects submenu.

![HyperPull](../_static/HyperPull.jpg)


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

- **Slow In** (Default: 1, Range: 0 to 1)
  If positive, causes the transition to start more gradually.

- **Slow Out** (Default: 1, Range: 0 to 1)
  If positive, causes the transition to end more gradually.

- **Dissolve Speed** (Default: 5, Range: 1 or greater)
  The speed of the dissolve between the foreground and background.

- **Z Dist From** (Default: 5, Range: 0.001 or greater)
  Scales the 'distance' of the foreground. Values greater than 1.0 move it farther away and make it smaller. Values less than 1.0 move the image closer and enlarge it. Zooming in slightly can sometimes be used to hide edge artifacts.

- **Center XY From** (X & Y, Default: [0 0], Range: any)
  The center of the fish-eye warping function, in screen coordinates relative to the center of the foreground.

- **Rotate From** (Default: 0, Range: any)
  Rotates the foreground about the center location by this many degrees. The angle ramps up to this value as the transition proceeds.

- **Wrap From** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  Determines the method for accessing outside the borders of the foreground image.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less visible with this method.

- **Z Dist To** (Default: 0.001, Range: 0.001 or greater)
  Scales the 'distance' of the background. Values greater than 1.0 move it farther away and make it smaller. Values less than 1.0 move the image closer and enlarge it. Zooming in slightly can sometimes be used to hide edge artifacts.

- **Center XY To** (X & Y, Default: [0 0], Range: any)
  The center of the fish-eye warping function, in screen coordinates relative to the center of the background.

- **Rotate To** (Default: 0, Range: any)
  Rotates the background about the center location by this many degrees. The angle ramps up to this value as the transition proceeds.

- **Wrap To** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  Determines the method for accessing outside the borders of the background image.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less visible with this method.

- **Motion Blur** (Check-box, Default: on)
  Enables motion blur.

- **Blur From Z Dist** (Default: 0.7, Range: 0.001 or greater)
  The 'distance' of the From transformation. Increase to zoom out, decrease to zoom in.

- **Blur From Rotate** (Default: 0, Range: any)
  The rotation angle of the From transformation, in degrees, about the center.

- **Blur To Z Dist** (Default: 0.9, Range: 0.001 or greater)
  The 'distance' of the To transformation. Increase to zoom out, or decrease to zoom in.

- **Blur To Rotate** (Default: 0, Range: any)
  The rotation angle of the To transformation, in degrees, about the center. Note that if the From and To Rotate angles are very different, the interpolation between them will become less accurate.

- **Camera Shake** (Check-box, Default: off)
  Enables camera shake.

- **Amplitude** (Default: 2, Range: 0 or greater)
  Scales the amplitude of the shaking motion.

- **Frequency** (Default: 2, Range: 0 or greater)
  Increase for faster shaking, decrease for slower shaking.

- **Glow Brights** (Check-box, Default: on)
  Enables glow brights.

- **Glow Brightness** (Default: 3, Range: 0 or greater)
  Overall maximum brightness of the glow.

- **Glow Threshold** (Default: 0.2, Range: 0 or greater)
  Parts of the source clip that are brighter than this value get glowed. A value of 0.9 makes only the brightest spots glow. A value of 0 makes every non-black area glow.

- **Glow Width** (Default: 0, Range: 0 or greater)
  The width of the glow.

- **Width X** (Default: 1, Range: 0 or greater)
  Scales the horizontal glow width. Set to 0 for vertical only.

- **Width Y** (Default: 1, Range: 0 or greater)
  Scales the vertical glow width. Set to 0 for horizontal only.

- **Glow Darks** (Check-box, Default: off)
  Enables glow darks.

- **Darkness** (Default: 0.5, Range: 0 or greater)
  The magnitude of the dark glows.

- **Dark Threshold** (Default: 0.5, Range: 0 or greater)
  Parts of the source clip that are brighter than this value get glowed. A value of 0.9 makes only the brightest spots glow. A value of 0 makes every non-black area glow.

- **Dark Width** (Default: 0, Range: 0 or greater)
  Scales the dark glow distance. Note that a zero glow width still affects the dark areas; set the darkness parameter to zero if you want to pass the Source through unchanged.

- **Dark Width X** (Default: 1, Range: 0 or greater)
  Scales the horizontal dark width. Set to 0 for vertical only.

- **Dark Width Y** (Default: 1, Range: 0 or greater)
  Scales the vertical dark width. Set to 0 for horizontal only.

- **Warp Chroma** (Check-box, Default: on)
  Enables warp chroma.

- **Warp Amount** (Default: 0.6, Range: 0 or greater)
  Adjusts the overall amount of chroma warping of the result. The amount of warping ramps up to this value as the transition proceeds. Setting this to zero disables warping and leaves the image chroma unchanged.

- **Steps** (Integer, Default: 10, Range: 3 to 100)
  The number of color samples along the chroma warp spectrum to include. More steps give a smoother result, but require more time to process.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Distortion Amount** (Default: -1, Range: any)
  The amplitude of the fish-eye warping.

- **Distort RGB Amount** (Default: 0.2, Range: any)
  Scales the magnitude of the lens distortion for all channels. Make negative to invert the direction of the distortions.

