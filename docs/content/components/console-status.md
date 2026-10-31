+++
title = "Console Status"
description = "App Store Connect’s status: a mark in its color before its word."

[extra]
category = "Console"
image = "pictures/console-status.png"
+++

A mark in the status’s color before its word, and a line under it that says more. The mark is never the only carrier: the word says the same.

```css
@import "ciderui/components/console/console-status/console-status.css";
```

{{ <cui_example file="console-status/status" title="Console statuses" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-status` | the mark and the word, 13px; `data-color` the mark’s color |
| `cui-console-status-mark` | the 12px circle, 6px before the word |
| `cui-console-status-detail` | the line under it, 11px, secondary |

{{ <cui_example file="console-status/detail" title="A console status with its detail" pad={true} /> }}

A chart’s key sets the mark to its series’ color with `--cui-console-status-color`.

{{ <cui_example file="console-status/key" title="A chart’s key" pad={true} /> }}

The crate writes it with `ConsoleStatus`.
