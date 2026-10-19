# apple.com, measured

What ciderui's tokens and components are held to: apple.com's computed styles, read off its live pages in a browser. Not its stylesheets: apple.com's CSS is Apple's, unlicensed and unversioned (its files carry cache hashes, not versions), so none of it is in this repository; what is recorded is what an element computes to (a color, a size, a radius), with the page, the element and the date beside it.

Measured 2026-10-19 in Playwright's Chromium 151 on Linux, with `--font-render-hinting=none` (`../browsers.mjs`; without it Linux hints every glyph to a whole pixel, `../tracking/`); apple.com draws its own SF Pro webfonts there (`../stack/`), so text-sized boxes are the Mac's.

## The pages

| spec | page | what |
| --- | --- | --- |
| `breakpoints` | https://www.apple.com/ | a headline and the global nav stepped across each boundary, 733 to 1442px |
| `home` | https://www.apple.com/ | the type its stylesheets define, the page, the global nav, the hero and its buttons, light and dark tiles, the gallery's dots, the footer |
| `product` | https://www.apple.com/iphone-18-pro/ | a dark page: the base type scale, the content's width, sections light, dark and alternate, a rounded tile, the paddles, the "more" link, a stat, the floating local nav |
| `store` | https://www.apple.com/store | the store's own type scale (`t-`), its shared buttons (`common-css@1.3.3`), the header, quick links, product nav, shelves, a card and its hover |
| `buy` | https://www.apple.com/shop/buy-iphone/iphone-18-pro | the option tiles (rest, focused, checked), the color swatches, a price |
| `bag` | https://www.apple.com/shop/bag | the empty bag, its form buttons, the ribbon, an accordion |
| `signin` | https://secure.store.apple.com/shop/account/home | Apple Account's sign-in widget (its iframe, idmsa.apple.com): a text field and its floating label, a checkbox, a link |
| `newsroom` | https://www.apple.com/newsroom/ | a light page under a dark global nav: the bar's dark material, its links and icons, its phone menu opened |
| `localnav` | https://www.apple.com/macbook-air/ | the local nav most product pages carry, light: the bar, its material and rule, the name, the links and the current page's mark, the Buy button, the phone's menu opened by its chevron; at 1600, 1200, 900, 834, 833 and 390px |
| `localnav-dark` | https://www.apple.com/macbook-pro/ | the same, dark |
| `ribbon` | https://www.apple.com/iphone/ | the ribbon, the message above the content: its ground, type, padding and width, its link, its motion (the animations and their keyframes), read once it has run at each width |

