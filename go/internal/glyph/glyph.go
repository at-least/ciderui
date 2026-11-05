// Package glyph is the components' own icons, each with the box of its one place
// in a part, so a site cannot pick one where it means a Lucide icon. Inline, not
// a data: image: a page whose Content-Security-Policy refuses data: images draws
// it too.
package glyph

import (
	"html"
	"io"

	g "maragu.dev/gomponents"
)

// Icon is an inline <svg>, a gomponents node.
type Icon struct {
	width   string
	height  string
	viewBox string
	paint   string
	inner   string
}

// Lucide is a Lucide icon; paint is the attributes on Lucide's own <svg>.
func Lucide(viewBox, paint, inner string) Icon {
	// the paint on the <svg>, not on each element: a class on the <svg>
	// (stroke-[1.5]) wins over a presentation attribute there, and the
	// elements inherit what wins
	return Icon{width: "1em", height: "1em", viewBox: viewBox, paint: paint, inner: inner}
}

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

var More = Icon{width: "0.442em", height: "0.54em", viewBox: "0 0 442 540", paint: ``, inner: `<path d="M107.7 47.7 332.3 272.5 107.7 497.3" fill="none" stroke="currentColor" stroke-width="67.9" stroke-linecap="round" stroke-linejoin="round"/>`}

var Search = Icon{width: "15", height: "44", viewBox: "0 0 15 44", paint: ``, inner: `<circle cx="5.95" cy="19.95" r="4.96" fill="none" stroke="currentColor" stroke-width="1.08"/><path d="M9.46 23.46 13.74 27.74" fill="none" stroke="currentColor" stroke-width="1.52" stroke-linecap="round"/>`}

var SearchCompact = Icon{width: "17", height: "48", viewBox: "0 0 17 48", paint: ``, inner: `<circle cx="6.94" cy="21.94" r="5.86" fill="none" stroke="currentColor" stroke-width="1.22"/><path d="M11.08 26.08 15.6 30.6" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/>`}

var Bag = Icon{width: "14", height: "44", viewBox: "0 0 14 44", paint: ``, inner: `<rect x="1.1" y="16.55" width="11.8" height="10.9" rx="1.5" fill="none" stroke="currentColor" stroke-width="1.1"/><path d="M4.55 16.55V16.05a2.45 2.45 0 0 1 4.9 0v.5" fill="none" stroke="currentColor" stroke-width="1.1"/>`}

var BagCompact = Icon{width: "17", height: "48", viewBox: "0 0 17 48", paint: ``, inner: `<rect x="1.55" y="17.5" width="13.9" height="12.95" rx="2" fill="none" stroke="currentColor" stroke-width="1.2"/><path d="M5.375 17.5V17.175a3.125 3.125 0 0 1 6.25 0v.325" fill="none" stroke="currentColor" stroke-width="1.2"/>`}

var Searchfield = Icon{width: "30", height: "32", viewBox: "0 0 30 32", paint: ``, inner: `<circle cx="13.475" cy="14.775" r="6.225" fill="none" stroke="currentColor" stroke-width="1.3"/><path d="M18.11 19.61 22.57 24.07" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round"/>`}

var SearchfieldCompact = Icon{width: "38", height: "40", viewBox: "0 0 38 40", paint: ``, inner: `<circle cx="17.225" cy="17.9" r="7.25" fill="none" stroke="currentColor" stroke-width="1.45"/><path d="M22.6 23.6 27.74 28.74" fill="none" stroke="currentColor" stroke-width="2.62" stroke-linecap="round"/>`}

var Arrow = Icon{width: "9", height: "16", viewBox: "0 0 9 16", paint: ``, inner: `<path d="M.75 7.75H8.25M5.25 4.75 8.25 7.75 5.25 10.75" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round" stroke-linejoin="round"/>`}

var ArrowCompact = Icon{width: "13", height: "25", viewBox: "0 0 13 25", paint: ``, inner: `<path d="M1.1 13H11.93" fill="none" stroke="currentColor" stroke-width="1.3" stroke-linecap="round"/><path d="M7.48 8.55 11.93 13 7.48 17.45" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round"/>`}

var Clear = Icon{width: "14", height: "14", viewBox: "0 0 14 14", paint: ``, inner: `<path d="M7 .0339A6.9661 6.9661 0 1 0 7 13.9661 6.9661 6.9661 0 1 0 7 .0339ZM7.7778 7 9.798 9.0201A.55.55 0 0 1 9.0201 9.798L7 7.7778 4.9799 9.798A.55.55 0 0 1 4.2021 9.0201L6.2222 7 4.2021 4.9799A.55.55 0 0 1 4.9799 4.2021L7 6.2222 9.0201 4.2021A.55.55 0 0 1 9.798 4.9799Z" fill="currentColor" fill-rule="evenodd"/>`}

