package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type ModalVariant struct{ page, dark bool }

var (
	// ModalVariantSheet: the store's, light on any page.
	ModalVariantSheet = ModalVariant{page: false, dark: false}
	// ModalVariantPage: the marketing pages'.
	ModalVariantPage = ModalVariant{page: true, dark: false}
	// ModalVariantPageDark: the marketing pages' on a dark product's page.
	ModalVariantPageDark = ModalVariant{page: true, dark: true}
)

type Modal struct {
	ID string
	// Labelledby: the id of a headline in Children.
	Labelledby string
	Variant    ModalVariant
	// CloseAriaLabel: the close button, in the page's language.
	CloseAriaLabel string
	Class          string
	Children       g.Node
}

func (m Modal) Render(w io.Writer) error {
	if m.ID == "" {
		panic("site.Modal needs an ID")
	}
	if m.Labelledby == "" {
		panic("site.Modal needs the id of its headline (Labelledby)")
	}
	if m.CloseAriaLabel == "" {
		panic("site.Modal needs a CloseAriaLabel")
	}
	block := "cui-modal"
	if m.Variant.dark {
		block = "cui-modal cui-dark"
	}
	return h.Dialog(class.Root(block, m.Class), g.If(m.Variant.page, h.Data("variant", "page")), h.ID(m.ID), h.Aria("labelledby", m.Labelledby),
		h.Div(h.Class("cui-modal-panel"),
			h.Form(h.Method("dialog"), h.Button(h.Class("cui-modal-close"), h.Aria("label", m.CloseAriaLabel), glyph.Close)),
			m.Children,
		),
		// after the panel: the dialog's opening focuses its first focusable
		// element, which must stay the close button
		h.Form(h.Method("dialog"), h.Button(h.Class("cui-modal-curtain"), h.TabIndex("-1"), h.Aria("hidden", "true"))),
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
