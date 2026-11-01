package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// OptionTile is cui-option (components/site/option): apple.com's buy page's
// choice, a tile on a radio, its title (a detail under it) at one side and a
// price at the other; checked, its edge the accent. Not Option: that is a
// [Select]'s choice.
type OptionTile struct {
	// ID is the radio's id, its label's for. Required: an empty one panics.
	ID string
	// Name is the radio's name, its group's.
	Name string
	// Value is the radio's value.
	Value string
	// Title is the title. Required: an empty one panics.
	Title string
	// Detail is a detail under the title; "" for none.
	Detail string
	// Price is the price at its end; "" for none.
	Price string
	// Checked is whether it is checked.
	Checked bool
	// Class is the page's utilities on the option; "" for none.
	Class string
}

// Render writes the option.
func (o OptionTile) Render(w io.Writer) error {
	if o.ID == "" {
		panic("site.OptionTile needs an ID")
	}
	if o.Title == "" {
		panic("site.OptionTile needs a Title")
	}
	return h.Div(class.Root("cui-option", o.Class),
		h.Input(h.Class("cui-option-input"), h.Type("radio"), h.ID(o.ID), h.Name(o.Name), h.Value(o.Value), g.If(o.Checked, h.Checked())),
		h.Label(h.Class("cui-option-label"), h.For(o.ID),
			h.Span(h.Class("cui-option-title"),
				g.Text(o.Title),
				g.If(o.Detail != "", h.Span(h.Class("cui-option-detail"), g.Text(o.Detail))),
			),
			g.If(o.Price != "", h.Span(h.Class("cui-option-price"), g.Text(o.Price))),
		),
	).Render(w)
}
