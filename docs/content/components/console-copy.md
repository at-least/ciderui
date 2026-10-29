+++
title = "Console Copy"
description = "A value to copy in a code listing, with a copy button in its corner."

[extra]
category = "Console"
image = "pictures/console-copy.png"
+++

A value to copy — a command, a URL, a key — in a code listing, with a round copy button in its top-right corner, as Apple’s developer documentation puts one on its code listings. The button is always shown, so a touch screen can use it. In a console card, whose gray is the listing’s, the copy takes the page’s white. The page’s script copies the value: while the button has `data-state="copied"` it shows a check, and its name should read “Copied”.

```css
@import "ciderui/components/console-listing/console-listing.css";
@import "ciderui/components/console-copy/console-copy.css";
```

{{ <cui_example file="console-copy/copy" title="A command to copy, and a URL just copied" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-copy` | the listing and its button |
| `cui-console-copy-button` | the button: name it “Copy” (`aria-label`) |
| `cui-console-copy-idle` | Lucide’s `copy`, shown before a copy |
| `cui-console-copy-done` | Lucide’s `check`, shown while copied |
| `data-state="copied"` | on the button, after a copy |

## In a card

{{ <cui_example file="console-copy/card" title="In a console card, on the page’s white" pad={true} /> }}
