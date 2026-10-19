+++
title = "Modal"
description = "apple.com’s store overlay on a dialog: a white panel over the dimmed page, opened and closed with no script."

[extra]
category = "Store"
image = "pictures/modal.png"
+++

The store’s overlay, the one a buy page’s “Show more” opens: a layer over the window, the page dimmed and held under it, a white panel with 18px corners scrolling over it, a close button at its corner. It is HTML’s own `<dialog>`, so it opens and closes with no script: the page’s button opens it with `command="show-modal"` and `commandfor`, its close button is a form’s (`method="dialog"`), and Escape closes it too, the focus back on the button that opened it.

```css
@import "ciderui/components/modal/modal.css";
```

{{ <cui_example file="modal/modal" title="A modal and the button that opens it" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-modal` | the `<dialog>`: the layer over the window, which scrolls the panel when it is taller; the page under it dimmed (its `::backdrop`, the scrim) and held; fading in, in 0.1s |
| `cui-modal-panel` | the panel: 816px at most, 40px from the top and as far below its end, its content 76px in; below 735px 640px at most, 20px from the top, running on to the window’s foot, 40px in |
| `cui-modal-close` | its close button at the panel’s top right corner, a 36px circle in a 44px target, its X `icon::CLOSE`; darker hovered and pressed |

Name the dialog by its headline (`aria-labelledby`). Opened, the focus is on its close button, its first control, as a dialog’s opening gives it. A modal keeps its light look in a dark block, as apple.com’s are on its dark pages. A click on the scrim closes nothing, where apple.com’s script closes the overlay; a page that wants that adds it.

A browser without the `command` attribute opens it from the page’s script, `showModal()`; Safari 26.5 and 27, Chrome and Chromium all have it. The crate writes it with `Modal`, and its opener with `ModalButton`.
