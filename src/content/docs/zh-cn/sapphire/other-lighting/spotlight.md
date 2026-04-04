---
title: SpotLight
---

## S_SpotLight

Lights the input clip using one or two spotlights.
For each enabled light, the intersection of a 3D light cone with the
image plane is calculated using the given light source position, aim
location, and beam angle. Ambient light can also be applied to
affect the entire source image evenly. A wide variety of lighting
shapes can be created by adjusting the parameters provided.

In the Sapphire Lighting effects submenu.

![SpotLight](../_static/SpotLight.jpg)


### Inputs:

- **Background**: The current layer. The clip to combine the light with.

- **Mask**: Defaults to None. If provided, the source spot colors are scaled by this input. A monochrome mask can be used to choose a subset of Source areas that will generate spots. A color mask can be used to selectively adjust the spot colors in different regions. The mask is applied to the source before the spots are generated so it will not clip the resulting spots.


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

- **Light1 Enable** (Check-box, Default: on)
  Turns on or off this spotlight.

- **Light1 Uses Mocha** (Check-box, Default: off)
  Controls whether the first light is controlled by the Light 1 parameter or follows the Light 1 tracked inside of Mocha.

- **Smooth Light1 Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Light1 Bright** (Default: 0.8, Range: any)
  Scales the brightness of this spotlight. This value can be made negative for a 'dark' spotlight effect.

- **Light1 Color** (Default rgb: [1 1 1])
  Determines the color of this spotlight.

- **Light1** (X & Y, Default: [-0.5 0.361], Range: any)
  The position of this light source relative to the image plane. This parameter can be adjusted using the Light1 Widget.

- **Light1 Z** (Default: 0.5, Range: 0.028 or greater)
  The distance of this light source from the image plane. Decreasing this brings the light source closer to the surface and causes the direction of the beam to be more parallel to the surface, which can stretch the spot into an ellipse or hyperbola shape.

- **Aim1** (X & Y, Default: [-0.167 0], Range: any)
  This spotlight is directed at this location on the image plane. If this is directly under the Light Source a circular spot will result. When moved away from the Light Source it can also cause the spot to change to an ellipse or hyperbola shape. This parameter can be adjusted using the Aim1 Widget.

- **Aim1 Uses Mocha** (Check-box, Default: off)
  Controls whether the first light's direction is controlled by the Aim 1 parameter or follows the Aim 1 tracked inside of Mocha.

- **Smooth Aim1 Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Spread Angle1** (Default: 45, Range: 0 to 360)
  The spread angle of this spotlight beam in degrees. Larger values open up the beam for a larger spot.

- **Softness1** (Default: 0.3, Range: 0.01 to 1)
  Determines the amount of penumbra or the softness of the spotlight edges, relative to the Spread Angle. Lower values make crisp edged shapes, higher values make softer shapes.

- **Falloff Power1** (Default: 0, Range: 0 or greater)
  Determines how much the spotlight brightness fades with distance. A value of 0 causes no fading, 1 fades the light as distance increases, and 2 fades it faster with distance. A value of 2 is correct for a physically realistic point light.

- **Light2 Enable** (Check-box, Default: off)
  Turns on or off the second spotlight.The remainder of the Light2 parameters are the same as those described above for Light1, but control the second spotlight instead.

- **Ambient Bright** (Default: 0.2, Range: any)
  The amount of ambient light included in the entire frame. This allows parts of the Background outside of the spotlights to still be visible if desired.

- **Ambient Color** (Default rgb: [1 1 1])
  Determines the color of the ambient light.

- **All Lights Bright** (Default: 1, Range: any)
  Scales the brightness of all the spotlights together.

- **All Lights Color** (Default rgb: [1 1 1])
  Scales the color of all the spotlights together.

- **All Aims Shift** (X & Y, Default: [0 0], Range: any)
  Adds this amount to all lights Aim parameters. This can be used to easily make all lights aim at the same location. This parameter can be adjusted using the All Aims Shift Widget.

- **All Shift** (X & Y, Default: [0 0], Range: any)
  Shifts the entire spotlight pattern without changing their shapes by adding this amount to all light and aim positions.

- **Combine** (Popup menu, Default: Mult)
  Determines how the light is combined with the Background.
  - **Lights Only**: gives only the light image with no Background.
  - **Mult**: the light is multiplied by the Background. This is
the effect that a real light would typically have.
  - **Add**: the light is added to the Background.
  - **Screen**: the light is blended with the Background using a
screen operation.
  - **Overlay**: the light is combined with the Background using an
overlay function.

- **Mask Use** (Popup menu, Default: Luma)
  Determines how the Mask input channels are used to make a monochrome mask.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Blur Mask** (Default: 0.05, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Mask** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Show Light1** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Light1 parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Aim1** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Aim1 parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Light2** (Check-box, Default: off)
  Turns on or off the screen user interface for adjusting the Light2 parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Aim2** (Check-box, Default: off)
  Turns on or off the screen user interface for adjusting the Aim2 parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show All Aims Shift** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the All Aims Shift parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

