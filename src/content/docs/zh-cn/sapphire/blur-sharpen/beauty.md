---
title: Beauty
---

## S_Beauty

Applies smoothing, color correction, soft focus, and glow to skin regions. Skin regions are determined
in one of four ways depending on the value of Enable Skin Detection and whether a second input is provided:Enable Skin Detection
OFF
, no second input: effect applies to the entire image.Enable Skin Detection
ON
, no second input: an internal matte is generated using the specified Skin Color,
Luma and Chroma Range parameters.Enable Skin Detection
OFF
, second input provided: The second input is used as an external matte. Effect
applies to bright areas of the matte (to apply to dark areas, see the Invert Matte parameter).Enable Skin Detection
ON
, second input provided: An internal matte is generated and multiplied by the
external matte.

In the Sapphire Blur+Sharpen effects submenu.

![Beauty](../_static/Beauty.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Matte**: Defaults to None. Garbage matte to combine with internal skin detection.


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

- **Enable Skin Detection** (Check-box, Default: on)
  Generate internal matte of skin regions based on Skin Color, Luma and Chroma Range, etc. When disabled, effect applies uniformly to matted input (see description above for full details).

- **Skin Color** (Default rgb: [0.749 0.498 0.345])
  Representative skin color to use during detection. This parameter can be adjusted using the Skin Color Widget.

- **Luma Range** (Default: 0.4, Range: 0 or greater)
  Difference in luma from Skin Color to consider skin.

- **Chroma Range** (Default: 0.2, Range: 0 or greater)
  Difference in chroma from Skin Color to consider skin. This parameter can be adjusted using the Chroma Range Widget.

- **Rel Orange** (Default: 1, Range: 0 to 1)
  Relative amount of orange in Chroma Range. The orange axis is along the human skin tone line common to all the human races, and is also known as the I-line in vectorscope terminology. Reducing this parameter can help eliminate blond or red hair (and similar highlights in other hair colors), some red/orange/yellow sports uniforms, warm backgrounds, etc. from the skin matte.

- **Rel Purple** (Default: 0.3, Range: 0 to 1)
  Relative amount of purple in Chroma Range. The purple axis is perpendicular to the orange axis, and is thus most un-skin-like. In most cases this parameter can be made smaller to eliminate lips, eyes, clothing, and jewelry from the skin matte. Increase this parameter to add purple and green shades to the skin detect matte, for example eye shadow, bad lighting, or alien skin tones (i.e. non-human).

- **Range Softness** (Default: 0.75, Range: 0 to 1)
  Controls the softness of the skin detection matte. A value of one means only pixels that exactly match Skin Color will produce a matte value of one with each other pixel's matte value being proportional to its distance from Skin Color. A value of zero means a hard matte where all pixels within the luma/chroma range of the Skin Color will produce a matte value of one.

- **Clip White** (Default: 1, Range: 0 to 1)
  Skin detection matte values greater than this value will be set to one.

- **Clip Black** (Default: 0, Range: 0 to 1)
  Skin detection matte values less than this value will be set to zero.

- **Post Blur** (Default: 0, Range: 0 or greater)
  Blur the skin detection matte by this amount.

- **Show** (Popup menu, Default: Final)
  Selects the type of output.
  - **Final**: Show the final output.
  - **Skin Detect Matte**: Show the matte generated by the internal skin detector.
  - **With Garbage Matte**: Show the combined result of the input garbage matte and the internal skin detector matte.
  - **Skin**: Show result of applying skin detection matte to source.
  - **Skin with Garbage Matte**: Show the result of applying the combined input garbage matte and the internal
skin detector matte to the Source.

- **Show Color Helper** (Check-box, Default: off)
  Display an interactive overlay to help set Skin Color, Chroma Range, Rel Orange, and Rel Purple. The overlay shows all possible colors that match the brightness (luma) of the Skin Color parameter. Orange is in the upper left corner, and purple is in the upper right (this orientation is similar to a traditional broadcast vectorscope). Colors matching the skin detection algorithm are highlighted. Changing Skin Color will move the highlighted region, adjusting Chroma Range changes the size of the highlighted region, and adjusting Rel Orange/Purple stretches the region along the diagonals of the square.

- **Matte Use** (Popup menu, Default: Luma)
  Determines how the Matte input channels are used to make a monochrome matte.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Matte** (Default: 0, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Suppress BG** (Check-box, Default: off)
  Only apply Beauty to the region specified by Face Center and related params.

- **Face Center** (X & Y, Default: [0 0], Range: any)
  Center position of face region when Suppress BG is enabled. This parameter can be adjusted using the Face Center Widget.

- **Face Softness** (Default: 0.1, Range: 0 or greater)
  Makes the face region softer when Suppress BG is enabled. This will provide a smoother transition from the face region to the background, but also possibly reduce the strength of Beauty in the face region.

- **Face Radius** (Default: 0.4, Range: 0 or greater)
  Size of face region when Suppress BG is enabled. This parameter can be adjusted using the Face Radius Widget.

- **Face Rel Height** (Default: 1.33, Range: 0.05 or greater)
  Relative height of face region when Suppress BG is enabled.

- **Face Rotate** (Default: 0, Range: any)
  Rotation of face region when Suppress BG is enabled. This parameter can be adjusted using the Face Rotate Widget.

- **Show Face Widget** (Check-box, Default: off)
  Display an interactive overlay to assist in placing and sizing the face region.

- **Pore Size** (Default: 0.01, Range: 0 or greater)
  Features smaller than this size (pores, etc.) will be preserved even when blurring.

- **Blur Amount** (Default: 0.056, Range: 0 or greater)
  Scales the width of the blur.

- **Edge Threshold** (Default: 0.1, Range: 0 or greater)
  Color regions separate by an edge larger than this value will not blur into each other.

- **Soften Shadows** (Default: 0.2, Range: -1 to 1)
  Postive values reduce the appearance of shadows, while negative values make shadows more pronounced. Reducing shadows can make the subject look younger, while darkening shadows will make them look older.

- **Shadow Thresh** (Default: 0.6, Range: 0 or greater)
  Dark regions less than this value will be enhanced/reduced by Soften Shadows.

- **Reduce Shine** (Default: 0, Range: 0 to 1)
  Darken bright, shiny areas. The darkening process can lead to a lack of color, use Shine Saturation to bring back a natural skin tone in the affected region.

- **Shine Saturation** (Default: 1, Range: 0 or greater)
  Scales the color saturation in bright regions. Useful for adding a natural skin tone to shiny areas that required darkening.

- **Shine Thresh** (Default: 0.9, Range: 0 or greater)
  Regions brighter than this value will be affected by Reduce Shine.

- **Hue Shift** (Default: 0, Range: any)
  Shifts the hue of the source colors, in revolutions from red to green to blue to red.

- **Saturation** (Default: 1.1, Range: -2 to 8)
  Scales the color saturation of the result. Increase for more intense colors. Set to 0 for monochrome. You can also invert the chroma of the result by making this negative.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Tint** (Default rgb: [1 1 1])
  Scales the result by this color, thus tinting the lighter regions.

- **Soft Focus** (Default: 0, Range: 0 or greater)
  Scales the width of the soft focus blur.

- **Glow Brightness** (Default: 0.1, Range: 0 or greater)
  Scales the brightness of the skin glow.

- **Glow Color** (Default rgb: [1 1 1])
  Scales the color the skin glow.

- **Glow Threshold** (Default: 0.2, Range: 0 or greater)
  Glows are generated from locations in the skin regions that are brighter than this value. A value of 0.9 causes glows at only the brightest spots. A value of 0 causes glows for every non-black area.

- **Glow Width** (Default: 0.1, Range: 0 or greater)
  Scales the skin glow distance.

- **Mix With Source** (Default: 0, Range: 0 to 1)
  Interpolates between the blurred result (0) and the original source (1). 0.1 can give a nice misty effect since it mixes only a little of the source in.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

