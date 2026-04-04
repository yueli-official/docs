---
title: DissolveAutoPaint
---

## S_DissolveAutoPaint

Fade in a 'paint-brushed' version of the starting clip. Decrease the complexity
of the painting until it is just a few colors, then transition to a 'paint-brushed' version of the second clip
which then grows in color and complexity until the second clip fades in.

In the Sapphire Transitions effects submenu.

![DissolveAutoPaint](../_static/DissolveAutoPaint.jpg)


### Inputs:

- **Foreground**: The current layer. Starts the transition with this clip.

- **Background**: Defaults to None. Ends the transition with this clip.


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

- **Dissolve Speed** (Default: 3, Range: 1 or greater)
  The speed of the dissolve between the From and To clips. When set to 1, the dissolve takes place over the entire duration of the effect. When set higher, the dissolve is shorter, although the glow ramp-up and ramp-down still takes the entire duration. Setting this to 10 can make the transition snappier and more like a flash-frame cut.

- **Paint Fade** (Default: 0.25, Range: 0 to 1)
  How much time the transition should take at the ends of the transition to fade the painted look in over the starting clip.

- **Style** (Popup menu, Default: Van Gogh)
  Selects the style of brush strokes.
  - **Van Gogh**: the stroke directions align with the edges found within the image.
  - **Hairy Paint**: the strokes are perpendicular to the edges within the image.
  - **Pointalize**: the strokes are cellular pointy shapes with no direction.

- **Min Brush Size** (Default: 0.04, Range: 0.0025 to 1)
  The size of the paint brush in the middle of the transition.

- **Max Brush Size** (Default: 0.4, Range: 0.0025 to 1)
  The size of the paint brush at the beginning and end of the transition.

- **Stroke Length** (Default: 2, Range: any)
  Determines the length of the brush strokes along the directions of edges in the source clip. If this is negative you can switch from VanGogh to HairyPaint styles and vice versa.

- **Stroke Align** (Default: 0.2, Range: 0 or greater)
  Increase to smooth out the directions of the strokes so nearby strokes are more parallel.

- **Smooth Colors** (Default: 0, Range: 0 or greater)
  Blurs the source by this amount before generating the brush strokes. Increase to cause the colors of nearby strokes to be more consistent.

- **Seed** (Default: 0, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Jitter Frames** (Integer, Default: 0, Range: 0 or greater)
  If this is 0, the locations of the strokes will remain the same for every frame processed. If it is 1, the locations of the stokes are re-randomized for each frame. If it is 2, they are re-randomized every second frame, and so on.

- **Sharpen** (Default: 1, Range: 0 or greater)
  The amount of post-process sharpening applied.

- **Sharpen Width** (Default: 0.2, Range: 0 or greater)
  The width at which to apply the post-process sharpening filter, relative to the stroke sizes. Higher values affect wider areas from the edges, lower values only affect areas near sharp edges.

