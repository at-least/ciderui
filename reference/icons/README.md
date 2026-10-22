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

ciderui's `MORE`: two strokes of that thickness, round at the ends and the join, from the ink's corners (inset by half the thickness) to its middle: (0.1077, 0.4823) to (0.3323, 0.2575) to (0.1077, 0.0327), above the baseline. Drawn on the same canvas at 400px, it shares 0.9481 of its ink with apple.com's glyph (intersection over union; the probe fails below 0.9, a mark set before the run). Measured 2026-10-19 on https://www.apple.com/iphone/ in Chromium 151 on Linux (`results/more.json`); again 2026-10-22, the chevron compared now the one the kit ships, read off the more link's example (`shippedIou`, 0.9481, the numbers' own the same), where the probe had compared one drawn from the numbers alone and passed whatever the kit shipped (its stroke halved in the example, the probe now fails: 0.447).

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

The global nav's search, opened, draws four more SVGs, each a regular one from 834px and a compact one below: the field's magnifier, the arrow before each quick link, the clear button's mark (shown once the field holds text), and the magnifier before each suggested search (shown as one types). `searchfield.mjs` reads the first three off the home page (the flyout's markup is in the page closed) and the last off the page's settings for its script (its `suggestedSearches` images, in a script of the page's own; its suggested links have the quick links' arrow), and compares them as `nav.mjs` does:

| icon | apple.com's ink | ciderui's |
| --- | --- | --- |
| the field's magnifier, 30 by 32 | 6.6 to 23.65 across, 7.9 to 25.2 down; the ring's outer edge 6.6 to 20.35 across and 7.9 to 21.65 down, 1.25 to 1.3 thick; the handle's run 3.1 across at 45° (2.2 thick), its middle on x = y − 1.5, its round end at 23.65, 25.2 | a circle at (13.475, 14.775), radius 6.225, stroke 1.3; the handle from (18.11, 19.61), where its round end meets the ring's inner edge, to (22.57, 24.07), stroke 2.2, round |
| the magnifier, 38 by 40 | 9.25 to 29.05, 9.9 to 30.05; the ring 9.25 to 25.2 across, 9.9 to 25.9 down, 1.45 thick; the handle's run 3.7 across (2.62 thick), its middle on x = y − 1 | a circle at (17.225, 17.9), radius 7.25, stroke 1.45; the handle from (22.6, 23.6) to (27.74, 28.74), stroke 2.62, round |
| the arrow, 9 by 16 | 0.25 to 8.75, 4.25 to 11.25; the shaft at 7.75 down, 1 thick; the head's strokes 1.4 across at 45° (1 thick), meeting at (8.25, 7.75) | the shaft from (0.75, 7.75) to (8.25, 7.75), the head from (5.25, 4.75) to it and back to (5.25, 10.75), stroke 1, round |
| the arrow, 13 by 25 | 0.45 to 12.55, 7.95 to 18.05; the shaft at 13 down, 1.3 thick; the head's strokes 1.7 across (1.2 thick), meeting at (11.93, 13) | the shaft from (1.1, 13), stroke 1.3; the head from (7.48, 8.55) to (11.93, 13) to (7.48, 17.45), stroke 1.2, round |
| the clear mark, 14 by 14 | a filled disc 0.05 to 13.95 (radius 6.9661 at (7, 7)), an X cut out of it: its outline's strokes 1.1 thick (half-round ends of radius 0.55, `a.55.55`, whose middles are at 4.591 and 9.409), meeting 0.778 from the middle; the runs across its middle 0.05 to 6.2 and 7.8 to 13.95 | the disc, and the X's outline cut out of it (`fill-rule="evenodd"`): ends' middles at (4.591, 4.591), (9.409, 9.409), (9.409, 4.591), (4.591, 9.409), half-width 0.55, one path |
| the clear mark, 16 by 16 | a disc 0 to 16 (radius about 8), the X's strokes about 1.34 thick, its ends drawn as curves (radius about 0.67), their middles about 2.81 from the disc's | the disc, radius 8, the X's outline: ends' middles 2.81 from (8, 8) along the diagonals, half-width 0.67 |
| a suggested search's magnifier, 9 by 16 | 0.2 to 8.75 across, 3.4 to 12.15 down; the ring 0.2 to 7.15 across and 3.45 to 10.35 down, 0.7 thick (its middle at (3.675, 6.9)); the handle's run 1.55 across at 45° (1.1 thick, `a.55.55`), its middle on y = x + 3.381, its round end at 8.75, 12.15 | a circle at (3.675, 6.9), radius 3.125, stroke 0.7; the handle from (6, 9.381), where its round end stays out of the ring's hole, to (8.2, 11.581), stroke 1.1, round |
| the magnifier, 13 by 25 | 0.5 to 12.6, 6.8 to 19.15; the ring 0.55 to 10.3 across, 6.85 to 16.6 down, 0.9 thick (its middle at (5.425, 11.725)); the handle's run 2.25 across (1.6 thick), its middle on y = x + 6.55 | a circle at (5.425, 11.725), radius 4.425, stroke 0.9; the handle from (8.692, 15.242) to (11.8, 18.35), stroke 1.6, round |

