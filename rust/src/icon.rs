//! Icons as inline SVG, ciderui's own, each drawn from the numbers of the
//! glyph or the drawing apple.com uses in its place (`reference/icons/`): its
//! box, the ink's place in it, the strokes' thickness; not Apple's outlines.
//! Each is sized by its own attributes: in em where it sits in text, so it
//! takes the text's size ([`MORE`]), in pixels where it sits in the global
//! nav's bar, as apple.com's do; and drawn in `currentColor`, so it takes the
//! text's color. A component's class (`cui-link-icon`) places it. Inline, so
//! a page whose Content-Security-Policy refuses `data:` images draws it too.
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

/// The global nav's search from 834px (`cui-globalnav-image` in a
/// `cui-globalnav-icon`), in place of apple.com's 15 by 44 SVG: a ring at
/// (5.95, 19.95), radius 4.96, 1.08 thick, and its handle from the ring at
/// 45° to (13.74, 27.74), 1.52 thick, round; 0.927 of its ink apple.com's.
pub const SEARCH: Icon = Icon {
    width: "15",
    height: "44",
    view_box: "0 0 15 44",
    inner: r#"<circle cx="5.95" cy="19.95" r="4.96" fill="none" stroke="currentColor" stroke-width="1.08"/><path d="M9.46 23.46 13.74 27.74" fill="none" stroke="currentColor" stroke-width="1.52" stroke-linecap="round"/>"#,
};

/// The global nav's search to 833px (`cui-globalnav-image-compact`), in
/// place of apple.com's 17 by 48 SVG: a ring at (6.94, 21.94), radius 5.86,
/// 1.22 thick, and its handle to (15.6, 30.6), 1.8 thick, round; 0.972 of
/// its ink apple.com's.
pub const SEARCH_COMPACT: Icon = Icon {
    width: "17",
    height: "48",
    view_box: "0 0 17 48",
    inner: r#"<circle cx="6.94" cy="21.94" r="5.86" fill="none" stroke="currentColor" stroke-width="1.22"/><path d="M11.08 26.08 15.6 30.6" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/>"#,
};

/// The global nav's bag from 834px (`cui-globalnav-image`), in place of
/// apple.com's 14 by 44 SVG: a body 11.8 by 10.9 from (1.1, 16.55), corners
/// 1.5, and its handle an arc of radius 2.45, both 1.1 thick; 0.948 of its
/// ink apple.com's.
pub const BAG: Icon = Icon {
    width: "14",
    height: "44",
    view_box: "0 0 14 44",
    inner: r#"<rect x="1.1" y="16.55" width="11.8" height="10.9" rx="1.5" fill="none" stroke="currentColor" stroke-width="1.1"/><path d="M4.55 16.55V16.05a2.45 2.45 0 0 1 4.9 0v.5" fill="none" stroke="currentColor" stroke-width="1.1"/>"#,
};

/// The global nav's bag to 833px (`cui-globalnav-image-compact`), in place
/// of apple.com's 17 by 48 SVG: a body 13.9 by 12.95 from (1.55, 17.5),
/// corners 2, and its handle an arc of radius 3.125, both 1.2 thick; 0.981
/// of its ink apple.com's.
pub const BAG_COMPACT: Icon = Icon {
    width: "17",
    height: "48",
    view_box: "0 0 17 48",
    inner: r#"<rect x="1.55" y="17.5" width="13.9" height="12.95" rx="2" fill="none" stroke="currentColor" stroke-width="1.2"/><path d="M5.375 17.5V17.175a3.125 3.125 0 0 1 6.25 0v.325" fill="none" stroke="currentColor" stroke-width="1.2"/>"#,
};

/// The global nav's menu button's lines (`cui-globalnav-menubutton`), as
/// apple.com's: two lines across 18 by 18, from 2 to 16 at 5 and 12 down,
/// 1.2 thick, round; in the opened menu's close button
/// (`cui-globalnav-closebutton`) `globalnav.css` turns them into its X (from
/// (3.5, 3.5) to (15, 15) and from (3.5, 15) to (15, 3.5)).
pub const MENU: Icon = Icon {
    width: "18",
    height: "18",
    view_box: "0 0 18 18",
    inner: r#"<path d="M2 5H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/><path d="M2 12H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/>"#,
};

/// The local nav's chevron (`cui-localnav-menubutton`), as apple.com's: a
/// line through (15.265, 0.835), (8, 8.167) and (0.735, 0.835) in 16 by 9,
/// 1.2 thick, round; `localnav.css` turns it up while the menu is open, which
/// is the line apple.com's animation ends on (`reference/icons/`).
pub const CHEVRON: Icon = Icon {
    width: "16",
    height: "9",
    view_box: "0 0 16 9",
    inner: r#"<path d="M15.265.835 8 8.167.735.835" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round"/>"#,
};

/// A gallery's previous paddle's chevron (`cui-gallery-paddle`), in place of
/// apple.com's 36 by 36 SVG: a line through (20, 12.5), (14.5, 18) and
/// (20, 23.5), 3 thick, round; 0.9995 of its ink apple.com's
/// (`reference/icons/`). Its size the paddle's (`gallery.css`).
pub const PADDLE_PREVIOUS: Icon = Icon {
    width: "36",
    height: "36",
    view_box: "0 0 36 36",
    inner: r#"<path d="M20 12.5 14.5 18 20 23.5" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>"#,
};

/// A gallery's next paddle's chevron: [`PADDLE_PREVIOUS`] turned, through
/// (16, 12.5), (21.5, 18) and (16, 23.5); 0.9993 of its ink apple.com's.
pub const PADDLE_NEXT: Icon = Icon {
    width: "36",
    height: "36",
    view_box: "0 0 36 36",
    inner: r#"<path d="M16 12.5 21.5 18 16 23.5" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>"#,
};

/// A checkbox's check (`cui-checkbox-indicator`), in place of apple.com's SF
/// Pro Icons glyph: a line through (3.3, 7.65), (6.25, 11.35) and
/// (11.15, 3.7) in the box's 14 by 14, 1.35 thick, round; 0.96 of its ink
/// apple.com's (`reference/icons/`).
pub const CHECK: Icon = Icon {
    width: "14",
    height: "14",
    view_box: "0 0 14 14",
    inner: r#"<path d="M3.3 7.65 6.25 11.35 11.15 3.7" fill="none" stroke="currentColor" stroke-width="1.35" stroke-linecap="round" stroke-linejoin="round"/>"#,
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
