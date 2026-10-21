+++
title = "Global Nav"
description = "apple.com’s bar across the top: a logo, its links and their flyouts, a search and a bag, a menu below 834px."

[extra]
category = "Navigation"
image = "pictures/globalnav.png"
+++

apple.com’s bar across the top of every page, on a `<nav>`: the site’s logo, its links and its icon links (a search, a bag) spread across 1024px at most, on a translucent material under a blur. Below 834px the links fold into a menu, which the bar’s button opens over the window as a sheet, with no script: the list is a `popover`. This page’s own bar is one.

```css
@import "ciderui/components/globalnav/globalnav.css";
```

{{ <cui_example file="globalnav/globalnav" title="A global nav" /> }}

The list is the menu. Closed, the page shows it as the bar’s row; opened, it is a sheet over the window with a close button of its own where the menu’s button was, its two lines turning into an X, and the focus moves to it (`autofocus`); Tab goes on to the links, and the X or Escape closes it, the focus back on the menu’s button. In a narrow window the example’s menu opens over its frame.

| class | what |
| --- | --- |
| `cui-globalnav` | the bar: 44px, 48px below 834px; `cui-dark` on it, a dark bar; `data-variant="clear"`, no ground and no blur, over a hero |
| `cui-globalnav-content` | 1024px at most, 22px in from 834px |
| `cui-globalnav-list` | the row closed, the menu opened: `id` and `popover` |
| `cui-globalnav-logo` | the site’s logo, a link home, named by its `aria-label` |
| `cui-globalnav-link` | a link: 12px in the bar, a 28px row in the menu |
| `cui-globalnav-icon` | an icon link, its image `cui-globalnav-image` from 834px and `cui-globalnav-image-compact` below |
| `cui-globalnav-menubutton` | the menu’s button, below 834px, beside the list: `popovertarget` |
| `cui-globalnav-closebutton` | the opened menu’s X, second in the list so the next stop is the first link |

The search can open over the page instead of leading to the site’s search page: its icon is then a button (`popovertarget`) opening another popover, the search, with no script. Its field takes the focus, and Return sends what is typed to the site’s search page (the form’s `action`, GET, as `q`); once something is typed, a clear button shows at its end, a reset button that empties it; Escape, its button again, a click on the curtain or on the bar closes it, the focus back on the button. From 834px it hangs from the bar’s bottom, placed there by anchor positioning, its panel over a curtain that blurs the page; below 834px it is a sheet over the window with an X of its own where the menu’s button is, the page held. In a narrow window the example’s search opens over its frame.

{{ <cui_example file="globalnav/search" title="A global nav and its search" /> }}

| class | what |
| --- | --- |
| `cui-globalnav-search` | the search, after the menu’s button: `id` and `popover`; a flyout under the bar from 834px, a sheet over the window below |
| `cui-globalnav-searchpanel` | its ground and its room: the field, then the quick links |
| `cui-globalnav-searchfield` | the form, `role="search"`: its magnifier and its input |
| `cui-globalnav-searchinput` | the field, `autofocus`: `type="text"`, as with `type="search"` Escape in it leaves the search open in Safari |
| `cui-globalnav-searchreset` | after the field: its clear button, `type="reset"`, its mark `icon::CLEAR` (`icon::CLEAR_COMPACT` below 834px); shown while the field holds text |
| `cui-globalnav-searchheading` | a section’s heading, which names its list: the quick links’, a suggestions’ section’s |
| `cui-globalnav-searchlinks` | its rows, each a `cui-globalnav-searchlink` after its image, on a ground hovered and focused |
| `cui-globalnav-curtain` | the curtain under the flyout from 834px, a button that closes it (`popovertargetaction="hide"`, out of the tab order) |

The clear button resets the form, so it gives the field back its `value` attribute: a page that fills the field with its query leaves the attribute out, or the button brings the query back. Clicked, the button hides itself and the focus is left on the page, where apple.com’s script puts it back in the field; Shift+Tab takes it there.

Its suggestions are a site’s script’s, as apple.com’s are (its script asks its service for them as one types): in the quick links’ place it writes sections, each a heading and its rows, a suggestion’s text in a `span` with what was typed in a `mark`. While a row marks what was typed, every row is muted and its marked part takes the rows’ color, hovered too; a section after another is 34px below it (48px below 834px). A suggested link keeps the arrow, a suggested search has a magnifier, `icon::SUGGESTION`. Rows the script writes once something is typed fade in where they are, as apple.com’s new results do: the heading from 40ms, each row 20ms after the one above; the search’s opening motion is its empty field’s. The example is the search as a script leaves it once “ma” is typed: what was typed is the field’s value, which the markup does not hold.

{{ <cui_example file="globalnav/suggestions" title="A global nav and its search’s suggestions" /> }}

