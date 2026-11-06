# The icons

The numbers ciderui's icons are drawn from. apple.com draws its small icons with its own icon font (SF Pro Icons) or its own SVG; both are Apple's, so neither is copied. What is read is a glyph's box in em, the place of its ink in it and its strokes' thickness; ciderui's icon is drawn from those numbers alone, and held to sharing most of the ink with apple.com's (`go/internal/glyph/glyph.go`).

## The global nav's icons

apple.com's nav draws its icons as inline SVG, a regular one in the bar from 834px (15 or 14 by 44) and a compact one on a phone (17 by 48). `nav.mjs` reads each off the home page, rasterizes it at 20 pixels to the unit (`svg.js`) and reads its ink: its box, and its runs along rows and columns, whose lengths are the strokes' thickness where a row crosses a stroke once (a 45° stroke's run is its thickness times √2). ciderui's are drawn from those numbers, in the same view boxes:

| icon | apple.com's ink | ciderui's |
| --- | --- | --- |
| the menu, 18 by 18 | two lines across, 2 to 16, at 5 and 12 down, 1.2 thick, round ends | the same two lines |
| the menu, open | the same lines from (3.5, 3.5) to (15, 15) and from (3.5, 15) to (15, 3.5): the end of the animation apple.com runs on them, in 0.24s | each closed line turned 45° about its middle, moved to (9.25, 9.25) and stretched to the X's length (1.1617 times: 16.26 against 14) |

| icon | shared ink |
| --- | --- |
| the menu | 1 |
| the menu, open | 0.9908 (its round ends, stretched with the line, reach 0.05 further) |

Intersection over union; the probe fails below 0.85, a mark set before the run. Measured 2026-10-19 on https://www.apple.com/ in Chromium 151 on Linux (`results/nav.json`). The open X is ciderui's lines as `globalnav.css` turns them, written here as SVG transforms with the same numbers; the component's test holds the rendered button to the same place (each line's box 11.5 by 11.5 from 3.5, 3.5).

## The field's icons

`form.mjs` reads the store's two form icons. The error's mark before a field's message is an SF Pro Icons glyph, U+F312, read as a glyph of the icon font is (drawn at 400px on the gift card's page, where the font is declared): its advance 1.1709em, its ink 0.0875 to 1.0825em across and 0.85 to −0.145em about the baseline; a ring, its rows' runs 0.0825em thick; in it a stem 0.265 to 0.605em up, 0.04 to 0.085em wide (tapering), and a dot 0.1075 to 0.2075em up. In the page's message row (12px, at four device pixels to the CSS pixel) its ink is 1.75 to 13.73px across and 2.75 to 14.73px down the 16px row (read again 2026-10-22, the same; a row with no red ink left of the text now fails the run, where its scan gave a box inside out, its left past its right, which the check took as a place: tried on pictures drawn for it, with a mark and without). The dropdown's chevron is an SVG its background draws, 21 by 21 at 25px: a line 1.45 thick (its rows' runs 2 to 2.1 across at 45°), its ink 5.3 to 15.7 across and 8.55 to 14.45 down, in `#6e6e73`.

| icon | ciderui's | shared ink |
| --- | --- | --- |
| `ALERT`, 1.171em by 0.995em (from the ink's top) | a ring at (585, 497.5) thousandths, radius 456.25, 82.5 thick; a stem from (584, 280) to (584, 550), 70 thick, round; a dot at (585, 692.5), radius 50 | 0.9695 |
| `DROPDOWN`, 21 by 21 drawn at 25 | a line through (6.075, 9.275), (10.5, 13.7) and (14.925, 9.275), 1.45 thick, round | 0.9689 |

Intersection over union; `form.mjs` fails below 0.85, a mark set before the run, or if the field's error example (`error.html`) draws others. Measured 2026-10-20 in Chromium 151 on Linux (`results/form.json`).

## Running it

```sh
PW=…/playwright-chromium/index.mjs node nav.mjs   # the global nav's icons
PW=…/playwright-chromium/index.mjs node form.mjs   # the field's error mark
```


## The accordion's and the footer's chevrons

Both are apple.com's lines, drawn as its pages draw them: the bag's accordion's a polyline through (15, 1.13), (8.5, 7.72) and (2, 1.13) in 17 by 8.85, 2.25 thick, round (`DISCLOSURE`); a phone's footer section's through (10.075, 0.675), (5.5, 5.323) and (0.925, 0.675) in 11 by 6, 1 thick, round (`FOOTER_CHEVRON`). Turned up, each is the line its animation ends on (the same points mirrored). `lines.mjs` reads both live and holds ciderui's to them: 1 and 1 of the ink shared (fails below 0.85), each in its component's example. Measured 2026-10-19 in Chromium 151 on Linux (`results/lines.json`).


Apple's developer documentation's icons (docc.mjs and its results) were measured here for the docs set, removed 2026-11-04 with it (the owner: 「ciderui也不再支援docs」); git history has them.

The more link's chevron and the probes of the global nav's search, flyouts and bag were measured here for parts removed by DESIGN.md decision 36, and went with them; git holds them.
