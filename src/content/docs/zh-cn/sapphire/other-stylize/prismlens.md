---
title: PrismLens
---

## S_PrismLens

Emulates recording footage through prismatic lenses with various different shapes.

In the Sapphire Stylize effects submenu.
### Inputs:

- **Source**: The current layer. The input clip that determines the prism locations and colors.

- **Reflection**: Defaults to None. The clip to use for the reflection. If no input is connected, the source will be used for the reflection.

- **Mask**: Defaults to None. If provided, the mask may control one of two things. By default, the mask will control the area affected by PrismLens. If 'Use Mask For' is changed to 'Shape', then the mask will be used on calculating the shape used for the Reflections.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: PrismLensSingle)
  The input clip that determines the prism locations and colors.
  - **PrismLensSingle**: Generates a single reflection.
  - **PrismLensLinear**: Generates a line of reflections similar to what would be generated from a straight prism.
  - **PrismLensRadial**: Generates several circles of reflections similar to what would be generated from a round prism.

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

- **Use Mask For** (Popup menu, Default: Effect)
  Specify whether the mask should be used to control the shape of the reflection or the region affected by the output of the effect.
  - **Shape**: uses the mask for the reflection shape.
  - **Effect**: uses the mask for the reflection output.

- **Shape Source** (Popup menu, Default: Shape Only)
  What input should be used to specify the reflection shape.
  - **Shape Only**: use the shape parameters and ignore the mask inputs.
  - **Shape and Mask**: combine the shape generated from the parameters and the mask inputs.
  - **Mask Only**: use the mask inputs to control the reflection shape and ignore the shape parameters.
  - **None**: don't restrict the borders of the reflection.

- **Show** (Popup menu, Default: Result)
  Selects the type of output
  - **Result**: show the final result.
  - **Source**: show the input image being used as the source.
  - **Reflection**: show the input image being used for the reflection.
  - **Shape Mask**: show the results of the shape parameters.
  - **All Masks**: show all masks combined.
  - **Gradient**: show the gradient alone before applying to the reflection.
  - **Single Reflection**: show a single reflection after its extracted.
  - **All Reflections**: show the reflection with no post processing.
  - **Final Reflections**: show the reflection with all post processing before the final composite step.
  - **Lights Only**: show the lights without the source or reflection.

- **Combine With Source** (Popup menu, Default: Composite)
  Determines how the reflection is combined with the Source.
  - **Composite**: composites the reflection over the source clip.
  - **Add**: the reflection color is added to the source clip. This
will have no effect if the vignette color is black.
  - **Screen**: the reflection color is combined with the source
clip using a screen operation. This will have no effect if the
reflection color is black.
  - **Reflections Only**: shows the reflection pattern without the
source clip. The output will be white where the amount of reflections
are greatest (e.g. where the source clip would be darkened
completely).

- **Geometry** (Popup menu, Default: Polygon)
  Specifies which shape to use for the reflection.
  - **Polygon**: set reflection to be a polygonal shape.
  - **Ellipse**: set the reflection to be an ellipse.

- **Sides** (Integer, Default: 4, Range: 3 or greater)
  The number of points in the reflection shape. Unless Pointiness is zero, the shape will have this many points around the edge.

- **Size** (Default: 0.724, Range: 0 or greater)
  The overall size of a single reflection.

- **Rel Width** (Default: 0.42, Range: 0 or greater)
  Increase to make the reflection wider.

- **Position** (X & Y, Default: [0.227 -0.024], Range: any)
  Location of the center of the reflection or reflection pattern.

- **Position Uses Mocha** (Check-box, Default: off)
  Use mocha to control the position of the reflection.

- **Position Mocha Shift** (X & Y, Default: [0 0], Range: any)
  Offset the position data out of mocha.

- **Rotate** (Default: -33, Range: any)
  Rotates the whole reflection around its center.

- **Softness** (Default: 0.2, Range: 0 or greater)
  Amount to blur the reflection by.

- **Copies** (Integer, Default: 5, Range: 1 to 10)
  Number of reflections in each line.

- **Angle** (Default: 0, Range: any)
  Angle of the line of reflections.

- **Distance** (Default: 0.2, Range: 0.01 to 1)
  Distance between reflections.

- **Start Offset** (Default: 0.2, Range: 0 to 1)
  Distance between center and first reflection.

- **Bias** (Default: 0, Range: -1 to 1)
  Size of spacing between reflections on the left and right lines. 0 means the spacing is equal, wheras a negative value will skew the spacing towards one side of the center and a positive value will skew the other direction.

- **Opacity Fade** (Default: 0, Range: 0 to 1)
  Scale the transparency of the reflection as they get farther from the center. The larger the value, the more transparent the reflection.

