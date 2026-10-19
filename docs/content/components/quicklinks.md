+++
title = "Quick Links"
description = "The store’s row of pill links under a title."

[extra]
category = "Store"
image = "pictures/quicklinks.png"
+++

The store’s quick links: a title over a row of pills, as far in as a shelf’s header. Each pill is a [Button](@/components/button.md), `data-variant="secondary-neutral"`, and the row wraps at every width.

```css
@import "ciderui/components/button/button.css";
@import "ciderui/components/quicklinks/quicklinks.css";
```

{{ <cui_example file="quicklinks/quicklinks" title="Quick links" /> }}

| class | what |
| --- | --- |
| `cui-quicklinks` | a `<nav>`, 980px at most, a shelf’s inset from the left (22px, to 140px from 1440px) |
| `cui-quicklinks-title` | 28px at 600 (21px below 735px), naming the nav by `aria-labelledby` |
| `cui-quicklinks-list` | the pills, 8px apart, rows 10px apart |

The crate writes it with `QuickLinks`.
