# What a tap leaves

The probe behind decision 5 (`../../DESIGN.md`): ciderui's hover is `@media (hover: hover) { &:hover }`, so a tapped element does not keep its hover, where apple.com's hover is a rule of its own and stays (`../apple.com/`, `touch-*:*` forced, `tap-*:*` tapped). The touch cases of the components' tests force `:hover` on a touch screen; this is what a tap itself leaves. Each element is tapped as a touch screen is, never clicked by a script: in Chromium and Chrome by the DevTools protocol's touch events, on a touch screen emulated from the load, which Chromium's gesture recognizer makes a tap of; on the iOS simulator by XCUITest's taps (`../xcuitest.mjs`), touches of the simulator's screen through its HID, as a finger's, where WebDriver's taps never click (`../popover/`). `page.html` holds a button whose hover is a rule of its own (apple.com's), one whose hover is behind `(hover: hover)` (ciderui's), somewhere else to tap, and the nav's menu, a popover its button opens from beside it with a close button inside; 800ms after each tap its script sends what the tap left. The claims:

1. a tap is a touch: its `pointerdown` and `pointerup` `touch` and trusted, then `mouseover`, `mousedown`, `mouseup` and a click on the button;
2. the screen has no hover: `(hover: hover)` does not match, `(pointer: coarse)` does;
3. the tapped button is left hovered (`:hover` 800ms on), and a hover rule of its own shows there (apple.com's way: `#0076df`);
4. a hover behind `(hover: hover)` does not: the tapped button hovered, at its rest `#0071e3` (ciderui's way);
5. no focus ring after a tap: `:focus-visible` on neither button;
6. the menu's button opens its popover by a tap;
7. its close button, tapped, closes it.

## The result

| browser | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Safari 26.5, iOS 26.5 simulator (iPhone 17), XCUITest's taps | ok | ok | ok | ok | ok | ok | ok |
| Chrome 156.0.8078.48, macOS 27.0.1, the DevTools protocol's touch events | ok | ok | ok | ok | ok | ok | ok |
| Chromium 151, Linux, the same | ok | ok | ok | ok | ok | ok | ok |
| Firefox 153.0 (Playwright's), Linux, Playwright's touchscreen taps on a page with a touch screen (`hasTouch`) | ok | ok | ok | ok | ok | ok | ok |

Measured 2026-10-21, Chromium again that day and Chrome (156, its update) 2026-10-22, after the run's wait for a report was shared with Firefox's taps, and Firefox that day (in its window on a virtual X display: `../browsers.mjs`) (`results/`). In Firefox, which Playwright gives no DevTools protocol (`newCDPSession` is Chromium's only), the taps are Playwright's touchscreen's on a page made with a touch screen from its load (`hasTouch`: Gecko's own touch emulation, which turns `(hover: hover)` off and `(pointer: coarse)` on): each leaves the touch's trusted pointer events, `touchstart` and `touchend`, then mouse events and a click, as in Chromium; its `navigator.maxTouchPoints` stays 0 there (Chromium's 1), which `page.html` does not read. Safari on a new simulator puts a tip over the page ("View Bookmarks, Share Menu, and Open Tabs"), a popover whose dismiss region covers the window and takes the next touch: the page got that tap's touch and pointer events, its button even hovered, and no mouse event or click (claim 1 failed so in four runs of five; in the three whose accessibility tree was printed before the first tap, the tip was there). So the test closes the tip first (`tip: "closed"` in the result) and fails if a popover of Safari's is up when a tap lands; four runs alike since. Recorded beside the claims: a tap on the gated button moved the hover off the plain one, and a tap somewhere else (on no control) moved it off the gated one, in each; Chromium and Chrome focus a tapped button (`:focus`, never `:focus-visible`), Safari on iOS does not; Safari's click after a tap is a `pointerType` `mouse`, Chromium's `touch`. The plain button is the gated one's control: the same tap shows a hover rule that is not behind the query, so the gated one's rest is the query's doing (in Chromium, with its query taken off, claim 4 fails, as claim 3 does with one put on the plain one, and claim 7 with its close button showing the menu; with no `popovertarget` on the menu's button the run stops at its close button, which has no box to tap). Not measured: a phone's own screen (the simulator's HID stands in for it), a touch screen's Firefox (Playwright's emulates one); Safari on the Mac has no touch screen.

## Running it

```sh
BROWSER=ios node run.mjs          # on the Mac: Xcode, XcodeGen (Homebrew)
BROWSER=chrome node run.mjs
PW=…/playwright-chromium/index.mjs BROWSER=chromium node run.mjs
BROWSER=firefox node run.mjs      # Playwright's Firefox: playwright-core install firefox
```

`run.mjs` serves `page.html` on 127.0.0.1, taps through the DevTools protocol (`../browsers.mjs`'s Playwright page), in Firefox through Playwright's touchscreen, or, for `ios`, creates an iPhone 17 simulator on the newest iOS runtime, boots it, generates the UI test's project from `../xcuitest/project.yml` into a temporary directory and runs it (`xcodebuild test`, the page's URL and the labels to tap in `TEST_RUNNER_` variables: `../xcuitest.mjs`, which the flyouts' probe shares), then shuts the simulator down and deletes it: the simulators already booted are left alone. It fails if a claim does not hold.
