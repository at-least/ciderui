+++
title = "Button"
description = "apple.com’s pill, on a link or a button: four colors, five sizes."

[extra]
category = "Controls"
image = "pictures/button.png"
+++

apple.com’s pill, from the store’s shared button styles and its marketing pages: a class on `<a href>` or `<button>`, its color a `data-variant`, its size a `data-size`.

```css
@import "ciderui/components/button/button.css";
```

{{ <cui_example file="button/default" title="A button" pad={true} /> }}

## Variants

| `data-variant` | what |
| --- | --- |
| (none) | the accent’s fill, white text |
| `secondary` | the link’s color as text and edge on no ground, filled with the accent hovered |
| `neutral` | the neutral’s fill: near-black, near-white in a dark block |
| `secondary-neutral` | the neutral’s color as text and edge, filled with it hovered: the store’s quick links |

{{ <cui_example file="button/variants" title="The four variants" pad={true} /> }}

## Sizes

apple.com’s pill comes in four heights: `data-size="sm"` 24px with 12px text, the default 36px with 14px, `lg` 44px and `xl` 56px with 17px. `cta`, a hero’s call to action, is 36px below 735px and 44px from it, as apple.com’s tile buttons step.

{{ <cui_example file="button/sizes" title="sm, the default, lg and xl" pad={true} /> }}

{{ <cui_example file="button/cta" title="A hero's calls to action" pad={true} /> }}

## In a dark block

Inside `cui-dark` the colors take their dark side: the accent stays, the link turns lighter, and the neutral turns near-white with black text, as on apple.com’s dark home page gallery.

{{ <cui_example file="button/dark" title="Buttons in a dark block" pad={true} /> }}

## Rectangle and block

`cui-btn-rect` gives it 12px corners and lets its text wrap; with `data-size="xl"` it is the store’s 56px form button. `cui-btn-block` is a block with 8px corners.

{{ <cui_example file="button/rect" title="The form buttons" pad={true} /> }}

{{ <cui_example file="button/block" title="A block button" pad={true} /> }}

## Disabled

A `disabled` button is dimmed to 0.42 and neither hovers nor presses.

{{ <cui_example file="button/disabled" title="Disabled buttons" pad={true} /> }}

No button transitions its colors, as none of apple.com’s does, and a focused one draws the page’s focus ring, 2px of the accent 3px out. Its width is the page’s.
