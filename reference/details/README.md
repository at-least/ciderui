# The accordion's `<details>`

Does `cui-accordion` open with no script, and does its answer's height move? `page.mjs` writes `page.html`, the accordion's example on the kit's CSS compiled as a site compiles it; `run.mjs` clicks its second question (by the page's script on the iOS simulator, as `../popover/` does) and reads the item 120ms and 700ms on.

| browser | opens, answer shown | chevron turned | its height |
| --- | --- | --- | --- |
| Chromium 151 on Linux | ok | ok | moves (56px, 86.8px at 120ms, 117.6px) |
| Chrome 155 on macOS | ok | ok | moves |
| Safari 26.5 on the iOS 26.5 simulator | ok | ok | at once: no `interpolate-size` (its full height at 120ms) |

Run 2026-10-19 (`results/`). Safari on macOS not run: its automation session still times out (`../popover/`).

```sh
node page.mjs && PW=…/playwright-chromium/index.mjs BROWSER=chromium node run.mjs
```
