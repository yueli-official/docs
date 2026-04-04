---
title: NightSky
---

## S_NightSky

生成从主要城市或指定经纬度观看的逼真星空夜景。星星使用星表数据库生成，因此主要星座会出现在预期位置。调整星等限制可以看到更多星星。为 Minute 参数制作动画可使星星随时间逼真地移动。

在 Sapphire Render 效果子菜单中。

![NightSky](../_static/NightSky.jpg)


### Inputs:

- **Background**: 当前图层。与夜景合成的素材。如果未提供背景，源也将用作背景。

- **Matte**: 默认为无。定义应渲染星星的区域。


### Parameters:

- **Load Preset** (Push-button)
  打开预设浏览器，浏览此效果的所有可用预设。

- **Save Preset** (Push-button)
  打开预设保存对话框，保存此效果的预设。

- **Mode** (Popup menu, Default: Night Sky)
  控制如何确定位置。
  - **Night Sky**: 通过调整 Latitude、Longitude 和 GMT Offset 参数来设置位置。
  - **Night Sky Locations**: 通过从 Location 列表中选择城市来设置位置。

- **Mocha Project** (Default: 0, Range: 0 or greater)
  打开 Mocha 窗口，用于跟踪素材和生成遮罩。

- **Blur Mocha** (Default: 0, Range: 0 or greater)
  在使用前按此数值模糊 Mocha 遮罩。可用于柔化遮罩的边缘或量化伪影，并平滑时间偏移。

- **Mocha Opacity** (Default: 1, Range: 0 to 1)
  控制 Mocha 遮罩的强度。较低的值会降低效果的强度。

- **Invert Mocha** (Check-box, Default: off)
  启用后，在应用效果之前反转 Mocha 遮罩的黑白。

- **Resize Mocha** (Default: 1, Range: 0 to 2)
  缩放 Mocha 遮罩。1.0 为原始大小。

- **Resize Rel X** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对水平大小。

- **Resize Rel Y** (Default: 1, Range: 0 to 2)
  Mocha 遮罩的相对垂直大小。

- **Shift Mocha** (X & Y, Default: [0 0], Range: any)
  偏移 Mocha 遮罩的位置。

- **Dilate Mocha** (Default: 0, Range: -100 to 100)
  在使用前按此像素值膨胀或收缩 Mocha 遮罩。

- **Dilation Quality** (Popup menu, Default: Fast)
  选择 Mocha 膨胀是在默认的快速模式下快速调整，还是在高质量模式下获得更好的效果。
  - **Fast**: 在快速模式下膨胀 Mocha 遮罩，用于快速调整。
  - **High**: 在高质量模式下膨胀 Mocha 遮罩，获得更好的遮罩形状。

- **Bypass Mocha** (Check-box, Default: off)
  忽略 Mocha 遮罩，将效果应用于整个源素材。

- **Show Mocha Only** (Check-box, Default: off)
  跳过效果，仅显示 Mocha 遮罩本身。

- **Combine Masks** (Popup menu, Default: Union)
  当同时提供 Mocha 遮罩和输入遮罩时，决定如何组合它们。
  - **Union**: 使用两个遮罩共同覆盖的区域。
  - **Intersect**: 使用两个遮罩之间重叠的区域。
  - **Mocha Only**: 忽略输入遮罩，仅使用 Mocha 遮罩。

- **Latitude** (Default: 42.3, Range: -90 to 90)
  用于指定摄像机位置的纬度。

- **Longitude** (Default: -71.1, Range: -180 to 180)
  用于指定摄像机位置的经度。

- **GMT Offset** (Default: -5, Range: -12 to 12)
  指定时间与协调世界时 (UTC) 或格林尼治标准时间 (GMT) 的偏移小时数。

