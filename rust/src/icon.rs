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

/// The global nav's search field's magnifier from 834px (`cui-globalnav-image`
/// in a `cui-globalnav-searchfield`), in place of apple.com's 30 by 32 SVG: a
/// ring at (13.475, 14.775), radius 6.225, 1.3 thick, and its handle from the
/// ring at 45° to (22.57, 24.07), 2.2 thick, round; 0.974 of its ink
/// apple.com's.
pub const SEARCHFIELD: Icon = Icon {
    width: "30",
    height: "32",
    view_box: "0 0 30 32",
    inner: r#"<circle cx="13.475" cy="14.775" r="6.225" fill="none" stroke="currentColor" stroke-width="1.3"/><path d="M18.11 19.61 22.57 24.07" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round"/>"#,
};

/// The search field's magnifier to 833px (`cui-globalnav-image-compact`), in
/// place of apple.com's 38 by 40 SVG: a ring at (17.225, 17.9), radius 7.25,
/// 1.45 thick, and its handle to (27.74, 28.74), 2.62 thick, round; 0.970 of
/// its ink apple.com's.
pub const SEARCHFIELD_COMPACT: Icon = Icon {
    width: "38",
    height: "40",
    view_box: "0 0 38 40",
    inner: r#"<circle cx="17.225" cy="17.9" r="7.25" fill="none" stroke="currentColor" stroke-width="1.45"/><path d="M22.6 23.6 27.74 28.74" fill="none" stroke="currentColor" stroke-width="2.62" stroke-linecap="round"/>"#,
};

/// A search quick link's arrow from 834px (`cui-globalnav-image` in a
/// `cui-globalnav-searchlink`), in place of apple.com's 9 by 16 SVG: a shaft
/// from (0.75, 7.75) to its head at (8.25, 7.75), the head's two strokes 3
/// back at 45°, 1 thick, round; 0.983 of its ink apple.com's.
pub const ARROW: Icon = Icon {
    width: "9",
    height: "16",
    view_box: "0 0 9 16",
    inner: r#"<path d="M.75 7.75H8.25M5.25 4.75 8.25 7.75 5.25 10.75" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round" stroke-linejoin="round"/>"#,
};

/// The quick link's arrow to 833px (`cui-globalnav-image-compact`), in place
/// of apple.com's 13 by 25 SVG: a shaft from (1.1, 13) to its head at
/// (11.93, 13), 1.3 thick, the head's strokes 4.45 back at 45°, 1.2 thick,
/// round; 0.995 of its ink apple.com's.
pub const ARROW_COMPACT: Icon = Icon {
    width: "13",
    height: "25",
    view_box: "0 0 13 25",
    inner: r#"<path d="M1.1 13H11.93" fill="none" stroke="currentColor" stroke-width="1.3" stroke-linecap="round"/><path d="M7.48 8.55 11.93 13 7.48 17.45" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round"/>"#,
};

/// The search field's clear mark from 834px (`cui-globalnav-image` in a
/// `cui-globalnav-searchreset`), in place of apple.com's 14 by 14 SVG: a disc
/// at (7, 7), radius 6.9661, and cut out of it an X, two strokes 1.1 thick
/// between round ends whose middles are 2.409 from the disc's middle along
/// each diagonal, drawn as one outline (evenodd); 0.9995 of its ink
/// apple.com's.
pub const CLEAR: Icon = Icon {
    width: "14",
    height: "14",
    view_box: "0 0 14 14",
    inner: r#"<path d="M7 .0339A6.9661 6.9661 0 1 0 7 13.9661 6.9661 6.9661 0 1 0 7 .0339ZM7.7778 7 9.798 9.0201A.55.55 0 0 1 9.0201 9.798L7 7.7778 4.9799 9.798A.55.55 0 0 1 4.2021 9.0201L6.2222 7 4.2021 4.9799A.55.55 0 0 1 4.9799 4.2021L7 6.2222 9.0201 4.2021A.55.55 0 0 1 9.798 4.9799Z" fill="currentColor" fill-rule="evenodd"/>"#,
};

