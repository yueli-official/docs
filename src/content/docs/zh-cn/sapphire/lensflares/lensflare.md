---
title: LensFlare
---

## S_LensFlare

Renders a lens flare image over the background clip,
aligning various flare elements between the hotspot and pivot
locations. Use the Lens menu to select different types of
lensflares.

In the Sapphire Lighting effects submenu.

![LensFlare](../_static/LensFlare.jpg)


### Inputs:

- **Background**: The current layer. The clip to apply the lens flare over.

- **Occlusion**: Defaults to None. Obscures and colorizes the flare. The brightness of the flare is reduced based on the opacity of this clip, and the color at the hotspot location is used to tin the flare . These behaviors can be adjusted with the Occlusion Softness, Occlusion From, Invert Occlusion, and Use Color parameters.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: 2D)
  Selects between several variations of the LensFlare effect.
  - **2D**: The hotspot is positioned manually. The flare may be occluded by a Matte input.
  - **3D**: One or more hotspots can be connected to Lights within a 3D composition.
The flare may be occluded by 3D layers within the comp.

- **Lens** (Default: 0, Range: 0 or greater)
  The type of lens flare to apply. Custom lens flare types can also be made, or existing types modified, by editing the flare in the flare designer.

- **Mocha Project** (Default: 0, Range: 0 or greater)
  Brings up the Mocha window for tracking footage and generating masks.

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  Blurs the Mocha Mask by this amount before using. This can be used to soften the edges or quantization artifacts of the mask, and smooth out the time displacements.

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  Controls the strength of the Mocha mask. Lower values reduce the intensity of the effect.

- **Invert Mocha** (Check-box, Default: off)
  If enabled, the black and white of the Mocha Mask are inverted before applying the effect.

- **Resize Mocha** (Default: 1, Range: 0 or greater)
  Scales the Mocha Mask. 1.0 is the original size.

- **Resize Rel X** (Default: 1, Range: 0 or greater)
  The relative horizontal size of the Mocha Mask.

- **Resize Rel Y** (Default: 1, Range: 0 or greater)
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

- **Scale Widths** (Default: 1.15, Range: 0 or greater)
  Scales the sizes of all the flare elements. This parameter can be adjusted using the Scale Widths Widget.

- **Rel Heights** (Default: 1, Range: 0 or greater)
  Scales the vertical dimension of all the flare elements, making them elliptical instead of circular. This can also be adjusted using the Scale Widths Widget.

- **Rays Rotate** (Default: 0, Range: any)
  Rotates the ray elements of the lens flare, if any, in degrees.

- **Hotspot** (X & Y, Default: [-0.444 0.176], Range: any)
  The location of the brightest spot in the flare in screen coordinates. It can be set by enabling and moving the hotspot widget.

- **Hotspot Uses Mocha** (Check-box, Default: off)
  Controls whether the LensFlare hotspot is controlled by the Hotspot parameter or follows the Hotspot tracked inside of Mocha.

- **Smooth Hotspot Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Hotspots** (Default: 0, Range: 0 or greater)
  The AE light(s) to attach the flare center to.

- **Pivot** (X & Y, Default: [0 0], Range: any)
  The elements of the flare will be in a line between the Hotspot and the Pivot locations. The Pivot location is in screen coordinates.

- **Pivot Uses Mocha** (Check-box, Default: off)
  Controls whether the LensFlare pivot is controlled by the Pivot parameter or follows the Pivot tracked inside of Mocha.

- **Smooth Pivot Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of all the flare elements.

- **Color** (Default rgb: [1 1 1])
  Scales the color of all flare elements.

- **Gamma** (Default: 1, Range: 0.1 or greater)
  Increasing gamma brightens the flare, and especially boosts the darker elements.

- **Saturation** (Default: 1, Range: -2 to 8)
  Scales the color saturation of the flare elements. Increase for more intense colors. Set to 0 for a monochrome lens flare.

- **Hue Shift** (Default: 0, Range: any)
  Shifts the hue of the flare, in revolutions from red to green to blue to red.

- **Hotspot Bright** (Default: 1, Range: 0 or greater)
  Scales the brightness of the hotspot elements only.

