# A nav's list as its own menu

The probe behind `cui-globalnav`'s phone menu, which works without a script. The nav's list is a `popover`: closed, the page's CSS shows it as the bar's row (the browser's own style hides a closed popover, and the page's wins over it); the menu's button is inside it (`popovertarget` naming it), so when the list opens into the top layer, over the window, the button is still on top, and its two lines turn into an X. The claims, each read by `popover.js` on `page.html` (the nav's techniques at their smallest) after the browser's own click on the button:

1. closed, the page shows the list as its row (`display: flex`, in the page, 48px);
2. its button opens it over the window (`position: fixed`, the window's height);
3. it grows from the bar: a frame, sampled each frame from the `toggle` event, between 48px and the window (a transition from the closed height);
4. its rows come in from transparent (`@starting-style`: the first frame's opacity 0, then 1);
5. the page is held while it is open (`:has()` on the root, `overflow: hidden`, as apple.com sets);
6. the lines are an X, each line's box 11.5 by 11.5 from (3.5, 3.5) in the button's 18px SVG (each closed line turned, moved and stretched by CSS transforms about its middle: apple.com's X, `../icons/`);
7. the same button closes it at once (every frame 48px: the transitions are declared on the open state alone, so nothing runs back);
8. Escape closes it.

## The result

| browser | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Chrome 155.0.8059.80, macOS 27.0.1 | ok | ok | ok | ok | ok | ok | ok | ok |
| Safari 26.5, iOS 26.5 simulator | ok | ok | ok | ok | ok | ok | ok | not run |
| Chromium 151, Linux | ok | ok | ok | ok | ok | ok | ok | ok |

Measured 2026-10-19 (`results/`). On the iOS simulator the button is clicked by the page's script (its `click()`, the activation a tap's click runs): safaridriver's taps there reached the button as `pointerdown` and `touchstart` and never a click, and its element click as nothing; a phone has no Escape. Not yet run: Safari on macOS, whose automation session timed out finding a Safari to drive (safaridriver: "waiting for its RWIApplication to appear"); not measured: Firefox.

## Running it

```sh
BROWSER=safari node run.mjs      # and ios, chrome
PW=…/playwright-chromium/index.mjs BROWSER=chromium node run.mjs
```

`run.mjs` serves `page.html` on 127.0.0.1, clicks and presses through `../browsers.mjs` (`click`, `press`: the browser's own input) and fails if a claim does not hold.