- **Size Fade** (Default: 0, Range: -3 to 3)
  Scale the size of the reflections as they get farther from the center. A positive value causes the size to decrease and a negative value causes the size to increase.

- **Reflect Both Directions** (Check-box, Default: on)
  Whether to draw a line of reflections on both sides of the center.

- **Radial Copies** (Integer, Default: 6, Range: 0 to 10)
  How many reflections should be generated in a single ring.

- **Start Angle** (Default: 15, Range: any)
  Angle of the first reflection in the rings.

- **Total Angle** (Default: 360, Range: 0 to 360)
  The total angle coverage of the reflection rings.

- **Rows** (Integer, Default: 1, Range: 1 to 3)
  How many rings to generate around the center of the reflection.

- **Row Distance** (Default: 0.2, Range: 0 to 1)
  The distance between the rows of reflections.

- **Row Start Offset** (Default: 0.2, Range: 0 to 1)
  The distance between the center and the first row of reflections.

- **Radial Opacity Fade** (Default: 0, Range: 0 to 1)
  Change the transparency of the reflections around the circle. A positive value makes the size shrink and a negative value makes the size grow.

- **Radial Size Fade** (Default: 0, Range: -3 to 3)
  Change the size of the reflections around the circle. A positive value makes the size shrink and a negative value makes the size grow.

- **Row Opacity Fade** (Default: 0, Range: 0 to 1)
  Change the transparency of the reflections radiating out from the center. A positive value makes the size shrink and a negative value makes the size grow.

- **Row Size Fade** (Default: 0, Range: -3 to 3)
  Change the size of the reflections radiating out from the center. A positive value makes the size shrink and a negative value makes the size grow.

- **Reflection Opacity** (Default: 0.9, Range: 0 to 1)
  Controls the transparency of each individual reflection before combined with the other reflections.

- **Source Position** (X & Y, Default: [-0.128 0.04], Range: any)
  Location of the source to use for the reflection image. Note, this should be used to specify the center of the reflection source, regardless of what method is used to specify the Reflection shape.

- **Source Pos Uses Mocha** (Check-box, Default: off)
  Use mocha to control the position to extract the reflection from the reflection input.

- **Source Pos Mocha Shift** (X & Y, Default: [0 0], Range: any)
  Offset the source position data out of mocha.

- **Source Z Dist** (Default: 1.05, Range: 0 or greater)
  Scales the 'distance' of the reflection. Values greater than 1.0 move it farther away and make it smaller. Values less then 1.0 move the image closer and enlarge it. Note that Scale X and Y also scale the size of the image, but in an inverse way and on each axis.

- **Source Rotate** (Default: -2, Range: any)
  Rotates the reflection by the specified angle in degrees.

- **Source Tilt** (Default: -10, Range: any)
  Rotates the reflection up or down in 3D about the horizontal axis. You can use Swivel and Tilt together to rotate about arbitrary diagonal axes.

- **Source Swivel** (Default: -12, Range: any)
  Rotates left or right in 3D about a vertical axis.

- **Flip Horizontal** (Check-box, Default: off)
  Flip the source within the reflection horizontally.

- **Flip Vertical** (Check-box, Default: off)
  Flip the source within the reflection vertically.

- **Blur Source** (Default: 0.0084, Range: 0 or greater)
  Blur the source within the reflection only.

- **Source Threshold** (Default: 0, Range: 0 or greater)
  Reflection is generated from locations in the Reflection input that are greater than the source threshold.

- **Threshold Softness** (Default: 0.05, Range: 0 or greater)
  Blur the reflection after it has been thresholded for softer edges.

- **Source Wrap** (Popup menu, Default: Reflect)
  Determines the method for accessing outside the borders of the source image.
  - **No**: gives black beyond the borders.
  - **Tile**: repeats a copy of the image.
  - **Reflect**: repeats a mirrored copy. Edges are often less
visible with this method.

- **CC Mix With Source** (Default: 0, Range: 0 to 1)
  Interpolates between the color corrected reflection and the original reflection.

- **CC Brightness** (Default: 1, Range: 0 or greater)
  Scales the color saturation of the reflection. Increase for more intense colors. Set to 0 for monochrome. You can also invert the chroma of the result by making this negative.

- **Saturation** (Default: 1, Range: -2 to 8)
  Scales the color saturation of the reflection. Increase for more intense colors. Set to 0 for monochrome. You can also invert the chroma of the result by making this negative.

- **Hue Shift** (Default: 0, Range: any)
  Shifts the hue of the reflection, in revolutions from red to green to blue to red.

- **Offset Darks** (Default: 0, Range: -8 to 2)
  Adds this gray value to the darker regions of the result. This can be negative to increase contrast.

- **Amount** (Default: 0.05, Range: 0 or greater)
  Adjusts the overall amount of warping by scaling the transformations. Setting this to zero disables both transforms and leaves the image unchanged.

