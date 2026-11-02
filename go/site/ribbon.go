package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Ribbon is cui-ribbon (components/site/ribbon): the message above a page's
// content, a band on the alternate ground, its text centered.
type Ribbon struct {
	// Class is the page's utilities on the ribbon; "" for none.
	Class string
	// Children is the message: a paragraph, a more link in it.
	Children g.Node
}

// Render writes the ribbon.
func (r Ribbon) Render(w io.Writer) error {
	return h.Div(class.Root("cui-ribbon", r.Class),
		h.Div(h.Class("cui-ribbon-content"), r.Children),
	).Render(w)
}
