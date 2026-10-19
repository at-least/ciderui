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
| the menu, open | the same lines from (3.5, 3.5) to (15, 15) and from (3.5, 15) to (15, 3.5): the end of the animation apple.com runs on them, in 0.24s | each closed line turned 45° about its middle, moved to (9.25, 9.25) and stretched to the X's length (1.1617 times: 16.26 against 14) |

| icon | shared ink |
| --- | --- |
| search | 0.9274 |
| bag | 0.9484 |
| search, compact | 0.9717 |
| bag, compact | 0.9807 |
| the menu | 1 |
| the menu, open | 0.9908 (its round ends, stretched with the line, reach 0.05 further) |

Intersection over union; the probe fails below 0.85, a mark set before the run. Measured 2026-10-19 on https://www.apple.com/ in Chromium 151 on Linux (`results/nav.json`). The open X is ciderui's lines as `globalnav.css` turns them, written here as SVG transforms with the same numbers; the component's test holds the rendered button to the same place (each line's box 11.5 by 11.5 from 3.5, 3.5).

## The search's icons

The global nav's search, opened, draws two more SVGs, each a regular one from 834px and a compact one below: the field's magnifier and the arrow before each quick link. `searchfield.mjs` reads them off the home page (the flyout's markup is in the page closed) and compares them as `nav.mjs` does:

| icon | apple.com's ink | ciderui's |
| --- | --- | --- |
| the field's magnifier, 30 by 32 | 6.6 to 23.65 across, 7.9 to 25.2 down; the ring's outer edge 6.6 to 20.35 across and 7.9 to 21.65 down, 1.25 to 1.3 thick; the handle's run 3.1 across at 45° (2.2 thick), its middle on x = y − 1.5, its round end at 23.65, 25.2 | a circle at (13.475, 14.775), radius 6.225, stroke 1.3; the handle from (18.11, 19.61), where its round end meets the ring's inner edge, to (22.57, 24.07), stroke 2.2, round |
| the magnifier, 38 by 40 | 9.25 to 29.05, 9.9 to 30.05; the ring 9.25 to 25.2 across, 9.9 to 25.9 down, 1.45 thick; the handle's run 3.7 across (2.62 thick), its middle on x = y − 1 | a circle at (17.225, 17.9), radius 7.25, stroke 1.45; the handle from (22.6, 23.6) to (27.74, 28.74), stroke 2.62, round |
| the arrow, 9 by 16 | 0.25 to 8.75, 4.25 to 11.25; the shaft at 7.75 down, 1 thick; the head's strokes 1.4 across at 45° (1 thick), meeting at (8.25, 7.75) | the shaft from (0.75, 7.75) to (8.25, 7.75), the head from (5.25, 4.75) to it and back to (5.25, 10.75), stroke 1, round |
| the arrow, 13 by 25 | 0.45 to 12.55, 7.95 to 18.05; the shaft at 13 down, 1.3 thick; the head's strokes 1.7 across (1.2 thick), meeting at (11.93, 13) | the shaft from (1.1, 13), stroke 1.3; the head from (7.48, 8.55) to (11.93, 13) to (7.48, 17.45), stroke 1.2, round |

| icon | shared ink |
| --- | --- |
| the magnifier | 0.9742 |
| the magnifier, compact | 0.9697 |
| the arrow | 0.9825 |
| the arrow, compact | 0.9952 |

Intersection over union; the probe fails below 0.85, a mark set before the run, or if the search's example (`components/globalnav/examples/search.html`) draws others. Measured 2026-10-19 on https://www.apple.com/ in Chromium 151 on Linux (`results/searchfield.json`).

## The local nav's chevron

`localnav.mjs` reads it off MacBook Air's page: a polyline in 16 by 9 through (15.265, 0.835), (8, 8.167) and (0.735, 0.835), 1.2 thick, round at the ends and the join; opened, apple.com's animation ends on the same line upside down, (15.265, 8.165), (8, 0.835), (0.735, 8.165), in 320ms, through a straight line at 4.5. ciderui's `CHEVRON` is that line; `localnav.css` turns it up with `scaleY(-1)` about its middle, which passes through the same straight line. Shared ink: 1 closed, 0.996 open (`results/localnav.json`, 2026-10-19, Chromium 151).

## Running it

```sh
PW=…/playwright-chromium/index.mjs node run.mjs   # the "more" chevron
PW=…/playwright-chromium/index.mjs node nav.mjs   # the global nav's icons
PW=…/playwright-chromium/index.mjs node localnav.mjs   # the local nav's chevron
PW=…/playwright-chromium/index.mjs node searchfield.mjs   # the search's magnifier and arrow
```

`glyph.js` runs on apple.com's page, where its icon font is declared: it loads the face, draws the character on a canvas, reads the numbers off the ink, draws the chevron they describe on another, and compares the two. `nav.mjs` reads the nav's SVG off the page and compares on a blank one. The browsers are `../browsers.mjs`'s.

## A gallery's paddles

apple.com's 36px paddles draw a chevron in a 36 by 36 SVG (`icon-control-chevronleft` and `-right`, the iPhone page's and iPhone 18 Pro's): the outline of a line 3 thick, round at its ends and its corner, its ink 14.5 to 23 across and 11 to 25 down (the next; the previous mirrored, 13 to 21.5 across). ciderui's `PADDLE_NEXT` is that line through (16, 12.5), (21.5, 18) and (16, 23.5), `PADDLE_PREVIOUS` through (20, 12.5), (14.5, 18) and (20, 23.5): they share 0.9993 and 0.9995 of their ink with apple.com's (intersection over union; `paddle.mjs` fails below 0.85, and holds each drawing to the gallery's example). The store's 56px paddles draw a longer one (its ink 9.35 to 26.5 down), which ciderui does not: its large paddle is the 36px chevron, larger. Measured 2026-10-19 on https://www.apple.com/iphone/ in Chromium 151 on Linux (`results/paddle.json`).