Each is read at 1600, 1200, 900 and 390px wide (one width inside each of apple.com's bands), `breakpoints` at the boundaries.

## What it showed

**Breakpoints** (`breakpoints`). The content changes at 734/735 and 1068/1069: a headline is 32px to 734px, 40px from 735px, 48px from 1069px; nothing it sets changes at 1440/1441. The global nav has its own: 48px tall to 833px, 44px from 834px.

**The palette** (every color the probes read; the probe beside it):

| color | value | where |
| --- | --- | --- |
| text | `#1d1d1f` | every page's body |
| secondary text | `#6e6e73` | `home:hero.callout`, `store:shelf.subheader` |
| tertiary, a control's edge | `#86868b` | `buy:option` border, `signin:field.input` border, `signin:checkbox` |
| page | `#ffffff` | the home page's and the bag's body (the store's home page is on `#f5f5f7`) |
| alternate ground | `#f5f5f7` | `home:hero.tile`, `home:footer`, `bag:ribbon` |
| divider | `#d2d2d7` | `bag:accordion.item` |
| accent fill | `#0071e3`, hover `#0076df`, pressed `#006edb` | every primary button |
| link | `#0066cc` | `home:button.secondary`, `bag:form-button.secondary` |
| neutral fill | `#1d1d1f`, hover `#272729`, pressed `#18181a` | `store:button.neutral` |
| field ground | `rgba(255,255,255,.8)` | `buy:option`, `signin:field.input`; the field's text `#494949` |
| dark page | `#000000` | `product:page.body`, `home:tile.dark` |
| dark alternate | `#1d1d1f` | `product:section.dark` (`.theme-dark.background-alt`) |
| dark text | `#f5f5f7` | `home:tile.dark-headline`, `product:section.dark` |
| dark link | `#2997ff` | `home:tile.dark-button-secondary`, `product:section.header-link` |
| tag | `#b64400` | `store:card.eyebrow` ("NEW", "PRE-ORDER NOW") |

apple.com has no dark mode: none of the 21 stylesheets these pages load reads `prefers-color-scheme`; dark is a section's (`.theme-dark`) or a page's.

**The type.** Two systems, each defined page by page:

- The marketing pages' `typography-*` classes, three tiers (to 734px, from 735px, from 1069px), each page defining its own subset and tuning it: `typography-headline` is 32/40/48px on the home page and not defined on the product page; `headline-elevated` is not defined on the home page and 40/48/64px on the product page. A class a page does not define reads as an unclassed paragraph does (`type.none`, the control every spec carries).
- The store's `t-*` classes, two tiers (to 734px, from 735px), one ladder: `headline-standalone` 96/100, `headline-super` 80/84, `headline-elevated` 64/68, `headline` 48/52, `headline-reduced` 40/44, `eyebrow-super` 32/36, `eyebrow-elevated` 28/32, `eyebrow` 24/28, `eyebrow-reduced` 21/25, `intro-elevated` 24/32 and `intro` 21/29 at 400, `body-tight` 17/21, `body-reduced` 14/20, `body-reduced-tight` 14/18, `caption` 12/16, `label` 24/28, `callout` 32/36, `tout` 19/23, `quote` 40/48, `manifesto` 32/40 (`store:type.*`).

Across both, a size has one line height and one letter-spacing, the cut following the size (SF Pro Text to 17px, SF Pro Display from 19px): 12/16 −0.01em, 14/20 (or 18) −0.016em, 17/25 (or 21) −0.022em, 19/23 +0.012em, 21/25 +0.011em, 24/28 +0.009em, 28/32 +0.007em, 32/36 +0.004em, 40/44 0, 48/52 −0.003em, 56/60 −0.005em, 64/68 −0.009em, 80/84 −0.015em, 96/100 −0.015em, at 600 from 19px for headlines and 400 for intros and body. Two sizes are tuned two ways: the home page sets 32px at +0.002em and 48px at −0.002em (`home:type.headline`, `home:type.callout`), the product page and the store +0.004em and −0.003em.

**The page's width.** The marketing pages' newer content is 87.5% of the window up to 1260px (`product:content`; 420px at most to 734px); the footer and the global nav keep the older 980px and 1024px (`home:footer.content`, `home:nav.content`), 22px from the edge (16px to 734px). The home page's promo grid is two columns from 735px with a 12px gap (`home:promo.section`).

**The global nav** (`home:nav*`, `home:menu*`). 44px (48px to 833px), `rgba(255,255,255,.8)` under `backdrop-filter: saturate(1.8) blur(20px)`; its content 1024px at most, 22px in; one row of links spread across it (`justify-content: space-between`, 8px out on each side), each 44px tall and 8px in, 12px text on a 12px line at 400, `rgba(0,0,0,.8)`, `#000` hovered, color in 0.32s `cubic-bezier(.4,0,.6,1)`; the logo, the search and the bag images 14, 15 and 14px wide, 44px tall. Fixed to the top on the home page; on the product page and the Newsroom placed at the top and scrolled away (`position: absolute`), the local nav the one that stays. An open menu's curtain is `rgba(232,232,237,.4)` under `blur(20px)`. On a dark page (`product:nav`) the bar is clear and its links `rgba(255,255,255,.8)`, `#fff` hovered; under the Newsroom's dark nav (`newsroom:nav`) the bar is `rgba(0,0,0,.8)` under the same blur, its links the same.

To 833px the bar holds the logo, the search and the bag, 48px squares of 17 by 48 images (`nav.*-compact`), and the menu's button, a 48px square at the right: two lines in 18 by 18, `#000` at `opacity: .8` (1 hovered; `#fff` on a dark nav), opacity and color in 0.32s. Opened (`menu.*`, the button clicked), the menu fills the window on `#fafafc` (`#161617` on a dark nav) and the page stops scrolling (`overflow: hidden` on the root and the body); its rows start 47px down, each 53px tall and 7px over the next (`margin-bottom: -7px`, a 46px pitch), its links 28px on a 28px line at 600, 3px 48px 4px in, `#333336` (`#e8e8ed` dark), `#000` hovered (`#fff`); the list 84px short of its end. Its motion, every animation the click starts recorded (`menu.open`'s `animations`): the ground from the bar's to the menu's and its height from the bar's to the window's in 450ms `cubic-bezier(.4,0,.6,1)`; the logo, search and bag out in 240ms (opacity, then `visibility: hidden` at the end, `steps(1)`); each row in from 8px up and transparent, in 240ms from 200ms, 20ms after the row above (200 to 400ms over eleven rows); the button's lines into an X in 0.24s (their SVG animation, `../icons/`).

**Buttons.** The pill, `border-radius: 980px`, a 1px transparent border:

| size | type | padding | height | where |
| --- | --- | --- | --- | --- |
| reduced | 12/16 | 3px 10px | 24px | `store:button.reduced` |
| default | 14/18 | 8px 15px | 36px | `store:button`, the marketing buttons to 734px |
| elevated | 17/20 | 11px 21px | 44px | `store:button.elevated`, the marketing buttons from 735px (`home:button`) |
| super | 17/20 | 17px 30px | 56px | `store:button.super` |

The fills are the accent's and the neutral's (above); a secondary button is the fill's color as its text and border on no ground, filled with the hovered fill when hovered (`home:button.secondary`, `store:button.secondary-neutral`); on a dark ground the secondary is the dark link `#2997ff` (`home:tile.dark-button-secondary`). The store's form buttons are rounded rectangles, `border-radius: 12px`, 17px text, 17px 30px, 56px tall (`bag:form-button`); `button-block` is the full width at `border-radius: 8px` (`store:button.block`). No button transitions its colors (0s).

**The focus ring** (`home:button`, `store:button`, `bag:form-button`, `buy:option.focus`): `outline: 2px solid #0071e3; outline-offset: 3px`; a text field's at `outline-offset: 1px` (`signin:field.input`).

**Fields** (`signin:field.*`). 56px tall, `border-radius: 12px`, a 1px `#86868b` border on `rgba(255,255,255,.8)`, 17px text, 16px in from the left; focused, a 2px `#0071e3` border (15px in, the box unchanged). The label is the placeholder at 17px `#6e6e73` until the field is focused, then 12px at the field's top (`top` 18px to 10px, in 0.125s). The checkbox is 16px, a 1px `#86868b` border, `border-radius: 3px`.

**Option tiles and swatches** (`buy:*`; the checked ones, `buy:option.checked` and `buy:swatch.checked*`, are apple.com's classes on a checked copy the probe adds beside the page's own, which is not checked until clicked). An option is a label on a radio: `border-radius: 12px`, a 1px `#86868b` border, 15px in, on `rgba(255,255,255,.8)`; checked, a 2px `#0071e3` border, 14px in, the box unchanged; focused, the focus ring. A color swatch is a 28px circle in a 36px ring, 2px of `#0071e3` when checked.

**Tiles and cards.** The home page's tiles are square-cornered, on `#f5f5f7` or `#000` (`home:hero.tile`, `home:tile.dark`), 692px tall from 1069px; the product page's rounded tiles `border-radius: 28px` (`product:tile.rounded`); the store's cards `border-radius: 18px` with `box-shadow: 2px 4px 12px rgba(0,0,0,.08)`, hovered `2px 4px 16px rgba(0,0,0,.16)` and `scale(1.01)` in 0.3s (`store:card`).

**The local nav** (`localnav:*`, `localnav-dark:*`). The bar most product pages carry under the global nav (of eleven product pages seen with a local nav on 2026-10-19, ten carry this one, MacBook Air's, MacBook Pro's, iPad Air's, iPad Pro's, AirPods Pro's, AirPods Max's, Apple Vision Pro's, Apple TV 4K's, iMac's and Mac Studio's; iPhone 18 Pro's is a newer floating one, below; the survey is not recorded): sticky at the top (`z-index: 9997`, under the global nav's 9999), 52px (48px to 833px: it steps with the global nav, not the content), `rgba(250,250,252,.8)` under `saturate(1.8) blur(20px)` (dark `rgba(22,22,23,.8)`; MacBook Pro's is clear at the top of its page until its script sees it stick), a 1px rule along its foot, `rgba(0,0,0,.16)` (dark `rgba(255,255,255,.24)`); its content 980px at most, 22px in (16px to 833px); the product's name 21px on 24px at 600 (19px to 833px), black at opacity .88 (white at .92); its links 12px on 22px, 24px apart, black at .8 (1 hovered, no transition), the current page's at 1 with a 1px line at .8 along the bar's foot under it; its Buy button the store's 24px pill (12px, 3px 10px), 24px after the links. To 833px the links fold into a menu: a 40 by 30 chevron (16 by 9, at .8) 4px into the 16px before the button opens it, a checkbox's label (apple.com's own way); opened, the bar and the menu on `#fafafc` (`#161617`), the rows 14px on a 34px row from 58px, at .8 (the current page's at 1, a 1.2 by 12px line 16px before it), the menu's foot 27px below the last, the chevron turned up; the ground and the menu's height in 320ms `cubic-bezier(.4,0,.6,1)`, each row in from 4px up and transparent, from 260ms, 20ms after the row above (`tray.open`'s animations).

**The ribbon** (`ribbon:*`, `bag:ribbon`). The message above a page's content: 16px above and below on `#f5f5f7`, 14px on 20px `#1d1d1f`, centered within 83.33% of the window, a more link in `#0066cc` after it (52px for one line, 72px for two at 390px). The iPhone page's drops in as the page shows it (`ribbon-drop`, `translateY(-100%)` to 0 in 0.8s ease-in-out) on the accent, `#0071e3` with `#f5f5f7` text and link, and 1.6s on turns to its ground in 0.8s (`ribbon-animate-*`, ease-in-out, `forwards`); its link hovered stops its animation, and its color falls back to the `#f5f5f7` it started from, on the `#f5f5f7` ground. The bag's is the ground alone, still. A new width runs the iPhone page's again, so the spec reads each width 3.5s after it (`settle`).

**Materials.** The global nav's (above); the product page's floating local nav, a bar `border-radius: 18px`, `rgba(18,18,18,.8)` under `blur(20px)` with a 1px `#282828` ring, 51.5px tall, its buttons 12px pills ringed in `#6e6e73` (`product:localnav.*`).

**Paddles and dots.** A gallery's paddles are circles, 36px on the product page (`rgba(66,66,69,.72)` on dark, hovered `rgba(73,73,76,.725)`, pressed `rgba(57,57,60,.74)`, in 0.1s linear), 56px on the store's shelves (`rgba(210,210,215,.64)` on light); a gallery's dots are 8px, the current one a 32px pill (`home:gallery.*`).

**The footer** (`home:footer.*`). `#f5f5f7`, 12/16 text at `rgba(0,0,0,.56)`, its links `.72` (underlined hovered), its column titles 600 at `.88`, its rules `rgba(0,0,0,.16)`.

**Accordions** (`bag:accordion.*`). A 1px `#d2d2d7` rule above each item; the question 32px from it (24px to 734px).

## Running it

```sh
PW=…/playwright-chromium/index.mjs PAGE=home node measure.mjs   # results/home.json
node report.mjs home [probe-prefix]                               # its values across the widths
```

`pages/<name>.mjs` is a page's spec: its URL and its probes. A probe selects an element of the page or adds markup to it (for a class the page defines but may not use), and reads the properties it names, its box, and with `states` the same again with each pseudo-class forced through the DevTools protocol. `measure.mjs` records a probe that matches nothing as null and fails the run, and records each element's markup without query strings (the store's carry session ids).

What the harness had to learn, each a reading that was wrong before it:

- A typography class a page does not define reads as body text; every spec carries the unclassed control `type.none`, and `report.mjs` prints such a class as not defined.
- Forcing a state off and re-enabling transitions in one frame starts a transition a later probe reads in flight (a paddle's icon read white instead of `rgba(255,255,255,.8)`); the harness flushes styles before transitions return.
- The sign-in field is focused on load, so its rest is read blurred (`blur`).
- The store's elements rename after hydration (`rf-ccard-content-eyebrow` in the served HTML is `rf-ccard-content-header-eyebrow` in the page).
- The sign-in widget's iframe shares the page's process here, so its elements are reached through the page's own DevTools session.

Not measured: a touch screen's hover, the modals, the checkout, the search field open, the global nav's flyouts (a desktop's menus under a link), and every page in Safari.
