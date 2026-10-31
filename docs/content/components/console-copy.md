+++
title = "Console Copy"
description = "A value to copy in a code listing, with a copy button in its corner."

[extra]
category = "Console"
image = "pictures/console-copy.png"
+++

A value to copy — a command, a URL, a key — in a code listing, with a round copy button in its top-right corner, as Apple’s developer documentation puts one on its code listings. The button is always shown, so a touch screen can use it. In a console card, whose gray is the listing’s, the copy takes the page’s white. The page’s script copies the value: while the button has `data-state="copied"` it shows a check, and its name should read “Copied”.

```css
@import "ciderui/components/console/console-listing/console-listing.css";
@import "ciderui/components/console/console-copy/console-copy.css";
```

{{ <cui_example file="console-copy/copy" title="A command to copy, and a URL just copied" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-copy` | the listing and its button |
| `cui-console-copy-button` | the button: name it “Copy” (`aria-label`) |
| `cui-console-copy-idle` | Lucide’s `copy`, shown before a copy |
| `cui-console-copy-done` | Lucide’s `check`, shown while copied |
| `data-state="copied"` | on the button, after a copy |
| `cui-console-copy-inline` | a copy button in a line of text, after the value it copies |

## In a card

{{ <cui_example file="console-copy/card" title="In a console card, on the page’s white" pad={true} /> }}

## In a line of text

A value inside a sentence — an issue’s ID, an event’s — takes a small copy button after it instead: its icon in the secondary color, a hit area that leaves the line’s height as it was, and a 44 pt target on a touch screen.

{{ <cui_example file="console-copy/inline" title="An issue ID to copy, an event ID just copied" pad={true} /> }}

The crate writes the button’s two icons with `ConsoleCopyIcons`; the button is the page’s own, with whatever its script reads.
