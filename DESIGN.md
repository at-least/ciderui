# ciderui: its design

ciderui is apple.com's look for websites: a Tailwind CSS v4 theme on a small token contract, apple.com's parts as CSS classes (`cui-btn`, `cui-field`, `cui-tile`, …), and a Go module that writes their markup as gomponents nodes and carries the stylesheets; ciderui-zola, a Zola theme in a repository of its own, is built on it. It is vpkit's architecture (`../vpkit`, `../vpkit-zola`) with apple.com where vpkit has VitePress. This holds what is decided and why: who it is for, the standard, what a component is, the rules every one follows, the tokens, the type, the page, the catalogue, the Go module, the documentation, how a component is proven, the browsers, and the decisions and assumptions behind them. What a value is, and where it was read, is in the file that holds it (a stylesheet's header, tokens.css, a test); this does not copy it.

*Measured* means read off a live page or a browser, recorded in `reference/` with its probe, page, browser and date; *stated* is the owner's; *guessed* is nobody's measurement.

## Who it is for

The owner's websites, and nobody else's: ciderui is not public (stated). The sites are in English (stated), and they are both kinds of apple.com: marketing pages (a product, its story, its tiles) and the store (product cards, options to choose, a bag, sign-in and other forms) (stated), and crashcart's console (decision 19). They render their pages in Go with gomponents and build their CSS with the Tailwind CLI (decision 26). Their static sites, and ciderui's own documentation, are built with ciderui-zola, as vpkit's are with vpkit-zola (stated: "也做一個 ciderui-zola"). ciderui's documentation reads as Apple's developer documentation (decision 23). With no consumer outside, a component takes the best shape and the sites follow it (decision 1).

## The standard

**apple.com, measured, for how a part looks** (decision 2). apple.com's own pages are the original, and they can be read: `reference/apple.com/` holds the computed styles of its pages (listed in its README) at 1600, 1200, 900 and 390px. apple.com's CSS is neither licensed nor versioned (its files carry cache hashes), so none of it is copied here and there is no upstream to re-sync: what is recorded is what elements compute to, dated, and a re-measure is a deliberate act that changes the record and the checks that read it.

**The HIG and the quietest page, for whether a part or an effect is there** (decision 25). Agents write the markup; people read the page, and nothing on it competes with the content for their attention. Where apple.com draws an effect that only draws the eye, the effect is left out, or kept to what the HIG asks of it: feedback for an action, a change of state made clear.

For what apple.com's pages are not, Apple's other sites are the standard, each read as apple.com's are:

- **The console** (decision 19): App Store Connect's newer look, its analytics pages' (stated, 2026-10-24: 「我們因為 ciderui沒有而自己做的元件很醜很不apple style，應該在ciderui新增這些元件，然後真的找apple的網站去量」), read in the owner's browser with their account (`reference/asc/`, light only, at 1920px; stated, 2026-10-23: 「App Store Connect 要量，可以用我的帳號」). That look is a design system of tokens its pages declare (`--tb-…`), recorded as its page computes them (`reference/asc/results/tokens.json`), with what its pages draw (`analytics*`, `hover-*`). Beside it, also read in the owner's browser: Apple Account (`reference/apple-account/`), Feedback Assistant (`reference/feedback-assistant/`), the Apple Developer account (`reference/developer-account/`, its appearance set light), Apple Developer Forums' post lists (`reference/developer-forums/`, signed out, 2026-10-28), and Apple Developer's System Status and releases list (`reference/apple.com/`'s `system-status*`, `releases`). The console is light (stated, 2026-10-23: 「淺色」), its dark derived (decision 18). CloudKit Console is left out: its classes are App Store Connect's toolkit's, and its paths carry the team's ID, which the reader does not mask.
- **The documentation** (decision 23): developer.apple.com's documentation (DocC), read as a page at its own steps (`reference/apple.com/`'s `docc*`). support.apple.com's articles (`support*`) are measured and used by no part.

What the measurements say about apple.com, and so about ciderui:

- **Dark is a section's.** None of the stylesheets its pages load reads `prefers-color-scheme`; a `.theme-dark` section is dark on a light page, a product's page dark as a whole. ciderui follows: dark is `color-scheme: dark` on any element (rule 11).
- **The type is the look, and it is responsive per class.** One class steps its size at a breakpoint by itself (`t-headline` is 48px from 735px, 32px below). Each size has one line height and one letter-spacing.
- **apple.com tunes page by page.** The marketing pages' `typography-*` classes are defined by each page for itself, and the same size is tuned two ways (32px at +0.002em on the home page, +0.004em on the store). ciderui curates one system from them (the type, below).
- **The palette is small**: a dozen colors and their alphas.
- **The breakpoints are apple.com's own**: 735px and 1069px for the content, 834px for the global nav, measured by stepping a headline and the nav across each boundary.

## What a component is

A component is a class, or a block of a few classes, that gives an element the look of one of apple.com's parts (or the console's, or DocC's), and nothing else: no JavaScript, no page layout, no behavior beyond what CSS reads off the element's own state. The page opens the dialog and checks the option; the component shows the state the page set. Behavior comes from HTML's own elements where it exists: `<dialog>` for a sheet, `<details>` for an accordion and a menu, `popover` for a flyout, scroll snapping for a gallery. Where its markup is more than a class on one element, the Go module writes it.

## Rules

Every component follows these. The numbers are cited across the repository; they do not change.

