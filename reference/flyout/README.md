# The global nav's flyouts, with no script

The probe behind the global nav's flyouts, apple.com's menus under its links (`../apple.com/`, `flyout:*`), which apple.com's script opens: the pointer resting 120ms on a link, and on a touch screen a tap on the link; the link's button for the keyboard; to 833px a row of the menu opening its flyout over it. `page.html` holds the techniques at their smallest (a sticky bar under a 70px banner, its list a closed popover shown as its row, as the nav's is; two links with a flyout each and one without), and nothing in it is a script but what reads it:

- From 834px with a pointer (`(hover: hover)`), a link's flyout, a popover after the link, is shown by the page closed (the browser's own style hides a closed popover; the page's wins), from the bar's top, behind the bar's links (theirs the higher `z-index`), hidden (`visibility`) until the pointer is on its item: `:hover`, with a transition's delay of 120ms, so that a pointer passing over a link shows nothing. Its panel is revealed down from the bar (`clip-path`), and under it the list's curtain (its `::after`) is fixed over the window, blurring the page, taking no pointer, so that the pointer on the curtain is off the item; the bar's own blur is off while a flyout shows (a `backdrop-filter` on the bar would hold the curtain to the bar, as its containing block, and keep it from the page, as its backdrop root), and back after the curtain has gone. The pointer moved on along the bar keeps the flyout it left until the next is shown (`:hover` on the list, a longer delay), and the pointer off the bar and the flyout hides it.
- The link's button (`popovertarget`) opens the flyout as a popover, from the bar's bottom (anchor positioning), its curtain a button that closes it; on a screen with no hover the button lies over the link, so a tap opens the flyout and does not follow the link, as apple.com's script does there.
- To 833px the list is the menu, a sheet its button opens; a row's button, over the row, opens its flyout over the open menu (a popover inside a popover: the menu stays open under it), a sheet with a back button, which closes the flyout, and a close button, which closes the menu and the flyout with it.

The claims, each read in the page (`page.html`'s `probe`: the flyouts' states and places, which show each frame and how far down the lowest reaches, when the pointer entered each item) and in pictures of the window:

1. the browser has anchor positioning and a pointer that hovers;
2. closed, no flyout shows, the curtain is hidden, the bar blurred;
3. the pointer resting on Store's link shows its flyout, no popover opened, 100 to 250ms after it reached the link;
4. it hangs from the bar's top, across the window;
5. the bar's links are over it (what the pointer finds at Mac's link is Mac's link);
6. the curtain shows, fixed, and the bar's blur is off;
7. in the picture the curtain blurs the page under it (along a row of the page's 4px stripes, the brightness's spread under a fifth of its spread closed), and the bar shows the flyout's ground (`#fafafc`);
8. the pointer moved on to Mac's link shows Mac's flyout and not Store's;
9. and in every frame between, a flyout shows to its full height (the one left kept until the next is shown);
10. the pointer off the bar to the page hides it, the curtain hidden, the bar blurred again;
11. swept along the bar past its links in 90ms, no flyout shows in any frame (each link's item entered);
12. Mac's button opens its flyout as a popover, from the bar's bottom, its curtain to the window's bottom;
13. Tab takes the focus to its first link;
14. Escape closes it, the focus back on its button;
15. a click on its curtain closes it and reaches nothing under it;
16. at a phone's width the menu's button opens the menu;
17. Mac's row's button opens Mac's flyout over the open menu, across the window;
18. its back button closes the flyout, the menu still open;
19. its close button closes both.

On the iPad (13-inch, 1032px wide), by XCUITest's taps:

1. the screen has no hover: `(hover: hover)` does not match;
2. a tap on Store's link (its button over it) opens its flyout as a popover, from the bar's bottom;
3. the link is not followed;
4. a tap on its curtain closes it and reaches nothing under it.

On the iPhone, by XCUITest's taps, claims 16 to 19.

## The result

| browser | 1–19 |
| --- | --- |
| Safari 27.0.1, macOS 27.0.1, its buttons by its keyboard | ok |
| Chrome 156.0.8078.48, macOS 27.0.1 | ok |
| Chromium 151, Linux | ok |

| browser | iPad 1–4 | 16–19 |
| --- | --- | --- |
| Safari 26.5, iOS 26.5 simulator (iPad Pro 13-inch (M5)), XCUITest's taps | ok | |
| Safari 26.5, iOS 26.5 simulator (iPhone 17), XCUITest's taps | | ok |

Measured 2026-10-21 (`results/`), at 1440 by 900 (Safari's window, its frame included: its page 812px tall) and 700 by 900 for a phone's layout (Safari's window goes no narrower), the simulators at their own sizes. The flyout showed 136ms after the pointer reached its link in Safari, 166ms in Chrome, 133ms in Chromium (a frame's time read each frame, from the item's `pointerover`). Under the curtain the stripes' spread fell from 127.3 to 8.4 in Safari, from 127.5 to 0 in Chrome and 1.1 in Chromium. In Chromium, with the dwell's delay taken off claims 3 and 11 fail (shown at 0ms, and in the sweep), with the longer delay on the list's `:hover` taken off claim 9 (the lowest reach fell to 237px of 328px as the one left shrank before the next had grown), with the bar's blur kept on claims 6 and 7, and with the bar's links under the flyouts claims 5 and 8. The iPad's Safari matches `(pointer: coarse)` and no hover, so the page's hover rules never apply there; it put no tip over the page (`tip: "none"`), the iPhone's did and the test closed it (`tip: "closed"`, `../tap/`).

What the probe found:

- **Safari's pointer moves.** safaridriver's pointer move of a duration is one jump at its end (one `pointermove`), so a sweep passed over no link; `../browsers.mjs`'s `move` gives it a move a frame. A move outside the page (below its 812px) does nothing at all, not even an error.
- **The bar's blur.** Kept on while a flyout shows, in Chromium, the curtain does not reach the page at all: the stripes' spread under it stayed 127.5, its tint not drawn there either, the curtain held to the bar (`backdrop-filter` makes the bar the containing block of its fixed descendants; it would be the backdrop root of the curtain's own blur too). So the bar drops it while a flyout shows, and takes it back 0.45s after one stops showing, the curtain gone by then.

Not measured: Firefox, a real iPhone or iPad; an iPad with a trackpad or mouse (its pointer hovers: the page's hover rules would apply); the flyouts' motion (the component's tests read it in Chromium).

## Running it

```sh
BROWSER=safari node run.mjs      # Safari quit first (../browsers.mjs); and chrome
BROWSER=ipad node run.mjs        # on the Mac: Xcode, XcodeGen (Homebrew); and iphone
PW=…/playwright-chromium/index.mjs BROWSER=chromium node run.mjs
```

`run.mjs` serves `page.html` on 127.0.0.1, moves the pointer, clicks and presses through `../browsers.mjs` (Safari's buttons by its keyboard: focused, Return, Option-Tab for Tab), takes its pictures (`shot()`), or taps through XCUITest on a simulator of its own (`../xcuitest.mjs`: an iPad Pro 13-inch or an iPhone 17, created and deleted, the booted ones left alone), and fails if a claim does not hold.