The bag can open over the page too, as apple.com’s does, where it would lead to the site’s bag: its icon is then a button opening another popover, the bag, with no script, an empty bag’s. Its message, a line under it beginning with a link (to sign in), and its links under a heading, each after its icon: the orders, the saved items, the account, signing in. It opens and closes as the search does: Escape, its button again, a click on the curtain or on the bar. From 834px it hangs from the bar’s bottom over a curtain that blurs the page, the focus staying on its button, Tab going on into it; below 834px it is a sheet over the window, its X taking the focus, the page held. What apple.com’s bag shows with something in it could not be read (its store refuses a headless browser’s Add to Bag), so the kit has the empty bag’s alone. In a narrow window the example’s bag opens over its frame.

{{ <cui_example file="globalnav/bag" title="A global nav and its bag" /> }}

| class | what |
| --- | --- |
| `cui-globalnav-bag` | the bag, after the menu’s button (and the search’s): `id` and `popover`; a flyout under the bar from 834px, a sheet over the window below |
| `cui-globalnav-bagpanel` | its ground and its room: its message, line, heading and links, and the sheet’s X |
| `cui-globalnav-bagmessage` | its message, an `h2` |
| `cui-globalnav-bagtext` | a line under it, its link underlined |
| `cui-globalnav-bagheading` | the heading over its links, an `h3`, which names their list |
| `cui-globalnav-baglinks` | its rows, each a `cui-globalnav-baglink` after its icon: `icon::ORDERS`, `icon::SAVES`, `icon::ACCOUNT` or `icon::SIGN_IN` (their compact forms below 834px) |
| `cui-globalnav-closebutton` | the sheet’s X below 834px, `autofocus` |
| `cui-globalnav-curtain` | the curtain under the flyout from 834px, a button that closes it |

A link can open a flyout, a menu under the bar, as apple.com’s do, with no script: in the link’s item, after the link, its button and the flyout, another popover. From 834px, with a pointer, the pointer resting on the link shows it, a panel from the bar down across the window, its groups in a row, the page under it blurred; moved on along the bar, the next link’s replaces it, and off the bar and the flyout, it goes. Under 120ms on a link shows nothing, so a pointer passing over the bar opens no flyout. The button, at the link’s end, shows its chevron only with the focus from the keyboard: Return opens the flyout, Tab goes into it, Escape closes it, the focus back on the button. On a screen with no hover the button lies over the link, so a tap on the link opens the flyout, as apple.com’s script makes it. Below 834px the menu’s row is the button: a tap opens the flyout over the menu, a sheet with a back button, which goes back to the menu, and an X, which closes both. The pointer shows the example’s flyouts where its frame is 834px wide or more.

{{ <cui_example file="globalnav/flyout" title="A global nav with flyouts" /> }}

| class | what |
| --- | --- |
| `cui-globalnav-flyoutbutton` | after the link: `popovertarget`, `aria-label` its link’s and “menu”; its chevron `icon::FLYOUT` from 834px, `icon::FLYOUT_COMPACT` below, at the row’s end, hovered |
| `cui-globalnav-flyout` | the flyout, after the button: `id` and `popover` |
| `cui-globalnav-flyoutpanel` | its ground and its room: the sheet’s back button and X, then its groups |
| `cui-globalnav-flyoutgroup` | a group: a heading over its links, in a row from 834px, one under another below |
| `cui-globalnav-flyoutheading` | a group’s heading, which names its list; a group of large links shows none below 834px |
| `cui-globalnav-flyoutlinks` | its rows, each a `cui-globalnav-flyoutlink`; `data-size="lg"` a large one (24px, 28px below 834px), first in its group |
| `cui-globalnav-backbutton` | the sheet’s back button below 834px, at the bar’s start: `popovertargetaction="hide"`, its chevron `icon::FLYOUT_BACK` |
| `cui-globalnav-closebutton` | the sheet’s X below 834px, hiding the menu, and the flyout with it |
| `cui-globalnav-curtain` | the curtain under the flyout its button opens, from 834px; the pointer’s flyout has the list’s own, which takes no click |

The nav has no place on the page of its own: a site that keeps it at the top adds `sticky top-0`. Its `z-index` is apple.com’s 9999, so that it and its flyouts are over the page, a local nav under it included; a `z-*` utility on it would replace that. The crate writes it with `GlobalNav`, its links `NavLink`, a flyout’s groups `FlyoutGroup`, its search `Search::Link` or `Search::Flyout`, its bag `Bag::Link` or `Bag::Flyout`, a bag’s links `BagLink`, its icons `icon::SEARCH`, `icon::BAG` and `icon::MENU`, their compact forms, the search’s `icon::SEARCHFIELD`, `icon::ARROW`, `icon::CLEAR` and `icon::SUGGESTION`, the bag’s `icon::ORDERS`, `icon::SAVES`, `icon::ACCOUNT` and `icon::SIGN_IN`, and a flyout’s `icon::FLYOUT`, `icon::FLYOUT_COMPACT` and `icon::FLYOUT_BACK`.