1. **One component, one directory, in its set.** `components/<set>/<name>/<name>.css`, the set `site` (apple.com's parts), `console` (`cui-console-*`, decision 19) or `docs` (`cui-docs-*`, decision 23), each set's entry (`site.css`, `console.css`, `docs.css`) importing every part of it by name order, in Tailwind's `components` layer, so a utility on the same element always wins (`class="cui-btn px-8"`); `examples/<id>.html`, the canonical markup, which the Go module's tests render and the cases measure; `<name>.test.mjs`, its proof; `golden.json`, its recorded values. Nothing is unlayered; nothing uses `!important`.

2. **The class says what it is; what varies is an attribute.** The block is `cui-<name>`, a part `cui-<name>-<part>`. A mutually exclusive axis is `data-variant` (`"secondary"`, `"neutral"`), `data-size` (`sm`, `lg`, `xl`) or `data-color`; an orthogonal switch is a class; a state is what the markup carries anyway: an ARIA attribute (`aria-current`, `aria-invalid`, `aria-expanded`), a control's own state (`checked`, `disabled`, read as `:checked` and `:disabled`), `[open]`, or `data-state` where none says it. A variant's rule is `.cui-x[data-variant="y"]`, so a state one attribute more wins over every variant whatever the order.

3. **Component-local variables.** The block sets its own `--cui-<name>-*` from the contract; a variant overrides the locals and nothing else, so a component nested in another keeps its colors and recoloring one block is one rule on it.

4. **Rules for what is inside sit behind `:where()`**, adding no specificity, so a component placed inside another keeps its own look.

5. **Hover is `@media (hover: hover) { &:hover }`**, so a tapped element does not keep its hover. Here ciderui departs from apple.com, whose hover stays on a touch screen nearly everywhere (`touch-*:*`, an emulated one; tapped, `tap-*:*`), its store's card's lift alone off (`html.no-touch`, its script's): each component's touch case reads its hover against its rest, its why naming apple.com's. A tapped element matches `:hover` after the tap: a hover rule of its own shows, one behind the query does not, in Safari on the iOS simulator (XCUITest's taps) and in Chrome and Chromium (the DevTools protocol's touch events) (`reference/tap/`).

6. **Size and placement are the page's**: no width, no margin, no position, with three kinds of exception, each named in its stylesheet's header: the type, whose classes carry their small-screen step; the sizes apple.com fixes (a control's height, a field's 56px, a tile's minimum height); and a containing block a component's own parts are placed in, which places nothing itself (the global nav's `position: relative`, for its flyouts, which a site's `sticky` replaces, and its `z-index`, apple.com's 9999). Utilities remove them.

7. **Tokens: only the contract's.** A component reads tokens.css's colors and its non-color tokens and no literal color; the token test holds every stylesheet to it.

8. **apple.com's control sizes, not one height.** apple.com's pill comes in four heights, measured: 24px (`data-size="sm"`), 36px (the default: the store's shared button styles' and most of its buttons), 44px (`lg`), 56px (`xl`); its marketing pages' tile buttons step at 735px, 36px below and 44px from it (`data-size="cta"`, a hero's call to action). A field is 56px, the store's form button 56px (decision 7).

9. **One focus ring.** `base.css` draws `:focus-visible` as 2px of the accent, 3px out, on every element, apple.com's on every control it was read on; a text field shows its 2px accent edge in its place (apple.com's field draws no ring for a click, `retail:field.input`, and its script's ring for the keyboard alone, which CSS cannot tell apart on a text field). A component draws no ring of its own; it restates the ring only to move it or, a field, to set it aside.

10. **Motion is apple.com's, where decision 25 keeps it, through the tokens.** No button transitions its colors (apple.com's 0s, every button measured). A part that moves moves at apple.com's durations and curves, measured from its recorded animations (`home:menu.open`, `flyout:nav`), its stylesheet's header naming them. `base.css` neutralizes every transition and animation under `prefers-reduced-motion`.

11. **Dark comes through the tokens, set on any element.** Every color is a `light-dark()` pair declared once on `:root`; `cui-dark` on an element sets `color-scheme: dark` and every pair inside it resolves dark, `cui-light` inside it light again (measured in Safari 27.0.1, Chrome 155, Chromium 151 and the iOS 26.5 simulator, `reference/scheme/`). On `<html>` it darkens the whole page. apple.com's dark product page also clears its global nav over the hero (`product:nav`), which is the global nav's own treatment, not the scheme's. There is no automatic dark mode: apple.com has none.

12. **Icons are inline SVG**, from the Go module's `icon` package, in `currentColor`. apple.com draws its chevrons and plus signs with its own icon font (`SF Pro Icons`), which is Apple's; ciderui draws them as SVG from the numbers of apple.com's glyphs (`reference/icons/`), so a page needs no font and no `data:` image for them (vpkit's decision 13). Every icon apple.com draws none of is Lucide's (decision 17), and never one in a component's place where apple.com draws its own.

13. **Proof.** Each check in a component's test names its source (test/harness.mjs's keys): `apple` (equal to what `reference/apple.com/results/` records for the probe it names, at the width it names: apple.com's pages and developer.apple.com's), `asc` (equal to what App Store Connect's page drew, `reference/asc/results/`, at every width), `tb` (equal to the App Store Connect toolkit's token it names, as the page computes it beside the element), `rest` (a touch screen's hovered or pressed element equal to itself at rest) or `golden` (equal to `golden.json`: a part with no instance to measure, or a departure, with the check's why). No list of known differences: every departure is a golden check with its reason beside it.

