# Apple Account, measured

account.apple.com's page for managing an Apple Account, the owner's example of the sites they asked for (2026-10-23: 「淺色，但找更多apple 網站，例如 https://account.apple.com/account/manage」): settings laid out as a side nav of sections and cards that each open their own sheet, for a site with settings of its own (crashcart: a project's settings, the account's users and keys). Its page holds the account's personal data and only a signed-in browser reaches it, so it is read as App Store Connect's are (`../asc/`): in the owner's own Chrome, signed in by the owner, through Claude in Chrome's javascript tool; `read.mjs` runs App Store Connect's reader and check over this folder's pages and results.

An element is recorded by its tag, classes and role, never its text, link or label. Read only: the page was reached by the address bar; nothing was typed, clicked, saved or sent. The page holds several sections' markup at once (five page headers): each probe reads the first match in the document's order, at the page's top, the Sign-In and Security section the page opens on; the payment, subscriptions and family sections' own pages were never opened.

Chrome 154 on Linux, its window 1920 by 898 at 2 device pixels; read 2026-10-23 (00:49 UTC), the owner's appearance set to light.

| spec | page | what |
| --- | --- | --- |
| `manage` | `/account/manage` | the page as it opens: its ground; the side nav (the person's block: the picture, the line with the name and the one under it; the list, an item, a link and the current one); the page's header (its title, its description); the section's grid (a column) and a card (the button it is, its ground and edge, its title, its line, the glyph at its right) |

## What it showed

**apple.com's page.** White, its text `rgb(29, 29, 31)` in SF Pro Text at 17/25 and -0.374px, antialiased.

**The side nav**, 290px. The person's block: an 80px picture, the name's line in SF Pro Display at 19/23, 600, the line under it at 14/20 in `rgb(110, 110, 115)`. The list's rows 32px: a link at 17/25, `rgb(29, 29, 31)`, 3px above and 4px below, never underlined; the current one `rgb(0, 113, 227)` at 700 (no ground on the item or its link, read by a script in the page, not recorded in a result; a mark drawn by a pseudo-element is unread).

**The header.** The title in SF Pro Display at 28/32, 700 (0.196px); the description 15px under it at 14/20 in `rgb(66, 66, 69)`; the cards 40px below.

**Cards**, two columns of 335px. Each card is a `button` (a 1px transparent edge around it) holding the card: 315 by 124px, white, a 1px `rgba(210, 210, 215, 0.63)` edge, 11px corners, no shadow; its title in SF Pro Display at 19/23, 600; its line at 14/20 in `rgb(110, 110, 115)`; at its right a 40px glyph filled `rgb(0, 113, 227)`.

## Not measured

- A card's sheet (every card is a button that opens one: no click); hover, focus and pressed.
- The other sections (personal information, devices, privacy not read; payment, subscriptions and family never opened).
- A narrower window: Claude in Chrome's resize did not take for App Store Connect, and this page was read at 1920px alone.

```sh
node read.mjs <spec>           # the script to run in the page
node read.mjs --check <spec>   # once results/<spec>.json is written
```