| icon | shared ink |
| --- | --- |
| the magnifier | 0.9742 |
| the magnifier, compact | 0.9697 |
| the arrow | 0.9825 |
| the arrow, compact | 0.9952 |
| the clear mark | 0.9995 |
| the clear mark, compact | 0.9928 |
| a suggested search's magnifier | 0.9639 |
| the magnifier, compact | 0.9792 |

Intersection over union; the probe fails below 0.85, a mark set before the run, or if the search's examples (`components/globalnav/examples/search.html`, `suggestions.html`) draw others. Measured 2026-10-19 on https://www.apple.com/ in Chromium 151 on Linux, the clear mark and a suggested search's magnifier 2026-10-21 (`results/searchfield.json`).

## The flyouts' chevrons

A link's flyout in the global nav (`../apple.com/`, `flyout:*`) draws three more SVGs: the chevron of the button beside the link, from 834px (shown when the button has the focus from the keyboard), the same button's on a phone, and the back button's on a phone. `flyout.mjs` reads them off the home page (the flyouts' markup is in the page closed), and the chevron a hovered row of the phone's menu shows, apple.com's drawn pointing back and turned by its stylesheet (`scaleX(-1)`, read at 390px), turned here the same:

| icon | apple.com's ink | ciderui's |
| --- | --- | --- |
| the button's chevron, 10 by 10 | the outline of a line 0.89 thick, round at its ends and its corner: 1.9 to 8.1 across, 3.25 to 6.75 down; its ends' middles at (2.33, 3.66) and (7.67, 3.66), its corner's at (5, 6.34) (its outer edge at 6.775, its inner at 5.725 down) | a line through (2.33, 3.66), (5, 6.34) and (7.67, 3.66), 0.88 thick, round |
| the button's chevron on a phone, 9 by 48 | the outline of a line 1.2 thick, round: 0.7 to 8.9 across, 16.6 to 31.4 down; its ends' middles at (1.3, 17.21) and (1.3, 30.79), its corner's at (8.3, 24.07) | a line through those three, 1.2 thick, round |
| the back button's, 9 by 48 | the same turned: its ends' middles at (7.7, 17.21) and (7.7, 30.79), its corner's at (0.7, 24.07) | a line through those three, 1.2 thick, round |
| a hovered row's, 9 by 48, turned | the back button's line, `scaleX(-1)` | the button's chevron on a phone |

| icon | shared ink |
| --- | --- |
| the button's chevron (`FLYOUT`) | 0.99 |
| on a phone (`FLYOUT_COMPACT`) | 0.9968 |
| the back button's (`FLYOUT_BACK`) | 0.9962 |
| a hovered row's (`FLYOUT_COMPACT`) | 0.9962 |

Intersection over union; `flyout.mjs` fails below 0.85, a mark set before the run, or if the flyout's example (`components/globalnav/examples/flyout.html`) draws others. Measured 2026-10-21 on https://www.apple.com/ in Chromium 151 on Linux (`results/flyout.json`).

## The field's icons

