---
title: UltraZap
---

## S_UltraZap

Generates lightning bolts along a spline and renders them over a
background. Increase Vary Endpoint to spread out the ends of the bolts.
Adjust the Glow Color for differently colored results. Adjust the Loop Speed
parameter to make it traverse the spline over time.

In the Sapphire Render effects submenu.

![UltraZap](../_static/UltraZap.jpg)


### Inputs:

- **Background**: The current layer. The clip to combine the glows with. If no background is given, the Source is also used as the Background.

- **Matte**: Defaults to None. If provided, defines the region where the lightning should be rendered. The glow on the lightning will bleed into areas with no lightning for a natural look.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: UltraZap)
  Selects between following the Sapphire spline tool, the AE paths, or the Mocha outlines.
  - **UltraZap**: Creates a zap along a spline.
  - **UltraZapAEPath**: Creates a zap along a single AE Path, all AE Paths, or the current AE text layer. This mode is not available in Premiere.
  - **UltraZapMocha**: Creates a zap along a single Mocha mask outline or all Mocha mask outlines.

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

- **Bypass Mocha** (Check-box, Default: on)
  Ignore the Mocha Mask and apply the effect to the entire source clip.

- **Show Mocha Only** (Check-box, Default: off)
  Bypass the effect and show the Mocha Mask itself.

- **Combine Masks** (Popup menu, Default: Union)
  Determines how to combine the Mocha Mask and Input Mask when both are supplied to the effect.
  - **Union**: Uses the area covered by both masks together.
  - **Intersect**: Uses the area that overlaps between the two masks.
  - **Mocha Only**: Ignore the Input Mask and only use the
Mocha Mask.

- **Start Point** (X & Y, Default: [-0.5 0.5], Range: any)
  The starting point of the bolts.

- **Point 1 Enable** (Check-box, Default: off)
  Turns on or off the first control point.

- **Control Point 1** (X & Y, Default: [-0.34 0.304], Range: any)
  First spline control point.

- **Point 2 Enable** (Check-box, Default: off)
  Turns on or off the second control point.

- **Control Point 2** (X & Y, Default: [0.1 0.2], Range: any)
  Second spline control point.

- **Point 3 Enable** (Check-box, Default: on)
  Turns on or off the third control point.

- **Control Point 3** (X & Y, Default: [0.4 0], Range: any)
  Third spline control point.

- **Point 4 Enable** (Check-box, Default: off)
  Turns on or off the fourth control point.

- **Control Point 4** (X & Y, Default: [0.45 -0.25], Range: any)
  Fourth spline control point.

- **End Point** (X & Y, Default: [0.5 -0.5], Range: any)
  The end point of the bolts. This parameter can be adjusted using the End Widget.

- **Start Uses Mocha** (Check-box, Default: off)
  Controls whether the staring point is controlled by the Start parameter or follows the Start point track from Mocha.

- **Smooth Start Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Point1 Uses Mocha** (Check-box, Default: off)
  Controls whether point 1 is controlled by the Control Point 1 parameter or follows the Control Point 1 track from Mocha.

- **Smooth Point1 Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Point2 Uses Mocha** (Check-box, Default: off)
  Controls whether point 2 is controlled by the Control Point 2 parameter or follows the Control Point 2 track from Mocha.

- **Smooth Point2 Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Point3 Uses Mocha** (Check-box, Default: off)
  Controls whether point 3 is controlled by the Control Point 3 parameter or follows the Control Point 3 track from Mocha.

- **Smooth Point3 Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Point4 Uses Mocha** (Check-box, Default: off)
  Controls whether point 4 is controlled by the Control Point 4 parameter or follows the Control Point 4 track from Mocha.

- **Smooth Point4 Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **End Uses Mocha** (Check-box, Default: off)
  Controls whether the end point is controlled by the End parameter or follows the End point track from Mocha.

- **Smooth End Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Follow Selection** (Popup menu, Default: text)
  Specify what the zap should trace.
  - **selected path**: Follows a single AE path specified in the
  - **all paths**: Follows all AE Masks applied to the current layer.
  - **text**: If UltraZap is applied to an AE text layer, this selection

- **Path To Follow** (Default: 0, Range: 0 or greater)
  AE path for the lightning to follow.

- **Bolts** (Integer, Default: 1, Range: 0 to 500)
  The number of lightning bolts to draw, each between the Start and End location.

- **Bolt Width** (Default: 0.06, Range: 0 or greater)
  The width of the lightning bolts.

- **Bolt Length** (Default: 1, Range: 0 to 1)
  The length of the bolts, beginning at Start Offset. If less than 1, the bolts will not be drawn all the way from start to end. This can be useful for animating a lightning strike.

