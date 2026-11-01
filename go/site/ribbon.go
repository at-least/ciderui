package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Ribbon is cui-ribbon (components/site/ribbon): the message above a page's
// content, a band on the alternate ground, its text centered; with Announce,
// it drops in on the accent as the page shows it and turns to its ground.
type Ribbon struct {
	// Announce is whether it drops in on the accent and turns
	// (data-variant="announce").
	Announce bool
	// Class is the page's utilities on the ribbon; "" for none.
	Class string
	// Children is the message: a paragraph, a more link in it.
	Children g.Node
}

// Render writes the ribbon.
func (r Ribbon) Render(w io.Writer) error {
	return h.Div(class.Root("cui-ribbon", r.Class), g.If(r.Announce, h.Data("variant", "announce")),
		h.Div(h.Class("cui-ribbon-content"), r.Children),
	).Render(w)
}
