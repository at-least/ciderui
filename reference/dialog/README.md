# A dialog as the store's overlay

The probe behind `cui-modal`, apple.com's store overlay (`../apple.com/`, `overlay:*`) on HTML's `<dialog>`, which needs no script: the page's button opens it with `command="show-modal"` and `commandfor`, a form's button (`method="dialog"`) closes it, Escape closes it. The dialog is the store's layer: fixed over the window, scrolling a panel taller than it, the page dimmed under it by its `::backdrop` and held (`:has()` on the root, `overflow: hidden`, as apple.com sets), fading in (`@starting-style`, 0.1s). The claims, each read by `dialog.js` on `page.html` (the modal's techniques at their smallest) after the browser's own click and keys:

1. the browser has the `command` attribute;
2. the page's button opens the dialog as a modal one (`:modal`), with no script;
3. it is the layer over the window (`position: fixed`, the window's box);
4. the page is dimmed under it (`::backdrop`, `rgba(0,0,0,.48)`);
5. the page is held while it is open, and not after;
6. the layer scrolls the panel taller than the window;
7. a point above the panel, in its 40px of room, hits the layer, not the page;
8. the focus is on its close button, the dialog's first control (as a dialog's opening gives it: in Chromium 151 `autofocus` and `tabindex="-1"` on the dialog leave it there too, where apple.com's script focuses the dialog itself);
9. it fades in: an open frame's opacity under 1, then 1 (sampled each frame from just before the click);
10. its close button closes it;
11. the focus is back on the button that opened it;
12. Escape closes it, the focus back on the button.

And the marketing pages' modal, as `modal.css`'s page variant makes it (`../apple.com/`, `ricmodal:*`): its `::backdrop` dims the page and blurs it (`backdrop-filter: blur(20px)`, in over 0.75s from clear, `@starting-style`), the panel centered (auto margins in the flex layer), its close button's form sticky, its top the layer's room taken back:

13. its button opens it, its `::backdrop`'s filter `blur(20px)`;
14. its panel centered: as far from the window's top as from its foot;
15. in a picture of the window, the page under it blurred and dimmed: along a row of the page's 4px stripes, the brightness's spread under a fifth of its spread closed, and its mean 30 lower;
16. its scrim comes in: a frame whose blur is between 0 and 20px (sampled each frame from just before the click);
17. made taller than the window and scrolled 400px, its close button stays 20px from the window's top;
18. its close button closes it.

## The result

| browser | 1–10 | 11 | 12 | 13–18 |
| --- | --- | --- | --- | --- |
| Safari 27.0.1, macOS 27.0.1, by its keyboard | ok | ok | ok | not run |
| Safari 26.5, iOS 26.5 simulator | ok | not run | not run | ok |
| Chrome 155.0.8059.85, macOS 27.0.1 | ok | ok | ok | ok |
| Chromium 151, Linux | ok | ok | ok | ok |

Measured 2026-10-19, claims 13 to 18 and the rows of the iOS simulator, Chrome and Chromium again on 2026-10-21 (`results/`), at 1440 by 900 (the simulator's phone at its own 402 by 714); Safari on the Mac's row is the run of 2026-10-19, as the owner's Safari was open on 2026-10-21 (a run takes no session in a Safari it did not launch, and it is not quit without the owner's leave). The pictures are read as `../search/` reads its own: the page's place in the picture found from a band's color with the modal closed (62px down on the simulator, under its status bar). The first version of claim 15 read the row with a panel taller than the window over it, its white raising the row's spread (88.9 of 127.5): the panel is now opened short, then made tall for claim 17. Clicked as `../popover/` clicks: by the page's script on the iOS simulator (safaridriver's taps never click there), so its opener is never focused and there is no focus to give back, and no key is pressed; by the keyboard in Safari on the Mac (the button focused, Return); by Playwright elsewhere. The first version read claim 7 at the window's left edge, which below 816px is inside the panel (it is the window's width there): it read the panel on the iOS simulator. Not measured: Firefox, a real iPhone; a click on the layer, which closes apple.com's overlay by its script and nothing here.

## Running it

```sh
BROWSER=safari node run.mjs      # Safari quit first (../browsers.mjs); and ios, chrome
PW=…/playwright-chromium/index.mjs BROWSER=chromium node run.mjs
```

`run.mjs` serves `page.html` on 127.0.0.1, clicks and presses through `../browsers.mjs` and fails if a claim does not hold.
