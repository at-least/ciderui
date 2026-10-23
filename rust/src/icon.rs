//! Icons as inline SVG, ciderui's own, each drawn from the numbers of the
//! glyph or the drawing apple.com uses in its place (`reference/icons/`): its
//! box, the ink's place in it, the strokes' thickness; not Apple's outlines.
//! Each is sized by its own attributes: in em where it sits in text, so it
//! takes the text's size ([`MORE`]), in pixels elsewhere (the global nav, a
//! gallery's paddles, a modal's X), as apple.com's do; and drawn in `currentColor`, so it takes the
//! text's color. A component's class (`cui-link-icon`) places it. Inline, so
//! a page whose Content-Security-Policy refuses `data:` images draws it too.
//! Every other icon is Lucide's ([`crate::lucide`]), an [`Icon`] too.
//!
//! ```
//! use hypertext::prelude::*;
//! use ciderui::icon::MORE;
//!
//! let html = rsx! { <a href="#learn">"Learn more"(MORE.class("cui-link-icon"))</a> }.render();
//! assert!(html.as_inner().contains(r#"<svg class="cui-link-icon" aria-hidden="true""#));
//! ```
use hypertext::{Buffer, Raw, prelude::*};

/// An icon: its box (in em or pixels), its view box, the attributes on its
/// `<svg>` that what is drawn in it inherits, and what is drawn in it.
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub struct Icon {
    width: &'static str,
    height: &'static str,
    view_box: &'static str,
    paint: &'static str,
    inner: &'static str,
}

impl Icon {
    /// A Lucide icon ([`crate::lucide`]): its elements, under the attributes
    /// Lucide's `<svg>` carries.
    pub(crate) const fn lucide(inner: &'static str) -> Self {
        Self {
            width: "1em",
            height: "1em",
            view_box: crate::lucide::VIEW_BOX,
            // on the <svg>, not on each element: a class on the <svg>
            // (stroke-[1.5]) wins over an attribute there, and its elements
            // inherit what wins
            paint: crate::lucide::PAINT,
            inner,
        }
    }

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
            r#" aria-hidden="true" width="{}" height="{}" viewBox="{}""#,
            self.width, self.height, self.view_box
        ));
        if !self.paint.is_empty() {
            svg.push(' ');
            svg.push_str(self.paint);
        }
        svg.push('>');
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
    paint: "",
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
    paint: "",
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
    paint: "",
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
    paint: "",
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
    paint: "",
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
    paint: "",
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
    paint: "",
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
    paint: "",
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
    paint: "",
    inner: r#"<path d="M1.1 13H11.93" fill="none" stroke="currentColor" stroke-width="1.3" stroke-linecap="round"/><path d="M7.48 8.55 11.93 13 7.48 17.45" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round"/>"#,
};

/// A suggested search's magnifier from 834px (`cui-globalnav-image` in a
/// `cui-globalnav-searchlink`, where a quick link has its arrow), in place of
/// apple.com's 9 by 16 SVG: a ring at (3.675, 6.9), radius 3.125, 0.7 thick,
/// and its handle at 45° from just inside the ring to (8.2, 11.581), 1.1
/// thick, round; 0.964 of its ink apple.com's.
pub const SUGGESTION: Icon = Icon {
    width: "9",
    height: "16",
    view_box: "0 0 9 16",
    paint: "",
    inner: r#"<circle cx="3.675" cy="6.9" r="3.125" fill="none" stroke="currentColor" stroke-width="0.7"/><path d="M6 9.381 8.2 11.581" fill="none" stroke="currentColor" stroke-width="1.1" stroke-linecap="round"/>"#,
};

/// A suggested search's magnifier to 833px (`cui-globalnav-image-compact`),
/// in place of apple.com's 13 by 25 SVG: a ring at (5.425, 11.725), radius
/// 4.425, 0.9 thick, and its handle to (11.8, 18.35), 1.6 thick, round; 0.979
/// of its ink apple.com's.
pub const SUGGESTION_COMPACT: Icon = Icon {
    width: "13",
    height: "25",
    view_box: "0 0 13 25",
    paint: "",
    inner: r#"<circle cx="5.425" cy="11.725" r="4.425" fill="none" stroke="currentColor" stroke-width="0.9"/><path d="M8.692 15.242 11.8 18.35" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round"/>"#,
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
    paint: "",
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
    paint: "",
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
    paint: "",
    inner: r#"<path d="M2 5H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/><path d="M2 12H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/>"#,
};

