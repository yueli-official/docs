---
title: Retime
---

## S_Retime

Retimes a clip using optical flow based motion estimation and frame interpolation.

In the Sapphire Time effects submenu.

![Retime](../_static/Retime.jpg)


### Inputs:

- **Source**: The current layer. The clip to be retimed.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Mode** (Popup menu, Default: Retime Speed)
  How the clip is to be retimed.
  - **Retime Speed**: The output clip is to run at a specified fraction of the speed of the input clip.
  - **Retime Longer**: The output clip is to be a specified percentage longer than the input clip.
  - **Retime Scale**: The output clip is to be a specified percentage times the length of the input clip.
  - **Retime FPS**: The output clip speed is to be suitable for running at a specified frames-per-second
given the input clip speed specified in frames-per-second.
  - **Retime Length**: Retimes a clip using optical flow motion estimation so that the output clip is a specified
length (in frames).
  - **Retime Variable Speed**: Retimes a clip using optical flow motion estimation so that the input frame for every output
frame may be controlled with a key-framed parameter. This allows vari-speed effects, backwards motion and
so forth.

- **Speed** (Default: 0.5, Range: 0.01 to 10)
  Used in Speed mode. The output clip runs at a speed which is this times the speed of the input clip. For example, 0.2 would make the output 5 times slower than the input.

- **Percent Longer** (Default: 100, Range: 0 to 1e+04)
  Used in Longer mode. How many percent longer the output is to be relative to the input. 0 percent makes the output the same length as the input. 100 percent would make the output twice as long as the input.

- **Percent** (Default: 200, Range: 1 to 1e+04)
  Used in Scale mode. How long the output clip is to be as a percentage of the input clip length. 100 percent makes the output the same length as the input. 200 percent would make the output twice as long as the input. 50 percent would make the output half as long as the input.

- **Input Fps** (Default: 30, Range: 1 to 1000)
  Used in FPS mode. The frame rate of the input clip.

- **Output Fps** (Default: 60, Range: 1 to 1000)
  Used in FPS mode. The frame rate of the output clip. If this is greater than the Input Fps parameter, the output clip will be longer than the input clip. If this is less than the Input Fps parameter, the output clip will be shorter than the input clip. If the output clip is played at the specified output frames-per-second, it will last as long (in seconds) as the input clip.

- **Output Length** (Integer, Default: 30, Range: 2 to 1e+04)
  Used in Length mode. The desired length of the output clip in frames.

- **Input Frame** (Default: 1, Range: 0 to 1e+04)
  Used in Curve mode. The input frame to output. This parameter should normally be key-framed, with the frame for which a key-frame is set designating an output frame, and the value set at that key-frame denoting the desired input frame at that output frame.

- **Result** (Popup menu, Default: Result)
  Output the retimed clip or a representation of the motion vectors.
  - **Result**: Output the retimed clip.
  - **Flow Vectors**: Output a representation of the motion vectors. The direction of the vector at a pixel
is shown as a color, and the length of the vector as the saturation of that color.

- **Start Frame** (Integer, Default: 0, Range: 0 or greater)
  Frame number of the first frame in the clip.

- **Motion Blur** (Default: 0, Range: 0 to 2)
  If greater than 0.0, blur along this fraction of the length of the motion vectors.

- **Field Dominance** (Popup menu, Default: Normal)
  For interlaced input clips, specifies the field order.
  - **Normal**: Use normal field order.
  - **Reverse**: Reverse the field order.

- **Flow Field Smoothness** (Default: 0.8, Range: 0 or greater)
  Optical flow parameter: The relative importance of adhering to the image data relative to the importance of keeping a smooth flow field. The higher this value, the more relative importance is assigned to the flow field smoothness.

- **Emphasize Edges** (Default: 0, Range: 0 to 1)
  Optical flow parameter: Specify the degree of structure/texture decomposition to apply to the input clip prior to solving for the motion vector field. This can be useful to reduce problems due to lighting changes and may give better localised flows and other improvements. It is not recommended when there is fast motion in the input clip. However, in some cases, it can be very helpful in getting acceptable results. Suggested usage is to first try a value of 0.0 (which turns off structure/texture decomposition). If the result is unsatisfactory, try 0.4. Finally, try 1.0.

- **Flow View Scale** (Default: 0.9, Range: 0 or greater)
  When outputting a representation of the motion vectors, this increases the saturation of the color corresponding to a unit length motion. This makes the representation more sensitive to show shorter vectors more clearly.

- **Crop Input Parameters** (Default: 0, Range: 0 or greater)
  These 4 parameters, Crop Top , Crop Bottom , Crop Left, and Crop Right , allow selecting a rectangular subsection of the input image to be processed. This can avoid artifacts due to pulling in black at the edges. This can happen when a lower resolution element is retimed in a higher resolution project.