- **Hotspot Color** (Default rgb: [1 1 1])
  Scales the color of the hotspot elements only.

- **Rays Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the ray elements only.

- **Rays Num Scale** (Default: 1, Range: 0 or greater)
  Increases or decreases the number of rays.

- **Rays Length** (Default: 1, Range: 0 or greater)
  Adjusts the length of the rays without changing their thickness, or changing the size of the other flare elements.

- **Rays Thickness** (Default: 1, Range: 0 or greater)
  Adjusts the thickness of the individual rays within the flare.

- **Other Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of all flare elements that are NOT at the hotspot location.

- **Other Color** (Default rgb: [1 1 1])
  Scales the color of all flare elements that are NOT at the hotspot location.

- **Other Width** (Default: 1, Range: 0 or greater)
  Scales the width of all flare elements that are NOT at the hotspot location.

- **Atmosphere Amp** (Default: 0, Range: 0 or greater)
  Atmosphere gives the effect of the flare shining through a dusty atmosphere and picking up light or getting shadowed. This parameter adjusts the amount, or amplitude, of the atmospheric effect. Zero gives a smoother flare, higher values give more dirty look.Atmosphere affects some parts of the flare but not others. This behavior can be adjusted for each element by toggling the Ignore Atmosphere setting within the Flare Designer. Typically, elements which originate within the camera (such as secondary reflections) should use Ignore Atmosphere, while elements which originate outside it (such as glows) should not.

- **Atmosphere Freq** (Default: 1, Range: 0.1 to 20)
  Controls the spatial frequency of the atmospheric noise. Turn this up higher to get finer details, turn down for broader overall variation.

- **Atmosphere Detail** (Default: 0.6, Range: 0 to 1)
  Controls the amount of fine detail in the atmosphere simulation. Decrease to get smoother atmosphere, increase for a more crunchy or grainy look.

- **Atmosphere Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator for the atmospheric noise. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Atmosphere Speed** (Default: 1, Range: any)
  The cloudy noise in the atmosphere evolves over time like real dust clouds; this parameter controls how fast the cloud pattern changes over time. Set to zero for a static pattern.

- **Flicker Amp** (Default: 0, Range: 0 or greater)
  The amount of random flickering of the flare brightness.

- **Flicker Speed** (Default: 1, Range: 0 or greater)
  The speed of random flickering.

- **Flicker Randomness** (Default: 0.6, Range: 0 to 1)
  Controls the variability of the flicker. When set to zero, the flare will flicker constantly, with a small amount of random variation. At higher values, the flickering will have longer steady spells, with the occasional large spike.

- **Atmosphere Amp** (Default: 0, Range: 0 or greater)
  Atmosphere gives the effect of the flare shining through a dusty atmosphere and picking up light or getting shadowed. This parameter adjusts the amount, or amplitude, of the atmospheric effect. Zero gives a smoother flare, higher values give more dirty look.Atmosphere affects some parts of the flare but not others. This behavior can be adjusted for each element by toggling the Ignore Atmosphere setting within the Flare Designer. Typically, elements which originate within the camera (such as secondary reflections) should use Ignore Atmosphere, while elements which originate outside it (such as glows) should not.

- **Atmosphere Freq** (Default: 1, Range: 0.1 to 20)
  Controls the spatial frequency of the atmospheric noise. Turn this up higher to get finer details, turn down for broader overall variation.

- **Atmosphere Detail** (Default: 0.6, Range: 0 to 1)
  Controls the amount of fine detail in the atmosphere simulation. Decrease to get smoother atmosphere, increase for a more crunchy or grainy look.

- **Atmosphere Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator for the atmospheric noise. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Atmosphere Speed** (Default: 1, Range: any)
  The cloudy noise in the atmosphere evolves over time like real dust clouds; this parameter controls how fast the cloud pattern changes over time. Set to zero for a static pattern.

- **Flicker Amp** (Default: 0, Range: 0 or greater)
  The amount of random flickering of the flare brightness.

- **Flicker Speed** (Default: 1, Range: 0 or greater)
  The speed of random flickering.

