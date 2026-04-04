---
title: Swish3D
---

## S_Swish3D

Dissolves between two input clips while performing 3D
moves on each. During the transition the From clip is transformed
by the Zdist, Rotate, Swivel, Tilt, Shift, Scale, and Shear
parameters, and the To clip is transformed by the opposite of these
values. The overall amount of motion for each image can
be scaled by the Rel Amp From and Rel Amp To parameters.

In the Sapphire Transitions effects submenu.

![Swish3D](../_static/Swish3D.jpg)


### Inputs:

- **Foreground**: The current layer. Starts the transition with this clip.

- **Background**: Defaults to None. Ends the transition with this clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: Blur Warp)
  Selects the type of motion blur to apply when moving the From and To clips.
  - **Blur Warp**: Normal motion blur, similar to BlurMotion effect.
  - **Chroma Warp**: Move the color channels by different amounts,
creating a color fringing effect similar to WarpChroma.

- **Transition Dir** (Popup menu, Default: Wipe Off to Bg)
  Selects the direction of the transition.
  - **Wipe Off to Bg**: transitions from the current layer to the Background.
  - **Wipe On from Bg**: transitions from the Background to the current layer.

- **Auto Trans** (Popup YES-NO, Default: No)
  If enabled, a transition is performed automatically between the first and last frames of the layer. If this is off, the transition is performed manually by animating the Swish3 Percent parameter.

- **Wipe Percent** (Default: 0, Range: 0 to 1)
  Auto Trans must be disabled for this parameter to be used. It determines the transition ratio between the From and To inputs, and would normally be animated from 0 to 100 to perform a complete transition. The curve controlling this parameter can be adjusted for more detailed control over the timing of the wipe.

- **Center** (X & Y, Default: [0 0], Range: any)
  The location of the d center in screen coordinates relative to the center of the frame. This parameter can be set by enabling and moving the Center Widget. Note that moving the d center can also cause the d size to change so that the current value of Wipe Amt remains correct.

- **Motion Blur** (Default: 1, Range: 0 or greater)
  Scales the amount of motion blur to use.

- **Z Dist** (Default: 0.5, Range: 0.001 or greater)
  The 'distance' to transform the From clip. Values greater than 1.0 move it farther away and make it smaller. Values less then 1.0 move the image closer and enlarge it. By default, the To clip is also transformed by the opposite of this value.

- **Rotate** (Default: 0, Range: any)
  Rotates by the specified angle in degrees.

- **Swivel** (Default: 0, Range: any)
  Rotates left or right in 3D about a vertical axis.

- **Tilt** (Default: 0, Range: any)
  Rotates up or down in 3D about a horizontal axis. You can use Swivel and Tilt together to rotate about arbitrary diagonal axes.

- **Perspective Amount** (Default: 1, Range: 0.25 to 4)
  Controls the amount of lens telescoping while applying Swivel and Tilt. Increase for more 3D perspective.

- **Shift** (X & Y, Default: [0 0], Range: any)
  Translation of the d pattern.

- **Scale** (Default: 1, Range: 0 to 2)
  Scales the size of the clips.

- **Scale Rel** (X & Y, Default: [1 1], Range: 0 to 2)
  Scales the relative horizontal or vertical size of the clips.

- **Shear** (X & Y, Default: [0 0], Range: any)
  Shears horizontally or vertically.

- **Rel Amp From** (Default: 1, Range: any)
  Scales the amount of transformation applied to the From clip. Set to zero to disable moving the From clip. Make negative to reverse the motion.

- **Rel Amp To** (Default: -1, Range: any)
  Scales the amount of transformation applied to the To clip. By default, the To clip is transformed in the opposite direction of the From clip. Set to zero to disable moving the To clip. Make positive to move the To clip in the same direction as the From clip.

- **Fade** (Popup menu, Default: From and To)
  Determines which clips are faded in or out during the transition.
  - **From and To**: Cross fades both clips during the transition.
  - **Only From**: Fades out the From clip and composites that
over the To clip. This causes the To clip to remain fully opaque in
areas where the From clip does not overlap with it.
  - **Only To**: Fades in the To clip and composites that over the
From clip. This causes the From clip to remain fully opaque in
areas where the To clip does not overlap with it.

- **Fade Mid Time** (Default: 0.5, Range: 0 to 1)
  The midpoint in time of the image dissolve. Decrease for an earlier dissolve or increase for a later dissolve. If this is 1.0 the From clip will remain fully opaque for the entire transition. You can use this in combination with the Combine parameter to create various reveals without fading either clip. For example set Dissolve Mid Time to 1.0, Combine to Fade From, and then Shift and/or Rotate to cause the From clip to move off the screen.

- **Slow In** (Default: 0.5, Range: 0 to 1)
  If positive, causes the transition to start more gradually.

- **Slow Out** (Default: 0.5, Range: 0 to 1)
  If positive, causes the transition to end more gradually.

- **Wrap From** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  Determines the method for accessing outside the borders of the From image.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less
visible with this method.

- **Wrap To** (X & Y, Popup menu, Default: [ Reflect Reflect ])
  Determines the method for accessing outside the borders of the To image.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less
visible with this method.

- **Filter** (Check-box, Default: on)
  If enabled, the image is adaptively filtered when it is resampled. This gives better quality results when the image is warped smaller.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result. This can be animated to brighten the result during the transition, but should typically start and end at 1.0 to avoid any pop at the start or end of the transition.

- **Mid Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result at the middle of the transition by this amount. Automatically ramps to this brightness and then back again during the transition.

- **Steps** (Integer, Default: 8, Range: 3 to 100)
  The number of spectrum samples to include along the path between the From (red) and To (blue) transformations. More steps give a smoother result, but require more time to process.

- **Color1** (Default rgb: [1 0 0])
  The color at the From transformation.

- **Color2** (Default rgb: [0 1 0])
  The color midway between the From and To transformations.

- **Color3** (Default rgb: [0 0 1])
  The color at the To transformation.

- **White Balance** (Check-box, Default: off)
  When enabled, the three colors are adjusted internally so they sum to white. In this case, the colors of unwarped regions are not affected and the average color of the result remains the same.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.
If your image has sharp color changes where the matte
channel also has sharp edges, you may get better results with Normal
mode.

- **Show To Shift** (Check-box, Default: off)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show To Transform** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the To Z Dist and To Rotate parameters.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show From Shift** (Check-box, Default: off)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show From Transform** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Center** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

