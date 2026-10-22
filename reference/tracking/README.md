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
| 32px | Display | +0.0040 | +0.0127 / +0.0127 / +0.0127 | +0.0138 / +0.0139 / +0.0138 | −0.0093 |
| 40px | Display | +0.0000 | +0.0093 / +0.0093 / +0.0093 | +0.0104 / +0.0105 / +0.0104 | −0.0099 |
| 48px | Display | −0.0030 | +0.0073 / +0.0073 / +0.0073 | +0.0085 / +0.0085 / +0.0085 | −0.0109 |
| 56px | Display | −0.0050 | +0.0054 / +0.0054 / +0.0054 | +0.0065 / +0.0066 / +0.0065 | −0.0110 |
| 64px | Display | −0.0090 | +0.0034 / +0.0034 / +0.0034 | +0.0046 / +0.0046 / +0.0046 | −0.0130 |
| 80px | Display | −0.0150 | +0.0000 / +0.0000 / +0.0000 | +0.0011 / +0.0012 / +0.0011 | −0.0156 |
| 96px | Display | −0.0150 | +0.0000 / +0.0000 / +0.0000 | +0.0011 / +0.0012 / +0.0011 | −0.0156 |

Measured 2026-10-19 in Safari 27.0.1 on macOS 27.0.1, Google Chrome 155.0.8059.80 on the same Mac, and Safari 26.5 in the iOS 26.5 simulator. Safari and Chrome are at most 0.00005em apart; iOS is macOS's Safari to the fourth decimal.

What else it showed:

- **The name.** `-apple-system` is Safari's alone (in Chrome it falls to another font, 0.065em off) and `BlinkMacSystemFont` is Chrome's alone (0.064em off in Safari); `system-ui` is the one name both take.
- **The control.** On Linux, Chromium 151's `system-ui` is Noto Sans, 0.0075em to 0.0439em off apple.com's cut at every size: the probe fails where the font is not SF (`results/chromium-measure.json`).
- **Safari's font loading.** On apple.com, `document.fonts.ready` never settles in Safari 27 (some of its icon faces stay loading), so the probe loads each face it measures explicitly and fails if one is not `loaded`.
- **Linux's hinting.** Playwright's Chromium on Linux hints every glyph to a whole-pixel advance (100 of Inter's 17px "a" are 1000px; 942.141px with `--font-render-hinting=none`), so no width measured there means anything without that flag, which `run.mjs` passes. With it, Linux's widths are the Mac's: the Inter runs below agree with Chrome's and Safari's on macOS within 0.00006em a character.

Not measured: a real iPhone (the simulator's iOS 26.5, not iOS 27), weights other than 400 and 600, Firefox on macOS (SF in Firefox: Playwright's Firefox does not run on the Mac over SSH; on Linux only Inter is measured in it), and the platforms where `system-ui` is not SF.

## Inter, where the system font is not SF

Where the system font is not SF, ciderui's text is Inter (4.001, the variable font with its `opsz` axis, 14 to 32, and `wght`; `fonts/`), and it takes the same letter-spacing: CSS has no letter-spacing per font of a stack. So the table is SF's, exact on Apple's platforms, and Inter is as near as it falls. Δ here is Inter at ciderui's letter-spacing against apple.com's cut at apple.com's, in em per character; "Inter alone" is the letter-spacing Inter would need to land on apple.com's widths by itself, for comparison:

| size | Δ at 400 | Δ at 600 | Inter alone at 400 | Inter alone at 600 |
| ---: | ---: | ---: | ---: | ---: |
| 12px | +0.0023 | −0.0075 | −0.0128 | −0.0030 |
| 14px | +0.0130 | +0.0032 | −0.0188 | −0.0090 |
| 17px | +0.0208 | +0.0120 | −0.0179 | −0.0091 |
| 19px | +0.0199 | +0.0098 | −0.0328 | −0.0227 |
| 21px | +0.0243 | +0.0149 | −0.0292 | −0.0198 |
| 24px | +0.0202 | +0.0120 | −0.0242 | −0.0160 |
| 28px | +0.0097 | +0.0030 | −0.0170 | −0.0103 |
| 32px | +0.0014 | −0.0038 | −0.0107 | −0.0055 |
| 40px | +0.0048 | −0.0004 | −0.0147 | −0.0095 |
| 48px | +0.0068 | +0.0016 | −0.0177 | −0.0125 |
| 56px | +0.0087 | +0.0035 | −0.0197 | −0.0145 |
| 64px | +0.0107 | +0.0055 | −0.0237 | −0.0185 |
| 80px | +0.0141 | +0.0089 | −0.0297 | −0.0245 |
| 96px | +0.0141 | +0.0089 | −0.0297 | −0.0245 |

Inter at 400 runs wider than apple.com's text by up to 0.0243em a character (21px), 0.0208em at 17px: a 60-character line of body text about 21px wider; 0.0141em at 80px and 96px. Measured 2026-10-19 in Safari 27.0.1 (the table), Chrome 155 on macOS and Chromium 151 on Linux, the three within 0.00006em of each other (`results/*-inter.json`); on 2026-10-22 in Playwright's Firefox 153 on Linux (in its window on a virtual X display: `../browsers.mjs`), within 0.0011em of Safari's, its worst 0.0244em; Inter is added to apple.com's page from its bytes (`inter.js`), `fonts/inter-roman-latin.woff2`.

## The method

On https://www.apple.com/, where apple.com's webfonts are already declared, each face measured is loaded with `document.fonts.load()` and must come back `loaded`. Then one 58-character string, `Sphinx of black quartz judge my vow 0123456789 AVATAR Wave`, is set in a hidden span, `white-space: nowrap`, at each of apple.com's sizes and at 400 and 600, and its width read:

- `measure.js`: in `"SF Pro Text"`, in `"SF Pro Display"` and in `system-ui`, all at `letter-spacing: 0`. It passes only if `system-ui` is within 0.003em a character of apple.com's cut at every size (apple.com's smallest nonzero letter-spacing is 0.002em); it fails on every platform, which is the finding.
- `verify.js`: apple.com's cut at apple.com's letter-spacing against `system-ui`, `-apple-system`, `BlinkMacSystemFont` and the kit's own stack (`tokens.css`'s `--cui-font`) at ciderui's. It passes if the kit's stack is within 0.0015em a character everywhere (0.14px a character at 96px), a tolerance set before the run (the table averages two weights): the stack is what a page draws in, whichever of its names the browser has (`-apple-system` in Safari; Chrome has none, its names falling to `BlinkMacSystemFont`, 0.0651em off apple.com's as the fallback). It passed on `system-ui` before 2026-10-22, a name the stack does not hold, so that a stack off apple.com's widths would have passed it. Run so 2026-10-22 in Chrome 156 on macOS, Safari 27.0.1 on macOS 27.0.1 and Safari 26.5 on the iOS 26.5 simulator: the stack 0.0006em off at worst in each (in Safari and on the simulator `-apple-system`, BlinkMacSystemFont there the fallback's, 0.0638em).

The sizes, the cut apple.com sets at each and its letter-spacing (`analyze.mjs`'s `SCALE`) are its pages' computed type (`../apple.com/`, 2026-10-19): the store's `t-` ladder and the home page's 56px, SF Pro Text up to 17px, SF Pro Display from 19px. Where apple.com tunes a size two ways the store's and the product page's is taken: 32px at +0.004em and 48px at −0.003em (the home page sets +0.002em and −0.002em).

## Running it

`run.mjs` runs one browser in one mode and writes `results/<browser>-<mode>.json`; `derive.mjs` writes `cui-tracking.json` from the Safari and Chrome measurements:

```sh
BROWSER=safari MODE=measure node run.mjs   # safaridriver; Safari's "Allow remote automation" on
BROWSER=ios    MODE=measure node run.mjs   # safaridriver on an iPhone simulator (Xcode)
BROWSER=chrome MODE=measure node run.mjs   # playwright-core, the installed Google Chrome
node derive.mjs
node tables.mjs                            # this README's tables
BROWSER=safari MODE=verify node run.mjs    # and ios, chrome
BROWSER=safari MODE=inter node run.mjs     # and chrome, chromium, firefox (Playwright's); INTER names another font file
PW=…/playwright-chromium/index.mjs BROWSER=chromium MODE=measure node run.mjs   # the Linux control
```

The browsers are `../browsers.mjs`'s. On the Mac it runs from a copy of the repository with `playwright-core` in a `node_modules` above it (`npm i playwright-core`, no browser download: Chrome is the installed one). `safari:deviceType: iPhone` takes whichever iPhone simulator safaridriver finds. The `chromium` runs pass `--font-render-hinting=none`.