var ClearCompact = Icon{width: "16", height: "16", viewBox: "0 0 16 16", paint: ``, inner: `<path d="M8 0A8 8 0 1 0 8 16 8 8 0 1 0 8 0ZM8.9475 8 11.2838 10.3362A.67.67 0 0 1 10.3362 11.2838L8 8.9475 5.6638 11.2838A.67.67 0 0 1 4.7162 10.3362L7.0525 8 4.7162 5.6638A.67.67 0 0 1 5.6638 4.7162L8 7.0525 10.3362 4.7162A.67.67 0 0 1 11.2838 5.6638Z" fill="currentColor" fill-rule="evenodd"/>`}

var Menu = Icon{width: "18", height: "18", viewBox: "0 0 18 18", paint: ``, inner: `<path d="M2 5H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/><path d="M2 12H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/>`}

var Flyout = Icon{width: "10", height: "10", viewBox: "0 0 10 10", paint: ``, inner: `<path d="M2.33 3.66 5 6.34 7.67 3.66" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".88"/>`}

var FlyoutCompact = Icon{width: "9", height: "48", viewBox: "0 0 9 48", paint: ``, inner: `<path d="M1.3 17.21 8.3 24.07 1.3 30.79" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"/>`}

var FlyoutBack = Icon{width: "9", height: "48", viewBox: "0 0 9 48", paint: ``, inner: `<path d="M7.7 17.21 .7 24.07 7.7 30.79" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"/>`}

var Chevron = Icon{width: "16", height: "9", viewBox: "0 0 16 9", paint: ``, inner: `<path d="M15.265.835 8 8.167.735.835" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round"/>`}

var PaddlePrevious = Icon{width: "36", height: "36", viewBox: "0 0 36 36", paint: ``, inner: `<path d="M20 12.5 14.5 18 20 23.5" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>`}

var PaddleNext = Icon{width: "36", height: "36", viewBox: "0 0 36 36", paint: ``, inner: `<path d="M16 12.5 21.5 18 16 23.5" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>`}

var Check = Icon{width: "14", height: "14", viewBox: "0 0 14 14", paint: ``, inner: `<path d="M3.3 7.65 6.25 11.35 11.15 3.7" fill="none" stroke="currentColor" stroke-width="1.35" stroke-linecap="round" stroke-linejoin="round"/>`}

var Disclosure = Icon{width: "17", height: "17", viewBox: "0 0 17 8.85", paint: ``, inner: `<path d="M15 1.13 8.5 7.72 2 1.13" fill="none" stroke="currentColor" stroke-width="2.25" stroke-linecap="round" stroke-linejoin="round"/>`}

var FooterChevron = Icon{width: "11", height: "6", viewBox: "0 0 11 6", paint: ``, inner: `<path d="M10.075.675 5.5 5.323.925.675" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round" stroke-linejoin="round"/>`}

var Alert = Icon{width: "1.171em", height: "0.995em", viewBox: "0 0 1171 995", paint: ``, inner: `<circle cx="585" cy="497.5" r="456.25" fill="none" stroke="currentColor" stroke-width="82.5"/><path d="M584 280V550" fill="none" stroke="currentColor" stroke-width="70" stroke-linecap="round"/><circle cx="585" cy="692.5" r="50" fill="currentColor"/>`}

var Dropdown = Icon{width: "25", height: "25", viewBox: "0 0 21 21", paint: ``, inner: `<path d="M6.075 9.275 10.5 13.7 14.925 9.275" fill="none" stroke="currentColor" stroke-width="1.45" stroke-linecap="round" stroke-linejoin="round"/>`}

var Close = Icon{width: "20", height: "20", viewBox: "0 0 20 20", paint: ``, inner: `<path d="M4.88 4.88 15.12 15.12M15.12 4.88 4.88 15.12" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round"/>`}

var CloseSmall = Icon{width: "20", height: "20", viewBox: "0 0 20 20", paint: ``, inner: `<path d="M5.5 5.5 14.5 14.5M14.5 5.5 5.5 14.5" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round"/>`}

var Orders = Icon{width: "11", height: "16", viewBox: "0 0 11 16", paint: ``, inner: `<path d="M6.033 2.843 9.198 4.628A1.0625 1.0625 0 0 1 9.738 5.553V9.063A1.0625 1.0625 0 0 1 9.198 9.989L6.033 11.774A1.0625 1.0625 0 0 1 4.989 11.774L1.825 9.989A1.0625 1.0625 0 0 1 1.285 9.063V5.553A1.0625 1.0625 0 0 1 1.825 4.628L4.989 2.843A1.0625 1.0625 0 0 1 6.033 2.843ZM1.43 5.018 5.511 7.375 9.593 5.018M5.511 7.375V11.91" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".875"/><path d="M3.849 3.988 7.254 5.878" fill="none" stroke="currentColor" stroke-width=".875"/>`}

