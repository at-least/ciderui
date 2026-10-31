+++
title = "Modal"
description = "apple.com’s store overlay on a dialog: a white panel over the dimmed page, opened and closed with no script; and its product pages’ modal, centered over the blurred page."

[extra]
category = "Store"
image = "pictures/modal.png"
+++

The store’s overlay, the one a buy page’s “Show more” opens: a layer over the window, the page dimmed and held under it, a white panel with 18px corners scrolling over it, a close button at its corner. It is HTML’s own `<dialog>`, so it opens and closes with no script: the page’s button opens it with `command="show-modal"` and `commandfor`, its close button is a form’s (`method="dialog"`), and Escape closes it too, the focus back on the button that opened it.

```css
@import "ciderui/components/site/button/button.css";
@import "ciderui/components/site/modal/modal.css";
```

{{ <cui_example file="modal/modal" title="A modal and the button that opens it" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-modal` | the `<dialog>`: the layer over the window, which scrolls the panel when it is taller; the page under it dimmed (its `::backdrop`, the scrim) and held; fading in, in 0.1s |
| `cui-modal-panel` | the panel: 816px at most, 40px from the top and as far below its end, its content 76px in; below 735px 640px at most, 20px from the top, running on to the window’s foot, 40px in |
| `data-variant="page"` | on the dialog: the product pages’ modal, its X `icon::CLOSE_SMALL`; with `cui-dark` on the dialog too, its dark one |
| `cui-modal-close` | its close button at the panel’s top right corner, a 36px circle in a 44px target, its X `icon::CLOSE`; darker hovered and pressed |

Name the dialog by its headline (`aria-labelledby`). Opened, the focus is on its close button, its first control, as a dialog’s opening gives it. A modal keeps its light look in a dark block, as apple.com’s are on its dark pages. A click on the scrim closes nothing, where apple.com’s script closes the overlay; a page that wants that adds it.

The product pages’ modal is the same dialog with `data-variant="page"`: the panel centered over the page dimmed and blurred, its corners 28px, 816px wide in a wide window, 692px from 735px and 480px at most below; its close button a dark circle with a white X, which stays at the top as a long panel scrolls. Its scrim comes in over 0.75s and its panel after it. Its dark one, as on apple.com’s dark product pages, takes `cui-dark` on the dialog: a black panel, its text light, its close button a light circle with a black X; in a dark block without it, the modal stays light.

{{ <cui_example file="modal/page" title="A product page’s modal and the button that opens it" pad={true} /> }}

A browser without the `command` attribute opens it from the page’s script, `showModal()`; Safari 26.5 and 27, Chrome, Chromium and Firefox 153 all have it. The crate writes it with `Modal` (its `page` the variant), and its opener with `ModalButton`, a button of type `button`: in a form, a submit button does nothing with its `commandfor`, opening no modal and sending no form (`reference/dialog/`).