14. **Ships complete.** The stylesheet with its header (what it is, from which page and probe, the markup, the attributes and parts, what is Apple's and what is ciderui's), its examples, its test with its golden values, its page in the documentation, and its Go component where its markup is more than one element. One commit per component, and one per step of any change, each with `npm test` green.

## The token contract

`tokens.css` is the contract: the colors every component reads, each one `light-dark()` pair, light first, or one value where both sides are the same; what each is for is its header's. `test/tokens.mjs` holds each value to the probe it was read from, every stylesheet to the contract alone, and each pair's contrast on every ground it is drawn on: text 4.5:1, a mark or a control's edge 3:1 (WCAG 2), a translucent color counted as what it draws over the ground under it.

- **apple.com's slots, `--cui-c-*`**: apple.com's values, measured, both sides. Where one falls short, it is fitted (decision 6): the dark secondary text is `#86868b` (apple.com's `#6e6e73` is 4.1:1 on black), the dark error Apple's dark red `#ff453a` (apple.com's `#e30000` is 4.3:1). A slot enters with the first component that reads it, its dark side measured then, so no stylesheet reads a guess. A field, a modal and the footer are light on a dark page too, as apple.com's are: their slots have one value.
- **The console's slots, `--cui-c-console-*`**: App Store Connect's toolkit's color roles, one slot per role a part reads, named by the role (`--cui-c-console-neutral-label-primary` is `--tb-color-role-neutral-label-primary`). Light is the role as its page computes it; dark is derived (decision 18); short ones fitted (decision 21). A console part reads these alone (the dialog's backdrop, `--cui-c-scrim`, is the one exception: the toolkit names none).
- **The docs set's slots, `--cui-c-docs-*`**: DocC's own colors, where they are not one of apple.com's slots on both sides (decision 23), measured light and dark.

The palette is fixed: no site redefines a color, and there are no themes (decision 15).

Not themed, in tokens.css too, each measured: the radii, the control sizes (rule 8), the focus ring (rule 9), the durations and curves (rule 10), the materials (the bar's `saturate(1.8) blur(20px)`, a curtain's `blur(20px)`), the shadows (apple.com's card's, the toolkit's), the content widths (apple.com's newer pages' 87.5% up to 1260px; the older 980px, its footer's; the global nav's 1024px; the Newsroom's article columns), the font stacks and the letter-spacing per size (the type).

## The type

**The stack** is `-apple-system, BlinkMacSystemFont, Inter, "Helvetica Neue", Helvetica, Arial, sans-serif`. SF cannot be shipped; on Apple's platforms the first two names are SF (`-apple-system` in Safari, `BlinkMacSystemFont` in Chrome; not `system-ui`: it is a font everywhere, so Inter after it would never be drawn). Elsewhere neither name is a font and the text is Inter, the kit's (`fonts.css`, Inter 4.001 with its `opsz` axis), fetched only there: in Safari 27.0.1, Chrome 155 and iOS 26.5's Safari the stack draws the system font and not one of Inter's faces is fetched; in Chromium on Linux it draws Inter with its Latin face alone (measured, `reference/stack/`; decisions 3, 4). Code's stack, `--cui-font-mono`, is SF Mono on Apple's platforms and the platform's own elsewhere, no font shipped.

**The ladder** is the store's `t-` classes (measured: one coherent ladder, where the marketing pages define theirs page by page), named `cui-<name>` (`cui-headline`, `cui-eyebrow`, `cui-body`, `cui-caption`, …), each with one step at 735px; `cui-body` is ciderui's (the store defines no `t-body`). The marketing pages' three-tier headlines (32, 48 and 56px) are the tile's, built with it (decision 10). The classes and their values are `type.css`'s; `test/type.mjs` holds each to the store's at 390, 734, 735, 900, 1200 and 1600px.

**The letter-spacing** is per size, and it is not apple.com's numbers: apple.com sets them on its SF Pro webfonts, which carry no tracking, while the system font carries Apple's size-specific tracking itself, the same in Safari and Chrome on macOS and in Safari on iOS; apple.com's values on the system font would track it twice. ciderui's are apple.com's minus the system font's own, measured size by size (`reference/tracking/cui-tracking.json`, the values in tokens.css). Where apple.com tunes a size two ways, the store's is taken. CSS has no letter-spacing per font of a stack, so Inter takes SF's values and runs wide of apple.com's text (decision 3).

**The classes are Tailwind utilities** (`type.css`, `@utility`), not component classes: measured in Tailwind 4.3.3 and Chromium 151, a class holds its own step at 735px, a variant composes with it (`cui-headline lg:cui-headline-super`), and a plain utility on the same element wins at every width (`cui-headline text-[20px]`), as rule 1 wants of a component.

## The page

**Breakpoints.** theme.css replaces Tailwind's with apple.com's, in pixels as apple.com's media queries are: `md` 735px, `lg` 1069px, `xl` 1441px (apple.com's stylesheets name 1441px; nothing measured changes there) (decision 9). A component's step and a page's `md:` switch at the same width. A set or a part may step at its own widths inside itself: the global nav at 834px, the console at 640px and 1024px (decision 20), the docs set at DocC's (decision 24).

**Sections.** A section is the page's ground (`bg`), the alternate (`bg-alt`) or dark (`cui-dark`, its alternate `#1d1d1f`, measured on the product page's highlights); its content is the newer width. Their padding is a section's own (apple.com's varies by section; measured per component).

## The catalogue

Every part, by set. Each row names the part; what it draws, its attributes and what is Apple's and what is ciderui's are its stylesheet's header (rule 14), its sources its test's (rule 13), its Go type in its set's package (the Go module, below).

