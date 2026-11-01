// Package icon is icons as inline SVG, ciderui's own, each drawn from the
// numbers of the glyph or the drawing apple.com uses in its place
// (`reference/icons/`): its box, the ink's place in it, the strokes'
// thickness; not Apple's outlines. Each is sized by its own attributes: in em
// where it sits in text, so it takes the text's size ([More]), in pixels
// elsewhere (the global nav, a gallery's paddles, a modal's X), as
// apple.com's do; and drawn in currentColor, so it takes the text's color. A
// component's class (`cui-link-icon`) places it. Inline, so a page whose
// Content-Security-Policy refuses data: images draws it too. Every other icon
// is Lucide's (package lucide), an [Icon] too.
package icon

import (
	"html"
	"io"

	g "maragu.dev/gomponents"
)

// Icon is an icon: its box (in em or pixels), its view box, the attributes on
// its <svg> that what is drawn in it inherits, and what is drawn in it. It is
// a gomponents node.
type Icon struct {
	width   string
	height  string
	viewBox string
	paint   string
	inner   string
}

// Lucide is a Lucide icon (package lucide, which build-lucide.mjs writes):
// 1em square, its elements under the attributes on Lucide's <svg>.
func Lucide(viewBox, paint, inner string) Icon {
	// the paint on the <svg>, not on each element: a class on the <svg>
	// (stroke-[1.5]) wins over an attribute there, and its elements inherit
	// what wins
	return Icon{width: "1em", height: "1em", viewBox: viewBox, paint: paint, inner: inner}
}

// Render writes the icon's <svg>.
func (i Icon) Render(w io.Writer) error {
	return i.write(w, "", false)
}

// Class is the icon with a class on its <svg>: a component's part
// (`cui-link-icon`) or a utility.
func (i Icon) Class(class string) g.Node {
	return g.NodeFunc(func(w io.Writer) error { return i.write(w, class, true) })
}

func (i Icon) write(w io.Writer, class string, classed bool) error {
	svg := "<svg"
	if classed {
		svg += ` class="` + html.EscapeString(class) + `"`
	}
	svg += ` aria-hidden="true" width="` + i.width + `" height="` + i.height + `" viewBox="` + i.viewBox + `"`
	if i.paint != "" {
		svg += " " + i.paint
	}
	_, err := io.WriteString(w, svg+">"+i.inner+"</svg>")
	return err
}

// More is the "more" chevron after a link's text (`cui-link`, `data-variant="more"`),
// in place of apple.com's SF Pro Icons U+F301: its advance, 0.4419em, as the
// box's width; from 0.01em below the baseline to 0.53em above it; two
// strokes 0.0679em thick from its ink's corners to the middle, round at the
// ends and the join, which share 0.948 of their ink with apple.com's glyph
// (`reference/icons/`). The view box is in thousandths of an em.
var More = Icon{width: "0.442em", height: "0.54em", viewBox: "0 0 442 540", paint: ``, inner: `<path d="M107.7 47.7 332.3 272.5 107.7 497.3" fill="none" stroke="currentColor" stroke-width="67.9" stroke-linecap="round" stroke-linejoin="round"/>`}

// Search is the global nav's search from 834px (`cui-globalnav-image` in a
// `cui-globalnav-icon`), in place of apple.com's 15 by 44 SVG: a ring at
// (5.95, 19.95), radius 4.96, 1.08 thick, and its handle from the ring at
// 45° to (13.74, 27.74), 1.52 thick, round; 0.927 of its ink apple.com's.
var Search = Icon{width: "15", height: "44", viewBox: "0 0 15 44", paint: ``, inner: `<circle cx="5.95" cy="19.95" r="4.96" fill="none" stroke="currentColor" stroke-width="1.08"/><path d="M9.46 23.46 13.74 27.74" fill="none" stroke="currentColor" stroke-width="1.52" stroke-linecap="round"/>`}

// SearchCompact is the global nav's search to 833px (`cui-globalnav-image-compact`), in
// place of apple.com's 17 by 48 SVG: a ring at (6.94, 21.94), radius 5.86,
// 1.22 thick, and its handle to (15.6, 30.6), 1.8 thick, round; 0.972 of
// its ink apple.com's.
var SearchCompact = Icon{width: "17", height: "48", viewBox: "0 0 17 48", paint: ``, inner: `<circle cx="6.94" cy="21.94" r="5.86" fill="none" stroke="currentColor" stroke-width="1.22"/><path d="M11.08 26.08 15.6 30.6" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/>`}

