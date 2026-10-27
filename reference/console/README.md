# The console parts in each browser

Do the console parts crashcart drew last (the tabs' phone row and its edge fade, the thumbnails, the pager, the page column, the toast) do in each browser what `test/components.mjs` holds them to in Chromium? `build.mjs` writes the pages: ciderui compiled as a site compiles it, each part's example in a frame of a set width (`index.html`), so its media queries read the frame's width in every browser, a phone's too in Safari, whose window is wider; the toast a page of its own. `run.mjs` reads its claims (its header) through `../browsers.mjs`. The tabs' fade is read off the window's pixels, not off its computed `mask-image`: the probe gives the bar a flat blue ground on white, and the red channel 1px under its top edge says where the mask lets the ground through.

| browser | 1 one row, scrolling | 2 fade: start, end, between | 3 focused tab's ring inside | 4 thumbnails (375, 1200) | 5 pager (375, 1200) | 6 page column (900, 1200, 1920) | 7 toast |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Chromium 151 on Linux | ok | ok (the edge 2px in: red 239 and 228; 12px in: 133 and 122) | ok | ok | ok | ok | ok (0.07 at 27ms, 1 at 213ms) |
| Firefox 153.0 (Playwright's) on Linux | ok | ok (the same reds) | ok | ok | ok | ok | ok (0.07 at 29ms, 1 at 215ms) |
| Safari 27, macOS | not run: the Mac did not answer (ping and ssh, 2026-10-28) | | | | | | |
| Safari 26.5, iOS simulator | not run: the same | | | | | | |
| Chrome 156, macOS | not run: the same | | | | | | |

Seen, not held (claim 3, `got.observed`): the last tab, 3.4px of it in the row, focused by script. Chromium 151 left it where it was, 3.4px of its 54.9px shown; Firefox 153 scrolled the row to show it whole. A probe with the row scrolled 0, 20 and 40px (its start 3, 23 and 43px inside the row's end) found both browsers leaving a tab 23px or 43px in view clipped. With `scroll-padding-inline: 24px` on the bar, both scrolled a tab 3px or 23px in view whole, and still left one 43px in view clipped by 12px.

Measured 2026-10-28 (`results/`).

## Running it

```sh
node build.mjs /tmp/ciderui-console
PW=…/playwright-chromium/index.mjs BROWSER=chromium SITE=/tmp/ciderui-console node run.mjs
BROWSER=firefox SITE=/tmp/ciderui-console node run.mjs   # Playwright's Firefox, Xvfb
BROWSER=safari SITE=/tmp/ciderui-console node run.mjs    # on the Mac, Safari quit first (../browsers.mjs); and ios, chrome
```