The site set, apple.com's parts (`components/site/`):

| component | the part |
| --- | --- |
| `cui-accordion` | apple.com's FAQ, on `<details>` |
| `cui-article` | the Newsroom's story: its header, copy, figures and footnotes; markdown's elements as they come |
| `cui-btn` (`components/site/button/`) | the pill, its sizes and variants; 12px corners (`cui-btn-rect`) and the block (`cui-btn-block`) |
| `cui-card` | the store's card, the whole card its headline's link |
| `cui-checkbox` | the store's sign-in checkbox |
| `cui-field` | the store's 56px text field, its textarea (`cui-field-box`) and its dropdown |
| `cui-footer` | the footer, light on a dark page too |
| `cui-footnote` | a footnote's mark |
| `cui-gallery` | a product page's gallery, its paddles and dots |
| `cui-globalnav` | the global nav: the bar, its menu below 834px, its search, its bag and its links' flyouts, with no script |
| `cui-link` | a link in running text, and the more link (`data-variant="more"`) |
| `cui-localnav` | a product page's local nav |
| `cui-modal` | the store's overlay and the product pages' modal (`data-variant="page"`), on `<dialog>` |
| `cui-option` | an option on a radio, the buy page's |
| `cui-posts` | the Newsroom's archive, and the related stories under an article (`data-size="sm"`) |
| `cui-price` | the buy page's price |
| `cui-productnav` | the store's shelf of small product cards |
| `cui-quicklinks` | the store's row of pill links under a title |
| `cui-ribbon` | the message above a page's content |
| `cui-section` | a band: its ground, its rhythm, its content's width, its header |
| `cui-shelf` | the store's row of cards across the window |
| `cui-stat` | a stat, a value between its captions |
| `cui-stories` | the Newsroom's index |
| `cui-swatch` | a color swatch on a radio, and a row of them (`cui-swatches`) |
| `cui-tile` | a home page tile, promo or full width |

The console set, App Store Connect's newer look's parts (`components/console/`):

| component | the part |
| --- | --- |
| `cui-console-appearance` | Light, Dark and Auto as one segmented control: DocC's footer control, for a site that offers the choice (the HIG's Dark Mode advises against one; crashcart's, the owner's call) |
| `cui-console-badge` | a short word on a tinted capsule |
| `cui-console-banner` | a message on a rounded, tinted panel |
| `cui-console-bar` | the band across the top of a console's page |
| `cui-console-btn` | the button, a class on the page's own `<button>` or `<a>` |
| `cui-console-card` | the analytics card: a metric, or a section's list or chart |
| `cui-console-chart` | what is around an analytics chart's data, and a sparkline |
| `cui-console-checkbox` | a checkbox |
| `cui-console-chip` | a key and its value; a filter's chip with its remove control |
| `cui-console-code` | inline code |
| `cui-console-content` | the page's column, with a `cui-console-sidebar` beside it from 1024px |
| `cui-console-copy` | a value to copy in a listing with its copy button, and a copy button in a line of text |
| `cui-console-dialog` | a `<dialog>` the page opens with `showModal()` |
| `cui-console-disclosure` | a `<details>` hiding what is not needed yet, the HIG's disclosure controls |
| `cui-console-dl` | an item's facts, labels beside their values |
| `cui-console-empty` | what a list or a chart shows with nothing in it |
| `cui-console-field` | a text field, a dropdown, a multi-line field, and a form field's label, help and error |
| `cui-console-file` | the native file input drawn as a text field |
| `cui-console-head` | what a page belongs to, by name, with its tabs |
| `cui-console-hint` | an info button opening a popover |
| `cui-console-link` | a link |
| `cui-console-list` | a grouped list |
| `cui-console-listing` | lines of code or of a log on a gray panel, DocC's listing at the console's size |
| `cui-console-menu` | the list of a section's pages down its left side |
| `cui-console-page` | the page's ground and text, and the console's type as utilities |
| `cui-console-pager` | Apple Developer Forums' pager under a long list |
| `cui-console-popover` | a small panel its button opens |
| `cui-console-progress` | the native `<progress>` as a bar |
| `cui-console-radio` | a radio |
| `cui-console-section-head` | a page's or a section's title with its controls |
| `cui-console-sidebar` | the app's sidebar, a sheet below 1024px |
| `cui-console-skip` | the skip link |
| `cui-console-status` | a status: a colored circle before its word |
| `cui-console-table` | the table, a class on the page's own `<table>` |
| `cui-console-tabs` | the tab bar, underlined (a page's sections) or filled (the HIG's segmented control) |
| `cui-console-thumb` | a screenshot shown small, a link to itself |
| `cui-console-toast` | a word that the page did what was asked (the toolkit has no toast) |

The docs set, Apple's developer documentation's parts (DocC on developer.apple.com, `components/docs/`), stepping at DocC's widths (decision 24):

| component | the part |
| --- | --- |
| `cui-docs-aside` | an aside: note, important, warning, tip, deprecated, experiment |
| `cui-docs-bar` | the documentation's bar |
| `cui-docs-content` | the article's prose, as markdown writes it |
| `cui-docs-copy` | a code listing with DocC's copy button |
| `cui-docs-hero` | the page's title block and breadcrumbs |
| `cui-docs-navigator` | the navigator and its filter |
| `cui-docs-page` | the page: the bar, the navigator's column (a sheet below 1024px), the content's columns |
| `cui-docs-topics` | See Also and Topics |

Not built: DocC's footer appearance control in the docs set, which needs the page's script (decision 11): a docs page is light, or dark by `cui-dark` (rule 11). The console's `cui-console-appearance` is that control at the console's size.

