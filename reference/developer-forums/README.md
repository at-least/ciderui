# Apple Developer Forums, measured

The one Apple console-like site found that pages a list: the forums' post lists (`/forums/latest`), read for crashcart's pager and its view switch (the owner, 2026-10-29: crashcart's own parts into ciderui, compared with Apple's sites, HIG-checked), where App Store Connect paged no list (`../asc/`). Signed out; read only: the pages reached by their address, nothing typed or posted.

Chrome 156 on macOS, its window 1920 by 907 at 2 device pixels; read 2026-10-28 19:54 UTC. `pages/pager.js` ran in the page; `results/pager.json` holds what it returned, its `page-1` and `tabs` entries read by smaller scripts in the same session.

## What it showed

**The pager.** A `<nav aria-label="Pagination">` under the list, 16px above and below, its items in a centered flex row 10px apart: the previous page's button, the page's number ("Page 2", 15px on 20px, `rgb(51, 51, 51)`, 15px either side of it), the next page's button. Each button an icon alone, a chevron in SF Pro Icons (11 by 10.5px), named by its `aria-label` ("Previous page", "Next page"): 39 by 28px, 10px in, 5px corners, white on `rgb(29, 29, 31)` (the site's neutral button). A direction with no page is not drawn (0 by 0px: page 1's previous and first-page buttons), never shown disabled. A first-page button exists, undrawn on pages 1 and 2.

**The view switch** (`nav.tabs-wrapper`): Topics and All Posts, the lists' two views, as pills: links 39px tall, 14px at 600, 7px corners; the current one (`aria-current="true"`) black on `rgb(226, 226, 228)`, the other `rgb(110, 110, 115)` on nothing. A choice of views drawn filled, not underlined.

## Not measured

- A phone's width; the last page (its next button's); a list with one page; hovered, pressed and focused states; the forums in dark (the site follows the system's).
