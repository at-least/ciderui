package console

import (
	"io"

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

// ConsoleFormField: the control is the page's, so the page writes on it
// id=ID, aria-describedby naming [ConsoleHelpID] and [ConsoleErrorID] where
// those lines are written, and aria-invalid="true" while there is an error.
type ConsoleFormField struct {
	ID    string
	Label g.Node
	// Children: the control, then a [ConsoleHelp] or a [ConsoleError].
	Children g.Node
}

func (f ConsoleFormField) Render(w io.Writer) error {
	if f.ID == "" {
		panic("console.ConsoleFormField needs an ID")
	}
	if f.Label == nil {
		panic("console.ConsoleFormField needs a Label")
	}
	return h.Div(h.Class("cui-console-form-field"),
		h.Label(h.Class("cui-console-label"), h.For(f.ID), f.Label),
		f.Children,
	).Render(w)
}

type ConsoleHelp struct {
	ControlID string
	Label     string
}

func (p ConsoleHelp) Render(w io.Writer) error {
	if p.Label == "" {
		panic("console.ConsoleHelp needs a Label")
	}
	return h.P(h.Class("cui-console-help"), h.ID(ConsoleHelpID(p.ControlID)), g.Text(p.Label)).Render(w)
}

// ConsoleError with no Label is a hidden slot the page's script fills and
// unhides when a request fails.
type ConsoleError struct {
	// ControlID: "" for a form's slot.
	ControlID string
	Label     string
	Class     string
}

func (e ConsoleError) Render(w io.Writer) error {
	if e.Label == "" {
		return h.P(class.Root("cui-console-error", e.Class), g.If(e.ControlID != "", h.ID(ConsoleErrorID(e.ControlID))), h.Role("alert"), g.Attr("hidden")).Render(w)
	}
	if e.ControlID == "" {
		panic("console.ConsoleError with a Label needs a ControlID")
	}
	return h.P(class.Root("cui-console-error", e.Class), h.ID(ConsoleErrorID(e.ControlID)), g.Text(e.Label)).Render(w)
}

func ConsoleHelpID(id string) string {
	return id + "-help"
}

func ConsoleErrorID(id string) string {
	return id + "-error"
}
