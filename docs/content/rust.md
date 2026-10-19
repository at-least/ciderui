+++
title = "Rust"
description = "The crate: the components’ markup as hypertext components, the icons, and the stylesheets."
+++

A site that renders its pages in Rust with [hypertext](https://crates.io/crates/hypertext) writes ciderui’s markup with the crate `ciderui` (in the repository’s `rust/`), and takes its stylesheets from the same crate, so its dependency, at one commit, pins both: the markup and the CSS cannot part.

```toml
[dependencies]
ciderui = { path = "../ciderui/rust", features = ["css"] }
hypertext = "0.12.1"
```

## The components

The crate holds a component for each part whose markup is more than a class on one element; a button, a link in running text or a price is a class string on the page’s own element. Each component’s fields are the part’s: its text, its links, its variant as an enum, and `class`, the page’s utilities on its root (`""` for none).

```rust
use hypertext::prelude::*;
use ciderui::{Ground, Rhythm, Section};

let html = rsx! {
    <Section ground=(Ground::Alt) rhythm=(Rhythm::Large) dark=true class="">
        <p>"Content"</p>
    </Section>
}
.render();
assert_eq!(
    html.as_inner(),
    concat!(
        r#"<section class="cui-section cui-dark" data-variant="alt" data-size="lg">"#,
        r#"<div class="cui-section-content"><p>Content</p></div></section>"#,
    ),
);
```

| component | the part |
| --- | --- |
| `MoreLink` | a more link and its chevron |
| `Section`, `SectionHeader` | a band and its header |
| `GlobalNav`, `LocalNav` | the global nav, a product’s local nav |
| `Ribbon` | the message above a page |
| `Tile`, `Card` | a home page tile, a store card |
| `Shelf`, `Gallery` | a row of cards, a product page’s gallery |
| `Field`, `Checkbox`, `OptionTile`, `Swatch` | a form’s controls |
| `Accordion`, `AccordionItem` | questions and answers |
| `QuickLinks`, `ProductNav`, `ProductNavItem` | the store’s links and its products |
| `Footer`, `FooterSection` | the footer |
| `Stat`, `Footnote` | a number, a footnote’s mark |
| `Article`, `Stories`, `Story`, `Posts`, `Post` | the Newsroom’s story, index and archive |

Each one’s test renders its example, the markup this documentation shows, and asserts the output equal to it.

## Icons

`ciderui::icon` holds ciderui’s icons as inline SVG in `currentColor`, each drawn from the numbers of the glyph apple.com draws in its place: the more link’s chevron, the global nav’s search, bag and menu, the paddles’ chevrons, a checkbox’s check.

```rust
use hypertext::prelude::*;
use ciderui::icon::MORE;

let html = rsx! { <a href="/iphone/">"Learn more"(MORE.class("cui-link-icon"))</a> }.render();
assert!(html.as_inner().contains(r#"<svg class="cui-link-icon" aria-hidden="true""#));
```

## The stylesheets

With the `css` feature the crate embeds every stylesheet, and writes them out in the kit’s layout before the site’s Tailwind build:

```rust
ciderui::css::write("target/ciderui").unwrap();
```

The site’s entry then imports from there:

```css
@import "tailwindcss";
@import "../target/ciderui/fonts.css";
@import "../target/ciderui/index.css";
@import "../target/ciderui/components/button/button.css";
```

The binary `ciderui-css <dir>` does the same from a shell. The Inter files are not embedded: copy ciderui’s `fonts/` beside the compiled stylesheet once.
