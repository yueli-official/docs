---
title: Feedback
---

## S_Feedback

The previous frames of the input clip are transformed and
combined with the current frame to give a variety of effects inspired by
video feedback. The output of each processed frame is stored and then
combined with the next frame. The feedback is reinitialized whenever any
non-consecutive frame is processed: either the first frame, reprocessing a
given frame, or jumping to another frame. You must process multiple frames
of a clip in a row to observe the effect, and clearing your image cache
before rendering may sometimes be necessary.

In the Sapphire Time effects submenu.

![Feedback](../_static/Feedback.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Mask**: Defaults to None. Interpolate between the result and the Source input. White areas use the result of the effect. Black areas use the Source clip.


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

- **Max Steps** (Integer, Default: 15, Range: 2 or greater)
  Adjusts the number of feedback cycles that are rendered.

- **Prev Brightness** (Default: 0.8, Range: 0 or greater)
  For each frame, the previous output is scaled by this amount before it is combined with the new input frame. Normally this value should be less than 1.0 which causes previous frames to fade out over time. A value of 1.0 causes no fading, and values greater than 1.0 cause previous frames to become brighter over time.

- **Prev Color** (Default rgb: [1 1 1])
  For each frame, the previous output is scaled by this color before it is combined with the new input frame. This is similar to Prev Brightness but affects the colors of the previous frames instead of just the brightness.

- **Prev Hue Shift** (Default: 0, Range: any)
  Shifts the hue of the previous frames' colors, for each new frame.

- **Combine New** (Popup menu, Default: Ave)
  Selects the method for combining previous frames with the current frame.
  - **Ave**: The current frame is averaged with the previous
output, smearing moving objects out over time. The output is scaled
by Fade and the input is scaled by 1.0-Fade for a weighted average,
so Fade must be less than 1.0 for this to work properly. Unlike the
other combine options, Ave should never affect the brightness of
stationary objects in the clip.
  - **Max**: The colors of the current frame and previous frames
are combined with a maximum function. This makes the output frame
at least as bright as the current frame, and will make brighter
'trails' for example if you have bright objects moving on a dark
background.
  - **Screen**: The colors of the current frame and previous
frames are combined with a blend function. This can be used to
accumulate the colors of a moving clip. However, non-black regions
will become brighter with each frame.
  - **Add**: The colors of the current frame and previous frames
are added. This can also be used to accumulate the colors of a
moving clip, with the non-black regions becoming brighter at each
frame.
  - **Over**: The current frame is composited over the previous
frames using its Alpha channel. This uses pre-multiplied
compositing, so where the alpha is black the Source image should
normally also be black. If the input clip contains no Alpha
channel, the luminance is used instead.
  - **Under**: The current frame is composited under the previous
frames.
  - **Min**: The colors of the current frame and previous frames
are combined with a minimum function. This makes the output frame
no brighter than the current frame, and will often fade quickly to a black frame.

- **New Color** (Default rgb: [1 1 1])
  Scales the color of the current frame. Set this to the complement of Old Color to offset overly colored trails.

- **New Opacity** (Default: 1, Range: 0 to 10)
  Scales the opacity and brightness of the current frame.

- **Blur Amount** (Default: 0, Range: 0 or greater)
  The previous frames are blurred by this amount for each new frame. This has no effect unless it is positive.

- **Diffuse Amount** (Default: 0, Range: 0 or greater)
  The previous frames are passed through a pixel-diffusion process of this magnitude, for each new frame. This has no effect unless it is positive.

- **Amount Rel** (X & Y, Default: [1 1], Range: 0 or greater)
  The relative amounts of horizontal and vertical blurring and/or diffusing. This has no effect unless Blur Amount or Diffuse Amount are positive.

- **Center** (X & Y, Default: [0 0], Range: any)
  The center position for rotation and scaling, in screen coordinates relative to the center of the frame. This parameter can be adjusted using the Center Widget.

- **Z Dist** (Default: 0.95, Range: 0.001 to 10)
  For each new frame, the 'distance' of the previous frames is scaled by this amount. This causes zooming during the feedback process. Values greater than 1.0 zoom out and make the previous frames smaller, and values less then 1.0 zoom in and enlarge them. This parameter can be adjusted using the Transform Widget.

- **Rotate** (Default: 3, Range: any)
  For each new frame, the amount of rotation in degrees to apply to the previous frames. This parameter can be adjusted using the Transform Widget.

- **Shift** (X & Y, Default: [0 0], Range: any)
  Shifts the previous frames by this amount for each new frame. If this is non-zero the Center location is less meaningful. This parameter can be adjusted using the Transform Widget.

- **Wrap** (X & Y, Popup menu, Default: [ No No ])
  Determines the method for accessing outside the borders of the source image.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less
visible with this method.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Mask Use** (Popup menu, Default: Luma)
  Determines how the Mask input channels are used to make a monochrome mask.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Mask** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Show Center** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Transform** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Z Dist and Rotate parameters.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Shift** (Check-box, Default: off)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

