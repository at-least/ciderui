# A link to a place on the page, under the local nav

Where does a link to a place on its page (a footnote's mark) land its target, on a page whose local nav sticks to the window's top? `apple.mjs` records apple.com's: three of its pages in Playwright's Chromium at 1440 and 390px, each scrolled 2000px down, its first footnote mark clicked, the note's top read 2s on against the nav's foot, with what sets the note's `scroll-margin-top` (the rules in the page's stylesheets that set it on the note, its inline style) and the root's `scroll-padding-top`.

| page | its local nav | 1440px | 390px |
| --- | --- | --- | --- |
| iPhone 18 Pro | a newer one (`nav#localnav`), 72px | at its foot (the note 71.94px down, the foot 72px) | at its foot (72.27px, 72px) |
| MacBook Air | the older one (`#ac-localnav`), 52px, 48px to 833px | under it (0.06px, the foot 52px) | under it (0.38px, 48px) |
| AirPods Pro | the same | under it (0.34px, 52px) | under it (−0.47px, 48px) |

On iPhone 18 Pro's page the note's `scroll-margin-top` is 0px before the click and 72px after: a `:target` rule in the page's own stylesheet sets it to the nav's height (`var(--localnav-total-height,0)`, 52px and 10px above and below); on the other two no rule sets it, 0px before and after. On all three no inline style sets it and the root's `scroll-padding-top` is `auto`. Recorded 2026-10-22 (`results/apple.json`).

ciderui has no local nav (decision 33). A local nav brought back wants iPhone 18 Pro's way, `:target`'s `scroll-margin-top` the bar's height: without it the note lands 0px down, under the bar.

```sh
node apple.mjs     # PW=…/playwright-chromium/index.mjs, as ../browsers.mjs takes it
```
