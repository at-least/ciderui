# The font stack

The probe behind ciderui's font stack, `-apple-system, BlinkMacSystemFont, Inter, "Helvetica Neue", Helvetica, Arial, sans-serif`, with the kit's `fonts.css` declaring Inter. The two claims it holds: on Apple's platforms the stack draws the system font and the page never fetches Inter; elsewhere it draws Inter, fetching its Latin face alone.

Why these two names and not `system-ui`: `system-ui` is a font on every platform (Segoe UI, Roboto, Noto Sans), so Inter after it would never be drawn anywhere; `-apple-system` is a font in Safari alone and `BlinkMacSystemFont` in Chrome alone, both SF on Apple's platforms and no font elsewhere (`../tracking/`: each lands on apple.com's widths within 0.0006em in its own browser, and is 0.064em off in the other's).

## The result

| browser | fetched by first paint | stack | system-ui | Inter | |
| --- | --- | ---: | ---: | ---: | --- |
| Safari 27.0.1, macOS 27.0.1 | nothing; Inter's 16 faces `unloaded` | 491.896px | 491.896px | 512.845px | the system font |
| Chrome 155.0.8059.80, macOS 27.0.1 | nothing; Inter's 16 faces `unloaded` | 491.906px | 491.906px | 512.859px | the system font |
| Safari 26.5, iOS 26.5 simulator | nothing; Inter's 16 faces `unloaded` | 491.896px | 491.896px | 512.845px | the system font |
| Chromium 151, Linux | `inter-roman-latin.woff2` alone | 512.813px | 504.797px | 512.813px | Inter |

Measured 2026-10-19 (`results/`). Not measured: Firefox, Windows, Android, a real iPhone.

iOS Safari's text autosizing drew the stack's line 1.31 times as wide (645.435px) until the page had `-webkit-text-size-adjust: 100%`, which Tailwind's preflight (4.3.3) sets on every ciderui page; `page.html` sets it as the preflight does.

## The method

`run.mjs` serves the repository on 127.0.0.1 and opens `page.html`, one line of text in the stack (17px) under the kit's `fonts.css`. `stack.js` waits two frames and a second, reads which `.woff2` files the page has fetched (resource timing) and the states of Inter's faces, then the line's width against the same text in `system-ui` and, loading it on purpose only then, in Inter. Widths equal within 0.01px are the same font.

```sh
BROWSER=safari EXPECT=sf node run.mjs      # and ios, chrome
PW=…/playwright-chromium/index.mjs BROWSER=chromium EXPECT=inter node run.mjs
```

`EXPECT=sf` passes if the line is the system font's width and not Inter's, no Inter face was fetched and every one is still `unloaded`; `EXPECT=inter` if the line is Inter's width and not the system font's, and Inter's Latin face was fetched. The browsers are `../browsers.mjs`'s; on the Mac it runs from a copy of the repository with `playwright-core` in a `node_modules` above it.
