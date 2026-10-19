# A nav's list as its own menu

The probe behind `cui-globalnav`'s phone menu, which works without a script. The nav's list is a `popover`: closed, the page's CSS shows it as the bar's row (the browser's own style hides a closed popover, and the page's wins over it), beside the menu's button (`popovertarget` naming it); opened into the top layer, over the window, it is a sheet with a close button of its own, whose two lines turn into an X. The claims, each read by `popover.js` on `page.html` (the nav's techniques at their smallest) after the browser's own click and keys:

1. closed, the page shows the list as its row (`display: flex`, in the page, 48px);
2. its button opens it over the window (`position: fixed`, the window's height);
3. it grows from the bar: a frame, sampled each frame from the `toggle` event, between 48px and the window (a transition from the closed height);
4. its rows come in from transparent (`@starting-style`: the first frame's opacity 0, then 1);
5. the page is held while it is open (`:has()` on the root, `overflow: hidden`, as apple.com sets);
6. the focus is on its close button (`autofocus`, which a popover's opening honors);
7. the close button's lines turn from two lines into an X (a frame with a line under 1px tall, `@starting-style`; then each line's box 11.5 by 11.5 from (3.5, 3.5) in its 18px SVG, each turned, moved and stretched by CSS transforms about its middle: apple.com's X, `../icons/`);
8. Tab goes from the close button to the first link (Option-Tab in Safari on the Mac, whose Tab skips links);
9. the close button closes it at once (every frame 48px: the transitions are declared on the open state alone, so nothing runs back);
10. the focus is back on the button;
11. Escape closes it, the focus back on the button.

The menu's button was first inside the list, so that one button would be on top in both states: in Chromium 151 a popover opened by a button inside it holds the focus on that button (Tab and Shift+Tab stay there; from a link, Tab reaches the button and stays), which `showPopover()` and a button outside the list do not; so the button opens the list from beside it, and the list closes from inside.

## The result

| browser | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Chrome 155.0.8059.80, macOS 27.0.1 | ok | ok | ok | ok | ok | ok | ok | ok | ok | ok | ok |
| Safari 27.0.1, macOS 27.0.1, by its keyboard | ok | ok | ok | ok | ok | ok | ok | ok | ok | ok | ok |
| Safari 26.5, iOS 26.5 simulator | ok | ok | ok | ok | ok | ok | ok | not run | ok | not run | not run |
| Chromium 151, Linux | ok | ok | ok | ok | ok | ok | ok | ok | ok | ok | ok |

Measured 2026-10-19 (`results/`). On the iOS simulator the buttons are clicked by the page's script (`click()`, the activation a tap's click runs): safaridriver's taps there reached a button as `pointerdown` and `touchstart` and never a click, and its element click as nothing; a phone has no Tab or Escape. Safari on the Mac is used as its keyboard user uses it: each button focused and Return pressed, Option-Tab for Tab. A click there focuses no button and clears the focus a field had (WebDriver's click on a button: the body focused after it either way), so a menu a click opened has no button to give the focus back to: clicked as Chrome is, the X and Escape closed it with the focus on the body and on the hidden X (10 and 11), and Tab, which skips links there, left the focus on the X (8). Safari on the Mac was first not run for want of a Tab key in `../browsers.mjs`, not, as said then, for a session that timed out (Safari's log has that run's session created); not measured: Firefox.

## Running it

```sh
BROWSER=safari node run.mjs      # and ios, chrome
PW=…/playwright-chromium/index.mjs BROWSER=chromium node run.mjs
```

`run.mjs` serves `page.html` on 127.0.0.1, clicks and presses through `../browsers.mjs` (`click`, `press`: the browser's own input) and fails if a claim does not hold.