/// The chevron of a flyout's button from 834px (`cui-globalnav-image` in a
/// `cui-globalnav-flyoutbutton`), shown when it has the focus from the
/// keyboard, in place of apple.com's 10 by 10 SVG: a line through (2.33,
/// 3.66), (5, 6.34) and (7.67, 3.66), 0.88 thick, round; 0.99 of its ink
/// apple.com's (`reference/icons/`).
pub const FLYOUT: Icon = Icon {
    width: "10",
    height: "10",
    view_box: "0 0 10 10",
    paint: "",
    inner: r#"<path d="M2.33 3.66 5 6.34 7.67 3.66" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".88"/>"#,
};

/// The same button's chevron to 833px (`cui-globalnav-image-compact`), at
/// the end of the menu's row, shown hovered, in place of apple.com's 9 by 48
/// SVG: a line through (1.3, 17.21), (8.3, 24.07) and (1.3, 30.79), 1.2
/// thick, round; 0.9968 of its ink apple.com's, and 0.9962 of the chevron a
/// hovered row of apple.com's shows (`reference/icons/`).
pub const FLYOUT_COMPACT: Icon = Icon {
    width: "9",
    height: "48",
    view_box: "0 0 9 48",
    paint: "",
    inner: r#"<path d="M1.3 17.21 8.3 24.07 1.3 30.79" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"/>"#,
};

/// A flyout's back button to 833px (`cui-globalnav-backbutton`), in place of
/// apple.com's 9 by 48 SVG: [`FLYOUT_COMPACT`] turned, through (7.7, 17.21),
/// (0.7, 24.07) and (7.7, 30.79); 0.9962 of its ink apple.com's
/// (`reference/icons/`).
pub const FLYOUT_BACK: Icon = Icon {
    width: "9",
    height: "48",
    view_box: "0 0 9 48",
    paint: "",
    inner: r#"<path d="M7.7 17.21 .7 24.07 7.7 30.79" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"/>"#,
};

/// The local nav's chevron (`cui-localnav-menubutton`), as apple.com's: a
/// line through (15.265, 0.835), (8, 8.167) and (0.735, 0.835) in 16 by 9,
/// 1.2 thick, round; `localnav.css` turns it up while the menu is open, which
/// is the line apple.com's animation ends on (`reference/icons/`).
pub const CHEVRON: Icon = Icon {
    width: "16",
    height: "9",
    view_box: "0 0 16 9",
    paint: "",
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
    paint: "",
    inner: r#"<path d="M20 12.5 14.5 18 20 23.5" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>"#,
};

/// A gallery's next paddle's chevron: [`PADDLE_PREVIOUS`] turned, through
/// (16, 12.5), (21.5, 18) and (16, 23.5); 0.9993 of its ink apple.com's.
pub const PADDLE_NEXT: Icon = Icon {
    width: "36",
    height: "36",
    view_box: "0 0 36 36",
    paint: "",
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
    paint: "",
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
    paint: "",
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
    paint: "",
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
    paint: "",
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
    paint: "",
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
    paint: "",
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
    paint: "",
    inner: r#"<path d="M5.5 5.5 14.5 14.5M14.5 5.5 5.5 14.5" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round"/>"#,
};

/// A bag's profile link to its orders (`cui-globalnav-baglink`, from 834px), in
/// place of apple.com's 11 by 16 SVG, a box: a hexagon through (5.511, 2.548),
/// (9.738, 4.933), (9.738, 9.684), (5.511, 12.068), (1.285, 9.684) and (1.285,
/// 4.933), its corners rounded 1.0625, the top's lower edges meeting at (5.511,
/// 7.375) over a line down to its foot, and a band across its top from (3.849,
/// 3.988) to (7.254, 5.878), 0.875 thick; 0.9915 of its ink apple.com's
/// (`reference/icons/`).
pub const ORDERS: Icon = Icon {
    width: "11",
    height: "16",
    view_box: "0 0 11 16",
    paint: "",
    inner: r#"<path d="M6.033 2.843 9.198 4.628A1.0625 1.0625 0 0 1 9.738 5.553V9.063A1.0625 1.0625 0 0 1 9.198 9.989L6.033 11.774A1.0625 1.0625 0 0 1 4.989 11.774L1.825 9.989A1.0625 1.0625 0 0 1 1.285 9.063V5.553A1.0625 1.0625 0 0 1 1.825 4.628L4.989 2.843A1.0625 1.0625 0 0 1 6.033 2.843ZM1.43 5.018 5.511 7.375 9.593 5.018M5.511 7.375V11.91" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".875"/><path d="M3.849 3.988 7.254 5.878" fill="none" stroke="currentColor" stroke-width=".875"/>"#,
};

