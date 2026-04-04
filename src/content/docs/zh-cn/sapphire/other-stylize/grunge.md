---
title: Grunge
---

## S_Grunge

Simulates many different kinds of grunge including dirt,
stains, flecks, grime, scratches, and paint. Up to three different
kinds of grunge can be combined. There are master controls for
adjusting all grunge together as well as a set of detailed controls
for adjusting the look of each of the grunge collections.

In the Sapphire Render effects submenu.

![Grunge](../_static/Grunge.jpg)


### Inputs:

- **Background**: The current layer. The clip to use as background.

- **Matte**: Defaults to None. If provided, the blur is only performed on regions of the source clip specified by the bright areas of this input. Pixels outside this matte are not blurred, and do not contribute to the resulting blurred pixels within it. This input can be affected using the Invert Matte, or Matte Use parameters.


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

- **Stamp Density** (Default: 500, Range: 0 or greater)
  The overall number of stamps over the frame. Increase for more stamps, decrease to get fewer.

- **Stamp Size** (Default: 1, Range: 0 or greater)
  Scales the overall size of the grunge stamps.

- **Stamp Opacity** (Default: 0.8, Range: 0 to 1)
  The overall opacity of the grunge stamps.

- **Stamp Brightness** (Default: 1, Range: 0 or greater)
  Controls the brightness of the grunge stamps.

- **Stamp Center** (X & Y, Default: [0 0], Range: any)
  The overall position of the center of the grunge stamps.

- **Stamp Uses Mocha** (Check-box, Default: off)
  Controls whether the stamp center is controlled by the Stamp Center parameter or follows the Stamp Center point tracked inside of Mocha.

- **Smooth Stamp Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Fade Softness** (Default: 0.3, Range: 0 to 1)
  Controls how fast individual grunge stamps fade. When set to zero, stamps will pop on and off. When set to one, they will fade in and out. Stamps will fade if they are blinking, along the edge of the grunge framing the background, or along the grey areas of the matte.

- **Blink Amount** (Default: 0, Range: 0 to 1)
  Controls how many stamps are blinking on a single frame. With a blink amount of 0 all stamps will be visible. With a blink amount of 1 approximately half of the stamps will be blinking on a given frame.

- **Blink Coherence** (Default: 0.5, Range: 0 to 1)
  Changes the pattern of the stamps that blink.

- **Blink Speed** (Default: 6, Range: 0 or greater)
  Controls how fast the stamps blink in and out.

- **Frame Amount** (Default: 0.35, Range: 0 to 1)
  Controls the brightness of the grunge inside the frame or border. If frame amount is set to anything other than 0, Grunge will create a picture frame border of grunge around the frame center. At a frame amount of 1, stamps inside the frame are completely invisible and at a frame amount of 0, there is no frame.

- **Frame Softness** (Default: 0.5, Range: 0 or greater)
  The width of the frame's soft edge. Larger values give the grunge frame a softness as the grunge brightness fades on the edges.

- **Frame Center** (X & Y, Default: [0 0], Range: any)
  The position of the center of the grunge frame. This parameter can be adjusted using the Frame Center Widget.

- **Frame Uses Mocha** (Check-box, Default: off)
  Controls whether the frame center is controlled by the Frame Center parameter or follows the Frame Center point tracked inside of Mocha.

- **Smooth Frame Track** (Integer, Default: 0, Range: 0 or greater)
  Controls how many points to average when stabilizing the Mocha point track.

- **Frame Radius** (Default: 1, Range: 0 or greater)
  Distance from the center to apply the grunge frame. This parameter can be adjusted using the Frame Radius Widget.

- **Frame Rel Height** (Default: 0.75, Range: 0.1 or greater)
  The relative vertical size of the frame shape. Increase for a taller shape, decrease for a wider one.

- **Invert Frame** (Check-box, Default: off)
  If enabled, shows the grunge in the center of the frame instead of at the edge of the frame.

