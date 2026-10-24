+++
title = "Console Card"
description = "App Store Connect’s analytics card: a metric, or a section’s list or chart."

[extra]
category = "Console"
image = "pictures/console-card.png"
+++

The analytics card: a gray rounded panel on the page holding a metric (its title, its value, a change up or down, a caption), or a section’s list or chart under its header.

```css
@import "ciderui/components/console-card/console-card.css";
```

{{ <cui_example file="console-card/metrics" title="Metric cards" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-card` | the card |
| `cui-console-card-header` | a row: its title, and a link at its end |
| `cui-console-card-title` | its title, 13px at 600 |
| `cui-console-card-figure` | a value and its change on one baseline |
| `cui-console-card-value` | the value, 41px at 700 |
| `cui-console-card-change` | the change; `data-color="success"` up, `"error"` down |
| `cui-console-card-caption` | a line under it |

## A list in a card

{{ <cui_example file="console-card/list" title="A card holding a table" pad={true} /> }}

The crate writes a metric’s card with `ConsoleMetric`.
