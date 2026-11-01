package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleSelect is cui-console-select (components/console/console-field): App
// Store Connect's newer look's dropdown, the page's own
// <select class="cui-console-field"> (its Children) in its box, with its
// chevron. A text field or a multi-line one is the page's own <input> or
// <textarea> with the class.
type ConsoleSelect struct {
	// Class is the page's utilities on its box; "" for none.
	Class string
	// Children is its <select>.
	Children g.Node
}

// Render writes the select's box.
func (s ConsoleSelect) Render(w io.Writer) error {
	return h.Span(class.Root("cui-console-select", s.Class),
		s.Children,
		lucide.ChevronsUpDown.Class("cui-console-select-chevron"),
	).Render(w)
}

// ConsoleFormField is cui-console-form-field (components/console/console-field):
// one field's box, its label (cui-console-label, 5px over it) first, then
// the body: the control the page writes, then a [ConsoleHelp] or an [ConsoleError] where
// there is one.
//
// The label's for is ID. The control is the page's, so the page gives it
// id, aria-describedby naming the help's and the error's ids ([ConsoleHelpID],
// [ConsoleErrorID]) where they are written, and aria-invalid="true" while there is
// an error.
type ConsoleFormField struct {
	// ID is the control's id: the label's for. Required: an empty one
	// panics.
	ID string
	// Label is the label's content, text or markup.
	Label g.Node
	// Children is the control, then its help or its error.
	Children g.Node
}

// Render writes the field's box.
func (f ConsoleFormField) Render(w io.Writer) error {
	if f.ID == "" {
		panic("console.ConsoleFormField needs an ID")
	}
	return h.Div(h.Class("cui-console-form-field"),
		h.Label(h.Class("cui-console-label"), h.For(f.ID), f.Label),
		f.Children,
	).Render(w)
}

// ConsoleHelp is cui-console-help: the line under a control saying what it takes,
// its id [ConsoleHelpID] for the control's aria-describedby.
type ConsoleHelp struct {
	// ID is the control's id.
	ID string
	// Text is the help. Required: an empty one panics.
	Text string
}

// Render writes the help line.
func (p ConsoleHelp) Render(w io.Writer) error {
	if p.Text == "" {
		panic("console.ConsoleHelp needs a Text")
	}
	return h.P(h.Class("cui-console-help"), h.ID(ConsoleHelpID(p.ID)), g.Text(p.Text)).Render(w)
}

// ConsoleError is cui-console-error: the line under a control saying what is
// wrong, its id [ConsoleErrorID] for the control's aria-describedby, written while
// there is an error; or, with no Text, its slot, empty and hidden, which the
// page's script fills when a request fails (role="alert": read out once it
// is filled).
type ConsoleError struct {
	// ID is the control's id; "" for a slot that is no one control's (a
	// form's). An error with a Text needs it: it panics without.
	ID string
	// Text is the error; "" for the slot.
	Text string
	// Class is the page's utilities on it; "" for none.
	Class string
}

// Render writes the error line or its slot.
func (e ConsoleError) Render(w io.Writer) error {
	if e.Text == "" {
		return h.P(class.Root("cui-console-error", e.Class), g.If(e.ID != "", h.ID(ConsoleErrorID(e.ID))), h.Role("alert"), g.Attr("hidden")).Render(w)
	}
	if e.ID == "" {
		panic("console.ConsoleError with a Text needs its control's ID")
	}
	return h.P(class.Root("cui-console-error", e.Class), h.ID(ConsoleErrorID(e.ID)), g.Text(e.Text)).Render(w)
}

// ConsoleHelpID is the help line's id for the control id.
func ConsoleHelpID(id string) string {
	return id + "-help"
}

// ConsoleErrorID is the error line's id for the control id.
func ConsoleErrorID(id string) string {
	return id + "-error"
}
