+++
title = "Console Chip"
description = "App Store Connect’s chip: a key and its value on the gray fill."

[extra]
category = "Console"
image = "pictures/console-chip.png"
+++

A key and its value on the gray fill, with 5px corners: a filter in force, with a control that removes it, or a tag that is a link to the list it filters.

```css
@import "ciderui/components/console-chip/console-chip.css";
```

{{ <cui_example file="console-chip/chips" title="Console chips" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-chip` | the chip, a `<span>` or a link: 13px, 2px by 4px in, on the gray fill; a link filled darker hovered |
| `cui-console-chip-key` | its key, before its value, secondary |
| `cui-console-chip-remove` | the control that removes a filter, its × 12px, its label its `aria-label` |

In Rust, `ciderui::ConsoleChip` (`href` for a link, `remove_href` and `remove_label` for a filter).
