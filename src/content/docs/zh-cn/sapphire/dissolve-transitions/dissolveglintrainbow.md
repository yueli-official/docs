---
title: DissolveGlintRainbow
---

## S_DissolveGlintRainbow

Transitions between two input clips using a bright glowing glint.
The clips dissolve into each other, while each one gets a glint which
ramps up and down over the duration of the effect.
The Dissolve Percent parameter should be animated
to control the transition speed.

In the Sapphire Transitions effects submenu.

![DissolveGlintRainbow](../_static/DissolveGlintRainbow.jpg)


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
  The speed of the dissolve between the From and To clips. When set to 1, the dissolve takes place over the entire duration of the effect. When set higher, the dissolve is shorter, although the glint ramp-up and ramp-down still takes the entire duration. Setting this to 10 can make the transition snappier and more like a flash-frame cut.

- **Glint Brightness** (Default: 1.5, Range: 0 or greater)
  The maximum brightness of the glint in the middle of the transition.

- **Glint Threshold** (Default: 0.7, Range: 0 or greater)
  Glints are generated from locations in the From and To clips there are brighter than this valuye. A value of 0.9 causes glints at only the brightest spots. A value of 0 causes glints for every non-black area.

- **Glint Threshold Blur** (Default: 0.0896, Range: 0 or greater)
  Increase to smooth out the areas creating glints. This can be used to eliminate glints generated from small speckles or to simply soften the glints. Increasing this may put more highlights below the threshold and darken the resulting glints, but you can decrease the Threshold parameter to compensate.

- **Glint Scale Colors** (Default rgb: [1 1 1])
  Scales the color of the glints. The colors and brightnesses of the glints are also affected by the From and To inputs.

- **Brightness X** (Default: 1, Range: 0 or greater)
  Scales the brightness of the horizontal glint rays.

- **Brightness Y** (Default: 1, Range: 0 or greater)
  Scales the brightness of the vertical glint rays.

- **Brightness Diag1** (Default: 1, Range: 0 or greater)
  Scales the brightness of the diagonal rays from top right to bottom left.

- **Brightness Diag2** (Default: 1, Range: 0 or greater)
  Scales the brightness of the diagonal rays from top left to bottom right.

- **Glint Size** (Default: 2, Range: 0 or greater)
  The maximum size of the glint at the middle of the transition.

- **Glint Shrink** (Default: 0.8, Range: 0 to 1)
  The fraction by which the glint size is reduced at the beginning and end of the transition.

- **Size X** (Default: 1, Range: 0 or greater)
  Scales the length of the horizontal glint rays.

- **Size Y** (Default: 1, Range: 0 or greater)
  Scales the length of the vertical glint rays.

- **Size Diag1** (Default: 0.75, Range: 0 or greater)
  Scales the length of the diagonal rays from top left to bottom right.

- **Size Diag2** (Default: 0.75, Range: 0 or greater)
  Scales the length of the diagonal rays from top right to bottom left.

- **Shift Out** (Default: 1, Range: any)
  Shifts the glint rays outwards from their source highlights by this amount relative to the glint size.

- **Shift Red** (Default: 0.3, Range: any)
  Shifts the red component of the glints in or out relative to the blue. The green is centered between blue and red for a complete spectrum.

- **Shift Blue** (Default: -0.3, Range: any)
  Shifts the blue component of the glints in or out relative to the red and green. This can be used with Shift Red to adjust the range of hues in the glints.


### Rel From Parameters:

Rel From Brightness:
*Default:
*1,
*Range:
*0 or greater.Relative brightness of the glint on the outgoing (From) clip.

Rel From Size:
*Default:
*1,
*Range:
*0 or greater.Relative size of the glint on the outgoing (From) clip.

From Offset Threshold:
*Default:
*0,
*Range:
*any.Extra threshold to apply to the glint on the outgoing (From) clip.

Rel From Color:
*Default rgb:
*[1 1 1].
Relative color of the glint on the outgoing (From) clip.

### Rel To Parameters:

Rel To Brightness:
*Default:
*1,
*Range:
*0 or greater.Relative brightness of the glint on the incoming (To) clip.

Rel To Size:
*Default:
*1,
*Range:
*0 or greater.Relative size of the glint on the incoming (To) clip.

To Offset Threshold:
*Default:
*0,
*Range:
*any.Extra threshold to apply to the glint on the incoming (To) clip.

Rel To Color:
*Default rgb:
*[1 1 1].Relative color of the glint on the incoming (To) clip.

Affect Alpha:
*Default:
*1,
*Range:
*0 or greater.If this value is positive the output Alpha channel will
include some opacity from the dissolves. The maximum of the red, green,
and blue dissolve brightness is scaled by this value and combined with
the background Alpha at each pixel.

Expand Borders:
*Check-box, Default:
*off.If enabled, transparent borders are added to the
input image before processing. This allows the result to include soft
edges beyond the original image size. When off, the effect only
occurs within the frame and the result will retain an edge
at the borders.

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

Swap Diagonals:
*Check-box, Default:
*off.Flips glints vertically if needed to achieve a consistent look.

Show Glint Size:
*Check-box, Default:
*on.
Turns on or off the screen user interface for adjusting the
Glint Size parameter.This parameter only appears on AE and Premiere,
where on-screen widgets are supported.
