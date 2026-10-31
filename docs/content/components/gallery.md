+++
title = "Gallery"
description = "A product page’s row of items across the window, its paddles and its dots."

[extra]
category = "Page"
image = "pictures/gallery.png"
+++

A product page’s gallery: a row of items across the whole window, each snapping into place where the content starts, 20px apart; below it its paddles at the content’s right edge, and its dots. The row scrolls with no script, by touch, trackpad and keys, and each dot is a link to its item; the paddles need the page’s script, which scrolls the row, disables the paddle at an end and marks the current dot, as it would open a dialog. The theme this documentation is built with ships that script, and it runs in the example below.

```css
@import "ciderui/components/site/card/card.css";
@import "ciderui/components/site/gallery/gallery.css";
```

{{ <cui_example file="gallery/gallery" title="A gallery of cards" /> }}

| class | what |
| --- | --- |
| `cui-gallery` | the gallery, dark in a dark block |
| `cui-gallery-scroller` | its items, each snapping at its start where the content does (87.5% of the window, 1260px at most); its scrollbar hidden |
| `cui-gallery-nav` | the paddles’ row |
| `cui-gallery-paddle` | a 36px paddle, translucent over the page, its chevron; `disabled` at an end; `data-size="lg"` 56px, the store’s |
| `cui-gallery-dots` | the dots’ row, centered |
| `cui-gallery-dot` | an item’s dot, a link to it: 8px; the current one, `aria-current="true"`, a 48px pill (32px below 735px) |

A paddle names the row it scrolls with `aria-controls`, the previous first. The crate writes it with `Gallery`, its chevrons `icon::PADDLE_PREVIOUS` and `icon::PADDLE_NEXT`.
