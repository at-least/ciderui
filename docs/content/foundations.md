+++
title = "Foundations"
description = "The colors, the type, the breakpoints and the rest every component is made of."
+++

Every component reads the same small contract, `tokens.css`: no component writes a color of its own, so a part placed in a dark block, or in another part, keeps to the page’s palette. The values are apple.com’s, read off its live pages; where one fell short of its contrast, it is fitted, and said so.

## The colors

Each color is one variable, a pair of a light and a dark value (`light-dark()`), light first:

| variable | light | dark | its job |
| --- | --- | --- | --- |
| `--cui-c-bg` | `#ffffff` | `#000000` | the page |
| `--cui-c-bg-alt` | `#f5f5f7` | `#1d1d1f` | the alternate ground: a band, the footer, a ribbon |
| `--cui-c-text` | `#1d1d1f` | `#f5f5f7` | body text |
| `--cui-c-text-2` | `#6e6e73` | `#86868b` | secondary text (dark: fitted, apple.com’s `#6e6e73` is 4.1:1 on black) |
| `--cui-c-text-3` | `#86868b` | `#6e6e73` | muted text, at 3:1, never body |
| `--cui-c-border` | `#86868b` | `#6e6e73` | a control’s edge: a field, an option, a checkbox |
| `--cui-c-divider` | `#d2d2d7` | `#424245` | a rule between things |
| `--cui-c-accent` | `#0071e3` | the same | the primary fill, the focus ring, a checked edge |
| `--cui-c-link` | `#0066cc` | `#2997ff` | a link; a secondary button’s text and edge |
| `--cui-c-neutral` | `#1d1d1f` | `#f5f5f7` | the neutral fill |

The rest are a part’s own: the global nav’s bar and its links, the local nav’s, a card’s text, a gallery’s paddles and dots, the footer’s. The palette is fixed: a site does not redefine a color, and there are no themes, as apple.com has one palette.

The semantic utilities read the same variables, so a page’s own markup follows a block’s scheme too: `text-body`, `text-secondary`, `text-muted`, `bg-page`, `bg-alt`, `border-border`, `text-accent`, `bg-accent`, `text-link`.

## Dark blocks

A block is dark when the page says so: `cui-dark` on any element sets its color scheme, and every color inside it takes its dark value; `cui-light` inside it is light again. On `<html>` it darkens the whole page. There is no automatic dark mode, since apple.com has none, and `dark:` is a variant for a dark block, not for the reader’s setting.

{{ <cui_example file="pages/dark" title="A dark block" /> }}

## The type

The type is apple.com’s store’s ladder, each class a Tailwind utility that steps its size by itself at 735px, so a variant composes with it (`cui-headline lg:cui-headline-super`) and a plain utility on the same element wins (`cui-headline text-[20px]`):

{{ <cui_example file="pages/type" title="The type's classes" pad={true} /> }}

| class | from 735px | below | weight |
| --- | --- | --- | --- |
| `cui-headline-standalone` | 96/100 | 48/52 | 600 |
| `cui-headline-super` | 80/84 | 48/52 | 600 |
| `cui-headline-elevated` | 64/68 | 40/44 | 600 |
| `cui-headline` | 48/52 | 32/36 | 600 |
| `cui-headline-reduced` | 40/44 | 28/32 | 600 |
| `cui-eyebrow-super` | 32/36 | 24/28 | 600 |
| `cui-eyebrow-elevated` | 28/32 | 21/25 | 600 |
| `cui-eyebrow` | 24/28 | 21/25 | 600 |
| `cui-eyebrow-reduced` | 21/25 | 19/23 | 600 |
| `cui-callout` | 32/36 | 24/28 | 600 |
| `cui-label` | 24/28 | 21/25 | 600 |
| `cui-quote` | 40/48 | 28/36 | 600 |
| `cui-quote-reduced` | 32/40 | 24/32 | 400 |
| `cui-manifesto` | 32/40 | 24/32 | 600 |
| `cui-intro-elevated` | 24/32 | 21/29 | 400 |
| `cui-intro` | 21/29 | 19/27 | 400 |
| `cui-tout` | 19/23 | 19/23 | 600 |
| `cui-body` | 17/25 | 17/25 | 400 |
| `cui-body-tight` | 17/21 | 17/21 | 400 |
| `cui-body-reduced` | 14/20 | 14/20 | 400 |
| `cui-body-reduced-tight` | 14/18 | 14/18 | 400 |
| `cui-caption` | 12/16 | 12/16 | 400 |

The font is the system’s on Apple’s platforms, SF, named by its two Apple-only names (`-apple-system`, `BlinkMacSystemFont`), then Inter. Each size has its own letter-spacing, apple.com’s less the tracking SF already carries, so the system font lands within 0.0006em a character of apple.com’s text; Inter takes the same values and runs a little wide of it.

## Breakpoints

apple.com’s, in pixels, in place of Tailwind’s: `md` 735px and `lg` 1069px, where the content steps, and `xl` 1441px. The global nav and the local nav fold their links into a menu below 834px, apple.com’s nav’s own step. A component’s step and a page’s `md:` switch at the same width.

## The rest

- **Radii**: the pill, `980px`; a field, an option and a form button, `12px`; a block button, `8px`; a card, `18px`; a rounded tile, `28px`; a checkbox, `3px`.
- **Control sizes**: apple.com’s pill in four heights, 24, 36, 44 and 56px; a field, 56px.
- **The focus ring**: 2px of the accent, 3px out, on every element (`base.css`); a text field shows its 2px accent edge instead.
- **Motion**: apple.com’s durations and its curve, from a paddle’s 0.1s to the menu’s 0.45s; every transition and animation stops under `prefers-reduced-motion`.
- **Widths**: the content’s 87.5% of the window, up to 1260px (420px at most below 735px); the older 980px, the footer’s; the global nav’s 1024px.
