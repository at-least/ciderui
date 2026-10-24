+++
title = "Console Thumbnail"
description = "App Store Connect’s screenshot thumbnail: 240px tall, 12px corners, a link to itself."

[extra]
category = "Console"
image = "pictures/console-thumb.png"
+++

A screenshot shown small, a link to itself whole, its name under it.

```css
@import "ciderui/components/console-thumb/console-thumb.css";
```

{{ <cui_example file="console-thumb/thumbs" title="Console thumbnails" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-thumbs` | a wrapping row, 12px between |
| `cui-console-thumb` | a `<figure>`: its link and image 240px tall, 12px corners |
| `cui-console-thumb-caption` | its name, 11px, secondary |
