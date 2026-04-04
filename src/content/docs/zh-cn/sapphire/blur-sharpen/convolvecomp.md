---
title: ConvolveComp
---

## S_ConvolveComp

Convolves front and back images with a kernel, and
composites them using a matte. Convolution is a mathematical
operator which uses one image, the kernel, as a filter shape for
another image (the source). Convolution effectively stamps a copy
of the kernel at each point of the source, using the source's
brightness at that point. The effect is that a copy of the kernel
will appear over all the bright spots of the source. A kernel image
shaped like a circle or polygon will give an effect similar to
RackDefocusComp; a kernel image shaped like a starburst can give
something like GlareComp.

The kernel size can vary between front and back so either or both
can be blurred.

In the Sapphire Blur+Sharpen effects submenu.

![ConvolveComp](../_static/ConvolveComp.jpg)


### Inputs:

- **Foreground**: The current layer. The clip to use as foreground.

- **Background**: Defaults to None. The clip to use as background.

- **Matte**: Defaults to None. The alpha channel of this input specifies the opacities of the Foreground input. If this input is not provided, the alpha channel of the Foreground input is used instead. This input can be affected by the Invert Matte or Matte Use parameters.

- **Kernel**: Defaults to None. The filter kernel or shape for the convolution. This should normally be all black around the edges (outside the specified Kernel Crop region), with a non-black central part. A larger shape normally produces blurrier results. Only the part of the kernel within the two Kernel Crop params is considered; the part outside that boundary is ignored.


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

- **Size Front** (Default: 1, Range: 0 or greater)
  Size Front resizes the kernel larger or smaller when convolving the Front clip. 1.0 is the original size. This parameter can be adjusted using the Size Front Widget.

- **Size Back** (Default: 0, Range: 0 or greater)
  Size Back resizes the kernel larger or smaller when convolving the Back clip. This parameter can be adjusted using the Size Back Widget.

- **Size Rel X** (Default: 1, Range: 0 or greater)
  Increase to make the kernel fatter or wider without changing its height. Decrease to shrink it horizontally, making it thinner.

- **Size Rel Y** (Default: 1, Range: 0 or greater)
  Increase to make the kernel taller without changing its wieght. Decrease to shrink it vertically, making it flatter.

- **Kernel Center** (X & Y, Default: [0 0], Range: any)
  The center point of the kernel; if you think of convolution as repeated stamping of the kernel at each point of the source, the center is where the stamp aligns with the source pixels it's stamped over. If you move the center to the right in the kernel, the whole result image will move to the left, and similarly up and down. This parameter is ignored if AutoCenter is on. It may be helpful to turn on Show Kernel while adjusting this parameter. Note that if Autocenter is off, the center point is always included in the kernel no matter what this param is set to. This parameter can be adjusted using the Kernel Center Widget.

- **Autocenter** (Check-box, Default: on)
  Automatically finds the center of the kernel image. Turning this on makes the effect ignore the Kernel Center parameter.

- **Use Color Kernel** (Check-box, Default: off)
  Use each color channel of the kernel independently. Turn this on if your kernel is not just black and white and you want the colors of the kernel to be used in the convolution. Turn off for fastest rendering.

- **Show Kernel** (Check-box, Default: off)
  Show the kernel over the result, for easier adjustment of kernel parameters. Turn this off for final rendering.

- **Use Gamma** (Default: 1, Range: 0.1 or greater)
  Values above 1 cause highlights in the source clip to keep their brightness after the convolution filter is applied.

- **Matte Gamma** (Default: 1, Range: 0.1 or greater)
  The gamma value to use for the defocus of the Matte.

- **Boost Highlights** (Default: 0, Range: 0 or greater)
  The amount to increase the luma of the highlights in the source clip. Increase this parameter to blow out the highlights without affecting the darks or mid-tones.

- **Hilight Threshold** (Default: 0.9, Range: 0 or greater)
  The minimum luma value for highlights. Pixels brighter than this will be brightened according to the Boost Highlights parameter.

- **Comp Premult** (Check-box, Default: on)
  Disable this if you have provided a separate Matte input and the Foreground pixel values have not been pre-multiplied by this Matte.

- **Front Brightness** (Default: 1, Range: 0 or greater)
  Scale the brightness of the convolved Front clip.

- **Front Opacity** (Default: 1, Range: 0 to 1)
  Scale the opacity of the front clip before compositing over the back.

- **Front Threshold** (Default: 0, Range: 0 or greater)
  In the Front clip, any source value below this will be treated as black. When combining the convolved result with the original, you can increase this value to only convolve bright areas of the source. Typically when using this parameter, you will also set Combine to Screen or Add to get a glare-like effect.

- **Threshold Add Color** (Default rgb: [0 0 0])
  This can be used to raise the threshold on a specific color and thereby reduce the convolved result generated on areas of the source clip containing that color.

- **Back Brightness** (Default: 1, Range: 0 or greater)
  Scale the brightness of the convolved Back clip.

- **Combine** (Popup menu, Default: Convolve Only)
  Determines how the front, back, and convolved images are combined.
  - **Convolve Only**: Convolve the Front and Back and composite them together.
