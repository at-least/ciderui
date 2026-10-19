+++
title = "apple.com’s parts, measured"
description = "ciderui’s catalogue is built: twenty-one of apple.com’s parts as classes, each held to the page it was read from."
date = 2026-10-19T17:31:36+08:00

[extra]
category = "Update"
image = "pictures/parts.png"
+++

Every row of ciderui’s catalogue is built: the button, the link, the section, the global nav, the local nav, the ribbon, the tile, the card, the shelf, the gallery, the field, the checkbox, the option, the swatch, the accordion, the quick links, the product nav, the footer, the price, the stat and the footnote’s mark.

## Read, not copied

apple.com’s stylesheets carry no license, so none of them is in ciderui. What it keeps is what apple.com’s elements compute to, read off its live pages (its home page, a product’s page, the store, a buy page, the bag) by a probe that names the page, the element, the width and the day. A component is written from those values, and its checks render it in Chromium and compare what it computes with what apple.com’s did: a color, a size, a corner, a duration. Where ciderui departs from apple.com, the check says why beside the value it holds instead.

## What the measurements decided

- **The type is the system font’s.** On Apple’s platforms the text is SF, so the letter-spacing is apple.com’s less the tracking SF already carries: apple.com sets its values on webfonts that carry none.
- **Dark is a block’s.** apple.com has no dark mode; a section, or a whole product page, is dark. So is ciderui’s, with `cui-dark` on any element.
- **No script.** A menu is a `popover`, an answer a `<details>`, a choice a radio: the parts that open do so with HTML’s own elements.

The [Components](@/components/_index.md) show each one live.
