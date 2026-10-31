+++
title = "Section"
description = "A band of the page: its ground, its rhythm, its content’s width and its header."

[extra]
category = "Page"
image = "pictures/section.png"
+++

A band of an apple.com page: its ground, the room above and below it, and its content, 87.5% of the window, at most 420px below 735px and 1260px from it, centered.

```css
@import "ciderui/components/site/link/link.css";
@import "ciderui/components/site/section/section.css";
```

{{ <cui_example file="section/section" title="A section" /> }}

| attribute | what |
| --- | --- |
| `data-variant="alt"` | the alternate ground (`#f5f5f7`, `#1d1d1f` in a dark block); without it, the page’s |
| `data-size="lg"` | iPhone 18 Pro’s rhythm, 112px, 120px from 735px, 144px from 1069px; without it, the iPhone page’s, 56px, 96px, 112px |
| `cui-dark` | a dark band |

The room is the section’s own: a page joining two sections takes one away with a utility (`pt-0`).

## Its header

`cui-section-header` holds a headline (`cui-section-headline`, 32px, 40px from 735px, 48px from 1069px) and the links beside it (`cui-section-links`, more links in a wrapping row), side by side from 735px, their last lines on one baseline, and one under the other below.

{{ <cui_example file="section/header" title="A section's header" pad={true} /> }}

The crate writes them with `Section` (its `Ground` and `Rhythm`) and `SectionHeader`.