// Bag is the global nav's bag from 834px (`cui-globalnav-image`), in place of
// apple.com's 14 by 44 SVG: a body 11.8 by 10.9 from (1.1, 16.55), corners
// 1.5, and its handle an arc of radius 2.45, both 1.1 thick; 0.948 of its
// ink apple.com's.
var Bag = Icon{width: "14", height: "44", viewBox: "0 0 14 44", paint: ``, inner: `<rect x="1.1" y="16.55" width="11.8" height="10.9" rx="1.5" fill="none" stroke="currentColor" stroke-width="1.1"/><path d="M4.55 16.55V16.05a2.45 2.45 0 0 1 4.9 0v.5" fill="none" stroke="currentColor" stroke-width="1.1"/>`}

// BagCompact is the global nav's bag to 833px (`cui-globalnav-image-compact`), in place
// of apple.com's 17 by 48 SVG: a body 13.9 by 12.95 from (1.55, 17.5),
// corners 2, and its handle an arc of radius 3.125, both 1.2 thick; 0.981
// of its ink apple.com's.
var BagCompact = Icon{width: "17", height: "48", viewBox: "0 0 17 48", paint: ``, inner: `<rect x="1.55" y="17.5" width="13.9" height="12.95" rx="2" fill="none" stroke="currentColor" stroke-width="1.2"/><path d="M5.375 17.5V17.175a3.125 3.125 0 0 1 6.25 0v.325" fill="none" stroke="currentColor" stroke-width="1.2"/>`}

// Searchfield is the global nav's search field's magnifier from 834px (`cui-globalnav-image`
// in a `cui-globalnav-searchfield`), in place of apple.com's 30 by 32 SVG: a
// ring at (13.475, 14.775), radius 6.225, 1.3 thick, and its handle from the
// ring at 45° to (22.57, 24.07), 2.2 thick, round; 0.974 of its ink
// apple.com's.
var Searchfield = Icon{width: "30", height: "32", viewBox: "0 0 30 32", paint: ``, inner: `<circle cx="13.475" cy="14.775" r="6.225" fill="none" stroke="currentColor" stroke-width="1.3"/><path d="M18.11 19.61 22.57 24.07" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round"/>`}

// SearchfieldCompact is the search field's magnifier to 833px (`cui-globalnav-image-compact`), in
// place of apple.com's 38 by 40 SVG: a ring at (17.225, 17.9), radius 7.25,
// 1.45 thick, and its handle to (27.74, 28.74), 2.62 thick, round; 0.970 of
// its ink apple.com's.
var SearchfieldCompact = Icon{width: "38", height: "40", viewBox: "0 0 38 40", paint: ``, inner: `<circle cx="17.225" cy="17.9" r="7.25" fill="none" stroke="currentColor" stroke-width="1.45"/><path d="M22.6 23.6 27.74 28.74" fill="none" stroke="currentColor" stroke-width="2.62" stroke-linecap="round"/>`}

// Arrow is a search quick link's arrow from 834px (`cui-globalnav-image` in a
// `cui-globalnav-searchlink`), in place of apple.com's 9 by 16 SVG: a shaft
// from (0.75, 7.75) to its head at (8.25, 7.75), the head's two strokes 3
// back at 45°, 1 thick, round; 0.983 of its ink apple.com's.
var Arrow = Icon{width: "9", height: "16", viewBox: "0 0 9 16", paint: ``, inner: `<path d="M.75 7.75H8.25M5.25 4.75 8.25 7.75 5.25 10.75" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round" stroke-linejoin="round"/>`}

// ArrowCompact is the quick link's arrow to 833px (`cui-globalnav-image-compact`), in place
// of apple.com's 13 by 25 SVG: a shaft from (1.1, 13) to its head at
// (11.93, 13), 1.3 thick, the head's strokes 4.45 back at 45°, 1.2 thick,
// round; 0.995 of its ink apple.com's.
var ArrowCompact = Icon{width: "13", height: "25", viewBox: "0 0 13 25", paint: ``, inner: `<path d="M1.1 13H11.93" fill="none" stroke="currentColor" stroke-width="1.3" stroke-linecap="round"/><path d="M7.48 8.55 11.93 13 7.48 17.45" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round"/>`}

// Suggestion is a suggested search's magnifier from 834px (`cui-globalnav-image` in a
// `cui-globalnav-searchlink`, where a quick link has its arrow), in place of
// apple.com's 9 by 16 SVG: a ring at (3.675, 6.9), radius 3.125, 0.7 thick,
// and its handle at 45° from just inside the ring to (8.2, 11.581), 1.1
// thick, round; 0.964 of its ink apple.com's.
var Suggestion = Icon{width: "9", height: "16", viewBox: "0 0 9 16", paint: ``, inner: `<circle cx="3.675" cy="6.9" r="3.125" fill="none" stroke="currentColor" stroke-width="0.7"/><path d="M6 9.381 8.2 11.581" fill="none" stroke="currentColor" stroke-width="1.1" stroke-linecap="round"/>`}

