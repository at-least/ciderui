+++
title = "Console Table"
description = "App Store Connect’s table: rows between hairlines, a row that is a link."

[extra]
category = "Console"
image = "pictures/console-table.png"
+++

App Store Connect’s table, as its analytics pages draw one: rows between hairlines, its header’s labels small and secondary, a figure at its end, a row that is a link filling faintly under the pointer. It is the page’s own `<table>` with the class.

```css
@import "ciderui/components/console/console-table/console-table.css";
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
| a checked `cui-console-checkbox` in a row | the row selected, as `aria-selected` |
| `data-sticky` | its header sticks to the top of its scroll container, at `--cui-console-table-sticky-top` |
| `data-stack` | below 640px each row a block of lines, a cell’s `data-label` before its value |
| `cui-console-table-scroll` | a box that scrolls a wide table sideways |
| `cui-console-table-sort` | a header’s label that sorts, a link; its `th` carries `aria-sort`, the sorted one’s label darker with its arrow (`cui-console-table-sort-icon`) |

## Large

`data-size="lg"`, the analytics overview’s plain table.

{{ <cui_example file="console-table/large" title="A large console table" pad={true} /> }}

## Sticky, sortable, selectable

A long list’s table: its header sticks under the page’s bar (`data-sticky`), its labels sort the list, a row’s checkbox selects it.

{{ <cui_example file="console-table/sticky" title="A sticky console table" pad={true} /> }}

## Stacked on a phone

`data-stack`: below 640px a row is a block, each cell a line of its `data-label` and its value.

{{ <cui_example file="console-table/stack" title="A stacking console table" pad={true} /> }}
