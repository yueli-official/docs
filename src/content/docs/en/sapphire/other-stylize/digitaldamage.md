---
title: DigitalDamage
---

## S_DigitalDamage

Simulates bad digital TV transmission with many
options, including freeze-frames, shifting and flowing blocks,
various kinds of blocky noise, and pixelization. Can give looks
similar to MPEG stream errors, digital dropouts, and satellite feed
data corruption.

In the Sapphire Stylize effects submenu.

![DigitalDamage](../_static/DigitalDamage.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.

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

- **Intensity** (Default: 1, Range: 0 or greater)
  Overall spatial intensity of the damage. Turn up to get more damage on each frame.

- **Time Intensity** (Default: 1, Range: 0 or greater)
  Temporal intensity. Normally, not all damage types are applied to each frame. Turning this up will damage more frames with each damage type.

- **Damage Size** (Default: 1, Range: 0.001 or greater)
  Turn up to increase the average size of the damage areas.

- **Damage Size Rel X** (Default: 1, Range: 0.001 or greater)
  Turn up to elongate the damage areas horizontally. This doesn't stretch the image, just changes the aspect ratio of the damage areas and noise patterns.

- **Freeze** (Check-box, Default: on)
  Enable freeze-frame damage.

- **Freeze Threshold** (Default: 0.09, Range: 0 to 1)
  Decrease for more frozen areas on each frame.

- **Freeze Saturation** (Default: 2.5, Range: 0 or greater)
  Boost the saturation of frozen areas, for a more damaged look.

- **Freeze Quality** (Default: 0.1, Range: 0 or greater)
  Reduce to give the frozen areas a JPEG-quantized look.

- **Freeze Errs** (Default: 0.05, Range: 0 or greater)
  Adds JPEG quantization errors to the frozen areas.

- **Freeze Frames** (Integer, Default: 10, Range: 0 to 20)
  Freeze every N frames. This doesn't freeze the whole image, but when there is freeze-damage present, it uses every Nth frame. Turn this up for a more extreme look. Turn it to 1 to get no freezing.

- **Freeze Blink Freq** (Default: 15, Range: 0.1 or greater)
  Controls how fast this type of damage blinks on and off.

- **Freeze Always** (Default: 0.3, Range: 0 to 1)
  Controls how often this type of damage occurs.

- **Shift** (Check-box, Default: on)
  Enable block-shifting damage.

- **Shift Amount** (Default: 0.1, Range: 0 or greater)
  Controls the intensity or amount of the block-shifting damage.

- **Shift Threshold** (Default: 0.3, Range: 0 to 1)
  Decrease for more block-shifted areas on each frame.

- **Shift Blink Freq** (Default: 10, Range: 0.1 or greater)
  Controls how fast this type of damage blinks on and off.

- **Shift Always** (Default: 0.3, Range: 0 to 1)
  Controls how often this type of damage occurs.

- **Brights Noise** (Check-box, Default: on)
  Enable noise that appears in the bright areas of the image.

- **Brights Threshold** (Default: 0.6, Range: 0 to 1)
  Areas brighter than this will be subject to brights-noise.

- **Brights Band Threshold** (Default: 0.4, Range: 0 to 1)
  This damage type occurs in bands; increase this param to make more damage bands, and thus increase the amount of overall damage.

- **Brights Band Freq** (Default: 6, Range: 0.1 or greater)
  Controls the average height of the damage bands; decrease for larger bands, increase for shorter, finer bands.

- **Brights Blink Freq** (Default: 10, Range: 0.1 or greater)
  Controls how fast this type of damage blinks on and off.

- **Brights Always** (Default: 0.23, Range: 0 to 1)
  Controls how often this type of damage occurs.

- **Pixelate** (Check-box, Default: on)
  Enable pixelation damage.

- **Pixelate Frequency** (Default: 40, Range: 1 or greater)
  Controls the size of the blocky pixels. Increase for more, smaller pixels; decrease for fewer, larger pixels.

- **Pixelate Hold** (Default: 0.95, Range: 0 to 1)
  Controls how the pixelate damage areas move. Increase to make the damage areas stay in one place over more frames; decrease to make it more random.

- **Pixelate Threshold** (Default: 0.1, Range: 0 to 1)
  Decrease for more overall pixelation per frame; increase for less.

- **Pixelate Overdrive** (Default: 1.6, Range: 0 or greater)
  Pixelate damage can invert and distort the damaged area; increase this param to make it look more damaged.

- **Pixelate Blink Freq** (Default: 10, Range: 0.1 or greater)
  Controls how fast this type of damage blinks on and off.

- **Pixelate Always** (Default: 0.15, Range: 0 to 1)
  Controls how often this type of damage occurs.

- **Block Noise** (Check-box, Default: on)
  Enable blocky-noise damage; this is commonly seen with bad satellite TV transmission. Bands and blocks of noise overlay and interact with the source footage.

- **Blocks Intensity** (Default: 1, Range: 0 or greater)
  Increase the intensity of the block damage.

- **Blocks Threshold** (Default: 0.4, Range: 0 to 1)
  Decrease for more overall damage per frame; increase for less.

- **Blocks Softness** (Default: 0.2, Range: 0 or greater)
  This param softens the damage pattern as it's increased.

- **Blocks Chroma** (Default: 0.95, Range: 0 to 1)
  Increase to overdrive the chroma of the block noise and make it look more damaged.

- **Blocks Blink Freq** (Default: 10, Range: 0.1 or greater)
  Controls how fast this type of damage blinks on and off.

- **Blocks Always** (Default: 0.15, Range: 0 to 1)
  Controls how often this type of damage occurs.

- **Blocks Affect Alpha** (Default: 0, Range: 0 or greater)
  Controls whether the blocky noise affects the alpha channel of the output. Normally you should set this to zero when using on a text layer or a key, so the blocks don't appear all over the image in the final result, but stay within the text.

- **Invert** (Check-box, Default: on)
  Enable image-inverting damage that inverts and recolors bands of the image.

- **Invert Threshold** (Default: 0.35, Range: 0 to 1)
  Decrease for more overall damage per frame; increase for less.

- **Invert Darken** (Default: 0.4, Range: 0 or greater)
  Increase to darken the inverted areas more; makes them stand out more and look more damaged.

- **Invert Pattern Freq** (Default: 2, Range: 0.1 or greater)
  Controls the spatial frequency of the invert damage pattern. Increase to make the inverted area pattern finer; decrease for larger areas of damage.

- **Invert Blink Freq** (Default: 10, Range: 0.1 or greater)
  Controls how fast this type of damage blinks on and off.

- **Invert Always** (Default: 0.21, Range: 0 to 1)
  Controls how often this type of damage occurs.

- **Flow** (Check-box, Default: on)
  Enable image-flow damage, similar to MPEG loss of I frames. Areas of the image sometimes freeze and start moving as a block.

- **Flow Block Freq** (Default: 4, Range: 0.1 or greater)
  Controls the spatial frequency of the flow damage pattern. Increase to make the flow areas smaller; decrease for larger areas of damage.

- **Flow Damage Amount** (Default: 1, Range: 0 or greater)
  Controls the amount of chroma damage in the flowed areas.

- **Flow Threshold** (Default: 0.6, Range: 0 to 1)
  Decrease for more overall damage per frame; increase for less.

- **Flow Speed** (Default: 1, Range: 0 or greater)
  Controls the motion speed of the flowed areas.

- **Flow Blink Freq** (Default: 8, Range: 0.1 or greater)
  Controls how fast this type of damage blinks on and off.

- **Flow Always** (Default: 0.2, Range: 0 to 1)
  Controls how often this type of damage occurs.

- **Seed** (Default: 1.23, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

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

