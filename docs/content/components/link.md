+++
title = "Link"
description = "apple.com’s link in running text, and its more link with a chevron."

[extra]
category = "Controls"
image = "pictures/link.png"
+++

A link in running text is apple.com’s own `a`: the link’s color, underlined at rest and the same hovered. Its size is the text’s around it.

```css
@import "ciderui/components/link/link.css";
```

{{ <cui_example file="link/inline" title="A link in a sentence" pad={true} /> }}

## The more link

`data-variant="more"` is the standalone link with a chevron after it, “Learn more ›”: not underlined until hovered, then its text alone. The text is `cui-link-text`, the chevron `cui-link-icon`, 0.24em after it, drawn from the numbers of apple.com’s glyph.

{{ <cui_example file="link/more" title="A more link" pad={true} /> }}

Write it on one line: a space between the text and the chevron would widen the gap. The crate writes it with `MoreLink`, its chevron `icon::MORE`.
