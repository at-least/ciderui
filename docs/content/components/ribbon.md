+++
title = "Ribbon"
description = "apple.com’s message above a page’s content, a more link after it."

[extra]
category = "Navigation"
image = "pictures/ribbon.png"
+++

apple.com’s message above a page’s content: a band on the alternate ground, its text centered within 83.33% of the window, a [more link](@/components/link.md) after it.

```css
@import "ciderui/components/link/link.css";
@import "ciderui/components/ribbon/ribbon.css";
```

{{ <cui_example file="ribbon/ribbon" title="A ribbon" /> }}

## Announced

`data-variant="announce"` is the iPhone page’s: as the page shows it, it drops in on the accent, then turns to its ground, in apple.com’s timing (a 0.8s drop, the colors 1.6s on). Under reduced motion it is still.

{{ <cui_example file="ribbon/announce" title="An announcing ribbon" /> }}

`cui-ribbon` is the band, `cui-ribbon-content` its ground and its 16px above and below. The crate writes it with `Ribbon`.
