# A dialog as the store's overlay

The probe behind a dialog with no script (`cui-console-dialog` uses it), apple.com's store overlay (`../apple.com/`, `overlay:*`) on HTML's `<dialog>`, which needs no script: the page's button opens it with `command="show-modal"` and `commandfor`, a form's button (`method="dialog"`) closes it, Escape closes it. The dialog is the store's layer: fixed over the window, scrolling a panel taller than it, the page dimmed under it by its `::backdrop` and held (`:has()` on the root, `overflow: hidden`, as apple.com sets), fading in (`@starting-style`, 0.1s). The claims, each read by `dialog.js` on `page.html` (the modal's techniques at their smallest) after the browser's own click and keys:

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

And a button in a form:

13. a button in a form, its type `button`, opens the dialog, the form not sent;
14. the control: one with no type, the form's submit button, does nothing at all with its `commandfor`, the dialog not opened and the form not sent (a submit button with no command sends it).

## The result

| browser | 1–10 | 11 | 12 | 13–14 |
| --- | --- | --- | --- | --- |
| Safari 27.0.1, macOS 27.0.1, by its keyboard | ok | ok | ok | ok |
| Safari 26.5, iOS 26.5 simulator | ok | not run | not run | ok |
| Chrome 156.0.8078.48, macOS 27.0.1 | ok | ok | ok | ok |
| Chromium 151, Linux | ok | ok | ok | ok |
| Firefox 153.0 (Playwright's), Linux, in its window on a virtual X display | ok | ok | ok | ok |

Measured 2026-10-19, and again in every row 2026-10-22 with claims 13 and 14 (`results/`; Safari not running before its run and in front of the Mac throughout, `results/front-safari.log`, the Mac idle; the simulator safaridriver booted, the owner's "agent-logout", shut down after), at 1440 by 900 (the simulator's phone at its own 402 by 714); the Safari each run launched quit after it. Clicked as `../popover/` clicks: by the page's script on the iOS simulator (safaridriver's taps never click there), so its opener is never focused and there is no focus to give back, and no key is pressed; by the keyboard in Safari on the Mac (the button focused, Return); by Playwright elsewhere. Claim 7's point is above the panel's middle, not at the window's left edge: below 816px the panel is the window's width, and a point at its edge hits the panel. Not measured: Firefox on macOS, a real iPhone; a click on the layer, which closes apple.com's overlay by its script and nothing here.

## Running it

```sh
BROWSER=safari node run.mjs      # Safari quit first (../browsers.mjs); and ios, chrome
PW=…/playwright-chromium/index.mjs BROWSER=chromium node run.mjs
BROWSER=firefox node run.mjs     # Playwright's Firefox (playwright-core install firefox), Xvfb and ImageMagick
```

`run.mjs` serves `page.html` on 127.0.0.1, clicks and presses through `../browsers.mjs` and fails if a claim does not hold.
