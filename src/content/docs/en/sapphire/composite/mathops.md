---
title: MathOps
---

## S_MathOps

Combines two clips using one of a variety of mathematical operations. The
colors of each input can also be adjusted using the lights, darks, and
saturation parameters.

In the Sapphire Composite effects submenu.

![MathOps](../_static/MathOps.jpg)


### Inputs:

- **SourceA**: The current layer. The first input clip to be processed.

- **SourceB**: Defaults to None. The second input clip to be processed.

- **Mask**: Defaults to None. Interpolate between the result and the Source input. White areas use the result of the effect. Black areas use the Source clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Operation** (Popup menu, Default: Add)
  Determines which mathematical operation is applied to combine the pixel colors of the two source inputs.
  - **Add**: A + B.
  - **Subtract**: A - B.
  - **Multiply**: A * B. This can be used as an 'intersection'
operation on matte images. The result only contains white where both
inputs are white.
  - **Divide**: A / B. This can be used to 'un-premultiply' an
image by using its matte as the second input.
  - **Screen**: A + B - AB. This can be useful in combining the
bright areas of two clips. It can also be used as a 'union'
operation on matte images. The result is white where either of the
input images is white.
  - **Average**: (A + B) / 2.
  - **Overlay**: combines A over B with an overlay function.
  - **Minimum**: the smallest value for each color channel of each
pixel. This can also be used as an 'intersection' operation with
slightly different results than Multiply.
  - **Maximum**: the largest value for each color channel of each
pixel. This can also be used as a 'union' operation with slightly
different results than Screen.
  - **Difference**: similar to Subtract but the absolute value of
the result is used, which tends to give more resulting colors in
bounds. This can be used to select the regions of two matte
images where one or the other is white, but not both.
  - **Xor**: performs an 'exclusive-or' operation on the colors of
the source clips. This can also be used to select the regions of two
matte images where one or the other is white, but not both, with
slightly different results than Difference.
  - **Xor Bits**: performs a bitwise exclusive-or on the colors of
the source clips. This can produce some interesting contour effects
although the results are often difficult to predict.
  - **And Bits**: performs a bitwise logical and on the colors of
the source clips. Similar to XorBits but tends to produce darker
results.
  - **Or Bits**: performs a bitwise logical or on the colors of
the source clips. Similar to XorBits but tends to produce brighter
results.
  - **Mod**: gives the remainder after dividing the colors of the
first source clip by the second. Set the A Scale parameter to a
high value for some unusual pixel banding effects.
  - **Round**: the colors of the first source clip are rounded using
the values of the second input as the step size.
  - **Bounce**: similar to Mod but the result contains fewer jagged
edges. Set the A Scale parameter to a high value for some striping
effects.

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

- **Swap Inputs** (Check-box, Default: off)
  If enabled, effectively swaps the A and B Source inputs, and can be helpful for non-commutative operations like subtract. Note that this also causes parameters labeled 'A' to affect the 'B' input instead, and vice versa.

- **A Lights** (Default: 1, Range: 0 or greater)
  Scales the brightness of SourceA before performing the operation.

- **A Darks** (Default: 0, Range: any)
  Offsets the darker regions of the first input before performing the effect. This can be negative to increase contrast.

- **A Saturation** (Default: 1, Range: 0 or greater)
  Adjusts the color intensity of SourceA before performing the operation. 0.0 makes it monochromatic, 1.0 has no effect.

- **B Lights** (Default: 1, Range: 0 or greater)
  Scales the brightness of SourceB before performing the operation.

- **B Darks** (Default: 0, Range: any)
  Offsets the darker regions of the second input before performing the effect. This can be negative to increase contrast.

- **B Saturation** (Default: 1, Range: 0 or greater)
  Adjusts the color intensity of SourceB before performing the operation. 0.0 makes it monochromatic, 1.0 has no effect.

- **Dest Lights** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result after performing the operation.

- **Dest Darks** (Default: 0, Range: any)
  Offsets the darker regions of the result after performing the effect. This can be negative to increase contrast.

- **Dest Saturation** (Default: 1, Range: 0 or greater)
  Scales the color intensity of the result after performing the operation. 0.0 makes it monochromatic, 1.0 has no effect.

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

