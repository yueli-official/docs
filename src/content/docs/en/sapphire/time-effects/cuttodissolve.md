---
title: CutToDissolve
---

## S_CutToDissolve

Turns a cut within a single clip into a dissolve.
No heads or tails are required; just set the cut point (frame) and
CutToDissolve creates a dissolve around that point.
Note that this effect does not take two clips; just a single clip
already containing a cut. The Cut Point param is key to making it
work; whatever frames are on either side of that will be treated as
the cut, and the dissolve will be created around them.

In the Sapphire Time effects submenu.

![CutToDissolve](../_static/CutToDissolve.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Cut Point** (Integer, Default: 5, Range: 2 or greater)
  The frame where the cut happens. Press the Find Cut pushbutton to automatically find the cut frame.

- **Find Cut** (Push-button)
  Press this button to automatically search through the clip and attempt to find a cut, starting at the current Cut Frame. The frame number will be stored in Cut Point above. If no cuts are found within a few seconds, the search will stop. Click Find Cut again to continue searching.

- **Dissolve Length** (Integer, Default: 6, Range: 2 or greater)
  The total length of the dissolve. Half will be on the left side of the cut, half on the right side.

- **Slow In Out** (Default: 2, Range: 0.1 to 10)
  Set to 0 for a linear dissolve, increase to 2 for a more subtle slow-in-out transition.

- **Gamma** (Default: 1, Range: 0.1 to 10)
  Set to 1 for a video dissolve, increase a little for a more filmic look.

- **Show** (Popup menu, Default: Result)
  This can help you find the cut frame; set it to Cut Frames to see a split-screen view of the last outgoing frame and first incoming, based on Cut Point. You can also use it to only show one side or the other, with the interpolated dissolve frames.
  - **Result**: Show the result clip, containing the dissolve.
  - **Cut Frames**: Show a split-screen of the two cut frames, no matter where the play head currently is.
  - **A**: Show the A (outgoing) side of the cut: as if the B side were black.
  - **B**: Show the B (incoming) side of the cut: as if the A side were black.

- **Split For Cut** (Default: 0, Range: -1 to 1)
  Where to split the split-screen view of the Cut Point, when in Show:Cut Frames mode. Normally this has no effect.

