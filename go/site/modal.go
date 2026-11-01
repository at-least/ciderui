package site

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Modal is cui-modal (components/site/modal): apple.com's store overlay on a
// <dialog>, the page dimmed and held under it, its content in a white panel
// with a close button at its corner (a form's, method="dialog"), which and
// Escape close it with no script; or, Page, the marketing pages' modal,
// centered over the page dimmed and blurred, its close button the neutral
// fill; cui-dark in its Class, its dark one. A [ModalButton] opens it.
type Modal struct {
	// ID is the dialog's id, which its [ModalButton] names. Required: an
	// empty one panics.
	ID string
	// Labelledby is the id of its headline, which names it
	// (aria-labelledby). Required: an empty one panics.
	Labelledby string
	// Page is whether it is the marketing pages' modal
	// (data-variant="page"), not the store's.
	Page bool
	// Class is the page's utilities on the dialog (cui-dark with Page, the
	// dark one); "" for none.
	Class string
	// Children is the panel's content, after its close button.
	Children g.Node
}

// Render writes the dialog.
func (m Modal) Render(w io.Writer) error {
	if m.ID == "" {
		panic("site.Modal needs an ID")
	}
	if m.Labelledby == "" {
		panic("site.Modal needs the id of its headline (Labelledby)")
	}
	closeIcon := icon.Close
	if m.Page {
		closeIcon = icon.CloseSmall
	}
	return h.Dialog(class.Root("cui-modal", m.Class), g.If(m.Page, h.Data("variant", "page")), h.ID(m.ID), h.Aria("labelledby", m.Labelledby),
		h.Div(h.Class("cui-modal-panel"),
			h.Form(h.Method("dialog"), h.Button(h.Class("cui-modal-close"), h.Aria("label", "Close"), closeIcon)),
			m.Children,
		),
	).Render(w)
}

// ModalButton is the button that opens a [Modal] with no script: commandfor
// naming the dialog, command="show-modal", its class the page's (a cui-btn);
// its type button, so that in a form it opens the dialog, where a form's
// submit button does nothing with its command (reference/dialog/).
type ModalButton struct {
	// Modal is the id of the [Modal] it opens. Required: an empty one
	// panics.
	Modal string
	// Class is its class: the page's (cui-btn, a link's).
	Class string
	// Children is its text.
	Children g.Node
}

// Render writes the button.
func (b ModalButton) Render(w io.Writer) error {
	if b.Modal == "" {
		panic("site.ModalButton needs its Modal's id")
	}
	return h.Button(h.Class(b.Class), h.Type("button"), g.Attr("commandfor", b.Modal), g.Attr("command", "show-modal"), b.Children).Render(w)
}