- **Chromatic Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the gradient image (both the Start Color and End Color).

- **Chromatic Steps** (Integer, Default: 12, Range: 3 to 100)
  The number of spectrum samples to include in the chromatic abberation on the reflection. More steps give a smoother result, but require more time to process.

- **Chromatic Shift** (X & Y, Default: [0 0], Range: any)
  Controls the direction of the chromatic abberation.

- **Chromatic Rotate** (Default: 0, Range: any)
  The rotation angle of the chromatic abberation.

- **Chromatic Z Dist** (Default: 1, Range: 0.001 or greater)
  The distance or scale of the chromatic abberation.

- **Chromatic Color1** (Default rgb: [1 0 0])
  The color at the start point in the chromatic distortion.

- **Chromatic Color2** (Default rgb: [0 1 0])
  The color at the midway point between the start and end chromatic distortion.

- **Chromatic Color3** (Default rgb: [0 0 1])
  The color at the end point in the chromatic distortion.

- **Glow Brightness** (Default: 1.56, Range: 0 or greater)
  Scales the brightness of the gradient image (both the Start Color and End Color).

- **Glow Threshold** (Default: 0.6, Range: 0 or greater)
  Prisms are generated from locations in the source clip that are brighter than this value. A value of 0.9 causes prisms at only the brightest spots. A value of 0 causes prisms for every non-black area.

- **Glow Width** (Default: 0.2, Range: 0 or greater)
  Scales the glow distance. This and all the width parameters can be adjusted using the Width Widget. Note that a zero glow width still enhances the bright areas; set the brightness parameter to zero if you want to pass the Source through unchanged.

- **Width Red** (Default: 1, Range: 0 or greater)
  Scales the red glow width. If the red, green, and blue widths are equal, the glows will match the color of the source clip. If they are not equal, the glows will vary in color with distance.

- **Width Green** (Default: 1, Range: 0 or greater)
  Scales the green glow width.

- **Width Blue** (Default: 1, Range: 0 or greater)
  Scales the blue glow width.

- **Gradient Opacity** (Default: 0.4, Range: 0 to 1)
  Control the transparency of the gradient before combined with the reflection.

- **Color 1** (Default rgb: [1 0 0])
  Color of the first control point of the gradient.

- **Point 1** (X & Y, Default: [0.599 -0.526], Range: any)
  Location of the first control point of the gradient.

- **Color 2** (Default rgb: [0 1 0])
  Color of the second control point of the gradient.

- **Point 2** (X & Y, Default: [-0.98 0.529], Range: any)
  Location of the second control point of the gradient.

- **Color 3** (Default rgb: [0 0 1])
  Color of the third control point of the gradient.

- **Point 3** (X & Y, Default: [0.958 0.529], Range: any)
  Location of the third control point of the gradient.

- **Combine** (Popup menu, Default: Screen)
  Determines how the gradient is combined with the background.
  - **Grad Only**: gives the gradient image alone with no background.
  - **Mult**: the background is multiplied by the gradient.
  - **Add**: the background is added to the gradient.
  - **Screen**: the background is blended with the gradient using a
screen operation.
  - **Difference**: the result is the difference between the background
and gradient.
  - **Overlay**: combines gradient and background using an overlay function.

- **Lights** (Default: 0, Range: 0 or greater)
  The type of lens flare to apply to generate a light pattern. Custom lens flare types can also be made, or existing types modified, by editing the flare with the Flare Designer.

- **Lights Brightness** (Default: 1, Range: 0 or greater)
  Scale the brightness of the lights.

- **Lights Size** (Default: 2.25, Range: 0 or greater)
  Scale the size of the lights.

- **Lights Start Position** (X & Y, Default: [-1 0], Range: any)
  The starting location of the hotspot of the lights pattern.

- **Lights Pos Uses Mocha** (Check-box, Default: off)
  Use the mocha track to control the position of the lights.

- **Shift Speed** (X & Y, Default: [0.167 -0.178], Range: any)
  How fast the hotspot should move across the clip.

- **Lights Pivot** (X & Y, Default: [0.565 0.328], Range: any)
  The location of the center of the lights pattern.

- **Pivot Uses Mocha** (Check-box, Default: off)
  Use the mocha track to control the pivot of the lights.

- **Mask Use** (Popup menu, Default: Luma)
  Determines how the Mask input channels are used to make a monochrome mask.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Mask** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the prisms. The maximum of the red, green, and blue prism brightness is scaled by this value and combined with the background Alpha at each pixel.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Position** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Position parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Source Position** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Source Position parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Point 1** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Point 1 parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Point 2** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Point 2 parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Point 3** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Point 3 parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Lights Start Position** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Lights Start Position parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Lights Pivot** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Lights Pivot parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

