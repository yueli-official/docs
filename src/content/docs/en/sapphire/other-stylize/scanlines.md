---
title: ScanLines
---

## S_ScanLines

Creates a version of the source clip with a scan line pattern
resembling a color TV monitor. Increase the Add Noise parameter to also add
a grainy effect to the result.

In the Sapphire Stylize effects submenu.

![ScanLines](../_static/ScanLines.jpg)


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

- **Lines Frequency** (Default: 50, Range: 1 or greater)
  The frequency of scan lines on the screen. Increase for more lines, decrease for fewer.

- **Lines Sharpness** (Default: 1, Range: 0 or greater)
  Scales the severity of the lines. Increase for sharper edges, or decrease for a more subtle effect. A sharpness of zero reduces the scan line effect to nothing.

- **Lines Angle** (Default: 0, Range: any)
  The angle in degrees of the scan lines. Set to 90 for vertical lines instead of horizontal. This parameter can be adjusted using the Lines Angle Widget.

- **Lines Shift** (Default: 0, Range: any)
  Offsets the position of the pattern of lines. A value of 1.0 shifts one entire scan line over, giving the same result as 0.

- **Shift Red** (Default: 0, Range: any)
  Shifts the red scan lines by this amount, relative to the other lines. Set the red, green, and blue shifts to -.33, .0, and .33 for an out-of-alignment television set look.

- **Shift Green** (Default: 0, Range: any)
  Shifts the green scan lines by this amount.

- **Shift Blue** (Default: 0, Range: any)
  Shifts the blur scan lines by this amount.

- **Add Noise** (Default: 0, Range: 0 or greater)
  If positive, this much color noise is added to the image.

- **Noise Freq Rel** (Default: 1, Range: 0.01 or greater)
  The frequency of the noise, relative to the frequency of lines. This has no effect unless the Add Noise parameter above is positive.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Scale Color** (Default rgb: [1 1 1])
  Scales the color of the result. For example, if it is yellow [1 1 0], the blue of the result will be 0.

- **Offset** (Default: 0, Range: any)
  Adds this gray value to the result (or subtracts if negative). 0 has no effect, .5 is middle gray, and 1 is white.

- **Gamma** (Default: 1.5, Range: 0.1 or greater)
  Scales the brightness of the image by a curve using this gamma value, allowing adjustment of the middle gray values in the scan lines. This can help make the average brightness of the output match the input.

- **Saturation** (Default: 1, Range: 0 or greater)
  Scales the color saturation. Increase for more intense colors. Set to 0 for monochrome.

- **Smooth Source** (Default: 0, Range: 0 or greater)
  If positive, the source clip is blurred by this amount before being processed.

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

- **Show Lines Angle** (Check-box, Default: off)
  Turns on or off the screen user interface for adjusting the Lines Angle parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

