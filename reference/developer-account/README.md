# The Apple Developer account, measured

developer.apple.com/account: the developer's account (its overview, cards of the program's resources and the membership's details) and Certificates, Identifiers & Profiles (lists of certificates, identifiers, devices, profiles and keys), for a site with settings and keys of its own (crashcart), at the owner's ask (2026-10-23: 「淺色，但找更多apple 網站」). Its pages hold the account's and the team's data and only a signed-in browser reaches them, so they are read as App Store Connect's are (`../asc/`): in the owner's own Chrome, signed in by the owner, through Claude in Chrome's javascript tool; `read.mjs` runs App Store Connect's reader and check over this folder's pages and results.

The site follows the system's dark (with the system set to dark, its body black and a dark-mode stylesheet among its own: read by a script in the page, not recorded in a result); it was read light because the owner set their appearance to light for the read (「done」), as they chose light (「淺色」). An element is recorded by its tag, classes and role, never its text, link or label. Read only: the pages were reached by the address bar; nothing was typed, clicked, saved or sent. The overview's emails card and the team's address were not read.

Chrome 154 on Linux, its window 1920 by 898 at 2 device pixels; read 2026-10-23 (00:54 and 00:56 UTC).

| spec | page | what |
| --- | --- | --- |
| `overview` | `/account` | the page and the app's ground; the tab nav (a tab, its glyph, a paddle, disabled while every tab shows); a card (its ground, corners and padding); the program card's title, a section's picture, title, text and link; the membership's rows (a row, its label and the word in it, its value) and the info button; the link that adds |
| `keys` | `/account/resources/authkeys/list` | the list's text; the side nav (an item and its word, the current one and its word); the list's header (its rule, the title, the add glyph, the filter's button and its chevron); the header row (a column's label, a sortable one) and a row (a cell, the glyph in one) |

## What it showed

**The overview**, the developer site's newer look. The app's ground `rgb(245, 245, 247)` (the body under it white), its text `rgb(29, 29, 31)` in SF Pro Text at 17/25. Cards white with 18px corners and no shadow, 51px in, 25.5px apart, 980px wide. A card's title in SF Pro Display at 24/28, 600; a section's 56px picture, its title at 19/23, 600, 5.7px under it, its text at 14/24 in `rgb(102, 102, 102)`, its links at 14/24 in `rgb(0, 102, 204)`, not underlined. The membership's details as rows, a label and a value: the label at 17/25, 700, in `rgb(102, 102, 102)`, its column 220px; the value at 17/25; 11.9px under each row, no rule. A 19px info button; the link that adds at 17/25 in `rgb(0, 102, 204)`, a glyph after it. The tab nav: 44px glyphs, each over its word at 11px in `rgb(45, 45, 45)`; a paddle 54px tall on `rgb(245, 245, 247)`, at opacity 0 while it is disabled.

**The keys list**, App Store Connect's older look, its toolkit's classes the same (`Box-sc-18eybku-0`, `Link-sc-1tc8rtf-0`, `SvgIcon__Icon-sc-194a9zm-0`): the system's font (`-apple-system, BlinkMacSystemFont, …`) at 13/20 in `rgb(51, 51, 51)`. The side nav, 183px: its items 26px, 5px apart, 3px by 7px in; the current one on `rgb(237, 237, 237)` with 4px corners, its word at 500. The list's header 50.5px over a 1px `rgb(229, 229, 229)` rule: its title at 24/30, 700, in `rgb(17, 17, 17)`; a 20px add glyph and the filter's button (16/18) in `rgb(0, 112, 201)`, its 9.6px chevron 6px after it. The header row 40.5px: its labels in `rgb(102, 102, 102)` on white, 10px above and below, a 1px `rgb(229, 229, 229)` rule under them; a sortable one with a pointer. A row 43px: its cells 10px above and below at 13/20, centered on the row, a 1px `rgb(229, 229, 229)` rule under each; a glyph 11 by 16px in one.

## Not measured

- Hovered and focused states, the filter's menu, a key's own page (and its one download), creating anything; the other lists (certificates, identifiers, devices, profiles), the keys list's twins.
- Dark, which the site has: read light only, as the owner chose.
- Any width but 1920px.

```sh
node read.mjs <spec>           # the script to run in the page
node read.mjs --check <spec>   # once results/<spec>.json is written
```