`form.mjs` reads the store's two form icons. The error's mark before a field's message is an SF Pro Icons glyph, U+F312, read as the "more" chevron is (drawn at 400px on the gift card's page, where the font is declared): its advance 1.1709em, its ink 0.0875 to 1.0825em across and 0.85 to −0.145em about the baseline; a ring, its rows' runs 0.0825em thick; in it a stem 0.265 to 0.605em up, 0.04 to 0.085em wide (tapering), and a dot 0.1075 to 0.2075em up. In the page's message row (12px, at four device pixels to the CSS pixel) its ink is 1.75 to 13.73px across and 2.75 to 14.73px down the 16px row (read again 2026-10-22, the same; a row with no red ink left of the text now fails the run, where its scan gave a box inside out, its left past its right, which the check took as a place: tried on pictures drawn for it, with a mark and without). The dropdown's chevron is an SVG its background draws, 21 by 21 at 25px: a line 1.45 thick (its rows' runs 2 to 2.1 across at 45°), its ink 5.3 to 15.7 across and 8.55 to 14.45 down, in `#6e6e73`.

| icon | ciderui's | shared ink |
| --- | --- | --- |
| `ALERT`, 1.171em by 0.995em (from the ink's top) | a ring at (585, 497.5) thousandths, radius 456.25, 82.5 thick; a stem from (584, 280) to (584, 550), 70 thick, round; a dot at (585, 692.5), radius 50 | 0.9695 |
| `DROPDOWN`, 21 by 21 drawn at 25 | a line through (6.075, 9.275), (10.5, 13.7) and (14.925, 9.275), 1.45 thick, round | 0.9689 |

Intersection over union; `form.mjs` fails below 0.85, a mark set before the run, or if the field's examples (`error.html`, `select.html`) draw others. Measured 2026-10-20 in Chromium 151 on Linux (`results/form.json`).

## The local nav's chevron

`localnav.mjs` reads it off MacBook Air's page: a polyline in 16 by 9 through (15.265, 0.835), (8, 8.167) and (0.735, 0.835), 1.2 thick, round at the ends and the join; opened, apple.com's animation ends on the same line upside down, (15.265, 8.165), (8, 0.835), (0.735, 8.165), in 320ms, through a straight line at 4.5. ciderui's `CHEVRON` is that line; `localnav.css` turns it up with `scaleY(-1)` about its middle, which passes through the same straight line. Shared ink: 1 closed, 0.996 open (`results/localnav.json`, 2026-10-19, Chromium 151).

## Running it

```sh
PW=…/playwright-chromium/index.mjs node run.mjs   # the "more" chevron
PW=…/playwright-chromium/index.mjs node nav.mjs   # the global nav's icons
PW=…/playwright-chromium/index.mjs node localnav.mjs   # the local nav's chevron
PW=…/playwright-chromium/index.mjs node searchfield.mjs   # the search's magnifier and arrow
PW=…/playwright-chromium/index.mjs node flyout.mjs   # the flyouts' chevrons
PW=…/playwright-chromium/index.mjs node form.mjs   # the field's error mark and the dropdown's chevron
PW=…/playwright-chromium/index.mjs node bag.mjs   # the bag's profile icons
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

The marketing pages' modal (the iPhone page's trade-in modal, `ric-modal`) draws its X white in a 36 by 36 SVG over its circle: one path, its ink 12 to 24 across and down, the same strokes as the store's (3 thick, the runs across its middle 15.85 to 20.15, as `CLOSE`'s drawn in the box's middle), each arm 0.6 shorter at its ends. `CLOSE` drawn there shares 0.8875 of its ink with it (its ink 11.4 to 24.6); ciderui's `CLOSE_SMALL` is the same two lines from (5.5, 5.5) to (14.5, 14.5) and from (14.5, 5.5) to (5.5, 14.5) in 20 by 20, and shares 0.9994 (`close.mjs` reads it too, and holds the drawing to the page modal's example). Measured 2026-10-21 in Chromium 151 on Linux (`results/close.json`'s `marketing`).

## The bag's profile icons

The global nav's bag, opened (`../apple.com/`, `bag-flyout:*`), draws an icon before each of its four profile links, a regular SVG from 834px (11 by 16) and a compact one to 833px (16 by 25), in its markup its script fetches: a box (`shippingbox`, Orders), a bookmark (Your Saves), a gear (Account) and a person in a circle (Sign in). Their fill, `6E6E73` without its `#`, is no color: the link's color draws them. `bag.mjs` reads them off the home page with the bag opened and compares as `nav.mjs` does; the box's, the bookmark's and the person's numbers are read off the outlines' points and arcs (the edges' slope, a stroke's two edges, a corner's radius), the gear's off its ink in rings about its middle (its ring's, hub's, spokes' and teeth's radii, their share of each ring, the teeth's places every 20°):

| icon | apple.com's | ciderui's |
| --- | --- | --- |
| the box, 11 by 16 | a hexagon 0.875 thick, its edges at 29.43°, its corners 1.5 round outside; its top face's lower edges meeting at 7.375 down over a line to its foot; a band across its top face from its upper left edge, 0.875 thick, cut square at its end | a hexagon through (5.511, 2.548), (9.738, 4.933), (9.738, 9.684), (5.511, 12.068), (1.285, 9.684) and (1.285, 4.933), its corners rounded 1.0625; the top face's edges from the side corners' arcs, (1.43, 5.018) and (9.593, 5.018), to (5.511, 7.375), and the line down to 11.91, the bottom corner's arc; the band from (3.849, 3.988) to (7.254, 5.878), square ends; 0.875 thick, round |
| the box, 16 by 25 | the same, 1.2 thick, its corners about 2.3 round outside | through (8, 5.74), (14, 9.125), (14, 15.875), (8, 19.259), (2, 15.875) and (2, 9.125), corners 1.7; the edges from (2.232, 9.26) and (13.768, 9.26) to (8, 12.594), the line to 19.007; the band from (5.561, 7.806) to (10.456, 10.522); 1.2 thick |
| the bookmark, 11 by 16 | its sides 0.82 thick (2.31 to 3.13 and 7.87 to 8.69), its top 0.73, its outer corners 1.59 round and inner 0.78; its notch 0.875 thick at 45°, its outer point at (5.5, 9.94), its legs' round ends 12.38 down | a line from (2.719, 11.97) up to 2.789, its corners rounded 1.184, across to 8.281 and down to 11.97, then up to (5.5, 9.3) and back, 0.82 thick, round |
| the bookmark, 16 by 25 | its sides and top 1.2 thick, its corners 2.3 round outside and 1.1 inside; its notch's points at 16.34 outside and 15.15 inside | the same from (4, 19.25) to 5.75, corners 1.7, across to 12, the notch up to (8, 15.75), 1.2 thick |
| the gear, 11 by 16 | about (5.5, 7.25): a hub 0.6 to 1.2 out, three spokes about 0.69 thick at 0°, 120° and 240°, a ring 3.66 to 4.45 out, 18 teeth every 20° to 5.2 out, about 0.7 wide | a hub circle of radius 0.9, 0.65 thick; the spokes from 1.2 to 3.7 out, 0.69 thick; a ring of radius 4.05, 0.79 thick; the teeth from 4.4 to 4.82 out, 0.68 thick, round |
| the gear, 16 by 25 | about (8, 12.5): a hub 0.8 to 2.05 out, the spokes about 1.21 thick, a ring 5.45 to 6.75, the teeth to 7.7, about 1.07 wide | a hub of radius 1.42, 1.25 thick; the spokes from 2 to 5.5, 1.21 thick; a ring of radius 6.1, 1.28 thick; the teeth from 6.7 to 7.15, 1.05 thick |
| the person, 11 by 16 | a ring about (5.5, 7.633), 4.731 out and 3.844 in; a head, an ellipse about (5.5, 6.71), 1.605 by 1.69; shoulders filled under an arc of radius 3.898 from (2.62, 10.163) to (8.38, 10.163) | a ring of radius 4.287, 0.887 thick; the head; the shoulders from (2.62, 10.163) up the arc to (8.38, 10.163) and back along the ring's inside, filled |
| the person, 16 by 25 | a ring about (7.99, 12.5), 7.1 out and 5.89 in; a head about (7.99, 11), 2.475 by 2.6; shoulders under curves from (3.529, 16.347) to (12.452, 16.347) through (7.99, 14.669) | a ring of radius 6.495, 1.21 thick; the head; the shoulders under an arc of radius 6.77 through the same points, filled |

| icon | shared ink |
| --- | --- |
| the box | 0.9915 |
| the box, compact | 0.9851 |
| the bookmark | 0.9518 |
| the bookmark, compact | 0.9723 |
| the gear | 0.9387 |
| the gear, compact | 0.9675 |
| the person | 0.972 |
| the person, compact | 0.986 |

Intersection over union; the probe fails below 0.85, a mark set before the run, or if the bag's example (`components/globalnav/examples/bag.html`) draws others. Measured 2026-10-22 on https://www.apple.com/ in Chromium 151 on Linux (`results/bag.json`).
