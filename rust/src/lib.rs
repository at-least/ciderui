//! ciderui's components as [hypertext] components: the markup the `cui-*`
//! classes expect, for a site that renders with hypertext's `rsx!` or
//! `maud!`; its icons as inline SVG ([`icon`]); and, with the `css` feature,
//! its stylesheets embedded ([`css`]), so a site's cargo dependency on this
//! crate, at a commit, is the one pin for both its markup and its CSS: the
//! site writes the stylesheets out (`ciderui::css::write`) and its Tailwind
//! entry imports the components it uses from there.
//!
//! Only the components whose markup is more than a class on one element are
//! here: a button is a class string, written on the element with the
//! attributes the page gives it (`<a class="cui-btn" data-size="cta">`); a
//! variant is an attribute, a state the one the markup carries anyway.

#![warn(missing_docs)]

mod card;
#[cfg(feature = "css")]
pub mod css;
mod gallery;
mod globalnav;
pub mod icon;
mod link;
mod localnav;
mod ribbon;
mod section;
mod shelf;
mod tile;

pub use card::Card;
pub use gallery::Gallery;
pub use globalnav::GlobalNav;
pub use link::MoreLink;
pub use localnav::LocalNav;
pub use ribbon::Ribbon;
pub use section::{Ground, Rhythm, Section, SectionHeader};
pub use shelf::Shelf;
pub use tile::{Tile, TileSize};

#[cfg(test)]
mod example;

use hypertext::{AttributeBuffer, context::AttributeValue, prelude::*};

/// A component's root class: its block, then the page's utilities (its
/// `class` field), if any.
struct Class<'a>(&'static str, &'a str);

impl Renderable<AttributeValue> for Class<'_> {
    fn render_to(&self, buffer: &mut AttributeBuffer) {
        self.0.render_to(buffer);
        if !self.1.is_empty() {
            " ".render_to(buffer);
            self.1.render_to(buffer);
        }
    }
}
