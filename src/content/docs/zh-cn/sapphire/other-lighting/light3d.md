---
title: Light3D
---

## S_Light3D

Performs 3D relighting with up to 4 individually controlled
light sources. The Source input is usually an ambient or diffuse
pass from a 3d renderer that shows the surface colors. The Normal
vector input determines the surface direction at each pixel.
The source and normals should be generated together by the 3d program
so they match.

In the Sapphire Lighting effects submenu.

![Light3D](../_static/Light3D.jpg)


### Inputs:

- **Source**: The current layer. The 3d surface colors.

- **Normals**: Defaults to None. Contains the normal vectors matching the Source clip. Typically the red channel will have the X component of the normal, green will have Y, and blue will have Z, but you can adjust this mapping using the Normal Offset and Invert parameters on the second page.

- **Matte**: Defaults to None. Used to interpolate between the original image and the result. Where the matte is black, no lighting is applied and the original Source image is visible.


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

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of all lights together.

- **Ambient Bright** (Default: 0.2, Range: any)
  The amount of ambient light included in the entire frame. This allows parts of the source where no light is falling to be visible.

- **Diffuse Bright** (Default: 0.5, Range: 0.1 or greater)
  Scales the diffuse light from all light sources.

- **Hilight Bright** (Default: 0.8, Range: 0 or greater)
  Scales the brightness of all specular highlights.

- **Hilight Size** (Default: 0.5, Range: 0.1 or greater)
  Adjusts the size of all specular highlights.


### Light1 Parameters:

Light1 Enable:
*Check-box, Default:
*on.Enables the first light source.

Light1 Dir:
*X & Y, Default:
*[-0.806 0.608],
*Range:
*any.The x and y position of the first light source. This parameter can be adjusted using the Light1 Dir Widget.

Light1 Z:
*Default:
*0.5,
*Range:
*any.The z position of the first light source.

Diffuse Bright 1:
*Default:
*0.5,
*Range:
*0.1 or greater.Scales the diffuse brightness for Light 1 only.

Hilight Bright 1:
*Default:
*1,
*Range:
*0 or greater.Scales the brightness of the specular highlights for Light 1 only.

Hilight Size 1:
*Default:
*1,
*Range:
*0 or greater.Adjusts the size of the specular highlights for Light 1 only.

Light1 Color:
*Default rgb:
*[1 1 1].
The color of the first light source.

### Light2 Parameters:

Light2 Enable:
*Check-box, Default:
*off.Enables the second light source.

Light2 Dir:
*X & Y, Default:
*[0.806 0.608],
*Range:
*any.The x and y position of the second light source. This parameter can be adjusted using the Light2 Dir Widget.

Light2 Z:
*Default:
*0.5,
*Range:
*any.The z position of the second light source.

Diffuse Bright 2:
*Default:
*0.5,
*Range:
*0.1 or greater.Scales the diffuse brightness for Light 2 only.

Hilight Bright 2:
*Default:
*1,
*Range:
*0 or greater.Scales the brightness of the specular highlights for Light 2 only.

Hilight Size 2:
*Default:
*1,
*Range:
*0 or greater.Adjusts the size of the specular highlights for Light 2 only.

Light2 Color:
*Default rgb:
*[1 1 1].
The color of the second light source.

### Light3 Parameters:

Light3 Enable:
*Check-box, Default:
*off.Enables the third light source.

Light3 Dir:
*X & Y, Default:
*[-0.806 -0.627],
*Range:
*any.The x and y position of the third light source. This parameter can be adjusted using the Light2 Dir Widget.

Light3 Z:
*Default:
*0.5,
*Range:
*any.The z position of the third light source.

Diffuse Bright 3:
*Default:
*0.5,
*Range:
*0.1 or greater.Scales the diffuse brightness for Light 3 only.

Hilight Bright 3:
*Default:
*1,
*Range:
*0 or greater.Scales the brightness of the specular highlights for Light 3 only.

Hilight Size 3:
*Default:
*1,
*Range:
*0 or greater.Adjusts the size of the specular highlights for Light 3 only.

Light3 Color:
*Default rgb:
*[1 1 1].
The color of the third light source.

### Light4 Parameters:

Light4 Enable:
*Check-box, Default:
*off.Enables the fourth light source.

Light4 Dir:
*X & Y, Default:
*[0.806 -0.627],
*Range:
*any.The x and y position of the fourth light source. This parameter can be adjusted using the Light2 Dir Widget.

Light4 Z:
*Default:
*0.5,
*Range:
*any.The z position of the fourth light source.

Diffuse Bright 4:
*Default:
*0.5,
*Range:
*0.1 or greater.Scales the diffuse brightness for Light 4 only.

Hilight Bright 4:
*Default:
*1,
*Range:
*0 or greater.Scales the brightness of the specular highlights for Light 4 only.

Hilight Size 4:
*Default:
*1,
*Range:
*0 or greater.Adjusts the size of the specular highlights for Light 4 only.

Light4 Color:
*Default rgb:
*[1 1 1].The color of the fourth light source.

Normal Offset:
*Default:
*-0.5,
*Range:
*any.Added to the values in the Normal input.

Normal X <-:
*Popup menu, Default: Red
*.Determines which color channel is used for the horizontal component of the
normal vectors.
*Red:
*Use Red channel.*Green:
*Use Green channel.*Blue:
*Use Blue channel.

Normal Y <-:
*Popup menu, Default: Green
*.Determines which color channel is used for the vertical component of the
normal vectors.
*Red:
*Use Red channel.*Green:
*Use Green channel.*Blue:
*Use Blue channel.

Normal Z <-:
*Popup menu, Default: Blue
*.Determines which color channel is used for the depth component of the
normal vectors.
*Red:
*Use Red channel.*Green:
*Use Green channel.*Blue:
*Use Blue channel.

Invert X:
*Check-box, Default:
*off.If checked, inverts the horizontal compononent of the normal vectors.

Invert Y:
*Check-box, Default:
*off.If checked, inverts the vertical compononent of the normal vectors.

Invert Z:
*Check-box, Default:
*off.If checked, inverts the depth compononent of the normal vectors.

Blur Matte:
*Default:
*0,
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

Show Light1 Dir:
*Check-box, Default:
*on.Turns on or off the screen user interface for adjusting the
Light1 Dir parameter.This parameter only appears on AE and Premiere,
where on-screen widgets are supported.

Show Light2 Dir:
*Check-box, Default:
*on.
Turns on or off the screen user interface for adjusting the
Light2 Dir parameter.This parameter only appears on AE and Premiere,
where on-screen widgets are supported.
