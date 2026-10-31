+++
title = "Console Link"
description = "App Store Connect’s link: blue, underlined under the pointer."

[extra]
category = "Console"
image = "pictures/console-link.png"
+++

App Store Connect’s link: the brand’s blue with no underline, underlined under the pointer. A name that is a link (an issue’s title) keeps the label’s color and turns blue.

```css
@import "ciderui/components/console/console-link/console-link.css";
```

{{ <cui_example file="console-link/link" title="Console links" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-link` | a link, on the page’s own `<a>` |
| `data-variant="label"` | in the label’s color, blue hovered |
| `aria-disabled="true"` | not available now (no `href`) |