/// The orders' box to 833px, in place of apple.com's 16 by 25 SVG: the same
/// hexagon through (8, 5.74), (14, 9.125), (14, 15.875), (8, 19.259), (2, 15.875)
/// and (2, 9.125), its corners rounded 1.7, the top's edges meeting at (8,
/// 12.594), the band from (5.561, 7.806) to (10.456, 10.522), 1.2 thick; 0.9851
/// of its ink apple.com's (`reference/icons/`).
pub const ORDERS_COMPACT: Icon = Icon {
    width: "16",
    height: "25",
    view_box: "0 0 16 25",
    paint: "",
    inner: r#"<path d="M8.835 6.211 13.135 8.637A1.7 1.7 0 0 1 14 10.118V14.882A1.7 1.7 0 0 1 13.135 16.363L8.835 18.788A1.7 1.7 0 0 1 7.165 18.788L2.865 16.363A1.7 1.7 0 0 1 2 14.882V10.118A1.7 1.7 0 0 1 2.865 8.637L7.165 6.211A1.7 1.7 0 0 1 8.835 6.211ZM2.232 9.26 8 12.594 13.768 9.26M8 12.594V19.007" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"/><path d="M5.561 7.806 10.456 10.522" fill="none" stroke="currentColor" stroke-width="1.2"/>"#,
};

/// A bag's profile link to its saved items (`cui-globalnav-baglink`, from
/// 834px), in place of apple.com's 11 by 16 SVG, a bookmark: its sides at 2.719
/// and 8.281 from 2.789 down to 11.97, its top corners rounded 1.184, its foot
/// a notch up to (5.5, 9.3), 0.82 thick, round; 0.9518 of its ink apple.com's
/// (`reference/icons/`).
pub const SAVES: Icon = Icon {
    width: "11",
    height: "16",
    view_box: "0 0 11 16",
    paint: "",
    inner: r#"<path d="M2.719 11.97V3.972A1.184 1.184 0 0 1 3.903 2.789H7.097A1.184 1.184 0 0 1 8.281 3.972V11.97L5.5 9.3Z" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".82"/>"#,
};

/// The saved items' bookmark to 833px, in place of apple.com's 16 by 25 SVG:
/// its sides at 4 and 12 from 5.75 down to 19.25, its top corners rounded 1.7,
/// the notch up to (8, 15.75), 1.2 thick, round; 0.9723 of its ink apple.com's
/// (`reference/icons/`).
pub const SAVES_COMPACT: Icon = Icon {
    width: "16",
    height: "25",
    view_box: "0 0 16 25",
    paint: "",
    inner: r#"<path d="M4 19.25V7.45A1.7 1.7 0 0 1 5.7 5.75H10.3A1.7 1.7 0 0 1 12 7.45V19.25L8 15.75Z" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"/>"#,
};

/// A bag's profile link to the account (`cui-globalnav-baglink`, from 834px),
/// in place of apple.com's 11 by 16 SVG, a gear: a ring at (5.5, 7.25), radius
/// 4.05, 0.79 thick, a hub of radius 0.9, 0.65 thick, three spokes from 1.2
/// to 3.7 out at 0°, 120° and 240°, 0.69 thick, and 18 teeth from 4.4 to 4.82
/// out, 0.68 thick, round; 0.9387 of its ink apple.com's (`reference/icons/`).
pub const ACCOUNT: Icon = Icon {
    width: "11",
    height: "16",
    view_box: "0 0 11 16",
    paint: "",
    inner: r#"<circle cx="5.5" cy="7.25" r="4.05" fill="none" stroke="currentColor" stroke-width=".79"/><circle cx="5.5" cy="7.25" r=".9" fill="none" stroke="currentColor" stroke-width=".65"/><path d="M6.7 7.25H9.2M4.9 8.289 3.65 10.454M4.9 6.211 3.65 4.046" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".69"/><path d="M9.9 7.25H10.32M9.635 8.755 10.029 8.899M8.871 10.078 9.192 10.348M7.7 11.061 7.91 11.424M6.264 11.583 6.337 11.997M4.736 11.583 4.663 11.997M3.3 11.061 3.09 11.424M2.129 10.078 1.808 10.348M1.365 8.755 .971 8.899M1.1 7.25H.68M1.365 5.745 .971 5.601M2.129 4.422 1.808 4.152M3.3 3.439 3.09 3.076M4.736 2.917 4.663 2.503M6.264 2.917 6.337 2.503M7.7 3.439 7.91 3.076M8.871 4.422 9.192 4.152M9.635 5.745 10.029 5.601" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".68"/>"#,
};

