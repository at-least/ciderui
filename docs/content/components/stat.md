+++
title = "Stat"
description = "A product page’s number, a caption above it and one below."

[extra]
category = "Page"
image = "pictures/stat.png"
+++

A number on a product page, between two captions in the secondary text: “Up to / 25 more hours / of video playback”.

```css
@import "ciderui/components/stat/stat.css";
```

{{ <cui_example file="stat/stat" title="A stat" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-stat` | the stat, a `<p>` |
| `cui-stat-caption` | a caption, 12px on 16px at 600 |
| `cui-stat-value` | the number, 28px at 600 |

## Large

`data-size="lg"` is AirPods Pro’s: the number 80px (64px below 1069px, 48px below 735px), its captions 19px (17px, 14px).

{{ <cui_example file="stat/large" title="A large stat" pad={true} /> }}

The crate writes it with `Stat`.
