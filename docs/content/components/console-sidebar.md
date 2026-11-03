+++
title = "Console Sidebar"
description = "App Store Connect’s app sidebar: a wordmark, a list, the account; a sheet on a narrow window."

[extra]
category = "Console"
image = "pictures/console-sidebar.png"
+++

The app’s sidebar, down the page’s left: the product’s wordmark at its top, its list (each page an item, the current one filled), headings and rules between runs, the account at its foot. From 1024px it stands in the page; below, it is a sheet over the page that the page’s button opens and its own button closes, as a popover, with no script. Where it stands is the page’s: a grid of the sidebar and the page’s content.

```css
@import "ciderui/components/console/console-sidebar/console-sidebar.css";
```

{{ <cui_example file="console-sidebar/sidebar" title="A console sidebar" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-sidebar` | the sidebar, a `<nav>` with an `id`, `popover` and an `aria-label` |
| `cui-console-sidebar-header` | its top: the wordmark (`cui-console-sidebar-wordmark`) and its close button (`cui-console-sidebar-close`) |
| `cui-console-sidebar-body` | its list, a [Console Menu](@/components/console-menu.md) at the sidebar’s size: `cui-console-menu` with `data-variant="sidebar"` on the same element, its items, headings and rules the menu’s |
| `cui-console-sidebar-footer`, `-account`, `-avatar` | the account at its foot |
| `cui-console-sidebar-open` | the page’s button that opens it below 1024px (`popovertarget` its id) |

## On a narrow window

This page’s frames are narrower than 1024px, so the sidebar above is a closed sheet. Here, the page’s button opens it.

{{ <cui_example file="console-sidebar/open" title="The sidebar’s sheet and the button that opens it" pad={true} /> }}

The Go module writes it with `console.ConsoleSidebar`, its list `console.ConsoleMenuItem`s.
