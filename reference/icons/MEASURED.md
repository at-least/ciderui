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

`form.mjs` reads the store's two form icons. The error's mark before a field's message is an SF Pro Icons glyph, U+F312, read as a glyph of the icon font is (drawn at 400px on the gift card's page, where the font is declared): its advance 1.1709em, its ink 0.0875 to 1.0825em across and 0.85 to −0.145em about the baseline; a ring, its rows' runs 0.0825em thick; in it a stem 0.265 to 0.605em up, 0.04 to 0.085em wide (tapering), and a dot 0.1075 to 0.2075em up. In the page's message row (12px, at four device pixels to the CSS pixel) its ink is 1.75 to 13.73px across and 2.75 to 14.73px down the 16px row (read again 2026-10-22, the same). The dropdown's chevron is an SVG its background draws, 21 by 21 at 25px: a line 1.45 thick (its rows' runs 2 to 2.1 across at 45°), its ink 5.3 to 15.7 across and 8.55 to 14.45 down, in `#6e6e73`.

| icon | ciderui's | shared ink |
| --- | --- | --- |
| `ALERT`, 1.171em by 0.995em (from the ink's top) | a ring at (585, 497.5) thousandths, radius 456.25, 82.5 thick; a stem from (584, 280) to (584, 550), 70 thick, round; a dot at (585, 692.5), radius 50 | 0.9695 |
| `DROPDOWN`, 21 by 21 drawn at 25 | a line through (6.075, 9.275), (10.5, 13.7) and (14.925, 9.275), 1.45 thick, round | 0.9689 |

Intersection over union; `form.mjs` fails if `ALERT` shares below 0.85, a mark set before the run, or if the field's error example (`error.html`) draws another; `DROPDOWN` it records alone, the dropdown having gone by decision 36. Measured 2026-10-20 in Chromium 151 on Linux (`results/form.json`).

## The accordion's chevron

apple.com's line, drawn as the bag's accordion draws it: a polyline through (15, 1.13), (8.5, 7.72) and (2, 1.13) in 17 by 8.85, 2.25 thick, round (`glyph.Disclosure`). Turned up, it is the line its animation ends on (the same points mirrored). `lines.mjs` reads it live and holds ciderui's to it: 1 of the ink shared (fails below 0.85), in the accordion's example. Measured 2026-10-19 in Chromium 151 on Linux (`results/lines.json`; it holds the footer's chevron too, a phone's footer section's, through (10.075, 0.675), (5.5, 5.323) and (0.925, 0.675) in 11 by 6, 1 thick, round, 1 of the ink shared, which went with the footer's directory).

## Running it

```sh
PW=…/playwright-chromium/index.mjs node nav.mjs     # the global nav's icons
PW=…/playwright-chromium/index.mjs node form.mjs    # the field's error mark
PW=…/playwright-chromium/index.mjs node lines.mjs   # the accordion's chevron
```

## The search field's icons

`search.mjs` reads the two icons of apple.com's search page's field (`../apple.com/`'s `search-serp`), each an SVG its button's background draws in `rgb(134, 134, 139)`, 25 by 25: the magnifier before the text and the reset after it. It rasterizes each at 20 pixels to the unit (`svg.js`) beside ciderui's own, drawn from their numbers, read along the ring's middle row (10.87):

| icon | apple.com's ink | ciderui's | shared ink |
| --- | --- | --- | --- |
| `MAGNIFIER` | a ring 3.85 to 17.5 across, 1.1 thick at its middle row; a handle to 21 by 21, round | a ring about (10.69, 10.87), radius 6.28, 1.1 thick; a line from (15.25, 15.35) to (20.15, 20.05), 1.7 thick, round | 0.9783 |
| `RESET` | a disc 3 to 22 each way; an X cut from it | a disc of radius 9.5 about the middle; an X cut to 3.5 each way of it, 1.1 thick, round | 0.9995 |

Intersection over union; the probe fails below 0.85, a mark set before the run, and failed at 0.1843 with the magnifier's pair given the reset as ciderui's. Measured 2026-11-12 in Chromium 151 on macOS (`results/search.json`); `go/internal/glyph`'s test holds `Magnifier` and `Reset` to the shapes recorded there.

## The pager's chevrons

`pager.mjs` reads Apple Developer Forums' pager's two chevrons (`../apple.com/`'s `forums-pager`), SF Symbols' chevron.left and chevron.right at medium weight, each a filled path in a 74.561 by 70.459 view box drawn past it (overflow visible), 0.74561em wide at the button's size. It rasterizes each over the ink's whole box (0 −10 74.561 90) at 4 pixels to the unit (`svg.js`) beside ciderui's own, drawn from their numbers:

| icon | Apple's ink | ciderui's | shared ink |
| --- | --- | --- | --- |
| `LEFT` | 6 to 55.5 across, −8 to 78.5 down | two strokes 11.5 thick, round, from (49.9, −2.3) to (11.78, 35.2) to (49.9, 72.7) | 0.9852 |
| `RIGHT` | 19 to 68.5 across, the same down | the left mirrored in the 74.561 advance | 0.9839 |

Intersection over union; the probe fails below 0.85, a mark set before the run, and failed at 0.1097 with the left's pair given the right as ciderui's. The forums' Content-Security-Policy refuses the `blob:` images `svg.js` draws, so the rasterizing runs on `about:blank`. Measured 2026-11-12 in Chromium 151 on macOS (`results/pager.json`); `go/internal/glyph`'s test holds `ChevronLeft` and `ChevronRight` to the shapes recorded there.

