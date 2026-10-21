# The documentation in each browser

The probe behind the claim that ciderui's documentation (`docs/`, built with ciderui-zola) works outside Chromium, where `test/docs.mjs` holds it. The site is built for `http://127.0.0.1:8737` and served there; `run.mjs` loads every page (37) through `../browsers.mjs` and reads, as `test/docs.mjs` does:

1. every page: nothing scrolls sideways, and every picture loads;
2. every example frame: it loads its stylesheet (`example.css`) and its pictures, is no wider than the page, and is as tall as its content, so it never scrolls;
3. a gallery's and a shelf's Next paddle (the theme's script, in their frames) scroll their row on, the previous paddle enabled (and the gallery's later dot current);
4. an example's link leaves its frame where it is (`about:srcdoc`);
5. below 834px, the global nav's menu opens over its frame, the frame grown to the window's height under the page's sticky local nav and scrolled there, and back to the bar's height closed;
6. below 834px, the local nav's menu opens below its bar, the frame grown to hold it;
7. the global nav's search opens in its frame by its button, its field focused, the frame grown to the window's height under the local nav; in a frame 834px wide or more it hangs from the bar's bottom (anchor positioning), its curtain to the frame's bottom, in a narrower one it is a sheet over the frame; its curtain or its X closes it, the frame the bar's height again;
8. the field's dropdown, with no script: nothing chosen, its label the placeholder (17px, 18px down) and its own text not shown; a choice made (its value set, as a pick sets it), its label at its top (12px, 10px down) and its text shown;
9. the field's textarea, with no script: at rest its label the placeholder and its box's edge 1px; its textarea focused (by the page's script), the box's edge 2px of the accent and the label at its top; then blurred and filled (its value set), the label at its top and the edge 1px;
10. the site's own search page, its query in its address as the global nav's search form sends it (`/search/?q=modal`): its field holding the query, and the rows ciderui-zola's script finds in Zola's index, the modal's page first;
11. the global nav's suggestions, on the home page: its search opened, "mo" put in its field and an input event sent, as typing sends one, then read in the frame its status line first says what they hold: the first suggested link the modal's page, "Mo" marked, the row `#6e6e73` and the mark `#333336` on no ground (ciderui's `:has()` on a `mark`), fading in as apple.com's new results do: at first sight still transparent, its transition `opacity 0.32s` from 60ms (`:not(:placeholder-shown)` in a `:has()`, which a value set by script turns on; the opening's motion is the empty field's), and 600ms on in full.

The frames' elements are clicked by the page's script (`click()`, the activation a tap's click runs): safaridriver's taps on the iOS simulator never click (`../popover/`).

## The result

| browser | window | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Safari 27.0.1, macOS 27.0.1 | 1440, 700 | ok | ok (110 frames) | ok | ok | ok | ok | ok (the flyout in its 978px frame, the sheet in its 412px one) | ok | ok | ok | ok (read 23ms, 23ms after the input) |
| Safari 26.5, iOS 26.5 simulator | 402 | ok | ok (55 frames) | ok | ok | ok | ok | ok (the sheet) | ok | ok | ok | ok (33ms) |
| Chrome 155.0.8059.85, macOS 27.0.1 | 1440, 700 | ok | ok (110 frames) | ok | ok | ok | ok | ok (the flyout, the sheet) | ok | ok | ok | ok (20ms, 12ms) |
| Chromium 151, Linux | 1440, 700 | ok | ok (110 frames) | ok | ok | ok | ok | ok (the flyout, the sheet) | ok | ok | ok | ok (29ms, 19ms) |

Measured 2026-10-21 (`results/`), last with the global nav's suggestions fading in (37 pages; ciderui-zola 97e0857, on ciderui a47c07f; before it, at once, with ciderui-zola 656e621, ciderui's suggestions example among the global nav's frames), Safari not running before the run and the one it launched quit after; before that with the site's search (ciderui-zola 7a5eabc, the search's clear button in the global nav's examples, the docs' own nav a search flyout), the same; before that the site with the field's textarea (the commit after ciderui 17a5efb) with ciderui-zola b6a2447, on the iOS simulator safaridriver chose ("agent-logout"); Safari on the Mac later that day, the site with the modal's page variant too (the commit after 4bbe42c: its frames 108, the two of the product page's modal among them), the owner's Safari quit first with their leave and the one the run launched quit after. On 2026-10-20 four runs there, on the one it chose then ("OwnMusic UITest"), stopped at the first pages: its Safari failed WebDriver's commands (`500 unknown error`), then refused a new session for minutes ("already paired with another WebDriver session"), safaridriver's HTTP service leaving `EXC_GUARD` fault reports; why was not found (no Safari or WebContent crash was logged; the Mac was low on free memory). First measured at ciderui b63888e with ciderui-zola 4bad1c7. Before ciderui-zola 4bad1c7, Safari on the iOS simulator failed 3 once, the shelf's: its row scrolled (400px, 18 scroll events) but its previous paddle stayed disabled, since the theme's script set the paddles in an animation frame and an iframe out of sight runs almost none there (2 in 1.5s, against 89 in sight; of 18 asked for as its row scrolled, none); the script now sets them in the scroll event itself. Not measured: Firefox, a real iPhone.

## Running it

```sh
zola --root docs build --base-url http://127.0.0.1:8737 --output-dir /tmp/ciderui-docs --force
BROWSER=safari SITE=/tmp/ciderui-docs node run.mjs     # Safari quit first (../browsers.mjs); and ios, chrome
PW=…/playwright-chromium/index.mjs BROWSER=chromium SITE=/tmp/ciderui-docs node run.mjs
```

Safari, the iOS simulator and Chrome run on the Mac, with the built site copied there; `run.mjs` fails if a claim does not hold, naming it.
