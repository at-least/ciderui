//! Icons as inline SVG, ciderui's own, each drawn from the numbers of the
//! glyph or the drawing apple.com uses in its place (`reference/icons/`): its
//! box in em, the ink's place in it, the strokes' thickness; not Apple's
//! outlines. Each is sized in em by its own attributes, so it takes the
//! text's size, and drawn in `currentColor`, so it takes the text's color; a
//! component's class (`cui-link-icon`) places it. Inline, so a page whose
//! Content-Security-Policy refuses `data:` images draws it too.
//!
//! ```
//! use hypertext::prelude::*;
//! use ciderui::icon::MORE;
//!
//! let html = rsx! { <a href="/iphone/">"Learn more"(MORE.class("cui-link-icon"))</a> }.render();
//! assert!(html.as_inner().contains(r#"<svg class="cui-link-icon" aria-hidden="true""#));
//! ```
use hypertext::{Buffer, Raw, prelude::*};

/// An icon: its box in em, its view box, and what is drawn in it.
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub struct Icon {
    width: &'static str,
    height: &'static str,
    view_box: &'static str,
    inner: &'static str,
}

impl Icon {
    /// The icon with a class on its `<svg>`: a component's part
    /// (`cui-link-icon`) or a utility.
    pub fn class<'a>(&self, class: &'a str) -> Classed<'a> {
        Classed { icon: *self, class }
    }

    // hypertext's rsx! knows the HTML elements, not SVG's: the markup is
    // written out, the class escaped as an attribute value
    fn write(&self, class: Option<&str>, buffer: &mut Buffer) {
        let mut svg = String::from("<svg");
        if let Some(class) = class {
            svg.push_str(" class=\"");
            for ch in class.chars() {
                match ch {
                    '&' => svg.push_str("&amp;"),
                    '"' => svg.push_str("&quot;"),
                    '<' => svg.push_str("&lt;"),
                    _ => svg.push(ch),
                }
            }
            svg.push('"');
        }
        svg.push_str(&format!(
            r#" aria-hidden="true" width="{}" height="{}" viewBox="{}">"#,
            self.width, self.height, self.view_box
        ));
        svg.push_str(self.inner);
        svg.push_str("</svg>");
        Raw::<_>::dangerously_create(svg).render_to(buffer);
    }
}

impl Renderable for Icon {
    fn render_to(&self, buffer: &mut Buffer) {
        self.write(None, buffer);
    }
}

/// An [`Icon`] with a class on its `<svg>` ([`Icon::class`]).
#[derive(Clone, Copy, Debug)]
pub struct Classed<'a> {
    icon: Icon,
    class: &'a str,
}

impl Renderable for Classed<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        self.icon.write(Some(self.class), buffer);
    }
}

/// The "more" chevron after a link's text (`cui-link`, `data-variant="more"`),
/// in place of apple.com's SF Pro Icons U+F301: its advance, 0.4419em, as the
/// box's width; from 0.01em below the baseline to 0.53em above it; two
/// strokes 0.0679em thick from its ink's corners to the middle, round at the
/// ends and the join, which share 0.948 of their ink with apple.com's glyph
/// (`reference/icons/`). The view box is in thousandths of an em.
pub const MORE: Icon = Icon {
    width: "0.442em",
    height: "0.54em",
    view_box: "0 0 442 540",
    inner: r#"<path d="M107.7 47.7 332.3 272.5 107.7 497.3" fill="none" stroke="currentColor" stroke-width="67.9" stroke-linecap="round" stroke-linejoin="round"/>"#,
};

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn renders_an_svg_in_em() {
        let html = rsx! { (MORE) }.render();
        assert_eq!(
            html.as_inner(),
            concat!(
                r#"<svg aria-hidden="true" width="0.442em" height="0.54em" viewBox="0 0 442 540">"#,
                r#"<path d="M107.7 47.7 332.3 272.5 107.7 497.3" fill="none" stroke="currentColor" stroke-width="67.9" stroke-linecap="round" stroke-linejoin="round"/></svg>"#,
            )
        );
    }

    #[test]
    fn escapes_its_class() {
        let html = rsx! { (MORE.class(r#"a"b<c&d"#)) }.render();
        assert!(
            html.as_inner()
                .starts_with(r#"<svg class="a&quot;b&lt;c&amp;d""#)
        );
    }
}
