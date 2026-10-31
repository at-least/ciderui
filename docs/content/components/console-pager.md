+++
title = "Console Pager"
description = "Under a long list, the buttons to the pages either side of what it shows."

[extra]
category = "Console"
image = "pictures/console-pager.png"
+++

Under a long list: the buttons to the pages either side of a line saying what it shows, as Apple Developer Forums’ pager draws them. Each button is a chevron alone, named “Previous page” or “Next page”, after a “First” button where the first page is further back; a direction with no page has no button.

```css
@import "ciderui/components/console/console-pager/console-pager.css";
```

{{ <cui_example file="console-pager/pager" title="Console pager" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-pager` | a `<nav>` named by its `aria-label`: one centered row |
| `cui-console-pager-summary` | what the list shows, 13px |
| `cui-console-btn cui-console-btn-icon` | the previous and the next page, a chevron with its `aria-label` |

## The first page

{{ <cui_example file="console-pager/first" title="The first page: no previous button" pad={true} /> }}

## A later page of a list paged by a cursor

A list paged by a cursor cannot step back one page: past its first page it offers the first page and the next.

{{ <cui_example file="console-pager/later" title="A later page: First and the next page" pad={true} /> }}

The crate writes it with `ConsolePager`.
