---
title: DissolveLensFlare
---

## S_DissolveLensFlare

Transitions between two input clips using an animated lens flare.
The clips dissolve into each other, while a lens flare moves along a straight
line. The lens flare grows and shrinks over the duration of the effect.
The Dissolve Percent parameter should be animated
to control the transition speed.

In the Sapphire Transitions effects submenu.

![DissolveLensFlare](../_static/DissolveLensFlare.jpg)


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
  The speed of the dissolve between the From and To clips. When set to 1, the dissolve takes place over the entire duration of the effect. When set higher, the dissolve is shorter, although the lens flare still changes size and brightness over the entire duration. Setting this to 10 can make the transition snappier and more like a flash-frame cut.

- **Hotspot Center** (X & Y, Default: [0 0], Range: any)
  The location through which the brightest spot of the flare passes at the center of the transition.

- **Hotspot Speed** (Default: 1, Range: 0 to 2)
  The speed at which the flare sweeps across the screen. Set this to zero to make the lens flare grow and shrink in place.

- **Hotspot Angle** (Default: -25, Range: any)
  The angle at which the flare sweeps across the screen.

- **Pivot** (X & Y, Default: [0 0], Range: any)
  The elements of the flare will be in a line between the Hotspot and the Pivot locations. The Pivot location is in screen coordinates.

- **Flare Brightness** (Default: 8, Range: 0 or greater)
  The maximum brightness of the flare at the center of the transition.

- **Flare Fade** (Default: 1, Range: 0 to 1)
  The fraction by which the brightness is reduced at the beginning and end of the transition.

- **Flare Width** (Default: 2.5, Range: 0 or greater)
  The maximum width of the flare at the center of the transition.

- **Flare Shrink** (Default: 0.5, Range: 0 to 1)
  The fraction by which the flare width is reduced at the beginning and end of the transition.

- **Rel Heights** (Default: 1, Range: 0 or greater)
  Scales the vertical dimension of all the flare elements, making them elliptical instead of circular. This can also be adjusted using the Scale Widths Widget.

- **Lens** (Default: 0, Range: 0 or greater)
  The type of lens flare to apply. Custom lens flare types can also be made, or existing types modified, by editing the flare in the flare designer.


### Flare Details Parameters:

Rays Rotate:
*Default:
*0,
*Range:
*any.Rotates the ray elements of the lens flare, if any,
in degrees.

Color:
*Default rgb:
*[1 1 1].Scales the color of all flare elements.

Gamma:
*Default:
*1,
*Range:
*0 or greater.Increasing gamma brightens the flare, and especially boosts
the darker elements.

Saturation:
*Default:
*1,
*Range:
*any.Scales the color saturation of the flare elements.
Increase for more intense colors. Set to 0 for a monochrome
lens flare.

Hue Shift:
*Default:
*0,
*Range:
*-1 to 1.Shifts the hue of the flare, in revolutions
from red to green to blue to red.

Hotspot Color:
*Default rgb:
*[1 1 1].Scales the color of the hotspot elements only.

Hotspot Brightness:
*Default:
*1,
*Range:
*0 or greater.Scales the brightness of the hotspot elements only.

Rays Brightness:
*Default:
*1,
*Range:
*0 or greater.Scales the brightness of the ray elements only.

Rays Num Scale:
*Default:
*1,
*Range:
*0 or greater.Increases or decreases the number of rays.

Rays Length:
*Default:
*1,
*Range:
*0 or greater.Adjusts the length of the rays without changing their thickness, or
changing the size of the other flare elements.

Rays Thickness:
*Default:
*1,
*Range:
*0 or greater.Adjusts the thickness of the individual rays within the flare.

Other Brightness:
*Default:
*1,
*Range:
*0 or greater.Scales the brightness of all flare elements that
are NOT at the hotspot location.

Other Width:
*Default:
*1,
*Range:
*0 or greater.Scales the width of all flare elements that
are NOT at the hotspot location.

Other Color:
*Default rgb:
*[1 1 1].Scales the color of all flare elements that
are NOT at the hotspot location.

Blur Flare:
*Default:
*0,
*Range:
*0 or greater.
If positive, the flare image is blurred by this amount
before being combined with the background.

### Other Parameters:

Bg Brightness:
*Default:
*1,
*Range:
*0 or greater.Scales the brightness of the background before
combining with the flare. If 0, the result will contain only the
flare image over black.

Combine:
*Popup menu, Default: Screen
*.Determines how the flare image is combined with the Background.
*Screen:
*performs a blend function which can help prevent
overly bright results.*Add:
*causes the flare image to be added to the background.

Tint Bg Whites:
*Check-box, Default:
*off.If this is enabled, the chroma of the flare
is added only after the result is clamped to the maximum brightness.
This allows the color of the flare image to still be visible even
over bright white backgrounds. For the majority of backgrounds
there will be no observable difference.

Affect Alpha:
*Default:
*1,
*Range:
*0 or greater.If this value is positive the output Alpha channel
will include some opacity from the flare. The maximum of the red,
green, and blue flare brightness is scaled by this value and combined
with the Background Alpha at each pixel.

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

Show Flare Width:
*Check-box, Default:
*on.Turns on or off the screen user interface for adjusting the
Hotspot Center parameter.This parameter only appears on AE and Premiere,
where on-screen widgets are supported.

Show Hotspot Center:
*Check-box, Default:
*on.Turns on or off the screen user interface for adjusting the
Hotspot Center parameter.This parameter only appears on AE and Premiere,
where on-screen widgets are supported.

Show Hotspot Angle:
*Check-box, Default:
*on.Turns on or off the screen user interface for adjusting the
Hotspot Center parameter.This parameter only appears on AE and Premiere,
where on-screen widgets are supported.

Show Rays Rotate:
*Check-box, Default:
*off.Turns on or off the screen user interface for adjusting the
Hotspot Center parameter.This parameter only appears on AE and Premiere,
where on-screen widgets are supported.

Show Pivot:
*Check-box, Default:
*on.
Turns on or off the screen user interface for adjusting the
Pivot parameter.This parameter only appears on AE and Premiere,
where on-screen widgets are supported.See general info for
[Motion Blur](/en/sapphire/#motion-blur)
