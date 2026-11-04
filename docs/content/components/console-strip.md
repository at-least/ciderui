+++
title = "Console Strip"
description = "A row of 16:9 thumbnails scrolled sideways, in named groups, one current, each marked kept, waiting or to redraw."

[extra]
category = "Console"
image = "pictures/console-strip.png"
+++

A storyboard under its player: every shot as a 16:9 thumbnail in one row that scrolls sideways and comes to rest on a shot, its scenes’ names above their shots. The shot playing is ringed; a corner or edge says what is left to do with each.

```css
@import "ciderui/components/console/console-strip/console-strip.css";
```

{{ <cui_example file="console-strip/strip" title="Console strip" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-strip` | a `<nav>`: one row of groups, scrolled sideways, snapping to each shot |
| `cui-console-strip-group` | a scene: its name, then its shots 12px apart; 24px to the next |
| `cui-console-strip-label` | the scene’s name, 11px, secondary, cut to its shots’ width |
| `cui-console-strip-item` | a shot’s link, its `<img>` 160px wide at 16:9 with 12px corners; `aria-current="true"` rings it in the brand blue |
| `cui-console-strip-mark` | a mark, `role="img"` with its `aria-label`: `data-state="doubt"` a dot in the warning color, `kept` a check, `redraw` a red edge inside the picture |

Each mark is a shape of its own, not a color alone, and in a high-contrast theme it keeps its shape; its `aria-label` is read after the picture’s `alt`, as the link’s name (“Shot 2.2 To redraw”).

The Go module writes the strip with `console.ConsoleStrip`, its scenes `console.ConsoleStripGroup` and its shots `console.ConsoleStripItem`, a mark one of `console.ConsoleStripMarkDoubt`, `ConsoleStripMarkKept` and `ConsoleStripMarkRedraw` with its `MarkAriaLabel`.
