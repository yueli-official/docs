---
title: UltraGlow
---

## S_UltraGlow

Generates varieties of glowing light from bright areas
of the source clip. Raise the threshold parameter to produce glows in
fewer areas. Adjust the Width RGB parameters to make glows with
different color falloffs, and adjust the Width XY parameters to make
horizontal or vertical glows. Adjust Glow Falloff and Glow Bias parameters
to control falloff distance and how far the hottest areas extend. Adjust
After Glow parameters to generate a secondary glow on the result of the
primary glow. Optionally enhance the edges or add highlights to the source
clip or combine the result with atmospheric noise.

In the Sapphire Lighting effects submenu.

![UltraGlow](../_static/UltraGlow.jpg)


### Inputs:

- **Source**: The current layer. The input clip that determines the glow locations and colors.

- **Background**: Defaults to None. The clip to combine the glows with. If no background is given, the Source is also used as the Background.

- **Matte**: Defaults to None. If provided, the source glow colors are scaled by this input. A monochrome matte can be used to choose a subset of Source areas that will generate glows. A color matte can be used to selectively adjust the glow colors in different regions. The matte is applied to the source before the glows are generated so it will not clip the resulting glows.


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

- **Brightness** (Default: 1.8, Range: 0 or greater)
  Scales the brightness of all the glows.

- **Color** (Default rgb: [1 1 1])
  Scales the color of the primary glow.

- **Threshold** (Default: 0.4, Range: 0 or greater)
  Glows are generated from locations in the source clip that are brighter than this value. A value of 0.9 causes glows at only the brightest spots. A value of 0 causes glows for every non-black area.

- **Threshold Add Color** (Default rgb: [0 0 0])
  This can be used to raise the threshold on a specific color and thereby reduce the glows generated on areas of the source clip containing that color.

- **Glow Width** (Default: 0.371, Range: 0 or greater)
  Scales the glow distance. This and all the width parameters can be adjusted using the Width Widget. Note that a zero glow width still enhances the bright areas; set the brightness parameter to zero if you want to pass the Source through unchanged.

- **Glow Falloff** (Default: 0.35, Range: -2 to 2)
  Boost or cut the distance that the glow extends.

- **Glow Bias** (Default: 0, Range: -3 to 3)
  Amount to grow the outskirts of the thresholded result, or shrink if negative.

- **Width X** (Default: 1, Range: 0 or greater)
  Scales the horizontal glow width. Set to 0 for vertical only.

- **Width Y** (Default: 1, Range: 0 or greater)
  Scales the vertical glow width. Set to 0 for horizontal only.

- **Width Red** (Default: 1, Range: 0 or greater)
  Scales the red glow width. If the red, green, and blue widths are equal, the glows will match the color of the source clip. If they are not equal, the glows will vary in color with distance.

- **Width Green** (Default: 1, Range: 0 or greater)
  Scales the green glow width.

- **Width Blue** (Default: 1, Range: 0 or greater)
  Scales the blue glow width.

- **Subpixel** (Check-box, Default: on)
  Enables glowing by subpixel widths. Use this for smoother animation of the Width parameters.

- **Show** (Popup menu, Default: Result)
  Selects the type of output
  - **Result**: Shows the final result of combining the glow, source, and background.
  - **Threshold**: Shows the thresholded image that is used to generate the glow.

- **After Glow Width** (Default: 0.808, Range: 0 or greater)
  Scales the glow distance for the secondary glow.

- **After Glow Color** (Default rgb: [1 1 1])
  Scales the color of the secondary glow.

- **After Glow Stretch X** (Default: 0.3, Range: 0 or greater)
  Scales the horizontal secondary glow width.

- **After Glow Stretch Y** (Default: 0.1, Range: 0 or greater)
  Scales the vertical secondary glow width.

- **Horizontal Streaks** (Default: 0.25, Range: 0 or greater)
  Scales the appearance of narrow trails in the horizontal direction.

- **Vertical Streaks** (Default: 0.25, Range: 0 or greater)
  Scales the appearance of narrow trails in the vertical direction.

- **Edge Detect** (Check-box, Default: off)
  Enables edge detection.

- **Edge Combine** (Popup menu, Default: Screen)
  Determines how the detected edges are combined with the Source.
  - **Screen**: detected edges are blended with the Source using a screen operation.
  - **Add**: detected edges are added to the Source.
  - **Edges Only**: gives only the detected edges with no Source.

