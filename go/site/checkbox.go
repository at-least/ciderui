package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Checkbox struct {
	ID      string
	Name    string
	Label   string
	Checked bool
	// Attrs: nodes appended to the <input> (disabled, required); nil for none.
	Attrs g.Node
	Class string
}

func (c Checkbox) Render(w io.Writer) error {
	if c.ID == "" {
		panic("site.Checkbox needs an ID")
	}
	if c.Name == "" {
		panic("site.Checkbox needs a Name")
	}
	if c.Label == "" {
		panic("site.Checkbox needs a Label")
	}
	return h.Div(class.Root("cui-checkbox", c.Class),
		h.Input(h.Class("cui-checkbox-input"), h.Type("checkbox"), h.ID(c.ID), h.Name(c.Name), g.If(c.Checked, h.Checked()), c.Attrs),
		h.Label(h.Class("cui-checkbox-label"), h.For(c.ID), h.Span(h.Class("cui-checkbox-indicator"), h.Aria("hidden", "true"), glyph.Check), g.Text(c.Label)),
	).Render(w)
}
