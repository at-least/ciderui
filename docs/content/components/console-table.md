+++
title = "Console Table"
description = "App Store Connect’s table: rows between hairlines, a row that is a link."

[extra]
category = "Console"
image = "pictures/console-table.png"
+++

App Store Connect’s table, as its analytics pages draw one: rows between hairlines, its header’s labels small and secondary, a figure at its end, a row that is a link filling faintly under the pointer. It is the page’s own `<table>` with the class.

```css
@import "ciderui/components/console-table/console-table.css";
```

{{ <cui_example file="console-table/table" title="A console table" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-table` | the table, 13px on 18px; rows 48px |
| `data-size="sm"`, `"lg"` | rows 32px or 56px |
| `cui-console-table-link` | a row’s link: the whole row, filled at 0.03 hovered |
| `aria-selected="true"` on a `<tr>` | a selected row, on the brand’s fill |
| `cui-console-table-figure` | a number’s cell, at its end at 500 |
| `cui-console-table-subtitle` | a line under a cell’s text, 12px, secondary |

## Large

`data-size="lg"`, the analytics overview’s plain table.

{{ <cui_example file="console-table/large" title="A large console table" pad={true} /> }}
