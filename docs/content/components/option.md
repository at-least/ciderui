+++
title = "Option"
description = "apple.com’s buy page choice: a tile on a radio, its title and its price; checked, the accent’s edge."

[extra]
category = "Store"
image = "pictures/option.png"
+++

A choice on apple.com’s buy page: a tile 83px tall at least, 12px corners and a gray edge, its title (and a detail under it) at one side and a price at the other. It is a radio: the input lies over the whole tile, transparent, and takes the click, the focus and the arrow keys; checked, the tile’s edge is 2px of the accent; focused from the keyboard, the ring.

```css
@import "ciderui/components/option/option.css";
```

{{ <cui_example file="option/option" title="An option with a price" pad={true} /> }}

{{ <cui_example file="option/checked" title="A checked option with a detail" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-option` | the option |
| `cui-option-input` | the `<input type="radio">`, over the tile |
| `cui-option-label` | its `<label>`, the tile, on the field’s ground |
| `cui-option-title` | 17px at 600; a detail inside it, `cui-option-detail`, 12px under it |
| `cui-option-price` | 12px at its end; without one, the title takes the row |

Options of one choice share a `name`. In a dark block an option keeps its light look, as a [Field](@/components/field.md) does. The crate writes it with `OptionTile`.
