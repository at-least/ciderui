+++
title = "Console Dialog"
description = "App Store Connect’s dialog: a white panel with 32px corners over the page."

[extra]
category = "Console"
image = "pictures/console-dialog.png"
+++

A question over the page: its title, what it asks, its buttons at its end. A `<dialog>` the page opens with `showModal()`, so the browser keeps the focus inside it and Escape closes it.

```css
@import "ciderui/components/console-dialog/console-dialog.css";
```

{{ <cui_example file="console-dialog/dialog" title="A console dialog" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-dialog` | the `<dialog>`: 540px wide, 32px corners, white, behind it the scrim |
| `cui-console-dialog-title` | its title, 18px at 600; the dialog’s `aria-labelledby` |
| `cui-console-dialog-actions` | its buttons, at its end |

In Rust, `ciderui::ConsoleDialog`.
