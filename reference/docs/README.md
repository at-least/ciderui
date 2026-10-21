# The documentation in each browser

The probe behind the claim that ciderui's documentation (`docs/`, built with ciderui-zola) works outside Chromium, where `test/docs.mjs` holds it. The site is built for `http://127.0.0.1:8737` and served there; `run.mjs` loads every page (36) through `../browsers.mjs` and reads, as `test/docs.mjs` does:

1. every page: nothing scrolls sideways, and every picture loads;
2. every example frame: it loads its stylesheet (`example.css`) and its pictures, is no wider than the page, and is as tall as its content, so it never scrolls;
3. a gallery's and a shelf's Next paddle (the theme's script, in their frames) scroll their row on, the previous paddle enabled (and the gallery's later dot current);
4. an example's link leaves its frame where it is (`about:srcdoc`);
5. below 834px, the global nav's menu opens over its frame, the frame grown to the window's height under the page's sticky local nav and scrolled there, and back to the bar's height closed;
6. below 834px, the local nav's menu opens below its bar, the frame grown to hold it;
7. the global nav's search opens in its frame by its button, its field focused, the frame grown to the window's height under the local nav; in a frame 834px wide or more it hangs from the bar's bottom (anchor positioning), its curtain to the frame's bottom, in a narrower one it is a sheet over the frame; its curtain or its X closes it, the frame the bar's height again;
8. the field's dropdown, with no script: nothing chosen, its label the placeholder (17px, 18px down) and its own text not shown; a choice made (its value set, as a pick sets it), its label at its top (12px, 10px down) and its text shown.

The frames' elements are clicked by the page's script (`click()`, the activation a tap's click runs): safaridriver's taps on the iOS simulator never click (`../popover/`).

## The result

| browser | window | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Safari 27.0.1, macOS 27.0.1 | 1440, 700 | ok | ok (104 frames) | ok | ok | ok | ok | ok (the flyout in its 978px frame, the sheet in its 412px one) | ok |
| Safari 26.5, iOS 26.5 simulator | 402 | ok | ok (52 frames) | ok | ok | ok | ok | ok (the sheet) | ok |
| Chrome 155.0.8059.80, macOS 27.0.1 | 1440, 700 | ok | ok (104 frames) | ok | ok | ok | ok | ok (the flyout, the sheet) | ok |
| Chromium 151, Linux | 1440, 700 | ok | ok (104 frames) | ok | ok | ok | ok | ok (the flyout, the sheet) | ok |

Measured 2026-10-20 (`results/`), the site with the field's states and dropdown (the commit after ciderui 31b8622) with ciderui-zola fd5c713; the iOS simulator's row on 2026-10-21, the site the same, on the simulator safaridriver chose that day ("agent-logout"). On 2026-10-20 four runs there, on the one it chose then ("OwnMusic UITest"), stopped at the first pages: its Safari failed WebDriver's commands (`500 unknown error`), then refused a new session for minutes ("already paired with another WebDriver session"), safaridriver's HTTP service leaving `EXC_GUARD` fault reports; why was not found (no Safari or WebContent crash was logged; the Mac was low on free memory). First measured at ciderui b63888e with ciderui-zola 4bad1c7. Before ciderui-zola 4bad1c7, Safari on the iOS simulator failed 3 once, the shelf's: its row scrolled (400px, 18 scroll events) but its previous paddle stayed disabled, since the theme's script set the paddles in an animation frame and an iframe out of sight runs almost none there (2 in 1.5s, against 89 in sight; of 18 asked for as its row scrolled, none); the script now sets them in the scroll event itself. Not measured: Firefox, a real iPhone.

## Running it

```sh
zola --root docs build --base-url http://127.0.0.1:8737 --output-dir /tmp/ciderui-docs --force
BROWSER=safari SITE=/tmp/ciderui-docs node run.mjs     # Safari quit first (../browsers.mjs); and ios, chrome
PW=…/playwright-chromium/index.mjs BROWSER=chromium SITE=/tmp/ciderui-docs node run.mjs
```

Safari, the iOS simulator and Chrome run on the Mac, with the built site copied there; `run.mjs` fails if a claim does not hold, naming it.
