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
	// Required: set it on one radio of the group.
	Required bool
	Disabled bool
	// Caption says what the choice means, under its label.
	Caption string
	Class   string
	// Children is its label.
	Children g.Node
}

func (r ConsoleRadio) Render(w io.Writer) error {
	if r.Name == "" {
		panic("console.ConsoleRadio needs a Name")
	}
	return h.Label(class.Root("cui-console-radio", r.Class),
		h.Input(h.Class("cui-console-radio-input"), h.Type("radio"), h.Name(r.Name), h.Value(r.Value), g.If(r.Checked, h.Checked()), g.If(r.Required, h.Required()), g.If(r.Disabled, h.Disabled())),
		r.Children,
		g.If(r.Caption != "", h.Span(h.Class("cui-console-radio-caption"), g.Text(r.Caption))),
	).Render(w)
}
