# The accordion's `<details>`

Does `cui-accordion` open with no script, and does its answer's height move? `page.mjs` writes `page.html`, the accordion's example on the kit's CSS compiled as a site compiles it; `run.mjs` clicks its second question (by the page's script on the iOS simulator, as `../popover/` does) and reads the item 120ms and 700ms on.

And is the focus ring of what is in an open answer shown whole? The answer's box clips what is outside it as its height moves, and the ring of what begins a line in it, the kit's one ring 3px out (`tokens.css`), is outside it. `page.mjs` writes `ring.html` too, 24px in from the window's edge: the bag's structure (the compact example as the large question's answer, opened, as the components' `nested` case) and the compact example with a link beginning its first answer. Where there is a keyboard (not on the simulator), `run.mjs` presses Tab (Option-Tab in Safari, whose Tab skips links) until the compact question in the large one's answer, then the link, has the focus, at 1200 and 390px, and reads the window's picture in the page (`ring.js`): the accent's pixels in the 8px beside it, left and right, where the ring's sides are drawn. A side is shown with at least a column of the screen's pixels the element's height in the accent. Its height is to be clipped all the same: `held.html` holds an open answer 10px tall, its text 13.6px down, and `run.mjs` reads the ink in its paragraph's box (pixels darker than mid-gray), none where it is clipped; `held-visible.html`, the same clipping nothing (`overflow: visible`), is the control, its text drawn.

| browser | opens, answer shown | chevron turned | its height | held 10px: its text | the question's ring in the large one's answer | the ring of a link beginning an answer |
| --- | --- | --- | --- | --- | --- | --- |
| Chromium 151 on Linux | ok | ok | moves (56px, 70.3px at 120ms, 117.6px) | clipped (no ink; 723 pixels of it with nothing clipped) | whole at 1200 and 390px | whole at 1200 and 390px |
| Chrome 156 on macOS | ok | ok | moves (56px, 70.3px at 120ms, 117.6px) | clipped (no ink; 685) | whole at 1200 and 390px | whole at 1200 and 390px |
| Safari 27.0.1 on macOS 27.0.1 | ok | ok | at once: no `interpolate-size` (56px, its full 117.6px at 120ms) | not read | not read | not read |
| Safari 26.5 on the iOS 26.5 simulator | ok | ok | at once: no `interpolate-size` (its full height at 120ms) | not read | no keyboard | no keyboard |
| Firefox 153.0 (Playwright's) on Linux | ok | ok | at once: no `interpolate-size` (56px, its full 117.6px at 120ms) | clipped (no ink; 741) | whole at 1200 and 390px | whole at 1200 and 390px |

Run 2026-10-22 in Chromium, Chrome and Firefox (`results/`), with the ring and the held answer; Safari and the simulator 2026-10-19, before them (Firefox is Playwright's build, Gecko with its automation patched in, in its window on a virtual X display: `../browsers.mjs`). Safari on the Mac clicks the question as Chrome does. Safari's ring is not read: the owner was at the Mac that day, and a Safari run takes its front, in a window not in front Safari matches no `:focus` (`../docs/`), so its ring would not be drawn (UNVERIFIED).

The ring first failed in Chromium and Firefox, the answer's box clipping both its axes (`overflow: clip`): its left side was not shown, 0 of the accent's pixels, nor its right at 390px, where the answer runs to the window's edge less its 24px (from 735px the answer stops 35px short of its end, room for the right side). Now the answer clips its height alone (`overflow: visible clip`): its sides shown whole, 142 of the accent's pixels a side for the question at 1200px, 136 at 390px, and its height clipped still, held 10px tall.

```sh
node page.mjs && PW=…/playwright-chromium/index.mjs BROWSER=chromium node run.mjs
BROWSER=safari node run.mjs      # Safari quit first (../browsers.mjs); and ios, chrome, firefox
```
