package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleDialog is cui-console-dialog (components/console/console-dialog): App
// Store Connect's newer look's modal, a <dialog> the page opens with
// showModal(): its title (the dialog's name, aria-labelledby), and what the
// page gives it — what it asks, its buttons in a
// <form method="dialog" class="cui-console-dialog-actions"> or the page's
// own form.
type ConsoleDialog struct {
	// ID is its id, which the page's script opens; its title's is the id
	// with -title. Required: an empty one panics.
	ID string
	// Title is its title, text or markup.
	Title g.Node
	// Open is whether it is rendered open (open): a page shown with it open,
	// no script.
	Open bool
	// Class is the page's utilities on it; "" for none.
	Class string
	// Children is what it asks and its buttons.
	Children g.Node
}

// Render writes the dialog.
func (d ConsoleDialog) Render(w io.Writer) error {
	if d.ID == "" {
		panic("console.ConsoleDialog needs an ID")
	}
	titleID := d.ID + "-title"
	return h.Dialog(class.Root("cui-console-dialog", d.Class), h.ID(d.ID), h.Aria("labelledby", titleID), g.If(d.Open, h.Open()),
		h.H2(h.Class("cui-console-dialog-title"), h.ID(titleID), d.Title),
		d.Children,
	).Render(w)
}
