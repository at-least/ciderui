package site

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Checkbox is cui-checkbox (components/site/checkbox): apple.com's 16px box
// before its label, the input itself kept under the box (it takes the focus,
// the click and the space bar); checked, the accent with a white check.
type Checkbox struct {
	// ID is the input's id, its label's for. Required: an empty one panics.
	ID string
	// Name is the input's name.
	Name string
	// Label is the label. Required: an empty one panics.
	Label string
	// Checked is whether it is checked.
	Checked bool
	// Class is the page's utilities on the checkbox; "" for none.
	Class string
}

// Render writes the checkbox.
func (c Checkbox) Render(w io.Writer) error {
	if c.ID == "" {
		panic("site.Checkbox needs an ID")
	}
	if c.Label == "" {
		panic("site.Checkbox needs a Label")
	}
	return h.Div(class.Root("cui-checkbox", c.Class),
		h.Input(h.Class("cui-checkbox-input"), h.Type("checkbox"), h.ID(c.ID), h.Name(c.Name), g.If(c.Checked, h.Checked())),
		h.Label(h.Class("cui-checkbox-label"), h.For(c.ID), h.Span(h.Class("cui-checkbox-indicator"), h.Aria("hidden", "true"), icon.Check), g.Text(c.Label)),
	).Render(w)
}
