+++
title = "Console Skip Link"
description = "The keyboard’s first stop, hidden until focused, then a filled button at the top left."

[extra]
category = "Console"
image = "pictures/console-skip.png"
+++

The first link on the page, for the keyboard: hidden until it is focused, then the filled button at the window’s top left, past the sidebar to the page.

```css
@import "ciderui/components/console-skip/console-skip.css";
```

{{ <cui_example file="console-skip/skip" title="Console skip link" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-skip` | on a filled `cui-console-btn`: unseen until focused, then fixed 16px from the top and left |
