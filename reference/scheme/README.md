# A section's color scheme

The probe behind ciderui's dark sections. apple.com's dark is a section's, not the page's (`../apple.com/`: no `prefers-color-scheme` in the stylesheets its pages load, `.theme-dark` sections on light pages); ciderui's colors are `light-dark()` pairs, so a dark section is `color-scheme: dark` on the section if `light-dark()` follows the color scheme an ancestor sets rather than the root's. The claim: a pair declared once on `:root` resolves dark inside a section that sets `color-scheme: dark`, and light again inside a light block within it.

## The result

| browser | page | dark section | in it | light inside dark |
| --- | --- | --- | --- | --- |
| Safari 27.0.1, macOS 27.0.1 | light | dark | dark | light |
| Chrome 155.0.8059.80, macOS 27.0.1 | light | dark | dark | light |
| Safari 26.5, iOS 26.5 simulator | light | dark | dark | light |
| Chromium 151, Linux | light | dark | dark | light |

Each element's color and ground are exactly the pair's side (`rgb(29, 29, 31)` on `rgb(255, 255, 255)`, `rgb(245, 245, 247)` on `rgb(0, 0, 0)`); measured 2026-10-19 (`results/`). The pairs are unregistered custom properties, whose computed value is the `light-dark()` itself, resolved where it is used. Not measured: Firefox.

## Running it

```sh
BROWSER=safari node run.mjs      # and ios, chrome
PW=…/playwright-chromium/index.mjs BROWSER=chromium node run.mjs
```

`run.mjs` serves `page.html` on 127.0.0.1, reads each element's color and ground (`scheme.js`) and fails if one is not its side of the pair. The browsers are `../browsers.mjs`'s.