- **Stamp1** (Popup menu, Default: Garage Floor)
  The style of grunge to apply. Up to three styles may be selected.
  - **None**: no grunge.
  - **Plaster**: large low detail grunge simulating plaster.
  - **Garage Floor**: large speckled grunge reminiscent of pavement.
  - **Speckles**: groups of small, similar sized, dots of grunge, similar to paint spray.
  - **Paint Spray**: groups of small and medium splashes of paint, similar to speckles.
  - **Paint Splatters**: long splashes of paint.
  - **Hairline Cracks**: long, narrow, wiggly cracks with no branches.
  - **Tile Cracks**: long, straight jagged cracks.
  - **Pavement Cracks**: very branchy cracks with very varied widths.
  - **Hairs**: curly and straight hairs of varying sizes.
  - **Scratches**: straight scratches biased towards the diagonals.
  - **Frost**: patches of frost.
  - **Glass Cracks**: patches of webbed glass cracks.
  - **Clouds**: puffs of clouds, similar to smoke and watercolor drops.
  - **Smoke**: whisps of smoke, similar to clouds and watercolor drops.
  - **Splotches**: paint splotches with lots of streaks of paint spreading outward.
  - **Corrosion**: patches of rust like damage.
  - **Watercolor Drops**: low detail splashes of grunge, similar to clouds and smoke.
  - **Dust**: small particles of grunge varying in shape, similar to flecks.
  - **Stains**: coffee and water stains.
  - **Flecks**: small to medium particles of grunge varying in shape and size, similar to dust.
  - **Circles**: large interrupted rings.
  - **Geometric Lines**: large, sharp lines and angular grunge.
  - **Liquid Lino**: wet looking patches.
  - **Geometric Star Map**: constellation looking grunge.
  - **Moon Dust**: streaks of small dots.
  - **Starfield Spray**: clusters of small dots.
  - **Spray Paint Stars**: clusters of dots and pointed stars.
  - **Spray Paint**: varied clusters of small dots.
  - **Splatter And Spray**: paint smears and splatters.
  - **Hand Painted**: varied patches of paint.
  - **Ink Drops**: varied ink drop patterns.
  - **Large Sponges**: blotted sponge patterns.
  - **Paper Sponges**: crinkly paper.
  - **Decaying Damask**: various faded damask patterns.
  - **Stylized Blooms**: stylized flowers in full bloom.
  - **Oil Paint Floral**: clusters of flowers.

- **Stamp1 Rel Density** (Default: 1, Range: 0 or greater)
  Scales the density for the specific stamp collection.

- **Stamp1 Color1** (Default rgb: [0.01 0.01 0.01])
  Beginning of the range of colors for the stamp collection.

- **Stamp1 Color2** (Default rgb: [0.3 0.3 0.3])
  End of the range of colors for the stamp collection. Each piece of grunge will have a random color between color1 and color2.

- **Stamp1 Rel Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightess for the specific stamp collection.

- **Stamp1 Rel Opacity** (Default: 1, Range: 0 or greater)
  Scales the opacity for the specific stamp collection.

- **Stamp1 Rel Size** (Default: 1, Range: 0 or greater)
  Scales the size for the specific stamp collection.

- **Vary Stamp1 Brightness** (Default: 0, Range: 0 or greater)
  Amount to vary the brightness from one piece of grunge to the next.

- **Vary Stamp1 Opacity** (Default: 0.75, Range: 0 or greater)
  Amount to vary the opacity from one piece of grunge to the next.

- **Vary Stamp1 Size** (Default: 0.5, Range: 0 or greater)
  Amount to vary the size from one piece of grunge to the next.

