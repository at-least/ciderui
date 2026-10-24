//! ciderui's components as [hypertext] components: the markup the `cui-*`
//! classes expect, for a site that renders with hypertext's `rsx!` or
//! `maud!`; its icons, and every Lucide icon, as inline SVG ([`icon`],
//! [`lucide`]); and, with the `css` feature, its stylesheets embedded
//! ([`css`]), so a site's cargo dependency on this crate, at a commit, is
//! the one pin for both its markup and its CSS: the
//! site writes the stylesheets out (`ciderui::css::write`) and its Tailwind
//! entry imports the components it uses from there.
//!
//! Only the components whose markup is more than a class on one element are
//! here: a button is a class string, written on the element with the
//! attributes the page gives it (`<a class="cui-btn" data-size="cta">`); a
//! variant is an attribute, a state the one the markup carries anyway.

#![warn(missing_docs)]

mod accordion;
mod article;
mod card;
mod checkbox;
mod console_badge;
mod console_banner;
mod console_card;
mod console_checkbox;
mod console_chip;
mod console_dialog;
mod console_empty;
mod console_field;
mod console_hint;
mod console_menu;
mod console_sidebar;
#[cfg(feature = "css")]
pub mod css;
mod field;
mod footer;
mod footnote;
mod gallery;
mod globalnav;
pub mod icon;
mod link;
mod localnav;
pub mod lucide;
mod modal;
mod option;
mod posts;
mod productnav;
mod quicklinks;
mod ribbon;
mod section;
mod shelf;
mod stat;
mod stories;
mod swatch;
mod tile;

pub use accordion::{Accordion, AccordionItem};
pub use article::Article;
pub use card::Card;
pub use checkbox::Checkbox;
pub use console_badge::{ConsoleBadge, ConsoleBadgeVariant, ConsoleColor};
pub use console_banner::ConsoleBanner;
pub use console_card::ConsoleMetric;
pub use console_checkbox::ConsoleCheckboxBox;
pub use console_chip::ConsoleChip;
pub use console_dialog::ConsoleDialog;
pub use console_empty::{ConsoleEmpty, ConsoleEmptySize};
pub use console_field::ConsoleSelect;
pub use console_hint::ConsoleHint;
pub use console_menu::{ConsoleMenu, ConsoleMenuGroup, ConsoleMenuItem};
pub use console_sidebar::{ConsoleSidebar, ConsoleSidebarItem};
pub use field::{Field, Select, TextArea};
pub use footer::{Footer, FooterSection};
pub use footnote::Footnote;
pub use gallery::Gallery;
pub use globalnav::{Bag, BagLink, FlyoutGroup, GlobalNav, NavLink, Search};
pub use link::MoreLink;
pub use localnav::LocalNav;
pub use modal::{Modal, ModalButton};
pub use option::OptionTile;
pub use posts::{Post, PostSize, Posts};
pub use productnav::{ProductNav, ProductNavItem};
pub use quicklinks::QuickLinks;
pub use ribbon::Ribbon;
pub use section::{Ground, Rhythm, Section, SectionHeader};
pub use shelf::Shelf;
pub use stat::Stat;
pub use stories::{Stories, Story, StorySize};
pub use swatch::Swatch;
pub use tile::{Tile, TileSize};

#[cfg(test)]
mod example;

use hypertext::{AttributeBuffer, context::AttributeValue, prelude::*};

/// A component's root class: its block (with its modifiers, `cui-dark`),
/// then the page's utilities (its `class` field), if any.
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