// SuggestionCompact is a suggested search's magnifier to 833px (`cui-globalnav-image-compact`),
// in place of apple.com's 13 by 25 SVG: a ring at (5.425, 11.725), radius
// 4.425, 0.9 thick, and its handle to (11.8, 18.35), 1.6 thick, round; 0.979
// of its ink apple.com's.
var SuggestionCompact = Icon{width: "13", height: "25", viewBox: "0 0 13 25", paint: ``, inner: `<circle cx="5.425" cy="11.725" r="4.425" fill="none" stroke="currentColor" stroke-width="0.9"/><path d="M8.692 15.242 11.8 18.35" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round"/>`}

// Clear is the search field's clear mark from 834px (`cui-globalnav-image` in a
// `cui-globalnav-searchreset`), in place of apple.com's 14 by 14 SVG: a disc
// at (7, 7), radius 6.9661, and cut out of it an X, two strokes 1.1 thick
// between round ends whose middles are 2.409 from the disc's middle along
// each diagonal, drawn as one outline (evenodd); 0.9995 of its ink
// apple.com's.
var Clear = Icon{width: "14", height: "14", viewBox: "0 0 14 14", paint: ``, inner: `<path d="M7 .0339A6.9661 6.9661 0 1 0 7 13.9661 6.9661 6.9661 0 1 0 7 .0339ZM7.7778 7 9.798 9.0201A.55.55 0 0 1 9.0201 9.798L7 7.7778 4.9799 9.798A.55.55 0 0 1 4.2021 9.0201L6.2222 7 4.2021 4.9799A.55.55 0 0 1 4.9799 4.2021L7 6.2222 9.0201 4.2021A.55.55 0 0 1 9.798 4.9799Z" fill="currentColor" fill-rule="evenodd"/>`}

// ClearCompact is the clear mark to 833px (`cui-globalnav-image-compact`), in place of
// apple.com's 16 by 16 SVG: a disc at (8, 8), radius 8, its X's strokes
// 1.34 thick, their round ends' middles 2.81 from the disc's middle; 0.9928
// of its ink apple.com's.
var ClearCompact = Icon{width: "16", height: "16", viewBox: "0 0 16 16", paint: ``, inner: `<path d="M8 0A8 8 0 1 0 8 16 8 8 0 1 0 8 0ZM8.9475 8 11.2838 10.3362A.67.67 0 0 1 10.3362 11.2838L8 8.9475 5.6638 11.2838A.67.67 0 0 1 4.7162 10.3362L7.0525 8 4.7162 5.6638A.67.67 0 0 1 5.6638 4.7162L8 7.0525 10.3362 4.7162A.67.67 0 0 1 11.2838 5.6638Z" fill="currentColor" fill-rule="evenodd"/>`}

// Menu is the global nav's menu button's lines (`cui-globalnav-menubutton`), as
// apple.com's: two lines across 18 by 18, from 2 to 16 at 5 and 12 down,
// 1.2 thick, round; in the opened menu's close button
// (`cui-globalnav-closebutton`) `globalnav.css` turns them into its X (from
// (3.5, 3.5) to (15, 15) and from (3.5, 15) to (15, 3.5)).
var Menu = Icon{width: "18", height: "18", viewBox: "0 0 18 18", paint: ``, inner: `<path d="M2 5H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/><path d="M2 12H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/>`}

// Flyout is the chevron of a flyout's button from 834px (`cui-globalnav-image` in a
// `cui-globalnav-flyoutbutton`), shown when it has the focus from the
// keyboard, in place of apple.com's 10 by 10 SVG: a line through (2.33,
// 3.66), (5, 6.34) and (7.67, 3.66), 0.88 thick, round; 0.99 of its ink
// apple.com's (`reference/icons/`).
var Flyout = Icon{width: "10", height: "10", viewBox: "0 0 10 10", paint: ``, inner: `<path d="M2.33 3.66 5 6.34 7.67 3.66" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".88"/>`}

// FlyoutCompact is the same button's chevron to 833px (`cui-globalnav-image-compact`), at
// the end of the menu's row, shown hovered, in place of apple.com's 9 by 48
// SVG: a line through (1.3, 17.21), (8.3, 24.07) and (1.3, 30.79), 1.2
// thick, round; 0.9968 of its ink apple.com's, and 0.9962 of the chevron a
// hovered row of apple.com's shows (`reference/icons/`).
var FlyoutCompact = Icon{width: "9", height: "48", viewBox: "0 0 9 48", paint: ``, inner: `<path d="M1.3 17.21 8.3 24.07 1.3 30.79" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"/>`}

