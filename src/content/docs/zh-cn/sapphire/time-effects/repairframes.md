---
title: RepairFrames
---

## S_RepairFrames

Repairs one or more frames of a clip by replacing them with a time-warped version of the surrounding frames.

In the Sapphire Time effects submenu.

![RepairFrames](../_static/RepairFrames.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **First Bad Frame** (Integer, Default: 5, Range: 2 or greater)
  The first bad frame to replace. Set Show:Bad Frame and scrub this param to help find the bad frame.

- **Bad Frame Count** (Integer, Default: 1, Range: 1 or greater)
  How many frames to repair, starting from First Bad Frame.

- **Show** (Popup menu, Default: Result)
  Set to Result for normal operation, showing the clip with the repaired frames. Set to Bad Frame only to help find the bad frames; once you've found the frames of interest, return to Result.
  - **Result**: Show the result clip, which is the source clip with the bad frames repaired.
  - **Bad Frame**: Show the first bad frame, no matter where the play head currently is.

