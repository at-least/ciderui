package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleRadio is cui-console-radio (components/console/console-radio): App Store
// Connect's newer look's radio, the input itself drawn, in its <label>
// before its label (Children), and the line under it that says what the
// choice means. Its group's <fieldset> and legend are the page's own.
type ConsoleRadio struct {
	// Name is its group's name. Required: an empty one panics.
	Name string
	// Value is its value.
	Value string
	// Checked is whether it is the chosen one.
	Checked bool
	// Required is whether its group needs a choice (required, on one radio
	// of it).
	Required bool
	// Disabled is whether it cannot be chosen.
	Disabled bool
	// Caption is the line under its label; "" for none.
	Caption string
	// Class is the page's utilities on its <label>; "" for none.
	Class string
	// Children is its label.
	Children g.Node
}

// Render writes the radio.
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