// FlyoutBack is a flyout's back button to 833px (`cui-globalnav-backbutton`), in place of
// apple.com's 9 by 48 SVG: [FlyoutCompact] turned, through (7.7, 17.21),
// (0.7, 24.07) and (7.7, 30.79); 0.9962 of its ink apple.com's
// (`reference/icons/`).
var FlyoutBack = Icon{width: "9", height: "48", viewBox: "0 0 9 48", paint: ``, inner: `<path d="M7.7 17.21 .7 24.07 7.7 30.79" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"/>`}

// Chevron is the local nav's chevron (`cui-localnav-menubutton`), as apple.com's: a
// line through (15.265, 0.835), (8, 8.167) and (0.735, 0.835) in 16 by 9,
// 1.2 thick, round; `localnav.css` turns it up while the menu is open, which
// is the line apple.com's animation ends on (`reference/icons/`).
var Chevron = Icon{width: "16", height: "9", viewBox: "0 0 16 9", paint: ``, inner: `<path d="M15.265.835 8 8.167.735.835" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round"/>`}

// PaddlePrevious is a gallery's previous paddle's chevron (`cui-gallery-paddle`), in place of
// apple.com's 36 by 36 SVG: a line through (20, 12.5), (14.5, 18) and
// (20, 23.5), 3 thick, round; 0.9995 of its ink apple.com's
// (`reference/icons/`). Its size the paddle's (`gallery.css`).
var PaddlePrevious = Icon{width: "36", height: "36", viewBox: "0 0 36 36", paint: ``, inner: `<path d="M20 12.5 14.5 18 20 23.5" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>`}

// PaddleNext is a gallery's next paddle's chevron: [PaddlePrevious] turned, through
// (16, 12.5), (21.5, 18) and (16, 23.5); 0.9993 of its ink apple.com's.
var PaddleNext = Icon{width: "36", height: "36", viewBox: "0 0 36 36", paint: ``, inner: `<path d="M16 12.5 21.5 18 16 23.5" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>`}

// Check is a checkbox's check (`cui-checkbox-indicator`), in place of apple.com's SF
// Pro Icons glyph: a line through (3.3, 7.65), (6.25, 11.35) and
// (11.15, 3.7) in the box's 14 by 14, 1.35 thick, round; 0.96 of its ink
// apple.com's (`reference/icons/`).
var Check = Icon{width: "14", height: "14", viewBox: "0 0 14 14", paint: ``, inner: `<path d="M3.3 7.65 6.25 11.35 11.15 3.7" fill="none" stroke="currentColor" stroke-width="1.35" stroke-linecap="round" stroke-linejoin="round"/>`}

// Disclosure is an accordion's chevron (`cui-accordion-question`), as apple.com's: a line
// through (15, 1.13), (8.5, 7.72) and (2, 1.13) in 17 by 8.85, 2.25 thick,
// round (`reference/icons/`); its box square, the line in its middle, sized
// by `accordion.css`, which turns it up while its answer is open.
var Disclosure = Icon{width: "17", height: "17", viewBox: "0 0 17 8.85", paint: ``, inner: `<path d="M15 1.13 8.5 7.72 2 1.13" fill="none" stroke="currentColor" stroke-width="2.25" stroke-linecap="round" stroke-linejoin="round"/>`}

// FooterChevron is the footer's chevron (`cui-footer-title`, to 833px), as apple.com's: a
// line through (10.075, 0.675), (5.5, 5.323) and (0.925, 0.675) in 11 by
// 6, 1 thick, round (`reference/icons/`); `footer.css` turns it up while
// its section is open, the line apple.com's animation ends on.
var FooterChevron = Icon{width: "11", height: "6", viewBox: "0 0 11 6", paint: ``, inner: `<path d="M10.075.675 5.5 5.323.925.675" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round" stroke-linejoin="round"/>`}

// Alert is a field's error mark (`cui-field-message`), in place of apple.com's SF
// Pro Icons glyph: a ring at (0.585em, 0.4975em down), radius 0.456em,
// 0.0825em thick, and in it a stem from 0.28 to 0.55em down, 0.07em thick,
// round, over a dot of radius 0.05em at 0.6925em down, in the glyph's
// advance (1.171em) by its ink's height (0.995em); 0.9695 of its ink
// apple.com's (`reference/icons/`). In em: it takes the message's size.
var Alert = Icon{width: "1.171em", height: "0.995em", viewBox: "0 0 1171 995", paint: ``, inner: `<circle cx="585" cy="497.5" r="456.25" fill="none" stroke="currentColor" stroke-width="82.5"/><path d="M584 280V550" fill="none" stroke="currentColor" stroke-width="70" stroke-linecap="round"/><circle cx="585" cy="692.5" r="50" fill="currentColor"/>`}

