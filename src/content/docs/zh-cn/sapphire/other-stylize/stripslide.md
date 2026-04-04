---
title: StripSlide
---

## S_StripSlide

Breaks a clip into strips and slides them off the
screen one at a time to reveal the Background.

In the Sapphire Stylize effects submenu.

![StripSlide](../_static/StripSlide.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Background**: Defaults to None. This clip is revealed as the Source slides away.


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

- **Combine Masks** (Popup menu, Default: Intersect)
  Determines how to combine the Mocha Mask and Input Mask when both are supplied to the effect.
  - **Union**: Uses the area covered by both masks together.
  - **Intersect**: Uses the area that overlaps between the two masks.
  - **Mocha Only**: Ignore the Input Mask and only use the
Mocha Mask.

- **Amount** (Default: 0.5, Range: 0 to 1)
  Controls the progress of the slide effect. At zero, the Source is fully visible, and at one the Background is fully visible.

- **Motion Blur** (Default: 0.3, Range: 0 or greater)
  Scales the amount of motion blur to use.

- **Strip Size** (Default: 0.1, Range: 0.01 or greater)
  The width of the strips. This parameter can affect the timing of the strips, so animating it is not recommended.

- **Randomize Size** (Default: 0, Range: 0 or greater)
  Makes some strips larger and some smaller, at random.

- **Strip Angle** (Default: 0, Range: any)
  Controls the angle along which the strips are divided, and also the direction in which they slide. This parameter can affect the timing of the strips, so animating it is not recommended.

- **Strip Shift** (Default: 0, Range: any)
  Adjusts the position of the strip boundaries. This parameter can affect the timing of the strips, so animating it is not recommended.

- **Speed** (Default: 10, Range: 1 or greater)
  The speed at which each strip moves. As speed increases, the delay between strips becomes larger. If the speed is low, many strips will be in motion at the same time, creating a wave or ripple effect. This parameter affects the timing of the strips, so animating it is not recommended.

- **Slow Start** (Default: 1, Range: 0 to 1)
  Controls the acceleration of each strip as it moves. If set to zero, the strip will start moving at full speed. With larger values, the strip will start moving more slowly and accelerate up to its full speed, resulting in smoother motion.

- **Order** (Popup menu, Default: Top Down)
  Controls the order in which the strips slide off screen.
  - **Top Down**: in order from top to bottom.
  - **Bottom Up**: in order from bottom to top.
  - **Random**: random order.
  - **Center Out**: outward from the center, alternating strips above and below the center.
  - **Edges In**: inward from the edges, alternating strips above and below the center.

- **Seed** (Default: 0.123, Range: 0 or greater)
  Initializes the random number generator for random strips sizes and order. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  These 4 parameters, Crop Top , Crop Bottom , Crop Left, and Crop Right , allow selecting a rectangular subsection of the input image to be processed. If the Wrap parameters are set to "No" the exposed borders will be transparent. If the Wrap is "Tile" or "Reflect" the source image is wrapped on the new cropped borders to fill the frame. This can make it easier to avoid artifacts due to distorting an image with bad edges.

