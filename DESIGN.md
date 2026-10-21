# ciderui: its design

ciderui is apple.com's look for websites: a Tailwind CSS v4 theme on a small token contract, apple.com's parts as CSS classes (`cui-btn`, `cui-field`, `cui-tile`, …), and a Rust crate that writes their markup with hypertext and carries the stylesheets; ciderui-zola, a Zola theme in a repository of its own, is built on it. It is vpkit's architecture (`../vpkit`, `../vpkit-zola`) with apple.com where vpkit has VitePress. This documents what is decided, before any component exists: what the standard is, what a component is, the rules every one follows, the token contract, the type, the catalogue, the crate, how a component is proven, and the decisions behind them.

A value called *measured* here was read off apple.com's live pages or a browser on 2026-10-19 and is in `reference/` with its probe, page, browser and date; *stated* is the owner's; *guessed* is nobody's measurement.

## Who it is for

The owner's websites, and nobody else's: ciderui is not public (stated). The sites are in English (stated), and they are both kinds of apple.com: marketing pages (a product, its story, its tiles) and the store (product cards, options to choose, a bag, sign-in and other forms) (stated). They render their pages in Rust with hypertext and build their CSS with the Tailwind CLI, as the owner's vpkit apps do (stated: "類似 ../vpkit，也要提供 rust hypertext 的 component"). Their static sites, and ciderui's own documentation, are built with ciderui-zola, as vpkit's are with vpkit-zola (stated: "也做一個 ciderui-zola"). With no consumer outside, a component takes the best shape and the sites follow it (decision 1).

## The standard: apple.com, measured

apple.com's own pages are the original, and they can be read: `reference/apple.com/` holds the computed styles of seven of them (the home page, iPhone 18 Pro's page, the store, a buy page, the bag, the store's sign-in) at 1600, 1200, 900 and 390px. Unlike VitePress for vpkit, apple.com's CSS is neither licensed nor versioned (its files carry cache hashes), so none of it is copied into this repository and there is no upstream to re-sync: what is recorded is what elements compute to, dated, and a re-measure is a deliberate act that changes the record and the checks that read it.

What the measurements say about apple.com, and so about ciderui:

- **Dark is a section's.** None of the 21 stylesheets the pages load reads `prefers-color-scheme`; a `.theme-dark` section is dark on a light page, a product's page dark as a whole. ciderui follows: dark is `color-scheme: dark` on any element (rule 11).
- **The type is the look, and it is responsive per class.** One class steps its size at a breakpoint by itself (`t-headline` is 48px from 735px, 32px below). Each size has one line height and one letter-spacing, and the letter-spacing follows the size as SF's tracking does.
- **apple.com tunes page by page.** The marketing pages' `typography-*` classes are defined by each page for itself (`typography-headline` is 32/40/48px on the home page and not defined on the product page) and the same size is tuned two ways (32px at +0.002em on the home page, +0.004em on the store). ciderui curates one system from them (the type, below), choosing where they disagree.
- **The palette is small**: a dozen colors and their alphas (the contract, below).
- **The breakpoints are apple.com's own**: 735px and 1069px for the content, 834px for the global nav, measured by stepping a headline and the nav across each boundary.

## What a component is

A component is a class, or a block of a few classes, that gives an element the look of one of apple.com's parts, and nothing else: no JavaScript, no page layout, no behavior beyond what CSS reads off the element's own state. The page opens the dialog and checks the option; the component shows the state the page set. Behavior comes from HTML's own elements where it exists: `<dialog>` for a sheet, `<details>` for an accordion and a menu, `popover` for a flyout, scroll snapping for a gallery. The components are apple.com's parts (a button, a tile, a field, an option, the global nav, the footer, the catalogue below) and the parts a site needs in apple.com's look where apple.com shows none.

## Rules

Every component follows these.

1. **One component, one directory.** `components/<name>/<name>.css`, in Tailwind's `components` layer, so a utility on the same element always wins (`class="cui-btn px-8"`); `examples/<id>.html`, the canonical markup, which the crate's tests render and the cases measure; `<name>.test.mjs`, its proof; `golden.json`, its recorded values. Nothing is unlayered; nothing uses `!important`.

2. **The class says what it is; what varies is an attribute.** The block is `cui-<name>`, a part `cui-<name>-<part>`. A mutually exclusive axis is `data-variant` (`data-variant="secondary"`, `"neutral"`) or `data-size` (`sm`, `lg`, `xl`); an orthogonal switch is a class; a state is what the markup carries anyway: an ARIA attribute (`aria-current`, `aria-invalid`, `aria-expanded`), a control's own state (`checked`, `disabled`, read as `:checked` and `:disabled`), `[open]`, or `data-state` where none says it. A variant's rule is `.cui-x[data-variant="y"]`, so a state one attribute more wins over every variant whatever the order.

3. **Component-local variables.** The block sets its own `--cui-<name>-*` from the contract; a variant overrides the locals and nothing else, so a component nested in another keeps its colors and recoloring one block is one rule on it.

4. **Rules for what is inside sit behind `:where()`**, adding no specificity, so a component placed inside another keeps its own look.

5. **Hover is `@media (hover: hover) { &:hover }`**, so a tapped element does not keep its hover.

