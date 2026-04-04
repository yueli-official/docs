---
title: RepairFrames
---

## S_RepairFrames

通过用周围帧的时间扭曲版本替换来修复素材中的一个或多个帧。

在 Sapphire Time effects 子菜单中。

![RepairFrames](../_static/RepairFrames.jpg)


### Inputs:

- **Source**: 当前图层。要处理的素材。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **First Bad Frame** (Integer, Default: 5, Range: 2 or greater)
  要替换的第一个坏帧。设置 Show:Bad Frame 并调整此参数以帮助找到坏帧。

- **Bad Frame Count** (Integer, Default: 1, Range: 1 or greater)
  从 First Bad Frame 开始要修复的帧数。

- **Show** (Popup menu, Default: Result)
  设为 Result 表示正常操作，显示修复帧后的素材。设为 Bad Frame 仅帮助找到坏帧；找到感兴趣的帧后，返回 Result。
  - **Result**: 显示结果素材，即修复了坏帧的源素材。
  - **Bad Frame**: 显示第一个坏帧，无论播放头当前在哪个位置。

