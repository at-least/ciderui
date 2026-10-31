+++
title = "Console Radio"
description = "App Store Connect’s radio: a 15px circle before its label, the blue when chosen."

[extra]
category = "Console"
image = "pictures/console-radio.png"
+++

One choice of a few: a 15px circle before its label, a line under the label saying what the choice means. The input itself is drawn, so it takes the focus, the click and the arrow keys as a radio does.

```css
@import "ciderui/components/console/console-radio/console-radio.css";
```

{{ <cui_example file="console-radio/radios" title="Console radios" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-radio` | the radio and its label, a `<label>`, 13px |
| `cui-console-radio-input` | the `<input type="radio">`: 15px, a 1.5px edge; chosen, the blue with a white mark |
| `cui-console-radio-caption` | a line under the label, secondary |

The crate writes it with `ConsoleRadio`.
