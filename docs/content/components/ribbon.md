+++
title = "Ribbon"
description = "apple.com’s message above a page’s content, a more link after it."

[extra]
category = "Navigation"
image = "pictures/ribbon.png"
+++

apple.com’s message above a page’s content: a band on the alternate ground, its text centered within 83.33% of the window, a [more link](@/components/link.md) after it.

```css
@import "ciderui/components/site/link/link.css";
@import "ciderui/components/site/ribbon/ribbon.css";
```

{{ <cui_example file="ribbon/ribbon" title="A ribbon" /> }}

`cui-ribbon` is the band, `cui-ribbon-content` its ground and its 16px above and below. It is still as the page shows it: apple.com’s iPhone page drops its ribbon in on the accent, an effect ciderui leaves out. The Go module writes it with `site.Ribbon`.
