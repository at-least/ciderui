+++
title = "Global Nav"
description = "apple.com’s bar across the top: a logo, its links, a search and a bag, a menu below 834px."

[extra]
category = "Navigation"
image = "pictures/globalnav.png"
+++

apple.com’s bar across the top of every page, on a `<nav>`: the site’s logo, its links and its icon links (a search, a bag) spread across 1024px at most, on a translucent material under a blur. Below 834px the links fold into a menu, which the bar’s button opens over the window as a sheet, with no script: the list is a `popover`. This page’s own bar is one.

```css
@import "ciderui/components/globalnav/globalnav.css";
```

{{ <cui_example file="globalnav/globalnav" title="A global nav" /> }}

The list is the menu. Closed, the page shows it as the bar’s row; opened, it is a sheet over the window with a close button of its own where the menu’s button was, its two lines turning into an X, and the focus moves to it (`autofocus`); Tab goes on to the links, and the X or Escape closes it, the focus back on the menu’s button. In a narrow window the example’s menu opens over its frame.

| class | what |
| --- | --- |
| `cui-globalnav` | the bar: 44px, 48px below 834px; `cui-dark` on it, a dark bar; `data-variant="clear"`, no ground and no blur, over a hero |
| `cui-globalnav-content` | 1024px at most, 22px in from 834px |
| `cui-globalnav-list` | the row closed, the menu opened: `id` and `popover` |
| `cui-globalnav-logo` | the site’s logo, a link home, named by its `aria-label` |
| `cui-globalnav-link` | a link: 12px in the bar, a 28px row in the menu |
| `cui-globalnav-icon` | an icon link, its image `cui-globalnav-image` from 834px and `cui-globalnav-image-compact` below |
| `cui-globalnav-menubutton` | the menu’s button, below 834px, beside the list: `popovertarget` |
| `cui-globalnav-closebutton` | the opened menu’s X, second in the list so the next stop is the first link |

The nav has no place on the page of its own: a site that keeps it at the top adds `sticky top-0`. The crate writes it with `GlobalNav`, its icons `icon::SEARCH`, `icon::BAG` and `icon::MENU` and their compact forms.