// Dropdown is a dropdown's chevron (`cui-field-chevron`), in place of the store's
// background SVG: a line through (6.075, 9.275), (10.5, 13.7) and
// (14.925, 9.275) in 21 by 21, 1.45 thick, round, drawn 25 by 25 as
// apple.com's; 0.9689 of its ink apple.com's (`reference/icons/`).
var Dropdown = Icon{width: "25", height: "25", viewBox: "0 0 21 21", paint: ``, inner: `<path d="M6.075 9.275 10.5 13.7 14.925 9.275" fill="none" stroke="currentColor" stroke-width="1.45" stroke-linecap="round" stroke-linejoin="round"/>`}

// Close is a modal's X (`cui-modal-close`), in place of the store overlay's SVG: two
// lines from (4.88, 4.88) to (15.12, 15.12) and from (15.12, 4.88) to
// (4.88, 15.12) in 20 by 20, 3 thick, round; 0.9927 of its ink apple.com's
// (`reference/icons/`).
var Close = Icon{width: "20", height: "20", viewBox: "0 0 20 20", paint: ``, inner: `<path d="M4.88 4.88 15.12 15.12M15.12 4.88 4.88 15.12" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round"/>`}

// CloseSmall is the marketing pages' modal's X (`cui-modal-close` in a
// `data-variant="page"` modal), in place of its SVG: [Close]'s strokes,
// each arm 0.62 shorter at its ends, from (5.5, 5.5) to (14.5, 14.5) and
// from (14.5, 5.5) to (5.5, 14.5); 0.9994 of its ink apple.com's
// (`reference/icons/`).
var CloseSmall = Icon{width: "20", height: "20", viewBox: "0 0 20 20", paint: ``, inner: `<path d="M5.5 5.5 14.5 14.5M14.5 5.5 5.5 14.5" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round"/>`}

// Orders is a bag's profile link to its orders (`cui-globalnav-baglink`, from 834px), in
// place of apple.com's 11 by 16 SVG, a box: a hexagon through (5.511, 2.548),
// (9.738, 4.933), (9.738, 9.684), (5.511, 12.068), (1.285, 9.684) and (1.285,
// 4.933), its corners rounded 1.0625, the top's lower edges meeting at (5.511,
// 7.375) over a line down to its foot, and a band across its top from (3.849,
// 3.988) to (7.254, 5.878), 0.875 thick; 0.9915 of its ink apple.com's
// (`reference/icons/`).
var Orders = Icon{width: "11", height: "16", viewBox: "0 0 11 16", paint: ``, inner: `<path d="M6.033 2.843 9.198 4.628A1.0625 1.0625 0 0 1 9.738 5.553V9.063A1.0625 1.0625 0 0 1 9.198 9.989L6.033 11.774A1.0625 1.0625 0 0 1 4.989 11.774L1.825 9.989A1.0625 1.0625 0 0 1 1.285 9.063V5.553A1.0625 1.0625 0 0 1 1.825 4.628L4.989 2.843A1.0625 1.0625 0 0 1 6.033 2.843ZM1.43 5.018 5.511 7.375 9.593 5.018M5.511 7.375V11.91" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".875"/><path d="M3.849 3.988 7.254 5.878" fill="none" stroke="currentColor" stroke-width=".875"/>`}

// OrdersCompact is the orders' box to 833px, in place of apple.com's 16 by 25 SVG: the same
// hexagon through (8, 5.74), (14, 9.125), (14, 15.875), (8, 19.259), (2, 15.875)
// and (2, 9.125), its corners rounded 1.7, the top's edges meeting at (8,
// 12.594), the band from (5.561, 7.806) to (10.456, 10.522), 1.2 thick; 0.9851
// of its ink apple.com's (`reference/icons/`).
var OrdersCompact = Icon{width: "16", height: "25", viewBox: "0 0 16 25", paint: ``, inner: `<path d="M8.835 6.211 13.135 8.637A1.7 1.7 0 0 1 14 10.118V14.882A1.7 1.7 0 0 1 13.135 16.363L8.835 18.788A1.7 1.7 0 0 1 7.165 18.788L2.865 16.363A1.7 1.7 0 0 1 2 14.882V10.118A1.7 1.7 0 0 1 2.865 8.637L7.165 6.211A1.7 1.7 0 0 1 8.835 6.211ZM2.232 9.26 8 12.594 13.768 9.26M8 12.594V19.007" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"/><path d="M5.561 7.806 10.456 10.522" fill="none" stroke="currentColor" stroke-width="1.2"/>`}

