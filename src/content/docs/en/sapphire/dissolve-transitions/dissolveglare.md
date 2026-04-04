---
title: DissolveGlare
---

## S_DissolveGlare

Transitions between two input clips using animated glares.
The clips dissolve into each other, and glares are added to the result.
The glare size and brightness ramps up and down over the duration of the effect.

In the Sapphire Transitions effects submenu.

![DissolveGlare](../_static/DissolveGlare.jpg)


### Inputs:

- **Foreground**: The current layer. Starts the transition with this clip.

- **Background**: Defaults to None. Ends the transition with this clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Transition Dir** (Popup menu, Default: Dissolve Off to Bg)
  Selects the direction of the transition.
  - **Dissolve Off to Bg**: transitions from the current layer to the Background.
  - **Dissolve On from Bg**: transitions from the Background to the current layer.

- **Auto Trans** (Popup YES-NO, Default: No)
  If enabled, a transition is performed automatically between the first and last frames of the layer. If this is off, the transition is performed manually by animating the Dissolve Percent parameter.

- **Dissolve Percent** (Default: 0, Range: 0 to 1)
  Auto Trans must be disabled for this parameter to be used. It determines the transition ratio between the Foreground and Background inputs, and would normally be animated from 0 to 100 to perform a complete transition. The curve controlling this parameter can be adjusted for more detailed control over the timing of the dissolve.

- **Dissolve Speed** (Default: 3, Range: 1 or greater)
  The speed of the dissolve between the From and To clips. When set to 1, the dissolve takes place over the entire duration of the effect. When set higher, the dissolve is shorter, although the edge rays ramp-up and ramp-down still takes the entire duration. Setting this to 10 can make the transition snappier and more like a flash-frame cut.

- **Size** (Default: 2.4, Range: 0 or greater)
  Scales the size of the glares.

- **Rel Height** (Default: 1, Range: 0 or greater)
  Scales the vertical dimension of the glares, making them elliptical instead of circular.

- **Style** (Default: 0, Range: 0 or greater)
  The style of glare to apply. Custom glare types can also be made, or existing types modified, by editing the "s_glares.text" file.

- **Convolve** (Check-box, Default: off)
  Determines the method for applying the glares to the Background.

- **Threshold** (Default: 0.5, Range: 0 or greater)
  Dissolves are generated from locations in the source clip that are brighter than this value. A value of 0.9 causes dissolves at only the brightest spots. A value of 0 causes dissolves for every non-black area.

- **Threshold Add Color** (Default rgb: [0 0 0])
  This can be used to raise the threshold on a specific color and thereby reduce the dissolves generated on areas of the source clip containing that color.

- **Threshold Blur** (Default: 0, Range: 0 or greater)
  Increase to smooth out the areas creating dissolves. This can be used to eliminate dissolves generated from small speckles or to simply soften the dissolves. Increasing this may put more highlights below the threshold and darken the resulting dissolves, but you can decrease the Threshold parameter to compensate.

- **Brightness** (Default: 3, Range: 0 or greater)
  Scales the brightness of all the dissolves.

- **Scale Colors** (Default rgb: [1 1 1])
  Scales the color of the dissolves. The colors and brightnesses of the dissolves are also affected by the Source and Matte inputs.

- **Saturation** (Default: 1, Range: -2 to 8)
  Scales the color saturation of the glare elements. Increase for more intense colors. Set to 0 for monochrome glares.

- **Rotate** (Default: 0, Range: any)
  Rotates the ray elements of the glares, if any, in degrees.

- **Rays Num Scale** (Default: 1, Range: 0 or greater)
  Increases or decreases the number of rays.

- **Rays Length** (Default: 1, Range: 0 or greater)
  Adjusts the length of the rays without changing their thickness.

- **Rays Thickness** (Default: 1, Range: 0 or greater)
  Adjusts the thickness of the individual rays.

- **Blur Glare** (Default: 0, Range: 0 or greater)
  The glare is blurred by this amount before being combined with the background.

- **Hue Shift** (Default: 0, Range: any)
  Shifts the hue of the glare, in revolutions from red to green to blue to red.

- **Glare Res** (Popup menu, Default: Full)
  Selects the resolution factor for the glares. Higher resolutions give sharper glares, lower resolutions give smoother glares and faster processing. This 'Res' factor only affects the glares: the background is still combined with the glares at full resolution.
  - **Full**: Full resolution is used.
  - **Half**: The glares are calculated at half resolution.
  - **Quarter**: The glares are calculated at quarter resolution.

- **Affect Alpha** (Default: 1, Range: 0 or greater)
  If this value is positive the output Alpha channel will include some opacity from the dissolves. The maximum of the red, green, and blue dissolve brightness is scaled by this value and combined with the background Alpha at each pixel.

- **Glare From Alpha** (Default: 0, Range: 0 to 1)
  Set to 1 to generate dissolves from the alpha channel of the source input instead of the RGB channels. In this case the dissolves will not pick up color from the source and will typically be brighter. Values between 0 and 1 interpolate between using the RGB and the Alpha.

- **Expand Borders** (Check-box, Default: off)
  If enabled, transparent borders are added to the input image before processing. This allows the result to include soft edges beyond the original image size. When off, the effect only occurs within the frame and the result will retain an edge at the borders.

- **Show Size** (Check-box, Default: on)
  Turns on or off the screen user interface widget for adjusting the Size and Rel Height parameters.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