- **Edge Smooth** (Default: 0, Range: 0 or greater)
  Increase for thicker and smoother edges.

- **Edge Mode** (Popup menu, Default: Reflect)
  Determines the behavior when accessing areas outside the source image.
  - **Transparent**: Areas outside the source image are treated as transparent, which can produce
transparency around the edges of the image.
Select this for fastest rendering.
  - **Reflect**: Reflects the image outside the border.

- **Edge Fill** (Check-box, Default: on)
  Make areas within detected edges opaque

- **Edge Thin** (Default: 0, Range: 0 or greater)
  Subtracts this value from the detected edge result. Increase to remove unwanted noise from minor edges.

- **Atmosphere** (Check-box, Default: off)
  Atmosphere gives the effect of the glow shining through a dusty atmosphere and picking up light or getting shadowed. This parameter adjusts the amount, or amplitude, of the atmospheric effect. Zero gives a smooth glow, higher values give more dusty look.

- **Atmosphere Amp** (Default: 1, Range: 0 or greater)
  Atmosphere gives the effect of the glow shining through a dusty atmosphere and picking up light or getting shadowed. This parameter adjusts the amount, or amplitude, of the atmospheric effect. Zero gives a smooth glow, higher values give more dusty look.

- **Atmosphere Freq** (Default: 11.6, Range: 0.1 to 20)
  Controls the spatial frequency of the atmospheric noise. Turn this up higher to get finer details, turn down for broader overall variation.

- **Atmosphere Detail** (Default: 0.506, Range: 0 to 1)
  Controls the amount of fine detail in the atmosphere simulation. Decrease to get smoother atmosphere, increase for a more crunchy or grainy look.

- **Atmosphere Speed** (Default: 1, Range: any)
  The cloudy noise in the atmosphere evolves over time like real dust clouds; this parameter controls how fast the cloud pattern changes over time. Set to zero for a static pattern.

- **Atmosphere Lights** (Default: 0.5, Range: 0 or greater)
  Scales the atmosphere layer by this value. Increase for a more intense result.

- **Atmosphere Darks** (Default: 0, Range: 0 or greater)
  Adds this gray value to the darker regions of the atmosphere layer. This can be negative to increase contrast.

- **Atmosphere Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator for the atmospheric noise. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Apply Pre-Glow** (Check-box, Default: off)
  Enables combining atmosphere with the Source prior to any glows.

- **Highlights** (Check-box, Default: off)
  Enables highlights using a selected texture pattern.

- **Highlights Texture** (Popup menu, Default: Plasma)
  Selects the texture used for highlights.
  - **Plasma**: Highlights with an electrical plasma texture.
  - **Micro**: Highlights with a magnified rough surface texture.

- **Highlights Freq** (Default: 1.2, Range: 0.01 or greater)
  The spatial frequency of the highlights. Increase to zoom out, decrease to zoom in.

- **Highlights Freq Rel X** (Default: 1, Range: 0.01 or greater)
  The relative horizontal frequency of the highlights. Increase to stretch vertically or decrease to stretch horizontally.

- **Highlights Octaves** (Integer, Default: 4, Range: 1 to 10)
  The number of octaves of highlights to include. Each octave is twice the frequency and half the amplitude of the previous.

- **Highlights Grad X** (Default: 0.1, Range: any)
  Determines the amplitude and direction of a gradient which orients the highlights. Increasing X makes them more vertical.

- **Highlights Grad Y** (Default: 0, Range: any)
  Determines the amplitude and direction of a gradient which orients the highlights. Increasing Y makes them more horizontal.

- **Highlights Layers** (Default: 4.5, Range: 0 or greater)
  The number of layers of highlights. Increase for a more striped effect.

- **Highlights Threshold** (Default: 0.5, Range: 0 or greater)
  Determines the thickness of the highlights. Increase for thinner lines, decrease for thicker and brighter ones.

- **Highlights Speed** (Default: 1, Range: any)
  Phase speed of the highlights. If non-zero, the lines are automatically animated to undulate at this rate.

- **Highlights Details** (Default: 0.43, Range: 0 to 1)
  Increases or decreases the amount of fine detail in the texture. Decrease to get a smoother look, increase to get a more high-frequency, noisy look.

