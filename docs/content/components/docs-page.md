+++
title = "Docs Page"
description = "A page of Apple’s developer documentation: the bar, the navigator beside the content or in a sheet, and the content’s column."

[extra]
category = "Documentation"
image = "pictures/docs-page.png"
+++

A page of Apple’s developer documentation, under the site’s global nav: the documentation’s bar across it, sticking at the top as the page scrolls; from 1024px the navigator in a column of its own beside the content, below that in a sheet the bar’s button opens, a popover with its own close button; and the content’s column, which the hero, the content and the topics each sit in.

```css
@import "ciderui/components/docs/docs-page/docs-page.css";
```

{{ <cui_example file="docs-page/page" title="A documentation page" /> }}

| class | what |
| --- | --- |
| `cui-docs-page` | the page: its bar over its columns |
| `cui-docs-page-sidebar` | the navigator’s place, a `popover` whose id the bar’s button names: a column from 1024px, a sheet over the window below |
| `cui-docs-page-close` | the sheet’s close button (`popovertargetaction="hide"`), below 1024px |
| `cui-docs-page-main` | the page’s `<main>` |
| `cui-docs-page-column` | the content’s column: 87.5% of the window to 735px, then 80px in on either side within 820px, 980px from 1251px and 1080px from 1500px |

The navigator’s column is 400px, Apple’s at a fresh load. The Go module writes the page, its bar, its navigator and its sheet’s button, the navigator’s id `docs.DocsNavigatorID`, with `docs.DocsPage`; the page writes its columns.
