+++
title = "Console Chart"
description = "App Store Connect’s analytics charts: a bar chart’s plot, scale, axis and legend, and a sparkline."

[extra]
category = "Console"
image = "pictures/console-chart.png"
+++

App Store Connect’s analytics charts. The page draws the data — its SVG, each mark’s fill, a series’ color — and these classes draw what is around it: the gridlines, the scale at the plot’s trailing end, the first and last labels under it, the legend, what the plot has none of, and a line round each mark so stacked marks stand apart. The svg’s `aria-label` says what it shows; the axis and the scale are `aria-hidden`, for sight alone.

```css
@import "ciderui/components/console-chart/console-chart.css";
```

{{ <cui_example file="console-chart/chart" title="A bar chart, two series stacked" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-chart` | the chart (a `<figure>`) |
| `cui-console-chart-plot` | the plot’s box |
| `cui-console-chart-svg` | the plot, as wide as its box; its height the svg’s |
| `cui-console-chart-grid` | a gridline, 1px |
| `cui-console-chart-bar` | a bar’s group: its `<title>`, its marks |
| `cui-console-chart-mark` | a mark, its `fill` the page’s |
| `cui-console-chart-scale` | the top gridline’s value; `data-at="half"` the middle one’s |
| `cui-console-chart-axis` | the first and last labels under the plot |
| `cui-console-chart-legend` | the series, each a `cui-console-status` |
| `cui-console-chart-empty` | what the plot has none of |

## Nothing to show

{{ <cui_example file="console-chart/empty" title="An empty chart" pad={true} /> }}

## Sparkline

{{ <cui_example file="console-chart/sparkline" title="A sparkline" pad={true} /> }}

| class or variable | what |
| --- | --- |
| `cui-console-chart-sparkline` | the svg |
| `cui-console-chart-sparkline-line` | its line, 2px, round |
| `cui-console-chart-sparkline-area` | its area |
| `--cui-console-chart-sparkline-color`, `--cui-console-chart-sparkline-fill` | the line’s and the area’s colors: the brand’s blue and its tint |
