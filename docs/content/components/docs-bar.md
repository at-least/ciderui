+++
title = "Docs Bar"
description = "The documentation’s bar in Apple’s developer documentation: the button that opens the navigator, then the documentation’s title."

[extra]
category = "Documentation"
image = "pictures/docs-bar.png"
+++

The bar across the top of Apple’s developer documentation, under the global nav: the button that opens the navigator, then the documentation’s title, a link home, over a hairline. From 1024px the navigator is a column of the page, and the button goes.

```css
@import "ciderui/components/docs/docs-bar/docs-bar.css";
```

{{ <cui_example file="docs-bar/bar" title="The documentation’s bar" /> }}

| class | what |
| --- | --- |
| `cui-docs-bar` | the bar: 52px, 48px below 768px, a hairline at its foot |
| `cui-docs-bar-toggle` | the button that opens the navigator, its `popovertarget` the navigator’s id; below 1024px alone |
| `cui-docs-bar-title` | the documentation’s title, a link home |

The Go module writes the bar and its button’s glyph with `docs.DocsBar`.