// Saves is a bag's profile link to its saved items (`cui-globalnav-baglink`, from
// 834px), in place of apple.com's 11 by 16 SVG, a bookmark: its sides at 2.719
// and 8.281 from 2.789 down to 11.97, its top corners rounded 1.184, its foot
// a notch up to (5.5, 9.3), 0.82 thick, round; 0.9518 of its ink apple.com's
// (`reference/icons/`).
var Saves = Icon{width: "11", height: "16", viewBox: "0 0 11 16", paint: ``, inner: `<path d="M2.719 11.97V3.972A1.184 1.184 0 0 1 3.903 2.789H7.097A1.184 1.184 0 0 1 8.281 3.972V11.97L5.5 9.3Z" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".82"/>`}

// SavesCompact is the saved items' bookmark to 833px, in place of apple.com's 16 by 25 SVG:
// its sides at 4 and 12 from 5.75 down to 19.25, its top corners rounded 1.7,
// the notch up to (8, 15.75), 1.2 thick, round; 0.9723 of its ink apple.com's
// (`reference/icons/`).
var SavesCompact = Icon{width: "16", height: "25", viewBox: "0 0 16 25", paint: ``, inner: `<path d="M4 19.25V7.45A1.7 1.7 0 0 1 5.7 5.75H10.3A1.7 1.7 0 0 1 12 7.45V19.25L8 15.75Z" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"/>`}

// Account is a bag's profile link to the account (`cui-globalnav-baglink`, from 834px),
// in place of apple.com's 11 by 16 SVG, a gear: a ring at (5.5, 7.25), radius
// 4.05, 0.79 thick, a hub of radius 0.9, 0.65 thick, three spokes from 1.2
// to 3.7 out at 0°, 120° and 240°, 0.69 thick, and 18 teeth from 4.4 to 4.82
// out, 0.68 thick, round; 0.9387 of its ink apple.com's (`reference/icons/`).
var Account = Icon{width: "11", height: "16", viewBox: "0 0 11 16", paint: ``, inner: `<circle cx="5.5" cy="7.25" r="4.05" fill="none" stroke="currentColor" stroke-width=".79"/><circle cx="5.5" cy="7.25" r=".9" fill="none" stroke="currentColor" stroke-width=".65"/><path d="M6.7 7.25H9.2M4.9 8.289 3.65 10.454M4.9 6.211 3.65 4.046" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".69"/><path d="M9.9 7.25H10.32M9.635 8.755 10.029 8.899M8.871 10.078 9.192 10.348M7.7 11.061 7.91 11.424M6.264 11.583 6.337 11.997M4.736 11.583 4.663 11.997M3.3 11.061 3.09 11.424M2.129 10.078 1.808 10.348M1.365 8.755 .971 8.899M1.1 7.25H.68M1.365 5.745 .971 5.601M2.129 4.422 1.808 4.152M3.3 3.439 3.09 3.076M4.736 2.917 4.663 2.503M6.264 2.917 6.337 2.503M7.7 3.439 7.91 3.076M8.871 4.422 9.192 4.152M9.635 5.745 10.029 5.601" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".68"/>`}

// AccountCompact is the account's gear to 833px, in place of apple.com's 16 by 25 SVG: a ring
// at (8, 12.5), radius 6.1, 1.28 thick, a hub of radius 1.42, 1.25 thick, the
// spokes from 2 to 5.5 out, 1.21 thick, the teeth from 6.7 to 7.15 out, 1.05
// thick, round; 0.9675 of its ink apple.com's (`reference/icons/`).
var AccountCompact = Icon{width: "16", height: "25", viewBox: "0 0 16 25", paint: ``, inner: `<circle cx="8" cy="12.5" r="6.1" fill="none" stroke="currentColor" stroke-width="1.28"/><circle cx="8" cy="12.5" r="1.42" fill="none" stroke="currentColor" stroke-width="1.25"/><path d="M10 12.5H13.5M7 14.232 5.25 17.263M7 10.768 5.25 7.737" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.21"/><path d="M14.7 12.5H15.15M14.296 14.792 14.719 14.945M13.132 16.807 13.477 17.096M11.35 18.302 11.575 18.692M9.163 19.098 9.242 19.541M6.837 19.098 6.758 19.541M4.65 18.302 4.425 18.692M2.868 16.807 2.523 17.096M1.704 14.792 1.281 14.945M1.3 12.5H.85M1.704 10.208 1.281 10.055M2.868 8.193 2.523 7.904M4.65 6.698 4.425 6.308M6.837 5.902 6.758 5.459M9.163 5.902 9.242 5.459M11.35 6.698 11.575 6.308M13.132 8.193 13.477 7.904M14.296 10.208 14.719 10.055" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.05"/>`}