- **Location** (Popup menu, Default: Boston)
  指定用于确定地球表面位置的城市。可用城市分布在各大洲。
  - **Anchorage**: 安克雷奇，阿拉斯加，美国，北美洲。
  - **Astana**: 阿斯塔纳，哈萨克斯坦，亚洲。
  - **Beijing**: 北京，中国，亚洲。
  - **Boston**: 波士顿，马萨诸塞州，美国，北美洲。
  - **Cairo**: 开罗，埃及，非洲。
  - **Caracas**: 加拉加斯，委内瑞拉，南美洲。
  - **Chicago**: 芝加哥，伊利诺伊州，美国，北美洲。
  - **Hong Kong**: 香港，中国，亚洲。
  - **Istanbul**: 伊斯坦布尔，土耳其，亚洲/欧洲。
  - **Johannesburg**: 约翰内斯堡，南非，非洲。
  - **Lagos**: 拉各斯，尼日利亚，非洲。
  - **Lima**: 利马，秘鲁，南美洲。
  - **London**: 伦敦，英格兰，欧洲。
  - **Los Angeles**: 洛杉矶，加利福尼亚州，美国，北美洲。
  - **Madrid**: 马德里，西班牙，欧洲。
  - **Mexico City**: 墨西哥城，墨西哥，北美洲。
  - **Moscow**: 莫斯科，俄罗斯，欧洲。
  - **Mumbai**: 孟买，印度，亚洲。
  - **Nairobi**: 内罗毕，肯尼亚，非洲。
  - **New York City**: 纽约市，纽约州，美国，北美洲。
  - **Nuuk**: 努克，格陵兰，北美洲。
  - **Perth**: 珀斯，澳大利亚，大洋洲。
  - **Punta Arenas**: 蓬塔阿雷纳斯，智利，南美洲。
  - **Rio de Janeiro**: 里约热内卢，巴西，南美洲。
  - **Central Siberia**: 西伯利亚中部，俄罗斯，欧洲。
  - **Stockholm**: 斯德哥尔摩，瑞典，欧洲。
  - **Sydney**: 悉尼，澳大利亚，大洋洲。
  - **Tokyo**: 东京，日本，亚洲。
  - **Vancouver**: 温哥华，加拿大，北美洲。
  - **Yellowknife**: 黄刀镇，加拿大，北美洲。
  - **Warsaw**: 华沙，波兰，欧洲。

- **Star Size** (Default: 0.05, Range: 0 or greater)
  零等星的大小，单位为像素。

- **Star Brightness** (Default: 1, Range: 0 or greater)
  星星的整体亮度。

- **Altitude** (Default: 30, Range: any)
  摄像机上下旋转。高度角为 0 时朝向地平线。90 度朝正上方。180 度向后看（且上下颠倒）。

- **Azimuth** (Default: -12, Range: any)
  摄像机左右旋转。方位角为零时朝向北方，正值向东（右）旋转。

- **Field Of View** (Default: 90, Range: 0.5 to 179)
  摄像机视场角。

- **Year** (Integer, Default: 2.02e+03, Range: 1900 to 2295)
  用于查找星星位置的年份。

- **Month** (Integer, Default: 12, Range: 1 to 12)
  用于查找星星位置的月份。

- **Day** (Integer, Default: 1, Range: 1 to 31)
  用于查找星星位置的日期。

- **Hour** (Integer, Default: 16, Range: 0 to 23)
  用于查找星星位置的小时部分。应以 24 小时制格式指定。

- **Minute** (Default: 0, Range: any)
  用于查找星星位置的分钟部分。如果想让星星在一段真实世界时间内移动，请为此参数制作动画。

- **Magnitude Limit** (Default: 6.5, Range: -2 to 10)
  根据视星等控制当前可见的星星。较亮的星星有较小的星等，较暗的星星有较大的星等，这与您可能认为的相反。天空中最亮的星星星等为 0 甚至 -1。肉眼可以看到星等 5 或 6 的星星，但后院望远镜可以看到更暗的星星，星等可达 12 或更高。Magnitude Limit 越大，屏幕上可见的星星越多。增大此参数将添加比当前可见星星更暗的星星。星等接近此参数值的星星会淡入淡出，以便可以为此参数制作动画。

- **Vary Size By Mag** (Default: 0.2, Range: 0 to 1)
  使较亮的星星更大，较暗的星星更小，以便较亮的星星看起来更亮，较暗的星星看起来更暗。值为零时所有星星大小相同，值为 1 时将近似天空中自然出现的视觉大小差异。无论此参数值如何，零等星将保持相同大小。负星等的星星（非常亮的星星）会随参数值增大而变大。正星等的星星（暗星）会随参数值增大而缩小。

- **Star Saturation** (Default: 1, Range: 0 or greater)
  缩放星星的颜色饱和度。设为 0 时所有星星为白色。增大可获得更浓烈的颜色。

- **Glare** (Default: 0, Range: 0 or greater)
  要应用的眩光样式。也可以通过编辑 "s_glares.text" 文件来制作自定义眩光类型或修改现有类型。

- **Glare Brightness** (Default: 0.5, Range: 0 or greater)
  眩光的整体亮度。

- **Glare Size** (Default: 0.12, Range: 0 or greater)
  缩放眩光的大小。

