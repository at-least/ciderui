+++
title = "Shelf"
description = "The store’s row of cards across the window, its header, its paddles."

[extra]
category = "Store"
image = "pictures/shelf.png"
+++

The store’s row of [cards](@/components/card.md): a header (a headline, its secondary text after it) over a scroller across the whole window, its cards snapping into place 20px apart, the first and the last as far in as the header: 22px in a window to 1024px, 140px from 1440px, between the two with the window. Its scrollbar is hidden.

```css
@import "ciderui/components/card/card.css";
@import "ciderui/components/shelf/shelf.css";
```

{{ <cui_example file="shelf/shelf" title="A shelf of cards" /> }}

| class | what |
| --- | --- |
| `cui-shelf` | the shelf |
| `cui-shelf-header` | its header, as far in as the first card |
| `cui-shelf-headline` | 28px at 600 (24px below 735px) |
| `cui-shelf-subhead` | the same in the secondary text, after a space |
| `cui-shelf-row` | the scroller and its paddles |
| `cui-shelf-scroller` | its items, each snapping at its start |

## Paddles

A shelf’s paddles are a gallery’s large ones, `cui-gallery-paddle cui-shelf-paddle` with `data-size="lg"`, at its sides over the cards; the one at an end, disabled, is hidden, and a screen with no hover shows none. As a [Gallery](@/components/gallery.md)’s, they need the page’s script, which runs here.

```css
@import "ciderui/components/gallery/gallery.css";
```

{{ <cui_example file="shelf/paddles" title="A shelf with paddles" /> }}

The crate writes it with `Shelf`.
