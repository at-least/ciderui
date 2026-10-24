+++
title = "Console Menu"
description = "App Store Connect’s list of a section’s pages down its left side."

[extra]
category = "Console"
image = "pictures/console-menu.png"
+++

The list of a section’s pages down the left of App Store Connect’s analytics: each page a pill, filled where the page is the current one and under the pointer; a heading over a run of them, a rule between runs, a group that opens with no script to show its own, set in from the left.

```css
@import "ciderui/components/console-menu/console-menu.css";
```

{{ <cui_example file="console-menu/menu" title="A console menu" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-menu` | the list, a `<nav>` with an `aria-label` |
| `cui-console-menu-item` | a page, an `<a>`: 36px, 13px at 500; `aria-current="page"` the current one; `aria-disabled="true"` (and no `href`) one that goes nowhere now |
| `cui-console-menu-icon` | its icon, 25px wide, before its label |
| `cui-console-menu-label` | its label, cut with an ellipsis |
| `cui-console-menu-secondary` | what follows it at its end, a count |
| `cui-console-menu-group` | a group, a `<details>`: its `<summary>` an item with a chevron (`cui-console-menu-chevron`), its items 48px in |
| `cui-console-menu-heading` | a heading over a run of items |
| `cui-console-menu-rule` | a rule between runs, an `<hr>` |

Below 640px its items are 37px at 16px. The crate writes it with `ConsoleMenu`, `ConsoleMenuItem` and `ConsoleMenuGroup`, an icon any of `ciderui::lucide`’s.
