+++
title = "Console Section Head"
description = "A title with its controls at the row’s end."

[extra]
category = "Console"
image = "pictures/console-section-head.png"
+++

A page’s or a section’s title with what acts on it at the end of the row — a time window, a main button — as App Store Connect puts the date control beside a section’s title. Where the row has no room, the controls move under the title.

```css
@import "ciderui/components/console/console-section-head/console-section-head.css";
```

{{ <cui_example file="console-section-head/section-head" title="A title, its line, the time window" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-section-head` | the row: its first child the title (and a line under it), the rest its controls |
