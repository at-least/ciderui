package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Swatch is an <li>: put it in a <ul class="cui-swatches">.
type Swatch struct {
	ID    string
	Name  string
	Value string
	// AriaLabel: the color's name, read by screen readers and shown by the
	// pointer's rest as the input's title.
	AriaLabel string
	Checked   bool
	// Attrs: nodes appended to the <input> (disabled, required); nil for none.
	Attrs g.Node
	Class string
	// Children: cui-swatch-color on an <img>, or on an element its style
	// colors.
	Children g.Node
}

func (s Swatch) Render(w io.Writer) error {
	if s.ID == "" {
		panic("site.Swatch needs an ID")
	}
	if s.Name == "" {
		panic("site.Swatch needs a Name")
	}
	if s.AriaLabel == "" {
		panic("site.Swatch needs an AriaLabel")
	}
	return h.Li(class.Root("cui-swatch", s.Class),
		// the title on the input: it lies over the ring, so it is what the pointer rests on
		h.Input(h.Class("cui-swatch-input"), h.Type("radio"), h.ID(s.ID), h.Name(s.Name), h.Value(s.Value), h.Title(s.AriaLabel), g.If(s.Checked, h.Checked()), s.Attrs),
		h.Label(h.Class("cui-swatch-ring"), h.For(s.ID),
			s.Children,
			h.Span(h.Class("cui-swatch-name"), g.Text(s.AriaLabel)),
		),
	).Render(w)
}
