# Forced colors

Does each of the kit's parts that draws a state or a color in a color of its own still show it in forced colors, a high-contrast theme, where the browser draws a page in the theme's few colors? `page.mjs` writes `page.html`, the kit's CSS compiled as a site compiles it, with the swatches (their colors an element's style, and again `<img>`), an option and a checkbox each unchecked and checked, the gallery, and the empty select four times (as it is; its empty choice's text ""; its label hidden; both). `run.mjs` reads each part's picture (Playwright's of the element, decoded in another page, unforced) in Playwright's Chromium with forced colors emulated (black on white, and with a dark color scheme white on black) and not, and in Playwright's Firefox with its own override of a page's colors on (Settings, Website contrast, Override colors: the pref `browser.display.document_color_use` 2) and not.

| | Chromium 151 | Chromium, forced | Chromium, forced, dark | Firefox 153 | Firefox, override |
| --- | --- | --- | --- | --- | --- |
| forced colors matched, the option and the checkbox drawn checked unlike unchecked (the controls) | ok | ok | ok | ok | ok |
| the checked swatch's ring drawn unlike each unchecked one | ok | ok | ok | ok | ok |
| the swatches' own colors drawn, three apart | ok | ok | ok | ok | ok |
| the gallery's dots drawn, the current one's pill more ink than a dot | ok | ok | ok | ok | ok |
| the empty select's own text not drawn where its label lies | ok | ok (drawn below it) | ok (drawn below it) | ok | ok (not drawn) |

Run 2026-10-22 (`results/`). Before the kit's rules for it, apple.com's look alone, every forced mode failed: each swatch's ring drawn alike (its edge the theme's text color, `rgb(0, 0, 0)`, checked or not, with `<img>` colors too) and every swatch's color the window's ground; the dots drew nothing (their ground forced to the window's, 0 pixels of ink); and in Chromium the empty select's "Select", its transparent color forced, was drawn under its label, the label's box rows 18 to 38 over it (300 pixels of ink, rows 31 to 42), where Firefox keeps the select's transparent color and draws none. The kit's rules (in each stylesheet's header):

- the swatch's ring and its color its own, as a picture's are (`forced-color-adjust: none`): the checked ring the accent and the colors shown;
- a dot a 4px edge in the theme's link color (`border: 4px solid`, in forced colors alone), the current one's pill longer;
- an empty select's label at its top in forced colors, as a filled one's: in Chromium "Select" drawn below it, in Firefox nothing drawn there.

apple.com's parts are the kit's here: not measured in forced colors on apple.com. Not in Safari: WebKit has no forced colors (macOS's Increase Contrast draws no page in a theme's colors).

```sh
node page.mjs && node run.mjs
```
