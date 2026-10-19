+++
title = "Posts"
description = "apple.com’s Newsroom archive: a month’s heading over ruled rows; smaller, the related stories under an article."

[extra]
category = "Newsroom"
image = "pictures/posts.png"
+++

The Newsroom’s archive: a heading over ruled rows, each row a link holding its picture, its category, its headline and its date. This documentation’s [news archive](@/news/archive/_index.md) is a `cui-posts` for each month.

```css
@import "ciderui/components/posts/posts.css";
```

{{ <cui_example file="posts/archive" title="A month of the archive" /> }}

| class | what |
| --- | --- |
| `cui-posts` | 980px, 692px from 735px, 87.5% below, centered |
| `cui-posts-heading` | 28px at 700 (24px from 735px, 21px below) |
| `cui-post` | a row, a link, ruled above in the divider’s color, 40px in above and below (32px, 28px); no hover |
| `cui-post-media` | its picture, 295 by 166px (264 by 148.5px from 735px, 105 by 59px below), covering |
| `cui-post-category` | 14px at 700, uppercase, the secondary text |
| `cui-post-headline` | 24px at 700 (21px, 17px) |
| `cui-post-date` | a `<time>`, 14px at 600, the secondary text |

## Related stories

`data-size="sm"` is the related stories under an article: in the article’s text column, ruled below, a 171 by 96px picture (93px square below 735px). The stories under each of this documentation’s news are one.

{{ <cui_example file="posts/related" title="Related stories" /> }}

The crate writes it with `Posts`, `Post` and `PostSize`.
