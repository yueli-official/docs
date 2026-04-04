---
title: Zap
---

## S_Zap

Generates lightning bolts between two points, and renders them over a
background. Increase the number of bolts to give a electrical plasma
effect. Increase Vary Endpoint to spread out the ends of the bolts.
Adjust the Glow Color for differently colored results. The Wiggle Speed
parameter causes the bolts to automatically undulate over time.

In the Sapphire Render effects submenu.

![Zap](../_static/Zap.jpg)


### Inputs:

- **Background**: The current layer. The clip to use as background.

- **Matte**: Defaults to None. If provided, defines the region where the lightning should be rendered. The glow on the lightning will bleed into areas with no lightning for a natural look.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: 2D)
  Selects between 2D and 3D modes.
  - **2D**: creates a zap along a spline.
  - **3D**: creates a three-dimensional zap.
  - **Follow Path**: creates a zap along an AE Path.

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

- **Path To Follow** (Default: 0, Range: 0 or greater)
  AE path for the lightning to follow.

- **Bolts** (Integer, Default: 1, Range: 1 to 500)
  The number of lightning bolts to draw, each between the Start and End location.

- **Start** (X & Y, Default: [-0.5 0.596], Range: any)
  The starting point of the bolts.

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

- **Point 2 Enable** (Check-box, Default: off)
  Turns on or off the second control point.

- **Control Point 2** (X & Y, Default: [0.1 0.25], Range: any)
  Second spline control point.

- **Point2 Uses Mocha** (Check-box, Default: off)
  Controls whether point 2 is controlled by the Control Point 2 parameter or follows the Control Point 2 track from Mocha.

- **Smooth Point2 Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Point 3 Enable** (Check-box, Default: on)
  Turns on or off the third control point.

- **Control Point 3** (X & Y, Default: [0.4 0], Range: any)
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

- **End** (X & Y, Default: [0.5 -0.596], Range: any)
  The end point of the bolts. This parameter can be adjusted using the End Widget.

- **End Uses Mocha** (Check-box, Default: off)
  Controls whether the end point is controlled by the End parameter or follows the End point track from Mocha.

- **Smooth End Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Start** (X & Y, Default: [-0.5 0.596], Range: any)
  The starting point of the bolts.

- **End** (X & Y, Default: [0.5 -0.596], Range: any)
  The end point of the bolts. This parameter can be adjusted using the End Widget.

- **Vary Endpoint** (Default: 0, Range: 0 or greater)
  Offsets the End location by a random amount within a circle of this radius. If Bolts is greater than 1, this can be useful to spread out the different End points. For example, you can create multiple radiating bolts by increasing this radius and placing the End point near the Start point. This parameter can also be adjusted using the End Widget, after it is made positive.

- **Bolt Width** (Default: 0.07, Range: 0 or greater)
  The width of the lightning bolts.

- **Vary Width** (Default: 0, Range: 0 to 1)
  The amount of random variation in the width of the bolts along their lengths.

- **End Pointiness** (Default: 0.1, Range: 0 to 1)
  Determines how pointed the end of the bolts are. If 0, the entire bolt will have equal width. If 1, the bolts will thin out along their entire length for a pointed end. If it is .5, the bolts will start thinning out half way between the start and end points.

- **Wiggle Start** (Default: 0, Range: 0 or greater)
  By default the bolts automatically wiggle over time. This parameter provides a starting offset for these bolt perturbations.

- **Wiggle Speed** (Default: 1, Range: 0 or greater)
  The speed at which the bolts are perturbed automatically over time. To animate changes in speed, set this to zero and animate the Wiggle Start parameter instead.

- **Jitter Frames** (Integer, Default: 0, Range: 0 or greater)
  If this is 0, the same random lightning will be used for every frame processed. If it is 1, different random lightning is used for each frame. If it is 2, new random lightning is used for every other frame, and so on.

- **Rand Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different random lightning bolts, and the same value should give a repeatable result.

- **Wrinkle Amp** (Default: 1, Range: 0 or greater)
  Scales the amount of wrinkles in the bolts. Decrease for straighter smoother bolts or increase for more kinky bolts.

- **Curve Amp** (Default: 0.5, Range: 0 or greater)
  Similar to Wrinkle Amp but affects the general path of the bolt. If decreased, the bolt will stay closer to the line between the Start and End points. If increased it can wander further away from this line. This differs from the Wrinkle Amp parameter in that it can be used to make straighter bolts while still keeping the wrinkles at the detailed level.

- **Branchiness** (Default: 1, Range: 0 to 20)
  Scales the number of additional bolts that branch from the main bolt. Set this to 0 for basic bolts with no extra branches.

- **Branch Angle** (Default: 65, Range: 0 to 180)
  The maximum angle of the random branches relative to the bolt they are branching off of. If this is 0 the branches will be more lined up with the main bolt. With larger values the branches will be more perpendicular to the main bolt.

- **Branch Length** (Default: 0.5, Range: 0 to 3)
  The scaled length of the branches relative to the distance between the Start and End points.


### Glow Parameters:

Glow Bright:
*Default:
*2,
*Range:
*0 or greater.Scales the brightness of the glow applied to the lightning.

Glow Color:
*Default rgb:
*[0.5 0.5 1].The color of the glow applied to the lightning.

