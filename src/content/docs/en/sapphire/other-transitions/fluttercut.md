---
title: FlutterCut
---

## S_FlutterCut

Transitions between two clips by rapidly cutting
back and forth between them, optionally inserting solid colored or
inverted frames as well. The cuts of each clip can get longer or shorter
over the length of the transition.

In the Sapphire Transitions effects submenu.

![FlutterCut](../_static/FlutterCut.jpg)


### Inputs:

- **Foreground**: The current layer. Starts the transition with this clip.

- **Background**: Defaults to None. Ends the transition with this clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **From Start Frames** (Integer, Default: 2, Range: 1 or greater)
  The number of frames of the From clip in the first cycle of the transition.

- **From End Frames** (Integer, Default: 2, Range: 1 or greater)
  The number of frames of the From clip in the last cycle of the transition.

- **From Acceleration** (Default: 1, Range: 1 or greater)
  The speed at which cut lengths change. When set to 1, the lengths will change gradually between From Start Frames and From End Frames, reaching the final value on the last cycle. As the value increases, the cut length will change more quickly and reach its final value more quickly.

- **To Start Frames** (Integer, Default: 2, Range: 1 or greater)
  The number of frames of the To clip in the first cycle of the transition.

- **To End Frames** (Integer, Default: 2, Range: 1 or greater)
  The number of frames of the To clip in the last cycle of the transition.

- **To Acceleration** (Default: 1, Range: 1 or greater)
  The speed at which cut lengths change. When set to 1, the lengths will change gradually between To Start Frames and To End Frames, reaching the final value on the last cycle. As the value increases, the cut length will change more quickly and reach its final value more quickly.


### Colored Frames Parameters:

Color1:
*Default rgb:
*[0 0 0].A solid color added to the pattern.

Color1 Frames:
*Integer, Default:
*0,
*Range:
*0 or greater.The number of Color 1 frames in each cycle. This stays constant
throughout the transition.

Color1 Position:
*Popup menu, Default: After Both
*.The position of the Color 1 frames within the pattern.
*Before From:
*at the beginning of the cycle, before the From clip.*After From:
*in the middle of the cycle, between the From and To clips.*After To:
*at the end of the cycle, after the To clip.*After Both:
*in the middle of the cycle, and again at the end (but not at the very end of the transition, after the last cycle).

Color2:
*Default rgb:
*[1 1 1].Another solid color added to the pattern.

Color2 Frames:
*Integer, Default:
*0,
*Range:
*0 or greater.The number of Color 2 frames in each cycle. This stays constant
throughout the transition.

Color2 Position:
*Popup menu, Default: After Both
*.The position of the Color 2 frames within the pattern.
*Before From:
*at the beginning of the cycle, before the From clip.*After From:
*in the middle of the cycle, between the From and To clips.*After To:
*at the end of the cycle, after the To clip.*After Both:
*in the middle of the cycle, and again at the end (but not at the very end of the transition, after the last cycle).

### Invert Parameters:

Invert:
*Popup menu, Default: None
*.Invert some frames.
*None:
*nothing is inverted*From Clip:
*only frames from the From clip are inverted.*To Clip:
*only frames to the To clip are inverted.*Both Clips:
*frames from the From and To clips are inverted,
but solid color frames are not.*Custom Pattern:
*overlays a custom pattern of inverted frames over the
From/To/color pattern. The pattern of inverted frames is controlled by Invert Length,
Normal Before, and Normal After.

### Invert Pattern Parameters:

Invert Length:
*Integer, Default:
*1,
*Range:
*0 or greater.The number of consecutive frames to invert. The
cycle of inverted frames is controlled by this parameter, Normal
Before, and Normal After, and is independent of the From/To/Color cycle.
Has no effect unless Invert is set to Custom Pattern.

Normal Before:
*Integer, Default:
*1,
*Range:
*0 or greater.Leave this number of normal, non-inverted frames before each group of
inverted frames. Has no effect unless Invert is set to Custom Pattern.

Normal After:
*Integer, Default:
*0,
*Range:
*0 or greater.
Leave this number of normal, non-inverted frames after each group of
inverted frames. Has no effect unless Invert is set to Custom Pattern.
