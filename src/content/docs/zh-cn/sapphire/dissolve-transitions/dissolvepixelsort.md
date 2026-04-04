---
title: DissolvePixelSort
---

## S_DissolvePixelSort

Transitions between two input clips while sorting the dissolve result.

In the Sapphire Transitions effects submenu.

![DissolvePixelSort](../_static/DissolvePixelSort.jpg)


### Inputs:

- **Foreground**: The current layer. Starts the transition with this clip.

- **Background**: Defaults to None. Ends the transition with this clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: DissolvePixelSortLinear)
  Selects between several sort patterns.
  - **DissolvePixelSortLinear**: Sort pixels along parallel lines.
  - **DissolvePixelSortRadial**: Sort pixels along lines emanating from a point.
  - **DissolvePixelSortCircular**: Sort pixels along concentric circles.

- **Transition Dir** (Popup menu, Default: Dissolve Off to Bg)
  Selects the direction of the transition.
  - **Dissolve Off to Bg**: transitions from the current layer to the Background.
  - **Dissolve On from Bg**: transitions from the Background to the current layer.

- **Auto Trans** (Popup YES-NO, Default: No)
  If enabled, a transition is performed automatically between the first and last frames of the layer. If this is off, the transition is performed manually by animating the Dissolve Percent parameter.

- **Dissolve Percent** (Default: 0, Range: 0 to 1)
  Auto Trans must be disabled for this parameter to be used. It determines the transition ratio between the Foreground and Background inputs, and would normally be animated from 0 to 100 to perform a complete transition. The curve controlling this parameter can be adjusted for more detailed control over the timing of the dissolve.

- **Dissolve Speed** (Default: 5, Range: 1 or greater)
  The speed of the dissolve between the From and To clips. When set to 1, the dissolve takes place over the entire duration of the effect. When set higher, the dissolve is shorter, although the pixel sort ramp-up and ramp-down still takes the entire duration. Setting this to 10 can make the transition snappier and more like a flash-frame cut.

- **Max Percentage** (Default: 1, Range: 0 or greater)
  The threshold at the point in the transition of the maximum or full sort. Set to 1 for maximum sort at the height of the transition.

- **Sort Angle** (Default: 0, Range: any)
  The angle to sort parallel lines at in Linear mode.

- **Center** (X & Y, Default: [0 0], Range: any)
  The center in Radial mode.

- **Start Angle** (Default: 0, Range: any)
  The angle to start sorting in Radial mode.

- **Degrees Sorted** (Default: 360, Range: 0 to 360)
  How many degrees around the center should be sorted in Radial mode.

- **Inner Radius** (Default: 0, Range: 0 or greater)
  The radius of unsorted pixels around the center in Radial mode.

- **Ray Length** (Default: 1, Range: 0 or greater)
  The length of the sorted pixel chunks in Radial mode.

- **Vary Radius** (Default: 0, Range: 0 to 1)
  How much to vary the starting pixel of the radiating lines in Radial mode.

- **Start Angle** (Default: 0, Range: any)
  The angle to start sorting in Circular mode.

- **Degrees Sorted** (Default: 360, Range: 0 to 360)
  How many degrees around the center should be sorted in Circular mode.

- **Circle Center** (X & Y, Default: [0 0], Range: any)
  The center of the concentric circles in Circular mode.

- **Vary Start** (Default: 0, Range: 0 to 1)
  How much to vary the start in Circular mode.

- **Inner Radius** (Default: 0, Range: 0 or greater)
  The minimum circle to sort in Circular mode.

- **Thickness** (Default: 1, Range: 0 or greater)
  The number of concentric circles to sort in Circular mode.

- **Sort Direction** (Popup menu, Default: sort above threshold)
  Controls whether pixels over the threshold or under the threshold will be sorted.
  - **sort below threshold**: Sort pixels with a value lower than the threshold.
  - **sort above threshold**: Sort pixels with a value higher than the threshold.

- **Reverse Sort Direction** (Check-box, Default: off)
  Determines whether to sort the pixels in an ascending or descending order.

- **Randomly Restart Sort** (Default: 100, Range: 0 to 1000)
  How often to break up the sorted pixel chunks.

- **Sort Type** (Popup menu, Default: monochrome)
  Determines how to analyze the colors of the pixels for sorting.
  - **monochrome**: Sort on the monochrome value of the pixels.
  - **average**: Sort on the average value of the color channels in the pixels.
  - **minimum**: Sort on the smallest channel in the pixels.
  - **maximum**: Sort on the largest channel in the pixels.
  - **red**: Sort on the red channel in the pixels.
  - **green**: Sort on the green channel in the pixels.
  - **blue**: Sort on the blue channel in the pixels.
  - **hue**: Sort on the hue of the pixels.
  - **saturation**: Sort on the saturation of the pixels.
  - **brightness**: Sort on the brightness of the pixels.

- **Seed** (Default: 0.273, Range: 0 or greater)
  Initializes the random number generator for the sort restarts.

- **Soften Threshold Mask** (Default: 0.1, Range: 0 or greater)
  Blur the mask generated by thresholding the image.

- **Downsample** (Check-box, Default: off)
  If checked, reduces output resolution according to Sort Resolution.

- **Sort Resolution** (Integer, Default: 720, Range: 1 or greater)
  Target output resolution in pixels when Downsample is checked.

- **Mix With Dissolve** (Default: 0, Range: 0 to 1)
  Softens the look of the pixel sort by interpolating the raw dissolve with the pixel sorted result.

- **Show** (Popup menu, Default: Result)
  Selects the output option.
  - **Result**: Shows the result of the pixel sort.
  - **Raw Sort Values**: Shows the values the pixel sort will use for sorting.
  - **Threshold Mask**: Shows the mask of pixels to sort based on the threshold parameter.
  - **Random Restart Noise**: Shows the pixels that will cause the sort lines to restart.
  - **Combined Mask**

