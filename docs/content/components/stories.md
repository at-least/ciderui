+++
title = "Stories"
description = "apple.com’s Newsroom index: a grid of stories, one, two or three across."

[extra]
category = "Newsroom"
image = "pictures/stories.png"
+++

The Newsroom’s index: a band on the alternate ground, its heading over a grid of stories, each a white tile rounded 32px, the whole of it its link: its picture, its category, its headline and its date. This documentation’s [News](@/news/_index.md) is one.

```css
@import "ciderui/components/site/stories/stories.css";
```

{{ <cui_example file="stories/stories" title="The latest stories" /> }}

| class | what |
| --- | --- |
| `cui-stories` | the band |
| `cui-stories-content` | 980px, 692px from 735px, 87.5% below |
| `cui-stories-heading` | 32px at 700 (28px from 735px, 24px below) |
| `cui-stories-list` | twelve columns 36px apart (24px below 1069px); below 735px one story under another |
| `cui-story` | a story, three across; no hover, no shadow |
| `cui-story-media` | its picture, 16:9, covering |
| `cui-story-content` | its text: its head (`cui-story-head`: `cui-story-category`, `cui-story-headline`) at its top, its date (`cui-story-date`) at its foot |

## Sizes

| `data-size` | what |
| --- | --- |
| (none) | three across |
| `md` | two across, 472px tall at least |
| `lg` | one across, from 1069px its picture beside its text |

From 735px to 1068px every story is two across, whatever its size, as apple.com’s script lays them out there. The Newsroom’s order is one large story, two medium, then the rest. The crate writes it with `Stories`, `Story` and `StorySize`.