Not measured: the store's checkout pages (my recommendation, which the owner took, 「都按照你的推薦」, 2026-10-21): they are not reachable without a bag, a headless browser's Add to Bag is refused (a 541), and getting past that would be getting past apple.com's bot protection; their form parts were read on the gift card's, the trade-in's and the feedback form's pages. So the global nav's bag is an empty bag's (`bag-flyout:*`): what it shows with something in it, and the count on the bag's icon, are not measured. Nor a phone's own screen: the taps are the iOS simulator's and Chromium's (`reference/tap/`); apple.com's pages are tapped in Chromium alone (`tap-*:*`). When a part is built, its probes are added where the reference's are not enough.

Not yet following decision 25 (read in their stylesheets, 2026-11-02): the card's lift (a deeper shadow and 1% larger, hovered), the shelf's paddle at an end shrinking to 0.8 as it fades out, and the global nav's rows and flyout rows coming in one after another, 20ms apart.

## The Go module

`go.mod` at the root is the module `github.com/at-least/ciderui`, for a site that renders with gomponents (decision 26). How a site uses it is the documentation's `docs/content/go.md`.

- **The components**, in their set's package (`site`, `console`, `docs`, under `go/`), for the parts whose markup is more than a class on one element; a button, a link in running text, a price, a console table is a class string on the page's own element. A type is named by its class without `cui-`, in CamelCase (`cui-card` is `site.Card`, `cui-console-badge` `console.ConsoleBadge`, `cui-docs-bar` `docs.DocsBar`): the set's prefix stays though the package names it, so one grep finds the class, the type and a site's uses. Where a part writes one piece of a larger block the type names that piece (`site.OptionTile`, `site.MoreLink`, `console.ConsoleMetric`, `console.ConsoleCheckboxBox`, `console.ConsoleListItem`, `console.ConsoleCopyIcons`).
- **Each component is a struct a call writes every field of**: `go tool exhaustruct` fails a literal of the module's types that leaves one out, in this repository's gate and in a site's, so a field added stops every call that does not set it. A variant is a struct of one unexported field, its zero value the default (the variant writing no attribute); a choice with data is a sealed interface, nil for none (`site.LinkedSearch` or `site.FlyoutSearch`); an absent text is `""`, an absent node nil. A precondition the type cannot hold is a panic naming the package's type.
- **The markup and the CSS cannot part**: each component's test renders its examples and holds the output equal to `examples/<id>.html`, its lines trimmed and joined (`go/internal/example`).
- **`icon`**: ciderui's icons (rule 12). **`lucide`**: every Lucide icon, named by Lucide's name in Go's capitals (`shopping-bag` is `lucide.ShoppingBag`), 1em square (decision 17). `scripts/build-lucide.mjs` writes `go/lucide/lucide.go` from lucide-static (package.json's exact version; its license `LICENSE-Lucide`), each icon's elements and its `<svg>`'s attributes held to its file there; `npm test` fails while the file differs from what the script writes.
- **The stylesheets**: the root package embeds every one; `ciderui.WriteCSS(dir)`, or `go run github.com/at-least/ciderui/go/cmd/ciderui-css <dir>`, writes them out in the kit's layout for a Tailwind build to import, so a site's dependency at a commit is the one pin for its markup and its CSS (decision 13). The font files are not embedded: `fonts.css` reads them beside the compiled stylesheet.
- **`ciderui.CheckMarkup`**: reads a rendered page and lists each ciderui class, and each `data-variant`, `data-size` or `data-color` value, that no stylesheet's rule selects, which would otherwise draw the default with no error; a site's tests call it on every page they render.

## The documentation

`docs/` is ciderui's documentation, a site built with ciderui-zola (decision 14): getting started, the foundations, the Go module and a page for each component with each of its examples live in a frame of its own over its markup, in Apple's developer documentation's look (decision 23: ciderui-zola's `docs.html` and `docs-section.html` on the docs set), and the home and the news in apple.com's and the Newsroom's (`docs/README.md`).

`npm test` holds it to the kit: the examples' copies and the frames' stylesheet are this commit's (`scripts/build-docs.mjs --check`); `test/docs.mjs` wants a page for every component showing every one of its examples, the frames' stylesheet holding every rule the examples' classes have, and, the site built with zola, every page at 1440 and 390px in Chromium with no error, no missing file or dead link and no sideways scroll, each frame as tall as its content with every picture loaded. Its pictures are ciderui drawn by itself (`scripts/docs-pictures.mjs`), held to what they were drawn from (`docs/pictures.json`, each scene's and the stylesheet's digest), not to Chromium's bytes.

Outside Chromium it is measured in Safari 27.0.1 on macOS, Safari 26.5 on the iOS simulator and Chrome 156 (`reference/docs/`: every page and frame, the gallery's and the shelf's paddles, the menus, the search, the flyouts and the bag in their frames), before the docs layout; the docs layout (its navigator's column and sheet, a section's chevron, a heading's link under the bar, the column's foot at the window's, the filter, a code block's copy button) in Chromium 151, Firefox 153 and Chrome 156 on Linux, in Docker, and in Safari 27.0 on macOS in a `tester` VM (2026-11-01); not yet on the iOS simulator.

## How a component is proven

`npm test` is the gate (CLAUDE.md, "This repository's gate"). `test/components.mjs` renders each case in headless Chromium, in test/Dockerfile's image, and compares computed styles by the check's source (rule 13). What the measurements taught the harness, which it keeps:

- **Linux's hinting.** Chromium on Linux hints every glyph to a whole-pixel advance unless launched with `--font-render-hinting=none`; with it, its text widths are the Mac's within 0.00006em a character (measured, `reference/tracking/`). Every run passes the flag.
- **The fonts differ between the two sides.** apple.com's values were read with its SF Pro webfonts; ciderui's cases render on Linux in Inter, fonts.css's, inlined into each case's page. So an `apple` check compares what does not depend on the text's width (the type's own values, colors, borders, radii, padding, heights where the line height fixes them), and a box that depends on the text's width is a `golden` check.
- **Transitions.** A state is forced with transitions off, and styles are flushed before they return, or a later read lands in flight.
- **States the protocol cannot force.** `:checked` is not a pseudo-class DevTools forces, so a checked state is proven by an example with `checked` in its markup, against the probe that measured apple.com's classes on a checked copy (`buy:option.checked`); such a probe says so in its name and its spec.
- **What only Safari or iOS can answer** (the system font's tracking, a section's color scheme, the stack's fetching, the nav's popover menu) is measured on macOS (`reference/browsers.mjs`'s `safari`, `ios`, `chrome`), not assumed from Chromium: Safari in a `tester` VM of its own (`reference/vm.mjs`), iOS's simulator and Chrome on a Mac with Xcode.
- **States a click makes.** `:popover-open` is not forced either: a case clicks the element that makes it (`click`), then moves the pointer off, which a click leaves over what it clicked, hovered; it presses keys (`press`: Tab, Escape) and reads where the focus is (`focused`); `wait: 0` reads its motion's start (`@starting-style` through the compiled stylesheet).
- **A scroller's snapping.** `scroll` scrolls a case's scroller to a place before its checks, at every width: between two snap positions it comes to rest at one, at its end its last item as far in as its first.
- **What the pointer finds.** A link laid over a box (a card's `::after`) has no box of its own a check can read: `hit` reads what the pointer finds at five points of the box, a tenth in from each corner and its middle.

The token test holds the contract and its contrast; `reference/touch/audit.mjs` holds every control's touch target; `test/go.mjs` holds the Go module to gofmt, go vet, exhaustruct and its tests, its examples included.

## Browsers

The kit needs `light-dark()` (vpkit's floor: Chrome 123, Safari 17.5, Firefox 120, stated there). Beyond it, each measured in the browsers below:

- the global nav's menu: `popover`, `:has()` and `@starting-style` (`reference/popover/`); its search: anchor positioning and `backdrop-filter` (`reference/search/`); its flyouts: a transition's delay on `:hover` for the pointer's rest, and the bar's blur off under their curtain (`reference/flyout/`: by the pointer in Safari, Chrome, Chromium and Firefox, by XCUITest's taps on the iPad and iPhone simulators); its `user-select: none`: Safari's `-webkit-user-select`, its field still taking text (`reference/select/`, by XCUITest's keyboard on the iPhone simulator);
- the modal: `<dialog>` and the `command` and `commandfor` attributes, to open with no script (`reference/dialog/`; a browser without them opens it by the page's `showModal()`);
- a tap opens the menu and leaves no hover behind the query (`reference/tap/`, on a touch screen: not Safari on the Mac);
- forced colors (a high-contrast theme): the swatch, the gallery's dots and an empty select still show their state and colors, by rules of the kit's own (`reference/forced/`: Chromium's emulation, light and dark, and Firefox's override; WebKit has no forced colors), and so do the console's chosen radio, current tab, page and selected row, its rules and hairlines and the sidebar's sheet (their tests' forced-colors cases).

ciderui's claims are measured in Safari 27.0.1 on macOS 27.0.1, Safari 26.5 on the iOS 26.5 simulator (an iPhone's; an iPad's too for the flyouts), Chrome 155 on macOS (156 for the flyouts and the documentation), Chromium 151 on Linux, and Playwright's Firefox 153 on Linux (2026-10-22: each claim holds, the accordion opening at once, as Firefox has no `interpolate-size`; the tracking and the stack Inter's, Linux having no SF). Not in Firefox on macOS (Playwright's does not run on the Mac over SSH; the body's `-moz-osx-font-smoothing: grayscale`, apple.com's, is for it, unmeasured), not on a real iPhone or iPad, not on Windows or Android.

The components' values are read in Chromium (`test/`). Read once in Firefox 153 (2026-10-22, by a harness outside the repository that puts each state on by real input), they differed in how a value reads, not in what is drawn: an SVG path's box (`getBoundingClientRect()` holds its stroke in Firefox), a text field's line-height (Firefox reports at least its font's `normal`; the line-height of a field of fixed height moves its text in neither), a grid's tracks (serialized evenly; the items' boxes within 0.011px); and in one thing a user meets: after a click on an empty part of the local nav's bar, Firefox starts Tab from the caret the click left in the title, past the title's link, where Chromium starts at the title (apple.com's local nav, read 2026-10-21, would do the same).

## Decisions

Numbered as cited across the repository; a new one takes the next number.

1. **The sites do not bind a component's shape.** Every consumer is the owner's; no compatibility is kept.
2. **apple.com, measured, is the standard for how a part looks.** Its live pages' computed values, recorded with their probes and dates; its CSS never copied. Whether a part or an effect is there is decision 25's.
3. **SF exact, Inter near.** The letter-spacing is fitted to the system font (within 0.0006em a character of apple.com's text at all 14 sizes in Safari and Chrome on macOS and Safari on iOS); Inter, where the system font is not SF, takes the same values and runs wide, by up to 0.0243em a character at 400 (21px), about 21px over a 60-character line (measured, `reference/tracking/`) (stated: the owner chose Inter over the platform's own font and over apple.com's Helvetica and Arial).
4. **The stack names the system font by its two Apple-only names**, so Inter is drawn, and fetched, only where SF is not.
5. **Dark is set on any element**, by `color-scheme`; no automatic dark mode.
6. **The palette is apple.com's, fitted where it falls short**: the dark secondary text is `#86868b`, apple.com's dark `#6e6e73` being below 4.5:1, and the dark error Apple's dark red, `#ff453a`, apple.com's `#e30000` being 4.3:1 on black; the token test holds the minimums.
7. **apple.com's control sizes** (24, 36, 44, 56px) where vpkit has one height; the default, 36px, is ciderui's recommendation, the owner deferring to it.
8. **One focus ring**, apple.com's: 2px of the accent, 3px out.
9. **apple.com's breakpoints**, in pixels, replace Tailwind's.
10. **The store's ladder is the type scale**, apple.com's one coherent ladder; the marketing pages' three-tier headlines are a tile's.
11. **No JavaScript**: HTML's own elements for behavior. A part that cannot work without a script (the gallery's and the shelf's paddles and dots, the navigator's filter, a copy button) shows the state the page's script sets, the script the page's (ciderui-zola ships the documentation's).
12. **Icons are inline SVG**, no icon font.
13. **The Go module carries the CSS**, one pin for a site's markup and its stylesheets (vpkit's decision 8).
14. **One product and a theme**: the kit at the root; ciderui-zola, the Zola theme built on it, in a repository of its own, beside this one (stated, 2026-10-19, `../ciderui-zola/DESIGN.md`), which ciderui's documentation is built with. No docs theme, where vpkit has `docs-theme/` (stated, 2026-10-19: 「docs-theme 不做」): Apple's documentation is a set of the kit's own (decision 23), and ciderui-zola's docs layout is built on it. ciderui-zola's first site is ciderui's documentation, with news in the Newsroom's look (`cui-article`, `cui-stories`, `cui-posts`).
15. **The palette is fixed** (stated: "顏色不能換"): no site color, no themes; apple.com has one palette.
16. **apple.com's look as it is, not iOS 26's** (stated, 2026-10-19: asked about iOS 26's style without its glass, 「維持 ios 18」): the bar translucent under a blur and the pill, as apple.com shows them measured; no Liquid Glass and no iOS 26 shapes.
17. **Lucide for every icon apple.com draws none of** (stated, 2026-10-23: 「ciderui 應該支援 https://lucide.dev/」; how, my recommendation): all of it, so a site finds any icon by its name on lucide.dev and draws it with no copy of its own; each name once, its aliases left out (one name per thing); 1em square where Lucide's files are 24px, so an icon takes its text's size, as ciderui's own icons in text are sized (`icon.More`, `icon.Alert`), a utility sizing it otherwise; its place the page's (rule 6): in a line of text the page adds `inline`, Tailwind's preflight making an `<svg>` a block; its attributes on its `<svg>`, so a utility there (`stroke-[1.5]`) reaches its strokes. All of it costs a site nothing: an icon it does not use is not in its binary (measured, below). Lucide's look is not apple.com's and is not held to it.
18. **A console's dark is derived** (stated, 2026-10-24, choosing it over a console light only: 「用 tb 深色色板推算」): App Store Connect ships its roles light only; each `--cui-c-console-*` slot's dark side is its role's step of the toolkit's dark palette (the light palette's step of the same lightness, chroma and hue), at the role's alpha, black and white each other's; a role that is no step (the theme's brand blue, `oklch(54.2% .162 251.9)`) takes the step of the ramp whose 500 step's hue is nearest its own, nearest its lightness (blue-600; not nearest in OKLab alone: that would be cyan-800, whose dark step is a pale 90% cyan). The values are Apple's; the mapping is ciderui's, and the one place a stylesheet reads what nobody measured; the token test holds the mapping and the contrast.
19. **The console's parts are one family, `cui-console-*`** (stated, 2026-10-24: 「統一 cui-console-*」), its colors `--cui-c-console-*`: one grep separates them from apple.com's parts, and apple.com's (`cui-btn`, `cui-field`) are untouched.
20. **The console steps at the toolkit's 640px and 1024px** (stated, 2026-10-24: 「tb 自己的 640/1024px」), its page scale's (desktop from 640px, mobile below), inside each part, as the global nav's 834px is its own.
21. **The console's colors are fitted to 4.5:1** (stated, 2026-10-24, asked whether to keep Apple's or fit them: 「調到 4.5:1」; decision 6 applied to the console): where a console color falls short of its job's contrast (text 4.5:1 on its ground, a label 4.5:1 on its fill, a field's edge 3:1 as a control's boundary), that side is the nearest step of the same ramp that reaches it, or the same step at another alpha. A part's text is held on a card's ground as well as the page's (a card is lighter than the page in dark). Each fitted side, its ramp and its reason are tokens.css's and `test/tokens.mjs`'s, which holds each to its ramp and its minimum on its grounds.
22. **The site and console sets** (stated, 2026-10-31: 「好，分兩套」, then 「Everything as first proposed」): apple.com's parts (its marketing pages', its store's, its Newsroom's, which share the global nav, the footer, the button and `--cui-c-*`) and the console's (decision 19), each its directory (`components/site/`, `components/console/`), its Go package (`site`, `console`; the names kept whole, the owner's pick over short ones) and its entry (`site.css`, `console.css`: every part of the set by name order, which `css_test.go` holds). The docs set is the third (decision 23).
23. **The docs set, `cui-docs-*`** (stated, 2026-11-01: asked whether ciderui should have a set for code docs, 「要，文件站改成 developer.apple.com 的樣子」): Apple's developer documentation's parts (DocC on developer.apple.com), its directory `components/docs/`, its Go package `docs`, its entry `docs.css`. Its colors are apple.com's slots where DocC's are the same values on both sides (its text, secondary text and links, the page's ground, a table's edge the divider: each measured equal light and dark) and `--cui-c-docs-*` for DocC's own. Its first site is ciderui's documentation, through ciderui-zola's docs layout: the documentation's pages in DocC's page; the home and the news stay the site set's. The DocC parts already in the console set (`cui-console-listing`, `cui-console-copy`, `cui-console-appearance`) stay there, DocC at the console's size; the docs set draws DocC at its own (its code 15px where the console's is 13px).
24. **The docs set steps at DocC's own widths** (its stylesheets' media queries, measured 2026-11-01: 735, 767, 1023, 1250 and 1499px, each read on both sides), inside each part (my recommendation: DocC's steps are what its parts are measured at; apple.com's 1069px is not one of them).
25. **The quietest layout, by the HIG** (stated, 2026-11-02: 「ciderui是為agent coding最佳化，可是設計出來的版面要最簡潔，不要用花俏的設計去搶使用者的注意力，所以即使是參考apple的設計，也要照著hig跟最簡潔來取捨」): agents write the markup, people read the page, and nothing on it competes with the content for their attention. apple.com, measured, stays the standard for how a part looks (decision 2), not for whether a part or an effect is there: where apple.com draws an effect that only draws the eye (an entrance animation, a lift on hover, a staggered reveal), the HIG and the simplest page decide, and the effect is left out or kept to what the HIG asks of it (feedback for an action, a change of state made clear). A component or a variant that exists only to call attention is not added. The parts that do not follow it yet are listed under the catalogue.
26. **Markup in Go, with gomponents** (stated, 2026-11-02: 「我希望都轉成go」, gomponents chosen over templ; then 「rust 清除掉就好」, the Rust crate removed): a component is a struct whose every field a call writes, exhaustruct holding it; plain Go, so gofmt, vet and gopls rename work on it, with no generate step (templ, tried on `site.Field`, dropped text after a raw node and needs a generate step whose stale output runs with no error).

## Assumptions

- **Measured**:
  - apple.com's values (`reference/apple.com/`, from 2026-10-19), App Store Connect's (`reference/asc/`), Apple's other sites' (`reference/apple-account/`, `reference/feedback-assistant/`, `reference/developer-account/`, `reference/developer-forums/`), DocC's page at its own steps (`reference/apple.com/`'s `docc-page*`, `docc-technote*`, 2026-11-01) and its navigator's width at a fresh load (400px up to 1300px wide, its script's).
  - The system font's tracking and its fix (`reference/tracking/`); the stack's fonts and fetching (`reference/stack/`); a section's color scheme (`reference/scheme/`); the palette's contrast, computed from the measured values.
  - The nav's popover menu and the accordion's `<details>` in each browser listed (`reference/popover/`, `reference/details/`); what a tap leaves and the menu opened by one (`reference/tap/`); the flyouts shown by the pointer's rest and opened by their buttons and taps (`reference/flyout/`); what of the nav can be selected and its field typed into (`reference/select/`).
  - The documentation built with ciderui-zola (`reference/docs/`; Safari on the Mac by its keyboard; on the iOS simulator no key is pressed, so the popover's Tab, Escape and focus-return claims are not run there).
  - The console's tabs' phone row and its edge fade, thumbnails, pager, page column and toast (`reference/console/`: Chromium 151 and Playwright's Firefox 153 on Linux, 2026-10-28; Safari 27.0 on macOS 27.0 in a `tester` VM, 2026-10-29; not yet in Chrome or on the iOS simulator).
  - Every control of every set 44pt at least each way on a touch screen, the HIG's default, hit-tested from each control's center (`reference/touch/`), desktop unchanged where apple.com and DocC draw smaller.
  - That a Lucide icon a site does not use is not in its binary (decision 17): a program rendering `lucide.Heart`, built with go1.27.3 on darwin/arm64 (2026-11-02), holds Heart's path and neither Anchor's nor Zap's.
- **Stated by the owner**: the sites are the owner's alone, in English, marketing and store, and crashcart's console; Go with gomponents and the Tailwind CLI; Inter where SF is not; the prefix `cui`; ciderui-zola, and no docs theme; the palette fixed; `cui-btn`'s default left to ciderui's recommendation; the look iOS 18's, not iOS 26's; Lucide supported; the console App Store Connect's, light, its dark derived and fitted to 4.5:1; ciderui's documentation in Apple's developer documentation's look; the quietest layout, by the HIG.
- **Guessed**: that the navigator's 400px, DocC's at a fresh load up to 1300px wide, serves every width from 1024px (DocC's grows to 543px on wider windows, by its script); that the toolkit's mobile page scale is its pages' below its 640px breakpoint (no page was read narrower than 1920px: decision 20); that a Lucide icon is wanted at its text's size more often than at Lucide's 24px (decision 17); that apple.com's look holds still long enough for its measurements to stay the standard; that iOS 27 on a phone, under a finger, behaves as the iOS 26.5 simulator under XCUITest's taps; that Firefox on macOS and Mozilla's own build behave as Playwright's Firefox on Linux.
