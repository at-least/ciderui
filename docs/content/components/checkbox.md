+++
title = "Checkbox"
description = "apple.com’s 16px box before its label; checked, the accent with a white check."

[extra]
category = "Controls"
image = "pictures/checkbox.png"
+++

The checkbox of Apple’s sign-in: a 16px box before its label, on the field’s ground with a gray edge; checked, the accent with a white check. The input itself is kept, transparent under the box, so it takes the focus, the click and the space bar as any checkbox does; focused from the keyboard, the ring is drawn on the box.

```css
@import "ciderui/components/checkbox/checkbox.css";
```

{{ <cui_example file="checkbox/checkbox" title="A checkbox" pad={true} /> }}

{{ <cui_example file="checkbox/checked" title="A checked checkbox" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-checkbox` | the checkbox and its label |
| `cui-checkbox-input` | the `<input type="checkbox">`, transparent, the box’s size and place |
| `cui-checkbox-label` | its `<label>`, 17px, 26px in |
| `cui-checkbox-indicator` | the box inside the label, holding the check (`icon::CHECK`) |

The crate writes it with `Checkbox`.
