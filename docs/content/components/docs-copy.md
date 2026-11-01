+++
title = "Docs Copy"
description = "A code listing with Apple’s developer documentation’s copy button, shown as the pointer comes onto it."

[extra]
category = "Documentation"
image = "pictures/docs-copy.png"
+++

A code listing with Apple’s developer documentation’s copy button at its corner. It shows as the pointer comes onto the listing, with the keyboard’s focus, and always on a touch screen. The page’s script copies the listing’s code and marks the button copied for a second; with no script the button stays hidden.

```css
@import "ciderui/components/docs/docs-copy/docs-copy.css";
```

{{ <cui_example file="docs-copy/copy" title="A listing and its copy button" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-docs-copy` | the listing’s box: a `<pre>` the docs content draws, then its button |
| `cui-docs-copy-button` | the button, `hidden` until the page’s script shows it |
| `data-state="copied"` | on the button, set by the page’s script once it has copied: the check in the link’s blue |

ciderui-zola’s docs layout gives every code block its button, and its script copies. The Go module writes the box, the button and its two glyphs with `docs.DocsCopy`.
