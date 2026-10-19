# The system font's tracking

The probe behind ciderui's letter-spacing. ciderui cannot ship SF Pro, so its text is `system-ui`, which is SF on Apple's platforms; apple.com sets its own SF Pro Text and SF Pro Display webfonts with a letter-spacing per size. The question: is `system-ui` the same font at the same widths as apple.com's cut, so apple.com's letter-spacing can be copied, or does it carry tracking of its own, which apple.com's letter-spacing would then double?

## The answer

`system-ui` carries Apple's size-specific tracking of its own, the same in Safari and Chrome on macOS and in Safari on iOS. Copying apple.com's letter-spacing onto it would track 17px text twice (−0.022em on top of the font's −0.025em). ciderui's letter-spacing is apple.com's minus that tracking, size by size, and lands `system-ui` on apple.com's widths within 0.0006em a character in all three.

Δ is how much wider `system-ui` is than apple.com's cut at that size, both at `letter-spacing: 0`, in em per character; ciderui's letter-spacing is apple.com's minus Δ, averaged over the two weights and the two macOS browsers (`cui-tracking.json`):

| size | apple.com cut | apple.com letter-spacing | Δ at 400 (Safari / Chrome / iOS) | Δ at 600 (Safari / Chrome / iOS) | ciderui letter-spacing |
| ---: | --- | ---: | ---: | ---: | ---: |
| 12px | Text | −0.0100 | +0.0000 / +0.0000 / +0.0000 | +0.0011 / +0.0011 / +0.0011 | −0.0105 |
| 14px | Text | −0.0160 | −0.0107 / −0.0107 / −0.0107 | −0.0097 / −0.0097 / −0.0097 | −0.0058 |
| 17px | Text | −0.0220 | −0.0254 / −0.0254 / −0.0254 | −0.0243 / −0.0243 / −0.0243 | +0.0029 |
| 19px | Display | +0.0120 | +0.0253 / +0.0253 / +0.0253 | +0.0244 / +0.0244 / +0.0244 | −0.0129 |
| 21px | Display | +0.0110 | +0.0160 / +0.0160 / +0.0160 | +0.0158 / +0.0158 / +0.0158 | −0.0049 |
| 24px | Display | +0.0090 | +0.0127 / +0.0126 / +0.0127 | +0.0134 / +0.0134 / +0.0134 | −0.0040 |
| 28px | Display | +0.0070 | +0.0137 / +0.0137 / +0.0137 | +0.0148 / +0.0149 / +0.0148 | −0.0073 |
| 32px | Display | +0.0020 | +0.0127 / +0.0127 / +0.0127 | +0.0138 / +0.0139 / +0.0138 | −0.0113 |
| 40px | Display | +0.0000 | +0.0093 / +0.0093 / +0.0093 | +0.0104 / +0.0105 / +0.0104 | −0.0099 |
| 48px | Display | −0.0020 | +0.0073 / +0.0073 / +0.0073 | +0.0085 / +0.0085 / +0.0085 | −0.0099 |
| 56px | Display | −0.0050 | +0.0054 / +0.0054 / +0.0054 | +0.0065 / +0.0066 / +0.0065 | −0.0110 |

Measured 2026-10-19 in Safari 27.0.1 on macOS 27.0.1, Google Chrome 155.0.8059.80 on the same Mac, and Safari 26.5 in the iOS 26.5 simulator. Safari and Chrome are at most 0.00005em apart; iOS is macOS's Safari to the fourth decimal.

What else it showed:

- **The name.** `-apple-system` is Safari's alone (in Chrome it falls to another font, 0.065em off) and `BlinkMacSystemFont` is Chrome's alone (0.064em off in Safari); `system-ui` is the one name both take.
- **The control.** On Linux, Chromium 151's `system-ui` is Noto Sans, 0.010em to 0.048em off apple.com's cut at every size: the probe fails where the font is not SF (`results/chromium-measure.json`).
- **Safari's font loading.** On apple.com, `document.fonts.ready` never settles in Safari 27 (some of its icon faces stay loading), so the probe loads each face it measures explicitly and fails if one is not `loaded`.

Not measured: a real iPhone (the simulator's iOS 26.5, not iOS 27), weights other than 400 and 600, Firefox, and the platforms where `system-ui` is not SF.

## The method

On https://www.apple.com/, where apple.com's webfonts are already declared, each face measured is loaded with `document.fonts.load()` and must come back `loaded`. Then one 58-character string, `Sphinx of black quartz judge my vow 0123456789 AVATAR Wave`, is set in a hidden span, `white-space: nowrap`, at each of apple.com's sizes and at 400 and 600, and its width read:

- `measure.js`: in `"SF Pro Text"`, in `"SF Pro Display"` and in `system-ui`, all at `letter-spacing: 0`. It passes only if `system-ui` is within 0.003em a character of apple.com's cut at every size (apple.com's smallest nonzero letter-spacing is 0.002em); it fails on every platform, which is the finding.
- `verify.js`: apple.com's cut at apple.com's letter-spacing against `system-ui`, `-apple-system` and `BlinkMacSystemFont` at ciderui's. It passes if `system-ui` is within 0.0015em a character everywhere (0.08px a character at 56px), a tolerance set before the run: the table averages two weights.

The sizes, the cut apple.com sets at each and its letter-spacing (`analyze.mjs`'s `SCALE`) are its stylesheets' typography rules, read 2026-10-19: SF Pro Text up to 17px, SF Pro Display from 19px.

## Running it

`run.mjs` runs one browser in one mode and writes `results/<browser>-<mode>.json`; `derive.mjs` writes `cui-tracking.json` from the Safari and Chrome measurements:

```sh
BROWSER=safari MODE=measure node run.mjs   # safaridriver; Safari's "Allow remote automation" on
BROWSER=ios    MODE=measure node run.mjs   # safaridriver on an iPhone simulator (Xcode)
BROWSER=chrome MODE=measure node run.mjs   # playwright-core, the installed Google Chrome
node derive.mjs
BROWSER=safari MODE=verify node run.mjs    # and ios, chrome
PW=…/playwright-chromium/index.mjs BROWSER=chromium MODE=measure node run.mjs   # the Linux control
```

On the Mac it runs from a copy of this directory with `playwright-core` installed beside it (`npm i playwright-core`, no browser download: Chrome is the installed one). `safari:deviceType: iPhone` takes whichever iPhone simulator safaridriver finds.
