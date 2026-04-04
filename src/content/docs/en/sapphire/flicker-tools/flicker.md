---
title: Flicker
---

## S_Flicker

Scales the colors of the source clip by different
amounts over time for a flickering effect. The pattern of flickering
can be random, a periodic wave, or a combination of the two.

In the Sapphire Time effects submenu.

![Flicker](../_static/Flicker.jpg)


### Inputs:

- **Source**: The current layer. The clip to be processed.


### Parameters:

- **Load Preset** (Push-button)
  Brings up the Preset Browser to browse all available presets for this effect.

- **Save Preset** (Push-button)
  Brings up the Preset Save dialog to save a preset for this effect.

- **Amplitude** (Default: 0.2, Range: 0 or greater)
  Scales the amplitude of all flickering.

- **Rand Luma Amp** (Default: 1, Range: 0 or greater)
  The amplitude of smooth but random flickering affecting the brightness.

- **Rand Color Amp** (Default: 0, Range: 0 or greater)
  The amplitude of random flickering affecting the color channels independently.

- **Rand Freq** (Default: 30, Range: 0 or greater)
  The frequency of the random flickering. Increase for more variation between frames. Decrease for slower flickering.

- **Wave Amp** (Default: 0, Range: 0 or greater)
  The amplitude of periodic wave flickering.

- **Wave Freq** (Default: 5, Range: 0 or greater)
  The frequency of the wave flickering. Increase for faster flickering, decrease for slower. This has no effect if Wave Amp is 0.

- **Wave R Phase** (Default: 0, Range: any)
  Shifts the wave pattern in time, for the red channel.

- **Wave G Phase** (Default: 0, Range: any)
  Shifts the wave pattern in time, for the green channel.

- **Wave B Phase** (Default: 0, Range: any)
  Shifts the wave pattern in time, for the blue channel.

- **Red Amp** (Default: 1, Range: 0 or greater)
  Scales the amount of flicker applied to the red channel.

- **Green Amp** (Default: 1, Range: 0 or greater)
  Scales the amount of flicker applied to the green channel.

- **Blue Amp** (Default: 1, Range: 0 or greater)
  Scales the amount of flicker applied to the blue channel.

- **Brightness** (Default: 1, Range: 0 or greater)
  Scales the brightness of the result.

- **Seed** (Default: 0.123, Range: 0 or greater)
  Used to initialize the random number generator. The actual seed value is not significant, but different seeds give different results and the same value should give a repeatable result.

