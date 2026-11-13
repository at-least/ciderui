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

type Icon struct {
	width   string
	height  string
	viewBox string
	paint   string
	inner   string
}

// Lucide is a Lucide icon; paint is the attributes on Lucide's own <svg>.
func Lucide(viewBox, paint, inner string) Icon {
	// @trap: the paint on the <svg>, not on each element: a class on the <svg>
	// (stroke-[1.5]) wins over a presentation attribute there, and the
	// elements inherit what wins
	return Icon{width: "1em", height: "1em", viewBox: viewBox, paint: paint, inner: inner}
}

func (i Icon) Render(w io.Writer) error {
	return i.write(w, "")
}

// Class is the icon with a class on its <svg>: a component's part
// (`cui-link-icon`) or a utility.
func (i Icon) Class(class string) g.Node {
	return g.NodeFunc(func(w io.Writer) error { return i.write(w, ` class="`+html.EscapeString(class)+`"`) })
}

func (i Icon) write(w io.Writer, classAttribute string) error {
	svg := "<svg" + classAttribute + ` aria-hidden="true" width="` + i.width + `" height="` + i.height + `" viewBox="` + i.viewBox + `"`
	if i.paint != "" {
		svg += " " + i.paint
	}
	_, err := io.WriteString(w, svg+">"+i.inner+"</svg>")
	return err
}

var Menu = Icon{width: "18", height: "18", viewBox: "0 0 18 18", paint: ``, inner: `<path d="M2 5H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/><path d="M2 12H16" fill="none" stroke="currentColor" stroke-width="1.2" stroke-linecap="round"/>`}

var Disclosure = Icon{width: "17", height: "17", viewBox: "0 0 17 8.85", paint: ``, inner: `<path d="M15 1.13 8.5 7.72 2 1.13" fill="none" stroke="currentColor" stroke-width="2.25" stroke-linecap="round" stroke-linejoin="round"/>`}

var Alert = Icon{width: "1.171em", height: "0.995em", viewBox: "0 0 1171 995", paint: ``, inner: `<circle cx="585" cy="497.5" r="456.25" fill="none" stroke="currentColor" stroke-width="82.5"/><path d="M584 280V550" fill="none" stroke="currentColor" stroke-width="70" stroke-linecap="round"/><circle cx="585" cy="692.5" r="50" fill="currentColor"/>`}