// SignIn is a bag's profile link to sign in (`cui-globalnav-baglink`, from 834px), in
// place of apple.com's 11 by 16 SVG, a person in a circle: a ring at (5.5,
// 7.633), radius 4.287, 0.887 thick; a head, an ellipse at (5.5, 6.71), 1.605
// by 1.69; the shoulders filled from (2.62, 10.163) up an arc of radius 3.898
// to (8.38, 10.163) and back along the ring's inside; 0.972 of its ink
// apple.com's (`reference/icons/`).
var SignIn = Icon{width: "11", height: "16", viewBox: "0 0 11 16", paint: ``, inner: `<circle cx="5.5" cy="7.633" r="4.287" fill="none" stroke="currentColor" stroke-width=".887"/><ellipse cx="5.5" cy="6.71" rx="1.605" ry="1.69" fill="currentColor"/><path d="M2.62 10.163A3.898 3.898 0 0 1 8.38 10.163 3.844 3.844 0 0 1 2.62 10.163Z" fill="currentColor"/>`}

// SignInCompact is the sign-in's person to 833px, in place of apple.com's 16 by 25 SVG: a
// ring at (7.99, 12.5), radius 6.495, 1.21 thick; the head at (7.99, 11), 2.475
// by 2.6; the shoulders from (3.529, 16.347) up an arc of radius 6.77 to
// (12.452, 16.347) and back along the ring's inside; 0.986 of its ink
// apple.com's (`reference/icons/`).
var SignInCompact = Icon{width: "16", height: "25", viewBox: "0 0 16 25", paint: ``, inner: `<circle cx="7.99" cy="12.5" r="6.495" fill="none" stroke="currentColor" stroke-width="1.21"/><ellipse cx="7.99" cy="11" rx="2.475" ry="2.6" fill="currentColor"/><path d="M3.529 16.347A6.77 6.77 0 0 1 12.452 16.347 5.89 5.89 0 0 1 3.529 16.347Z" fill="currentColor"/>`}

// DocsSidebar is apple's developer documentation's bar's button that opens the navigator
// (`cui-docs-bar-toggle`), in place of DocC's 14 by 14 SVG drawn at 19px: a
// rounded rectangle 0.975 thick from (0.625, 2) to (13.375, 12), its
// corners 1.27 at the stroke's middle, a bar 0.95 wide down it at 5.3, and
// three lines in the left pane at 4.635, 6.056 and 7.473 down, from 2.325
// to 4.125 across, 0.656 thick, round; 0.987 of its ink DocC's
// (`reference/icons/`).
var DocsSidebar = Icon{width: "19", height: "19", viewBox: "0 0 14 14", paint: ``, inner: `<rect x="1.1125" y="2.4875" width="11.775" height="9.025" rx="1.27" fill="none" stroke="currentColor" stroke-width="0.975"/><path d="M5.3 2.4875V11.5125" fill="none" stroke="currentColor" stroke-width="0.95"/><path d="M2.653 4.635H3.797M2.653 6.056H3.797M2.653 7.473H3.797" fill="none" stroke="currentColor" stroke-width="0.656" stroke-linecap="round"/>`}

// DocsChevron is apple's developer documentation's navigator's toggle, the chevron of a
// row that opens (`cui-docs-navigator-toggle`), in place of DocC's 14 by 14
// SVG drawn at 10px: two strokes 1.63 thick from (4.81, 1.81) to (10.17, 7)
// to (4.81, 12.32), round at the ends and the join; 0.975 of its ink
// DocC's.
var DocsChevron = Icon{width: "10", height: "10", viewBox: "0 0 14 14", paint: ``, inner: `<path d="M4.81 1.81 10.17 7 4.81 12.32" fill="none" stroke="currentColor" stroke-width="1.63" stroke-linecap="round" stroke-linejoin="round"/>`}