- **Stamp2** (Popup menu, Default: Paint Spray)
  The style of grunge to apply. Up to three styles may be selected.
  - **None**: no grunge.
  - **Plaster**: large low detail grunge simulating plaster.
  - **Garage Floor**: large speckled grunge reminiscent of pavement.
  - **Speckles**: groups of small, similar sized, dots of grunge, similar to paint spray.
  - **Paint Spray**: groups of small and medium splashes of paint, similar to speckles.
  - **Paint Splatters**: long splashes of paint.
  - **Hairline Cracks**: long, narrow, wiggly cracks with no branches.
  - **Tile Cracks**: long, straight jagged cracks.
  - **Pavement Cracks**: very branchy cracks with very varied widths.
  - **Hairs**: curly and straight hairs of varying sizes.
  - **Scratches**: straight scratches biased towards the diagonals.
  - **Frost**: patches of frost.
  - **Glass Cracks**: patches of webbed glass cracks.
  - **Clouds**: puffs of clouds, similar to smoke and watercolor drops.
  - **Smoke**: whisps of smoke, similar to clouds and watercolor drops.
  - **Splotches**: paint splotches with lots of streaks of paint spreading outward.
  - **Corrosion**: patches of rust like damage.
  - **Watercolor Drops**: low detail splashes of grunge, similar to clouds and smoke.
  - **Dust**: small particles of grunge varying in shape, similar to flecks.
  - **Stains**: coffee and water stains.
  - **Flecks**: small to medium particles of grunge varying in shape and size, similar to dust.
  - **Circles**: large interrupted rings.
  - **Geometric Lines**: large, sharp lines and angular grunge.
  - **Liquid Lino**: wet looking patches.
  - **Geometric Star Map**: constellation looking grunge.
  - **Moon Dust**: streaks of small dots.
  - **Starfield Spray**: clusters of small dots.
  - **Spray Paint Stars**: clusters of dots and pointed stars.
  - **Spray Paint**: varied clusters of small dots.
  - **Splatter And Spray**: paint smears and splatters.
  - **Hand Painted**: varied patches of paint.
  - **Ink Drops**: varied ink drop patterns.
  - **Large Sponges**: blotted sponge patterns.
  - **Paper Sponges**: crinkly paper.
  - **Decaying Damask**: various faded damask patterns.
  - **Stylized Blooms**: stylized flowers in full bloom.
  - **Oil Paint Floral**: clusters of flowers.

- **Stamp2 Rel Density** (Default: 0.5, Range: 0 or greater)
  Scales the density for the specific stamp collection.

- **Stamp2 Color1** (Default rgb: [0.1 0 0])
  Beginning of the range of colors for the stamp collection.

- **Stamp2 Color2** (Default rgb: [0.2 0.1 0.05])
  End of the range of colors for the stamp collection. Each piece of grunge will have a random color between color1 and color2.

- **Stamp2 Rel Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightess for the specific stamp collection.

- **Stamp2 Rel Opacity** (Default: 1, Range: 0 or greater)
  Scales the opacity for the specific stamp collection.

- **Stamp2 Rel Size** (Default: 1, Range: 0 or greater)
  Scales the size for the specific stamp collection.

- **Vary Stamp2 Brightness** (Default: 0, Range: 0 or greater)
  Amount to vary the brightness from one piece of grunge to the next.

- **Vary Stamp2 Opacity** (Default: 0.75, Range: 0 or greater)
  Amount to vary the opacity from one piece of grunge to the next.

- **Vary Stamp2 Size** (Default: 0.5, Range: 0 or greater)
  Amount to vary the size from one piece of grunge to the next.

