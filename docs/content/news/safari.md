+++
title = "Safari on the Mac, by its keyboard"
description = "The global nav’s menu and the accordion, measured in Safari 27 on macOS, as its keyboard users use them."
date = 2026-10-19T18:10:25+08:00

[extra]
category = "Update"
image = "pictures/accordion.png"
+++

What only a browser can answer, ciderui measures in that browser. The global nav’s menu, a `popover`, and the accordion, a `<details>`, now have Safari 27.0.1 on macOS among their measurements, beside Chrome and Chromium and Safari on the iOS simulator.

## Driving Safari

Safari takes an automation session only in a Safari that its driver launched itself, so a run now quits the Safari it launched when it ends. And Safari on the Mac behaves as its keyboard users know it: a click focuses no button, and Tab skips links, which Option-Tab reaches. So the menu’s checks drive it that way, a button pressed with Return and the links reached with Option-Tab, and all eleven of them hold: the menu opens as a sheet, the focus moves to its close button, Tab goes on to its links, and Escape closes it with the focus back on its button.

## The accordion

Safari opens an answer at once: it has no `interpolate-size`, by which Chrome moves the answer’s height to `auto`. Nothing waits on the motion, so the accordion reads the same in both.
