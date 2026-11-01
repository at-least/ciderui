package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Swatch is cui-swatch (components/site/swatch): apple.com's buy page's color
// choice, an <li> of a <ul class="cui-swatches">: a 28px circle of the color
// (its Children: the page's <img class="cui-swatch-color">, or an element its
// style colors) in a 36px ring, on a radio; checked, the ring the accent; the
// color's name hidden but read.
type Swatch struct {
	// ID is the radio's id, its ring's for. Required: an empty one panics.
	ID string
	// Name is the radio's name, its group's.
	Name string
	// Value is the radio's value.
	Value string
	// Label is the color's name, read but not shown. Required: an empty one
	// panics.
	Label string
	// Checked is whether it is checked.
	Checked bool
	// Class is the page's utilities on the item; "" for none.
	Class string
	// Children is the color: cui-swatch-color on an image, or on an element
	// its style colors.
	Children g.Node
}

// Render writes the swatch.
func (s Swatch) Render(w io.Writer) error {
	if s.ID == "" {
		panic("site.Swatch needs an ID")
	}
	if s.Label == "" {
		panic("site.Swatch needs a Label")
	}
	return h.Li(class.Root("cui-swatch", s.Class),
		h.Input(h.Class("cui-swatch-input"), h.Type("radio"), h.ID(s.ID), h.Name(s.Name), h.Value(s.Value), g.If(s.Checked, h.Checked())),
		h.Label(h.Class("cui-swatch-ring"), h.For(s.ID),
			s.Children,
			h.Span(h.Class("cui-swatch-name"), g.Text(s.Label)),
		),
	).Render(w)
}