- **Stamp3** (Popup menu, Default: None)
  The style of grunge to apply. Up to three styles may be selected.
  - **None**: no grunge.
  - **Plaster**: large low detail grunge simulating plaster.
  - **Garage Floor**: large speckled grunge reminiscent of pavement.
  - **Speckles**: groups of small, similar sized, dots of grunge, similar to paint spray.
  - **Paint Spray**: groups of small and medium splashes of paint, similar to speckles.
  - **Paint Splatters**: long splashes of paint.
  - **Hairline Cracks**: long, narrow, wiggly cracks with no branches.
  - **Tile Cracks**: long, straight jagged cracks.
  - **Pavement Cracks**: very branchy cracks with very varied widths.
  - **Hairs**: curly and straight hairs of varying sizes.
  - **Scratches**: straight scratches biased towards the diagonals.
  - **Frost**: patches of frost.
  - **Glass Cracks**: patches of webbed glass cracks.
  - **Clouds**: puffs of clouds, similar to smoke and watercolor drops.
  - **Smoke**: whisps of smoke, similar to clouds and watercolor drops.
  - **Splotches**: paint splotches with lots of streaks of paint spreading outward.
  - **Corrosion**: patches of rust like damage.
  - **Watercolor Drops**: low detail splashes of grunge, similar to clouds and smoke.
  - **Dust**: small particles of grunge varying in shape, similar to flecks.
  - **Stains**: coffee and water stains.
  - **Flecks**: small to medium particles of grunge varying in shape and size, similar to dust.
  - **Circles**: large interrupted rings.
  - **Geometric Lines**: large, sharp lines and angular grunge.
  - **Liquid Lino**: wet looking patches.
  - **Geometric Star Map**: constellation looking grunge.
  - **Moon Dust**: streaks of small dots.
  - **Starfield Spray**: clusters of small dots.
  - **Spray Paint Stars**: clusters of dots and pointed stars.
  - **Spray Paint**: varied clusters of small dots.
  - **Splatter And Spray**: paint smears and splatters.
  - **Hand Painted**: varied patches of paint.
  - **Ink Drops**: varied ink drop patterns.
  - **Large Sponges**: blotted sponge patterns.
  - **Paper Sponges**: crinkly paper.
  - **Decaying Damask**: various faded damask patterns.
  - **Stylized Blooms**: stylized flowers in full bloom.
  - **Oil Paint Floral**: clusters of flowers.

- **Stamp3 Rel Density** (Default: 0.5, Range: 0 or greater)
  Scales the density for the specific stamp collection.

- **Stamp3 Color1** (Default rgb: [0 0 0])
  Beginning of the range of colors for the stamp collection.

- **Stamp3 Color2** (Default rgb: [0.25 0.25 0.25])
  End of the range of colors for the stamp collection. Each piece of grunge will have a random color between color1 and color2.

- **Stamp3 Rel Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightess for the specific stamp collection.

- **Stamp3 Rel Opacity** (Default: 1, Range: 0 or greater)
  Scales the opacity for the specific stamp collection.

- **Stamp3 Rel Size** (Default: 1, Range: 0 or greater)
  Scales the size for the specific stamp collection.

- **Vary Stamp3 Brightness** (Default: 0, Range: 0 or greater)
  Amount to vary the brightness from one piece of grunge to the next.

- **Vary Stamp3 Opacity** (Default: 0.75, Range: 0 or greater)
  Amount to vary the opacity from one piece of grunge to the next.

- **Vary Stamp3 Size** (Default: 0.5, Range: 0 or greater)
  Amount to vary the size from one piece of grunge to the next.

- **Emboss Bumps Scale** (Default: 0.15, Range: 0 or greater)
  Scales the amplitude of the bumps.

- **Emboss Light Angle** (Default: 135, Range: any)
  Adjusts the angle of light for the emboss. This parameter can be adjusted using the Emboss Light Angle Widget.

- **Emboss Smooth** (Default: 0.0001, Range: 0 or greater)
  Smooths the small details of the image so they don't get embossed as strongly as the big features. Set to 0 to emboss all details. Increase to smooth out more details.

- **Emboss Threshold** (Default: 0.5, Range: -0.5 to 1.5)
  Any grunge brighter the threshold will be a bump and any grunge darker than the threshold will be a pit.

- **Blur Grunge** (Default: 0, Range: 0 or greater)
  Blurs the grunge. Increase for more blur. Doesn't affect the background.

- **Blur Grunge Rel** (X & Y, Default: [1 1], Range: 0 or greater)
  Scales the width of the blur.

- **Shake Amplitude** (Default: 0, Range: 0 or greater)
  The amount of random shaking motion that is applied to the stamps.