## The checkbox's check

apple.com's sign-in widget draws its check as an SF Pro Icons glyph (`.form-checkbox-indicator::after`, 18px on 28px, white on the checked box). `check.mjs` captures the checked box at 20 pixels to the unit and reads its white ink inside the 14 by 14 content box: 2.65 to 11.85 across, 2.95 to 11.9 down, its strokes about 1.33 thick (a row's run across the long arm 1.55 at 58°, across the short one 1.7 at 53°). ciderui's `CHECK` is a line 1.35 thick, round at its ends and its corner, through (3.3, 7.65), (6.25, 11.35) and (11.15, 3.7), the best of a small grid of points about those the ink gives: it shares 0.9603 of its ink with apple.com's (intersection over union; the probe fails below 0.85, and holds the drawing to the checkbox's example). Measured 2026-10-19 on https://secure.store.apple.com/shop/account/home in Chromium 151 on Linux (`results/check.json`).

## The accordion's and the footer's chevrons

Both are apple.com's lines, drawn as its pages draw them: the bag's accordion's a polyline through (15, 1.13), (8.5, 7.72) and (2, 1.13) in 17 by 8.85, 2.25 thick, round (`DISCLOSURE`); a phone's footer section's through (10.075, 0.675), (5.5, 5.323) and (0.925, 0.675) in 11 by 6, 1 thick, round (`FOOTER_CHEVRON`). Turned up, each is the line its animation ends on (the same points mirrored). `lines.mjs` reads both live and holds ciderui's to them: 1 and 1 of the ink shared (fails below 0.85), each in its component's example. Measured 2026-10-19 in Chromium 151 on Linux (`results/lines.json`).


## A modal's X

The store's overlay draws its close button's X as a 20px SVG with no view box (`.rc-overlay-close svg`, its units the button's pixels), the outline of two lines 3 thick, round at their ends: its ink 3.35 to 16.7 across and 3.3 to 16.6 down, each round end's middle at 4.88 and 15.12 (the arcs' points, `a1.5 1.5` between them). ciderui's `CLOSE` is those two lines, from (4.88, 4.88) to (15.12, 15.12) and from (15.12, 4.88) to (4.88, 15.12), 3 thick, round, in 20 by 20: it shares 0.9927 of its ink with apple.com's (intersection over union; `close.mjs` fails below 0.85, and holds the drawing to the modal's example), the runs across its middle and its arms the same to the twentieth of a unit. Measured 2026-10-19 on iPhone 18 Pro's buy page, its first "Show more" opened, in Chromium 151 on Linux (`results/close.json`).
