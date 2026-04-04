---
title: Vignette
---

## S_Vignette

Darkens the border areas of the source clip to create a
vignette effect. Use the Squareness, Radius, and Edge Softness
parameters to affect the shape of the vignette. Use the Opacity and
Color parameters to adjust its strength and color.

In the Sapphire Stylize effects submenu.

![Vignette](../_static/Vignette.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: Vignette)
  Selects between several variations on generating the Vignette shape.
  - **Vignette**: The Vignette shape and location are defined by specific parameters.
  - **VignetteMocha**: The Vignette shape and location are defined using a Mocha Mask.

- **Center** (X & Y, Default: [0 0], Range: any)
  The center location of the vignette effect. This parameter can be adjusted using the Center Widget.

- **Squareness** (Default: 0, Range: 0 to 1)
  Determines how square the vignette shape is. Set to 1.0 for a square or rectangle shape. Set to 0 for a circle or ellipse. Values in between give rectangles with rounded corners by varying amounts.

- **Radius** (Default: 0.9, Range: 0 or greater)
  Distance from the center to apply the vignette. This parameter can be adjusted using the Radius Widget.

- **Rel Height** (Default: 0.75, Range: 0.05 or greater)
  The relative vertical size of the vignette shape. Increase for a taller shape, decrease for a wider one.

- **Rel Width** (Default: 1, Range: 0.05 or greater)
  The relative horizontal size of the vignette shape. Increase for a wider shape, decrease for a taller one.

- **Rotate** (Default: 0, Range: any)
  Rotation in degrees of the vignette shape. Note that rotation will have no effect if Squareness is zero, and Rel Width and Rel Height are equal. This parameter can be adjusted using the Rotate Widget.

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

- **Edge Softness** (Default: 1, Range: 0 or greater)
  The width of the vignette's soft edge. Larger values give softer, less visible edges.

- **Smooth Curve** (Default: 0.4, Range: 0 to 1)
  If zero, a linear gradient is used across the screen in the soft edge area. Increase this value to use a smoother 'S' shaped curve for interpolation which can reduce the visual perception of the gradient's start and end locations.

- **Color** (Default rgb: [0 0 0])
  The color of the vignette.

- **Opacity** (Default: 1, Range: 0 or greater)
  The opacity of the vignette; animate to 0 to fade the vignette out.

- **Blur Amount** (Default: 0, Range: 0 or greater)
  Blurs the borders of the image in addition to darkening them.

- **Blur Inside** (Check-box, Default: off)
  If checked, the center (undarkened) area of the image is blurred instead of the border.

- **Source Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the source clip. To see only the vignette, set this to zero.

- **Combine** (Popup menu, Default: Composite)
  Determines how the vignette is combined with the Source.
  - **Composite**: composites the vignette over the source clip.
  - **Mult**: the vignette color is multiplied by the source clip.
If the Color is not black, this will selectively colorize the
vignette area.
  - **Add**: the vignette color is added to the source clip. This
will have no effect if the vignette color is black.
  - **Screen**: the vignette color is combined with the source
clip using a screen operation. This will have no effect if the
vignette color is black.
  - **Subtract Inv**: the inverse of the vignette color is
subtracted from the source clip. Inverse means white for black,
yellow for blue, and so on. This mode looks similar to Mult, but a
bit more severe; it crushes the blacks and leaves the highlights
more. This will have no effect if the vignette color is white.
  - **Vignette Only**: shows the vignette pattern without the
source clip. The output will be white where the amount of vignetting
is greatest (e.g. where the source clip would be darkened
completely).
  - **Vignette Only Inv**: shows the inverted vignette pattern
without the source clip. The output will be white where there is no
vignetting (e.g. where the source clip would not be darkened at
all).

- **Show Radius** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Rotate** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Center** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

