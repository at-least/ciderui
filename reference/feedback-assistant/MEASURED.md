# Feedback Assistant, measured

Apple's tracker for the reports developers send it: a list of items, each opened beside it, for a site that is one (crashcart, the owner's error tracker: its issues and an issue). Its pages hold the account's reports and only a signed-in browser reaches them (signed out, it sends the browser to Apple's sign-in), so they are read as App Store Connect's are (`../asc/`): in the owner's own Chrome, signed in by the owner, who asked for more of Apple's sites to be measured (2026-10-23: 「淺色，但找更多apple 網站」), through Claude in Chrome's javascript tool. `read.mjs` runs App Store Connect's reader and check over this folder's pages and results.

An element is recorded by its tag, classes and role, never its text, link or label; the reader writes a run of six digits or more as `<app>` (App Store Connect's word for it), here the feedback's number. Read only: the pages were reached by the address bar; nothing was typed, clicked, saved or sent. The list's other items are Apple's announcements, unread, and opening one marks it read, so none was opened; the feedback opened is the owner's own, already read.

Chrome 154 on Linux at 2 device pixels; read 2026-10-23: `item` at 00:21 UTC in a 1920 by 954 window, `list` at 00:51 UTC in a 1920 by 898 one, after the owner signed in again (the session had ended between the two).

| spec | page | what |
| --- | --- | --- |
| `item` | `/feedback/<app>` | a feedback opened: its card in the list, now the current one; the byline (the monogram, the name, the date); the title's row and the actions button; the status box (a label and its value); a section (its lead, a label, its text); an attachment (its glyph, its link, the line under it); the reply area at the foot (its field, the attach glyph, the send button, disabled while the field is empty) |
| `list` | `/` | the list as it opens, nothing chosen: the page's ground; the bar (Apple's dark local nav: its ground and the line under it, the title, the compose and help glyphs, the account's menu); the list's pane: its search field and its placeholder, the sort and filter row, an item (an announcement, unread: its card, the unread mark, its title, its kind, the line under it, its date); the other pane's empty state |

## What it showed

**Light only.** With the system set to dark, the list page's 32 readable stylesheets hold no `prefers-color-scheme` rule (one more, from another origin, is unread). Read by a script in the page, not recorded in a result.

**Its own accent, a purple.** The actions button's text `rgb(109, 29, 169)`, the attach glyph the same; the status box's edge `rgb(215, 171, 236)`; the current item's card `rgb(246, 235, 250)`. The send button, disabled while its field is empty, declares `rgb(177, 84, 224)` as its ground and is drawn at 0.3 opacity, so what shows is paler; its enabled look is unread.

**Type.** SF Pro Text at 17/25 and -0.374px in `rgb(51, 51, 51)`; the title 24px on 35.3px at 600; a section's lead 21px on 30.9px; a section's label, a date and an attachment's line `rgb(135, 135, 135)` (the status's label `rgb(139, 135, 135)`); the name in the byline at 700; the date and the attachment's line 13.6px on 20px.

**The current item.** Its card 405 by 90px on the purple tint, a 1px `rgb(229, 229, 229)` edge at its left and a 1px `rgb(237, 237, 237)` rule under it.

**The bar**, Apple's dark local nav: 60px, held at the top (`position: sticky`), `rgba(34, 34, 34, 0.7)` under `saturate(1.8) blur(20px)`, a 1px `rgba(51, 51, 51, 0.95)` line under it; its title white in SF Pro Display at 26px, 500; the compose glyph white at 26px; the help glyph and the account's menu white.

**The list's pane**, 420px on the page's `rgb(250, 250, 250)`, a 1px `rgb(229, 229, 229)` edge at its right. Its search field a tags input (`react-tagsinput`), 50px in a 56px row, no edge, no ground, at 16px in the browser's sans-serif (no font set), its text `rgb(119, 119, 119)`, its placeholder `rgb(117, 117, 117)`. Under it the sort and filter row, 37px on `rgb(250, 250, 250)` between 1px `rgb(229, 229, 229)` rules, its words and glyphs at 15px in `rgb(136, 136, 136)`. An item: 90px (at least), 10px in (15px at its right), a 1px `rgb(229, 229, 229)` edge at its left and a 1px `rgb(237, 237, 237)` rule under it, no ground; an unread one marked by an 11px purple dot, `rgb(176, 81, 222)`; its title 17/25 at 700 in `rgb(39, 39, 39)` on one line, cut with an ellipsis; its kind at 14px and the line under it at 13px, both `rgb(136, 136, 136)`; its date at 14px, `rgb(136, 136, 136)`, at the row's end.

**The empty pane.** With nothing chosen, the other pane holds one line centered in it, 34px on 50px in `rgb(214, 214, 214)`.

**The detail.** The byline: a 47px circle on `rgb(169, 169, 169)` holding white initials (Helvetica, 17px), the name and the date beside it. The title's row on white over a 1px `rgb(229, 229, 229)` rule, 17px in; beside the title the actions button, its text alone (no ground, no edge), 35px tall. The status box: `rgb(250, 250, 250)`, a 1px edge in the purple's tint, 4px corners, 12px in, 17px above and below. A section 15px in, each label 17px above its text. An attachment: a 25px glyph, its name at 16px, not underlined, its size and date under it. The reply area held at the window's foot (`position: sticky`, 136px at least) on `rgb(243, 243, 243)` under a 1px `rgb(229, 229, 229)` rule: its field 64px, a 1px `rgb(204, 204, 204)` edge, 5px corners, 8px in, in the browser's monospace at 16px (the page sets no font on it), not resizable; the send button 100 by 37.7px, 4px corners, white text at 18.7px, at 0.3 opacity while it is disabled.

## Not measured

- An unread item, the compose form, the sort's and filter's menus, hovered and focused states; any width but 1920px.

Claude in Chrome's output hid the values of every probe named `author…` (as it hides a hex digest), so the byline's probes are named `byline`.

```sh
node read.mjs <spec>           # the script to run in the page
node read.mjs --check <spec>   # once results/<spec>.json is written
```
