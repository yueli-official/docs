---
title: BokehLights
---

## S_BokehLights

Generates random, defocused lights that move around the screen.

In the Sapphire Lighting effects submenu.

![BokehLights](../_static/BokehLights.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Matte**: Defaults to None. The color of each light is scaled by the color of this clip at the center of the light. A black and white mask can be used to create lights that are obscured by foreground objects. A color mask will colorize the lights, which can give the appearance of the lights passing behind a partially-transparent object.


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

- **Brightness** (Default: 0.5, Range: 0 or greater)
  The overall brightness of the lights.

- **Color** (Default rgb: [1 1 0.3])
  The overall color of the lights.

- **Vary Hue** (Default: 0, Range: 0 to 1)
  Randomly varies the hue of each light.

- **Vary Saturation** (Default: 0, Range: 0 to 1)
  Randomly varies the saturation of each light.

- **Vary Brightness** (Default: 1, Range: 0 to 1)
  Randomly varies the brightness of each light.

- **Size** (Default: 0.3, Range: 0 or greater)
  The overall size of the defocused lights. This parameter can be adjusted using the Size Widget.

- **Rel Height** (Default: 1, Range: 0.01 or greater)
  The relative height of the iris shape. If it is not 1, circles become ellipses, etc.

- **Vary Size** (Default: 0.2, Range: 0 or greater)
  Randomly varies the size of the lights, to simulate lights that are at different distances from the camera.

- **Softness** (Default: 0.01, Range: 0.001 or greater)
  The softness of the light sources. Increase this for blurrier lights.

- **Lights** (Integer, Default: 30, Range: 0 or greater)
  The number of lights.

- **Drift Speed** (Default: 0.2, Range: 0 or greater)
  The speed at which lights move around the screen.

- **Drift Distance** (Default: 0.5, Range: 0 or greater)
  The maximum distance that each light will move.

- **Drift Size Speed** (Default: 0.1, Range: 0 or greater)
  The speed at which lights change their size.

- **Drift Size Distance** (Default: 0.3, Range: 0 or greater)
  The maximum amount by which each light's size will change.

- **Drift Smoothness** (Default: 0.65, Range: 0 to 1)
  Controls the amount of high frequency variation in each light's motion. Increase this for a gentle drifting motion. Decrease for a jerky shaking motion.

- **Shift Speed** (X & Y, Default: [0 0], Range: any)
  Translation speed of the lights. If non-zero, the result is automatically animated to shift at this rate. The result of animated Speed values may not be intuitive, so for variable speed motion it is usually best to set this to 0 and animate the Shift Start values instead.

- **Shift Start** (X & Y, Default: [0 0], Range: any)
  Translation offset of the lights.

- **Use Source Color** (Default: 0.25, Range: 0 to 1)
  Scales the lights by a smoothed version of the Source clip. Increase this to help the lights blend with the background.

- **Smooth Source Color** (Default: 0.4, Range: 0 or greater)
  The amount to blur the Source clip before scaling the lights. Has no effect if Use Source Color is zero.

- **Shape** (Popup menu, Default: 7 sides)
  Determines the shape of the simulated camera iris.
  - **Circle**: round.
  - **3 sides**: triangle.
  - **4 sides**: square.
  - **5 sides**: pentagon.
  - **6 sides**: hexagon.
  - **7 sides**: etc.

- **Roundness** (Default: 0.3, Range: any)
  Modifies the shape of the simulated camera iris. A value of 1 produces a circle; 0 gives a flat-sided polygon with a number of sides given by the Shape parameter. Less than 0 causes the sides to squeeze inward giving a star shape, while a value greater than 1 causes the corners to squeeze inward, giving a flowery shape. Has no effect if the Shape is set to Circle.

- **Rotate** (Default: 0, Range: any)
  Rotates the iris shape.

- **Bokeh** (Default: 0.5, Range: any)
  Softens the outer edge of the iris shape, which gives a softer look to the defocused highlights. A negative value darkens the center of the iris shape, producing a ring-like defocus shape.

- **Lens Noise** (Default: 0.5, Range: 0 or greater)
  Increase to add noise to the iris shape, dirtying up the defocus a little. Can make the result more realistic. Turn up past 1 for a more stylistic result.

- **Noise Freq** (Default: 20, Range: 0.01 or greater)
  The frequency of the added noise. Ignored if Lens Noise is zero.

- **Noise Freq Rel X** (Default: 1, Range: 0.01 or greater)
  The relative horizontal frequency of the added iris noise. Increase to stretch it vertically or decrease to stretch it horizontally.

- **Chroma Distort** (Default: 0.05, Range: any)
  Adds some chromatic aberration around the edges of the image; red and blue wavelengths of light refract differently in real lenses, producing fringes of color where the rays strike the lens at oblique angles.

- **Color Fringing** (Default: 0, Range: any)
  Color Fringing produces rings of color around every object in the image by varying the focal distance for each color channel. It gives a different style of chromatic aberration from Chroma Distort because it's not just in the image corners.

- **Flicker Amp** (Default: 0.2, Range: 0 or greater)
  The amount of random flickering of the lights.

- **Flicker Speed** (Default: 0.5, Range: 0 or greater)
  The speed of random flickering.

- **Flicker Randomness** (Default: 0.7, Range: 0 to 1)
  Controls the variability of the flicker. When set to zero, the lights will flicker constantly, with a small amount of random variation. At higher values, the flickering will have longer steady spells, with the occasional large spike.

- **Seed** (Default: 0.123, Range: 0 or greater)
  Initializes the random number generator for light positioning, size, and color variation. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Combine** (Popup menu, Default: Screen)
  Determines how the lights are combined with the Source clip.
  - **Screen**: performs a blend function which can help prevent
overly bright results.
  - **Add**: the lights are added to the source.
  - **Lights Only**: gives only the lights with no background.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the lights. The maximum of the red, green, and blue light brightness is scaled by this value and combined with the Source Alpha at each pixel.

- **Blur Matte** (Default: 0, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Matte Type** (Popup menu, Default: Luma)
  This setting is ignored unless the Mask input is provided.
  - **Luma**: uses the luminance of the Mask input to scale the brightness of the lights.
  - **Color**: uses the RGB channels of the Mask input to scale the colors of the lgiths.
  - **Alpha**: uses the alpha channel of the Mask input to scale the
brightness of the lights.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Size** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Size parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

