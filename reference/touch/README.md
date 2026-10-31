# Touch targets of the console parts

Is every console part a finger presses big enough? The HIG's control sizes (Accessibility, Mobility, read 2026-10-29 from developer.apple.com's JSON for the page): on iOS 44x44pt by default, 28x28pt at least; on macOS 28x28pt, 20x20pt at least. App Store Connect's toolkit draws its controls 30px whatever the screen (its tokens name no larger control on its mobile page scale, `../asc/results/tokens.json`), and crashcart, the console's site, had held its own controls to 44pt on a touch screen until it moved onto ciderui's console parts (its HIG pass, 2026-10-21).

`audit.mjs` opens every console part's example at 390px with touch emulated (`test/browser.mjs`: `(hover: none)` and `(pointer: coarse)` match), takes each control a finger can press, and prints its box and its target: the larger of its box and a `::before` or `::after` it draws with `position: absolute`, a checkbox's or a radio's `<label>`. It exits 1 while a target is under 28pt.

```sh
node scripts/docker.mjs node reference/touch/audit.mjs
```

## What it showed (Chromium 151, 2026-10-29)

**Before** (ciderui 167a5c2): 23 targets under 28pt: the small buttons (22px; the banner's), the chips (22px) and a chip's remove × (12px), the checkboxes (15px, their labels 18px), the table's sort links (20px), the sidebar's open button (16 by 30px), a radio with no caption (18px), the small field (22px), the links in running text (17px), the skip link while hidden (1px). Between 28 and 44pt: the default buttons, tabs and fields (30px), the menu's items (37px), the hint (32 by 44px since 044dfc2).

**After**: on a touch screen a button of any size, a chip that is a link, a remove ×, a sort link: a transparent `::after` 44px at least each way, round it (the hint's way, apple.com's modal's close button's); a checkbox's transparent input 44px square on its 15px box; a radio's label 44px tall at least (not a larger target round it: stacked 12px apart, one reached the radio above); since 2026-10-31 a checkbox's `<label>` 44px tall at least too and a disclosure's line 44px each way, for the same reason (stacked 12px and 15px apart, the next one's target reached 14px into the checkbox above and 12px toward the summary above, and an opened summary's into its content; `audit.mjs` then: 0 under 28pt). Under 28pt still: the links in running text (17px, exempt as text is: WCAG 2.5.8's inline exception), the small field (`data-size="sm"`, a page's choice), and the skip link while hidden (clipped, `clip-path: inset(50%)`, which clips its target's hits too).

Left between 28 and 44pt, at the HIG's minimum and App Store Connect's own size: the tabs (30px; an underlined row on a phone scrolls sideways, and its scroll box would clip a target drawn outside it), the fields (30px), the menu's items (37px).
