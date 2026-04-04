---
title: ParallaxStrips
---

## S_ParallaxStrips

Applies a collection of 3d refracting glass strips to
break up the image. The image is shifted within each strip, and the strips
move over time. The strips gradually fade in or out, so the transition to
the source is seamless.

In the Sapphire Distort effects submenu.

![ParallaxStrips](../_static/ParallaxStrips.jpg)


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

- **Mode** (Popup menu, Default: Automatic)
  Sets whether the effect evolves automatically with time, or can be controlled manually.
  - **Automatic**: Automatically move the strips according to the time
within the clip. Manual Amount is ignored in this mode.
  - **Manual**: Use Manual Amount to specify control the evolution of the effect.
0 is the beginning, and 1 is the end.

- **Manual Amount** (Default: 1, Range: 0 or greater)
  In Manual mode, this controls the strength and evolution of the effect. With Fade:End selected, 1 puts the strips at their starting position, with the most shift (i.e. the effect is at its strongest). 0 gives the original source: it puts the strips at their end position, and they become invisible because the amount they shift the underlying image goes to zero. With Fade:Start, the values are reversed, so 0 is strongest and 1 fades out.

- **Ensure Full Coverage** (Push-button)
  Pressing this button adjusts the number of strips to the minimum needed to fully cover the last frame.

- **N Strips** (Integer, Default: 50, Range: 1 to 1000)
  Number of refracting strips to apply. The strips are positioned randomly all over the image.

- **Size** (Default: 0.35, Range: 0 or greater)
  Size of the strips, in image-widths.

- **Rel Height** (Default: 0.3, Range: 0.001 or greater)
  Height of the strips, relative to their width. Increase to make the strips taller.

- **Size Vary** (X & Y, Default: [0.1 0.1], Range: 0 to 1)
  Increase to make each strip randomly larger or smaller.

- **Angle** (Default: 0, Range: any)
  Angle of the strips; 0 is horizontal. The strips move along their angle, and also shift the image along the same angle.

- **Depth** (Default: 2, Range: 0 or greater)
  Make the frontmost strips larger and move faster, so it appears they're in front, giving a 3d look.

- **Strip Speed** (Default: 0.4, Range: any)
  Sets how fast the strips move along their major axis. Note that this doesn't affect how the image refracts, or shifts, within the strip, just how fast the strip itself moves.

- **Strip Speed Vary** (Default: 0, Range: 0 or greater)
  Increase to give each strip a bit of randomness in its speed.

- **Shift Amount** (Default: 0.6, Range: any)
  Sets how much the image shifts, or refracts, within each strip. Shifting is always along the major axis of the strip. As the effect progresses, the shift amount progressively goes to zero, seamlessly transitioning to the original clip. See the Fade and Slow Fade params for details.

- **Shift Vary** (Default: 0, Range: 0 or greater)
  Increase to make the amount of shift in each strip more random.

- **All Strips Shift** (X & Y, Default: [0 0], Range: any)
  Move all strips around on the screen.

- **Z Dist** (Default: 1, Range: 0.001 or greater)
  Zoom in or out on the source image, before applying the parallax strips.

- **Show** (Popup menu, Default: Result)
  Show the effect result, or the strips themselves, which is useful during effect setup.
  - **Result**: Show the result of the effect.
  - **Strips Over Source**: Show each strip as a gray rectangle, with brightness set by depth.
Uncovered areas show the source image.
  - **Strips Over Black**: Show each strip as a gray rectangle, with brightness set by depth.
Uncovered areas show as black.

- **Slow Fade** (Default: 0.9, Range: 0 to 2)
  Increase to make the fade in our out slower. Set to 0 for a linear fade.

- **Full Height** (Check-box, Default: off)
  Turn on to make the strips always full height, ignoring Rel Height. This can make a nice sliding strip effect.

- **Fade** (Popup menu, Default: End)
  The effect fades out at one end, to allow a seamless start or end.
  - **Start**: Make the effect fade in slowly at the start.
  - **End**: Make the effect fade out slowly at the end.

- **Wrap** (Popup menu, Default: Reflect)
  Determines the method for accessing outside the borders of the source image.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less
visible with this method.

- **Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Flip Tiles** (Check-box, Default: off)
  Flips tiles vertically if needed to achieve a consistent look.

- **Mask Use** (Popup menu, Default: Luma)
  Determines how the Mask input channels are used to make a monochrome mask.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Mask** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

