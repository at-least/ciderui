+++
title = "Console Toast"
description = "A word that the page did what was asked, floating for a moment on the popover’s ground."

[extra]
category = "Console"
image = "pictures/console-toast.png"
+++

A short word that the page did what was asked (“Saved”), at the window’s bottom right for a moment. The page sets `data-shown` to show it and removes it to fade it out; a `role="status"` region says it to a screen reader.

```css
@import "ciderui/components/console/console-toast/console-toast.css";
```

{{ <cui_example file="console-toast/toast" title="Console toast" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-toast` | fixed at the bottom right, the popover’s ground; shown with `data-shown` |
