+++
title = "Console Menu"
description = "App Store Connect’s list of a section’s pages down its left side."

[extra]
category = "Console"
image = "pictures/console-menu.png"
+++

The list of a section’s pages down the left of App Store Connect’s analytics: each page a pill, filled where the page is the current one and under the pointer; a heading over a run of them, a rule between runs, a group that opens with no script to show its own, set in from the left.

```css
@import "ciderui/components/console/console-menu/console-menu.css";
```

{{ <cui_example file="console-menu/menu" title="A console menu" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-menu` | the list, a `<nav>` with an `aria-label` |
| `data-variant="sidebar"` | the list in a [Console Sidebar](@/components/console-sidebar.md), at the toolkit’s sidebar size: items 44px at 15px, 48px at 18px below 640px, their icons 24px wide |
| `cui-console-menu-item` | a page, an `<a>`: 36px, 13px at 500; `aria-current="page"` the current one; `aria-disabled="true"` (and no `href`) one that goes nowhere now |
| `cui-console-menu-icon` | its icon, 25px wide, before its label |
| `cui-console-menu-label` | its label, cut with an ellipsis |
| `cui-console-menu-secondary` | what follows it at its end, a count |
| `cui-console-menu-group` | a group, a `<details>`: its `<summary>` an item with a chevron (`cui-console-menu-chevron`), its items 48px in |
| `cui-console-menu-heading` | a heading over a run of items |
| `cui-console-menu-rule` | a rule between runs, an `<hr>` |

Below 640px its items are 37px at 16px. The Go module writes it with `console.ConsoleMenu`, `console.ConsoleMenuItem` and `console.ConsoleMenuGroup`, an icon any of package `lucide`’s.

## Choices

A list to choose one from, in a [Console Popover](@/components/console-popover.md): a [Console Player](@/components/console-player.md)’s subtitles. Each choice is a `<button role="menuitemradio">` item, its `value` what the page’s script reads, its check in the icon’s place, shown on the one with `aria-checked="true"` and kept hidden on the others so every label starts in line. A click closes the popover by itself (`popovertarget` and `popovertargetaction="hide"`); the script moves `aria-checked` to the choice and moves the focus among the choices by the arrow keys. A language’s name carries its `lang`.

{{ <cui_example file="console-menu/choices" title="A console menu of choices" pad={true} /> }}

Two lists side by side, each its title over its choices, are one menu of columns (`cui-console-menu-columns`, each column a `cui-console-menu` with its `data-column`), as the Apple TV app’s Audio beside its Subtitles: a rule between them, each column its own check, the popover as wide as they ask (at most 85% of the window).

{{ <cui_example file="console-menu/columns" title="A console menu of choices in two columns" pad={true} /> }}

The Go module writes it with `console.ConsoleChoiceMenu` (the popover’s id, its columns), `console.ConsoleChoiceColumn` (its name, its title, its choices; one is checked) and `console.ConsoleChoice` (the label, the value, the language, checked).
