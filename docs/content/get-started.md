+++
title = "Get Started"
description = "ciderui in a site: Tailwind CSS v4, its stylesheets imported, its classes on your markup."
+++

ciderui is apple.com’s look for a website built with Tailwind CSS v4: a small contract of variables (the colors, the sizes, the motion), the type as utilities, and apple.com’s parts as classes, each in a stylesheet of its own that a site imports when it uses the part. Everything it draws was measured on apple.com’s live pages; none of apple.com’s CSS is copied.

It has no JavaScript. A part that opens does so with HTML’s own elements (`<details>`, `popover`, a checkbox and its label); what only a script can do, a gallery’s paddles, is the page’s, and the component shows the state the page sets.

## Install

ciderui is not published. A site takes it as a checkout beside its own, a file dependency, with the Tailwind CLI:

```json
{
  "devDependencies": {
    "@tailwindcss/cli": "4.3.3",
    "ciderui": "file:../ciderui",
    "tailwindcss": "4.3.3"
  }
}
```

## The stylesheet

Import it after Tailwind, its fonts first, then the set the site uses: `site.css`, apple.com’s parts, or `console.css`, the console’s:

```css
@import "tailwindcss";
@import "ciderui/fonts.css";
@import "ciderui";
@import "ciderui/site.css";
```

A site that uses a few parts can import those alone, `ciderui/components/site/globalnav/globalnav.css` in place of `ciderui/site.css`.

`ciderui` is four parts, which a project can also take one by one:

| file | what it holds |
| --- | --- |
| `theme.css` | apple.com’s breakpoints, in place of Tailwind’s; the `dark:` variant; the semantic utilities (`text-body`, `bg-alt`, `text-link`, …) |
| `tokens.css` | the `--cui-*` variables: the colors, each a light and a dark value, and the radii, sizes, durations and materials |
| `base.css` | the page’s text and ground, the focus ring, dark and light blocks, reduced motion |
| `type.css` | the type’s classes, `cui-headline` to `cui-caption` |

Each component sits in Tailwind’s `components` layer, so a utility on the same element always wins: `class="cui-btn px-8"` gets the wider padding.

## The fonts

On Apple’s platforms the text is the system font, SF, which a page cannot ship; elsewhere it is Inter, which `fonts.css` declares and the browser fetches only where SF is not. Its `url()`s are relative to the compiled stylesheet: serve ciderui’s `fonts/` directory beside it.

## A first page

A component is a class on the element that is the part, its variants attributes, its states the ones the markup carries anyway (`aria-current`, `checked`, `disabled`, `[open]`):

```html
<section class="cui-section" data-variant="alt">
  <div class="cui-section-content">
    <h2 class="cui-headline">Hello.</h2>
    <a class="cui-btn" data-size="lg" href="/buy">Buy</a>
  </div>
</section>
```

The pages of [Components](@/components/_index.md) show each part live, over its markup. [Foundations](@/foundations.md) has the colors, the type and the breakpoints they are made of.

## Dark

There is no dark mode: apple.com has none. A block is dark when the page says so, `cui-dark` on any element, from a section to the whole page on `<html>`; inside it, every color takes its dark value, and `cui-light` makes a block inside it light again.

## In Rust, and with Zola

A site that renders its pages in Rust with hypertext writes the components’ markup with [the crate](@/rust.md), which also carries the stylesheets, so one dependency pins both. A static site is built with ciderui-zola, the Zola theme on ciderui, which ships its stylesheet, its fonts and its one script prebuilt: this documentation is built with it.
