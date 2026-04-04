---
title: Aurora
---

## S_Aurora

Generates a two colored swirl of light along a user controlled spline
reminiscent of the Aurora Borealis (Northern Lights).

In the Sapphire Render effects submenu.

![Aurora](../_static/Aurora.jpg)


### Inputs:

- **Background**: The current layer. The clip to use as background.

- **Mask**: Defaults to None. Defines the area where aurora should be rendered.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: Aurora)
  Selects between 2D and 3D modes.
  - **Aurora**: creates an aurora along a spline.
  - **Follow Path**: creates an aurora along an AE Path.

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

- **Start** (X & Y, Default: [-0.7 0.4], Range: any)
  The starting point of the Aurora.

- **Start Uses Mocha** (Check-box, Default: off)
  Controls whether the staring point is controlled by the Start parameter or follows the Start point track from Mocha.

- **Smooth Start Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Point 1 Enable** (Check-box, Default: off)
  Turns on or off the first control point.

- **Control Point 1** (X & Y, Default: [-0.33 0.4], Range: any)
  First spline control point.

- **Point1 Uses Mocha** (Check-box, Default: off)
  Controls whether point 1 is controlled by the Control Point 1 parameter or follows the Control Point 1 track from Mocha.

- **Smooth Point1 Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Point 2 Enable** (Check-box, Default: on)
  Turns on or off the second control point.

- **Control Point 2** (X & Y, Default: [0.04 0.2], Range: any)
  Second spline control point.

- **Point2 Uses Mocha** (Check-box, Default: off)
  Controls whether point 2 is controlled by the Control Point 2 parameter or follows the Control Point 2 track from Mocha.

- **Smooth Point2 Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Point 3 Enable** (Check-box, Default: on)
  Turns on or off the third control point.

- **Control Point 3** (X & Y, Default: [-0.1 -0.4], Range: any)
  Third spline control point.

- **Point3 Uses Mocha** (Check-box, Default: off)
  Controls whether point 3 is controlled by the Control Point 3 parameter or follows the Control Point 3 track from Mocha.

- **Smooth Point3 Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Point 4 Enable** (Check-box, Default: off)
  Turns on or off the fourth control point.

- **Control Point 4** (X & Y, Default: [0.45 -0.33], Range: any)
  Fourth spline control point.

- **Point4 Uses Mocha** (Check-box, Default: off)
  Controls whether point 4 is controlled by the Control Point 4 parameter or follows the Control Point 4 track from Mocha.

- **Smooth Point4 Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **End** (X & Y, Default: [0.8 -0.5], Range: any)
  The ending point of the Aurora.

- **End Uses Mocha** (Check-box, Default: off)
  Controls whether the end point is controlled by the End parameter or follows the End point track from Mocha.

- **Smooth End Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Path To Follow** (Default: 0, Range: 0 or greater)
  AE path for the Aurora to follow.

- **Start Color** (Default rgb: [0 0.8 0.3])
  Sets the color at the starting control point.

- **End Color** (Default rgb: [0.6 0.2 0.8])
  Sets the color at the ending control point.

- **Color Phase** (Default: 0, Range: 0 or greater)
  Adjusts the phase of the gradient between Start Color and End Color. Use this to animate the movement of the colors along the Aurora.

- **Stroke Size** (Default: 0.25, Range: 0 or greater)
  Influences the width of the Aurora along the spline. This parameter controls the size of the underlying color gradient before it's distorted.

- **Brightness** (Default: 0.3, Range: 0 or greater)
  Scales the brightness of the Aurora.

- **Softness** (Default: 0.075, Range: 0 or greater)
  The amount of blur applied to the Aurora. Set to 0 to get a colored point cloud.

- **Softness Rel Y** (Default: 10, Range: 0 or greater)
  The relative vertical amount of softness.

- **Swirl Complexity** (Integer, Default: 3, Range: 1 or greater)
  Specifies how many layers should be rendered in the Aurora. The more layers rendered, the more complex pattern generated along the spline.

- **Swirl Magnitude** (Default: 0.7, Range: 0 or greater)
  The magnitude or amplitude of the swirls along the spline. Setting this to 0 will render a color gradient along the spline.

- **Magnitude Rel Y** (Default: 1.25, Range: 0 or greater)
  The relative vertical magnitude of the swirls.

- **Swirl Frequency** (Default: 2, Range: 0 to 50)
  The frequency of the swirls along the spline.

- **Frequency Rel Y** (Default: 10, Range: 0 or greater)
  The relative vertical frequency along the spline.

- **Swirl Speed X** (Default: 0.1, Range: any)
  The speed at which the swirls move horizontally.

- **Swirl Speed Y** (Default: 0, Range: any)
  The speed at which the swirls move vertically.

- **Light Brightness** (Default: 1, Range: 0 or greater)
  Lights a circular area of the Aurora. Set to 0 to disable the light. Increase value to increase the intensity of the light.

- **Light Pos** (X & Y, Default: [-0.3 -0.5], Range: any)
  The position of the center of the light.

- **Light Color** (Default rgb: [1 1 1])
  The color of the light.

- **Ambient Light** (Default: 1, Range: 0 or greater)
  The level of illumination outside the light.

- **Light Radius** (Default: 1, Range: 0 or greater)
  Distance from the center of the light to the edge of the brightest section.

- **Light Softness** (Default: 2, Range: 0 or greater)
  How quickly the edge of the light should taper off of to darkness.

- **Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the Background input.

- **Combine** (Popup menu, Default: Screen)
  Determines how the Aurora is combined with the source image.
  - **Screen**: the Aurora is blended with the source using a screen operation.
  - **Add**: the Aurora is added to the source.
  - **Aurora Only**: gives the Aurora with no source.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the aurora. The maximum of the red, green, and blue light leak brightness is scaled by this value and combined with the Background Alpha at each pixel.

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

- **Show Spline** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Start parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Light Pos** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Light Pos parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

