+++
title = "Local Nav"
description = "A product’s bar under the global nav: its name, its pages, a button, stuck as the page scrolls."

[extra]
category = "Navigation"
image = "pictures/localnav.png"
+++

The bar most of apple.com’s product pages carry under the global nav, on a `<nav>`: the product’s name, its pages’ links with the current one marked, and a button, stuck to the top of the window as the page scrolls. Below 834px the links fold into a menu its chevron opens below the bar, with no script: a checkbox and its label, apple.com’s own way. The bar over this page, “Components”, is one with its name alone.

```css
@import "ciderui/components/localnav/localnav.css";
```

{{ <cui_example file="localnav/localnav" title="A local nav" /> }}

| class | what |
| --- | --- |
| `cui-localnav` | the bar: 52px, 48px below 834px, sticky, on its material with a rule along its foot; `cui-dark` on it, a dark bar |
| `cui-localnav-content` | 980px at most |
| `cui-localnav-title` | the product’s name, a link: 21px, 19px below 834px |
| `cui-localnav-toggle` | the menu’s checkbox, before the menu, so the next stop opened is the first link |
| `cui-localnav-menu`, `cui-localnav-list` | the links’ place: in the bar, or below it as the menu |
| `cui-localnav-link` | a link: 12px in the bar, a 34px row in the menu; `aria-current="page"` the current page, a line under it |
| `cui-localnav-actions` | the chevron (`cui-localnav-menubutton`, the checkbox’s label) and the button, a `cui-btn` with `data-size="sm"` |

The crate writes it with `LocalNav`, its chevron `icon::CHEVRON`.