Use this option for a blur or defocus-like effect
  - **Screen**: Composite the Front over the convolved back, then screen with the convolved
front. Use this option for a glow or glare-like effect.
  - **Add**: Composite the Front over the convolved back, then add the convolved
front.
  - **Difference**: Composite the Front over the convolved back, then show the difference
with the convolved front.

- **Edge Mode** (X & Y, Popup menu, Default: [ Transparent Transparent ])
  Determines the behavior when accessing areas outside the source image.
  - **Transparent**: Areas outside the source image are treated as transparent, which can produce
transparency around the edges of the image.
Select this for fastest rendering.
  - **Repeat**: Repeats the last pixel outside the border of the image.
  - **Reflect**: Reflects the image outside the border.

- **Kernel Threshold** (Default: 0.001, Range: 0 or greater)
  Any kernel value below this will be treated as black. It's important for the edges of the kernel image to be completely black, or the result will have a grayish cast to it. If your kernel image may have a little noise in the black areas, turn up threshold a little to remove that background noise.

- **Clamp Below Thresh** (Check-box, Default: on)
  When turned on, values below the threshold are clamped to zero. This usually gives the best result. For certain special cases with partially-negative kernels, turning this off gives you additional flexibility in designing your kernel.

- **Kernel Crop1** (X & Y, Default: [-0.997 -0.747], Range: any)
  The upper left corner of the kernel area. Parts of the kernel image outside the rectangle defined by Kernel Crop1 and Kernel Crop2 are assumed to be black. Making this area smaller to avoid processing the kernel's black edges can speed up the convolution somewhat. It may be helpful to turn on Show Kernel while adjusting this parameter. Note that if Autocenter is off, the center point is always included in the kernel no matter what this param is set to.

- **Kernel Crop2** (X & Y, Default: [0.997 0.747], Range: any)
  The lower right corner of the kernel area.

- **Autoscale Mode** (Popup menu, Default: Max Channel)
  In convolution, either a larger or brighter kernel will make the result image brighter. The kernel must be auto-scaled or normalized so the result is, on average, as bright as the input. The autoscaling can be done in several ways, each of which is best in certain circumstances. With a monochrome kernel or with Color Kernel turned off, Max Channel, Luma, and Indep Channels all give the same result.
  - **Max Channel**: Autoscales the kernel by summing the
elements of each channel, and using whichever is brightest as the
overall kernel scale factor. This normalizes a dim kernel to full
brightness, and generally preserves the color of the kernel, but
allows brightness variations in the dimmer channels to show in the
result.
  - **Luma**: Autoscales the kernel by summing the
luminances of each kernel pixel. This method preserves changes in
the kernel's hue, but normalizes the luma, so a brighter or darker
kernel will have no effect. Use the Scale parameter to adjust the
result brightness.
  - **Indep Channels**: Independently normalizes each
color channel of the kernel. A colored kernel will give a
white/gray result with this method. Use this method if your kernel
channels are independent of each other (i.e. different things going
on in each of R, G, and B) but you want normalized results in each
channel.
  - **Count Nonzero**: Count how many kernel pixels are
nonzero (brighter than black), but otherwise ignore how bright they
are. This method is best if you want variations in kernel hue and
luma to show up in the result. But blurring the kernel will give a
dimmer result, since there will be more nonzero pixels.
  - **Kernel Size**: Ignore the pixel
entirely;
only use the size of the kernel rectangle to auto-scale. Use this
if you want all kernel variations to show up in the result, but
don't use it if you intend to animate Kernel Crop1 and Crop2, as
that would affect the result's brightness.

- **Matte Use** (Popup menu, Default: Alpha)
  Determines how the Matte input channels are used to make a monochrome matte.
  - **Luma**: the luminance of the RGB channels is used.
  - **Alpha**: only the Alpha channel is used.

- **Invert Matte** (Check-box, Default: off)
  If on, inverts the Matte input so the effect is applied to areas where the Matte is black instead of white. This has no effect unless the Matte input is provided.

- **Soft Borders** (Check-box, Default: off)
  If enabled, transparent borders are added to the input image before processing. This allows the result to include soft edges beyond the original image size. When off, the effect only occurs within the frame and the result will retain an edge at the borders.

- **Opacity** (Popup menu, Default: Normal)
  Determines the method used for dealing with opacity/transparency.
  - **All Opaque**: Use this option to render slightly faster when
the input image is fully opaque with no transparency (alpha=1).
  - **Normal**: Process opacity normally.
  - **As Premult**: Process as if the image is already in
premultiplied form (colors have been scaled by opacity). This option
also renders slightly faster than Normal mode, but the results will
also be in premultiplied form, which is sometimes less correct.

- **Show Size Front** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Size Front parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Size Back** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Size Back parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Kernel Center** (Check-box, Default: on)
  Turns on or off the screen user interface for adjusting the Kernel Center parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

- **Show Kernel Crop** (Check-box, Default: off)
  Turns on or off the screen user interface for adjusting the Kernel Crop1 parameter.This parameter only appears on AE and Premiere, where on-screen widgets are supported.

