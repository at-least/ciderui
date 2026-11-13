package console

import (
	"io"

	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleSelect boxes the page's own <select>: a [ConsoleField]'s Control, or
// a <select class="cui-console-field-input"> of the page's outside one. A text
// field is the page's own <input> or <textarea> and needs no type here.
type ConsoleSelect struct {
	Children g.Node
}

func (s ConsoleSelect) Render(w io.Writer) error {
	return h.Span(h.Class("cui-console-select"),
		s.Children,
		lucide.ChevronsUpDown.Class("cui-console-select-chevron"),
	).Render(w)
}

// ConsoleField is a control under its label, and an error under it. Control
// writes attrs on the page's own <input>, <textarea> or <select>: the page
// cannot forget them or get an id wrong.
type ConsoleField struct {
	ID      string
	Label   string
	Error   string
	Control func(attrs g.Node) g.Node
}

func (f ConsoleField) Render(w io.Writer) error {
	need("console.ConsoleField", "an ID", f.ID)
	need("console.ConsoleField", "a Label", f.Label)
	if f.Control == nil {
		panic("console.ConsoleField needs a Control")
	}
	errorID := f.ID + "-error"
	return h.Div(h.Class("cui-console-field"),
		h.Label(h.Class("cui-console-field-label"), h.For(f.ID), g.Text(f.Label)),
		f.Control(g.Group{
			h.Class("cui-console-field-input"),
			h.ID(f.ID),
			g.If(f.Error != "", g.Group{h.Aria("invalid", "true"), h.Aria("describedby", errorID)}),
		}),
		g.If(f.Error != "", h.P(h.Class("cui-console-field-error"), h.ID(errorID), g.Text(f.Error))),
	).Render(w)
}
