# Touch targets of the console parts

Is every console part a finger presses big enough? The HIG's control sizes (Accessibility, Mobility, read 2026-10-29 from developer.apple.com's JSON for the page): on iOS 44x44pt by default, 28x28pt at least; on macOS 28x28pt, 20x20pt at least. App Store Connect's toolkit draws its controls 30px whatever the screen (its tokens name no larger control on its mobile page scale, `../asc/results/tokens.json`), and crashcart, the console's site, had held its own controls to 44pt on a touch screen until it moved onto ciderui's console parts (its HIG pass, 2026-10-21).

`audit.mjs` opens every console part's example at 390px with touch emulated (`test/browser.mjs`: `(hover: none)` and `(pointer: coarse)` match), 48px by 24px of room round it, takes each control a finger can press and can see (`checkVisibility()`: not one in a closed `<details>`; one hidden until focused, the skip link, read focused), and prints its box and its target: hit-tested from its center (a checkbox's or a radio's `<label>`'s), each way, as far as a tap still lands on the control, every 0.5px. A `::before` or `::after` drawn round it counts as far as it is not clipped (a scroll box) or covered (a neighbour's target); each way reads up to 0.5px short, so 43 holds a 44px target. It exits 1 while a target is under 44pt.

```sh
node scripts/docker.mjs node reference/touch/audit.mjs             # a step of npm test
VERBOSE=1 node scripts/docker.mjs node reference/touch/audit.mjs   # every control's line
```

## What it showed (Chromium 151, 2026-10-29)

**Before** (ciderui 167a5c2): 23 targets under 28pt: the small buttons (22px; the banner's), the chips (22px) and a chip's remove × (12px), the checkboxes (15px, their labels 18px), the table's sort links (20px), the sidebar's open button (16 by 30px), a radio with no caption (18px), the small field (22px), the links in running text (17px), the skip link while hidden (1px). Between 28 and 44pt: the default buttons, tabs and fields (30px), the menu's items (37px), the hint (32 by 44px since 044dfc2).

**After**: on a touch screen a button of any size, a chip that is a link, a remove ×, a sort link: a transparent `::after` 44px at least each way, round it (the hint's way, apple.com's modal's close button's); a checkbox's transparent input 44px square on its 15px box; a radio's label 44px tall at least (not a larger target round it: stacked 12px apart, one reached the radio above); since 2026-10-31 a checkbox's `<label>` 44px tall at least too and a disclosure's line 44px each way, for the same reason (stacked 12px and 15px apart, the next one's target reached 14px into the checkbox above and 12px toward the summary above, and an opened summary's into its content; `audit.mjs` then: 0 under 28pt). Under 28pt still: the links in running text (17px, exempt as text is: WCAG 2.5.8's inline exception), the small field (`data-size="sm"`, a page's choice), and the skip link while hidden (clipped, `clip-path: inset(50%)`, which clips its target's hits too).

Left between 28 and 44pt, at the HIG's minimum and App Store Connect's own size: the tabs (30px; an underlined row on a phone scrolls sideways, and its scroll box would clip a target drawn outside it), the fields (30px), the menu's items (37px).

## Every control at 44pt (Chromium 151, 2026-11-01)

The owner asked for every console control at the HIG's 44pt default on a touch screen ("fix all", then "Grow to 44"). `audit.mjs` then read a control's computed `::after` size, which a neighbour's target or a scroll box can cut: hit-tested instead (above), 86 controls were under 44pt, buttons that drew a 44px target among them (30px buttons 8px apart, 38px from center to center: each one's target covered its neighbour's, 36.5px left of it). A target drawn round a control cannot hold 44pt where controls sit closer than that, nor a tab in its scrolling row, nor a field (an `<input>` draws no `::after`): on a touch screen the control itself is drawn 44px at least each way, every size but xl (50px): a button, a tab (a pill's corners 22px), a field and a file field, a menu item, a chip that is a link or holds a remove × and the × at its end, an appearance choice's width (its height a target round it), the hint (its circle where it was), the head's name. A table cell holding a checkbox takes 14.5px either side for its input's 44px. Then: 126 controls, 0 under 44pt.

## Labels, links in text, and the other sets (Chromium 151, 2026-11-01)

A control's target counts every label it has (`control.labels`: one wrapping it or naming it by `for=`), so a hidden toggle is read at its chevron's label and a field under its floating label at either. A link inside a sentence (an inline `<a>` whose nearest block holds other text: an article's, a footnote mark's) is listed `text`, not counted: WCAG 2.5.8's inline exception. What cannot be tapped now is left out: hidden (`visibility`, a closed `<details>`), disabled, or taking no pointer with no label to take it. Then the site's controls under 44pt, at apple.com's own sizes, grew on a touch screen too (the owner: "Grow to 44"): the footer's directory and legal links, section titles and locale; the gallery's paddles, and its dots' row with room for each dot's target (the dots drawn as before); the local nav's title and menu button (its chevron where it was); the swatches, with room round each ring. Every set is in the gate: site 121 controls, the console 123, 0 under 44pt (2026-11-04).

```sh
node scripts/docker.mjs sh -c 'SETS=site node reference/touch/audit.mjs'   # some of the sets
```
