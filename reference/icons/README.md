# The icons

The numbers ciderui's icons are drawn from. apple.com draws its small icons with its own icon font (SF Pro Icons) or its own SVG; both are Apple's, so neither is copied. What is read is a glyph's box in em, the place of its ink in it and its strokes' thickness; ciderui's icon is drawn from those numbers alone, and held to sharing most of the ink with apple.com's (`rust/src/icon.rs`).

## The "more" chevron

| | em |
| --- | --- |
| the character | SF Pro Icons U+F301, in `.icon-after.more::after` |
| advance | 0.4419 |
| ink, left to right | 0.0725 to 0.3675 |
| ink, above the baseline | −0.0025 to 0.5175 |
| strokes' thickness | 0.0679 |

ciderui's `MORE`: two strokes of that thickness, round at the ends and the join, from the ink's corners (inset by half the thickness) to its middle: (0.1077, 0.4823) to (0.3323, 0.2575) to (0.1077, 0.0327), above the baseline. Drawn on the same canvas at 400px, it shares 0.9481 of its ink with apple.com's glyph (intersection over union; the probe fails below 0.9, a mark set before the run). Measured 2026-10-19 on https://www.apple.com/iphone/ in Chromium 151 on Linux (`results/more.json`).

## The global nav's icons

apple.com's nav draws its icons as inline SVG, a regular one in the bar from 834px (15 or 14 by 44) and a compact one on a phone (17 by 48). `nav.mjs` reads each off the home page, rasterizes it at 20 pixels to the unit (`svg.js`) and reads its ink: its box, and its runs along rows and columns, whose lengths are the strokes' thickness where a row crosses a stroke once (a 45° stroke's run is its thickness times √2). ciderui's are drawn from those numbers, in the same view boxes:

| icon | apple.com's ink | ciderui's |
| --- | --- | --- |
| search, 15 by 44 | 0.45 to 14.5 across, 14.45 to 28.5 down; the ring 1.08 thick, the handle 1.52 at 45° | a circle at (5.95, 19.95), radius 4.96, stroke 1.08; the handle from the ring's middle at 45° to (13.74, 27.74), stroke 1.52, round |
| bag, 14 by 44 | 0.55 to 13.45, 13.05 to 28; strokes 1.1 | a rectangle from (1.1, 16.55), 11.8 by 10.9, corners 1.5; the handle an arc of radius 2.45 on two legs from the body's top edge |
| search, 17 by 48 | 0.45 to 16.5, 15.45 to 31.5; the ring 1.2 to 1.25 thick (its middle at 1.075 and 12.8 across, 16.075 and 27.8 down), the handle's run 2.55 across, 1.8 thick | a circle at (6.94, 21.94), radius 5.86, stroke 1.22; the handle from (11.08, 26.08), the ring's middle at 45°, to (15.6, 30.6), stroke 1.8, round |
| bag, 17 by 48 | 0.95 to 16.05, 13.45 to 31.05; the body's sides 1.2 thick, its outer corner meeting the side 2.6 from it; the handle's legs 1.25 thick, their middles at 5.375 and 11.625 | a rectangle from (1.55, 17.5), 13.9 by 12.95, corners 2, stroke 1.2; the handle an arc of radius 3.125 on legs from the body's top edge |
| the menu, 18 by 18 | two lines across, 2 to 16, at 5 and 12 down, 1.2 thick, round ends | the same two lines |
| the menu, open | the same lines from (3.5, 3.5) to (15, 15) and from (3.5, 15) to (15, 3.5): the end of the animation apple.com runs on them | each closed line turned 45° about its middle, moved to (9.25, 9.25) and stretched to the X's length (1.1617 times: 16.26 against 14) |

| icon | shared ink |
| --- | --- |
| search | 0.9274 |
| bag | 0.9484 |
| search, compact | 0.9717 |
| bag, compact | 0.9807 |
| the menu | 1 |
| the menu, open | 0.9908 (its round ends, stretched with the line, reach 0.05 further) |

Intersection over union; the probe fails below 0.85, a mark set before the run. Measured 2026-10-19 on https://www.apple.com/ in Chromium 151 on Linux (`results/nav.json`). The open X is ciderui's lines as `globalnav.css` turns them, written here as SVG transforms with the same numbers; the component's test holds the rendered button to the same place (each line's box 11.5 by 11.5 from 3.5, 3.5).

## Running it

```sh
PW=…/playwright-chromium/index.mjs node run.mjs   # the "more" chevron
PW=…/playwright-chromium/index.mjs node nav.mjs   # the global nav's icons
```

`glyph.js` runs on apple.com's page, where its icon font is declared: it loads the face, draws the character on a canvas, reads the numbers off the ink, draws the chevron they describe on another, and compares the two. `nav.mjs` reads the nav's SVG off the page and compares on a blank one. The browsers are `../browsers.mjs`'s.