- **Highlights Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the highlights.

- **Highlights Lights** (Default: 1, Range: 0 or greater)
  Scales the highlights layer by this value. Increase for a more intense result.

- **Highlights Darks** (Default: 0, Range: 0 or greater)
  Adds this gray value to the darker regions of the highlights layer. This can be negative to increase contrast.

- **Highlights Blur** (Default: 0, Range: 0 or greater)
  Soften the edges of the highlights.

- **Highlights Combine** (Popup menu, Default: Multiply)
  Determines which blending method is used to combine the highlights with the background.
  - **Multiply**: The default method 'intersects' the highlights with the background.
  - **Highlights Only**: Display only the highlights to suppress outlines.
  - **Dissolve**: Randomly replaces background pixels with highlights.
  - **Screen**: Display the 'union' of the highlights with the background
  - **Overlay**: Combines highlights and background using an overlay function.
  - **Soft Light**: Darkens or lightens the background depending on the highlights.
  - **Hard Light**: Similar to overlay but with highlights and background swapped.
  - **Color Dodge**: Brightens the background depending on the highlights.
  - **Color Burn**: Darkens the background depending on the highlights.
  - **Darken**: The minimum of highlights and background. This can
also be used as an 'intersection' operation with slightly different
results than Multiply.
  - **Lighten**: the maximum of highlights and background. This
can also be used as a 'union' operation with slightly different
results than Screen.
  - **Add**: Adds the highlights to the background.
  - **Subtract**: Subtracts the highlights from the background.
  - **Difference**: Similar to Subtract but the absolute value of
the result is used, which tends to give more resulting colors in
bounds.
  - **Exclusion**: Similar to Difference but with smoother results.
  - **Hue**: Combines the hue of the highlights with the saturation
and luminance of the background.
  - **Saturation**: Combines the saturation of the highlights with
the hue and luminance of the background.
  - **Chroma**: Combines the hue and saturation of the highlights
with the luminance of the background.
  - **Luminance**: Combines the luminance of the highlights with
the hue and saturation of the background.
  - **Linear Dodge**: Adds highlights and background and clamps the
result at white.
  - **Linear Burn**: Adds highlights and background but offsets to
make the result darker.
  - **Linear Light**: Performs a linear burn or linear dodge
depending on if the highlights are more or less than 50 percent gray.
  - **Vivid Light**: Performs a color burn or color dodge
depending on if the highlights are more or less than 50 percent gray.
  - **Pin Light**: Performs a lighten or darken depending on if the
highlights are more or less than 50 percent gray.

- **Combine** (Popup menu, Default: Screen)
  Determines how the glow is combined with the Source or Background. This parameter has no effect if Light BG is set to 1.
  - **Mult**: the source or background is multiplied by the glow.
  - **Add**: the glow is added to the source or background.
  - **Screen**: the glow is blended with the source or background using a screen operation.
  - **Difference**: the result is the difference between the glow and the
source or background.
  - **Overlay**: the glow is combined with the source or background using an overlay function.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the glows. The maximum of the red, green, and blue glow brightness is scaled by this value and combined with the background Alpha at each pixel.

- **Glow From Alpha** (Default: 0, Range: 0 to 1)
  Set to 1 to generate glows from the alpha channel of the source input instead of the RGB channels. In this case the glows will not pick up color from the source and will typically be brighter. Values between 0 and 1 interpolate between using the RGB and the Alpha.

- **Glow Under Source** (Default: 0, Range: 0 to 1)
  Set to 1 to composite the Source input over the glows.

- **Light Background** (Default: 0, Range: 0 to 1)
  Increase this to give a look of the glow casting light onto the background image. To see this more clearly you can also lower the Background Scale parameter or raise the Brightness parameter.

- **Source Opacity** (Default: 1, Range: 0 to 1)
  Scales the opacity of the Source input when combined with the glows. This does not affect the generation of the glows themselves.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the background. This parameter only has an effect if the background input is provided, and is visible due to a partially transparent Source image or a reduced Source Opacity parameter value.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Expand Borders** (Check-box, Default: on)
  If enabled, transparent borders are added to the input image before processing. This allows the result to include soft edges beyond the original image size. When off, the effect only occurs within the frame and the result will retain an edge at the borders.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Glow Width** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Glow Width parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

