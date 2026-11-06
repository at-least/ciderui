package console

import (
	"io"
	"strings"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleSelect boxes the page's own <select>. A text field is the page's own
// <input> or <textarea> with cui-console-field and needs no type here.
type ConsoleSelect struct {
	Class string
	// Children: <select class="cui-console-field">.
	Children g.Node
}

func (s ConsoleSelect) Render(w io.Writer) error {
	return h.Span(class.Root("cui-console-select", s.Class),
		s.Children,
		lucide.ChevronsUpDown.Class("cui-console-select-chevron"),
	).Render(w)
}

// ConsoleFormField is a control under its label, a help line and an error under
// it. Control gets what it must carry (its id, what describes it, whether it is
// invalid) to write on the page's own <input>, <textarea> or <select>: the page
// cannot forget them or get an id wrong.
type ConsoleFormField struct {
	ID    string
	Label string
	// Help: a line that says what to write; "" for none.
	Help string
	// Error: "" while there is none.
	Error   string
	Control func(attrs g.Node) g.Node
}

func (f ConsoleFormField) Render(w io.Writer) error {
	if f.ID == "" {
		panic("console.ConsoleFormField needs an ID")
	}
	if f.Label == "" {
		panic("console.ConsoleFormField needs a Label")
	}
	if f.Control == nil {
		panic("console.ConsoleFormField needs a Control")
	}
	helpID, errorID := f.ID+"-help", f.ID+"-error"
	var described []string
	if f.Help != "" {
		described = append(described, helpID)
	}
	if f.Error != "" {
		described = append(described, errorID)
	}
	return h.Div(h.Class("cui-console-form-field"),
		h.Label(h.Class("cui-console-label"), h.For(f.ID), g.Text(f.Label)),
		f.Control(g.Group{
			h.ID(f.ID),
			g.If(f.Error != "", h.Aria("invalid", "true")),
			g.If(len(described) > 0, h.Aria("describedby", strings.Join(described, " "))),
		}),
		g.If(f.Help != "", h.P(h.Class("cui-console-help"), h.ID(helpID), g.Text(f.Help))),
		g.If(f.Error != "", h.P(h.Class("cui-console-error"), h.ID(errorID), g.Text(f.Error))),
	).Render(w)
}