6. **Size and placement are the page's**: no width, no margin, no position, with two kinds of exception, each named in its stylesheet's header: the type, whose classes carry their small-screen step (the type, below), and the sizes apple.com fixes (a control's height, a field's 56px, a tile's minimum height). Utilities remove them.

7. **Tokens: only the contract's.** A component reads tokens.css's colors and its non-color tokens and no literal color; the token test holds every stylesheet to it.

8. **apple.com's control sizes, not one height.** apple.com's pill comes in four heights, measured: 24px (`data-size="sm"`, 12px text), 36px (the default, 14px: the store's shared button styles' default and most of the store's buttons), 44px (`lg`, 17px), 56px (`xl`, 17px); and its marketing pages' tile buttons step at 735px, 36px below and 44px from it, which is `data-size="cta"`, a hero's call to action. A field is 56px; the store's form button 56px. This is where ciderui departs from vpkit's one control height (decision 7; the default is ciderui's recommendation, the owner deferring to it).

9. **One focus ring.** `base.css` draws `:focus-visible { outline: 2px solid var(--cui-c-accent); outline-offset: 3px }` on every element, apple.com's on every control it was read on (the home page's buttons, the store's, the bag's form buttons, the buy page's options); a text field shows its 2px accent edge in its place (`cui-field`: apple.com's field draws no ring for a click, `retail:field.input`, and its script's ring for the keyboard alone, which CSS cannot tell apart on a text field). A component draws no ring of its own; it restates the ring only to move it or, a field, to set it aside.

10. **apple.com's motion, through the tokens.** No button transitions its colors (apple.com's 0s, every button measured); the global nav's links fade their color in 0.32s `cubic-bezier(.4,0,.6,1)`, its menu grows over the window in 0.45s, its rows come in over 0.24s each from 0.2s, 20ms apart, its X turns in 0.24s (`home:menu.open`'s recorded animations); a paddle in 0.1s linear, a gallery's dot its color in 0.25s linear; a card's shadow and scale in 0.3s `cubic-bezier(0,0,.5,1)` (apple.com's `all`, which would also move its size as a window crosses 735px); a field's label in 0.125s. `base.css` neutralizes every transition and animation under `prefers-reduced-motion`.

11. **Dark comes through the tokens, set on any element.** Every color is a `light-dark()` pair declared once on `:root`; `cui-dark` on an element sets `color-scheme: dark` and every pair inside it resolves dark, `cui-light` inside it light again (measured in Safari 27.0.1, Chrome 155 and Chromium 151 and on the iOS 26.5 simulator, `reference/scheme/`). On `<html>` it darkens the whole page; apple.com's dark product page also clears its global nav over the hero (`product:nav`: no ground, no material), which is the global nav's own treatment, built with it, not the scheme's. There is no automatic dark mode: apple.com has none.

12. **Icons are inline SVG**, from the crate, in `currentColor`. apple.com draws its chevrons and plus signs with its own icon font (`SF Pro Icons`: `product:link.more-icon`, `buy:accordion.icon`), which is Apple's; ciderui draws them as SVG, so a page needs no font and no `data:` image for them (vpkit's decision 13).

13. **Proof.** A check's source is `apple.com` (its computed value equals the one `reference/apple.com/results/` records for the probe it names, at the width it names, for what is meant to be apple.com's: the colors at rest, hovered and pressed, the type, the borders, the radii, the padding, the heights, the motion), `rest` (a touch screen's hovered element against itself at rest) or `golden` (equal to `golden.json`: a part apple.com has no instance of, or a departure from it, with the check's `why`); a source comparing one rendering of ciderui with another (inside a dark block against alone) is added with the first component that needs it. No list of known differences: every departure is a golden check with its reason beside it.

14. **Ships complete.** The stylesheet with its header (what it is, from which page and probe, the markup, the attributes and parts, what is apple.com's and what is ciderui's), its examples, its test with its golden values, and the crate's component where its markup is more than one element. One commit per component, and one per step of any change, each with the tests green.

## The token contract

`tokens.css` is the contract: the colors every component reads, each one `light-dark()` pair, light first. The values are apple.com's (measured, `reference/apple.com/README.md`'s palette), except where marked:

| token | light | dark | its job |
| --- | --- | --- | --- |
| `bg` | `#ffffff` | `#000000` | the page |
| `bg-alt` | `#f5f5f7` | `#1d1d1f` | the alternate ground: a tile, the footer, a ribbon; a dark section's alternate |
| `text` | `#1d1d1f` | `#f5f5f7` | body text |
| `text-2` | `#6e6e73` | `#86868b` (fitted) | secondary text |
| `text-3` | `#86868b` | `#6e6e73` | muted text, at 3:1, never body |
| `border` | `#86868b` | `#6e6e73` | a control's edge: a field, an option, a checkbox; the dark ring of the floating nav's buttons |
| `divider` | `#d2d2d7` (`bag:accordion.item`) | `#424245` (`product:divider`, a rule on iPhone 18 Pro's dark page) | a rule between things: an accordion's items |
| `accent`, `accent-hover`, `accent-active` | `#0071e3`, `#0076df`, `#006edb` | the same | the primary fill, white on it; the focus ring; a checked option's edge |
| `link` | `#0066cc` | `#2997ff` | a link; a secondary button's text and edge |
| `neutral`, `neutral-hover`, `neutral-active` | `#1d1d1f`, `#272729`, `#18181a` | `#f5f5f7`, `#ffffff`, `#ededf2` (a dark gallery item's, `home:gallery.button-neutral`) | the neutral fill |
| `on-accent`, `on-neutral` | `#ffffff`, `#ffffff` | `#ffffff`, `#000000` | text on the accent's and the neutral's fills |
| `field` | `rgba(255,255,255,.8)` (`retail:field.input`) | the same (`retail-dark:field.input`: apple.com's field in a dark color scheme) | a field's and an option's ground; a field keeps its light look in a dark block |
| `field-error` | `#fff2f4` (`giftcard:error`) | the same (`retail-dark:field.error`) | a field's ground in error |
| `field-disabled` | `rgba(0,0,0,.01)` (`giftcard:field.disabled`) | the same (`retail-dark:field.disabled`) | a disabled field's ground |
| `field-prefix` | `rgba(0,0,0,.02)` (`giftcard:prefix.currency`) | the same (a field is light on any page) | the ground of a currency or a country code before a field |
| `border-disabled` | `rgba(134,134,139,.42)` (`giftcard:field.disabled`) | the same (`retail-dark:field.disabled`) | a disabled field's edge |
| `error` | `#e30000` (`giftcard:error.message`) | `#ff453a`, fitted: apple.com's is 4.3:1 on black, below 4.5:1 (Apple's dark red) | a field's edge, label and message in error |
| `bar` | `rgba(255,255,255,.8)` | `rgba(0,0,0,.8)` (the Newsroom's dark nav, `newsroom:nav`) | the global nav's ground, under its material |
| `on-bar`, `on-bar-hover` | `rgba(0,0,0,.8)`, `#000000` | `rgba(255,255,255,.8)`, `#ffffff` | the nav's links and icons, and hovered; the menu's button is `on-bar-hover` at 0.8 opacity |
| `menu`, `on-menu` | `#fafafc`, `#333336` | `#161617`, `#e8e8ed` | the nav's menu opened over the window on a phone, and its links; the local nav's menu |
| `bar-2`, `on-bar-2` | `rgba(250,250,252,.8)`, `rgba(0,0,0,.88)` | `rgba(22,22,23,.8)`, `rgba(255,255,255,.92)` | the local nav's material and its product's name (apple.com's black at opacity .88, white at .92, as one color) |
| `separator` | `rgba(0,0,0,.16)` | `rgba(255,255,255,.24)` | a rule on a bar (the local nav's foot); a footer's rules |
| `tile` | `#f5f5f7` | `#000000` | a tile's ground (`home:hero.tile`, `home:tile.dark`) |
| `on-card` | `#1d1d1f` | `#ffffff` (a dark card's, `store:card.dark-header`) | a card's text, on the page's ground: its headline, description and price |
| `eyebrow`, `tag` | `#424245`, `#b64400` | `#f5f5f7`, `#ff791b` (`store:card.dark-eyebrow`, `.dark-tag`) | a card's eyebrow, and a tag in its place ("NEW") |
| `paddle`, `-hover`, `-active` | `rgba(210,210,215,.64)`, `rgba(223,223,227,.698)`, `rgba(193,193,198,.655)` (`iphone:paddle`) | `rgba(66,66,69,.72)`, `rgba(73,73,76,.725)`, `rgba(57,57,60,.74)` (`product:paddle.next`) | a gallery's paddles, translucent over the page |
| `on-paddle`, `-hover` | `rgba(0,0,0,.56)`, `rgba(0,0,0,.64)` | `rgba(255,255,255,.8)`, `#ffffff` | a paddle's chevron, and hovered or pressed |
| `on-footer`, `on-footer-link`, `on-footer-title`, `footer-rule` | `rgba(0,0,0,.56)`, `.72`, `.88`, `.48` (`home:footer*`) | the same (`product:footer*`: apple.com's footer is light on a dark page) | the footer's text, links and titles on the alternate ground; its legal links' separators |
| `dot`, `dot-hover` | `rgba(29,29,31,.6)`, `rgba(0,0,0,.54)` (`iphone:gallery.dot`) | `rgba(255,255,255,.36)`, `rgba(255,255,255,.48)` (`product:gallery.dot`) | a gallery's dots, and hovered |
| `scrim` | `rgba(0,0,0,.48)` (`overlay:overlay`) | the same (a modal is light on any page) | the layer a modal lays over the page |
| `curtain` | `rgba(232,232,237,.4)` (`search:curtain`) | `rgba(0,0,0,.4)` (the Newsroom's dark nav, `search-dark:curtain`) | the layer over the page under the global nav's search, under a blur |
| `close`, `close-hover`, `close-active`; `on-close`, `on-close-hover` | `#e8e8ed`, `#ececf0`, `#dfdfe4`; `#6e6e73`, `rgba(0,0,0,.72)` (`overlay:close.circle*`) | the same | a modal's close button's circle, hovered and pressed, and its X |

tokens.css holds the slots measured on both sides (`bg`, `bg-alt`, `text`, `text-2`, `text-3`, `border`, the accent's three and `on-accent`, `link`, the neutral's three and `on-neutral`, `bar`, `on-bar`, `on-bar-hover`, `menu`, `on-menu`, `curtain`, `bar-2`, `on-bar-2`, `separator`, `tile`, `on-card`, `eyebrow`, `tag`, the paddle's three, `on-paddle` and its hover, `dot` and its hover, `field`, `field-error`, `field-disabled`, `field-prefix`, `border-disabled`, `error` (its dark side fitted), `divider`, `on-footer`, `on-footer-link`, `on-footer-title`, `footer-rule`, `scrim`, `close` and its hover and press, `on-close` and its hover); a slot with a guessed side enters with the first component that reads it, that side measured then, so no stylesheet ever reads a guess. `test/tokens.mjs` holds each to the probe beside it.

Their contrast, computed (WCAG 2): `text` 16.8:1 on `bg` and 15.5:1 on `bg-alt` (dark: 19.3:1, 15.5:1); `text-2` 5.1:1 and 4.7:1 (dark: 5.8:1, 4.7:1); `link` 5.6:1 and 5.1:1 (dark: 7.0:1, 5.6:1); white on `accent` 4.7:1, on `accent-hover` 4.5:1, on `accent-active` 4.9:1; `border` 3.6:1 on `bg` and 3.3:1 on `bg-alt` (dark: 4.1:1, 3.3:1); a card's text on its ground, `bg`: `on-card` 16.8:1 (dark 21:1), `eyebrow` 10.0:1 (dark 19.3:1), `tag` 5.5:1 (dark 8.0:1); a paddle's chevron on the paddle over `bg` 4.6:1 and over `bg-alt` 4.6:1 (dark: 9.1:1, 8.1:1), a dot 4.5:1 and 4.3:1 (dark: 3.1:1, 3.3:1), both marks, not text, held to WCAG 2's 3:1; `on-bar` on the bar over `bg` 12.6:1 and over `bg-alt` 12.4:1 (dark: 13.1:1, 12.7:1), a translucent color counted as what it draws over the one under it; `on-menu` on `menu` 12.1:1 (dark: 14.8:1); the local nav's name on `bar-2` over `bg` 16.1:1 (dark 15.9:1), its links 12.4:1 (dark 12.1:1). apple.com's dark secondary text, `#6e6e73` (the product page's stat captions), is 4.1:1 on black and 3.3:1 on `#1d1d1f`, below 4.5:1, so the contract's dark `text-2` is `#86868b` (5.8:1 and 4.7:1) and apple.com's `#6e6e73` is dark `text-3` and `border` (decision 6). The token test holds these minimums on both sides: body and secondary text 4.5:1 on `bg` and `bg-alt`, `text-3` and `border` 3:1, white 4.5:1 on the accent's three grounds, the nav's text 4.5:1 on its bar over both grounds and on its menu, a card's text 4.5:1 on `bg`, a gallery's marks 3:1 on both grounds.

The palette is fixed: no site redefines a color, and there are no themes (stated: "顏色不能換"); apple.com has one palette. The token test holds tokens.css to the values above.

Not themed, in tokens.css too:

- **The radii**: the pill `980px`; a field, an option and a form button `12px`; a block button `8px`; a card `18px`; a rounded tile `28px`; a checkbox `3px`; a swatch and a paddle `50%` (measured, each with its probe).
- **The control sizes**: 24, 36, 44 and 56px (rule 8).
- **The focus ring**: 2px, the accent, 3px out; a text field shows its 2px accent edge instead (rule 9).
- **The durations**: 0.1s, 0.125s, 0.24s, 0.25s, 0.3s, 0.32s, 0.45s with apple.com's curve (rule 10).
- **The materials**: the bar's `saturate(1.8) blur(20px)`, the floating nav's and a curtain's `blur(20px)` (measured).
- **The shadows**: a card's `2px 4px 12px rgba(0,0,0,.08)`, hovered `2px 4px 16px rgba(0,0,0,.16)`, `--cui-shadow-card` and `-hover` (measured, `store:card`; a dark card's the same, `store:card.dark`).
- **The widths**: the content's 87.5% up to 1260px (420px at most below 735px), apple.com's newer pages'; the older 980px with 22px edges (16px below 735px), its footer's; the global nav's 1024px (measured).
- **The font stack** (the type, below), and code's, `--cui-font-mono`: `ui-monospace, "SF Mono", Menlo, Consolas, "Liberation Mono", monospace`, SF Mono on Apple's platforms, the platform's own elsewhere, no font shipped (apple.com's pages show no code: ciderui's, for an article's code).
- **The widths**, the Newsroom's: an article's text 653px from 1069px, 576px from 735px, 87.5% of 414px at most below; its figures 980px, 692px, 414px at most (measured, `newsroom-article:column.*`).

## The type

**The stack** is `-apple-system, BlinkMacSystemFont, Inter, "Helvetica Neue", Helvetica, Arial, sans-serif`. SF cannot be shipped; on Apple's platforms the first two names are SF (`-apple-system` in Safari, `BlinkMacSystemFont` in Chrome; `system-ui` is both, but is a font everywhere, so Inter after it would never be drawn). Elsewhere neither name is a font and the text is Inter, the kit's (`fonts.css`, Inter 4.001 with its `opsz` axis), fetched only there: in Safari 27.0.1, Chrome 155 and iOS 26.5's Safari the stack draws the system font and not one of Inter's 16 faces is fetched; in Chromium on Linux it draws Inter with its Latin face alone (measured, `reference/stack/`). The owner chose Inter over the platform's own font and over apple.com's Helvetica and Arial (stated).

**The ladder** is the store's `t-` classes (measured, `reference/apple.com/`: one coherent ladder where the marketing pages define theirs page by page), named `cui-<name>`, each with one step at 735px:

| class | from 735px | below | weight |
| --- | --- | --- | --- |
| `cui-headline-standalone` | 96/100 | 48/52 | 600 |
| `cui-headline-super` | 80/84 | 48/52 | 600 |
| `cui-headline-elevated` | 64/68 | 40/44 | 600 |
| `cui-headline` | 48/52 | 32/36 | 600 |
| `cui-headline-reduced` | 40/44 | 28/32 | 600 |
| `cui-eyebrow-super` | 32/36 | 24/28 | 600 |
| `cui-eyebrow-elevated` | 28/32 | 21/25 | 600 |
| `cui-eyebrow` | 24/28 | 21/25 | 600 |
| `cui-eyebrow-reduced` | 21/25 | 19/23 | 600 |
| `cui-callout` | 32/36 | 24/28 | 600 |
| `cui-label` | 24/28 | 21/25 | 600 |
| `cui-quote` | 40/48 | 28/36 | 600 |
| `cui-quote-reduced` | 32/40 | 24/32 | 400 |
| `cui-manifesto` | 32/40 | 24/32 | 600 |
| `cui-intro-elevated` | 24/32 | 21/29 | 400 |
| `cui-intro` | 21/29 | 19/27 | 400 |
| `cui-tout` | 19/23 | 19/23 | 600 |
| `cui-body` | 17/25 | 17/25 | 400 (the page's body text: the store defines no `t-body`) |
| `cui-body-tight` | 17/21 | 17/21 | 400 |
| `cui-body-reduced` | 14/20 | 14/20 | 400 |
| `cui-body-reduced-tight` | 14/18 | 14/18 | 400 |
| `cui-caption` | 12/16 | 12/16 | 400 |

The marketing pages' three-tier headlines (the home page's tile headline: 32px below 735px, 48px from 735px, 56px from 1069px) are the tile's, built with the tile (the catalogue).

**The letter-spacing** is per size, and it is not apple.com's numbers: apple.com sets them on its SF Pro webfonts, which carry no tracking, while the system font carries Apple's size-specific tracking itself, the same in Safari and Chrome on macOS and in Safari on iOS (17px: −0.025em of its own). apple.com's values on the system font would track it twice. ciderui's are apple.com's minus the system font's own, measured size by size (`reference/tracking/cui-tracking.json`), and land its text within 0.0006em a character of apple.com's at all 14 sizes in all three browsers:

| size | 12 | 14 | 17 | 19 | 21 | 24 | 28 | 32 | 40 | 48 | 56 | 64 | 80 | 96 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| em | −0.0105 | −0.0058 | +0.0029 | −0.0129 | −0.0049 | −0.0040 | −0.0073 | −0.0093 | −0.0099 | −0.0109 | −0.0110 | −0.0130 | −0.0156 | −0.0156 |

Where apple.com tunes a size two ways, the store's is taken (32px +0.004em, 48px −0.003em; the home page's +0.002em, −0.002em). CSS has no letter-spacing per font of a stack, so Inter takes SF's values and runs wide of apple.com's text: by up to 0.0243em a character at 400 (21px), 0.0208em at 17px, about 21px over a 60-character line (measured, `reference/tracking/`). That is the price of SF exact on Apple's platforms, the owner's choice of Inter (decision 3).

**The classes are Tailwind utilities** (`type.css`, `@utility`), not component classes: measured in Tailwind 4.3.3 and Chromium 151, a class holds its own step at 735px, a variant composes with it (`cui-headline lg:cui-headline-super` is 32, 48 and 80px at 390, 900 and 1069px), and a plain utility on the same element wins at every width (`cui-headline text-[20px]` is 20px), as rule 1 wants of a component. `test/type.mjs` holds each class to the store's at 390, 734, 735, 900, 1200 and 1600px.

## The page

**Breakpoints.** theme.css replaces Tailwind's with apple.com's, in pixels as apple.com's media queries are: `md` 735px, `lg` 1069px, `xl` 1441px (apple.com's stylesheets name 1441px; nothing measured changes there), and `nav` 834px, the global nav's. A component's step and a page's `md:` switch at the same width.

**Sections.** A section is the page's ground (`bg`), the alternate (`bg-alt`) or dark (`cui-dark`, its alternate `bg-alt` dark: `#1d1d1f`, measured on the product page's highlights); its content is the newer width (87.5% to 1260px). Their padding is a section's own (apple.com's varies by section; measured per component).

## The catalogue

What apple.com's pages are made of, from the measurements, as the components to build, "(built)" beside each that is. Each is held to the probes named (`reference/apple.com/results/<page>.json`):

| component | apple.com's | from |
| --- | --- | --- |
| `cui-btn` (built) | the pill: five sizes (`data-size` sm, lg, xl, cta), the accent, secondary, neutral and secondary-neutral (`data-variant`), on light and in a dark block, disabled; 12px corners (`cui-btn-rect`; with `data-size="xl"` the store's form button) and the block (`cui-btn-block`), orthogonal classes | `store:button*`, `home:button*`, `home:tile.dark-button*`, `home:gallery.button-neutral`, `bag:form-button*` |
| `cui-link` (built) | a link in running text (underlined at rest, as apple.com's own `a` is), and the more link (`data-variant="more"`, its text underlined hovered, its chevron, the crate's `MoreLink` and `icon::MORE`) | `iphone:link.*`, `product:link.*`; the chevron `reference/icons/` |
| `cui-section` (built) | a band: its ground (`data-variant="alt"`, dark with `cui-dark`), its rhythm (the iPhone page's, `data-size="lg"` iPhone 18 Pro's), its content's width (`cui-section-content`), its header, headline and links (the crate's `Section`, `SectionHeader`) | `iphone:section.*`, `product:section.*`, `*:content` |
| `cui-globalnav` (built) | the bar (44px, 48px below 834px) on its material, light, dark (`cui-dark`) or clear (`data-variant="clear"`); the site's logo, its links, its search and bag (`icon::SEARCH`, `icon::BAG` and their compact forms); below 834px its menu, the list a popover its button opens from beside it, a sheet over the window with a close button of its own, which takes the focus, its lines turning into an X (`icon::MENU`; the crate's `GlobalNav`); its search a link to the site's search page, or a button opening the search, another popover, no script: its field focused (`icon::SEARCHFIELD`), sent GET to the site's search page, its quick links (`icon::ARROW`); from 834px a flyout hung from the bar's bottom (anchor positioning) over the curtain (`curtain`), a button that closes it, as a click on the bar does; below 834px a sheet over the window with its own X, the page held (`Search::Link`, `Search::Flyout`); not placed on the page (a site adds `sticky top-0`); no desktop menus under a link | `home:nav*`, `home:menu*`, `newsroom:nav*`, `newsroom:menu*`, `product:nav`, `search:*`, `search-dark:*`; the icons `reference/icons/`; the menu `reference/popover/`; the search `reference/search/` |
| `cui-localnav` (built) | the local nav most product pages carry (not iPhone 18 Pro's newer floating one): the bar, sticky, 52px (48px below 834px) on its material (`bar-2`, dark with `cui-dark`) with a rule along its foot; the product's name, its pages' links (`aria-current="page"` marked by a line), its button (`cui-btn`, `data-size="sm"`); below 834px a menu its chevron opens below the bar, a checkbox and its label as apple.com's (`icon::CHEVRON`; the crate's `LocalNav`) | `localnav:*`, `localnav-dark:*`; the chevron `reference/icons/` |
| `cui-ribbon` (built) | the message above the content on the alternate ground, a more link after it; `data-variant="announce"` drops in on the accent and turns to its ground (the crate's `Ribbon`) | `ribbon:*`, `bag:ribbon` |
| `cui-tile` (built) | a tile on its ground (`tile`, black with `cui-dark`): a promo tile (580, 490, 500px) or, `data-size="lg"`, a full-width one (692, 650, 500px); its headline, subhead and callout in their tiers, its buttons; its picture filling it, the whole tile a link under its buttons; 28px corners with `cui-tile-rounded` (the crate's `Tile`) | `home:hero.*`, `home:tile.*`, `home:type.tile-*`, `home:type.promo-*`, `product:tile.rounded` |
| `cui-card` (built) | the store's card: 400 by 500px (309 by 450px below 735px) on the page's ground, dark with `cui-dark`, 18px corners, its shadow, lifted hovered or with the focus in it; its picture filling it (`cui-card-media`), its eyebrow or tag (`data-variant="tag"`), headline, description and price; the whole card its headline's link, with no script (the crate's `Card`) | `store:card*` |
| `cui-shelf` (built) | the store's row of cards: its header, a headline and its secondary text after it, over a scroller across the window, its items snapping into place 20px apart, the first and the last as far in as the header (22px to 1024px, 140px from 1440px, a straight line between), its scrollbar hidden; its paddles a gallery's large ones (`cui-shelf-paddle` on `cui-gallery-paddle`, `data-size="lg"`) at its sides, the one at an end hidden, none on a screen with no hover, their state the page's script's (the crate's `Shelf`) | `store:shelf*` |
| `cui-gallery` (built) | a product page's gallery: a row of items across the window, each snapping where the content starts, 20px apart; its paddles' row at the content's right edge, 36px paddles (56px, the store's, `data-size="lg"`) translucent over the page, their chevrons `icon::PADDLE_PREVIOUS` and `icon::PADDLE_NEXT`, disabled at an end; its dots, 8px, the current a 48px pill (32px below 735px), each a link to its item; the paddles' and dots' state the page's script sets, as it opens a dialog (the crate's `Gallery`) | `product:gallery.*`, `product:paddle*`, `iphone:paddle*`, `iphone:gallery.*`, `store:shelf.paddle`; the chevrons `reference/icons/` |
| `cui-field` (built) | apple.com's own 56px text field (the store's form-textbox): its edge, its translucent ground, its label the placeholder until the field is focused or filled, then small at its top, with no script (its placeholder a space); focused, its 2px accent edge; light in a dark block, on a ground of its own; in error (`aria-invalid`, the page's to set) its edge and label red on a pink ground and its message under it after a mark (`icon::ALERT`); disabled; a currency or a country code at its start (`cui-field-prefix`); a textarea in the field's box (`cui-field-box`, `cui-field-textarea`), 3 lines, taller by its handle; the store's dropdown, a `select` with its look and a chevron (`icon::DROPDOWN`), its label the placeholder while its empty choice is chosen (the crate's `Field`, `TextArea`, `Select`) | `retail:field*`, `retail-dark:field*`, `giftcard:*`, `tradein:dropdown*`, `retail:textarea*`, `feedback:*`; the icons `reference/icons/`; the dropdown in each browser `reference/docs/` |
| `cui-checkbox` (built) | the 16px box before its label, the input itself kept under it, transparent; checked, the accent with a white check (`icon::CHECK`); focused from the keyboard, the ring on the box (the crate's `Checkbox`) | `signin:checkbox*`; the check `reference/icons/` |
| `cui-option` (built) | an option on a radio: a tile 83px tall at least, 12px corners, its edge, its title (a detail under it) and its price; the radio itself over it, transparent; checked, a 2px accent edge; focused from the keyboard, the ring; light in a dark block, as a field (the crate's `OptionTile`) | `buy:option*` (checked: apple.com's classes on a checked copy the probe adds, `buy:option.checked`) |
| `cui-swatch` (built) | a color swatch, the page's color or image 28px in its 36px ring, on a radio; checked, the ring the accent; the color's name hidden but read; a row of them that wraps (`cui-swatches`; the crate's `Swatch`) | `buy:swatch*` (checked: the same, `buy:swatch.checked*`) |
| `cui-accordion` (built) | apple.com's FAQ on `<details>`, no script: a rule above each question (`divider`), its chevron at the end (`icon::DISCLOSURE`) in the muted text, the secondary hovered, turned up open; its answer below; `data-size="lg"` the bag's large question; its height moving where the browser can (the crate's `Accordion`, `AccordionItem`) | `bag:accordion*`, `bag-accordion:*`; `reference/details/` |
| `cui-quicklinks`, `cui-productnav` (built) | the store's row of pill links under a title, as far in as a shelf's header (`cui-btn`, `data-variant="secondary-neutral"`, at 38px; the crate's `QuickLinks`); its shelf of small cards, each a product's image above its name, a link, on `cui-shelf` (its row, snapping, paddles; the crate's `ProductNav`, `ProductNavItem`) | `store:quicklink*`, `store:productnav*` |
| `cui-footer` (built) | the footer on the alternate ground, light on a dark page too: its footnotes over a rule, its directory of columns (to 833px a column of sections opening one by one by a checkbox and its label, `icon::FOOTER_CHEVRON`), the shop line over a rule, the legal row (the copyright, its links 9px either side of a short rule, the locale); four tokens of its own (the crate's `Footer`, `FooterSection`) | `home:footer*`, `home-footer:*`, `product:footer*`; the chevron `reference/icons/` |
| `cui-price`, `cui-stat`, `cui-footnote` (built) | the buy page's price, 17/25 at 600 (12/16 below 735px), its amounts kept whole (`cui-price-amount`); a stat, a value between its captions, iPhone 18 Pro's or, `data-size="lg"`, AirPods Pro's in three tiers (the crate's `Stat`); a footnote's mark, a link at its text's size on its baseline, underlined, its digits the font's numerators (`font-feature-settings: "numr"`, SF's and Inter's; the crate's `Footnote`) | `buy:price.*`, `product:stat.*`, `airpods:stat.*`, `product:footnote*`, `airpods:footnote*` |
| `cui-article` (built) | apple.com's Newsroom story: its header (an eyebrow of its category over its date, the headline at 700, a subhead at 500) over its copy, the text's column narrower than its figures', 19/27 (17/25 below 1069px), paragraphs their line height apart, subheads 24/28 at 700, links underlined, a footnote's mark 0.6em, a figure and its caption, its footnotes; markdown's plain elements styled as they come, h3 to h6, code, a block of it, a quotation, a table and a rule ciderui's, and a word too long for its line broken where it must (apple.com's breaks none) (the crate's `Article`) | `newsroom-article:*` |
| `cui-stories` (built) | apple.com's Newsroom index: a band on the alternate ground, its heading at 700 over a grid of twelve columns 36px apart (24px below 1069px), each story a white tile rounded 32px (24px), the whole of it its link (no hover, no shadow): its picture 16:9, its category, its headline at 700, its date at its foot; three across by default, `data-size="md"` two across, `"lg"` one across beside its picture; two across whatever its size from 735px, one under another below (the crate's `Stories`, `Story`, `StorySize`) | `newsroom-index:*` |
| `cui-posts` (built) | apple.com's Newsroom archive: a month's heading at 700 over rows ruled above in the divider's color, each a link (no hover): its picture (295 by 166, 264 by 148.5, 105 by 59) before its category (uppercase), its headline at 700 and its date, the last row no room below; `data-size="sm"` the related stories under an article, in its text column, ruled below, a 171 by 96 picture (93px square below 735px) (the crate's `Posts`, `Post`, `PostSize`) | `newsroom-archive:*`, `newsroom-article:related.*` |
| `cui-modal` (built) | the store's overlay on `<dialog>`, no script: the page's button opens it (`command="show-modal"`), the layer fixed over the window scrolling a white panel (18px corners, 816px at most, 640px below 735px, its foot square there), the page dimmed (`scrim`) and held, fading in; its close button at the panel's corner (a 36px circle, a 44px target, `icon::CLOSE`), a form's (`method="dialog"`), and Escape closing it; light in a dark block; `data-variant="page"`, the product pages' modal: the page dimmed and blurred (`blur(20px)` on its `::backdrop`), the panel centered, 28px corners, 816/692/480px, its close button the neutral fill (`icon::CLOSE_SMALL`) sticking at the top as it scrolls, its scrim and panel coming in (the crate's `Modal`, its `page`, `ModalButton`) | `overlay:*`, `ricmodal:*`; the X's `reference/icons/`; `reference/dialog/` |

No docs theme (decision 14): support.apple.com's user guides are not measured.

Not measured yet: the search field with text in it (apple.com's suggestions and clear button), the checkout's own pages (not reachable without a bag: its form parts were read on the gift card's, the trade-in's and the feedback form's pages), a dropdown in error, a dark product page's modal, a touch screen's hover. When a part is built, its probes are added where the inventory's are not enough (rule 14: a component's test names its probes).

## The crate

`rust/` is the crate `ciderui`, for a site that renders with hypertext (0.12.1, as vpkit's apps do; the newest on crates.io on 2026-10-19):

- hypertext components for the markup that is more than a class on one element, in the catalogue's order: the more link (`MoreLink`, its chevron), a section and its header (`Section`, `SectionHeader`), the global nav (`GlobalNav`, its search a `Search`), the local nav (`LocalNav`), the ribbon (`Ribbon`), a tile (`Tile`), a card (`Card`), a shelf (`Shelf`), a gallery (`Gallery`), a field and its label (`Field`), a textarea (`TextArea`), the dropdown (`Select`), a checkbox (`Checkbox`), an option (`OptionTile`), a swatch (`Swatch`), an accordion (`Accordion`, `AccordionItem`), the quick links (`QuickLinks`), the product nav (`ProductNav`, `ProductNavItem`), the footer (`Footer`, `FooterSection`), a stat (`Stat`), a footnote's mark (`Footnote`), an article (`Article`), the stories (`Stories`, `Story`), the posts (`Posts`, `Post`), the modal (`Modal`, and `ModalButton`, its opener: hypertext's `<button>` has no `command` or `commandfor`). A button, a link in running text or a price (its amounts spans in the page's own sentence) is a class string on the page's own elements. Each component's test renders its example and asserts the output equal to it, so the markup and the CSS cannot part. A component writes ciderui's classes and nothing else.
- `icon`: the icons as inline SVG in `currentColor` (rule 12).
- `css` (a feature): every stylesheet embedded, written out by `ciderui-css`, so a site's cargo dependency at a commit is the one pin for its markup and its CSS (vpkit's decision 8). The font files are not embedded; `fonts.css` reads them beside the compiled stylesheet.

## The documentation

`docs/` is ciderui's documentation, a site built with ciderui-zola (decision 14): getting started, the foundations, the crate, a page for each component with each of its examples live in a frame of its own over its markup, and news in the Newsroom's look (`docs/README.md`). `npm test` holds it to the kit: the examples' copies and the frames' stylesheet are this commit's (`scripts/build-docs.mjs --check`), and `test/docs.mjs` wants a page for every component showing every one of its examples, the frames' stylesheet holding every rule the examples' classes have, and, the site built with zola, every page at 1440 and 390px in Chromium with no error, no missing file or dead link and no sideways scroll, each frame as tall as its content with every picture loaded. Its pictures are ciderui drawn by itself (`scripts/docs-pictures.mjs`), held to what they were drawn from (`docs/pictures.json`, each scene's and the stylesheet's digest: `--check` in `npm test`), not to Chromium's bytes. Outside Chromium it is measured in Safari 27.0.1 on macOS, Safari 26.5 on the iOS simulator and Chrome 155 (`reference/docs/`: every page and frame, the gallery's and the shelf's paddles, the menus in their frames).

## How a component is proven

`test/components.mjs` renders each case in headless Chromium and compares computed styles by the check's source (rule 13). What the measurements taught the harness, which it keeps:

- **Linux's hinting.** Playwright's Chromium on Linux hints every glyph to a whole-pixel advance unless launched with `--font-render-hinting=none`; with it, its text widths are the Mac's within 0.00006em a character (measured, `reference/tracking/`). Every run passes the flag.
- **The fonts differ between the two sides.** apple.com's values were read with its SF Pro webfonts; ciderui's cases render on Linux in Inter (the stack falls through). So an `apple.com` check compares what does not depend on the text's width (the type's own values, colors, borders, radii, padding, heights where the line height fixes them), and a box that depends on the text's width is a `reference` or `golden` check.
- **Transitions.** A state is forced with transitions off, and styles are flushed before they return, or a later read lands in flight.
- **States the protocol cannot force.** `:checked` is not a pseudo-class DevTools forces, so a checked state is proven by an example with `checked` in its markup, against the probe that measured apple.com's classes on a checked copy (`buy:option.checked`); such a probe says so in its name and its spec.
- **What only Safari or iOS can answer** (the system font's tracking, a section's color scheme, the stack's fetching, the nav's popover menu) is measured on the Mac (`reference/browsers.mjs`'s `safari`, `ios`, `chrome`), not assumed from Chromium.
- **States a click makes.** `:popover-open` is not forced either: a case clicks the element that makes it (`click`, the nav's menu button), then moves the pointer off, which a click leaves over what it clicked, hovered; it presses keys (`press`: Tab, Escape) and reads where the focus is (`focused`); `wait: 0` reads its motion's start (the rows still transparent: `@starting-style` through the compiled stylesheet).
- **A scroller's snapping.** `scroll` scrolls a case's scroller to a place before its checks, at every width: between two snap positions it comes to rest at one (a shelf's second card at the inset), at its end its last item as far in as its first.
- **What the pointer finds.** A link laid over a box (a card's, its `::after` over the card) has no box of its own a check can read: `hit` reads what the pointer finds at five points of the box, a tenth in from each corner and its middle.

The token test holds the contract (every component reads only it) and its contrast minimums; the Rust test holds the crate to `cargo fmt`, clippy with warnings as errors, its tests and its documentation, as vpkit's does.

## Browsers

The kit needs `light-dark()` (vpkit's floor: Chrome 123, Safari 17.5, Firefox 120, stated there); the global nav's menu needs `popover`, `:has()` and `@starting-style` besides (measured in the browsers below, `reference/popover/`), its search anchor positioning and `backdrop-filter` (`reference/search/`), and the modal `<dialog>` and the `command` and `commandfor` attributes, to open with no script (`reference/dialog/`; a browser without them opens it by the page's `showModal()`). ciderui's claims are measured in Safari 27.0.1 on macOS 27.0.1, Safari 26.5 on the iOS 26.5 simulator, Chrome 155 on macOS and Chromium 151 on Linux; not in Firefox, not on a real iPhone, not on Windows or Android.

## Decisions

1. **The sites do not bind a component's shape.** Every consumer is the owner's; no compatibility is kept.
2. **apple.com, measured, is the standard.** Its live pages' computed values, recorded with their probes and dates; its CSS never copied.
3. **SF exact, Inter near.** The letter-spacing is fitted to the system font (within 0.0006em of apple.com's text); Inter, where the system font is not SF, takes the same values and runs up to 0.0243em a character wide (the owner chose Inter).
4. **The stack names the system font by its two Apple-only names**, so Inter is drawn, and fetched, only where SF is not.
5. **Dark is set on any element**, by `color-scheme`; no automatic dark mode.
6. **The palette is apple.com's, fitted where it falls short**: the dark secondary text is `#86868b`, apple.com's dark `#6e6e73` being below 4.5:1; the token test holds the minimums.
7. **apple.com's control sizes** (24, 36, 44, 56px) where vpkit has one height.
8. **One focus ring**, apple.com's: 2px of the accent, 3px out.
9. **apple.com's breakpoints**, in pixels, replace Tailwind's.
10. **The store's ladder is the type scale**, apple.com's one coherent ladder; the marketing pages' three-tier headlines are a tile's.
11. **No JavaScript**: HTML's own elements for behavior.
12. **Icons are inline SVG**, no icon font.
13. **The crate carries the CSS**, one pin for a site's markup and its stylesheets.
14. **One product and a theme**: the site kit at the root; ciderui-zola, the Zola theme built on it, in a repository of its own, which ciderui's documentation is built with (stated). No docs theme, where vpkit has `docs-theme/` (stated, 2026-10-19: 「docs-theme 不做」). ciderui-zola's first site is ciderui's documentation, with news in apple.com's Newsroom's look (`cui-article`, `cui-stories`, `cui-posts`), and it ships the one script the gallery's and the shelf's paddles and dots need, which ciderui leaves to the page; its repository is local, beside this one (stated, 2026-10-19, `../ciderui-zola/DESIGN.md`).
15. **The palette is fixed** (stated): no site color, no themes.
16. **apple.com's look as it is, not iOS 26's** (stated, 2026-10-19: asked about iOS 26's style without its glass, the owner answered 「維持 ios 18」): the bar translucent under a blur and the pill, as apple.com shows them measured today; no Liquid Glass and no iOS 26 shapes.

## Assumptions

- Measured: apple.com's values in `reference/apple.com/` (eight pages, 2026-10-19); the system font's tracking and its fix (`reference/tracking/`); the stack's fonts and fetching (`reference/stack/`); a section's color scheme (`reference/scheme/`); the palette's contrast, computed from the measured values; the nav's popover menu and the accordion's `<details>` in each browser of the list above (`reference/popover/`, `reference/details/`), and the documentation built with ciderui-zola (`reference/docs/`; Safari on the Mac by its keyboard, Return and Option-Tab; on the iOS simulator no key is pressed, so the popover's Tab, Escape and focus-return claims are not run there).
- Stated by the owner: the sites are the owner's alone, in English, marketing and store; hypertext and the Tailwind CLI; Inter where SF is not; the prefix `cui`; ciderui-zola, and no docs theme; the palette fixed; `cui-btn`'s default left to ciderui's recommendation; the look iOS 18's, not iOS 26's.
- Guessed: that apple.com's look holds still long enough for its measurements to stay the standard; that a tap opens the menu on an iPhone as a script's click does on the simulator; that iOS 27 on a phone behaves as the iOS 26.5 simulator; that Firefox behaves as the others.
