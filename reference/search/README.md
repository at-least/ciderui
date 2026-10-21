# A popover as the global nav's search

The probe behind the global nav's search, apple.com's flyout under the bar and its sheet on a phone (`../apple.com/`, `search:*`), on HTML's popover, which needs no script: the bar's search button opens it (`popovertarget`), its field takes the focus (`autofocus`), Escape and a click outside it close it. From 834px the flyout hangs from the bar's bottom, placed there by anchor positioning (`anchor-name` on the bar, `top: anchor(bottom)` on the flyout), with the curtain under its panel to the window's bottom; to 833px it is a sheet over the window and the bar, with a close button of its own, the page held. The claims, each read by `search.js` on `page.html` (the search's techniques at their smallest: a sticky bar under a 70px banner, its list a closed popover shown as its row, as the nav's is) after the browser's own clicks and keys, and in pictures of the window:

1. the browser has anchor positioning;
2. closed, the bar's list is its row and the flyout is not shown;
3. the search button, in the bar's closed list, opens the flyout, with no script;
4. the focus is on its field (`autofocus`);
5. it hangs from the bar's bottom, under the banner (114px down), across the window;
6. its curtain fills the window under its panel;
7. the page scrolled by 200px, the bar stuck at the top, the flyout follows it (44px down);
8. in the picture the bar is its own color, not under the curtain, and the panel `#fafafc`;
9. in the picture the curtain blurs the page under it: along a row of the page's 4px stripes, the brightness's spread under a fifth of its spread closed (`backdrop-filter: blur(20px)` on the curtain);
10. the curtain's middle hits the curtain, not the page; the bar hits its link;
11. a popover's `::backdrop` takes no click (`pointer-events: none`), so the curtain is a button of its own (`popovertargetaction="hide"`, out of the tab order and hidden from assistive technology);
12. a click on the bar's link closes it (outside it: light dismiss) and reaches the link;
13. a click on the curtain closes it and reaches nothing under it;
14. Escape closes it, the focus back on the search button;
15. its button closes it again;
16. Return in its field sends the form to the site's search page (GET, `?q=`; the form has no submit button, so Return sends it);
17. as a sheet its button opens it over the window and the bar;
18. as a sheet the focus is on its field;
19. as a sheet the page is held while it is open, and not after;
20. as a sheet its close button closes it;
21. as a sheet the focus is back on the search button;
22. empty, its field's clear button is hidden (`visibility: hidden`);
23. with text in the field, its clear button shows, with no script (`:has(input:not(:placeholder-shown))`; the text set by the page's script, as the CSS reads the value);
24. its clear button, a reset button (`type="reset"`), clicked empties the field and hides itself, the flyout still open;
25. then Shift+Tab takes the focus back to the field;
26. and Escape still closes it.

## The result

| browser | 1–11 | 12 | 13 | 14 | 15 | 16 | 17–20 | 21 | 22–23 | 24 | 25–26 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Safari 27.0.1, macOS 27.0.1, its buttons by its keyboard | ok | ok | ok | ok | ok | ok | ok | ok | ok | ok | ok |
| Safari 26.5, iOS 26.5 simulator | ok | not run | ok, by script | not run | ok | not run | ok | not run | ok | ok, by script | not run |
| Chrome 155.0.8059.85, macOS 27.0.1 | ok | ok | ok | ok | ok | ok | ok | ok | ok | ok | ok |
| Chromium 151, Linux | ok | ok | ok | ok | ok | ok | ok | ok | ok | ok | ok |
| Firefox 153.0 (Playwright's), Linux, in its window on a virtual X display | ok | ok | ok | ok | ok | ok | ok | ok | ok | ok | ok |

Measured 2026-10-21 (`results/`; 1 to 21 first on 2026-10-19, in Chrome 155.0.8059.80), at 1440 by 900 (the simulator's phone at its own 402 by 714). Clicked as `../popover/` clicks: on the iOS simulator by the page's script (safaridriver's taps never click there), so there is no click outside the flyout and no key, and its button is never focused to give the focus back to; in Safari on the Mac its buttons by the keyboard (a button focused, Return), the bar's link and the curtain by its own click; by Playwright elsewhere.

What the probe found:

- **The field is a text field, not a search field.** With `type="search"` Escape in the focused field leaves the flyout open in Safari 27 on the Mac (twice, the focus still in the field; claims 14 to 16 failed, the last two because it was still open), and closes it in Chrome and Chromium; with `type="text"` Escape closes it in all three. apple.com's field has no type.
- **The curtain is not the popover's `::backdrop`.** A popover's backdrop has `pointer-events: none` from the browser's own stylesheet (claim 11, in all four), which the page cannot override: a click on it closes the flyout and goes on to whatever is under it (in Chromium 151 with `pointer-events: auto !important` on the backdrop too, a link under it was followed). apple.com's curtain takes the click. So the flyout runs to the window's bottom, its panel at the top and the curtain, a button that hides it, under the panel; the bar, above the flyout, is outside it, and a click there both closes it and reaches the bar's link.
- **The clear button leaves the focus on the page.** Clicked (the browser's own click), the focused button hides itself, and the focus goes to the document (`BODY`) in Safari 27, Chrome and Chromium, where apple.com's script puts it back in the field; Shift+Tab from there is the field again (claim 25), the place it left kept, and Escape still closes the flyout. Clicked by script on the simulator, the focus stays in the field.
- Safari on the Mac draws a 17px scrollbar: the flyout is the page's width (`clientWidth`), not the window's.
- The simulator's picture has the phone's status bar above the page (62px): the page's place in it is found from the bar's color with the flyout closed.

In Firefox (2026-10-22) the pictures are the virtual display's pixels (`../browsers.mjs`): Playwright's picture of a Firefox page leaves the curtain's blur out (the stripes' spread 76.5 under it, its ground drawn, the stripes sharp), so claim 9 failed in the first run, 2026-10-21; read from the display the spread is 0, from 127.5 closed.

Not measured: Firefox on macOS, a real iPhone or iPad; a tap outside the flyout on a phone (no tap clicks on the simulator); the flyout's motion (the component's tests read it in Chromium).

## Running it

```sh
BROWSER=safari node run.mjs      # Safari quit first (../browsers.mjs); and ios, chrome
PW=…/playwright-chromium/index.mjs BROWSER=chromium node run.mjs
BROWSER=firefox node run.mjs     # Playwright's Firefox (playwright-core install firefox), Xvfb and ImageMagick
```

`run.mjs` serves `page.html` on 127.0.0.1 (`?sheet` the same page as a sheet), clicks and presses through `../browsers.mjs`, takes its pictures (`shot()`) and fails if a claim does not hold.
