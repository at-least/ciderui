+++
title = "Docs Navigator"
description = "The tree of a documentation’s pages in Apple’s developer documentation: groups, pages, and the current one marked."

[extra]
category = "Documentation"
image = "pictures/docs-navigator.png"
+++

The tree of a documentation’s pages, as Apple’s developer documentation shows it beside the page: its title over its rows, a group’s title over the pages after it, each page after its glyph, the current page’s row marked. A page with pages under it opens and closes by its chevron, a checkbox and its label, with no script.

```css
@import "ciderui/components/docs/docs-navigator/docs-navigator.css";
```

{{ <cui_example file="docs-navigator/navigator" title="A documentation’s navigator" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-docs-navigator` | the navigator, a `<nav>` named Navigator |
| `cui-docs-navigator-title` | the documentation’s title, a link to its first page |
| `cui-docs-navigator-list` | a level of rows, each further level 15px further in |
| `cui-docs-navigator-group` | a group’s title |
| `cui-docs-navigator-item` | a page: its checkbox (`cui-docs-navigator-toggle`) where it has pages under it, its row, and those pages |
| `cui-docs-navigator-row` | the row: its chevron (`cui-docs-navigator-chevron`, the checkbox’s label), its glyph (`cui-docs-navigator-icon`) and its link (`cui-docs-navigator-link`), which takes the whole row |
| `aria-current="page"` | on the current page’s link: its row on the navigator’s gray |
| `cui-docs-navigator-filter` | the filter at the navigator’s foot, `hidden` until the page’s script shows it: its field (`cui-docs-navigator-field`), its glyph and its input (`cui-docs-navigator-input`) |
| `cui-docs-navigator-match` | a `<mark>` around the part of a row’s name the filter matches |
| `cui-docs-navigator-empty` | the line saying the filter matches nothing, `hidden` until then |
| `cui-docs-navigator-status` | read aloud, not shown: what the filter found |

The filter is the page’s script’s: it hides the rows whose names do not hold what is typed, keeping the pages above a match open, and marks the match; ciderui-zola’s does. Its width, its place beside the page and its scroll are the page’s. The crate writes the navigator, each page’s glyph and each chevron’s checkbox with `DocsNavigator`, its `DocsNavigatorItem`s and `DocsNavigatorPage`s.
