---
title: RackDefocus
---

## S_RackDefocus

Generates a defocused version of the source clip
using a 'circle of confusion' convolution. This effect is often
preferable to a gaussian blur for simulating a real defocused camera
lens, because bright spots can be defocused into clean shapes
instead of being smoothed away. The iris shape can be controlled
using Points, Pointiness and Rotate, and the Use Gamma parameter can
adjust the relative brightness of the blurred highlights.

In the Sapphire Blur+Sharpen effects submenu.

![RackDefocus](../_static/RackDefocus.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Mask**: Defaults to None. Interpolate between the result and the Source input. White areas use the result of the effect. Black areas use the Source clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: Defocus Color)
  Selects between full color or monochrome defocus.
  - **Defocus Color**: defocuses all channels of the source input.
  - **Defocus Mono**: makes the source monochrome and then
applies the defocus (faster).

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

- **Defocus Width** (Default: 0.088, Range: 0 or greater)
  The width of the defocus. This parameter can be adjusted using the Defocus Width Widget.

- **Rel Height** (Default: 1, Range: 0.01 or greater)
  The relative height of the iris shape. If it is not 1, circles become ellipses, etc.

- **Shape** (Popup menu, Default: Circle)
  Determines the shape of the simulated camera iris.
  - **Circle**: round.
  - **3 sides**: triangle.
  - **4 sides**: square.
  - **5 sides**: pentagon.
  - **6 sides**: hexagon.
  - **7 sides**: etc.

- **Show Shape** (Check-box, Default: off)
  Show the iris shape instead of the defocused image.

- **Roundness** (Default: 0, Range: any)
  Modifies the shape of the simulated camera iris. A value of 1 produces a circle; 0 gives a flat-sided polygon with a number of sides given by the Shape parameter. Less than 0 causes the sides to squeeze inward giving a star shape, while a value greater than 1 causes the corners to squeeze inward, giving a flowery shape. Has no effect if the Shape is set to Circle.

- **Rotate** (Default: 0, Range: any)
  Rotates the iris shape.

- **Bokeh** (Default: 0, Range: any)
  Softens the outer edge of the iris shape, which gives a softer look to the defocused highlights. A negative value darkens the center of the iris shape, producing a ring-like defocus shape.

- **Lens Noise** (Default: 0, Range: 0 or greater)
  Increase to add noise to the iris shape, dirtying up the defocus a little. Can make the result more realistic. Turn up past 1 for a more stylistic result.

- **Noise Freq** (Default: 40, Range: 0.01 or greater)
  The frequency of the added noise. Ignored if Lens Noise is zero.

- **Noise Freq Rel X** (Default: 1, Range: 0.01 or greater)
  The relative horizontal frequency of the added iris noise. Increase to stretch it vertically or decrease to stretch it horizontally.

- **Noise Seed** (Default: 0.123, Range: 0 or greater)
  The seed value for the added noise. To make the noise appear different on each frame, animate this to be different on each frame. The actual value doesn't matter; only that it's different.

- **Gauss Blur** (Default: 0, Range: 0 or greater)
  If positive, a gaussian blur is also applied which smooths out the edges of the shapes. This might also darken the highlights because Gamma is not considered in the gaussian blur.

- **Use Gamma** (Default: 1, Range: 0.1 or greater)
  Values above 1 cause highlights in the source clip to keep their brightness after the defocus is applied.

- **Boost Highlights** (Default: 0, Range: 0 or greater)
  The amount to increase the luma of the highlights in the source clip. Increase this parameter to blow out the highlights without affecting the darks or mid-tones.

- **Hilight Threshold** (Default: 0.9, Range: 0 or greater)
  The minimum luma value for highlights. Pixels brighter than this will be brightened according to the Boost Highlights parameter.

- **Chroma Distort** (Default: 0, Range: any)
  Adds some chromatic aberration around the edges of the image; red and blue wavelengths of light refract differently in real lenses, producing fringes of color where the rays strike the lens at oblique angles.

- **Color Fringing** (Default: 0, Range: any)
  Color Fringing produces rings of color around every object in the image by varying the focal distance for each color channel. It gives a different style of chromatic aberration from Chroma Distort because it's not just in the image corners.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Offset Darks** (Default: 0, Range: any)
  Adds this gray value to the darker regions of the result. This can be negative to increase contrast.

- **Mix With Source** (Default: 0, Range: 0 to 1)
  Interpolates between the defocused result and the original source. Set this to 1 for the original source.

- **Edge Mode** (Popup menu, Default: Reflect)
  Determines the behavior when accessing areas outside the source image.
  - **Transparent**: Areas outside the source image are treated as transparent, which can produce
transparency around the edges of the image.
Select this for fastest rendering.
  - **Repeat**: Repeats the last pixel outside the border of the image.
  - **Reflect**: Reflects the image outside the border.

- **Soft Borders** (Check-box, Default: off)
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

- **Mask Use** (Popup menu, Default: Luma)
  Determines how the Mask input channels are used to make a monochrome mask.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Mask** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Show Defocus Width** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Defocus Width parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