- **Shake Amplitude Rel X** (Default: 1, Range: 0 or greater)
  The relative horizontal amount of shaking.

- **Shake Amplitude Rel Y** (Default: 1, Range: 0 or greater)
  The relative vertical amount of shaking.

- **Shake Frequency** (Default: 60, Range: 0 or greater)
  Increase for faster shaking, decrease for slower shaking. (Be careful if you animate frequency values because the resulting shake frequency is also affected by the rate of change of the value.)

- **Per Stamp Amplitude** (Default: 0, Range: 0 or greater)
  The amount of random shaking applied independently to each stamp.

- **Per Stamp Frequency** (Default: 5, Range: 0 or greater)
  The frequency of per-stamp shaking. Increase for faster shaking, decrease for slower shaking.

- **Combine** (Popup menu, Default: Comp)
  Determines how the grunge image is combined with the background.
  - **Grunge Only**: gives only the grunge image with no background.
  - **Comp**: composites the grunge image over the background.
  - **Mult**: this can be used as an 'intersection' operation
on matte images. White is the identity for Multiply, where one
image contains white the other is not affected, so the result only
contains white where both inputs are white.
  - **Add**: causes the grunge image to be added to the background.
  - **Screen**: performs a blend function which can help prevent
overly bright results.
  - **Difference**: similar to Subtract but the absolute value of
the result is used, which tends to give more resulting colors in
bounds. This can be used to select the regions of two matte
images where one or the other is white, but not both.
  - **Subtract**: subtracts the grunge image from the background.
  - **Overlay**: combines foreground and background using an
overlay function.
  - **Hard Light**: similar to overlay but with foreground and
background swapped.
  - **Soft Light**: darkens or lightens the background depending on
the foreground.
  - **Color Dodge**: brightens the background depending on the grunge image.
  - **Color Burn**: darkens the background depending on the grunge image.
  - **Darken**: the minimum of grunge image and background. This can
also be used as an 'intersection' operation with slightly different
results than Multiply.
  - **Lighten**: the maximum of grunge image and background. This
can also be used as a 'union' operation with slightly different
results than Screen.
  - **Exclusion**: similar to Difference but with smoother results.
  - **Linear Dodge**: adds the grunge image to the background and clamps the
result at white.
  - **Linear Burn**: adds the grunge image to the background but offsets to
make the result darker. Similar to multiply in that combining with
white gives no change and combining with black gives black.
  - **Linear Light**: performs a linear burn or linear dodge
depending on if the foreground is more or less than 50 percent gray.

- **Use Bg Alpha** (Check-box, Default: off)
  When enabled, ignore the alpha generated by the grunge and use the alpha from the background.

- **Scale Background** (Default: 1, Range: 0 or greater)
  Scales the brightness of the background before combining with the grunges. If 0, the result will contain only the grunge image over black.

- **Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator for the stamp positioning, size, and variation. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

- **Blur Matte** (Default: 0, Range: 0 or greater)
  Blurs the Matte input by this amount before using. This can provide a smoother transition between the matted and unmatted areas. It has no effect unless the Matte input is provided.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Matte Use** (Popup menu, Default: Luma)
  Determines how the Matte input channels are used to make a monochrome matte.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Crop Grunge To Matte** (Check-box, Default: off)
  Controls whether the grunge has a sharp edge along the edge of the mask. If enabled, the grunge will crop to the mask. This allows grunge to be applied to specific sections of the input clip. If disabled, each grunge stamp will be checked against the mask. If most of the grunge stamp falls inside the mask, the entire grunge stamp will be added. If most of the grunge stamp falls outside the mask, the entire grunge stamp will be hidden. This allows for a Vignette like application of grunge.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Flip Stamps Vertically** (Check-box, Default: off)
  Flips stamps vertically if needed to achieve a consistent look.

- **Show Emboss Light Angle** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Frame Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Frame Radius** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Frame Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Frame Center** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Frame Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

