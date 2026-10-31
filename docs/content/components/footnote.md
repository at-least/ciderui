+++
title = "Footnote"
description = "A footnote’s mark: a link at the text’s size, its digits the font’s numerators."

[extra]
category = "Page"
image = "pictures/footnote.png"
+++

A product page’s footnote mark: a `<sup>` holding a link to its note, at the text’s size and on its baseline, underlined in the text’s color, its digits drawn small and raised by the font itself (`font-feature-settings: "numr"`, which SF and Inter both have).

```css
@import "ciderui/components/site/footnote/footnote.css";
```

{{ <cui_example file="footnote/footnote" title="A footnote's mark" pad={true} /> }}

The notes themselves are the footer’s, `cui-footer-notes` on the [Footer](@/components/footer.md). Give the link an `aria-label` (“Footnote 1”), since a lone digit says little. The crate writes it with `Footnote`.
