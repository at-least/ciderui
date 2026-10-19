+++
title = "Tile"
description = "apple.com’s home page tile: a headline, a subhead and buttons over a picture, the whole tile a link."

[extra]
category = "Page"
image = "pictures/tile.png"
+++

A tile of apple.com’s home page: a block of its own ground with a headline, a subhead, a callout and buttons centered at its top, its picture filling it. The whole tile is a link, its buttons above it.

```css
@import "ciderui/components/button/button.css";
@import "ciderui/components/tile/tile.css";
```

{{ <cui_example file="tile/tile" title="A promo tile" /> }}

| class | what |
| --- | --- |
| `cui-tile` | the tile, on the alternate ground (black in a dark block): 580px tall at least from 1069px, 490px from 735px, 500px below |
| `cui-tile-rounded` | 28px corners, the product page’s |
| `cui-tile-media` | its picture, filling it |
| `cui-tile-link` | the whole tile a link, named by its `aria-label` |
| `cui-tile-content` | the text and buttons |
| `cui-tile-headline` | 40px (32px below 1069px) |
| `cui-tile-subhead` | 21px (19px) |
| `cui-tile-callout` | 14px in the secondary text |
| `cui-tile-ctas` | its buttons in a row, above the link |

## Full width

`data-size="lg"` is the home page’s full-width tile: 692px tall at least, 650px from 735px, 500px below, its type a step larger (a 56px headline), its buttons a hero’s, `data-size="cta"`.

{{ <cui_example file="tile/hero" title="A full-width tile" /> }}

The crate writes it with `Tile` and `TileSize`.
