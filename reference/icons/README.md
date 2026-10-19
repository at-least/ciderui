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

## Running it

```sh
PW=…/playwright-chromium/index.mjs node run.mjs
```

`glyph.js` runs on apple.com's page, where its icon font is declared: it loads the face, draws the character on a canvas, reads the numbers off the ink, draws the chevron they describe on another, and compares the two. The browsers are `../browsers.mjs`'s.