- **Flicker Randomness** (Default: 0.6, Range: 0 to 1)
  Controls the variability of the flicker. When set to zero, the flare will flicker constantly, with a small amount of random variation. At higher values, the flickering will have longer steady spells, with the occasional large spike.

- **Blur Flare** (Default: 0, Range: 0 or greater)
  If positive, the flare image is blurred by this amount before being combined with the background.

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the background before combining with the flare. If 0, the result will contain only the flare image over black.

- **Combine** (Popup menu, Default: Screen)
  Determines how the flare image is combined with the Background.
  - **Screen**: performs a blend function which can help prevent
overly bright results.
  - **Add**: causes the flare image to be added to the background.
  - **Flare Only**: gives only the flare image with no background.

- **Tint Bg Whites** (Check-box, Default: off)
  If this is enabled, the chroma of the flare is added only after the result is clamped to the maximum brightness. This allows the color of the flare image to still be visible even over bright white backgrounds. For the majority of backgrounds there will be no observable difference.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the flare. The maximum of the red, green, and blue flare brightness is scaled by this value and combined with the Background Alpha at each pixel.


### Edge Triggers Parameters:

Edge Width:
*Default:
*0.3,
*Range:
*0 or greater.Creates a trigger zone at the edge of the screen which
affects the brightness and size of the flare. This parameter
controls the width of the zone. Setting it to zero will disable Edge
Triggers.

Edge Falloff:
*Default:
*2,
*Range:
*0.01 or greater.Controls the speed with which the intensity of the
trigger decreases when moving away from the edge. A value of 1
results in a linear ramp. Values greater than one result in a
steeper initial drop which gradually levels out. Values less than
one result in a slope that starts gradually and gets steeper at the end.

Shift Out:
*Default:
*0,
*Range:
*any.Shifts the trigger zone outward from the edge of the
screen, placing the peak off-screen. Negative values will shift the
trigger zone inward toward the center of the image.

One Way:
*Check-box, Default:
*off.If this box is checked, the trigger will stay at maximum
intensity when the Hotspot moves off-screen, instead of ramping back
down as it moves farther from the edge.

Edge Scale Brightness:
*Default:
*1.5,
*Range:
*0 or greater.Scales the Brightness by this amount when
the trigger is at peak intensity. For example, a value of 3 will
cause the flare to ramp up to 3 times its normal brightness as it
passes through the trigger zone. If less than 1, the trigger will
dim the flare instead of brightening it.

Edge Scale Widths:
*Default:
*1.2,
*Range:
*0 or greater.Scales the width of the flare by this amount when
the trigger is at peak intensity. For example, a value of 3 will
cause the flare to expand to 3 times its normal size as it
passes through the trigger zone. If less than 1, the trigger will
shrink the flare instead of growing it.

Show Edge Zones:
*Check-box, Default:
*off.Overlays the output with a grayscale image which
shows the location intensity of the Edge Triggers.

Center Radius:
*Default:
*0.3,
*Range:
*0 or greater.Creates a trigger zone centered on the Pivot which
affects the brightness and size of the flare. This parameter
controls the radius of the zone. Setting it to zero will disable Center
Triggers.

Center Falloff:
*Default:
*1,
*Range:
*0.01 or greater.Controls the speed with which the intensity of the
trigger decreases when moving away from the edge. A value of 1
results in a linear ramp. Values greater than one result in a
steeper initial drop which gradually levels out. Values less than
one result in a slope that starts gradually and gets steeper at the end.

Center Scale Brightness:
*Default:
*1.5,
*Range:
*0 or greater.Scales the Brightness by this amount when
the trigger is at peak intensity. For example, a value of 3 will
cause the flare to ramp up to 3 times its normal brightness as it
passes through the trigger zone. If less than 1, the trigger will
dim the flare instead of brightening it.

Center Scale Widths:
*Default:
*1.2,
*Range:
*0 or greater.Scales the width of the flare by this amount when
the trigger is at peak intensity. For example, a value of 3 will
cause the flare to expand to 3 times its normal size as it
passes through the trigger zone. If less than 1, the trigger will
shrink the flare instead of growing it.

