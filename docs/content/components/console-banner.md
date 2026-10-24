+++
title = "Console Banner"
description = "App Store Connect’s banner: a message on a rounded panel tinted in its color."

[extra]
category = "Console"
image = "pictures/console-banner.png"
+++

A message on a rounded panel tinted in its color: an icon at its start, a title, a line or two under it, its actions, a button that closes it (which the page’s script wires).

```css
@import "ciderui/components/console-banner/console-banner.css";
```

{{ <cui_example file="console-banner/banner" title="A console banner" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-banner` | the panel, `role="status"` (an error one `role="alert"`); the gray fill |
| `data-color` | `info`, `success`, `warning`, `error`: tinted in its color |
| `cui-console-banner-icon` | its icon, a Lucide icon |
| `cui-console-banner-title` | its title, 13px at 500 |
| `cui-console-banner-body` | a line under it |
| `cui-console-banner-caption` | a small line, secondary |
| `cui-console-banner-actions` | its buttons |
| `cui-console-banner-dismiss` | its close button, on a plain small icon button |

## Colors

{{ <cui_example file="console-banner/colors" title="Console banners in each color" pad={true} /> }}

The crate writes it with `ConsoleBanner`.
