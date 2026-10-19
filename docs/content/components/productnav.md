+++
title = "Product Nav"
description = "The store’s shelf of small cards, each a product’s picture over its name."

[extra]
category = "Store"
image = "pictures/productnav.png"
+++

The row at the top of apple.com’s store: small cards, each a product’s picture above its name, the whole card a link to its products. It is a [Shelf](@/components/shelf.md) with no header, so it scrolls and snaps as one, its paddles the shelf’s.

```css
@import "ciderui/components/shelf/shelf.css";
@import "ciderui/components/productnav/productnav.css";
```

{{ <cui_example file="productnav/productnav" title="A product nav" /> }}

| class | what |
| --- | --- |
| `cui-productnav` | on a `cui-shelf`: 62px below it (30px below 735px), its cards 10px apart |
| `cui-productnav-item` | a card, a link |
| `cui-productnav-image` | its picture, 78px tall (60px below 735px), as wide as its shape makes it |
| `cui-productnav-name` | 14px on 20px at 600, centered |

The crate writes it with `ProductNav` and `ProductNavItem`.
