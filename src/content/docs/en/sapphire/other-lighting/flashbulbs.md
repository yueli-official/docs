---
title: Flashbulbs
---

## S_Flashbulbs

Simulates lots of flashbulbs going off. With many small
flashes, can look like a stadium scene. With a few large flashes,
works well on a celebrity red carpet clip.

In the Sapphire Lighting effects submenu.

![Flashbulbs](../_static/Flashbulbs.jpg)


### Inputs:

- **Background**: The current layer. The clip to use as background.

- **Matte**: Defaults to None. Used to restrict the flares to a certain area of the image. Areas where this is white will get flashbulbs; black areas will be get no flashes.


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

- **Combine Masks** (Popup menu, Default: Union)
  Determines how to combine the Mocha Mask and Input Mask when both are supplied to the effect.
  - **Union**: Uses the area covered by both masks together.
  - **Intersect**: Uses the area that overlaps between the two masks.
  - **Mocha Only**: Ignore the Input Mask and only use the
Mocha Mask.

- **Flash Style** (Default: 0, Range: 0 or greater)
  Style of flashbulb to use. Several styles are available, or you can try some of the glares for a different look.

- **Brightness** (Default: 5, Range: 0 or greater)
  Overall brightness of the flashes.

- **Vary Brightness** (Default: 0.2, Range: 0 to 1)
  Increase to vary the brightness of each flashbulb in each frame.

- **Flashes** (Integer, Default: 2, Range: 0 or greater)
  Approximate number of flashes per frame.

- **Flash Randomness** (Default: 0.5, Range: 0 to 1)
  Increase to get more flashes on some frames (up to the values of Flashes) and fewer on others.

- **Flash Size** (Default: 1, Range: 0 or greater)
  Average size of flashes. This parameter can be adjusted using the Flash Size Widget.

- **Flash Size Rel** (X & Y, Default: [1 1], Range: 0 or greater)
  Use to squash or stretch flashes. This parameter can be adjusted using the Flash Size Widget.

- **Flash Gamma** (Default: 1, Range: 0.1 or greater)
  Brightens or darkens the midtones of the flashes. Can give a round, hard-edged look, or make the flashes more soft and subtle.

- **Hold Frames** (Integer, Default: 1, Range: 0 or greater)
  Each flash trails off slightly in time, to simulate persistence of vision as well as the effect of the filament cooling off in old-time flashbulbs. Hold Frames controls how long that trail lasts.

- **Flash Decay Rate** (Default: 0.1, Range: 0 to 1)
  How quickly the flashes decay over the Hold Frames time. Increase to make them stay on screen brighter, for longer; decrease to make them disappear quickly. Note that you may have to increase Hold Frames to see long-lived flash trails.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the background before combining with the flashbulbss. If 0, the result will contain only the flashbulbs image over black.

- **Combine** (Popup menu, Default: Add)
  Determines how the flash image is combined with the background.
  - **Screen**: blends the flashes with the background, which can help prevent
overly bright results.
  - **Add**: causes the flash image to be added to the background.
  - **Flashes Only**: shows the flashes over a transparent black background.

- **Seed** (Default: 0.1, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the flashes. The maximum of the red, green, and blue flash brightness is scaled by this value and combined with the background Alpha at each pixel.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Matte Use** (Popup menu, Default: Luma)
  Determines how the Matte input channels are used to make a monochrome matte.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Flip Vertically** (Check-box, Default: off)
  Flips flashes vertically if needed to achieve a consistent look.

- **Show Flash Size** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Flash Size parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

