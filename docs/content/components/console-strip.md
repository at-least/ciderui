+++
title = "Console Strip"
description = "A row of 16:9 thumbnails scrolled sideways, in named groups, one current, each captioned and marked kept or to redraw, a doubt beside either."

[extra]
category = "Console"
image = "pictures/console-strip.png"
+++

A storyboard under its player: every shot as a 16:9 thumbnail in one row that scrolls sideways and comes to rest on a shot, its code under it, its scenes’ names above their shots. A scene’s name stays at the row’s start while its scene scrolls past. The shot playing is ringed; a corner or edge says what is left to do with each.

```css
@import "ciderui/components/console/console-strip/console-strip.css";
```

{{ <cui_example file="console-strip/strip" title="Console strip" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-strip` | a `<nav>`: one row of groups, scrolled sideways, snapping to each shot |
| `cui-console-strip-group` | a scene: its name, then its shots 12px apart; 24px to the next |
| `cui-console-strip-head` | a `<p>`: the track of the scene’s name, its shots’ width |
| `cui-console-strip-label` | a `<span>` in the head: the scene’s name, 11px, secondary, cut to its shots’ width; it sticks at the row’s start while its scene is in view |
| `cui-console-strip-item` | a shot’s link, its `<img>` 160px wide at 16:9 with 12px corners; `aria-current="true"` rings it in the brand blue; `data-variant="sheet"` crops its picture to the left edge, a character sheet’s face; `data-state="unused"` dims its picture, brighter under the pointer, its lines kept legible |
| `cui-console-strip-caption` | a `<span>` after the `<img>`: the shot’s code, 11px, secondary, 5px under the picture and cut to its width |
| `cui-console-strip-detail` | a `<span>` after the caption: a second line under it, in its type |
| `cui-console-strip-mark` | a mark, `role="img"` with its `aria-label`: `data-state="doubt"` a dot in the warning color, `kept` a check, `redraw` a red edge inside the picture; a doubt comes before the other mark, and a check after a dot stands beside it |

Each mark is a shape of its own, not a color alone, and in a high-contrast theme it keeps its shape. A shot can be current, waiting and to redraw at once (shot 2.2). The caption and each mark’s `aria-label` are read after the picture’s `alt`, as the link’s name (“Shot 2.2 S02-02 Waiting for you To redraw”).

The Go module writes the strip with `console.ConsoleStrip`, its scenes `console.ConsoleStripGroup` and its shots `console.ConsoleStripItem`: its `Caption` and `Detail`, `Variant` (`console.ConsoleStripItemVariantFrame` or `ConsoleStripItemVariantSheet`), `Unused`, a doubt by its `DoubtAriaLabel`, and a mark, `console.ConsoleStripMarkKept` or `ConsoleStripMarkRedraw`, with its `MarkAriaLabel`.

## Assets

The same strip lists a story’s assets by kind: each with its name and how many shots use it, a character sheet cropped to its face, one nothing uses dimmed, its picture alone.

{{ <cui_example file="console-strip/assets" title="Console strip of assets" pad={true} /> }}
