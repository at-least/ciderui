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

#[cfg(feature = "css")]
pub mod css;
pub mod icon;

#[cfg(test)]
mod example;
