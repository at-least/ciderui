+++
title = "Card"
description = "The store’s card: a picture, a headline and a price, the whole card a link, lifted as the pointer comes over it."

[extra]
category = "Store"
image = "pictures/card.png"
+++

The store’s card: 400 by 500px (309 by 450px below 735px) on the page’s ground, 18px corners and a soft shadow, its picture filling it and its text at its top. The whole card is its headline’s link, with no script: the link’s `::after` lies over the card. Hovered, or with the focus in it, its shadow deepens and it grows 1%, in 0.3s.

```css
@import "ciderui/components/site/card/card.css";
```

{{ <cui_example file="card/card" title="A card with a tag" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-card` | the card; `cui-dark` on it, a black card with white text |
| `cui-card-media` | its picture, filling it, under its text |
| `cui-card-content` | its text, 28px in, 340px wide at most |
| `cui-card-eyebrow` | 12px at 600 over the headline; `data-variant="tag"` a tag in its place, in the tag’s orange |
| `cui-card-headline` | 28px at 600 (21px below 735px) |
| `cui-card-link` | the headline’s link, the whole card’s |
| `cui-card-description` | 14px at 600 |
| `cui-card-price` | 14px at 400 |

## An eyebrow

A plain eyebrow is the dark gray of apple.com’s, the headline as far down with no eyebrow above it as with one.

{{ <cui_example file="card/eyebrow" title="A card with an eyebrow" pad={true} /> }}

Its size is the store’s most common card; a site gives it another width with a utility (`w-full`), its picture cropped to it. A row of cards is a [Shelf](@/components/shelf.md). The crate writes it with `Card`.
