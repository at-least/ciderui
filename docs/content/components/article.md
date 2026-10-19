+++
title = "Article"
description = "apple.com’s Newsroom story: its header over its copy, markdown’s elements styled as they come."

[extra]
category = "Newsroom"
image = "pictures/article.png"
+++

The Newsroom’s story: a header (an eyebrow of its category over its date, the headline, a subhead) over its copy, the text in a column narrower than its figures, its footnotes at its end. The copy’s plain elements are styled as they come, so the HTML markdown writes reads as the Newsroom’s with no class on it: every page of this documentation is one.

```css
@import "ciderui/components/article/article.css";
```

{{ <cui_example file="article/article" title="An article" /> }}

## Its columns

The text is 653px wide from 1069px, 576px from 735px, and below that 87.5% of the window up to 414px; its figures 980px, 692px and 414px at most, centered on the same line. The header is 52px down (40px from 735px, 32px below), the copy 44px under it (36px, 32px).

## Its header

| class | what |
| --- | --- |
| `cui-article-header` | the header, in the text’s column |
| `cui-article-eyebrow` | its category (`cui-article-category`, 12px at 700, uppercase) over its date (`cui-article-date`, a `<time>`, 14px at 600), the secondary text |
| `cui-article-headline` | 48px at 700 (40px from 735px, 32px below) |
| `cui-article-subhead` | 24px at 500 (21px below 1069px) |

## Its copy

`cui-article-body` styles its children as markdown writes them:

| element | how |
| --- | --- |
| `p` | 19px on 27px (17px on 25px below 1069px), a line apart |
| `h2`, `h3` | 24px and 19px at 700, room above them |
| `h4` to `h6` | the copy’s size at 700 |
| `ul`, `ol` | the copy, its marks outside |
| `a`, `strong` | the link’s color, underlined; 700 |
| `sup` | a footnote’s mark, 0.6em, raised |
| `code`, `pre` | the mono font; a block of it on the alternate ground |
| `blockquote` | indented, in the secondary text |
| `table`, `hr` | rows and a rule in the divider’s color |
| `figure`, `figcaption` | the figures’ column; the caption in the text’s, 12px at 600 |

A word too long for its line, an import in code, a bare URL, a hash in a table’s cell, breaks where it must, so it stays in the column; apple.com’s copy, prose with no such word, breaks none.

`cui-article-footnotes` is the notes’ section at the end, 12px in the secondary text, numbered. The elements apple.com’s copy has none of (h3 to h6, code, a quotation, a table, a rule) are ciderui’s. A component inside the copy, any element with a `cui-` class, keeps its own look: the copy’s rules match plain elements alone.

The crate writes it with `Article`, its copy the page’s; ciderui-zola writes a page’s markdown into it.