Glow Width:
*Default:
*0.224,
*Range:
*0 or greater.The width of the glow applied to the lightning.

Glow Width Red:
*Default:
*0.5,
*Range:
*0 or greater.The relative red width of the glow.

Glow Width Grn:
*Default:
*1,
*Range:
*0 or greater.The relative green width of the glow.

Glow Width Blue:
*Default:
*1.5,
*Range:
*0 or greater.The relative blue width of the glow.

Affect Alpha:
*Default:
*1,
*Range:
*0 or greater.
If this value is positive the output Alpha channel will
include some opacity from the lightning and its glow. The maximum
of the red, green, and blue brightness is scaled by this value and
combined with the background Alpha at each pixel.

### Other Parameters:

Zap Bright:
*Default:
*1,
*Range:
*0 or greater.Scales the brightness of the lightning bolts.

Zap Color:
*Default rgb:
*[1 1 1].The color of the lightning. If you want to keep the
lightning bolt itself bright white, you can still affect the
perceived color by adjusting the Glow Color instead.

Start Offset:
*Default:
*0,
*Range:
*0 to 1.The offset from the start point to begin drawing the
bolts. This can be useful for animating a lightning strike.

Length:
*Default:
*1,
*Range:
*0 to 1.The length of the bolts, beginning at Start Offset. If
less than 1, the bolts will not be drawn all the way from start to end.
This can be useful for animating a lightning strike.

Bg Brightness:
*Default:
*1,
*Range:
*0 or greater.Scales the brightness of the background before
combining with the lightning. If 0, the result will contain only the
lightning image over black.

Combine:
*Popup menu, Default: Screen
*.Determines how the lightning and glow are combined with the Background.
*Screen:
*performs a blend function which can help prevent
overly bright results.*Add:
*causes the lightning to be added to the background. This gives
brighter glows over light backgrounds.*Zap Only:
*displays the lightning over black, ignoring the source image.

Atmosphere Amp:
*Default:
*0,
*Range:
*0 or greater.Atmosphere gives the effect of the zap shining through
a dusty atmosphere and picking up light or getting shadowed.
This parameter adjusts the amount, or amplitude, of the atmospheric
effect. Zero gives a smooth zap, higher values give more dusty look.

Atmosphere Freq:
*Default:
*2,
*Range:
*0.1 to 20.Controls the spatial frequency of the atmospheric
noise. Turn this up higher to get finer details, turn down for
broader overall variation.

Atmosphere Detail:
*Default:
*0.7,
*Range:
*0 to 1.Controls the amount of fine detail in the
atmosphere simulation. Decrease to get smoother atmosphere,
increase for a more crunchy or grainy look.

Atmosphere Seed:
*Default:
*0.123,
*Range:
*0 or greater.Used to initialize the random number generator for
the atmospheric noise. The actual seed value is not significant, but
different seeds give different results and the same value should
give a repeatable result.

Atmosphere Speed:
*Default:
*1,
*Range:
*any.The cloudy noise in the atmosphere evolves over
time like real dust clouds; this parameter controls how fast the
cloud pattern changes over time. Set to zero for a static pattern.

Opacity:
*Popup menu, Default: Normal
*.Determines the method used for dealing with
opacity/transparency.
*All Opaque:
*Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).*Normal:
*Process opacity normally.*As Premult:
*Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

Show:
*Popup menu, Default: Result
*.Selects what the effect will output.
*Result:
*shows the normal lightning result over the background.*ZBuffer:
*shows a depth map of the lightning which can be used
for compositing or to control other effects.

Swivel Zap:
*Default:
*-45,
*Range:
*any.In 3D mode, rotates the lightning left or right about a vertical
axis.

Tilt Zap:
*Default:
*-15,
*Range:
*any.In 3D mode, rotates the lightning up or down about a horizontal
axis. You can use Swivel and Tilt together to rotate about
arbitrary diagonal axes.

Camera Zoom:
*Default:
*0,
*Range:
*-5 to 1.In 3D mode, zooms in or out on the lightning.

Glow Fade:
*Default:
*0.2,
*Range:
*0 or greater.In 3D mode, fades out the glow on more distant parts of the lightning.

Blur Matte:
*Default:
*0.05,
*Range:
*0 or greater.Blurs the Matte input by this amount before using. This
can provide a smoother transition between the matted and unmatted
areas. It has no effect unless the Matte input is provided.

Invert Matte:
*Check-box, Default:
*off.If on, inverts the Matte input so the effect is applied
to areas where the Matte is black instead of white. This has no effect
unless the Matte input is provided.

Matte Use:
*Popup menu, Default: Luma
*.Determines how the Matte input channels are used to make
a monochrome matte.
*Luma:
*the luminance of the RGB channels is used.*Alpha:
*only the Alpha channel is used.

Show Spline:
*Check-box, Default:
*on.Turns on or off the screen user interface for adjusting the
Start parameter.This parameter only appears on AE and Premiere,
where on-screen widgets are supported.

Show Vary Endpoint:
*Check-box, Default:
*on.
Turns on or off the screen user interface for adjusting the
End parameter.This parameter only appears on AE and Premiere,
where on-screen widgets are supported.See general info for
[Motion Blur](/en/sapphire/#motion-blur)