- **Wrinkle Amp** (Default: 1, Range: 0 or greater)
  Scales the amount of wrinkles in the bolts. Decrease for straighter smoother bolts or increase for more kinky bolts.

- **Curve Amp** (Default: 0.5, Range: 0 or greater)
  Similar to Wrinkle Amp but affects the general path of the bolt. If decreased, the bolt will stay closer to the line between the Start and End points. If increased it can wander further away from this line. This differs from the Wrinkle Amp parameter in that it can be used to make straighter bolts while still keeping the wrinkles at the detailed level.

- **Taper Start** (Default: 0.025, Range: 0 to 1)
  Determines how pointed the start of the bolts are. If 0, the entire bolt will have equal width. If 1, the bolts will thin out along their entire length for a pointed end. If it is .5, the bolts will start thin and grow to the width parameter at the center of the bolt.

- **Taper End** (Default: 0.25, Range: 0 to 1)
  Determines how pointed the end of the bolts are. If 0, the entire bolt will have equal width. If 1, the bolts will thin out along their entire length for a pointed end. If it is .5, the bolts will start thinning out half way between the start and end points.

- **Feather Bolt** (Default: 0, Range: 0 or greater)
  Add softness to the bolt before glowing.

- **Zap Edge Blur** (Default: 0, Range: 0 or greater)
  Only blur the edges of the zap - useful if there is noise on the zap.

- **Zap Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the lightning bolts.

- **Start Color** (Default rgb: [0.87 0.969 1])
  The color at the start of the main bolt.

- **End Color** (Default rgb: [0 0.333 1])
  The color at the tip of the main bolt.

- **Vary Bolt Hue** (Default: 0, Range: -1 to 1)
  This varies the start and end color between multiple bolts. This does not change the color along a single bolt.

- **Vary Width** (Default: 0, Range: 0 to 1)
  The amount of random variation in the width of the bolts along their lengths.

- **Vary Endpoint** (Default: 0, Range: 0 or greater)
  Offsets the End location by a random amount within a circle of this radius. If Bolts is greater than 1, this can be useful to spread out the different End points. For example, you can create multiple radiating bolts by increasing this radius and placing the End point near the Start point. This parameter can also be adjusted using the End Widget, after it is made positive.

- **Random Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different random lightning bolts, and the same value should give a repeatable result.

- **Start Offset** (Default: 0, Range: 0 to 1)
  The offset from the start point to begin drawing the bolts. This can be useful for animating a lightning strike.

- **Loop Speed** (Default: 0, Range: any)
  How fast to animate where the zap starts along the spline.

- **Loop Start** (Default: 0, Range: any)
  Where along the spline the zap starts for animation purposes.

- **Vary Loop Start** (Default: 0, Range: -1 to 1)
  This varies where multiple bolts start along the spline from one another.

- **Zap Direction** (Popup menu, Default: all same)
  Control which direction along the spline the zap renders.
  - **all same**: Normal zap direction for all splines.
  - **all reverse**: Flip the direction of all splines.
  - **even reverse**: Reverse the direction of even splines.
  - **odd reverse**: Reverse the direction of odd splines.
  - **random reverse**: Randomly pick normal or reverse for each spline.

- **Wiggle Start** (Default: 0, Range: 0 or greater)
  By default the bolts automatically wiggle over time. This parameter provides a starting offset for these bolt perturbations.

- **Wiggle Speed** (Default: 1, Range: 0 or greater)
  The speed at which the bolts are perturbed automatically over time. To animate changes in speed, set this to zero and animate the Wiggle Start parameter instead.

- **Jitter Frames** (Integer, Default: 0, Range: 0 or greater)
  If this is 0, the same random lightning will be used for every frame processed. If it is 1, different random lightning is used for each frame. If it is 2, new random lightning is used for every other frame, and so on.

- **Bolt Noise** (Default: 0.25, Range: 0 or greater)
  Apply a noise pattern to the bolt before the glow.

- **Bolt Noise Freq** (Default: 0.05, Range: 0.01 or greater)
  The frequency of the bolt noise.

- **Bolt Noise Speed** (Default: 1, Range: any)
  The speed of the bolt noise.

- **Branches Only** (Check-box, Default: off)
  Hide the core bolt and only show the branches.

- **Branchiness** (Default: 4, Range: 0 to 20)
  Scales the number of additional bolts that branch from the main bolt. Set this to 0 for basic bolts with no extra branches.

- **Branch Angle** (Default: 45, Range: 0 to 180)
  The maximum angle of the random branches relative to the bolt they are branching off of. If this is 0 the branches will be more lined up with the main bolt. With larger values the branches will be more perpendicular to the main bolt.

