+++
title = "Swatch"
description = "apple.com’s color choice: a 28px circle of the color in a 36px ring, on a radio."

[extra]
category = "Store"
image = "pictures/swatch.png"
+++

A finish on apple.com’s buy page: a 28px circle of the color in a 36px ring, on a radio; checked, the ring is the accent. The color’s name is hidden but read.

```css
@import "ciderui/components/site/swatch/swatch.css";
```

{{ <cui_example file="swatch/swatches" title="Three finishes, the first chosen" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-swatches` | the row, which wraps: 5px apart, 8px between rows |
| `cui-swatch` | one |
| `cui-swatch-input` | the `<input type="radio">`, over the ring, transparent |
| `cui-swatch-ring` | its `<label>`, the ring: clear, the accent checked, the focus ring from the keyboard |
| `cui-swatch-color` | the color: the page’s image, or an element its `style` colors |
| `cui-swatch-name` | the color’s name, for assistive technology |

The crate writes it with `Swatch`.
