package site

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Modal struct {
	ID string
	// Labelledby: the id of a headline in Children.
	Labelledby string
	Page       bool
	Dark       bool
	Class      string
	Children   g.Node
}

func (m Modal) Render(w io.Writer) error {
	if m.ID == "" {
		panic("site.Modal needs an ID")
	}
	if m.Labelledby == "" {
		panic("site.Modal needs the id of its headline (Labelledby)")
	}
	if m.Dark && !m.Page {
		panic("site.Modal is Dark only with Page")
	}
	closeIcon := icon.Close
	if m.Page {
		closeIcon = icon.CloseSmall
	}
	block := "cui-modal"
	if m.Dark {
		block = "cui-modal cui-dark"
	}
	return h.Dialog(class.Root(block, m.Class), g.If(m.Page, h.Data("variant", "page")), h.ID(m.ID), h.Aria("labelledby", m.Labelledby),
		h.Div(h.Class("cui-modal-panel"),
			h.Form(h.Method("dialog"), h.Button(h.Class("cui-modal-close"), h.Aria("label", "Close"), closeIcon)),
			m.Children,
		),
	).Render(w)
}

type ModalButton struct {
	Modal string
	// Class: the button's whole class (cui-btn).
	Class    string
	Children g.Node
}

func (b ModalButton) Render(w io.Writer) error {
	if b.Modal == "" {
		panic("site.ModalButton needs its Modal's id")
	}
	// type="button": inside a form a submit button ignores its command
	// (reference/dialog/)
	return h.Button(h.Class(b.Class), h.Type("button"), g.Attr("commandfor", b.Modal), g.Attr("command", "show-modal"), b.Children).Render(w)
}
