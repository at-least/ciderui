# App Store Connect, measured

Apple's own console, for a site that is one (crashcart, the owner's error tracker: its tables, statuses, sidebar and charts), where apple.com's pages show none of these. Its pages hold an account's data, and only a signed-in browser reaches them (a headless one lands on `/login?targetUrl=%2Fapps&authResult=FAILED`), so they were read in the owner's own Chrome, signed in by the owner, who asked for them to be measured with their account (2026-10-23).

`node read.mjs <spec>` prints a script; Claude in Chrome's javascript tool runs it in the page; what it hands back is written by hand to `results/<spec>.json`, and `node read.mjs --check <spec>` holds the file to the page's digest of it (a value copied wrong fails), to the spec's probes and the reader's source (a script pasted wrong fails), and to the account's privacy (an id or an address fails). `read.mjs`'s command line is `cli(root)`, which another site's folder calls with its own pages and results (`../feedback-assistant/`). An element is recorded by its tag, classes and role, never its text, link or label; the app's and the team's ids are `<app>` and `<team>`. Read only: the pages were reached by their own links and the address bar; nothing was typed, dismissed, saved or submitted.

Chrome 154 on Linux, its window 1920 by 954 at 2 device pixels; read 2026-10-23 (17:02 to 17:07 UTC on the 22nd).

