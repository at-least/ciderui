# A link to a place on the page, under the local nav

Where does a link to a place on its page (a footnote's mark) land its target, on a page whose local nav sticks to the window's top? `apple.mjs` records apple.com's: three of its pages in Playwright's Chromium at 1440 and 390px, each scrolled 2000px down, its first footnote mark clicked, the note's top read 2s on against the nav's foot, with what sets the note's `scroll-margin-top` (the rules in the page's stylesheets that set it on the note, its inline style) and the root's `scroll-padding-top`.

| page | its local nav | 1440px | 390px |
| --- | --- | --- | --- |
| iPhone 18 Pro | a newer one (`nav#localnav`), 72px | at its foot (the note 71.94px down, the foot 72px) | at its foot (72.27px, 72px) |
| MacBook Air | the one ciderui's is (`#ac-localnav`), 52px, 48px to 833px | under it (0.06px, the foot 52px) | under it (0.38px, 48px) |
| AirPods Pro | the same | under it (0.34px, 52px) | under it (−0.47px, 48px) |

On iPhone 18 Pro's page the note's `scroll-margin-top` is 0px before the click and 72px after: a `:target` rule in the page's own stylesheet sets it to the nav's height (`var(--localnav-total-height,0)`, 52px and 10px above and below); on the other two no rule sets it, 0px before and after. On all three no inline style sets it and the root's `scroll-padding-top` is `auto`. Recorded 2026-10-22 (`results/apple.json`).

ciderui's local nav lands the target at its foot, as iPhone 18 Pro's page does: while the page has one, `:target`'s `scroll-margin-top` is the bar's height, 48px, 52px from 834px (`components/localnav/localnav.css`). Its claims:

- the local nav's cases `target` and `target-phone` (`test/components.mjs`, Chromium): a mark under the bar clicked, its note lands 52px down at 834px and 48px at 833px; before the rule it landed 0px down, under the bar;
- `test/docs.mjs` (Chromium): on the documentation's news article, its markdown footnote's mark clicked, the note lands below the page's local nav at 1440 and 390px; before the rule, at 0px under the 52px bar at 1440px, and 42.16px under the 48px one at 390px;
- `../docs/`'s claim 14 in every browser it runs in, the same page: the note at the nav's foot, or below it where the page ends first. Before the rule (the site built with ciderui-zola's stylesheet as it was) Chromium read the note at 0px at 1440px, the nav's foot 52px.

| browser | 1440px | 700px (a phone's width on the simulator) |
| --- | --- | --- |
| Safari 27.0.1, macOS 27.0.1 | at its foot (52px, 52px) | at its foot (48.16px, 48px) |
| Safari 26.5, iOS 26.5 simulator | | at 402px, at its foot (48.16px, 48px) |
| Chrome 156, macOS 27.0.1 | at its foot (52px, 52px) | below it, the page at its end first (74.16px, 48px) |
| Chromium 151, Linux | at its foot (52px, 52px) | below it, the page at its end first (74.16px, 48px) |
| Firefox 153.0 (Playwright's), Linux | at its foot (52px, 52px) | below it, the page at its end first (74.17px, 48px) |

```sh
node apple.mjs     # PW=…/playwright-chromium/index.mjs, as ../browsers.mjs takes it
```