- **Branch Length** (Default: 0.5, Range: 0 to 3)
  The scaled length of the branches relative to the distance between the Start and End points.

- **Branch Color** (Default rgb: [0.8 0.867 1])
  The color at the tips of the branches.

- **Vary Branch Hue** (Default: 0, Range: -1 to 1)
  This varies the end color between multiple branches.

- **Branch Style** (Popup menu, Default: width uses dist)
  Specify how the branch width should be calculated.
  - **width uses dist**: Brandh width is calculated based on
how far along the bolt the branch splits off.
  - **width uses bolt width**: Brandh width is calculated based on
how wide the bolt is as the point where the branch splits off.

- **Branch Width** (Default: 0.5, Range: 0 or greater)
  Width of branches relative to the main bolt.

- **Secondary Bolts** (Integer, Default: 0, Range: 0 or greater)
  The number of independent forked lightning bolts to draw for the secondary zap.

- **Secondary Bolt Width** (Default: 0.05, Range: 0 or greater)
  The width of the secondary bolts.

- **Secondary Length** (Default: 0.5, Range: 0 to 1)
  The length of the secondary bolts, beginning at Start Offset. If less than 1, the bolts will not be drawn all the way from start to end. This can be useful for animating a lightning strike.

- **Secondary Wrinkle Amp** (Default: 0.75, Range: 0 or greater)
  Scales the amount of wrinkles in the secondary bolts. Decrease for straighter, smoother bolts or increase for more kinky bolts.

- **Secondary Curve Amp** (Default: 0.35, Range: 0 or greater)
  Similar to Wrinkle Amp but affects the general path of the secondary bolt. If decreased, the bolt will stay closer to the line between the Start and End points. If increased it can wander further away from this line. This differs from the Wrinkle Amp parameter in that it can be used to make straighter bolts while still keeping the wrinkles at the detailed level.

- **Secondary Taper Start** (Default: 0, Range: 0 to 1)
  Determines how pointed the start of the secondary bolts are. If 0, the entire bolt will have equal width. If 1, the bolts will thin out along their entire length for a pointed end. If it is .5, the bolts will start thin and grow to the width parameter at the center of the bolt.

- **Secondary Taper End** (Default: 0.25, Range: 0 to 1)
  Determines how pointed the end of the secondary bolts are. If 0, the entire bolt will have equal width. If 1, the bolts will thin out along their entire length for a pointed end. If it is .5, the bolts will start thinning out half way between the start and end points.

- **Secondary Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the seoncary zap.

- **Secondary Start Color** (Default rgb: [1 1 1])
  The color at the start of the secondary bolt.

- **Secondary End Color** (Default rgb: [0.557 0 0.557])
  The color at the tip of the secondary bolt.

- **Secondary Vary Bolt Hue** (Default: 0, Range: -1 to 1)
  This varies the color between multiple secondary bolts. This does not change the color along a single bolt.

- **Secondary Vary Width** (Default: 0, Range: 0 to 1)
  The amount of random variation in the width of the secondary bolts along their lengths.

- **Secondary Vary Endpoint** (Default: 0, Range: 0 or greater)
  Offsets the End location by a random amount within a circle of this radius. If Bolts is greater than 1, this can be useful to spread out the different End points. For example, you can create multiple radiating bolts by increasing this radius and placing the End point near the Start point.

- **Secondary Seed** (Default: 0.456, Range: 0 or greater)
  A separate random seed for the secondary bolt so that it moves differently from the primary bolt.

- **Secondary Start Offset** (Default: 0, Range: 0 to 1)
  The offset from the start point to begin drawing the bolts. This can be useful for animating a lightning strike.

- **Secondary Loop Speed** (Default: 0, Range: any)
  How fast to animate where the secondary zap starts along the spline.

- **Secondary Loop Start** (Default: 0, Range: any)
  Where along the spline the zap starts.

- **Secondary Vary Loop Start** (Default: 0, Range: -1 to 1)
  This varies where multiple bolts start along the spline from one another.

- **Secondary Direction** (Popup menu, Default: all same)
  Control which direction along the secondary spline the zap renders.
  - **all same**: Normal zap direction for all secondary splines.
  - **all reverse**: Flip the direction of all secondary splines.
  - **even reverse**: Reverse the direction of even secondary splines.
  - **odd reverse**: Reverse the direction of odd secondary splines.
  - **random reverse**: Randomly pick normal or reverse for each secondary spline.

- **Secondary Wiggle Start** (Default: 0, Range: 0 or greater)
  By default the bolts automatically wiggle over time. This parameter provides a starting offset for these bolt perturbations.