/// The account's gear to 833px, in place of apple.com's 16 by 25 SVG: a ring
/// at (8, 12.5), radius 6.1, 1.28 thick, a hub of radius 1.42, 1.25 thick, the
/// spokes from 2 to 5.5 out, 1.21 thick, the teeth from 6.7 to 7.15 out, 1.05
/// thick, round; 0.9675 of its ink apple.com's (`reference/icons/`).
pub const ACCOUNT_COMPACT: Icon = Icon {
    width: "16",
    height: "25",
    view_box: "0 0 16 25",
    paint: "",
    inner: r#"<circle cx="8" cy="12.5" r="6.1" fill="none" stroke="currentColor" stroke-width="1.28"/><circle cx="8" cy="12.5" r="1.42" fill="none" stroke="currentColor" stroke-width="1.25"/><path d="M10 12.5H13.5M7 14.232 5.25 17.263M7 10.768 5.25 7.737" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.21"/><path d="M14.7 12.5H15.15M14.296 14.792 14.719 14.945M13.132 16.807 13.477 17.096M11.35 18.302 11.575 18.692M9.163 19.098 9.242 19.541M6.837 19.098 6.758 19.541M4.65 18.302 4.425 18.692M2.868 16.807 2.523 17.096M1.704 14.792 1.281 14.945M1.3 12.5H.85M1.704 10.208 1.281 10.055M2.868 8.193 2.523 7.904M4.65 6.698 4.425 6.308M6.837 5.902 6.758 5.459M9.163 5.902 9.242 5.459M11.35 6.698 11.575 6.308M13.132 8.193 13.477 7.904M14.296 10.208 14.719 10.055" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.05"/>"#,
};

/// A bag's profile link to sign in (`cui-globalnav-baglink`, from 834px), in
/// place of apple.com's 11 by 16 SVG, a person in a circle: a ring at (5.5,
/// 7.633), radius 4.287, 0.887 thick; a head, an ellipse at (5.5, 6.71), 1.605
/// by 1.69; the shoulders filled from (2.62, 10.163) up an arc of radius 3.898
/// to (8.38, 10.163) and back along the ring's inside; 0.972 of its ink
/// apple.com's (`reference/icons/`).
pub const SIGN_IN: Icon = Icon {
    width: "11",
    height: "16",
    view_box: "0 0 11 16",
    paint: "",
    inner: r#"<circle cx="5.5" cy="7.633" r="4.287" fill="none" stroke="currentColor" stroke-width=".887"/><ellipse cx="5.5" cy="6.71" rx="1.605" ry="1.69" fill="currentColor"/><path d="M2.62 10.163A3.898 3.898 0 0 1 8.38 10.163 3.844 3.844 0 0 1 2.62 10.163Z" fill="currentColor"/>"#,
};

/// The sign-in's person to 833px, in place of apple.com's 16 by 25 SVG: a
/// ring at (7.99, 12.5), radius 6.495, 1.21 thick; the head at (7.99, 11), 2.475
/// by 2.6; the shoulders from (3.529, 16.347) up an arc of radius 6.77 to
/// (12.452, 16.347) and back along the ring's inside; 0.986 of its ink
/// apple.com's (`reference/icons/`).
pub const SIGN_IN_COMPACT: Icon = Icon {
    width: "16",
    height: "25",
    view_box: "0 0 16 25",
    paint: "",
    inner: r#"<circle cx="7.99" cy="12.5" r="6.495" fill="none" stroke="currentColor" stroke-width="1.21"/><ellipse cx="7.99" cy="11" rx="2.475" ry="2.6" fill="currentColor"/><path d="M3.529 16.347A6.77 6.77 0 0 1 12.452 16.347 5.89 5.89 0 0 1 3.529 16.347Z" fill="currentColor"/>"#,
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

    // lucide-static's icons/x.svg: its <svg>'s paint and its paths, at 1em
    // where it has 24
    #[test]
    fn renders_a_lucide_icon_with_lucides_paint() {
        let html = rsx! { (crate::lucide::X) }.render();
        assert_eq!(
            html.as_inner(),
            concat!(
                r#"<svg aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">"#,
                r#"<path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>"#,
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
