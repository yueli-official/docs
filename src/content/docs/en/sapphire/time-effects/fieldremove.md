---
title: FieldRemove
---

## S_FieldRemove

Adaptively removes video field interlacing artifacts
from areas with motion, without blurring the stationary parts of the
image. A 'Motion Matte' is generated internally and the moving areas
are deinterlaced with the usual loss of vertical resolution, but the
stationary areas are not deinterlaced and should remain sharp.

In the Sapphire Time effects submenu.

![FieldRemove](../_static/FieldRemove.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

- **Mask**: Defaults to None. Interpolate between the result and the Source input. White areas use the result of the effect. Black areas use the Source clip.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: Same Speed)
  Selects speed-change options.
  - **Same Speed**: No change in speed.
  - **NTSC to Film**: Converts 60 field/sec input to 24 frame/sec
output. Every 5 frames of input are converted to 4 frames of
output, so in this mode only 4/5 of your output clip will be
useful.
  - **Half Speed**: Every field of input is converted to one frame
of output. In this mode, you should normally first pad your input
clip to make it twice as long, so the correct number of output frames
will be generated.

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

- **Scale Mo Matte** (Default: 4, Range: 0 or greater)
  Increase to remove more field artifacts, or decrease to remove fewer and keep the image sharper.

- **Threshold Matte** (Default: 0.05, Range: 0 or greater)
  This value is subtracted from the Motion Matte and can be increased to reduce unwanted deinterlacing due just to noise.

- **Blur Mo Matte** (Default: 0.112, Range: 0 or greater)
  Determines how much the Motion Matte is smoothed out to avoid sharp transitions between the interlaced and deinterlaced areas.

- **Show** (Popup menu, Default: Result)
  Selects the output option.
  - **Result**: output the deinterlaced result normally.
  - **MotionMatte**: this allows viewing the Motion Matte itself, and can
be helpful when adjusting the other parameters above.

- **Use Field** (Popup menu, Default: Lower)
  Selects which field to preserve in areas with field artifacts. This parameter only has an affect when using Same Speed mode.
  - **Lower**: keeps the lower field.
  - **Upper**: keeps the upper field.
  - **Merge**: Uses the average of both fields.

- **Field Dominance** (Popup menu, Default: Lower First)
  Selects the ordering of the output fields. This parameter only has an affect when NOT using Same Speed mode.
  - **Lower First**: The lower field is first in time.
  - **Upper First**: The upper field is first in time.

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

