# The console parts in each browser

Do the console parts crashcart drew last (the tabs' phone row and its edge fade, the page column) do in each browser what `test/components.mjs` holds them to in Chromium? `build.mjs` writes the pages: ciderui compiled as a site compiles it, each part's example in a frame of a set width (the index.html it writes), so its media queries read the frame's width in every browser, a phone's too in Safari, whose window is wider; the toast a page of its own. `run.mjs` reads its claims (its header) through `../browsers.mjs`. The tabs' fade is read off the window's pixels, not off its computed `mask-image`: the probe gives the bar a flat blue ground, once on a `cui-light` page (white behind it, the red channel) and once on a `cui-dark` one (black behind it, the blue channel), and the pixel 1px under its top edge says where the mask lets the page through. Each page sets its scheme, so the system's appearance picks neither: on air (Dark) the one page that followed it was black behind the bar, and the fade read as no fade (red 0 at every point).

| browser | 1 one row, scrolling | 2 fade: start, end, between | 3 focused tab brought in whole, its ring inside | 4 page column (900, 1200, 1920) |
| --- | --- | --- | --- | --- |
| Chromium 151 on Linux | ok | ok (the edge 2px in: red 239 and 228; 12px in: 133 and 122) | ok (54.9px of 54.9px) | ok |
| Firefox 153.0 (Playwright's) on Linux | ok | ok (the same reds) | ok (54.5px of 54.9px) | ok |
| Safari 27.0, macOS 27.0 (a tester VM, `../vm.mjs`, 2026-10-29) | ok | ok (the edge 2px in: red 229 at the end) | ok (54.5px of 54.9px) | ok |
| Safari 27.0.1, macOS 27.0.1 (air over plain ssh, by hand, 2026-11-07) | ok | ok, light and dark (the edge 2px in: 231 and 235 light, 231 and 236 dark) | ok | ok |
| Safari 26.5, iOS simulator | not run: the same | | | | |
| Chrome 156, macOS | not run: the same | | | | |

Claim 3 failed first in Chromium: the last tab, 3.4px of it in the row, focused by script, was left where it was (3.375px of 54.9px shown); Firefox 153 brought it in. A browser's focus scrolls only a tab it judges out of view: with the row scrolled 0, 20 and 40px (the tab's start 3, 23 and 43px inside the row's end), both browsers left the 23px and 43px ones clipped. The phone row's `scroll-padding-inline: 24px` (the fade's width) made both bring in the 3px and 23px ones whole; the 43px one, its start outside the padding, they still leave clipped by 12px, which the host's script brings in: crashcart's, on focus, `scrollIntoView` to the nearest edge (its e2e/layout.mjs: each of six tabs whole once focused, the row at either end, in Chromium). Here the 43px one so scrolled lands whole at the row's end (Chromium and Firefox: the row scrolled to its end, 52px and 51px), the side the row's end state leaves unfaded.

Measured 2026-10-28 (`results/`). The thumbnails, pager and toast read here went by DESIGN.md decision 36.

## Running it

```sh
node build.mjs /tmp/ciderui-console
PW=…/playwright-chromium/index.mjs BROWSER=chromium SITE=/tmp/ciderui-console node run.mjs
BROWSER=firefox SITE=/tmp/ciderui-console node run.mjs   # Playwright's Firefox, Xvfb
BROWSER=safari SITE=/tmp/ciderui-console node ../vm.mjs node run.mjs   # in a tester VM of its own (../vm.mjs)
ssh air "pkill -x Safari; cd $PWD && BROWSER=safari SITE=/tmp/ciderui-console caffeinate -du node run.mjs"   # on air, the repository and SITE rsynced to the same paths; caffeinate: with air's display asleep Safari's picture came back all black
BROWSER=ios SITE=/tmp/ciderui-console node run.mjs     # on a Mac with Xcode; and chrome
```