// DocsArticle is apple's developer documentation's glyph for an article, before a row of
// the navigator (17px) and a link to it (15px), in place of DocC's 14 by 14
// SVG: a page 1.05 thick from (1.6, 0) to (12.6, 14), its corners 1.35, its
// top right corner folded down from (7.515, 0.525) to (12.075, 5.085), the
// fold's edge 0.975 thick at 7.09 across and 5.59 down, two lines of text at
// 8.28 and 10.57 down from 4.13 to 9.9, 0.8 thick, round; 0.982 of its ink
// DocC's.
var DocsArticle = Icon{width: "17", height: "17", viewBox: "0 0 14 14", paint: ``, inner: `<path d="M7.515 0.525H3.475A1.35 1.35 0 0 0 2.125 1.875V12.125A1.35 1.35 0 0 0 3.475 13.475H10.725A1.35 1.35 0 0 0 12.075 12.125V5.085Z" fill="none" stroke="currentColor" stroke-width="1.05" stroke-linejoin="round"/><path d="M7.09 0.6V4.9A0.69 0.69 0 0 0 7.78 5.59H12" fill="none" stroke="currentColor" stroke-width="0.975"/><path d="M4.53 8.28H9.5M4.53 10.57H9.5" fill="none" stroke="currentColor" stroke-width="0.8" stroke-linecap="round"/>`}

// DocsCollection is apple's developer documentation's glyph for a collection, a page of
// pages, before a row of the navigator (17px) and a link to it (15px), in
// place of DocC's 14 by 14 SVG: three dots 1.72 across at 0.875 across and
// 2.81, 7.01 and 11.21 down, each before a line from 3.375 to 14 across,
// 1.125 thick, round; 0.992 of its ink DocC's.
var DocsCollection = Icon{width: "17", height: "17", viewBox: "0 0 14 14", paint: ``, inner: `<circle cx="0.875" cy="2.81" r="0.86" fill="currentColor"/><circle cx="0.875" cy="7.01" r="0.86" fill="currentColor"/><circle cx="0.875" cy="11.21" r="0.86" fill="currentColor"/><path d="M3.94 2.81H13.44M3.94 7.01H13.44M3.94 11.21H13.44" fill="none" stroke="currentColor" stroke-width="1.125" stroke-linecap="round"/>`}

// DocsClose is apple's developer documentation's button that closes the navigator's
// sheet (`cui-docs-page-close`), in place of DocC's 14 by 14 SVG drawn at
// 19px: two strokes 1.58 thick, from (1.79, 1.79) to (11.235, 11.235) and
// from (11.235, 1.79) to (1.79, 11.235), round; 0.9985 of its ink DocC's.
var DocsClose = Icon{width: "19", height: "19", viewBox: "0 0 14 14", paint: ``, inner: `<path d="M1.79 1.79 11.235 11.235M11.235 1.79 1.79 11.235" fill="none" stroke="currentColor" stroke-width="1.58" stroke-linecap="round"/>`}

// DocsCopy is apple's developer documentation's copy button's glyph
// (`cui-docs-copy-button`), in place of DocC's 24 by 24 SVG drawn at
// 25.5px: a sheet behind, its top and left edges 2 thick from (16, 2) round
// a corner of 1 to (3, 17), and a sheet in front, a rectangle 2 thick from
// (7, 6) to (20, 22), corners 1; 0.9997 of its ink DocC's.
var DocsCopy = Icon{width: "25.5", height: "25.5", viewBox: "0 0 24 24", paint: ``, inner: `<path d="M16 2H4A1 1 0 0 0 3 3V17" fill="none" stroke="currentColor" stroke-width="2"/><rect x="7" y="6" width="13" height="16" rx="1" fill="none" stroke="currentColor" stroke-width="2"/>`}

// DocsCopied is apple's developer documentation's copy button's glyph once it has
// copied, in place of DocC's 24 by 24 check drawn at 25.5px: two strokes
// 1.98 thick, square, from (4.1, 12.7) to (9, 17.6) to (20.3, 6.3); 0.9992
// of its ink DocC's.
var DocsCopied = Icon{width: "25.5", height: "25.5", viewBox: "0 0 24 24", paint: ``, inner: `<path d="M4.1 12.7 9 17.6 20.3 6.3" fill="none" stroke="currentColor" stroke-width="1.98"/>`}

// DocsFilter is apple's developer documentation's navigator's filter's glyph
// (`cui-docs-navigator-filter-icon`), in place of DocC's 21 by 21 SVG
// drawn at 21px: three lines 1 thick, round, at 6.3, 10.5 and 14.7 down,
// from 2.75 to 18.3, 4.6 to 16.4 and 6.5 to 14.7 across; 0.9983 of its ink
// DocC's.
var DocsFilter = Icon{width: "21", height: "21", viewBox: "0 0 21 21", paint: ``, inner: `<path d="M2.75 6.3H18.3M4.6 10.5H16.4M6.5 14.7H14.7" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round"/>`}