/// The clear mark to 833px (`cui-globalnav-image-compact`), in place of
/// apple.com's 16 by 16 SVG: a disc at (8, 8), radius 8, its X's strokes
/// 1.34 thick, their round ends' middles 2.81 from the disc's middle; 0.9928
/// of its ink apple.com's.
pub const CLEAR_COMPACT: Icon = Icon {
    width: "16",
    height: "16",
    view_box: "0 0 16 16",
    inner: r#"<path d="M8 0A8 8 0 1 0 8 16 8 8 0 1 0 8 0ZM8.9475 8 11.2838 10.3362A.67.67 0 0 1 10.3362 11.2838L8 8.9475 5.6638 11.2838A.67.67 0 0 1 4.7162 10.3362L7.0525 8 4.7162 5.6638A.67.67 0 0 1 5.6638 4.7162L8 7.0525 10.3362 4.7162A.67.67 0 0 1 11.2838 5.6638Z" fill="currentColor" fill-rule="evenodd"/>"#,
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

/// An accordion's chevron (`cui-accordion-question`), as apple.com's: a line
/// through (15, 1.13), (8.5, 7.72) and (2, 1.13) in 17 by 8.85, 2.25 thick,
/// round (`reference/icons/`); its box square, the line in its middle, sized
/// by `accordion.css`, which turns it up while its answer is open.
pub const DISCLOSURE: Icon = Icon {
    width: "17",
    height: "17",
    view_box: "0 0 17 8.85",
    inner: r#"<path d="M15 1.13 8.5 7.72 2 1.13" fill="none" stroke="currentColor" stroke-width="2.25" stroke-linecap="round" stroke-linejoin="round"/>"#,
};

/// The footer's chevron (`cui-footer-title`, to 833px), as apple.com's: a
/// line through (10.075, 0.675), (5.5, 5.323) and (0.925, 0.675) in 11 by
/// 6, 1 thick, round (`reference/icons/`); `footer.css` turns it up while
/// its section is open, the line apple.com's animation ends on.
pub const FOOTER_CHEVRON: Icon = Icon {
    width: "11",
    height: "6",
    view_box: "0 0 11 6",
    inner: r#"<path d="M10.075.675 5.5 5.323.925.675" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round" stroke-linejoin="round"/>"#,
};

/// A field's error mark (`cui-field-message`), in place of apple.com's SF
/// Pro Icons glyph: a ring at (0.585em, 0.4975em down), radius 0.456em,
/// 0.0825em thick, and in it a stem from 0.28 to 0.55em down, 0.07em thick,
/// round, over a dot of radius 0.05em at 0.6925em down, in the glyph's
/// advance (1.171em) by its ink's height (0.995em); 0.9695 of its ink
/// apple.com's (`reference/icons/`). In em: it takes the message's size.
pub const ALERT: Icon = Icon {
    width: "1.171em",
    height: "0.995em",
    view_box: "0 0 1171 995",
    inner: r#"<circle cx="585" cy="497.5" r="456.25" fill="none" stroke="currentColor" stroke-width="82.5"/><path d="M584 280V550" fill="none" stroke="currentColor" stroke-width="70" stroke-linecap="round"/><circle cx="585" cy="692.5" r="50" fill="currentColor"/>"#,
};

/// A dropdown's chevron (`cui-field-chevron`), in place of the store's
/// background SVG: a line through (6.075, 9.275), (10.5, 13.7) and
/// (14.925, 9.275) in 21 by 21, 1.45 thick, round, drawn 25 by 25 as
/// apple.com's; 0.9689 of its ink apple.com's (`reference/icons/`).
pub const DROPDOWN: Icon = Icon {
    width: "25",
    height: "25",
    view_box: "0 0 21 21",
    inner: r#"<path d="M6.075 9.275 10.5 13.7 14.925 9.275" fill="none" stroke="currentColor" stroke-width="1.45" stroke-linecap="round" stroke-linejoin="round"/>"#,
};

/// A modal's X (`cui-modal-close`), in place of the store overlay's SVG: two
/// lines from (4.88, 4.88) to (15.12, 15.12) and from (15.12, 4.88) to
/// (4.88, 15.12) in 20 by 20, 3 thick, round; 0.9927 of its ink apple.com's
/// (`reference/icons/`).
pub const CLOSE: Icon = Icon {
    width: "20",
    height: "20",
    view_box: "0 0 20 20",
    inner: r#"<path d="M4.88 4.88 15.12 15.12M15.12 4.88 4.88 15.12" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round"/>"#,
};

/// The marketing pages' modal's X (`cui-modal-close` in a
/// `data-variant="page"` modal), in place of its SVG: [`CLOSE`]'s strokes,
/// each arm 0.62 shorter at its ends, from (5.5, 5.5) to (14.5, 14.5) and
/// from (14.5, 5.5) to (5.5, 14.5); 0.9994 of its ink apple.com's
/// (`reference/icons/`).
pub const CLOSE_SMALL: Icon = Icon {
    width: "20",
    height: "20",
    view_box: "0 0 20 20",
    inner: r#"<path d="M5.5 5.5 14.5 14.5M14.5 5.5 5.5 14.5" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round"/>"#,
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