| spec | page | what |
| --- | --- | --- |
| `apps` | `/apps` | the page's ground; the bar (amp-nav's header, in its shadow root): its ground, logo, links and the current one; the sticky header: its title, add button, status filter and rule; the banners, plain and informational; an app's tile and its version's status |
| `testflight` | `/teams/<team>/apps/<app>/testflight/ios` | the app's tabs; the sidebar: a heading, a link, the current one, the rule; the page's title; a group of builds opened; its table: the header, a column's label, a row, a cell, the version's link, a status (its mark, its word, the line under it), a count, the info button |
| `crashes` | `/teams/<team>/apps/<app>/testflight/crashes` | the table with no crash: its header over the empty row's message |
| `analytics` | `/apps/<app>/analytics/overview` | the newer look (classes `rt-…`): the menu, the date's control, a metric's card, a sparkline, a bar chart, a plain table |
| `version` | `/apps/<app>/distribution/ios/version/inflight` | the tabs (links, and the media's buttons), the header's buttons (primary, neutral, disabled), a help button, a field and a disabled one, a label, a multi-line field and its count, the checkbox and the radio |
| `tokens` | `/apps/<app>/analytics/overview` | the newer look's toolkit's tokens (`--tb-…`, on `:root`), each as the page computes it: the type, sizes, spaces, radii, borders, shadows, timing, breakpoints, every color role and both palettes, and the components the console parts are built from (their names in `pages/tokens.mjs`) |
| `analytics-page` | `/apps/<app>/analytics/overview` | what `analytics` left: a metric's card, a section's heading and its See All link, the menu's box, a group's button and chevron, an item in the open group, a disabled item |
| `button`, `hover-button` | `/apps/<app>/analytics/acquisition/sources` | the toolkit's gray button (its rounded rectangle, 30px, 13px at 500) at rest and with the pointer resting on it |
| `hover-menu`, `hover-row`, `hover-link`, `hover-date` | `/apps/<app>/analytics/overview` | each with the pointer resting on it (Claude in Chrome's hover): a menu item that is not the current one, the plain table's row, a See All link, the date control's period |

## What it showed

**Light only.** With the system set to dark, no `prefers-color-scheme` rule is in the page's stylesheets (10 on the apps page, 16 on the analytics page) or in the shadow roots of the bar and the footer (`amp-nav`, 234 rules; `amp-footer`, 34); the `--tb-color-…` tokens (827 declarations) are set on `:root` alone, the `html.dark` rules are an alert box's (`tb-alertbox`), and `html` has no `dark` class. Read by a script in the page, not recorded in a result.

**Two looks.** The older, on most pages: SF Pro Text at 13px/18px, `rgb(51, 51, 51)` text, links and the primary button `rgb(0, 112, 201)`, rules `rgb(214, 214, 214)`, 4px corners, 28px buttons, 30px fields; titles in SF Pro Display at 700, 34px/41px (the apps') and 24px/30px (a page's). The newer, on the analytics page (classes `rt-…`): the system font (`-system-ui, -apple-system, …`), text `oklch(0 0 0 / 0.85)`, 36px menu items with 18px corners (the current one on `oklch(0.575 0.012 286 / 0.08)`), its blue `oklch(0.542 0.162 251.9)` (the date's control, the sparkline's line).

**The bar.** amp-nav's header: 50px on `rgb(241, 241, 241)`, no edge or shadow, its content 1442px at most; its links black at 14px/19px, 400, the current one at 600.

**A status.** A 12px mark in the status's color, 6px before its word (13px/18px, `rgb(51, 51, 51)`), and under it a line at 11px/13px, `rgb(136, 136, 136)` (a build's); in an app's tile, the mark before a 12px/16px line in `rgb(136, 136, 136)`. Seen: a build's green, `rgb(0, 125, 27)`, and a version's yellow, `rgb(255, 204, 0)`.

**Tables.** Collapsed borders. The header row on white over a 1px `rgb(214, 214, 214)` rule, 37.5px over the builds and 36px over the crashes, its labels 12px/16px at 500, `rgb(102, 102, 102)`, upper case. A row 59px between 1px `rgb(214, 214, 214)` rules; its cells 16px by 10px in (the first none at its left), 13px/18px `rgb(51, 51, 51)`, centered on the row; a count right-aligned. The newer look's plain table: rows 57px, its cells 19px above and below (the first none at its left and 8px at its right), a 1px `oklch(0.359 0.012 285.8 / 0.13)` rule under each.

**An empty table.** Its one row holds a box on `rgb(250, 250, 250)`, 60px above and below its message, 13px/18px `rgb(51, 51, 51)`, centered.

**Banners.** A 1px edge and 4px corners: the plain one on `rgb(250, 250, 250)`, its edge `rgb(214, 214, 214)`, its icon `rgb(51, 51, 51)`; the informational one on `rgb(235, 244, 251)`, its edge `rgb(204, 226, 244)`, its icon and dismiss mark `rgb(0, 112, 201)`; a title at 15px/20px, 600, its text 13px/18px, an 8px dismiss mark 15px from the corner.

**Navigation.** The app's tabs at 20px/22px, 500, in the blue, the current one on a 2px rule in it. The sidebar, 215px: headings at 15px/30px, 600; links 30px tall at 13px, 10px in, the current one on `rgb(242, 242, 242)` with 4px corners; a 1px `rgb(214, 214, 214)` rule between sections. A section's tabs at 15px/20px, 600, in the blue, the current one on a 2px blue rule (the media's at 18px/22px, 500, `rgb(51, 51, 51)`, the current one on a 2px rule in it).

**Controls.** Buttons 28px with 4px corners at 13px: the primary blue with white text at 500, 11px in; the disabled one white, its text and edge `rgb(204, 204, 204)`; a neutral one on `rgba(214, 214, 214, 0.5)` at 400, 10px in. A help button, a 17px circle on `rgb(242, 242, 242)`, its glyph `rgb(136, 136, 136)`. Fields 30px tall (a multi-line one 90px) with a 1px `rgb(214, 214, 214)` edge and 4px corners, 8px in; a disabled one on `rgb(250, 250, 250)`, its text `rgb(136, 136, 136)`. Labels at 13px/18px, 500; a count under a field `rgb(136, 136, 136)`, right-aligned. The checkbox and the radio are the browser's own, 13px.

**Charts** (the analytics page). A sparkline: a 2px line in the blue, round at its ends and joins, over an area filled from `rgb(192, 227, 255)` at its top to `rgb(249, 252, 255)`. A bar chart: grid lines 1px `rgb(214, 214, 214)`; the axes' labels `rgb(136, 136, 136)`, 11px at 600 up the side and 13px at 400 along the foot; bars filled from `rgb(66, 147, 217)` to `rgb(40, 133, 211)`. A metric: its title at 13px/18px, 600, its value at 41px/45px, 700, its change at 18px/22px, 500 (here `oklch(0.654 0.232 28.8)`). The date's control: a 30px pill of three buttons in the blue (15px corners at its two ends), its period in white at 13px, 600.

**The toolkit.** The newer look is a design system of tokens on `:root` (`--tb-…`: `--tb-meta-theme` "apps", `--tb-meta-mode` "light"), loaded on the analytics pages alone (Trends, Users and Access, Business and the other pages carry none, read 2026-10-24), and its components' classes say their state (`rt-Item--effect--hover`, `--selected--…`, `--effect--disabled`). Its colors are roles (`neutral`, `brand`, `error`, `warning`, `success`, `info`; each an accent, labels, fills, borders and grounds at fixed alphas) over two palettes, light and dark, of the same steps (`coolGray-800` is `oklch(35.9% 0.012 285.8)` light and `oklch(87.9% 0.015 286.1)` dark); the roles are light only, their dark side unshipped. `tokens.json` holds 3479 of them; `read.mjs`'s `TOKENS` reads them and hands the result back as a file the page downloads (too long for the tool's output); the hover results were handed back the same way. Its type: the system font with `letter-spacing: normal`, body 13/18, h4 18/22, h1Plus 41/45. Its controls: 22, 30, 40 and 50px (`--tb-size-control-*`), capsules by default. Its breakpoints: 640, 1024, 1440, 1920px.

**Hovered** (`hover-*`): a menu item on the neutral fill at 0.08 (`oklch(0.575 0.012 286 / 0.08)`), its label unchanged; the plain table's row on it at 0.03; a See All link underlined, its color unchanged; the date control's period `oklch(0.589 0.148 248.4)`, lighter than its rest. A metric's card is `oklch(0.968 0.001 286.4)` with 10px corners, 15px above its title, no edge or shadow; a group's item is 48px in; a disabled item's label the neutral label at 0.39.

## Not measured

- A narrower window: asked for 420 by 900, Claude in Chrome reported the window resized and the page's viewport stayed 1920 by 954 (why not found); App Store Connect's own steps (`xl-device` on `html`) are unread.
- States: focused and pressed (hovered only on the analytics page, `hover-*`); a menu or popover opened (the status filter's, an info button's), a dialog. The newer look's Badge, Banner, Tabs, Chip, Popover, EmptyState, inputs and Toggle have no instance in the account: their tokens alone.
- Other statuses' colors: one app, its builds' green and its version's yellow.
- Xcode Cloud; Users and Access (its table the builds' component); the sales, payments and business pages, not opened.

```sh
node read.mjs <spec>           # the script to run in the page
node read.mjs --check <spec>   # once results/<spec>.json is written
```
