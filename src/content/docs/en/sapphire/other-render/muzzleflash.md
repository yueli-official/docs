---
title: MuzzleFlash
---

## S_MuzzleFlash

Simulates the flash and smoke that is generated when a gun is fired. The flash from
several types of gun can be simulated. All guns have a primary flash, and guns with suppressors
may have secondary flashes. The gun may easily be fired repeatedly.
Note: All time units in this effect are in frames.

In the Sapphire Render effects submenu.

![MuzzleFlash](../_static/MuzzleFlash.jpg)


### Inputs:

- **Background**: The current layer. The clip to use as background.

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

- **Location** (X & Y, Default: [0 0], Range: any)
  The position (on the image) of the end of the gun barrel. This parameter can be adjusted using the MuzzleFlash Widget.

- **Location Uses Mocha** (Check-box, Default: off)
  Controls whether the location is controlled by the Location parameter or follows the Location tracked inside of Mocha.

- **Smooth Location Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Elevation** (Default: 0, Range: -360 to 360)
  The angle of the gun barrel in the vertical plane relative to horizontal. This parameter can be adjusted using the MuzzleFlash Widget.

- **Rotation** (Default: 0, Range: -360 to 360)
  The angle of the gun barrel in the horizontal plane relative to due east. This parameter can be adjusted using the MuzzleFlash Widget.

- **Radius** (Default: 0.2, Range: 0.05 or greater)
  The overall size of the gun blast. This parameter can be adjusted using the MuzzleFlash Widget.

- **Brightness** (Default: 1, Range: 0 or greater)
  The overall brightness of the flash.

- **Shift Out** (Default: 0, Range: 0 or greater)
  The muzzle flash will be drawn this far from the end of the gun barrel.

- **Gun** (Popup menu, Default: M16)
  The type of gun being fired.
  - **AK47**: AK47 assault rifle.
  - **Beretta**: Beretta pistol.
  - **Colt45**: Colt 45 revolver.
  - **M16**: M16 assault rifle.
  - **12Gauge**: Shotgun.

- **Variant** (Popup menu, Default: Two)
  Variations on the muzzle flash pattern.
  - **One**: First variation.
  - **Two**: Second variation.
  - **Three**: Third variation.
  - **Random**: Randomly pick one of the variations each time a shot is fired.

- **Seed** (Default: 0.123, Range: 0 or greater)
  Random number seed. Many aspects of the muzzle flash appearence depend on the seed. To get a different result for a given set of control values, try changing this seed value.

- **Animate Seed** (Check-box, Default: off)
  Selecting this will subtly change the shape of each flash generated so that they differ from frame to frame.

- **Octaves** (Integer, Default: 3, Range: 1 to 6)
  Increasing this increases the detail in the smoke.

- **Blur** (Default: 0.0015, Range: 0 or greater)
  How much blur is to be applied to the muzzle flash.

- **Mid Density** (Default: 0.4, Range: 0.01 to 0.99)
  How dense the muzzle flash should be at half its radius.

- **Mid Color** (Default rgb: [1 0.792 0.6])
  The color of the muzzle flash at half its radius.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the muzzle flash. The maximum of the red, green, and blue muzzle flash brightness is scaled by this value and combined with the Background Alpha at each pixel.

- **Combine** (Popup menu, Default: Screen)
  Determines how the muzzle flash is combined with the Background.
  - **Screen**: The muzzle flash is blended with the Background using a screen operation.
  - **Add**: The muzzle flash is added to the Background.
  - **MuzzleFlash Only**: Gives only the muzzle flash with no Background.

- **Primary Length** (Default: 0.21, Range: 0 or greater)
  Length of the primary flash. This is always aligned with the gun barrel and is present in all guns.

- **Primary Width** (Default: 0.035, Range: 0 or greater)
  Width of the primary flash.

- **Primary Brightness** (Default: 0.6, Range: 0 or greater)
  Overall density and brightness of the primary flash.

- **Primary Puff** (Default: 0.8, Range: 0 or greater)
  How puffy (as opposed to smooth) the primary flash should appear.

- **Primary Detail** (Default: 0.5, Range: 0 or greater)
  How much detail should appear in the primary flash specifically.

- **Secondary Number** (Integer, Default: 6, Range: 0 to 10)
  How many secondary flashes should be generated for those guns that have secondary flashes. These flashes appear from holes in the suppressor and only guns with suppressors have them. The flashes appear at an angle in the vertical plane with respect to the gun barrel direction. This control specifies how many holes there are in the suppressor. The holes are assumed to be evenly distributed around the circumference of the suppressor.

- **Secondary Angle** (Default: 66, Range: 0 to 89)
  The angle the secondary flashes make in the vertical plane with respect to the gun barrel.

- **Secondary Length** (Default: 0.11, Range: 0 or greater)
  The length of the secondary flashes.

- **Secondary Width** (Default: 0.02, Range: 0 or greater)
  The width of the secondary flashes.

- **Secondary Brightness** (Default: 0.6, Range: 0 or greater)
  The overall density and brightness of the secondary flashes.

- **Secondary Puff** (Default: 0.7, Range: 0 or greater)
  How puffy (as opposed to smooth) the secondary flashes should appear.

- **Secondary Detail** (Default: 0.8, Range: 0 or greater)
  How much detail should appear in the secondary flashes specifically.

- **Time Mode** (Popup menu, Default: Always)
  How to draw flashes as the frame number changes.
  - **Always**: Always draw the flash. Whether or not the flash appears must be controlled manually by key framing the controls.
  - **Repeat**: Fire the weapon at regular intervals between a start and end time. This is appropriate for automatic weapons,
but can also be used to generate a single flash for a non-automatic weapon. The flash is drawn on a single frame.
  - **Repeat+Fade**: As Repeat, but the flash will fade out over a number of frames.

- **Start Time** (Integer, Default: 0, Range: 0 or greater)
  In Repeat and Repeat+Fade mode, this is the first frame on which a flash will appear.

- **Duration** (Integer, Default: 10, Range: 0 to 500)
  In Repeat and Repeat+Fade mode, this is the length of time after start time in which a flash may appear.

- **Gap Time** (Integer, Default: 3, Range: 1 to 50)
  In Repeat and Repeat+Fade mode, this is number of frames between flashes.

- **Fade Time** (Integer, Default: 1, Range: 1 to 5)
  In Repeat+Fade mode, this is time in frames over which the flash will fade away.

- **Glow Color** (Default rgb: [0.1 0.1 1])
  The color of an overall glow around the bright parts of the muzzle flash.

- **Glow Bright** (Default: 3, Range: 0 to 8)
  The brightness of an overall glow around the muzzle flash.

- **Glow Width** (Default: 1, Range: 0 to 2)
  The size of an overall glow around the muzzle flash.

- **Glow Threshold** (Default: 0.1, Range: 0.01 to 1)
  Muzzle flash intensities above this threshold will glow.

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

- **Show MuzzleFlash** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Location parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

