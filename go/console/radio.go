package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleRadio is one radio; its group's <fieldset> and legend are the
// page's own.
type ConsoleRadio struct {
	Name    string
	Value   string
	Checked bool
	// Caption says what the choice means, under its label.
	Caption string
	// Attrs: the input's other attributes (required on one radio of the
	// group, disabled); nil for none.
	Attrs g.Node
	Class string
	// Children is its label.
	Children g.Node
}

func (r ConsoleRadio) Render(w io.Writer) error {
	if r.Name == "" {
		panic("console.ConsoleRadio needs a Name")
	}
	if r.Children == nil {
		panic("console.ConsoleRadio needs its label as Children")
	}
	return h.Label(class.Root("cui-console-radio", r.Class),
		h.Input(h.Class("cui-console-radio-input"), h.Type("radio"), h.Name(r.Name), h.Value(r.Value), g.If(r.Checked, h.Checked()), r.Attrs),
		r.Children,
		g.If(r.Caption != "", h.Span(h.Class("cui-console-radio-caption"), g.Text(r.Caption))),
	).Render(w)
}
