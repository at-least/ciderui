+++
title = "Console Empty State"
description = "App Store Connect’s empty state: what a list or a chart shows when it has nothing."

[extra]
category = "Console"
image = "pictures/console-empty.png"
+++

Centered where a list or a chart has nothing to show: an icon, a heading, a line saying what to do next, its action.

```css
@import "ciderui/components/console-empty/console-empty.css";
```

{{ <cui_example file="console-empty/empty" title="A console empty state" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-empty` | the empty state, centered |
| `data-size="sm"`, `"lg"` | in a card; a whole page’s |
| `data-variant="backed"` | on a faint ground of its own, 32px corners |
| `cui-console-empty-icon` | its icon, a Lucide icon |
| `cui-console-empty-heading` | its heading |
| `cui-console-empty-caption` | what to do next |
| `cui-console-empty-actions` | its buttons |

## Small, backed

{{ <cui_example file="console-empty/small" title="A small backed empty state" pad={true} /> }}

## Large

{{ <cui_example file="console-empty/large" title="A large empty state" pad={true} /> }}

The crate writes it with `ConsoleEmpty`.
