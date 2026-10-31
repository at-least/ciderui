+++
title = "Console Progress"
description = "App Store Connect’s progress bar: how much of a whole a value is."

[extra]
category = "Console"
image = "pictures/console-progress.png"
+++

How much of a whole a value is: a share, or a task done so far. The native `<progress>`, drawn as a thin rounded bar.

```css
@import "ciderui/components/console/console-progress/console-progress.css";
```

{{ <cui_example file="console-progress/progress" title="Console progress bars" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-progress` | the `<progress>`: 4px high, its track gray, the done part blue |
| `data-size="sm"` / `"lg"` | 3px / 5px high |
| `data-color="neutral"` | the done part black (white in dark) |

When the figure is already in text beside the bar, give the bar `aria-hidden="true"`.
