package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleDialog is a modal the page's script opens with showModal().
type ConsoleDialog struct {
	ID    string
	Title g.Node
	// Open renders it open, for a page that shows it with no script.
	Open  bool
	Class string
	// Children: what it asks, then its buttons in a
	// <form method="dialog" class="cui-console-dialog-actions"> or the page's
	// own form.
	Children g.Node
}

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
