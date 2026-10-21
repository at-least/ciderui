# What of the nav can be selected

The probe behind the global nav's `user-select: none`, apple.com's on `#globalnav` (`../apple.com/`, `nav-select:*`): nothing in the nav can be selected, and its search's field, inside it, still takes text. `page.html` is the technique at its smallest: a bar with `-webkit-user-select: none; user-select: none` (what `globalnav.css` compiles to), its list a closed popover shown as its row, a link, a button opening a link's flyout (a heading in it) and one opening the search (a heading and a text field, `autofocus`), the page's own text under them; `?plain` is the same page with the bar selectable, so that each claim of nothing selected has its control. The claims, each read by `run.mjs` from what the page reports (`cuiState()`):

1. the bar computes `user-select: none` (`-webkit-user-select` where the browser has no other);
2. where the bar can be selected, Select All on the page (the platform's command key and A) takes the bar's text with the page's;
3. where it cannot, Select All takes the page's text alone;
4. where the bar can be selected, a double-click on a link's flyout's heading selects its text;
5. where it cannot, the same double-click selects nothing;
6. the search button opens its flyout, the focus on its field;
7. where the bar can be selected, a double-click on the search's heading selects its text;
8. where it cannot, the same double-click selects nothing;
9. the field clicked, keys typed go into it ("mac");
10. Select All in the field selects its text whole.

On the iOS simulator, by XCUITest (`../xcuitest.mjs`: its taps, its keyboard and a long press), each page on a simulator of its own:

11. the search button, tapped, opens its flyout, the focus on its field;
12. XCUITest's keyboard types into the field ("mac");
13. the link's flyout's button, tapped, closes the search and opens its flyout, the field left;
14. where the bar can be selected, a long press on the flyout's heading selects its text;
15. where it cannot, the same long press selects nothing.

## The result

| browser | 1–10 | 11–15 |
| --- | --- | --- |
| Safari 27.0.1, macOS 27.0.1, its buttons by its keyboard | ok | not run |
| Safari 26.5, iOS 26.5 simulator (iPhone 17), XCUITest | not run | ok |
| Chrome 156.0.8078.48, macOS 27.0.1 | ok | not run |
| Chromium 151, Linux | ok | not run |
| Firefox 153.0 (Playwright's), Linux, in its window on a virtual X display | ok | not run |

Measured 2026-10-22 (`results/`). In Safari on the Mac the buttons are pressed as `../search/` presses them, by the keyboard (a button focused, Return), the double-clicks and keys WebDriver's (`../browsers.mjs`); Playwright's elsewhere.

What the probe found:

- **Safari has no `user-select`.** In Safari 27.0.1 and in Safari 26.5 on the iOS simulator `CSS.supports('user-select', 'none')` is false and `-webkit-user-select` is the property, so an unprefixed declaration alone would be dropped. Tailwind's CLI (4.3.3) writes the prefixed one beside it when it minifies (`--minify`, as ciderui's tests and ciderui-zola's build run it), and `user-select: none` alone without: so `globalnav.css` declares both, as apple.com's stylesheet does.
- **The field needs no exception.** apple.com's field computes `none` in Chromium, which carries the value down, and takes keys (`nav-select:search.typed`); here, inside the bar, it takes keys and its text is selected whole by Select All in all four desktop browsers, and takes XCUITest's keys on iOS. Safari computes `-webkit-user-select: text` for the field inside the `none` bar (WebKit's own for what can be edited) and `text` for the page; Firefox computes `none` on the bar alone, `auto` below it.
- **A double-click selects no link's text in Chromium.** In Chromium 151 a double-click on a link with an `href` selects nothing, the bar selectable or not (an `<a>` without one, and a `<span>`, select their word), so the double-clicks are on the nav's text that is no link, a flyout's heading and the search's; Select All shows the rule on the bar's link.
- **On iOS a long press selected nothing while the field was focused.** With the search's field focused, a long press on the page's own text selected nothing (once, the press closing the search), and one on the search's heading selected nothing on either page (once each); a tap on the heading left the field focused. So the long press is on a link's flyout's heading, the search closed by that flyout's button, where on `?plain` it selects.

Not measured: Firefox on macOS, a real iPhone or iPad; a selection dragged from the page's text across the nav.

## Running it

```sh
BROWSER=safari node run.mjs      # Safari quit first (../browsers.mjs); and chrome
BROWSER=ios node run.mjs         # on the Mac: Xcode, XcodeGen (Homebrew)
PW=…/playwright-chromium/index.mjs BROWSER=chromium node run.mjs
BROWSER=firefox node run.mjs     # Playwright's Firefox: playwright-core install firefox
```

`run.mjs` serves `page.html` on 127.0.0.1 and fails if a claim does not hold, printing what each step left. For `ios` it creates an iPhone 17 simulator on the newest iOS runtime that has one, once for `?plain` and once for the page, runs the UI test (`../xcuitest/`) with the steps `Search`, `type:mac`, `More`, `press:Explore`, and shuts each down and deletes it after: the simulators already booted are left alone.
