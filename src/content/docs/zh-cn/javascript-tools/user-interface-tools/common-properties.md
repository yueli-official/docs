---
title: 通用属性
---
# 通用属性

所有类型的用户界面元素，包括窗口、容器和控件，共享许多相同的属性，尽管某些属性在不同类型的对象中含义略有不同。下表总结了哪些属性用于哪些对象类型。

| 属性 | [窗口](.././window-object) | [面板](../control-objects#panel) | [选项卡面板](../control-objects#tabbedpanel) | [选项卡](../control-objects#tab) | [组](../control-objects#group) | [按钮](../control-objects#button) | [复选框](../control-objects#checkbox) | [下拉列表](../control-objects#dropdownlist) | [编辑文本](../control-objects#edittext) | [Flash播放器](../control-objects#flashplayer) | [图标按钮](../control-objects#iconbutton) | [图像](../control-objects#image) | [列表框](../control-objects#listbox) | [列表项](../types-of-controls#listitem) | [进度条](../control-objects#progressbar) | [单选按钮](../control-objects#radiobutton) | [滚动条](../control-objects#scrollbar) | [滑块](../control-objects#slider) | [静态文本](../control-objects#statictext) | [树视图](../control-objects#treeview) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [`active`](../control-objects#active) | √ | | | | | √ | √ | √ | √ | √ | √ | √ | √ | | | √ | √ | √ | √ | √ |
| [`alignChildren`](../window-object#alignchildren) | √ | √ | √ | √ | √ | | | | | | | | | | | | | | | |
| [`alignment`](../control-objects#alignment) | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | | √ | √ | √ | √ | √ | √ |
| [`bounds`](../control-objects#bounds) | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | | √ | √ | √ | √ | √ | √ |
| [`cancelElement`](../window-object#cancelelement) | √ | | | | | | | | | | | | | | | | | | | |
| [`characters`](../control-objects#characters) | | √ | √ | √ | √ | | | | | | | | | | | | | | | |
| [`checked`](../control-objects#checked) | | | | | | | | | | | | | | √ | | | | | | |
| [`children`](../window-object#children) | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | | √ | √ | √ | √ | √ | √ |
| [`columns`](../control-objects#columns) | | | | | | | | | | | | | √ | | | | | | | |
| [`defaultElement`](../window-object#defaultelement) | √ | | | | | | | | | | | | | | | | | | | |
| [`enabled`](../control-objects#enabled) | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ |
| [`expanded`](../control-objects#expanded) | | | | | | | | | | | | | | √ | | | | | | |
| [`frameBounds`](../window-object#framebounds) | √ | | | | | | | | | | | | | | | | | | | |
| [`frameLocation`](../window-object#framelocation) | √ | | | | | | | | | | | | | | | | | | | |
| [`frameSize`](../window-object#framesize) | √ | | | | | | | | | | | | | | | | | | | |
| [`graphics`](../control-objects#graphics) | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | | √ | √ | √ | √ | √ | √ |
| [`helpTip`](../control-objects#helptip) | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | | √ | √ | √ | √ | √ | √ |
| [`icon`](../control-objects#icon) | | | | | | | | | | | √ | √ | | √ | | | | | | |
| [`image`](../control-objects#image) | | | | | | | | | | | √ | √ | | √ | | | | | | |
| [`index`](../control-objects#index) | | | | | | | | | | | | | | √ | | | | | | |
| [`items`](../control-objects#items) | | | | | | | | √ | | | | | √ | | | | | | | √ |
| [`itemSize`](../control-objects#itemsize) | | | | | | | | √ | | | | | √ | | | | | | | √ |
| [`jumpdelta`](../control-objects#jumpdelta) | | | | | | | | | | | | | | | | | √ | | | |
| [`justify`](../control-objects#justify) | | | | | | √ | √ | | √ | | | | | | | √ | | | √ | |
| [`layout`](../window-object#layout) | √ | √ | √ | √ | √ | | | | | | | | | | | | | | | |
| [`location`](../control-objects#location) | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | | √ | √ | √ | √ | √ | √ |
| [`margins`](../window-object#margins) | √ | √ | √ | √ | √ | | | | | | | | | | | | | | | |
| [`maximumSize`](../control-objects#maximumsize) | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | | √ | √ | √ | √ | √ | √ |
| [`maxvalue`](../control-objects#maxvalue) | | | | | | | | | | | | | | | √ | | √ | √ | | |
| [`minimumSize`](../control-objects#minimumsize) | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | | √ | √ | √ | √ | √ | √ |
| [`minvalue`](../control-objects#minvalue) | | | | | | | | | | | | | | | √ | | √ | √ | | |
| [`orientation`](../window-object#orientation) | √ | √ | √ | √ | √ | | | | | | | | | | | | | | | |
| [`parent`](../control-objects#parent) | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ |
| [`preferredSize`](../control-objects#preferredsize) | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | | √ | √ | √ | √ | √ | √ |
| [`properties`](../control-objects#properties) | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ |
| [`resizeable`](../window-object#creation-properties) | √ | | | | | | | | | | | | | | | | | | | |
| [`selected`](../control-objects#selected) | | | | | | | | | | | | | | √ | | | | | | |
| [`selection`](../control-objects#selection) | | | √ | | | | | √ | | | | | √ | | | | | | | √ |
| [`shortcutKey`](../control-objects#shortcutkey) | √ | | | | | √ | √ | √ | √ | √ | √ | √ | √ | | | √ | √ | √ | √ | √ |
| [`size`](../control-objects#size) | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | √ | | √ | √ | √ | √ | √ | √ |
| [`spacing`](../window-object#spacing) | √ | √ | √ | √ | √ | | | | | | | | | | | | | | | |
| [`stepdelta`](../control-objects#stepdelta) | | | | | | | | | | | | | | | | | √ | | | |
| [`subitems`](../control-objects#subitems) | | | | |