var OrdersCompact = Icon{width: "16", height: "25", viewBox: "0 0 16 25", paint: ``, inner: `<path d="M8.835 6.211 13.135 8.637A1.7 1.7 0 0 1 14 10.118V14.882A1.7 1.7 0 0 1 13.135 16.363L8.835 18.788A1.7 1.7 0 0 1 7.165 18.788L2.865 16.363A1.7 1.7 0 0 1 2 14.882V10.118A1.7 1.7 0 0 1 2.865 8.637L7.165 6.211A1.7 1.7 0 0 1 8.835 6.211ZM2.232 9.26 8 12.594 13.768 9.26M8 12.594V19.007" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"/><path d="M5.561 7.806 10.456 10.522" fill="none" stroke="currentColor" stroke-width="1.2"/>`}

var Saves = Icon{width: "11", height: "16", viewBox: "0 0 11 16", paint: ``, inner: `<path d="M2.719 11.97V3.972A1.184 1.184 0 0 1 3.903 2.789H7.097A1.184 1.184 0 0 1 8.281 3.972V11.97L5.5 9.3Z" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".82"/>`}

var SavesCompact = Icon{width: "16", height: "25", viewBox: "0 0 16 25", paint: ``, inner: `<path d="M4 19.25V7.45A1.7 1.7 0 0 1 5.7 5.75H10.3A1.7 1.7 0 0 1 12 7.45V19.25L8 15.75Z" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"/>`}

var Account = Icon{width: "11", height: "16", viewBox: "0 0 11 16", paint: ``, inner: `<circle cx="5.5" cy="7.25" r="4.05" fill="none" stroke="currentColor" stroke-width=".79"/><circle cx="5.5" cy="7.25" r=".9" fill="none" stroke="currentColor" stroke-width=".65"/><path d="M6.7 7.25H9.2M4.9 8.289 3.65 10.454M4.9 6.211 3.65 4.046" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".69"/><path d="M9.9 7.25H10.32M9.635 8.755 10.029 8.899M8.871 10.078 9.192 10.348M7.7 11.061 7.91 11.424M6.264 11.583 6.337 11.997M4.736 11.583 4.663 11.997M3.3 11.061 3.09 11.424M2.129 10.078 1.808 10.348M1.365 8.755 .971 8.899M1.1 7.25H.68M1.365 5.745 .971 5.601M2.129 4.422 1.808 4.152M3.3 3.439 3.09 3.076M4.736 2.917 4.663 2.503M6.264 2.917 6.337 2.503M7.7 3.439 7.91 3.076M8.871 4.422 9.192 4.152M9.635 5.745 10.029 5.601" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".68"/>`}

var AccountCompact = Icon{width: "16", height: "25", viewBox: "0 0 16 25", paint: ``, inner: `<circle cx="8" cy="12.5" r="6.1" fill="none" stroke="currentColor" stroke-width="1.28"/><circle cx="8" cy="12.5" r="1.42" fill="none" stroke="currentColor" stroke-width="1.25"/><path d="M10 12.5H13.5M7 14.232 5.25 17.263M7 10.768 5.25 7.737" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.21"/><path d="M14.7 12.5H15.15M14.296 14.792 14.719 14.945M13.132 16.807 13.477 17.096M11.35 18.302 11.575 18.692M9.163 19.098 9.242 19.541M6.837 19.098 6.758 19.541M4.65 18.302 4.425 18.692M2.868 16.807 2.523 17.096M1.704 14.792 1.281 14.945M1.3 12.5H.85M1.704 10.208 1.281 10.055M2.868 8.193 2.523 7.904M4.65 6.698 4.425 6.308M6.837 5.902 6.758 5.459M9.163 5.902 9.242 5.459M11.35 6.698 11.575 6.308M13.132 8.193 13.477 7.904M14.296 10.208 14.719 10.055" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.05"/>`}

var SignIn = Icon{width: "11", height: "16", viewBox: "0 0 11 16", paint: ``, inner: `<circle cx="5.5" cy="7.633" r="4.287" fill="none" stroke="currentColor" stroke-width=".887"/><ellipse cx="5.5" cy="6.71" rx="1.605" ry="1.69" fill="currentColor"/><path d="M2.62 10.163A3.898 3.898 0 0 1 8.38 10.163 3.844 3.844 0 0 1 2.62 10.163Z" fill="currentColor"/>`}

var SignInCompact = Icon{width: "16", height: "25", viewBox: "0 0 16 25", paint: ``, inner: `<circle cx="7.99" cy="12.5" r="6.495" fill="none" stroke="currentColor" stroke-width="1.21"/><ellipse cx="7.99" cy="11" rx="2.475" ry="2.6" fill="currentColor"/><path d="M3.529 16.347A6.77 6.77 0 0 1 12.452 16.347 5.89 5.89 0 0 1 3.529 16.347Z" fill="currentColor"/>`}