- **Rel Height** (Default: 1, Range: 0 or greater)
  缩放眩光的垂直尺寸，使其变为椭圆形而非圆形。

- **Glare Color** (Default rgb: [1 1 1])
  缩放眩光的颜色。

- **Glare Rotate** (Default: 0, Range: any)
  旋转眩光的射线元素（如有），单位为度。

- **Rays Length** (Default: 1, Range: 0 or greater)
  调整射线的长度而不改变其粗细。

- **Glare Star Mag** (Default: 1.8, Range: -2 to 10)
  根据星等指定哪些星星应产生眩光。亮于此值的星星（即较低星等）将获得眩光。

- **Streak Length** (Default: 0.02, Range: 0 or greater)
  从最亮星星向外辐射的条纹或射线的长度。

- **Streak Brightness** (Default: 0.1, Range: 0 or greater)
  从最亮星星向外辐射的条纹的亮度。

- **Streak Number** (Default: 8, Range: 0 to 8)
  从最亮星星向外辐射的条纹数量。

- **Streak Rotation** (Default: 0, Range: any)
  从星星发出的条纹的旋转角度。当 Streak Symmetry 设为 1 时，值为 0 表示第一条条纹从星星垂直向上。

- **Streak Symmetry** (Default: 0.8, Range: 0 to 1)
  射线绘制的对称程度。这会同时影响间距和长度。

- **Streak Star Mag** (Default: 3, Range: -2 to 10)
  根据星等指定哪些星星应有条纹。亮于此值的星星（即较低星等）将获得条纹。

- **Twinkle Amount** (Default: 1, Range: 0 to 1)
  星星闪烁时应变暗多少。闪烁旨在模拟星星亮度不频繁但剧烈的变化（例如云朵从其前方经过）。

- **Twinkle Frequency** (Default: 5, Range: 0.01 or greater)
  星星闪烁的频率。

- **Twinkle Always** (Default: 0.3, Range: 0 to 1)
  星星平均闪烁时间的百分比。设为 1 可持续闪烁。设为较小值可使闪烁更间歇。

- **Flicker Amount** (Default: 0.3, Range: 0 to 1)
  星星闪烁时应变暗多少。闪烁旨在模拟频繁而细微的亮度变化，如大气扰动引起的变化。

- **Flicker Frequency** (Default: 30, Range: 0.01 or greater)
  星星闪烁的频率。

- **Bg Brightness** (Default: 1, Range: 0 or greater)
  缩放背景的亮度。此参数仅在提供背景输入且由于源图像部分透明或降低了 Source Opacity 参数值而可见时才有效。

- **Combine** (Popup menu, Default: Screen)
  决定星星如何与源图像合成。
  - **Stars Only**: 仅显示星星，不包含源。
  - **Mult**: 将星星与源相乘。
  - **Add**: 将星星叠加到源上。
  - **Screen**: 使用滤色操作将星星与源混合。
  - **Difference**: 结果为星星与源之间的差异。
  - **Overlay**: 使用叠加函数将星星与源合成。

- **Blur Mask** (Default: 0, Range: 0 or greater)
  在使用前按此数值模糊蒙版输入。可在蒙版区域和非蒙版区域之间提供更平滑的过渡。除非提供了蒙版输入，否则无效。

- **Invert Mask** (Check-box, Default: off)
  开启后，反转蒙版输入，使效果应用于蒙版为黑色而非白色的区域。除非提供了蒙版输入，否则无效。

- **Mask Type** (Popup menu, Default: Luma)
  除非提供了遮罩输入，否则此设置被忽略。
  - **Luma**: 使用遮罩输入的亮度来缩放灯光的亮度。
  - **Color**: 使用遮罩输入的 RGB 通道来缩放灯光的颜色。
  - **Alpha**: 使用遮罩输入的 Alpha 通道来缩放灯光的亮度。

- **Seed** (Default: 0.123, Range: 0 or greater)
  用于初始化随机数生成器。实际种子值本身并不重要，但不同的种子会产生不同的结果，相同的值应产生可重复的结果。

- **Opacity** (Popup menu, Default: Normal)
  决定处理不透明度/透明度的方法。
  - **All Opaque**: 当输入图像完全不透明且无透明度 (alpha=1) 时，使用此选项可略微加快渲染速度。
  - **Normal**: 正常处理不透明度。
  - **As Premult**: 按预乘形式处理图像（颜色已按不透明度缩放）。此选项的渲染速度也比 Normal 模式略快，但结果也将是预乘形式，有时不够精确。