- **Secondary Wiggle Speed** (Default: 1, Range: 0 or greater)
  The speed at which the bolts are perturbed automatically over time. To animate changes in speed, set this to zero and animate the Wiggle Start parameter instead.

- **Secondary Jitter Frames** (Integer, Default: 0, Range: 0 or greater)
  If this is 0, the noise texture will remain the same for every frame processed. If it is 1, a new noise texture is used for each frame. If it is 2, a new noise texture is used for every other frame, and so on.

- **Secondary Branchiness** (Default: 4, Range: 0 to 20)
  Scales the number of additional bolts that branch from the main bolt. Set this to 0 for basic bolts with no extra branches.

- **Secondary Branch Angle** (Default: 45, Range: 0 to 180)
  The maximum angle of the random branches relative to the bolt they are branching off of. If this is 0 the branches will be more lined up with the main bolt. With larger values the branches will be more perpendicular to the main bolt.

- **Secondary Branch Width** (Default: 0.5, Range: 0 or greater)
  Width of secondary branches compared to the secondary bolt.

- **Secondary Branch Length** (Default: 0.7, Range: 0 to 3)
  The scaled length of the secondary branches relative to the distance between the Start and End points.

- **Secondary Branch Color** (Default rgb: [1 0.8 1])
  The color at the tips of the secondary branches.

- **Secondary Vary Branch Hue** (Default: 0, Range: -1 to 1)
  This varies the end color between multiple branches.

- **Glow Brightness** (Default: 1.5, Range: 0 or greater)
  Scales the brightness of the glow applied to the lightning.

- **Glow Threshold** (Default: 0, Range: 0 or greater)
  Threshold for what colors are glowed from the zap.

- **Glow Color** (Default rgb: [0.97 0.87 1])
  The color of the glow applied to the lightning.

- **Glow Width** (Default: 0.3, Range: 0 or greater)
  The width of the glow applied to the lightning.

- **Glow Width Red** (Default: 1, Range: 0 or greater)
  The relative red width of the glow.

- **Glow Width Green** (Default: 1, Range: 0 or greater)
  The relative green width of the glow.

- **Glow Width Blue** (Default: 1, Range: 0 or greater)
  The relative blue width of the glow.

- **Glow Falloff** (Default: 0.05, Range: -2 to 2)
  Boost or cut the distance that the glow extends.

- **Glow Bias** (Default: 0.5, Range: -3 to 3)
  Amount to grow the outskirts of the thresholded result, or shrink if negative.

- **After Glow Width** (Default: 0.6, Range: 0 or greater)
  Scales the glow distance for the secondary glow.

- **After Glow Color** (Default rgb: [1 1 1])
  Scales the color of the secondary glow.

- **Glow Highlights** (Check-box, Default: off)
  Enables highlights using an electrical plasma texture.

- **Glow Highlights Frequency** (Default: 1.2, Range: 0.01 or greater)
  The spatial frequency of the highlights. Increase to zoom out, decrease to zoom in.

- **Glow Highlights Speed** (Default: 1, Range: any)
  Phase speed of the highlights. If non-zero, the lines are automatically animated to undulate at this rate.

- **Glow Atmosphere** (Check-box, Default: off)
  Turn atmoosphere on and off.

- **Glow Atmosphere Amp** (Default: 1, Range: 0 or greater)
  Atmosphere gives the effect of rays shining through a dusty atmosphere and picking up light or getting shadowed. This parameter adjusts the amount, or amplitude, of the atmospheric effect. Zero gives smooth rays, higher values give more dusty look.

- **Glow Atmosphere Freq** (Default: 0.5, Range: 0.1 or greater)
  Controls the frequency of the atmospheric noise.

- **Glow Atmosphere Speed** (Default: 1, Range: any)
  The cloudy noise in the atmosphere evolves over time like real dust clouds; this parameter controls how fast the cloud pattern changes over time. Set to zero for a static pattern.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the lightning and its glow. The maximum of the red, green, and blue brightness is scaled by this value and combined with the background Alpha at each pixel.

- **Combine** (Popup menu, Default: Screen)
  Determines how the lightning and glow are combined with the Background.
  - **Screen**: performs a blend function which can help prevent
overly bright results.
  - **Add**: causes the lightning to be added to the background. This gives
brighter glows over light backgrounds.
  - **Zap Only**: displays the lightning over black, ignoring the source image.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the background before combining with the lightning. If 0, the result will contain only the lightning image over black.

- **Blur Matte** (Default: 0.05, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Matte Use** (Popup menu, Default: Luma)
  Determines how the Matte input channels are used to make a monochrome matte.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Show Spline** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Start Point parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

