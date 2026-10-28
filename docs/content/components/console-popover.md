+++
title = "Console Popover"
description = "App Store Connect’s popover: a small panel a button opens over the page."

[extra]
category = "Console"
image = "pictures/console-popover.png"
+++

App Store Connect’s popover: a small panel a button opens over the page, holding a few related controls or a line of text. It is the button’s `popovertarget`: a click or a tap opens it, Escape or a click elsewhere closes it, with no script. It opens under its button, centered on it, and above it where there is no room below.

```css
@import "ciderui/components/console-popover/console-popover.css";
```

{{ <cui_example file="console-popover/popover" title="Console popover" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-popover` | the page’s own element with `popover` and an `id` its button names: 16px in, 13px, 240px to 350px wide, its parts 8px apart |

The hint’s popover is one too.