Show Center Zone:
*Check-box, Default:
*off.
Overlays the output with a grayscale image which
shows the location intensity of the Center Trigger.

### Occlusion Parameters:

Occlusion Softness:
*Default:
*0.0224,
*Range:
*0 or greater.Increase this value to make the flare to fade
out gradually as the Hotspot moves behind an object. Set it to zero
to have the flare wink out suddenly. With large softness values, the
Hotspot position is automatically adjusted toward non-occluded
areas. This accurately simulates a large light source which is
partially obscured, and it prevents the Hotspot from appearing in
front of an occluding object.

Occlusion From:
*Popup menu, Default: Alpha
*.Selects the channels of the Occlusion image which control occlusion of the flare.
*None:
*The flare is not occluded at all.*Luma:
*The flare is occluded by the luminance of the
Occlusion clip. Black areas will show the flare, and white areas will obscure it.*Alpha:
*The flare is occluded by the alpha channel of the Occlusion clip. Use this if you
have an RGBA clip of an object that should appear in front of the flare.

Invert Occlusion:
*Check-box, Default:
*off.If enabled, inverts the Occlusion clip so the flare is obscured by black areas
instead of white.

Use Color:
*Check-box, Default:
*on.If enabled, colorizes the flare based on the
color of the Occlusion clip. Use this to create stained glass effects when
the Hotspot passes behind a transparent object.

Diffraction Glow:
*Default:
*0,
*Range:
*0 or greater.Creates a glow on the edges of occluding objects that are near the hotspot, to simulate
light bleeding around the edges of the object. This parameter controls the brightness of the glow.

Glow Width:
*Default:
*0.4,
*Range:
*0 or greater.The width of the Diffraction Glow.

Glow Color:
*Default rgb:
*[1 1 1].The color of the Diffraction Glow.

Glow Radius:
*Default:
*0.2,
*Range:
*0 or greater.
The distance from the hotspot at which Diffraction Glow is visible. The glow will fall off
softly as distance from the hotspot increases. Objects outside this radius won't create any glow.

### Occlusion Triggers Parameters:

Occlusion Scale Brightness:
*Default:
*1,
*Range:
*0 or greater.Enables a Trigger that adjusts the
brighness of the flare as it's occluded. The brightness will ramp up
to this value when the Hotspot is 50 percent occluded, then ramp back down
to its normal brightness. At the same time, the flare will be fading
out due to the occlusion, so its brightness will drop off quickly at
the end. Works best when Occlusion Softness is turned up as well.

Occlusion Scale Widths:
*Default:
*1,
*Range:
*0 or greater.Enables a Trigger that adjusts the size of
the flare as it's occluded. The flare will expand up to this amount
when the Hotspot is 50 percent occluded, then shrink back down to its
normal size. At the same time, the flare will be fading out due to
the occlusion, so it will appear to shrink quickly at the end. Works
best when Occlusion Softness is turned up as well.

Occlusion Layers:
*Default:
*0,
*Range:
*0 or greater.The AE layers to occlude the flare hotspot by.

Performance:
*Popup menu, Default: full flare
*.Determine whether to render all elements or only select
elements. Certain elements are selected in the Flare Designer to be
important for the look of the Flare. Rendering with priority only
should give the look and feel of the true LensFlare for previewing
purposes but render quicker than the full flare.
*full flare:
*Render all LensFlare elements.*priority only:
*Only render a subset of the LensFlare
elements for increased performance.

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

Show Scale Widths:
*Check-box, Default:
*on.Turns on or off the screen interface widget for
adjusting the Scale Widths and Rel Heights parameters.This parameter only appears on AE and Premiere,
where on-screen widgets are supported.

Show Hotspot:
*Check-box, Default:
*on.Turns on or off the screen user interface for adjusting the
Hotspot parameter.This parameter only appears on AE and Premiere,
where on-screen widgets are supported.

Show Pivot:
*Check-box, Default:
*on.
Turns on or off the screen user interface for adjusting the
Pivot parameter.This parameter only appears on AE and Premiere,
where on-screen widgets are supported.See general info for
[Motion Blur](/en/sapphire/#motion-blur)
