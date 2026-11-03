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
	// Label: the color's name, read by screen readers but not shown.
	Label   string
	Checked bool
	Class   string
	// Children: cui-swatch-color on an <img>, or on an element its style
	// colors.
	Children g.Node
}

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
